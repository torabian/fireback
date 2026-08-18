//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x PassportAwareDeleteActionRequest) IsCli() bool {
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

// PassportAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PassportAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PassportAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPassportAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// PassportAwareDeleteActionCliHandler builds a full *cli.Command for the
// PassportAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PassportAwareDeleteActionRequest the same way
// PassportAwareDeleteActionHandler (Gin) and PassportAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PassportAwareDeleteActionCliHandler(
	handler func(c PassportAwareDeleteActionRequest) (*PassportAwareDeleteActionResponse, error),
) *cli.Command {
	meta := PassportAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PassportAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PassportAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastPassportAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PassportAwareDeleteActionCli is a high-level convenience wrapper around
// PassportAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PassportAwareDeleteActionGin
// registers a route on a Gin engine.
func PassportAwareDeleteActionCli(
	app *cli.Command,
	handler func(c PassportAwareDeleteActionRequest) (*PassportAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, PassportAwareDeleteActionCliHandler(handler))
}
