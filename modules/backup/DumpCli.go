// DumpCli.go adds `backup dump` and `backup restore-dump` - the plain
// pg_dump/mysqldump/sqlite-VACUUM path (Dump.go/RestoreDump.go), separate
// from wal-g's own `backup push`/`backup restore` (Cli.go), which stay
// Postgres-only, point-in-time restore commands. These two work against
// whatever DB_VENDOR the app is actually configured for.
package backup

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli/v3"
)

func dumpRestoreCommands() []*cli.Command {
	return []*cli.Command{dumpCmd, restoreDumpCmd}
}

// resolveDatabase returns c's --database flag if set, otherwise resolves
// one interactively: sqlite has exactly one candidate (cfg.Database, no
// prompt needed), a single-database postgres/mysql connection is used
// without asking, and only a genuine multi-database connection prompts
// (fireback.AskForSelect, wired by modules/fireback/clitools).
func resolveDatabase(ctx context.Context, cfg *DumpConfig, flagValue string) (string, error) {
	if flagValue != "" {
		return flagValue, nil
	}

	names, err := ListDatabases(ctx, cfg)
	if err != nil {
		return "", fmt.Errorf("listing databases to choose from: %w", err)
	}
	switch len(names) {
	case 0:
		return "", errors.New("no databases found on this connection")
	case 1:
		return names[0], nil
	default:
		if fireback.AskForSelect == nil {
			return "", fmt.Errorf("multiple databases found (%v) and no interactive terminal available - pass --database explicitly", names)
		}
		selected := fireback.AskForSelect("Select database to dump", names)
		if selected == "" {
			return "", errors.New("no database selected")
		}
		return selected, nil
	}
}

var dumpCmd = &cli.Command{
	Name: "dump",
	Usage: "Dump one database (postgres/mysql/sqlite) to a zip - to disk, or streamed on the fly over HTTP behind a one-time hash with --hash. " +
		"Run with no flags at all for an interactive wizard (disk vs HTTP, database, output name, size/disk-space estimate).",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "database", Usage: "database to dump; omit to be prompted when more than one exists"},
		&cli.StringFlag{Name: "output", Usage: "output zip path; defaults to <dump-dir>/<database>_<date>.zip (see BACKUP_DUMP_DIR)"},
		&cli.BoolFlag{Name: "hash", Usage: "don't write to disk - register a one-time HTTP download hash instead. The actual dump only runs when something fetches it (see backup restore-dump --hash), streamed live as it's produced, never written to disk here or on the server"},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		cfg, err := LoadDumpConfig()
		if err != nil {
			return err
		}

		// No flags at all -> the interactive wizard. Any flag present at
		// all (even just --database) means "scripted/cron use", which
		// keeps behaving exactly as before: no prompts, no wizard output.
		interactive := !c.IsSet("database") && !c.IsSet("output") && !c.IsSet("hash")

		useHash := c.Bool("hash")
		if interactive {
			useHash, err = askDiskOrHttp()
			if err != nil {
				return err
			}
		}

		database, err := resolveDatabase(ctx, cfg, c.String("database"))
		if err != nil {
			return err
		}

		if interactive {
			printSizeAndSpaceEstimate(ctx, cfg, database, useHash)
		}

		if useHash {
			return dumpHash(cfg, database)
		}

		output := c.String("output")
		if output == "" {
			suggested, err := DefaultDumpPath(cfg, database)
			if err != nil {
				return err
			}
			if interactive && fireback.AskForInput != nil {
				output = fireback.AskForInput("Output file", suggested)
			} else {
				output = suggested
			}
		}

		f, err := os.Create(output)
		if err != nil {
			return fmt.Errorf("creating %s: %w", output, err)
		}
		defer f.Close()

		if err := DumpDatabase(ctx, cfg, database, f); err != nil {
			os.Remove(output)
			return err
		}

		fmt.Printf("dumped %s to %s\n", database, output)
		return nil
	},
}

func askDiskOrHttp() (useHash bool, err error) {
	if fireback.AskForSelect == nil {
		return false, errors.New("no interactive terminal available - pass --database/--output explicitly, or --hash, instead of running with no flags")
	}
	const diskOption = "Write to a local zip file"
	const httpOption = "Stream it over HTTP behind a one-time hash"
	switch fireback.AskForSelect("Where should the backup go?", []string{diskOption, httpOption}) {
	case diskOption:
		return false, nil
	case httpOption:
		return true, nil
	default:
		return false, errors.New("no selection made")
	}
}

// printSizeAndSpaceEstimate is purely informational (never blocks the
// dump) - it prints EstimateDatabaseSize's rough figure, and, only when
// writing to local disk (useHash false - a --hash dump is streamed live
// and never touches local disk here), AvailableDiskSpace for cfg.DumpDir,
// with a warning line if the estimate clearly exceeds what's free.
func printSizeAndSpaceEstimate(ctx context.Context, cfg *DumpConfig, database string, useHash bool) {
	size, sizeErr := EstimateDatabaseSize(ctx, cfg, database)
	if sizeErr != nil {
		fmt.Printf("(couldn't estimate backup size: %v)\n", sizeErr)
	} else {
		fmt.Printf("estimated size: ~%s (%s - the actual dump is often smaller)\n", humanBytes(size), sizeEstimateBasis(cfg.Vendor))
	}

	if useHash {
		fmt.Println("streamed directly over HTTP as it's produced - not written to local disk at any point")
		return
	}

	dir := cfg.DumpDir
	if dir == "" {
		dir = "."
	}
	free, spaceErr := AvailableDiskSpace(dir)
	if spaceErr != nil {
		fmt.Printf("(couldn't determine available disk space for %s: %v)\n", dir, spaceErr)
		return
	}
	fmt.Printf("available disk space (%s): %s\n", dir, humanBytes(free))

	if sizeErr == nil && size > 0 && size > free {
		fmt.Println("WARNING: the estimated backup size is larger than the available disk space - this dump may fail partway through")
	}
}

func sizeEstimateBasis(vendor DumpVendor) string {
	switch vendor {
	case VendorSqlite:
		return "current file size"
	default:
		return "database size on disk"
	}
}

func humanBytes(n int64) string {
	const unit = 1024
	if n < unit {
		return fmt.Sprintf("%d B", n)
	}
	div, exp := int64(unit), 0
	for v := n / unit; v >= unit; v /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(n)/float64(div), "KMGTPE"[exp])
}

// dumpHash registers a one-time dump job directly in the local, OS-private
// job store (see HashRegistry.go) and prints the hash/URL to fetch it from.
// It never runs pg_dump/mysqldump/sqlite itself - the actual dump only
// happens when *something* GETs that URL (see DumpHttp.go's
// fetchDumpHandler), streamed live as it's produced, which is what makes
// the hash single-use and lets it be handed to whoever needs the file
// without giving them database credentials or any other access.
//
// Registering no longer requires talking to a running server over HTTP (an
// earlier version of this called out to POST /backup/dumps, authenticated
// with a since-removed BACKUP_API_TOKEN) - the CLI just writes the job
// where DumpHttp.go's GET handler will later look for it. The one
// requirement this carries forward: whatever process later serves the GET
// (`app start`) must run on this same host, as this same OS user - see
// jobStoreDir's own doc comment.
func dumpHash(cfg *DumpConfig, database string) error {
	ttl := time.Duration(cfg.HashTTLSeconds) * time.Second
	if ttl <= 0 {
		ttl = 30 * time.Minute
	}

	hash, err := registerDumpJob(database, ttl)
	if err != nil {
		return err
	}

	fc := fireback.LoadConfiguration()
	url := fmt.Sprintf("http://%s:%v/backup/dumps/%s/raw", fc.Host, fc.Port, hash)

	fmt.Printf("database:   %s\n", database)
	fmt.Printf("hash:       %s\n", hash)
	fmt.Printf("url:        %s\n", url)
	fmt.Printf("expires at: %s (sooner if fetched - a hash is disabled the instant it's used)\n", time.Now().Add(ttl).Format("2006-01-02T15:04:05Z07:00"))
	fmt.Println("only works while this app's own HTTP server (`app start`/`app s`) is running - the dump itself starts the moment the URL above is fetched, streamed live, never written to disk anywhere first.")
	fmt.Printf("fetch once with: curl -OJ %s\n", url)
	return nil
}

var restoreDumpCmd = &cli.Command{
	Name:  "restore-dump",
	Usage: "Restore a dump (from --file or --hash) into a database - postgres/mysql: created automatically if it doesn't exist yet; refuses if it does, unless --force. sqlite: --database is the destination file path, same refuse-unless-exists-and---force rule",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "database", Usage: "database to restore into; omit to be prompted for it"},
		&cli.StringFlag{Name: "file", Usage: "local dump zip to restore from (as produced by `backup dump`)"},
		&cli.StringFlag{Name: "hash", Usage: "fetch the dump from a running server's --hash URL instead of a local --file"},
		&cli.StringFlag{Name: "server", Usage: "base URL to fetch --hash from; defaults to http://<HOST>:<PORT> from this app's own config"},
		&cli.BoolFlag{Name: "force", Usage: "restore into an already-existing database/file instead of refusing - does not drop/clean anything first, a real conflict (e.g. a table that already exists) still fails normally"},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		cfg, err := LoadDumpConfig()
		if err != nil {
			return err
		}

		database := c.String("database")
		if database == "" {
			if fireback.AskForInput == nil {
				return errors.New("--database is required (no interactive terminal available)")
			}
			database = fireback.AskForInput("Database to restore into", cfg.Database)
		}
		if database == "" {
			return errors.New("a target database name is required")
		}

		file := c.String("file")
		hash := c.String("hash")
		if file == "" && hash == "" {
			return errors.New("provide either --file <path> or --hash <hash> to restore from")
		}

		var src io.Reader
		if hash != "" {
			server := c.String("server")
			if server == "" {
				fc := fireback.LoadConfiguration()
				server = fmt.Sprintf("http://%s:%v", fc.Host, fc.Port)
			}
			resp, err := http.Get(server + "/backup/dumps/" + hash + "/raw")
			if err != nil {
				return fmt.Errorf("fetching %s: %w", server, err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				body, _ := io.ReadAll(resp.Body)
				return fmt.Errorf("server returned %s: %s", resp.Status, string(body))
			}
			src = resp.Body
		} else {
			f, err := os.Open(file)
			if err != nil {
				return err
			}
			defer f.Close()
			src = f
		}

		if err := RestoreDump(ctx, cfg, database, src, c.Bool("force")); err != nil {
			return err
		}

		fmt.Printf("restored into %s\n", database)
		return nil
	},
}
