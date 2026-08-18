//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x ClassicPassportOtpActionRequest) IsCli() bool {
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

// ClassicPassportOtpActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the ClassicPassportOtpAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func ClassicPassportOtpActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetClassicPassportOtpActionReqCliFlags(""))...)
	return flags
}

// ClassicPassportOtpActionCliHandler builds a full *cli.Command for the
// ClassicPassportOtpAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a ClassicPassportOtpActionRequest the same way
// ClassicPassportOtpActionHandler (Gin) and ClassicPassportOtpActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func ClassicPassportOtpActionCliHandler(
	handler func(c ClassicPassportOtpActionRequest) (*ClassicPassportOtpActionResponse, error),
) *cli.Command {
	meta := ClassicPassportOtpActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: ClassicPassportOtpActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := ClassicPassportOtpActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastClassicPassportOtpActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// ClassicPassportOtpActionCli is a high-level convenience wrapper around
// ClassicPassportOtpActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way ClassicPassportOtpActionGin
// registers a route on a Gin engine.
func ClassicPassportOtpActionCli(
	app *cli.Command,
	handler func(c ClassicPassportOtpActionRequest) (*ClassicPassportOtpActionResponse, error),
) {
	app.Commands = append(app.Commands, ClassicPassportOtpActionCliHandler(handler))
}
