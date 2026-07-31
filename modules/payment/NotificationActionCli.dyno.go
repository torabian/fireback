//go:build !wasm

package payment

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x NotificationActionRequest) IsCli() bool {
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

// NotificationActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the NotificationAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func NotificationActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetNotificationActionReqCliFlags(""))...)
	return flags
}

// NotificationActionCliHandler builds a full *cli.Command for the
// NotificationAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a NotificationActionRequest the same way
// NotificationActionHandler (Gin) and NotificationActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func NotificationActionCliHandler(
	handler func(c NotificationActionRequest) (*NotificationActionResponse, error),
) *cli.Command {
	meta := NotificationActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: NotificationActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := NotificationActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastNotificationActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// NotificationActionCli is a high-level convenience wrapper around
// NotificationActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way NotificationActionGin
// registers a route on a Gin engine.
func NotificationActionCli(
	app *cli.Command,
	handler func(c NotificationActionRequest) (*NotificationActionResponse, error),
) {
	app.Commands = append(app.Commands, NotificationActionCliHandler(handler))
}
