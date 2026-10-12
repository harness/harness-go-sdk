package chaos

import (
	"bytes"
	"encoding/json"
)

const (
	// ScanTypePipelineExecution is a pipeline scan. The UI creates one of these
	// from a successful pipeline execution.
	ScanTypePipelineExecution = "PipelineExecution"
	// ScanTypeDiscoveryAgent is an infrastructure scan. The UI creates one of
	// these from a service-discovery agent watching a Kubernetes infrastructure.
	ScanTypeDiscoveryAgent = "DiscoveryAgent"
	// ScanTypeProbeExecution is a probe-execution scan. The public API accepts
	// it; the Terraform scan resource does not expose it yet.
	ScanTypeProbeExecution = "ProbeExecution"

	DetectionMethodAI           = "AI"
	DetectionMethodRiskAnalyzer = "RiskAnalyzer"

	ScanStatusQueued     = "QUEUED"
	ScanStatusCollecting = "COLLECTING"
	ScanStatusCollected  = "COLLECTED"
	ScanStatusAnalysing  = "ANALYSING"
	ScanStatusCompleted  = "COMPLETED"
	ScanStatusErrored    = "ERRORED"
	ScanStatusAborted    = "ABORTED"
)

// Scope identifies the Harness account, and optionally org and project, for a chaos API call.
type Scope struct {
	AccountID      string
	OrgID          string
	ProjectID      string
	ParentUniqueID string
	CorrelationID  string
}

// UserRef is a user on a scan response. The API returns either a user object
// or, when user lookup is skipped, a plain user id string.
type UserRef struct {
	UserID   string `json:"userID,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

// UnmarshalJSON accepts both {"userID":"..."} and "user-id".
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

// DisplayName prefers a human-readable user label.
func (u UserRef) DisplayName() string {
	if u.Username != "" {
		return u.Username
	}
	if u.Email != "" {
		return u.Email
	}
	return u.UserID
}

// Scan is a risk scan. Pipeline and infrastructure scans are the same object;
// ScanType plus Source selects which one it is.
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

	CollectedData *CollectionSummary `json:"collectedData,omitempty"`
	RiskSummary   *RiskSummary       `json:"riskSummary,omitempty"`
	Report        *ReportInfo        `json:"report,omitempty"`
	Summary       *HeatmapSummary    `json:"summary,omitempty"`
	ErrorMessage  string             `json:"errorMessage,omitempty"`
	RetryCount    int                `json:"retryCount,omitempty"`
	CreatedAt     int64              `json:"createdAt,omitempty"`
	UpdatedAt     int64              `json:"updatedAt,omitempty"`
	CompletedAt   int64              `json:"completedAt,omitempty"`
	CreatedBy     UserRef            `json:"createdBy,omitempty"`
	UpdatedBy     UserRef            `json:"updatedBy,omitempty"`
}

// ScanSource is the origin of a scan. Exactly one of Pipeline, Infrastructure,
// or ProbeExecution is set for a valid create request.
type ScanSource struct {
	Pipeline       *PipelineScan       `json:"pipeline,omitempty"`
	Infrastructure *InfrastructureScan `json:"discoveryAgent,omitempty"`
	ProbeExecution *ProbeExecutionScan `json:"probeExecution,omitempty"`
}

// PipelineScan identifies the pipeline execution being scanned.
type PipelineScan struct {
	PipelineIdentity    string                   `json:"pipelineIdentity"`
	PipelineExecutionID string                   `json:"pipelineExecutionID,omitempty"`
	NotifyID            string                   `json:"notifyID,omitempty"`
	Infrastructures     []PipelineInfrastructure `json:"infrastructures,omitempty"`
}

// PipelineInfrastructure is a CD environment and infrastructure pair observed on collected manifests.
type PipelineInfrastructure struct {
	EnvironmentRef         string `json:"environmentRef,omitempty"`
	InfrastructureIdentity string `json:"infrastructureIdentity,omitempty"`
	ChaosOnboarded         bool   `json:"chaosOnboarded"`
	ChaosSupported         bool   `json:"chaosSupported"`
}

// InfrastructureScan identifies the discovery agent whose Kubernetes infrastructure is scanned.
// The API field name is discoveryAgent; the product surface calls this an infrastructure scan.
type InfrastructureScan struct {
	AgentIdentity             string `json:"agentIdentity"`
	EnvironmentIdentity       string `json:"environmentIdentity"`
	EnvironmentParentUniqueID string `json:"environmentParentUniqueId,omitempty"`
}

// ProbeExecutionScan identifies a probe execution that triggered a scan.
type ProbeExecutionScan struct {
	ProbeIdentity   string                `json:"probeIdentity,omitempty"`
	ProbeName       string                `json:"probeName,omitempty"`
	ExecutionSource *ProbeExecutionSource `json:"executionSource,omitempty"`
}

// ProbeExecutionSource is a union: experiment or pipeline step.
type ProbeExecutionSource struct {
	Experiment   *ProbeExperimentSource   `json:"experiment,omitempty"`
	PipelineStep *ProbePipelineStepSource `json:"pipelineStep,omitempty"`
}

// ProbeExperimentSource identifies the experiment run that executed the probe.
type ProbeExperimentSource struct {
	ExperimentID    string `json:"experimentID,omitempty"`
	ExperimentRunID string `json:"experimentRunID,omitempty"`
}

// ProbePipelineStepSource identifies the pipeline step that executed the probe.
type ProbePipelineStepSource struct {
	PipelineID    string `json:"pipelineID,omitempty"`
	PipelineRunID string `json:"pipelineRunID,omitempty"`
}

// ScanCreateRequest is the body for POST /v3/risk-scans.
type ScanCreateRequest struct {
	Identity        string     `json:"identity"`
	Name            string     `json:"name"`
	Description     string     `json:"description,omitempty"`
	Tags            []string   `json:"tags,omitempty"`
	ScanType        string     `json:"scanType"`
	DetectionMethod string     `json:"detectionMethod,omitempty"`
	Source          ScanSource `json:"source"`
}

// ScanUpdateRequest updates user-facing metadata. Scan type and source are immutable.
// A nil Tags slice is omitted. A non-nil empty slice is sent as [] so tags can be cleared
// without a name-only update wiping them.
type ScanUpdateRequest struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Tags        []string `json:"-"`
}

// MarshalJSON omits a nil tag list and encodes an empty list as [].
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

// UnmarshalJSON keeps a missing tag list nil and an empty list empty.
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

// ScanListRequest filters GET /v3/risk-scans.
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

// ScanListResponse is the body of GET /v3/risk-scans.
type ScanListResponse struct {
	CorrelationID string     `json:"correlationID,omitempty"`
	Data          []Scan     `json:"data"`
	Pagination    Pagination `json:"pagination"`
}

// Pagination is the shared list page descriptor.
type Pagination struct {
	Index      int64 `json:"index,omitempty"`
	Limit      int64 `json:"limit,omitempty"`
	TotalItems int64 `json:"totalItems,omitempty"`
	TotalPages int64 `json:"totalPages,omitempty"`
}

// CollectionSummary describes resources gathered for a scan.
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

// CollectionError is a single collection failure.
type CollectionError struct {
	Source  string `json:"source,omitempty"`
	Message string `json:"message,omitempty"`
}

// RiskSummary is the aggregate produced when analysis finishes.
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

// ReportInfo tracks PDF report generation for a scan.
type ReportInfo struct {
	Status       string `json:"status,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	GeneratedAt  int64  `json:"generatedAt,omitempty"`
	ExpiresAt    int64  `json:"expiresAt,omitempty"`
}

// HeatmapSummary is the severity rollup returned with a scan.
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

// ScanReport is GET /v3/risk-scans/{identity}/report.
// Scan is the same document Get returns. ScannedRisks is the report embedding
// for this scan, not the durable scanned-risk object.
type ScanReport struct {
	Scan         *Scan         `json:"scan,omitempty"`
	ScannedRisks []ScanRisk    `json:"scannedRisks"`
	Summary      ReportSummary `json:"summary"`
}

// ReportSummary counts risks and occurrences on a scan report.
type ReportSummary struct {
	TotalRisks       int            `json:"totalRisks,omitempty"`
	TotalOccurrences int            `json:"totalOccurrences,omitempty"`
	BySeverity       map[string]int `json:"bySeverity,omitempty"`
	ByStatus         map[string]int `json:"byStatus,omitempty"`
}

// ScanRisk is one risk embedded in a scan report, including the occurrences
// found by that scan. Identity matches ScannedRisk.Identity so a later risk
// resource can address the same object.
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

// ScanRiskOccurrence is a single detection of a risk inside a scan.
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

// OccurrenceContext holds the evidence and guidance shown for an occurrence.
type OccurrenceContext struct {
	Evidence              []string `json:"evidence,omitempty"`
	Remediation           []string `json:"remediation,omitempty"`
	ResilienceTestingReco []string `json:"resilienceTestingRec,omitempty"`
}
