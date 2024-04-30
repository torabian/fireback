package workspaces
import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/urfave/cli"
)
func CastTestMailFromCli (c *cli.Context) *TestMailDto {
	template := &TestMailDto{}
      if c.IsSet("sender-id") {
        value := c.String("sender-id")
        template.SenderId = &value
      }
      if c.IsSet("to-name") {
        value := c.String("to-name")
        template.ToName = &value
      }
      if c.IsSet("to-email") {
        value := c.String("to-email")
        template.ToEmail = &value
      }
      if c.IsSet("subject") {
        value := c.String("subject")
        template.Subject = &value
      }
      if c.IsSet("content") {
        value := c.String("content")
        template.Content = &value
      }
	return template
}
var TestMailDtoCommonCliFlagsOptional = []cli.Flag{
  &cli.StringFlag{
    Name:     "wid",
    Required: false,
    Usage:    "Provide workspace id, if you want to change the data workspace",
  },
  &cli.StringFlag{
    Name:     "uid",
    Required: false,
    Usage:    "uniqueId (primary key)",
  },
  &cli.StringFlag{
    Name:     "pid",
    Required: false,
    Usage:    " Parent record id of the same type",
  },
    &cli.StringFlag{
      Name:     "sender-id",
      Required: false,
      Usage:    "senderId",
    },
    &cli.StringFlag{
      Name:     "to-name",
      Required: false,
      Usage:    "toName",
    },
    &cli.StringFlag{
      Name:     "to-email",
      Required: false,
      Usage:    "toEmail",
    },
    &cli.StringFlag{
      Name:     "subject",
      Required: false,
      Usage:    "subject",
    },
    &cli.StringFlag{
      Name:     "content",
      Required: false,
      Usage:    "content",
    },
}
type TestMailDto struct {
    SenderId   *string `json:"senderId" yaml:"senderId"       `
    // Datenano also has a text representation
    ToName   *string `json:"toName" yaml:"toName"       `
    // Datenano also has a text representation
    ToEmail   *string `json:"toEmail" yaml:"toEmail"       `
    // Datenano also has a text representation
    Subject   *string `json:"subject" yaml:"subject"       `
    // Datenano also has a text representation
    Content   *string `json:"content" yaml:"content"       `
    // Datenano also has a text representation
}
func (x* TestMailDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x* TestMailDto) JsonPrint()  {
    fmt.Println(x.Json())
}
