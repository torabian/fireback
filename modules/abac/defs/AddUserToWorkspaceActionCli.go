//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x AddUserToWorkspaceActionRequest) IsCli() bool {
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

// AddUserToWorkspaceActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the AddUserToWorkspaceAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func AddUserToWorkspaceActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetAddUserToWorkspaceActionReqCliFlags(""))...)
	return flags
}

// AddUserToWorkspaceActionCliHandler builds a full *cli.Command for the
// AddUserToWorkspaceAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a AddUserToWorkspaceActionRequest the same way
// AddUserToWorkspaceActionHandler (Gin) and AddUserToWorkspaceActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func AddUserToWorkspaceActionCliHandler(
	handler func(c AddUserToWorkspaceActionRequest) (*AddUserToWorkspaceActionResponse, error),
) *cli.Command {
	meta := AddUserToWorkspaceActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: AddUserToWorkspaceActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := AddUserToWorkspaceActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastAddUserToWorkspaceActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// AddUserToWorkspaceActionCli is a high-level convenience wrapper around
// AddUserToWorkspaceActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way AddUserToWorkspaceActionGin
// registers a route on a Gin engine.
func AddUserToWorkspaceActionCli(
	app *cli.Command,
	handler func(c AddUserToWorkspaceActionRequest) (*AddUserToWorkspaceActionResponse, error),
) {
	app.Commands = append(app.Commands, AddUserToWorkspaceActionCliHandler(handler))
}
