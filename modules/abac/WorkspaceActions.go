package abac

import (
	"reflect"
	"slices"

	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/abac/queries"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"
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
}{
	CannotDeleteRootWorkspace: fireback.ErrorItem{
		"$": "CannotDeleteRootWorkspace", "en": "The root workspace cannot be deleted.",
	},
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
	// Bug fix: this never stamped WorkspaceId onto the created entity, so its
	// workspace_id column stayed null and it never matched its own creator's
	// workspace_id = query.WorkspaceId browse filter (see PERM_ROOT_WORKSPACE_QUERY's
	// AllowOnRoot: true - every call here runs in "root", same as query.WorkspaceId).
	entity := abacdefs.WorkspaceEntity{
		Description: c.Body.Description,
		Name:        c.Body.Name,
		TypeId:      c.Body.TypeId,
		WorkspaceId: emigo.NullableOf(query.WorkspaceId),
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
	query.UniqueId = c.Params.UniqueId
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

// WorkspaceCliFn mirrors the old Module3-generated grouped "workspace" cli command
// (minus the import/export/dev commands, which had no hand-written equivalent to
// recover), plus the "cte" subcommand for the recursive tree query. WorkspaceCliCommands
// (see WorkspaceCli.go) carries every entity-scoped cli group that doesn't have its own
// top-level command elsewhere (publicAuthentication, timezoneGroup, workspaceType,
// workspaceConfig, workspaceInvite, workspaceRole, userWorkspace, publicJoinKey, ...).
func WorkspaceCliFn() *cli.Command {
	commands := []*cli.Command{
		abacdefs.WorkspaceBrowseActionCliHandler(WorkspaceBrowseAction),
		abacdefs.WorkspaceGetActionCliHandler(WorkspaceGetAction),
		abacdefs.WorkspaceCreateActionCliHandler(WorkspaceCreateAction),
		abacdefs.WorkspaceUpdateActionCliHandler(WorkspaceUpdateAction),
		abacdefs.WorkspaceAwareDeletePreviewActionCliHandler(WorkspaceAwareDeletePreviewAction),
		abacdefs.WorkspaceAwareDeleteActionCliHandler(WorkspaceAwareDeleteAction),
	}
	commands = append(commands, WorkspaceCliCommands...)
	return &cli.Command{
		Name:        "workspace",
		Aliases:     []string{"ws"},
		Description: `Fireback general user role, workspaces services.`,
		Usage:       `Fireback general user role, workspaces services.`,
		Commands:    commands,
	}
}
