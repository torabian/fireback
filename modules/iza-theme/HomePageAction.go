package izaTheme

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	HomePageImpl = func(c HomePageActionRequest, query fireback.QueryDSL) (*HomePageActionResponse, error) {

		fireback.RenderPage(fireback.ResolveScreens(embeddedScreens), c.GinCtx, "index.html", nil)

		return nil, nil
	}
}
