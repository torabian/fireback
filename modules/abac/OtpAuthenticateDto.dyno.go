package abac

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

// The base class definition for otpAuthenticateDto
type OtpAuthenticateDto struct {
	Value    string `json:"value" validate:"required" yaml:"value"`
	Otp      string `json:"otp" yaml:"otp"`
	Type     string `json:"type" validate:"required" yaml:"type"`
	Password string `json:"password" validate:"required" yaml:"password"`
}

func (x *OtpAuthenticateDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
