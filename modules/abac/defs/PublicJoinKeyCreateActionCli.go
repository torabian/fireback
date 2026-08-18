//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x PublicJoinKeyCreateActionRequest) IsCli() bool {
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

// PublicJoinKeyCreateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PublicJoinKeyCreateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PublicJoinKeyCreateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPublicJoinKeyDtoCliFlags(""))...)
	return flags
}

// PublicJoinKeyCreateActionCliHandler builds a full *cli.Command for the
// PublicJoinKeyCreateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PublicJoinKeyCreateActionRequest the same way
// PublicJoinKeyCreateActionHandler (Gin) and PublicJoinKeyCreateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PublicJoinKeyCreateActionCliHandler(
	handler func(c PublicJoinKeyCreateActionRequest) (*PublicJoinKeyCreateActionResponse, error),
) *cli.Command {
	meta := PublicJoinKeyCreateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PublicJoinKeyCreateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PublicJoinKeyCreateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastPublicJoinKeyDtoFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PublicJoinKeyCreateActionCli is a high-level convenience wrapper around
// PublicJoinKeyCreateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PublicJoinKeyCreateActionGin
// registers a route on a Gin engine.
func PublicJoinKeyCreateActionCli(
	app *cli.Command,
	handler func(c PublicJoinKeyCreateActionRequest) (*PublicJoinKeyCreateActionResponse, error),
) {
	app.Commands = append(app.Commands, PublicJoinKeyCreateActionCliHandler(handler))
}
