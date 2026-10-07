package nextgen

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

/*
Anthropic Model Connector — Unit Tests (TestUnitAnthropicModel_*)
Tests JSON serialization/deserialization round-trips for all four auth types
(Token, BedrockApiKey, Vertex, CloudProvider), real API response parsing,
and CreateConnector request construction via a mock HTTP transport.

No env vars or network access required.

Run all unit tests:

    go test -v -run TestUnitAnthropicModel -count=1 \
      ./harness/nextgen/ -args -account dummy -api-key dummy

Run a single unit test:

    go test -v -run TestUnitAnthropicModel_UnmarshalToken -count=1 \
      ./harness/nextgen/ -args -account dummy -api-key dummy

Note: -args -account dummy -api-key dummy is needed because TestMain in this
package gates all tests behind those flags.
*/

func TestUnitAnthropicModel_UnmarshalToken(t *testing.T) {
	raw := `{
		"name": "test-token",
		"identifier": "testtoken",
		"type": "Anthropic",
		"spec": {
			"url": "https://api.anthropic.com",
			"authentication": {
				"type": "Token",
				"spec": { "tokenRef": "account.my_key" }
			},
			"delegateSelectors": [],
			"executeOnDelegate": true,
			"ignoreTestConnection": false,
			"model": "claude-opus-4-5-20251101",
			"harnessManaged": false
		}
	}`

	var info ConnectorInfo
	err := json.Unmarshal([]byte(raw), &info)
	require.NoError(t, err)
	require.Equal(t, ConnectorTypes.AnthropicModel, info.Type_)
	require.NotNil(t, info.AnthropicModel)
	require.Equal(t, "https://api.anthropic.com", info.AnthropicModel.Url)
	require.Equal(t, "claude-opus-4-5-20251101", info.AnthropicModel.Model)
	require.Equal(t, AnthropicModelAuthTypes.Token, info.AnthropicModel.Authentication.Type_)
	require.NotNil(t, info.AnthropicModel.Authentication.Token)
	require.Equal(t, "account.my_key", info.AnthropicModel.Authentication.Token.TokenRef)
}

func TestUnitAnthropicModel_UnmarshalBedrockApiKey(t *testing.T) {
	raw := `{
		"name": "test-bedrock",
		"identifier": "testbedrock",
		"type": "Anthropic",
		"spec": {
			"authentication": {
				"type": "BedrockApiKey",
				"spec": { "apiKeyRef": "account.aws_key", "region": "us-east-1" }
			},
			"delegateSelectors": [],
			"executeOnDelegate": true,
			"ignoreTestConnection": false,
			"model": "global.anthropic.claude-opus-4-6-v1",
			"harnessManaged": false
		}
	}`

	var info ConnectorInfo
	err := json.Unmarshal([]byte(raw), &info)
	require.NoError(t, err)
	require.Equal(t, AnthropicModelAuthTypes.BedrockApiKey, info.AnthropicModel.Authentication.Type_)
	require.NotNil(t, info.AnthropicModel.Authentication.BedrockApiKey)
	require.Equal(t, "account.aws_key", info.AnthropicModel.Authentication.BedrockApiKey.ApiKeyRef)
	require.Equal(t, "us-east-1", info.AnthropicModel.Authentication.BedrockApiKey.Region)
}

func TestUnitAnthropicModel_UnmarshalVertex(t *testing.T) {
	raw := `{
		"name": "test-vertex",
		"identifier": "testvertex",
		"type": "Anthropic",
		"spec": {
			"authentication": {
				"type": "Vertex",
				"spec": {
					"serviceAccountKeyRef": "account.gcp_sa_key",
					"projectId": "my-gcp-project",
					"region": "us-central1"
				}
			},
			"delegateSelectors": [],
			"executeOnDelegate": false,
			"ignoreTestConnection": false,
			"model": "claude-sonnet-4-6",
			"harnessManaged": false
		}
	}`

	var info ConnectorInfo
	err := json.Unmarshal([]byte(raw), &info)
	require.NoError(t, err)
	require.Equal(t, AnthropicModelAuthTypes.Vertex, info.AnthropicModel.Authentication.Type_)
	require.NotNil(t, info.AnthropicModel.Authentication.Vertex)
	require.Equal(t, "account.gcp_sa_key", info.AnthropicModel.Authentication.Vertex.ServiceAccountKeyRef)
	require.Equal(t, "my-gcp-project", info.AnthropicModel.Authentication.Vertex.ProjectId)
	require.Equal(t, "us-central1", info.AnthropicModel.Authentication.Vertex.Region)
}

func TestUnitAnthropicModel_UnmarshalCloudProvider(t *testing.T) {
	raw := `{
		"name": "test-cp",
		"identifier": "testcp",
		"type": "Anthropic",
		"spec": {
			"authentication": {
				"type": "CloudProvider",
				"spec": { "type": "AWS", "connectorRef": "nwt_aws_oidc_dev" }
			},
			"delegateSelectors": [],
			"executeOnDelegate": true,
			"ignoreTestConnection": true,
			"model": "claude-opus-4-6",
			"harnessManaged": false
		}
	}`

	var info ConnectorInfo
	err := json.Unmarshal([]byte(raw), &info)
	require.NoError(t, err)
	require.Equal(t, AnthropicModelAuthTypes.CloudProvider, info.AnthropicModel.Authentication.Type_)
	require.NotNil(t, info.AnthropicModel.Authentication.CloudProvider)
	require.Equal(t, AnthropicModelCloudProviderTypes.AWS, info.AnthropicModel.Authentication.CloudProvider.Type_)
	require.Equal(t, "nwt_aws_oidc_dev", info.AnthropicModel.Authentication.CloudProvider.ConnectorRef)
}

func TestUnitAnthropicModel_MarshalRoundTrip(t *testing.T) {
	original := ConnectorInfo{
		Name:       "round-trip-test",
		Identifier: "roundtriptest",
		Type_:      ConnectorTypes.AnthropicModel,
		AnthropicModel: &AnthropicModelConnector{
			Url:               "https://api.anthropic.com",
			Model:             "claude-opus-4-6",
			ExecuteOnDelegate: true,
			Authentication: &AnthropicModelAuthentication{
				Type_: AnthropicModelAuthTypes.Token,
				Token: &AnthropicModelTokenSpec{TokenRef: "account.test_key"},
			},
		},
	}

	data, err := json.Marshal(&original)
	require.NoError(t, err)

	var decoded ConnectorInfo
	err = json.Unmarshal(data, &decoded)
	require.NoError(t, err)

	require.Equal(t, original.Name, decoded.Name)
	require.Equal(t, original.Identifier, decoded.Identifier)
	require.Equal(t, original.Type_, decoded.Type_)
	require.NotNil(t, decoded.AnthropicModel)
	require.Equal(t, original.AnthropicModel.Url, decoded.AnthropicModel.Url)
	require.Equal(t, original.AnthropicModel.Model, decoded.AnthropicModel.Model)
	require.Equal(t, original.AnthropicModel.ExecuteOnDelegate, decoded.AnthropicModel.ExecuteOnDelegate)
	require.Equal(t, original.AnthropicModel.Authentication.Type_, decoded.AnthropicModel.Authentication.Type_)
	require.Equal(t, original.AnthropicModel.Authentication.Token.TokenRef, decoded.AnthropicModel.Authentication.Token.TokenRef)
}

func TestUnitAnthropicModel_UnmarshalRealAPIResponse(t *testing.T) {
	raw := `{
		"status": "SUCCESS",
		"data": {
			"connector": {
				"name": "senthi-anthropic-testconn-01",
				"identifier": "senthianthropictestconn01",
				"description": "",
				"accountIdentifier": "AM8HCbDiTXGQNrTIhNl7qQ",
				"orgIdentifier": "default",
				"projectIdentifier": "senthilproj01",
				"tags": {},
				"type": "Anthropic",
				"spec": {
					"url": "https://api.anthropic.com",
					"authentication": {
						"type": "Token",
						"spec": { "tokenRef": "test_mock_secret_ref" }
					},
					"delegateSelectors": [],
					"executeOnDelegate": false,
					"ignoreTestConnection": true,
					"model": "claude-opus-4-5-20251101",
					"harnessManaged": false
				}
			},
			"createdAt": 1790701106741,
			"lastModifiedAt": 1790701106736
		}
	}`

	type apiResp struct {
		Data struct {
			Connector ConnectorInfo `json:"connector"`
		} `json:"data"`
	}

	var resp apiResp
	err := json.Unmarshal([]byte(raw), &resp)
	require.NoError(t, err)

	c := resp.Data.Connector
	require.Equal(t, "senthi-anthropic-testconn-01", c.Name)
	require.Equal(t, ConnectorTypes.AnthropicModel, c.Type_)
	require.NotNil(t, c.AnthropicModel)
	require.Equal(t, "https://api.anthropic.com", c.AnthropicModel.Url)
	require.True(t, c.AnthropicModel.IgnoreTestConnection)
	require.False(t, c.AnthropicModel.HarnessManaged)
	require.Equal(t, AnthropicModelAuthTypes.Token, c.AnthropicModel.Authentication.Type_)
	require.Equal(t, "test_mock_secret_ref", c.AnthropicModel.Authentication.Token.TokenRef)
}

func TestUnitAnthropicModel_CreateConnectorRequest(t *testing.T) {
	cfg := NewConfiguration()
	cfg.BasePath = "https://app.harness.io"

	var captured *http.Request
	var capturedBody []byte
	cfg.HTTPClient.HTTPClient.Transport = roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		captured = r
		capturedBody, _ = io.ReadAll(r.Body)
		return &http.Response{
			StatusCode: 200,
			Body: io.NopCloser(strings.NewReader(`{
				"status": "SUCCESS",
				"data": {
					"connector": {
						"name": "mock-anthropic",
						"identifier": "mockanthropic",
						"type": "Anthropic",
						"spec": {
							"authentication": { "type": "Token", "spec": { "tokenRef": "account.key" } },
							"executeOnDelegate": false,
							"ignoreTestConnection": true,
							"model": "claude-opus-4-6",
							"harnessManaged": false
						}
					}
				}
			}`)),
			Header: http.Header{"Content-Type": []string{"application/json"}},
		}, nil
	})

	client := NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), ContextAPIKey, APIKey{Key: "test-key"})

	conn := Connector{
		Connector: &ConnectorInfo{
			Name:       "mock-anthropic",
			Identifier: "mockanthropic",
			Type_:      ConnectorTypes.AnthropicModel,
			AnthropicModel: &AnthropicModelConnector{
				Model:                "claude-opus-4-6",
				IgnoreTestConnection: true,
				Authentication: &AnthropicModelAuthentication{
					Type_: AnthropicModelAuthTypes.Token,
					Token: &AnthropicModelTokenSpec{TokenRef: "account.key"},
				},
			},
		},
	}

	resp, httpResp, err := client.ConnectorsApi.CreateConnector(ctx, conn, "test_account", &ConnectorsApiCreateConnectorOpts{})
	require.NoError(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, 200, httpResp.StatusCode)
	require.Equal(t, "mock-anthropic", resp.Data.Connector.Name)
	require.Equal(t, ConnectorTypes.AnthropicModel, resp.Data.Connector.Type_)

	require.NotNil(t, captured)
	require.Equal(t, http.MethodPost, captured.Method)
	require.Equal(t, "/ng/api/connectors", captured.URL.Path)
	require.Equal(t, "test_account", captured.URL.Query().Get("accountIdentifier"))
	require.Equal(t, "test-key", captured.Header.Get("x-api-key"))

	require.Contains(t, string(capturedBody), `"type":"Anthropic"`)
	require.Contains(t, string(capturedBody), `"tokenRef":"account.key"`)
}
