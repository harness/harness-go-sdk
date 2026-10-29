package nextgen

import (
	"encoding/json"
)

func (f *Failure) UnmarshalJSON(data []byte) error {
	type Alias Failure

	aux := &struct {
		*Alias
		Code json.RawMessage `json:"code,omitempty"`
	}{
		Alias: (*Alias)(f),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	if len(aux.Code) == 0 {
		return nil
	}

	if err := json.Unmarshal(aux.Code, &f.Code); err == nil {
		return nil
	}

	var number json.Number
	if err := json.Unmarshal(aux.Code, &number); err != nil {
		return err
	}
	f.Code = number.String()

	return nil
}
