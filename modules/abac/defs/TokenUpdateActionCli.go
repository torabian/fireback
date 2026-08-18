//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetTokenUpdateActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func TokenUpdateActionPathParameterFromCli(c *cli.Command) TokenUpdateActionPathParameter {
	return TokenUpdateActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x TokenUpdateActionRequest) IsCli() bool {
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

// TokenUpdateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the TokenUpdateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func TokenUpdateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetTokenOptionalDtoCliFlags(""))...)
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetTokenUpdateActionPathParameterCliFlags(""))...)
	return flags
}

// TokenUpdateActionCliHandler builds a full *cli.Command for the
// TokenUpdateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a TokenUpdateActionRequest the same way
// TokenUpdateActionHandler (Gin) and TokenUpdateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func TokenUpdateActionCliHandler(
	handler func(c TokenUpdateActionRequest) (*TokenUpdateActionResponse, error),
) *cli.Command {
	meta := TokenUpdateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: TokenUpdateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := TokenUpdateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastTokenOptionalDtoFromCli(c),
			Params:      TokenUpdateActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// TokenUpdateActionCli is a high-level convenience wrapper around
// TokenUpdateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way TokenUpdateActionGin
// registers a route on a Gin engine.
func TokenUpdateActionCli(
	app *cli.Command,
	handler func(c TokenUpdateActionRequest) (*TokenUpdateActionResponse, error),
) {
	app.Commands = append(app.Commands, TokenUpdateActionCliHandler(handler))
}
