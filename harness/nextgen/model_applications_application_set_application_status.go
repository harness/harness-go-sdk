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

type ApplicationsApplicationSetApplicationStatus struct {
	Application        string  `json:"application,omitempty"`
	LastTransitionTime *V1Time `json:"lastTransitionTime,omitempty"`
	Message            string  `json:"message,omitempty"`
	Status             string  `json:"status,omitempty"`
	Step               string  `json:"step,omitempty"`
	// TargetRevision tracks the desired revisions the Application should be synced to.
	Targetrevisions []string `json:"targetrevisions,omitempty"`
}
