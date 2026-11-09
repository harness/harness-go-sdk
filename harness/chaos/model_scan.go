package chaos

import (
	"bytes"
	"encoding/json"
)

const (
	ScanTypePipelineExecution = "PipelineExecution"
	ScanTypeDiscoveryAgent    = "DiscoveryAgent"
	ScanTypeProbeExecution    = "ProbeExecution"
	ScanTypeCodeRepository    = "CodeRepository"

	DetectionMethodAI           = "AI"
	DetectionMethodRiskAnalyzer = "RiskAnalyzer"
	DetectionMethodExternal     = "External"

	ScanStatusQueued     = "QUEUED"
	ScanStatusCollecting = "COLLECTING"
	ScanStatusCollected  = "COLLECTED"
	ScanStatusAnalysing  = "ANALYSING"
	ScanStatusCompleted  = "COMPLETED"
	ScanStatusErrored    = "ERRORED"
	ScanStatusAborted    = "ABORTED"
)

type Scope struct {
	AccountID      string
	OrgID          string
	ProjectID      string
	ParentUniqueID string
	CorrelationID  string
}

// UserRef is a user object or a plain user id.
type UserRef struct {
	UserID   string `json:"userID,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

func (u *UserRef) UnmarshalJSON(data []byte) error {
	*u = UserRef{}
	data = bytes.TrimSpace(data)
	if len(data) == 0 || string(data) == "null" {
		return nil
	}
	if data[0] == '"' {
		var id string
		if err := json.Unmarshal(data, &id); err != nil {
			return err
		}
		u.UserID = id
		return nil
	}
	type alias UserRef
	var decoded alias
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	*u = UserRef(decoded)
	return nil
}

func (u UserRef) DisplayName() string {
	if u.Username != "" {
		return u.Username
	}
	if u.Email != "" {
		return u.Email
	}
	return u.UserID
}

type Scan struct {
	ID             string   `json:"id,omitempty"`
	AccountID      string   `json:"accountID,omitempty"`
	OrgID          string   `json:"orgID,omitempty"`
	ProjectID      string   `json:"projectID,omitempty"`
	ParentUniqueID string   `json:"parentUniqueID,omitempty"`
	UniqueID       string   `json:"uniqueID,omitempty"`
	Identity       string   `json:"identity,omitempty"`
	Name           string   `json:"name,omitempty"`
	Description    string   `json:"description,omitempty"`
	Tags           []string `json:"tags,omitempty"`

	ScanType        string     `json:"scanType,omitempty"`
	Status          string     `json:"status,omitempty"`
	Result          string     `json:"result,omitempty"`
	DetectionMethod string     `json:"detectionMethod,omitempty"`
	Source          ScanSource `json:"source"`

	CollectedData        *CollectionSummary    `json:"collectedData,omitempty"`
	RiskSummary          *RiskSummary          `json:"riskSummary,omitempty"`
	RecommendationCounts *RecommendationCounts `json:"recommendationCounts,omitempty"`
	Report               *ReportInfo           `json:"report,omitempty"`
	Summary              *HeatmapSummary       `json:"summary,omitempty"`
	ErrorMessage         string                `json:"errorMessage,omitempty"`
	RetryCount           int                   `json:"retryCount,omitempty"`
	CreatedAt            int64                 `json:"createdAt,omitempty"`
	UpdatedAt            int64                 `json:"updatedAt,omitempty"`
	CompletedAt          int64                 `json:"completedAt,omitempty"`
	CreatedBy            UserRef               `json:"createdBy,omitempty"`
	UpdatedBy            UserRef               `json:"updatedBy,omitempty"`
}

type ScanSource struct {
	Pipeline       *PipelineScan       `json:"pipeline,omitempty"`
	DiscoveryAgent *DiscoveryAgentScan `json:"discoveryAgent,omitempty"`
	ProbeExecution *ProbeExecutionScan `json:"probeExecution,omitempty"`
	CodeRepository *CodeRepositoryScan `json:"codeRepository,omitempty"`
}

type CodeRepositoryScan struct {
	RepoURL        string          `json:"repoURL,omitempty"`
	Branch         string          `json:"branch,omitempty"`
	CommitSHA      string          `json:"commitSHA,omitempty"`
	Provider       string          `json:"provider,omitempty"`
	LinkedServices []LinkedService `json:"linkedServices,omitempty"`
}

type LinkedService struct {
	Module                    string                   `json:"module,omitempty"`
	ChaosServiceIdentity      string                   `json:"chaosServiceIdentity,omitempty"`
	ChaosServiceName          string                   `json:"chaosServiceName,omitempty"`
	MatchedByRepository       bool                     `json:"matchedByRepository,omitempty"`
	LinkedToDiscoveredService *bool                    `json:"linkedToDiscoveredService,omitempty"`
	LinkedDiscoveredService   *LinkedDiscoveredService `json:"linkedDiscoveredService,omitempty"`
}

type LinkedDiscoveredService struct {
	Identity            string `json:"identity,omitempty"`
	EnvironmentIdentity string `json:"environmentIdentity,omitempty"`
	AgentIdentity       string `json:"agentIdentity,omitempty"`
	ExternalID          string `json:"externalID,omitempty"`
}

type RecommendationCounts struct {
	Total int `json:"total,omitempty"`
}

type CodeLocation struct {
	File      string `json:"file,omitempty"`
	Snippet   string `json:"snippet,omitempty"`
	StartLine int    `json:"startLine,omitempty"`
	EndLine   int    `json:"endLine,omitempty"`
}

type PipelineScan struct {
	PipelineIdentity    string                   `json:"pipelineIdentity"`
	PipelineExecutionID string                   `json:"pipelineExecutionID,omitempty"`
	NotifyID            string                   `json:"notifyID,omitempty"`
	Infrastructures     []PipelineInfrastructure `json:"infrastructures,omitempty"`
}

type PipelineInfrastructure struct {
	EnvironmentRef         string `json:"environmentRef,omitempty"`
	InfrastructureIdentity string `json:"infrastructureIdentity,omitempty"`
	ChaosOnboarded         bool   `json:"chaosOnboarded"`
	ChaosSupported         bool   `json:"chaosSupported"`
}

type DiscoveryAgentScan struct {
	AgentIdentity             string `json:"agentIdentity"`
	EnvironmentIdentity       string `json:"environmentIdentity"`
	EnvironmentParentUniqueID string `json:"environmentParentUniqueId,omitempty"`
}

type ProbeExecutionScan struct {
	ProbeIdentity   string                `json:"probeIdentity,omitempty"`
	ProbeName       string                `json:"probeName,omitempty"`
	ExecutionSource *ProbeExecutionSource `json:"executionSource,omitempty"`
}

type ProbeExecutionSource struct {
	Experiment   *ProbeExperimentSource   `json:"experiment,omitempty"`
	PipelineStep *ProbePipelineStepSource `json:"pipelineStep,omitempty"`
}

type ProbeExperimentSource struct {
	ExperimentID    string `json:"experimentID,omitempty"`
	ExperimentRunID string `json:"experimentRunID,omitempty"`
}

type ProbePipelineStepSource struct {
	PipelineID    string `json:"pipelineID,omitempty"`
	PipelineRunID string `json:"pipelineRunID,omitempty"`
}

type PipelineScanRequest struct {
	Identity            string
	Name                string
	Description         string
	Tags                []string
	DetectionMethod     string
	PipelineIdentity    string
	PipelineExecutionID string
	NotifyID            string
}

type InfrastructureScanRequest struct {
	Identity         string
	Name             string
	Description      string
	Tags             []string
	DetectionMethod  string
	InfrastructureID string
	EnvironmentID    string
}

type ScanCreateRequest struct {
	Identity        string     `json:"identity"`
	Name            string     `json:"name"`
	Description     string     `json:"description,omitempty"`
	Tags            []string   `json:"tags,omitempty"`
	ScanType        string     `json:"scanType"`
	DetectionMethod string     `json:"detectionMethod,omitempty"`
	Source          ScanSource `json:"source"`
}

// ScanUpdateRequest updates name, description, and tags. A nil tag list is omitted.
type ScanUpdateRequest struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Tags        []string `json:"-"`
}

func (r ScanUpdateRequest) MarshalJSON() ([]byte, error) {
	payload := map[string]interface{}{}
	if r.Name != nil {
		payload["name"] = *r.Name
	}
	if r.Description != nil {
		payload["description"] = *r.Description
	}
	if r.Tags != nil {
		payload["tags"] = r.Tags
	}
	return json.Marshal(payload)
}

func (r *ScanUpdateRequest) UnmarshalJSON(data []byte) error {
	var payload struct {
		Name        *string   `json:"name"`
		Description *string   `json:"description"`
		Tags        *[]string `json:"tags"`
	}
	if err := json.Unmarshal(data, &payload); err != nil {
		return err
	}
	r.Name = payload.Name
	r.Description = payload.Description
	if payload.Tags == nil {
		r.Tags = nil
		return nil
	}
	r.Tags = *payload.Tags
	return nil
}

type ScanListRequest struct {
	Scope
	Page                int
	Limit               int
	Search              string
	SortField           string
	SortAscending       *bool
	ScanType            string
	Status              string
	PipelineIdentity    string
	AgentIdentity       string
	EnvironmentIdentity string
	Tags                string
	StartTime           int64
	EndTime             int64
}

type ScanListResponse struct {
	CorrelationID string     `json:"correlationID,omitempty"`
	Data          []Scan     `json:"data"`
	Pagination    Pagination `json:"pagination"`
}

type Pagination struct {
	Index      int64 `json:"index,omitempty"`
	Limit      int64 `json:"limit,omitempty"`
	TotalItems int64 `json:"totalItems,omitempty"`
	TotalPages int64 `json:"totalPages,omitempty"`
}

type CollectionSummary struct {
	ScanType           string            `json:"scanType,omitempty"`
	CollectedAt        int64             `json:"collectedAt,omitempty"`
	TotalResources     int               `json:"totalResources,omitempty"`
	PrimaryResources   int               `json:"primaryResources,omitempty"`
	DependentResources int               `json:"dependentResources,omitempty"`
	ServiceGroups      int               `json:"serviceGroups,omitempty"`
	Namespaces         []string          `json:"namespaces,omitempty"`
	Errors             []CollectionError `json:"errors,omitempty"`
}

type CollectionError struct {
	Source  string `json:"source,omitempty"`
	Message string `json:"message,omitempty"`
}

type RiskSummary struct {
	RiskCount        int            `json:"riskCount,omitempty"`
	AnalysedAt       int64          `json:"analysedAt,omitempty"`
	SeverityCounts   map[string]int `json:"severityCounts,omitempty"`
	CategoryCounts   map[string]int `json:"categoryCounts,omitempty"`
	ServicesAnalysed int            `json:"servicesAnalysed,omitempty"`
	ServicesAtRisk   int            `json:"servicesAtRisk,omitempty"`
	Score            int            `json:"score,omitempty"`
	ScoreMax         int            `json:"scoreMax,omitempty"`
}

type ReportInfo struct {
	Status       string `json:"status,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	GeneratedAt  int64  `json:"generatedAt,omitempty"`
	ExpiresAt    int64  `json:"expiresAt,omitempty"`
}

type HeatmapSummary struct {
	Score           uint `json:"score,omitempty"`
	MaxScore        uint `json:"maxScore,omitempty"`
	ServicesScanned int  `json:"servicesScanned,omitempty"`
	TotalRisks      int  `json:"totalRisks,omitempty"`
	RuleCount       int  `json:"ruleCount,omitempty"`
	Critical        int  `json:"critical,omitempty"`
	High            int  `json:"high,omitempty"`
	Medium          int  `json:"medium,omitempty"`
	Low             int  `json:"low,omitempty"`
}

type ScanReport struct {
	Scan         *Scan         `json:"scan,omitempty"`
	ScannedRisks []ScanRisk    `json:"scannedRisks"`
	Summary      ReportSummary `json:"summary"`
}

type ReportSummary struct {
	TotalRisks       int            `json:"totalRisks,omitempty"`
	TotalOccurrences int            `json:"totalOccurrences,omitempty"`
	BySeverity       map[string]int `json:"bySeverity,omitempty"`
	ByStatus         map[string]int `json:"byStatus,omitempty"`
}

type ScanRisk struct {
	Identity        string                 `json:"identity,omitempty"`
	Name            string                 `json:"name,omitempty"`
	Description     string                 `json:"description,omitempty"`
	Tags            []string               `json:"tags,omitempty"`
	Fingerprint     string                 `json:"fingerprint,omitempty"`
	Severity        string                 `json:"severity,omitempty"`
	Status          string                 `json:"status,omitempty"`
	ServiceIdentity string                 `json:"serviceIdentity,omitempty"`
	ServiceName     string                 `json:"serviceName,omitempty"`
	ServiceType     string                 `json:"serviceType,omitempty"`
	Metadata        map[string]interface{} `json:"metadata,omitempty"`
	Occurrences     []ScanRiskOccurrence   `json:"occurrences,omitempty"`
}

type ScanRiskOccurrence struct {
	OccurrenceID         string                `json:"occurrenceId,omitempty"`
	DetectedAt           int64                 `json:"detectedAt,omitempty"`
	SourceType           string                `json:"sourceType,omitempty"`
	SourceID             string                `json:"sourceId,omitempty"`
	Context              OccurrenceContext     `json:"context"`
	ResolutionID         string                `json:"resolutionId,omitempty"`
	ResolvedAt           int64                 `json:"resolvedAt,omitempty"`
	ResolvedBy           string                `json:"resolvedBy,omitempty"`
	ResolutionMsg        string                `json:"resolutionMsg,omitempty"`
	ServiceName          string                `json:"serviceName,omitempty"`
	ProbeExecutionSource *ProbeExecutionSource `json:"probeExecutionSource,omitempty"`
}

type OccurrenceContext struct {
	Evidence              []string       `json:"evidence,omitempty"`
	Remediation           []string       `json:"remediation,omitempty"`
	ResilienceTestingReco []string       `json:"resilienceTestingRec,omitempty"`
	CodeLocations         []CodeLocation `json:"codeLocations,omitempty"`
}
