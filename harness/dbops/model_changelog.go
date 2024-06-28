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

// location of the changelog file containing schema changes in a git repository
type Changelog struct {
	// identifier of the harness git connector
	Connector string `json:"connector"`
	// path to the change log file
	Location string `json:"location"`
	// repo name of the git based connector when ConnectionType is Account
	Repo string `json:"repo,omitempty"`
}