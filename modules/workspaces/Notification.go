package workspaces

// General definition of the notification, which could be sent over web, push notification, socket
// and more.
type Notification struct {

	// Name of the notification, or kinda identifier of the notification
	Name string

	// Content of the notification
	Payload []byte

	// The permissions it requires
	Permissions []string
}

type Event struct {

	// Name of the notification, or kinda identifier of the notification
	Name string

	// Content of the notification
	Payload []byte

	// The permissions it requires
	Permissions []string
}
