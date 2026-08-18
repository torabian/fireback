package abacdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for capabilityOptionalDto
type CapabilityOptionalDto struct {
	UniqueId    emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Name        emigo.Nullable[string] `json:"name" yaml:"name"`
	Description emigo.Nullable[string] `json:"description" yaml:"description"`
}

func (x *CapabilityOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
