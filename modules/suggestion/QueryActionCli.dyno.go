//go:build !wasm

package suggestion

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x QueryActionRequest) IsCli() bool {
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

// QueryActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the QueryAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func QueryActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetQueryActionReqCliFlags(""))...)
	return flags
}

// QueryActionCliHandler builds a full *cli.Command for the
// QueryAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a QueryActionRequest the same way
// QueryActionHandler (Gin) and QueryActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func QueryActionCliHandler(
	handler func(c QueryActionRequest) (*QueryActionResponse, error),
) *cli.Command {
	meta := QueryActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: QueryActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := QueryActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastQueryActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// QueryActionCli is a high-level convenience wrapper around
// QueryActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way QueryActionGin
// registers a route on a Gin engine.
func QueryActionCli(
	app *cli.Command,
	handler func(c QueryActionRequest) (*QueryActionResponse, error),
) {
	app.Commands = append(app.Commands, QueryActionCliHandler(handler))
}
