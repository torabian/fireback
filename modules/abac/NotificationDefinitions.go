package workspaces

import "encoding/json"

type EmailMessageContent struct {
	FromName  string
	FromEmail string
	ToName    string
	ToEmail   string
	Subject   string
	Content   string
}

func (x *EmailMessageContent) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
