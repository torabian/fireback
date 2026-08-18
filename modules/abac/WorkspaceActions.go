package abac

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/abac/queries"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	"gorm.io/gorm"
)

// workspace had permRewrite root.modules -> root.manage, and
// security: { writeOnRoot: true, readOnRoot: true } in the old yaml: the old generated
// code set AllowOnRoot: true on every action (Query/Get/Create/Update/Delete) -
// preserved here exactly.
var workspacePerms = NewCrudPermissionSet("root.manage", "workspace", "workspace")
var PERM_ROOT_WORKSPACE = workspacePerms.Wildcard
var PERM_ROOT_WORKSPACE_QUERY = workspacePerms.Query
var PERM_ROOT_WORKSPACE_CREATE = workspacePerms.Create
var PERM_ROOT_WORKSPACE_UPDATE = workspacePerms.Update
var PERM_ROOT_WORKSPACE_DELETE = workspacePerms.Delete
var ALL_WORKSPACE_PERMISSIONS = workspacePerms.All

// WorkspaceMessages mirrors the entity-scoped messages pattern WorkspaceTypeMessages/
// RoleMessages use - there's no Emi equivalent for entity-scoped messages, so it's
// hand-declared here.
var WorkspaceMessages = struct {
	CannotDeleteRootWorkspace fireback.ErrorItem
	UniqueIdAlreadyExists     fireback.ErrorItem
}{
	CannotDeleteRootWorkspace: fireback.ErrorItem{
		"$": "CannotDeleteRootWorkspace", "en": "The root workspace cannot be deleted.",
	},
	UniqueIdAlreadyExists: fireback.ErrorItem{
		"$": "WorkspaceUniqueIdAlreadyExists", "en": "This id is already used by another workspace.",
	},
}

// workspaceUniqueIdTaken mirrors capabilityUniqueIdTaken (CapabilityActions.go) - used
// by WorkspaceCreateAction to give a clean 400 instead of a raw unique-constraint SQL
// error when a caller-supplied uniqueId collides with an existing workspace.
func workspaceUniqueIdTaken(uniqueId string) (bool, error) {
	var count int64
	err := fireback.GetDbRef().Model(&abacdefs.WorkspaceEntity{}).
		Where("unique_id = ?", uniqueId).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// WorkspaceTreeNode wraps WorkspaceEntity with a Children slice for the CTE tree result -
// see AppMenuTreeNode's doc comment for why this is a separate wrapper type rather than
// a field on the Emi-generated WorkspaceEntity struct.
type WorkspaceTreeNode struct {
	abacdefs.WorkspaceEntity
	Children []*WorkspaceTreeNode `json:"children,omitempty" yaml:"children,omitempty"`
}

func (dto *WorkspaceTreeNode) Size() int {
	size := len(dto.Children)
	for _, c := range dto.Children {
		size += c.Size()
	}
	return size
}

func (dto *WorkspaceTreeNode) Add(nodes ...*WorkspaceTreeNode) bool {
	size := dto.Size()
	for _, n := range nodes {
		if !n.ParentId.IsNull() && n.ParentId.OrDefault("") == dto.UniqueId {
			dto.Children = append(dto.Children, n)
		} else {
			for _, c := range dto.Children {
				if c.Add(n) {
					break
				}
			}
		}
	}
	return dto.Size() == size+len(nodes)
}

// workspaceActionsBundle extends the generic CRUD bundle with the CteQuery function the
// old <Entity>ActionsSig also had for this entity - see appMenuActionsBundle's doc
// comment.
type workspaceActionsBundle struct {
	EntityActionsBundle[abacdefs.WorkspaceEntity]
	CteQuery func(query fireback.QueryDSL) ([]*WorkspaceTreeNode, *fireback.QueryResultMeta, *fireback.IError)
}

var WorkspaceActions = workspaceActionsBundle{
	EntityActionsBundle: NewEntityActionsBundle[abacdefs.WorkspaceEntity](),
	CteQuery:            WorkspaceActionCteQueryFn,
}

// WorkspaceActionCteQueryFn fetches every abacdefs.WorkspaceEntity row matching query (via the
// recursive WorkspaceCte.vsql - see modules/abac/queries) and assembles them into a
// forest of WorkspaceTreeNode, one root per row with no parentId.
func WorkspaceActionCteQueryFn(query fireback.QueryDSL) ([]*WorkspaceTreeNode, *fireback.QueryResultMeta, *fireback.IError) {
	refl := reflect.ValueOf(&abacdefs.WorkspaceEntity{})
	items, meta, err := fireback.ContextAwareVSqlOperation[abacdefs.WorkspaceEntity](
		refl, &queries.QueriesFs, "WorkspaceCte.vsql", query,
	)
	if err != nil {
		return nil, meta, err
	}

	nodes := make([]*WorkspaceTreeNode, len(items))
	for i, item := range items {
		nodes[i] = &WorkspaceTreeNode{WorkspaceEntity: *item}
	}

	var tree []*WorkspaceTreeNode
	for _, node := range nodes {
		if !node.ParentId.IsSet() || node.ParentId.IsNull() {
			root := node
			root.Add(nodes...)
			tree = append(tree, root)
		}
	}
	return tree, meta, nil
}

func WorkspaceBrowseAction(c abacdefs.WorkspaceBrowseActionRequest) (*abacdefs.WorkspaceBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_WORKSPACE_QUERY}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := WorkspaceActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.WorkspaceBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func WorkspaceGetAction(c abacdefs.WorkspaceGetActionRequest) (*abacdefs.WorkspaceGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_WORKSPACE_QUERY}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := WorkspaceActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.WorkspaceGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func WorkspaceCreateAction(c abacdefs.WorkspaceCreateActionRequest) (*abacdefs.WorkspaceCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_WORKSPACE_CREATE}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	// name is validate:"required" (see WorkspaceDto.go) but nothing was actually
	// enforcing it - same fix as PassportCreateAction.go's own comment
	// (EmailProviderActions.go's EmailProviderCreateAction, TableViewSizingCreateAction).
	// Deliberately validating just this one field rather than the whole WorkspaceDto
	// (which also carries a required typeId): typeId isn't surfaced in the workspace
	// form yet (WorkspaceEditForm.tsx) and enforcing it here would break existing
	// workspace creation - both via the UI and manage-workspaces.cy.ts, neither of
	// which currently ever sets it.
	if err2 := fireback.CommonStructValidatorPointer(&struct {
		Name string `validate:"required"`
	}{Name: c.Body.Name}, false); err2 != nil {
		return nil, err2
	}
	// Bug fix: this never stamped WorkspaceId onto the created entity, so its
	// workspace_id column stayed null and it never matched its own creator's
	// workspace_id = query.WorkspaceId browse filter (see PERM_ROOT_WORKSPACE_QUERY's
	// AllowOnRoot: true - every call here runs in "root", same as query.WorkspaceId).
	entity := abacdefs.WorkspaceEntity{
		Description: c.Body.Description,
		Name:        c.Body.Name,
		TypeId:      c.Body.TypeId,
		ParentId:    c.Body.ParentId,
		WorkspaceId: emigo.NullableOf(query.WorkspaceId),
	}
	// UniqueId is optional on create (see WorkspaceDto.go) - like
	// CapabilityCreateAction, take it from the body when the caller sets one,
	// falling back to gen_random_uuid() (WorkspaceEntity's gorm tag) otherwise.
	// Editing an existing workspace's id is not possible either way:
	// WorkspaceUpdateAction below only ever reads a uniqueId from the URL path
	// parameter, never from its body.
	if v, ok := c.Body.UniqueId.Get(); ok && v != nil && *v != "" {
		taken, err3 := workspaceUniqueIdTaken(*v)
		if err3 != nil {
			return nil, fireback.GormErrorToIError(err3)
		}
		if taken {
			return nil, &fireback.IError{
				Message:  WorkspaceMessages.UniqueIdAlreadyExists,
				HttpCode: 400,
				Errors: []*fireback.IErrorItem{
					{Location: "uniqueId", Message: &WorkspaceMessages.UniqueIdAlreadyExists},
				},
			}
		}
		entity.UniqueId = *v
	}
	created, err2 := WorkspaceActions.Create(&entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.WorkspaceCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func WorkspaceUpdateAction(c abacdefs.WorkspaceUpdateActionRequest) (*abacdefs.WorkspaceUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_WORKSPACE_UPDATE}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	// Present for consistency with other *UpdateAction handlers and to catch any
	// future non-required constraint added to WorkspaceOptionalDto. It does not
	// (and structurally cannot) re-enforce name's required tag here the way
	// WorkspaceCreateAction now does on create: emigo.Nullable[string]'s own
	// "an empty string on the wire means the field was omitted" convention already
	// makes fields.Name below only ever get set from a genuinely non-empty value -
	// there is no wire representation of "explicitly blank out the name" to reject
	// in the first place.
	if err2 := fireback.CommonStructValidatorPointer(&c.Body, false); err2 != nil {
		return nil, err2
	}
	query.UniqueId = c.Params.UniqueId
	// uniqueId is deliberately never read from c.Body here - only from the URL path
	// parameter above - so an existing workspace's id can never be changed via this
	// action, regardless of what a caller puts in the request body.
	fields := &abacdefs.WorkspaceEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.Description.Get(); ok {
		fields.Description = *v
	}
	if v, ok := c.Body.Name.Get(); ok {
		fields.Name = *v
	}
	if v, ok := c.Body.TypeId.Get(); ok {
		fields.TypeId = *v
	}
	if v, ok := c.Body.ParentId.Get(); ok {
		fields.ParentId = emigo.NullableOf(*v)
	}
	updated, err2 := WorkspaceActionUpdate(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.WorkspaceUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func WorkspaceAwareDeletePreviewAction(c abacdefs.WorkspaceAwareDeletePreviewActionRequest) (*abacdefs.WorkspaceAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_WORKSPACE_DELETE}, AllowOnRoot: true}); err != nil {
		return nil, err
	}
	uniqueIds := abacdefs.WorkspaceAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	// The root workspace is the one every other workspace/role/permission bootstraps
	// from (see RepairTheWorkspaces, UserCli.go) - deleting it would leave the whole
	// install unusable, so it's rejected outright, same as RoleActions.go already does
	// for the root role. Checked here too (not just in WorkspaceAwareDeleteAction),
	// so the confirmation step itself already surfaces the problem instead of letting
	// an admin preview a "safe" delete that the real delete then rejects.
	if slices.Contains(uniqueIds, ROOT_VAR) {
		return nil, &fireback.IError{Message: WorkspaceMessages.CannotDeleteRootWorkspace, HttpCode: 400}
	}
	preview, err2 := abacdefs.WorkspaceEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.WorkspaceAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func WorkspaceAwareDeleteAction(c abacdefs.WorkspaceAwareDeleteActionRequest) (*abacdefs.WorkspaceAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_WORKSPACE_DELETE}, AllowOnRoot: true}); err != nil {
		return nil, err
	}
	// See WorkspaceAwareDeletePreviewAction's own comment - the root workspace must
	// never be deletable. Rejecting the whole batch (rather than silently dropping
	// just "root" from the list and deleting the rest) so a request that includes it
	// fails loudly instead of quietly succeeding on everything else.
	if slices.Contains(c.Body.UniqueIds, ROOT_VAR) {
		return nil, &fireback.IError{Message: WorkspaceMessages.CannotDeleteRootWorkspace, HttpCode: 400}
	}
	if err2 := abacdefs.WorkspaceEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.WorkspaceAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

/**
*  Call this when you are going to initialize a server, it will create root workspaces
*  It will create root workspace, assign the role to it.
 */
func RepairTheWorkspaces() error {
	{

		if role := GetRoleByUniqueId("root"); role == nil || role.UniqueId == "" {
			if _, err2 := CreateRootRoleInWorkspace("root"); err2 != nil {
				if !strings.Contains(err2.Error(), "Duplicate") {

					fmt.Println(err2)
				}
			}
		}
	}
	{
		item := &abacdefs.WorkspaceTypeEntity{}
		err := fireback.GetDbRef().Model(&abacdefs.WorkspaceTypeEntity{}).Where(&abacdefs.WorkspaceTypeEntity{UniqueId: "root"}).First(item).Error
		system := "system"
		if err == gorm.ErrRecordNotFound {
			err = fireback.GetDbRef().Create(&abacdefs.WorkspaceTypeEntity{WorkspaceId: emigo.NullableOf(system), UniqueId: "root", RoleId: ROOT_VAR}).Error
			if err != nil {
				return err
			}
		}
	}

	{

		item := &abacdefs.WorkspaceEntity{}
		err := fireback.GetDbRef().Model(&abacdefs.WorkspaceEntity{}).Where(&abacdefs.WorkspaceEntity{UniqueId: "root"}).First(item).Error

		description := "The root system which holds entire software data tree"
		if err == gorm.ErrRecordNotFound {
			err = fireback.GetDbRef().Create(&abacdefs.WorkspaceEntity{
				UniqueId: "root", Name: ROOT_VAR, Description: description,
				WorkspaceId: emigo.NullableOf(ROOT_VAR),
				TypeId:      ROOT_VAR,
			}).Error

			if err != nil {
				return err
			}

			_, err2 := CreateRootRoleInWorkspace("root")

			if err2 != nil && !strings.Contains(err2.Error(), "Duplicate") {
				return err2
			}

		}

		ws := GetWorkspaceByUniqueId("root")
		if ws == nil || ws.UniqueId != "root" {
			return errors.New(("ROOT_WORKSPACE_DOES_NOT_EXISTS"))
		}
	}

	{
		item := &abacdefs.WorkspaceEntity{}
		err := fireback.GetDbRef().Model(&abacdefs.WorkspaceEntity{}).Where(&abacdefs.WorkspaceEntity{UniqueId: "system"}).First(item).Error
		system := "system"
		if err == gorm.ErrRecordNotFound {
			description := "The workspace content which applies to everyworkspace"
			err = fireback.GetDbRef().Create(&abacdefs.WorkspaceEntity{WorkspaceId: emigo.NullableOf(system), UniqueId: "system", Name: system, Description: description}).Error

			if err != nil {
				return err
			}
		}

		ws := GetWorkspaceByUniqueId("system")
		if ws == nil || ws.UniqueId != "system" {
			return errors.New(("SYSTEM_WORKSPACE_DOES_NOT_EXISTS"))
		}
	}

	return nil
}

func CreateRootRoleInWorkspace(workspaceId string) (*abacdefs.RoleEntity, error) {
	sampleName := "Root Access"
	entity := &abacdefs.RoleEntity{
		UniqueId:           "root",
		WorkspaceId:        emigo.NullableOf(ROOT_VAR),
		Name:               sampleName,
		CapabilitiesListId: RoleCapabilitiesListIdOf([]string{ROOT_ALL_ACCESS}),
	}

	err := fireback.GetDbRef().
		Where(&abacdefs.RoleEntity{UniqueId: "root"}).
		FirstOrCreate(&entity).Error

	return entity, err
}
