package abac

import "encoding/json"

// The base class definition for emailAccountSigninDto
type EmailAccountSigninDto struct {
	Email    string `json:"email" validate:"required" yaml:"email"`
	Password string `json:"password" validate:"required" yaml:"password"`
}

func (x *EmailAccountSigninDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
