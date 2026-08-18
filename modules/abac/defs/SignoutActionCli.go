//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x SignoutActionRequest) IsCli() bool {
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

// SignoutActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the SignoutAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func SignoutActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetSignoutActionReqCliFlags(""))...)
	return flags
}

// SignoutActionCliHandler builds a full *cli.Command for the
// SignoutAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a SignoutActionRequest the same way
// SignoutActionHandler (Gin) and SignoutActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func SignoutActionCliHandler(
	handler func(c SignoutActionRequest) (*SignoutActionResponse, error),
) *cli.Command {
	meta := SignoutActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: SignoutActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := SignoutActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastSignoutActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// SignoutActionCli is a high-level convenience wrapper around
// SignoutActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way SignoutActionGin
// registers a route on a Gin engine.
func SignoutActionCli(
	app *cli.Command,
	handler func(c SignoutActionRequest) (*SignoutActionResponse, error),
) {
	app.Commands = append(app.Commands, SignoutActionCliHandler(handler))
}
