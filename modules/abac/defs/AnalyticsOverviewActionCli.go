//go:build !wasm

package abacdefs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x AnalyticsOverviewActionRequest) IsCli() bool {
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

// AnalyticsOverviewActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the AnalyticsOverviewAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func AnalyticsOverviewActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	return flags
}

// AnalyticsOverviewActionCliHandler builds a full *cli.Command for the
// AnalyticsOverviewAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a AnalyticsOverviewActionRequest the same way
// AnalyticsOverviewActionHandler (Gin) and AnalyticsOverviewActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func AnalyticsOverviewActionCliHandler(
	handler func(c AnalyticsOverviewActionRequest) (*AnalyticsOverviewActionResponse, error),
) *cli.Command {
	meta := AnalyticsOverviewActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: AnalyticsOverviewActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := AnalyticsOverviewActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// AnalyticsOverviewActionCli is a high-level convenience wrapper around
// AnalyticsOverviewActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way AnalyticsOverviewActionGin
// registers a route on a Gin engine.
func AnalyticsOverviewActionCli(
	app *cli.Command,
	handler func(c AnalyticsOverviewActionRequest) (*AnalyticsOverviewActionResponse, error),
) {
	app.Commands = append(app.Commands, AnalyticsOverviewActionCliHandler(handler))
}
