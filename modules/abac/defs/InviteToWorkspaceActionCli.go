//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x InviteToWorkspaceActionRequest) IsCli() bool {
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

// InviteToWorkspaceActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the InviteToWorkspaceAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func InviteToWorkspaceActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWorkspaceInvitationDtoCliFlags(""))...)
	return flags
}

// InviteToWorkspaceActionCliHandler builds a full *cli.Command for the
// InviteToWorkspaceAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a InviteToWorkspaceActionRequest the same way
// InviteToWorkspaceActionHandler (Gin) and InviteToWorkspaceActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func InviteToWorkspaceActionCliHandler(
	handler func(c InviteToWorkspaceActionRequest) (*InviteToWorkspaceActionResponse, error),
) *cli.Command {
	meta := InviteToWorkspaceActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: InviteToWorkspaceActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := InviteToWorkspaceActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastWorkspaceInvitationDtoFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// InviteToWorkspaceActionCli is a high-level convenience wrapper around
// InviteToWorkspaceActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way InviteToWorkspaceActionGin
// registers a route on a Gin engine.
func InviteToWorkspaceActionCli(
	app *cli.Command,
	handler func(c InviteToWorkspaceActionRequest) (*InviteToWorkspaceActionResponse, error),
) {
	app.Commands = append(app.Commands, InviteToWorkspaceActionCliHandler(handler))
}
