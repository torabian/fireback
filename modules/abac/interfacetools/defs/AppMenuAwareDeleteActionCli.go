//go:build !wasm

package interfacetoolsdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x AppMenuAwareDeleteActionRequest) IsCli() bool {
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

// AppMenuAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the AppMenuAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func AppMenuAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetAppMenuAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// AppMenuAwareDeleteActionCliHandler builds a full *cli.Command for the
// AppMenuAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a AppMenuAwareDeleteActionRequest the same way
// AppMenuAwareDeleteActionHandler (Gin) and AppMenuAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func AppMenuAwareDeleteActionCliHandler(
	handler func(c AppMenuAwareDeleteActionRequest) (*AppMenuAwareDeleteActionResponse, error),
) *cli.Command {
	meta := AppMenuAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: AppMenuAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := AppMenuAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastAppMenuAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// AppMenuAwareDeleteActionCli is a high-level convenience wrapper around
// AppMenuAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way AppMenuAwareDeleteActionGin
// registers a route on a Gin engine.
func AppMenuAwareDeleteActionCli(
	app *cli.Command,
	handler func(c AppMenuAwareDeleteActionRequest) (*AppMenuAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, AppMenuAwareDeleteActionCliHandler(handler))
}
