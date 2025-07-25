/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the Notification Management Service
 *
 * API version: 1.0
 * Contact: contact.harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// Notification resource and list of events
type NotificationResourceDto struct {
	ResourceName string   `json:"resource_name,omitempty"`
	Events       []string `json:"events,omitempty"`
	Enabled      bool     `json:"enabled,omitempty"`
}
