package workspaces
import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
)
func CastExchangeKeyInformationFromCli (c *cli.Context) *ExchangeKeyInformationDto {
	template := &ExchangeKeyInformationDto{}
      if c.IsSet("key") {
        value := c.String("key")
        template.Key = &value
      }
      if c.IsSet("visibility") {
        value := c.String("visibility")
        template.Visibility = &value
      }
	return template
}
var ExchangeKeyInformationDtoCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "key",
      Required: false,
      Usage:    "key",
    },
    &cli.StringFlag{
      Name:     "visibility",
      Required: false,
      Usage:    "visibility",
    },
}
type ExchangeKeyInformationDto struct {
    Key   *string `json:"key" yaml:"key"       `
    // Datenano also has a text representation
    Visibility   *string `json:"visibility" yaml:"visibility"       `
    // Datenano also has a text representation
}
func (x* ExchangeKeyInformationDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
func (x* ExchangeKeyInformationDto) JsonPrint()  {
    fmt.Println(x.Json())
}
