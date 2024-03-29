package workspaces
import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/urfave/cli"
)
func CastUserRoleWorkspaceFromCli (c *cli.Context) *UserRoleWorkspaceDto {
	template := &UserRoleWorkspaceDto{}
      if c.IsSet("role-id") {
        value := c.String("role-id")
        template.RoleId = &value
      }
	return template
}
var UserRoleWorkspaceDtoCommonCliFlagsOptional = []cli.Flag{
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
}
type UserRoleWorkspaceDto struct {
    RoleId   *string `json:"roleId" yaml:"roleId"       `
    // Datenano also has a text representation
    Capabilities   []string `json:"capabilities" yaml:"capabilities"       `
    // Datenano also has a text representation
}
func (x* UserRoleWorkspaceDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x* UserRoleWorkspaceDto) JsonPrint()  {
    fmt.Println(x.Json())
}
