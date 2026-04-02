package abac

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

func GetExchangeKeyInformationDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "key",
			Type: "string",
		},
		{
			Name: prefix + "visibility",
			Type: "string",
		},
	}
}
func CastExchangeKeyInformationDtoFromCli(c emigo.CliCastable) ExchangeKeyInformationDto {
	data := ExchangeKeyInformationDto{}
	if c.IsSet("key") {
		data.Key = c.String("key")
	}
	if c.IsSet("visibility") {
		data.Visibility = c.String("visibility")
	}
	return data
}

// The base class definition for exchangeKeyInformationDto
type ExchangeKeyInformationDto struct {
	Key        string `json:"key" yaml:"key"`
	Visibility string `json:"visibility" yaml:"visibility"`
}

func (x *ExchangeKeyInformationDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
