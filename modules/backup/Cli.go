// Cli.go exposes this module's operations as urfave/cli/v3 commands via
// Commands(), the same convention modules/storage/Cli.go uses. Unlike that
// module, nothing here opens its own database connection - every operation
// either drives wal-g (a go.mod dependency, run in-process via a re-exec of
// this same binary - see Exec.go/Engine.go; it handles the Postgres/storage
// connection itself, configured entirely through env vars, see Config.go)
// or drives a local Postgres process directly via pg_ctl for restores. So
// there's no withPool-style wrapper: LoadModuleConfig reads connection
// details from fireback's config when running inside a full app, or from
// the standard PGHOST/PGPORT/... env vars when run standalone
// (cmd/backup-cli), exactly mirroring how wal-g itself is configured.
package backup

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/urfave/cli/v3"
)

func Commands() []*cli.Command {
	cmds := []*cli.Command{
		initCmd,
		enableCmd,
		pushCmd,
		listCmd,
		restoreCmd,
		verifyCmd,
		checkCmd,
		pruneCmd,
		downloadCmd,
		archivePushCmd,
		archiveGetCmd,
	}
	// dump/restore-dump (Dump.go/RestoreDump.go): plain pg_dump/mysqldump/
	// sqlite snapshots, for whichever vendor DB_VENDOR is actually set to -
	// unlike every command above, which is wal-g and Postgres-only.
	return append(cmds, dumpRestoreCommands()...)
}

func loadEngine() (*ModuleConfig, *Engine, error) {
	cfg, err := LoadModuleConfig()
	if err != nil {
		return nil, nil, err
	}
	return cfg, NewEngine(cfg), nil
}

var initCmd = &cli.Command{
	Name:  "init",
	Usage: "Check that this environment is ready for wal-g backups (embedded wal-g runs, config resolvable)",
	Action: func(ctx context.Context, c *cli.Command) error {
		cfg, err := LoadModuleConfig()
		if err != nil {
			return err
		}
		if err := os.MkdirAll(cfg.FilePrefix, 0o755); err != nil {
			return fmt.Errorf("creating/checking WALG_FILE_PREFIX %q: %w", cfg.FilePrefix, err)
		}

		engine := NewEngine(cfg)
		if err := engine.Run(ctx, "--version"); err != nil {
			return fmt.Errorf("running embedded wal-g: %w", err)
		}

		fmt.Printf("ok: embedded wal-g runs, storage prefix %s is writable, Postgres target %s@%s:%s/%s\n",
			cfg.FilePrefix, cfg.PgUser, cfg.PgHost, cfg.PgPort, cfg.PgDatabase)
		fmt.Println("reminder: Postgres itself still needs wal_level=replica, archive_mode=on, and archive_command set - run `backup enable` to configure these automatically, or see modules/backup/README.md for the manual steps")
		return nil
	},
}

var pushCmd = &cli.Command{
	Name:  "push",
	Usage: "Take a base backup now (full or delta, depending on WALG_DELTA_MAX_STEPS)",
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "full", Usage: "force a full backup regardless of delta settings"},
		&cli.StringFlag{
			Name:    "pgdata",
			Sources: cli.EnvVars("PGDATA"),
			Usage:   "Postgres data directory to back up - this command must run where that directory is locally readable (co-located with Postgres, e.g. inside the postgres container), not remotely",
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		pgdata := c.String("pgdata")
		if pgdata == "" {
			return errors.New("--pgdata (or PGDATA) is required: backup-push needs local filesystem access to Postgres's data directory to support delta backups")
		}

		cfg, engine, err := loadEngine()
		if err != nil {
			return err
		}

		full := c.Bool("full")
		if err := engine.BackupPush(ctx, pgdata, full); err != nil {
			return err
		}

		backups, err := engine.BackupList(ctx)
		if err != nil || len(backups) == 0 {
			// The push itself succeeded; a listing failure here shouldn't
			// be reported as the push having failed.
			return nil
		}
		newest := backups[0]
		for _, b := range backups {
			if b.Time.After(newest.Time) {
				newest = b
			}
		}

		catalog, err := LoadCatalog(cfg)
		if err == nil {
			backupType := "delta"
			if full {
				backupType = "full"
			}
			catalog.RecordPush(newest.BackupName, backupType)
			_ = catalog.Save(cfg)
		}

		fmt.Printf("pushed %s\n", newest.BackupName)
		return nil
	},
}

var listCmd = &cli.Command{
	Name:  "list",
	Usage: "List available backups",
	Action: func(ctx context.Context, c *cli.Command) error {
		_, engine, err := loadEngine()
		if err != nil {
			return err
		}
		backups, err := engine.BackupList(ctx)
		if err != nil {
			return err
		}
		if len(backups) == 0 {
			fmt.Println("no backups found")
			return nil
		}
		for _, b := range backups {
			fmt.Printf("%s\t%s\n", b.BackupName, b.Time.Format(time.RFC3339))
		}
		return nil
	},
}

var restoreCmd = &cli.Command{
	Name:  "restore",
	Usage: "Restore: fetch the right base backup and configure Postgres to replay WAL - to --at if given, or as fully up to date as available WAL allows otherwise",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "at", Usage: "target time, RFC3339 (e.g. 2026-07-30T12:00:00Z). Omit to recreate the database as fully up to date as possible instead of restoring to a specific past instant"},
		&cli.StringFlag{Name: "target", Required: true, Usage: "empty/nonexistent data directory to restore into"},
		&cli.BoolFlag{Name: "start", Usage: "also start Postgres against --target and wait for recovery to finish"},
		&cli.BoolFlag{Name: "promote", Usage: "auto-promote once --at is reached, instead of pausing recovery for inspection first (ignored without --at, which always auto-promotes)"},
		&cli.IntFlag{Name: "port", Value: 5433, Usage: "port to start Postgres on when --start is set"},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		var targetTime *time.Time
		if c.IsSet("at") {
			t, err := time.Parse(time.RFC3339, c.String("at"))
			if err != nil {
				return fmt.Errorf("--at must be RFC3339, e.g. 2026-07-30T12:00:00Z: %w", err)
			}
			targetTime = &t
		}

		_, engine, err := loadEngine()
		if err != nil {
			return err
		}

		_, err = Restore(ctx, engine, RestoreOptions{
			TargetTime: targetTime,
			TargetDir:  c.String("target"),
			Start:      c.Bool("start"),
			Promote:    c.Bool("promote"),
			Port:       int(c.Int("port")),
		})
		return err
	},
}

var verifyCmd = &cli.Command{
	Name:  "verify",
	Usage: "Restore drill: fetch the latest backup into a scratch dir, replay all WAL, and confirm Postgres actually starts",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "scratch", Value: filepath.Join(os.TempDir(), "nima-backup-verify"), Usage: "scratch data directory - wiped before and after"},
		&cli.IntFlag{Name: "port", Value: 5433, Usage: "port to start the drill Postgres instance on"},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		cfg, engine, err := loadEngine()
		if err != nil {
			return err
		}

		verifyErr := Verify(ctx, engine, c.String("scratch"), int(c.Int("port")))

		backups, listErr := engine.BackupList(ctx)
		if listErr == nil && len(backups) > 0 {
			newest := backups[0]
			for _, b := range backups {
				if b.Time.After(newest.Time) {
					newest = b
				}
			}
			if catalog, err := LoadCatalog(cfg); err == nil {
				notes := ""
				if verifyErr != nil {
					notes = verifyErr.Error()
				}
				catalog.RecordVerify(newest.BackupName, verifyErr == nil, notes)
				_ = catalog.Save(cfg)
			}
		}

		if verifyErr != nil {
			return fmt.Errorf("verify failed: %w", verifyErr)
		}
		fmt.Println("verify ok: latest backup restored and reached a promoted state")
		return nil
	},
}

var checkCmd = &cli.Command{
	Name:  "check",
	Usage: "Fast health check for cron/monitoring: exits non-zero if backups are stale or WAL archiving has gaps",
	Flags: []cli.Flag{
		&cli.DurationFlag{Name: "max-age", Value: 26 * time.Hour, Usage: "how old the newest backup is allowed to be before this fails"},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		_, engine, err := loadEngine()
		if err != nil {
			return err
		}
		result, err := Check(ctx, engine, c.Duration("max-age"))
		if err != nil {
			return err
		}
		if !result.OK {
			for _, reason := range result.Reasons {
				fmt.Fprintln(os.Stderr, "check:", reason)
			}
			return errors.New("backup check failed")
		}
		fmt.Println("check ok")
		return nil
	},
}

var pruneCmd = &cli.Command{
	Name:  "prune",
	Usage: "Enforce retention: keep the newest --retain full backups (and what's needed to restore from them), delete the rest",
	Flags: []cli.Flag{
		&cli.IntFlag{Name: "retain", Usage: "full backups to keep; defaults to BACKUP_RETAIN_FULL or 4"},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		cfg, engine, err := loadEngine()
		if err != nil {
			return err
		}
		retain := cfg.RetainFullBackups
		if c.IsSet("retain") {
			retain = int(c.Int("retain"))
		}
		return Prune(ctx, engine, retain)
	},
}

var downloadCmd = &cli.Command{
	Name:      "download",
	Usage:     "Copy a backup to a destination path, ready to zip up and move off-box",
	ArgsUsage: "<backup-name|LATEST> <dest>",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "with-history",
			Usage: "also bring along every WAL segment through now, so the copy can restore to any instant up to when you ran this, not just to <backup-name>'s own completion (needed to later 'recreate the database' from this copy)",
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		name := c.Args().Get(0)
		dest := c.Args().Get(1)
		if name == "" || dest == "" {
			return errors.New("usage: backup download <backup-name|LATEST> <dest>")
		}
		_, engine, err := loadEngine()
		if err != nil {
			return err
		}
		if err := Download(ctx, engine, name, dest, c.Bool("with-history")); err != nil {
			return err
		}
		fmt.Printf("copied %s to %s\n", name, dest)
		return nil
	},
}

// archivePushCmd/archiveGetCmd are the intended archive_command/
// restore_command targets (instead of pointing Postgres at bare wal-g),
// so a failure is captured in .nima-archive.log for `backup check`-adjacent
// debugging while still preserving wal-g's own exit code - critical for
// wal-fetch's documented exit-74 "not found yet, keep retrying" contract,
// which a generic os.Exit(1) would silently break.
var archivePushCmd = &cli.Command{
	Name:      "archive-push",
	Usage:     "archive_command target: wal-g wal-push wrapper with failure logging",
	ArgsUsage: "<wal-file-path>",
	Action: func(ctx context.Context, c *cli.Command) error {
		path := c.Args().First()
		if path == "" {
			return errors.New("usage: backup archive-push <wal-file-path>")
		}
		cfg, engine, err := loadEngine()
		if err != nil {
			return err
		}
		if err := engine.WalPush(ctx, path); err != nil {
			appendArchiveLog(cfg, fmt.Sprintf("archive-push %s failed: %v", path, err))
			return exitWithSameCode(err)
		}
		return nil
	},
}

var archiveGetCmd = &cli.Command{
	Name:      "archive-get",
	Usage:     "restore_command target: wal-g wal-fetch wrapper with failure logging",
	ArgsUsage: "<wal-name> <dest-path>",
	Action: func(ctx context.Context, c *cli.Command) error {
		walName := c.Args().Get(0)
		dest := c.Args().Get(1)
		if walName == "" || dest == "" {
			return errors.New("usage: backup archive-get <wal-name> <dest-path>")
		}
		cfg, engine, err := loadEngine()
		if err != nil {
			return err
		}
		if err := engine.WalFetch(ctx, walName, dest); err != nil {
			appendArchiveLog(cfg, fmt.Sprintf("archive-get %s failed: %v", walName, err))
			return exitWithSameCode(err)
		}
		return nil
	},
}

func appendArchiveLog(cfg *ModuleConfig, msg string) {
	f, err := os.OpenFile(filepath.Join(cfg.FilePrefix, ".nima-archive.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return
	}
	defer f.Close()
	fmt.Fprintf(f, "%s %s\n", time.Now().UTC().Format(time.RFC3339), msg)
}

// exitWithSameCode terminates the process immediately with err's own exit
// code when it wraps an *exec.ExitError (e.g. wal-fetch's 74), instead of
// letting urfave/cli's default error handling collapse every failure to
// exit code 1.
func exitWithSameCode(err error) error {
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		os.Exit(exitErr.ExitCode())
	}
	return err
}
