//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x RemoveUserFromWorkspaceActionRequest) IsCli() bool {
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

// RemoveUserFromWorkspaceActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the RemoveUserFromWorkspaceAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func RemoveUserFromWorkspaceActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetRemoveUserFromWorkspaceActionReqCliFlags(""))...)
	return flags
}

// RemoveUserFromWorkspaceActionCliHandler builds a full *cli.Command for the
// RemoveUserFromWorkspaceAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a RemoveUserFromWorkspaceActionRequest the same way
// RemoveUserFromWorkspaceActionHandler (Gin) and RemoveUserFromWorkspaceActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func RemoveUserFromWorkspaceActionCliHandler(
	handler func(c RemoveUserFromWorkspaceActionRequest) (*RemoveUserFromWorkspaceActionResponse, error),
) *cli.Command {
	meta := RemoveUserFromWorkspaceActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: RemoveUserFromWorkspaceActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := RemoveUserFromWorkspaceActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastRemoveUserFromWorkspaceActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// RemoveUserFromWorkspaceActionCli is a high-level convenience wrapper around
// RemoveUserFromWorkspaceActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way RemoveUserFromWorkspaceActionGin
// registers a route on a Gin engine.
func RemoveUserFromWorkspaceActionCli(
	app *cli.Command,
	handler func(c RemoveUserFromWorkspaceActionRequest) (*RemoveUserFromWorkspaceActionResponse, error),
) {
	app.Commands = append(app.Commands, RemoveUserFromWorkspaceActionCliHandler(handler))
}
