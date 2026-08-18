//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetPublicAuthenticationAwareDeletePreviewActionQueryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "qs-unique-ids",
			Type: "slice",
		},
	}
}

// PublicAuthenticationAwareDeletePreviewActionQueryFromCli extracts and casts query parameters the same way
// PublicAuthenticationAwareDeletePreviewActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func PublicAuthenticationAwareDeletePreviewActionQueryFromCli(c *cli.Command) PublicAuthenticationAwareDeletePreviewActionQuery {
	data := PublicAuthenticationAwareDeletePreviewActionQuery{}
	values := url.Values{}
	if c.IsSet("qs-unique-ids") {
		raw := c.String("qs-unique-ids")
		emigo.InflatePossibleSlice(raw, &data.UniqueIds)
		values.Set("uniqueIds", raw)
	}
	data.SetValues(values)
	return data
}
func (x PublicAuthenticationAwareDeletePreviewActionRequest) IsCli() bool {
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

// PublicAuthenticationAwareDeletePreviewActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PublicAuthenticationAwareDeletePreviewAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PublicAuthenticationAwareDeletePreviewActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPublicAuthenticationAwareDeletePreviewActionQueryCliFlags(""))...)
	return flags
}

// PublicAuthenticationAwareDeletePreviewActionCliHandler builds a full *cli.Command for the
// PublicAuthenticationAwareDeletePreviewAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PublicAuthenticationAwareDeletePreviewActionRequest the same way
// PublicAuthenticationAwareDeletePreviewActionHandler (Gin) and PublicAuthenticationAwareDeletePreviewActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PublicAuthenticationAwareDeletePreviewActionCliHandler(
	handler func(c PublicAuthenticationAwareDeletePreviewActionRequest) (*PublicAuthenticationAwareDeletePreviewActionResponse, error),
) *cli.Command {
	meta := PublicAuthenticationAwareDeletePreviewActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PublicAuthenticationAwareDeletePreviewActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PublicAuthenticationAwareDeletePreviewActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		req.QueryParams = PublicAuthenticationAwareDeletePreviewActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PublicAuthenticationAwareDeletePreviewActionCli is a high-level convenience wrapper around
// PublicAuthenticationAwareDeletePreviewActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PublicAuthenticationAwareDeletePreviewActionGin
// registers a route on a Gin engine.
func PublicAuthenticationAwareDeletePreviewActionCli(
	app *cli.Command,
	handler func(c PublicAuthenticationAwareDeletePreviewActionRequest) (*PublicAuthenticationAwareDeletePreviewActionResponse, error),
) {
	app.Commands = append(app.Commands, PublicAuthenticationAwareDeletePreviewActionCliHandler(handler))
}
