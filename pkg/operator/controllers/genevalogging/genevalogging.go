package genevalogging

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"strings"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	projectv1 "github.com/openshift/api/project/v1"
	securityv1 "github.com/openshift/api/security/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	arov1alpha1 "github.com/Azure/ARO-RP/pkg/operator/apis/aro.openshift.io/v1alpha1"
	"github.com/Azure/ARO-RP/pkg/util/version"
)

const (
	GenevaCertName = "gcscert.pem"
	GenevaKeyName  = "gcskey.pem"
)

func (g *GenevaloggingReconciler) securityContextConstraints(ctx context.Context, name, serviceAccountName string) (*securityv1.SecurityContextConstraints, error) {
	scc, err := g.securitycli.SecurityV1().SecurityContextConstraints().Get(ctx, "privileged", metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	scc.ObjectMeta = metav1.ObjectMeta{
		Name: name,
	}
	scc.Groups = []string{}
	scc.Users = []string{serviceAccountName}
	return scc, nil
}

func (g *GenevaloggingReconciler) daemonset(cluster *arov1alpha1.Cluster) (*appsv1.DaemonSet, error) {
	r, err := azure.ParseResourceID(cluster.Spec.ResourceID)
	if err != nil {
		return nil, err
	}

	return &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "mdsd",
			Namespace: kubeNamespace,
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": "mdsd"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels:      map[string]string{"app": "mdsd"},
					Annotations: map[string]string{"scheduler.alpha.kubernetes.io/critical-pod": ""},
				},
				Spec: corev1.PodSpec{
					Volumes: []corev1.Volume{
						{
							Name: "log",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: "/var/log",
								},
							},
						},
						{
							Name: "fluent",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: "/var/lib/fluent",
								},
							},
						},
						{
							Name: "fluent-config",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "fluent-config",
									},
								},
							},
						},
						{
							Name: "machine-id",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: "/etc/machine-id",
								},
							},
						},
						{
							Name: "certificates",
							VolumeSource: corev1.VolumeSource{
								Secret: &corev1.SecretVolumeSource{
									SecretName: certificatesSecretName,
								},
							},
						},
					},
					ServiceAccountName:       "geneva",
					DeprecatedServiceAccount: "geneva",
					Tolerations: []corev1.Toleration{
						{
							Effect:   corev1.TaintEffectNoExecute,
							Operator: corev1.TolerationOpExists,
						},
						{
							Effect:   corev1.TaintEffectNoSchedule,
							Operator: corev1.TolerationOpExists,
						},
					},
					Containers: []corev1.Container{
						{
							Name:  "fluentbit",
							Image: version.FluentbitImage(cluster.Spec.ACRDomain),
							Command: []string{
								"/opt/td-agent-bit/bin/td-agent-bit",
							},
							Args: []string{
								"-c",
								"/etc/td-agent-bit/fluent.conf",
							},
							// TODO: specify requests/limits
							SecurityContext: &corev1.SecurityContext{
								Privileged: to.BoolPtr(true),
								RunAsUser:  to.Int64Ptr(0),
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "fluent-config",
									ReadOnly:  true,
									MountPath: "/etc/td-agent-bit",
								},
								{
									Name:      "machine-id",
									ReadOnly:  true,
									MountPath: "/etc/machine-id",
								},
								{
									Name:      "log",
									ReadOnly:  true,
									MountPath: "/var/log",
								},
								{
									Name:      "fluent",
									MountPath: "/var/lib/fluent",
								},
							},
						},
						{
							Name:  "mdsd",
							Image: version.MdsdImage(cluster.Spec.ACRDomain),
							Command: []string{
								"/usr/sbin/mdsd",
							},
							Args: []string{
								"-A",
								"-D",
								"-f",
								"24224",
								"-r",
								"/var/run/mdsd/default",
							},
							Env: []corev1.EnvVar{
								{
									Name:  "MONITORING_GCS_ENVIRONMENT",
									Value: cluster.Spec.GenevaLogging.MonitoringGCSEnvironment,
								},
								{
									Name:  "MONITORING_GCS_ACCOUNT",
									Value: "AROClusterLogs",
								},
								{
									Name:  "MONITORING_GCS_REGION",
									Value: cluster.Spec.Location,
								},
								{
									Name:  "MONITORING_GCS_CERT_CERTFILE",
									Value: "/etc/mdsd.d/secret/" + GenevaCertName,
								},
								{
									Name:  "MONITORING_GCS_CERT_KEYFILE",
									Value: "/etc/mdsd.d/secret/" + GenevaKeyName,
								},
								{
									Name:  "MONITORING_GCS_NAMESPACE",
									Value: ClusterLogsNamespace,
								},
								{
									Name:  "MONITORING_CONFIG_VERSION",
									Value: cluster.Spec.GenevaLogging.ConfigVersion,
								},
								{
									Name:  "MONITORING_USE_GENEVA_CONFIG_SERVICE",
									Value: "true",
								},
								{
									Name:  "MONITORING_TENANT",
									Value: cluster.Spec.Location,
								},
								{
									Name:  "MONITORING_ROLE",
									Value: "cluster",
								},
								{
									Name: "MONITORING_ROLE_INSTANCE",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											APIVersion: "v1",
											FieldPath:  "spec.nodeName",
										},
									},
								},
								{
									Name:  "RESOURCE_ID",
									Value: strings.ToLower(cluster.Spec.ResourceID),
								},
								{
									Name:  "SUBSCRIPTION_ID",
									Value: strings.ToLower(r.SubscriptionID),
								},
								{
									Name:  "RESOURCE_GROUP",
									Value: strings.ToLower(r.ResourceGroup),
								},
								{
									Name:  "RESOURCE_NAME",
									Value: strings.ToLower(r.ResourceName),
								},
							},
							Resources: corev1.ResourceRequirements{
								Limits: corev1.ResourceList{
									corev1.ResourceCPU:    resource.MustParse("200m"),
									corev1.ResourceMemory: resource.MustParse("1000Mi"),
								},
								Requests: corev1.ResourceList{
									corev1.ResourceCPU:    resource.MustParse("10m"),
									corev1.ResourceMemory: resource.MustParse("100Mi"),
								},
							},
							SecurityContext: &corev1.SecurityContext{
								Privileged: to.BoolPtr(true),
								RunAsUser:  to.Int64Ptr(0),
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "certificates",
									MountPath: "/etc/mdsd.d/secret",
								},
							},
						},
					},
				},
			},
		},
	}, nil
}

func (g *GenevaloggingReconciler) resources(ctx context.Context, cluster *arov1alpha1.Cluster, gcscert, gcskey []byte) ([]runtime.Object, error) {
	scc, err := g.securityContextConstraints(ctx, "privileged-genevalogging", kubeServiceAccount)
	if err != nil {
		return nil, err
	}

	daemonset, err := g.daemonset(cluster)
	if err != nil {
		return nil, err
	}

	return []runtime.Object{
		&corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name:        kubeNamespace,
				Annotations: map[string]string{projectv1.ProjectNodeSelector: ""},
			},
		},
		&corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      certificatesSecretName,
				Namespace: kubeNamespace,
			},
			Data: map[string][]byte{
				GenevaCertName: gcscert,
				GenevaKeyName:  gcskey,
			},
		},
		&corev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "fluent-config",
				Namespace: kubeNamespace,
			},
			Data: map[string]string{
				"fluent.conf":  fluentConf,
				"parsers.conf": parsersConf,
			},
		},
		&corev1.ServiceAccount{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "geneva",
				Namespace: kubeNamespace,
			},
		},
		scc,
		daemonset,
	}, nil
}
