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

type GovernanceMetadata struct {
	UnknownFields             *UnknownFieldSet             `json:"unknownFields,omitempty"`
	TypeBytes                 *ByteString                  `json:"typeBytes,omitempty"`
	SerializedSize            int32                        `json:"serializedSize,omitempty"`
	ParserForType             *ParserGovernanceMetadata    `json:"parserForType,omitempty"`
	DefaultInstanceForType    *GovernanceMetadata          `json:"defaultInstanceForType,omitempty"`
	Status                    string                       `json:"status,omitempty"`
	Deny                      bool                         `json:"deny,omitempty"`
	Initialized               bool                         `json:"initialized,omitempty"`
	Timestamp                 int64                        `json:"timestamp,omitempty"`
	Message                   string                       `json:"message,omitempty"`
	Id                        string                       `json:"id,omitempty"`
	Type_                     string                       `json:"type,omitempty"`
	MessageBytes              *ByteString                  `json:"messageBytes,omitempty"`
	OrgId                     string                       `json:"orgId,omitempty"`
	StatusBytes               *ByteString                  `json:"statusBytes,omitempty"`
	AccountId                 string                       `json:"accountId,omitempty"`
	AccountIdBytes            *ByteString                  `json:"accountIdBytes,omitempty"`
	OrgIdBytes                *ByteString                  `json:"orgIdBytes,omitempty"`
	ProjectId                 string                       `json:"projectId,omitempty"`
	ProjectIdBytes            *ByteString                  `json:"projectIdBytes,omitempty"`
	Created                   int64                        `json:"created,omitempty"`
	IdBytes                   *ByteString                  `json:"idBytes,omitempty"`
	DetailsList               []PolicySetMetadata          `json:"detailsList,omitempty"`
	DetailsOrBuilderList      []PolicySetMetadataOrBuilder `json:"detailsOrBuilderList,omitempty"`
	DetailsCount              int32                        `json:"detailsCount,omitempty"`
	EntityBytes               *ByteString                  `json:"entityBytes,omitempty"`
	Action                    string                       `json:"action,omitempty"`
	ActionBytes               *ByteString                  `json:"actionBytes,omitempty"`
	Entity                    string                       `json:"entity,omitempty"`
	InitializationErrorString string                       `json:"initializationErrorString,omitempty"`
	DescriptorForType         *Descriptor                  `json:"descriptorForType,omitempty"`
	AllFields                 map[string]interface{}       `json:"allFields,omitempty"`
	MemoizedSerializedSize    int32                        `json:"memoizedSerializedSize,omitempty"`
}