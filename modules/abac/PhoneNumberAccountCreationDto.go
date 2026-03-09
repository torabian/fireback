package abac

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

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

// The base class definition for phoneNumberAccountCreationDto
type PhoneNumberAccountCreationDto struct {
	PhoneNumber string `yaml:"phoneNumber" json:"phoneNumber"`
}

func (x *PhoneNumberAccountCreationDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
