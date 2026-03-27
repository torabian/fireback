package izaTheme

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

func GetProductSummaryDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "name",
			Type: "string",
		},
	}
}
func CastProductSummaryDtoFromCli(c emigo.CliCastable) ProductSummaryDto {
	data := ProductSummaryDto{}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}

// The base class definition for productSummaryDto
type ProductSummaryDto struct {
	// Product name, used in titles.
	Name string `json:"name" yaml:"name"`
}

func (x *ProductSummaryDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
