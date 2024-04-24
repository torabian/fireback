package demo
import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/urfave/cli"
)
func CastUserActivityFocusFromCli (c *cli.Context) *UserActivityFocusDto {
	template := &UserActivityFocusDto{}
	return template
}
var UserActivityFocusDtoCommonCliFlagsOptional = []cli.Flag{
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
}
type UserActivityFocusDto struct {
    Ids   []string `json:"ids" yaml:"ids"       `
    // Datenano also has a text representation
}
func (x* UserActivityFocusDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x* UserActivityFocusDto) JsonPrint()  {
    fmt.Println(x.Json())
}
