//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x TokenAwareDeleteActionRequest) IsCli() bool {
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

// TokenAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the TokenAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func TokenAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetTokenAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// TokenAwareDeleteActionCliHandler builds a full *cli.Command for the
// TokenAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a TokenAwareDeleteActionRequest the same way
// TokenAwareDeleteActionHandler (Gin) and TokenAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func TokenAwareDeleteActionCliHandler(
	handler func(c TokenAwareDeleteActionRequest) (*TokenAwareDeleteActionResponse, error),
) *cli.Command {
	meta := TokenAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: TokenAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := TokenAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastTokenAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// TokenAwareDeleteActionCli is a high-level convenience wrapper around
// TokenAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way TokenAwareDeleteActionGin
// registers a route on a Gin engine.
func TokenAwareDeleteActionCli(
	app *cli.Command,
	handler func(c TokenAwareDeleteActionRequest) (*TokenAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, TokenAwareDeleteActionCliHandler(handler))
}
