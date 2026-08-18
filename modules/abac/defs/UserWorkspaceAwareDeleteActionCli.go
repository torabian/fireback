//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x UserWorkspaceAwareDeleteActionRequest) IsCli() bool {
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

// UserWorkspaceAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the UserWorkspaceAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func UserWorkspaceAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetUserWorkspaceAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// UserWorkspaceAwareDeleteActionCliHandler builds a full *cli.Command for the
// UserWorkspaceAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a UserWorkspaceAwareDeleteActionRequest the same way
// UserWorkspaceAwareDeleteActionHandler (Gin) and UserWorkspaceAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func UserWorkspaceAwareDeleteActionCliHandler(
	handler func(c UserWorkspaceAwareDeleteActionRequest) (*UserWorkspaceAwareDeleteActionResponse, error),
) *cli.Command {
	meta := UserWorkspaceAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: UserWorkspaceAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := UserWorkspaceAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastUserWorkspaceAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// UserWorkspaceAwareDeleteActionCli is a high-level convenience wrapper around
// UserWorkspaceAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way UserWorkspaceAwareDeleteActionGin
// registers a route on a Gin engine.
func UserWorkspaceAwareDeleteActionCli(
	app *cli.Command,
	handler func(c UserWorkspaceAwareDeleteActionRequest) (*UserWorkspaceAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, UserWorkspaceAwareDeleteActionCliHandler(handler))
}
