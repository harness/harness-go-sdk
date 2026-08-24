/*
Required env vars for integration tests:
========================================
  HARNESS_ACCOUNT_ID            - Account identifier
  HARNESS_PLATFORM_API_KEY      - API key for authentication
  IDP_TEST_PLUGIN_ID            - Plugin ID to test with (e.g. "harness-proxy")
  IDP_TEST_PLUGIN_NAME          - Plugin display name (e.g. "Configure Backend Proxies")
  IDP_TEST_SECRET_ID            - Harness secret identifier that exists in the account
  IDP_TEST_PLUGIN_CONFIGS_FILE  - (optional) Path to YAML file with plugin configs; if empty, SaveOrUpdate test is skipped
                                  Example file content (for harness-proxy plugin):
                                    proxy:
                                      endpoints:
                                        /test:
                                          target: https://app.harness.io/gateway
                                          pathRewrite:
                                            api/proxy/test/?: /
                                          headers:
                                            x-api-key: ${TEST_ENV_VAR}

Run all tests
=============
	go test -v -run TestPluginAppConfig ./harness/idp

Run only specific test e.g below
================================
	TF_LOG=DEBUG go test -v -run TestPluginAppConfigIntegration/SaveOrUpdate ./harness/idp/...
*/

package idp

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPluginAppConfigSaveOrUpdateUnit(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/v1/app-config", r.URL.Path)
		require.NotEmpty(t, r.Header.Get("x-api-key"))
		require.NotEmpty(t, r.Header.Get("Harness-Account"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := PluginAppConfigResponse{
			AppConfig: &PluginAppConfigResponseData{
				ConfigId:   "test-plugin",
				ConfigName: "Test Plugin",
				Configs:    "test:\n  key: value",
				Enabled:    true,
				Created:    1700000000000,
				Updated:    1700000000000,
				EnvVariables: []PluginAppConfigEnvVar{
					{
						Identifier:              "abc123",
						EnvName:                 "TEST_TOKEN",
						Type:                    "Secret",
						HarnessSecretIdentifier: "my_secret",
					},
				},
				Proxy: []PluginAppConfigProxy{},
			},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.BasePath = server.URL + "/v1"
	client := NewAPIClient(cfg)
	client.AccountId = "test-account"
	ctx := context.WithValue(context.Background(), ContextAPIKey, APIKey{Key: "test-key"})

	req := PluginAppConfigRequest{
		AppConfig: PluginAppConfig{
			ConfigId:   "test-plugin",
			ConfigName: "Test Plugin",
			Enabled:    true,
			Configs:    "test:\n  key: value",
			EnvVariables: []PluginAppConfigEnvVar{
				{
					EnvName:                 "TEST_TOKEN",
					Type:                    "Secret",
					HarnessSecretIdentifier: "my_secret",
				},
			},
			Proxy: []PluginAppConfigProxy{},
		},
	}

	resp, err := client.PluginAppConfigApi.SaveOrUpdate(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, resp.AppConfig)
	require.Equal(t, "test-plugin", resp.AppConfig.ConfigId)
	require.Equal(t, "Test Plugin", resp.AppConfig.ConfigName)
	require.True(t, resp.AppConfig.Enabled)
	require.Len(t, resp.AppConfig.EnvVariables, 1)
	require.Equal(t, "TEST_TOKEN", resp.AppConfig.EnvVariables[0].EnvName)
}

func TestPluginAppConfigGetPluginInfoUnit(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/v1/plugins-info/harness-proxy", r.URL.Path)
		require.NotEmpty(t, r.Header.Get("x-api-key"))
		require.NotEmpty(t, r.Header.Get("Harness-Account"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		config := "proxy:\n  endpoints:\n    /test:\n      target: https://example.com"
		resp := PluginInfoResponse{
			Plugin: &PluginInfoData{
				PluginDetails: &PluginDetails{
					Id:         "harness-proxy",
					Name:       "Configure Backend Proxies",
					CreatedBy:  "Harness",
					Category:   "Utility",
					Enabled:    true,
					PluginType: "DEFAULT",
				},
				Config: &config,
				EnvVariables: []PluginAppConfigEnvVar{
					{
						Identifier:              "env1",
						EnvName:                 "PROXY_TOKEN",
						Type:                    "Secret",
						HarnessSecretIdentifier: "my_token",
					},
				},
				Proxy: []PluginAppConfigProxy{},
				Saved: true,
			},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.BasePath = server.URL + "/v1"
	client := NewAPIClient(cfg)
	client.AccountId = "test-account"
	ctx := context.WithValue(context.Background(), ContextAPIKey, APIKey{Key: "test-key"})

	resp, err := client.PluginAppConfigApi.GetPluginInfo(ctx, "harness-proxy")
	require.NoError(t, err)
	require.NotNil(t, resp.Plugin)
	require.NotNil(t, resp.Plugin.PluginDetails)
	require.Equal(t, "harness-proxy", resp.Plugin.PluginDetails.Id)
	require.True(t, resp.Plugin.PluginDetails.Enabled)
	require.True(t, resp.Plugin.Saved)
	require.NotNil(t, resp.Plugin.Config)
	require.Contains(t, *resp.Plugin.Config, "proxy:")
	require.Len(t, resp.Plugin.EnvVariables, 1)
}

func TestPluginAppConfigToggleUnit(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/v1/plugin-toggle/harness-proxy", r.URL.Path)
		require.Equal(t, "true", r.URL.Query().Get("enabled"))
		require.NotEmpty(t, r.Header.Get("x-api-key"))
		require.NotEmpty(t, r.Header.Get("Harness-Account"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := PluginAppConfigResponse{
			AppConfig: &PluginAppConfigResponseData{
				ConfigId:          "harness-proxy",
				ConfigName:        "Configure Backend Proxies",
				Enabled:           true,
				EnabledDisabledAt: 1700000000000,
			},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.BasePath = server.URL + "/v1"
	client := NewAPIClient(cfg)
	client.AccountId = "test-account"
	ctx := context.WithValue(context.Background(), ContextAPIKey, APIKey{Key: "test-key"})

	resp, err := client.PluginAppConfigApi.Toggle(ctx, "harness-proxy", true)
	require.NoError(t, err)
	require.NotNil(t, resp.AppConfig)
	require.Equal(t, "harness-proxy", resp.AppConfig.ConfigId)
	require.True(t, resp.AppConfig.Enabled)
}

func TestPluginAppConfigListPluginsUnit(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/v1/plugins-info", r.URL.Path)
		require.NotEmpty(t, r.Header.Get("x-api-key"))
		require.NotEmpty(t, r.Header.Get("Harness-Account"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := []PluginListItem{
			{Plugin: &PluginListDetails{Id: "harness-proxy", Name: "Configure Backend Proxies", Enabled: true, PluginType: "DEFAULT"}},
			{Plugin: &PluginListDetails{Id: "pager-duty", Name: "PagerDuty", Enabled: false, PluginType: "DEFAULT"}},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.BasePath = server.URL + "/v1"
	client := NewAPIClient(cfg)
	client.AccountId = "test-account"
	ctx := context.WithValue(context.Background(), ContextAPIKey, APIKey{Key: "test-key"})

	resp, err := client.PluginAppConfigApi.ListPlugins(ctx)
	require.NoError(t, err)
	require.Len(t, resp, 2)
	require.Equal(t, "harness-proxy", resp[0].Plugin.Id)
	require.True(t, resp[0].Plugin.Enabled)
	require.Equal(t, "pager-duty", resp[1].Plugin.Id)
	require.False(t, resp[1].Plugin.Enabled)
}

// Integration tests
//
// Required env vars:
//   HARNESS_ACCOUNT_ID            - Account identifier
//   HARNESS_PLATFORM_API_KEY      - API key for authentication
//   IDP_TEST_PLUGIN_ID            - Plugin ID to test with (e.g. "harness-proxy")
//   IDP_TEST_PLUGIN_NAME          - Plugin display name (e.g. "Configure Backend Proxies")
//   IDP_TEST_SECRET_ID            - Harness secret identifier that exists in the account
//   IDP_TEST_PLUGIN_CONFIGS_FILE  - (optional) Path to YAML file with plugin configs; if empty, SaveOrUpdate test is skipped
//                                   Example file content (for harness-proxy plugin):
//                                     proxy:
//                                       endpoints:
//                                         /test:
//                                           target: https://app.harness.io/gateway
//                                           pathRewrite:
//                                             api/proxy/test/?: /
//                                           headers:
//                                             x-api-key: ${TEST_ENV_VAR}

func getIntegrationTestConfig(t *testing.T) (pluginId, pluginName, secretId, configs string) {
	t.Helper()
	pluginId = os.Getenv("IDP_TEST_PLUGIN_ID")
	pluginName = os.Getenv("IDP_TEST_PLUGIN_NAME")
	secretId = os.Getenv("IDP_TEST_SECRET_ID")
	if pluginId == "" || pluginName == "" || secretId == "" {
		t.Skip("IDP_TEST_PLUGIN_ID, IDP_TEST_PLUGIN_NAME, and IDP_TEST_SECRET_ID must be set")
	}
	configsFile := os.Getenv("IDP_TEST_PLUGIN_CONFIGS_FILE")
	if configsFile != "" {
		data, err := os.ReadFile(configsFile)
		if err != nil {
			t.Fatalf("failed to read IDP_TEST_PLUGIN_CONFIGS_FILE=%s: %v", configsFile, err)
		}
		configs = string(data)
	}
	return
}

func TestPluginAppConfigIntegration(t *testing.T) {
	client, ctx := getIdpClientWithContext(t)
	pluginId, pluginName, secretId, configs := getIntegrationTestConfig(t)

	// 1. List all plugins
	t.Run("ListPlugins", func(t *testing.T) {
		plugins, err := client.PluginAppConfigApi.ListPlugins(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, plugins, "expected at least one plugin")
		t.Logf("Found %d plugins", len(plugins))

		var found bool
		for _, p := range plugins {
			if p.Plugin != nil && p.Plugin.Id == pluginId {
				found = true
				break
			}
		}
		require.True(t, found, "expected %s in plugin list", pluginId)
	})

	// 2. Get plugin info
	t.Run("GetPluginInfo", func(t *testing.T) {
		resp, err := client.PluginAppConfigApi.GetPluginInfo(ctx, pluginId)
		require.NoError(t, err)
		require.NotNil(t, resp.Plugin)
		require.NotNil(t, resp.Plugin.PluginDetails)
		require.Equal(t, pluginId, resp.Plugin.PluginDetails.Id)
		t.Logf("Plugin enabled=%v, saved=%v", resp.Plugin.PluginDetails.Enabled, resp.Plugin.Saved)
		if resp.Plugin.Config != nil {
			t.Logf("Config length: %d bytes", len(*resp.Plugin.Config))
		}
	})

	// 3. Save or update plugin config
	t.Run("SaveOrUpdate", func(t *testing.T) {
		if configs == "" {
			t.Skip("IDP_TEST_PLUGIN_CONFIGS not set, skipping SaveOrUpdate")
		}
		req := PluginAppConfigRequest{
			AppConfig: PluginAppConfig{
				ConfigId:   pluginId,
				ConfigName: pluginName,
				Enabled:    true,
				Configs:    configs,
				EnvVariables: []PluginAppConfigEnvVar{
					{
						EnvName:                 "TEST_ENV_VAR",
						Type:                    "Secret",
						HarnessSecretIdentifier: secretId,
					},
				},
				Proxy: []PluginAppConfigProxy{},
			},
		}
		resp, err := client.PluginAppConfigApi.SaveOrUpdate(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, resp.AppConfig)
		require.Equal(t, pluginId, resp.AppConfig.ConfigId)
		require.Equal(t, pluginName, resp.AppConfig.ConfigName)
		t.Logf("Saved config, updated=%d", resp.AppConfig.Updated)
	})

	// 4. Toggle plugin (enable)
	t.Run("ToggleEnable", func(t *testing.T) {
		resp, err := client.PluginAppConfigApi.Toggle(ctx, pluginId, true)
		require.NoError(t, err)
		require.NotNil(t, resp.AppConfig)
		require.Equal(t, pluginId, resp.AppConfig.ConfigId)
		require.True(t, resp.AppConfig.Enabled)
		t.Logf("Toggled to enabled, enabled_disabled_at=%d", resp.AppConfig.EnabledDisabledAt)
	})

	// 5. Toggle plugin (disable) — restore state
	t.Run("ToggleDisable", func(t *testing.T) {
		resp, err := client.PluginAppConfigApi.Toggle(ctx, pluginId, false)
		require.NoError(t, err)
		require.NotNil(t, resp.AppConfig)
		require.False(t, resp.AppConfig.Enabled)
		t.Logf("Toggled to disabled, enabled_disabled_at=%d", resp.AppConfig.EnabledDisabledAt)
	})
}
