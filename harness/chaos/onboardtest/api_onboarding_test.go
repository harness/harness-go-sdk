package onboardtest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/harness/harness-go-sdk/harness/chaos"
)

func TestOpenAndFinishServiceOnboarding(t *testing.T) {
	var created bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("accountIdentifier") != "acc" || r.URL.Query().Get("projectIdentifier") != "proj" {
			t.Errorf("scope query = %s", r.URL.RawQuery)
		}
		if onboardingPrelude(w, r) {
			return
		}
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/chaos/manager/api/rest/v2/onboarding":
			var body chaos.V2OnboardingV2OnboardingRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatal(err)
			}
			if body.EnvironmentRef != "hpb" || body.InfrastructureRef != "payments" {
				t.Fatalf("refs = %s %s", body.EnvironmentRef, body.InfrastructureRef)
			}
			if body.Mode == nil || *body.Mode != chaos.MANUAL_V2OnboardingOnboardingMode {
				t.Fatalf("mode = %v", body.Mode)
			}
			if body.ManualMethod == nil || *body.ManualMethod != chaos.SERVICE_V2OnboardingManualMethod {
				t.Fatalf("method = %v", body.ManualMethod)
			}
			writeOnboarding(w, "ob-1", "QUEUED", "")
		case r.Method == http.MethodGet && r.URL.Path == "/chaos/manager/api/rest/v2/onboarding/ob-1":
			writeOnboarding(w, "ob-1", "COMPLETED", "RUNNING")
		case r.Method == http.MethodGet && r.URL.Path == "/servicediscovery/api/v1/agents/payments/discoveredservices":
			writeJSON(w, `{"items":[{"id":"svc-a","identity":"accesscontrol","name":"Access Control","type":"Kubernetes","spec":{"kubernetes":{"namespace":"chaos"}}}]}`)
		case r.Method == http.MethodPost && r.URL.Path == "/chaos/manager/api/v3/chaos-services":
			var body struct {
				Identity           string `json:"identity"`
				Name               string `json:"name"`
				AgentID            string `json:"agentId"`
				EnvironmentID      string `json:"environmentId"`
				InfrastructureID   string `json:"infrastructureId"`
				InfrastructureType string `json:"infrastructureType"`
				ExternalServiceID  string `json:"externalServiceId"`
				OnboardingID       string `json:"onboardingId"`
				Probes             []struct {
					ProbeID string `json:"probeId"`
					Inputs  []struct {
						Name  string `json:"name"`
						Value string `json:"value"`
					} `json:"inputs"`
				} `json:"probes"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatal(err)
			}
			if body.ExternalServiceID != "svc-a" || body.Name != "Access Control" || body.AgentID != "payments" {
				t.Fatalf("service = %+v", body)
			}
			if body.Identity != "accesscontrol-chaos-payments" || body.InfrastructureType != "KubernetesV2" {
				t.Fatalf("identity = %s type = %s", body.Identity, body.InfrastructureType)
			}
			if body.EnvironmentID != "hpb" || body.InfrastructureID != "payments" || body.OnboardingID != "ob-1" {
				t.Fatalf("scope = %s %s onboarding = %s", body.EnvironmentID, body.InfrastructureID, body.OnboardingID)
			}
			if len(body.Probes) != 2 || body.Probes[0].ProbeID != "http-check" || body.Probes[1].ProbeID != "pod-status" {
				t.Fatalf("probes = %+v", body.Probes)
			}
			if len(body.Probes[0].Inputs) != 1 || body.Probes[0].Inputs[0].Name != "url" || body.Probes[0].Inputs[0].Value != "http://svc" {
				t.Fatalf("inputs = %+v", body.Probes[0].Inputs)
			}
			if body.Probes[1].Inputs != nil {
				t.Fatalf("second probe inputs = %+v", body.Probes[1].Inputs)
			}
			created = true
			writeJSON(w, `{"identity":"accesscontrol-chaos-payments","name":"Access Control","externalServiceId":"svc-a"}`)
		default:
			t.Fatalf("unexpected %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := chaos.NewOnboardingClient(chaos.NewAPIClient(&chaos.Configuration{
		BasePath:   server.URL + "/chaos/manager/api",
		HTTPClient: server.Client(),
	}))
	client.PollInterval = time.Millisecond
	req := chaos.ServiceOnboardingRequest{
		EnvironmentID:    "hpb",
		InfrastructureID: "payments",
		ServiceID:        "accesscontrol",
		Probes: []chaos.ServiceProbe{
			{ProbeID: "http-check", Inputs: map[string]string{"url": "http://svc"}},
			{ProbeID: "pod-status"},
		},
	}
	scope := chaos.Scope{AccountID: "acc", ProjectID: "proj"}
	opened, _, err := client.OpenServiceOnboarding(context.Background(), scope, req)
	if err != nil {
		t.Fatal(err)
	}
	got, _, err := client.FinishServiceOnboarding(context.Background(), scope, opened, req)
	if err != nil {
		t.Fatal(err)
	}
	if !created {
		t.Fatal("chaos service was not created")
	}
	if got.OnboardingID != "ob-1" || got.DiscoveryStatus != "COMPLETED" || got.ChaosServiceID != "accesscontrol-chaos-payments" {
		t.Fatalf("%+v", got)
	}
	if got.InstallCommand == "" {
		t.Fatal("install command was empty")
	}
}

func TestGetServiceOnboarding(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/rest/v2/onboarding/ob-1" {
			t.Fatalf("unexpected %s %s", r.Method, r.URL.Path)
		}
		writeOnboarding(w, "ob-1", "COMPLETED", "RUNNING")
	}))
	defer server.Close()

	client := chaos.NewOnboardingClient(chaos.NewAPIClient(&chaos.Configuration{BasePath: server.URL, HTTPClient: server.Client()}))
	got, _, err := client.Get(context.Background(), chaos.Scope{AccountID: "acc"}, "ob-1")
	if err != nil {
		t.Fatal(err)
	}
	if got.DiscoveryStatus != "COMPLETED" || got.ServiceStatus != "RUNNING" {
		t.Fatalf("%+v", got)
	}
}

func TestCreateServiceOnboardingDiscoveryErrored(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if onboardingPrelude(w, r) {
			return
		}
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected %s %s", r.Method, r.URL.Path)
		}
		writeOnboarding(w, "ob-1", "ERRORED", "")
	}))
	defer server.Close()

	client := chaos.NewOnboardingClient(chaos.NewAPIClient(&chaos.Configuration{BasePath: server.URL, HTTPClient: server.Client()}))
	req := chaos.ServiceOnboardingRequest{EnvironmentID: "hpb", InfrastructureID: "payments", ServiceID: "svc-a"}
	opened, _, err := client.OpenServiceOnboarding(context.Background(), chaos.Scope{AccountID: "acc"}, req)
	if err != nil {
		t.Fatal(err)
	}
	got, _, err := client.FinishServiceOnboarding(context.Background(), chaos.Scope{AccountID: "acc"}, opened, req)
	if err == nil || !strings.Contains(err.Error(), "discovery for onboarding ob-1 is ERRORED") {
		t.Fatal(err)
	}
	if got == nil || got.OnboardingID != "ob-1" {
		t.Fatalf("session = %+v", got)
	}
}

func TestCreateServiceOnboardingServiceNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if onboardingPrelude(w, r) {
			return
		}
		switch {
		case r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/rest/v2/onboarding"):
			writeOnboarding(w, "ob-1", "COMPLETED", "")
		case r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/onboarding/ob-1"):
			writeOnboarding(w, "ob-1", "COMPLETED", "")
		case r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/discoveredservices"):
			writeJSON(w, `{"items":[{"id":"gone","removed":true}]}`)
		default:
			t.Fatalf("unexpected %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := chaos.NewOnboardingClient(chaos.NewAPIClient(&chaos.Configuration{
		BasePath:   server.URL + "/chaos/manager/api",
		HTTPClient: server.Client(),
	}))
	req := chaos.ServiceOnboardingRequest{EnvironmentID: "hpb", InfrastructureID: "payments", ServiceID: "svc-a"}
	opened, _, err := client.OpenServiceOnboarding(context.Background(), chaos.Scope{AccountID: "acc"}, req)
	if err != nil {
		t.Fatal(err)
	}
	got, _, err := client.FinishServiceOnboarding(context.Background(), chaos.Scope{AccountID: "acc"}, opened, req)
	if err == nil || !strings.Contains(err.Error(), "discovered service svc-a was not found") {
		t.Fatal(err)
	}
	if got == nil || got.OnboardingID != "ob-1" {
		t.Fatalf("session = %+v", got)
	}
}

func TestCreateServiceOnboardingRequiresInfrastructure(t *testing.T) {
	client := chaos.NewOnboardingClient(chaos.NewAPIClient(&chaos.Configuration{BasePath: "http://unused"}))
	_, _, err := client.OpenServiceOnboarding(context.Background(), chaos.Scope{}, chaos.ServiceOnboardingRequest{EnvironmentID: "prod"})
	if err == nil || err.Error() != "environment id and infrastructure id are required" {
		t.Fatal(err)
	}
}

func TestOnboardServicesAPIErrorReturnsSession(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if onboardingPrelude(w, r) {
			return
		}
		switch {
		case r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/rest/v2/onboarding"):
			writeOnboarding(w, "ob-1", "COMPLETED", "")
		case r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/onboarding/ob-1"):
			writeOnboarding(w, "ob-1", "COMPLETED", "")
		case r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/discoveredservices"):
			writeJSON(w, `{"items":[{"id":"svc-a","name":"auth"}]}`)
		case r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/v3/chaos-services"):
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"message":"discovery step is not completed yet"}`))
		default:
			t.Fatalf("unexpected %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := chaos.NewOnboardingClient(chaos.NewAPIClient(&chaos.Configuration{
		BasePath:   server.URL + "/chaos/manager/api",
		HTTPClient: server.Client(),
	}))
	req := chaos.ServiceOnboardingRequest{EnvironmentID: "hpb", InfrastructureID: "payments", ServiceID: "svc-a"}
	opened, _, err := client.OpenServiceOnboarding(context.Background(), chaos.Scope{AccountID: "acc"}, req)
	if err != nil {
		t.Fatal(err)
	}
	got, resp, err := client.FinishServiceOnboarding(context.Background(), chaos.Scope{AccountID: "acc"}, opened, req)
	if err == nil || !strings.Contains(err.Error(), "discovery step is not completed yet") {
		t.Fatal(err)
	}
	if resp == nil || resp.StatusCode != http.StatusBadRequest || got == nil || got.OnboardingID != "ob-1" {
		t.Fatalf("session=%+v resp=%v", got, resp)
	}
}

func TestFindServiceOnboardingEscapesRegex(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("search") != `payments\.prod\+1` {
			t.Fatalf("search = %q", r.URL.Query().Get("search"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":[]}`))
	}))
	defer server.Close()

	client := chaos.NewOnboardingClient(chaos.NewAPIClient(&chaos.Configuration{BasePath: server.URL, HTTPClient: server.Client()}))
	got, _, err := client.FindServiceOnboarding(context.Background(), chaos.Scope{AccountID: "acc"}, "payments.prod+1")
	if err != nil {
		t.Fatal(err)
	}
	if got != nil {
		t.Fatalf("session = %+v", got)
	}
}

func TestCreateChaosServiceIdentity(t *testing.T) {
	var gotIdentity, gotAgent, gotInfra, gotOnboarding string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/onboarding/ob-1"):
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]interface{}{
				"onboardingID": "ob-1", "environmentRef": "hpb", "infrastructureRef": "payments",
				"manualMethod":            "Service",
				"discoveredServiceStatus": map[string]string{"status": "COMPLETED", "sDAgentIdentity": "agent-1"},
			})
		case r.Method == http.MethodGet && r.URL.Path == "/servicediscovery/api/v1/agents/agent-1/discoveredservices":
			writeJSON(w, `{"items":[{"id":"svc-a","name":"Cart Service","type":"LinuxVM","spec":{"kubernetes":{"namespace":"shop"}}}]}`)
		case r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/v3/chaos-services"):
			var body struct {
				Identity           string `json:"identity"`
				AgentID            string `json:"agentId"`
				InfrastructureID   string `json:"infrastructureId"`
				InfrastructureType string `json:"infrastructureType"`
				OnboardingID       string `json:"onboardingId"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatal(err)
			}
			if body.InfrastructureType != "Linux" {
				t.Fatalf("infrastructure type = %s", body.InfrastructureType)
			}
			gotIdentity, gotAgent, gotInfra, gotOnboarding = body.Identity, body.AgentID, body.InfrastructureID, body.OnboardingID
			writeJSON(w, `{"identity":"cartservice-agent-1"}`)
		default:
			t.Fatalf("unexpected %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := chaos.NewOnboardingClient(chaos.NewAPIClient(&chaos.Configuration{
		BasePath:   server.URL + "/chaos/manager/api",
		HTTPClient: server.Client(),
	}))
	got, _, err := client.CreateChaosService(context.Background(), chaos.Scope{AccountID: "acc"}, "ob-1", chaos.ServiceOnboardingRequest{
		ServiceID: "svc-a",
	})
	if err != nil {
		t.Fatal(err)
	}
	if gotIdentity != "cartservice-agent-1" || gotAgent != "custom-discovery-agent" || gotInfra != "payments" || gotOnboarding != "ob-1" {
		t.Fatalf("identity=%s agent=%s infra=%s onboarding=%s", gotIdentity, gotAgent, gotInfra, gotOnboarding)
	}
	if got.Identity != "cartservice-agent-1" {
		t.Fatalf("%+v", got)
	}
}

func TestCreateChaosServiceByName(t *testing.T) {
	var externalID string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/onboarding/ob-1"):
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]interface{}{
				"onboardingID": "ob-1", "environmentRef": "hpb", "infrastructureRef": "payments",
				"manualMethod":            "Service",
				"discoveredServiceStatus": map[string]string{"status": "COMPLETED", "sDAgentIdentity": "payments"},
			})
		case r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/discoveredservices"):
			writeJSON(w, `{"items":[{"id":"6a9fe56a53e356f4436659e7","identity":"bankofanthosconfig","name":"bank-of-anthos-config","type":"Kubernetes","spec":{"kubernetes":{"namespace":"banking"}}}]}`)
		case r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/v3/chaos-services"):
			var body struct {
				ExternalServiceID string `json:"externalServiceId"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatal(err)
			}
			externalID = body.ExternalServiceID
			writeJSON(w, `{"identity":"bank-of-anthos-config-banking-payments"}`)
		default:
			t.Fatalf("unexpected %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := chaos.NewOnboardingClient(chaos.NewAPIClient(&chaos.Configuration{
		BasePath:   server.URL + "/chaos/manager/api",
		HTTPClient: server.Client(),
	}))
	_, _, err := client.CreateChaosService(context.Background(), chaos.Scope{AccountID: "acc"}, "ob-1", chaos.ServiceOnboardingRequest{
		ServiceName: "bank-of-anthos-config",
	})
	if err != nil {
		t.Fatal(err)
	}
	if externalID != "6a9fe56a53e356f4436659e7" {
		t.Fatalf("external service id = %s", externalID)
	}

	dup := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/onboarding/ob-1"):
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]interface{}{
				"onboardingID": "ob-1", "environmentRef": "hpb", "infrastructureRef": "payments",
				"discoveredServiceStatus": map[string]string{"status": "COMPLETED", "sDAgentIdentity": "payments"},
			})
		case r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/discoveredservices"):
			writeJSON(w, `{"items":[{"id":"a","identity":"bankofanthosconfig","name":"bank-of-anthos-config"},{"id":"b","identity":"bankofanthosconfig2","name":"bank-of-anthos-config"}]}`)
		default:
			t.Fatalf("unexpected %s %s", r.Method, r.URL.Path)
		}
	}))
	defer dup.Close()
	dupClient := chaos.NewOnboardingClient(chaos.NewAPIClient(&chaos.Configuration{
		BasePath:   dup.URL + "/chaos/manager/api",
		HTTPClient: dup.Client(),
	}))
	_, _, err = dupClient.CreateChaosService(context.Background(), chaos.Scope{AccountID: "acc"}, "ob-1", chaos.ServiceOnboardingRequest{
		ServiceName: "bank-of-anthos-config",
	})
	if err == nil || !strings.Contains(err.Error(), "bankofanthosconfig") || !strings.Contains(err.Error(), "discovered_service_id") {
		t.Fatal(err)
	}
}

func onboardingPrelude(w http.ResponseWriter, r *http.Request) bool {
	path := r.URL.Path
	switch {
	case r.Method == http.MethodGet && strings.HasSuffix(path, "/rest/v2/onboarding"):
		writeJSON(w, `{"data":[]}`)
	case r.Method == http.MethodGet && strings.Contains(path, "/rest/v2/infrastructure/"):
		writeJSON(w, `{"identity":"payments","infraNamespace":"chaos","serviceAccount":"litmus","environmentID":"hpb"}`)
	case r.Method == http.MethodPost && strings.HasSuffix(path, "/rest/v2/infrastructure/yaml"):
		writeJSON(w, `{"yaml":"apiVersion: v1\nkind: List\n"}`)
	case r.Method == http.MethodPost && strings.Contains(path, "/onboarding-confirm-discovery/"):
		writeJSON(w, `{"agentCreated":true,"agentLinked":true,"discoveryCompleted":true}`)
	default:
		return false
	}
	return true
}

func writeJSON(w http.ResponseWriter, body string) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(body))
}

func writeOnboarding(w http.ResponseWriter, id, discovery, service string) {
	w.Header().Set("Content-Type", "application/json")
	body := map[string]interface{}{
		"onboardingID":            id,
		"environmentRef":          "hpb",
		"infrastructureRef":       "payments",
		"mode":                    "Manual",
		"manualMethod":            "Service",
		"status":                  "PENDING",
		"discoveredServiceStatus": map[string]string{"status": discovery},
		"onboardingChaosServiceStatus": map[string]interface{}{
			"status":            service,
			"chaosServiceCount": 1,
		},
	}
	_ = json.NewEncoder(w).Encode(body)
}
