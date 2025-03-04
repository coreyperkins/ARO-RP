package adminactions

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"encoding/json"
	"fmt"

	mgmtcompute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-06-01/compute"
	"github.com/Azure/go-autorest/autorest/azure"

	"github.com/Azure/ARO-RP/pkg/util/arm"
	"github.com/Azure/ARO-RP/pkg/util/azureclient"
	"github.com/Azure/ARO-RP/pkg/util/stringutils"
	"github.com/Azure/ARO-RP/pkg/util/subnet"
)

func (a *azureActions) ResourcesList(ctx context.Context) ([]byte, error) {
	clusterRGName := stringutils.LastTokenByte(a.oc.Properties.ClusterProfile.ResourceGroupID, '/')

	resources, err := a.resources.ListByResourceGroup(ctx, clusterRGName, "", "", nil)
	if err != nil {
		return nil, err
	}

	armResources := make([]arm.Resource, 0, len(resources)+3)
	armResources, err = a.appendAzureNetworkResources(ctx, armResources)
	if err != nil {
		a.log.Warnf("error when getting network resources: %s", err)
	}

	for _, res := range resources {
		apiVersion := azureclient.APIVersion(*res.Type)
		if apiVersion == "" {
			return nil, fmt.Errorf("API version not found for type %s", *res.Type)
		}
		switch *res.Type {
		case "Microsoft.Compute/virtualMachines":
			vm, err := a.virtualMachines.Get(ctx, clusterRGName, *res.Name, mgmtcompute.InstanceView)
			if err != nil {
				return nil, err
			}
			armResources = append(armResources, arm.Resource{
				Resource: vm,
			})
		default:
			gr, err := a.resources.GetByID(ctx, *res.ID, apiVersion)
			if err != nil {
				return nil, err
			}
			armResources = append(armResources, arm.Resource{
				Resource: gr,
			})
		}
	}

	return json.Marshal(armResources)
}

func (a *azureActions) appendAzureNetworkResources(ctx context.Context, armResources []arm.Resource) ([]arm.Resource, error) {
	vNetID, _, err := subnet.Split(a.oc.Properties.MasterProfile.SubnetID)
	if err != nil {
		return nil, err
	}

	r, err := azure.ParseResourceID(vNetID)
	if err != nil {
		return armResources, err
	}

	vnet, err := a.virtualNetworks.Get(ctx, r.ResourceGroup, r.ResourceName, "")
	if err != nil {
		return armResources, err
	}
	armResources = append(armResources, arm.Resource{
		Resource: vnet,
	})
	if vnet.Subnets != nil {
		for _, snet := range *vnet.Subnets {
			//we already have the VNet resource, filtering subnets instead of fetching them individually with a SubnetClient
			interestingSubnet := (*snet.ID == a.oc.Properties.MasterProfile.SubnetID)
			for _, wProfile := range a.oc.Properties.WorkerProfiles {
				interestingSubnet = interestingSubnet || (*snet.ID == wProfile.SubnetID)
			}
			if !interestingSubnet {
				continue
			}
			//by this time the snet subnet is used in a Master or Worker profile
			if snet.RouteTable != nil {
				r, err := azure.ParseResourceID(*snet.RouteTable.ID)
				if err != nil {
					a.log.Warnf("skipping route table '%s' due to ID parse error: %s", *snet.RouteTable.ID, err)
					continue
				}
				rt, err := a.routeTables.Get(ctx, r.ResourceGroup, r.ResourceName, "")
				if err != nil {
					a.log.Warnf("skipping route table '%s' due to Get error: %s", *snet.RouteTable.ID, err)
					continue
				}
				armResources = append(armResources, arm.Resource{
					Resource: rt,
				})
			}
		}
	}

	return armResources, nil
}
