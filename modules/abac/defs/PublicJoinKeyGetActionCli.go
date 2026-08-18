//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetPublicJoinKeyGetActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func PublicJoinKeyGetActionPathParameterFromCli(c *cli.Command) PublicJoinKeyGetActionPathParameter {
	return PublicJoinKeyGetActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x PublicJoinKeyGetActionRequest) IsCli() bool {
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

// PublicJoinKeyGetActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PublicJoinKeyGetAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PublicJoinKeyGetActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPublicJoinKeyGetActionPathParameterCliFlags(""))...)
	return flags
}

// PublicJoinKeyGetActionCliHandler builds a full *cli.Command for the
// PublicJoinKeyGetAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PublicJoinKeyGetActionRequest the same way
// PublicJoinKeyGetActionHandler (Gin) and PublicJoinKeyGetActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PublicJoinKeyGetActionCliHandler(
	handler func(c PublicJoinKeyGetActionRequest) (*PublicJoinKeyGetActionResponse, error),
) *cli.Command {
	meta := PublicJoinKeyGetActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PublicJoinKeyGetActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PublicJoinKeyGetActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Params:      PublicJoinKeyGetActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PublicJoinKeyGetActionCli is a high-level convenience wrapper around
// PublicJoinKeyGetActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PublicJoinKeyGetActionGin
// registers a route on a Gin engine.
func PublicJoinKeyGetActionCli(
	app *cli.Command,
	handler func(c PublicJoinKeyGetActionRequest) (*PublicJoinKeyGetActionResponse, error),
) {
	app.Commands = append(app.Commands, PublicJoinKeyGetActionCliHandler(handler))
}
