package abac

/*
*	Generated by fireback 1.2.3
*	Written by Ali Torabi.
*	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
 */
import (
	"encoding/json"
	"fmt"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli"
	"strings"
)

func CastExchangeKeyInformationFromCli(c *cli.Context) *ExchangeKeyInformationDto {
	template := &ExchangeKeyInformationDto{}
	fireback.HandleXsrc(c, template)
	if c.IsSet("key") {
		template.Key = c.String("key")
	}
	if c.IsSet("visibility") {
		template.Visibility = c.String("visibility")
	}
	return template
}

var ExchangeKeyInformationDtoCommonCliFlagsOptional = []cli.Flag{
	&cli.StringFlag{
		Name:     "x-src",
		Required: false,
		Usage:    `Import the body of the request from a file (e.g. json/yaml) on the disk`,
	},
	&cli.StringFlag{
		Name:     "wid",
		Required: false,
		Usage:    "Provide workspace id, if you want to change the data workspace",
	},
	&cli.StringFlag{
		Name:     "uid",
		Required: false,
		Usage:    "Unique Id - external unique hash to query entity",
	},
	&cli.StringFlag{
		Name:     "pid",
		Required: false,
		Usage:    " Parent record id of the same type",
	},
	&cli.StringFlag{
		Name:     "key",
		Required: false,
		Usage:    `key (string)`,
	},
	&cli.StringFlag{
		Name:     "visibility",
		Required: false,
		Usage:    `visibility (string)`,
	},
}

type ExchangeKeyInformationDto struct {
	Key        string `json:"key" xml:"key" yaml:"key"        `
	Visibility string `json:"visibility" xml:"visibility" yaml:"visibility"        `
}
type ExchangeKeyInformationDtoList struct {
	Items []*ExchangeKeyInformationDto
}

func NewExchangeKeyInformationDtoList(items []*ExchangeKeyInformationDto) *ExchangeKeyInformationDtoList {
	return &ExchangeKeyInformationDtoList{
		Items: items,
	}
}
func (x *ExchangeKeyInformationDtoList) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
func (x *ExchangeKeyInformationDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x *ExchangeKeyInformationDto) JsonPrint() {
	fmt.Println(x.Json())
	// Somehow to make the import always needed, makes no sense.
	_ = fireback.Body
}

// This is an experimental way to create new dtos, with exluding the pointers as helper.
func NewExchangeKeyInformationDto(
	Key string,
	Visibility string,
) ExchangeKeyInformationDto {
	return ExchangeKeyInformationDto{
		Key:        Key,
		Visibility: Visibility,
	}
}
