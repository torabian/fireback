//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x WorkspaceCreateActionRequest) IsCli() bool {
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

// WorkspaceCreateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the WorkspaceCreateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func WorkspaceCreateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWorkspaceDtoCliFlags(""))...)
	return flags
}

// WorkspaceCreateActionCliHandler builds a full *cli.Command for the
// WorkspaceCreateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a WorkspaceCreateActionRequest the same way
// WorkspaceCreateActionHandler (Gin) and WorkspaceCreateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func WorkspaceCreateActionCliHandler(
	handler func(c WorkspaceCreateActionRequest) (*WorkspaceCreateActionResponse, error),
) *cli.Command {
	meta := WorkspaceCreateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: WorkspaceCreateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := WorkspaceCreateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastWorkspaceDtoFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// WorkspaceCreateActionCli is a high-level convenience wrapper around
// WorkspaceCreateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way WorkspaceCreateActionGin
// registers a route on a Gin engine.
func WorkspaceCreateActionCli(
	app *cli.Command,
	handler func(c WorkspaceCreateActionRequest) (*WorkspaceCreateActionResponse, error),
) {
	app.Commands = append(app.Commands, WorkspaceCreateActionCliHandler(handler))
}
