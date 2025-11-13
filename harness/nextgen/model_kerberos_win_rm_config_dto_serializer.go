package nextgen

import (
	"encoding/json"
)

func (a *KerberosWinRmConfigDto) UnmarshalJSON(data []byte) error {

	type Alias KerberosWinRmConfigDto

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	// Unmarshal the spec into the appropriate typed field based on TgtGenerationMethod
	if len(a.Spec) > 0 && a.TgtGenerationMethod != "" {
		switch a.TgtGenerationMethod {
		case string(TgtGenerationMethodTypes.TGTKeyTabFilePathSpecDTO):
			var keyTabSpec TgtKeyTabFilePathSpecDto
			err = json.Unmarshal(a.Spec, &keyTabSpec)
			if err == nil {
				a.KeyTabFilePathSpec = &keyTabSpec
			}
		case string(TgtGenerationMethodTypes.TGTPasswordSpecDTO):
			var passwordSpec TgtPasswordSpecDto
			err = json.Unmarshal(a.Spec, &passwordSpec)
			if err == nil {
				a.PasswordSpec = &passwordSpec
			}
		}
	}

	return err
}

func (a *KerberosWinRmConfigDto) MarshalJSON() ([]byte, error) {
	type Alias KerberosWinRmConfigDto

	// Marshal the appropriate typed field into Spec based on TgtGenerationMethod
	if a.TgtGenerationMethod != "" {
		var spec []byte
		var err error

		switch a.TgtGenerationMethod {
		case string(TgtGenerationMethodTypes.TGTKeyTabFilePathSpecDTO):
			if a.KeyTabFilePathSpec != nil {
				spec, err = json.Marshal(a.KeyTabFilePathSpec)
			}
		case string(TgtGenerationMethodTypes.TGTPasswordSpecDTO):
			if a.PasswordSpec != nil {
				spec, err = json.Marshal(a.PasswordSpec)
			}
		}

		if err != nil {
			return nil, err
		}

		if spec != nil {
			a.Spec = json.RawMessage(spec)
		}
	}

	return json.Marshal((*Alias)(a))
}
