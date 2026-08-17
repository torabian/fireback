//go:build !wasm

package interfacetoolsdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetAppMenuGetActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func AppMenuGetActionPathParameterFromCli(c *cli.Command) AppMenuGetActionPathParameter {
	return AppMenuGetActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x AppMenuGetActionRequest) IsCli() bool {
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

// AppMenuGetActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the AppMenuGetAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func AppMenuGetActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetAppMenuGetActionPathParameterCliFlags(""))...)
	return flags
}

// AppMenuGetActionCliHandler builds a full *cli.Command for the
// AppMenuGetAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a AppMenuGetActionRequest the same way
// AppMenuGetActionHandler (Gin) and AppMenuGetActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func AppMenuGetActionCliHandler(
	handler func(c AppMenuGetActionRequest) (*AppMenuGetActionResponse, error),
) *cli.Command {
	meta := AppMenuGetActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: AppMenuGetActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := AppMenuGetActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Params:      AppMenuGetActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// AppMenuGetActionCli is a high-level convenience wrapper around
// AppMenuGetActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way AppMenuGetActionGin
// registers a route on a Gin engine.
func AppMenuGetActionCli(
	app *cli.Command,
	handler func(c AppMenuGetActionRequest) (*AppMenuGetActionResponse, error),
) {
	app.Commands = append(app.Commands, AppMenuGetActionCliHandler(handler))
}
