//go:build !wasm

package fireback

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x GetCapabilitiesActionRequest) IsCli() bool {
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

// GetCapabilitiesActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the GetCapabilitiesAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func GetCapabilitiesActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	return flags
}

// GetCapabilitiesActionCliHandler builds a full *cli.Command for the
// GetCapabilitiesAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a GetCapabilitiesActionRequest the same way
// GetCapabilitiesActionHandler (Gin) and GetCapabilitiesActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func GetCapabilitiesActionCliHandler(
	handler func(c GetCapabilitiesActionRequest) (*GetCapabilitiesActionResponse, error),
) *cli.Command {
	meta := GetCapabilitiesActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: GetCapabilitiesActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := GetCapabilitiesActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// GetCapabilitiesActionCli is a high-level convenience wrapper around
// GetCapabilitiesActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way GetCapabilitiesActionGin
// registers a route on a Gin engine.
func GetCapabilitiesActionCli(
	app *cli.Command,
	handler func(c GetCapabilitiesActionRequest) (*GetCapabilitiesActionResponse, error),
) {
	app.Commands = append(app.Commands, GetCapabilitiesActionCliHandler(handler))
}
