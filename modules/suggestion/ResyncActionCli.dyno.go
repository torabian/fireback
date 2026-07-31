//go:build !wasm

package suggestion

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x ResyncActionRequest) IsCli() bool {
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

// ResyncActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the ResyncAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func ResyncActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	return flags
}

// ResyncActionCliHandler builds a full *cli.Command for the
// ResyncAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a ResyncActionRequest the same way
// ResyncActionHandler (Gin) and ResyncActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func ResyncActionCliHandler(
	handler func(c ResyncActionRequest) (*ResyncActionResponse, error),
) *cli.Command {
	meta := ResyncActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: ResyncActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := ResyncActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// ResyncActionCli is a high-level convenience wrapper around
// ResyncActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way ResyncActionGin
// registers a route on a Gin engine.
func ResyncActionCli(
	app *cli.Command,
	handler func(c ResyncActionRequest) (*ResyncActionResponse, error),
) {
	app.Commands = append(app.Commands, ResyncActionCliHandler(handler))
}
