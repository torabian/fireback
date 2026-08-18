//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x WorkspaceTypeAwareDeleteActionRequest) IsCli() bool {
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

// WorkspaceTypeAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the WorkspaceTypeAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func WorkspaceTypeAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWorkspaceTypeAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// WorkspaceTypeAwareDeleteActionCliHandler builds a full *cli.Command for the
// WorkspaceTypeAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a WorkspaceTypeAwareDeleteActionRequest the same way
// WorkspaceTypeAwareDeleteActionHandler (Gin) and WorkspaceTypeAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func WorkspaceTypeAwareDeleteActionCliHandler(
	handler func(c WorkspaceTypeAwareDeleteActionRequest) (*WorkspaceTypeAwareDeleteActionResponse, error),
) *cli.Command {
	meta := WorkspaceTypeAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: WorkspaceTypeAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := WorkspaceTypeAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastWorkspaceTypeAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// WorkspaceTypeAwareDeleteActionCli is a high-level convenience wrapper around
// WorkspaceTypeAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way WorkspaceTypeAwareDeleteActionGin
// registers a route on a Gin engine.
func WorkspaceTypeAwareDeleteActionCli(
	app *cli.Command,
	handler func(c WorkspaceTypeAwareDeleteActionRequest) (*WorkspaceTypeAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, WorkspaceTypeAwareDeleteActionCliHandler(handler))
}
