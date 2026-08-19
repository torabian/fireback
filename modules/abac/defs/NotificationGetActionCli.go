package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetNotificationGetActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func NotificationGetActionPathParameterFromCli(c *cli.Command) NotificationGetActionPathParameter {
	return NotificationGetActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x NotificationGetActionRequest) IsCli() bool {
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

// NotificationGetActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the NotificationGetAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func NotificationGetActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetNotificationGetActionPathParameterCliFlags(""))...)
	return flags
}

// NotificationGetActionCliHandler builds a full *cli.Command for the
// NotificationGetAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a NotificationGetActionRequest the same way
// NotificationGetActionHandler (Gin) and NotificationGetActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func NotificationGetActionCliHandler(
	handler func(c NotificationGetActionRequest) (*NotificationGetActionResponse, error),
) *cli.Command {
	meta := NotificationGetActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: NotificationGetActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := NotificationGetActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Params:      NotificationGetActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// NotificationGetActionCli is a high-level convenience wrapper around
// NotificationGetActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way NotificationGetActionGin
// registers a route on a Gin engine.
func NotificationGetActionCli(
	app *cli.Command,
	handler func(c NotificationGetActionRequest) (*NotificationGetActionResponse, error),
) {
	app.Commands = append(app.Commands, NotificationGetActionCliHandler(handler))
}
