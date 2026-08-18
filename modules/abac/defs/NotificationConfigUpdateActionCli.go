//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetNotificationConfigUpdateActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func NotificationConfigUpdateActionPathParameterFromCli(c *cli.Command) NotificationConfigUpdateActionPathParameter {
	return NotificationConfigUpdateActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x NotificationConfigUpdateActionRequest) IsCli() bool {
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

// NotificationConfigUpdateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the NotificationConfigUpdateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func NotificationConfigUpdateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetNotificationConfigOptionalDtoCliFlags(""))...)
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetNotificationConfigUpdateActionPathParameterCliFlags(""))...)
	return flags
}

// NotificationConfigUpdateActionCliHandler builds a full *cli.Command for the
// NotificationConfigUpdateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a NotificationConfigUpdateActionRequest the same way
// NotificationConfigUpdateActionHandler (Gin) and NotificationConfigUpdateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func NotificationConfigUpdateActionCliHandler(
	handler func(c NotificationConfigUpdateActionRequest) (*NotificationConfigUpdateActionResponse, error),
) *cli.Command {
	meta := NotificationConfigUpdateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: NotificationConfigUpdateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := NotificationConfigUpdateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastNotificationConfigOptionalDtoFromCli(c),
			Params:      NotificationConfigUpdateActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// NotificationConfigUpdateActionCli is a high-level convenience wrapper around
// NotificationConfigUpdateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way NotificationConfigUpdateActionGin
// registers a route on a Gin engine.
func NotificationConfigUpdateActionCli(
	app *cli.Command,
	handler func(c NotificationConfigUpdateActionRequest) (*NotificationConfigUpdateActionResponse, error),
) {
	app.Commands = append(app.Commands, NotificationConfigUpdateActionCliHandler(handler))
}
