//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetTokenGetActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func TokenGetActionPathParameterFromCli(c *cli.Command) TokenGetActionPathParameter {
	return TokenGetActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x TokenGetActionRequest) IsCli() bool {
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

// TokenGetActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the TokenGetAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func TokenGetActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetTokenGetActionPathParameterCliFlags(""))...)
	return flags
}

// TokenGetActionCliHandler builds a full *cli.Command for the
// TokenGetAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a TokenGetActionRequest the same way
// TokenGetActionHandler (Gin) and TokenGetActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func TokenGetActionCliHandler(
	handler func(c TokenGetActionRequest) (*TokenGetActionResponse, error),
) *cli.Command {
	meta := TokenGetActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: TokenGetActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := TokenGetActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Params:      TokenGetActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// TokenGetActionCli is a high-level convenience wrapper around
// TokenGetActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way TokenGetActionGin
// registers a route on a Gin engine.
func TokenGetActionCli(
	app *cli.Command,
	handler func(c TokenGetActionRequest) (*TokenGetActionResponse, error),
) {
	app.Commands = append(app.Commands, TokenGetActionCliHandler(handler))
}
