//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x WorkspaceConfigCreateActionRequest) IsCli() bool {
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

// WorkspaceConfigCreateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the WorkspaceConfigCreateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func WorkspaceConfigCreateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWorkspaceConfigDtoCliFlags(""))...)
	return flags
}

// WorkspaceConfigCreateActionCliHandler builds a full *cli.Command for the
// WorkspaceConfigCreateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a WorkspaceConfigCreateActionRequest the same way
// WorkspaceConfigCreateActionHandler (Gin) and WorkspaceConfigCreateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func WorkspaceConfigCreateActionCliHandler(
	handler func(c WorkspaceConfigCreateActionRequest) (*WorkspaceConfigCreateActionResponse, error),
) *cli.Command {
	meta := WorkspaceConfigCreateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: WorkspaceConfigCreateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := WorkspaceConfigCreateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastWorkspaceConfigDtoFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// WorkspaceConfigCreateActionCli is a high-level convenience wrapper around
// WorkspaceConfigCreateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way WorkspaceConfigCreateActionGin
// registers a route on a Gin engine.
func WorkspaceConfigCreateActionCli(
	app *cli.Command,
	handler func(c WorkspaceConfigCreateActionRequest) (*WorkspaceConfigCreateActionResponse, error),
) {
	app.Commands = append(app.Commands, WorkspaceConfigCreateActionCliHandler(handler))
}
