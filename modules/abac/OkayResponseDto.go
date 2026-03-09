package abac

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

func GetOkayResponseDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "status",
			Type: "int",
		},
	}
}
func CastOkayResponseDtoFromCli(c emigo.CliCastable) OkayResponseDto {
	data := OkayResponseDto{}
	if c.IsSet("status") {
		data.Status = int(c.Int64("status"))
	}
	return data
}

// The base class definition for okayResponseDto
type OkayResponseDto struct {
	Status int `json:"status" yaml:"status"`
}

func (x *OkayResponseDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
