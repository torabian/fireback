package workspaces
import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/urfave/cli"
)
func CastAcceptInviteFromCli (c *cli.Context) *AcceptInviteDto {
	template := &AcceptInviteDto{}
      if c.IsSet("invite-unique-id") {
        value := c.String("invite-unique-id")
        template.InviteUniqueId = &value
      }
      if c.IsSet("visibility") {
        value := c.String("visibility")
        template.Visibility = &value
      }
	return template
}
var AcceptInviteDtoCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "invite-unique-id",
      Required: false,
      Usage:    "inviteUniqueId",
    },
    &cli.StringFlag{
      Name:     "visibility",
      Required: false,
      Usage:    "visibility",
    },
    &cli.Int64Flag{
      Name:     "updated",
      Required: false,
      Usage:    "updated",
    },
    &cli.Int64Flag{
      Name:     "created",
      Required: false,
      Usage:    "created",
    },
}
type AcceptInviteDto struct {
    InviteUniqueId   *string `json:"inviteUniqueId" yaml:"inviteUniqueId"       `
    // Datenano also has a text representation
    Visibility   *string `json:"visibility" yaml:"visibility"       `
    // Datenano also has a text representation
    Updated   *int64 `json:"updated" yaml:"updated"       `
    // Datenano also has a text representation
    Created   *int64 `json:"created" yaml:"created"       `
    // Datenano also has a text representation
}
func (x* AcceptInviteDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x* AcceptInviteDto) JsonPrint()  {
    fmt.Println(x.Json())
}
