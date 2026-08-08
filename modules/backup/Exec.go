package backup

import (
	"os"

	walgpg "github.com/wal-g/wal-g/cmd/pg"
)

// walgExecMarker is the hidden first argument this binary recognizes to run
// wal-g's own command tree directly in the current process, then exit -
// see MaybeRunEmbeddedWalg and Engine.command.
const walgExecMarker = "__walg-exec"

// MaybeRunEmbeddedWalg makes wal-g a real go.mod dependency
// (github.com/wal-g/wal-g/cmd/pg) instead of a separately installed binary.
// There is no external wal-g executable anywhere: Engine drives wal-g by
// re-executing this same binary with a hidden marker argument, and this
// function is what makes that re-exec do something.
//
// The re-exec (rather than just calling wal-g's command tree in-process
// wherever Engine needs it) exists because wal-g's cobra command tree is
// only safe to run once per process - confirmed by actually trying it
// twice: the second call panics with "flag redefined", since wal-g's own
// global command/flag state is never reset between runs (its own
// maintainers only ever call it once, from their own main()). Giving every
// wal-g operation a fresh process keeps each invocation isolated exactly
// the way it needs to be, without a second binary needing to exist on disk.
//
// Callers (cmd/backup-cli/main.go, cmd/nima-server/main.go) must call this
// as the very first thing in main(), before constructing their own
// cli.Command tree - wal-g's own flags (--json, --full, etc.) must never
// reach this module's own CLI flag parser.
func MaybeRunEmbeddedWalg() {
	if len(os.Args) < 2 {
		return
	}

	switch os.Args[1] {
	case walgExecMarker:
		walgpg.GetCmd().SetArgs(os.Args[2:])
	case "wal-prefetch":
		// wal-g prefetches upcoming WAL segments during wal-fetch by
		// re-exec'ing os.Args[0] with a bare "wal-prefetch ..." command
		// line itself (internal/databases/postgres/wal_prefetcher.go) -
		// found by actually running a restore and seeing "No help topic
		// for 'wal-prefetch'" once this module's own CLI parser (which
		// only knows the __walg-exec marker) saw that call instead. wal-g
		// has no notion of our marker convention when it spawns this
		// itself, so it has to be recognized here too, bare, or prefetch
		// silently stops working the moment wal-g stops being a separate
		// binary.
		walgpg.GetCmd().SetArgs(os.Args[1:])
	default:
		return
	}

	walgpg.Execute() // os.Exit(1) on failure; otherwise just returns
	os.Exit(0)
}
