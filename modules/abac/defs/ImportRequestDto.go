package abacdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

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
