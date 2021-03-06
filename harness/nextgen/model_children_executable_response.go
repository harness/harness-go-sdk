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

type ChildrenExecutableResponse struct {
	UnknownFields             *UnknownFieldSet                  `json:"unknownFields,omitempty"`
	Initialized               bool                              `json:"initialized,omitempty"`
	ParserForType             *ParserChildrenExecutableResponse `json:"parserForType,omitempty"`
	SerializedSize            int32                             `json:"serializedSize,omitempty"`
	DefaultInstanceForType    *ChildrenExecutableResponse       `json:"defaultInstanceForType,omitempty"`
	ChildrenList              []Child                           `json:"childrenList,omitempty"`
	ChildrenOrBuilderList     []ChildOrBuilder                  `json:"childrenOrBuilderList,omitempty"`
	ChildrenCount             int32                             `json:"childrenCount,omitempty"`
	AllFields                 map[string]interface{}            `json:"allFields,omitempty"`
	DescriptorForType         *Descriptor                       `json:"descriptorForType,omitempty"`
	InitializationErrorString string                            `json:"initializationErrorString,omitempty"`
	MemoizedSerializedSize    int32                             `json:"memoizedSerializedSize,omitempty"`
}
