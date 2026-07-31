//go:build !wasm

package payment

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x CheckProductInventoryActionRequest) IsCli() bool {
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

// CheckProductInventoryActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the CheckProductInventoryAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func CheckProductInventoryActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	return flags
}

// CheckProductInventoryActionCliHandler builds a full *cli.Command for the
// CheckProductInventoryAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a CheckProductInventoryActionRequest the same way
// CheckProductInventoryActionHandler (Gin) and CheckProductInventoryActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func CheckProductInventoryActionCliHandler(
	handler func(c CheckProductInventoryActionRequest) (*CheckProductInventoryActionResponse, error),
) *cli.Command {
	meta := CheckProductInventoryActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: CheckProductInventoryActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := CheckProductInventoryActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// CheckProductInventoryActionCli is a high-level convenience wrapper around
// CheckProductInventoryActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way CheckProductInventoryActionGin
// registers a route on a Gin engine.
func CheckProductInventoryActionCli(
	app *cli.Command,
	handler func(c CheckProductInventoryActionRequest) (*CheckProductInventoryActionResponse, error),
) {
	app.Commands = append(app.Commands, CheckProductInventoryActionCliHandler(handler))
}
