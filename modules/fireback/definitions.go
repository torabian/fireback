package fireback

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type ActivePortType string

type ActiveImportJob struct {
	Filename        string
	TotalLines      int32
	LinesProccessed int32
	SecondsLeft     int32
	TimeLeft        string
	Ref             string
}

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
type ForgetPasswordRequest struct {
	Enabled      bool   `yaml:"enabled,omitempty"`
	url          string `yaml:"url,omitempty"`
	Subject      string `yaml:"subject,omitempty"`
	TemplateFile string `yaml:"templateFile,omitempty"`
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

type AppConfig struct {
	Name        string       `yaml:"name,omitempty"`
	WorkspaceAs string       `yaml:"workspaceAs,omitempty"`
	Token       string       `yaml:"token,omitempty"`
	Worker      WorkerConfig `yaml:"worker,omitempty"`
	CliLanguage string       `yaml:"cliLanguage,omitempty"`
	CliRegion   string       `yaml:"cliRegion,omitempty"`
	SelfHosted  bool         `yaml:"selfHosted,omitempty"`
	License     `yaml:"license,omitempty"`
	Service     `yaml:"service,omitempty"`
	Log         `yaml:"log,omitempty"`
	Database    `yaml:"database,omitempty"`
	Gin         `yaml:"gin,omitempty"`
	Headers     `yaml:"headers,omitempty"`
	Drive       `yaml:"drive,omitempty"`
}

func (x *AppConfig) Yaml() string {
	if x != nil {
		str, _ := yaml.Marshal(x)
		return (string(str))
	}
	return ""
}

func (x *AppConfig) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
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
