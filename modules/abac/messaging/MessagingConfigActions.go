package messaging

import (
	"errors"

	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	messagingdefs "github.com/torabian/fireback/modules/abac/messaging/defs"
	"github.com/torabian/fireback/modules/fireback"
	"gorm.io/gorm"
)

func MessagingConfigUpdateAction(c messagingdefs.MessagingConfigUpdateActionRequest) (*messagingdefs.MessagingConfigUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ResolveStrategy: "workspace",
		AllowOnRoot:     true,
	})
	if err != nil {
		return nil, err
	}
	updated, err2 := messagingConfigUpsert(*query, c.Body)
	if err2 != nil {
		return nil, err2
	}
	return &messagingdefs.MessagingConfigUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func MessagingConfigGetAction(c abacdefs.WorkspaceConfigDistinctGetActionRequest) (*messagingdefs.MessagingConfigGetActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ResolveStrategy: "workspace",
		AllowOnRoot:     true,
	}); err != nil {
		return nil, err
	}

	var item messagingdefs.MessagingConfigEntity
	err2 := fireback.GetDbRef().Where(&messagingdefs.MessagingConfigEntity{WorkspaceId: "root"}).First(&item).Error
	if err2 != nil {
		if errors.Is(err2, gorm.ErrRecordNotFound) {
			return &messagingdefs.MessagingConfigGetActionResponse{Payload: fireback.GResponseSingleItem(&abacdefs.WorkspaceConfigEntity{})}, nil
		}
		return nil, fireback.GormErrorToIError(err2)
	}

	return &messagingdefs.MessagingConfigGetActionResponse{Payload: fireback.GResponseSingleItem(&item)}, nil
}

func messagingConfigUpsert(query fireback.QueryDSL, body messagingdefs.MessagingConfigOptionalDto) (*messagingdefs.MessagingConfigEntity, *fireback.IError) {
	dbref := fireback.GetDbRef()
	var item messagingdefs.MessagingConfigEntity
	if err2 := dbref.
		Where(&messagingdefs.MessagingConfigEntity{WorkspaceId: "root"}).
		FirstOrCreate(&item).Error; err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}

	changes := messagingConfigChangesFromOptionalDto(body)
	if len(changes) > 0 {
		if err2 := dbref.Model(&item).Updates(changes).Error; err2 != nil {
			return nil, fireback.GormErrorToIError(err2)
		}
	}

	var updated messagingdefs.MessagingConfigEntity
	if err2 := dbref.Where("id = ?", item.Id).First(&updated).Error; err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &updated, nil
}

func messagingConfigChangesFromOptionalDto(body messagingdefs.MessagingConfigOptionalDto) map[string]interface{} {
	changes := map[string]interface{}{}

	if body.GeneralEmailProviderId.IsSet() {
		changes["GeneralEmailProviderId"] = body.GeneralEmailProviderId
	}
	if body.GeneralGsmProviderId.IsSet() {
		changes["GeneralGsmProviderId"] = body.GeneralGsmProviderId
	}
	if body.InviteToWorkspaceContentId.IsSet() {
		changes["InviteToWorkspaceContentId"] = body.InviteToWorkspaceContentId
	}
	if body.EmailOtpContentId.IsSet() {
		changes["EmailOtpContentId"] = body.EmailOtpContentId
	}
	if body.SmsOtpContentId.IsSet() {
		changes["SmsOtpContentId"] = body.SmsOtpContentId
	}
	return changes
}
