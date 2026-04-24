package fireback

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

func GetDeleteRequestDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "query",
			Type: "string",
		},
	}
}
func CastDeleteRequestDtoFromCli(c emigo.CliCastable) DeleteRequestDto {
	data := DeleteRequestDto{}
	if c.IsSet("query") {
		data.Query = c.String("query")
	}
	return data
}

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
