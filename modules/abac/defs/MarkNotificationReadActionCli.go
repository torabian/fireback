//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x MarkNotificationReadActionRequest) IsCli() bool {
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

// MarkNotificationReadActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the MarkNotificationReadAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func MarkNotificationReadActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetMarkNotificationReadActionReqCliFlags(""))...)
	return flags
}

// MarkNotificationReadActionCliHandler builds a full *cli.Command for the
// MarkNotificationReadAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a MarkNotificationReadActionRequest the same way
// MarkNotificationReadActionHandler (Gin) and MarkNotificationReadActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func MarkNotificationReadActionCliHandler(
	handler func(c MarkNotificationReadActionRequest) (*MarkNotificationReadActionResponse, error),
) *cli.Command {
	meta := MarkNotificationReadActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: MarkNotificationReadActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := MarkNotificationReadActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastMarkNotificationReadActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// MarkNotificationReadActionCli is a high-level convenience wrapper around
// MarkNotificationReadActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way MarkNotificationReadActionGin
// registers a route on a Gin engine.
func MarkNotificationReadActionCli(
	app *cli.Command,
	handler func(c MarkNotificationReadActionRequest) (*MarkNotificationReadActionResponse, error),
) {
	app.Commands = append(app.Commands, MarkNotificationReadActionCliHandler(handler))
}
