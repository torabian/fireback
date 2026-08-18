//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetTokenAwareDeletePreviewActionQueryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "qs-unique-ids",
			Type: "slice",
		},
	}
}

// TokenAwareDeletePreviewActionQueryFromCli extracts and casts query parameters the same way
// TokenAwareDeletePreviewActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func TokenAwareDeletePreviewActionQueryFromCli(c *cli.Command) TokenAwareDeletePreviewActionQuery {
	data := TokenAwareDeletePreviewActionQuery{}
	values := url.Values{}
	if c.IsSet("qs-unique-ids") {
		raw := c.String("qs-unique-ids")
		emigo.InflatePossibleSlice(raw, &data.UniqueIds)
		values.Set("uniqueIds", raw)
	}
	data.SetValues(values)
	return data
}
func (x TokenAwareDeletePreviewActionRequest) IsCli() bool {
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

// TokenAwareDeletePreviewActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the TokenAwareDeletePreviewAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func TokenAwareDeletePreviewActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetTokenAwareDeletePreviewActionQueryCliFlags(""))...)
	return flags
}

// TokenAwareDeletePreviewActionCliHandler builds a full *cli.Command for the
// TokenAwareDeletePreviewAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a TokenAwareDeletePreviewActionRequest the same way
// TokenAwareDeletePreviewActionHandler (Gin) and TokenAwareDeletePreviewActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func TokenAwareDeletePreviewActionCliHandler(
	handler func(c TokenAwareDeletePreviewActionRequest) (*TokenAwareDeletePreviewActionResponse, error),
) *cli.Command {
	meta := TokenAwareDeletePreviewActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: TokenAwareDeletePreviewActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := TokenAwareDeletePreviewActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		req.QueryParams = TokenAwareDeletePreviewActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// TokenAwareDeletePreviewActionCli is a high-level convenience wrapper around
// TokenAwareDeletePreviewActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way TokenAwareDeletePreviewActionGin
// registers a route on a Gin engine.
func TokenAwareDeletePreviewActionCli(
	app *cli.Command,
	handler func(c TokenAwareDeletePreviewActionRequest) (*TokenAwareDeletePreviewActionResponse, error),
) {
	app.Commands = append(app.Commands, TokenAwareDeletePreviewActionCliHandler(handler))
}
