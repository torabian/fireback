package abac

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
	"github.com/torabian/fireback/modules/fireback/complexes"
)

// The base class definition for userOptionalDto
type UserOptionalDto struct {
	UniqueId  emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	FirstName emigo.Nullable[string] `json:"firstName" validate:"required" yaml:"firstName"`
	LastName  emigo.Nullable[string] `json:"lastName" validate:"required" yaml:"lastName"`
	Photo     emigo.Nullable[string] `json:"photo" yaml:"photo"`
	Gender    emigo.Nullable[int]    `json:"gender" yaml:"gender"`
	Title     emigo.Nullable[string] `json:"title" yaml:"title"`
	BirthDate complexes.XDate        `json:"birthDate" yaml:"birthDate"`
	Avatar    emigo.Nullable[string] `json:"avatar" yaml:"avatar"`
	// User last connecting ip address
	LastIpAddress emigo.Nullable[string] `json:"lastIpAddress" yaml:"lastIpAddress"`
	// User primary address location. Can be useful for simple projects that a user is associated with a single address.
	PrimaryAddress emigo.Nullable[UserOptionalDtoPrimaryAddress] `json:"primaryAddress" yaml:"primaryAddress"`
	WorkspaceId    emigo.Nullable[string]                        `json:"workspaceId" yaml:"workspaceId"`
	UserId         emigo.Nullable[string]                        `json:"userId" yaml:"userId"`
	CreatedAt      abaccomplexes.PlainTime                       `json:"createdAt" yaml:"createdAt"`
	UpdatedAt      abaccomplexes.PlainTime                       `json:"updatedAt" yaml:"updatedAt"`
}

// The base class definition for primaryAddress
type UserOptionalDtoPrimaryAddress struct {
	// Street address, building number
	AddressLine1 emigo.Nullable[string] `json:"addressLine1" yaml:"addressLine1"`
	// Apartment, suite, floor (optional)
	AddressLine2 emigo.Nullable[string] `json:"addressLine2" yaml:"addressLine2"`
	// City or locality
	City emigo.Nullable[string] `json:"city" yaml:"city"`
	// State, region, or province
	StateOrProvince emigo.Nullable[string] `json:"stateOrProvince" yaml:"stateOrProvince"`
	// ZIP or postal code
	PostalCode emigo.Nullable[string] `json:"postalCode" yaml:"postalCode"`
	// ISO 3166-1 alpha-2 (e.g., "US", "DE")
	CountryCode emigo.Nullable[string] `json:"countryCode" yaml:"countryCode"`
}

func (x *UserOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetUserOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "first-name",
			Type: "string?",
		},
		{
			Name: prefix + "last-name",
			Type: "string?",
		},
		{
			Name: prefix + "photo",
			Type: "string?",
		},
		{
			Name: prefix + "gender",
			Type: "int?",
		},
		{
			Name: prefix + "title",
			Type: "string?",
		},
		{
			Name: prefix + "birth-date",
			Type: "complex",
		},
		{
			Name: prefix + "avatar",
			Type: "string?",
		},
		{
			Name:        prefix + "last-ip-address",
			Type:        "string?",
			Description: "User last connecting ip address",
		},
		{
			Name:        prefix + "primary-address",
			Type:        "object?",
			Description: "User primary address location. Can be useful for simple projects that a user is associated with a single address.",
		},
		{
			Name: prefix + "workspace-id",
			Type: "string?",
		},
		{
			Name: prefix + "user-id",
			Type: "string?",
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
func CastUserOptionalDtoFromCli(c emigo.CliCastable) UserOptionalDto {
	data := UserOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("first-name") {
		emigo.ParseNullable(c.String("first-name"), &data.FirstName)
	}
	if c.IsSet("last-name") {
		emigo.ParseNullable(c.String("last-name"), &data.LastName)
	}
	if c.IsSet("photo") {
		emigo.ParseNullable(c.String("photo"), &data.Photo)
	}
	if c.IsSet("gender") {
		emigo.ParseNullable(c.String("gender"), &data.Gender)
	}
	if c.IsSet("title") {
		emigo.ParseNullable(c.String("title"), &data.Title)
	}
	if c.IsSet("birth-date") {
		if u, ok := any(&data.BirthDate).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("birth-date")))
		}
	}
	if c.IsSet("avatar") {
		emigo.ParseNullable(c.String("avatar"), &data.Avatar)
	}
	if c.IsSet("last-ip-address") {
		emigo.ParseNullable(c.String("last-ip-address"), &data.LastIpAddress)
	}
	if c.IsSet("primary-address") {
		emigo.ParseNullable(c.String("primary-address"), &data.PrimaryAddress)
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
func GetUserOptionalDtoPrimaryAddressCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "address-line1",
			Type:        "string?",
			Description: "Street address, building number",
		},
		{
			Name:        prefix + "address-line2",
			Type:        "string?",
			Description: "Apartment, suite, floor (optional)",
		},
		{
			Name:        prefix + "city",
			Type:        "string?",
			Description: "City or locality",
		},
		{
			Name:        prefix + "state-or-province",
			Type:        "string?",
			Description: "State, region, or province",
		},
		{
			Name:        prefix + "postal-code",
			Type:        "string?",
			Description: "ZIP or postal code",
		},
		{
			Name:        prefix + "country-code",
			Type:        "string?",
			Description: "ISO 3166-1 alpha-2 (e.g., \"US\", \"DE\")",
		},
	}
}
func CastUserOptionalDtoPrimaryAddressFromCli(c emigo.CliCastable) UserOptionalDtoPrimaryAddress {
	data := UserOptionalDtoPrimaryAddress{}
	if c.IsSet("address-line1") {
		emigo.ParseNullable(c.String("address-line1"), &data.AddressLine1)
	}
	if c.IsSet("address-line2") {
		emigo.ParseNullable(c.String("address-line2"), &data.AddressLine2)
	}
	if c.IsSet("city") {
		emigo.ParseNullable(c.String("city"), &data.City)
	}
	if c.IsSet("state-or-province") {
		emigo.ParseNullable(c.String("state-or-province"), &data.StateOrProvince)
	}
	if c.IsSet("postal-code") {
		emigo.ParseNullable(c.String("postal-code"), &data.PostalCode)
	}
	if c.IsSet("country-code") {
		emigo.ParseNullable(c.String("country-code"), &data.CountryCode)
	}
	return data
}
