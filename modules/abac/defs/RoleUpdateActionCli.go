//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetRoleUpdateActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func RoleUpdateActionPathParameterFromCli(c *cli.Command) RoleUpdateActionPathParameter {
	return RoleUpdateActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x RoleUpdateActionRequest) IsCli() bool {
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

// RoleUpdateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the RoleUpdateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func RoleUpdateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetRoleOptionalDtoCliFlags(""))...)
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetRoleUpdateActionPathParameterCliFlags(""))...)
	return flags
}

// RoleUpdateActionCliHandler builds a full *cli.Command for the
// RoleUpdateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a RoleUpdateActionRequest the same way
// RoleUpdateActionHandler (Gin) and RoleUpdateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func RoleUpdateActionCliHandler(
	handler func(c RoleUpdateActionRequest) (*RoleUpdateActionResponse, error),
) *cli.Command {
	meta := RoleUpdateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: RoleUpdateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := RoleUpdateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastRoleOptionalDtoFromCli(c),
			Params:      RoleUpdateActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// RoleUpdateActionCli is a high-level convenience wrapper around
// RoleUpdateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way RoleUpdateActionGin
// registers a route on a Gin engine.
func RoleUpdateActionCli(
	app *cli.Command,
	handler func(c RoleUpdateActionRequest) (*RoleUpdateActionResponse, error),
) {
	app.Commands = append(app.Commands, RoleUpdateActionCliHandler(handler))
}
