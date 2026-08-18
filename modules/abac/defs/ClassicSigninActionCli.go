//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x ClassicSigninActionRequest) IsCli() bool {
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

// ClassicSigninActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the ClassicSigninAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func ClassicSigninActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetClassicSigninActionReqCliFlags(""))...)
	return flags
}

// ClassicSigninActionCliHandler builds a full *cli.Command for the
// ClassicSigninAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a ClassicSigninActionRequest the same way
// ClassicSigninActionHandler (Gin) and ClassicSigninActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func ClassicSigninActionCliHandler(
	handler func(c ClassicSigninActionRequest) (*ClassicSigninActionResponse, error),
) *cli.Command {
	meta := ClassicSigninActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: ClassicSigninActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := ClassicSigninActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastClassicSigninActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// ClassicSigninActionCli is a high-level convenience wrapper around
// ClassicSigninActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way ClassicSigninActionGin
// registers a route on a Gin engine.
func ClassicSigninActionCli(
	app *cli.Command,
	handler func(c ClassicSigninActionRequest) (*ClassicSigninActionResponse, error),
) {
	app.Commands = append(app.Commands, ClassicSigninActionCliHandler(handler))
}
