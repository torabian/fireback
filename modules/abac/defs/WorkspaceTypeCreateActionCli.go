//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x WorkspaceTypeCreateActionRequest) IsCli() bool {
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

// WorkspaceTypeCreateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the WorkspaceTypeCreateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func WorkspaceTypeCreateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWorkspaceTypeDtoCliFlags(""))...)
	return flags
}

// WorkspaceTypeCreateActionCliHandler builds a full *cli.Command for the
// WorkspaceTypeCreateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a WorkspaceTypeCreateActionRequest the same way
// WorkspaceTypeCreateActionHandler (Gin) and WorkspaceTypeCreateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func WorkspaceTypeCreateActionCliHandler(
	handler func(c WorkspaceTypeCreateActionRequest) (*WorkspaceTypeCreateActionResponse, error),
) *cli.Command {
	meta := WorkspaceTypeCreateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: WorkspaceTypeCreateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := WorkspaceTypeCreateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastWorkspaceTypeDtoFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// WorkspaceTypeCreateActionCli is a high-level convenience wrapper around
// WorkspaceTypeCreateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way WorkspaceTypeCreateActionGin
// registers a route on a Gin engine.
func WorkspaceTypeCreateActionCli(
	app *cli.Command,
	handler func(c WorkspaceTypeCreateActionRequest) (*WorkspaceTypeCreateActionResponse, error),
) {
	app.Commands = append(app.Commands, WorkspaceTypeCreateActionCliHandler(handler))
}
