//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetWorkspaceConfigUpdateActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func WorkspaceConfigUpdateActionPathParameterFromCli(c *cli.Command) WorkspaceConfigUpdateActionPathParameter {
	return WorkspaceConfigUpdateActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x WorkspaceConfigUpdateActionRequest) IsCli() bool {
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

// WorkspaceConfigUpdateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the WorkspaceConfigUpdateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func WorkspaceConfigUpdateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWorkspaceConfigOptionalDtoCliFlags(""))...)
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWorkspaceConfigUpdateActionPathParameterCliFlags(""))...)
	return flags
}

// WorkspaceConfigUpdateActionCliHandler builds a full *cli.Command for the
// WorkspaceConfigUpdateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a WorkspaceConfigUpdateActionRequest the same way
// WorkspaceConfigUpdateActionHandler (Gin) and WorkspaceConfigUpdateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func WorkspaceConfigUpdateActionCliHandler(
	handler func(c WorkspaceConfigUpdateActionRequest) (*WorkspaceConfigUpdateActionResponse, error),
) *cli.Command {
	meta := WorkspaceConfigUpdateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: WorkspaceConfigUpdateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := WorkspaceConfigUpdateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastWorkspaceConfigOptionalDtoFromCli(c),
			Params:      WorkspaceConfigUpdateActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// WorkspaceConfigUpdateActionCli is a high-level convenience wrapper around
// WorkspaceConfigUpdateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way WorkspaceConfigUpdateActionGin
// registers a route on a Gin engine.
func WorkspaceConfigUpdateActionCli(
	app *cli.Command,
	handler func(c WorkspaceConfigUpdateActionRequest) (*WorkspaceConfigUpdateActionResponse, error),
) {
	app.Commands = append(app.Commands, WorkspaceConfigUpdateActionCliHandler(handler))
}
