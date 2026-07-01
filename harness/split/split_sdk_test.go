// Integration tests for the Split API client. Skipped when credentials are missing or -short is set.
//
// Run with: go test -v ./harness/split/... -run TestSplitSDK
// Requires: HARNESS_ACCOUNT_ID, HARNESS_PLATFORM_API_KEY (or use -account and -api-key flags).
// Skip integration tests: go test -short ./harness/split/...
package split_test

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/split"
)

var (
	integrationAccountID = flag.String("account", "", "Harness account ID (or set HARNESS_ACCOUNT_ID)")
	integrationApiKey    = flag.String("api-key", "", "Harness/Split API key (or set HARNESS_PLATFORM_API_KEY)")
)

func TestMain(m *testing.M) {
	flag.Parse()
	if *integrationAccountID == "" {
		*integrationAccountID = os.Getenv("HARNESS_ACCOUNT_ID")
	}
	if *integrationApiKey == "" {
		*integrationApiKey = os.Getenv("HARNESS_PLATFORM_API_KEY")
	}
	if *integrationAccountID == "" || *integrationApiKey == "" {
		fmt.Println("Skipping Split SDK integration tests: HARNESS_ACCOUNT_ID and HARNESS_PLATFORM_API_KEY (or -account and -api-key) not set")
		os.Exit(0)
	}
	os.Exit(m.Run())
}

func setupIntegrationClient(t *testing.T) *split.APIClient {
	t.Helper()
	cfg := split.NewDefaultConfiguration()
	cfg.AccountId = *integrationAccountID
	cfg.ApiKey = *integrationApiKey
	if cfg.ApiKey != "" {
		cfg.AddDefaultHeader("x-api-key", cfg.ApiKey)
	}
	return split.NewAPIClient(cfg)
}

// TestSplitSDK_WorkspacesResolveAndList runs against the real Split API when credentials are set and -short is not used.
func TestSplitSDK_WorkspacesResolveAndList(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}
	client := setupIntegrationClient(t)
	// List workspaces (does not require org/project)
	list, err := client.Workspaces.List(nil)
	if err != nil {
		t.Fatalf("Workspaces.List: %v", err)
	}
	t.Logf("Workspaces count: %d", len(list))
	if len(list) > 0 {
		t.Logf("First workspace: %s (%s)", list[0].Name, list[0].ID)
		// Resolve by org/project if available
		if list[0].OrganizationIdentifier != "" && list[0].ProjectIdentifier != "" {
			id, err := client.Workspaces.ResolveWorkspaceID(list[0].OrganizationIdentifier, list[0].ProjectIdentifier)
			if err != nil {
				t.Logf("ResolveWorkspaceID: %v (optional)", err)
			} else {
				t.Logf("ResolveWorkspaceID: %s", id)
			}
		}
	}
}
