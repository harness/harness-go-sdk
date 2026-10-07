package nextgen

import (
	"encoding/json"
	"fmt"
)

type OpenAIModelConnector struct {
	Authentication       *OpenAIModelAuthentication `json:"authentication"`
	Url                  string                     `json:"baseUrl,omitempty"`
	DelegateSelectors    []string                   `json:"delegateSelectors,omitempty"`
	ExecuteOnDelegate    bool                       `json:"executeOnDelegate"`
	IgnoreTestConnection bool                       `json:"ignoreTestConnection"`
	Model                string                     `json:"model,omitempty"`
	HarnessManaged       bool                       `json:"harnessManaged"`
}

type OpenAIModelAuthType string

var OpenAIModelAuthTypes = struct {
	Token  OpenAIModelAuthType
	Vertex OpenAIModelAuthType
}{
	Token:  "Token",
	Vertex: "Vertex",
}

type OpenAIModelAuthentication struct {
	Type_  OpenAIModelAuthType `json:"type"`
	Token  *OpenAIModelTokenSpec  `json:"-"`
	Vertex *OpenAIModelVertexSpec `json:"-"`
	Spec   json.RawMessage       `json:"spec,omitempty"`
}

type OpenAIModelTokenSpec struct {
	TokenRef string `json:"tokenRef"`
}

type OpenAIModelVertexSpec struct {
	ServiceAccountKeyRef string `json:"serviceAccountKeyRef"`
	ProjectId            string `json:"projectId"`
	Region               string `json:"region"`
}

func (a *OpenAIModelAuthentication) UnmarshalJSON(data []byte) error {
	type Alias OpenAIModelAuthentication

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
	case OpenAIModelAuthTypes.Token:
		err = json.Unmarshal(aux.Spec, &a.Token)
	case OpenAIModelAuthTypes.Vertex:
		err = json.Unmarshal(aux.Spec, &a.Vertex)
	default:
		panic(fmt.Sprintf("unknown openai model auth type %s", a.Type_))
	}

	return err
}

func (a *OpenAIModelAuthentication) MarshalJSON() ([]byte, error) {
	type Alias OpenAIModelAuthentication

	var spec []byte
	var err error

	switch a.Type_ {
	case OpenAIModelAuthTypes.Token:
		spec, err = json.Marshal(a.Token)
	case OpenAIModelAuthTypes.Vertex:
		spec, err = json.Marshal(a.Vertex)
	default:
		panic(fmt.Sprintf("unknown openai model auth type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
