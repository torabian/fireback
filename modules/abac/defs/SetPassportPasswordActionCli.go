//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x SetPassportPasswordActionRequest) IsCli() bool {
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

// SetPassportPasswordActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the SetPassportPasswordAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func SetPassportPasswordActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetSetPassportPasswordActionReqCliFlags(""))...)
	return flags
}

// SetPassportPasswordActionCliHandler builds a full *cli.Command for the
// SetPassportPasswordAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a SetPassportPasswordActionRequest the same way
// SetPassportPasswordActionHandler (Gin) and SetPassportPasswordActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func SetPassportPasswordActionCliHandler(
	handler func(c SetPassportPasswordActionRequest) (*SetPassportPasswordActionResponse, error),
) *cli.Command {
	meta := SetPassportPasswordActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: SetPassportPasswordActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := SetPassportPasswordActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastSetPassportPasswordActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// SetPassportPasswordActionCli is a high-level convenience wrapper around
// SetPassportPasswordActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way SetPassportPasswordActionGin
// registers a route on a Gin engine.
func SetPassportPasswordActionCli(
	app *cli.Command,
	handler func(c SetPassportPasswordActionRequest) (*SetPassportPasswordActionResponse, error),
) {
	app.Commands = append(app.Commands, SetPassportPasswordActionCliHandler(handler))
}
