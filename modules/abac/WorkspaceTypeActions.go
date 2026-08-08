package abac

import (
	"regexp"
	"strings"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
)

// workspaceType had permRewrite root.modules -> root.manage, and
// security: { writeOnRoot: true, readOnRoot: true } in the old yaml: the old generated
// code set AllowOnRoot: true on every action (Query/Get/Create/Update/Delete) -
// preserved here exactly. features.mock/features.msync were both explicitly false in
// the old yaml, so no mock/msync support is needed here.
var workspaceTypePerms = NewCrudPermissionSet("root.manage", "workspace-type", "workspace type")
var PERM_ROOT_WORKSPACE_TYPE = workspaceTypePerms.Wildcard
var PERM_ROOT_WORKSPACE_TYPE_QUERY = workspaceTypePerms.Query
var PERM_ROOT_WORKSPACE_TYPE_CREATE = workspaceTypePerms.Create
var PERM_ROOT_WORKSPACE_TYPE_UPDATE = workspaceTypePerms.Update
var PERM_ROOT_WORKSPACE_TYPE_DELETE = workspaceTypePerms.Delete
var ALL_WORKSPACE_TYPE_PERMISSIONS = workspaceTypePerms.All

// WorkspaceTypeMessages mirrors the entity-scoped messages: block workspaceType had in
// the old yaml - there's no Emi equivalent for entity-scoped messages, so it's
// hand-declared here.
var WorkspaceTypeMessages = struct {
	RoleIsNecessary               fireback.ErrorItem
	RoleIsNotAccessible           fireback.ErrorItem
	OnlyRootRoleIsAccepted        fireback.ErrorItem
	RoleNeedsToHaveCapabilities   fireback.ErrorItem
	RoleCannotBeSuperAdmin        fireback.ErrorItem
	CannotCreateWorkspaceType     fireback.ErrorItem
	CannotModifyWorkspaceType     fireback.ErrorItem
	CannotDeleteRootWorkspaceType fireback.ErrorItem
	SlugMustStartWithSlash        fireback.ErrorItem
	SlugInvalidFormat             fireback.ErrorItem
	SlugTooLong                   fireback.ErrorItem
	SlugAlreadyTaken              fireback.ErrorItem
}{
	RoleIsNecessary: fireback.ErrorItem{
		"$": "RoleIsNecessary", "en": "Role needs to be defined and exist.",
	},
	RoleIsNotAccessible: fireback.ErrorItem{
		"$": "RoleIsNotAccessible", "en": "Role is not accessible unfortunately. Make sure you the role chose exists.",
	},
	OnlyRootRoleIsAccepted: fireback.ErrorItem{
		"$": "OnlyRootRoleIsAccepted", "en": "You can only select a role which is created or belong to 'root' workspace.",
	},
	RoleNeedsToHaveCapabilities: fireback.ErrorItem{
		"$": "RoleNeedsToHaveCapabilities", "en": "Role needs to have at least one capability before could be assigned.",
	},
	RoleCannotBeSuperAdmin: fireback.ErrorItem{
		"$": "RoleCannotBeSuperAdmin", "en": "The 'root' role, or any role holding the root.* wildcard capability, cannot be assigned to a workspace type.",
	},
	CannotCreateWorkspaceType: fireback.ErrorItem{
		"$": "CannotCreateWorkspaceType", "en": "You cannot create workspace type due to some validation errors.",
	},
	CannotModifyWorkspaceType: fireback.ErrorItem{
		"$": "CannotModifyWorkspaceType", "en": "You cannot modify workspace type due to some validation errors.",
	},
	CannotDeleteRootWorkspaceType: fireback.ErrorItem{
		"$": "CannotDeleteRootWorkspaceType", "en": "The 'root' workspace type is required by the system and cannot be deleted.",
	},
	SlugMustStartWithSlash: fireback.ErrorItem{
		"$": "SlugMustStartWithSlash", "en": "Slug must start with '/'.",
	},
	SlugInvalidFormat: fireback.ErrorItem{
		"$": "SlugInvalidFormat", "en": "Slug can only contain lowercase letters (a-z) and dashes after the leading '/'.",
	},
	SlugTooLong: fireback.ErrorItem{
		"$": "SlugTooLong", "en": "Slug cannot be longer than 50 characters.",
	},
	SlugAlreadyTaken: fireback.ErrorItem{
		"$": "SlugAlreadyTaken", "en": "This slug is already used by another workspace type.",
	},
}

var WorkspaceTypeActions = NewEntityActionsBundle[WorkspaceTypeEntity]()

func workspaceTypeSecurity(perm fireback.PermissionInfo) *fireback.SecurityModel {
	return &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{perm}, AllowOnRoot: true}
}

func WorkspaceTypeBrowseAction(c WorkspaceTypeBrowseActionRequest) (*WorkspaceTypeBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, workspaceTypeSecurity(PERM_ROOT_WORKSPACE_TYPE_QUERY))
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := WorkspaceTypeActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &WorkspaceTypeBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func WorkspaceTypeGetAction(c WorkspaceTypeGetActionRequest) (*WorkspaceTypeGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, workspaceTypeSecurity(PERM_ROOT_WORKSPACE_TYPE_QUERY))
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := WorkspaceTypeActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &WorkspaceTypeGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func WorkspaceTypeCreateAction(c WorkspaceTypeCreateActionRequest) (*WorkspaceTypeCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, workspaceTypeSecurity(PERM_ROOT_WORKSPACE_TYPE_CREATE))
	if err != nil {
		return nil, err
	}
	entity := &WorkspaceTypeEntity{
		Title:       c.Body.Title,
		Description: c.Body.Description,
		Slug:        c.Body.Slug,
		RoleId:      c.Body.RoleId,
		// workspaceTypeSecurity's AllowOnRoot already rejects this whole action unless
		// query.WorkspaceId == "root" (see ResolveActionContext ->
		// WithAuthorizationPureDefault), so this is always "root" here - set explicitly
		// rather than left unset/null, since nothing else ever populated it before.
		WorkspaceId: emigo.NullableOf(query.WorkspaceId),
	}
	created, err2 := WorkspaceTypeActionCreate(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &WorkspaceTypeCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func WorkspaceTypeUpdateAction(c WorkspaceTypeUpdateActionRequest) (*WorkspaceTypeUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, workspaceTypeSecurity(PERM_ROOT_WORKSPACE_TYPE_UPDATE))
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &WorkspaceTypeEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.Title.Get(); ok {
		fields.Title = *v
	}
	if v, ok := c.Body.Description.Get(); ok {
		fields.Description = *v
	}
	if v, ok := c.Body.Slug.Get(); ok {
		fields.Slug = *v
	}
	if v, ok := c.Body.RoleId.Get(); ok {
		fields.RoleId = *v
	}
	updated, err2 := WorkspaceTypeActionUpdate(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &WorkspaceTypeUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func WorkspaceTypeAwareDeletePreviewAction(c WorkspaceTypeAwareDeletePreviewActionRequest) (*WorkspaceTypeAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, workspaceTypeSecurity(PERM_ROOT_WORKSPACE_TYPE_DELETE)); err != nil {
		return nil, err
	}
	uniqueIds := WorkspaceTypeAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	if err := RejectRootWorkspaceTypeDeletion(uniqueIds); err != nil {
		return nil, err
	}
	preview, err2 := WorkspaceTypeEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &WorkspaceTypeAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func WorkspaceTypeAwareDeleteAction(c WorkspaceTypeAwareDeleteActionRequest) (*WorkspaceTypeAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, workspaceTypeSecurity(PERM_ROOT_WORKSPACE_TYPE_DELETE)); err != nil {
		return nil, err
	}
	if err := RejectRootWorkspaceTypeDeletion(c.Body.UniqueIds); err != nil {
		return nil, err
	}
	if err2 := WorkspaceTypeEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &WorkspaceTypeAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

// RejectRootWorkspaceTypeDeletion blocks the seeded "root" workspace type (created in
// UserCli.go alongside the root workspace/user, RoleId: ROOT_VAR) from ever being
// deleted - it's a system fixture every root-workspace signup/login flow depends on,
// not a regular user-created workspace type.
func RejectRootWorkspaceTypeDeletion(uniqueIds []string) *fireback.IError {
	for _, id := range uniqueIds {
		if id == ROOT_VAR {
			return &fireback.IError{
				Message:  WorkspaceTypeMessages.CannotDeleteRootWorkspaceType,
				HttpCode: 400,
			}
		}
	}
	return nil
}

// --- Hand business logic recovered from the pre-migration WorkspaceTypeEntity.go ---

func WorkspaceTypeActionCreate(
	dto *WorkspaceTypeEntity, query fireback.QueryDSL,
) (*WorkspaceTypeEntity, *fireback.IError) {

	if errors := ValidateTheWorkspaceTypeEntity(dto); len(errors) > 0 {
		return nil, &fireback.IError{
			Message:  WorkspaceTypeMessages.CannotCreateWorkspaceType,
			HttpCode: 400,
			Errors:   errors,
		}
	}

	created, err := WorkspaceTypeActions.Create(dto, query)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func WorkspaceTypeActionUpdate(
	query fireback.QueryDSL,
	fields *WorkspaceTypeEntity,
) (*WorkspaceTypeEntity, *fireback.IError) {

	if errors := ValidateTheWorkspaceTypeEntity(fields); len(errors) > 0 {
		return nil, &fireback.IError{
			Message:  WorkspaceTypeMessages.CannotModifyWorkspaceType,
			HttpCode: 400,
			Errors:   errors,
		}
	}

	return WorkspaceTypeActions.Update(query, fields)
}

func ValidateRoleAndItsExistence(roleId emigo.Nullable[string]) (*RoleEntity, []*fireback.IErrorItem) {
	items := []*fireback.IErrorItem{}

	// Bug fix: this condition was inverted (`!roleId.IsNull()`), so every workspace
	// type create/update - even with a perfectly valid roleId - was rejected here
	// before ever reaching the actual existence/capability checks below. It should
	// reject when roleId is *absent*, not when it's present.
	if roleId.IsNull() {
		items = append(items, &fireback.IErrorItem{
			Location: "roleId",
			Message:  &WorkspaceTypeMessages.RoleIsNecessary,
		})

		return nil, items
	}

	if role, err := RoleActions.GetOne(fireback.QueryDSL{UniqueId: roleId.OrDefault("")}); err != nil {
		items = append(items, &fireback.IErrorItem{
			Location: "roleId",
			Message:  &WorkspaceTypeMessages.RoleIsNotAccessible,
		})
		return nil, items
	} else {
		if role == nil {
			items = append(items, &fireback.IErrorItem{
				Location: "roleId",
				Message:  &WorkspaceTypeMessages.RoleIsNotAccessible,
			})

			return nil, items
		} else {
			if len(RoleCapabilitiesListIdGet(role)) == 0 {
				items = append(items, &fireback.IErrorItem{
					Location: "roleId",
					Message:  &WorkspaceTypeMessages.RoleNeedsToHaveCapabilities,
				})
				return nil, items
			}

			return role, nil
		}
	}
}

// Before write or update we need some extra validation for this.
// It's important to check if the role actually exists, and has some previliges
// before making it available
func ValidateTheWorkspaceTypeEntity(fields *WorkspaceTypeEntity) []*fireback.IErrorItem {
	items := []*fireback.IErrorItem{}
	role, roleErrors := ValidateRoleAndItsExistence(emigo.NullableOf(fields.RoleId))
	if len(roleErrors) != 0 {
		return roleErrors
	}

	if !role.WorkspaceId.IsSet() || role.WorkspaceId.OrDefault("") != ROOT_VAR {
		items = append(items, &fireback.IErrorItem{
			Location: "roleId",
			Message:  &WorkspaceTypeMessages.OnlyRootRoleIsAccepted,
		})

		return items
	}

	// A workspace type effectively hands every user of that workspace super-admin
	// powers if its role is (or acts like) "root": either the literal reserved 'root'
	// role itself, or any role that's been granted the root.* wildcard capability.
	if role.UniqueId == ROOT_VAR || roleHasRootWildcardCapability(role) {
		items = append(items, &fireback.IErrorItem{
			Location: "roleId",
			Message:  &WorkspaceTypeMessages.RoleCannotBeSuperAdmin,
		})

		return items
	}

	// fields.UniqueId is only ever populated on an update (WorkspaceTypeUpdateAction
	// seeds it from the path parameter; WorkspaceTypeCreateAction never sets it) - so
	// it doubles as "am I updating, and if so, which record can this slug collide with
	// itself on" for the uniqueness check below.
	items = append(items, ValidateWorkspaceTypeSlug(fields.Slug, fields.UniqueId)...)

	return items
}

// roleHasRootWildcardCapability reports whether role has been granted the root.*
// wildcard capability - the same capability the reserved 'root' role itself carries,
// and effectively equivalent to it for authorization purposes even if the role wasn't
// literally created as 'root'.
func roleHasRootWildcardCapability(role *RoleEntity) bool {
	for _, capability := range RoleCapabilitiesListIdGet(role) {
		if capability == "root.*" {
			return true
		}
	}
	return false
}

// workspaceTypeSlugPattern matches a leading "/" followed by one or more lowercase
// letters and/or dashes, and nothing else - e.g. "/school", "/back-office".
var workspaceTypeSlugPattern = regexp.MustCompile(`^/[a-z-]+$`)

// ValidateWorkspaceTypeSlug enforces every rule the slug field needs: present, at most
// 50 characters, starts with "/", only lowercase a-z and dashes after that, and unique
// across every other workspace type. excludeUniqueId should be the record's own
// uniqueId on an update (so its own unchanged slug doesn't collide with itself), or ""
// on create.
func ValidateWorkspaceTypeSlug(slug string, excludeUniqueId string) []*fireback.IErrorItem {
	items := []*fireback.IErrorItem{}

	if len(slug) > 50 {
		items = append(items, &fireback.IErrorItem{
			Location: "slug",
			Message:  &WorkspaceTypeMessages.SlugTooLong,
		})
		return items
	}

	if !strings.HasPrefix(slug, "/") {
		items = append(items, &fireback.IErrorItem{
			Location: "slug",
			Message:  &WorkspaceTypeMessages.SlugMustStartWithSlash,
		})
		return items
	}

	if !workspaceTypeSlugPattern.MatchString(slug) {
		items = append(items, &fireback.IErrorItem{
			Location: "slug",
			Message:  &WorkspaceTypeMessages.SlugInvalidFormat,
		})
		return items
	}

	taken, err := workspaceTypeSlugTaken(slug, excludeUniqueId)
	if err != nil {
		items = append(items, &fireback.IErrorItem{
			Location: "slug",
			Message:  &fireback.FirebackMessages.DatabaseOperationError,
		})
		return items
	}
	if taken {
		items = append(items, &fireback.IErrorItem{
			Location: "slug",
			Message:  &WorkspaceTypeMessages.SlugAlreadyTaken,
		})
	}

	return items
}

// workspaceTypeSlugTaken reports whether slug is already used by some other workspace
// type. This is a friendly pre-check ahead of the gorm:"unique" constraint on
// WorkspaceTypeEntity.Slug (see Abac.emi.yml) - that constraint is still the real
// backstop against a race between this check and the insert, but hitting it directly
// would surface as an opaque database error instead of this field-scoped one.
func workspaceTypeSlugTaken(slug string, excludeUniqueId string) (bool, error) {
	var count int64
	q := fireback.GetDbRef().Model(&WorkspaceTypeEntity{}).Where("slug = ?", slug)
	if excludeUniqueId != "" {
		q = q.Where("unique_id <> ?", excludeUniqueId)
	}
	if err := q.Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func WorkspaceTypeActionPublicQuery(query fireback.QueryDSL) ([]*QueryWorkspaceTypesPubliclyActionRes, *fireback.QueryResultMeta, *fireback.IError) {
	// Make this API public, so the signup screen can get it.
	// At this moment, we just move things back as are, but maybe later we need
	// to add some limits on what kind of information is going out.
	query.WorkspaceId = "root"
	query.UserId = "root"

	items, qr, err := WorkspaceTypeActions.Query(query)

	var all []*QueryWorkspaceTypesPubliclyActionRes

	for _, item := range items {
		if item.UniqueId == "root" {
			continue
		}

		all = append(all, &QueryWorkspaceTypesPubliclyActionRes{
			Title:       item.Title,
			Description: item.Description,
			UniqueId:    item.UniqueId,
			Slug:        item.Slug,
		})
	}

	return all, qr, err
}
