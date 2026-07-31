//go:build !wasm

package abac

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x ClassicPassportRequestOtpActionRequest) IsCli() bool {
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

// ClassicPassportRequestOtpActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the ClassicPassportRequestOtpAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func ClassicPassportRequestOtpActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetClassicPassportRequestOtpActionReqCliFlags(""))...)
	return flags
}

// ClassicPassportRequestOtpActionCliHandler builds a full *cli.Command for the
// ClassicPassportRequestOtpAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a ClassicPassportRequestOtpActionRequest the same way
// ClassicPassportRequestOtpActionHandler (Gin) and ClassicPassportRequestOtpActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func ClassicPassportRequestOtpActionCliHandler(
	handler func(c ClassicPassportRequestOtpActionRequest) (*ClassicPassportRequestOtpActionResponse, error),
) *cli.Command {
	meta := ClassicPassportRequestOtpActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: ClassicPassportRequestOtpActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := ClassicPassportRequestOtpActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastClassicPassportRequestOtpActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// ClassicPassportRequestOtpActionCli is a high-level convenience wrapper around
// ClassicPassportRequestOtpActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way ClassicPassportRequestOtpActionGin
// registers a route on a Gin engine.
func ClassicPassportRequestOtpActionCli(
	app *cli.Command,
	handler func(c ClassicPassportRequestOtpActionRequest) (*ClassicPassportRequestOtpActionResponse, error),
) {
	app.Commands = append(app.Commands, ClassicPassportRequestOtpActionCliHandler(handler))
}
