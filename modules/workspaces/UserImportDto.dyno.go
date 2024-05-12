package workspaces
import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/urfave/cli"
)
type UserImportPassports struct {
    Value   *string `json:"value" yaml:"value"       `
    // Datenano also has a text representation
    Password   *string `json:"password" yaml:"password"       `
    // Datenano also has a text representation
}
func ( x * UserImportPassports) RootObjectName() string {
	return "UserImportDto"
}
type UserImportAddress struct {
    Street   *string `json:"street" yaml:"street"       `
    // Datenano also has a text representation
    ZipCode   *string `json:"zipCode" yaml:"zipCode"       `
    // Datenano also has a text representation
    City   *string `json:"city" yaml:"city"       `
    // Datenano also has a text representation
    Country   *string `json:"country" yaml:"country"       `
    // Datenano also has a text representation
}
func ( x * UserImportAddress) RootObjectName() string {
	return "UserImportDto"
}
func CastUserImportFromCli (c *cli.Context) *UserImportDto {
	template := &UserImportDto{}
      if c.IsSet("avatar") {
        value := c.String("avatar")
        template.Avatar = &value
      }
      if c.IsSet("person-id") {
        value := c.String("person-id")
        template.PersonId = &value
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
    Usage:    "uniqueId (primary key)",
  },
  &cli.StringFlag{
    Name:     "pid",
    Required: false,
    Usage:    " Parent record id of the same type",
  },
    &cli.StringFlag{
      Name:     "avatar",
      Required: false,
      Usage:    "avatar",
    },
    &cli.StringSliceFlag{
      Name:     "passports",
      Required: false,
      Usage:    "passports",
    },
    &cli.StringFlag{
      Name:     "person-id",
      Required: false,
      Usage:    "person",
    },
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
      Name:     "street",
      Required: false,
      Usage:    "street",
    },
    &cli.StringFlag{
      Name:     "zip-code",
      Required: false,
      Usage:    "zipCode",
    },
    &cli.StringFlag{
      Name:     "city",
      Required: false,
      Usage:    "city",
    },
    &cli.StringFlag{
      Name:     "country",
      Required: false,
      Usage:    "country",
    },
}
type UserImportDto struct {
    Avatar   *string `json:"avatar" yaml:"avatar"       `
    // Datenano also has a text representation
    Passports   []*  UserImportPassports `json:"passports" yaml:"passports"    gorm:"foreignKey:LinkerId;references:UniqueId"     `
    // Datenano also has a text representation
    Person   *  PersonEntity `json:"person" yaml:"person"    gorm:"foreignKey:PersonId;references:UniqueId"     `
    // Datenano also has a text representation
        PersonId *string `json:"personId" yaml:"personId"`
    Address   *  UserImportAddress `json:"address" yaml:"address"    gorm:"foreignKey:LinkerId;references:UniqueId"     `
    // Datenano also has a text representation
}
func (x* UserImportDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x* UserImportDto) JsonPrint()  {
    fmt.Println(x.Json())
}
// This is an experimental way to create new dtos, with exluding the pointers as helper.
func NewUserImportDto(
	Avatar string,
) UserImportDto {
    return UserImportDto{
	Avatar: &Avatar,
    }
}