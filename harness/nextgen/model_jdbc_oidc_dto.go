package nextgen

import (
	"encoding/json"
	"fmt"
)

type JDBCOidcProviderType string

var JDBCOidcProviderTypes = struct {
	Gcp JDBCOidcProviderType
}{
	Gcp: "Gcp",
}

func (e JDBCOidcProviderType) String() string {
	return string(e)
}

type JdbcOidcDto struct {
	ProviderType JDBCOidcProviderType `json:"type"`
	GcpOidcSpec  *JdbcGcpOidcSpecDto  `json:"-"`
	Spec         json.RawMessage      `json:"spec,omitempty"`
}

type JdbcGcpOidcSpecDto struct {
	ProjectNumber       string `json:"projectNumber"`
	WorkloadPoolId      string `json:"workloadPoolId"`
	ProviderId          string `json:"providerId"`
	ServiceAccountEmail string `json:"serviceAccountEmail"`
}

func (o *JdbcOidcDto) UnmarshalJSON(data []byte) error {
	type Alias JdbcOidcDto
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(o),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	switch o.ProviderType {
	case JDBCOidcProviderTypes.Gcp:
		if o.Spec != nil {
			return json.Unmarshal(o.Spec, &o.GcpOidcSpec)
		}
	default:
		return fmt.Errorf("unknown jdbc oidc provider type %s", o.ProviderType)
	}

	return nil
}

func (o *JdbcOidcDto) MarshalJSON() ([]byte, error) {
	type Alias JdbcOidcDto

	var spec []byte
	var err error

	switch o.ProviderType {
	case JDBCOidcProviderTypes.Gcp:
		spec, err = json.Marshal(o.GcpOidcSpec)
	default:
		return nil, fmt.Errorf("unknown jdbc oidc provider type %s", o.ProviderType)
	}

	if err != nil {
		return nil, err
	}

	o.Spec = json.RawMessage(spec)
	return json.Marshal((*Alias)(o))
}
