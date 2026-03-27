package abac

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

func GetTestMail4DtoCliFlags(prefix string) []emigo.CliFlag {
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
func CastTestMail4DtoFromCli(c emigo.CliCastable) TestMail4Dto {
	data := TestMail4Dto{}
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

// The base class definition for testMail4Dto
type TestMail4Dto struct {
	SenderId string `yaml:"senderId" json:"senderId"`
	ToName   string `yaml:"toName" json:"toName"`
	ToEmail  string `json:"toEmail" yaml:"toEmail"`
	Subject  string `json:"subject" yaml:"subject"`
	Content  string `json:"content" yaml:"content"`
}

func (x *TestMail4Dto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
