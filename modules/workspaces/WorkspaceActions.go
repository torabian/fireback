package workspaces

import (
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	queries "github.com/torabian/fireback/modules/workspaces/queries"
	"gorm.io/gorm"
)

func GetUserInvitesAction(query QueryDSL) ([]*WorkspaceInviteEntity, *QueryResultMeta, error) {

	query.WorkspaceId = ""
	query.UserId = ""
	refl := reflect.ValueOf(&WorkspaceInviteEntity{})
	return QueryEntitiesPointer[WorkspaceInviteEntity](query, refl)
}

func GetWorkspaceInvitesAction(query QueryDSL) ([]*WorkspaceInviteEntity, *QueryResultMeta, error) {

	// This entity is an exception, it's uniqueid is the workspace id, so we swap that.
	// query.UniqueId = query.WorkspaceId
	// query.WorkspaceId = ""
	// query.InternalQuery = strings.ReplaceAll(query.InternalQuery, "workspace_id", "unique_id")

	refl := reflect.ValueOf(&WorkspaceInviteEntity{})
	return QueryEntitiesPointer[WorkspaceInviteEntity](query, refl)
}

// Use this one for internal purposes, it queries everything
func GetSystemWorkspacesAction(query QueryDSL) ([]*WorkspaceEntity, *QueryResultMeta, error) {
	refl := reflect.ValueOf(&WorkspaceEntity{})
	return QueryEntitiesPointer[WorkspaceEntity](query, refl)
}

func GetUserWorkspacesAction(query QueryDSL) ([]*WorkspaceEntity, *QueryResultMeta, error) {

	// This entity is an exception, it's uniqueid is the workspace id, so we swap that.
	// query.UniqueId = query.WorkspaceId
	query.InternalQuery = strings.ReplaceAll(query.InternalQuery, "workspace_id", "unique_id")
	query.Query = "linker_id = \"" + query.WorkspaceId + "\""
	query.WorkspaceId = ""

	// fmt.Println("Internal query:", query.InternalQuery)
	refl := reflect.ValueOf(&WorkspaceEntity{})
	return QueryEntitiesPointer[WorkspaceEntity](query, refl)

	// var items []*UserRoleWorkspaceEntity
	// q := GetDbRef().
	// 	Preload("Role").
	// 	Preload("Role.Capabilities").
	// 	Preload("Workspace").
	// 	Preload("User").
	// 	Where("user_id = ?", query.UserId)

	// q.Find(&items)

	// return items, int64(len(items)), nil
}

func GetWorkspaceAction(query QueryDSL) (*WorkspaceEntity, *IError) {

	refl := reflect.ValueOf(&WorkspaceEntity{})
	return GetOneEntity[WorkspaceEntity](query, refl)

}

var ForcedCapabilities []*CapabilityEntity = nil

func UpdateWorkspaceAction(query QueryDSL, dto *WorkspaceEntity) (*WorkspaceEntity, *IError) {

	// query.UniqueId = query.WorkspaceId
	// query.WorkspaceId = ""
	// query.InternalQuery = strings.ReplaceAll(query.InternalQuery, "workspace_id", "unique_id")

	// refl := reflect.ValueOf(&WorkspaceEntity{})
	return UpdateEntity(query, dto)
	// var items []*UserRoleWorkspaceEntity
	// q := GetDbRef().
	// 	Preload("Role").
	// 	Preload("Role.Capabilities").
	// 	Preload("Workspace").
	// 	Preload("User").
	// 	Where("workspace_id = ?", query.UniqueId)

	// q.Find(&items)

	// return &WorkspaceDto{Relations: items}, nil
}

func LimitCapabilitiesByVariant(urw []*WorkspaceRoleEntity) {
	if ForcedCapabilities != nil {
		for _, item := range urw {
			item.Role.Capabilities = ForcedCapabilities
		}
	}
}

func GetUserWorkspaces(q QueryDSL) ([]*WorkspaceEntity, error) {

	var workspaces []*WorkspaceEntity
	if items, _, err := UserWorkspaceActionQuery(q); err != nil {
		return nil, err
	} else {
		for _, item := range items {
			workspaces = append(workspaces, item.Workspace)
		}
	}

	// Needs to be moved to the roles instead
	// LimitCapabilitiesByVariant(workspaces)

	return workspaces, nil
}

func GetRolesInsideWorkspaceById(workspaceId string) ([]*RoleEntity, *IError) {
	var items []*RoleEntity

	GetDbRef().
		Where(RoleEntity{WorkspaceId: &workspaceId}).
		Find(&items)

	return items, nil
}

func GetRoles(q QueryDSL) ([]*RoleEntity, error) {
	roles := []*RoleEntity{}

	if items, _, err := WorkspaceRoleActionQuery(q); err != nil {
		return nil, err
	} else {
		for _, item := range items {
			roles = append(roles, item.Role)
		}
	}

	// Needs to be moved to the roles instead
	// LimitCapabilitiesByVariant(roles)

	return roles, nil
}

// Case 1: User is a ROOT, can access every entity, and every row
// Case 2: User is not a ROOT, but can everyones passports, and modify them
// Case 3: User is a Workspace owner, can do whatever, but only in his workspace
// Case 4: User is inside a workspace, only can do certain actions

func appendAccessLevelToSQL(acl *UserAccessLevelDto) {

	sql := ""

	if !Contains(acl.Workspaces, "*") && len(acl.Workspaces) > 0 && !Contains(acl.Workspaces, ROOT_VAR) {
		sql += "workspace_id in (\"" + strings.Join(acl.Workspaces, "\",\"") + "\") or visibility = \"A\""
	}

	acl.SQL = &sql
}

type UserRoleWorkspacePermission struct {
	WorkspaceId  string `gorm:"workspace_id" json:"workspaceId"`
	UserId       string `gorm:"user_id" json:"userId"`
	RoleId       string `gorm:"role_id" json:"roleId"`
	CapabilityId string `gorm:"capability_id" json:"capabilityId"`
}

func GetUserAccessLevels(query QueryDSL) (*UserAccessLevelDto, *IError) {

	access := &UserAccessLevelDto{
		Workspaces: []string{"system"},
	}

	query.ItemsPerPage = 1000

	items, _, err := UnsafeQuerySqlFromFs[UserRoleWorkspacePermission](
		&queries.QueriesFs, "UserRolePermission", query,
	)

	if err != nil {
		return nil, CastToIError(err)
	}

	for _, item := range items {
		access.Workspaces = append(access.Workspaces, item.WorkspaceId)
		access.Capabilities = append(access.Capabilities, item.CapabilityId)
	}

	appendAccessLevelToSQL(access)

	return access, nil
}

// func AddUserToWorkspace(UserID string, invite *WorkspaceInviteEntity) error {
// user, _ := UserActionGetOne(QueryDSL{UniqueId: UserID})

// var item UserRoleWorkspaceEntity = UserRoleWorkspaceEntity{
// 	User:        user,
// 	RoleId:      invite.RoleId,
// 	WorkspaceId: invite.WorkspaceId,
// 	UniqueId:    UUID(),
// }

// return GetDbRef().Create(&item).Error
// }

func AssignRoleToUserAction(dto *AssignRoleDto, query QueryDSL) (bool, *IError) {
	// func AssignRoleToUserAction(userUniqueId string, workspace WorkspaceEntity, roleUniqueId string) error {

	// user, _ := UserActionGetOne(QueryDSL{UniqueId: *dto.UserId})
	// role, _ := RoleActionGetOne(QueryDSL{
	// 	UniqueId: *dto.RoleId,
	// })

	// var item UserRoleWorkspaceEntity = UserRoleWorkspaceEntity{
	// 	User:        user,
	// 	Role:        role,
	// 	WorkspaceId: &query.WorkspaceId,
	// 	UniqueId:    UUID(),
	// }

	// err := GetDbRef().Create(&item).Error

	// if err != nil {
	// 	return false, GormErrorToIError(err)
	// }

	return true, nil
}

func AssignRoleToUser(userUniqueId string, workspace WorkspaceEntity, roleUniqueId string) error {

	// user, _ := UserActionGetOne(QueryDSL{UniqueId: userUniqueId})

	// // role := GetRoleByUniqueId(roleUniqueId)
	// role, _ := RoleActionGetOne(QueryDSL{
	// 	UniqueId: roleUniqueId,
	// })

	// var item UserRoleWorkspaceEntity = UserRoleWorkspaceEntity{
	// 	User:      user,
	// 	Role:      role,
	// 	Workspace: &workspace,
	// 	UniqueId:  UUID(),
	// }

	// return GetDbRef().Create(&item).Error

	return nil
}

func GetUserPendingInvitations(UserID string) []PendingWorkspaceInviteEntity {
	var items []PendingWorkspaceInviteEntity

	GetDbRef().Raw(`
select 
	fb_workspace_invite_entities.unique_id,
    fb_passport_entities.value,
    fb_passport_entities.type,
    fb_workspace_invite_entities.workspace_id,
    fb_role_entities.name role_name,
    fb_workspace_entities.name workspace_name,
	fb_workspace_invite_entities.cover_letter,
	fb_workspace_invite_entities.role_id

from fb_passport_entities  
  left join fb_workspace_invite_entities  on 
  fb_workspace_invite_entities.email == fb_passport_entities.value
  left join fb_workspace_entities on 
  fb_workspace_entities.unique_id == fb_workspace_invite_entities.workspace_id
  left join fb_role_entities on 
  fb_role_entities.unique_id == fb_workspace_invite_entities.role_id
  where fb_passport_entities.user_id = ? and fb_workspace_invite_entities.unique_id is not null`, UserID).Scan(&items)

	return items
}

func GetInvite(inviteUniqueId string) (*WorkspaceInviteEntity, error) {
	var item WorkspaceInviteEntity
	err := GetDbRef().Where(&WorkspaceInviteEntity{UniqueId: inviteUniqueId}).First(&item).Error
	return &item, err
}

/**
*	This is used when a user already exists in our system, and we want he accepts the invite.
 */
func AcceptInvitationAction(dto *AcceptInviteDto, query QueryDSL) (*OkayResponseDto, *IError) {
	// invite := &WorkspaceInviteEntity{}
	// GetDbRef().Preload("Workspace").Find(&invite, &WorkspaceInviteEntity{
	// 	UniqueId: *dto.InviteUniqueId,
	// })

	// err := AddUserToWorkspace(query.UserId, invite)

	// if err != nil {
	// 	return nil, GormErrorToIError(err)
	// }

	// GetDbRef().Model(&WorkspaceInviteEntity{}).Delete(&WorkspaceInviteEntity{
	// 	UniqueId: *dto.InviteUniqueId,
	// })

	return &OkayResponseDto{}, nil
}

func DeleteUserWorkspace(workspace WorkspaceEntity) bool {
	return false
	// GetDbRef().Delete(&workspace)
	// var items []UserRoleWorkspaceEntity
	// return GetDbRef().Delete(&items).Where("WorkspaceId = ?", workspace.UniqueId).Error == nil
}

func GetWorkspaceByUniqueId(Id string) *WorkspaceEntity {
	workspace := &WorkspaceEntity{}
	GetDbRef().Where(WorkspaceEntity{UniqueId: Id}).First(&workspace)

	return workspace
}

func GetRoleByUniqueId(Id string) *RoleEntity {
	workspace := &RoleEntity{}
	GetDbRef().Where(RoleEntity{UniqueId: Id}).First(&workspace)

	return workspace
}

func (x *WorkspaceEntity) HasValidationErrors(isPatch bool) *IError {
	return CommonStructValidatorPointer(x, isPatch)
}

func RunTransaction[T any](
	entity *T, query QueryDSL,
	fn func(tx *gorm.DB) error,
) (*T, *IError) {

	vf := GetRef(query).Transaction(fn)

	if vf != nil {
		return nil, CastToIError(vf)
	}
	return entity, nil
}

func WorkspaceActionCreate(entity *WorkspaceEntity, query QueryDSL) (*WorkspaceEntity, *IError) {

	// Workspace is a bit different entity.
	// We always set the workspace id of the workspace same as unique id
	if entity.UniqueId == "" {
		entity.UniqueId = UUID()
		query.WorkspaceId = entity.UniqueId
	}

	return WorkspaceActionCreateFn(entity, query)
	// if entity == nil {
	// 	return nil, CreateIErrorString("ENTITY_NEEDED", []string{}, 403)
	// }

	// // Validate the entity first
	// if err := entity.HasValidationErrors(false); err != nil {
	// 	return nil, err
	// }

	// if entity.UniqueId == "" {
	// 	entity.UniqueId = UUID()
	// }

	// // @todo I have some doubts here about the level of workspaces.
	// // If you are creating a workspace, then it's gonna be under current
	// // Workspace? Can you create parallel workspaces

	// // @todo - check if user has access to this workspace even
	// // which he is requesting, or does it even exist

	// if entity.ParentId == nil {
	// 	entity.ParentId = &query.WorkspaceId
	// }

	// return RunTransaction(entity, query, func(tx *gorm.DB) error {
	// 	query.Tx = tx

	// 	user, err3 := UserActionGetOne(query)
	// 	if err3 != nil {
	// 		return err3
	// 	}

	// 	err := tx.Create(&entity).Error
	// 	if err != nil {
	// 		return err
	// 	}

	// 	capabilities := []*CapabilityEntity{
	// 		{UniqueId: "root/*"},
	// 	}

	// 	adminName := "Administrator"
	// 	roleD := &RoleEntity{
	// 		UniqueId:     UUID(),
	// 		Name:         &adminName,
	// 		Capabilities: capabilities,
	// 	}

	// 	role, err5 := RoleActionCreate(roleD, query)
	// 	if err5 != nil {
	// 		return err5
	// 	}

	// 	var linker UserRoleWorkspaceEntity = UserRoleWorkspaceEntity{
	// 		User:      user,
	// 		Role:      role,
	// 		Workspace: entity,
	// 		UniqueId:  UUID(),
	// 	}

	// 	return tx.Create(&linker).Error
	// })

}

func WorkspaceActionCreateChild(entity *WorkspaceEntity, query QueryDSL) (*WorkspaceEntity, *IError) {
	if entity == nil {
		return nil, CreateIErrorString("ENTITY_NEEDED", []string{}, 403)
	}

	// Validate the entity first
	if err := entity.HasValidationErrors(false); err != nil {
		return nil, err
	}

	if entity.UniqueId == "" {
		entity.UniqueId = UUID()
	}

	if entity.ParentId == nil {
		entity.ParentId = &query.WorkspaceId
	}

	return RunTransaction(entity, query, func(tx *gorm.DB) error {
		query.Tx = tx

		err := tx.Create(&entity).Error
		if err != nil {
			return err
		}

		return nil
	})

}

func GetAllWorkspaces(c *gin.Context) []WorkspaceEntity {

	var items []WorkspaceEntity
	GetDbRef().Find(&items)

	return items
}

// func GetUserRoles(user UserEntity) []UserRoleWorkspaceEntity {

// 	var roles []UserRoleWorkspaceEntity
// 	GetDbRef().Preload("Role").Where(UserRoleWorkspaceEntity{UserID: user.UniqueId}).Find(&roles)

// 	return roles
// }

func WorkspaceConfigurationActionGetOne(query QueryDSL) (*WorkspaceConfigEntity, *IError) {
	var item WorkspaceConfigEntity
	q := GetDbRef()
	err := q.Where("workspace_id = ?", query.WorkspaceId).First(&item).Error

	// if err == gorm.ErrRecordNotFound {
	// 	q.Create(&WorkspaceConfigEntity{
	// 		WorkspaceId: &query.WorkspaceId,
	// 	})
	// }
	if err != nil {
		return nil, GormErrorToIError(err)
	}
	return &item, nil
}

func UpdateWorkspaceConfigurationAction(
	query QueryDSL,
	config *WorkspaceConfigEntity,
) (*WorkspaceConfigEntity, *IError) {

	result := GetDbRef().
		Model(&config).
		Where("workspace_id = ?", query.WorkspaceId).
		UpdateColumns(&config)

	if result.RowsAffected == 0 {
		config.WorkspaceId = &query.WorkspaceId
		GetDbRef().Model(&config).Create(config)
	}

	return config, nil
}

func SyncPermissionsInDatabase(x *XWebServer, db *gorm.DB) {

	for _, item := range x.Modules {

		if item.BackupTables != nil && len(item.BackupTables) > 0 {
			for _, table := range item.BackupTables {

				GetDbRef().Model(&BackupTableMetaEntity{}).Create(&BackupTableMetaEntity{
					UniqueId:      table.EntityName,
					TableNameInDb: &table.TableNameInDb,
				})
			}
		}

		// Insert the permissions into the database
		item.PermissionsProvider = append(item.PermissionsProvider, "root/*")

		for _, perm := range item.PermissionsProvider {
			hasChildren := HasChildren(perm, item.PermissionsProvider)
			UpsertPermission(perm, hasChildren, db)
		}

	}

}

// func WorkspaceActionQuery(query QueryDSL) ([]*WorkspaceEntity, *QueryResultMeta, error) {

// 	result, qrm, err := UnsafeQuerySqlFromFs[WorkspaceEntity](
// 		&queries.QueriesFs, "queryWorkspaces", query,
// 	)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	return result, qrm, err
// }

func ClassicSignupAction(dto *ClassicSignupActionReqDto, q QueryDSL) (*UserSessionDto, *IError) {
	if err := ClassicSignupActionReqValidator(dto); err != nil {
		return nil, err
	}

	// if *dto.Type == "phonenumber" {

	// }

	// if *dto.Type == "email" {

	// }

	ClearShot(dto.Value)
	user, role, workspace, passport := GetEmailPassportSignupMechanism(dto)

	return UnsafeGenerateUser(&GenerateUserDto{

		createUser:      true,
		createWorkspace: true,
		createRole:      true,
		createPassport:  true,

		user:      user,
		role:      role,
		workspace: workspace,
		passport:  passport,

		// We want always to be able to login regardless
		restricted: true,
	}, q)
}

func init() {
	CreateWorkspaceActionImp = CreateWorkspaceAction
	CheckClassicPassportActionImp = CheckClassicPassportAction
	ClassicSignupActionImp = ClassicSignupAction
	ClassicSigninActionImp = ClassicSigninAction
	ClassicPassportOtpActionImp = ClassicPassportOtpAction
	GsmSendSmsWithProviderActionImp = GsmSendSmsWithProvider
	GsmSendSmsActionImp = GsmSendSmsAction
	InviteToWorkspaceActionImp = InviteToWorkspaceAction
}

func InviteToWorkspaceAction(req *WorkspaceInviteEntity, q QueryDSL) (*WorkspaceInviteEntity, *IError) {
	if err := WorkspaceInviteValidator(req, false); err != nil {
		return nil, err
	}

	var invite WorkspaceInviteEntity = WorkspaceInviteEntity{
		Value:       req.Value,
		WorkspaceId: &q.WorkspaceId,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		RoleId:      req.RoleId,
		UniqueId:    UUID(),
	}

	if err := GetRef(q).Create(&invite).Error; err != nil {
		return &invite, GormErrorToIError(err)
	}

	// @todo: Detect the type of passport, and

	method, _ := GetTypeByValue(*req.Value)

	if method == PASSPORT_METHOD_EMAIL {
		if err7 := SendInviteEmail(q, &invite); err7 != nil {
			return nil, GormErrorToIError(err7)
		}
	}
	if method == PASSPORT_METHOD_PHONE {
		inviteBody := "You are invite " + *invite.FirstName + " " + *invite.LastName
		if _, err7 := GsmSendSmsAction(&GsmSendSmsActionReqDto{ToNumber: req.Value, Body: &inviteBody}, q); err7 != nil {
			return nil, GormErrorToIError(err7)
		}
	}

	return &invite, nil
}

func GsmSendSmsAction(req *GsmSendSmsActionReqDto, q QueryDSL) (*GsmSendSmsActionResDto, *IError) {

	if err := GsmSendSmsActionReqValidator(req); err != nil {
		return nil, err
	}
	if res, err := GsmSendSMSUsingNotificationConfig(*req.Body, []string{*req.ToNumber}); err != nil {
		return nil, err
	} else {
		return &GsmSendSmsActionResDto{
			QueueId: res.QueueId,
		}, nil
	}
}

func GsmSendSmsWithProvider(req *GsmSendSmsWithProviderActionReqDto, q QueryDSL) (*GsmSendSmsWithProviderActionResDto, *IError) {

	if err := GsmSendSmsWithProviderActionReqValidator(req); err != nil {
		return nil, err
	}

	return GsmSendSMS(*req.GsmProviderId, *req.Body, []string{*req.ToNumber})
}

func GetTypeByValue(value string) (string, *IError) {
	if len(value) > 2 && (value[0:2] == "00" || value[0:1] == "+") {
		return PASSPORT_METHOD_PHONE, nil
	} else if strings.Contains(value, "@") {
		return PASSPORT_METHOD_EMAIL, nil
	}
	return "", CreateIErrorString(WorkspacesMessageCode.PassportNotAvailable, []string{}, 403)
}

func ClassicPassportOtpAction(req *ClassicPassportOtpActionReqDto, q QueryDSL) (
	*ClassicPassportOtpActionResDto, *IError,
) {
	if err := ClassicPassportOtpActionReqValidator(req); err != nil {
		return nil, err
	}

	var secondsToUnblock int64 = 120
	passport, user, err := UnsafeGetUserByPassportValue(*req.Value, q)
	if err != nil {
		return nil, err
	}

	olderEntity := &ForgetPasswordEntity{}
	GetDbRef().Where(&ForgetPasswordEntity{PassportId: &passport.UniqueId}).Find(olderEntity)

	if olderEntity.UniqueId != "" {
		if req.Otp != nil {

			if *req.Otp == *olderEntity.Otp {
				session := &UserSessionDto{}

				if token, err := user.AuthorizeWithToken(q); err != nil {
					return nil, CastToIError(err)
				} else {
					session.Token = &token
				}

				if err != nil {
					return nil, GormErrorToIError(err)
				}

				// Delete the session so user cannot login again
				err2 := GetDbRef().Where(&ForgetPasswordEntity{PassportId: &passport.UniqueId}).Delete(&ForgetPasswordEntity{}).Error

				if err2 != nil {
					return nil, GormErrorToIError(err)
				}

				return &ClassicPassportOtpActionResDto{
					Session: session,
				}, nil
			}
		}

		if time.Now().UnixNano() < olderEntity.BlockedUntil {

			remaining := (olderEntity.BlockedUntil - time.Now().UnixNano()) / 1000000000

			return &ClassicPassportOtpActionResDto{
					ValidUntil:       &olderEntity.ValidUntil,
					BlockedUntil:     &olderEntity.BlockedUntil,
					SecondsToUnblock: &remaining,
				}, CreateIErrorString(
					PassportMessageCode.OTARequestBlockedUntil, []string{}, 403,
				)
		} else {
			GetDbRef().Where(&ForgetPasswordEntity{PassportId: &passport.UniqueId}).Delete(&ForgetPasswordEntity{})
		}
	}

	{

		if passport == nil || user == nil || user.UniqueId == "" {
			return nil, CreateIErrorString(PassportMessageCode.UserDoesNotExist, []string{}, 403)
		}

		uid := UUID()
		otp := GenerateRandomKey(6)
		url := "http://localhost:8888/reset-password?session=" + uid
		item := &ForgetPasswordEntity{
			User:                user,
			Passport:            passport,
			UniqueId:            uid,
			ValidUntil:          time.Now().Add(time.Second * time.Duration(secondsToUnblock)).UnixNano(),
			BlockedUntil:        time.Now().Add(time.Second * time.Duration(secondsToUnblock)).UnixNano(),
			SecondsToUnblock:    &secondsToUnblock,
			Otp:                 &otp,
			RecoveryAbsoluteUrl: &url,
		}

		if err := GetDbRef().Create(item).Error; err != nil {
			return nil, GormErrorToIError(err)
		}

		passportType, _ := GetTypeByValue(*passport.Value)

		if passportType == "phonenumber" {
			if result, err := ResolveRegionalContentTemplate(&RegionalContentRequest{
				LanguageId:       q.Language,
				Region:           "any",
				RegionContentKey: SMS_OTP,
			}, q); err != nil {
				return nil, err
			} else {
				body, err3 := result.CompileContent(map[string]string{"Otp": otp})
				if err3 != nil {
					return nil, CastToIError(err3)
				}

				if _, err2 := GsmSendSMSUsingNotificationConfig(body, []string{*passport.Value}); err2 != nil {
					return nil, GormErrorToIError(err2)
				}
			}
		}

		if passportType == "email" {
			if result, err := ResolveRegionalContentTemplate(&RegionalContentRequest{
				LanguageId:       q.Language,
				Region:           "any",
				RegionContentKey: EMAIL_OTP,
			}, q); err != nil {
				return nil, err
			} else {
				var body = ""
				var title = ""
				if body0, err3 := result.CompileContent(map[string]string{"Otp": otp}); err3 != nil {
					return nil, CastToIError(err3)
				} else {
					body = body0
				}

				if title0, err3 := result.CompileTitle(map[string]string{"Otp": otp}); err3 != nil {
					return nil, CastToIError(err3)
				} else {
					title = title0
				}

				msg := EmailMessageContent{
					Subject:   title,
					Content:   body,
					ToEmail:   *passport.Value,
					FromName:  "Account Center",
					FromEmail: "accountcenter@gmail.com",
					ToName:    user.FullName(),
				}

				if _, err2 := SendEmailUsingNotificationConfig(&msg, GENERAL_SENDER); err2 != nil {
					return nil, GormErrorToIError(err2)
				}
			}
		}

		return &ClassicPassportOtpActionResDto{
			ValidUntil:       &item.ValidUntil,
			BlockedUntil:     &item.BlockedUntil,
			SecondsToUnblock: &secondsToUnblock,
		}, nil
	}
}

func ClearShot(str *string) {
	v := strings.ToLower(strings.TrimSpace(*str))
	*str = v
}

func UnsafeGetUserByPassportValue(value string, q QueryDSL) (*PassportEntity, *UserEntity, *IError) {

	// Check the passport if exists
	var item PassportEntity
	if err := GetRef(q).Model(&PassportEntity{}).Where(&PassportEntity{Value: &value}).First(&item).Error; err != nil || item.Value == nil {

		return nil, nil, CreateIErrorString(WorkspacesMessageCode.PassportNotAvailable, []string{}, 403)
	}

	var user UserEntity
	if err := GetRef(q).Model(&UserEntity{}).Where(&UserEntity{UniqueId: *item.UserId}).First(&user).Error; err != nil {
		return nil, nil, CreateIErrorString(WorkspacesMessageCode.PassportNotAvailable, []string{}, 403)
	}

	return &item, &user, nil
}

func ClassicSigninAction(req *ClassicSigninActionReqDto, q QueryDSL) (*UserSessionDto, *IError) {
	if err := ClassicSigninActionReqValidator(req); err != nil {
		return nil, err
	}

	session := &UserSessionDto{}

	ClearShot(req.Value)

	var password = ""
	if passport, user, err := UnsafeGetUserByPassportValue(*req.Value, q); err != nil {
		return nil, err
	} else {
		session.User = user
		password = *passport.Password
	}

	if !CheckPasswordHash(*req.Password, password) {
		return nil, CreateIErrorString(WorkspacesMessageCode.PassportNotAvailable, []string{}, 403)
	}

	if session.User == nil {
		return nil, CreateIErrorString(WorkspacesMessageCode.PassportUserNotAvailable, []string{}, 403)
	}

	// Authorize the session, put the token
	if token, err := session.User.AuthorizeWithToken(q); err != nil {
		return nil, CastToIError(err)
	} else {
		session.Token = &token
	}

	return session, nil

}

var TRUE = true
var FALSE = false

func CheckClassicPassportAction(req *CheckClassicPassportActionReqDto, q QueryDSL) (*CheckClassicPassportActionResDto, *IError) {
	if err := CheckClassicPassportActionReqValidator(req); err != nil {
		return nil, err
	}

	ClearShot(req.Value)

	var item PassportEntity
	if err := GetRef(q).Model(&PassportEntity{}).Where(&PassportEntity{Value: req.Value}).First(&item).Error; err == nil && item.Value != nil {
		if *item.Value == *req.Value {
			return &CheckClassicPassportActionResDto{
				Exists: &TRUE,
			}, nil
		}
	}

	return &CheckClassicPassportActionResDto{
		Exists: &FALSE,
	}, nil
}

/**
*	Creates a workspace, considering the parent workspace,
*	Who creates it, and might accept even manager and roles in the first
**/
func CreateWorkspaceAction(req *CreateWorkspaceActionReqDto, q QueryDSL) (*WorkspaceEntity, *IError) {

	context := &GenerateUserDto{
		createUser:      false,
		createWorkspace: true,
		workspace: &WorkspaceEntity{
			Name: req.Name,
		},
		user: &UserEntity{
			UniqueId: q.UserId,
			UserId:   &q.UserId,
		},
		restricted: true,
		// createRole: true,
		// role: &RoleEntity{
		// 	Name: "role",
		// },
	}
	session := &UserSessionDto{}
	if err := CreateWorkspaceAndAssignUser(context, q, session); err != nil {
		return nil, err
	} else {
		return context.workspace, nil
	}

}
