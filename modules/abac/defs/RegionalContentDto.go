package abacdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for regionalContentDto
type RegionalContentDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	// The template body sent to the user - supports Go template syntax to insert dynamic values, such as {{.Otp}} for the one-time password.
	Content string `json:"content" validate:"required" yaml:"content"`
	// Region or locale this content applies to, for example any, us, eu, or asia/*. Use any unless you need to target a specific region.
	Region string `json:"region" validate:"required" yaml:"region"`
	// Optional subject line - only used for email-type content.
	Title string `json:"title" yaml:"title"`
	// Language code this content is written in, for example en, fa, or pl. Falls back to English if nothing matches a user's language.
	LanguageId string `json:"languageId" validate:"required" yaml:"languageId"`
	// Which kind of message this content is used for.
	KeyGroup string `json:"keyGroup" validate:"required" yaml:"keyGroup"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// The unique-id of the user which created/owns the record.
	UserId    emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *RegionalContentDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
