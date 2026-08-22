package wallet

import (
	"net/http"

	"github.com/torabian/fireback/modules/fireback"
)

// emiIErrorPayload turns a *fireback.IError into the (status, payload) pair to set
// directly on an Emi-generated action's Response struct. See
// modules/category/ActionErrors.go for the full rationale - duplicated per-package since
// it isn't exported from fireback itself.
func emiIErrorPayload(ierr *fireback.IError, query *fireback.QueryDSL) (int, any) {
	status := int(ierr.HttpCode)
	if status == 0 {
		status = http.StatusInternalServerError
	}
	return status, map[string]any{"error": ierr.ToPublicEndUser(query)}
}
