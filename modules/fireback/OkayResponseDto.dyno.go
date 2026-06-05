package fireback

import "encoding/json"

// import emigo "github.com/torabian/emi/emigo"

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
