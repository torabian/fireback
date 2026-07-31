//go:build !wasm

package abac

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x UserInvitationsActionRequest) IsCli() bool {
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

// UserInvitationsActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the UserInvitationsAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func UserInvitationsActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	return flags
}

// UserInvitationsActionCliHandler builds a full *cli.Command for the
// UserInvitationsAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a UserInvitationsActionRequest the same way
// UserInvitationsActionHandler (Gin) and UserInvitationsActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func UserInvitationsActionCliHandler(
	handler func(c UserInvitationsActionRequest) (*UserInvitationsActionResponse, error),
) *cli.Command {
	meta := UserInvitationsActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: UserInvitationsActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := UserInvitationsActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// UserInvitationsActionCli is a high-level convenience wrapper around
// UserInvitationsActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way UserInvitationsActionGin
// registers a route on a Gin engine.
func UserInvitationsActionCli(
	app *cli.Command,
	handler func(c UserInvitationsActionRequest) (*UserInvitationsActionResponse, error),
) {
	app.Commands = append(app.Commands, UserInvitationsActionCliHandler(handler))
}
