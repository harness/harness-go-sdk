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

// Get migration state request input
type MigrationStateIn struct {
	InstanceFilter *MigrationStateInstanceIn `json:"instanceFilter,omitempty"`
}