package abac

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

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
