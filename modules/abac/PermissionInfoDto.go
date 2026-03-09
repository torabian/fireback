package abac

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

func GetPermissionInfoDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "name",
			Type: "string",
		},
		{
			Name: prefix + "description",
			Type: "string",
		},
		{
			Name: prefix + "complete-key",
			Type: "string",
		},
	}
}
func CastPermissionInfoDtoFromCli(c emigo.CliCastable) PermissionInfoDto {
	data := PermissionInfoDto{}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("description") {
		data.Description = c.String("description")
	}
	if c.IsSet("complete-key") {
		data.CompleteKey = c.String("complete-key")
	}
	return data
}

// The base class definition for permissionInfoDto
type PermissionInfoDto struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
	CompleteKey string `json:"completeKey" yaml:"completeKey"`
}

func (x *PermissionInfoDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
