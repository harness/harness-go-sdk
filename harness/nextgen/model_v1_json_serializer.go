package nextgen

import "encoding/json"

func (j *V1Json) UnmarshalJSON(data []byte) error {
	var wrapper struct {
		Raw string `json:"raw,omitempty"`
	}
	if err := json.Unmarshal(data, &wrapper); err == nil && wrapper.Raw != "" {
		j.Raw = wrapper.Raw
		return nil
	}
	j.Raw = string(data)
	return nil
}
