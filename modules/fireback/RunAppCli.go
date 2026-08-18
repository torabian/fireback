// RunApp is CLI dispatch (app.Run(context.Background(), os.Args), GetCliCommands -
// CliActions.go, also !wasm) - the wasm binary never calls it, wiring its own
// request handling directly instead (see cmd/fireback-wasm/main.go). Split out of
// FirebackApp.go, which also holds SERVER_INSTANCE/GooseZapLogger - both needed
// under wasm too, unlike this.
//go:build !wasm

package fireback

import (
	"context"
	"log"
	"os"

	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"
)

func RunApp(xapp *application.Application) {

	app := &cli.Command{
		EnableShellCompletion: true,
		Name:                  xapp.Title,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "al",
				Usage: "Set's the language of the query, equal to accept-language header in http requests",
				Value: "en-us",
			},
		},
		Commands: GetCliCommands(xapp),
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
