package backup

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/urfave/cli/v3"
)

var enableCmd = &cli.Command{
	Name: "enable",
	Usage: "Configure Postgres itself for wal-g archiving (wal_level, archive_mode, archive_command, archive_timeout) - " +
		"the one-shot version of README.md's \"Postgres-side configuration\" section. Requires connecting as a superuser.",
	Flags: []cli.Flag{
		&cli.DurationFlag{Name: "archive-timeout", Value: 60 * time.Second, Usage: "worst-case RPO bound: how long Postgres waits before archiving a WAL segment even on a quiet database"},
		&cli.BoolFlag{Name: "restart", Usage: "also restart Postgres (pg_ctl restart) once settings are applied - wal_level/archive_mode need a restart to actually take effect; without this flag you restart it yourself"},
		&cli.StringFlag{
			Name:    "pgdata",
			Sources: cli.EnvVars("PGDATA"),
			Usage:   "Postgres data directory - only required together with --restart",
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		cfg, err := LoadModuleConfig()
		if err != nil {
			return err
		}

		changed, needsRestart, err := EnablePostgresArchiving(ctx, cfg, int(c.Duration("archive-timeout").Seconds()))
		if err != nil {
			return err
		}

		if len(changed) == 0 {
			fmt.Println("already configured: wal_level, archive_mode, archive_command, and archive_timeout all already match")
			return nil
		}

		for _, ch := range changed {
			if ch.Pending {
				fmt.Printf("%-16s %q -> %q (already staged in postgresql.auto.conf, waiting for restart)\n", ch.Name, ch.Old, ch.New)
			} else {
				fmt.Printf("%-16s %q -> %q\n", ch.Name, ch.Old, ch.New)
			}
		}

		if !needsRestart {
			fmt.Println("applied and reloaded - no restart needed (archive_command/archive_timeout take effect from the reload alone)")
			return nil
		}

		if !c.Bool("restart") {
			fmt.Println("wal_level/archive_mode need a full Postgres restart (not just a reload) to actually take effect - restart it yourself when ready, or re-run with --restart --pgdata <dir> to have this do it now")
			return nil
		}

		pgdata := c.String("pgdata")
		if pgdata == "" {
			return errors.New("--restart requires --pgdata (or PGDATA) - the settings above are already saved and will take effect on your next restart regardless")
		}

		fmt.Printf("restarting Postgres (pg_ctl restart -D %s)...\n", pgdata)
		if err := restartPostgres(ctx, pgdata); err != nil {
			return fmt.Errorf("restarting Postgres: %w (the settings above are still saved and will take effect on the next successful restart)", err)
		}
		fmt.Println("restarted - wal_level/archive_mode are now active")
		return nil
	},
}

// restartPostgres runs `pg_ctl restart -D dataDir -w`, waiting for Postgres
// to report ready before returning - the same pg_ctl-based convenience
// StartPostgresAndWait/StopPostgres (Restore.go) already use for restore
// drills, just the "restart" variant instead of a bare start/stop.
func restartPostgres(ctx context.Context, dataDir string) error {
	cmd := exec.CommandContext(ctx, "pg_ctl", "restart", "-D", dataDir, "-w")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
