//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetRoleGetActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func RoleGetActionPathParameterFromCli(c *cli.Command) RoleGetActionPathParameter {
	return RoleGetActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x RoleGetActionRequest) IsCli() bool {
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

// RoleGetActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the RoleGetAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func RoleGetActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetRoleGetActionPathParameterCliFlags(""))...)
	return flags
}

// RoleGetActionCliHandler builds a full *cli.Command for the
// RoleGetAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a RoleGetActionRequest the same way
// RoleGetActionHandler (Gin) and RoleGetActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func RoleGetActionCliHandler(
	handler func(c RoleGetActionRequest) (*RoleGetActionResponse, error),
) *cli.Command {
	meta := RoleGetActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: RoleGetActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := RoleGetActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Params:      RoleGetActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// RoleGetActionCli is a high-level convenience wrapper around
// RoleGetActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way RoleGetActionGin
// registers a route on a Gin engine.
func RoleGetActionCli(
	app *cli.Command,
	handler func(c RoleGetActionRequest) (*RoleGetActionResponse, error),
) {
	app.Commands = append(app.Commands, RoleGetActionCliHandler(handler))
}
