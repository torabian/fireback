//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetPublicAuthenticationUpdateActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func PublicAuthenticationUpdateActionPathParameterFromCli(c *cli.Command) PublicAuthenticationUpdateActionPathParameter {
	return PublicAuthenticationUpdateActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x PublicAuthenticationUpdateActionRequest) IsCli() bool {
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

// PublicAuthenticationUpdateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PublicAuthenticationUpdateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PublicAuthenticationUpdateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPublicAuthenticationOptionalDtoCliFlags(""))...)
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPublicAuthenticationUpdateActionPathParameterCliFlags(""))...)
	return flags
}

// PublicAuthenticationUpdateActionCliHandler builds a full *cli.Command for the
// PublicAuthenticationUpdateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PublicAuthenticationUpdateActionRequest the same way
// PublicAuthenticationUpdateActionHandler (Gin) and PublicAuthenticationUpdateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PublicAuthenticationUpdateActionCliHandler(
	handler func(c PublicAuthenticationUpdateActionRequest) (*PublicAuthenticationUpdateActionResponse, error),
) *cli.Command {
	meta := PublicAuthenticationUpdateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PublicAuthenticationUpdateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PublicAuthenticationUpdateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastPublicAuthenticationOptionalDtoFromCli(c),
			Params:      PublicAuthenticationUpdateActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PublicAuthenticationUpdateActionCli is a high-level convenience wrapper around
// PublicAuthenticationUpdateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PublicAuthenticationUpdateActionGin
// registers a route on a Gin engine.
func PublicAuthenticationUpdateActionCli(
	app *cli.Command,
	handler func(c PublicAuthenticationUpdateActionRequest) (*PublicAuthenticationUpdateActionResponse, error),
) {
	app.Commands = append(app.Commands, PublicAuthenticationUpdateActionCliHandler(handler))
}
