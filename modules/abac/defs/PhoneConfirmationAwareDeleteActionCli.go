//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x PhoneConfirmationAwareDeleteActionRequest) IsCli() bool {
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

// PhoneConfirmationAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PhoneConfirmationAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PhoneConfirmationAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPhoneConfirmationAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// PhoneConfirmationAwareDeleteActionCliHandler builds a full *cli.Command for the
// PhoneConfirmationAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PhoneConfirmationAwareDeleteActionRequest the same way
// PhoneConfirmationAwareDeleteActionHandler (Gin) and PhoneConfirmationAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PhoneConfirmationAwareDeleteActionCliHandler(
	handler func(c PhoneConfirmationAwareDeleteActionRequest) (*PhoneConfirmationAwareDeleteActionResponse, error),
) *cli.Command {
	meta := PhoneConfirmationAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PhoneConfirmationAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PhoneConfirmationAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastPhoneConfirmationAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PhoneConfirmationAwareDeleteActionCli is a high-level convenience wrapper around
// PhoneConfirmationAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PhoneConfirmationAwareDeleteActionGin
// registers a route on a Gin engine.
func PhoneConfirmationAwareDeleteActionCli(
	app *cli.Command,
	handler func(c PhoneConfirmationAwareDeleteActionRequest) (*PhoneConfirmationAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, PhoneConfirmationAwareDeleteActionCliHandler(handler))
}
