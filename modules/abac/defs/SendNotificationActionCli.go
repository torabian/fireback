//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x SendNotificationActionRequest) IsCli() bool {
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

// SendNotificationActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the SendNotificationAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func SendNotificationActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetSendNotificationActionReqCliFlags(""))...)
	return flags
}

// SendNotificationActionCliHandler builds a full *cli.Command for the
// SendNotificationAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a SendNotificationActionRequest the same way
// SendNotificationActionHandler (Gin) and SendNotificationActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func SendNotificationActionCliHandler(
	handler func(c SendNotificationActionRequest) (*SendNotificationActionResponse, error),
) *cli.Command {
	meta := SendNotificationActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: SendNotificationActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := SendNotificationActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastSendNotificationActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// SendNotificationActionCli is a high-level convenience wrapper around
// SendNotificationActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way SendNotificationActionGin
// registers a route on a Gin engine.
func SendNotificationActionCli(
	app *cli.Command,
	handler func(c SendNotificationActionRequest) (*SendNotificationActionResponse, error),
) {
	app.Commands = append(app.Commands, SendNotificationActionCliHandler(handler))
}
