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

type OverviewResponseOutput struct {
	DbSchemaIdentifier   string `json:"dbSchemaIdentifier"`
	DbSchemaName         string `json:"dbSchemaName"`
	DbInstanceIdentifier string `json:"dbInstanceIdentifier"`
	DbInstanceName       string `json:"dbInstanceName"`
	LastDeployed         int64  `json:"lastDeployed"`
}