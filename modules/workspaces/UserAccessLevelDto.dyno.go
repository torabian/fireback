package workspaces

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/urfave/cli"
)

func CastUserAccessLevelFromCli(c *cli.Context) *UserAccessLevelDto {
	template := &UserAccessLevelDto{}
	if c.IsSet("user-role-workspace-permissions") {
		value := c.String("user-role-workspace-permissions")
		template.UserRoleWorkspacePermissionsListId = strings.Split(value, ",")
	}
	if c.IsSet("sql") {
		value := c.String("sql")
		template.SQL = &value
	}
	return template
}

var UserAccessLevelDtoCommonCliFlagsOptional = []cli.Flag{
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
		Name:     "user-role-workspace-permissions",
		Required: false,
		Usage:    "userRoleWorkspacePermissions",
	},
	&cli.StringFlag{
		Name:     "sql",
		Required: false,
		Usage:    "SQL",
	},
}

type UserAccessLevelDto struct {
	Capabilities []string `json:"capabilities" yaml:"capabilities"       `
	// Datenano also has a text representation
	UserRoleWorkspacePermissions []*UserRoleWorkspacePermissionDto `json:"userRoleWorkspacePermissions" yaml:"userRoleWorkspacePermissions"    gorm:"many2many:_userRoleWorkspacePermissions;foreignKey:UniqueId;references:UniqueId"     `
	// Datenano also has a text representation
	UserRoleWorkspacePermissionsListId []string `json:"userRoleWorkspacePermissionsListId" yaml:"userRoleWorkspacePermissionsListId" gorm:"-" sql:"-"`
	Workspaces                         []string `json:"workspaces" yaml:"workspaces"       `
	// Datenano also has a text representation
	SQL *string `json:"SQL" yaml:"SQL"       `
	// Datenano also has a text representation
}

func (x *UserAccessLevelDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x *UserAccessLevelDto) JsonPrint() {
	fmt.Println(x.Json())
}

// This is an experimental way to create new dtos, with exluding the pointers as helper.
func NewUserAccessLevelDto(
	SQL string,
) UserAccessLevelDto {
	return UserAccessLevelDto{
		SQL: &SQL,
	}
}
