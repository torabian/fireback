//go:build !wasm

package interfacetoolsdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetAppMenuUpdateActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func AppMenuUpdateActionPathParameterFromCli(c *cli.Command) AppMenuUpdateActionPathParameter {
	return AppMenuUpdateActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x AppMenuUpdateActionRequest) IsCli() bool {
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

// AppMenuUpdateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the AppMenuUpdateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func AppMenuUpdateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetAppMenuOptionalDtoCliFlags(""))...)
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetAppMenuUpdateActionPathParameterCliFlags(""))...)
	return flags
}

// AppMenuUpdateActionCliHandler builds a full *cli.Command for the
// AppMenuUpdateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a AppMenuUpdateActionRequest the same way
// AppMenuUpdateActionHandler (Gin) and AppMenuUpdateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func AppMenuUpdateActionCliHandler(
	handler func(c AppMenuUpdateActionRequest) (*AppMenuUpdateActionResponse, error),
) *cli.Command {
	meta := AppMenuUpdateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: AppMenuUpdateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := AppMenuUpdateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastAppMenuOptionalDtoFromCli(c),
			Params:      AppMenuUpdateActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// AppMenuUpdateActionCli is a high-level convenience wrapper around
// AppMenuUpdateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way AppMenuUpdateActionGin
// registers a route on a Gin engine.
func AppMenuUpdateActionCli(
	app *cli.Command,
	handler func(c AppMenuUpdateActionRequest) (*AppMenuUpdateActionResponse, error),
) {
	app.Commands = append(app.Commands, AppMenuUpdateActionCliHandler(handler))
}
