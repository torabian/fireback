package abac

import (
	"github.com/torabian/fireback/modules/fireback"
)

func init() {

	AppMenuActions.CteQuery = func(query fireback.QueryDSL) ([]*AppMenuEntity, *fireback.QueryResultMeta, *fireback.IError) {
		result, qrm, err := AppMenuActionCteQueryFn(query)
		return result, qrm, err
	}
}
