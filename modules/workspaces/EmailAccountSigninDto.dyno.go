package workspaces
import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
)
func CastEmailAccountSigninFromCli (c *cli.Context) *EmailAccountSigninDto {
	template := &EmailAccountSigninDto{}
      if c.IsSet("email") {
        value := c.String("email")
        template.Email = &value
      }
      if c.IsSet("password") {
        value := c.String("password")
        template.Password = &value
      }
	return template
}
var EmailAccountSigninDtoCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "email",
      Required: true,
      Usage:    "email",
    },
    &cli.StringFlag{
      Name:     "password",
      Required: true,
      Usage:    "password",
    },
}
type EmailAccountSigninDto struct {
    Email   *string `json:"email" yaml:"email"  validate:"required"       `
    // Datenano also has a text representation
    Password   *string `json:"password" yaml:"password"  validate:"required"       `
    // Datenano also has a text representation
}
func (x* EmailAccountSigninDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
func (x* EmailAccountSigninDto) JsonPrint()  {
    fmt.Println(x.Json())
}
