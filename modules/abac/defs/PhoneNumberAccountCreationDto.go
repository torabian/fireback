package abacdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for phoneNumberAccountCreationDto
type PhoneNumberAccountCreationDto struct {
	PhoneNumber string `json:"phoneNumber" yaml:"phoneNumber"`
}

func (x *PhoneNumberAccountCreationDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetPhoneNumberAccountCreationDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "phone-number",
			Type: "string",
		},
	}
}
func CastPhoneNumberAccountCreationDtoFromCli(c emigo.CliCastable) PhoneNumberAccountCreationDto {
	data := PhoneNumberAccountCreationDto{}
	if c.IsSet("phone-number") {
		data.PhoneNumber = c.String("phone-number")
	}
	return data
}
