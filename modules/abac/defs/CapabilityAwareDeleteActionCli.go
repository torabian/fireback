//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x CapabilityAwareDeleteActionRequest) IsCli() bool {
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

// CapabilityAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the CapabilityAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func CapabilityAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetCapabilityAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// CapabilityAwareDeleteActionCliHandler builds a full *cli.Command for the
// CapabilityAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a CapabilityAwareDeleteActionRequest the same way
// CapabilityAwareDeleteActionHandler (Gin) and CapabilityAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func CapabilityAwareDeleteActionCliHandler(
	handler func(c CapabilityAwareDeleteActionRequest) (*CapabilityAwareDeleteActionResponse, error),
) *cli.Command {
	meta := CapabilityAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: CapabilityAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := CapabilityAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastCapabilityAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// CapabilityAwareDeleteActionCli is a high-level convenience wrapper around
// CapabilityAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way CapabilityAwareDeleteActionGin
// registers a route on a Gin engine.
func CapabilityAwareDeleteActionCli(
	app *cli.Command,
	handler func(c CapabilityAwareDeleteActionRequest) (*CapabilityAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, CapabilityAwareDeleteActionCliHandler(handler))
}
