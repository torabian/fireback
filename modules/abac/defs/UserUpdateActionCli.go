//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetUserUpdateActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func UserUpdateActionPathParameterFromCli(c *cli.Command) UserUpdateActionPathParameter {
	return UserUpdateActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x UserUpdateActionRequest) IsCli() bool {
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

// UserUpdateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the UserUpdateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func UserUpdateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetUserOptionalDtoCliFlags(""))...)
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetUserUpdateActionPathParameterCliFlags(""))...)
	return flags
}

// UserUpdateActionCliHandler builds a full *cli.Command for the
// UserUpdateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a UserUpdateActionRequest the same way
// UserUpdateActionHandler (Gin) and UserUpdateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func UserUpdateActionCliHandler(
	handler func(c UserUpdateActionRequest) (*UserUpdateActionResponse, error),
) *cli.Command {
	meta := UserUpdateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: UserUpdateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := UserUpdateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastUserOptionalDtoFromCli(c),
			Params:      UserUpdateActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// UserUpdateActionCli is a high-level convenience wrapper around
// UserUpdateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way UserUpdateActionGin
// registers a route on a Gin engine.
func UserUpdateActionCli(
	app *cli.Command,
	handler func(c UserUpdateActionRequest) (*UserUpdateActionResponse, error),
) {
	app.Commands = append(app.Commands, UserUpdateActionCliHandler(handler))
}
