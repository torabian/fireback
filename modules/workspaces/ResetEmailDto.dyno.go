package workspaces
import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
)
func CastResetEmailFromCli (c *cli.Context) *ResetEmailDto {
	template := &ResetEmailDto{}
      if c.IsSet("password") {
        value := c.String("password")
        template.Password = &value
      }
	return template
}
var ResetEmailDtoCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "password",
      Required: false,
      Usage:    "password",
    },
}
type ResetEmailDto struct {
    Password   *string `json:"password" yaml:"password"       `
    // Datenano also has a text representation
}
func (x* ResetEmailDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
func (x* ResetEmailDto) JsonPrint()  {
    fmt.Println(x.Json())
}
