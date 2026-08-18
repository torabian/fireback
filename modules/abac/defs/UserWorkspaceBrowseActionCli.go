//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
	"strconv"
)

func GetUserWorkspaceBrowseActionQueryCliFlags(prefix string) []emigo.CliFlag {
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

// UserWorkspaceBrowseActionQueryFromCli extracts and casts query parameters the same way
// UserWorkspaceBrowseActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func UserWorkspaceBrowseActionQueryFromCli(c *cli.Command) UserWorkspaceBrowseActionQuery {
	data := UserWorkspaceBrowseActionQuery{}
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
func (x UserWorkspaceBrowseActionRequest) IsCli() bool {
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

// UserWorkspaceBrowseActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the UserWorkspaceBrowseAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func UserWorkspaceBrowseActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetUserWorkspaceBrowseActionQueryCliFlags(""))...)
	return flags
}

// UserWorkspaceBrowseActionCliHandler builds a full *cli.Command for the
// UserWorkspaceBrowseAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a UserWorkspaceBrowseActionRequest the same way
// UserWorkspaceBrowseActionHandler (Gin) and UserWorkspaceBrowseActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func UserWorkspaceBrowseActionCliHandler(
	handler func(c UserWorkspaceBrowseActionRequest) (*UserWorkspaceBrowseActionResponse, error),
) *cli.Command {
	meta := UserWorkspaceBrowseActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: UserWorkspaceBrowseActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := UserWorkspaceBrowseActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		req.QueryParams = UserWorkspaceBrowseActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// UserWorkspaceBrowseActionCli is a high-level convenience wrapper around
// UserWorkspaceBrowseActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way UserWorkspaceBrowseActionGin
// registers a route on a Gin engine.
func UserWorkspaceBrowseActionCli(
	app *cli.Command,
	handler func(c UserWorkspaceBrowseActionRequest) (*UserWorkspaceBrowseActionResponse, error),
) {
	app.Commands = append(app.Commands, UserWorkspaceBrowseActionCliHandler(handler))
}
