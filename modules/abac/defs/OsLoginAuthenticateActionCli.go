//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x OsLoginAuthenticateActionRequest) IsCli() bool {
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

// OsLoginAuthenticateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the OsLoginAuthenticateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func OsLoginAuthenticateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	return flags
}

// OsLoginAuthenticateActionCliHandler builds a full *cli.Command for the
// OsLoginAuthenticateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a OsLoginAuthenticateActionRequest the same way
// OsLoginAuthenticateActionHandler (Gin) and OsLoginAuthenticateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func OsLoginAuthenticateActionCliHandler(
	handler func(c OsLoginAuthenticateActionRequest) (*OsLoginAuthenticateActionResponse, error),
) *cli.Command {
	meta := OsLoginAuthenticateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: OsLoginAuthenticateActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := OsLoginAuthenticateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// OsLoginAuthenticateActionCli is a high-level convenience wrapper around
// OsLoginAuthenticateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way OsLoginAuthenticateActionGin
// registers a route on a Gin engine.
func OsLoginAuthenticateActionCli(
	app *cli.Command,
	handler func(c OsLoginAuthenticateActionRequest) (*OsLoginAuthenticateActionResponse, error),
) {
	app.Commands = append(app.Commands, OsLoginAuthenticateActionCliHandler(handler))
}
