//go:build !wasm

package abac

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x AcceptInviteActionRequest) IsCli() bool {
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

// AcceptInviteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the AcceptInviteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func AcceptInviteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetAcceptInviteActionReqCliFlags(""))...)
	return flags
}

// AcceptInviteActionCliHandler builds a full *cli.Command for the
// AcceptInviteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a AcceptInviteActionRequest the same way
// AcceptInviteActionHandler (Gin) and AcceptInviteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func AcceptInviteActionCliHandler(
	handler func(c AcceptInviteActionRequest) (*AcceptInviteActionResponse, error),
) *cli.Command {
	meta := AcceptInviteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: AcceptInviteActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := AcceptInviteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastAcceptInviteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// AcceptInviteActionCli is a high-level convenience wrapper around
// AcceptInviteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way AcceptInviteActionGin
// registers a route on a Gin engine.
func AcceptInviteActionCli(
	app *cli.Command,
	handler func(c AcceptInviteActionRequest) (*AcceptInviteActionResponse, error),
) {
	app.Commands = append(app.Commands, AcceptInviteActionCliHandler(handler))
}
