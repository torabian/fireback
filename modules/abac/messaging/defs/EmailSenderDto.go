package messagingdefs

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for emailSenderDto
type EmailSenderDto struct {
	UniqueId         emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	FromName         string                 `json:"fromName" validate:"required" yaml:"fromName"`
	FromEmailAddress string                 `json:"fromEmailAddress" validate:"required" yaml:"fromEmailAddress"`
	ReplyTo          string                 `json:"replyTo" validate:"required" yaml:"replyTo"`
	NickName         string                 `json:"nickName" validate:"required" yaml:"nickName"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// The unique-id of the user which created/owns the record.
	UserId    emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *EmailSenderDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEmailSenderDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "from-name",
			Type: "string",
		},
		{
			Name: prefix + "from-email-address",
			Type: "string",
		},
		{
			Name: prefix + "reply-to",
			Type: "string",
		},
		{
			Name: prefix + "nick-name",
			Type: "string",
		},
		{
			Name:        prefix + "workspace-id",
			Type:        "string?",
			Description: "The unique-id of the workspace which content belongs to.",
		},
		{
			Name:        prefix + "user-id",
			Type:        "string?",
			Description: "The unique-id of the user which created/owns the record.",
		},
		{
			Name: prefix + "created-at",
			Type: "complex",
		},
		{
			Name: prefix + "updated-at",
			Type: "complex",
		},
	}
}
func CastEmailSenderDtoFromCli(c emigo.CliCastable) EmailSenderDto {
	data := EmailSenderDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("from-name") {
		data.FromName = c.String("from-name")
	}
	if c.IsSet("from-email-address") {
		data.FromEmailAddress = c.String("from-email-address")
	}
	if c.IsSet("reply-to") {
		data.ReplyTo = c.String("reply-to")
	}
	if c.IsSet("nick-name") {
		data.NickName = c.String("nick-name")
	}
	if c.IsSet("workspace-id") {
		emigo.ParseNullable(c.String("workspace-id"), &data.WorkspaceId)
	}
	if c.IsSet("user-id") {
		emigo.ParseNullable(c.String("user-id"), &data.UserId)
	}
	if c.IsSet("created-at") {
		if u, ok := any(&data.CreatedAt).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("created-at")))
		}
	}
	if c.IsSet("updated-at") {
		if u, ok := any(&data.UpdatedAt).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("updated-at")))
		}
	}
	return data
}
