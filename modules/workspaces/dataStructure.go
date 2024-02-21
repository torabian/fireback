package workspaces

import (
	"encoding/json"
)

func Flatten(m map[string]interface{}) map[string]interface{} {
	o := make(map[string]interface{})
	for k, v := range m {
		switch child := v.(type) {
		case map[string]interface{}:
			nm := Flatten(child)
			for nk, nv := range nm {
				o[k+"."+nk] = nv
			}
		default:
			o[k] = v
		}
	}
	return o
}

func CastJsonDataTypeTo[T any](x *JSON) *T {

	var d T
	if x == nil {
		return nil
	}

	data, err := x.MarshalJSON()

	if err != nil {
		return nil
	}

	json.Unmarshal(data, &d)

	return &d

}
