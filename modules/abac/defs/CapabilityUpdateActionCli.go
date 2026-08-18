//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetCapabilityUpdateActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func CapabilityUpdateActionPathParameterFromCli(c *cli.Command) CapabilityUpdateActionPathParameter {
	return CapabilityUpdateActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x CapabilityUpdateActionRequest) IsCli() bool {
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

// CapabilityUpdateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the CapabilityUpdateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func CapabilityUpdateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetCapabilityOptionalDtoCliFlags(""))...)
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetCapabilityUpdateActionPathParameterCliFlags(""))...)
	return flags
}

// CapabilityUpdateActionCliHandler builds a full *cli.Command for the
// CapabilityUpdateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a CapabilityUpdateActionRequest the same way
// CapabilityUpdateActionHandler (Gin) and CapabilityUpdateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func CapabilityUpdateActionCliHandler(
	handler func(c CapabilityUpdateActionRequest) (*CapabilityUpdateActionResponse, error),
) *cli.Command {
	meta := CapabilityUpdateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: CapabilityUpdateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := CapabilityUpdateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastCapabilityOptionalDtoFromCli(c),
			Params:      CapabilityUpdateActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// CapabilityUpdateActionCli is a high-level convenience wrapper around
// CapabilityUpdateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way CapabilityUpdateActionGin
// registers a route on a Gin engine.
func CapabilityUpdateActionCli(
	app *cli.Command,
	handler func(c CapabilityUpdateActionRequest) (*CapabilityUpdateActionResponse, error),
) {
	app.Commands = append(app.Commands, CapabilityUpdateActionCliHandler(handler))
}
