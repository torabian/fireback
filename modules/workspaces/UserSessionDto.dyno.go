package workspaces
import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/urfave/cli"
)
func CastUserSessionFromCli (c *cli.Context) *UserSessionDto {
	template := &UserSessionDto{}
      if c.IsSet("passport-id") {
        value := c.String("passport-id")
        template.PassportId = &value
      }
      if c.IsSet("token") {
        value := c.String("token")
        template.Token = &value
      }
      if c.IsSet("exchange-key") {
        value := c.String("exchange-key")
        template.ExchangeKey = &value
      }
      if c.IsSet("user-workspaces") {
        value := c.String("user-workspaces")
        template.UserWorkspacesListId = strings.Split(value, ",")
      }
      if c.IsSet("user-id") {
        value := c.String("user-id")
        template.UserId = &value
      }
      if c.IsSet("user-id") {
        value := c.String("user-id")
        template.UserId = &value
      }
	return template
}
var UserSessionDtoCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "passport-id",
      Required: false,
      Usage:    "passport",
    },
    &cli.StringFlag{
      Name:     "token",
      Required: false,
      Usage:    "token",
    },
    &cli.StringFlag{
      Name:     "exchange-key",
      Required: false,
      Usage:    "exchangeKey",
    },
    &cli.StringSliceFlag{
      Name:     "user-workspaces",
      Required: false,
      Usage:    "userWorkspaces",
    },
    &cli.StringFlag{
      Name:     "user-id",
      Required: false,
      Usage:    "user",
    },
    &cli.StringFlag{
      Name:     "user-id",
      Required: false,
      Usage:    "userId",
    },
}
type UserSessionDto struct {
    Passport   *  PassportEntity `json:"passport" yaml:"passport"    gorm:"foreignKey:PassportId;references:UniqueId"     `
    // Datenano also has a text representation
        PassportId *string `json:"passportId" yaml:"passportId"`
    Token   *string `json:"token" yaml:"token"       `
    // Datenano also has a text representation
    ExchangeKey   *string `json:"exchangeKey" yaml:"exchangeKey"       `
    // Datenano also has a text representation
    UserWorkspaces   []*  UserWorkspaceEntity `json:"userWorkspaces" yaml:"userWorkspaces"    gorm:"many2many:_userWorkspaces;foreignKey:UniqueId;references:UniqueId"     `
    // Datenano also has a text representation
    UserWorkspacesListId []string `json:"userWorkspacesListId" yaml:"userWorkspacesListId" gorm:"-" sql:"-"`
    User   *  UserEntity `json:"user" yaml:"user"    gorm:"foreignKey:UserId;references:UniqueId"     `
    // Datenano also has a text representation
    UserId   *string `json:"userId" yaml:"userId"       `
    // Datenano also has a text representation
}
func (x* UserSessionDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x* UserSessionDto) JsonPrint()  {
    fmt.Println(x.Json())
}
