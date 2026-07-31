//go:build !wasm

package abac

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x ConfirmClassicPassportTotpActionRequest) IsCli() bool {
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

// ConfirmClassicPassportTotpActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the ConfirmClassicPassportTotpAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func ConfirmClassicPassportTotpActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetConfirmClassicPassportTotpActionReqCliFlags(""))...)
	return flags
}

// ConfirmClassicPassportTotpActionCliHandler builds a full *cli.Command for the
// ConfirmClassicPassportTotpAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a ConfirmClassicPassportTotpActionRequest the same way
// ConfirmClassicPassportTotpActionHandler (Gin) and ConfirmClassicPassportTotpActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func ConfirmClassicPassportTotpActionCliHandler(
	handler func(c ConfirmClassicPassportTotpActionRequest) (*ConfirmClassicPassportTotpActionResponse, error),
) *cli.Command {
	meta := ConfirmClassicPassportTotpActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: ConfirmClassicPassportTotpActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := ConfirmClassicPassportTotpActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastConfirmClassicPassportTotpActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// ConfirmClassicPassportTotpActionCli is a high-level convenience wrapper around
// ConfirmClassicPassportTotpActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way ConfirmClassicPassportTotpActionGin
// registers a route on a Gin engine.
func ConfirmClassicPassportTotpActionCli(
	app *cli.Command,
	handler func(c ConfirmClassicPassportTotpActionRequest) (*ConfirmClassicPassportTotpActionResponse, error),
) {
	app.Commands = append(app.Commands, ConfirmClassicPassportTotpActionCliHandler(handler))
}
