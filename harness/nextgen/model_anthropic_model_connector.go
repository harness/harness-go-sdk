package nextgen

import (
	"encoding/json"
	"fmt"
)

// Connector type enum value: "Anthropic"
// This is the Anthropic Model Connector (Harness calls Claude API).
// Distinct from AnthropicCost (billing) and Anthropic Harness Connector (MCP/OAuth).

type AnthropicModelConnector struct {
	Authentication       *AnthropicModelAuthentication `json:"authentication"`
	Url                  string                        `json:"url,omitempty"`
	DelegateSelectors    []string                      `json:"delegateSelectors,omitempty"`
	ExecuteOnDelegate    bool                          `json:"executeOnDelegate"`
	IgnoreTestConnection bool                          `json:"ignoreTestConnection"`
	Model                string                        `json:"model,omitempty"`
	HarnessManaged       bool                          `json:"harnessManaged"`
}

// Authentication

type AnthropicModelAuthType string

var AnthropicModelAuthTypes = struct {
	Token         AnthropicModelAuthType
	BedrockApiKey AnthropicModelAuthType
	Vertex        AnthropicModelAuthType
	CloudProvider AnthropicModelAuthType
}{
	Token:         "Token",
	BedrockApiKey: "BedrockApiKey",
	Vertex:        "Vertex",
	CloudProvider: "CloudProvider",
}

type AnthropicModelAuthentication struct {
	Type_         AnthropicModelAuthType           `json:"type"`
	Token         *AnthropicModelTokenSpec         `json:"-"`
	BedrockApiKey *AnthropicModelBedrockApiKeySpec `json:"-"`
	Vertex        *AnthropicModelVertexSpec        `json:"-"`
	CloudProvider *AnthropicModelCloudProviderSpec `json:"-"`
	Spec          json.RawMessage                  `json:"spec,omitempty"`
}

type AnthropicModelTokenSpec struct {
	TokenRef string `json:"tokenRef"`
}

type AnthropicModelBedrockApiKeySpec struct {
	ApiKeyRef string `json:"apiKeyRef"`
	Region    string `json:"region"`
}

type AnthropicModelVertexSpec struct {
	ServiceAccountKeyRef string `json:"serviceAccountKeyRef"`
	ProjectId            string `json:"projectId"`
	Region               string `json:"region"`
}

// Cloud provider (references an existing AWS/GCP/Azure connector)

type AnthropicModelCloudProviderType string

var AnthropicModelCloudProviderTypes = struct {
	AWS   AnthropicModelCloudProviderType
	GCP   AnthropicModelCloudProviderType
	Azure AnthropicModelCloudProviderType
}{
	AWS:   "AWS",
	GCP:   "GCP",
	Azure: "Azure",
}

type AnthropicModelCloudProviderSpec struct {
	Type_        AnthropicModelCloudProviderType `json:"type"`
	ConnectorRef string                          `json:"connectorRef"`
}

// Serialization

func (a *AnthropicModelAuthentication) UnmarshalJSON(data []byte) error {
	type Alias AnthropicModelAuthentication

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	switch a.Type_ {
	case AnthropicModelAuthTypes.Token:
		err = json.Unmarshal(aux.Spec, &a.Token)
	case AnthropicModelAuthTypes.BedrockApiKey:
		err = json.Unmarshal(aux.Spec, &a.BedrockApiKey)
	case AnthropicModelAuthTypes.Vertex:
		err = json.Unmarshal(aux.Spec, &a.Vertex)
	case AnthropicModelAuthTypes.CloudProvider:
		err = json.Unmarshal(aux.Spec, &a.CloudProvider)
	default:
		panic(fmt.Sprintf("unknown anthropic model auth type %s", a.Type_))
	}

	return err
}

func (a *AnthropicModelAuthentication) MarshalJSON() ([]byte, error) {
	type Alias AnthropicModelAuthentication

	var spec []byte
	var err error

	switch a.Type_ {
	case AnthropicModelAuthTypes.Token:
		spec, err = json.Marshal(a.Token)
	case AnthropicModelAuthTypes.BedrockApiKey:
		spec, err = json.Marshal(a.BedrockApiKey)
	case AnthropicModelAuthTypes.Vertex:
		spec, err = json.Marshal(a.Vertex)
	case AnthropicModelAuthTypes.CloudProvider:
		spec, err = json.Marshal(a.CloudProvider)
	default:
		panic(fmt.Sprintf("unknown anthropic model auth type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
