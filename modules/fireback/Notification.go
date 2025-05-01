package fireback

import (
	"encoding/json"
)

// General definition of the notification, which could be sent over web, push notification, socket
// and more.
type Notification struct {

	// Name of the notification, or kinda identifier of the notification
	Name string `json:"name"`

	// Content of the notification
	Payload interface{} `json:"payload"`

	// A mechanism to understand the event has invalidated an specific cache key
	// so instead of using payload to detect changes, refresh the resource instead
	// it simplifies some apps in case developer doesn't want to clearly react on the events
	CacheKey string `json:"cacheKey"`
}

type Event struct {

	// Name of the notification, or kinda identifier of the notification
	Name string `json:"name"`

	// Content of the notification
	Payload interface{} `json:"payload"`

	// The source of the event if it's a instance
	InstanceSourceId string `json:"instanceSourceId"`

	// Security model of the event
	Security *SecurityModel `json:"security"`

	// Indicates if the event has occured in the internal project.
	// Because your project might be listening to other environments or events as well
	Internal bool `json:"internal"`

	// A mechanism to understand the event has invalidated an specific cache key
	// so instead of using payload to detect changes, refresh the resource instead
	// it simplifies some apps in case developer doesn't want to clearly react on the events
	CacheKey string `json:"cacheKey"`

	// Event Context can be anything, and it would be used by the event router implementation
	// to understand how the event happened or based on what happened.
	// For example in ABAC module, we put userId and workspaceId inside of it, to just show the
	// message for correct users. You can implement it anyway you want, and access it in the event router
	// Obviously you could use Payload for putting the content as well, this is just another extra separated method
	SourceContext interface{} `json:"sourceContext"`
}

func (x *Event) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
func (x *Event) ToNotification() Notification {

	return Notification{
		Name:     x.Name,
		Payload:  x.Payload,
		CacheKey: x.CacheKey,
	}
}
