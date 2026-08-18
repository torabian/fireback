//go:build !wasm

package interfacetoolsdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x TimezoneGroupCreateActionRequest) IsCli() bool {
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

// TimezoneGroupCreateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the TimezoneGroupCreateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func TimezoneGroupCreateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetTimezoneGroupDtoCliFlags(""))...)
	return flags
}

// TimezoneGroupCreateActionCliHandler builds a full *cli.Command for the
// TimezoneGroupCreateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a TimezoneGroupCreateActionRequest the same way
// TimezoneGroupCreateActionHandler (Gin) and TimezoneGroupCreateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func TimezoneGroupCreateActionCliHandler(
	handler func(c TimezoneGroupCreateActionRequest) (*TimezoneGroupCreateActionResponse, error),
) *cli.Command {
	meta := TimezoneGroupCreateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: TimezoneGroupCreateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := TimezoneGroupCreateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastTimezoneGroupDtoFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// TimezoneGroupCreateActionCli is a high-level convenience wrapper around
// TimezoneGroupCreateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way TimezoneGroupCreateActionGin
// registers a route on a Gin engine.
func TimezoneGroupCreateActionCli(
	app *cli.Command,
	handler func(c TimezoneGroupCreateActionRequest) (*TimezoneGroupCreateActionResponse, error),
) {
	app.Commands = append(app.Commands, TimezoneGroupCreateActionCliHandler(handler))
}
