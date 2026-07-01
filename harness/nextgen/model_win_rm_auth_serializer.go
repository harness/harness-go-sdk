package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *WinRmAuth) UnmarshalJSON(data []byte) error {

	type Alias WinRmAuth

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
	case "Kerberos":
		err = json.Unmarshal(aux.Spec, &a.KerberosConfig)
	case "NTLM":
		err = json.Unmarshal(aux.Spec, &a.NtlmConfig)
	default:
		panic(fmt.Sprintf("unknown WinRM authentication type %s", a.Type_))
	}

	return err
}

func (a *WinRmAuth) MarshalJSON() ([]byte, error) {
	type Alias WinRmAuth

	var spec []byte
	var err error

	switch a.Type_ {
	case "Kerberos":
		spec, err = json.Marshal(a.KerberosConfig)
	case "NTLM":
		spec, err = json.Marshal(a.NtlmConfig)
	default:
		panic(fmt.Sprintf("unknown WinRM authentication type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
