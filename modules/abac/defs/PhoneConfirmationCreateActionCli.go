//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x PhoneConfirmationCreateActionRequest) IsCli() bool {
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

// PhoneConfirmationCreateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PhoneConfirmationCreateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PhoneConfirmationCreateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPhoneConfirmationDtoCliFlags(""))...)
	return flags
}

// PhoneConfirmationCreateActionCliHandler builds a full *cli.Command for the
// PhoneConfirmationCreateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PhoneConfirmationCreateActionRequest the same way
// PhoneConfirmationCreateActionHandler (Gin) and PhoneConfirmationCreateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PhoneConfirmationCreateActionCliHandler(
	handler func(c PhoneConfirmationCreateActionRequest) (*PhoneConfirmationCreateActionResponse, error),
) *cli.Command {
	meta := PhoneConfirmationCreateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PhoneConfirmationCreateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PhoneConfirmationCreateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastPhoneConfirmationDtoFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PhoneConfirmationCreateActionCli is a high-level convenience wrapper around
// PhoneConfirmationCreateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PhoneConfirmationCreateActionGin
// registers a route on a Gin engine.
func PhoneConfirmationCreateActionCli(
	app *cli.Command,
	handler func(c PhoneConfirmationCreateActionRequest) (*PhoneConfirmationCreateActionResponse, error),
) {
	app.Commands = append(app.Commands, PhoneConfirmationCreateActionCliHandler(handler))
}
