//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x RegionalContentAwareDeleteActionRequest) IsCli() bool {
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

// RegionalContentAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the RegionalContentAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func RegionalContentAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetRegionalContentAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// RegionalContentAwareDeleteActionCliHandler builds a full *cli.Command for the
// RegionalContentAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a RegionalContentAwareDeleteActionRequest the same way
// RegionalContentAwareDeleteActionHandler (Gin) and RegionalContentAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func RegionalContentAwareDeleteActionCliHandler(
	handler func(c RegionalContentAwareDeleteActionRequest) (*RegionalContentAwareDeleteActionResponse, error),
) *cli.Command {
	meta := RegionalContentAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: RegionalContentAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := RegionalContentAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastRegionalContentAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// RegionalContentAwareDeleteActionCli is a high-level convenience wrapper around
// RegionalContentAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way RegionalContentAwareDeleteActionGin
// registers a route on a Gin engine.
func RegionalContentAwareDeleteActionCli(
	app *cli.Command,
	handler func(c RegionalContentAwareDeleteActionRequest) (*RegionalContentAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, RegionalContentAwareDeleteActionCliHandler(handler))
}
