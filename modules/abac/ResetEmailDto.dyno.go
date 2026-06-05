package abac

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for resetEmailDto
type ResetEmailDto struct {
	Password string `json:"password" yaml:"password"`
}

func (x *ResetEmailDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
