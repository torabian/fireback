package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x MyNotificationsActionRequest) IsCli() bool {
	if x.CliCtx == nil {
		return false
	}
	v := reflect.ValueOf(x.CliCtx)
	switch v.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface, reflect.Func, reflect.Chan:
		return !v.IsNil()
	}
	return true
}

// MyNotificationsActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the MyNotificationsAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func MyNotificationsActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	return flags
}

// MyNotificationsActionCliHandler builds a full *cli.Command for the
// MyNotificationsAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a MyNotificationsActionRequest the same way
// MyNotificationsActionHandler (Gin) and MyNotificationsActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func MyNotificationsActionCliHandler(
	handler func(c MyNotificationsActionRequest) (*MyNotificationsActionResponse, error),
) *cli.Command {
	meta := MyNotificationsActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: MyNotificationsActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := MyNotificationsActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// MyNotificationsActionCli is a high-level convenience wrapper around
// MyNotificationsActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way MyNotificationsActionGin
// registers a route on a Gin engine.
func MyNotificationsActionCli(
	app *cli.Command,
	handler func(c MyNotificationsActionRequest) (*MyNotificationsActionResponse, error),
) {
	app.Commands = append(app.Commands, MyNotificationsActionCliHandler(handler))
}
