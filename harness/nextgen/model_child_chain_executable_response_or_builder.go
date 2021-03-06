/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type ChildChainExecutableResponseOrBuilder struct {
	PassThroughData           *ByteString            `json:"passThroughData,omitempty"`
	LastLink                  bool                   `json:"lastLink,omitempty"`
	Suspend                   bool                   `json:"suspend,omitempty"`
	NextChildId               string                 `json:"nextChildId,omitempty"`
	NextChildIdBytes          *ByteString            `json:"nextChildIdBytes,omitempty"`
	PreviousChildId           string                 `json:"previousChildId,omitempty"`
	PreviousChildIdBytes      *ByteString            `json:"previousChildIdBytes,omitempty"`
	AllFields                 map[string]interface{} `json:"allFields,omitempty"`
	DescriptorForType         *Descriptor            `json:"descriptorForType,omitempty"`
	UnknownFields             *UnknownFieldSet       `json:"unknownFields,omitempty"`
	InitializationErrorString string                 `json:"initializationErrorString,omitempty"`
	DefaultInstanceForType    *Message               `json:"defaultInstanceForType,omitempty"`
	Initialized               bool                   `json:"initialized,omitempty"`
}
