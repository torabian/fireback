package home

import (
	"embed"
	"io/fs"

	"github.com/torabian/fireback/modules/fireback"
)

var ScreensFs fs.FS // use interface, not embed.FS

//go:embed all:screens
var embeddedScreens embed.FS

func init() {
	HomePageActionImp = HomePageAction
}

func HomePageAction(
	q fireback.QueryDSL) (*fireback.XHtml,
	*fireback.IError,
) {

	return &fireback.XHtml{
		Params:       nil,
		TemplateName: "home.html",
		ScreensFs:    fireback.ResolveScreens(embeddedScreens),
	}, nil
}
