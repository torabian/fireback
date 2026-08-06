package abac

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for emailSenderOptionalDto
type EmailSenderOptionalDto struct {
	UniqueId         emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	FromName         emigo.Nullable[string] `json:"fromName" yaml:"fromName"`
	FromEmailAddress emigo.Nullable[string] `json:"fromEmailAddress" yaml:"fromEmailAddress"`
	ReplyTo          emigo.Nullable[string] `json:"replyTo" yaml:"replyTo"`
	NickName         emigo.Nullable[string] `json:"nickName" yaml:"nickName"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// The unique-id of the user which created/owns the record.
	UserId    emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *EmailSenderOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEmailSenderOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "from-name",
			Type: "string?",
		},
		{
			Name: prefix + "from-email-address",
			Type: "string?",
		},
		{
			Name: prefix + "reply-to",
			Type: "string?",
		},
		{
			Name: prefix + "nick-name",
			Type: "string?",
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
func CastEmailSenderOptionalDtoFromCli(c emigo.CliCastable) EmailSenderOptionalDto {
	data := EmailSenderOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("from-name") {
		emigo.ParseNullable(c.String("from-name"), &data.FromName)
	}
	if c.IsSet("from-email-address") {
		emigo.ParseNullable(c.String("from-email-address"), &data.FromEmailAddress)
	}
	if c.IsSet("reply-to") {
		emigo.ParseNullable(c.String("reply-to"), &data.ReplyTo)
	}
	if c.IsSet("nick-name") {
		emigo.ParseNullable(c.String("nick-name"), &data.NickName)
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
