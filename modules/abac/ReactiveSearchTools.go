package abac

import (
	"github.com/torabian/fireback/modules/abac/interfacetools"
	"github.com/torabian/fireback/modules/fireback"
)

func QueryMenusReact(query fireback.QueryDSL, chanStream chan *fireback.ReactiveSearchResultDto) {
	actionFnNavigate := "navigate"

	query.Query = "label %" + query.SearchPhrase + "%"
	items, _, _ := interfacetools.AppMenuActions.Query(query)

	for _, item := range items {
		if !item.ParentId.IsSet() {
			continue
		}

		uid := fireback.UUID()
		// item.Label is a complexes.TString (locale -> text map) now; ReactiveSearchResultDto
		// needs a single display string, so resolve it the same way TString.String() always
		// has (DefaultLocale, falling back to whatever locale is present).
		label := item.Label.String()
		chanStream <- &fireback.ReactiveSearchResultDto{
			Phrase:      label,
			Description: label,
			Icon:        item.Icon,
			Group:       item.ParentId.OrDefault(""),
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
		loc := "/selfservice/role/" + item.UniqueId

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
