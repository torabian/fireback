package workspaces
import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/urfave/cli"
)
func CastReactiveSearchResultFromCli (c *cli.Context) *ReactiveSearchResultDto {
	template := &ReactiveSearchResultDto{}
      if c.IsSet("unique-id") {
        value := c.String("unique-id")
        template.UniqueId = &value
      }
      if c.IsSet("phrase") {
        value := c.String("phrase")
        template.Phrase = &value
      }
      if c.IsSet("icon") {
        value := c.String("icon")
        template.Icon = &value
      }
      if c.IsSet("description") {
        value := c.String("description")
        template.Description = &value
      }
      if c.IsSet("group") {
        value := c.String("group")
        template.Group = &value
      }
      if c.IsSet("ui-location") {
        value := c.String("ui-location")
        template.UiLocation = &value
      }
      if c.IsSet("action-fn") {
        value := c.String("action-fn")
        template.ActionFn = &value
      }
	return template
}
var ReactiveSearchResultDtoCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "unique-id",
      Required: false,
      Usage:    "uniqueId",
    },
    &cli.StringFlag{
      Name:     "phrase",
      Required: false,
      Usage:    "phrase",
    },
    &cli.StringFlag{
      Name:     "icon",
      Required: false,
      Usage:    "icon",
    },
    &cli.StringFlag{
      Name:     "description",
      Required: false,
      Usage:    "description",
    },
    &cli.StringFlag{
      Name:     "group",
      Required: false,
      Usage:    "group",
    },
    &cli.StringFlag{
      Name:     "ui-location",
      Required: false,
      Usage:    "uiLocation",
    },
    &cli.StringFlag{
      Name:     "action-fn",
      Required: false,
      Usage:    "actionFn",
    },
}
type ReactiveSearchResultDto struct {
    UniqueId   *string `json:"uniqueId" yaml:"uniqueId"       `
    // Datenano also has a text representation
    Phrase   *string `json:"phrase" yaml:"phrase"       `
    // Datenano also has a text representation
    Icon   *string `json:"icon" yaml:"icon"       `
    // Datenano also has a text representation
    Description   *string `json:"description" yaml:"description"       `
    // Datenano also has a text representation
    Group   *string `json:"group" yaml:"group"       `
    // Datenano also has a text representation
    UiLocation   *string `json:"uiLocation" yaml:"uiLocation"       `
    // Datenano also has a text representation
    ActionFn   *string `json:"actionFn" yaml:"actionFn"       `
    // Datenano also has a text representation
}
func (x* ReactiveSearchResultDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x* ReactiveSearchResultDto) JsonPrint()  {
    fmt.Println(x.Json())
}
