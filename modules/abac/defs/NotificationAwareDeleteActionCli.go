//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x NotificationAwareDeleteActionRequest) IsCli() bool {
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

// NotificationAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the NotificationAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func NotificationAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetNotificationAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// NotificationAwareDeleteActionCliHandler builds a full *cli.Command for the
// NotificationAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a NotificationAwareDeleteActionRequest the same way
// NotificationAwareDeleteActionHandler (Gin) and NotificationAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func NotificationAwareDeleteActionCliHandler(
	handler func(c NotificationAwareDeleteActionRequest) (*NotificationAwareDeleteActionResponse, error),
) *cli.Command {
	meta := NotificationAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: NotificationAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := NotificationAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastNotificationAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// NotificationAwareDeleteActionCli is a high-level convenience wrapper around
// NotificationAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way NotificationAwareDeleteActionGin
// registers a route on a Gin engine.
func NotificationAwareDeleteActionCli(
	app *cli.Command,
	handler func(c NotificationAwareDeleteActionRequest) (*NotificationAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, NotificationAwareDeleteActionCliHandler(handler))
}
