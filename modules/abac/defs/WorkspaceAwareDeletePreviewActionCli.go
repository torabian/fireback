//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetWorkspaceAwareDeletePreviewActionQueryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "qs-unique-ids",
			Type: "slice",
		},
	}
}

// WorkspaceAwareDeletePreviewActionQueryFromCli extracts and casts query parameters the same way
// WorkspaceAwareDeletePreviewActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func WorkspaceAwareDeletePreviewActionQueryFromCli(c *cli.Command) WorkspaceAwareDeletePreviewActionQuery {
	data := WorkspaceAwareDeletePreviewActionQuery{}
	values := url.Values{}
	if c.IsSet("qs-unique-ids") {
		raw := c.String("qs-unique-ids")
		emigo.InflatePossibleSlice(raw, &data.UniqueIds)
		values.Set("uniqueIds", raw)
	}
	data.SetValues(values)
	return data
}
func (x WorkspaceAwareDeletePreviewActionRequest) IsCli() bool {
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

// WorkspaceAwareDeletePreviewActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the WorkspaceAwareDeletePreviewAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func WorkspaceAwareDeletePreviewActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWorkspaceAwareDeletePreviewActionQueryCliFlags(""))...)
	return flags
}

// WorkspaceAwareDeletePreviewActionCliHandler builds a full *cli.Command for the
// WorkspaceAwareDeletePreviewAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a WorkspaceAwareDeletePreviewActionRequest the same way
// WorkspaceAwareDeletePreviewActionHandler (Gin) and WorkspaceAwareDeletePreviewActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func WorkspaceAwareDeletePreviewActionCliHandler(
	handler func(c WorkspaceAwareDeletePreviewActionRequest) (*WorkspaceAwareDeletePreviewActionResponse, error),
) *cli.Command {
	meta := WorkspaceAwareDeletePreviewActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: WorkspaceAwareDeletePreviewActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := WorkspaceAwareDeletePreviewActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		req.QueryParams = WorkspaceAwareDeletePreviewActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// WorkspaceAwareDeletePreviewActionCli is a high-level convenience wrapper around
// WorkspaceAwareDeletePreviewActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way WorkspaceAwareDeletePreviewActionGin
// registers a route on a Gin engine.
func WorkspaceAwareDeletePreviewActionCli(
	app *cli.Command,
	handler func(c WorkspaceAwareDeletePreviewActionRequest) (*WorkspaceAwareDeletePreviewActionResponse, error),
) {
	app.Commands = append(app.Commands, WorkspaceAwareDeletePreviewActionCliHandler(handler))
}
