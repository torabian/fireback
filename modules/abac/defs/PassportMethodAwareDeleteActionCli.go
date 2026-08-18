//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x PassportMethodAwareDeleteActionRequest) IsCli() bool {
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

// PassportMethodAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PassportMethodAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PassportMethodAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPassportMethodAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// PassportMethodAwareDeleteActionCliHandler builds a full *cli.Command for the
// PassportMethodAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PassportMethodAwareDeleteActionRequest the same way
// PassportMethodAwareDeleteActionHandler (Gin) and PassportMethodAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PassportMethodAwareDeleteActionCliHandler(
	handler func(c PassportMethodAwareDeleteActionRequest) (*PassportMethodAwareDeleteActionResponse, error),
) *cli.Command {
	meta := PassportMethodAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PassportMethodAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PassportMethodAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastPassportMethodAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PassportMethodAwareDeleteActionCli is a high-level convenience wrapper around
// PassportMethodAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PassportMethodAwareDeleteActionGin
// registers a route on a Gin engine.
func PassportMethodAwareDeleteActionCli(
	app *cli.Command,
	handler func(c PassportMethodAwareDeleteActionRequest) (*PassportMethodAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, PassportMethodAwareDeleteActionCliHandler(handler))
}
