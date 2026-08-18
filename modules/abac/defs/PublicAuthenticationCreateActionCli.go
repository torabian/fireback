//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x PublicAuthenticationCreateActionRequest) IsCli() bool {
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

// PublicAuthenticationCreateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PublicAuthenticationCreateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PublicAuthenticationCreateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPublicAuthenticationDtoCliFlags(""))...)
	return flags
}

// PublicAuthenticationCreateActionCliHandler builds a full *cli.Command for the
// PublicAuthenticationCreateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PublicAuthenticationCreateActionRequest the same way
// PublicAuthenticationCreateActionHandler (Gin) and PublicAuthenticationCreateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PublicAuthenticationCreateActionCliHandler(
	handler func(c PublicAuthenticationCreateActionRequest) (*PublicAuthenticationCreateActionResponse, error),
) *cli.Command {
	meta := PublicAuthenticationCreateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PublicAuthenticationCreateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PublicAuthenticationCreateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastPublicAuthenticationDtoFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PublicAuthenticationCreateActionCli is a high-level convenience wrapper around
// PublicAuthenticationCreateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PublicAuthenticationCreateActionGin
// registers a route on a Gin engine.
func PublicAuthenticationCreateActionCli(
	app *cli.Command,
	handler func(c PublicAuthenticationCreateActionRequest) (*PublicAuthenticationCreateActionResponse, error),
) {
	app.Commands = append(app.Commands, PublicAuthenticationCreateActionCliHandler(handler))
}
