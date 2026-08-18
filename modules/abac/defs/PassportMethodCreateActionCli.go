//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x PassportMethodCreateActionRequest) IsCli() bool {
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

// PassportMethodCreateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PassportMethodCreateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PassportMethodCreateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPassportMethodDtoCliFlags(""))...)
	return flags
}

// PassportMethodCreateActionCliHandler builds a full *cli.Command for the
// PassportMethodCreateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PassportMethodCreateActionRequest the same way
// PassportMethodCreateActionHandler (Gin) and PassportMethodCreateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PassportMethodCreateActionCliHandler(
	handler func(c PassportMethodCreateActionRequest) (*PassportMethodCreateActionResponse, error),
) *cli.Command {
	meta := PassportMethodCreateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PassportMethodCreateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PassportMethodCreateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastPassportMethodDtoFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PassportMethodCreateActionCli is a high-level convenience wrapper around
// PassportMethodCreateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PassportMethodCreateActionGin
// registers a route on a Gin engine.
func PassportMethodCreateActionCli(
	app *cli.Command,
	handler func(c PassportMethodCreateActionRequest) (*PassportMethodCreateActionResponse, error),
) {
	app.Commands = append(app.Commands, PassportMethodCreateActionCliHandler(handler))
}
