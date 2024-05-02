package workspaces
import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/urfave/cli"
)
func CastOtpAuthenticateFromCli (c *cli.Context) *OtpAuthenticateDto {
	template := &OtpAuthenticateDto{}
      if c.IsSet("value") {
        value := c.String("value")
        template.Value = &value
      }
      if c.IsSet("otp") {
        value := c.String("otp")
        template.Otp = &value
      }
      if c.IsSet("type") {
        value := c.String("type")
        template.Type = &value
      }
      if c.IsSet("password") {
        value := c.String("password")
        template.Password = &value
      }
	return template
}
var OtpAuthenticateDtoCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "value",
      Required: true,
      Usage:    "value",
    },
    &cli.StringFlag{
      Name:     "otp",
      Required: false,
      Usage:    "otp",
    },
    &cli.StringFlag{
      Name:     "type",
      Required: true,
      Usage:    "type",
    },
    &cli.StringFlag{
      Name:     "password",
      Required: true,
      Usage:    "password",
    },
}
type OtpAuthenticateDto struct {
    Value   *string `json:"value" yaml:"value"  validate:"required"       `
    // Datenano also has a text representation
    Otp   *string `json:"otp" yaml:"otp"       `
    // Datenano also has a text representation
    Type   *string `json:"type" yaml:"type"  validate:"required"       `
    // Datenano also has a text representation
    Password   *string `json:"password" yaml:"password"  validate:"required"       `
    // Datenano also has a text representation
}
func (x* OtpAuthenticateDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x* OtpAuthenticateDto) JsonPrint()  {
    fmt.Println(x.Json())
}
// This is an experimental way to create new dtos, with exluding the pointers as helper.
func NewOtpAuthenticateDto(
	Value string,
	Otp string,
	Type string,
	Password string,
) OtpAuthenticateDto {
    return OtpAuthenticateDto{
	Value: &Value,
	Otp: &Otp,
	Type: &Type,
	Password: &Password,
    }
}