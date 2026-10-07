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
Anthropic Model Connector — Acceptance Tests (TestAccAnthropicModel*_*)
Tests Create, Get, and Delete of the Anthropic Model Connector against a live
Harness account. Covers three auth types (Token, BedrockApiKey, CloudProvider)
at three scopes (Account, Org, Project) = 9 tests total.

Each scope uses its own API key env var:
  - Account scope: HARNESS_PLATFORM_ACCOUNT_API_KEY
  - Org scope:     HARNESS_PLATFORM_ORG_API_KEY
  - Project scope: HARNESS_PLATFORM_PROJECT_API_KEY

Tests skip automatically if the required env vars are not set.

Connector fields set by tests:
  - Name, Identifier:      auto-generated with hsf98_ prefix + random suffix
  - Description:           "<connector-name> hsf98 testing only"
  - Tags:                  hsf98, testing
  - Url:                   https://api.anthropic.com (hardcoded)
  - Model:                 claude-opus-4-6 (hardcoded)
  - ExecuteOnDelegate:     true (hardcoded)
  - IgnoreTestConnection:  true (hardcoded)
  - DelegateSelectors:     from ANTHROPIC_MODEL_DELEGATE_SELECTORS (optional, comma-separated)

Optional env vars (apply to all scopes):

    # Comma-separated delegate selectors (optional, omit to use any delegate)
    export ANTHROPIC_MODEL_DELEGATE_SELECTORS="delegate01,delegate02"

    # Set to "true" to skip cleanup deletion (for manual UI verification)
    export IS_SKIP_DELETION_TEST="true"

---------------------------------------------------------------------------
# Account-Scope Tests
---------------------------------------------------------------------------

    export HARNESS_ACCOUNT_ID='<account-id>'
    export HARNESS_PLATFORM_ACCOUNT_API_KEY='<pat-or-sa-key>'
    # Token auth
    export ANTHROPIC_MODEL_TOKEN_REF="account.<account-level-secret>"
    # BedrockApiKey auth
    export ANTHROPIC_MODEL_BEDROCK_API_KEY_REF="account.<account-level-secret>"
    export ANTHROPIC_MODEL_BEDROCK_REGION="us-east-1"
    # CloudProvider auth
    export ANTHROPIC_MODEL_CLOUD_PROVIDER_TYPE="AWS"
    export ANTHROPIC_MODEL_CLOUD_CONNECTOR_REF="account.<account-level-cloud-connector>"
	ANTHROPIC_MODEL_DELEGATE_SELECTORS="delegate1,delegate2"

    go test -v -run "TestAccAnthropicModel.*_AccountScope" -count=1 -timeout 120s \
      ./harness/nextgen/ -args -account dummy -api-key dummy 2>&1 | tee /tmp/hsf-98.out

---------------------------------------------------------------------------
# Org-Scope Tests
---------------------------------------------------------------------------

    export HARNESS_ACCOUNT_ID='<account-id>'
    export HARNESS_PLATFORM_ORG_API_KEY='<pat-or-sa-key>'
    export ANTHROPIC_MODEL_ORG_SCOPE_ORG_ID="<org-id>"
    # Token auth
    export ANTHROPIC_MODEL_TOKEN_REF="org.<org-level-secret>"
    # BedrockApiKey auth
    export ANTHROPIC_MODEL_BEDROCK_API_KEY_REF="org.<org-level-secret>"
    export ANTHROPIC_MODEL_BEDROCK_REGION="us-east-1"
    # CloudProvider auth
    export ANTHROPIC_MODEL_CLOUD_PROVIDER_TYPE="AWS"
    export ANTHROPIC_MODEL_CLOUD_CONNECTOR_REF="org.<org-level-cloud-connector>"
	export ANTHROPIC_MODEL_DELEGATE_SELECTORS="delegate1,delegate2" # set to empty and check

    go test -v -run "TestAccAnthropicModel.*_OrgScope" -count=1 -timeout 120s \
      ./harness/nextgen/ -args -account dummy -api-key dummy 2>&1 | tee /tmp/hsf-98.out

---------------------------------------------------------------------------
# Project-Scope Tests
---------------------------------------------------------------------------

    export HARNESS_ACCOUNT_ID='<account-id>'
    export HARNESS_PLATFORM_PROJECT_API_KEY='<pat-or-sa-key>'
    export ANTHROPIC_MODEL_PROJECT_SCOPE_ORG_ID="<org-id>"
    export ANTHROPIC_MODEL_PROJECT_SCOPE_PROJECT_ID="<project-id>"
    # Token auth
    export ANTHROPIC_MODEL_TOKEN_REF="<project-level-secret>"
    # BedrockApiKey auth
    export ANTHROPIC_MODEL_BEDROCK_API_KEY_REF="<project-level-secret>"
    export ANTHROPIC_MODEL_BEDROCK_REGION="us-east-1"
    # CloudProvider auth
    export ANTHROPIC_MODEL_CLOUD_PROVIDER_TYPE="AWS"
    export ANTHROPIC_MODEL_CLOUD_CONNECTOR_REF="<project-level-cloud-connector>"
	export ANTHROPIC_MODEL_DELEGATE_SELECTORS="delegate1,delegate2"

    go test -v -run "TestAccAnthropicModel.*_ProjectScope" -count=1 -timeout 120s \
      ./harness/nextgen/ -args -account dummy -api-key dummy

*/

type anthropicModelTestEnv struct {
	accountID string
	apiKey    string
	orgID     string
	projectID string
}

func getAnthropicModelTestEnv(t *testing.T, scope string) anthropicModelTestEnv {
	t.Helper()

	var apiKeyEnvVar, orgEnvVar, projectEnvVar string
	switch scope {
	case "account":
		apiKeyEnvVar = "HARNESS_PLATFORM_ACCOUNT_API_KEY"
	case "org":
		apiKeyEnvVar = "HARNESS_PLATFORM_ORG_API_KEY"
		orgEnvVar = "ANTHROPIC_MODEL_ORG_SCOPE_ORG_ID"
	case "project":
		apiKeyEnvVar = "HARNESS_PLATFORM_PROJECT_API_KEY"
		orgEnvVar = "ANTHROPIC_MODEL_PROJECT_SCOPE_ORG_ID"
		projectEnvVar = "ANTHROPIC_MODEL_PROJECT_SCOPE_PROJECT_ID"
	default:
		t.Fatalf("unknown scope %q: must be account, org, or project", scope)
	}

	env := anthropicModelTestEnv{
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

func newAnthropicModelClient(t *testing.T, env anthropicModelTestEnv) (*APIClient, context.Context) {
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

func createAnthropicModelConnector(
	t *testing.T,
	client *APIClient,
	ctx context.Context,
	accountID string,
	name string,
	auth *AnthropicModelAuthentication,
	orgID string,
	projectID string,
) ResponseDtoConnectorResponse {
	t.Helper()

	var delegateSelectors []string
	if ds := os.Getenv("ANTHROPIC_MODEL_DELEGATE_SELECTORS"); ds != "" {
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
		Description:       fmt.Sprintf("%s hsf98 testing only", name),
		OrgIdentifier:     orgID,
		ProjectIdentifier: projectID,
		Tags:              map[string]string{"hsf98": "true", "testing": "true"},
		Type_:             ConnectorTypes.AnthropicModel,
		AnthropicModel: &AnthropicModelConnector{
			Url:                  "https://api.anthropic.com",
			Model:                "claude-opus-4-6",
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
	require.Equal(t, ConnectorTypes.AnthropicModel, resp.Data.Connector.Type_)

	t.Logf("Created connector: name=%s scope=account:%s/org:%s/project:%s",
		name, accountID, orgID, projectID)

	return resp
}

func deleteAnthropicModelConnector(
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

func getAnthropicModelConnector(
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

func TestAccAnthropicModelToken_AccountScope(t *testing.T) {
	env := getAnthropicModelTestEnv(t, "account")
	tokenRef := os.Getenv("ANTHROPIC_MODEL_TOKEN_REF")
	if tokenRef == "" {
		t.Skip("ANTHROPIC_MODEL_TOKEN_REF must be set")
	}

	client, ctx := newAnthropicModelClient(t, env)
	name := fmt.Sprintf("hsf98_token_acct_%s", utils.RandStringBytes(6))

	auth := &AnthropicModelAuthentication{
		Type_: AnthropicModelAuthTypes.Token,
		Token: &AnthropicModelTokenSpec{TokenRef: tokenRef},
	}

	createResp := createAnthropicModelConnector(t, client, ctx, env.accountID, name, auth, "", "")
	defer deleteAnthropicModelConnector(t, client, ctx, env.accountID, name, "", "")

	require.NotNil(t, createResp.Data.Connector.AnthropicModel)
	require.Equal(t, AnthropicModelAuthTypes.Token, createResp.Data.Connector.AnthropicModel.Authentication.Type_)

	getResp := getAnthropicModelConnector(t, client, ctx, env.accountID, name, "", "")
	require.Equal(t, name, getResp.Data.Connector.Name)
	require.Equal(t, AnthropicModelAuthTypes.Token, getResp.Data.Connector.AnthropicModel.Authentication.Type_)
}

func TestAccAnthropicModelToken_OrgScope(t *testing.T) {
	env := getAnthropicModelTestEnv(t, "org")
	tokenRef := os.Getenv("ANTHROPIC_MODEL_TOKEN_REF")
	if tokenRef == "" || env.orgID == "" {
		t.Skip("ANTHROPIC_MODEL_TOKEN_REF and ANTHROPIC_MODEL_ORG_SCOPE_ORG_ID must be set")
	}

	client, ctx := newAnthropicModelClient(t, env)
	name := fmt.Sprintf("hsf98_token_org_%s", utils.RandStringBytes(6))

	auth := &AnthropicModelAuthentication{
		Type_: AnthropicModelAuthTypes.Token,
		Token: &AnthropicModelTokenSpec{TokenRef: tokenRef},
	}

	createResp := createAnthropicModelConnector(t, client, ctx, env.accountID, name, auth, env.orgID, "")
	defer deleteAnthropicModelConnector(t, client, ctx, env.accountID, name, env.orgID, "")

	require.NotNil(t, createResp.Data.Connector.AnthropicModel)
	require.Equal(t, AnthropicModelAuthTypes.Token, createResp.Data.Connector.AnthropicModel.Authentication.Type_)

	getResp := getAnthropicModelConnector(t, client, ctx, env.accountID, name, env.orgID, "")
	require.Equal(t, name, getResp.Data.Connector.Name)
}

func TestAccAnthropicModelToken_ProjectScope(t *testing.T) {
	env := getAnthropicModelTestEnv(t, "project")
	tokenRef := os.Getenv("ANTHROPIC_MODEL_TOKEN_REF")
	if tokenRef == "" || env.orgID == "" || env.projectID == "" {
		t.Skip("ANTHROPIC_MODEL_TOKEN_REF, ANTHROPIC_MODEL_PROJECT_SCOPE_ORG_ID, and ANTHROPIC_MODEL_PROJECT_SCOPE_PROJECT_ID must be set")
	}

	client, ctx := newAnthropicModelClient(t, env)
	name := fmt.Sprintf("hsf98_token_proj_%s", utils.RandStringBytes(6))

	auth := &AnthropicModelAuthentication{
		Type_: AnthropicModelAuthTypes.Token,
		Token: &AnthropicModelTokenSpec{TokenRef: tokenRef},
	}

	createResp := createAnthropicModelConnector(t, client, ctx, env.accountID, name, auth, env.orgID, env.projectID)
	defer deleteAnthropicModelConnector(t, client, ctx, env.accountID, name, env.orgID, env.projectID)

	require.NotNil(t, createResp.Data.Connector.AnthropicModel)

	getResp := getAnthropicModelConnector(t, client, ctx, env.accountID, name, env.orgID, env.projectID)
	require.Equal(t, name, getResp.Data.Connector.Name)
}

// ---------------------------------------------------------------------------
// BedrockApiKey Auth
// ---------------------------------------------------------------------------

func TestAccAnthropicModelBedrock_AccountScope(t *testing.T) {
	env := getAnthropicModelTestEnv(t, "account")
	apiKeyRef := os.Getenv("ANTHROPIC_MODEL_BEDROCK_API_KEY_REF")
	region := os.Getenv("ANTHROPIC_MODEL_BEDROCK_REGION")
	if apiKeyRef == "" || region == "" {
		t.Skip("ANTHROPIC_MODEL_BEDROCK_API_KEY_REF and ANTHROPIC_MODEL_BEDROCK_REGION must be set")
	}

	client, ctx := newAnthropicModelClient(t, env)
	name := fmt.Sprintf("hsf98_bedrock_acct_%s", utils.RandStringBytes(6))

	auth := &AnthropicModelAuthentication{
		Type_: AnthropicModelAuthTypes.BedrockApiKey,
		BedrockApiKey: &AnthropicModelBedrockApiKeySpec{
			ApiKeyRef: apiKeyRef,
			Region:    region,
		},
	}

	createResp := createAnthropicModelConnector(t, client, ctx, env.accountID, name, auth, "", "")
	defer deleteAnthropicModelConnector(t, client, ctx, env.accountID, name, "", "")

	require.NotNil(t, createResp.Data.Connector.AnthropicModel)
	require.Equal(t, AnthropicModelAuthTypes.BedrockApiKey, createResp.Data.Connector.AnthropicModel.Authentication.Type_)

	getResp := getAnthropicModelConnector(t, client, ctx, env.accountID, name, "", "")
	require.Equal(t, name, getResp.Data.Connector.Name)
	require.Equal(t, AnthropicModelAuthTypes.BedrockApiKey, getResp.Data.Connector.AnthropicModel.Authentication.Type_)
}

func TestAccAnthropicModelBedrock_OrgScope(t *testing.T) {
	env := getAnthropicModelTestEnv(t, "org")
	apiKeyRef := os.Getenv("ANTHROPIC_MODEL_BEDROCK_API_KEY_REF")
	region := os.Getenv("ANTHROPIC_MODEL_BEDROCK_REGION")
	if apiKeyRef == "" || region == "" || env.orgID == "" {
		t.Skip("ANTHROPIC_MODEL_BEDROCK_API_KEY_REF, ANTHROPIC_MODEL_BEDROCK_REGION, and ANTHROPIC_MODEL_ORG_SCOPE_ORG_ID must be set")
	}

	client, ctx := newAnthropicModelClient(t, env)
	name := fmt.Sprintf("hsf98_bedrock_org_%s", utils.RandStringBytes(6))

	auth := &AnthropicModelAuthentication{
		Type_: AnthropicModelAuthTypes.BedrockApiKey,
		BedrockApiKey: &AnthropicModelBedrockApiKeySpec{
			ApiKeyRef: apiKeyRef,
			Region:    region,
		},
	}

	createResp := createAnthropicModelConnector(t, client, ctx, env.accountID, name, auth, env.orgID, "")
	defer deleteAnthropicModelConnector(t, client, ctx, env.accountID, name, env.orgID, "")

	require.NotNil(t, createResp.Data.Connector.AnthropicModel)

	getResp := getAnthropicModelConnector(t, client, ctx, env.accountID, name, env.orgID, "")
	require.Equal(t, name, getResp.Data.Connector.Name)
}

func TestAccAnthropicModelBedrock_ProjectScope(t *testing.T) {
	env := getAnthropicModelTestEnv(t, "project")
	apiKeyRef := os.Getenv("ANTHROPIC_MODEL_BEDROCK_API_KEY_REF")
	region := os.Getenv("ANTHROPIC_MODEL_BEDROCK_REGION")
	if apiKeyRef == "" || region == "" || env.orgID == "" || env.projectID == "" {
		t.Skip("ANTHROPIC_MODEL_BEDROCK_API_KEY_REF, ANTHROPIC_MODEL_BEDROCK_REGION, ANTHROPIC_MODEL_PROJECT_SCOPE_ORG_ID, and ANTHROPIC_MODEL_PROJECT_SCOPE_PROJECT_ID must be set")
	}

	client, ctx := newAnthropicModelClient(t, env)
	name := fmt.Sprintf("hsf98_bedrock_proj_%s", utils.RandStringBytes(6))

	auth := &AnthropicModelAuthentication{
		Type_: AnthropicModelAuthTypes.BedrockApiKey,
		BedrockApiKey: &AnthropicModelBedrockApiKeySpec{
			ApiKeyRef: apiKeyRef,
			Region:    region,
		},
	}

	createResp := createAnthropicModelConnector(t, client, ctx, env.accountID, name, auth, env.orgID, env.projectID)
	defer deleteAnthropicModelConnector(t, client, ctx, env.accountID, name, env.orgID, env.projectID)

	require.NotNil(t, createResp.Data.Connector.AnthropicModel)

	getResp := getAnthropicModelConnector(t, client, ctx, env.accountID, name, env.orgID, env.projectID)
	require.Equal(t, name, getResp.Data.Connector.Name)
}

// ---------------------------------------------------------------------------
// CloudProvider Auth (AWS)
// ---------------------------------------------------------------------------

func TestAccAnthropicModelCloudProvider_AccountScope(t *testing.T) {
	env := getAnthropicModelTestEnv(t, "account")
	cpType := os.Getenv("ANTHROPIC_MODEL_CLOUD_PROVIDER_TYPE")
	connRef := os.Getenv("ANTHROPIC_MODEL_CLOUD_CONNECTOR_REF")
	if cpType == "" || connRef == "" {
		t.Skip("ANTHROPIC_MODEL_CLOUD_PROVIDER_TYPE and ANTHROPIC_MODEL_CLOUD_CONNECTOR_REF must be set")
	}

	client, ctx := newAnthropicModelClient(t, env)
	name := fmt.Sprintf("hsf98_cp_acct_%s", utils.RandStringBytes(6))

	auth := &AnthropicModelAuthentication{
		Type_: AnthropicModelAuthTypes.CloudProvider,
		CloudProvider: &AnthropicModelCloudProviderSpec{
			Type_:        AnthropicModelCloudProviderType(cpType),
			ConnectorRef: connRef,
		},
	}

	createResp := createAnthropicModelConnector(t, client, ctx, env.accountID, name, auth, "", "")
	defer deleteAnthropicModelConnector(t, client, ctx, env.accountID, name, "", "")

	require.NotNil(t, createResp.Data.Connector.AnthropicModel)
	require.Equal(t, AnthropicModelAuthTypes.CloudProvider, createResp.Data.Connector.AnthropicModel.Authentication.Type_)

	getResp := getAnthropicModelConnector(t, client, ctx, env.accountID, name, "", "")
	require.Equal(t, name, getResp.Data.Connector.Name)
	require.Equal(t, AnthropicModelAuthTypes.CloudProvider, getResp.Data.Connector.AnthropicModel.Authentication.Type_)
}

func TestAccAnthropicModelCloudProvider_OrgScope(t *testing.T) {
	env := getAnthropicModelTestEnv(t, "org")
	cpType := os.Getenv("ANTHROPIC_MODEL_CLOUD_PROVIDER_TYPE")
	connRef := os.Getenv("ANTHROPIC_MODEL_CLOUD_CONNECTOR_REF")
	if cpType == "" || connRef == "" || env.orgID == "" {
		t.Skip("ANTHROPIC_MODEL_CLOUD_PROVIDER_TYPE, ANTHROPIC_MODEL_CLOUD_CONNECTOR_REF, and ANTHROPIC_MODEL_ORG_SCOPE_ORG_ID must be set")
	}

	client, ctx := newAnthropicModelClient(t, env)
	name := fmt.Sprintf("hsf98_cp_org_%s", utils.RandStringBytes(6))

	auth := &AnthropicModelAuthentication{
		Type_: AnthropicModelAuthTypes.CloudProvider,
		CloudProvider: &AnthropicModelCloudProviderSpec{
			Type_:        AnthropicModelCloudProviderType(cpType),
			ConnectorRef: connRef,
		},
	}

	createResp := createAnthropicModelConnector(t, client, ctx, env.accountID, name, auth, env.orgID, "")
	defer deleteAnthropicModelConnector(t, client, ctx, env.accountID, name, env.orgID, "")

	require.NotNil(t, createResp.Data.Connector.AnthropicModel)

	getResp := getAnthropicModelConnector(t, client, ctx, env.accountID, name, env.orgID, "")
	require.Equal(t, name, getResp.Data.Connector.Name)
}

func TestAccAnthropicModelCloudProvider_ProjectScope(t *testing.T) {
	env := getAnthropicModelTestEnv(t, "project")
	cpType := os.Getenv("ANTHROPIC_MODEL_CLOUD_PROVIDER_TYPE")
	connRef := os.Getenv("ANTHROPIC_MODEL_CLOUD_CONNECTOR_REF")
	if cpType == "" || connRef == "" || env.orgID == "" || env.projectID == "" {
		t.Skip("ANTHROPIC_MODEL_CLOUD_PROVIDER_TYPE, ANTHROPIC_MODEL_CLOUD_CONNECTOR_REF, ANTHROPIC_MODEL_PROJECT_SCOPE_ORG_ID, and ANTHROPIC_MODEL_PROJECT_SCOPE_PROJECT_ID must be set")
	}

	client, ctx := newAnthropicModelClient(t, env)
	name := fmt.Sprintf("hsf98_cp_proj_%s", utils.RandStringBytes(6))

	auth := &AnthropicModelAuthentication{
		Type_: AnthropicModelAuthTypes.CloudProvider,
		CloudProvider: &AnthropicModelCloudProviderSpec{
			Type_:        AnthropicModelCloudProviderType(cpType),
			ConnectorRef: connRef,
		},
	}

	createResp := createAnthropicModelConnector(t, client, ctx, env.accountID, name, auth, env.orgID, env.projectID)
	defer deleteAnthropicModelConnector(t, client, ctx, env.accountID, name, env.orgID, env.projectID)

	require.NotNil(t, createResp.Data.Connector.AnthropicModel)

	getResp := getAnthropicModelConnector(t, client, ctx, env.accountID, name, env.orgID, env.projectID)
	require.Equal(t, name, getResp.Data.Connector.Name)
}
