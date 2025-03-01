package workspaces

import (
	"github.com/stretchr/testify/assert"
)

func AppMenuActionCreate(
	dto *AppMenuEntity, query QueryDSL,
) (*AppMenuEntity, *IError) {
	return AppMenuActionCreateFn(dto, query)
}

func AppMenuActionUpdate(
	query QueryDSL,
	fields *AppMenuEntity,
) (*AppMenuEntity, *IError) {
	return AppMenuActionUpdateFn(query, fields)
}

var AppMenuTests = []Test{
	{
		Name: "Creation of app menu polyglot, and removing them",
		Function: func(t *TestContext) error {

			// 1. Create the app menu first
			id := UUID()

			label := "Warsaw"

			menu, err := AppMenuActionCreate(&AppMenuEntity{
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
			menuUpdated1, err2 := AppMenuActionUpdate(t.F, &AppMenuEntity{
				UniqueId: menu.UniqueId,
				Label:    newLabel,
			})

			assert.Nil(t, err2, "There should be no error while updating the menu item")
			assert.Equal(t, 3, len(menuUpdated1.Translations), "There has to be now 3 items")

			affected, err3 := AppMenuActionRemove(QueryDSL{
				Query: "unique_id = " + menuUpdated1.UniqueId,
			})

			assert.Nil(t, err3, "There should be no issue while deleting app menu")
			assert.Equal(t, int64(1), affected, "Only one row has to be deleted")

			return nil
		},
	},
}
