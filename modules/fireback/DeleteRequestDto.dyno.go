package fireback

import "encoding/json"

// The base class definition for deleteRequestDto
type DeleteRequestDto struct {
	// The query selector which would be used to delete the content.
	Query string `json:"query" yaml:"query"`
}

func (x *DeleteRequestDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
