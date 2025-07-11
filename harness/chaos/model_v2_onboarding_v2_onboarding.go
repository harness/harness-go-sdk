/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type V2OnboardingV2Onboarding struct {
	AccountID string `json:"accountID"`
	AutoCreateNetworkMapStatus *V2OnboardingAutoCreatedNetworkMapStatus `json:"autoCreateNetworkMapStatus,omitempty"`
	ChaosAdvanceConfiguration *V2OnboardingChaosInfraAdvanceConfiguration `json:"chaosAdvanceConfiguration,omitempty"`
	CreatedAt int32 `json:"createdAt,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	DdciID string `json:"ddciID,omitempty"`
	DiscoveredServiceStatus *V2OnboardingDiscoveredService `json:"discoveredServiceStatus,omitempty"`
	DiscoveryAdvanceConfiguration *DatabaseAgentConfiguration `json:"discoveryAdvanceConfiguration,omitempty"`
	EnvironmentRef string `json:"environmentRef,omitempty"`
	InfrastructureRef string `json:"infrastructureRef,omitempty"`
	IsRemoved bool `json:"isRemoved"`
	Mode *V2OnboardingOnboardingMode `json:"mode,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	OnboardingID string `json:"onboardingID,omitempty"`
	OrgID string `json:"orgID,omitempty"`
	ProjectID string `json:"projectID,omitempty"`
	RunSafeExperimentStatus *V2OnboardingRunSafeExperimentStatus `json:"runSafeExperimentStatus,omitempty"`
	ServiceAccount string `json:"serviceAccount,omitempty"`
	Status *V2OnboardingV2OnboardingStatus `json:"status,omitempty"`
	TargetNetworkMapStatus *V2OnboardingTargetNetworkMapStatus `json:"targetNetworkMapStatus,omitempty"`
	UpdatedAt int32 `json:"updatedAt,omitempty"`
	UpdatedBy string `json:"updatedBy,omitempty"`
}
