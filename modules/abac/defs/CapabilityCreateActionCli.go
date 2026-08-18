//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x CapabilityCreateActionRequest) IsCli() bool {
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

// CapabilityCreateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the CapabilityCreateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func CapabilityCreateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetCapabilityDtoCliFlags(""))...)
	return flags
}

// CapabilityCreateActionCliHandler builds a full *cli.Command for the
// CapabilityCreateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a CapabilityCreateActionRequest the same way
// CapabilityCreateActionHandler (Gin) and CapabilityCreateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func CapabilityCreateActionCliHandler(
	handler func(c CapabilityCreateActionRequest) (*CapabilityCreateActionResponse, error),
) *cli.Command {
	meta := CapabilityCreateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: CapabilityCreateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := CapabilityCreateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastCapabilityDtoFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// CapabilityCreateActionCli is a high-level convenience wrapper around
// CapabilityCreateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way CapabilityCreateActionGin
// registers a route on a Gin engine.
func CapabilityCreateActionCli(
	app *cli.Command,
	handler func(c CapabilityCreateActionRequest) (*CapabilityCreateActionResponse, error),
) {
	app.Commands = append(app.Commands, CapabilityCreateActionCliHandler(handler))
}
