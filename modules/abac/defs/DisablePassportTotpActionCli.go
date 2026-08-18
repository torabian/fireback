//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x DisablePassportTotpActionRequest) IsCli() bool {
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

// DisablePassportTotpActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the DisablePassportTotpAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func DisablePassportTotpActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetDisablePassportTotpActionReqCliFlags(""))...)
	return flags
}

// DisablePassportTotpActionCliHandler builds a full *cli.Command for the
// DisablePassportTotpAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a DisablePassportTotpActionRequest the same way
// DisablePassportTotpActionHandler (Gin) and DisablePassportTotpActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func DisablePassportTotpActionCliHandler(
	handler func(c DisablePassportTotpActionRequest) (*DisablePassportTotpActionResponse, error),
) *cli.Command {
	meta := DisablePassportTotpActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: DisablePassportTotpActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := DisablePassportTotpActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastDisablePassportTotpActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// DisablePassportTotpActionCli is a high-level convenience wrapper around
// DisablePassportTotpActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way DisablePassportTotpActionGin
// registers a route on a Gin engine.
func DisablePassportTotpActionCli(
	app *cli.Command,
	handler func(c DisablePassportTotpActionRequest) (*DisablePassportTotpActionResponse, error),
) {
	app.Commands = append(app.Commands, DisablePassportTotpActionCliHandler(handler))
}
