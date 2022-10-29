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

type CertificatesRepositoryCertificate struct {
	ServerName string `json:"serverName,omitempty"`
	CertType string `json:"certType,omitempty"`
	CertSubType string `json:"certSubType,omitempty"`
	CertData string `json:"certData,omitempty"`
	CertInfo string `json:"certInfo,omitempty"`
}