//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x WhoamiActionRequest) IsCli() bool {
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

// WhoamiActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the WhoamiAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func WhoamiActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	return flags
}

// WhoamiActionCliHandler builds a full *cli.Command for the
// WhoamiAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a WhoamiActionRequest the same way
// WhoamiActionHandler (Gin) and WhoamiActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func WhoamiActionCliHandler(
	handler func(c WhoamiActionRequest) (*WhoamiActionResponse, error),
) *cli.Command {
	meta := WhoamiActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: WhoamiActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := WhoamiActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// WhoamiActionCli is a high-level convenience wrapper around
// WhoamiActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way WhoamiActionGin
// registers a route on a Gin engine.
func WhoamiActionCli(
	app *cli.Command,
	handler func(c WhoamiActionRequest) (*WhoamiActionResponse, error),
) {
	app.Commands = append(app.Commands, WhoamiActionCliHandler(handler))
}
