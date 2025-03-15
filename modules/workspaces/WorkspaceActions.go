package workspaces

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"strings"

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
	if items, _, err := UserWorkspaceActions.Query(q); err != nil {
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
		Where(RoleEntity{WorkspaceId: NewString(workspaceId)}).
		Find(&items)

	return items, nil
}

func GetRoles(q QueryDSL) ([]*RoleEntity, error) {
	roles := []*RoleEntity{}

	if items, _, err := WorkspaceRoleActions.Query(q); err != nil {
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

// type UserRoleWorkspacePermission struct {
// 	WorkspaceId  string `gorm:"workspace_id" json:"workspaceId"`
// 	UserId       string `gorm:"user_id" json:"userId"`
// 	RoleId       string `gorm:"role_id" json:"roleId"`
// 	CapabilityId string `gorm:"capability_id" json:"capabilityId"`
// 	Type         string `gorm:"type" json:"type"`
// }

type UserAccessPerWorkspaceDto map[string]*struct {

	// The access which are available to this workspace, not to the specific user.
	// Even a user has access to many things, these accesses need to reduce those
	WorkspacesAccesses []string

	// The permissions which user has access to
	UserRoles map[string]*struct {
		Name     string
		Accesses []string
	}
}

func (x *UserAccessPerWorkspaceDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

func GetUserAccessLevels(query QueryDSL) (*UserAccessLevelDto, *IError) {

	access := &UserAccessLevelDto{}
	query.ItemsPerPage = 1000

	items, _, err := UnsafeQuerySqlFromFs[UserRoleWorkspacePermissionDto](
		&queries.QueriesFs, "UserRolePermission", query,
	)

	if err != nil {
		return nil, CastToIError(err)
	}

	ws := UserAccessPerWorkspaceDto{}

	for _, item := range items {
		if ws[item.WorkspaceId] == nil {
			ws[item.WorkspaceId] = &struct {
				WorkspacesAccesses []string
				UserRoles          map[string]*struct {
					Name     string
					Accesses []string
				}
			}{}
		}

		if item.Type == "account_restrict" {
			if ws[item.WorkspaceId].UserRoles[item.RoleId] == nil {
				ws[item.WorkspaceId].UserRoles = map[string]*struct {
					Name     string
					Accesses []string
				}{}
				ws[item.WorkspaceId].UserRoles[item.RoleId] = &struct {
					Name     string
					Accesses []string
				}{}
			}
			ws[item.WorkspaceId].UserRoles[item.RoleId].Accesses = append(ws[item.WorkspaceId].UserRoles[item.RoleId].Accesses, item.CapabilityId)
			ws[item.WorkspaceId].UserRoles[item.RoleId].Name = item.RoleName
		}

		if item.Type == "workspace_restrict" {
			ws[item.WorkspaceId].WorkspacesAccesses = append(ws[item.WorkspaceId].WorkspacesAccesses, item.CapabilityId)
		}
	}

	access.UserAccessPerWorkspace = &ws

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

// Seems not used :((((()))))
func GetUserPendingInvitations(UserID string) []PendingWorkspaceInviteEntity {
	var items []PendingWorkspaceInviteEntity

	GetDbRef().Raw(`
select 
	workspace_invite_entities.unique_id,
    passport_entities.value,
    passport_entities.type,
    workspace_invite_entities.workspace_id,
    role_entities.name role_name,
    workspace_entities.name workspace_name,
	workspace_invite_entities.cover_letter,
	workspace_invite_entities.role_id

from passport_entities  
  left join workspace_invite_entities  on 
  workspace_invite_entities.email == passport_entities.value
  left join workspace_entities on 
  workspace_entities.unique_id == workspace_invite_entities.workspace_id
  left join role_entities on 
  role_entities.unique_id == workspace_invite_entities.role_id
  where passport_entities.user_id = ? and workspace_invite_entities.unique_id is not null`, UserID).Scan(&items)

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
	// 		{UniqueId: ROOT_ALL_ACCESS},
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
		return nil, Create401Error(&WorkspacesMessages.BodyIsMissing, []string{})
	}

	// Validate the entity first
	if err := entity.HasValidationErrors(false); err != nil {
		return nil, err
	}

	if entity.UniqueId == "" {
		entity.UniqueId = UUID()
	}

	if !entity.ParentId.Valid {
		entity.ParentId = NewString(query.WorkspaceId)
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

func PermissionInfoToString(items []PermissionInfo) []string {
	res := []string{}

	for _, j := range items {
		res = append(res, j.CompleteKey)
	}

	return res
}

func SyncPermissionsInDatabase(x *FirebackApp, db *gorm.DB) {

	for _, item := range x.Modules {

		if item.BackupTables != nil && len(item.BackupTables) > 0 {
			for _, table := range item.BackupTables {

				GetDbRef().Model(&BackupTableMetaEntity{}).Create(&BackupTableMetaEntity{
					UniqueId:      table.EntityName,
					TableNameInDb: table.TableNameInDb,
				})
			}
		}

		// Insert the permissions into the database
		item.PermissionsProvider = append(item.PermissionsProvider, PermissionInfo{
			CompleteKey: ROOT_ALL_ACCESS,
		}, PermissionInfo{
			CompleteKey: ROOT_ALL_MODULES,
		})

		for _, perm := range item.PermissionsProvider {
			hasChildren := HasChildren(perm.CompleteKey, PermissionInfoToString(item.PermissionsProvider))
			UpsertPermission(&perm, hasChildren, db)
		}

		for _, bundle := range item.EntityBundles {
			for _, perm := range bundle.Permissions {
				hasChildren := HasChildren(perm.CompleteKey, PermissionInfoToString(bundle.Permissions))
				UpsertPermission(&perm, hasChildren, db)
			}
		}

	}

}

func ClearShot(str *string) {
	v := strings.ToLower(strings.TrimSpace(*str))
	*str = v
}

func UnsafeGetUserByPassportValue(value string, q QueryDSL) (*PassportEntity, *UserEntity, *IError) {

	// Check the passport if exists
	var item PassportEntity
	if err := GetRef(q).Model(&PassportEntity{}).Where(&PassportEntity{Value: value}).First(&item).Error; err != nil || item.Value == "" {

		return nil, nil, Create401Error(&WorkspacesMessages.PassportNotAvailable, []string{})
	}

	var user UserEntity
	if err := GetRef(q).Model(&UserEntity{}).Where(&UserEntity{UniqueId: item.UserId.String}).First(&user).Error; err != nil {
		return nil, nil, Create401Error(&WorkspacesMessages.PassportNotAvailable, []string{})
	}

	return &item, &user, nil
}

var TRUE = true
var FALSE = false

func validateRecaptcha(token string, RECAPTCHA_SECRET_KEY string) error {
	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify",
		url.Values{"secret": {RECAPTCHA_SECRET_KEY}, "response": {token}})

	if err != nil {
		return errors.New("failed to connect to reCAPTCHA service")
	}
	defer resp.Body.Close()

	var googleResp struct {
		Success bool `json:"success"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&googleResp); err != nil || !googleResp.Success {
		return errors.New("captcha verification failed")
	}

	return nil
}
