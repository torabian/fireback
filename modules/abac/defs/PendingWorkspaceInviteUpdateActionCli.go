//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetPendingWorkspaceInviteUpdateActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func PendingWorkspaceInviteUpdateActionPathParameterFromCli(c *cli.Command) PendingWorkspaceInviteUpdateActionPathParameter {
	return PendingWorkspaceInviteUpdateActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x PendingWorkspaceInviteUpdateActionRequest) IsCli() bool {
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

// PendingWorkspaceInviteUpdateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PendingWorkspaceInviteUpdateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PendingWorkspaceInviteUpdateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPendingWorkspaceInviteOptionalDtoCliFlags(""))...)
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPendingWorkspaceInviteUpdateActionPathParameterCliFlags(""))...)
	return flags
}

// PendingWorkspaceInviteUpdateActionCliHandler builds a full *cli.Command for the
// PendingWorkspaceInviteUpdateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PendingWorkspaceInviteUpdateActionRequest the same way
// PendingWorkspaceInviteUpdateActionHandler (Gin) and PendingWorkspaceInviteUpdateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PendingWorkspaceInviteUpdateActionCliHandler(
	handler func(c PendingWorkspaceInviteUpdateActionRequest) (*PendingWorkspaceInviteUpdateActionResponse, error),
) *cli.Command {
	meta := PendingWorkspaceInviteUpdateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PendingWorkspaceInviteUpdateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PendingWorkspaceInviteUpdateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastPendingWorkspaceInviteOptionalDtoFromCli(c),
			Params:      PendingWorkspaceInviteUpdateActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PendingWorkspaceInviteUpdateActionCli is a high-level convenience wrapper around
// PendingWorkspaceInviteUpdateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PendingWorkspaceInviteUpdateActionGin
// registers a route on a Gin engine.
func PendingWorkspaceInviteUpdateActionCli(
	app *cli.Command,
	handler func(c PendingWorkspaceInviteUpdateActionRequest) (*PendingWorkspaceInviteUpdateActionResponse, error),
) {
	app.Commands = append(app.Commands, PendingWorkspaceInviteUpdateActionCliHandler(handler))
}
