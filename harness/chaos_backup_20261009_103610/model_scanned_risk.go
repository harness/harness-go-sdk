package chaos

// Scanned risk status values returned by GET /v3/scanned-risks.
const (
	ScannedRiskStatusOpen       = "Open"
	ScannedRiskStatusInProgress = "InProgress"
	ScannedRiskStatusResolved   = "Resolved"
	ScannedRiskStatusIgnored    = "Ignored"
)

// ScannedRisk is a durable risk produced by one or more scans.
//
// Read operations exist today (Get, List, ListOccurrences). Create, Update, and
// Delete are not implemented: the public API does not expose them yet. This
// type is the payload those methods should accept and return, keyed by Identity,
// so adding them does not require a new schema.
type ScannedRisk struct {
	AccountID      string          `json:"accountID,omitempty"`
	OrgID          string          `json:"orgID,omitempty"`
	ProjectID      string          `json:"projectID,omitempty"`
	ParentUniqueID string          `json:"parentUniqueId,omitempty"`
	UniqueID       string          `json:"uniqueId,omitempty"`
	Identity       string          `json:"identity,omitempty"`
	Name           string          `json:"name,omitempty"`
	Description    string          `json:"description,omitempty"`
	Tags           []string        `json:"tags,omitempty"`
	Spec           ScannedRiskSpec `json:"spec"`
}

// ScannedRiskSpec is the mutable and descriptive body of a scanned risk.
// Status, name, description, and tags are the fields a future update is
// expected to change; the rest identify how and where the risk was detected.
type ScannedRiskSpec struct {
	RiskRuleID      string        `json:"riskRuleId,omitempty"`
	RiskRuleName    string        `json:"riskRuleName,omitempty"`
	Impact          string        `json:"impact,omitempty"`
	TargetService   TargetService `json:"targetService"`
	Fingerprint     string        `json:"fingerprint,omitempty"`
	Severity        string        `json:"severity,omitempty"`
	SeverityOrder   int           `json:"severityOrder,omitempty"`
	Status          string        `json:"status,omitempty"`
	OccurrenceCount int64         `json:"occurrenceCount,omitempty"`
	FirstDetected   int64         `json:"firstDetected,omitempty"`
	LastDetected    int64         `json:"lastDetected,omitempty"`
	ValidationType  string        `json:"validationType,omitempty"`
	ValidationTypes []string      `json:"validationTypes,omitempty"`
	Category        string        `json:"category,omitempty"`
	Sources         []string      `json:"sources,omitempty"`
}

// TargetService is the service a scanned risk applies to.
type TargetService struct {
	Type              string         `json:"type,omitempty"`
	Name              string         `json:"name,omitempty"`
	Identity          string         `json:"identity,omitempty"`
	Namespace         string         `json:"namespace,omitempty"`
	ResourceKind      string         `json:"resourceKind,omitempty"`
	ResourceName      string         `json:"resourceName,omitempty"`
	ExternalServiceID string         `json:"externalServiceID,omitempty"`
	Details           ServiceDetails `json:"details"`
}

// ServiceDetails carries the kind-specific service reference.
type ServiceDetails struct {
	DiscoveredService *DiscoveredServiceRef `json:"discoveredService,omitempty"`
	ChaosService      *ChaosServiceRef      `json:"chaosService,omitempty"`
	CDService         *CDServiceRef         `json:"cdService,omitempty"`
}

// DiscoveredServiceRef points at a service found by a discovery agent.
type DiscoveredServiceRef struct {
	Identity            string `json:"identity,omitempty"`
	UniqueID            string `json:"uniqueID,omitempty"`
	ParentUniqueID      string `json:"parentUniqueID,omitempty"`
	EnvironmentIdentity string `json:"environmentIdentity,omitempty"`
	AgentIdentity       string `json:"agentIdentity,omitempty"`
	ExternalID          string `json:"externalID,omitempty"`
}

// ChaosServiceRef points at a chaos service.
type ChaosServiceRef struct {
	Name     string `json:"name,omitempty"`
	Identity string `json:"identity,omitempty"`
}

// CDServiceRef points at a CD service deployed to an environment and infrastructure.
type CDServiceRef struct {
	Identity               string `json:"identity,omitempty"`
	UniqueID               string `json:"uniqueID,omitempty"`
	ParentUniqueID         string `json:"parentUniqueID,omitempty"`
	EnvironmentIdentity    string `json:"environmentIdentity,omitempty"`
	InfrastructureIdentity string `json:"infrastructureIdentity,omitempty"`
	ServiceIdentity        string `json:"serviceIdentity,omitempty"`
}

// ScannedRiskGetResponse is the body of GET /v3/scanned-risks/{identity}.
type ScannedRiskGetResponse struct {
	ScannedRisk ScannedRisk `json:"scannedRisk"`
}

// ScannedRiskListRequest filters GET /v3/scanned-risks.
type ScannedRiskListRequest struct {
	Scope
	Page                int
	Limit               int
	Search              string
	SortField           string
	SortAscending       *bool
	Severity            string
	Source              string
	ValidationType      string
	ServiceIdentity     string
	AgentIdentity       string
	EnvironmentIdentity string
	RiskRuleID          string
	Tags                string
	StartTime           int64
	EndTime             int64
	IncludeAllScope     *bool
	ServiceFilter       ScannedRiskServiceFilter
}

// ScannedRiskServiceFilter is the serviceFilterV2* query on GET /v3/scanned-risks.
// Kind is DiscoveredService, ChaosService, or CDService. The other fields apply
// only for the matching kind, and only when set.
type ScannedRiskServiceFilter struct {
	Kind                           string
	DiscoveredServiceAgentID       string
	DiscoveredServiceEnvironmentID string
	DiscoveredServiceExternalID    string
	ChaosServiceIdentity           string
	CDServiceIdentity              string
	CDServiceEnvironmentID         string
	CDServiceInfrastructureID      string
}

// ScannedRiskOccurrencesRequest filters GET /v3/scanned-risks/{identity}/occurrences.
type ScannedRiskOccurrencesRequest struct {
	Scope
	Page           int
	Limit          int
	Search         string
	SortField      string
	SortAscending  *bool
	Source         string
	ScanType       string
	ValidationType string
	StartTime      int64
	EndTime        int64
}

// ScannedRiskListResponse is the body of GET /v3/scanned-risks.
type ScannedRiskListResponse struct {
	CorrelationID string        `json:"correlationID,omitempty"`
	Data          []ScannedRisk `json:"data"`
	Pagination    Pagination    `json:"pagination"`
}

// ScannedRiskOccurrence is one detection returned by the occurrences API.
// It is separate from ScanRiskOccurrence, which is the scan-report embedding.
type ScannedRiskOccurrence struct {
	DetectionRunID       string                `json:"detectionRunId,omitempty"`
	ScanID               string                `json:"scanId,omitempty"`
	ScanName             string                `json:"scanName,omitempty"`
	ScanType             string                `json:"scanType,omitempty"`
	DetectionMethod      string                `json:"detectionMethod,omitempty"`
	Source               string                `json:"source,omitempty"`
	ValidationType       string                `json:"validationType,omitempty"`
	Evidence             []string              `json:"evidence,omitempty"`
	Remediation          []string              `json:"remediation,omitempty"`
	ResilienceTestingRec []string              `json:"resilienceTestingRec,omitempty"`
	TimeDetected         int64                 `json:"timeDetected,omitempty"`
	EvidenceResource     *EvidenceResource     `json:"evidenceResource,omitempty"`
	ProbeExecutionSource *ProbeExecutionSource `json:"probeExecutionSource,omitempty"`
}

// EvidenceResource is the manifest that surfaced a risk occurrence.
type EvidenceResource struct {
	Kind      string `json:"kind,omitempty"`
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Role      string `json:"role,omitempty"`
}

// ScannedRiskOccurrencesResponse is the body of GET /v3/scanned-risks/{identity}/occurrences.
type ScannedRiskOccurrencesResponse struct {
	CorrelationID string                  `json:"correlationID,omitempty"`
	Data          []ScannedRiskOccurrence `json:"data"`
	Pagination    Pagination              `json:"pagination"`
}
