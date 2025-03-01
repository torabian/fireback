package workspaces

func QueryMenusReact(query QueryDSL, chanStream chan *ReactiveSearchResultDto) {
	actionFnNavigate := "navigate"

	query.Query = "label %" + query.SearchPhrase + "%"
	items, _, _ := AppMenuActionQuery(query)

	for _, item := range items {
		if !item.ParentId.Valid {
			continue
		}

		uid := UUID()
		chanStream <- &ReactiveSearchResultDto{
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
func QueryRolesReact(query QueryDSL, chanStream chan *ReactiveSearchResultDto) {
	actionFnNavigate := "navigate"

	query.Query = "name %" + query.SearchPhrase + "%"
	items, _, _ := RoleActionQuery(query)

	roles := "roles"
	for _, item := range items {
		loc := "/role/" + item.UniqueId

		uid := UUID()

		chanStream <- &ReactiveSearchResultDto{
			Phrase:      item.Name,
			Description: item.Name,
			Group:       roles,
			ActionFn:    actionFnNavigate,
			UiLocation:  loc,
			UniqueId:    uid,
		}
	}

}
