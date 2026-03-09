package abac

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

func GetReactiveSearchResultDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string",
		},
		{
			Name: prefix + "phrase",
			Type: "string",
		},
		{
			Name: prefix + "icon",
			Type: "string",
		},
		{
			Name: prefix + "description",
			Type: "string",
		},
		{
			Name: prefix + "group",
			Type: "string",
		},
		{
			Name: prefix + "ui-location",
			Type: "string",
		},
		{
			Name: prefix + "action-fn",
			Type: "string",
		},
	}
}
func CastReactiveSearchResultDtoFromCli(c emigo.CliCastable) ReactiveSearchResultDto {
	data := ReactiveSearchResultDto{}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("phrase") {
		data.Phrase = c.String("phrase")
	}
	if c.IsSet("icon") {
		data.Icon = c.String("icon")
	}
	if c.IsSet("description") {
		data.Description = c.String("description")
	}
	if c.IsSet("group") {
		data.Group = c.String("group")
	}
	if c.IsSet("ui-location") {
		data.UiLocation = c.String("ui-location")
	}
	if c.IsSet("action-fn") {
		data.ActionFn = c.String("action-fn")
	}
	return data
}

// The base class definition for reactiveSearchResultDto
type ReactiveSearchResultDto struct {
	UniqueId    string `json:"uniqueId" yaml:"uniqueId"`
	Phrase      string `json:"phrase" yaml:"phrase"`
	Icon        string `yaml:"icon" json:"icon"`
	Description string `json:"description" yaml:"description"`
	Group       string `json:"group" yaml:"group"`
	UiLocation  string `json:"uiLocation" yaml:"uiLocation"`
	ActionFn    string `json:"actionFn" yaml:"actionFn"`
}

func (x *ReactiveSearchResultDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
