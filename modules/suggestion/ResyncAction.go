package suggestion

import (
	"fmt"

	"github.com/schollz/progressbar/v3"
	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	// Override the implementation with our actual code.
	ResyncActionImp = ResyncAction
}

func ResyncAction(
	q fireback.QueryDSL,
) (string, *fireback.IError) {
	q.ItemsPerPage = 10
	q.StartIndex = 0
	stream, meta, err := ContentEntityStream(q)
	if err != nil {
		return "", fireback.GormErrorToIError(err)
	}

	bar := progressbar.Default(int64(meta.TotalAvailableItems))
	inserted := 0
	db := fireback.GetRef(q)

	// Clean existing virtual table entries
	if err := db.Exec("DELETE FROM content_virtual").Error; err != nil {
		return "", fireback.GormErrorToIError(err)
	}

	for batch := range stream {
		for _, item := range batch {
			query := "INSERT INTO content_virtual(rowid, title, excerpt, content_type) VALUES (?, ?, ? , ?)"
			if err := db.Exec(query, item.ID, item.Title, item.Excerpt, item.ContentType).Error; err != nil {
				fmt.Println("Insertion failed for ID:", item.ID, err.Error())
				continue
			}
			inserted++
		}

		bar.Add(inserted)
	}

	return fmt.Sprintf("Resynced %d content items into FTS table (content_virtual). Total available: %d", inserted, meta.TotalAvailableItems), nil
}
