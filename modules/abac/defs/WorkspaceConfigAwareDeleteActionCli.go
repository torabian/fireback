//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x WorkspaceConfigAwareDeleteActionRequest) IsCli() bool {
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

// WorkspaceConfigAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the WorkspaceConfigAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func WorkspaceConfigAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWorkspaceConfigAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// WorkspaceConfigAwareDeleteActionCliHandler builds a full *cli.Command for the
// WorkspaceConfigAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a WorkspaceConfigAwareDeleteActionRequest the same way
// WorkspaceConfigAwareDeleteActionHandler (Gin) and WorkspaceConfigAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func WorkspaceConfigAwareDeleteActionCliHandler(
	handler func(c WorkspaceConfigAwareDeleteActionRequest) (*WorkspaceConfigAwareDeleteActionResponse, error),
) *cli.Command {
	meta := WorkspaceConfigAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: WorkspaceConfigAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := WorkspaceConfigAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastWorkspaceConfigAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// WorkspaceConfigAwareDeleteActionCli is a high-level convenience wrapper around
// WorkspaceConfigAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way WorkspaceConfigAwareDeleteActionGin
// registers a route on a Gin engine.
func WorkspaceConfigAwareDeleteActionCli(
	app *cli.Command,
	handler func(c WorkspaceConfigAwareDeleteActionRequest) (*WorkspaceConfigAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, WorkspaceConfigAwareDeleteActionCliHandler(handler))
}
