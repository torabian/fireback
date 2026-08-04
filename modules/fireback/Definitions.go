package fireback

import (
	"encoding/json"
)

type Database struct {
	Username string `yaml:"username,omitempty"`
	Port     string `yaml:"port,omitempty"`
	Host     string `yaml:"host,omitempty"`
	Password string `yaml:"password,omitempty"`
	Database string `yaml:"database,omitempty"`
	Vendor   string `yaml:"vendor,omitempty"`
	Dsn      string `yaml:"dsn,omitempty"`
}

type Service struct {
	MacIdentifier     string `yaml:"macIdentifier,omitempty"`
	WindowsIdentifier string `yaml:"windowsIdentifier,omitempty"`
	DebianIdentifier  string `yaml:"DebianIdentifier,omitempty"`
}

type QueryResultMeta struct {
	TotalItems          int64   `json:"totalItems" yaml:"totalItems"`
	TotalAvailableItems int64   `json:"totalAvailableItems" yaml:"totalAvailableItems"`
	Cursor              *string `json:"cursor" yaml:"cursor"`
}

func (x *QueryResultMeta) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
