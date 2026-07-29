package fireback

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
)

func ResolveActionContext(request emigo.EmiRequestContexts, securityModel *SecurityModel) (*QueryDSL, error) {
	GinCtx := request.GetGinCtx()
	CliCtx := request.GetCliCtx()

	var qsdl QueryDSL

	if value, ok := GinCtx.(*gin.Context); ok {
		qsdl = ExtractQueryDslFromGinContext(value)
	}

	// Only handles the gin security, and this is a problem needs to handle cli as well
	// perfectly
	if securityModel != nil && !IsNilish(GinCtx) {
		// Fireback no longer plans to check security anymore. Do it manually,
		// on every action. This is fr transparency.
		if !AuthorizeRequest(securityModel, GinCtx.(*gin.Context)) {
			return nil, errors.New("Authorization general failed")
		}

		// Important because now we have more details of security
		if value, ok := GinCtx.(*gin.Context); ok {
			qsdl = ExtractQueryDslFromGinContext(value)
		}
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
