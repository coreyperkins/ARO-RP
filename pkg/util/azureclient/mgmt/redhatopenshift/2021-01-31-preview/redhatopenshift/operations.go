package redhatopenshift

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"crypto/tls"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"

	mgmtredhatopenshift20210131preview "github.com/Azure/ARO-RP/pkg/client/services/redhatopenshift/mgmt/2021-01-31-preview/redhatopenshift"
	"github.com/Azure/ARO-RP/pkg/util/deployment"
)

// OperationsClient is a minimal interface for azure OperationsClient
type OperationsClient interface {
	OperationsClientAddons
}

type operationsClient struct {
	mgmtredhatopenshift20210131preview.OperationsClient
}

var _ OperationsClient = &operationsClient{}

// NewOperationsClient creates a new OperationsClient
func NewOperationsClient(environment *azure.Environment, subscriptionID string, authorizer autorest.Authorizer) OperationsClient {
	var client mgmtredhatopenshift20210131preview.OperationsClient
	if deployment.NewMode() == deployment.Development {
		client = mgmtredhatopenshift20210131preview.NewOperationsClientWithBaseURI("https://localhost:8443", subscriptionID)
		client.Sender = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
	} else {
		client = mgmtredhatopenshift20210131preview.NewOperationsClientWithBaseURI(environment.ResourceManagerEndpoint, subscriptionID)
		client.Authorizer = authorizer
	}

	return &operationsClient{
		OperationsClient: client,
	}
}
