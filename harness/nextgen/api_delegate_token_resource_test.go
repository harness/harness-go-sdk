package nextgen_test

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/stretchr/testify/require"
)

/*
	    # To run only TestDelegateTokenCreateRevokeDelete use following
		go test ./harness/nextgen \
		  -run '^TestDelegateTokenCreateRevokeDelete$' \
		  -v -count=1 \
		  -args \
		  -account <ACCOUNT_ID> \
		  -api-key <API_KEY> \
		  -org <ORG_ID> \
		  -project <PROJECT_ID>
*/

var (
	accountID = flag.String("account", "", "Harness account ID")
	apiKey    = flag.String("api-key", "", "Harness API key")
	endpoint  = flag.String("endpoint", "https://app.harness.io/gateway", "Harness gateway base URL")
	orgID     = flag.String("org", "", "Organization identifier")
	projectID = flag.String("project", "", "Project identifier")
)

func TestMain(m *testing.M) {
	flag.Parse()
	if *accountID == "" || *apiKey == "" {
		fmt.Println("Skipping tests: account ID or API key not provided")
		os.Exit(0)
	}
	os.Exit(m.Run())
}

func setupClient() (*nextgen.APIClient, context.Context) {
	cfg := nextgen.NewConfiguration()
	cfg.BasePath = *endpoint

	client := nextgen.NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), nextgen.ContextAPIKey, nextgen.APIKey{Key: *apiKey})
	return client, ctx
}

func createTokenOpts() *nextgen.DelegateTokenResourceApiCreateDelegateTokenOpts {
	if *orgID == "" && *projectID == "" {
		return nil
	}

	opts := &nextgen.DelegateTokenResourceApiCreateDelegateTokenOpts{}
	if *orgID != "" {
		opts.OrgIdentifier = optional.NewString(*orgID)
	}
	if *projectID != "" {
		opts.ProjectIdentifier = optional.NewString(*projectID)
	}
	return opts
}

func revokeTokenOpts() *nextgen.DelegateTokenResourceApiRevokeCgDelegateTokenOpts {
	if *orgID == "" && *projectID == "" {
		return nil
	}

	opts := &nextgen.DelegateTokenResourceApiRevokeCgDelegateTokenOpts{}
	if *orgID != "" {
		opts.OrgIdentifier = optional.NewString(*orgID)
	}
	if *projectID != "" {
		opts.ProjectIdentifier = optional.NewString(*projectID)
	}
	return opts
}

func deleteTokenOpts() *nextgen.DelegateTokenResourceApiDeleteCgDelegateTokenOpts {
	if *orgID == "" && *projectID == "" {
		return nil
	}

	opts := &nextgen.DelegateTokenResourceApiDeleteCgDelegateTokenOpts{}
	if *orgID != "" {
		opts.OrgIdentifier = optional.NewString(*orgID)
	}
	if *projectID != "" {
		opts.ProjectIdentifier = optional.NewString(*projectID)
	}
	return opts
}

func checkDelegateTokenTest(t *testing.T, testType string, resp nextgen.RestResponseDelegateTokenDetails, httpResp *http.Response, err error) {
	require.NoError(t, err)
	require.NotNil(t, httpResp)
	require.Less(t, httpResp.StatusCode, 300, testType)
}

func TestDelegateTokenCreateRevokeDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client, ctx := setupClient()

	tokenName := fmt.Sprintf("test-sdk-delegate-token-%d", time.Now().UnixNano())

	createResp, createHTTPResp, err := client.DelegateTokenResourceApi.CreateDelegateToken(ctx, *accountID, tokenName, createTokenOpts())
	checkDelegateTokenTest(t, "create", createResp, createHTTPResp, err)

	revokeResp, revokeHTTPResp, err := client.DelegateTokenResourceApi.RevokeCgDelegateToken(ctx, *accountID, tokenName, revokeTokenOpts())
	checkDelegateTokenTest(t, "revoke", revokeResp, revokeHTTPResp, err)

	const maxRetry = 3
	var lastErr error
	var deleteHTTPResp *http.Response

	for i := 0; i < maxRetry; i++ {
		httpResp, err := client.DelegateTokenResourceApi.DeleteCgDelegateToken(ctx, *accountID, tokenName, deleteTokenOpts())
		if err == nil && httpResp != nil && httpResp.StatusCode < 300 {
			require.NoError(t, err)
			require.NotNil(t, httpResp)
			require.Less(t, httpResp.StatusCode, 300, "delete")
			lastErr = nil
			break
		}
		lastErr = err
		deleteHTTPResp = httpResp
		time.Sleep(5 * time.Second)
	}

	if lastErr != nil {
		require.NotNil(t, deleteHTTPResp)
	}
	require.NoError(t, lastErr)
}
