//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetUserProfileGetActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func UserProfileGetActionPathParameterFromCli(c *cli.Command) UserProfileGetActionPathParameter {
	return UserProfileGetActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x UserProfileGetActionRequest) IsCli() bool {
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

// UserProfileGetActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the UserProfileGetAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func UserProfileGetActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetUserProfileGetActionPathParameterCliFlags(""))...)
	return flags
}

// UserProfileGetActionCliHandler builds a full *cli.Command for the
// UserProfileGetAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a UserProfileGetActionRequest the same way
// UserProfileGetActionHandler (Gin) and UserProfileGetActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func UserProfileGetActionCliHandler(
	handler func(c UserProfileGetActionRequest) (*UserProfileGetActionResponse, error),
) *cli.Command {
	meta := UserProfileGetActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: UserProfileGetActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := UserProfileGetActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Params:      UserProfileGetActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// UserProfileGetActionCli is a high-level convenience wrapper around
// UserProfileGetActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way UserProfileGetActionGin
// registers a route on a Gin engine.
func UserProfileGetActionCli(
	app *cli.Command,
	handler func(c UserProfileGetActionRequest) (*UserProfileGetActionResponse, error),
) {
	app.Commands = append(app.Commands, UserProfileGetActionCliHandler(handler))
}
