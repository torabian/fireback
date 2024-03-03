package workspaces
import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/urfave/cli"
)
func CastImportRequestFromCli (c *cli.Context) *ImportRequestDto {
	template := &ImportRequestDto{}
      if c.IsSet("file") {
        value := c.String("file")
        template.File = &value
      }
	return template
}
var ImportRequestDtoCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "file",
      Required: false,
      Usage:    "file",
    },
}
type ImportRequestDto struct {
    File   *string `json:"file" yaml:"file"       `
    // Datenano also has a text representation
}
func (x* ImportRequestDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x* ImportRequestDto) JsonPrint()  {
    fmt.Println(x.Json())
}
