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

import (
	"time"
)

// Cloud Cost Report Schedule contains definition of 'how often' and 'to whom' the Report will be sent
type CeReportSchedule struct {
	Uuid             string        `json:"uuid,omitempty"`
	Name             string        `json:"name,omitempty"`
	Enabled          bool          `json:"enabled,omitempty"`
	Description      string        `json:"description,omitempty"`
	ViewsId          []string      `json:"viewsId"`
	UserCron         string        `json:"userCron,omitempty"`
	Recipients       []string      `json:"recipients,omitempty"`
	AccountId        string        `json:"accountId,omitempty"`
	CreatedAt        int64         `json:"createdAt,omitempty"`
	LastUpdatedAt    int64         `json:"lastUpdatedAt,omitempty"`
	UserCronTimeZone string        `json:"userCronTimeZone,omitempty"`
	CreatedBy        *EmbeddedUser `json:"createdBy,omitempty"`
	LastUpdatedBy    *EmbeddedUser `json:"lastUpdatedBy,omitempty"`
	NextExecution    time.Time     `json:"nextExecution,omitempty"`
}