package abac

import "encoding/json"

// The base class definition for exchangeKeyInformationDto
type ExchangeKeyInformationDto struct {
	Key        string `json:"key" yaml:"key"`
	Visibility string `json:"visibility" yaml:"visibility"`
}

func (x *ExchangeKeyInformationDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
