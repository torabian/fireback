//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetPublicJoinKeyUpdateActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func PublicJoinKeyUpdateActionPathParameterFromCli(c *cli.Command) PublicJoinKeyUpdateActionPathParameter {
	return PublicJoinKeyUpdateActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x PublicJoinKeyUpdateActionRequest) IsCli() bool {
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

// PublicJoinKeyUpdateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PublicJoinKeyUpdateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PublicJoinKeyUpdateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPublicJoinKeyOptionalDtoCliFlags(""))...)
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPublicJoinKeyUpdateActionPathParameterCliFlags(""))...)
	return flags
}

// PublicJoinKeyUpdateActionCliHandler builds a full *cli.Command for the
// PublicJoinKeyUpdateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PublicJoinKeyUpdateActionRequest the same way
// PublicJoinKeyUpdateActionHandler (Gin) and PublicJoinKeyUpdateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PublicJoinKeyUpdateActionCliHandler(
	handler func(c PublicJoinKeyUpdateActionRequest) (*PublicJoinKeyUpdateActionResponse, error),
) *cli.Command {
	meta := PublicJoinKeyUpdateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PublicJoinKeyUpdateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PublicJoinKeyUpdateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastPublicJoinKeyOptionalDtoFromCli(c),
			Params:      PublicJoinKeyUpdateActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PublicJoinKeyUpdateActionCli is a high-level convenience wrapper around
// PublicJoinKeyUpdateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PublicJoinKeyUpdateActionGin
// registers a route on a Gin engine.
func PublicJoinKeyUpdateActionCli(
	app *cli.Command,
	handler func(c PublicJoinKeyUpdateActionRequest) (*PublicJoinKeyUpdateActionResponse, error),
) {
	app.Commands = append(app.Commands, PublicJoinKeyUpdateActionCliHandler(handler))
}
