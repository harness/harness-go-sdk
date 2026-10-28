package scantest

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/harness/harness-go-sdk/harness/chaos"
)

func TestUserRefUnmarshal(t *testing.T) {
	var object UserRef
	if err := json.Unmarshal([]byte(`{"userID":"u1","username":"ada","email":"ada@example.com"}`), &object); err != nil {
		t.Fatal(err)
	}
	if object.DisplayName() != "ada" {
		t.Fatalf("display name = %q", object.DisplayName())
	}

	var plain UserRef
	if err := json.Unmarshal([]byte(`"user-123"`), &plain); err != nil {
		t.Fatal(err)
	}
	if plain.UserID != "user-123" || plain.DisplayName() != "user-123" {
		t.Fatalf("plain user ref = %+v", plain)
	}

	// A reused value must not keep fields from the previous shape.
	if err := json.Unmarshal([]byte(`"user-123"`), &object); err != nil {
		t.Fatal(err)
	}
	if object.UserID != "user-123" || object.Username != "" || object.Email != "" {
		t.Fatalf("reused user ref = %+v", object)
	}
}

func TestScanUpdateRequestOmitsNilTags(t *testing.T) {
	name := "Renamed"
	omitted, err := json.Marshal(ScanUpdateRequest{Name: &name})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(omitted), "tags") {
		t.Fatalf("nil tags were sent: %s", omitted)
	}

	cleared, err := json.Marshal(ScanUpdateRequest{Name: &name, Tags: []string{}})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(cleared), `"tags":[]`) {
		t.Fatalf("empty tags = %s", cleared)
	}

	var decoded ScanUpdateRequest
	if err := json.Unmarshal(cleared, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.Tags == nil || len(decoded.Tags) != 0 || decoded.Name == nil || *decoded.Name != name {
		t.Fatalf("decoded update = %+v", decoded)
	}
}

func TestScanClientPipelineAndInfrastructure(t *testing.T) {
	var gotMethod, gotPath, gotQuery string
	var gotBody map[string]interface{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		gotQuery = r.URL.RawQuery
		body, _ := io.ReadAll(r.Body)
		gotBody = nil
		if len(body) > 0 {
			if err := json.Unmarshal(body, &gotBody); err != nil {
				t.Errorf("decode body: %v", err)
			}
		}
		if r.Header.Get("X-API-Key") != "test-key" {
			t.Errorf("api key = %q", r.Header.Get("X-API-Key"))
		}

		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/v3/risk-scans":
			_, _ = w.Write([]byte(`{"identity":"scan-1","name":"Scan 1","scanType":"PipelineExecution","status":"QUEUED","createdBy":"user-1"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/v3/risk-scans/scan-1":
			_, _ = w.Write([]byte(`{"identity":"scan-1","name":"Scan 1","scanType":"DiscoveryAgent","status":"COMPLETED","source":{"discoveryAgent":{"agentIdentity":"agent-1","environmentIdentity":"env-1"}},"createdBy":{"userID":"user-1","username":"ada"},"summary":{"totalRisks":2,"critical":1}}`))
		case r.Method == http.MethodGet && r.URL.Path == "/v3/risk-scans/scan-1/report":
			_, _ = w.Write([]byte(`{"scannedRisks":[{"identity":"risk-1","name":"No replicas","severity":"High","status":"Open","serviceIdentity":"svc-1","occurrences":[{"occurrenceId":"occ-1","detectedAt":10,"context":{"evidence":["one"]}}]}],"summary":{"totalRisks":1,"totalOccurrences":1}}`))
		case r.Method == http.MethodPut && r.URL.Path == "/v3/risk-scans/scan-1":
			_, _ = w.Write([]byte(`{"status":"ok"}`))
		case r.Method == http.MethodDelete && r.URL.Path == "/v3/risk-scans/scan-1":
			_, _ = w.Write([]byte(`{"status":"ok"}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	client := NewAPIClient(&Configuration{
		BasePath:   server.URL,
		ApiKey:     "test-key",
		AccountId:  "acc",
		HTTPClient: server.Client(),
	})
	ctx := context.WithValue(context.Background(), ContextAPIKey, APIKey{Key: "test-key"})
	scans := NewScanClient(client)
	scope := Scope{AccountID: "acc", OrgID: "org", ProjectID: "proj"}

	created, _, err := scans.Create(ctx, scope, ScanCreateRequest{
		Identity:        "scan-1",
		Name:            "Scan 1",
		ScanType:        ScanTypePipelineExecution,
		DetectionMethod: DetectionMethodRiskAnalyzer,
		Source: ScanSource{Pipeline: &PipelineScan{
			PipelineIdentity:    "pipe-1",
			PipelineExecutionID: "exec-1",
		}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if created.Identity != "scan-1" || created.CreatedBy.UserID != "user-1" {
		t.Fatalf("created = %+v", created)
	}
	if gotMethod != http.MethodPost || gotPath != "/v3/risk-scans" {
		t.Fatalf("create request %s %s", gotMethod, gotPath)
	}
	if gotQuery != "accountIdentifier=acc&organizationIdentifier=org&projectIdentifier=proj" {
		t.Fatalf("query = %s", gotQuery)
	}
	source := gotBody["source"].(map[string]interface{})
	pipeline := source["pipeline"].(map[string]interface{})
	if pipeline["pipelineIdentity"] != "pipe-1" || pipeline["pipelineExecutionID"] != "exec-1" {
		t.Fatalf("pipeline source = %#v", source)
	}
	if _, ok := source["discoveryAgent"]; ok {
		t.Fatalf("pipeline create included discovery agent: %#v", source)
	}

	got, _, err := scans.Get(ctx, scope, "scan-1")
	if err != nil {
		t.Fatal(err)
	}
	if got.ScanType != ScanTypeDiscoveryAgent || got.Source.Infrastructure == nil || got.Source.Infrastructure.AgentIdentity != "agent-1" {
		t.Fatalf("get = %+v", got)
	}
	if got.CreatedBy.DisplayName() != "ada" || got.Summary == nil || got.Summary.TotalRisks != 2 {
		t.Fatalf("get metadata = %+v summary=%+v", got.CreatedBy, got.Summary)
	}

	report, _, err := scans.GetReport(ctx, scope, "scan-1")
	if err != nil {
		t.Fatal(err)
	}
	if len(report.ScannedRisks) != 1 || report.ScannedRisks[0].Identity != "risk-1" {
		t.Fatalf("report = %+v", report)
	}
	if len(report.ScannedRisks[0].Occurrences) != 1 || report.ScannedRisks[0].Occurrences[0].Context.Evidence[0] != "one" {
		t.Fatalf("occurrences = %+v", report.ScannedRisks[0].Occurrences)
	}

	name := "Renamed"
	if _, err := scans.Update(ctx, scope, "scan-1", ScanUpdateRequest{Name: &name, Tags: []string{"a"}}); err != nil {
		t.Fatal(err)
	}
	if gotBody["name"] != "Renamed" {
		t.Fatalf("update body = %#v", gotBody)
	}
	if _, ok := gotBody["scanType"]; ok {
		t.Fatalf("update must not send scan type: %#v", gotBody)
	}

	if _, err := scans.Delete(ctx, scope, "scan-1"); err != nil {
		t.Fatal(err)
	}
	if gotMethod != http.MethodDelete {
		t.Fatalf("delete method = %s", gotMethod)
	}
}

func TestScanClientDecodesJSONWithoutContentType(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"identity":"scan-1","name":"Scan","scanType":"PipelineExecution","status":"QUEUED"}`))
	}))
	defer server.Close()

	client := NewAPIClient(&Configuration{BasePath: server.URL, HTTPClient: server.Client()})
	got, _, err := NewScanClient(client).Get(context.Background(), Scope{AccountID: "acc"}, "scan-1")
	if err != nil {
		t.Fatal(err)
	}
	if got.Identity != "scan-1" || got.Status != ScanStatusQueued {
		t.Fatalf("scan = %+v", got)
	}
}

func TestScanClientRejectsNilAPI(t *testing.T) {
	if _, _, err := NewScanClient(nil).Get(context.Background(), Scope{}, "scan-1"); err == nil {
		t.Fatal("expected error")
	}
}

func TestScanClientCreateError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"message":"pipeline identity is required","correlationID":"c1"}`))
	}))
	defer server.Close()

	client := NewAPIClient(&Configuration{BasePath: server.URL, HTTPClient: server.Client()})
	_, resp, err := NewScanClient(client).Create(context.Background(), Scope{AccountID: "acc"}, ScanCreateRequest{})
	if err == nil {
		t.Fatal("expected error")
	}
	if resp == nil || resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %v", resp)
	}
	if !strings.Contains(err.Error(), "pipeline identity is required") {
		t.Fatalf("error = %q", err)
	}
}

func TestAPIDetectionMethod(t *testing.T) {
	cases := map[string]string{
		"ai":           DetectionMethodAI,
		"AI":           DetectionMethodAI,
		"ruleBased":    DetectionMethodRiskAnalyzer,
		"RiskAnalyzer": DetectionMethodRiskAnalyzer,
	}
	for in, want := range cases {
		got, err := APIDetectionMethod(in)
		if err != nil || got != want {
			t.Fatalf("APIDetectionMethod(%q) = %q, %v", in, got, err)
		}
	}
	if got, err := APIDetectionMethod(""); err != nil || got != "" {
		t.Fatalf("empty method = %q, %v", got, err)
	}
	if _, err := APIDetectionMethod("rule"); err == nil || !strings.Contains(err.Error(), "ai or ruleBased") {
		t.Fatalf("error = %v", err)
	}
}

func TestCreatePipelineScanSendsRuleBasedAsRiskAnalyzer(t *testing.T) {
	var gotBody map[string]interface{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		if err := json.Unmarshal(body, &gotBody); err != nil {
			t.Errorf("decode body: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"identity":"scan-1","scanType":"PipelineExecution","status":"QUEUED"}`))
	}))
	defer server.Close()

	client := NewAPIClient(&Configuration{BasePath: server.URL, HTTPClient: server.Client()})
	_, _, err := NewScanClient(client).CreatePipelineScan(context.Background(), Scope{AccountID: "acc"}, PipelineScanRequest{
		Identity:            "scan-1",
		Name:                "Cart",
		DetectionMethod:     "ruleBased",
		PipelineIdentity:    "cart_deploy",
		PipelineExecutionID: "exec-1",
	})
	if err != nil {
		t.Fatal(err)
	}
	if gotBody["detectionMethod"] != DetectionMethodRiskAnalyzer || gotBody["scanType"] != ScanTypePipelineExecution {
		t.Fatalf("body = %#v", gotBody)
	}
}

func TestCreateInfrastructureScanReadsLinkedAgent(t *testing.T) {
	var gotBody map[string]interface{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/chaos/manager/api/rest/v2/infrastructure/payments-infra":
			_, _ = w.Write([]byte(`{"identity":"payments-infra","name":"payments","discoveryAgentID":"payments-agent"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/servicediscovery/api/v1/agents/payments-agent":
			if r.URL.Query().Get("environmentIdentifier") != "prod" {
				t.Errorf("agent query = %s", r.URL.RawQuery)
			}
			_, _ = w.Write([]byte(`{"identity":"payments-agent","name":"payments","environmentIdentifier":"prod","parentUniqueId":"parent-1"}`))
		case r.Method == http.MethodPost && r.URL.Path == "/chaos/manager/api/v3/risk-scans":
			body, _ := io.ReadAll(r.Body)
			if err := json.Unmarshal(body, &gotBody); err != nil {
				t.Errorf("decode body: %v", err)
			}
			_, _ = w.Write([]byte(`{"identity":"scan-1","scanType":"DiscoveryAgent","status":"QUEUED","detectionMethod":"AI"}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	client := NewAPIClient(&Configuration{BasePath: server.URL + "/chaos/manager/api", HTTPClient: server.Client()})
	created, _, err := NewScanClient(client).CreateInfrastructureScan(context.Background(), Scope{AccountID: "acc", OrgID: "org", ProjectID: "proj"}, InfrastructureScanRequest{
		Identity:         "scan-1",
		Name:             "Infra",
		DetectionMethod:  "ai",
		InfrastructureID: "payments-infra",
		EnvironmentID:    "prod",
	})
	if err != nil {
		t.Fatal(err)
	}
	if created.ScanType != ScanTypeDiscoveryAgent {
		t.Fatalf("scan = %+v", created)
	}
	source := gotBody["source"].(map[string]interface{})
	agent := source["discoveryAgent"].(map[string]interface{})
	if agent["agentIdentity"] != "payments-agent" || agent["environmentIdentity"] != "prod" || agent["environmentParentUniqueId"] != "parent-1" {
		t.Fatalf("source = %#v", source)
	}
	if gotBody["detectionMethod"] != DetectionMethodAI {
		t.Fatalf("detection method = %#v", gotBody["detectionMethod"])
	}
}

func TestCreateInfrastructureScanUsesOnboardingAgentIdentity(t *testing.T) {
	var gotBody map[string]interface{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/chaos/manager/api/rest/v2/infrastructure/payments-infra":
			_, _ = w.Write([]byte(`{"identity":"payments-infra","name":"payments","environmentID":"prod"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/servicediscovery/api/v1/agents/payments-infra":
			if r.URL.Query().Get("environmentIdentifier") != "prod" {
				t.Errorf("agent query = %s", r.URL.RawQuery)
			}
			_, _ = w.Write([]byte(`{"identity":"payments-infra","name":"Discovery payments-infra","environmentIdentifier":"prod","parentUniqueId":"parent-1"}`))
		case r.Method == http.MethodPost && r.URL.Path == "/chaos/manager/api/v3/risk-scans":
			body, _ := io.ReadAll(r.Body)
			if err := json.Unmarshal(body, &gotBody); err != nil {
				t.Errorf("decode body: %v", err)
			}
			_, _ = w.Write([]byte(`{"identity":"scan-1","scanType":"DiscoveryAgent","status":"QUEUED"}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	client := NewAPIClient(&Configuration{BasePath: server.URL + "/chaos/manager/api", HTTPClient: server.Client()})
	_, _, err := NewScanClient(client).CreateInfrastructureScan(context.Background(), Scope{AccountID: "acc", OrgID: "org", ProjectID: "proj"}, InfrastructureScanRequest{
		Identity:         "scan-1",
		Name:             "Infra",
		InfrastructureID: "payments-infra",
		EnvironmentID:    "prod",
	})
	if err != nil {
		t.Fatal(err)
	}
	source := gotBody["source"].(map[string]interface{})
	agent := source["discoveryAgent"].(map[string]interface{})
	if agent["agentIdentity"] != "payments-infra" || agent["environmentIdentity"] != "prod" || agent["environmentParentUniqueId"] != "parent-1" {
		t.Fatalf("source = %#v", source)
	}
}

func TestCreateInfrastructureScanRejectsMissingDiscoveryAgent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/chaos/manager/api/rest/v2/infrastructure/payments-infra" {
			_, _ = w.Write([]byte(`{"identity":"payments-infra","name":"payments","discoveryAgentID":""}`))
			return
		}
		http.NotFound(w, r)
	}))
	defer server.Close()

	client := NewAPIClient(&Configuration{BasePath: server.URL + "/chaos/manager/api", HTTPClient: server.Client()})
	_, _, err := NewScanClient(client).CreateInfrastructureScan(context.Background(), Scope{AccountID: "acc", OrgID: "org", ProjectID: "proj"}, InfrastructureScanRequest{
		Identity:         "scan-1",
		Name:             "Infra",
		InfrastructureID: "payments-infra",
		EnvironmentID:    "prod",
	})
	if err == nil || !strings.Contains(err.Error(), "no discovery agent") {
		t.Fatalf("error = %v", err)
	}
}

func TestCreateInfrastructureScanRequiresEnvironment(t *testing.T) {
	client := NewAPIClient(&Configuration{BasePath: "https://app.harness.io/gateway/chaos/manager/api", HTTPClient: http.DefaultClient})
	_, _, err := NewScanClient(client).CreateInfrastructureScan(context.Background(), Scope{AccountID: "acc"}, InfrastructureScanRequest{
		Identity:         "scan-1",
		Name:             "Payments",
		InfrastructureID: "payments",
	})
	if err == nil || !strings.Contains(err.Error(), "environment id") {
		t.Fatalf("error = %v", err)
	}
}
