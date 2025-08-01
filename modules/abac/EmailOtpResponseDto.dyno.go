package abac

/*
*	Generated by fireback 1.2.3
*	Written by Ali Torabi.
*	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
 */
import (
	"encoding/json"
	"fmt"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli"
	"strings"
)

func CastEmailOtpResponseFromCli(c *cli.Context) *EmailOtpResponseDto {
	template := &EmailOtpResponseDto{}
	fireback.HandleXsrc(c, template)
	if c.IsSet("request-id") {
		template.RequestId = fireback.NewStringAutoNull(c.String("request-id"))
	}
	if c.IsSet("user-session-id") {
		template.UserSessionId = fireback.NewStringAutoNull(c.String("user-session-id"))
	}
	return template
}

var EmailOtpResponseDtoCommonCliFlagsOptional = []cli.Flag{
	&cli.StringFlag{
		Name:     "x-src",
		Required: false,
		Usage:    `Import the body of the request from a file (e.g. json/yaml) on the disk`,
	},
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
		Name:     "request-id",
		Required: false,
		Usage:    `request (one)`,
	},
	&cli.StringFlag{
		Name:     "user-session-id",
		Required: false,
		Usage:    `userSession (one)`,
	},
}

type EmailOtpResponseDto struct {
	Request       *PublicAuthenticationEntity `json:"request" xml:"request" yaml:"request"    gorm:"foreignKey:RequestId;references:UniqueId"      `
	RequestId     fireback.String             `json:"requestId" yaml:"requestId" xml:"requestId"  `
	UserSession   *UserSessionDto             `json:"userSession" xml:"userSession" yaml:"userSession"    gorm:"foreignKey:UserSessionId;references:UniqueId"      `
	UserSessionId fireback.String             `json:"userSessionId" yaml:"userSessionId" xml:"userSessionId"  `
}
type EmailOtpResponseDtoList struct {
	Items []*EmailOtpResponseDto
}

func NewEmailOtpResponseDtoList(items []*EmailOtpResponseDto) *EmailOtpResponseDtoList {
	return &EmailOtpResponseDtoList{
		Items: items,
	}
}
func (x *EmailOtpResponseDtoList) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
func (x *EmailOtpResponseDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x *EmailOtpResponseDto) JsonPrint() {
	fmt.Println(x.Json())
	// Somehow to make the import always needed, makes no sense.
	_ = fireback.Body
}

// This is an experimental way to create new dtos, with exluding the pointers as helper.
func NewEmailOtpResponseDto() EmailOtpResponseDto {
	return EmailOtpResponseDto{}
}
