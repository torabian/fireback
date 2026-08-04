package fireback

import (
	"encoding/json"
	reflect "reflect"
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

// @meta(include)
type Gin struct {
	Mode string `yaml:"mode,omitempty"`
}

// @meta(include)
type Headers struct {
	AccessControlAllowOrigin  string `yaml:"access-control-allow-origin,omitempty"`
	AccessControlAllowHeaders string `yaml:"access-control-allow-headers,omitempty"`
}

// @meta(include)
type Drive struct {
	Storage string `yaml:"storage,omitempty"`
	Port    string `yaml:"port,omitempty"`
	Enabled bool   `yaml:"enabled,omitempty"`
}

type Log struct {
	StdErr string `yaml:"stderr,omitempty"`
	StdOut string `yaml:"stdout,omitempty"`
}

type Service struct {
	MacIdentifier     string `yaml:"macIdentifier,omitempty"`
	WindowsIdentifier string `yaml:"windowsIdentifier,omitempty"`
	DebianIdentifier  string `yaml:"DebianIdentifier,omitempty"`
}

type License struct {
	MacIdentifier string `yaml:"macIdentifier,omitempty"`
}

type WorkerConfig struct {
	Type        string `yaml:"type,omitempty"`
	Address     string `yaml:"address,omitempty"`
	Concurrency int64  `yaml:"concurrency,omitempty"`
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

// Checks if the underlying value is a nil error.
// Using IError sometimes causes the err to not be nil, but underlying value is nil
func IsErr(err error) bool {
	return err != nil && !reflect.ValueOf(err).IsNil()
}
