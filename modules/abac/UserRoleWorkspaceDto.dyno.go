package abac

import (
	"encoding/json"

	emigo "github.com/torabian/emi/emigo"
)

func GetUserRoleWorkspaceDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "role-id",
			Type: "string",
		},
		{
			Name: prefix + "capabilities",
			Type: "slice",
		},
	}
}
func CastUserRoleWorkspaceDtoFromCli(c emigo.CliCastable) UserRoleWorkspaceDto {
	data := UserRoleWorkspaceDto{}
	if c.IsSet("role-id") {
		data.RoleId = c.String("role-id")
	}
	if c.IsSet("capabilities") {
		emigo.InflatePossibleSlice(c.String("capabilities"), &data.Capabilities)
	}
	return data
}

// The base class definition for userRoleWorkspaceDto
type UserRoleWorkspaceDto struct {
	RoleId       string   `json:"roleId" yaml:"roleId"`
	Capabilities []string `json:"capabilities" yaml:"capabilities"`
}

func (x *UserRoleWorkspaceDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
