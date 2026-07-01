package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *JdbcAuthenticationDto) UnmarshalJSON(data []byte) error {

	type Alias JdbcAuthenticationDto

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
	case JDBCAuthTypes.UsernamePassword:
		err = json.Unmarshal(aux.Spec, &a.UsernamePassword)
	case JDBCAuthTypes.ServiceAccount:
		err = json.Unmarshal(aux.Spec, &a.ServiceAccount)
	case JDBCAuthTypes.KeyPair:
		err = json.Unmarshal(aux.Spec, &a.KeyPair)
	case JDBCAuthTypes.InheritFromDelegate:
		a.InheritFromDelegate = &JdbcDelegateAccessDto{}
		if len(aux.Spec) > 0 {
			err = json.Unmarshal(aux.Spec, a.InheritFromDelegate)
		}
	case JDBCAuthTypes.Oidc:
		err = json.Unmarshal(aux.Spec, &a.Oidc)
	default:
		panic(fmt.Sprintf("unknown jdbc auth method type %s", a.Type_))
	}

	return err
}

func (a *JdbcAuthenticationDto) MarshalJSON() ([]byte, error) {
	type Alias JdbcAuthenticationDto

	var spec []byte
	var err error

	switch a.Type_ {
	case JDBCAuthTypes.UsernamePassword:
		spec, err = json.Marshal(a.UsernamePassword)
	case JDBCAuthTypes.ServiceAccount:
		spec, err = json.Marshal(a.ServiceAccount)
	case JDBCAuthTypes.KeyPair:
		spec, err = json.Marshal(a.KeyPair)
	case JDBCAuthTypes.InheritFromDelegate:
		if a.InheritFromDelegate != nil {
			spec, err = json.Marshal(a.InheritFromDelegate)
		} else {
			spec = []byte("{}")
		}
	case JDBCAuthTypes.Oidc:
		spec, err = json.Marshal(a.Oidc)
	default:
		panic(fmt.Sprintf("unknown jdbc auth method type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
