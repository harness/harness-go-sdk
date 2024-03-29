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

// This is the view of the UserGroup entity defined in Harness
type UserGroup struct {
	AccountIdentifier    string                         `json:"accountIdentifier,omitempty"`
	OrgIdentifier        string                         `json:"orgIdentifier,omitempty"`
	ProjectIdentifier    string                         `json:"projectIdentifier,omitempty"`
	Identifier           string                         `json:"identifier"`
	Name                 string                         `json:"name"`
	Users                []string                       `json:"users,omitempty"`
	NotificationConfigs  []NotificationSettingConfigDto `json:"notificationConfigs,omitempty"`
	IsSsoLinked          bool                           `json:"isSsoLinked,omitempty"`
	LinkedSsoId          string                         `json:"linkedSsoId,omitempty"`
	LinkedSsoType        string                         `json:"linkedSsoType,omitempty"`
	LinkedSsoDisplayName string                         `json:"linkedSsoDisplayName,omitempty"`
	SsoGroupId           string                         `json:"ssoGroupId,omitempty"`
	SsoGroupName         string                         `json:"ssoGroupName,omitempty"`
	ExternallyManaged    bool                           `json:"externallyManaged,omitempty"`
	Description          string                         `json:"description,omitempty"`
	Tags                 map[string]string              `json:"tags,omitempty"`
	SsoLinked            bool                           `json:"ssoLinked,omitempty"`
}
