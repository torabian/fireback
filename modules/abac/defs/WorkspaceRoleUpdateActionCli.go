//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetWorkspaceRoleUpdateActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func WorkspaceRoleUpdateActionPathParameterFromCli(c *cli.Command) WorkspaceRoleUpdateActionPathParameter {
	return WorkspaceRoleUpdateActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x WorkspaceRoleUpdateActionRequest) IsCli() bool {
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

// WorkspaceRoleUpdateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the WorkspaceRoleUpdateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func WorkspaceRoleUpdateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWorkspaceRoleOptionalDtoCliFlags(""))...)
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWorkspaceRoleUpdateActionPathParameterCliFlags(""))...)
	return flags
}

// WorkspaceRoleUpdateActionCliHandler builds a full *cli.Command for the
// WorkspaceRoleUpdateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a WorkspaceRoleUpdateActionRequest the same way
// WorkspaceRoleUpdateActionHandler (Gin) and WorkspaceRoleUpdateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func WorkspaceRoleUpdateActionCliHandler(
	handler func(c WorkspaceRoleUpdateActionRequest) (*WorkspaceRoleUpdateActionResponse, error),
) *cli.Command {
	meta := WorkspaceRoleUpdateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: WorkspaceRoleUpdateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := WorkspaceRoleUpdateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastWorkspaceRoleOptionalDtoFromCli(c),
			Params:      WorkspaceRoleUpdateActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// WorkspaceRoleUpdateActionCli is a high-level convenience wrapper around
// WorkspaceRoleUpdateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way WorkspaceRoleUpdateActionGin
// registers a route on a Gin engine.
func WorkspaceRoleUpdateActionCli(
	app *cli.Command,
	handler func(c WorkspaceRoleUpdateActionRequest) (*WorkspaceRoleUpdateActionResponse, error),
) {
	app.Commands = append(app.Commands, WorkspaceRoleUpdateActionCliHandler(handler))
}
