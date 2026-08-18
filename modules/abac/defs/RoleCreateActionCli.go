//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x RoleCreateActionRequest) IsCli() bool {
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

// RoleCreateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the RoleCreateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func RoleCreateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetRoleDtoCliFlags(""))...)
	return flags
}

// RoleCreateActionCliHandler builds a full *cli.Command for the
// RoleCreateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a RoleCreateActionRequest the same way
// RoleCreateActionHandler (Gin) and RoleCreateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func RoleCreateActionCliHandler(
	handler func(c RoleCreateActionRequest) (*RoleCreateActionResponse, error),
) *cli.Command {
	meta := RoleCreateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: RoleCreateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := RoleCreateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastRoleDtoFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// RoleCreateActionCli is a high-level convenience wrapper around
// RoleCreateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way RoleCreateActionGin
// registers a route on a Gin engine.
func RoleCreateActionCli(
	app *cli.Command,
	handler func(c RoleCreateActionRequest) (*RoleCreateActionResponse, error),
) {
	app.Commands = append(app.Commands, RoleCreateActionCliHandler(handler))
}
