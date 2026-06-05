package fireback

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for capabilityInfoDto
type CapabilityInfoDto struct {
	UniqueId string                              `json:"uniqueId" yaml:"uniqueId"`
	Name     string                              `json:"name" yaml:"name"`
	Children emigo.Collection[CapabilityInfoDto] `json:"children" yaml:"children"`
}

func (x *CapabilityInfoDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
