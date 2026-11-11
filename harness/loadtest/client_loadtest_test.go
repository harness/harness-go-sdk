package loadtest

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestClientUsesConfiguredBasePath(t *testing.T) {
	var gotPath string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"identity":"checkout","name":"Checkout","toolType":"K6"}`))
	}))
	defer server.Close()

	client := NewLoadTestClient(NewAPIClient(&Configuration{BasePath: server.URL + "/loadTest/manager/api", HTTPClient: server.Client()}))
	_, _, err := client.Get(t.Context(), Scope{AccountID: "acc"}, "checkout")
	if err != nil {
		t.Fatal(err)
	}
	if gotPath != "/loadTest/manager/api/v1/load-tests/checkout" {
		t.Fatalf("path = %s", gotPath)
	}
}

func TestCreateLoadTestRewritesBaseAndSendsScope(t *testing.T) {
	var gotPath, gotQuery string
	var gotBody map[string]interface{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		gotPath = r.URL.Path
		gotQuery = r.URL.RawQuery
		body, _ := io.ReadAll(r.Body)
		if err := json.Unmarshal(body, &gotBody); err != nil {
			t.Errorf("decode body: %v", err)
		}
		if r.Header.Get("X-API-Key") != "test-key" {
			t.Errorf("api key = %q", r.Header.Get("X-API-Key"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"identity":"checkout","name":"Checkout","toolType":"K6","uniqueId":"uid-1"}`))
	}))
	defer server.Close()

	client := NewAPIClient(&Configuration{
		BasePath:   server.URL + "/loadTest/manager/api",
		ApiKey:     "test-key",
		HTTPClient: server.Client(),
	})
	_, ctx := client.WithAuthContext(t.Context())
	created, _, err := NewLoadTestClient(client).Create(ctx, Scope{
		AccountID: "acc",
		OrgID:     "default",
		ProjectID: "demo",
	}, LoadTestCreateRequest{
		Identity:              "checkout",
		Name:                  "Checkout",
		ToolType:              LoadTestToolK6,
		EnvironmentIdentifier: "prod",
		InfraIdentifier:       "payments",
		ToolConfig: map[string]interface{}{
			"mode":   "script",
			"script": map[string]interface{}{"content": "a2s="},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if gotPath != "/loadTest/manager/api/v1/load-tests" {
		t.Fatalf("path = %s", gotPath)
	}
	if gotQuery != "accountIdentifier=acc&organizationIdentifier=default&projectIdentifier=demo" {
		t.Fatalf("query = %s", gotQuery)
	}
	toolConfig, _ := gotBody["toolConfig"].(map[string]interface{})
	if toolConfig["mode"] != "script" {
		t.Fatalf("k6 toolConfig = %#v", gotBody["toolConfig"])
	}
	if _, nested := toolConfig["k6"]; nested {
		t.Fatalf("k6 toolConfig must be flat, got %#v", toolConfig)
	}
	if created.UniqueID != "uid-1" || created.Identity != "checkout" {
		t.Fatalf("created = %+v", created)
	}
}

func TestCreateLocustNestsToolConfig(t *testing.T) {
	var gotBody map[string]interface{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(body, &gotBody)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"identity":"browse","toolType":"Locust"}`))
	}))
	defer server.Close()

	client := NewAPIClient(&Configuration{BasePath: server.URL + "/loadTest/manager/api", HTTPClient: server.Client()})
	_, _, err := NewLoadTestClient(client).Create(t.Context(), Scope{AccountID: "acc"}, LoadTestCreateRequest{
		Identity:              "browse",
		Name:                  "Browse",
		ToolType:              LoadTestToolLocust,
		EnvironmentIdentifier: "prod",
		InfraIdentifier:       "linux-1",
		TargetType:            LoadTestTargetLinux,
		ToolConfig: map[string]interface{}{
			"locust": map[string]interface{}{"mode": "script", "script": map[string]interface{}{"content": "cHk="}},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	toolConfig, _ := gotBody["toolConfig"].(map[string]interface{})
	locust, _ := toolConfig["locust"].(map[string]interface{})
	if locust["mode"] != "script" {
		t.Fatalf("locust toolConfig = %#v", gotBody["toolConfig"])
	}
}

func TestCreateRejectsJMeter(t *testing.T) {
	client := NewLoadTestClient(NewAPIClient(&Configuration{BasePath: "http://unused/loadTest/manager/api"}))
	_, _, err := client.Create(t.Context(), Scope{AccountID: "acc"}, LoadTestCreateRequest{
		Identity:              "checkout",
		Name:                  "Checkout",
		ToolType:              LoadTestToolJMeter,
		TargetType:            LoadTestTargetKubernetes,
		EnvironmentIdentifier: "prod",
		InfraIdentifier:       "payments",
	})
	if err == nil || !strings.Contains(err.Error(), "JMeter") {
		t.Fatal(err)
	}
}

func TestGetReadsJMeter(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"identity":"checkout","name":"Checkout","toolType":"JMeter","toolConfig":{"jmeter":{"mode":"script"}}}`))
	}))
	defer server.Close()

	client := NewAPIClient(&Configuration{BasePath: server.URL + "/loadTest/manager/api", HTTPClient: server.Client()})
	got, _, err := NewLoadTestClient(client).Get(t.Context(), Scope{AccountID: "acc"}, "checkout")
	if err != nil {
		t.Fatal(err)
	}
	if got.ToolType != LoadTestToolJMeter {
		t.Fatalf("toolType = %s", got.ToolType)
	}
	jmeter, _ := got.ToolConfig["jmeter"].(map[string]interface{})
	if jmeter["mode"] != "script" {
		t.Fatalf("toolConfig = %#v", got.ToolConfig)
	}
}

func TestCreateManagedInfraRejectsIdentifiers(t *testing.T) {
	client := NewLoadTestClient(NewAPIClient(&Configuration{BasePath: "http://unused/loadTest/manager/api"}))
	_, _, err := client.Create(t.Context(), Scope{AccountID: "acc"}, LoadTestCreateRequest{
		Identity:        "managed",
		Name:            "Managed",
		ToolType:        LoadTestToolK6,
		ManagedInfra:    true,
		InfraIdentifier: "should-not-be-set",
	})
	if err == nil {
		t.Fatal("expected infra identifiers to be rejected with managed infra")
	}
}

func TestUpdateLeavesUnsetDurationAndManagedInfra(t *testing.T) {
	var gotBody map[string]interface{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(body, &gotBody)
		_, _ = w.Write([]byte(`{"identity":"checkout","name":"Checkout"}`))
	}))
	defer server.Close()

	client := NewLoadTestClient(NewAPIClient(&Configuration{
		BasePath:   server.URL + "/loadTest/manager/api",
		HTTPClient: server.Client(),
	}))
	_, _, err := client.Update(t.Context(), Scope{AccountID: "acc"}, "checkout", LoadTestUpdateRequest{
		Name: "Checkout",
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := gotBody["maxDurationSec"]; !ok || gotBody["maxDurationSec"] != nil {
		t.Fatalf("maxDurationSec = %#v", gotBody["maxDurationSec"])
	}
	if _, ok := gotBody["managedInfra"]; !ok || gotBody["managedInfra"] != nil {
		t.Fatalf("managedInfra = %#v", gotBody["managedInfra"])
	}

	duration := 0
	managed := false
	_, _, err = client.Update(t.Context(), Scope{AccountID: "acc"}, "checkout", LoadTestUpdateRequest{
		Name:           "Checkout",
		MaxDurationSec: &duration,
		ManagedInfra:   &managed,
	})
	if err != nil {
		t.Fatal(err)
	}
	if gotBody["maxDurationSec"] != float64(0) {
		t.Fatalf("maxDurationSec = %#v", gotBody["maxDurationSec"])
	}
	if gotBody["managedInfra"] != false {
		t.Fatalf("managedInfra = %#v", gotBody["managedInfra"])
	}
}

func TestIdentityIsEscapedInThePath(t *testing.T) {
	var gotRequest string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		gotRequest = r.RequestURI
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	client := NewLoadTestClient(NewAPIClient(&Configuration{
		BasePath:   server.URL + "/loadTest/manager/api",
		HTTPClient: server.Client(),
	}))
	_, _, err := client.Delete(t.Context(), Scope{AccountID: "acc"}, "checkout?v2")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(gotRequest, "/loadTest/manager/api/v1/load-tests/checkout%3Fv2?") {
		t.Fatalf("request = %s", gotRequest)
	}
}

func TestUpdateAndDeleteLoadTest(t *testing.T) {
	var method, path string
	var gotBody map[string]interface{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		method = r.Method
		path = r.URL.Path
		if r.Method == http.MethodPut {
			body, _ := io.ReadAll(r.Body)
			_ = json.Unmarshal(body, &gotBody)
			_, _ = w.Write([]byte(`{"identity":"checkout","name":"Checkout v2","toolType":"K6"}`))
			return
		}
		_, _ = w.Write([]byte(`{"success":true,"message":"deleted"}`))
	}))
	defer server.Close()

	client := NewLoadTestClient(NewAPIClient(&Configuration{
		BasePath:   server.URL + "/loadTest/manager/api",
		HTTPClient: server.Client(),
	}))
	managed := false
	updated, _, err := client.Update(t.Context(), Scope{AccountID: "acc", OrgID: "default", ProjectID: "demo"}, "checkout", LoadTestUpdateRequest{
		Name:                  "Checkout v2",
		Tags:                  []string{},
		Variables:             []LoadTestVariable{},
		ManagedInfra:          &managed,
		EnvironmentIdentifier: "prod",
		InfraIdentifier:       "payments",
		ToolConfig:            map[string]interface{}{"mode": "script"},
		CleanupPolicy:         CleanupPolicyRetain,
	})
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodPut || path != "/loadTest/manager/api/v1/load-tests/checkout" {
		t.Fatalf("update %s %s", method, path)
	}
	if _, ok := gotBody["tags"].([]interface{}); !ok {
		t.Fatalf("tags = %#v", gotBody["tags"])
	}
	if _, ok := gotBody["variables"].([]interface{}); !ok {
		t.Fatalf("variables = %#v", gotBody["variables"])
	}
	if updated.Name != "Checkout v2" {
		t.Fatalf("updated = %+v", updated)
	}

	deleted, _, err := client.Delete(t.Context(), Scope{AccountID: "acc"}, "checkout")
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodDelete || !deleted.Success {
		t.Fatalf("delete method %s response %+v", method, deleted)
	}
}

func TestLoadTestRuns(t *testing.T) {
	var method, path, query string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		method = r.Method
		path = r.URL.Path
		query = r.URL.RawQuery
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/loadTest/manager/api/v1/load-tests/checkout/runs":
			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write([]byte(`{"identity":"run-1","loadTestIdentity":"checkout","status":"Pending"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/loadTest/manager/api/v1/runs/run-1":
			_, _ = w.Write([]byte(`{"identity":"run-1","status":"Running"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/loadTest/manager/api/v1/load-tests/checkout/runs":
			_, _ = w.Write([]byte(`{"items":[{"identity":"run-1","status":"Running"}],"pagination":{"totalItems":1}}`))
		case r.Method == http.MethodPost && r.URL.Path == "/loadTest/manager/api/v1/runs/run-1/stop":
			_, _ = w.Write([]byte(`{"success":true}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	client := NewLoadTestClient(NewAPIClient(&Configuration{
		BasePath:   server.URL + "/loadTest/manager/api",
		HTTPClient: server.Client(),
	}))
	scope := Scope{AccountID: "acc", OrgID: "default", ProjectID: "demo"}
	run, _, err := client.CreateRun(t.Context(), scope, "checkout", LoadTestRunCreateRequest{Identity: "run-1"})
	if err != nil {
		t.Fatal(err)
	}
	if run.Status != LoadTestRunStatusPending || path != "/loadTest/manager/api/v1/load-tests/checkout/runs" {
		t.Fatalf("create run path %s status %s", path, run.Status)
	}
	got, _, err := client.GetRun(t.Context(), scope, "run-1")
	if err != nil || got.Status != LoadTestRunStatusRunning {
		t.Fatalf("get run %+v err %v path %s", got, err, path)
	}
	list, _, err := client.ListRuns(t.Context(), "checkout", ListLoadTestRunsRequest{Scope: scope, Status: "Running"})
	if err != nil || len(list.Items) != 1 || !strings.Contains(query, "status=Running") {
		t.Fatalf("list %+v query %s err %v", list, query, err)
	}
	stopped, _, err := client.StopRun(t.Context(), scope, "checkout", "run-1")
	if err != nil || !stopped.Success || method != http.MethodPost || !strings.Contains(query, "loadTestIdentity=checkout") {
		t.Fatalf("stop %+v %s %s err %v", stopped, method, query, err)
	}
}
