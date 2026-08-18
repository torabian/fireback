//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x CapabilitiesTreeActionRequest) IsCli() bool {
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

// CapabilitiesTreeActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the CapabilitiesTreeAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func CapabilitiesTreeActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	return flags
}

// CapabilitiesTreeActionCliHandler builds a full *cli.Command for the
// CapabilitiesTreeAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a CapabilitiesTreeActionRequest the same way
// CapabilitiesTreeActionHandler (Gin) and CapabilitiesTreeActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func CapabilitiesTreeActionCliHandler(
	handler func(c CapabilitiesTreeActionRequest) (*CapabilitiesTreeActionResponse, error),
) *cli.Command {
	meta := CapabilitiesTreeActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: CapabilitiesTreeActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := CapabilitiesTreeActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// CapabilitiesTreeActionCli is a high-level convenience wrapper around
// CapabilitiesTreeActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way CapabilitiesTreeActionGin
// registers a route on a Gin engine.
func CapabilitiesTreeActionCli(
	app *cli.Command,
	handler func(c CapabilitiesTreeActionRequest) (*CapabilitiesTreeActionResponse, error),
) {
	app.Commands = append(app.Commands, CapabilitiesTreeActionCliHandler(handler))
}
