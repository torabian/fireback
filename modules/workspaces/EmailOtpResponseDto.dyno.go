package workspaces
import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
)
func CastEmailOtpResponseFromCli (c *cli.Context) *EmailOtpResponseDto {
	template := &EmailOtpResponseDto{}
      if c.IsSet("request-id") {
        value := c.String("request-id")
        template.RequestId = &value
      }
      if c.IsSet("user-session-id") {
        value := c.String("user-session-id")
        template.UserSessionId = &value
      }
	return template
}
var EmailOtpResponseDtoCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "request-id",
      Required: false,
      Usage:    "request",
    },
    &cli.StringFlag{
      Name:     "user-session-id",
      Required: false,
      Usage:    "userSession",
    },
}
type EmailOtpResponseDto struct {
    Request   *  ForgetPasswordEntity `json:"request" yaml:"request"    gorm:"foreignKey:RequestId;references:UniqueId"     `
    // Datenano also has a text representation
        RequestId *string `json:"requestId" yaml:"requestId"`
    UserSession   *  UserSessionDto `json:"userSession" yaml:"userSession"    gorm:"foreignKey:UserSessionId;references:UniqueId"     `
    // Datenano also has a text representation
        UserSessionId *string `json:"userSessionId" yaml:"userSessionId"`
}
func (x* EmailOtpResponseDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
func (x* EmailOtpResponseDto) JsonPrint()  {
    fmt.Println(x.Json())
}
