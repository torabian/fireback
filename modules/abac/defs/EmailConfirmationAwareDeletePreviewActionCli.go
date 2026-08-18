//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetEmailConfirmationAwareDeletePreviewActionQueryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "qs-unique-ids",
			Type: "slice",
		},
	}
}

// EmailConfirmationAwareDeletePreviewActionQueryFromCli extracts and casts query parameters the same way
// EmailConfirmationAwareDeletePreviewActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func EmailConfirmationAwareDeletePreviewActionQueryFromCli(c *cli.Command) EmailConfirmationAwareDeletePreviewActionQuery {
	data := EmailConfirmationAwareDeletePreviewActionQuery{}
	values := url.Values{}
	if c.IsSet("qs-unique-ids") {
		raw := c.String("qs-unique-ids")
		emigo.InflatePossibleSlice(raw, &data.UniqueIds)
		values.Set("uniqueIds", raw)
	}
	data.SetValues(values)
	return data
}
func (x EmailConfirmationAwareDeletePreviewActionRequest) IsCli() bool {
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

// EmailConfirmationAwareDeletePreviewActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the EmailConfirmationAwareDeletePreviewAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func EmailConfirmationAwareDeletePreviewActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetEmailConfirmationAwareDeletePreviewActionQueryCliFlags(""))...)
	return flags
}

// EmailConfirmationAwareDeletePreviewActionCliHandler builds a full *cli.Command for the
// EmailConfirmationAwareDeletePreviewAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a EmailConfirmationAwareDeletePreviewActionRequest the same way
// EmailConfirmationAwareDeletePreviewActionHandler (Gin) and EmailConfirmationAwareDeletePreviewActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func EmailConfirmationAwareDeletePreviewActionCliHandler(
	handler func(c EmailConfirmationAwareDeletePreviewActionRequest) (*EmailConfirmationAwareDeletePreviewActionResponse, error),
) *cli.Command {
	meta := EmailConfirmationAwareDeletePreviewActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: EmailConfirmationAwareDeletePreviewActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := EmailConfirmationAwareDeletePreviewActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		req.QueryParams = EmailConfirmationAwareDeletePreviewActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// EmailConfirmationAwareDeletePreviewActionCli is a high-level convenience wrapper around
// EmailConfirmationAwareDeletePreviewActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way EmailConfirmationAwareDeletePreviewActionGin
// registers a route on a Gin engine.
func EmailConfirmationAwareDeletePreviewActionCli(
	app *cli.Command,
	handler func(c EmailConfirmationAwareDeletePreviewActionRequest) (*EmailConfirmationAwareDeletePreviewActionResponse, error),
) {
	app.Commands = append(app.Commands, EmailConfirmationAwareDeletePreviewActionCliHandler(handler))
}
