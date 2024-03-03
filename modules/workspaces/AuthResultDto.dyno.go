package workspaces
import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/urfave/cli"
)
func CastAuthResultFromCli (c *cli.Context) *AuthResultDto {
	template := &AuthResultDto{}
      if c.IsSet("workspace-id") {
        value := c.String("workspace-id")
        template.WorkspaceId = &value
      }
      if c.IsSet("internal-sql") {
        value := c.String("internal-sql")
        template.InternalSql = &value
      }
      if c.IsSet("user-id") {
        value := c.String("user-id")
        template.UserId = &value
      }
      if c.IsSet("user-id") {
        value := c.String("user-id")
        template.UserId = &value
      }
      if c.IsSet("access-level-id") {
        value := c.String("access-level-id")
        template.AccessLevelId = &value
      }
	return template
}
var AuthResultDtoCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "workspace-id",
      Required: false,
      Usage:    "workspaceId",
    },
    &cli.StringFlag{
      Name:     "internal-sql",
      Required: false,
      Usage:    "internalSql",
    },
    &cli.StringFlag{
      Name:     "user-id",
      Required: false,
      Usage:    "userId",
    },
    &cli.StringFlag{
      Name:     "user-id",
      Required: false,
      Usage:    "user",
    },
    &cli.StringFlag{
      Name:     "access-level-id",
      Required: false,
      Usage:    "accessLevel",
    },
}
type AuthResultDto struct {
    WorkspaceId   *string `json:"workspaceId" yaml:"workspaceId"       `
    // Datenano also has a text representation
    InternalSql   *string `json:"internalSql" yaml:"internalSql"       `
    // Datenano also has a text representation
    UserId   *string `json:"userId" yaml:"userId"       `
    // Datenano also has a text representation
    UserHas   []string `json:"userHas" yaml:"userHas"       `
    // Datenano also has a text representation
    User   *  UserEntity `json:"user" yaml:"user"    gorm:"foreignKey:UserId;references:UniqueId"     `
    // Datenano also has a text representation
    AccessLevel   *  UserAccessLevelDto `json:"accessLevel" yaml:"accessLevel"    gorm:"foreignKey:AccessLevelId;references:UniqueId"     `
    // Datenano also has a text representation
        AccessLevelId *string `json:"accessLevelId" yaml:"accessLevelId"`
}
func (x* AuthResultDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x* AuthResultDto) JsonPrint()  {
    fmt.Println(x.Json())
}
