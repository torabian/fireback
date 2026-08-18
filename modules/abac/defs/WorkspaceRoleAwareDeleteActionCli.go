//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x WorkspaceRoleAwareDeleteActionRequest) IsCli() bool {
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

// WorkspaceRoleAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the WorkspaceRoleAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func WorkspaceRoleAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWorkspaceRoleAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// WorkspaceRoleAwareDeleteActionCliHandler builds a full *cli.Command for the
// WorkspaceRoleAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a WorkspaceRoleAwareDeleteActionRequest the same way
// WorkspaceRoleAwareDeleteActionHandler (Gin) and WorkspaceRoleAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func WorkspaceRoleAwareDeleteActionCliHandler(
	handler func(c WorkspaceRoleAwareDeleteActionRequest) (*WorkspaceRoleAwareDeleteActionResponse, error),
) *cli.Command {
	meta := WorkspaceRoleAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: WorkspaceRoleAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := WorkspaceRoleAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastWorkspaceRoleAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// WorkspaceRoleAwareDeleteActionCli is a high-level convenience wrapper around
// WorkspaceRoleAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way WorkspaceRoleAwareDeleteActionGin
// registers a route on a Gin engine.
func WorkspaceRoleAwareDeleteActionCli(
	app *cli.Command,
	handler func(c WorkspaceRoleAwareDeleteActionRequest) (*WorkspaceRoleAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, WorkspaceRoleAwareDeleteActionCliHandler(handler))
}
