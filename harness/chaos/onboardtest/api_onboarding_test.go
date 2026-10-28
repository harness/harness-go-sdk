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

func TestCreateServiceOnboarding(t *testing.T) {
	var onboarded, confirmed bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("accountIdentifier") != "acc" || r.URL.Query().Get("projectIdentifier") != "proj" {
			t.Errorf("scope query = %s", r.URL.RawQuery)
		}
		if onboardingPrelude(w, r) {
			return
		}
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/rest/v2/onboarding":
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
		case r.Method == http.MethodGet && r.URL.Path == "/rest/v2/onboarding/ob-1":
			writeOnboarding(w, "ob-1", "COMPLETED", "RUNNING")
		case r.Method == http.MethodPost && r.URL.Path == "/rest/v2/onboarding-onboard-services/ob-1":
			var onboard struct {
				ExternalServiceIDs []string `json:"externalServiceIds"`
			}
			if err := json.NewDecoder(r.Body).Decode(&onboard); err != nil {
				t.Fatal(err)
			}
			if len(onboard.ExternalServiceIDs) != 2 || onboard.ExternalServiceIDs[0] != "svc-a" || onboard.ExternalServiceIDs[1] != "svc-b" {
				t.Fatalf("onboard body = %+v", onboard)
			}
			onboarded = true
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"onboardedCount":2,"onboardedServiceIds":["svc-a","svc-b"]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/rest/v2/onboarding-confirm-service-onboarding/ob-1":
			confirmed = true
			w.WriteHeader(http.StatusOK)
		case r.Method == http.MethodGet && r.URL.Path == "/rest/v2/onboarding/ob-1/services":
			writeJSON(w, `{"services":[{"externalServiceId":"svc-a"},{"externalServiceId":"svc-b"}],"pagination":{"totalPages":1}}`)
		default:
			t.Fatalf("unexpected %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := chaos.NewOnboardingClient(chaos.NewAPIClient(&chaos.Configuration{BasePath: server.URL, HTTPClient: server.Client()}))
	client.PollInterval = time.Millisecond
	got, _, err := client.CreateServiceOnboarding(context.Background(), chaos.Scope{AccountID: "acc", ProjectID: "proj"}, chaos.ServiceOnboardingRequest{
		EnvironmentID:    "hpb",
		InfrastructureID: "payments",
		ServiceIDs:       []string{"svc-a", "svc-b"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if !onboarded || !confirmed {
		t.Fatalf("onboarded=%v confirmed=%v", onboarded, confirmed)
	}
	if got.OnboardingID != "ob-1" || got.DiscoveryStatus != "COMPLETED" || got.ServiceStatus != "RUNNING" || got.OnboardedCount != 2 {
		t.Fatalf("%+v", got)
	}
	if got.InstallCommand == "" {
		t.Fatal("install command was empty")
	}
	if len(got.OnboardedServiceIDs) != 2 {
		t.Fatalf("onboarded ids = %v", got.OnboardedServiceIDs)
	}
}

func TestCreateServiceOnboardingAllDiscoveredServices(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if onboardingPrelude(w, r) {
			return
		}
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/chaos/manager/api/rest/v2/onboarding":
			writeOnboarding(w, "ob-1", "COMPLETED", "")
		case r.Method == http.MethodGet && r.URL.Path == "/servicediscovery/api/v1/agents/payments/discoveredservices":
			if r.URL.Query().Get("all") != "true" || r.URL.Query().Get("environmentIdentifier") != "hpb" {
				t.Fatalf("discovery query = %s", r.URL.RawQuery)
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"items":[{"id":"svc-a"},{"id":"","removed":false},{"id":"gone","removed":true}]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/chaos/manager/api/rest/v2/onboarding-onboard-services/ob-1":
			var onboard struct {
				ExternalServiceIDs []string `json:"externalServiceIds"`
			}
			if err := json.NewDecoder(r.Body).Decode(&onboard); err != nil {
				t.Fatal(err)
			}
			if len(onboard.ExternalServiceIDs) != 1 || onboard.ExternalServiceIDs[0] != "svc-a" {
				t.Fatalf("onboarded = %+v", onboard)
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"onboardedCount":1,"onboardedServiceIds":["svc-a"]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/chaos/manager/api/rest/v2/onboarding-confirm-service-onboarding/ob-1":
			w.WriteHeader(http.StatusOK)
		case r.Method == http.MethodGet && r.URL.Path == "/chaos/manager/api/rest/v2/onboarding/ob-1":
			writeOnboarding(w, "ob-1", "COMPLETED", "RUNNING")
		case r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/services"):
			writeJSON(w, `{"services":[{"externalServiceId":"svc-a"}],"pagination":{"totalPages":1}}`)
		default:
			t.Fatalf("unexpected %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := chaos.NewOnboardingClient(chaos.NewAPIClient(&chaos.Configuration{
		BasePath:   server.URL + "/chaos/manager/api",
		HTTPClient: server.Client(),
	}))
	got, _, err := client.CreateServiceOnboarding(context.Background(), chaos.Scope{AccountID: "acc"}, chaos.ServiceOnboardingRequest{
		EnvironmentID:    "hpb",
		InfrastructureID: "payments",
	})
	if err != nil {
		t.Fatal(err)
	}
	if got.OnboardingID != "ob-1" || len(got.OnboardedServiceIDs) != 1 || got.OnboardedServiceIDs[0] != "svc-a" {
		t.Fatalf("%+v", got)
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
	if got.DiscoveryStatus != "COMPLETED" || got.ServiceStatus != "RUNNING" || got.OnboardedCount != 1 {
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
	got, _, err := client.CreateServiceOnboarding(context.Background(), chaos.Scope{AccountID: "acc"}, chaos.ServiceOnboardingRequest{
		EnvironmentID:    "hpb",
		InfrastructureID: "payments",
		ServiceIDs:       []string{"svc-a"},
	})
	if err == nil || !strings.Contains(err.Error(), "discovery for onboarding ob-1 is ERRORED") {
		t.Fatal(err)
	}
	if got == nil || got.OnboardingID != "ob-1" {
		t.Fatalf("session = %+v", got)
	}
}

func TestCreateServiceOnboardingNoDiscoveredServices(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if onboardingPrelude(w, r) {
			return
		}
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/chaos/manager/api/rest/v2/onboarding":
			writeOnboarding(w, "ob-1", "COMPLETED", "")
		case r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/onboarding/ob-1"):
			writeOnboarding(w, "ob-1", "COMPLETED", "")
		case r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/discoveredservices"):
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"items":[{"id":"gone","removed":true}]}`))
		default:
			t.Fatalf("unexpected %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := chaos.NewOnboardingClient(chaos.NewAPIClient(&chaos.Configuration{
		BasePath:   server.URL + "/chaos/manager/api",
		HTTPClient: server.Client(),
	}))
	got, _, err := client.CreateServiceOnboarding(context.Background(), chaos.Scope{AccountID: "acc"}, chaos.ServiceOnboardingRequest{
		EnvironmentID:    "hpb",
		InfrastructureID: "payments",
	})
	if err == nil || !strings.Contains(err.Error(), "no services") {
		t.Fatal(err)
	}
	if got == nil || got.OnboardingID != "ob-1" {
		t.Fatalf("session = %+v", got)
	}
}

func TestCreateServiceOnboardingRequiresInfrastructure(t *testing.T) {
	client := chaos.NewOnboardingClient(chaos.NewAPIClient(&chaos.Configuration{BasePath: "http://unused"}))
	_, _, err := client.CreateServiceOnboarding(context.Background(), chaos.Scope{}, chaos.ServiceOnboardingRequest{EnvironmentID: "prod"})
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
		case r.Method == http.MethodPost && r.URL.Path == "/rest/v2/onboarding":
			writeOnboarding(w, "ob-1", "COMPLETED", "")
		case r.Method == http.MethodGet && r.URL.Path == "/rest/v2/onboarding/ob-1":
			writeOnboarding(w, "ob-1", "COMPLETED", "")
		case r.Method == http.MethodPost && strings.Contains(r.URL.Path, "onboard-services"):
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"message":"discovery step is not completed yet"}`))
		default:
			t.Fatalf("unexpected %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := chaos.NewOnboardingClient(chaos.NewAPIClient(&chaos.Configuration{BasePath: server.URL, HTTPClient: server.Client()}))
	got, resp, err := client.CreateServiceOnboarding(context.Background(), chaos.Scope{AccountID: "acc"}, chaos.ServiceOnboardingRequest{
		EnvironmentID:    "hpb",
		InfrastructureID: "payments",
		ServiceIDs:       []string{"svc-a"},
	})
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

func TestOnboardServicesRequiresIDs(t *testing.T) {
	client := chaos.NewOnboardingClient(chaos.NewAPIClient(&chaos.Configuration{BasePath: "http://unused"}))
	_, _, err := client.OnboardServices(context.Background(), chaos.Scope{}, "ob-1", nil)
	if err == nil || err.Error() != "at least one service id is required" {
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
