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
	"strconv"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	appsv1alpha1 "github.com/arturoguerra/cloudflare-tunnel-operator/apis/apps/v1alpha1"
	"github.com/arturoguerra/cloudflare-tunnel-operator/apis/apps/v1alpha1/cloudflared/mutate"
)

// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete

// CreateConfigMapNamespaceCloudflaredConfig creates the ConfigMap resource with name cloudflared-config.
func CreateConfigMapNamespaceCloudflaredConfig(
	parent *appsv1alpha1.CloudflareTunnel,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name":      "cloudflared-config",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"data": map[string]interface{}{
				// controlled by field: tunnelID
				// controlled by field: service.protocol
				// controlled by field: service.name
				// controlled by field: service.port
				"config.yaml": `tunnel: ` + parent.Spec.TunnelID + `
credentials-file: /etc/cloudflared/tunnel.json
ingress:
- service: ` + parent.Spec.Service.Protocol + `://` + parent.Spec.Service.Name + `:` + strconv.Itoa(parent.Spec.Service.Port) + `
  originRequest:
    noTLSVerify: true`,
			},
		},
	}

	return mutate.MutateConfigMapNamespaceCloudflaredConfig(resourceObj, parent, reconciler, req)
}
