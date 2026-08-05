package fireback

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback/complexes"
)

// The base class definition for webPushConfigOptionalDto
type WebPushConfigOptionalDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	// The json content of the web push after getting it from browser
	Subscription complexes.JSON `json:"subscription" yaml:"subscription"`
}

func (x *WebPushConfigOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWebPushConfigOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name:        prefix + "subscription",
			Type:        "complex",
			Description: "The json content of the web push after getting it from browser",
		},
	}
}
func CastWebPushConfigOptionalDtoFromCli(c emigo.CliCastable) WebPushConfigOptionalDto {
	data := WebPushConfigOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("subscription") {
		if u, ok := any(&data.Subscription).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("subscription")))
		}
	}
	return data
}
