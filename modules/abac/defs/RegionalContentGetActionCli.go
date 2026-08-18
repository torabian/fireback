//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetRegionalContentGetActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func RegionalContentGetActionPathParameterFromCli(c *cli.Command) RegionalContentGetActionPathParameter {
	return RegionalContentGetActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x RegionalContentGetActionRequest) IsCli() bool {
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

// RegionalContentGetActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the RegionalContentGetAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func RegionalContentGetActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetRegionalContentGetActionPathParameterCliFlags(""))...)
	return flags
}

// RegionalContentGetActionCliHandler builds a full *cli.Command for the
// RegionalContentGetAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a RegionalContentGetActionRequest the same way
// RegionalContentGetActionHandler (Gin) and RegionalContentGetActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func RegionalContentGetActionCliHandler(
	handler func(c RegionalContentGetActionRequest) (*RegionalContentGetActionResponse, error),
) *cli.Command {
	meta := RegionalContentGetActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: RegionalContentGetActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := RegionalContentGetActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Params:      RegionalContentGetActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// RegionalContentGetActionCli is a high-level convenience wrapper around
// RegionalContentGetActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way RegionalContentGetActionGin
// registers a route on a Gin engine.
func RegionalContentGetActionCli(
	app *cli.Command,
	handler func(c RegionalContentGetActionRequest) (*RegionalContentGetActionResponse, error),
) {
	app.Commands = append(app.Commands, RegionalContentGetActionCliHandler(handler))
}
