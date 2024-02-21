package workspaces
import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
)
func CastClassicAuthFromCli (c *cli.Context) *ClassicAuthDto {
	template := &ClassicAuthDto{}
      if c.IsSet("value") {
        value := c.String("value")
        template.Value = &value
      }
      if c.IsSet("password") {
        value := c.String("password")
        template.Password = &value
      }
      if c.IsSet("first-name") {
        value := c.String("first-name")
        template.FirstName = &value
      }
      if c.IsSet("last-name") {
        value := c.String("last-name")
        template.LastName = &value
      }
      if c.IsSet("invite-id") {
        value := c.String("invite-id")
        template.InviteId = &value
      }
      if c.IsSet("public-join-key-id") {
        value := c.String("public-join-key-id")
        template.PublicJoinKeyId = &value
      }
      if c.IsSet("workspace-type-id") {
        value := c.String("workspace-type-id")
        template.WorkspaceTypeId = &value
      }
	return template
}
var ClassicAuthDtoCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "password",
      Required: true,
      Usage:    "password",
    },
    &cli.StringFlag{
      Name:     "first-name",
      Required: true,
      Usage:    "firstName",
    },
    &cli.StringFlag{
      Name:     "last-name",
      Required: true,
      Usage:    "lastName",
    },
    &cli.StringFlag{
      Name:     "invite-id",
      Required: false,
      Usage:    "inviteId",
    },
    &cli.StringFlag{
      Name:     "public-join-key-id",
      Required: false,
      Usage:    "publicJoinKeyId",
    },
    &cli.StringFlag{
      Name:     "workspace-type-id",
      Required: false,
      Usage:    "workspaceTypeId",
    },
}
type ClassicAuthDto struct {
    Value   *string `json:"value" yaml:"value"  validate:"required"       `
    // Datenano also has a text representation
    Password   *string `json:"password" yaml:"password"  validate:"required"       `
    // Datenano also has a text representation
    FirstName   *string `json:"firstName" yaml:"firstName"  validate:"required"       `
    // Datenano also has a text representation
    LastName   *string `json:"lastName" yaml:"lastName"  validate:"required"       `
    // Datenano also has a text representation
    InviteId   *string `json:"inviteId" yaml:"inviteId"       `
    // Datenano also has a text representation
    PublicJoinKeyId   *string `json:"publicJoinKeyId" yaml:"publicJoinKeyId"       `
    // Datenano also has a text representation
    WorkspaceTypeId   *string `json:"workspaceTypeId" yaml:"workspaceTypeId"       `
    // Datenano also has a text representation
}
func (x* ClassicAuthDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
func (x* ClassicAuthDto) JsonPrint()  {
    fmt.Println(x.Json())
}
