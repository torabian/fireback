//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x PreferenceAwareDeleteActionRequest) IsCli() bool {
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

// PreferenceAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PreferenceAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PreferenceAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPreferenceAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// PreferenceAwareDeleteActionCliHandler builds a full *cli.Command for the
// PreferenceAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PreferenceAwareDeleteActionRequest the same way
// PreferenceAwareDeleteActionHandler (Gin) and PreferenceAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PreferenceAwareDeleteActionCliHandler(
	handler func(c PreferenceAwareDeleteActionRequest) (*PreferenceAwareDeleteActionResponse, error),
) *cli.Command {
	meta := PreferenceAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PreferenceAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PreferenceAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastPreferenceAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PreferenceAwareDeleteActionCli is a high-level convenience wrapper around
// PreferenceAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PreferenceAwareDeleteActionGin
// registers a route on a Gin engine.
func PreferenceAwareDeleteActionCli(
	app *cli.Command,
	handler func(c PreferenceAwareDeleteActionRequest) (*PreferenceAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, PreferenceAwareDeleteActionCliHandler(handler))
}
