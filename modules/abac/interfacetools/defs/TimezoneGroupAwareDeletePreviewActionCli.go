//go:build !wasm

package interfacetoolsdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func GetTimezoneGroupAwareDeletePreviewActionQueryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "qs-unique-ids",
			Type: "slice",
		},
	}
}

// TimezoneGroupAwareDeletePreviewActionQueryFromCli extracts and casts query parameters the same way
// TimezoneGroupAwareDeletePreviewActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func TimezoneGroupAwareDeletePreviewActionQueryFromCli(c *cli.Command) TimezoneGroupAwareDeletePreviewActionQuery {
	data := TimezoneGroupAwareDeletePreviewActionQuery{}
	values := url.Values{}
	if c.IsSet("qs-unique-ids") {
		raw := c.String("qs-unique-ids")
		emigo.InflatePossibleSlice(raw, &data.UniqueIds)
		values.Set("uniqueIds", raw)
	}
	data.SetValues(values)
	return data
}
func (x TimezoneGroupAwareDeletePreviewActionRequest) IsCli() bool {
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

// TimezoneGroupAwareDeletePreviewActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the TimezoneGroupAwareDeletePreviewAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func TimezoneGroupAwareDeletePreviewActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetTimezoneGroupAwareDeletePreviewActionQueryCliFlags(""))...)
	return flags
}

// TimezoneGroupAwareDeletePreviewActionCliHandler builds a full *cli.Command for the
// TimezoneGroupAwareDeletePreviewAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a TimezoneGroupAwareDeletePreviewActionRequest the same way
// TimezoneGroupAwareDeletePreviewActionHandler (Gin) and TimezoneGroupAwareDeletePreviewActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func TimezoneGroupAwareDeletePreviewActionCliHandler(
	handler func(c TimezoneGroupAwareDeletePreviewActionRequest) (*TimezoneGroupAwareDeletePreviewActionResponse, error),
) *cli.Command {
	meta := TimezoneGroupAwareDeletePreviewActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: TimezoneGroupAwareDeletePreviewActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := TimezoneGroupAwareDeletePreviewActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		req.QueryParams = TimezoneGroupAwareDeletePreviewActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// TimezoneGroupAwareDeletePreviewActionCli is a high-level convenience wrapper around
// TimezoneGroupAwareDeletePreviewActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way TimezoneGroupAwareDeletePreviewActionGin
// registers a route on a Gin engine.
func TimezoneGroupAwareDeletePreviewActionCli(
	app *cli.Command,
	handler func(c TimezoneGroupAwareDeletePreviewActionRequest) (*TimezoneGroupAwareDeletePreviewActionResponse, error),
) {
	app.Commands = append(app.Commands, TimezoneGroupAwareDeletePreviewActionCliHandler(handler))
}
