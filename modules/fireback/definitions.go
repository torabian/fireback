package fireback

import (
	"encoding/json"
	"time"

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

type InformationProviderService struct {
	InternalUsageService         bool
	CpuInformationService        bool
	SmcSensorsInformationService bool
	MacBatteryInformationService bool
}

var currentAvailableServices InformationProviderService = InformationProviderService{
	InternalUsageService:         true,
	CpuInformationService:        true,
	SmcSensorsInformationService: true,
	MacBatteryInformationService: true,
}

type HostMemoryStat struct {
	DeviceID  int
	Alloc     uint64
	HeapInuse uint64
	HeapIdle  uint64
	Frees     uint64
}

type DeviceMessage struct {
	RawMessage   string
	Source       string
	FinalMessage map[string]interface{}
	DeviceID     int
	Timestamp    time.Time
}

// @meta(include)
type PublicServer struct {
	Enabled bool   `yaml:"enabled,omitempty"`
	Port    string `yaml:"port,omitempty"`
	Host    string `yaml:"host,omitempty"`
}

// @meta(include)
type BackOfficeServer struct {
	Enabled bool   `yaml:"enabled,omitempty"`
	Port    string `yaml:"port,omitempty"`
	Host    string `yaml:"host,omitempty"`
}

// @meta(include)
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
type MailServer struct {
	Provider      string `yaml:"provider,omitempty"`
	ApiKey        string `yaml:"apiKey,omitempty"`
	SenderName    string `yaml:"senderName,omitempty"`
	SenderAddress string `yaml:"senderAddress,omitempty"`
}

// @meta(include)
type Development struct {
	ForwardMails     string `yaml:"forwardMails,omitempty"`
	ForwardMailsName string `yaml:"forwardMailsName,omitempty"`
}

// @meta(include)
type Gin struct {
	Mode string `yaml:"mode,omitempty"`
}

// @meta(include)
type LoginSMS struct {
	Enabled      bool   `yaml:"enabled,omitempty"`
	url          string `yaml:"smsRedirectUrl,omitempty"`
	Subject      string `yaml:"subject,omitempty"`
	TemplateFile string `yaml:"templateFile,omitempty"`
}

// @meta(include)
type PhoneNumberTemplates struct {
	LoginSMS `yaml:"loginSms,omitempty"`
}

// @meta(include)
type ConfirmMail struct {
	Enabled      bool   `yaml:"enabled,omitempty"`
	url          string `yaml:"url,omitempty"`
	Subject      string `yaml:"subject,omitempty"`
	TemplateFile string `yaml:"templateFile,omitempty"`
}

// @meta(include)
type ForgetPasswordRequest struct {
	Enabled      bool   `yaml:"enabled,omitempty"`
	url          string `yaml:"url,omitempty"`
	Subject      string `yaml:"subject,omitempty"`
	TemplateFile string `yaml:"templateFile,omitempty"`
}

// @meta(include)
type AcceptWorkspaceInvitation struct {
	Enabled      bool   `yaml:"enabled,omitempty"`
	url          string `yaml:"url,omitempty"`
	Subject      string `yaml:"subject,omitempty"`
	TemplateFile string `yaml:"templateFile,omitempty"`
}

// @meta(include)
type MailTemplates struct {
	ConfirmMail               `yaml:"confirmMail,omitempty"`
	ForgetPasswordRequest     `yaml:"forgetPasswordRequest,omitempty"`
	AcceptWorkspaceInvitation `yaml:"acceptWorkspaceInvitation,omitempty"`
}

// @meta(include)
type SmartUI struct {
	Enabled           bool   `yaml:"enabled,omitempty"`
	RedirectOnSuccess string `yaml:"redirectOnSuccess,omitempty"`
	HomeUrl           string `yaml:"homeUrl,omitempty"`
}

// @meta(include)
type Google struct {
	GoogleSigninEnabled  bool   `yaml:"googleSigninEnabled,omitempty"`
	GoogleSigninClientId string `yaml:"googleSigninClientId,omitempty"`
	ApiKey               string `yaml:"apiKey,omitempty"`
}

// @meta(include)
type Authentication struct {
	Google `yaml:"google,omitempty"`
}

// @meta(include)
type Headers struct {
	AccessControlAllowOrigin  string `yaml:"access-control-allow-origin,omitempty"`
	AccessControlAllowHeaders string `yaml:"access-control-allow-headers,omitempty"`
}

// @meta(include)
type Drive struct {
	Storage         string `yaml:"storage,omitempty"`
	Port            string `yaml:"port,omitempty"`
	Enabled         bool   `yaml:"enabled,omitempty"`
	ImageTypeConfig `yaml:"imageTypeConfig,omitempty"`
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

type Mqtt struct {
	Name            string `yaml:"name,omitempty"`
	ClientId        string `yaml:"clientId,omitempty"`
	Host            string `yaml:"host,omitempty"`
	Username        string `yaml:"username,omitempty"`
	Password        string `yaml:"password,omitempty"`
	MqttVersion     string `yaml:"mqttVersion,omitempty"`
	LastWillTopic   string `yaml:"lastWillTopic,omitempty"`
	LastWillPayload string `yaml:"lastWillPayload,omitempty"`
	Port            int64  `yaml:"port,omitempty"`
	KeepAlive       int64  `yaml:"keepAlive,omitempty"`
	ConnectTimeout  int64  `yaml:"connectTimeout,omitempty"`
	LastWillQoS     int64  `yaml:"lastWillQos,omitempty"`
	SSL             bool   `yaml:"ssl,omitempty"`
	AutoReconnect   bool   `yaml:"autoReconnect,omitempty"`
	CleanSession    bool   `yaml:"cleanSession,omitempty"`
	LastWillRetail  bool   `yaml:"lastWillRetain,omitempty"`
}

type ImageCropSize struct {
	Width   int
	Height  int
	Quality uint
	Name    string
}

type ImageTypeConfig struct {
	Sizes []ImageCropSize `yaml:"sizes,omitempty"`
}

type AppConfig struct {
	Name                 string       `yaml:"name,omitempty"`
	WorkspaceAs          string       `yaml:"workspaceAs,omitempty"`
	Token                string       `yaml:"token,omitempty"`
	Worker               WorkerConfig `yaml:"worker,omitempty"`
	CliLanguage          string       `yaml:"cliLanguage,omitempty"`
	CliRegion            string       `yaml:"cliRegion,omitempty"`
	SelfHosted           bool         `yaml:"selfHosted,omitempty"`
	License              `yaml:"license,omitempty"`
	Service              `yaml:"service,omitempty"`
	Log                  `yaml:"log,omitempty"`
	PublicServer         `yaml:"publicServer,omitempty"`
	BackOfficeServer     `yaml:"backOfficeServer,omitempty"`
	Mqtt                 Mqtt `yaml:"mqtt,omitempty"`
	Database             `yaml:"database,omitempty"`
	MailServer           `yaml:"mailServer,omitempty"`
	Development          `yaml:"development,omitempty"`
	Gin                  `yaml:"gin,omitempty"`
	PhoneNumberTemplates `yaml:"phoneNumberTemplates,omitempty"`
	MailTemplates        `yaml:"mailTemplates,omitempty"`
	SmartUI              `yaml:"smartUI,omitempty"`
	Authentication       `yaml:"authentication,omitempty"`
	Headers              `yaml:"headers,omitempty"`
	Drive                `yaml:"drive,omitempty"`
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
