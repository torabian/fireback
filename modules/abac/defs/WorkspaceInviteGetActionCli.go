//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetWorkspaceInviteGetActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func WorkspaceInviteGetActionPathParameterFromCli(c *cli.Command) WorkspaceInviteGetActionPathParameter {
	return WorkspaceInviteGetActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x WorkspaceInviteGetActionRequest) IsCli() bool {
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

// WorkspaceInviteGetActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the WorkspaceInviteGetAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func WorkspaceInviteGetActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWorkspaceInviteGetActionPathParameterCliFlags(""))...)
	return flags
}

// WorkspaceInviteGetActionCliHandler builds a full *cli.Command for the
// WorkspaceInviteGetAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a WorkspaceInviteGetActionRequest the same way
// WorkspaceInviteGetActionHandler (Gin) and WorkspaceInviteGetActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func WorkspaceInviteGetActionCliHandler(
	handler func(c WorkspaceInviteGetActionRequest) (*WorkspaceInviteGetActionResponse, error),
) *cli.Command {
	meta := WorkspaceInviteGetActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: WorkspaceInviteGetActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := WorkspaceInviteGetActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Params:      WorkspaceInviteGetActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// WorkspaceInviteGetActionCli is a high-level convenience wrapper around
// WorkspaceInviteGetActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way WorkspaceInviteGetActionGin
// registers a route on a Gin engine.
func WorkspaceInviteGetActionCli(
	app *cli.Command,
	handler func(c WorkspaceInviteGetActionRequest) (*WorkspaceInviteGetActionResponse, error),
) {
	app.Commands = append(app.Commands, WorkspaceInviteGetActionCliHandler(handler))
}
