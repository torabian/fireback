package abacdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for capabilityDto
type CapabilityDto struct {
	UniqueId    emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Name        string                 `json:"name" yaml:"name"`
	Description string                 `json:"description" yaml:"description"`
}

func (x *CapabilityDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
