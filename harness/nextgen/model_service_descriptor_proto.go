/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type ServiceDescriptorProto struct {
	UnknownFields             *UnknownFieldSet                 `json:"unknownFields,omitempty"`
	MethodOrBuilderList       []MethodDescriptorProtoOrBuilder `json:"methodOrBuilderList,omitempty"`
	SerializedSize            int32                            `json:"serializedSize,omitempty"`
	ParserForType             *ParserServiceDescriptorProto    `json:"parserForType,omitempty"`
	DefaultInstanceForType    *ServiceDescriptorProto          `json:"defaultInstanceForType,omitempty"`
	Options                   *ServiceOptions                  `json:"options,omitempty"`
	Initialized               bool                             `json:"initialized,omitempty"`
	Name                      string                           `json:"name,omitempty"`
	MethodList                []MethodDescriptorProto          `json:"methodList,omitempty"`
	NameBytes                 *ByteString                      `json:"nameBytes,omitempty"`
	OptionsOrBuilder          *ServiceOptionsOrBuilder         `json:"optionsOrBuilder,omitempty"`
	MethodCount               int32                            `json:"methodCount,omitempty"`
	InitializationErrorString string                           `json:"initializationErrorString,omitempty"`
	DescriptorForType         *Descriptor                      `json:"descriptorForType,omitempty"`
	AllFields                 map[string]interface{}           `json:"allFields,omitempty"`
	MemoizedSerializedSize    int32                            `json:"memoizedSerializedSize,omitempty"`
}