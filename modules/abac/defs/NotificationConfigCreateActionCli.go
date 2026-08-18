//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x NotificationConfigCreateActionRequest) IsCli() bool {
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

// NotificationConfigCreateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the NotificationConfigCreateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func NotificationConfigCreateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetNotificationConfigDtoCliFlags(""))...)
	return flags
}

// NotificationConfigCreateActionCliHandler builds a full *cli.Command for the
// NotificationConfigCreateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a NotificationConfigCreateActionRequest the same way
// NotificationConfigCreateActionHandler (Gin) and NotificationConfigCreateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func NotificationConfigCreateActionCliHandler(
	handler func(c NotificationConfigCreateActionRequest) (*NotificationConfigCreateActionResponse, error),
) *cli.Command {
	meta := NotificationConfigCreateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: NotificationConfigCreateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := NotificationConfigCreateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastNotificationConfigDtoFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// NotificationConfigCreateActionCli is a high-level convenience wrapper around
// NotificationConfigCreateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way NotificationConfigCreateActionGin
// registers a route on a Gin engine.
func NotificationConfigCreateActionCli(
	app *cli.Command,
	handler func(c NotificationConfigCreateActionRequest) (*NotificationConfigCreateActionResponse, error),
) {
	app.Commands = append(app.Commands, NotificationConfigCreateActionCliHandler(handler))
}
