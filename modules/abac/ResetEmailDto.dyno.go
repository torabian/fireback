package abac

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

func GetResetEmailDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "password",
			Type: "string",
		},
	}
}
func CastResetEmailDtoFromCli(c emigo.CliCastable) ResetEmailDto {
	data := ResetEmailDto{}
	if c.IsSet("password") {
		data.Password = c.String("password")
	}
	return data
}

// The base class definition for resetEmailDto
type ResetEmailDto struct {
	Password string `json:"password" yaml:"password"`
}

func (x *ResetEmailDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
