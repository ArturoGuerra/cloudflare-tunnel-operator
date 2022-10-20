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

// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;patch;delete

// CreateSecretNamespaceCloudflaredAuth creates the Secret resource with name cloudflared-auth.
func CreateSecretNamespaceCloudflaredAuth(
	parent *appsv1alpha1.CloudflareTunnel,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Secret",
			"metadata": map[string]interface{}{
				"name":      "cloudflared-auth",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"type": "Opaque",
			"data": map[string]interface{}{
				// controlled by field: tunnelCredentials
				"tunnel.json": "" + parent.Spec.TunnelCredentials + "",
			},
		},
	}

	return mutate.MutateSecretNamespaceCloudflaredAuth(resourceObj, parent, reconciler, req)
}
