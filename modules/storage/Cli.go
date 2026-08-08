// cli.go exposes the operations in admin.go as urfave/cli/v3 commands via
// Commands(), so they can be mounted into any cli.App - the standalone
// cmd/fileupload-cli binary is just a thin wrapper around it, and the same
// slice can be appended to a full fireback app's own command tree (e.g.
// cmd/nima-server) to get these for free there too.
package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/urfave/cli/v3"
)

// Commands returns the usage/upload/delete subcommands as a group, ready to
// drop into any cli.App's Commands slice.
func Commands() []*cli.Command {
	return []*cli.Command{usageCmd, uploadCmd, deleteCmd}
}

// databaseURLFlag is attached to every command in Commands so the
// standalone fileupload-cli binary works on its own; withPool falls back to
// fireback's own config when it's left unset, which is what happens when
// Commands is mounted into a fireback app's CLI instead.
func databaseURLFlag() cli.Flag {
	return &cli.StringFlag{
		Name:    "database-url",
		Sources: cli.EnvVars("DATABASE_URL"),
		Usage:   "Postgres connection string; falls back to the fireback app's own database config if omitted",
	}
}

// withPool resolves a *pgxpool.Pool and hands it to fn, trying in order:
//
//  1. Pool() - the shared pool MountAll already opened, if this process is
//     also running (or has already started) the fileupload module inside a
//     live app. Reused as-is, not closed afterwards.
//  2. --database-url, or the DATABASE_URL env var it falls back to - for the
//     standalone fileupload-cli binary pointed at a database with no app
//     running against it.
//  3. NewPgxPool (webserver.go), the same fireback-config-derived connection
//     mountOnFirebackApp uses - so when Commands is mounted into a fireback
//     app's own CLI (e.g. cmd/nima-server), it resolves the same database
//     that app's web server would, with no flag needed.
//
// A pool opened for cases 2/3 is closed once fn returns.
func withPool(ctx context.Context, c *cli.Command, fn func(ctx context.Context, pool *pgxpool.Pool) error) error {
	if pool := Pool(); pool != nil {
		return fn(ctx, pool)
	}

	if connString := c.String("database-url"); connString != "" {
		pool, err := pgxpool.New(ctx, connString)
		if err != nil {
			return fmt.Errorf("connecting to postgres: %w", err)
		}
		defer pool.Close()
		return fn(ctx, pool)
	}

	pool, err := NewPgxPool()
	if err != nil {
		return fmt.Errorf("connecting to postgres: %w", err)
	}
	if pool == nil {
		return errors.New("no database found: pass --database-url / set DATABASE_URL, or run this inside a fireback app configured for postgres")
	}
	defer pool.Close()

	return fn(ctx, pool)
}

var usageCmd = &cli.Command{
	Name:  "usage",
	Usage: "Print how many bytes of storage a user or workspace has in use",
	Flags: []cli.Flag{
		databaseURLFlag(),
		&cli.StringFlag{Name: "user", Usage: "user_id to sum usage for"},
		&cli.StringFlag{Name: "workspace", Usage: "workspace_id to sum usage for"},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		userId := c.String("user")
		workspaceId := c.String("workspace")
		if (userId == "") == (workspaceId == "") {
			return errors.New("pass exactly one of --user or --workspace")
		}

		return withPool(ctx, c, func(ctx context.Context, pool *pgxpool.Pool) error {
			var (
				used int64
				err  error
			)
			if userId != "" {
				used, err = UsedBytes(ctx, pool, userId)
			} else {
				used, err = WorkspaceUsedBytes(ctx, pool, workspaceId)
			}
			if err != nil {
				return err
			}

			fmt.Printf("%d\n", used)
			return nil
		})
	},
}

var uploadCmd = &cli.Command{
	Name:      "upload",
	Usage:     "Upload a local file directly into fileupload storage",
	ArgsUsage: "<path>",
	Flags: []cli.Flag{
		databaseURLFlag(),
		&cli.StringFlag{Name: "user", Usage: "user_id to record as the upload's owner"},
		&cli.StringFlag{Name: "workspace", Usage: "workspace_id to record on the upload"},
		&cli.StringFlag{Name: "access-level", Usage: "access_level to record on the upload"},
		&cli.StringFlag{Name: "claim", Usage: `claim the upload immediately, e.g. --claim "score:abc123", so the reaper never deletes it`},
		&cli.StringSliceFlag{Name: "meta", Usage: `extra metadata as key=value, repeatable, e.g. --meta title=Invoice`},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		path := c.Args().First()
		if path == "" {
			return errors.New("usage: fileupload-cli upload [flags] <path>")
		}

		meta := map[string]string{}
		for _, kv := range c.StringSlice("meta") {
			key, value, ok := splitKV(kv)
			if !ok {
				return fmt.Errorf("--meta %q must be in key=value form", kv)
			}
			meta[key] = value
		}

		return withPool(ctx, c, func(ctx context.Context, pool *pgxpool.Pool) error {
			store := NewStore(pool)

			rec, err := UploadFile(ctx, pool, store, path, UploadFileOptions{
				UserId:      c.String("user"),
				WorkspaceId: c.String("workspace"),
				AccessLevel: c.String("access-level"),
				ClaimedBy:   c.String("claim"),
				MetaData:    meta,
			})
			if err != nil {
				return err
			}

			fmt.Printf("%s\t%d bytes\n", rec.ID, rec.Size)
			return nil
		})
	},
}

var deleteCmd = &cli.Command{
	Name:      "delete",
	Usage:     "Delete an upload by id",
	ArgsUsage: "<id>",
	Flags: []cli.Flag{
		databaseURLFlag(),
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		id := c.Args().First()
		if id == "" {
			return errors.New("usage: fileupload-cli delete <id>")
		}

		return withPool(ctx, c, func(ctx context.Context, pool *pgxpool.Pool) error {
			store := NewStore(pool)
			if err := DeleteFile(ctx, store, id); err != nil {
				return err
			}

			fmt.Printf("deleted %s\n", id)
			return nil
		})
	},
}

// splitKV splits "key=value" into its two parts. ok is false if s has no
// '=', or an empty key.
func splitKV(s string) (key, value string, ok bool) {
	for i := 0; i < len(s); i++ {
		if s[i] == '=' {
			if i == 0 {
				return "", "", false
			}
			return s[:i], s[i+1:], true
		}
	}
	return "", "", false
}
