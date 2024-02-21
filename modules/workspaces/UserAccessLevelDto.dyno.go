package workspaces
import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
)
func CastUserAccessLevelFromCli (c *cli.Context) *UserAccessLevelDto {
	template := &UserAccessLevelDto{}
      if c.IsSet("sql") {
        value := c.String("sql")
        template.SQL = &value
      }
	return template
}
var UserAccessLevelDtoCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "sql",
      Required: false,
      Usage:    "SQL",
    },
}
type UserAccessLevelDto struct {
    Capabilities   []string `json:"capabilities" yaml:"capabilities"       `
    // Datenano also has a text representation
    Workspaces   []string `json:"workspaces" yaml:"workspaces"       `
    // Datenano also has a text representation
    SQL   *string `json:"SQL" yaml:"SQL"       `
    // Datenano also has a text representation
}
func (x* UserAccessLevelDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
func (x* UserAccessLevelDto) JsonPrint()  {
    fmt.Println(x.Json())
}
