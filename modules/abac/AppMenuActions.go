package abac

import (
	"reflect"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/queries"
	appMenuSeeders "github.com/torabian/fireback/modules/abac/seeders/AppMenu"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli/v3"
)

func AppMenuSyncSeeders() {
	fireback.SeederFromFSImport(
		fireback.QueryDSL{WorkspaceId: fireback.USER_SYSTEM},
		AppMenuActions.Create,
		reflect.ValueOf(&AppMenuEntity{}).Elem(),
		&appMenuSeeders.ViewsFs,
		[]string{},
		true,
	)
}

// appMenu had no security: block in the old yaml, so its old generated code used the
// default SecurityModel on every action (no AllowOnRoot) - preserved here exactly.
var appMenuPerms = NewCrudPermissionSet("root.modules", "app-menu", "app menu")
var PERM_ROOT_APP_MENU = appMenuPerms.Wildcard
var PERM_ROOT_APP_MENU_QUERY = appMenuPerms.Query
var PERM_ROOT_APP_MENU_CREATE = appMenuPerms.Create
var PERM_ROOT_APP_MENU_UPDATE = appMenuPerms.Update
var PERM_ROOT_APP_MENU_DELETE = appMenuPerms.Delete
var ALL_APP_MENU_PERMISSIONS = appMenuPerms.All

// AppMenuTreeNode wraps AppMenuEntity with a Children slice for the CTE tree result.
// AppMenuEntity itself can't carry a Children field (it's an Emi-generated struct,
// regenerated from Abac.emi.yml on every recompile) - embedding it here keeps the
// JSON response shape identical to the old Module3 output (Children flattens onto the
// same object via anonymous embedding) without touching Emi or the generated entity.
type AppMenuTreeNode struct {
	AppMenuEntity
	Children []*AppMenuTreeNode `json:"children,omitempty" yaml:"children,omitempty"`
}

func (dto *AppMenuTreeNode) Size() int {
	size := len(dto.Children)
	for _, c := range dto.Children {
		size += c.Size()
	}
	return size
}

func (dto *AppMenuTreeNode) Add(nodes ...*AppMenuTreeNode) bool {
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

// appMenuActionsBundle extends the generic CRUD bundle with the CteQuery function the
// old <Entity>ActionsSig also had for this entity - CTE/tree support has no Emi
// equivalent, so it's a hand-added field here, not something the generic
// EntityActionsBundle[T] provides.
type appMenuActionsBundle struct {
	EntityActionsBundle[AppMenuEntity]
	CteQuery func(query fireback.QueryDSL) ([]*AppMenuTreeNode, *fireback.QueryResultMeta, *fireback.IError)
}

var AppMenuActions = appMenuActionsBundle{
	EntityActionsBundle: NewEntityActionsBundle[AppMenuEntity](),
	CteQuery:             AppMenuActionCteQueryFn,
}

// AppMenuActionCteQueryFn fetches every AppMenuEntity row matching query (via the
// recursive AppMenuCte.vsql - see modules/abac/queries) and assembles them into a
// forest of AppMenuTreeNode, one root per row with no parentId.
func AppMenuActionCteQueryFn(query fireback.QueryDSL) ([]*AppMenuTreeNode, *fireback.QueryResultMeta, *fireback.IError) {
	refl := reflect.ValueOf(&AppMenuEntity{})
	items, meta, err := fireback.ContextAwareVSqlOperation[AppMenuEntity](
		refl, &queries.QueriesFs, "AppMenuCte.vsql", query,
	)
	if err != nil {
		return nil, meta, err
	}

	nodes := make([]*AppMenuTreeNode, len(items))
	for i, item := range items {
		nodes[i] = &AppMenuTreeNode{AppMenuEntity: *item}
	}

	var tree []*AppMenuTreeNode
	for _, node := range nodes {
		if !node.ParentId.IsSet() || node.ParentId.IsNull() {
			root := node
			root.Add(nodes...)
			tree = append(tree, root)
		}
	}
	return tree, meta, nil
}

func AppMenuBrowseAction(c AppMenuBrowseActionRequest) (*AppMenuBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_APP_MENU_QUERY}})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := AppMenuActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &AppMenuBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func AppMenuGetAction(c AppMenuGetActionRequest) (*AppMenuGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_APP_MENU_QUERY}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := AppMenuActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &AppMenuGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func AppMenuCreateAction(c AppMenuCreateActionRequest) (*AppMenuCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_APP_MENU_CREATE}})
	if err != nil {
		return nil, err
	}
	entity := &AppMenuEntity{
		Label:         c.Body.Label,
		Href:          c.Body.Href,
		Icon:          c.Body.Icon,
		ActiveMatcher: c.Body.ActiveMatcher,
		CapabilityId:  c.Body.CapabilityId,
		ParentId:      c.Body.ParentId,
	}
	created, err2 := AppMenuActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &AppMenuCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func AppMenuUpdateAction(c AppMenuUpdateActionRequest) (*AppMenuUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_APP_MENU_UPDATE}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &AppMenuEntity{UniqueId: c.Params.UniqueId}
	// complex fields (TString included) have no Nullable/partial-update variant, so
	// unlike the plain-string fields below, this is always assigned wholesale - see
	// the same pattern for BirthDate/CreatedAt/UpdatedAt elsewhere in this migration.
	fields.Label = c.Body.Label
	if v, ok := c.Body.Href.Get(); ok {
		fields.Href = *v
	}
	if v, ok := c.Body.Icon.Get(); ok {
		fields.Icon = *v
	}
	if v, ok := c.Body.ActiveMatcher.Get(); ok {
		fields.ActiveMatcher = *v
	}
	if v, ok := c.Body.CapabilityId.Get(); ok {
		fields.CapabilityId = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.ParentId.Get(); ok {
		fields.ParentId = emigo.NullableOf(*v)
	}
	updated, err2 := AppMenuActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &AppMenuUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func AppMenuAwareDeletePreviewAction(c AppMenuAwareDeletePreviewActionRequest) (*AppMenuAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_APP_MENU_DELETE}}); err != nil {
		return nil, err
	}
	uniqueIds := AppMenuAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := AppMenuEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &AppMenuAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func AppMenuAwareDeleteAction(c AppMenuAwareDeleteActionRequest) (*AppMenuAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_APP_MENU_DELETE}}); err != nil {
		return nil, err
	}
	if err2 := AppMenuEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &AppMenuAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

// AppMenuCliFn mirrors the old Module3-generated grouped "appmenu" cli command (minus
// the import/export/mock/dev commands, which had no hand-written equivalent to recover).
func AppMenuCliFn() *cli.Command {
	return &cli.Command{
		Name:        "appmenu",
		Description: `Manages the menus in the app, (for example tab views, sidebar items, etc.)`,
		Usage:       `Manages the menus in the app, (for example tab views, sidebar items, etc.)`,
		Commands: []*cli.Command{
			AppMenuBrowseActionCliHandler(AppMenuBrowseAction),
			AppMenuGetActionCliHandler(AppMenuGetAction),
			AppMenuCreateActionCliHandler(AppMenuCreateAction),
			AppMenuUpdateActionCliHandler(AppMenuUpdateAction),
			AppMenuAwareDeletePreviewActionCliHandler(AppMenuAwareDeletePreviewAction),
			AppMenuAwareDeleteActionCliHandler(AppMenuAwareDeleteAction),
		},
	}
}
