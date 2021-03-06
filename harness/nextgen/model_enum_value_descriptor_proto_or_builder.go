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

type EnumValueDescriptorProtoOrBuilder struct {
	Options                   *EnumValueOptions          `json:"options,omitempty"`
	OptionsOrBuilder          *EnumValueOptionsOrBuilder `json:"optionsOrBuilder,omitempty"`
	NameBytes                 *ByteString                `json:"nameBytes,omitempty"`
	Name                      string                     `json:"name,omitempty"`
	Number                    int32                      `json:"number,omitempty"`
	AllFields                 map[string]interface{}     `json:"allFields,omitempty"`
	DescriptorForType         *Descriptor                `json:"descriptorForType,omitempty"`
	UnknownFields             *UnknownFieldSet           `json:"unknownFields,omitempty"`
	InitializationErrorString string                     `json:"initializationErrorString,omitempty"`
	DefaultInstanceForType    *Message                   `json:"defaultInstanceForType,omitempty"`
	Initialized               bool                       `json:"initialized,omitempty"`
}
