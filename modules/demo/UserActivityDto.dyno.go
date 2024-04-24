package demo
import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/urfave/cli"
)
type UserActivityActivities struct {
    UniqueId   *string `json:"uniqueId" yaml:"uniqueId"       `
    // Datenano also has a text representation
    Activity   *int64 `json:"activity" yaml:"activity"       `
    // Datenano also has a text representation
}
func ( x * UserActivityActivities) RootObjectName() string {
	return "UserActivityDto"
}
func CastUserActivityFromCli (c *cli.Context) *UserActivityDto {
	template := &UserActivityDto{}
	return template
}
var UserActivityDtoCommonCliFlagsOptional = []cli.Flag{
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
    &cli.StringSliceFlag{
      Name:     "activities",
      Required: false,
      Usage:    "activities",
    },
}
type UserActivityDto struct {
    Activities   []*  UserActivityActivities `json:"activities" yaml:"activities"    gorm:"foreignKey:LinkerId;references:UniqueId"     `
    // Datenano also has a text representation
}
func (x* UserActivityDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x* UserActivityDto) JsonPrint()  {
    fmt.Println(x.Json())
}
