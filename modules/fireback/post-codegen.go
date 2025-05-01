package fireback

import "encoding/json"

type PostmanProtocolBehavior struct {
	StrictSSL          bool  `json:"strictSSL,omitempty"`
	FollowRedirects    bool  `json:"followRedirects,omitempty"`
	MaxRedirects       int64 `json:"maxRedirects,omitempty"`
	DisableBodyPruning bool  `json:"disableBodyPruning,omitempty"`
	DisableUrlEncoding bool  `json:"disableUrlEncoding,omitempty"`
}

type PostmanCollection struct {
	Info                    PostmanInfo             `json:"info"`
	Item                    []PostmanItem           `json:"item"`
	Auth                    *PostmanAuth            `json:"auth,omitempty"`
	Variable                []PostmanVariable       `json:"variable,omitempty"`
	Event                   []PostmanEvent          `json:"event,omitempty"`
	ProtocolProfileBehavior PostmanProtocolBehavior `json:"protocolProfileBehavior,omitempty"`
}

func (x *PostmanCollection) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

type PostmanInfo struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Schema      string `json:"schema"`
}

type PostmanItem struct {
	Id                      string                  `json:"id"`
	Name                    string                  `json:"name"`
	Description             string                  `json:"description,omitempty"`
	Request                 PostmanRequest          `json:"request"`
	Response                []PostmanResponse       `json:"response,omitempty"`
	Event                   []PostmanEvent          `json:"event,omitempty"`
	ProtocolProfileBehavior PostmanProtocolBehavior `json:"protocolProfileBehavior,omitempty"`
	Item                    []PostmanItem           `json:"item,omitempty"`
}

type PostmanRequest struct {
	Method      string          `json:"method"`
	Header      []PostmanHeader `json:"header,omitempty"`
	Body        PostmanBody     `json:"body,omitempty"`
	Url         PostmanUrl      `json:"url"`
	Description string          `json:"description,omitempty"`
}

type PostmanResponse struct {
	Id              string          `json:"id"`
	Name            string          `json:"name,omitempty"`
	OriginalRequest PostmanRequest  `json:"originalRequest,omitempty"`
	ResponseTime    int64           `json:"responseTime"`
	Header          []PostmanHeader `json:"header,omitempty"`
	Body            string          `json:"body,omitempty"`
	Status          string          `json:"status,omitempty"`
}

type PostmanHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PostmanBodyOptionRaw struct {
	Language string `json:"language"`
}

type PostmanBodyOption struct {
	Raw PostmanBodyOptionRaw `json:"raw"`
}

type PostmanBody struct {
	Mode string `json:"mode"`
	Raw  string `json:"raw,omitempty"`
	Form []struct {
		Key         string `json:"key"`
		Value       string `json:"value,omitempty"`
		ContentType string `json:"content_type,omitempty"`
	} `json:"formdata,omitempty"`
	Options PostmanBodyOption `json:"options,omitempty"`
}

type PostmanUrl struct {
	Raw       string          `json:"raw"`
	Port      string          `json:"port"`
	Host      []string        `json:"host"`
	Path      []string        `json:"path,omitempty"`
	Query     []PostmanQuery  `json:"query,omitempty"`
	Variables []PostmanUrlVar `json:"variable,omitempty"`
	Protocol  string          `json:"protocol"`
}

type PostmanQuery struct {
	Key         string `json:"key"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
}

type PostmanUrlVar struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

type PostmanAuth struct {
	Type   string              `json:"type"`
	ApiKey []PostmanAuthApiKey `json:"apikey"`
}

type PostmanAuthApiKey struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Key   string `json:"key"`
}

type PostmanVariable struct {
	Type        string `json:"type"`
	Key         string `json:"key"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
}

type PostmanEvent struct {
	Listen string        `json:"listen,omitempty"`
	Script PostmanScript `json:"script"`
}

type PostmanScript struct {
	Exec []string `json:"exec,omitempty"`
}
