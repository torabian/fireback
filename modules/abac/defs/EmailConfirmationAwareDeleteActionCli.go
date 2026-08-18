//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x EmailConfirmationAwareDeleteActionRequest) IsCli() bool {
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

// EmailConfirmationAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the EmailConfirmationAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func EmailConfirmationAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetEmailConfirmationAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// EmailConfirmationAwareDeleteActionCliHandler builds a full *cli.Command for the
// EmailConfirmationAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a EmailConfirmationAwareDeleteActionRequest the same way
// EmailConfirmationAwareDeleteActionHandler (Gin) and EmailConfirmationAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func EmailConfirmationAwareDeleteActionCliHandler(
	handler func(c EmailConfirmationAwareDeleteActionRequest) (*EmailConfirmationAwareDeleteActionResponse, error),
) *cli.Command {
	meta := EmailConfirmationAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: EmailConfirmationAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := EmailConfirmationAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastEmailConfirmationAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// EmailConfirmationAwareDeleteActionCli is a high-level convenience wrapper around
// EmailConfirmationAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way EmailConfirmationAwareDeleteActionGin
// registers a route on a Gin engine.
func EmailConfirmationAwareDeleteActionCli(
	app *cli.Command,
	handler func(c EmailConfirmationAwareDeleteActionRequest) (*EmailConfirmationAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, EmailConfirmationAwareDeleteActionCliHandler(handler))
}
