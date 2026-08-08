package backup

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Engine drives wal-g by re-executing this same binary with a hidden
// "__walg-exec" marker argument (see MaybeRunEmbeddedWalg in Exec.go),
// rather than shelling out to a separately installed wal-g binary. wal-g is
// a real go.mod dependency (github.com/wal-g/wal-g/cmd/pg); the re-exec
// exists only because its cobra command tree isn't safe to invoke more than
// once per process (see Exec.go), so every operation needs its own fresh
// process - this is what gives it one, using no binary but this one.
type Engine struct {
	cfg *ModuleConfig
}

func NewEngine(cfg *ModuleConfig) *Engine {
	return &Engine{cfg: cfg}
}

func (e *Engine) env() []string {
	env := os.Environ()
	env = append(env,
		"WALG_FILE_PREFIX="+e.cfg.FilePrefix,
		"WALG_COMPRESSION_METHOD="+e.cfg.CompressionMethod,
		"WALG_DELTA_MAX_STEPS="+e.cfg.DeltaMaxSteps,
		"PGHOST="+e.cfg.PgHost,
		"PGPORT="+e.cfg.PgPort,
		"PGUSER="+e.cfg.PgUser,
		"PGPASSWORD="+e.cfg.PgPassword,
		"PGDATABASE="+e.cfg.PgDatabase,
	)
	return env
}

// command builds the re-exec: this same binary, invoked with the hidden
// "__walg-exec" marker followed by wal-g's own args - MaybeRunEmbeddedWalg
// recognizes the marker and hands args straight to wal-g's command tree.
func (e *Engine) command(ctx context.Context, args ...string) (*exec.Cmd, error) {
	exePath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("resolving own executable to run wal-g %s: %w", strings.Join(args, " "), err)
	}
	cmd := exec.CommandContext(ctx, exePath, append([]string{walgExecMarker}, args...)...)
	cmd.Env = e.env()
	return cmd, nil
}

// Run invokes wal-g with args, streaming its stdout/stderr through to ours -
// base backups and restores can run long, so progress needs to be visible
// rather than buffered and dumped at the end.
func (e *Engine) Run(ctx context.Context, args ...string) error {
	cmd, err := e.command(ctx, args...)
	if err != nil {
		return err
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("wal-g %s: %w", strings.Join(args, " "), err)
	}
	return nil
}

// RunJSON invokes wal-g with args and decodes its stdout as JSON into v.
// This is only used for backup-list, whose --json output (BackupTime in
// wal-g's own source) is a small, stable shape - everything else goes
// through Run so its native progress/error output reaches the operator
// unmodified. wal-g's own INFO/ERROR logging goes to stderr (confirmed:
// only the JSON payload itself lands on stdout), so this and Run's direct
// passthrough never risk mixing the two.
func (e *Engine) RunJSON(ctx context.Context, v any, args ...string) error {
	cmd, err := e.command(ctx, args...)
	if err != nil {
		return err
	}
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("wal-g %s: %w: %s", strings.Join(args, " "), err, stderr.String())
	}
	if err := json.Unmarshal(stdout.Bytes(), v); err != nil {
		return fmt.Errorf("parsing wal-g %s output: %w", strings.Join(args, " "), err)
	}
	return nil
}

// BackupInfo mirrors wal-g's own BackupTime struct (internal/backup_time.go)
// as returned by `backup-list --json` - deliberately not `--detail`, whose
// JSON shape has changed across wal-g releases.
type BackupInfo struct {
	BackupName  string    `json:"backup_name"`
	Time        time.Time `json:"time"`
	WalFileName string    `json:"wal_file_name"`
	StorageName string    `json:"storage_name"`
}

// BackupList returns every backup wal-g knows about, oldest first (wal-g's
// own ordering).
func (e *Engine) BackupList(ctx context.Context) ([]BackupInfo, error) {
	var out []BackupInfo
	if err := e.RunJSON(ctx, &out, "backup-list", "--json"); err != nil {
		return nil, err
	}
	return out, nil
}

// BackupPush triggers a base backup of pgDataDir. Whether it becomes a full
// or delta backup is governed by WALG_DELTA_MAX_STEPS, unless full is true,
// which forces `--full-backup` regardless.
func (e *Engine) BackupPush(ctx context.Context, pgDataDir string, full bool) error {
	args := []string{"backup-push", pgDataDir}
	if full {
		args = append(args, "--full")
	}
	return e.Run(ctx, args...)
}

// BackupFetch extracts the named backup (or "LATEST") into destDir, which
// wal-g creates if it doesn't already exist.
func (e *Engine) BackupFetch(ctx context.Context, destDir, name string) error {
	return e.Run(ctx, "backup-fetch", destDir, name)
}

// WalFetch downloads a single WAL segment - this is the restore_command
// entry point during recovery, not something the CLI calls directly, but
// it's exposed here so ArchiveGet (Cli.go) can wrap it.
func (e *Engine) WalFetch(ctx context.Context, walName, destPath string) error {
	return e.Run(ctx, "wal-fetch", walName, destPath)
}

// WalPush uploads a single WAL segment - the archive_command entry point.
func (e *Engine) WalPush(ctx context.Context, walPath string) error {
	return e.Run(ctx, "wal-push", walPath)
}

// DeleteRetainFull keeps the newest n full backups (and everything needed to
// restore from the oldest of them forward - deltas and WAL), deleting
// everything older. --confirm is required or wal-g only dry-runs.
func (e *Engine) DeleteRetainFull(ctx context.Context, n int) error {
	return e.Run(ctx, "delete", "retain", "FULL", fmt.Sprintf("%d", n), "--confirm")
}

// CopyBackup uses wal-g's own `copy` command to produce a self-contained,
// restorable copy of backupName - its base backup plus WAL - at destPrefix,
// without decrypting/recompressing/re-encrypting anything. This is what
// `backup download` uses: the result is plain files under destPrefix that
// can be moved anywhere (offsite, another host) and restored from directly
// with `backup-fetch`/`wal-fetch` pointed at that prefix.
//
// Without withHistory, only the WAL needed to bring backupName itself to a
// consistent state is copied - restoring it lands you at that backup's own
// completion instant, not anything after. With withHistory, every WAL
// segment archived from backupName's recovery point through "now" (when
// this command runs) comes along too - the copy is then restorable to any
// instant up to when you made it, which is what "download a backup and be
// able to recreate the database" actually needs.
//
// wal-g's `copy` takes storage configs as JSON files rather than env vars,
// so this writes two throwaway ones (each just {"WALG_FILE_PREFIX": ...})
// and cleans them up afterwards.
func (e *Engine) CopyBackup(ctx context.Context, backupName, destPrefix string, withHistory bool) error {
	fromPath, err := writeTempStorageConfig(e.cfg.FilePrefix)
	if err != nil {
		return fmt.Errorf("writing source storage config: %w", err)
	}
	defer os.Remove(fromPath)

	toPath, err := writeTempStorageConfig(destPrefix)
	if err != nil {
		return fmt.Errorf("writing destination storage config: %w", err)
	}
	defer os.Remove(toPath)

	args := []string{"copy", "--from=" + fromPath, "--to=" + toPath, "--backup-name=" + backupName}
	if withHistory {
		args = append(args, "--with-history")
	}
	return e.Run(ctx, args...)
}

func writeTempStorageConfig(filePrefix string) (string, error) {
	f, err := os.CreateTemp("", "nima-walg-storage-*.json")
	if err != nil {
		return "", err
	}
	defer f.Close()

	data, err := json.Marshal(map[string]string{"WALG_FILE_PREFIX": filePrefix})
	if err != nil {
		return "", err
	}
	if _, err := f.Write(data); err != nil {
		return "", err
	}
	return f.Name(), nil
}
