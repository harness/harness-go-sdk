/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type V1LabelSelectorRequirement struct {
	// key is the label key that the selector applies to. +patchMergeKey=key +patchStrategy=merge
	Key string `json:"key,omitempty"`
	// operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
	Operator *AllOfv1LabelSelectorRequirementOperator `json:"operator,omitempty"`
	// values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch. +optional
	Values []string `json:"values,omitempty"`
}