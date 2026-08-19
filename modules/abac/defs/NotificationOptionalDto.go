package abacdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for notificationOptionalDto
type NotificationOptionalDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	// UniqueId of the user this notification was sent to.
	UserId emigo.Nullable[string] `json:"userId" validate:"required" yaml:"userId"`
	// UniqueId of the (root) user who sent this notification.
	SenderId emigo.Nullable[string] `json:"senderId" yaml:"senderId"`
	// Short notification title.
	Title emigo.Nullable[string] `json:"title" validate:"required" yaml:"title"`
	// Notification message body.
	Body emigo.Nullable[string] `json:"body" validate:"required" yaml:"body"`
	// Whether the recipient has read this notification yet.
	IsRead emigo.Nullable[bool] `json:"isRead" yaml:"isRead"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	CreatedAt   abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt   abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *NotificationOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
