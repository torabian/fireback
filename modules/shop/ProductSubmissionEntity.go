package shop

import (
	"strings"

	"github.com/torabian/fireback/modules/workspaces"
)

func ProductSubmissionActionCreate(
	dto *ProductSubmissionEntity, query workspaces.QueryDSL,
) (*ProductSubmissionEntity, *workspaces.IError) {
	FormatComputedValue(dto, query)
	return ProductSubmissionActionCreateFn(dto, query)
}
func ProductSubmissionActionUpdate(
	query workspaces.QueryDSL,
	fields *ProductSubmissionEntity,
) (*ProductSubmissionEntity, *workspaces.IError) {
	FormatComputedValue(fields, query)
	return ProductSubmissionActionUpdateFn(query, fields)
}

func FormatComputedValue(fields *ProductSubmissionEntity, query workspaces.QueryDSL) string {
	if fields.Price == nil || len(fields.Price.Variations) == 0 {
		return ""
	}

	value := []string{}
	for _, variation := range fields.Price.Variations {
		if variation.Amount == nil || variation.CurrencyId == nil {
			continue
		}
		// ac := accounting.DefaultAccounting(strings.ToUpper(*variation.CurrencyId), 2)
		// value = append(value, ac.FormatMoney(ac))

	}

	return strings.Join(value, ", ")
}
