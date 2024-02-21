package workspaces
import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
)
func CastPhoneNumberAccountCreationFromCli (c *cli.Context) *PhoneNumberAccountCreationDto {
	template := &PhoneNumberAccountCreationDto{}
      if c.IsSet("phone-number") {
        value := c.String("phone-number")
        template.PhoneNumber = &value
      }
	return template
}
var PhoneNumberAccountCreationDtoCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "phone-number",
      Required: false,
      Usage:    "phoneNumber",
    },
}
type PhoneNumberAccountCreationDto struct {
    PhoneNumber   *string `json:"phoneNumber" yaml:"phoneNumber"       `
    // Datenano also has a text representation
}
func (x* PhoneNumberAccountCreationDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
func (x* PhoneNumberAccountCreationDto) JsonPrint()  {
    fmt.Println(x.Json())
}
