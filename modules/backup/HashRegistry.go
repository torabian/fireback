package backup

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/torabian/fireback/modules/fireback"
)

// dumpJob is a pending "backup dump --hash" request: dump database on the
// first (and only) fetch of its hash, not before - see DumpHttp.go's GET
// handler.
//
// Persisted as a small JSON file under jobStoreDir() rather than kept only
// in one process's memory: `backup dump --hash` is typically a short-lived
// CLI invocation, completely separate from the long-running server process
// that later serves the GET fetch (see DumpHttp.go) - there's no shared
// memory between them, only a shared filesystem. This is also what makes
// registering a job no longer need its own auth check (there used to be a
// BACKUP_API_TOKEN bearer requirement here) - the OS-level privacy of
// jobStoreDir() *is* the access control now: whoever can write there is
// already as trusted as whoever runs `backup dump` locally.
type dumpJob struct {
	Database  string    `json:"database"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}

// jobStoreDir returns the OS-appropriate, per-user-private directory
// pending dump jobs are persisted under. os.UserConfigDir() already
// resolves this per OS - Linux: $XDG_CONFIG_HOME or ~/.config; macOS:
// ~/Library/Application Support; Windows: %AppData% - each private to the
// current OS user by that platform's own convention/default ACLs, which is
// exactly the "store it somewhere in system safely, depending on OS"
// this needs, without hand-rolling a per-OS path table.
//
// Because this is scoped to the OS user, the CLI process registering a job
// and the server process later serving it must run on the same host as the
// same OS user - true for every deployment this module already assumes
// elsewhere (wal-g's own push/restore have the same "co-located"
// requirement, see README.md).
func jobStoreDir() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("resolving this OS's per-user config directory for pending dump jobs: %w", err)
	}
	dir := filepath.Join(base, "fireback", "backup-jobs")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return "", fmt.Errorf("creating %s: %w", dir, err)
	}
	return dir, nil
}

func jobPath(dir, hash string) string {
	return filepath.Join(dir, hash+".json")
}

// registerDumpJob persists a new single-use job for database, valid for
// ttl, and returns the hash it's filed under.
func registerDumpJob(database string, ttl time.Duration) (string, error) {
	dir, err := jobStoreDir()
	if err != nil {
		return "", err
	}

	// Opportunistic, so the store can't grow unbounded from links nobody
	// ever fetched, without needing a background goroutine/cron of its own.
	sweepExpiredJobs(dir)

	hash := fireback.GenerateSecureToken(48)
	now := time.Now()
	job := dumpJob{Database: database, CreatedAt: now, ExpiresAt: now.Add(ttl)}

	data, err := json.Marshal(job)
	if err != nil {
		return "", err
	}
	if err := os.WriteFile(jobPath(dir, hash), data, 0o600); err != nil {
		return "", fmt.Errorf("persisting dump job: %w", err)
	}
	return hash, nil
}

// claimDumpJob reads and immediately deletes hash's job file - single-use:
// a hash is disabled the moment it's fetched, successfully or not. A
// missing, corrupt, or expired file are all treated identically (not
// found) - deliberately, so there's nothing for a caller to probe to tell
// a stale hash from a bad one.
func claimDumpJob(hash string) (*dumpJob, bool) {
	dir, err := jobStoreDir()
	if err != nil {
		return nil, false
	}

	path := jobPath(dir, hash)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, false
	}
	os.Remove(path) // disabled now, regardless of what's below

	var job dumpJob
	if err := json.Unmarshal(data, &job); err != nil {
		return nil, false
	}
	if time.Now().After(job.ExpiresAt) {
		return nil, false
	}
	return &job, true
}

func sweepExpiredJobs(dir string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	now := time.Now()
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		path := filepath.Join(dir, e.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		var job dumpJob
		if err := json.Unmarshal(data, &job); err != nil {
			os.Remove(path) // not a job file this version recognizes - stray/corrupt, safe to drop
			continue
		}
		if now.After(job.ExpiresAt) {
			os.Remove(path)
		}
	}
}
