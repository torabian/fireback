//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x UserWorkspaceCreateActionRequest) IsCli() bool {
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

// UserWorkspaceCreateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the UserWorkspaceCreateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func UserWorkspaceCreateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetUserWorkspaceDtoCliFlags(""))...)
	return flags
}

// UserWorkspaceCreateActionCliHandler builds a full *cli.Command for the
// UserWorkspaceCreateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a UserWorkspaceCreateActionRequest the same way
// UserWorkspaceCreateActionHandler (Gin) and UserWorkspaceCreateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func UserWorkspaceCreateActionCliHandler(
	handler func(c UserWorkspaceCreateActionRequest) (*UserWorkspaceCreateActionResponse, error),
) *cli.Command {
	meta := UserWorkspaceCreateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: UserWorkspaceCreateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := UserWorkspaceCreateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastUserWorkspaceDtoFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// UserWorkspaceCreateActionCli is a high-level convenience wrapper around
// UserWorkspaceCreateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way UserWorkspaceCreateActionGin
// registers a route on a Gin engine.
func UserWorkspaceCreateActionCli(
	app *cli.Command,
	handler func(c UserWorkspaceCreateActionRequest) (*UserWorkspaceCreateActionResponse, error),
) {
	app.Commands = append(app.Commands, UserWorkspaceCreateActionCliHandler(handler))
}
