package scantest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	. "github.com/harness/harness-go-sdk/harness/chaos"
	"gopkg.in/yaml.v3"
)

// TestScanContractMatchesHCESaaSSwagger checks the hand-written scan client
// against the chaos-manager OpenAPI spec in hce-saas. It does not import
// hce-saas and it does not call a live account.
//
// The spec is found from HCE_SAAS_SWAGGER, HCE_SAAS_ROOT, or a sibling
// checkout at ../hce-saas relative to this module. CI without that checkout
// skips this test. The fixture test below always runs.
func TestScanContractMatchesHCESaaSSwagger(t *testing.T) {
	path := hceSwaggerPath()
	if path == "" {
		t.Skip("hce-saas swagger.yaml not found; set HCE_SAAS_SWAGGER or check out hce-saas next to harness-go-sdk")
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var doc swaggerDoc
	if err := yaml.Unmarshal(raw, &doc); err != nil {
		t.Fatalf("parse %s: %v", path, err)
	}

	assertRequiredCovered(t, doc, "riskscans.CreateRequest", ScanCreateRequest{})
	for _, model := range []struct {
		definition string
		value      interface{}
	}{
		{"riskscans.CreateRequest", ScanCreateRequest{}},
		{"riskscans.UpdateRequest", []string{"name", "description", "tags"}},
		{"riskscans.RiskScanResponse", Scan{}},
		{"risk.ScanSource", ScanSource{}},
		{"risk.ScanSourceCodeRepository", CodeRepositoryScan{}},
		{"risk.LinkedService", LinkedService{}},
		{"risk.CodeLocation", CodeLocation{}},
		{"riskscans.RecommendationCounts", RecommendationCounts{}},
		{"risk.ScanSourcePipeline", PipelineScan{}},
		{"risk.ScanSourceServiceDiscovery", InfrastructureScan{}},
		{"risk.ScanSourceProbeExecution", ProbeExecutionScan{}},
		{"risk.ProbeExecutionSource", ProbeExecutionSource{}},
		{"risk.ProbeExecutionExperimentSource", ProbeExperimentSource{}},
		{"risk.ProbeExecutionPipelineStepSource", ProbePipelineStepSource{}},
		{"risk.PipelineInfrastructure", PipelineInfrastructure{}},
		{"risk.CollectionSummary", CollectionSummary{}},
		{"risk.CollectionError", CollectionError{}},
		{"risk.RiskSummary", RiskSummary{}},
		{"risk.ReportInfo", ReportInfo{}},
		{"riskscans.HeatmapSummary", HeatmapSummary{}},
		{"riskscans.ListResponse", ScanListResponse{}},
		{"shared.PaginationResponse", Pagination{}},
		{"riskscans.ReportResponse", ScanReport{}},
		{"riskscans.ReportSummary", ReportSummary{}},
		{"riskscans.ScannedRiskWithOccurrences", ScanRisk{}},
		{"riskscans.OccurrenceDetails", ScanRiskOccurrence{}},
		{"riskscans.OccurrenceContext", OccurrenceContext{}},
		{"hcesdk.UserDetails", UserRef{}},
	} {
		assertJSONFields(t, doc, model.definition, model.value)
		assertSwaggerPropertiesPresent(t, doc, model.definition, model.value)
	}

	assertEnumCovers(t, doc, "risk.ScanType", ScanTypePipelineExecution, ScanTypeDiscoveryAgent, ScanTypeProbeExecution, ScanTypeCodeRepository)
	assertEnumCovers(t, doc, "risk.DetectionMethod", DetectionMethodAI, DetectionMethodRiskAnalyzer, DetectionMethodExternal)
	assertEnumCovers(t, doc, "risk.ScanStatus",
		ScanStatusQueued, ScanStatusCollecting, ScanStatusCollected, ScanStatusAnalysing,
		ScanStatusCompleted, ScanStatusErrored, ScanStatusAborted)
}

// TestScanResponseFixtures unmarshals chaos-manager JSON shapes that the
// swagger types allow, including the two createdBy encodings and false booleans.
func TestScanResponseFixtures(t *testing.T) {
	var scan Scan
	if err := json.Unmarshal([]byte(`{
		"identity":"scan-1",
		"scanType":"PipelineExecution",
		"detectionMethod":"RiskAnalyzer",
		"status":"COMPLETED",
		"source":{"pipeline":{"pipelineIdentity":"pipe","infrastructures":[{"environmentRef":"prod","chaosOnboarded":false,"chaosSupported":false}]}},
		"createdBy":{"userID":"u1","username":"ada"},
		"summary":{"score":0,"totalRisks":1},
		"riskSummary":{"severityCounts":{"High":1},"categoryCounts":{"Availability":1}}
	}`), &scan); err != nil {
		t.Fatal(err)
	}
	if scan.CreatedBy.DisplayName() != "ada" {
		t.Fatalf("createdBy = %+v", scan.CreatedBy)
	}
	infra := scan.Source.Pipeline.Infrastructures[0]
	if infra.ChaosOnboarded || infra.ChaosSupported || infra.EnvironmentRef != "prod" {
		t.Fatalf("infrastructure = %+v", infra)
	}
	if scan.Summary == nil || scan.Summary.Score != 0 || scan.Summary.TotalRisks != 1 {
		t.Fatalf("summary = %+v", scan.Summary)
	}
	if scan.RiskSummary == nil || scan.RiskSummary.SeverityCounts["High"] != 1 {
		t.Fatalf("risk summary = %+v", scan.RiskSummary)
	}

	var plain Scan
	if err := json.Unmarshal([]byte(`{"identity":"scan-2","createdBy":"user-9"}`), &plain); err != nil {
		t.Fatal(err)
	}
	if plain.CreatedBy.UserID != "user-9" || plain.CreatedBy.Username != "" {
		t.Fatalf("string createdBy = %+v", plain.CreatedBy)
	}

	var report ScanReport
	if err := json.Unmarshal([]byte(`{"scannedRisks":[],"summary":{"totalRisks":0},"scan":{"identity":"scan-1"}}`), &report); err != nil {
		t.Fatal(err)
	}
	if report.ScannedRisks == nil || len(report.ScannedRisks) != 0 {
		t.Fatalf("empty report = %+v", report)
	}
}

type swaggerDoc struct {
	Definitions map[string]swaggerDefinition `yaml:"definitions"`
}

type swaggerDefinition struct {
	Required   []string                   `yaml:"required"`
	Properties map[string]swaggerProperty `yaml:"properties"`
	Enum       []string                   `yaml:"enum"`
}

type swaggerProperty struct {
	Ref string `yaml:"$ref"`
}

func hceSwaggerPath() string {
	if path := os.Getenv("HCE_SAAS_SWAGGER"); path != "" {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	if root := os.Getenv("HCE_SAAS_ROOT"); root != "" {
		path := filepath.Join(root, "graphql", "server", "docs", "swagger.yaml")
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	for i := 0; i < 8; i++ {
		path := filepath.Join(dir, "hce-saas", "graphql", "server", "docs", "swagger.yaml")
		if _, err := os.Stat(path); err == nil {
			return path
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return ""
}

func assertJSONFields(t *testing.T, doc swaggerDoc, definition string, value interface{}) {
	t.Helper()
	def, ok := doc.Definitions[definition]
	if !ok {
		t.Errorf("swagger definition %s is missing", definition)
		return
	}
	for _, name := range contractNames(value) {
		if _, ok := def.Properties[name]; !ok {
			t.Errorf("SDK field %q is not a property of swagger %s", name, definition)
		}
	}
}

func assertSwaggerPropertiesPresent(t *testing.T, doc swaggerDoc, definition string, value interface{}) {
	t.Helper()
	def, ok := doc.Definitions[definition]
	if !ok {
		t.Errorf("swagger definition %s is missing", definition)
		return
	}
	have := map[string]struct{}{}
	for _, name := range contractNames(value) {
		have[name] = struct{}{}
	}
	for name := range def.Properties {
		if _, ok := have[name]; !ok {
			t.Errorf("swagger %s property %q is missing from the SDK type", definition, name)
		}
	}
}

func assertRequiredCovered(t *testing.T, doc swaggerDoc, definition string, value interface{}) {
	t.Helper()
	def, ok := doc.Definitions[definition]
	if !ok {
		t.Errorf("swagger definition %s is missing", definition)
		return
	}
	have := map[string]struct{}{}
	for _, name := range jsonFieldNames(value) {
		have[name] = struct{}{}
	}
	for _, name := range def.Required {
		if _, ok := have[name]; !ok {
			t.Errorf("swagger %s requires %q and the SDK create body cannot send it", definition, name)
		}
	}
}

func assertEnumCovers(t *testing.T, doc swaggerDoc, definition string, values ...string) {
	t.Helper()
	def, ok := doc.Definitions[definition]
	if !ok {
		t.Errorf("swagger definition %s is missing", definition)
		return
	}
	have := map[string]struct{}{}
	for _, value := range def.Enum {
		have[value] = struct{}{}
	}
	for _, value := range values {
		if _, ok := have[value]; !ok {
			t.Errorf("SDK value %q is not in swagger %s enum %v", value, definition, def.Enum)
		}
	}
}

func contractNames(value interface{}) []string {
	switch typed := value.(type) {
	case string:
		return []string{typed}
	case []string:
		return typed
	default:
		return jsonFieldNames(value)
	}
}

func jsonFieldNames(value interface{}) []string {
	fieldType := reflect.TypeOf(value)
	if fieldType.Kind() == reflect.Ptr {
		fieldType = fieldType.Elem()
	}
	names := make([]string, 0, fieldType.NumField())
	for i := 0; i < fieldType.NumField(); i++ {
		tag := fieldType.Field(i).Tag.Get("json")
		name := strings.Split(tag, ",")[0]
		if name == "" || name == "-" {
			continue
		}
		names = append(names, name)
	}
	return names
}
