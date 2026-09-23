package nextgen

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"github.com/stretchr/testify/require"
)

// AutoStopping Rules V2 Integration Tests
// ========================================
//
// These tests perform real CRUD operations against the Harness API.
// A TestMain in this package requires -account and -api-key flags (see
// api_delegate_token_resource_test.go), so those must always be passed via -args.
//

/*
Same-Account Tests
==================
These test CRUD for rules where the proxy and target are in the same account.
The proxy has NO cloud_account_id set.

Environment variables:

  export HARNESS_ACCOUNT_ID="<harness-account-id>"
  export HARNESS_API_KEY="<harness-api-key>"
  export CONNECTOR_ID="<connector-id-for-the-account>"
  export ACCESS_POINT_ID="<ap-xxxxx-autostopping-proxy-id>"
  export TARGET_INSTANCE_ID="<i-xxxxx-ec2-instance-id>"
  export TARGET_VM_REGION="<aws-region>"

Run all same-account tests:

  TF_LOG=DEBUG go test -v -run TestAutoStoppingSameAccount -count=1 ./harness/nextgen/ \
    -args -account "$HARNESS_ACCOUNT_ID" -api-key "$HARNESS_API_KEY"

Run only HTTP or TCP:

  TF_LOG=DEBUG go test -v -run TestAutoStoppingSameAccountHTTP -count=1 ./harness/nextgen/ \
    -args -account "$HARNESS_ACCOUNT_ID" -api-key "$HARNESS_API_KEY"

  TF_LOG=DEBUG go test -v -run TestAutoStoppingSameAccountTCP -count=1 ./harness/nextgen/ \
    -args -account "$HARNESS_ACCOUNT_ID" -api-key "$HARNESS_API_KEY"
*/

/*
Cross-Account Tests
===================
These test CRUD for rules where the proxy is in Account A and the target VM
is in Account B. The proxy has cloud_account_id set to Account A's connector.

Environment variables (same as above, plus PROXY_CONNECTOR_ID):

  export HARNESS_ACCOUNT_ID="<harness-account-id>"
  export HARNESS_API_KEY="<harness-api-key>"
  export CONNECTOR_ID="<connector-id-for-target-account>"
  export ACCESS_POINT_ID="<ap-xxxxx-autostopping-proxy-id-in-proxy-account>"
  export TARGET_INSTANCE_ID="<i-xxxxx-ec2-instance-id-in-target-account>"
  export TARGET_VM_REGION="<aws-region-of-target>"
  export PROXY_CONNECTOR_ID="<connector-id-for-proxy-account>"

Run all cross-account tests:

  TF_LOG=DEBUG go test -v -run TestAutoStoppingCrossAccount -count=1 ./harness/nextgen/ \
    -args -account "$HARNESS_ACCOUNT_ID" -api-key "$HARNESS_API_KEY"

Run only HTTP or TCP:

  TF_LOG=DEBUG go test -v -run TestAutoStoppingCrossAccountHTTP -count=1 ./harness/nextgen/ \
    -args -account "$HARNESS_ACCOUNT_ID" -api-key "$HARNESS_API_KEY"

  TF_LOG=DEBUG go test -v -run TestAutoStoppingCrossAccountTCP -count=1 ./harness/nextgen/ \
    -args -account "$HARNESS_ACCOUNT_ID" -api-key "$HARNESS_API_KEY"
*/


// testEnv holds env vars for same-account tests.
type testEnv struct {
	AccountId        string
	ApiKey           string
	ConnectorId      string // CONNECTOR_ID — connector for the account where proxy + target live
	AccessPointId    string // ACCESS_POINT_ID — proxy access point
	TargetInstanceId string // TARGET_INSTANCE_ID — EC2 instance
	TargetRegion     string // TARGET_VM_REGION — region of the target instance
}

// crossAccountEnv extends testEnv with the proxy connector for cross-account tests.
type crossAccountEnv struct {
	testEnv
	ProxyConnectorId string // PROXY_CONNECTOR_ID — connector for the account where the proxy lives
}

func getTestEnv(t *testing.T) testEnv {
	t.Helper()
	env := testEnv{
		AccountId:        os.Getenv("HARNESS_ACCOUNT_ID"),
		ApiKey:           os.Getenv("HARNESS_API_KEY"),
		ConnectorId:      os.Getenv("CONNECTOR_ID"),
		AccessPointId:    os.Getenv("ACCESS_POINT_ID"),
		TargetInstanceId: os.Getenv("TARGET_INSTANCE_ID"),
		TargetRegion:     os.Getenv("TARGET_VM_REGION"),
	}
	if env.AccountId == "" || env.ApiKey == "" ||
		env.ConnectorId == "" || env.AccessPointId == "" ||
		env.TargetInstanceId == "" || env.TargetRegion == "" {
		t.Skip("Skipping: set HARNESS_ACCOUNT_ID, HARNESS_API_KEY, " +
			"CONNECTOR_ID, ACCESS_POINT_ID, TARGET_INSTANCE_ID, TARGET_VM_REGION")
	}
	return env
}

func getCrossAccountEnv(t *testing.T) crossAccountEnv {
	t.Helper()
	base := getTestEnv(t)
	env := crossAccountEnv{
		testEnv:          base,
		ProxyConnectorId: os.Getenv("PROXY_CONNECTOR_ID"),
	}
	if env.ProxyConnectorId == "" {
		t.Skip("Skipping cross-account test: set PROXY_CONNECTOR_ID")
	}
	return env
}

func newClient(env testEnv) (*APIClient, context.Context) {
	cfg := NewConfiguration()
	cfg.AccountId = env.AccountId
	cfg.ApiKey = env.ApiKey
	client := NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), ContextAPIKey, APIKey{Key: env.ApiKey})
	return client, ctx
}

func uniqueName(prefix string) string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return fmt.Sprintf("ccm35727-%s-%s", prefix, string(b))
}

// deleteRule is a cleanup helper that logs but does not fail on error.
func deleteRule(t *testing.T, client *APIClient, ctx context.Context, accountId string, ruleId int64) {
	t.Helper()
	_, err := client.CloudCostAutoStoppingRulesApi.DeleteAutoStoppingRule(ctx, float64(ruleId), accountId, accountId)
	if err != nil {
		t.Logf("Warning: failed to delete rule %d: %v", ruleId, err)
	} else {
		t.Logf("Deleted rule %d", ruleId)
	}
}

// ---------------------------------------------------------------------------
// Same-account HTTP: Create → Read → Update → Read → Delete
// ---------------------------------------------------------------------------
func TestAutoStoppingSameAccountHTTP(t *testing.T) {
	env := getTestEnv(t)
	client, ctx := newClient(env)

	ruleName := uniqueName("sdk-same-http")

	rule := SaveServiceRequestV2{
		Service: &ServiceV2{
			Name:              ruleName,
			AccountIdentifier: env.AccountId,
			Kind:              "instance",
			CloudAccountId:    env.ConnectorId,
			IdleTimeMins:      5,
			Fulfilment:        "ondemand",
			HealthCheck: &HealthCheck{
				Protocol:       "http",
				Path:           "/",
				Port:           80,
				Timeout:        30,
				StatusCodeFrom: 200,
				StatusCodeTo:   299,
			},
			Routing: &RoutingDataV2{
				Instance: &InstanceBasedRoutingDataV2{
					Filter: &FilterObject{
						Ids:     []string{env.TargetInstanceId},
						Regions: []string{env.TargetRegion},
					},
				},
				Http: &HttpProxy{
					Proxy: &Proxy{
						Id: env.AccessPointId,
						// No CloudAccountId — same account
					},
					Ports: []PortConfig{
						{
							Protocol:       "http",
							TargetProtocol: "http",
							Port:           80,
							TargetPort:     80,
							Action:         "forward",
							RoutingRules:   []RoutingRule{},
						},
					},
				},
			},
			Opts: &Opts{},
			Metadata: &ServiceMetadata{
				CloudProviderDetails: &ServiceMetadataCloudProviderDetails{
					Name: "sdk-test-same-account",
				},
			},
			Disabled: false,
		},
		Deps:     []ServiceDep{},
		ApplyNow: false,
	}

	// CREATE
	createResp, httpResp, err := client.CloudCostAutoStoppingRulesV2Api.CreateAutoStoppingRuleV2(ctx, rule, env.AccountId, env.AccountId)
	require.NoError(t, err)
	require.Equal(t, 201, httpResp.StatusCode)
	require.NotNil(t, createResp.Response)
	ruleId := createResp.Response.Id
	t.Logf("Created same-account HTTP rule: id=%d", ruleId)
	defer deleteRule(t, client, ctx, env.AccountId, ruleId)
	require.NotZero(t, ruleId)

	svc := createResp.Response
	require.Equal(t, ruleName, svc.Name)
	require.Equal(t, env.ConnectorId, svc.CloudAccountId)
	require.NotNil(t, svc.Routing.Http)
	require.NotNil(t, svc.Routing.Http.Proxy)
	require.Equal(t, env.AccessPointId, svc.Routing.Http.Proxy.Id)
	require.Equal(t, env.ConnectorId, svc.Routing.Http.Proxy.CloudAccountId, "same-account: API auto-fills proxy cloud_account_id with the connector")

	// LIST — verify rule appears in list
	listResp, _, err := client.CloudCostAutoStoppingRulesApi.ListAutoStoppingRules(ctx, env.AccountId, env.AccountId)
	require.NoError(t, err)
	found := false
	for _, s := range listResp.Response {
		if s.Id == ruleId {
			require.Equal(t, ruleName, s.Name)
			found = true
			break
		}
	}
	require.True(t, found, "rule %d not found in list", ruleId)

	// READ
	getResp, httpResp, err := client.CloudCostAutoStoppingRulesV2Api.GetAutoStoppingRuleV2(ctx, env.AccountId, float64(ruleId), env.AccountId)
	require.NoError(t, err)
	require.Equal(t, 200, httpResp.StatusCode)
	readSvc := getResp.Response.Service
	require.Equal(t, ruleName, readSvc.Name)
	require.Equal(t, env.AccessPointId, readSvc.Routing.Http.Proxy.Id)
	require.Equal(t, env.ConnectorId, readSvc.Routing.Http.Proxy.CloudAccountId)

	// UPDATE — change name and idle time
	updatedName := ruleName + "-updated"
	rule.Service.Name = updatedName
	rule.Service.IdleTimeMins = 10
	rule.Service.Id = ruleId
	updateResp, httpResp, err := client.CloudCostAutoStoppingRulesV2Api.UpdateAutoStoppingRuleV2(ctx, rule, env.AccountId, env.AccountId, strconv.FormatInt(ruleId, 10))
	require.NoError(t, err)
	require.Equal(t, 201, httpResp.StatusCode)
	require.Equal(t, updatedName, updateResp.Response.Name)
	require.Equal(t, 10, updateResp.Response.IdleTimeMins)
	require.Equal(t, env.ConnectorId, updateResp.Response.Routing.Http.Proxy.CloudAccountId)

	// READ after update
	getResp2, _, err := client.CloudCostAutoStoppingRulesV2Api.GetAutoStoppingRuleV2(ctx, env.AccountId, float64(ruleId), env.AccountId)
	require.NoError(t, err)
	require.Equal(t, updatedName, getResp2.Response.Service.Name)
	require.Equal(t, 10, getResp2.Response.Service.IdleTimeMins)

	// DELETE is handled by defer
}

// ---------------------------------------------------------------------------
// Same-account TCP: Create → Read → Update → Read → Delete
// ---------------------------------------------------------------------------
func TestAutoStoppingSameAccountTCP(t *testing.T) {
	env := getTestEnv(t)
	client, ctx := newClient(env)

	ruleName := uniqueName("sdk-same-tcp")

	rule := SaveServiceRequestV2{
		Service: &ServiceV2{
			Name:              ruleName,
			AccountIdentifier: env.AccountId,
			Kind:              "instance",
			CloudAccountId:    env.ConnectorId,
			IdleTimeMins:      5,
			Fulfilment:        "ondemand",
			Routing: &RoutingDataV2{
				Instance: &InstanceBasedRoutingDataV2{
					Filter: &FilterObject{
						Ids:     []string{env.TargetInstanceId},
						Regions: []string{env.TargetRegion},
					},
				},
				Tcp: &TcpProxy{
					Proxy: &Proxy{
						Id: env.AccessPointId,
						// No CloudAccountId — same account
					},
					SshConf: &ServiceRoutingTcpPort{
						Source: 2222,
						Target: 22,
					},
				},
			},
			Opts: &Opts{},
			Metadata: &ServiceMetadata{
				CloudProviderDetails: &ServiceMetadataCloudProviderDetails{
					Name: "sdk-test-same-account",
				},
			},
			Disabled: false,
		},
		Deps:     []ServiceDep{},
		ApplyNow: false,
	}

	// CREATE
	createResp, httpResp, err := client.CloudCostAutoStoppingRulesV2Api.CreateAutoStoppingRuleV2(ctx, rule, env.AccountId, env.AccountId)
	require.NoError(t, err)
	require.Equal(t, 201, httpResp.StatusCode)
	require.NotNil(t, createResp.Response)
	ruleId := createResp.Response.Id
	t.Logf("Created same-account TCP rule: id=%d", ruleId)
	defer deleteRule(t, client, ctx, env.AccountId, ruleId)
	require.NotZero(t, ruleId)

	svc := createResp.Response
	require.Equal(t, env.ConnectorId, svc.CloudAccountId)
	require.NotNil(t, svc.Routing.Tcp)
	require.NotNil(t, svc.Routing.Tcp.Proxy)
	require.Equal(t, env.AccessPointId, svc.Routing.Tcp.Proxy.Id)
	require.Equal(t, env.ConnectorId, svc.Routing.Tcp.Proxy.CloudAccountId, "same-account: API auto-fills proxy cloud_account_id with the connector")

	// LIST — verify rule appears in list
	listResp, _, err := client.CloudCostAutoStoppingRulesApi.ListAutoStoppingRules(ctx, env.AccountId, env.AccountId)
	require.NoError(t, err)
	found := false
	for _, s := range listResp.Response {
		if s.Id == ruleId {
			require.Equal(t, ruleName, s.Name)
			found = true
			break
		}
	}
	require.True(t, found, "rule %d not found in list", ruleId)

	// READ
	getResp, _, err := client.CloudCostAutoStoppingRulesV2Api.GetAutoStoppingRuleV2(ctx, env.AccountId, float64(ruleId), env.AccountId)
	require.NoError(t, err)
	readSvc := getResp.Response.Service
	require.Equal(t, env.AccessPointId, readSvc.Routing.Tcp.Proxy.Id)
	require.Equal(t, env.ConnectorId, readSvc.Routing.Tcp.Proxy.CloudAccountId)

	// UPDATE — change idle time
	updatedName := ruleName + "-updated"
	rule.Service.Name = updatedName
	rule.Service.IdleTimeMins = 10
	rule.Service.Id = ruleId
	updateResp, httpResp, err := client.CloudCostAutoStoppingRulesV2Api.UpdateAutoStoppingRuleV2(ctx, rule, env.AccountId, env.AccountId, strconv.FormatInt(ruleId, 10))
	require.NoError(t, err)
	require.Equal(t, 201, httpResp.StatusCode)
	require.Equal(t, updatedName, updateResp.Response.Name)
	require.Equal(t, env.ConnectorId, updateResp.Response.Routing.Tcp.Proxy.CloudAccountId)

	// READ after update
	getResp2, _, err := client.CloudCostAutoStoppingRulesV2Api.GetAutoStoppingRuleV2(ctx, env.AccountId, float64(ruleId), env.AccountId)
	require.NoError(t, err)
	require.Equal(t, updatedName, getResp2.Response.Service.Name)
	require.Equal(t, 10, getResp2.Response.Service.IdleTimeMins)

	// DELETE is handled by defer
}

// ---------------------------------------------------------------------------
// Cross-account HTTP: Create → Read → Update → Read → Delete
// ---------------------------------------------------------------------------
func TestAutoStoppingCrossAccountHTTP(t *testing.T) {
	env := getCrossAccountEnv(t)
	client, ctx := newClient(env.testEnv)

	ruleName := uniqueName("sdk-xacct-http")

	rule := SaveServiceRequestV2{
		Service: &ServiceV2{
			Name:              ruleName,
			AccountIdentifier: env.AccountId,
			Kind:              "instance",
			CloudAccountId:    env.ConnectorId,
			IdleTimeMins:      5,
			Fulfilment:        "ondemand",
			HealthCheck: &HealthCheck{
				Protocol:       "http",
				Path:           "/",
				Port:           80,
				Timeout:        30,
				StatusCodeFrom: 200,
				StatusCodeTo:   299,
			},
			Routing: &RoutingDataV2{
				Instance: &InstanceBasedRoutingDataV2{
					Filter: &FilterObject{
						Ids:     []string{env.TargetInstanceId},
						Regions: []string{env.TargetRegion},
					},
				},
				Http: &HttpProxy{
					Proxy: &Proxy{
						Id:             env.AccessPointId,
						CloudAccountId: env.ProxyConnectorId, // cross-account
					},
					Ports: []PortConfig{
						{
							Protocol:       "http",
							TargetProtocol: "http",
							Port:           80,
							TargetPort:     80,
							Action:         "forward",
							RoutingRules:   []RoutingRule{},
						},
					},
				},
			},
			Opts: &Opts{},
			Metadata: &ServiceMetadata{
				CloudProviderDetails: &ServiceMetadataCloudProviderDetails{
					Name: "sdk-test-target-account",
				},
			},
			Disabled: false,
		},
		Deps:     []ServiceDep{},
		ApplyNow: false,
	}

	// CREATE
	createResp, httpResp, err := client.CloudCostAutoStoppingRulesV2Api.CreateAutoStoppingRuleV2(ctx, rule, env.AccountId, env.AccountId)
	require.NoError(t, err)
	require.Equal(t, 201, httpResp.StatusCode)
	require.NotNil(t, createResp.Response)
	ruleId := createResp.Response.Id
	t.Logf("Created cross-account HTTP rule: id=%d", ruleId)
	defer deleteRule(t, client, ctx, env.AccountId, ruleId)
	require.NotZero(t, ruleId)

	svc := createResp.Response
	require.Equal(t, env.ConnectorId, svc.CloudAccountId)
	require.NotNil(t, svc.Routing.Http)
	require.NotNil(t, svc.Routing.Http.Proxy)
	require.Equal(t, env.AccessPointId, svc.Routing.Http.Proxy.Id)
	require.Equal(t, env.ProxyConnectorId, svc.Routing.Http.Proxy.CloudAccountId, "cross-account: proxy cloud_account_id must be set")

	// LIST — verify rule appears in list
	listResp, _, err := client.CloudCostAutoStoppingRulesApi.ListAutoStoppingRules(ctx, env.AccountId, env.AccountId)
	require.NoError(t, err)
	found := false
	for _, s := range listResp.Response {
		if s.Id == ruleId {
			require.Equal(t, ruleName, s.Name)
			found = true
			break
		}
	}
	require.True(t, found, "rule %d not found in list", ruleId)

	// READ
	getResp, httpResp, err := client.CloudCostAutoStoppingRulesV2Api.GetAutoStoppingRuleV2(ctx, env.AccountId, float64(ruleId), env.AccountId)
	require.NoError(t, err)
	require.Equal(t, 200, httpResp.StatusCode)
	readSvc := getResp.Response.Service
	require.Equal(t, ruleName, readSvc.Name)
	require.Equal(t, env.ConnectorId, readSvc.CloudAccountId)
	require.Equal(t, env.AccessPointId, readSvc.Routing.Http.Proxy.Id)
	require.Equal(t, env.ProxyConnectorId, readSvc.Routing.Http.Proxy.CloudAccountId, "read-back: proxy cloud_account_id must survive round-trip")

	// UPDATE — change name and idle time, keep cross-account proxy
	updatedName := ruleName + "-updated"
	rule.Service.Name = updatedName
	rule.Service.IdleTimeMins = 10
	rule.Service.Id = ruleId
	updateResp, httpResp, err := client.CloudCostAutoStoppingRulesV2Api.UpdateAutoStoppingRuleV2(ctx, rule, env.AccountId, env.AccountId, strconv.FormatInt(ruleId, 10))
	require.NoError(t, err)
	require.Equal(t, 201, httpResp.StatusCode)
	require.Equal(t, updatedName, updateResp.Response.Name)
	require.Equal(t, 10, updateResp.Response.IdleTimeMins)
	require.Equal(t, env.ProxyConnectorId, updateResp.Response.Routing.Http.Proxy.CloudAccountId, "cross-account: proxy cloud_account_id must persist after update")

	// READ after update
	getResp2, _, err := client.CloudCostAutoStoppingRulesV2Api.GetAutoStoppingRuleV2(ctx, env.AccountId, float64(ruleId), env.AccountId)
	require.NoError(t, err)
	require.Equal(t, updatedName, getResp2.Response.Service.Name)
	require.Equal(t, 10, getResp2.Response.Service.IdleTimeMins)
	require.Equal(t, env.ProxyConnectorId, getResp2.Response.Service.Routing.Http.Proxy.CloudAccountId)

	// DELETE is handled by defer
}

// ---------------------------------------------------------------------------
// Cross-account TCP: Create → Read → Update → Read → Delete
// ---------------------------------------------------------------------------
func TestAutoStoppingCrossAccountTCP(t *testing.T) {
	env := getCrossAccountEnv(t)
	client, ctx := newClient(env.testEnv)

	ruleName := uniqueName("sdk-xacct-tcp")

	rule := SaveServiceRequestV2{
		Service: &ServiceV2{
			Name:              ruleName,
			AccountIdentifier: env.AccountId,
			Kind:              "instance",
			CloudAccountId:    env.ConnectorId,
			IdleTimeMins:      5,
			Fulfilment:        "ondemand",
			Routing: &RoutingDataV2{
				Instance: &InstanceBasedRoutingDataV2{
					Filter: &FilterObject{
						Ids:     []string{env.TargetInstanceId},
						Regions: []string{env.TargetRegion},
					},
				},
				Tcp: &TcpProxy{
					Proxy: &Proxy{
						Id:             env.AccessPointId,
						CloudAccountId: env.ProxyConnectorId, // cross-account
					},
					SshConf: &ServiceRoutingTcpPort{
						Source: 2222,
						Target: 22,
					},
					CustomPorts: []ServiceRoutingTcpPort{
						{
							Source: 8080,
							Target: 80,
						},
					},
				},
			},
			Opts: &Opts{},
			Metadata: &ServiceMetadata{
				CloudProviderDetails: &ServiceMetadataCloudProviderDetails{
					Name: "sdk-test-target-account",
				},
			},
			Disabled: false,
		},
		Deps:     []ServiceDep{},
		ApplyNow: false,
	}

	// CREATE
	createResp, httpResp, err := client.CloudCostAutoStoppingRulesV2Api.CreateAutoStoppingRuleV2(ctx, rule, env.AccountId, env.AccountId)
	require.NoError(t, err)
	require.Equal(t, 201, httpResp.StatusCode)
	require.NotNil(t, createResp.Response)
	ruleId := createResp.Response.Id
	t.Logf("Created cross-account TCP rule: id=%d", ruleId)
	defer deleteRule(t, client, ctx, env.AccountId, ruleId)
	require.NotZero(t, ruleId)

	svc := createResp.Response
	require.Equal(t, env.ConnectorId, svc.CloudAccountId)
	require.NotNil(t, svc.Routing.Tcp)
	require.NotNil(t, svc.Routing.Tcp.Proxy)
	require.Equal(t, env.AccessPointId, svc.Routing.Tcp.Proxy.Id)
	require.Equal(t, env.ProxyConnectorId, svc.Routing.Tcp.Proxy.CloudAccountId, "cross-account: proxy cloud_account_id must be set")

	// LIST — verify rule appears in list
	listResp, _, err := client.CloudCostAutoStoppingRulesApi.ListAutoStoppingRules(ctx, env.AccountId, env.AccountId)
	require.NoError(t, err)
	found := false
	for _, s := range listResp.Response {
		if s.Id == ruleId {
			require.Equal(t, ruleName, s.Name)
			found = true
			break
		}
	}
	require.True(t, found, "rule %d not found in list", ruleId)

	// READ
	getResp, _, err := client.CloudCostAutoStoppingRulesV2Api.GetAutoStoppingRuleV2(ctx, env.AccountId, float64(ruleId), env.AccountId)
	require.NoError(t, err)
	readSvc := getResp.Response.Service
	require.Equal(t, ruleName, readSvc.Name)
	require.Equal(t, env.AccessPointId, readSvc.Routing.Tcp.Proxy.Id)
	require.Equal(t, env.ProxyConnectorId, readSvc.Routing.Tcp.Proxy.CloudAccountId, "read-back: proxy cloud_account_id must survive round-trip")

	// UPDATE — change name and idle time, keep cross-account proxy
	updatedName := ruleName + "-updated"
	rule.Service.Name = updatedName
	rule.Service.IdleTimeMins = 10
	rule.Service.Id = ruleId
	updateResp, httpResp, err := client.CloudCostAutoStoppingRulesV2Api.UpdateAutoStoppingRuleV2(ctx, rule, env.AccountId, env.AccountId, strconv.FormatInt(ruleId, 10))
	require.NoError(t, err)
	require.Equal(t, 201, httpResp.StatusCode)
	require.Equal(t, updatedName, updateResp.Response.Name)
	require.Equal(t, env.ProxyConnectorId, updateResp.Response.Routing.Tcp.Proxy.CloudAccountId, "cross-account: proxy cloud_account_id must persist after update")

	// READ after update
	getResp2, _, err := client.CloudCostAutoStoppingRulesV2Api.GetAutoStoppingRuleV2(ctx, env.AccountId, float64(ruleId), env.AccountId)
	require.NoError(t, err)
	require.Equal(t, updatedName, getResp2.Response.Service.Name)
	require.Equal(t, 10, getResp2.Response.Service.IdleTimeMins)
	require.Equal(t, env.ProxyConnectorId, getResp2.Response.Service.Routing.Tcp.Proxy.CloudAccountId)

	// DELETE is handled by defer
}
