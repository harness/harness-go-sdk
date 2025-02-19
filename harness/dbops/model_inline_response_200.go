/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the DB Service
 *
 * API version: 1.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dbops

type InlineResponse200 struct {
	Image            string            `json:"image,omitempty"`
	DbImages         *interface{}      `json:"dbImages,omitempty"`
	DownloadImages   *interface{}      `json:"downloadImages,omitempty"`
	Settings         map[string]string `json:"settings,omitempty"`
	DefaultConfigs   map[string]string `json:"defaultConfigs,omitempty"`
	Schema           *DbSchemaOut      `json:"schema,omitempty"`
	Instance         *DbInstanceOut    `json:"instance,omitempty"`
	PropertiesResult map[string]string `json:"propertiesResult,omitempty"`
}
