//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetWorkspaceTypeUpdateActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func WorkspaceTypeUpdateActionPathParameterFromCli(c *cli.Command) WorkspaceTypeUpdateActionPathParameter {
	return WorkspaceTypeUpdateActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x WorkspaceTypeUpdateActionRequest) IsCli() bool {
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

// WorkspaceTypeUpdateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the WorkspaceTypeUpdateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func WorkspaceTypeUpdateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWorkspaceTypeOptionalDtoCliFlags(""))...)
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWorkspaceTypeUpdateActionPathParameterCliFlags(""))...)
	return flags
}

// WorkspaceTypeUpdateActionCliHandler builds a full *cli.Command for the
// WorkspaceTypeUpdateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a WorkspaceTypeUpdateActionRequest the same way
// WorkspaceTypeUpdateActionHandler (Gin) and WorkspaceTypeUpdateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func WorkspaceTypeUpdateActionCliHandler(
	handler func(c WorkspaceTypeUpdateActionRequest) (*WorkspaceTypeUpdateActionResponse, error),
) *cli.Command {
	meta := WorkspaceTypeUpdateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: WorkspaceTypeUpdateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := WorkspaceTypeUpdateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastWorkspaceTypeOptionalDtoFromCli(c),
			Params:      WorkspaceTypeUpdateActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// WorkspaceTypeUpdateActionCli is a high-level convenience wrapper around
// WorkspaceTypeUpdateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way WorkspaceTypeUpdateActionGin
// registers a route on a Gin engine.
func WorkspaceTypeUpdateActionCli(
	app *cli.Command,
	handler func(c WorkspaceTypeUpdateActionRequest) (*WorkspaceTypeUpdateActionResponse, error),
) {
	app.Commands = append(app.Commands, WorkspaceTypeUpdateActionCliHandler(handler))
}
