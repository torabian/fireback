//go:build !wasm

package interfacetoolsdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
	"strconv"
)

func GetTimezoneGroupBrowseActionQueryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "qs-filter",
			Type: "string",
		},
		{
			Name: prefix + "qs-sort",
			Type: "string",
		},
		{
			Name: prefix + "qs-start-index",
			Type: "int",
		},
		{
			Name: prefix + "qs-items-per-page",
			Type: "int",
		},
		{
			Name: prefix + "qs-cursor",
			Type: "string",
		},
	}
}

// TimezoneGroupBrowseActionQueryFromCli extracts and casts query parameters the same way
// TimezoneGroupBrowseActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func TimezoneGroupBrowseActionQueryFromCli(c *cli.Command) TimezoneGroupBrowseActionQuery {
	data := TimezoneGroupBrowseActionQuery{}
	values := url.Values{}
	if c.IsSet("qs-filter") {
		data.Filter = c.String("qs-filter")
		values.Set("filter", data.Filter)
	}
	if c.IsSet("qs-sort") {
		data.Sort = c.String("qs-sort")
		values.Set("sort", data.Sort)
	}
	if c.IsSet("qs-start-index") {
		data.StartIndex = int(c.Int64("qs-start-index"))
		values.Set("startIndex", strconv.FormatInt(int64(data.StartIndex), 10))
	}
	if c.IsSet("qs-items-per-page") {
		data.ItemsPerPage = int(c.Int64("qs-items-per-page"))
		values.Set("itemsPerPage", strconv.FormatInt(int64(data.ItemsPerPage), 10))
	}
	if c.IsSet("qs-cursor") {
		data.Cursor = c.String("qs-cursor")
		values.Set("cursor", data.Cursor)
	}
	data.SetValues(values)
	return data
}
func (x TimezoneGroupBrowseActionRequest) IsCli() bool {
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

// TimezoneGroupBrowseActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the TimezoneGroupBrowseAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func TimezoneGroupBrowseActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetTimezoneGroupBrowseActionQueryCliFlags(""))...)
	return flags
}

// TimezoneGroupBrowseActionCliHandler builds a full *cli.Command for the
// TimezoneGroupBrowseAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a TimezoneGroupBrowseActionRequest the same way
// TimezoneGroupBrowseActionHandler (Gin) and TimezoneGroupBrowseActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func TimezoneGroupBrowseActionCliHandler(
	handler func(c TimezoneGroupBrowseActionRequest) (*TimezoneGroupBrowseActionResponse, error),
) *cli.Command {
	meta := TimezoneGroupBrowseActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: TimezoneGroupBrowseActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := TimezoneGroupBrowseActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		req.QueryParams = TimezoneGroupBrowseActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// TimezoneGroupBrowseActionCli is a high-level convenience wrapper around
// TimezoneGroupBrowseActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way TimezoneGroupBrowseActionGin
// registers a route on a Gin engine.
func TimezoneGroupBrowseActionCli(
	app *cli.Command,
	handler func(c TimezoneGroupBrowseActionRequest) (*TimezoneGroupBrowseActionResponse, error),
) {
	app.Commands = append(app.Commands, TimezoneGroupBrowseActionCliHandler(handler))
}
