//go:build !wasm

package payment

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetPayInvoiceActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func PayInvoiceActionPathParameterFromCli(c *cli.Command) PayInvoiceActionPathParameter {
	return PayInvoiceActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func GetPayInvoiceActionQueryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "qs-invoice-id",
			Type:        "string",
			Description: "Created invoice to be payed via strip",
		},
	}
}

// PayInvoiceActionQueryFromCli extracts and casts query parameters the same way
// PayInvoiceActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func PayInvoiceActionQueryFromCli(c *cli.Command) PayInvoiceActionQuery {
	data := PayInvoiceActionQuery{}
	values := url.Values{}
	if c.IsSet("qs-invoice-id") {
		data.InvoiceId = c.String("qs-invoice-id")
		values.Set("invoiceId", data.InvoiceId)
	}
	data.SetValues(values)
	return data
}
func (x PayInvoiceActionRequest) IsCli() bool {
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

// PayInvoiceActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PayInvoiceAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PayInvoiceActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPayInvoiceActionPathParameterCliFlags(""))...)
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPayInvoiceActionQueryCliFlags(""))...)
	return flags
}

// PayInvoiceActionCliHandler builds a full *cli.Command for the
// PayInvoiceAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PayInvoiceActionRequest the same way
// PayInvoiceActionHandler (Gin) and PayInvoiceActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PayInvoiceActionCliHandler(
	handler func(c PayInvoiceActionRequest) (*PayInvoiceActionResponse, error),
) *cli.Command {
	meta := PayInvoiceActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PayInvoiceActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PayInvoiceActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Params:      PayInvoiceActionPathParameterFromCli(c),
		}
		req.QueryParams = PayInvoiceActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PayInvoiceActionCli is a high-level convenience wrapper around
// PayInvoiceActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PayInvoiceActionGin
// registers a route on a Gin engine.
func PayInvoiceActionCli(
	app *cli.Command,
	handler func(c PayInvoiceActionRequest) (*PayInvoiceActionResponse, error),
) {
	app.Commands = append(app.Commands, PayInvoiceActionCliHandler(handler))
}
