package abac

import "encoding/json"

// The base class definition for importRequestDto
type ImportRequestDto struct {
	File string `json:"file" yaml:"file"`
}

func (x *ImportRequestDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
