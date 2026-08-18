//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetRegionalContentUpdateActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func RegionalContentUpdateActionPathParameterFromCli(c *cli.Command) RegionalContentUpdateActionPathParameter {
	return RegionalContentUpdateActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x RegionalContentUpdateActionRequest) IsCli() bool {
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

// RegionalContentUpdateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the RegionalContentUpdateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func RegionalContentUpdateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetRegionalContentOptionalDtoCliFlags(""))...)
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetRegionalContentUpdateActionPathParameterCliFlags(""))...)
	return flags
}

// RegionalContentUpdateActionCliHandler builds a full *cli.Command for the
// RegionalContentUpdateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a RegionalContentUpdateActionRequest the same way
// RegionalContentUpdateActionHandler (Gin) and RegionalContentUpdateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func RegionalContentUpdateActionCliHandler(
	handler func(c RegionalContentUpdateActionRequest) (*RegionalContentUpdateActionResponse, error),
) *cli.Command {
	meta := RegionalContentUpdateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: RegionalContentUpdateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := RegionalContentUpdateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastRegionalContentOptionalDtoFromCli(c),
			Params:      RegionalContentUpdateActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// RegionalContentUpdateActionCli is a high-level convenience wrapper around
// RegionalContentUpdateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way RegionalContentUpdateActionGin
// registers a route on a Gin engine.
func RegionalContentUpdateActionCli(
	app *cli.Command,
	handler func(c RegionalContentUpdateActionRequest) (*RegionalContentUpdateActionResponse, error),
) {
	app.Commands = append(app.Commands, RegionalContentUpdateActionCliHandler(handler))
}
