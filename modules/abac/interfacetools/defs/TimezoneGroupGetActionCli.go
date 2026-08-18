//go:build !wasm

package interfacetoolsdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetTimezoneGroupGetActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func TimezoneGroupGetActionPathParameterFromCli(c *cli.Command) TimezoneGroupGetActionPathParameter {
	return TimezoneGroupGetActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x TimezoneGroupGetActionRequest) IsCli() bool {
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

// TimezoneGroupGetActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the TimezoneGroupGetAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func TimezoneGroupGetActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetTimezoneGroupGetActionPathParameterCliFlags(""))...)
	return flags
}

// TimezoneGroupGetActionCliHandler builds a full *cli.Command for the
// TimezoneGroupGetAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a TimezoneGroupGetActionRequest the same way
// TimezoneGroupGetActionHandler (Gin) and TimezoneGroupGetActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func TimezoneGroupGetActionCliHandler(
	handler func(c TimezoneGroupGetActionRequest) (*TimezoneGroupGetActionResponse, error),
) *cli.Command {
	meta := TimezoneGroupGetActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: TimezoneGroupGetActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := TimezoneGroupGetActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Params:      TimezoneGroupGetActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// TimezoneGroupGetActionCli is a high-level convenience wrapper around
// TimezoneGroupGetActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way TimezoneGroupGetActionGin
// registers a route on a Gin engine.
func TimezoneGroupGetActionCli(
	app *cli.Command,
	handler func(c TimezoneGroupGetActionRequest) (*TimezoneGroupGetActionResponse, error),
) {
	app.Commands = append(app.Commands, TimezoneGroupGetActionCliHandler(handler))
}
