package abac

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

func GetTestMailDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "sender-id",
			Type: "string",
		},
		{
			Name: prefix + "to-name",
			Type: "string",
		},
		{
			Name: prefix + "to-email",
			Type: "string",
		},
		{
			Name: prefix + "subject",
			Type: "string",
		},
		{
			Name: prefix + "content",
			Type: "string",
		},
	}
}
func CastTestMailDtoFromCli(c emigo.CliCastable) TestMailDto {
	data := TestMailDto{}
	if c.IsSet("sender-id") {
		data.SenderId = c.String("sender-id")
	}
	if c.IsSet("to-name") {
		data.ToName = c.String("to-name")
	}
	if c.IsSet("to-email") {
		data.ToEmail = c.String("to-email")
	}
	if c.IsSet("subject") {
		data.Subject = c.String("subject")
	}
	if c.IsSet("content") {
		data.Content = c.String("content")
	}
	return data
}

// The base class definition for testMailDto
type TestMailDto struct {
	SenderId string `json:"senderId" yaml:"senderId"`
	ToName   string `json:"toName" yaml:"toName"`
	ToEmail  string `json:"toEmail" yaml:"toEmail"`
	Subject  string `json:"subject" yaml:"subject"`
	Content  string `json:"content" yaml:"content"`
}

func (x *TestMailDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
