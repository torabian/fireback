package workspaces

import (
	"time"

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
	Enabled  bool   `yaml:"enabled"`
	Port     string `yaml:"port"`
	GrpcPort string `yaml:"grpcPort"`
	Host     string `yaml:"host"`
}

// @meta(include)
type BackOfficeServer struct {
	Enabled bool   `yaml:"enabled"`
	Port    string `yaml:"port"`
	Host    string `yaml:"host"`
}

// @meta(include)
type Database struct {
	Username string `yaml:"username"`
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Vendor   string `yaml:"vendor"`
	Dsn      string `yaml:"dsn"`
}

// @meta(include)
type MailServer struct {
	Provider      string `yaml:"provider"`
	ApiKey        string `yaml:"apiKey"`
	SenderName    string `yaml:"senderName"`
	SenderAddress string `yaml:"senderAddress"`
}

// @meta(include)
type Development struct {
	ForwardMails     string `yaml:"forwardMails"`
	ForwardMailsName string `yaml:"forwardMailsName"`
}

// @meta(include)
type Gin struct {
	Mode string `yaml:"mode"`
}

// @meta(include)
type LoginSMS struct {
	Enabled      bool   `yaml:"enabled"`
	url          string `yaml:"smsRedirectUrl"`
	Subject      string `yaml:"subject"`
	TemplateFile string `yaml:"templateFile"`
}

// @meta(include)
type PhoneNumberTemplates struct {
	LoginSMS `yaml:"loginSms"`
}

// @meta(include)
type ConfirmMail struct {
	Enabled      bool   `yaml:"enabled"`
	url          string `yaml:"url"`
	Subject      string `yaml:"subject"`
	TemplateFile string `yaml:"templateFile"`
}

// @meta(include)
type ForgetPasswordRequest struct {
	Enabled      bool   `yaml:"enabled"`
	url          string `yaml:"url"`
	Subject      string `yaml:"subject"`
	TemplateFile string `yaml:"templateFile"`
}

// @meta(include)
type AcceptWorkspaceInvitation struct {
	Enabled      bool   `yaml:"enabled"`
	url          string `yaml:"url"`
	Subject      string `yaml:"subject"`
	TemplateFile string `yaml:"templateFile"`
}

// @meta(include)
type MailTemplates struct {
	ConfirmMail               `yaml:"confirmMail"`
	ForgetPasswordRequest     `yaml:"forgetPasswordRequest"`
	AcceptWorkspaceInvitation `yaml:"acceptWorkspaceInvitation"`
}

// @meta(include)
type SmartUI struct {
	Enabled           bool   `yaml:"enabled"`
	RedirectOnSuccess string `yaml:"redirectOnSuccess"`
	HomeUrl           string `yaml:"homeUrl"`
}

// @meta(include)
type Google struct {
	GoogleSigninEnabled  bool   `yaml:"googleSigninEnabled"`
	GoogleSigninClientId string `yaml:"googleSigninClientId"`
	ApiKey               string `yaml:"apiKey"`
}

// @meta(include)
type Authentication struct {
	Google `yaml:"google"`
}

// @meta(include)
type Headers struct {
	AccessControlAllowOrigin  string `yaml:"access-control-allow-origin"`
	AccessControlAllowHeaders string `yaml:"access-control-allow-headers"`
}

// @meta(include)
type Drive struct {
	Storage string `yaml:"storage"`
	Port    string `yaml:"port"`
	Enabled bool   `yaml:"enabled"`
}

type Log struct {
	StdErr string `yaml:"stderr"`
	StdOut string `yaml:"stdout"`
}

type Service struct {
	MacIdentifier     string `yaml:"macIdentifier"`
	WindowsIdentifier string `yaml:"windowsIdentifier"`
	DebianIdentifier  string `yaml:"DebianIdentifier"`
}

type License struct {
	MacIdentifier string `yaml:"macIdentifier"`
}

type Mqtt struct {
	Name            string `yaml:"name"`
	ClientId        string `yaml:"clientId"`
	Host            string `yaml:"host"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	MqttVersion     string `yaml:"mqttVersion"`
	LastWillTopic   string `yaml:"lastWillTopic"`
	LastWillPayload string `yaml:"lastWillPayload"`
	Port            int64  `yaml:"port"`
	KeepAlive       int64  `yaml:"keepAlive"`
	ConnectTimeout  int64  `yaml:"connectTimeout"`
	LastWillQoS     int64  `yaml:"lastWillQos"`
	SSL             bool   `yaml:"ssl"`
	AutoReconnect   bool   `yaml:"autoReconnect"`
	CleanSession    bool   `yaml:"cleanSession"`
	LastWillRetail  bool   `yaml:"lastWillRetain"`
}

// @meta(include)
type AppConfig struct {
	Name                 string `yaml:"name"`
	WorkspaceAs          string `yaml:"workspaceAs"`
	UserAs               string `yaml:"userAs"`
	CliLanguage          string `yaml:"cliLanguage"`
	CliRegion            string `yaml:"cliRegion"`
	SelfHosted           bool   `yaml:"selfHosted"`
	License              `yaml:"license"`
	Service              `yaml:"service"`
	Log                  `yaml:"log"`
	PublicServer         `yaml:"publicServer"`
	BackOfficeServer     `yaml:"backOfficeServer"`
	Mqtt                 Mqtt `yaml:"mqtt"`
	Database             `yaml:"database"`
	MailServer           `yaml:"mailServer"`
	Development          `yaml:"development"`
	Gin                  `yaml:"gin"`
	PhoneNumberTemplates `yaml:"phoneNumberTemplates"`
	MailTemplates        `yaml:"mailTemplates"`
	SmartUI              `yaml:"smartUI"`
	Authentication       `yaml:"authentication"`
	Headers              `yaml:"headers"`
	Drive                `yaml:"drive"`
}

type QueryResultMeta struct {
	TotalItems          int64 `json:"totalItems"`
	TotalAvailableItems int64 `json:"totalAvailableItems"`
}
