package abac

import (
	"time"

	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
)

// MyNotificationsAction is the self-service counterpart to NotificationBrowseAction
// (root-only, workspace-scoped) - same relationship UserPassportsAction has to the
// generic PassportBrowseAction. ResolveStrategyUser resolves the caller by their own
// session, not by workspace membership/capability, so any authenticated user can call
// this regardless of which workspace they're currently in - a notification is
// addressed to a user, not scoped to a workspace, the same way UserPassportsAction
// already treats "my own passports".
func MyNotificationsAction(c abacdefs.MyNotificationsActionRequest) (*abacdefs.MyNotificationsActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ResolveStrategy: fireback.ResolveStrategyUser,
	})
	if err != nil {
		return nil, err
	}
	q := *query

	items := []abacdefs.NotificationEntity{}
	err2 := fireback.GetRef(q).
		Where(abacdefs.NotificationEntity{UserId: q.UserId}).
		Order("created_at desc").
		Limit(100).
		Find(&items).Error
	if err2 != nil {
		return nil, fireback.CastToIError(err2)
	}

	result := []abacdefs.MyNotificationsActionRes{}
	for _, item := range items {
		result = append(result, abacdefs.MyNotificationsActionRes{
			UniqueId:  item.UniqueId,
			Title:     item.Title,
			Body:      item.Body,
			IsRead:    item.IsRead,
			CreatedAt: item.CreatedAt.Format(time.RFC3339),
		})
	}

	return &abacdefs.MyNotificationsActionResponse{
		Payload: fireback.GResponseQuery(result, nil, &q),
	}, nil
}

// MarkNotificationReadAction flips isRead on one of the caller's own notifications.
// Looked up by (uniqueId, userId) together - not uniqueId alone, then a separate
// ownership check - so a uniqueId belonging to someone else 404s exactly like one
// that doesn't exist at all, rather than leaking whether it exists via a different
// error (e.g. a 403). Same ResourceNotFound treatment RoleGetAction's own
// visibility rule uses for "exists, but not for you to see".
func MarkNotificationReadAction(c abacdefs.MarkNotificationReadActionRequest) (*abacdefs.MarkNotificationReadActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ResolveStrategy: fireback.ResolveStrategyUser,
	})
	if err != nil {
		return nil, err
	}
	q := *query

	req := c.Body
	if verr := fireback.CommonStructValidatorPointer(&req, false); verr != nil {
		return nil, verr
	}

	var entity abacdefs.NotificationEntity
	findErr := fireback.GetRef(q).
		Where(abacdefs.NotificationEntity{UniqueId: req.UniqueId, UserId: q.UserId}).
		First(&entity).Error
	if findErr != nil {
		return nil, &fireback.IError{
			Message:  fireback.FirebackMessages.ResourceNotFound,
			HttpCode: 404,
		}
	}

	if !entity.IsRead {
		if updErr := fireback.GetRef(q).Model(&entity).Update("is_read", true).Error; updErr != nil {
			return nil, fireback.CastToIError(updErr)
		}
	}

	return &abacdefs.MarkNotificationReadActionResponse{
		Payload: fireback.GResponseSingleItem(struct{}{}),
	}, nil
}
