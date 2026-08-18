//go:build !wasm

package interfacetoolsdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x AppMenuCreateActionRequest) IsCli() bool {
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

// AppMenuCreateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the AppMenuCreateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func AppMenuCreateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetAppMenuDtoCliFlags(""))...)
	return flags
}

// AppMenuCreateActionCliHandler builds a full *cli.Command for the
// AppMenuCreateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a AppMenuCreateActionRequest the same way
// AppMenuCreateActionHandler (Gin) and AppMenuCreateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func AppMenuCreateActionCliHandler(
	handler func(c AppMenuCreateActionRequest) (*AppMenuCreateActionResponse, error),
) *cli.Command {
	meta := AppMenuCreateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: AppMenuCreateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := AppMenuCreateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastAppMenuDtoFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// AppMenuCreateActionCli is a high-level convenience wrapper around
// AppMenuCreateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way AppMenuCreateActionGin
// registers a route on a Gin engine.
func AppMenuCreateActionCli(
	app *cli.Command,
	handler func(c AppMenuCreateActionRequest) (*AppMenuCreateActionResponse, error),
) {
	app.Commands = append(app.Commands, AppMenuCreateActionCliHandler(handler))
}
