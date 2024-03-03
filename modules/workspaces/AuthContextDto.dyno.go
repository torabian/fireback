package workspaces
import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/urfave/cli"
)
func CastAuthContextFromCli (c *cli.Context) *AuthContextDto {
	template := &AuthContextDto{}
      if c.IsSet("workspace-id") {
        value := c.String("workspace-id")
        template.WorkspaceId = &value
      }
      if c.IsSet("token") {
        value := c.String("token")
        template.Token = &value
      }
	return template
}
var AuthContextDtoCommonCliFlagsOptional = []cli.Flag{
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
    &cli.BoolFlag{
      Name:     "skip-workspace-id",
      Required: false,
      Usage:    "skipWorkspaceId",
    },
    &cli.StringFlag{
      Name:     "workspace-id",
      Required: false,
      Usage:    "workspaceId",
    },
    &cli.StringFlag{
      Name:     "token",
      Required: false,
      Usage:    "token",
    },
}
type AuthContextDto struct {
    SkipWorkspaceId   *bool `json:"skipWorkspaceId" yaml:"skipWorkspaceId"       `
    // Datenano also has a text representation
    WorkspaceId   *string `json:"workspaceId" yaml:"workspaceId"       `
    // Datenano also has a text representation
    Token   *string `json:"token" yaml:"token"       `
    // Datenano also has a text representation
    Capabilities   []string `json:"capabilities" yaml:"capabilities"       `
    // Datenano also has a text representation
}
func (x* AuthContextDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x* AuthContextDto) JsonPrint()  {
    fmt.Println(x.Json())
}
