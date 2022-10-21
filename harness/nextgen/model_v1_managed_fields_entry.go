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

// ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.
type V1ManagedFieldsEntry struct {
	// Manager is an identifier of the workflow managing these fields.
	Manager string `json:"manager,omitempty"`
	// Operation is the type of operation which lead to this ManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'.
	Operation string `json:"operation,omitempty"`
	// APIVersion defines the version of this resource that this field set applies to. The format is \"group/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted.
	ApiVersion string      `json:"apiVersion,omitempty"`
	Time       *V1Time     `json:"time,omitempty"`
	FieldsType string      `json:"fieldsType,omitempty"`
	FieldsV1   *V1FieldsV1 `json:"fieldsV1,omitempty"`
	// Subresource is the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource.
	Subresource string `json:"subresource,omitempty"`
}