package abac

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

func GetOtpAuthenticateDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "value",
			Type: "string",
		},
		{
			Name: prefix + "otp",
			Type: "string",
		},
		{
			Name: prefix + "type",
			Type: "string",
		},
		{
			Name: prefix + "password",
			Type: "string",
		},
	}
}
func CastOtpAuthenticateDtoFromCli(c emigo.CliCastable) OtpAuthenticateDto {
	data := OtpAuthenticateDto{}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	if c.IsSet("otp") {
		data.Otp = c.String("otp")
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	if c.IsSet("password") {
		data.Password = c.String("password")
	}
	return data
}

// The base class definition for otpAuthenticateDto
type OtpAuthenticateDto struct {
	Value    string `yaml:"value" validate:"required" json:"value"`
	Otp      string `json:"otp" yaml:"otp"`
	Type     string `validate:"required" json:"type" yaml:"type"`
	Password string `validate:"required" json:"password" yaml:"password"`
}

func (x *OtpAuthenticateDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
