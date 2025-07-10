package suggestion

import (
	"log"
	"strings"

	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	// Override the implementation with our actual code.
	QueryActionImp = QueryAction
}

func QueryAction(
	req *QueryActionReqDto,
	q fireback.QueryDSL) (*QueryActionResDto,
	*fireback.IError,
) {

	db := fireback.GetRef(q)

	// Generalize FTS config (could be passed in future, for now hardcoded)
	ftsConfig := FtsConfig{
		FtsTable: "content_virtual",
		Columns:  []string{"title", "excerpt", "content_type"},
	}

	// Use ItemsPerPage and StartIndex from req if set, else from q, else defaults
	itemsPerPage := req.ItemsPerPage
	if itemsPerPage <= 0 {
		itemsPerPage = q.ItemsPerPage
	}
	if itemsPerPage <= 0 {
		itemsPerPage = 10
	}
	startIndex := req.StartIndex
	if startIndex < 0 {
		startIndex = q.StartIndex
	}
	if startIndex < 0 {
		startIndex = 0
	}

	var results []*QueryResDtoItems

	affixes := []string{}
	if strings.TrimSpace(req.Phrase) != "" {
		affixes = append(affixes, "MATCH ? ")
	}
	affixes = append(affixes, "LIMIT ? OFFSET ?")

	values := []interface{}{}
	if strings.TrimSpace(req.Phrase) != "" {
		values = append(values, req.Phrase)
	}

	values = append(values, itemsPerPage, startIndex)

	query := "SELECT rowid, " + strings.Join(ftsConfig.Columns, ", ") + " FROM " + ftsConfig.FtsTable + " WHERE " + ftsConfig.FtsTable + " " + strings.Join(affixes, " ")
	err := db.Raw(query, values...).Scan(&results).Error
	if err != nil {
		log.Fatal("Query failed:", err)
	}

	return &QueryActionResDto{Items: results}, nil
}
