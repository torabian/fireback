package abac

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for gsmProviderOptionalDto
type GsmProviderOptionalDto struct {
	UniqueId         emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	ApiKey           emigo.Nullable[string] `json:"apiKey" yaml:"apiKey"`
	MainSenderNumber emigo.Nullable[string] `json:"mainSenderNumber" yaml:"mainSenderNumber"`
	Type             emigo.Nullable[string] `json:"type" yaml:"type"`
	InvokeUrl        emigo.Nullable[string] `json:"invokeUrl" yaml:"invokeUrl"`
	InvokeBody       emigo.Nullable[string] `json:"invokeBody" yaml:"invokeBody"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// The unique-id of the user which created/owns the record.
	UserId    emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *GsmProviderOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetGsmProviderOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "api-key",
			Type: "string?",
		},
		{
			Name: prefix + "main-sender-number",
			Type: "string?",
		},
		{
			Name: prefix + "type",
			Type: "enum?",
		},
		{
			Name: prefix + "invoke-url",
			Type: "string?",
		},
		{
			Name: prefix + "invoke-body",
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
func CastGsmProviderOptionalDtoFromCli(c emigo.CliCastable) GsmProviderOptionalDto {
	data := GsmProviderOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("api-key") {
		emigo.ParseNullable(c.String("api-key"), &data.ApiKey)
	}
	if c.IsSet("main-sender-number") {
		emigo.ParseNullable(c.String("main-sender-number"), &data.MainSenderNumber)
	}
	if c.IsSet("type") {
		emigo.ParseNullable(c.String("type"), &data.Type)
	}
	if c.IsSet("invoke-url") {
		emigo.ParseNullable(c.String("invoke-url"), &data.InvokeUrl)
	}
	if c.IsSet("invoke-body") {
		emigo.ParseNullable(c.String("invoke-body"), &data.InvokeBody)
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
