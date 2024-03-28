package licenses
import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/urfave/cli"
)
func CastLicenseFromActivationKeyFromCli (c *cli.Context) *LicenseFromActivationKeyDto {
	template := &LicenseFromActivationKeyDto{}
      if c.IsSet("activation-key-id") {
        value := c.String("activation-key-id")
        template.ActivationKeyId = &value
      }
      if c.IsSet("machine-id") {
        value := c.String("machine-id")
        template.MachineId = &value
      }
	return template
}
var LicenseFromActivationKeyDtoCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "activation-key-id",
      Required: false,
      Usage:    "activationKeyId",
    },
    &cli.StringFlag{
      Name:     "machine-id",
      Required: false,
      Usage:    "machineId",
    },
}
type LicenseFromActivationKeyDto struct {
    ActivationKeyId   *string `json:"activationKeyId" yaml:"activationKeyId"       `
    // Datenano also has a text representation
    MachineId   *string `json:"machineId" yaml:"machineId"       `
    // Datenano also has a text representation
}
func (x* LicenseFromActivationKeyDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x* LicenseFromActivationKeyDto) JsonPrint()  {
    fmt.Println(x.Json())
}
