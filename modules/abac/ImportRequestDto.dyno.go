package abac

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

func GetImportRequestDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "file",
			Type: "string",
		},
	}
}
func CastImportRequestDtoFromCli(c emigo.CliCastable) ImportRequestDto {
	data := ImportRequestDto{}
	if c.IsSet("file") {
		data.File = c.String("file")
	}
	return data
}

// The base class definition for importRequestDto
type ImportRequestDto struct {
	File string `json:"file" yaml:"file"`
}

func (x *ImportRequestDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
