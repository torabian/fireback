package abac

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

func GetAssignRoleDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "role-id",
			Type: "string",
		},
		{
			Name: prefix + "user-id",
			Type: "string",
		},
		{
			Name: prefix + "visibility",
			Type: "string",
		},
		{
			Name: prefix + "updated",
			Type: "int64",
		},
		{
			Name: prefix + "created",
			Type: "int64",
		},
	}
}
func CastAssignRoleDtoFromCli(c emigo.CliCastable) AssignRoleDto {
	data := AssignRoleDto{}
	if c.IsSet("role-id") {
		data.RoleId = c.String("role-id")
	}
	if c.IsSet("user-id") {
		data.UserId = c.String("user-id")
	}
	if c.IsSet("visibility") {
		data.Visibility = c.String("visibility")
	}
	if c.IsSet("updated") {
		data.Updated = int64(c.Int64("updated"))
	}
	if c.IsSet("created") {
		data.Created = int64(c.Int64("created"))
	}
	return data
}

// The base class definition for assignRoleDto
type AssignRoleDto struct {
	RoleId     string `json:"roleId" yaml:"roleId"`
	UserId     string `yaml:"userId" json:"userId"`
	Visibility string `json:"visibility" yaml:"visibility"`
	Updated    int64  `json:"updated" yaml:"updated"`
	Created    int64  `json:"created" yaml:"created"`
}

func (x *AssignRoleDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
