//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x UserAwareDeleteActionRequest) IsCli() bool {
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

// UserAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the UserAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func UserAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetUserAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// UserAwareDeleteActionCliHandler builds a full *cli.Command for the
// UserAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a UserAwareDeleteActionRequest the same way
// UserAwareDeleteActionHandler (Gin) and UserAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func UserAwareDeleteActionCliHandler(
	handler func(c UserAwareDeleteActionRequest) (*UserAwareDeleteActionResponse, error),
) *cli.Command {
	meta := UserAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: UserAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := UserAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastUserAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// UserAwareDeleteActionCli is a high-level convenience wrapper around
// UserAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way UserAwareDeleteActionGin
// registers a route on a Gin engine.
func UserAwareDeleteActionCli(
	app *cli.Command,
	handler func(c UserAwareDeleteActionRequest) (*UserAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, UserAwareDeleteActionCliHandler(handler))
}
