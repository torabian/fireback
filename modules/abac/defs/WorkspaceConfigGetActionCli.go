//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetWorkspaceConfigGetActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func WorkspaceConfigGetActionPathParameterFromCli(c *cli.Command) WorkspaceConfigGetActionPathParameter {
	return WorkspaceConfigGetActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x WorkspaceConfigGetActionRequest) IsCli() bool {
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

// WorkspaceConfigGetActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the WorkspaceConfigGetAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func WorkspaceConfigGetActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWorkspaceConfigGetActionPathParameterCliFlags(""))...)
	return flags
}

// WorkspaceConfigGetActionCliHandler builds a full *cli.Command for the
// WorkspaceConfigGetAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a WorkspaceConfigGetActionRequest the same way
// WorkspaceConfigGetActionHandler (Gin) and WorkspaceConfigGetActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func WorkspaceConfigGetActionCliHandler(
	handler func(c WorkspaceConfigGetActionRequest) (*WorkspaceConfigGetActionResponse, error),
) *cli.Command {
	meta := WorkspaceConfigGetActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: WorkspaceConfigGetActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := WorkspaceConfigGetActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Params:      WorkspaceConfigGetActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// WorkspaceConfigGetActionCli is a high-level convenience wrapper around
// WorkspaceConfigGetActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way WorkspaceConfigGetActionGin
// registers a route on a Gin engine.
func WorkspaceConfigGetActionCli(
	app *cli.Command,
	handler func(c WorkspaceConfigGetActionRequest) (*WorkspaceConfigGetActionResponse, error),
) {
	app.Commands = append(app.Commands, WorkspaceConfigGetActionCliHandler(handler))
}
