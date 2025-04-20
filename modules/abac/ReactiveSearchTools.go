package abac

import "github.com/torabian/fireback/modules/fireback"

func QueryMenusReact(query fireback.QueryDSL, chanStream chan *fireback.ReactiveSearchResultDto) {
	actionFnNavigate := "navigate"

	query.Query = "label %" + query.SearchPhrase + "%"
	items, _, _ := AppMenuActions.Query(query)

	for _, item := range items {
		if !item.ParentId.Valid {
			continue
		}

		uid := fireback.UUID()
		chanStream <- &fireback.ReactiveSearchResultDto{
			Phrase:      item.Label,
			Description: item.Label,
			Icon:        item.Icon,
			Group:       item.ParentId.String,
			ActionFn:    actionFnNavigate,
			UiLocation:  item.Href,
			UniqueId:    uid,
		}
	}

}
func QueryRolesReact(query fireback.QueryDSL, chanStream chan *fireback.ReactiveSearchResultDto) {
	actionFnNavigate := "navigate"

	query.Query = "name %" + query.SearchPhrase + "%"
	items, _, _ := RoleActions.Query(query)

	roles := "roles"
	for _, item := range items {
		loc := "/role/" + item.UniqueId

		uid := fireback.UUID()

		chanStream <- &fireback.ReactiveSearchResultDto{
			Phrase:      item.Name,
			Description: item.Name,
			Group:       roles,
			ActionFn:    actionFnNavigate,
			UiLocation:  loc,
			UniqueId:    uid,
		}
	}

}
