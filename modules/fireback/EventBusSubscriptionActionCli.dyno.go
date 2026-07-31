//go:build !wasm

package fireback

import (
	"bufio"
	"context"
	"github.com/urfave/cli/v3"
	"os"
)

// EventBusSubscriptionActionCliFlags returns the query-parameter flags the
// EventBusSubscriptionAction action can bind from urfave v3.
//
func EventBusSubscriptionActionCliFlags() []cli.Flag {
	flags := []cli.Flag{}
	return flags
}

// EventBusSubscriptionActionCliReactiveHandler builds a full *cli.Command for the
// EventBusSubscriptionAction reactive action: stdin becomes the read side (one frame per
// line), and whatever the factory's returned channel produces is written straight to
// stdout - so the generated command composes as one leg of a Linux pipe
// (`producer | app the-action | consumer`). Piping ends the command the same way an
// EOF on a socket would: stdin closing (or scanner error) ends the read loop and the
// command returns.
func EventBusSubscriptionActionCliReactiveHandler(
	factory func(session EventBusSubscriptionActionSession) (chan []byte, error),
) *cli.Command {
	meta := EventBusSubscriptionActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: EventBusSubscriptionActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		done := make(chan bool)
		read := make(chan EventBusSubscriptionActionReadChan)
		session := EventBusSubscriptionActionSession{
			Ctx:  c,
			Done: done,
			Read: read,
		}
		out, err := factory(session)
		if err != nil {
			return err
		}
		go func() {
			for {
				select {
				case <-done:
					return
				case data, more := <-out:
					os.Stdout.Write(data)
					if !more {
						return
					}
				}
			}
		}()
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := make([]byte, len(scanner.Bytes()))
			copy(line, scanner.Bytes())
			read <- EventBusSubscriptionActionReadChan{Data: line}
		}
		return scanner.Err()
	}
	return cmd
}

// EventBusSubscriptionActionCliReactive is a high-level convenience wrapper around
// EventBusSubscriptionActionCliReactiveHandler. It registers the generated command as a
// subcommand of an existing urfave v3 *cli.Command, the same way EventBusSubscriptionActionGin
// registers a route on a Gin engine.
func EventBusSubscriptionActionCliReactive(
	app *cli.Command,
	factory func(session EventBusSubscriptionActionSession) (chan []byte, error),
) {
	app.Commands = append(app.Commands, EventBusSubscriptionActionCliReactiveHandler(factory))
}
