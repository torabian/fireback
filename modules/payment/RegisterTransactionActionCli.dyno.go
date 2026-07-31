//go:build !wasm

package payment

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x RegisterTransactionActionRequest) IsCli() bool {
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

// RegisterTransactionActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the RegisterTransactionAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func RegisterTransactionActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetRegisterTransactionActionReqCliFlags(""))...)
	return flags
}

// RegisterTransactionActionCliHandler builds a full *cli.Command for the
// RegisterTransactionAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a RegisterTransactionActionRequest the same way
// RegisterTransactionActionHandler (Gin) and RegisterTransactionActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func RegisterTransactionActionCliHandler(
	handler func(c RegisterTransactionActionRequest) (*RegisterTransactionActionResponse, error),
) *cli.Command {
	meta := RegisterTransactionActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: RegisterTransactionActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := RegisterTransactionActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastRegisterTransactionActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// RegisterTransactionActionCli is a high-level convenience wrapper around
// RegisterTransactionActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way RegisterTransactionActionGin
// registers a route on a Gin engine.
func RegisterTransactionActionCli(
	app *cli.Command,
	handler func(c RegisterTransactionActionRequest) (*RegisterTransactionActionResponse, error),
) {
	app.Commands = append(app.Commands, RegisterTransactionActionCliHandler(handler))
}
