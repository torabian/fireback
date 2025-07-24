package fireback

import (
	"encoding/json"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func MarshalWithWhitelist(obj any, allowed []string) ([]byte, error) {
	original, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	// Remove nulls
	var m map[string]interface{}
	if err := json.Unmarshal(original, &m); err != nil {
		return nil, err
	}

	// Re-marshal after pruning nulls
	cleaned, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	// Build whitelisted JSON
	result := "{}"
	for _, path := range allowed {
		val := gjson.GetBytes(cleaned, path)
		if val.Exists() {
			result, err = sjson.SetRaw(result, path, val.Raw)
			if err != nil {
				return nil, err
			}
		}
	}

	return []byte(result), nil
}

func removeNulls(m map[string]interface{}) {
	for k, v := range m {
		switch val := v.(type) {
		case nil:
			delete(m, k)
		case map[string]interface{}:
			removeNulls(val)
		case []interface{}:
			for _, item := range val {
				if subMap, ok := item.(map[string]interface{}); ok {
					removeNulls(subMap)
				}
			}
		}
	}
}
