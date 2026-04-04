package abac

import (
	"encoding/json"

	emigo "github.com/torabian/emi/emigo"
)

func GetOkayResponseDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{}
}
func CastOkayResponseDtoFromCli(c emigo.CliCastable) OkayResponseDto {
	data := OkayResponseDto{}
	return data
}

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
