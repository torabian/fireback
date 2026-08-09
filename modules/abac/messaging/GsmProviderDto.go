package messaging

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for gsmProviderDto
type GsmProviderDto struct {
	UniqueId         emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	ApiKey           string                 `json:"apiKey" yaml:"apiKey"`
	MainSenderNumber string                 `json:"mainSenderNumber" validate:"required" yaml:"mainSenderNumber"`
	Type             string                 `json:"type" validate:"required" yaml:"type"`
	InvokeUrl        string                 `json:"invokeUrl" yaml:"invokeUrl"`
	InvokeBody       string                 `json:"invokeBody" yaml:"invokeBody"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// The unique-id of the user which created/owns the record.
	UserId    emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *GsmProviderDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetGsmProviderDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "api-key",
			Type: "string",
		},
		{
			Name: prefix + "main-sender-number",
			Type: "string",
		},
		{
			Name: prefix + "type",
			Type: "enum",
		},
		{
			Name: prefix + "invoke-url",
			Type: "string",
		},
		{
			Name: prefix + "invoke-body",
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
func CastGsmProviderDtoFromCli(c emigo.CliCastable) GsmProviderDto {
	data := GsmProviderDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("api-key") {
		data.ApiKey = c.String("api-key")
	}
	if c.IsSet("main-sender-number") {
		data.MainSenderNumber = c.String("main-sender-number")
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	if c.IsSet("invoke-url") {
		data.InvokeUrl = c.String("invoke-url")
	}
	if c.IsSet("invoke-body") {
		data.InvokeBody = c.String("invoke-body")
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
