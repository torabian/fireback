//go:build !wasm

package interfacetoolsdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x TableViewSizingCreateActionRequest) IsCli() bool {
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

// TableViewSizingCreateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the TableViewSizingCreateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func TableViewSizingCreateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetTableViewSizingDtoCliFlags(""))...)
	return flags
}

// TableViewSizingCreateActionCliHandler builds a full *cli.Command for the
// TableViewSizingCreateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a TableViewSizingCreateActionRequest the same way
// TableViewSizingCreateActionHandler (Gin) and TableViewSizingCreateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func TableViewSizingCreateActionCliHandler(
	handler func(c TableViewSizingCreateActionRequest) (*TableViewSizingCreateActionResponse, error),
) *cli.Command {
	meta := TableViewSizingCreateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: TableViewSizingCreateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := TableViewSizingCreateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastTableViewSizingDtoFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// TableViewSizingCreateActionCli is a high-level convenience wrapper around
// TableViewSizingCreateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way TableViewSizingCreateActionGin
// registers a route on a Gin engine.
func TableViewSizingCreateActionCli(
	app *cli.Command,
	handler func(c TableViewSizingCreateActionRequest) (*TableViewSizingCreateActionResponse, error),
) {
	app.Commands = append(app.Commands, TableViewSizingCreateActionCliHandler(handler))
}
