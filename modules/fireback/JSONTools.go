package fireback

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/itchyny/gojq"
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

func SearchJqInArray[T any](queryJson string, items []T) ([]T, error) {

	if queryJson == "" {
		return items, nil
	}

	prejson, _ := json.MarshalIndent(items, "", "  ")

	var input interface{}
	if err := json.Unmarshal(prejson, &input); err != nil {
		log.Fatal(err)
	}

	res := []T{}
	query, err := gojq.Parse(queryJson)
	if err != nil {
		log.Fatal(err)
	}

	iter := query.Run(input)

	for {
		v, ok := iter.Next()
		if !ok {
			break
		}

		if err, isErr := v.(error); isErr {
			return nil, err
		}

		// If the query result is a slice, unpack it
		if arr, isSlice := v.([]interface{}); isSlice {
			for _, item := range arr {
				// Marshal & unmarshal to convert from `interface{}` to `T`
				data, _ := json.Marshal(item)
				var out T
				json.Unmarshal(data, &out)
				res = append(res, out)
			}
		} else {
			// Fallback single item
			data, _ := json.Marshal(v)
			var out T
			json.Unmarshal(data, &out)
			res = append(res, out)
		}
	}

	return res, nil
}

// Paginates an array slice, and creates QRM object out of it, and returns the items
func QueryArrayInMemory[T any](items []T, q QueryDSL) ([]T, *QueryResultMeta, *IError) {
	totalAvailableItems := len(items)
	items, err := SearchJqInArray(q.JqQuery, items)

	if err != nil {
		return nil, nil, CastToIError(err)
	}

	cursorIn := q.Cursor
	limit := q.ItemsPerPage

	start := 0

	if cursorIn != nil && *cursorIn != "" {
		if strings.HasPrefix(*cursorIn, "index(") && strings.HasSuffix(*cursorIn, ")") {
			inner := strings.TrimSuffix(strings.TrimPrefix(*cursorIn, "index("), ")")
			if i, err := strconv.Atoi(inner); err == nil && i >= 0 {
				start = i
			}
		}
	}

	if start >= len(items) {
		return []T{}, &QueryResultMeta{
			TotalItems:          int64(len(items)),
			TotalAvailableItems: int64(totalAvailableItems),
			Cursor:              nil,
		}, nil
	}

	end := start + limit
	if end > len(items) {
		end = len(items)
	}

	var nextCursor *string
	if end < len(items) {
		cursorStr := fmt.Sprintf("index(%d)", end)
		nextCursor = &cursorStr
	}

	return items[start:end], &QueryResultMeta{
		TotalItems:          int64(len(items[start:end])),
		TotalAvailableItems: int64(len(items)),
		Cursor:              nextCursor,
	}, nil
}
