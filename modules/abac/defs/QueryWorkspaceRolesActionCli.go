//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x QueryWorkspaceRolesActionRequest) IsCli() bool {
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

// QueryWorkspaceRolesActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the QueryWorkspaceRolesAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func QueryWorkspaceRolesActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetQueryWorkspaceRolesActionReqCliFlags(""))...)
	return flags
}

// QueryWorkspaceRolesActionCliHandler builds a full *cli.Command for the
// QueryWorkspaceRolesAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a QueryWorkspaceRolesActionRequest the same way
// QueryWorkspaceRolesActionHandler (Gin) and QueryWorkspaceRolesActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func QueryWorkspaceRolesActionCliHandler(
	handler func(c QueryWorkspaceRolesActionRequest) (*QueryWorkspaceRolesActionResponse, error),
) *cli.Command {
	meta := QueryWorkspaceRolesActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: QueryWorkspaceRolesActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := QueryWorkspaceRolesActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastQueryWorkspaceRolesActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// QueryWorkspaceRolesActionCli is a high-level convenience wrapper around
// QueryWorkspaceRolesActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way QueryWorkspaceRolesActionGin
// registers a route on a Gin engine.
func QueryWorkspaceRolesActionCli(
	app *cli.Command,
	handler func(c QueryWorkspaceRolesActionRequest) (*QueryWorkspaceRolesActionResponse, error),
) {
	app.Commands = append(app.Commands, QueryWorkspaceRolesActionCliHandler(handler))
}
