package fireback

import (
	"encoding/json"
)

// The base class definition for reactiveSearchResultDto
type ReactiveSearchResultDto struct {
	UniqueId    string `json:"uniqueId" yaml:"uniqueId"`
	Phrase      string `json:"phrase" yaml:"phrase"`
	Icon        string `json:"icon" yaml:"icon"`
	Description string `json:"description" yaml:"description"`
	Group       string `json:"group" yaml:"group"`
	UiLocation  string `json:"uiLocation" yaml:"uiLocation"`
	ActionFn    string `json:"actionFn" yaml:"actionFn"`
}

func (x *ReactiveSearchResultDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
