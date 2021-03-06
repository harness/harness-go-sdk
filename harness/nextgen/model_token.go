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

// This has the API Key Token details defined in Harness.
type Token struct {
	// Identifier of the Token
	Identifier string `json:"identifier,omitempty"`
	// Name of the Token
	Name string `json:"name,omitempty"`
	// This is the time from which the Token is valid. The time is in milliseconds.
	ValidFrom int64 `json:"validFrom,omitempty"`
	// This is the time till which the Token is valid. The time is in milliseconds.
	ValidTo int64 `json:"validTo,omitempty"`
	// Scheduled expiry time in milliseconds.
	ScheduledExpireTime int64 `json:"scheduledExpireTime,omitempty"`
	// Boolean value to indicate if Token is valid or not.
	Valid bool `json:"valid,omitempty"`
	// Account Identifier for the Entity.
	AccountIdentifier string `json:"accountIdentifier,omitempty"`
	// Project Identifier for the Entity.
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
	// Organization Identifier for the Entity.
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	// This is the API Key Id within which the Token is created.
	ApiKeyIdentifier string `json:"apiKeyIdentifier,omitempty"`
	// This is the ID of the Parent entity from which the Token inherits its role bindings.
	ParentIdentifier string `json:"parentIdentifier,omitempty"`
	// Type of the API Key
	ApiKeyType string `json:"apiKeyType,omitempty"`
	// Description of the Token
	Description string `json:"description,omitempty"`
	// Tags for the Token
	Tags map[string]string `json:"tags,omitempty"`
	// Email Id of the user who created the Token.
	Email string `json:"email,omitempty"`
	// Name of the user who created the Token.
	Username        string `json:"username,omitempty"`
	EncodedPassword string `json:"encodedPassword,omitempty"`
}
