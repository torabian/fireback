//go:build !wasm

package abac

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x SendEmailWithProviderActionRequest) IsCli() bool {
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

// SendEmailWithProviderActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the SendEmailWithProviderAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func SendEmailWithProviderActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetSendEmailWithProviderActionReqCliFlags(""))...)
	return flags
}

// SendEmailWithProviderActionCliHandler builds a full *cli.Command for the
// SendEmailWithProviderAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a SendEmailWithProviderActionRequest the same way
// SendEmailWithProviderActionHandler (Gin) and SendEmailWithProviderActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func SendEmailWithProviderActionCliHandler(
	handler func(c SendEmailWithProviderActionRequest) (*SendEmailWithProviderActionResponse, error),
) *cli.Command {
	meta := SendEmailWithProviderActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: SendEmailWithProviderActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := SendEmailWithProviderActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastSendEmailWithProviderActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// SendEmailWithProviderActionCli is a high-level convenience wrapper around
// SendEmailWithProviderActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way SendEmailWithProviderActionGin
// registers a route on a Gin engine.
func SendEmailWithProviderActionCli(
	app *cli.Command,
	handler func(c SendEmailWithProviderActionRequest) (*SendEmailWithProviderActionResponse, error),
) {
	app.Commands = append(app.Commands, SendEmailWithProviderActionCliHandler(handler))
}
