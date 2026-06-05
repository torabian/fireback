package abac

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for phoneNumberAccountCreationDto
type PhoneNumberAccountCreationDto struct {
	PhoneNumber string `json:"phoneNumber" yaml:"phoneNumber"`
}

func (x *PhoneNumberAccountCreationDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
