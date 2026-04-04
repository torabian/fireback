package abac

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

func GetUserImportDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "avatar",
			Type: "string",
		},
		{
			Name: prefix + "passports",
			Type: "array",
		},
		{
			Name:     prefix + "address",
			Type:     "object",
			Children: GetUserImportDtoAddressCliFlags("address-"),
		},
	}
}
func CastUserImportDtoFromCli(c emigo.CliCastable) UserImportDto {
	data := UserImportDto{}
	if c.IsSet("avatar") {
		data.Avatar = c.String("avatar")
	}
	if c.IsSet("passports") {
		data.Passports = emigo.CapturePossibleArray(CastUserImportDtoPassportsFromCli, "passports", c)
	}
	if c.IsSet("address") {
		data.Address = CastUserImportDtoAddressFromCli(c)
	}
	return data
}

// The base class definition for userImportDto
type UserImportDto struct {
	Avatar    string                   `json:"avatar" yaml:"avatar"`
	Passports []UserImportDtoPassports `json:"passports" yaml:"passports"`
	Address   UserImportDtoAddress     `json:"address" yaml:"address"`
}

func GetUserImportDtoPassportsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "value",
			Type: "string",
		},
		{
			Name: prefix + "password",
			Type: "string",
		},
	}
}
func CastUserImportDtoPassportsFromCli(c emigo.CliCastable) UserImportDtoPassports {
	data := UserImportDtoPassports{}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	if c.IsSet("password") {
		data.Password = c.String("password")
	}
	return data
}

// The base class definition for passports
type UserImportDtoPassports struct {
	Value    string `json:"value" yaml:"value"`
	Password string `json:"password" yaml:"password"`
}

func GetUserImportDtoAddressCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "street",
			Type: "string",
		},
		{
			Name: prefix + "zip-code",
			Type: "string",
		},
		{
			Name: prefix + "city",
			Type: "string",
		},
		{
			Name: prefix + "country",
			Type: "string",
		},
	}
}
func CastUserImportDtoAddressFromCli(c emigo.CliCastable) UserImportDtoAddress {
	data := UserImportDtoAddress{}
	if c.IsSet("street") {
		data.Street = c.String("street")
	}
	if c.IsSet("zip-code") {
		data.ZipCode = c.String("zip-code")
	}
	if c.IsSet("city") {
		data.City = c.String("city")
	}
	if c.IsSet("country") {
		data.Country = c.String("country")
	}
	return data
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
