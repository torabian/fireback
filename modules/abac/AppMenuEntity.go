package abac

import (
	"github.com/stretchr/testify/assert"
	"github.com/torabian/fireback/modules/workspaces"
)

var AppMenuTests = []workspaces.Test{
	{
		Name: "Creation of app menu polyglot, and removing them",
		Function: func(t *workspaces.TestContext) error {

			// 1. Create the app menu first
			id := workspaces.UUID()

			label := "Warsaw"

			menu, err := AppMenuActions.Create(&AppMenuEntity{
				UniqueId: id,
				Label:    label,
				Translations: []*AppMenuEntityPolyglot{
					{
						LanguageId: "fa",
						Label:      "Warshoo",
						LinkerId:   id,
					},
					{
						LanguageId: "pl",
						Label:      "Warszawa",
						LinkerId:   id,
					},
				},
			}, t.F)

			assert.Nil(t, err, "There should be no error while creating the menu item")
			assert.Equal(t, 2, len(menu.Translations), "Two polyglot items for fa and pl must be present")
			assert.Equal(t, "Warszawa", menu.Translations[1].Label, "Polish version has correct value")
			assert.Equal(t, "Warshoo", menu.Translations[0].Label, "Farsi version has correct value")

			// 2. Try to update the menu with polyglot

			newLabel := "This is updated english label"
			menuUpdated1, err2 := AppMenuActions.Update(t.F, &AppMenuEntity{
				UniqueId: menu.UniqueId,
				Label:    newLabel,
			})

			assert.Nil(t, err2, "There should be no error while updating the menu item")
			assert.Equal(t, 3, len(menuUpdated1.Translations), "There has to be now 3 items")

			affected, err3 := AppMenuActions.Remove(workspaces.QueryDSL{
				Query: "unique_id = " + menuUpdated1.UniqueId,
			})

			assert.Nil(t, err3, "There should be no issue while deleting app menu")
			assert.Equal(t, int64(1), affected, "Only one row has to be deleted")

			return nil
		},
	},
}

func init() {

	AppMenuActions.CteQuery = func(query workspaces.QueryDSL) ([]*AppMenuEntity, *workspaces.QueryResultMeta, error) {
		result, qrm, err := AppMenuActionCteQueryFn(query)
		return result, qrm, err
	}
}
