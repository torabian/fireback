package abac

import "encoding/json"

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
