package har

// LifecycleRuleAction defines the action for a lifecycle rule.
type LifecycleRuleAction string

const (
	LifecycleRuleActionDELETE  LifecycleRuleAction = "DELETE"
	LifecycleRuleActionPROTECT LifecycleRuleAction = "PROTECT"
)

// LifecycleRuleApplyToMode defines how the rule applies to registries.
type LifecycleRuleApplyToMode string

const (
	LifecycleRuleApplyToModeALLINSCOPE LifecycleRuleApplyToMode = "ALL_IN_SCOPE"
	LifecycleRuleApplyToModeEXPLICIT   LifecycleRuleApplyToMode = "EXPLICIT"
)

// LifecycleRuleCriteriaMatch defines the match mode for criteria.
type LifecycleRuleCriteriaMatch string

const (
	LifecycleRuleCriteriaMatchALL LifecycleRuleCriteriaMatch = "ALL"
	LifecycleRuleCriteriaMatchANY LifecycleRuleCriteriaMatch = "ANY"
)

// LifecycleRuleCriteriaType defines the type of cleanup criteria.
type LifecycleRuleCriteriaType string

const (
	LifecycleRuleCriteriaTypeAGEBASED  LifecycleRuleCriteriaType = "AGE_BASED"
	LifecycleRuleCriteriaTypeKEEPLASTN LifecycleRuleCriteriaType = "KEEP_LAST_N"
	LifecycleRuleCriteriaTypeUNUSEDFOR LifecycleRuleCriteriaType = "UNUSED_FOR"
)

// LifecycleRuleCriteriaUnit defines the time unit for age/unused-for criteria.
type LifecycleRuleCriteriaUnit string

const (
	LifecycleRuleCriteriaUnitDAYS   LifecycleRuleCriteriaUnit = "DAYS"
	LifecycleRuleCriteriaUnitMONTHS LifecycleRuleCriteriaUnit = "MONTHS"
	LifecycleRuleCriteriaUnitYEARS  LifecycleRuleCriteriaUnit = "YEARS"
)

// LifecycleExecutionStatus defines the status of a lifecycle execution.
type LifecycleExecutionStatus string

const (
	LifecycleExecutionStatusPENDING    LifecycleExecutionStatus = "PENDING"
	LifecycleExecutionStatusINPROGRESS LifecycleExecutionStatus = "INPROGRESS"
	LifecycleExecutionStatusSUCCESS    LifecycleExecutionStatus = "SUCCESS"
	LifecycleExecutionStatusFAILED     LifecycleExecutionStatus = "FAILED"
)

// LifecycleExecutionTriggerType defines how a lifecycle execution was triggered.
type LifecycleExecutionTriggerType string

const (
	LifecycleExecutionTriggerTypeSCHEDULED LifecycleExecutionTriggerType = "SCHEDULED"
	LifecycleExecutionTriggerTypeDRYRUN    LifecycleExecutionTriggerType = "DRY_RUN"
)

// LifecycleRuleApplyTo defines which registries the rule applies to.
type LifecycleRuleApplyTo struct {
	Mode       LifecycleRuleApplyToMode `json:"mode"`
	Registries []string                 `json:"registries,omitempty"`
}

// LifecycleRuleSchedule defines the cron schedule for a lifecycle rule.
type LifecycleRuleSchedule struct {
	Expression string `json:"expression"`
	Timezone   string `json:"timezone"`
}

// LifecycleRuleCriteriaConfig defines the configuration for a single criteria item.
type LifecycleRuleCriteriaConfig struct {
	Unit  *LifecycleRuleCriteriaUnit `json:"unit,omitempty"`
	Value *int64                     `json:"value,omitempty"`
}

// LifecycleRuleCriteriaItem is a single criterion within a lifecycle rule.
type LifecycleRuleCriteriaItem struct {
	Config LifecycleRuleCriteriaConfig `json:"config"`
	Type   LifecycleRuleCriteriaType   `json:"type"`
}

// LifecycleRuleCriteria defines the full criteria for a lifecycle rule.
type LifecycleRuleCriteria struct {
	Match LifecycleRuleCriteriaMatch  `json:"match"`
	Rules []LifecycleRuleCriteriaItem `json:"rules"`
}

// LifecycleRuleFilterConfigDocker holds filter config for DOCKER package type.
type LifecycleRuleFilterConfigDocker struct {
	PackageType               string   `json:"packageType"` // always "DOCKER"
	PackageNameAllowedPattern []string `json:"packageNameAllowedPattern,omitempty"`
	TagNameAllowedPattern     []string `json:"tagNameAllowedPattern,omitempty"`
}

// LifecycleRuleFilterConfigMaven holds filter config for MAVEN package type.
type LifecycleRuleFilterConfigMaven struct {
	PackageType               string   `json:"packageType"` // always "MAVEN"
	GroupIdAllowedPattern     []string `json:"groupIdAllowedPattern,omitempty"`
	PackageNameAllowedPattern []string `json:"packageNameAllowedPattern,omitempty"`
	VersionNameAllowedPattern []string `json:"versionNameAllowedPattern,omitempty"`
}

// LifecycleRuleFilterConfigHuggingFace holds filter config for HUGGINGFACE package type.
type LifecycleRuleFilterConfigHuggingFace struct {
	PackageType               string   `json:"packageType"` // always "HUGGINGFACE"
	ModelAllowedPattern       []string `json:"modelAllowedPattern,omitempty"`
	DatasetAllowedPattern     []string `json:"datasetAllowedPattern,omitempty"`
	VersionNameAllowedPattern []string `json:"versionNameAllowedPattern,omitempty"`
}

// LifecycleRuleFilterConfigGeneric holds filter config for all other package types.
type LifecycleRuleFilterConfigGeneric struct {
	PackageType               string   `json:"packageType"`
	PackageNameAllowedPattern []string `json:"packageNameAllowedPattern,omitempty"`
	VersionNameAllowedPattern []string `json:"versionNameAllowedPattern,omitempty"`
}

// LifecycleRuleRequest is the request body for creating or updating a lifecycle rule.
type LifecycleRuleRequest struct {
	Name        string                      `json:"name"`
	Action      LifecycleRuleAction         `json:"action"`
	ApplyTo     LifecycleRuleApplyTo        `json:"applyTo"`
	Description *string                     `json:"description,omitempty"`
	PackageType *string                     `json:"packageType,omitempty"`
	Criteria    *LifecycleRuleCriteria      `json:"criteria,omitempty"`
	FilterConfig interface{}                `json:"filterConfig,omitempty"`
	Schedule    *LifecycleRuleSchedule      `json:"schedule,omitempty"`
}

// UserInfo contains basic user information returned in API responses.
type UserInfo struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

// LifecycleRuleResponse is the API response for a single lifecycle rule.
type LifecycleRuleResponse struct {
	Id                 string                 `json:"id"`
	Name               string                 `json:"name"`
	Action             LifecycleRuleAction    `json:"action"`
	ApplyTo            LifecycleRuleApplyTo   `json:"applyTo"`
	Description        *string                `json:"description,omitempty"`
	PackageType        *string                `json:"packageType,omitempty"`
	Enabled            bool                   `json:"enabled"`
	HasLocalAssignment bool                   `json:"hasLocalAssignment"`
	Criteria           *LifecycleRuleCriteria `json:"criteria,omitempty"`
	FilterConfig       interface{}            `json:"filterConfig,omitempty"`
	Schedule           *LifecycleRuleSchedule `json:"schedule,omitempty"`
	OrgIdentifier      *string                `json:"orgIdentifier,omitempty"`
	ProjectIdentifier  *string                `json:"projectIdentifier,omitempty"`
	CreatedAt          *int64                 `json:"createdAt,omitempty"`
	UpdatedAt          *int64                 `json:"updatedAt,omitempty"`
	LastRunAt          *int64                 `json:"lastRunAt,omitempty"`
	NextRunAt          *int64                 `json:"nextRunAt,omitempty"`
	CreatedBy          *UserInfo              `json:"createdBy,omitempty"`
	UpdatedBy          *UserInfo              `json:"updatedBy,omitempty"`
}

// LifecycleExecutionResponse is a single lifecycle execution record.
type LifecycleExecutionResponse struct {
	Id                 string                        `json:"id"`
	PolicyId           string                        `json:"policyId"`
	PolicyName         string                        `json:"policyName"`
	Status             LifecycleExecutionStatus      `json:"status"`
	TriggerType        LifecycleExecutionTriggerType `json:"triggerType"`
	CreatedAt          int64                         `json:"createdAt"`
	StartedAt          *int64                        `json:"startedAt,omitempty"`
	CompletedAt        *int64                        `json:"completedAt,omitempty"`
	VersionsDeleted    int64                         `json:"versionsDeleted"`
	Protected          int64                         `json:"protected"`
	PackagesAffected   int64                         `json:"packagesAffected"`
	RegistriesAffected int64                         `json:"registriesAffected"`
	StorageReclaimed   string                        `json:"storageReclaimed"`
	Message            *string                       `json:"message,omitempty"`
	OrgIdentifier      *string                       `json:"orgIdentifier,omitempty"`
	ProjectIdentifier  *string                       `json:"projectIdentifier,omitempty"`
}

// LifecycleDryRunRequest is the request body for triggering a dry-run execution.
type LifecycleDryRunRequest struct {
	PolicyId string `json:"policyId"`
}

// LifecycleStatsResponse holds aggregate lifecycle statistics.
type LifecycleStatsResponse struct {
	ArtifactsCleaned     int64  `json:"artifactsCleaned"`
	ExecutionsLast30Days int64  `json:"executionsLast30Days"`
	StorageReclaimed     string `json:"storageReclaimed"`
}

// ListLifecycleRulesResponseWrapper wraps a paginated lifecycle rules list response.
type ListLifecycleRulesResponseWrapper struct {
	HasMore bool                    `json:"hasMore"`
	Items   []LifecycleRuleResponse `json:"items"`
	Page    int64                   `json:"page"`
	Size    int64                   `json:"size"`
}
