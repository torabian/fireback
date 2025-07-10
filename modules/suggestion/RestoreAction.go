package suggestion

import (
	"log"

	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	// Override the implementation with our actual code.
	RestoreActionImp = RestoreAction
}

// RestoreFtsAction uses the package-level config for flexibility
func RestoreAction(
	q fireback.QueryDSL,
) (string, *fireback.IError) {
	db := fireback.GetRef(q)

	// Default config, can be set by SetFtsConfig
	var ftsConfig = FtsConfig{
		BaseTable: "content_entities",
		FtsTable:  "content_virtual",
		Columns:   []string{"title", "excerpt", "content_type"},
		IdColumn:  "id",
	}

	if err := CreateFts5TableAndTriggers(db, ftsConfig); err != nil {
		log.Fatal("Failed to restore FTS5:", err)
	}

	return "OKAY", nil
}
