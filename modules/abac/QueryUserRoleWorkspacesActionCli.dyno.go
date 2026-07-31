//go:build !wasm

package abac

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x QueryUserRoleWorkspacesActionRequest) IsCli() bool {
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

// QueryUserRoleWorkspacesActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the QueryUserRoleWorkspacesAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func QueryUserRoleWorkspacesActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	return flags
}

// QueryUserRoleWorkspacesActionCliHandler builds a full *cli.Command for the
// QueryUserRoleWorkspacesAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a QueryUserRoleWorkspacesActionRequest the same way
// QueryUserRoleWorkspacesActionHandler (Gin) and QueryUserRoleWorkspacesActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func QueryUserRoleWorkspacesActionCliHandler(
	handler func(c QueryUserRoleWorkspacesActionRequest) (*QueryUserRoleWorkspacesActionResponse, error),
) *cli.Command {
	meta := QueryUserRoleWorkspacesActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: QueryUserRoleWorkspacesActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := QueryUserRoleWorkspacesActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// QueryUserRoleWorkspacesActionCli is a high-level convenience wrapper around
// QueryUserRoleWorkspacesActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way QueryUserRoleWorkspacesActionGin
// registers a route on a Gin engine.
func QueryUserRoleWorkspacesActionCli(
	app *cli.Command,
	handler func(c QueryUserRoleWorkspacesActionRequest) (*QueryUserRoleWorkspacesActionResponse, error),
) {
	app.Commands = append(app.Commands, QueryUserRoleWorkspacesActionCliHandler(handler))
}
