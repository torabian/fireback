//go:build !wasm

package abac

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x SendEmailActionRequest) IsCli() bool {
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

// SendEmailActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the SendEmailAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func SendEmailActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetSendEmailActionReqCliFlags(""))...)
	return flags
}

// SendEmailActionCliHandler builds a full *cli.Command for the
// SendEmailAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a SendEmailActionRequest the same way
// SendEmailActionHandler (Gin) and SendEmailActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func SendEmailActionCliHandler(
	handler func(c SendEmailActionRequest) (*SendEmailActionResponse, error),
) *cli.Command {
	meta := SendEmailActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: SendEmailActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := SendEmailActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastSendEmailActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// SendEmailActionCli is a high-level convenience wrapper around
// SendEmailActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way SendEmailActionGin
// registers a route on a Gin engine.
func SendEmailActionCli(
	app *cli.Command,
	handler func(c SendEmailActionRequest) (*SendEmailActionResponse, error),
) {
	app.Commands = append(app.Commands, SendEmailActionCliHandler(handler))
}
