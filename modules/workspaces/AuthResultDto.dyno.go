package workspaces

/*
*	Generated by fireback 1.2.0
*	Written by Ali Torabi.
*	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
 */
import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"strings"
)

func CastAuthResultFromCli(c *cli.Context) *AuthResultDto {
	template := &AuthResultDto{}
	if c.IsSet("workspace-id") {
		template.WorkspaceId = c.String("workspace-id")
	}
	if c.IsSet("user-role-workspace-permissions") {
		value := c.String("user-role-workspace-permissions")
		template.UserRoleWorkspacePermissionsListId = strings.Split(value, ",")
	}
	if c.IsSet("internal-sql") {
		template.InternalSql = c.String("internal-sql")
	}
	if c.IsSet("user-id") {
		template.UserId = NewStringAutoNull(c.String("user-id"))
	}
	if c.IsSet("user-id") {
		template.UserId = NewStringAutoNull(c.String("user-id"))
	}
	if c.IsSet("access-level-id") {
		template.AccessLevelId = NewStringAutoNull(c.String("access-level-id"))
	}
	return template
}

var AuthResultDtoCommonCliFlagsOptional = []cli.Flag{
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
		Name:     "workspace-id",
		Required: false,
		Usage:    `workspaceId (string)`,
	},
	&cli.StringSliceFlag{
		Name:     "user-role-workspace-permissions",
		Required: false,
		Usage:    `userRoleWorkspacePermissions (many2many)`,
	},
	&cli.StringFlag{
		Name:     "internal-sql",
		Required: false,
		Usage:    `internalSql (string)`,
	},
	&cli.StringFlag{
		Name:     "user-id",
		Required: false,
		Usage:    `userId (string?)`,
	},
	&cli.StringFlag{
		Name:     "user-id",
		Required: false,
		Usage:    `user (one)`,
	},
	&cli.StringFlag{
		Name:     "access-level-id",
		Required: false,
		Usage:    `accessLevel (one)`,
	},
}

type AuthResultDto struct {
	WorkspaceId                        string                            `json:"workspaceId" yaml:"workspaceId"        `
	UserRoleWorkspacePermissions       []*UserRoleWorkspacePermissionDto `json:"userRoleWorkspacePermissions" yaml:"userRoleWorkspacePermissions"    gorm:"many2many:_userRoleWorkspacePermissions;foreignKey:UniqueId;references:UniqueId"      `
	UserRoleWorkspacePermissionsListId []string                          `json:"userRoleWorkspacePermissionsListId" yaml:"userRoleWorkspacePermissionsListId" gorm:"-" sql:"-"`
	InternalSql                        string                            `json:"internalSql" yaml:"internalSql"        `
	UserId                             String                            `json:"userId" yaml:"userId"        `
	UserHas                            []string                          `json:"userHas" yaml:"userHas"        `
	WorkspaceHas                       []string                          `json:"workspaceHas" yaml:"workspaceHas"        `
	User                               *UserEntity                       `json:"user" yaml:"user"    gorm:"foreignKey:UserId;references:UniqueId"      `
	AccessLevel                        *UserAccessLevelDto               `json:"accessLevel" yaml:"accessLevel"    gorm:"foreignKey:AccessLevelId;references:UniqueId"      `
	AccessLevelId                      String                            `json:"accessLevelId" yaml:"accessLevelId"`
}
type AuthResultDtoList struct {
	Items []*AuthResultDto
}

func NewAuthResultDtoList(items []*AuthResultDto) *AuthResultDtoList {
	return &AuthResultDtoList{
		Items: items,
	}
}
func (x *AuthResultDtoList) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
func (x *AuthResultDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x *AuthResultDto) JsonPrint() {
	fmt.Println(x.Json())
}

// This is an experimental way to create new dtos, with exluding the pointers as helper.
func NewAuthResultDto(
	WorkspaceId string,
	InternalSql string,
) AuthResultDto {
	return AuthResultDto{
		WorkspaceId: WorkspaceId,
		InternalSql: InternalSql,
	}
}
