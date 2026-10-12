/*
 * Chaos Manager
 *
 * ChaosManager serves as the primary API server for the chaos module, responsible for transferring and handling requests to their respective operations
 *
 * API version: 1.67.16
 */
package chaos

// Object represents a generic JSON object
// swagger:model
// +k8s:openapi-gen=true
type Object map[string]interface{}

// IntOrString is a type that can hold an int32 or a string.
// +protobuf=true
// +k8s:openapi-gen=true
// +kubebuilder:validation:Type=string
// +kubebuilder:validation:Format=int-or-string
type IntOrString struct {
	Type   IntstrType `json:"type"`
	IntVal int32      `json:"intVal,omitempty" protobuf:"varint,1,opt,name=intVal"`
	StrVal string     `json:"strVal,omitempty" protobuf:"bytes,2,opt,name=strVal"`
}

// IntstrType represents the stored type of IntOrString.
type IntstrType string

const (
	IntstrInt    IntstrType = "int"
	IntstrString IntstrType = "string"
)
