//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetPassportMethodUpdateActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func PassportMethodUpdateActionPathParameterFromCli(c *cli.Command) PassportMethodUpdateActionPathParameter {
	return PassportMethodUpdateActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x PassportMethodUpdateActionRequest) IsCli() bool {
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

// PassportMethodUpdateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PassportMethodUpdateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PassportMethodUpdateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPassportMethodOptionalDtoCliFlags(""))...)
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPassportMethodUpdateActionPathParameterCliFlags(""))...)
	return flags
}

// PassportMethodUpdateActionCliHandler builds a full *cli.Command for the
// PassportMethodUpdateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PassportMethodUpdateActionRequest the same way
// PassportMethodUpdateActionHandler (Gin) and PassportMethodUpdateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PassportMethodUpdateActionCliHandler(
	handler func(c PassportMethodUpdateActionRequest) (*PassportMethodUpdateActionResponse, error),
) *cli.Command {
	meta := PassportMethodUpdateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PassportMethodUpdateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PassportMethodUpdateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastPassportMethodOptionalDtoFromCli(c),
			Params:      PassportMethodUpdateActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PassportMethodUpdateActionCli is a high-level convenience wrapper around
// PassportMethodUpdateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PassportMethodUpdateActionGin
// registers a route on a Gin engine.
func PassportMethodUpdateActionCli(
	app *cli.Command,
	handler func(c PassportMethodUpdateActionRequest) (*PassportMethodUpdateActionResponse, error),
) {
	app.Commands = append(app.Commands, PassportMethodUpdateActionCliHandler(handler))
}
