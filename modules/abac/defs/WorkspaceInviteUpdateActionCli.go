//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetWorkspaceInviteUpdateActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func WorkspaceInviteUpdateActionPathParameterFromCli(c *cli.Command) WorkspaceInviteUpdateActionPathParameter {
	return WorkspaceInviteUpdateActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x WorkspaceInviteUpdateActionRequest) IsCli() bool {
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

// WorkspaceInviteUpdateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the WorkspaceInviteUpdateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func WorkspaceInviteUpdateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWorkspaceInviteOptionalDtoCliFlags(""))...)
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWorkspaceInviteUpdateActionPathParameterCliFlags(""))...)
	return flags
}

// WorkspaceInviteUpdateActionCliHandler builds a full *cli.Command for the
// WorkspaceInviteUpdateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a WorkspaceInviteUpdateActionRequest the same way
// WorkspaceInviteUpdateActionHandler (Gin) and WorkspaceInviteUpdateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func WorkspaceInviteUpdateActionCliHandler(
	handler func(c WorkspaceInviteUpdateActionRequest) (*WorkspaceInviteUpdateActionResponse, error),
) *cli.Command {
	meta := WorkspaceInviteUpdateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: WorkspaceInviteUpdateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := WorkspaceInviteUpdateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastWorkspaceInviteOptionalDtoFromCli(c),
			Params:      WorkspaceInviteUpdateActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// WorkspaceInviteUpdateActionCli is a high-level convenience wrapper around
// WorkspaceInviteUpdateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way WorkspaceInviteUpdateActionGin
// registers a route on a Gin engine.
func WorkspaceInviteUpdateActionCli(
	app *cli.Command,
	handler func(c WorkspaceInviteUpdateActionRequest) (*WorkspaceInviteUpdateActionResponse, error),
) {
	app.Commands = append(app.Commands, WorkspaceInviteUpdateActionCliHandler(handler))
}
