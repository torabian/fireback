// Cli.go adds `internalstats watch` on top of the CLI commands InternalStatsModule.go
// already gets for free from the generated defs (`internalstats snapshot` - see
// internalstatsdefs.InternalStatsSnapshotActionCliHandler in InternalStatsModule.go).
// watch renders the same StreamSnapshots feed the reactive websocket action serves
// (Stream.go) as a colored, left-label/right-value table, redrawn in place on every
// tick - the terminal equivalent of subscribing to InternalStatsStream, minus the
// network hop, since this already runs on the box being measured. Like every other
// local admin command in this codebase (`backup dump`, `config list`, ...) it doesn't
// go through AuthorizeFn - that gate exists for the HTTP/websocket routes, not for an
// operator already holding a shell on this machine.
package internalstats

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/urfave/cli/v3"

	internalstatsdefs "github.com/torabian/fireback/modules/internalstats/defs"
)

var (
	colorCategory = color.New(color.FgCyan, color.Bold)
	colorHeader   = color.New(color.FgHiWhite, color.Bold)
	colorDim      = color.New(color.FgHiBlack)
	colorOk       = color.New(color.FgGreen)
	colorWarn     = color.New(color.FgYellow)
	colorCritical = color.New(color.FgRed, color.Bold)
)

func watchCommand(interval time.Duration) *cli.Command {
	return &cli.Command{
		Name:  "watch",
		Usage: fmt.Sprintf("Live-refreshing colored table of server stats, redrawn every %s (Ctrl+C to quit)", interval),
		Action: func(cliCtx context.Context, c *cli.Command) error {
			ctx, cancel := context.WithCancel(cliCtx)
			defer cancel()

			sigCh := make(chan os.Signal, 1)
			signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
			defer signal.Stop(sigCh)
			go func() {
				<-sigCh
				cancel()
			}()

			for snapshot := range StreamSnapshots(ctx, interval) {
				clearScreen()
				renderTable(os.Stdout, snapshot, interval)
			}

			return nil
		},
	}
}

// clearScreen moves the cursor home and clears the terminal (standard ANSI, the same
// trick `top`/`htop` use) so each tick redraws in place instead of scrolling.
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

const (
	labelWidth = 22
	valueWidth = 16
)

func renderTable(w *os.File, snapshot *internalstatsdefs.InternalStatsSnapshotActionRes, interval time.Duration) {
	colorHeader.Fprintf(w, "INTERNAL STATS  ")
	fmt.Fprintln(w, snapshot.Hostname)
	colorDim.Fprintf(w, "updated %s · refreshing every %s · Ctrl+C to quit\n", snapshot.GeneratedAt, interval)

	lastCategory := ""
	for _, item := range snapshot.Items.Items {
		if item.Category != lastCategory {
			fmt.Fprintln(w)
			colorCategory.Fprintln(w, strings.ToUpper(item.Category))
			lastCategory = item.Category
		}

		// Pad the plain text first, then hand the already-padded string to color -
		// coloring after padding keeps column alignment correct, since color only
		// wraps invisible ANSI escapes around the text without changing what's
		// visibly printed (padding a string that already contains escape codes
		// would count them as visible width and misalign every column).
		label := fmt.Sprintf("  %-*s", labelWidth, item.Label)
		value := fmt.Sprintf("%*s", valueWidth, item.Value)

		fmt.Fprint(w, label)
		severityColor(item.Severity).Fprintln(w, value)
	}
}

func severityColor(severity string) *color.Color {
	switch severity {
	case SeverityOk:
		return colorOk
	case SeverityWarn:
		return colorWarn
	case SeverityCritical:
		return colorCritical
	default:
		return colorDim
	}
}
