package abac

import "encoding/json"

// The base class definition for testMailDto
type TestMailDto struct {
	SenderId string `json:"senderId" yaml:"senderId"`
	ToName   string `json:"toName" yaml:"toName"`
	ToEmail  string `json:"toEmail" yaml:"toEmail"`
	Subject  string `json:"subject" yaml:"subject"`
	Content  string `json:"content" yaml:"content"`
}

func (x *TestMailDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
