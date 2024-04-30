package workspaces
import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/urfave/cli"
)
func CastAssignRoleFromCli (c *cli.Context) *AssignRoleDto {
	template := &AssignRoleDto{}
      if c.IsSet("role-id") {
        value := c.String("role-id")
        template.RoleId = &value
      }
      if c.IsSet("user-id") {
        value := c.String("user-id")
        template.UserId = &value
      }
      if c.IsSet("visibility") {
        value := c.String("visibility")
        template.Visibility = &value
      }
	return template
}
var AssignRoleDtoCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "role-id",
      Required: false,
      Usage:    "roleId",
    },
    &cli.StringFlag{
      Name:     "user-id",
      Required: false,
      Usage:    "userId",
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
type AssignRoleDto struct {
    RoleId   *string `json:"roleId" yaml:"roleId"       `
    // Datenano also has a text representation
    UserId   *string `json:"userId" yaml:"userId"       `
    // Datenano also has a text representation
    Visibility   *string `json:"visibility" yaml:"visibility"       `
    // Datenano also has a text representation
    Updated   *int64 `json:"updated" yaml:"updated"       `
    // Datenano also has a text representation
    Created   *int64 `json:"created" yaml:"created"       `
    // Datenano also has a text representation
}
func (x* AssignRoleDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x* AssignRoleDto) JsonPrint()  {
    fmt.Println(x.Json())
}
