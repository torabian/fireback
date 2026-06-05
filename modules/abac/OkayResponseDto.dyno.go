package abac

import "encoding/json"

// The base class definition for okayResponseDto
type OkayResponseDto struct {
}

func (x *OkayResponseDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
