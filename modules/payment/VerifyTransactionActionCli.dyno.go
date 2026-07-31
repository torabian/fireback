//go:build !wasm

package payment

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x VerifyTransactionActionRequest) IsCli() bool {
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

// VerifyTransactionActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the VerifyTransactionAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func VerifyTransactionActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetVerifyTransactionActionReqCliFlags(""))...)
	return flags
}

// VerifyTransactionActionCliHandler builds a full *cli.Command for the
// VerifyTransactionAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a VerifyTransactionActionRequest the same way
// VerifyTransactionActionHandler (Gin) and VerifyTransactionActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func VerifyTransactionActionCliHandler(
	handler func(c VerifyTransactionActionRequest) (*VerifyTransactionActionResponse, error),
) *cli.Command {
	meta := VerifyTransactionActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: VerifyTransactionActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := VerifyTransactionActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastVerifyTransactionActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// VerifyTransactionActionCli is a high-level convenience wrapper around
// VerifyTransactionActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way VerifyTransactionActionGin
// registers a route on a Gin engine.
func VerifyTransactionActionCli(
	app *cli.Command,
	handler func(c VerifyTransactionActionRequest) (*VerifyTransactionActionResponse, error),
) {
	app.Commands = append(app.Commands, VerifyTransactionActionCliHandler(handler))
}
