//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x WorkspaceInviteAwareDeleteActionRequest) IsCli() bool {
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

// WorkspaceInviteAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the WorkspaceInviteAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func WorkspaceInviteAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWorkspaceInviteAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// WorkspaceInviteAwareDeleteActionCliHandler builds a full *cli.Command for the
// WorkspaceInviteAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a WorkspaceInviteAwareDeleteActionRequest the same way
// WorkspaceInviteAwareDeleteActionHandler (Gin) and WorkspaceInviteAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func WorkspaceInviteAwareDeleteActionCliHandler(
	handler func(c WorkspaceInviteAwareDeleteActionRequest) (*WorkspaceInviteAwareDeleteActionResponse, error),
) *cli.Command {
	meta := WorkspaceInviteAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: WorkspaceInviteAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := WorkspaceInviteAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastWorkspaceInviteAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// WorkspaceInviteAwareDeleteActionCli is a high-level convenience wrapper around
// WorkspaceInviteAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way WorkspaceInviteAwareDeleteActionGin
// registers a route on a Gin engine.
func WorkspaceInviteAwareDeleteActionCli(
	app *cli.Command,
	handler func(c WorkspaceInviteAwareDeleteActionRequest) (*WorkspaceInviteAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, WorkspaceInviteAwareDeleteActionCliHandler(handler))
}
