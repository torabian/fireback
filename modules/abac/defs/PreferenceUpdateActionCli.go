//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetPreferenceUpdateActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func PreferenceUpdateActionPathParameterFromCli(c *cli.Command) PreferenceUpdateActionPathParameter {
	return PreferenceUpdateActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x PreferenceUpdateActionRequest) IsCli() bool {
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

// PreferenceUpdateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PreferenceUpdateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PreferenceUpdateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPreferenceOptionalDtoCliFlags(""))...)
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPreferenceUpdateActionPathParameterCliFlags(""))...)
	return flags
}

// PreferenceUpdateActionCliHandler builds a full *cli.Command for the
// PreferenceUpdateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PreferenceUpdateActionRequest the same way
// PreferenceUpdateActionHandler (Gin) and PreferenceUpdateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PreferenceUpdateActionCliHandler(
	handler func(c PreferenceUpdateActionRequest) (*PreferenceUpdateActionResponse, error),
) *cli.Command {
	meta := PreferenceUpdateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PreferenceUpdateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PreferenceUpdateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastPreferenceOptionalDtoFromCli(c),
			Params:      PreferenceUpdateActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PreferenceUpdateActionCli is a high-level convenience wrapper around
// PreferenceUpdateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PreferenceUpdateActionGin
// registers a route on a Gin engine.
func PreferenceUpdateActionCli(
	app *cli.Command,
	handler func(c PreferenceUpdateActionRequest) (*PreferenceUpdateActionResponse, error),
) {
	app.Commands = append(app.Commands, PreferenceUpdateActionCliHandler(handler))
}
