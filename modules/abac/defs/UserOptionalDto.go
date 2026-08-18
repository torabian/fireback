package abacdefs

import (
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
	// Contact phone number for this user (separate from any passport used to sign in).
	PhoneNumber emigo.Nullable[string] `json:"phoneNumber" yaml:"phoneNumber"`
	// The user's job title/role, e.g. "Support Engineer".
	JobTitle emigo.Nullable[string] `json:"jobTitle" yaml:"jobTitle"`
	// The company or organization the user is associated with.
	Company emigo.Nullable[string] `json:"company" yaml:"company"`
	// Free-form biography/notes about the user.
	Bio       emigo.Nullable[string]  `json:"bio" yaml:"bio"`
	UserId    emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
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
