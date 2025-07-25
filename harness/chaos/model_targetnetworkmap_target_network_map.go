/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type TargetnetworkmapTargetNetworkMap struct {
	AccountID string `json:"accountID"`
	AgentIdentity string `json:"agentIdentity,omitempty"`
	AverageResiliencyScore float64 `json:"averageResiliencyScore,omitempty"`
	CreatedAt int32 `json:"createdAt,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	EnvironmentRef string `json:"environmentRef,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	ExperimentCreationMode *TargetnetworkmapExperimentCreationMode `json:"experimentCreationMode,omitempty"`
	Id string `json:"id,omitempty"`
	Identity string `json:"identity,omitempty"`
	InfrastructureId string `json:"infrastructureId,omitempty"`
	IsRemoved bool `json:"isRemoved"`
	Name string `json:"name,omitempty"`
	OnboardingID string `json:"onboardingID,omitempty"`
	OrgID string `json:"orgID,omitempty"`
	ProjectID string `json:"projectID,omitempty"`
	RecentExperimentRunsDetails []TargetnetworkmapExperimentRunDetail `json:"recentExperimentRunsDetails,omitempty"`
	ResiliencyCoverage float64 `json:"resiliencyCoverage,omitempty"`
	ServiceIDs []string `json:"serviceIDs,omitempty"`
	Status *TargetnetworkmapStatus `json:"status,omitempty"`
	TargetNetworkMapID string `json:"targetNetworkMapID,omitempty"`
	TotalExperimentCount int32 `json:"totalExperimentCount,omitempty"`
	TotalExperimentRunCount int32 `json:"totalExperimentRunCount,omitempty"`
	UpdatedAt int32 `json:"updatedAt,omitempty"`
	UpdatedBy string `json:"updatedBy,omitempty"`
}
