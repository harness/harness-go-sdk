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

// Deployment info of Change Sets across instances
type MigrationStateChangeSet struct {
	// Identifier of the Change Set
	ChangeSet string `json:"changeSet"`
	// If Change Set is successfully deployed to all instances
	Synced    bool                     `json:"synced"`
	Instances []MigrationStateInstance `json:"instances"`
	Updated   int64                    `json:"updated"`
	Author    string                   `json:"author"`
	FileName  string                   `json:"fileName"`
}
