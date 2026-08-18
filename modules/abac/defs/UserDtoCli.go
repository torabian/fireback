//go:build !wasm

package abacdefs

import (
	"encoding"
	"github.com/torabian/emi/emigo"
)

func GetUserDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "first-name",
			Type: "string",
		},
		{
			Name: prefix + "last-name",
			Type: "string",
		},
		{
			Name: prefix + "photo",
			Type: "string",
		},
		{
			Name: prefix + "gender",
			Type: "int?",
		},
		{
			Name: prefix + "title",
			Type: "string",
		},
		{
			Name: prefix + "birth-date",
			Type: "complex",
		},
		{
			Name: prefix + "avatar",
			Type: "string",
		},
		{
			Name:        prefix + "last-ip-address",
			Type:        "string",
			Description: "User last connecting ip address",
		},
		{
			Name:        prefix + "primary-address",
			Type:        "object?",
			Description: "User primary address location. Can be useful for simple projects that a user is associated with a single address.",
		},
		{
			Name:        prefix + "phone-number",
			Type:        "string?",
			Description: "Contact phone number for this user (separate from any passport used to sign in).",
		},
		{
			Name:        prefix + "job-title",
			Type:        "string?",
			Description: "The user's job title/role, e.g. \"Support Engineer\".",
		},
		{
			Name:        prefix + "company",
			Type:        "string?",
			Description: "The company or organization the user is associated with.",
		},
		{
			Name:        prefix + "bio",
			Type:        "string?",
			Description: "Free-form biography/notes about the user.",
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
func CastUserDtoFromCli(c emigo.CliCastable) UserDto {
	data := UserDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("first-name") {
		data.FirstName = c.String("first-name")
	}
	if c.IsSet("last-name") {
		data.LastName = c.String("last-name")
	}
	if c.IsSet("photo") {
		data.Photo = c.String("photo")
	}
	if c.IsSet("gender") {
		emigo.ParseNullable(c.String("gender"), &data.Gender)
	}
	if c.IsSet("title") {
		data.Title = c.String("title")
	}
	if c.IsSet("birth-date") {
		if u, ok := any(&data.BirthDate).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("birth-date")))
		}
	}
	if c.IsSet("avatar") {
		data.Avatar = c.String("avatar")
	}
	if c.IsSet("last-ip-address") {
		data.LastIpAddress = c.String("last-ip-address")
	}
	if c.IsSet("primary-address") {
		emigo.ParseNullable(c.String("primary-address"), &data.PrimaryAddress)
	}
	if c.IsSet("phone-number") {
		emigo.ParseNullable(c.String("phone-number"), &data.PhoneNumber)
	}
	if c.IsSet("job-title") {
		emigo.ParseNullable(c.String("job-title"), &data.JobTitle)
	}
	if c.IsSet("company") {
		emigo.ParseNullable(c.String("company"), &data.Company)
	}
	if c.IsSet("bio") {
		emigo.ParseNullable(c.String("bio"), &data.Bio)
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
func GetUserDtoPrimaryAddressCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "address-line1",
			Type:        "string",
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
func CastUserDtoPrimaryAddressFromCli(c emigo.CliCastable) UserDtoPrimaryAddress {
	data := UserDtoPrimaryAddress{}
	if c.IsSet("address-line1") {
		data.AddressLine1 = c.String("address-line1")
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
