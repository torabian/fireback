package suggestion

import (
	"fmt"

	"gorm.io/gorm"
)

// FtsConfig holds configuration for FTS5 table and triggers
type FtsConfig struct {
	BaseTable string   // e.g., "content_entities"
	FtsTable  string   // e.g., "content_virtual"
	Columns   []string // e.g., []string{"title", "body"}
	IdColumn  string   // e.g., "id"
}

func CreateFts5TableAndTriggers(db *gorm.DB, config FtsConfig) error {
	// Build FTS5 table creation SQL
	ftsCols := ""
	for i, col := range config.Columns {
		if i > 0 {
			ftsCols += ", "
		}
		ftsCols += col
	}
	createTableSQL := "CREATE VIRTUAL TABLE " + config.FtsTable + " USING fts5(" + ftsCols + ")"

	// Drop the virtual table if it exists
	if err := db.Exec("DROP TABLE IF EXISTS " + config.FtsTable).Error; err != nil {
		return err
	}
	// Recreate the virtual table
	if err := db.Exec(createTableSQL).Error; err != nil {
		return err
	}

	// Build triggers for each column
	insertCols := ""
	insertVals := ""
	updateSet := ""
	for i, col := range config.Columns {
		if i > 0 {
			insertCols += ", "
			insertVals += ", "
			updateSet += ", "
		}
		insertCols += col
		insertVals += "new." + col
		updateSet += col + " = new." + col
	}

	triggerSQL := []string{
		// INSERT trigger
		"CREATE TRIGGER IF NOT EXISTS " + config.BaseTable + "_ai AFTER INSERT ON " + config.BaseTable + " BEGIN\n" +
			"INSERT INTO " + config.FtsTable + "(rowid, " + insertCols + ") VALUES (new." + config.IdColumn + ", " + insertVals + ");\nEND;",
		// UPDATE trigger
		"CREATE TRIGGER IF NOT EXISTS " + config.BaseTable + "_au AFTER UPDATE ON " + config.BaseTable + " BEGIN\n" +
			"UPDATE " + config.FtsTable + " SET " + updateSet + " WHERE rowid = new." + config.IdColumn + ";\nEND;",
		// DELETE trigger
		"CREATE TRIGGER IF NOT EXISTS " + config.BaseTable + "_ad AFTER DELETE ON " + config.BaseTable + " BEGIN\n" +
			"DELETE FROM " + config.FtsTable + " WHERE rowid = old." + config.IdColumn + ";\nEND;",
	}

	for _, sql := range triggerSQL {
		fmt.Println(sql)
		if err := db.Exec(sql).Error; err != nil {
			return err
		}
	}
	return nil
}
