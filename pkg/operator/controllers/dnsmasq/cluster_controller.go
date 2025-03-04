package dnsmasq

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"

	"github.com/openshift/installer/pkg/aro/dnsmasq"
	mcoclient "github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	arov1alpha1 "github.com/Azure/ARO-RP/pkg/operator/apis/aro.openshift.io/v1alpha1"
	aroclient "github.com/Azure/ARO-RP/pkg/operator/clientset/versioned"
	"github.com/Azure/ARO-RP/pkg/operator/controllers"
	"github.com/Azure/ARO-RP/pkg/util/dynamichelper"
)

type ClusterReconciler struct {
	log *logrus.Entry

	arocli aroclient.Interface
	mcocli mcoclient.Interface
	dh     dynamichelper.Interface
}

func NewClusterReconciler(log *logrus.Entry, arocli aroclient.Interface, mcocli mcoclient.Interface, dh dynamichelper.Interface) *ClusterReconciler {
	return &ClusterReconciler{
		log:    log,
		arocli: arocli,
		mcocli: mcocli,
		dh:     dh,
	}
}

// Reconcile watches the ARO object, and if it changes, reconciles all the
// 99-%s-aro-dns machineconfigs
func (r *ClusterReconciler) Reconcile(request ctrl.Request) (ctrl.Result, error) {
	// TODO(mj): controller-runtime master fixes the need for this (https://github.com/kubernetes-sigs/controller-runtime/blob/master/pkg/reconcile/reconcile.go#L93) but it's not yet released.
	ctx := context.Background()
	if request.Name != arov1alpha1.SingletonClusterName {
		return reconcile.Result{}, nil
	}

	mcps, err := r.mcocli.MachineconfigurationV1().MachineConfigPools().List(ctx, metav1.ListOptions{})
	if err != nil {
		r.log.Error(err)
		return reconcile.Result{}, err
	}

	roles := make([]string, 0, len(mcps.Items))
	for _, mcp := range mcps.Items {
		roles = append(roles, mcp.Name)
	}

	err = reconcileMachineConfigs(ctx, r.arocli, r.dh, roles...)
	if err != nil {
		r.log.Error(err)
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}

// SetupWithManager setup our mananger
func (r *ClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&arov1alpha1.Cluster{}).
		Named(controllers.DnsmasqClusterControllerName).
		Complete(r)
}

func reconcileMachineConfigs(ctx context.Context, arocli aroclient.Interface, dh dynamichelper.Interface, roles ...string) error {
	instance, err := arocli.AroV1alpha1().Clusters().Get(ctx, arov1alpha1.SingletonClusterName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	var resources []runtime.Object
	for _, role := range roles {
		resource, err := dnsmasq.MachineConfig(instance.Spec.Domain, instance.Spec.APIIntIP, instance.Spec.IngressIP, role)
		if err != nil {
			return err
		}

		resources = append(resources, resource)
	}

	err = dynamichelper.SetControllerReferences(resources, instance)
	if err != nil {
		return err
	}

	err = dynamichelper.Prepare(resources)
	if err != nil {
		return err
	}

	return dh.Ensure(ctx, resources...)
}
