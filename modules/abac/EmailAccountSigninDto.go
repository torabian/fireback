package abac

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

func GetEmailAccountSigninDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "email",
			Type: "string",
		},
		{
			Name: prefix + "password",
			Type: "string",
		},
	}
}
func CastEmailAccountSigninDtoFromCli(c emigo.CliCastable) EmailAccountSigninDto {
	data := EmailAccountSigninDto{}
	if c.IsSet("email") {
		data.Email = c.String("email")
	}
	if c.IsSet("password") {
		data.Password = c.String("password")
	}
	return data
}

// The base class definition for emailAccountSigninDto
type EmailAccountSigninDto struct {
	Email    string `validate:"required" json:"email" yaml:"email"`
	Password string `validate:"required" json:"password" yaml:"password"`
}

func (x *EmailAccountSigninDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
