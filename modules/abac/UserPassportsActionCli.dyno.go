//go:build !wasm

package abac

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x UserPassportsActionRequest) IsCli() bool {
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

// UserPassportsActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the UserPassportsAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func UserPassportsActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	return flags
}

// UserPassportsActionCliHandler builds a full *cli.Command for the
// UserPassportsAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a UserPassportsActionRequest the same way
// UserPassportsActionHandler (Gin) and UserPassportsActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func UserPassportsActionCliHandler(
	handler func(c UserPassportsActionRequest) (*UserPassportsActionResponse, error),
) *cli.Command {
	meta := UserPassportsActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: UserPassportsActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := UserPassportsActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// UserPassportsActionCli is a high-level convenience wrapper around
// UserPassportsActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way UserPassportsActionGin
// registers a route on a Gin engine.
func UserPassportsActionCli(
	app *cli.Command,
	handler func(c UserPassportsActionRequest) (*UserPassportsActionResponse, error),
) {
	app.Commands = append(app.Commands, UserPassportsActionCliHandler(handler))
}
