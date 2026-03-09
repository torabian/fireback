package abac

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

func GetEmailOtpResponseDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "request",
			Type: "one",
		},
		{
			Name: prefix + "user-session",
			Type: "one",
		},
	}
}
func CastEmailOtpResponseDtoFromCli(c emigo.CliCastable) EmailOtpResponseDto {
	data := EmailOtpResponseDto{}
	return data
}

// The base class definition for emailOtpResponseDto
type EmailOtpResponseDto struct {
	Request     PublicAuthenticationEntity `yaml:"request" json:"request"`
	UserSession UserSessionDto             `yaml:"userSession" json:"userSession"`
}

func (x *EmailOtpResponseDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
