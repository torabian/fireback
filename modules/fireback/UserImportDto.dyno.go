package fireback

/*
*	Generated by fireback 1.2.1
*	Written by Ali Torabi.
*	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
 */
import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/urfave/cli"
)

type UserImportPassports struct {
	Value    string `json:"value" yaml:"value"        `
	Password string `json:"password" yaml:"password"        `
}

func (x *UserImportPassports) RootObjectName() string {
	return "UserImportDto"
}

type UserImportAddress struct {
	Street  string `json:"street" yaml:"street"        `
	ZipCode string `json:"zipCode" yaml:"zipCode"        `
	City    string `json:"city" yaml:"city"        `
	Country string `json:"country" yaml:"country"        `
}

func (x *UserImportAddress) RootObjectName() string {
	return "UserImportDto"
}
func CastUserImportFromCli(c *cli.Context) *UserImportDto {
	template := &UserImportDto{}
	if c.IsSet("avatar") {
		template.Avatar = c.String("avatar")
	}
	return template
}

var UserImportDtoCommonCliFlagsOptional = []cli.Flag{
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
		Name:     "avatar",
		Required: false,
		Usage:    `avatar (string)`,
	},
	&cli.StringSliceFlag{
		Name:     "passports",
		Required: false,
		Usage:    `passports (array)`,
	},
	&cli.StringFlag{
		Name:     "address-wid",
		Required: false,
		Usage:    "Provide workspace id, if you want to change the data workspace",
	},
	&cli.StringFlag{
		Name:     "address-uid",
		Required: false,
		Usage:    "Unique Id - external unique hash to query entity",
	},
	&cli.StringFlag{
		Name:     "address-pid",
		Required: false,
		Usage:    " Parent record id of the same type",
	},
	&cli.StringFlag{
		Name:     "address-street",
		Required: false,
		Usage:    `street (string)`,
	},
	&cli.StringFlag{
		Name:     "address-zip-code",
		Required: false,
		Usage:    `zipCode (string)`,
	},
	&cli.StringFlag{
		Name:     "address-city",
		Required: false,
		Usage:    `city (string)`,
	},
	&cli.StringFlag{
		Name:     "address-country",
		Required: false,
		Usage:    `country (string)`,
	},
}

type UserImportDto struct {
	Avatar    string                 `json:"avatar" yaml:"avatar"        `
	Passports []*UserImportPassports `json:"passports" yaml:"passports"    gorm:"foreignKey:LinkerId;references:UniqueId;constraint:OnDelete:CASCADE"      `
	Address   *UserImportAddress     `json:"address" yaml:"address"    gorm:"foreignKey:LinkerId;references:UniqueId;constraint:OnDelete:CASCADE"      `
}
type UserImportDtoList struct {
	Items []*UserImportDto
}

func NewUserImportDtoList(items []*UserImportDto) *UserImportDtoList {
	return &UserImportDtoList{
		Items: items,
	}
}
func (x *UserImportDtoList) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
func (x *UserImportDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x *UserImportDto) JsonPrint() {
	fmt.Println(x.Json())
}

// This is an experimental way to create new dtos, with exluding the pointers as helper.
func NewUserImportDto(
	Avatar string,
) UserImportDto {
	return UserImportDto{
		Avatar: Avatar,
	}
}
