package interfacetools

import (
	"log"
	"reflect"

	"github.com/torabian/emi/emigo"
	interfacetoolsdefs "github.com/torabian/fireback/modules/abac/interfacetools/defs"
	"github.com/torabian/fireback/modules/abac/interfacetools/queries"
	appMenuSeeders "github.com/torabian/fireback/modules/abac/interfacetools/seeders/AppMenu"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"
	"gorm.io/gorm"
)

func AppMenuSyncSeeders() {
	fireback.SeederFromFSImport(
		fireback.QueryDSL{WorkspaceId: fireback.USER_SYSTEM},
		AppMenuActions.Create,
		reflect.ValueOf(&interfacetoolsdefs.AppMenuEntity{}).Elem(),
		&appMenuSeeders.ViewsFs,
		[]string{},
		true,
	)
}

// SyncAppMenu upserts a single AppMenu row by uniqueId: if it has one, any existing row
// with the same uniqueId is removed first, then the item is (re)created fresh - so
// re-running this against an already-synced menu replaces it instead of piling up
// duplicates, which is what fireback.SeederFromFSImport/AppMenuSyncSeeders above (a bare
// insert, no de-duplication at all) would do on a second run. Used by
// ModuleSetup/InterfaceToolsModuleConfig.ExtraAppMenus to keep vendor/host-app-provided
// menu definitions (e.g. abac.Menu) in sync as part of this module's own migration,
// rather than requiring a separate manual "seeders" step.
//
// If the item has no uniqueId there's nothing to de-duplicate against, so it's just
// inserted directly. If removing the existing row fails (a genuine DB error - GORM's
// Delete does not error just because there was nothing to delete), the item is skipped
// entirely rather than risking a duplicate insert alongside whatever's already there.
func SyncAppMenu(item *interfacetoolsdefs.AppMenuEntity) error {
	return fireback.GetDbRef().Transaction(func(tx *gorm.DB) error {
		if item.UniqueId != "" {
			if err := tx.Where("unique_id = ?", item.UniqueId).Delete(&interfacetoolsdefs.AppMenuEntity{}).Error; err != nil {
				return err
			}
		}
		return tx.Create(item).Error
	})
}

// SyncAppMenus calls SyncAppMenu for every item, logging (rather than aborting on) any
// individual failure - one broken item shouldn't stop every other, unrelated one from
// syncing.
func SyncAppMenus(items []*interfacetoolsdefs.AppMenuEntity) {
	for _, item := range items {
		if err := SyncAppMenu(item); err != nil {
			log.Println("Skipping app menu item", item.UniqueId, "- sync failed:", err)
		}
	}
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
// regenerated from InterfaceTools.emi.yml on every recompile) - embedding it here keeps
// the JSON response shape identical to the old Module3 output (Children flattens onto
// the same object via anonymous embedding) without touching Emi or the generated entity.
type AppMenuTreeNode struct {
	interfacetoolsdefs.AppMenuEntity
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
	EntityActionsBundle[interfacetoolsdefs.AppMenuEntity]
	CteQuery func(query fireback.QueryDSL) ([]*AppMenuTreeNode, *fireback.QueryResultMeta, *fireback.IError)
}

var AppMenuActions = appMenuActionsBundle{
	EntityActionsBundle: NewEntityActionsBundle[interfacetoolsdefs.AppMenuEntity](),
	CteQuery:            AppMenuActionCteQueryFn,
}

// AppMenuActionCteQueryFn fetches every interfacetoolsdefs.AppMenuEntity row matching query (via the
// recursive AppMenuCte.vsql - see modules/abac/interfacetools/queries) and assembles
// them into a forest of AppMenuTreeNode, one root per row with no parentId.
func AppMenuActionCteQueryFn(query fireback.QueryDSL) ([]*AppMenuTreeNode, *fireback.QueryResultMeta, *fireback.IError) {
	refl := reflect.ValueOf(&interfacetoolsdefs.AppMenuEntity{})
	items, meta, err := fireback.ContextAwareVSqlOperation[interfacetoolsdefs.AppMenuEntity](
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

func AppMenuBrowseAction(c interfacetoolsdefs.AppMenuBrowseActionRequest) (*interfacetoolsdefs.AppMenuBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_APP_MENU_QUERY}})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := AppMenuActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &interfacetoolsdefs.AppMenuBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func AppMenuGetAction(c interfacetoolsdefs.AppMenuGetActionRequest) (*interfacetoolsdefs.AppMenuGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_APP_MENU_QUERY}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := AppMenuActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &interfacetoolsdefs.AppMenuGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func AppMenuCreateAction(c interfacetoolsdefs.AppMenuCreateActionRequest) (*interfacetoolsdefs.AppMenuCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_APP_MENU_CREATE}})
	if err != nil {
		return nil, err
	}
	entity := &interfacetoolsdefs.AppMenuEntity{
		Label:         c.Body.Label,
		Href:          c.Body.Href,
		Icon:          c.Body.Icon,
		ActiveMatcher: c.Body.ActiveMatcher,
		CapabilityId:  c.Body.CapabilityId,
		ParentId:      c.Body.ParentId,
		// Without this, the row is invisible to AppMenuBrowseAction's workspace-scoped
		// query (see GetSqlContext) - same workspace-stamping fix as
		// modules/abac/messaging/EmailProviderActions.go's EmailProviderCreateAction.
		WorkspaceId: emigo.NullableOf(query.WorkspaceId),
		UserId:      emigo.NullableOf(query.UserId),
	}
	created, err2 := AppMenuActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &interfacetoolsdefs.AppMenuCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func AppMenuUpdateAction(c interfacetoolsdefs.AppMenuUpdateActionRequest) (*interfacetoolsdefs.AppMenuUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_APP_MENU_UPDATE}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &interfacetoolsdefs.AppMenuEntity{UniqueId: c.Params.UniqueId}
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
	return &interfacetoolsdefs.AppMenuUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func AppMenuAwareDeletePreviewAction(c interfacetoolsdefs.AppMenuAwareDeletePreviewActionRequest) (*interfacetoolsdefs.AppMenuAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_APP_MENU_DELETE}}); err != nil {
		return nil, err
	}
	uniqueIds := interfacetoolsdefs.AppMenuAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := interfacetoolsdefs.AppMenuEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &interfacetoolsdefs.AppMenuAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func AppMenuAwareDeleteAction(c interfacetoolsdefs.AppMenuAwareDeleteActionRequest) (*interfacetoolsdefs.AppMenuAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_APP_MENU_DELETE}}); err != nil {
		return nil, err
	}
	if err2 := interfacetoolsdefs.AppMenuEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &interfacetoolsdefs.AppMenuAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

// AppMenuCliFn mirrors the old Module3-generated grouped "appmenu" cli command (minus
// the import/export/mock/dev commands, which had no hand-written equivalent to recover).
func AppMenuCliFn() *cli.Command {
	return &cli.Command{
		Name:        "appmenu",
		Description: `Manages the menus in the app, (for example tab views, sidebar items, etc.)`,
		Usage:       `Manages the menus in the app, (for example tab views, sidebar items, etc.)`,
		Commands: []*cli.Command{
			interfacetoolsdefs.AppMenuBrowseActionCliHandler(AppMenuBrowseAction),
			interfacetoolsdefs.AppMenuGetActionCliHandler(AppMenuGetAction),
			interfacetoolsdefs.AppMenuCreateActionCliHandler(AppMenuCreateAction),
			interfacetoolsdefs.AppMenuUpdateActionCliHandler(AppMenuUpdateAction),
			interfacetoolsdefs.AppMenuAwareDeletePreviewActionCliHandler(AppMenuAwareDeletePreviewAction),
			interfacetoolsdefs.AppMenuAwareDeleteActionCliHandler(AppMenuAwareDeleteAction),
			interfacetoolsdefs.CteAppMenusActionCliHandler(CteAppMenusAction),
		},
	}
}
