package izaTheme

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	ProductSingleScreenImpl = func(c ProductSingleScreenActionRequest, query fireback.QueryDSL) (*ProductSingleScreenActionResponse, error) {
		fireback.RenderPage(fireback.ResolveScreens(embeddedScreens), c.GinCtx, "product-info.html", nil)

		return nil, nil
	}
}
