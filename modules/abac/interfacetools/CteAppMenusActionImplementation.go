package interfacetools

import (
	interfacetoolsdefs "github.com/torabian/fireback/modules/abac/interfacetools/defs"
	"github.com/torabian/fireback/modules/fireback"
)

// CteAppMenusAction restores the old Module3-generated GET /cte-app-menus route lost
// during the migration to Emi - AppMenu rows assembled into a tree (each item's
// children nested under it) instead of Browse's flat list. See
// AppMenuActions.CteQuery / AppMenuTreeNode for the actual recursive query + tree
// assembly logic.
func CteAppMenusAction(c interfacetoolsdefs.CteAppMenusActionRequest) (*interfacetoolsdefs.CteAppMenusActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{})
	if err != nil {
		return nil, err
	}
	tree, qrm, err2 := AppMenuActions.CteQuery(*query)
	if err2 != nil {
		return nil, err2
	}
	return &interfacetoolsdefs.CteAppMenusActionResponse{Payload: fireback.GResponseQuery(tree, qrm, query)}, nil
}
