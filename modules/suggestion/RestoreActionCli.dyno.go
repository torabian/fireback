//go:build !wasm

package suggestion

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x RestoreActionRequest) IsCli() bool {
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

// RestoreActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the RestoreAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func RestoreActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	return flags
}

// RestoreActionCliHandler builds a full *cli.Command for the
// RestoreAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a RestoreActionRequest the same way
// RestoreActionHandler (Gin) and RestoreActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func RestoreActionCliHandler(
	handler func(c RestoreActionRequest) (*RestoreActionResponse, error),
) *cli.Command {
	meta := RestoreActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: RestoreActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := RestoreActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// RestoreActionCli is a high-level convenience wrapper around
// RestoreActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way RestoreActionGin
// registers a route on a Gin engine.
func RestoreActionCli(
	app *cli.Command,
	handler func(c RestoreActionRequest) (*RestoreActionResponse, error),
) {
	app.Commands = append(app.Commands, RestoreActionCliHandler(handler))
}
