//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x RoleAwareDeleteActionRequest) IsCli() bool {
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

// RoleAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the RoleAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func RoleAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetRoleAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// RoleAwareDeleteActionCliHandler builds a full *cli.Command for the
// RoleAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a RoleAwareDeleteActionRequest the same way
// RoleAwareDeleteActionHandler (Gin) and RoleAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func RoleAwareDeleteActionCliHandler(
	handler func(c RoleAwareDeleteActionRequest) (*RoleAwareDeleteActionResponse, error),
) *cli.Command {
	meta := RoleAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: RoleAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := RoleAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastRoleAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// RoleAwareDeleteActionCli is a high-level convenience wrapper around
// RoleAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way RoleAwareDeleteActionGin
// registers a route on a Gin engine.
func RoleAwareDeleteActionCli(
	app *cli.Command,
	handler func(c RoleAwareDeleteActionRequest) (*RoleAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, RoleAwareDeleteActionCliHandler(handler))
}
