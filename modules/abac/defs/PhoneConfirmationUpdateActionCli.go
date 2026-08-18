//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetPhoneConfirmationUpdateActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func PhoneConfirmationUpdateActionPathParameterFromCli(c *cli.Command) PhoneConfirmationUpdateActionPathParameter {
	return PhoneConfirmationUpdateActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x PhoneConfirmationUpdateActionRequest) IsCli() bool {
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

// PhoneConfirmationUpdateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PhoneConfirmationUpdateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PhoneConfirmationUpdateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPhoneConfirmationOptionalDtoCliFlags(""))...)
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPhoneConfirmationUpdateActionPathParameterCliFlags(""))...)
	return flags
}

// PhoneConfirmationUpdateActionCliHandler builds a full *cli.Command for the
// PhoneConfirmationUpdateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PhoneConfirmationUpdateActionRequest the same way
// PhoneConfirmationUpdateActionHandler (Gin) and PhoneConfirmationUpdateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PhoneConfirmationUpdateActionCliHandler(
	handler func(c PhoneConfirmationUpdateActionRequest) (*PhoneConfirmationUpdateActionResponse, error),
) *cli.Command {
	meta := PhoneConfirmationUpdateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PhoneConfirmationUpdateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PhoneConfirmationUpdateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastPhoneConfirmationOptionalDtoFromCli(c),
			Params:      PhoneConfirmationUpdateActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PhoneConfirmationUpdateActionCli is a high-level convenience wrapper around
// PhoneConfirmationUpdateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PhoneConfirmationUpdateActionGin
// registers a route on a Gin engine.
func PhoneConfirmationUpdateActionCli(
	app *cli.Command,
	handler func(c PhoneConfirmationUpdateActionRequest) (*PhoneConfirmationUpdateActionResponse, error),
) {
	app.Commands = append(app.Commands, PhoneConfirmationUpdateActionCliHandler(handler))
}
