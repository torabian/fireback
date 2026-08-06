package abac

import (
	"reflect"

	"github.com/torabian/fireback/modules/abac/queries"
	"github.com/torabian/fireback/modules/fireback"
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

// WorkspaceTreeNode wraps WorkspaceEntity with a Children slice for the CTE tree result -
// see AppMenuTreeNode's doc comment for why this is a separate wrapper type rather than
// a field on the Emi-generated WorkspaceEntity struct.
type WorkspaceTreeNode struct {
	WorkspaceEntity
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
	EntityActionsBundle[WorkspaceEntity]
	CteQuery func(query fireback.QueryDSL) ([]*WorkspaceTreeNode, *fireback.QueryResultMeta, *fireback.IError)
}

var WorkspaceActions = workspaceActionsBundle{
	EntityActionsBundle: NewEntityActionsBundle[WorkspaceEntity](),
	CteQuery:             WorkspaceActionCteQueryFn,
}

// WorkspaceActionCteQueryFn fetches every WorkspaceEntity row matching query (via the
// recursive WorkspaceCte.vsql - see modules/abac/queries) and assembles them into a
// forest of WorkspaceTreeNode, one root per row with no parentId.
func WorkspaceActionCteQueryFn(query fireback.QueryDSL) ([]*WorkspaceTreeNode, *fireback.QueryResultMeta, *fireback.IError) {
	refl := reflect.ValueOf(&WorkspaceEntity{})
	items, meta, err := fireback.ContextAwareVSqlOperation[WorkspaceEntity](
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

func WorkspaceBrowseAction(c WorkspaceBrowseActionRequest) (*WorkspaceBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_WORKSPACE_QUERY}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := WorkspaceActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &WorkspaceBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func WorkspaceGetAction(c WorkspaceGetActionRequest) (*WorkspaceGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_WORKSPACE_QUERY}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := WorkspaceActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &WorkspaceGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func WorkspaceCreateAction(c WorkspaceCreateActionRequest) (*WorkspaceCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_WORKSPACE_CREATE}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	entity := &WorkspaceEntity{
		Description: c.Body.Description,
		Name:        c.Body.Name,
		TypeId:      c.Body.TypeId,
	}
	created, err2 := WorkspaceActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &WorkspaceCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func WorkspaceUpdateAction(c WorkspaceUpdateActionRequest) (*WorkspaceUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_WORKSPACE_UPDATE}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &WorkspaceEntity{UniqueId: c.Params.UniqueId}
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
	return &WorkspaceUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func WorkspaceAwareDeletePreviewAction(c WorkspaceAwareDeletePreviewActionRequest) (*WorkspaceAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_WORKSPACE_DELETE}, AllowOnRoot: true}); err != nil {
		return nil, err
	}
	uniqueIds := WorkspaceAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := WorkspaceEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &WorkspaceAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func WorkspaceAwareDeleteAction(c WorkspaceAwareDeleteActionRequest) (*WorkspaceAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_WORKSPACE_DELETE}, AllowOnRoot: true}); err != nil {
		return nil, err
	}
	if err2 := WorkspaceEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &WorkspaceAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

// WorkspaceCliFn mirrors the old Module3-generated grouped "workspace" cli command
// (minus the import/export/dev commands, which had no hand-written equivalent to
// recover), plus the "cte" subcommand for the recursive tree query. WorkspaceCliCommands
// (see WorkspaceCli.go) carries every entity-scoped cli group that doesn't have its own
// top-level command elsewhere (publicAuthentication, timezoneGroup, workspaceType,
// workspaceConfig, workspaceInvite, workspaceRole, userWorkspace, publicJoinKey, ...).
func WorkspaceCliFn() *cli.Command {
	commands := []*cli.Command{
		WorkspaceBrowseActionCliHandler(WorkspaceBrowseAction),
		WorkspaceGetActionCliHandler(WorkspaceGetAction),
		WorkspaceCreateActionCliHandler(WorkspaceCreateAction),
		WorkspaceUpdateActionCliHandler(WorkspaceUpdateAction),
		WorkspaceAwareDeletePreviewActionCliHandler(WorkspaceAwareDeletePreviewAction),
		WorkspaceAwareDeleteActionCliHandler(WorkspaceAwareDeleteAction),
		fireback.GetCommonCteQuery(WorkspaceActions.CteQuery),
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
