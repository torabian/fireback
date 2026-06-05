package abac

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for userImportDto
type UserImportDto struct {
	Avatar    string                              `json:"avatar" yaml:"avatar"`
	Passports emigo.Array[UserImportDtoPassports] `json:"passports" yaml:"passports"`
	Address   UserImportDtoAddress                `json:"address" yaml:"address"`
}

// The base class definition for passports
type UserImportDtoPassports struct {
	Value    string `json:"value" yaml:"value"`
	Password string `json:"password" yaml:"password"`
}

// The base class definition for address
type UserImportDtoAddress struct {
	Street  string `json:"street" yaml:"street"`
	ZipCode string `json:"zipCode" yaml:"zipCode"`
	City    string `json:"city" yaml:"city"`
	Country string `json:"country" yaml:"country"`
}

func (x *UserImportDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
