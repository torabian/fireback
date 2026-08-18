//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x PendingWorkspaceInviteAwareDeleteActionRequest) IsCli() bool {
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

// PendingWorkspaceInviteAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PendingWorkspaceInviteAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PendingWorkspaceInviteAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPendingWorkspaceInviteAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// PendingWorkspaceInviteAwareDeleteActionCliHandler builds a full *cli.Command for the
// PendingWorkspaceInviteAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PendingWorkspaceInviteAwareDeleteActionRequest the same way
// PendingWorkspaceInviteAwareDeleteActionHandler (Gin) and PendingWorkspaceInviteAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PendingWorkspaceInviteAwareDeleteActionCliHandler(
	handler func(c PendingWorkspaceInviteAwareDeleteActionRequest) (*PendingWorkspaceInviteAwareDeleteActionResponse, error),
) *cli.Command {
	meta := PendingWorkspaceInviteAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PendingWorkspaceInviteAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PendingWorkspaceInviteAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastPendingWorkspaceInviteAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PendingWorkspaceInviteAwareDeleteActionCli is a high-level convenience wrapper around
// PendingWorkspaceInviteAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PendingWorkspaceInviteAwareDeleteActionGin
// registers a route on a Gin engine.
func PendingWorkspaceInviteAwareDeleteActionCli(
	app *cli.Command,
	handler func(c PendingWorkspaceInviteAwareDeleteActionRequest) (*PendingWorkspaceInviteAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, PendingWorkspaceInviteAwareDeleteActionCliHandler(handler))
}
