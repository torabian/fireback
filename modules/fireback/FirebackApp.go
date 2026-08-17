package fireback

import (
	"context"
	"log"
	"os"

	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"
	"go.uber.org/zap"
)

var SERVER_INSTANCE string = UUID_Long()

type GooseZapLogger struct {
	Logger *zap.Logger
}

func (l GooseZapLogger) Printf(format string, v ...interface{}) {
	l.Logger.Sugar().Infof(format, v...)
}

func (l GooseZapLogger) Fatalf(format string, v ...interface{}) {
	l.Logger.Sugar().Fatalf(format, v...)
}

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
