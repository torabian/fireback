//go:build !wasm

package abac

import (
	"log"
	"strings"

	"github.com/torabian/fireback/modules/abac/interfacetools"
	"github.com/torabian/fireback/modules/fireback"
	reactivesearchdefs "github.com/torabian/fireback/modules/reactivesearch/defs"
)

func QueryMenusReact(query fireback.QueryDSL, chanStream chan *reactivesearchdefs.ReactiveSearchResultDto) {
	actionFnNavigate := "navigate"

	// Bug fix: this used to set query.Query = "label %<phrase>%", which the generic
	// goven-based Query() (see fireback.CrudCoreActions.go's queryAdaptor.Parse) turns
	// into a raw `label LIKE ?` SQL clause. AppMenuEntity.Label is a complexes.TString
	// (a locale -> text map), stored as a jsonb/JSON column, not plain text - Postgres
	// rejects LIKE against jsonb outright ("operator does not exist: jsonb ~~
	// unknown"), and the error was silently discarded here (the old `_, _ :=`), so menu
	// search always returned zero results, for every phrase, ever since Label stopped
	// being a plain string column (see Menu.go's own TStringFrom labels). Fetch every
	// menu item the caller can see instead (a generous ItemsPerPage - AppMenu items
	// number in the dozens, never the default page size of 10) and match the phrase in
	// Go against the same resolved display string item.Label.String() already used
	// below for the result itself - portable across every DB vendor this app supports,
	// and matches what's actually shown to the user, instead of just one locale's raw
	// JSON key.
	query.Query = ""
	if query.ItemsPerPage < 500 {
		query.ItemsPerPage = 500
	}

	phrase := strings.ToLower(query.SearchPhrase)

	items, _, err := interfacetools.AppMenuActions.Query(query)
	if err != nil {
		log.Println("reactivesearch: querying app menu items failed:", err)
		return
	}

	for _, item := range items {
		if !item.ParentId.IsSet() {
			continue
		}

		// item.Label is a complexes.TString (locale -> text map) now; ReactiveSearchResultDto
		// needs a single display string, so resolve it the same way TString.String() always
		// has (DefaultLocale, falling back to whatever locale is present).
		label := item.Label.String()

		if phrase != "" && !strings.Contains(strings.ToLower(label), phrase) {
			continue
		}

		uid := fireback.UUID()
		chanStream <- &reactivesearchdefs.ReactiveSearchResultDto{
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
func QueryRolesReact(query fireback.QueryDSL, chanStream chan *reactivesearchdefs.ReactiveSearchResultDto) {
	actionFnNavigate := "navigate"

	query.Query = "name %" + query.SearchPhrase + "%"
	items, _, _ := RoleActions.Query(query)

	roles := "roles"
	for _, item := range items {
		loc := "/selfservice/role/" + item.UniqueId

		uid := fireback.UUID()

		chanStream <- &reactivesearchdefs.ReactiveSearchResultDto{
			Phrase:      item.Name,
			Description: item.Name,
			Group:       roles,
			ActionFn:    actionFnNavigate,
			UiLocation:  loc,
			UniqueId:    uid,
		}
	}

}
