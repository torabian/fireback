//go:build !wasm

package abac

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x ChangePasswordActionRequest) IsCli() bool {
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

// ChangePasswordActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the ChangePasswordAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func ChangePasswordActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetChangePasswordActionReqCliFlags(""))...)
	return flags
}

// ChangePasswordActionCliHandler builds a full *cli.Command for the
// ChangePasswordAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a ChangePasswordActionRequest the same way
// ChangePasswordActionHandler (Gin) and ChangePasswordActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func ChangePasswordActionCliHandler(
	handler func(c ChangePasswordActionRequest) (*ChangePasswordActionResponse, error),
) *cli.Command {
	meta := ChangePasswordActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: ChangePasswordActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := ChangePasswordActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastChangePasswordActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// ChangePasswordActionCli is a high-level convenience wrapper around
// ChangePasswordActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way ChangePasswordActionGin
// registers a route on a Gin engine.
func ChangePasswordActionCli(
	app *cli.Command,
	handler func(c ChangePasswordActionRequest) (*ChangePasswordActionResponse, error),
) {
	app.Commands = append(app.Commands, ChangePasswordActionCliHandler(handler))
}
