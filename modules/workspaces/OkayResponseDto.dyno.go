package workspaces
import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
)
func CastOkayResponseFromCli (c *cli.Context) *OkayResponseDto {
	template := &OkayResponseDto{}
	return template
}
var OkayResponseDtoCommonCliFlagsOptional = []cli.Flag{
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
type OkayResponseDto struct {
}
func (x* OkayResponseDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
func (x* OkayResponseDto) JsonPrint()  {
    fmt.Println(x.Json())
}
