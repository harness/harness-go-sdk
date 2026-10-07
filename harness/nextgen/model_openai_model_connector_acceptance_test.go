package nextgen

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

/*
OpenAI Model Connector — Acceptance Tests (TestAccOpenAIModel*_*)
Tests Create, Get, and Delete of the OpenAI Model Connector against a live
Harness account. Covers two auth types (Token, Vertex)
at three scopes (Account, Org, Project) = 6 tests total.

Optional env vars (apply to all scopes):

    # Comma-separated delegate selectors (optional, omit to use any delegate)
    export OPENAI_MODEL_DELEGATE_SELECTORS="delegate01,delegate02"

    # Set to "true" to skip cleanup deletion (for manual UI verification)
    export IS_SKIP_DELETION_TEST="true"

---------------------------------------------------------------------------
# Account-Scope Tests
---------------------------------------------------------------------------

    export HARNESS_ACCOUNT_ID='<account-id>'
    export HARNESS_PLATFORM_ACCOUNT_API_KEY='<pat-or-sa-key>'
    # Token auth
    export OPENAI_MODEL_TOKEN_REF="account.<account-level-secret>"
    # Vertex auth
    export OPENAI_MODEL_VERTEX_SERVICE_ACCOUNT_KEY_REF="account.<account-level-secret>"
    export OPENAI_MODEL_VERTEX_PROJECT_ID="my-gcp-project"
    export OPENAI_MODEL_VERTEX_REGION="us-central1"
    export OPENAI_MODEL_DELEGATE_SELECTORS="delegate1,delegate2"

    go test -v -run "TestAccOpenAIModel.*_AccountScope" -count=1 -timeout 120s \
      ./harness/nextgen/ -args -account dummy -api-key dummy

---------------------------------------------------------------------------
# Org-Scope Tests
---------------------------------------------------------------------------

    export HARNESS_ACCOUNT_ID='<account-id>'
    export HARNESS_PLATFORM_ORG_API_KEY='<pat-or-sa-key>'
    export OPENAI_MODEL_ORG_SCOPE_ORG_ID="<org-id>"
    # Token auth
    export OPENAI_MODEL_TOKEN_REF="org.<org-level-secret>"
    # Vertex auth
    export OPENAI_MODEL_VERTEX_SERVICE_ACCOUNT_KEY_REF="org.<org-level-secret>"
    export OPENAI_MODEL_VERTEX_PROJECT_ID="my-gcp-project"
    export OPENAI_MODEL_VERTEX_REGION="us-central1"
    export OPENAI_MODEL_DELEGATE_SELECTORS="delegate1,delegate2"

    go test -v -run "TestAccOpenAIModel.*_OrgScope" -count=1 -timeout 120s \
      ./harness/nextgen/ -args -account dummy -api-key dummy

---------------------------------------------------------------------------
# Project-Scope Tests
---------------------------------------------------------------------------

    export HARNESS_ACCOUNT_ID='<account-id>'
    export HARNESS_PLATFORM_PROJECT_API_KEY='<pat-or-sa-key>'
    export OPENAI_MODEL_PROJECT_SCOPE_ORG_ID="<org-id>"
    export OPENAI_MODEL_PROJECT_SCOPE_PROJECT_ID="<project-id>"
    # Token auth
    export OPENAI_MODEL_TOKEN_REF="<project-level-secret>"
    # Vertex auth
    export OPENAI_MODEL_VERTEX_SERVICE_ACCOUNT_KEY_REF="<project-level-secret>"
    export OPENAI_MODEL_VERTEX_PROJECT_ID="my-gcp-project"
    export OPENAI_MODEL_VERTEX_REGION="us-central1"
    export OPENAI_MODEL_DELEGATE_SELECTORS="delegate1,delegate2"

    go test -v -run "TestAccOpenAIModel.*_ProjectScope" -count=1 -timeout 120s \
      ./harness/nextgen/ -args -account dummy -api-key dummy

*/

type openaiModelTestEnv struct {
	accountID string
	apiKey    string
	orgID     string
	projectID string
}

func getOpenAIModelTestEnv(t *testing.T, scope string) openaiModelTestEnv {
	t.Helper()

	var apiKeyEnvVar, orgEnvVar, projectEnvVar string
	switch scope {
	case "account":
		apiKeyEnvVar = "HARNESS_PLATFORM_ACCOUNT_API_KEY"
	case "org":
		apiKeyEnvVar = "HARNESS_PLATFORM_ORG_API_KEY"
		orgEnvVar = "OPENAI_MODEL_ORG_SCOPE_ORG_ID"
	case "project":
		apiKeyEnvVar = "HARNESS_PLATFORM_PROJECT_API_KEY"
		orgEnvVar = "OPENAI_MODEL_PROJECT_SCOPE_ORG_ID"
		projectEnvVar = "OPENAI_MODEL_PROJECT_SCOPE_PROJECT_ID"
	default:
		t.Fatalf("unknown scope %q: must be account, org, or project", scope)
	}

	env := openaiModelTestEnv{
		accountID: os.Getenv("HARNESS_ACCOUNT_ID"),
		apiKey:    os.Getenv(apiKeyEnvVar),
	}
	if orgEnvVar != "" {
		env.orgID = os.Getenv(orgEnvVar)
	}
	if projectEnvVar != "" {
		env.projectID = os.Getenv(projectEnvVar)
	}

	if env.accountID == "" || env.apiKey == "" {
		t.Skipf("HARNESS_ACCOUNT_ID and %s must be set for acceptance tests", apiKeyEnvVar)
	}
	return env
}

func newOpenAIModelClient(t *testing.T, env openaiModelTestEnv) (*APIClient, context.Context) {
	t.Helper()
	cfg := NewConfiguration()
	cfg.AccountId = env.accountID
	cfg.ApiKey = env.apiKey
	cfg.HTTPClient.HTTPClient.Transport = &loggingRoundTripper{
		t:       t,
		wrapped: cfg.HTTPClient.HTTPClient.Transport,
	}
	client := NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), ContextAPIKey, APIKey{Key: env.apiKey})
	return client, ctx
}

func createOpenAIModelConnector(
	t *testing.T,
	client *APIClient,
	ctx context.Context,
	accountID string,
	name string,
	auth *OpenAIModelAuthentication,
	orgID string,
	projectID string,
) ResponseDtoConnectorResponse {
	t.Helper()

	var delegateSelectors []string
	if ds := os.Getenv("OPENAI_MODEL_DELEGATE_SELECTORS"); ds != "" {
		for _, s := range strings.Split(ds, ",") {
			s = strings.TrimSpace(s)
			if s != "" {
				delegateSelectors = append(delegateSelectors, s)
			}
		}
	}

	connInfo := &ConnectorInfo{
		Name:              name,
		Identifier:        name,
		Description:       fmt.Sprintf("%s hsf97 testing only", name),
		OrgIdentifier:     orgID,
		ProjectIdentifier: projectID,
		Tags:              map[string]string{"hsf97": "true", "testing": "true"},
		Type_:             ConnectorTypes.OpenAIModel,
		OpenAIModel: &OpenAIModelConnector{
			Url:                  "https://api.openai.com",
			Model:                "gpt-4o",
			DelegateSelectors:    delegateSelectors,
			ExecuteOnDelegate:    true,
			IgnoreTestConnection: true,
			Authentication:       auth,
		},
	}

	resp, httpResp, err := client.ConnectorsApi.CreateConnector(
		ctx,
		Connector{Connector: connInfo},
		accountID,
		&ConnectorsApiCreateConnectorOpts{},
	)
	require.NoError(t, err, "CreateConnector failed")
	require.NotNil(t, httpResp)
	require.Equal(t, 200, httpResp.StatusCode)
	require.Equal(t, name, resp.Data.Connector.Name)
	require.Equal(t, ConnectorTypes.OpenAIModel, resp.Data.Connector.Type_)

	t.Logf("Created connector: name=%s scope=account:%s/org:%s/project:%s",
		name, accountID, orgID, projectID)

	return resp
}

func deleteOpenAIModelConnector(
	t *testing.T,
	client *APIClient,
	ctx context.Context,
	accountID string,
	identifier string,
	orgID string,
	projectID string,
) {
	t.Helper()
	if strings.EqualFold(os.Getenv("IS_SKIP_DELETION_TEST"), "true") {
		t.Logf("Skipping deletion (IS_SKIP_DELETION_TEST=true): %s", identifier)
		return
	}
	opts := &ConnectorsApiDeleteConnectorOpts{
		ForceDelete: optional.NewBool(true),
	}
	if orgID != "" {
		opts.OrgIdentifier = optional.NewString(orgID)
	}
	if projectID != "" {
		opts.ProjectIdentifier = optional.NewString(projectID)
	}
	_, _, err := client.ConnectorsApi.DeleteConnector(ctx, accountID, identifier, opts)
	if err != nil {
		t.Logf("Warning: cleanup delete failed for %s: %v", identifier, err)
	} else {
		t.Logf("Cleaned up connector: %s", identifier)
	}
}

func getOpenAIModelConnector(
	t *testing.T,
	client *APIClient,
	ctx context.Context,
	accountID string,
	identifier string,
	orgID string,
	projectID string,
) ResponseDtoConnectorResponse {
	t.Helper()
	opts := &ConnectorsApiGetConnectorOpts{}
	if orgID != "" {
		opts.OrgIdentifier = optional.NewString(orgID)
	}
	if projectID != "" {
		opts.ProjectIdentifier = optional.NewString(projectID)
	}
	resp, httpResp, err := client.ConnectorsApi.GetConnector(ctx, accountID, identifier, opts)
	require.NoError(t, err, "GetConnector failed")
	require.NotNil(t, httpResp)
	require.Equal(t, 200, httpResp.StatusCode)
	return resp
}

// ---------------------------------------------------------------------------
// Token Auth
// ---------------------------------------------------------------------------

func TestAccOpenAIModelToken_AccountScope(t *testing.T) {
	env := getOpenAIModelTestEnv(t, "account")
	tokenRef := os.Getenv("OPENAI_MODEL_TOKEN_REF")
	if tokenRef == "" {
		t.Skip("OPENAI_MODEL_TOKEN_REF must be set")
	}

	client, ctx := newOpenAIModelClient(t, env)
	name := fmt.Sprintf("hsf97_token_acct_%s", utils.RandStringBytes(6))

	auth := &OpenAIModelAuthentication{
		Type_: OpenAIModelAuthTypes.Token,
		Token: &OpenAIModelTokenSpec{TokenRef: tokenRef},
	}

	createResp := createOpenAIModelConnector(t, client, ctx, env.accountID, name, auth, "", "")
	defer deleteOpenAIModelConnector(t, client, ctx, env.accountID, name, "", "")

	require.NotNil(t, createResp.Data.Connector.OpenAIModel)
	require.Equal(t, OpenAIModelAuthTypes.Token, createResp.Data.Connector.OpenAIModel.Authentication.Type_)

	getResp := getOpenAIModelConnector(t, client, ctx, env.accountID, name, "", "")
	require.Equal(t, name, getResp.Data.Connector.Name)
	require.Equal(t, OpenAIModelAuthTypes.Token, getResp.Data.Connector.OpenAIModel.Authentication.Type_)
}

func TestAccOpenAIModelToken_OrgScope(t *testing.T) {
	env := getOpenAIModelTestEnv(t, "org")
	tokenRef := os.Getenv("OPENAI_MODEL_TOKEN_REF")
	if tokenRef == "" || env.orgID == "" {
		t.Skip("OPENAI_MODEL_TOKEN_REF and OPENAI_MODEL_ORG_SCOPE_ORG_ID must be set")
	}

	client, ctx := newOpenAIModelClient(t, env)
	name := fmt.Sprintf("hsf97_token_org_%s", utils.RandStringBytes(6))

	auth := &OpenAIModelAuthentication{
		Type_: OpenAIModelAuthTypes.Token,
		Token: &OpenAIModelTokenSpec{TokenRef: tokenRef},
	}

	createResp := createOpenAIModelConnector(t, client, ctx, env.accountID, name, auth, env.orgID, "")
	defer deleteOpenAIModelConnector(t, client, ctx, env.accountID, name, env.orgID, "")

	require.NotNil(t, createResp.Data.Connector.OpenAIModel)

	getResp := getOpenAIModelConnector(t, client, ctx, env.accountID, name, env.orgID, "")
	require.Equal(t, name, getResp.Data.Connector.Name)
}

func TestAccOpenAIModelToken_ProjectScope(t *testing.T) {
	env := getOpenAIModelTestEnv(t, "project")
	tokenRef := os.Getenv("OPENAI_MODEL_TOKEN_REF")
	if tokenRef == "" || env.orgID == "" || env.projectID == "" {
		t.Skip("OPENAI_MODEL_TOKEN_REF, OPENAI_MODEL_PROJECT_SCOPE_ORG_ID, and OPENAI_MODEL_PROJECT_SCOPE_PROJECT_ID must be set")
	}

	client, ctx := newOpenAIModelClient(t, env)
	name := fmt.Sprintf("hsf97_token_proj_%s", utils.RandStringBytes(6))

	auth := &OpenAIModelAuthentication{
		Type_: OpenAIModelAuthTypes.Token,
		Token: &OpenAIModelTokenSpec{TokenRef: tokenRef},
	}

	createResp := createOpenAIModelConnector(t, client, ctx, env.accountID, name, auth, env.orgID, env.projectID)
	defer deleteOpenAIModelConnector(t, client, ctx, env.accountID, name, env.orgID, env.projectID)

	require.NotNil(t, createResp.Data.Connector.OpenAIModel)

	getResp := getOpenAIModelConnector(t, client, ctx, env.accountID, name, env.orgID, env.projectID)
	require.Equal(t, name, getResp.Data.Connector.Name)
}

// ---------------------------------------------------------------------------
// Vertex Auth
// ---------------------------------------------------------------------------

func TestAccOpenAIModelVertex_AccountScope(t *testing.T) {
	env := getOpenAIModelTestEnv(t, "account")
	saKeyRef := os.Getenv("OPENAI_MODEL_VERTEX_SERVICE_ACCOUNT_KEY_REF")
	projectID := os.Getenv("OPENAI_MODEL_VERTEX_PROJECT_ID")
	region := os.Getenv("OPENAI_MODEL_VERTEX_REGION")
	if saKeyRef == "" || projectID == "" || region == "" {
		t.Skip("OPENAI_MODEL_VERTEX_SERVICE_ACCOUNT_KEY_REF, OPENAI_MODEL_VERTEX_PROJECT_ID, and OPENAI_MODEL_VERTEX_REGION must be set")
	}

	client, ctx := newOpenAIModelClient(t, env)
	name := fmt.Sprintf("hsf97_vertex_acct_%s", utils.RandStringBytes(6))

	auth := &OpenAIModelAuthentication{
		Type_: OpenAIModelAuthTypes.Vertex,
		Vertex: &OpenAIModelVertexSpec{
			ServiceAccountKeyRef: saKeyRef,
			ProjectId:            projectID,
			Region:               region,
		},
	}

	createResp := createOpenAIModelConnector(t, client, ctx, env.accountID, name, auth, "", "")
	defer deleteOpenAIModelConnector(t, client, ctx, env.accountID, name, "", "")

	require.NotNil(t, createResp.Data.Connector.OpenAIModel)
	require.Equal(t, OpenAIModelAuthTypes.Vertex, createResp.Data.Connector.OpenAIModel.Authentication.Type_)

	getResp := getOpenAIModelConnector(t, client, ctx, env.accountID, name, "", "")
	require.Equal(t, name, getResp.Data.Connector.Name)
	require.Equal(t, OpenAIModelAuthTypes.Vertex, getResp.Data.Connector.OpenAIModel.Authentication.Type_)
}

func TestAccOpenAIModelVertex_OrgScope(t *testing.T) {
	env := getOpenAIModelTestEnv(t, "org")
	saKeyRef := os.Getenv("OPENAI_MODEL_VERTEX_SERVICE_ACCOUNT_KEY_REF")
	projectID := os.Getenv("OPENAI_MODEL_VERTEX_PROJECT_ID")
	region := os.Getenv("OPENAI_MODEL_VERTEX_REGION")
	if saKeyRef == "" || projectID == "" || region == "" || env.orgID == "" {
		t.Skip("OPENAI_MODEL_VERTEX_SERVICE_ACCOUNT_KEY_REF, OPENAI_MODEL_VERTEX_PROJECT_ID, OPENAI_MODEL_VERTEX_REGION, and OPENAI_MODEL_ORG_SCOPE_ORG_ID must be set")
	}

	client, ctx := newOpenAIModelClient(t, env)
	name := fmt.Sprintf("hsf97_vertex_org_%s", utils.RandStringBytes(6))

	auth := &OpenAIModelAuthentication{
		Type_: OpenAIModelAuthTypes.Vertex,
		Vertex: &OpenAIModelVertexSpec{
			ServiceAccountKeyRef: saKeyRef,
			ProjectId:            projectID,
			Region:               region,
		},
	}

	createResp := createOpenAIModelConnector(t, client, ctx, env.accountID, name, auth, env.orgID, "")
	defer deleteOpenAIModelConnector(t, client, ctx, env.accountID, name, env.orgID, "")

	require.NotNil(t, createResp.Data.Connector.OpenAIModel)

	getResp := getOpenAIModelConnector(t, client, ctx, env.accountID, name, env.orgID, "")
	require.Equal(t, name, getResp.Data.Connector.Name)
}

func TestAccOpenAIModelVertex_ProjectScope(t *testing.T) {
	env := getOpenAIModelTestEnv(t, "project")
	saKeyRef := os.Getenv("OPENAI_MODEL_VERTEX_SERVICE_ACCOUNT_KEY_REF")
	projectID := os.Getenv("OPENAI_MODEL_VERTEX_PROJECT_ID")
	region := os.Getenv("OPENAI_MODEL_VERTEX_REGION")
	if saKeyRef == "" || projectID == "" || region == "" || env.orgID == "" || env.projectID == "" {
		t.Skip("OPENAI_MODEL_VERTEX_SERVICE_ACCOUNT_KEY_REF, OPENAI_MODEL_VERTEX_PROJECT_ID, OPENAI_MODEL_VERTEX_REGION, OPENAI_MODEL_PROJECT_SCOPE_ORG_ID, and OPENAI_MODEL_PROJECT_SCOPE_PROJECT_ID must be set")
	}

	client, ctx := newOpenAIModelClient(t, env)
	name := fmt.Sprintf("hsf97_vertex_proj_%s", utils.RandStringBytes(6))

	auth := &OpenAIModelAuthentication{
		Type_: OpenAIModelAuthTypes.Vertex,
		Vertex: &OpenAIModelVertexSpec{
			ServiceAccountKeyRef: saKeyRef,
			ProjectId:            projectID,
			Region:               region,
		},
	}

	createResp := createOpenAIModelConnector(t, client, ctx, env.accountID, name, auth, env.orgID, env.projectID)
	defer deleteOpenAIModelConnector(t, client, ctx, env.accountID, name, env.orgID, env.projectID)

	require.NotNil(t, createResp.Data.Connector.OpenAIModel)

	getResp := getOpenAIModelConnector(t, client, ctx, env.accountID, name, env.orgID, env.projectID)
	require.Equal(t, name, getResp.Data.Connector.Name)
}
