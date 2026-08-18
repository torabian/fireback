//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x WorkspaceConfigDistinctUpdateActionRequest) IsCli() bool {
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

// WorkspaceConfigDistinctUpdateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the WorkspaceConfigDistinctUpdateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func WorkspaceConfigDistinctUpdateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWorkspaceConfigOptionalDtoCliFlags(""))...)
	return flags
}

// WorkspaceConfigDistinctUpdateActionCliHandler builds a full *cli.Command for the
// WorkspaceConfigDistinctUpdateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a WorkspaceConfigDistinctUpdateActionRequest the same way
// WorkspaceConfigDistinctUpdateActionHandler (Gin) and WorkspaceConfigDistinctUpdateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func WorkspaceConfigDistinctUpdateActionCliHandler(
	handler func(c WorkspaceConfigDistinctUpdateActionRequest) (*WorkspaceConfigDistinctUpdateActionResponse, error),
) *cli.Command {
	meta := WorkspaceConfigDistinctUpdateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: WorkspaceConfigDistinctUpdateActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := WorkspaceConfigDistinctUpdateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastWorkspaceConfigOptionalDtoFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// WorkspaceConfigDistinctUpdateActionCli is a high-level convenience wrapper around
// WorkspaceConfigDistinctUpdateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way WorkspaceConfigDistinctUpdateActionGin
// registers a route on a Gin engine.
func WorkspaceConfigDistinctUpdateActionCli(
	app *cli.Command,
	handler func(c WorkspaceConfigDistinctUpdateActionRequest) (*WorkspaceConfigDistinctUpdateActionResponse, error),
) {
	app.Commands = append(app.Commands, WorkspaceConfigDistinctUpdateActionCliHandler(handler))
}
