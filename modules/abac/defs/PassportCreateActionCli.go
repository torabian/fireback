//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x PassportCreateActionRequest) IsCli() bool {
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

// PassportCreateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PassportCreateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PassportCreateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPassportDtoCliFlags(""))...)
	return flags
}

// PassportCreateActionCliHandler builds a full *cli.Command for the
// PassportCreateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PassportCreateActionRequest the same way
// PassportCreateActionHandler (Gin) and PassportCreateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PassportCreateActionCliHandler(
	handler func(c PassportCreateActionRequest) (*PassportCreateActionResponse, error),
) *cli.Command {
	meta := PassportCreateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PassportCreateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PassportCreateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastPassportDtoFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PassportCreateActionCli is a high-level convenience wrapper around
// PassportCreateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PassportCreateActionGin
// registers a route on a Gin engine.
func PassportCreateActionCli(
	app *cli.Command,
	handler func(c PassportCreateActionRequest) (*PassportCreateActionResponse, error),
) {
	app.Commands = append(app.Commands, PassportCreateActionCliHandler(handler))
}
