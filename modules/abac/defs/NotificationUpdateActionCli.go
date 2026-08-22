//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetNotificationUpdateActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func NotificationUpdateActionPathParameterFromCli(c *cli.Command) NotificationUpdateActionPathParameter {
	return NotificationUpdateActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x NotificationUpdateActionRequest) IsCli() bool {
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

// NotificationUpdateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the NotificationUpdateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func NotificationUpdateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetNotificationOptionalDtoCliFlags(""))...)
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetNotificationUpdateActionPathParameterCliFlags(""))...)
	return flags
}

// NotificationUpdateActionCliHandler builds a full *cli.Command for the
// NotificationUpdateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a NotificationUpdateActionRequest the same way
// NotificationUpdateActionHandler (Gin) and NotificationUpdateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func NotificationUpdateActionCliHandler(
	handler func(c NotificationUpdateActionRequest) (*NotificationUpdateActionResponse, error),
) *cli.Command {
	meta := NotificationUpdateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: NotificationUpdateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := NotificationUpdateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastNotificationOptionalDtoFromCli(c),
			Params:      NotificationUpdateActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// NotificationUpdateActionCli is a high-level convenience wrapper around
// NotificationUpdateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way NotificationUpdateActionGin
// registers a route on a Gin engine.
func NotificationUpdateActionCli(
	app *cli.Command,
	handler func(c NotificationUpdateActionRequest) (*NotificationUpdateActionResponse, error),
) {
	app.Commands = append(app.Commands, NotificationUpdateActionCliHandler(handler))
}
