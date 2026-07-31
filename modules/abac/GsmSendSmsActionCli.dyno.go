//go:build !wasm

package abac

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x GsmSendSmsActionRequest) IsCli() bool {
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

// GsmSendSmsActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the GsmSendSmsAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func GsmSendSmsActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetGsmSendSmsActionReqCliFlags(""))...)
	return flags
}

// GsmSendSmsActionCliHandler builds a full *cli.Command for the
// GsmSendSmsAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a GsmSendSmsActionRequest the same way
// GsmSendSmsActionHandler (Gin) and GsmSendSmsActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func GsmSendSmsActionCliHandler(
	handler func(c GsmSendSmsActionRequest) (*GsmSendSmsActionResponse, error),
) *cli.Command {
	meta := GsmSendSmsActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: GsmSendSmsActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := GsmSendSmsActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastGsmSendSmsActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// GsmSendSmsActionCli is a high-level convenience wrapper around
// GsmSendSmsActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way GsmSendSmsActionGin
// registers a route on a Gin engine.
func GsmSendSmsActionCli(
	app *cli.Command,
	handler func(c GsmSendSmsActionRequest) (*GsmSendSmsActionResponse, error),
) {
	app.Commands = append(app.Commands, GsmSendSmsActionCliHandler(handler))
}
