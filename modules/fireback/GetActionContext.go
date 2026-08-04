package fireback

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
)

// ResolveActionContextFromGinContext resolves the QueryDSL for a plain
// *gin.Context, enforcing the security model if one is given. This holds the
// gin-specific logic shared by ResolveActionContext so it isn't duplicated
// wherever only a *gin.Context (and no emigo.EmiRequestContexts) is available.
func ResolveActionContextFromGinContext(ginCtx *gin.Context, securityModel *SecurityModel) (*QueryDSL, error) {
	qsdl := ExtractQueryDslFromGinContext(ginCtx)

	// Only handles the gin security, and this is a problem needs to handle cli as well
	// perfectly
	if securityModel != nil && !IsNilish(ginCtx) {
		// Fireback no longer plans to check security anymore. Do it manually,
		// on every action. This is fr transparency.
		if !AuthorizeRequest(securityModel, ginCtx) {
			return nil, errors.New("Authorization general failed")
		}

		// Important because now we have more details of security
		qsdl = ExtractQueryDslFromGinContext(ginCtx)
	}

	return &qsdl, nil
}

func ResolveActionContext(request emigo.EmiRequestContexts, securityModel *SecurityModel) (*QueryDSL, error) {
	GinCtx := request.GetGinCtx()
	CliCtx := request.GetCliCtx()

	var qsdl QueryDSL

	if value, ok := GinCtx.(*gin.Context); ok {
		resolved, err := ResolveActionContextFromGinContext(value, securityModel)
		if err != nil {
			return nil, err
		}
		qsdl = *resolved
	}

	if !IsNilish(CliCtx) {
		// Fireback no longer plans to check security anymore. Do it manually,
		// on every action. This is fr transparency.
		if value, ok := CliCtx.(*cli.Command); ok {
			qsdl = CommonCliQueryDSLBuilderAuthorize(value, securityModel)
		}
	}

	return &qsdl, nil
}
