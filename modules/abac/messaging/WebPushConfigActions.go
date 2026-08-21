package messaging

import (
	"github.com/torabian/emi/emigo"
	messagingdefs "github.com/torabian/fireback/modules/abac/messaging/defs"
	"github.com/torabian/fireback/modules/fireback"
)

// webPushConfig is strictly self-service: every authenticated user manages only their
// own browser push subscriptions, so unlike the other messaging entities (which are
// root-only, see emailProviderSecurity/gsmProviderSecurity) there is no capability
// permission gate here - just a valid token, with results scoped to the caller via
// ResolveStrategyUser. See modules/fireback/CrudCoreActions.go's QueryEntitiesPointer,
// which filters by `user_id = query.UserId` whenever query.ResolveStrategy is set to
// this.
func webPushConfigSecurity() *fireback.SecurityModel {
	return &fireback.SecurityModel{ResolveStrategy: fireback.ResolveStrategyUser}
}

var WebPushConfigActions = NewEntityActionsBundle[messagingdefs.WebPushConfigEntity]()

func WebPushConfigBrowseAction(c messagingdefs.WebPushConfigBrowseActionRequest) (*messagingdefs.WebPushConfigBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, webPushConfigSecurity())
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := WebPushConfigActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &messagingdefs.WebPushConfigBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

// webPushConfigGetOwn fetches a webPushConfig by uniqueId and verifies it belongs to
// the resolved query.UserId - GetOneEntity itself only filters by uniqueId, so without
// this check any authenticated user could look up (or, via the callers below,
// update/delete) another user's push subscription by guessing its uniqueId.
func webPushConfigGetOwn(query fireback.QueryDSL) (*messagingdefs.WebPushConfigEntity, *fireback.IError) {
	item, err := WebPushConfigActions.GetOne(query)
	if err != nil {
		return nil, err
	}
	if owner, ok := item.UserId.Get(); !ok || owner == nil || *owner != query.UserId {
		return nil, &fireback.IError{
			Message:  fireback.FirebackMessages.ResourceNotFound,
			HttpCode: 404,
		}
	}
	return item, nil
}

func WebPushConfigGetAction(c messagingdefs.WebPushConfigGetActionRequest) (*messagingdefs.WebPushConfigGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, webPushConfigSecurity())
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := webPushConfigGetOwn(*query)
	if err2 != nil {
		return nil, err2
	}
	return &messagingdefs.WebPushConfigGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func WebPushConfigCreateAction(c messagingdefs.WebPushConfigCreateActionRequest) (*messagingdefs.WebPushConfigCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, webPushConfigSecurity())
	if err != nil {
		return nil, err
	}
	if err2 := fireback.CommonStructValidatorPointer(&c.Body, false); err2 != nil {
		return nil, err2
	}
	entity := &messagingdefs.WebPushConfigEntity{
		Subscription: c.Body.Subscription,
		UserId:       emigo.NullableOf(query.UserId),
		WorkspaceId:  emigo.NullableOf(query.WorkspaceId),
	}
	created, err2 := WebPushConfigActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &messagingdefs.WebPushConfigCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func WebPushConfigUpdateAction(c messagingdefs.WebPushConfigUpdateActionRequest) (*messagingdefs.WebPushConfigUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, webPushConfigSecurity())
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	if _, err2 := webPushConfigGetOwn(*query); err2 != nil {
		return nil, err2
	}
	fields := &messagingdefs.WebPushConfigEntity{UniqueId: c.Params.UniqueId, Subscription: c.Body.Subscription}
	updated, err2 := WebPushConfigActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &messagingdefs.WebPushConfigUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

// webPushConfigOwnedUniqueIds narrows a requested uniqueIds list down to the ones the
// caller actually owns, so a bulk delete/preview can never touch another user's
// subscription even if extra ids are passed in.
func webPushConfigOwnedUniqueIds(query fireback.QueryDSL, uniqueIds []string) []string {
	owned := []string{}
	for _, id := range uniqueIds {
		q := query
		q.UniqueId = id
		if _, err := webPushConfigGetOwn(q); err == nil {
			owned = append(owned, id)
		}
	}
	return owned
}

func WebPushConfigAwareDeletePreviewAction(c messagingdefs.WebPushConfigAwareDeletePreviewActionRequest) (*messagingdefs.WebPushConfigAwareDeletePreviewActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, webPushConfigSecurity())
	if err != nil {
		return nil, err
	}
	uniqueIds := messagingdefs.WebPushConfigAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	owned := webPushConfigOwnedUniqueIds(*query, uniqueIds)
	preview, err2 := messagingdefs.WebPushConfigEntityActions.AwareDeletePreview(fireback.GetDbRef(), owned)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &messagingdefs.WebPushConfigAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func WebPushConfigAwareDeleteAction(c messagingdefs.WebPushConfigAwareDeleteActionRequest) (*messagingdefs.WebPushConfigAwareDeleteActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, webPushConfigSecurity())
	if err != nil {
		return nil, err
	}
	owned := webPushConfigOwnedUniqueIds(*query, c.Body.UniqueIds)
	if err2 := messagingdefs.WebPushConfigEntityActions.AwareDelete(fireback.GetDbRef(), owned); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &messagingdefs.WebPushConfigAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}
