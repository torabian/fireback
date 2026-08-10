package abacdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

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
func GetOkayResponseDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{}
}
func CastOkayResponseDtoFromCli(c emigo.CliCastable) OkayResponseDto {
	data := OkayResponseDto{}
	return data
}
