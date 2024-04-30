package licenses
import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/urfave/cli"
)
func CastLicenseFromPlanIdFromCli (c *cli.Context) *LicenseFromPlanIdDto {
	template := &LicenseFromPlanIdDto{}
      if c.IsSet("machine-id") {
        value := c.String("machine-id")
        template.MachineId = &value
      }
      if c.IsSet("email") {
        value := c.String("email")
        template.Email = &value
      }
      if c.IsSet("owner") {
        value := c.String("owner")
        template.Owner = &value
      }
	return template
}
var LicenseFromPlanIdDtoCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "machine-id",
      Required: false,
      Usage:    "machineId",
    },
    &cli.StringFlag{
      Name:     "email",
      Required: false,
      Usage:    "email",
    },
    &cli.StringFlag{
      Name:     "owner",
      Required: false,
      Usage:    "owner",
    },
}
type LicenseFromPlanIdDto struct {
    MachineId   *string `json:"machineId" yaml:"machineId"       `
    // Datenano also has a text representation
    Email   *string `json:"email" yaml:"email"       `
    // Datenano also has a text representation
    Owner   *string `json:"owner" yaml:"owner"       `
    // Datenano also has a text representation
}
func (x* LicenseFromPlanIdDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x* LicenseFromPlanIdDto) JsonPrint()  {
    fmt.Println(x.Json())
}
