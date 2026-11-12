package loadtest

// Load test values used by the Terraform resource. Requests are validated here
// and then sent through the generated Load Test Manager client.

const (
	LoadTestToolLocust  = "Locust"
	LoadTestToolJMeter  = "JMeter"
	LoadTestToolK6      = "K6"
	LoadTestToolGatling = "Gatling" // Legacy: read-only, rejected on create
	LoadTestToolCustom  = "Custom"  // Legacy: read-only, rejected on create

	LoadTestTargetLinux      = "linux"
	LoadTestTargetKubernetes = "kubernetes"

	CleanupPolicyDelete = "delete"
	CleanupPolicyRetain = "retain"

	LoadTestRunStatusPreparingInfra = "PreparingInfra"
	LoadTestRunStatusPending        = "Pending"
	LoadTestRunStatusRunning        = "Running"
	LoadTestRunStatusStopping       = "Stopping"
	LoadTestRunStatusStopped        = "Stopped"
	LoadTestRunStatusFinished       = "Finished"
	LoadTestRunStatusFailed         = "Failed"
)

// LoadTestCreateRequest is the body for POST /v1/load-tests.
type LoadTestCreateRequest struct {
	Identity              string                 `json:"identity"`
	Name                  string                 `json:"name"`
	Description           string                 `json:"description,omitempty"`
	ToolType              string                 `json:"toolType"`
	InfraType             string                 `json:"infraType,omitempty"`
	InfraIdentifier       string                 `json:"infraIdentifier,omitempty"`
	EnvironmentIdentifier string                 `json:"environmentIdentifier,omitempty"`
	Tags                  []string               `json:"tags,omitempty"`
	ManagedInfra          bool                   `json:"managedInfra,omitempty"`
	ServiceReferences     []string               `json:"serviceReferences,omitempty"`
	TargetType            string                 `json:"targetType,omitempty"`
	ScriptSource          string                 `json:"scriptSource,omitempty"`
	ScriptContent         string                 `json:"scriptContent,omitempty"`
	MaxDurationSec        *int                   `json:"maxDurationSec,omitempty"`
	Variables             []LoadTestVariable     `json:"variables,omitempty"`
	ToolConfig            map[string]interface{} `json:"toolConfig,omitempty"`
	UseSampleTest         bool                   `json:"useSampleTest,omitempty"`
	IsAIUsed              bool                   `json:"isAIUsed,omitempty"`
	YAML                  string                 `json:"yaml,omitempty"`
	CleanupPolicy         string                 `json:"cleanupPolicy,omitempty"`
	Resources             *LoadTestResources     `json:"resources,omitempty"`
	SkipScriptValidation  bool                   `json:"skipScriptValidation,omitempty"`
}

// LoadTestUpdateRequest is the body for PUT /v1/load-tests/{identity}.
// Tags and variables are always sent: an empty array replaces the stored list.
// Other empty fields are omitted, which tells the API to keep the stored value.
type LoadTestUpdateRequest struct {
	Name                  string                 `json:"name,omitempty"`
	Description           string                 `json:"description,omitempty"`
	Tags                  []string               `json:"tags"`
	ServiceReferences     []string               `json:"serviceReferences,omitempty"`
	InfraIdentifier       string                 `json:"infraIdentifier,omitempty"`
	EnvironmentIdentifier string                 `json:"environmentIdentifier,omitempty"`
	ManagedInfra          *bool                  `json:"managedInfra,omitempty"`
	TargetType            string                 `json:"targetType,omitempty"`
	ScriptSource          string                 `json:"scriptSource,omitempty"`
	ScriptContent         string                 `json:"scriptContent,omitempty"`
	MaxDurationSec        *int                   `json:"maxDurationSec,omitempty"`
	Variables             []LoadTestVariable     `json:"variables"`
	ToolConfig            map[string]interface{} `json:"toolConfig,omitempty"`
	IsAIUsed              bool                   `json:"isAIUsed,omitempty"`
	YAML                  string                 `json:"yaml,omitempty"`
	CleanupPolicy         string                 `json:"cleanupPolicy,omitempty"`
	Resources             *LoadTestResources     `json:"resources,omitempty"`
	SkipScriptValidation  bool                   `json:"skipScriptValidation,omitempty"`
}

// LoadTestTemplate is one revision of a hub template. Variables are read from
// the template variables endpoint, not this object.
type LoadTestTemplate struct {
	UniqueID      string                 `json:"uniqueId,omitempty"`
	Identity      string                 `json:"identity,omitempty"`
	HubIdentity   string                 `json:"hubIdentity,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Revision      string                 `json:"revision,omitempty"`
	ToolType      string                 `json:"toolType,omitempty"`
	ToolConfig    map[string]interface{} `json:"toolConfig,omitempty"`
	Tags          []string               `json:"tags,omitempty"`
	EnvID         string                 `json:"envId,omitempty"`
	InfraID       string                 `json:"infraId,omitempty"`
	InfraType     string                 `json:"infraType,omitempty"`
	ScriptContent string                 `json:"scriptContent,omitempty"`
	ScriptSource  string                 `json:"scriptSource,omitempty"`
}

// LoadTestTemplateRequest creates a hub template.
type LoadTestTemplateRequest struct {
	Identity             string                 `json:"identity,omitempty"`
	Name                 string                 `json:"name,omitempty"`
	Description          string                 `json:"description,omitempty"`
	Revision             string                 `json:"revision,omitempty"`
	ToolType             string                 `json:"toolType,omitempty"`
	ToolConfig           map[string]interface{} `json:"toolConfig,omitempty"`
	Variables            []LoadTestVariable     `json:"variables,omitempty"`
	Tags                 []string               `json:"tags,omitempty"`
	EnvID                string                 `json:"envId,omitempty"`
	InfraID              string                 `json:"infraId,omitempty"`
	InfraType            string                 `json:"infraType,omitempty"`
	ScriptContent        string                 `json:"scriptContent,omitempty"`
	ScriptSource         string                 `json:"scriptSource,omitempty"`
	SkipScriptValidation bool                   `json:"skipScriptValidation,omitempty"`
	YAML                 string                 `json:"yaml,omitempty"`
}

// LoadTestTemplateRevisionRequest creates one revision. Omitted fields are inherited.
type LoadTestTemplateRevisionRequest struct {
	Name                 string                 `json:"name,omitempty"`
	Description          string                 `json:"description,omitempty"`
	Revision             string                 `json:"revision,omitempty"`
	ToolConfig           map[string]interface{} `json:"toolConfig,omitempty"`
	Variables            []LoadTestVariable     `json:"variables,omitempty"`
	Tags                 []string               `json:"tags,omitempty"`
	EnvID                string                 `json:"envId,omitempty"`
	InfraID              string                 `json:"infraId,omitempty"`
	ScriptContent        string                 `json:"scriptContent,omitempty"`
	ScriptSource         string                 `json:"scriptSource,omitempty"`
	SkipScriptValidation bool                   `json:"skipScriptValidation,omitempty"`
	YAML                 string                 `json:"yaml,omitempty"`
}

// LoadTestTemplateUpdateRequest edits the latest revision in place.
// Tags and variables are always sent: an empty array replaces the stored list.
type LoadTestTemplateUpdateRequest struct {
	Name                 string                 `json:"name,omitempty"`
	Description          string                 `json:"description,omitempty"`
	ToolConfig           map[string]interface{} `json:"toolConfig,omitempty"`
	Variables            []LoadTestVariable     `json:"variables"`
	Tags                 []string               `json:"tags"`
	EnvID                string                 `json:"envId,omitempty"`
	InfraID              string                 `json:"infraId,omitempty"`
	InfraType            string                 `json:"infraType,omitempty"`
	SkipScriptValidation bool                   `json:"skipScriptValidation,omitempty"`
	YAML                 string                 `json:"yaml,omitempty"`
}

// TemplateReferencedByItem is one load test that references a hub template.
type TemplateReferencedByItem struct {
	Identity          string `json:"identity,omitempty"`
	Kind              string `json:"kind,omitempty"`
	Name              string `json:"name,omitempty"`
	OrgIdentifier     string `json:"orgIdentifier,omitempty"`
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
}

// TemplateReferencedByResponse is one page of load tests that reference a template.
type TemplateReferencedByResponse struct {
	Data       []TemplateReferencedByItem `json:"data,omitempty"`
	Pagination Pagination                 `json:"pagination,omitempty"`
}

// Scope is the account, org, and project for a load-test call.
type Scope struct {
	AccountID      string
	OrgID          string
	ProjectID      string
	ParentUniqueID string
	CorrelationID  string
}

// UserRef is the user object returned on createdByUserDetails.
type UserRef struct {
	UserID   string `json:"userID,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

// Pagination is the list page returned with load tests and runs.
type Pagination struct {
	Index      int64 `json:"index,omitempty"`
	Limit      int64 `json:"limit,omitempty"`
	TotalItems int64 `json:"totalItems,omitempty"`
	TotalPages int64 `json:"totalPages,omitempty"`
}

// LoadTest is a load test returned by Load Test Manager.
type LoadTest struct {
	UniqueID                 string                 `json:"uniqueId,omitempty"`
	Identity                 string                 `json:"identity,omitempty"`
	ParentUniqueID           string                 `json:"parentUniqueId,omitempty"`
	Name                     string                 `json:"name,omitempty"`
	Description              string                 `json:"description,omitempty"`
	ToolType                 string                 `json:"toolType,omitempty"`
	ToolLanguage             string                 `json:"toolLanguage,omitempty"`
	InfraType                string                 `json:"infraType,omitempty"`
	Tags                     []string               `json:"tags,omitempty"`
	ServiceReferences        []string               `json:"serviceReferences,omitempty"`
	AccountIdentifier        string                 `json:"accountIdentifier,omitempty"`
	OrganizationIdentifier   string                 `json:"organizationIdentifier,omitempty"`
	ProjectIdentifier        string                 `json:"projectIdentifier,omitempty"`
	InfraIdentifier          string                 `json:"infraIdentifier,omitempty"`
	EnvironmentIdentifier    string                 `json:"environmentIdentifier,omitempty"`
	TargetType               string                 `json:"targetType,omitempty"`
	ScriptSource             string                 `json:"scriptSource,omitempty"`
	ScriptContent            string                 `json:"scriptContent,omitempty"`
	IsAPIGenerated           bool                   `json:"isApiGenerated,omitempty"`
	IsSampleTest             bool                   `json:"isSampleTest,omitempty"`
	JSONScript               string                 `json:"jsonScript,omitempty"`
	JSONSpecData             *LoadTestJSONSpec      `json:"jsonSpecData,omitempty"`
	LatestRevisionIdentifier string                 `json:"latestRevisionIdentifier,omitempty"`
	ManagedInfra             bool                   `json:"managedInfra,omitempty"`
	TargetUsers              string                 `json:"targetUsers,omitempty"`
	DurationSeconds          string                 `json:"durationSeconds,omitempty"`
	MaxDurationSec           *int                   `json:"maxDurationSec,omitempty"`
	CleanupPolicy            string                 `json:"cleanupPolicy,omitempty"`
	Resources                *LoadTestResources     `json:"resources,omitempty"`
	ToolConfig               map[string]interface{} `json:"toolConfig,omitempty"`
	RecentRuns               []LoadTestRecentRun    `json:"recentRuns,omitempty"`
	LastExecuted             string                 `json:"lastExecuted,omitempty"`
	CreatedAt                string                 `json:"createdAt,omitempty"`
	CreatedBy                string                 `json:"createdBy,omitempty"`
	UpdatedAt                string                 `json:"updatedAt,omitempty"`
	UpdatedBy                string                 `json:"updatedBy,omitempty"`
	CreatedByUserDetails     *UserRef               `json:"createdByUserDetails,omitempty"`
	UpdatedByUserDetails     *UserRef               `json:"updatedByUserDetails,omitempty"`
	IsAIUsed                 bool                   `json:"isAIUsed,omitempty"`
	AISummary                string                 `json:"aiSummary,omitempty"`
	GenerateAISummary        bool                   `json:"generateAISummary,omitempty"`
	TemplateReference        *LoadTestTemplateRef   `json:"templateReference,omitempty"`
	ImportType               string                 `json:"importType,omitempty"`
	TemplateUpdateAvailable  bool                   `json:"templateUpdateAvailable,omitempty"`
	YAML                     string                 `json:"yaml,omitempty"`
}

// LoadTestVariablesResponse is the body for GET /v1/load-tests/{identity}/variables.
// Variables are the custom variables. Inputs are the tool tunables.
type LoadTestVariablesResponse struct {
	Inputs           []map[string]interface{} `json:"inputs,omitempty"`
	LoadTestIdentity string                   `json:"loadTestIdentity,omitempty"`
	Variables        []LoadTestVariable       `json:"variables,omitempty"`
}

// LoadTestListResponse is the body for GET /v1/load-tests.
type LoadTestListResponse struct {
	Items      []LoadTest `json:"items,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
}

// LoadTestSuccessResponse is the body for a successful delete or stop.
type LoadTestSuccessResponse struct {
	Message string `json:"message,omitempty"`
	Success bool   `json:"success,omitempty"`
}

// LoadTestResources is per-pod CPU and memory for Kubernetes load tests.
type LoadTestResources struct {
	Limits   map[string]string `json:"limits,omitempty"`
	Requests map[string]string `json:"requests,omitempty"`
}

// LoadTestVariable is a custom variable stored on the load test.
type LoadTestVariable struct {
	Name        string                  `json:"name"`
	Value       interface{}             `json:"value"`
	Path        string                  `json:"path,omitempty"`
	Type        string                  `json:"type,omitempty"`
	Description string                  `json:"description,omitempty"`
	Required    bool                    `json:"required"`
	Reference   *LoadTestInputReference `json:"reference,omitempty"`
}

// LoadTestInputReference points a variable or run value at another Harness resource.
type LoadTestInputReference struct {
	Kind string `json:"kind,omitempty"`
	Type string `json:"type,omitempty"`
}

// LoadTestJSONSpec is the optional JSON-authored API scenario stored on a load test.
type LoadTestJSONSpec struct {
	Config    map[string]interface{}   `json:"config,omitempty"`
	Endpoints []map[string]interface{} `json:"endpoints,omitempty"`
}

// LoadTestFromTemplateRequest is the body for POST /v1/load-tests/from-template.
type LoadTestFromTemplateRequest struct {
	Identity              string                 `json:"identity"`
	Name                  string                 `json:"name"`
	Description           string                 `json:"description,omitempty"`
	ImportType            string                 `json:"importType"`
	EnvironmentIdentifier string                 `json:"environmentIdentifier"`
	InfraIdentifier       string                 `json:"infraIdentifier"`
	TargetType            string                 `json:"targetType,omitempty"`
	ServiceReferences     []string               `json:"serviceReferences,omitempty"`
	TemplateReference     *LoadTestTemplateRef   `json:"templateReference"`
	ToolConfig            map[string]interface{} `json:"toolConfig,omitempty"`
	Variables             []LoadTestVariable     `json:"variables,omitempty"`
}

// LoadTestTemplateRef is the template a load test was created from.
type LoadTestTemplateRef struct {
	HubIdentity            string `json:"hubIdentity,omitempty"`
	Identity               string `json:"identity,omitempty"`
	OrganizationIdentifier string `json:"organizationIdentifier,omitempty"`
	ProjectIdentifier      string `json:"projectIdentifier,omitempty"`
	Revision               string `json:"revision,omitempty"`
}

// LoadTestRecentRun is a short run summary embedded on a load test.
type LoadTestRecentRun struct {
	UniqueID             string   `json:"uniqueId,omitempty"`
	Identity             string   `json:"identity,omitempty"`
	Name                 string   `json:"name,omitempty"`
	Status               string   `json:"status,omitempty"`
	TargetUsers          int      `json:"targetUsers,omitempty"`
	SpawnRate            float64  `json:"spawnRate,omitempty"`
	DurationSeconds      *int     `json:"durationSeconds,omitempty"`
	StartedAt            string   `json:"startedAt,omitempty"`
	FinishedAt           string   `json:"finishedAt,omitempty"`
	CreatedAt            string   `json:"createdAt,omitempty"`
	CreatedBy            string   `json:"createdBy,omitempty"`
	CreatedByUserDetails *UserRef `json:"createdByUserDetails,omitempty"`
	IsAIUsed             bool     `json:"isAIUsed,omitempty"`
}

// LoadTestRunCreateRequest is the body for POST /v1/load-tests/{identity}/runs.
// Identity is the new run id, chosen by the caller.
type LoadTestRunCreateRequest struct {
	Identity      string                 `json:"identity"`
	Name          string                 `json:"name,omitempty"`
	Values        []LoadTestRunValue     `json:"values,omitempty"`
	RuntimeValues map[string]interface{} `json:"runtimeValues,omitempty"`
	IsAIUsed      bool                   `json:"isAIUsed,omitempty"`
}

// LoadTestRunValue overrides one variable or input for a single run.
type LoadTestRunValue struct {
	Name      string                  `json:"name"`
	Value     interface{}             `json:"value"`
	Reference *LoadTestInputReference `json:"reference,omitempty"`
}

// LoadTestRun is a load test execution. Fields match internal_api.LoadTestRunResponse.
type LoadTestRun struct {
	AccountIdentifier        string                   `json:"accountIdentifier,omitempty"`
	Action                   string                   `json:"action,omitempty"`
	AIInsights               map[string]interface{}   `json:"aiInsights,omitempty"`
	CreatedAt                string                   `json:"createdAt,omitempty"`
	CreatedBy                string                   `json:"createdBy,omitempty"`
	CreatedByUserDetails     *UserRef                 `json:"createdByUserDetails,omitempty"`
	DDCILogStreamID          string                   `json:"ddciLogStreamId,omitempty"`
	DelegateLogUID           string                   `json:"delegateLogUid,omitempty"`
	DurationSeconds          *int                     `json:"durationSeconds,omitempty"`
	EnvironmentIdentifier    string                   `json:"environmentIdentifier,omitempty"`
	ErrorCode                string                   `json:"errorCode,omitempty"`
	ErrorMessage             string                   `json:"errorMessage,omitempty"`
	Errors                   []map[string]interface{} `json:"errors,omitempty"`
	FinishedAt               string                   `json:"finishedAt,omitempty"`
	Identity                 string                   `json:"identity,omitempty"`
	InfraIdentifier          string                   `json:"infraIdentifier,omitempty"`
	InfraPreparation         map[string]interface{}   `json:"infraPreparation,omitempty"`
	IsAIUsed                 bool                     `json:"isAIUsed,omitempty"`
	LastMetrics              map[string]interface{}   `json:"lastMetrics,omitempty"`
	LoadArgs                 string                   `json:"loadArgs,omitempty"`
	LoadTestIdentity         string                   `json:"loadTestIdentity,omitempty"`
	LoadTestName             string                   `json:"loadTestName,omitempty"`
	LogStreamID              string                   `json:"logStreamId,omitempty"`
	ManagedInfra             bool                     `json:"managedInfra,omitempty"`
	Name                     string                   `json:"name,omitempty"`
	NotifyID                 string                   `json:"notifyId,omitempty"`
	OrganizationIdentifier   string                   `json:"organizationIdentifier,omitempty"`
	ParentUniqueID           string                   `json:"parentUniqueId,omitempty"`
	ProjectIdentifier        string                   `json:"projectIdentifier,omitempty"`
	RampUpTimeSec            int                      `json:"rampUpTimeSec,omitempty"`
	ResolvedExpressionValues map[string]interface{}   `json:"resolvedExpressionValues,omitempty"`
	ResolvedRuntimeValues    map[string]interface{}   `json:"resolvedRuntimeValues,omitempty"`
	ResolvedVariables        map[string]interface{}   `json:"resolvedVariables,omitempty"`
	RunSequence              int                      `json:"runSequence,omitempty"`
	ScriptEntrypoint         string                   `json:"scriptEntrypoint,omitempty"`
	ScriptImage              string                   `json:"scriptImage,omitempty"`
	ScriptImagePullSecret    string                   `json:"scriptImagePullSecret,omitempty"`
	ScriptSource             string                   `json:"scriptSource,omitempty"`
	SpawnRate                float64                  `json:"spawnRate,omitempty"`
	StartedAt                string                   `json:"startedAt,omitempty"`
	Status                   string                   `json:"status,omitempty"`
	TargetType               string                   `json:"targetType,omitempty"`
	TargetUsers              int                      `json:"targetUsers,omitempty"`
	ThresholdResults         []map[string]interface{} `json:"thresholdResults,omitempty"`
	ThresholdVerdict         string                   `json:"thresholdVerdict,omitempty"`
	ToolLanguage             string                   `json:"toolLanguage,omitempty"`
	ToolType                 string                   `json:"toolType,omitempty"`
	UniqueID                 string                   `json:"uniqueId,omitempty"`
	UpdatedAt                string                   `json:"updatedAt,omitempty"`
	UpdatedBy                string                   `json:"updatedBy,omitempty"`
	UpdatedByUserDetails     *UserRef                 `json:"updatedByUserDetails,omitempty"`
	WorkerCount              int                      `json:"workerCount,omitempty"`
}

// LoadTestRunListResponse is the body for GET /v1/load-tests/{identity}/runs.
type LoadTestRunListResponse struct {
	Items      []LoadTestRun `json:"items,omitempty"`
	Pagination Pagination    `json:"pagination,omitempty"`
}

// ListLoadTestsRequest filters GET /v1/load-tests.
type ListLoadTestsRequest struct {
	Scope                 Scope
	Page                  int
	Limit                 int
	Search                string
	SortField             string
	SortAscending         *bool
	ToolType              string
	Tags                  []string
	EnvironmentIdentifier string
}

// ListLoadTestRunsRequest filters GET /v1/load-tests/{identity}/runs.
type ListLoadTestRunsRequest struct {
	Scope         Scope
	Page          int
	Limit         int
	Search        string
	SortField     string
	SortAscending *bool
	Status        string
}

// CompositeLoadTestRequest is POST /v1/composite-load-tests.
// Identifier is the stage id. The pipeline id is that value plus "_pipeline".
type CompositeLoadTestRequest struct {
	Name        string                 `json:"name"`
	Identifier  string                 `json:"identifier"`
	Description string                 `json:"description,omitempty"`
	Tags        map[string]string      `json:"tags,omitempty"`
	LoadTest    CompositeLoadTestRef   `json:"loadTest"`
	Probe       CompositeLoadTestProbe `json:"probe"`
}

// CompositeLoadTestUpdateRequest is PUT /v1/composite-load-tests/{identifier}.
type CompositeLoadTestUpdateRequest struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Tags        map[string]string      `json:"tags,omitempty"`
	LoadTest    CompositeLoadTestRef   `json:"loadTest"`
	Probe       CompositeLoadTestProbe `json:"probe"`
}

// CompositeLoadTest is one composite load test read from its pipeline.
type CompositeLoadTest struct {
	Identifier         string                 `json:"identifier,omitempty"`
	PipelineIdentifier string                 `json:"pipelineIdentifier,omitempty"`
	Name               string                 `json:"name,omitempty"`
	Description        string                 `json:"description,omitempty"`
	Tags               map[string]string      `json:"tags,omitempty"`
	LoadTest           CompositeLoadTestRef   `json:"loadTest,omitempty"`
	Probe              CompositeLoadTestProbe `json:"probe,omitempty"`
}

// CompositeLoadTestCreated is the create response. Identifier is the stage id.
type CompositeLoadTestCreated struct {
	Identifier         string `json:"identifier,omitempty"`
	Name               string `json:"name,omitempty"`
	PipelineIdentifier string `json:"pipelineIdentifier,omitempty"`
	StageIdentifier    string `json:"stageIdentifier,omitempty"`
}

// CompositeLoadTestRef is the load test identity in the same project.
type CompositeLoadTestRef struct {
	LoadTestRef string `json:"loadTestRef,omitempty"`
}

// CompositeLoadTestProbe is the chaos probe stored on the composite stage.
type CompositeLoadTestProbe struct {
	Identity       string `json:"identity,omitempty"`
	InfraReference string `json:"infraReference,omitempty"`
	Duration       string `json:"duration,omitempty"`
}
