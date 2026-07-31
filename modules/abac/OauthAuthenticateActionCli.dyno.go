//go:build !wasm

package abac

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x OauthAuthenticateActionRequest) IsCli() bool {
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

// OauthAuthenticateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the OauthAuthenticateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func OauthAuthenticateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetOauthAuthenticateActionReqCliFlags(""))...)
	return flags
}

// OauthAuthenticateActionCliHandler builds a full *cli.Command for the
// OauthAuthenticateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a OauthAuthenticateActionRequest the same way
// OauthAuthenticateActionHandler (Gin) and OauthAuthenticateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func OauthAuthenticateActionCliHandler(
	handler func(c OauthAuthenticateActionRequest) (*OauthAuthenticateActionResponse, error),
) *cli.Command {
	meta := OauthAuthenticateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: OauthAuthenticateActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := OauthAuthenticateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastOauthAuthenticateActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// OauthAuthenticateActionCli is a high-level convenience wrapper around
// OauthAuthenticateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way OauthAuthenticateActionGin
// registers a route on a Gin engine.
func OauthAuthenticateActionCli(
	app *cli.Command,
	handler func(c OauthAuthenticateActionRequest) (*OauthAuthenticateActionResponse, error),
) {
	app.Commands = append(app.Commands, OauthAuthenticateActionCliHandler(handler))
}
