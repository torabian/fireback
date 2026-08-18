//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x SendPassportResetEmailActionRequest) IsCli() bool {
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

// SendPassportResetEmailActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the SendPassportResetEmailAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func SendPassportResetEmailActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetSendPassportResetEmailActionReqCliFlags(""))...)
	return flags
}

// SendPassportResetEmailActionCliHandler builds a full *cli.Command for the
// SendPassportResetEmailAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a SendPassportResetEmailActionRequest the same way
// SendPassportResetEmailActionHandler (Gin) and SendPassportResetEmailActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func SendPassportResetEmailActionCliHandler(
	handler func(c SendPassportResetEmailActionRequest) (*SendPassportResetEmailActionResponse, error),
) *cli.Command {
	meta := SendPassportResetEmailActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: SendPassportResetEmailActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := SendPassportResetEmailActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastSendPassportResetEmailActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// SendPassportResetEmailActionCli is a high-level convenience wrapper around
// SendPassportResetEmailActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way SendPassportResetEmailActionGin
// registers a route on a Gin engine.
func SendPassportResetEmailActionCli(
	app *cli.Command,
	handler func(c SendPassportResetEmailActionRequest) (*SendPassportResetEmailActionResponse, error),
) {
	app.Commands = append(app.Commands, SendPassportResetEmailActionCliHandler(handler))
}
