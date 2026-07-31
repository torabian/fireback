//go:build !wasm

package abac

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x ClassicSignupActionRequest) IsCli() bool {
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

// ClassicSignupActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the ClassicSignupAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func ClassicSignupActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetClassicSignupActionReqCliFlags(""))...)
	return flags
}

// ClassicSignupActionCliHandler builds a full *cli.Command for the
// ClassicSignupAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a ClassicSignupActionRequest the same way
// ClassicSignupActionHandler (Gin) and ClassicSignupActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func ClassicSignupActionCliHandler(
	handler func(c ClassicSignupActionRequest) (*ClassicSignupActionResponse, error),
) *cli.Command {
	meta := ClassicSignupActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: ClassicSignupActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := ClassicSignupActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastClassicSignupActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// ClassicSignupActionCli is a high-level convenience wrapper around
// ClassicSignupActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way ClassicSignupActionGin
// registers a route on a Gin engine.
func ClassicSignupActionCli(
	app *cli.Command,
	handler func(c ClassicSignupActionRequest) (*ClassicSignupActionResponse, error),
) {
	app.Commands = append(app.Commands, ClassicSignupActionCliHandler(handler))
}
