//go:build !wasm

package abac

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x QueryWorkspaceTypesPubliclyActionRequest) IsCli() bool {
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

// QueryWorkspaceTypesPubliclyActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the QueryWorkspaceTypesPubliclyAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func QueryWorkspaceTypesPubliclyActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	return flags
}

// QueryWorkspaceTypesPubliclyActionCliHandler builds a full *cli.Command for the
// QueryWorkspaceTypesPubliclyAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a QueryWorkspaceTypesPubliclyActionRequest the same way
// QueryWorkspaceTypesPubliclyActionHandler (Gin) and QueryWorkspaceTypesPubliclyActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func QueryWorkspaceTypesPubliclyActionCliHandler(
	handler func(c QueryWorkspaceTypesPubliclyActionRequest) (*QueryWorkspaceTypesPubliclyActionResponse, error),
) *cli.Command {
	meta := QueryWorkspaceTypesPubliclyActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: QueryWorkspaceTypesPubliclyActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := QueryWorkspaceTypesPubliclyActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// QueryWorkspaceTypesPubliclyActionCli is a high-level convenience wrapper around
// QueryWorkspaceTypesPubliclyActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way QueryWorkspaceTypesPubliclyActionGin
// registers a route on a Gin engine.
func QueryWorkspaceTypesPubliclyActionCli(
	app *cli.Command,
	handler func(c QueryWorkspaceTypesPubliclyActionRequest) (*QueryWorkspaceTypesPubliclyActionResponse, error),
) {
	app.Commands = append(app.Commands, QueryWorkspaceTypesPubliclyActionCliHandler(handler))
}
