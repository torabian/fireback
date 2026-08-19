package abac

import (
	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

var notificationPerms = NewCrudPermissionSet("root.modules", "notification", "notification")
var PERM_ROOT_NOTIFICATION = notificationPerms.Wildcard
var PERM_ROOT_NOTIFICATION_QUERY = notificationPerms.Query
var PERM_ROOT_NOTIFICATION_CREATE = notificationPerms.Create
var PERM_ROOT_NOTIFICATION_DELETE = notificationPerms.Delete
var ALL_NOTIFICATION_PERMISSIONS = notificationPerms.All

var NotificationActions = NewEntityActionsBundle[abacdefs.NotificationEntity]()

// NotificationBrowseAction/NotificationGetAction are root-only (AllowOnRoot) - see
// SendNotificationAction's own comment for why every row is stamped with the sending
// root workspace's own WorkspaceId ("root"). A recipient-facing "my notifications"
// listing (querying by userId instead of workspace, no AllowOnRoot) would be a
// natural follow-up but is out of scope here: this only covers root sending and
// root auditing what was sent.
func NotificationBrowseAction(c abacdefs.NotificationBrowseActionRequest) (*abacdefs.NotificationBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_NOTIFICATION_QUERY}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := NotificationActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.NotificationBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func NotificationGetAction(c abacdefs.NotificationGetActionRequest) (*abacdefs.NotificationGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_NOTIFICATION_QUERY}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := NotificationActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.NotificationGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func NotificationAwareDeletePreviewAction(c abacdefs.NotificationAwareDeletePreviewActionRequest) (*abacdefs.NotificationAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_NOTIFICATION_DELETE}, AllowOnRoot: true}); err != nil {
		return nil, err
	}
	uniqueIds := abacdefs.NotificationAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := abacdefs.NotificationEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.NotificationAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func NotificationAwareDeleteAction(c abacdefs.NotificationAwareDeleteActionRequest) (*abacdefs.NotificationAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_NOTIFICATION_DELETE}, AllowOnRoot: true}); err != nil {
		return nil, err
	}
	if err2 := abacdefs.NotificationEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.NotificationAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

// SendNotificationAction lets a root administrator send a notification (title + body)
// to one or more existing users at once - one NotificationEntity row is created per
// recipient, all sharing the same title/body and senderId (the root user's own
// uniqueId, from query.UserId). Unlike CreatePassportForUserAction (which acts on
// exactly one target user and 404s if it doesn't exist), a userId here that doesn't
// resolve to an existing user is skipped rather than failing the whole batch - the
// point of a bulk send is that one bad id in a list of a hundred shouldn't block the
// other ninety-nine, and skippedUserIds in the response tells the caller exactly
// which ones didn't land.
//
// Every created row is stamped with the sending root workspace's own WorkspaceId
// ("root" - see query.WorkspaceId, guaranteed by AllowOnRoot below), the same
// workspace-stamping every entity's Create action needs so it stays visible to
// NotificationBrowseAction's own workspace-scoped query afterwards (see
// PreferenceCreateAction for the same fix on a different entity) - not the
// recipient's own workspace, which the recipient may not even share with root.
func SendNotificationAction(c abacdefs.SendNotificationActionRequest) (*abacdefs.SendNotificationActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ActionRequires: []application.PermissionInfo{PERM_ROOT_NOTIFICATION_CREATE},
		AllowOnRoot:    true,
	})
	if err != nil {
		return nil, err
	}

	req := c.Body
	if verr := fireback.CommonStructValidatorPointer(&req, false); verr != nil {
		return nil, verr
	}

	var senderId emigo.Nullable[string]
	if query.UserId != "" {
		senderId = emigo.NullableOf(query.UserId)
	}

	sentCount := int64(0)
	skipped := []string{}

	for _, userId := range req.UserIds {
		if _, userErr := UserActions.GetOne(fireback.QueryDSL{UniqueId: userId}); userErr != nil {
			skipped = append(skipped, userId)
			continue
		}

		entity := &abacdefs.NotificationEntity{
			UniqueId:    fireback.UUID(),
			UserId:      userId,
			SenderId:    senderId,
			Title:       req.Title,
			Body:        req.Body,
			WorkspaceId: emigo.NullableOf(query.WorkspaceId),
		}
		if _, createErr := NotificationActions.Create(entity, *query); createErr != nil {
			skipped = append(skipped, userId)
			continue
		}

		sentCount++
	}

	return &abacdefs.SendNotificationActionResponse{
		Payload: fireback.GResponseSingleItem(abacdefs.SendNotificationActionRes{
			SentCount:      sentCount,
			SkippedUserIds: skipped,
		}),
	}, nil
}
