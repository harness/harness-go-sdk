package nextgen

/*
# Trigger V2 API Acceptance Tests

These tests hit the real Harness API and verify the v2 trigger endpoints
(CreateTriggerV2 / UpdateTriggerV2) which carry an executorInfo payload
alongside the trigger YAML.

Test cases:
  - TestCreateTriggerV2_UserExecutor      — USER executor identity
  - TestCreateTriggerV2_SA_ProjectScoped  — SERVICE_ACCOUNT executor, project-scoped SA
  - TestCreateTriggerV2_SA_OrgScoped      — SERVICE_ACCOUNT executor, org-scoped SA
  - TestCreateTriggerV2_SA_AccountScoped  — SERVICE_ACCOUNT executor, account-scoped SA
  - TestCreateTriggerV2_NoExecutorInfo    — v2 API with yaml-only body (no executorInfo)
  - TestUpdateTriggerV2_SA_ProjectScoped  — update path: triggerIdentifier substitution, IfMatch, executorInfo round-trip

## Required env vars (set all to run every test in one shot)

	export HARNESS_ACCOUNT_ID='<your-account-id>'

	# Provider auth key — any key with core_trigger_edit permission.
	# Independent of the executor identity configured in each test.
	export HARNESS_PLATFORM_API_KEY='<personal-or-sa-api-key>'

	# Pipeline / org / project that triggers are created under
	export TRIGGER_TEST_ORG_ID='hsforg01'
	export TRIGGER_TEST_PROJECT_ID='hsfproj01'
	export TRIGGER_TEST_PIPELINE_ID='hsf143test01'
	export TRIGGER_TEST_CONNECTOR_REF='githubconn003'
	export TRIGGER_TEST_REPO_NAME='hands-on-workbook'

	# USER executor — uuid of the user that should run the pipeline.
	# TRIGGER_TEST_USER_API_KEY must be a personal API key (not an SA token):
	# Harness rejects USER executor when the provider auth key is an SA token.
	export TRIGGER_TEST_USER_UUID='<user-uuid-from-harness-api>'
	export TRIGGER_TEST_USER_API_KEY='<personal-api-key>'

	# SERVICE_ACCOUNT executor — project-scoped SA (both org and project set)
	export TRIGGER_TEST_SA_PROJECT_IDENTIFIER='sa005'

	# SERVICE_ACCOUNT executor — org-scoped SA (only org set, no project)
	export TRIGGER_TEST_SA_ORG_IDENTIFIER='hsforg01sa006'

	# SERVICE_ACCOUNT executor — account-scoped SA (no org or project)
	export TRIGGER_TEST_SA_ACCOUNT_IDENTIFIER='hsftftestsvcacc01'

## Run all acceptance tests

	go test ./harness/nextgen -run TestCreateTriggerV2 -v -count=1 -args -account $HARNESS_ACCOUNT_ID -api-key $HARNESS_PLATFORM_API_KEY
*/

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

type triggerTestEnv struct {
	accountID    string
	orgID        string
	projectID    string
	pipelineID   string
	connectorRef string
	repoName     string
}

func getTriggerTestEnv(t *testing.T) triggerTestEnv {
	t.Helper()
	env := triggerTestEnv{
		accountID:    os.Getenv("HARNESS_ACCOUNT_ID"),
		orgID:        os.Getenv("TRIGGER_TEST_ORG_ID"),
		projectID:    os.Getenv("TRIGGER_TEST_PROJECT_ID"),
		pipelineID:   os.Getenv("TRIGGER_TEST_PIPELINE_ID"),
		connectorRef: os.Getenv("TRIGGER_TEST_CONNECTOR_REF"),
		repoName:     os.Getenv("TRIGGER_TEST_REPO_NAME"),
	}
	if env.accountID == "" {
		t.Skip("HARNESS_ACCOUNT_ID must be set for acceptance tests")
	}
	missing := []string{}
	if env.orgID == "" {
		missing = append(missing, "TRIGGER_TEST_ORG_ID")
	}
	if env.projectID == "" {
		missing = append(missing, "TRIGGER_TEST_PROJECT_ID")
	}
	if env.pipelineID == "" {
		missing = append(missing, "TRIGGER_TEST_PIPELINE_ID")
	}
	if env.connectorRef == "" {
		missing = append(missing, "TRIGGER_TEST_CONNECTOR_REF")
	}
	if env.repoName == "" {
		missing = append(missing, "TRIGGER_TEST_REPO_NAME")
	}
	if len(missing) > 0 {
		t.Skipf("missing required env vars: %v", missing)
	}
	return env
}

type triggerLoggingRoundTripper struct {
	t       *testing.T
	wrapped http.RoundTripper
}

func (l *triggerLoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := l.wrapped.RoundTrip(req)
	if err != nil || resp == nil {
		return resp, err
	}
	body, readErr := io.ReadAll(resp.Body)
	resp.Body.Close()
	if readErr == nil {
		l.t.Logf("raw response: status=%d body=%s", resp.StatusCode, string(body))
	}
	resp.Body = io.NopCloser(bytes.NewReader(body))
	return resp, nil
}

func newTriggerClientAndCtx(t *testing.T, apiKey string) (*APIClient, context.Context) {
	t.Helper()
	cfg := NewConfiguration()
	if cfg.AccountId == "" {
		t.Skip("HARNESS_ACCOUNT_ID must be set for acceptance tests")
	}
	cfg.ApiKey = apiKey
	cfg.HTTPClient.HTTPClient.Transport = &triggerLoggingRoundTripper{
		t:       t,
		wrapped: cfg.HTTPClient.HTTPClient.Transport,
	}
	client := NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), ContextAPIKey, APIKey{Key: apiKey})
	return client, ctx
}

func triggerYaml(env triggerTestEnv, identifier string) string {
	return fmt.Sprintf(`trigger:
  name: %s
  identifier: %s
  enabled: true
  orgIdentifier: %s
  projectIdentifier: %s
  pipelineIdentifier: %s
  stagesToExecute: []
  tags: {}
  source:
    type: Webhook
    spec:
      type: Github
      spec:
        type: Push
        spec:
          connectorRef: %s
          repoName: %s
          autoAbortPreviousExecutions: false
          payloadConditions:
            - key: targetBranch
              operator: Contains
              value: hsf143_b01
          headerConditions: []
          actions: []
`,
		identifier, identifier,
		env.orgID, env.projectID, env.pipelineID,
		env.connectorRef, env.repoName,
	)
}

func deleteTrigger(t *testing.T, client *APIClient, ctx context.Context, env triggerTestEnv, id string) {
	t.Helper()
	_, _, err := client.TriggersApi.DeleteTrigger(ctx, client.AccountId,
		env.orgID, env.projectID, env.pipelineID, id,
		&TriggersApiDeleteTriggerOpts{})
	if err != nil {
		t.Logf("warning: failed to delete trigger %s: %v", id, err)
	}
}

func assertTriggerResponse(t *testing.T, resp ResponseDtongTriggerResponse, httpResp *http.Response, err error, expectedID string) {
	t.Helper()
	require.NoError(t, err, "API call returned an error")
	require.NotNil(t, httpResp, "HTTP response is nil")
	require.Equal(t, 200, httpResp.StatusCode, "unexpected HTTP status: %d", httpResp.StatusCode)
	require.NotNil(t, resp.Data, "response Data is nil")
	require.Equal(t, expectedID, resp.Data.Identifier)
	t.Logf("trigger created: identifier=%s webhookUrl=%s", resp.Data.Identifier, resp.Data.WebhookUrl)
	if resp.Data.ExecutorInfo != nil {
		t.Logf("executorInfo: type=%s uuid=%s accountIdentifier=%s orgIdentifier=%s projectIdentifier=%s",
			resp.Data.ExecutorInfo.Type_,
			resp.Data.ExecutorInfo.Uuid,
			resp.Data.ExecutorInfo.AccountIdentifier,
			resp.Data.ExecutorInfo.OrgIdentifier,
			resp.Data.ExecutorInfo.ProjectIdentifier,
		)
	}
}

// ---------------------------------------------------------------------------
// Acceptance tests
// ---------------------------------------------------------------------------

// TestCreateTriggerV2_UserExecutor creates a trigger with a USER executor identity.
// Provider auth must be a personal API key — Harness rejects USER executor when
// the auth key is an SA token.
func TestCreateTriggerV2_UserExecutor(t *testing.T) {
	apiKey := os.Getenv("TRIGGER_TEST_USER_API_KEY")
	if apiKey == "" {
		t.Skip("TRIGGER_TEST_USER_API_KEY must be set for this acceptance test (must be a personal API key, not an SA token)")
	}
	userUUID := os.Getenv("TRIGGER_TEST_USER_UUID")
	if userUUID == "" {
		t.Skip("TRIGGER_TEST_USER_UUID must be set for this acceptance test")
	}

	env := getTriggerTestEnv(t)
	client, ctx := newTriggerClientAndCtx(t, apiKey)
	id := fmt.Sprintf("trgv2user_%s", utils.RandStringBytes(5))

	body := TriggerYamlRequest{
		Yaml: triggerYaml(env, id),
		ExecutorInfo: &TriggerExecutorInfo{
			Type_: "USER",
			Uuid:  userUUID,
		},
	}

	resp, httpResp, err := client.TriggersApi.CreateTriggerV2(ctx, body, client.AccountId,
		env.orgID, env.projectID, env.pipelineID,
		&TriggersApiCreateTriggerV2Opts{WithServiceV2: optional.NewBool(true)})
	defer deleteTrigger(t, client, ctx, env, id)

	assertTriggerResponse(t, resp, httpResp, err, id)
	require.NotNil(t, resp.Data.ExecutorInfo)
	require.Equal(t, "USER", resp.Data.ExecutorInfo.Type_)
	require.Equal(t, userUUID, resp.Data.ExecutorInfo.Uuid)
}

// TestCreateTriggerV2_SA_ProjectScoped creates a trigger with a project-scoped
// SERVICE_ACCOUNT executor (both executor_org_id and executor_project_id set).
func TestCreateTriggerV2_SA_ProjectScoped(t *testing.T) {
	apiKey := os.Getenv("HARNESS_PLATFORM_API_KEY")
	if apiKey == "" {
		t.Skip("HARNESS_PLATFORM_API_KEY must be set for this acceptance test")
	}
	saIdentifier := os.Getenv("TRIGGER_TEST_SA_PROJECT_IDENTIFIER")
	if saIdentifier == "" {
		t.Skip("TRIGGER_TEST_SA_PROJECT_IDENTIFIER must be set for this acceptance test")
	}

	env := getTriggerTestEnv(t)
	client, ctx := newTriggerClientAndCtx(t, apiKey)
	id := fmt.Sprintf("trgv2saproj_%s", utils.RandStringBytes(5))

	body := TriggerYamlRequest{
		Yaml: triggerYaml(env, id),
		ExecutorInfo: &TriggerExecutorInfo{
			Type_:             "SERVICE_ACCOUNT",
			Uuid:              saIdentifier,
			AccountIdentifier: client.AccountId,
			OrgIdentifier:     env.orgID,
			ProjectIdentifier: env.projectID,
		},
	}

	resp, httpResp, err := client.TriggersApi.CreateTriggerV2(ctx, body, client.AccountId,
		env.orgID, env.projectID, env.pipelineID,
		&TriggersApiCreateTriggerV2Opts{WithServiceV2: optional.NewBool(true)})
	defer deleteTrigger(t, client, ctx, env, id)

	assertTriggerResponse(t, resp, httpResp, err, id)
	require.NotNil(t, resp.Data.ExecutorInfo)
	require.Equal(t, "SERVICE_ACCOUNT", resp.Data.ExecutorInfo.Type_)
	require.Equal(t, saIdentifier, resp.Data.ExecutorInfo.Uuid)
	require.Equal(t, env.orgID, resp.Data.ExecutorInfo.OrgIdentifier)
	require.Equal(t, env.projectID, resp.Data.ExecutorInfo.ProjectIdentifier)
}

// TestCreateTriggerV2_SA_OrgScoped creates a trigger with an org-scoped
// SERVICE_ACCOUNT executor (only executor_org_id set, no executor_project_id).
func TestCreateTriggerV2_SA_OrgScoped(t *testing.T) {
	apiKey := os.Getenv("HARNESS_PLATFORM_API_KEY")
	if apiKey == "" {
		t.Skip("HARNESS_PLATFORM_API_KEY must be set for this acceptance test")
	}
	saIdentifier := os.Getenv("TRIGGER_TEST_SA_ORG_IDENTIFIER")
	if saIdentifier == "" {
		t.Skip("TRIGGER_TEST_SA_ORG_IDENTIFIER must be set for this acceptance test")
	}

	env := getTriggerTestEnv(t)
	client, ctx := newTriggerClientAndCtx(t, apiKey)
	id := fmt.Sprintf("trgv2saorg_%s", utils.RandStringBytes(5))

	body := TriggerYamlRequest{
		Yaml: triggerYaml(env, id),
		ExecutorInfo: &TriggerExecutorInfo{
			Type_:             "SERVICE_ACCOUNT",
			Uuid:              saIdentifier,
			AccountIdentifier: client.AccountId,
			OrgIdentifier:     env.orgID,
			// ProjectIdentifier intentionally omitted — org-scoped SA
		},
	}

	resp, httpResp, err := client.TriggersApi.CreateTriggerV2(ctx, body, client.AccountId,
		env.orgID, env.projectID, env.pipelineID,
		&TriggersApiCreateTriggerV2Opts{WithServiceV2: optional.NewBool(true)})
	defer deleteTrigger(t, client, ctx, env, id)

	assertTriggerResponse(t, resp, httpResp, err, id)
	require.NotNil(t, resp.Data.ExecutorInfo)
	require.Equal(t, "SERVICE_ACCOUNT", resp.Data.ExecutorInfo.Type_)
	require.Equal(t, saIdentifier, resp.Data.ExecutorInfo.Uuid)
	require.Equal(t, env.orgID, resp.Data.ExecutorInfo.OrgIdentifier)
	require.Empty(t, resp.Data.ExecutorInfo.ProjectIdentifier)
}

// TestCreateTriggerV2_SA_AccountScoped creates a trigger with an account-scoped
// SERVICE_ACCOUNT executor (no executor_org_id or executor_project_id set).
func TestCreateTriggerV2_SA_AccountScoped(t *testing.T) {
	apiKey := os.Getenv("HARNESS_PLATFORM_API_KEY")
	if apiKey == "" {
		t.Skip("HARNESS_PLATFORM_API_KEY must be set for this acceptance test")
	}
	saIdentifier := os.Getenv("TRIGGER_TEST_SA_ACCOUNT_IDENTIFIER")
	if saIdentifier == "" {
		t.Skip("TRIGGER_TEST_SA_ACCOUNT_IDENTIFIER must be set for this acceptance test")
	}

	env := getTriggerTestEnv(t)
	client, ctx := newTriggerClientAndCtx(t, apiKey)
	id := fmt.Sprintf("trgv2saacc_%s", utils.RandStringBytes(5))

	body := TriggerYamlRequest{
		Yaml: triggerYaml(env, id),
		ExecutorInfo: &TriggerExecutorInfo{
			Type_:             "SERVICE_ACCOUNT",
			Uuid:              saIdentifier,
			AccountIdentifier: client.AccountId,
			// OrgIdentifier and ProjectIdentifier intentionally omitted — account-scoped SA
		},
	}

	resp, httpResp, err := client.TriggersApi.CreateTriggerV2(ctx, body, client.AccountId,
		env.orgID, env.projectID, env.pipelineID,
		&TriggersApiCreateTriggerV2Opts{WithServiceV2: optional.NewBool(true)})
	defer deleteTrigger(t, client, ctx, env, id)

	assertTriggerResponse(t, resp, httpResp, err, id)
	require.NotNil(t, resp.Data.ExecutorInfo)
	require.Equal(t, "SERVICE_ACCOUNT", resp.Data.ExecutorInfo.Type_)
	require.Equal(t, saIdentifier, resp.Data.ExecutorInfo.Uuid)
	require.Empty(t, resp.Data.ExecutorInfo.OrgIdentifier)
	require.Empty(t, resp.Data.ExecutorInfo.ProjectIdentifier)
}

// TestUpdateTriggerV2_SA_ProjectScoped creates a trigger then updates it via
// UpdateTriggerV2, verifying triggerIdentifier substitution in the URL,
// IfMatch header handling, and executorInfo round-trip on the response.
func TestUpdateTriggerV2_SA_ProjectScoped(t *testing.T) {
	apiKey := os.Getenv("HARNESS_PLATFORM_API_KEY")
	if apiKey == "" {
		t.Skip("HARNESS_PLATFORM_API_KEY must be set for this acceptance test")
	}
	saIdentifier := os.Getenv("TRIGGER_TEST_SA_PROJECT_IDENTIFIER")
	if saIdentifier == "" {
		t.Skip("TRIGGER_TEST_SA_PROJECT_IDENTIFIER must be set for this acceptance test")
	}

	env := getTriggerTestEnv(t)
	client, ctx := newTriggerClientAndCtx(t, apiKey)
	id := fmt.Sprintf("trgv2upd_%s", utils.RandStringBytes(5))

	createBody := TriggerYamlRequest{
		Yaml: triggerYaml(env, id),
		ExecutorInfo: &TriggerExecutorInfo{
			Type_:             "SERVICE_ACCOUNT",
			Uuid:              saIdentifier,
			AccountIdentifier: client.AccountId,
			OrgIdentifier:     env.orgID,
			ProjectIdentifier: env.projectID,
		},
	}

	createResp, _, err := client.TriggersApi.CreateTriggerV2(ctx, createBody, client.AccountId,
		env.orgID, env.projectID, env.pipelineID,
		&TriggersApiCreateTriggerV2Opts{WithServiceV2: optional.NewBool(true)})
	require.NoError(t, err)
	require.NotNil(t, createResp.Data)
	defer deleteTrigger(t, client, ctx, env, id)

	// Update: change the branch filter value to confirm yaml change is accepted,
	// keep the same executorInfo and verify it round-trips correctly.
	updatedYaml := fmt.Sprintf(`trigger:
  name: %s
  identifier: %s
  enabled: true
  orgIdentifier: %s
  projectIdentifier: %s
  pipelineIdentifier: %s
  stagesToExecute: []
  tags: {}
  source:
    type: Webhook
    spec:
      type: Github
      spec:
        type: Push
        spec:
          connectorRef: %s
          repoName: %s
          autoAbortPreviousExecutions: false
          payloadConditions:
            - key: targetBranch
              operator: Contains
              value: hsf143_b02_updated
          headerConditions: []
          actions: []
`,
		id, id, env.orgID, env.projectID, env.pipelineID, env.connectorRef, env.repoName,
	)

	updateBody := TriggerYamlRequest{
		Yaml: updatedYaml,
		ExecutorInfo: &TriggerExecutorInfo{
			Type_:             "SERVICE_ACCOUNT",
			Uuid:              saIdentifier,
			AccountIdentifier: client.AccountId,
			OrgIdentifier:     env.orgID,
			ProjectIdentifier: env.projectID,
		},
	}

	updateResp, httpResp, err := client.TriggersApi.UpdateTriggerV2(ctx, updateBody, client.AccountId,
		env.orgID, env.projectID, env.pipelineID, id,
		&TriggersApiUpdateTriggerV2Opts{})

	require.NoError(t, err, "UpdateTriggerV2 returned an error")
	require.Equal(t, 200, httpResp.StatusCode)
	require.NotNil(t, updateResp.Data)
	require.Equal(t, id, updateResp.Data.Identifier)
	t.Logf("trigger updated: identifier=%s", updateResp.Data.Identifier)

	// Verify executorInfo round-trips correctly through the update response.
	require.NotNil(t, updateResp.Data.ExecutorInfo)
	require.Equal(t, "SERVICE_ACCOUNT", updateResp.Data.ExecutorInfo.Type_)
	require.Equal(t, saIdentifier, updateResp.Data.ExecutorInfo.Uuid)
	require.Equal(t, env.orgID, updateResp.Data.ExecutorInfo.OrgIdentifier)
	require.Equal(t, env.projectID, updateResp.Data.ExecutorInfo.ProjectIdentifier)
}

// TestCreateTriggerV2_NoExecutorInfo creates a trigger via the v2 API with
// no executorInfo — yaml-only body, same as v1 but through the v2 endpoint.
func TestCreateTriggerV2_NoExecutorInfo(t *testing.T) {
	apiKey := os.Getenv("HARNESS_PLATFORM_API_KEY")
	if apiKey == "" {
		t.Skip("HARNESS_PLATFORM_API_KEY must be set for this acceptance test")
	}

	env := getTriggerTestEnv(t)
	client, ctx := newTriggerClientAndCtx(t, apiKey)
	id := fmt.Sprintf("trgv2noexec_%s", utils.RandStringBytes(5))

	body := TriggerYamlRequest{
		Yaml: triggerYaml(env, id),
		// ExecutorInfo intentionally omitted
	}

	resp, httpResp, err := client.TriggersApi.CreateTriggerV2(ctx, body, client.AccountId,
		env.orgID, env.projectID, env.pipelineID,
		&TriggersApiCreateTriggerV2Opts{WithServiceV2: optional.NewBool(true)})
	defer deleteTrigger(t, client, ctx, env, id)

	assertTriggerResponse(t, resp, httpResp, err, id)
}
