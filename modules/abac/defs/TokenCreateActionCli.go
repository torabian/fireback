//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x TokenCreateActionRequest) IsCli() bool {
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

// TokenCreateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the TokenCreateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func TokenCreateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetTokenDtoCliFlags(""))...)
	return flags
}

// TokenCreateActionCliHandler builds a full *cli.Command for the
// TokenCreateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a TokenCreateActionRequest the same way
// TokenCreateActionHandler (Gin) and TokenCreateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func TokenCreateActionCliHandler(
	handler func(c TokenCreateActionRequest) (*TokenCreateActionResponse, error),
) *cli.Command {
	meta := TokenCreateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: TokenCreateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := TokenCreateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastTokenDtoFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// TokenCreateActionCli is a high-level convenience wrapper around
// TokenCreateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way TokenCreateActionGin
// registers a route on a Gin engine.
func TokenCreateActionCli(
	app *cli.Command,
	handler func(c TokenCreateActionRequest) (*TokenCreateActionResponse, error),
) {
	app.Commands = append(app.Commands, TokenCreateActionCliHandler(handler))
}
