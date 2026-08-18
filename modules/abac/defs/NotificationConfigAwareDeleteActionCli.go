//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x NotificationConfigAwareDeleteActionRequest) IsCli() bool {
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

// NotificationConfigAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the NotificationConfigAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func NotificationConfigAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetNotificationConfigAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// NotificationConfigAwareDeleteActionCliHandler builds a full *cli.Command for the
// NotificationConfigAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a NotificationConfigAwareDeleteActionRequest the same way
// NotificationConfigAwareDeleteActionHandler (Gin) and NotificationConfigAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func NotificationConfigAwareDeleteActionCliHandler(
	handler func(c NotificationConfigAwareDeleteActionRequest) (*NotificationConfigAwareDeleteActionResponse, error),
) *cli.Command {
	meta := NotificationConfigAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: NotificationConfigAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := NotificationConfigAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastNotificationConfigAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// NotificationConfigAwareDeleteActionCli is a high-level convenience wrapper around
// NotificationConfigAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way NotificationConfigAwareDeleteActionGin
// registers a route on a Gin engine.
func NotificationConfigAwareDeleteActionCli(
	app *cli.Command,
	handler func(c NotificationConfigAwareDeleteActionRequest) (*NotificationConfigAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, NotificationConfigAwareDeleteActionCliHandler(handler))
}
