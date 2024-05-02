package workspaces
import (
	"encoding/json"
	"fmt"
	"strings"
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
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x* ResetEmailDto) JsonPrint()  {
    fmt.Println(x.Json())
}
// This is an experimental way to create new dtos, with exluding the pointers as helper.
func NewResetEmailDto(
	Password string,
) ResetEmailDto {
    return ResetEmailDto{
	Password: &Password,
    }
}