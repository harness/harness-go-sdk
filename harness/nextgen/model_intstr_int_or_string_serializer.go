package nextgen

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func (i *IntstrIntOrString) UnmarshalJSON(data []byte) error {
	var raw interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch v := raw.(type) {
	case float64:
		i.Type_ = "0"
		i.IntVal = int32(v)
	case string:
		if n, err := strconv.ParseInt(v, 10, 32); err == nil {
			i.Type_ = "0"
			i.IntVal = int32(n)
		} else {
			i.Type_ = "1"
			i.StrVal = v
		}
	case map[string]interface{}:
		if t, ok := v["type"]; ok {
			i.Type_ = fmt.Sprintf("%v", t)
		}
		if iv, ok := v["intVal"]; ok {
			if f, ok := iv.(float64); ok {
				i.IntVal = int32(f)
			}
		}
		if sv, ok := v["strVal"]; ok {
			if s, ok := sv.(string); ok {
				i.StrVal = s
			}
		}
	default:
		return fmt.Errorf("IntstrIntOrString: unsupported type %T", raw)
	}
	return nil
}
