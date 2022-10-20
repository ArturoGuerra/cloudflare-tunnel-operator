/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cloudflared

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	appsv1alpha1 "github.com/arturoguerra/cloudflare-tunnel-operator/apis/apps/v1alpha1"
	"github.com/arturoguerra/cloudflare-tunnel-operator/apis/apps/v1alpha1/cloudflared/mutate"
)

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceCloudflared creates the Deployment resource with name cloudflared.
func CreateDeploymentNamespaceCloudflared(
	parent *appsv1alpha1.CloudflareTunnel,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app": "cloudflared",
				},
				"name":      "cloudflared",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"replicas": parent.Spec.Replicas, //  controlled by field: replicas
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "cloudflared",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app": "cloudflared",
						},
					},
					"spec": map[string]interface{}{
						"containers": []interface{}{
							map[string]interface{}{
								"args": []interface{}{
									"tunnel",
									"--no-autoupdate",
									"run",
								},
								"image":           "cloudflare/cloudflared:latest",
								"imagePullPolicy": "Always",
								"name":            "cloudflared-tunnel",
								"resources": map[string]interface{}{
									"limits": map[string]interface{}{
										"cpu":    "100m",
										"memory": "128Mi",
									},
									"requests": map[string]interface{}{
										"cpu":    "100m",
										"memory": "128Mi",
									},
								},
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"mountPath": "/etc/cloudflared/tunnel.json",
										"name":      "cloudflared-auth",
										"readOnly":  true,
										"subPath":   "tunnel.json",
									},
									map[string]interface{}{
										"mountPath": "/etc/cloudflared/cert.pem",
										"name":      "cloudflared-auth",
										"readOnly":  true,
										"subPath":   "cert.pem",
									},
									map[string]interface{}{
										"mountPath": "/etc/cloudflared/config.yaml",
										"name":      "cloudflared-config",
										"readOnly":  true,
										"subPath":   "config.yaml",
									},
								},
							},
						},
						"volumes": []interface{}{
							map[string]interface{}{
								"name": "cloudflared-auth",
								"secret": map[string]interface{}{
									"secretName": "cloudflared-auth",
								},
							},
							map[string]interface{}{
								"name": "cloudflared-config",
								"configMap": map[string]interface{}{
									"items": []interface{}{
										map[string]interface{}{
											"key":  "config.yaml",
											"path": "config.yaml",
										},
									},
									"name": "cloudflared-config",
								},
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateDeploymentNamespaceCloudflared(resourceObj, parent, reconciler, req)
}
