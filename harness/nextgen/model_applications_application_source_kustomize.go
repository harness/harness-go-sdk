/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type ApplicationsApplicationSourceKustomize struct {
	NamePrefix             string            `json:"namePrefix,omitempty"`
	NameSuffix             string            `json:"nameSuffix,omitempty"`
	Images                 []string          `json:"images,omitempty"`
	CommonLabels           map[string]string `json:"commonLabels,omitempty"`
	Version                string            `json:"version,omitempty"`
	CommonAnnotations      map[string]string `json:"commonAnnotations,omitempty"`
	ForceCommonLabels      bool              `json:"forceCommonLabels,omitempty"`
	ForceCommonAnnotations bool              `json:"forceCommonAnnotations,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Replicas []ApplicationsKustomizeReplicas `json:"replicas,omitempty"`
}
