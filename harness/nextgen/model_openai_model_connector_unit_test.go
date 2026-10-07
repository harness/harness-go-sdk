/* ---------------------------------------------------------------------------
OpenAI Model Connector — Unit Tests
---------------------------------------------------------------------------

Run all unit tests:

  go test -v -run "TestUnitOpenAIModel" -count=1 -timeout 30s \
    ./harness/nextgen/ -args -account dummy -api-key dummy

Run a single test:

  go test -v -run "TestUnitOpenAIModel_UnmarshalToken" -count=1 -timeout 30s \
    ./harness/nextgen/ -args -account dummy -api-key dummy

*/

package nextgen

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnitOpenAIModel_UnmarshalToken(t *testing.T) {
	jsonPayload := `{
		"connector": {
			"name": "test-openai-token",
			"identifier": "test_openai_token",
			"type": "OpenAI",
			"spec": {
				"baseUrl": "https://api.openai.com",
				"model": "gpt-4o",
				"executeOnDelegate": true,
				"ignoreTestConnection": false,
				"delegateSelectors": ["d1"],
				"authentication": {
					"type": "Token",
					"spec": {
						"tokenRef": "account.openai_key"
					}
				}
			}
		}
	}`

	type envelope struct {
		Connector ConnectorInfo `json:"connector"`
	}

	var env envelope
	err := json.Unmarshal([]byte(jsonPayload), &env)
	require.NoError(t, err)

	c := env.Connector
	assert.Equal(t, ConnectorTypes.OpenAIModel, c.Type_)
	assert.NotNil(t, c.OpenAIModel)
	assert.Equal(t, "https://api.openai.com", c.OpenAIModel.Url)
	assert.Equal(t, "gpt-4o", c.OpenAIModel.Model)
	assert.True(t, c.OpenAIModel.ExecuteOnDelegate)
	assert.False(t, c.OpenAIModel.IgnoreTestConnection)
	assert.Equal(t, []string{"d1"}, c.OpenAIModel.DelegateSelectors)

	auth := c.OpenAIModel.Authentication
	assert.Equal(t, OpenAIModelAuthTypes.Token, auth.Type_)
	assert.NotNil(t, auth.Token)
	assert.Equal(t, "account.openai_key", auth.Token.TokenRef)
}

func TestUnitOpenAIModel_UnmarshalVertex(t *testing.T) {
	jsonPayload := `{
		"connector": {
			"name": "test-openai-vertex",
			"identifier": "test_openai_vertex",
			"type": "OpenAI",
			"spec": {
				"baseUrl": "https://api.openai.com",
				"model": "gpt-4o",
				"executeOnDelegate": true,
				"ignoreTestConnection": true,
				"authentication": {
					"type": "Vertex",
					"spec": {
						"serviceAccountKeyRef": "account.gcp_sa_key",
						"projectId": "my-gcp-project",
						"region": "us-central1"
					}
				}
			}
		}
	}`

	type envelope struct {
		Connector ConnectorInfo `json:"connector"`
	}

	var env envelope
	err := json.Unmarshal([]byte(jsonPayload), &env)
	require.NoError(t, err)

	c := env.Connector
	assert.Equal(t, ConnectorTypes.OpenAIModel, c.Type_)
	assert.NotNil(t, c.OpenAIModel)

	auth := c.OpenAIModel.Authentication
	assert.Equal(t, OpenAIModelAuthTypes.Vertex, auth.Type_)
	assert.NotNil(t, auth.Vertex)
	assert.Equal(t, "account.gcp_sa_key", auth.Vertex.ServiceAccountKeyRef)
	assert.Equal(t, "my-gcp-project", auth.Vertex.ProjectId)
	assert.Equal(t, "us-central1", auth.Vertex.Region)
}

func TestUnitOpenAIModel_MarshalRoundTrip(t *testing.T) {
	original := &ConnectorInfo{
		Name:       "rt-openai",
		Identifier: "rt_openai",
		Type_:      ConnectorTypes.OpenAIModel,
		OpenAIModel: &OpenAIModelConnector{
			Url:               "https://api.openai.com",
			Model:             "gpt-4o",
			ExecuteOnDelegate: true,
			Authentication: &OpenAIModelAuthentication{
				Type_: OpenAIModelAuthTypes.Token,
				Token: &OpenAIModelTokenSpec{
					TokenRef: "account.my_token",
				},
			},
		},
	}

	data, err := json.Marshal(original)
	require.NoError(t, err)

	var decoded ConnectorInfo
	err = json.Unmarshal(data, &decoded)
	require.NoError(t, err)

	assert.Equal(t, original.Name, decoded.Name)
	assert.Equal(t, original.Identifier, decoded.Identifier)
	assert.Equal(t, ConnectorTypes.OpenAIModel, decoded.Type_)
	require.NotNil(t, decoded.OpenAIModel)
	assert.Equal(t, "https://api.openai.com", decoded.OpenAIModel.Url)
	assert.Equal(t, "gpt-4o", decoded.OpenAIModel.Model)
	assert.Equal(t, OpenAIModelAuthTypes.Token, decoded.OpenAIModel.Authentication.Type_)
	assert.Equal(t, "account.my_token", decoded.OpenAIModel.Authentication.Token.TokenRef)
}

func TestUnitOpenAIModel_UnmarshalRealAPIResponse(t *testing.T) {
	apiResponse := `{
		"status": "SUCCESS",
		"data": {
			"connector": {
				"name": "my-openai",
				"identifier": "my_openai",
				"description": "OpenAI connector",
				"orgIdentifier": "default",
				"projectIdentifier": "proj1",
				"tags": {"env": "test"},
				"type": "OpenAI",
				"spec": {
					"baseUrl": "https://api.openai.com",
					"model": "gpt-4o",
					"delegateSelectors": ["mydelegate"],
					"executeOnDelegate": true,
					"ignoreTestConnection": true,
					"authentication": {
						"type": "Token",
						"spec": {
							"tokenRef": "org.openai_secret"
						}
					}
				}
			}
		}
	}`

	type dataWrapper struct {
		Connector ConnectorInfo `json:"connector"`
	}
	type apiResp struct {
		Status string      `json:"status"`
		Data   dataWrapper `json:"data"`
	}

	var resp apiResp
	err := json.Unmarshal([]byte(apiResponse), &resp)
	require.NoError(t, err)

	assert.Equal(t, "SUCCESS", resp.Status)
	c := resp.Data.Connector
	assert.Equal(t, "my-openai", c.Name)
	assert.Equal(t, "my_openai", c.Identifier)
	assert.Equal(t, ConnectorTypes.OpenAIModel, c.Type_)
	assert.NotNil(t, c.OpenAIModel)
	assert.Equal(t, "https://api.openai.com", c.OpenAIModel.Url)
	assert.Equal(t, "gpt-4o", c.OpenAIModel.Model)
	assert.Equal(t, []string{"mydelegate"}, c.OpenAIModel.DelegateSelectors)
	assert.True(t, c.OpenAIModel.ExecuteOnDelegate)
	assert.True(t, c.OpenAIModel.IgnoreTestConnection)
	assert.Equal(t, OpenAIModelAuthTypes.Token, c.OpenAIModel.Authentication.Type_)
	assert.Equal(t, "org.openai_secret", c.OpenAIModel.Authentication.Token.TokenRef)
}

func TestUnitOpenAIModel_CreateConnectorRequest(t *testing.T) {
	connector := ConnectorInfo{
		Name:        "openai-test",
		Identifier:  "openai_test",
		Description: "unit test connector",
		Type_:       ConnectorTypes.OpenAIModel,
		OpenAIModel: &OpenAIModelConnector{
			Url:               "https://api.openai.com",
			Model:             "gpt-4o",
			ExecuteOnDelegate: true,
			DelegateSelectors: []string{"d1"},
			Authentication: &OpenAIModelAuthentication{
				Type_: OpenAIModelAuthTypes.Vertex,
				Vertex: &OpenAIModelVertexSpec{
					ServiceAccountKeyRef: "account.gcp_key",
					ProjectId:            "my-project",
					Region:               "us-central1",
				},
			},
		},
	}

	data, err := json.Marshal(connector)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "https://app.harness.io/gateway/ng/api/connectors", nil)
	require.NoError(t, err)
	assert.NotNil(t, req)
	assert.NotEmpty(t, data)
}
