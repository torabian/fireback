// Black-box-ish package tests for storage against a real Postgres instance -
// this module talks to Postgres directly via pgx (large objects), so there's
// nothing meaningful to mock; these exercise Store/Admin/Claim/Quota/Queries/
// Reaper exactly the way an app using this module would, no HTTP involved
// (Handler.go/Download.go's gin wiring isn't covered here).
//
// Every test skips (rather than fails) when Postgres isn't reachable, the
// same convention modules/abac/tests uses for its own server-dependent
// tests - so `go test ./modules/storage/...` stays green in environments with
// no local Postgres (large objects are a Postgres-only feature to begin
// with, see NewPgxPool's own doc comment).
package storage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tus/tusd/pkg/handler"
)

// testDSN is the Postgres connection string these tests run against.
// Override with STORAGE_TEST_POSTGRES_DSN to point at a different instance;
// the default matches this repo's own local dev Postgres (see repo root
// .env / docker-compose) with a dedicated, disposable database name -
// never point this at a database anything else cares about, TestMain
// unlinks every large object it finds there on every run.
func testDSN() string {
	if v := os.Getenv("STORAGE_TEST_POSTGRES_DSN"); v != "" {
		return v
	}
	return "postgres://changeme:changeme@localhost:6900/fireback_storage_test?sslmode=disable"
}

// TestMain gives every run of this package's tests a clean tus_uploads table
// with no leftover large objects from a previous run, best-effort - if
// Postgres isn't reachable at all, this is silently skipped and every
// individual test reports its own skip reason via newTestPool.
func TestMain(m *testing.M) {
	dsn := testDSN()
	if err := Migrate(dsn); err == nil {
		if pool, err := pgxpool.New(context.Background(), dsn); err == nil {
			ctx := context.Background()
			// lo_unlink is a plain SQL-callable function - no explicit large
			// object API transaction needed beyond the implicit one Exec
			// already runs each statement in.
			_, _ = pool.Exec(ctx, `SELECT lo_unlink(oid) FROM tus_uploads`)
			_, _ = pool.Exec(ctx, `TRUNCATE tus_uploads`)
			pool.Close()
		}
	}

	os.Exit(m.Run())
}

// newTestPool connects to testDSN, having already run migrations against it
// (see Migrate) - skips the calling test if Postgres isn't reachable there.
func newTestPool(t *testing.T) *pgxpool.Pool {
	t.Helper()

	dsn := testDSN()
	if err := Migrate(dsn); err != nil {
		t.Skipf("storage tests need a reachable Postgres at %s (override with STORAGE_TEST_POSTGRES_DSN): migrate failed: %v", dsn, err)
	}

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		t.Skipf("storage tests need a reachable Postgres at %s: %v", dsn, err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		t.Skipf("storage tests need a reachable Postgres at %s: ping failed: %v", dsn, err)
	}

	t.Cleanup(pool.Close)
	return pool
}

// createAndWrite creates a new upload of the given size, writes all of
// content to it in one WriteChunk call, and returns the resulting
// handler.Upload plus its FileInfo. Every id is a fresh random UUID, so
// tests never collide with each other even when run against a database that
// already has rows in it from a previous run.
func createAndWrite(t *testing.T, ctx context.Context, store Store, content []byte, meta handler.MetaData) (handler.Upload, handler.FileInfo) {
	t.Helper()

	if meta == nil {
		meta = handler.MetaData{}
	}

	upload, err := store.NewUpload(ctx, handler.FileInfo{
		Size:     int64(len(content)),
		MetaData: meta,
	})
	if err != nil {
		t.Fatalf("NewUpload: %v", err)
	}

	info, err := upload.GetInfo(ctx)
	if err != nil {
		t.Fatalf("GetInfo: %v", err)
	}

	if len(content) > 0 {
		n, err := upload.WriteChunk(ctx, 0, bytesReader(content))
		if err != nil {
			t.Fatalf("WriteChunk: %v", err)
		}
		if n != int64(len(content)) {
			t.Fatalf("WriteChunk wrote %d bytes, want %d", n, len(content))
		}
	}

	info, err = upload.GetInfo(ctx)
	if err != nil {
		t.Fatalf("GetInfo after write: %v", err)
	}
	return upload, info
}

func bytesReader(b []byte) io.Reader { return &byteSliceReader{b: b} }

// byteSliceReader avoids pulling in "bytes" just for a one-shot io.Reader.
type byteSliceReader struct {
	b   []byte
	pos int
}

func (r *byteSliceReader) Read(p []byte) (int, error) {
	if r.pos >= len(r.b) {
		return 0, io.EOF
	}
	n := copy(p, r.b[r.pos:])
	r.pos += n
	return n, nil
}

func readAll(t *testing.T, r io.Reader) []byte {
	t.Helper()
	data, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("reading: %v", err)
	}
	return data
}

// --- lifecycle ---------------------------------------------------------

func TestPostgres_UploadLifecycle(t *testing.T) {
	pool := newTestPool(t)
	store := NewStore(pool)
	ctx := context.Background()

	content := []byte("hello large object world")
	upload, info := createAndWrite(t, ctx, store, content, handler.MetaData{"filename": "hello.txt"})

	if info.Offset != int64(len(content)) {
		t.Fatalf("offset = %d, want %d", info.Offset, len(content))
	}

	reader, err := upload.GetReader(ctx)
	if err != nil {
		t.Fatalf("GetReader: %v", err)
	}
	got := readAll(t, reader)
	if closer, ok := reader.(io.Closer); ok {
		if err := closer.Close(); err != nil {
			t.Fatalf("closing reader: %v", err)
		}
	}
	if string(got) != string(content) {
		t.Fatalf("read back %q, want %q", got, content)
	}

	rec, err := GetFile(ctx, store, info.ID)
	if err != nil {
		t.Fatalf("GetFile: %v", err)
	}
	if rec == nil {
		t.Fatal("GetFile returned nil for an upload that was just created")
	}
	if !rec.Completed {
		t.Fatal("rec.Completed = false, want true (all declared bytes were written)")
	}
	if rec.Size != int64(len(content)) {
		t.Fatalf("rec.Size = %d, want %d", rec.Size, len(content))
	}
	if rec.MetaData["filename"] != "hello.txt" {
		t.Fatalf("rec.MetaData[filename] = %q, want %q", rec.MetaData["filename"], "hello.txt")
	}

	if err := DeleteFile(ctx, store, info.ID); err != nil {
		t.Fatalf("DeleteFile: %v", err)
	}

	if rec, err := GetFile(ctx, store, info.ID); err != nil {
		t.Fatalf("GetFile after delete: %v", err)
	} else if rec != nil {
		t.Fatal("GetFile still returns a row after DeleteFile")
	}

	if _, err := store.GetUpload(ctx, info.ID); !errors.Is(err, handler.ErrNotFound) {
		t.Fatalf("GetUpload after delete = %v, want handler.ErrNotFound", err)
	}

	if err := DeleteFile(ctx, store, info.ID); !errors.Is(err, ErrUploadNotFound) {
		t.Fatalf("DeleteFile on already-deleted id = %v, want ErrUploadNotFound", err)
	}
}

// TestPostgres_ResumableWrite writes an upload in two separate WriteChunk
// calls (as a resumed tus upload would, after a HEAD to discover the current
// offset) and checks completion only flips true once every declared byte has
// arrived, and that the two chunks come back correctly reassembled.
func TestPostgres_ResumableWrite(t *testing.T) {
	pool := newTestPool(t)
	store := NewStore(pool)
	ctx := context.Background()

	first := []byte("first-half-")
	second := []byte("second-half")
	full := append(append([]byte{}, first...), second...)

	upload, err := store.NewUpload(ctx, handler.FileInfo{Size: int64(len(full))})
	if err != nil {
		t.Fatalf("NewUpload: %v", err)
	}
	info, _ := upload.GetInfo(ctx)

	if _, err := upload.WriteChunk(ctx, 0, bytesReader(first)); err != nil {
		t.Fatalf("WriteChunk (first half): %v", err)
	}

	if rec, err := GetFile(ctx, store, info.ID); err != nil {
		t.Fatalf("GetFile mid-upload: %v", err)
	} else if rec.Completed {
		t.Fatal("rec.Completed = true after only the first half was written")
	} else if rec.Offset != int64(len(first)) {
		t.Fatalf("rec.Offset = %d, want %d", rec.Offset, len(first))
	}

	// Reload, the way a resumed client would after a HEAD request - the
	// in-memory `upload` above still has a stale info.Offset of 0.
	resumed, err := store.GetUpload(ctx, info.ID)
	if err != nil {
		t.Fatalf("GetUpload (resume): %v", err)
	}
	resumedInfo, _ := resumed.GetInfo(ctx)
	if resumedInfo.Offset != int64(len(first)) {
		t.Fatalf("resumed offset = %d, want %d", resumedInfo.Offset, len(first))
	}

	if _, err := resumed.WriteChunk(ctx, resumedInfo.Offset, bytesReader(second)); err != nil {
		t.Fatalf("WriteChunk (second half): %v", err)
	}

	rec, err := GetFile(ctx, store, info.ID)
	if err != nil {
		t.Fatalf("GetFile after both halves: %v", err)
	}
	if !rec.Completed {
		t.Fatal("rec.Completed = false after every declared byte was written")
	}

	reader, err := resumed.GetReader(ctx)
	if err != nil {
		t.Fatalf("GetReader: %v", err)
	}
	got := readAll(t, reader)
	if closer, ok := reader.(io.Closer); ok {
		closer.Close()
	}
	if string(got) != string(full) {
		t.Fatalf("reassembled content = %q, want %q", got, full)
	}
}

// TestPostgres_OpenRange checks Store.OpenRange (used by Download.go's Range
// support) seeks correctly into an already-completed upload's large object.
func TestPostgres_OpenRange(t *testing.T) {
	pool := newTestPool(t)
	store := NewStore(pool)
	ctx := context.Background()

	content := []byte("0123456789")
	_, info := createAndWrite(t, ctx, store, content, nil)

	r, err := store.OpenRange(ctx, info.ID, 3)
	if err != nil {
		t.Fatalf("OpenRange: %v", err)
	}
	got := readAll(t, r)
	r.Close()

	if string(got) != "3456789" {
		t.Fatalf("OpenRange(offset=3) = %q, want %q", got, "3456789")
	}

	if _, err := store.OpenRange(ctx, "does-not-exist", 0); !errors.Is(err, handler.ErrNotFound) {
		t.Fatalf("OpenRange on missing id = %v, want handler.ErrNotFound", err)
	}
}

// --- claiming ------------------------------------------------------------

func TestPostgres_ClaimAndRelease(t *testing.T) {
	pool := newTestPool(t)
	store := NewStore(pool)
	ctx := context.Background()

	_, info := createAndWrite(t, ctx, store, []byte("claim me"), nil)

	rec, err := ClaimFile(ctx, store, info.ID, "score:abc123")
	if err != nil {
		t.Fatalf("ClaimFile: %v", err)
	}
	if rec.ClaimedBy == nil || *rec.ClaimedBy != "score:abc123" {
		t.Fatalf("ClaimedBy = %v, want %q", rec.ClaimedBy, "score:abc123")
	}
	if rec.ClaimedAt == nil {
		t.Fatal("ClaimedAt is nil after a successful claim")
	}

	// Claiming again with the same owner is a no-op success.
	if _, err := ClaimFile(ctx, store, info.ID, "score:abc123"); err != nil {
		t.Fatalf("re-claiming with the same owner: %v", err)
	}

	if err := ReleaseFile(ctx, store, info.ID, "score:abc123"); err != nil {
		t.Fatalf("ReleaseFile: %v", err)
	}
	rec, err = GetFile(ctx, store, info.ID)
	if err != nil {
		t.Fatalf("GetFile after release: %v", err)
	}
	if rec.ClaimedBy != nil {
		t.Fatalf("ClaimedBy = %v after ReleaseFile, want nil", rec.ClaimedBy)
	}

	if _, err := ClaimFile(ctx, store, "does-not-exist", "x"); !errors.Is(err, ErrUploadNotFound) {
		t.Fatalf("ClaimFile on missing id = %v, want ErrUploadNotFound", err)
	}

	incomplete, err := store.NewUpload(ctx, handler.FileInfo{Size: 10})
	if err != nil {
		t.Fatalf("NewUpload (incomplete): %v", err)
	}
	incompleteInfo, _ := incomplete.GetInfo(ctx)
	if _, err := ClaimFile(ctx, store, incompleteInfo.ID, "x"); !errors.Is(err, ErrUploadNotCompleted) {
		t.Fatalf("ClaimFile on an incomplete upload = %v, want ErrUploadNotCompleted", err)
	}
}

// ReleaseFile clearing someone else's claim should be a no-op, not an error
// and not an actual release - one owner can't release another's claim.
func TestPostgres_ReleaseFile_WrongOwnerIsNoop(t *testing.T) {
	pool := newTestPool(t)
	store := NewStore(pool)
	ctx := context.Background()

	_, info := createAndWrite(t, ctx, store, []byte("x"), nil)
	if _, err := ClaimFile(ctx, store, info.ID, "owner-a"); err != nil {
		t.Fatalf("ClaimFile: %v", err)
	}

	if err := ReleaseFile(ctx, store, info.ID, "owner-b"); err != nil {
		t.Fatalf("ReleaseFile (wrong owner): %v", err)
	}

	rec, err := GetFile(ctx, store, info.ID)
	if err != nil {
		t.Fatalf("GetFile: %v", err)
	}
	if rec.ClaimedBy == nil || *rec.ClaimedBy != "owner-a" {
		t.Fatalf("ClaimedBy = %v, want still %q (release by a different owner must be a no-op)", rec.ClaimedBy, "owner-a")
	}
}

// --- quota / usage ---------------------------------------------------------

func TestPostgres_UsedBytes(t *testing.T) {
	pool := newTestPool(t)
	store := NewStore(pool)
	ctx := context.Background()

	userId := "user-" + uuid.NewString()
	sizes := []int{100, 250}
	var total int64
	for _, size := range sizes {
		content := make([]byte, size)
		upload, err := store.NewUpload(ctx, handler.FileInfo{
			Size:     int64(size),
			MetaData: handler.MetaData{metaKeyUserId: userId},
		})
		if err != nil {
			t.Fatalf("NewUpload: %v", err)
		}
		if _, err := upload.WriteChunk(ctx, 0, bytesReader(content)); err != nil {
			t.Fatalf("WriteChunk: %v", err)
		}
		total += int64(size)
	}

	used, err := UsedBytes(ctx, store, userId)
	if err != nil {
		t.Fatalf("UsedBytes: %v", err)
	}
	if used != total {
		t.Fatalf("UsedBytes = %d, want %d", used, total)
	}

	if used, err := UsedBytes(ctx, store, "user-with-nothing-"+uuid.NewString()); err != nil {
		t.Fatalf("UsedBytes (no uploads): %v", err)
	} else if used != 0 {
		t.Fatalf("UsedBytes (no uploads) = %d, want 0", used)
	}
}

func TestPostgres_WorkspaceUsedBytes(t *testing.T) {
	pool := newTestPool(t)
	store := NewStore(pool)
	ctx := context.Background()

	workspaceId := "workspace-" + uuid.NewString()
	content := make([]byte, 512)
	upload, err := store.NewUpload(ctx, handler.FileInfo{
		Size:     int64(len(content)),
		MetaData: handler.MetaData{metaKeyWorkspaceId: workspaceId},
	})
	if err != nil {
		t.Fatalf("NewUpload: %v", err)
	}
	if _, err := upload.WriteChunk(ctx, 0, bytesReader(content)); err != nil {
		t.Fatalf("WriteChunk: %v", err)
	}

	used, err := WorkspaceUsedBytes(ctx, store, workspaceId)
	if err != nil {
		t.Fatalf("WorkspaceUsedBytes: %v", err)
	}
	if used != int64(len(content)) {
		t.Fatalf("WorkspaceUsedBytes = %d, want %d", used, len(content))
	}
}

// --- admin (CLI/internal path) --------------------------------------------

func TestPostgres_Admin_UploadFile_DeleteFile(t *testing.T) {
	pool := newTestPool(t)
	store := NewStore(pool)
	ctx := context.Background()

	dir := t.TempDir()
	path := filepath.Join(dir, "report.csv")
	content := []byte("a,b,c\n1,2,3\n")
	if err := os.WriteFile(path, content, 0o600); err != nil {
		t.Fatalf("writing temp file: %v", err)
	}

	rec, err := UploadFile(ctx, store, path, UploadFileOptions{
		UserId:      "admin-user",
		WorkspaceId: "admin-workspace",
		ClaimedBy:   "test:admin-upload",
	})
	if err != nil {
		t.Fatalf("UploadFile: %v", err)
	}
	if rec.Size != int64(len(content)) {
		t.Fatalf("rec.Size = %d, want %d", rec.Size, len(content))
	}
	if rec.MetaData["filename"] != "report.csv" {
		t.Fatalf("rec.MetaData[filename] = %q, want %q", rec.MetaData["filename"], "report.csv")
	}
	if rec.ClaimedBy == nil || *rec.ClaimedBy != "test:admin-upload" {
		t.Fatalf("ClaimedBy = %v, want %q", rec.ClaimedBy, "test:admin-upload")
	}
	if rec.UserId == nil || *rec.UserId != "admin-user" {
		t.Fatalf("UserId = %v, want %q", rec.UserId, "admin-user")
	}

	if err := DeleteFile(ctx, store, rec.ID); err != nil {
		t.Fatalf("DeleteFile: %v", err)
	}
}

func TestPostgres_Admin_UploadFile_RejectsDirectory(t *testing.T) {
	pool := newTestPool(t)
	store := NewStore(pool)
	ctx := context.Background()

	if _, err := UploadFile(ctx, store, t.TempDir(), UploadFileOptions{}); err == nil {
		t.Fatal("UploadFile on a directory should return an error")
	}
}

// --- listing ---------------------------------------------------------------

func TestPostgres_ListFiles(t *testing.T) {
	pool := newTestPool(t)
	store := NewStore(pool)
	ctx := context.Background()

	var ids []string
	for i := range 3 {
		_, info := createAndWrite(t, ctx, store, fmt.Appendf(nil, "file-%d", i), nil)
		ids = append(ids, info.ID)
		time.Sleep(5 * time.Millisecond) // keep created_at strictly increasing
	}

	files, err := ListFiles(ctx, store, len(ids)+10, 0)
	if err != nil {
		t.Fatalf("ListFiles: %v", err)
	}

	seen := map[string]bool{}
	for _, f := range files {
		seen[f.ID] = true
	}
	for _, id := range ids {
		if !seen[id] {
			t.Fatalf("ListFiles didn't include upload %s", id)
		}
	}

	// Newest first.
	page, err := ListFiles(ctx, store, 1, 0)
	if err != nil {
		t.Fatalf("ListFiles (limit 1): %v", err)
	}
	if len(page) != 1 || page[0].ID != ids[len(ids)-1] {
		t.Fatalf("ListFiles(limit=1) = %+v, want the most recently created upload (%s) first", page, ids[len(ids)-1])
	}
}

// --- reaper ------------------------------------------------------------

// backdate directly rewrites completed_at/updated_at so SweepOrphaned's TTL
// checks see an upload as older than it actually is, without having to
// actually sleep in the test.
func backdate(t *testing.T, ctx context.Context, pool *pgxpool.Pool, id string, age time.Duration) {
	t.Helper()
	past := time.Now().Add(-age)
	if _, err := pool.Exec(ctx, `UPDATE tus_uploads SET completed_at = $1, updated_at = $1 WHERE id = $2`, past, id); err != nil {
		t.Fatalf("backdating %s: %v", id, err)
	}
}

func TestPostgres_SweepOrphaned(t *testing.T) {
	pool := newTestPool(t)
	store := NewStore(pool)
	ctx := context.Background()

	const ttl = time.Hour

	// 1. Completed, unclaimed, old -> swept.
	_, orphanInfo := createAndWrite(t, ctx, store, []byte("orphan"), nil)
	backdate(t, ctx, pool, orphanInfo.ID, 2*ttl)

	// 2. Completed, unclaimed, fresh -> kept (hasn't aged out yet).
	_, freshInfo := createAndWrite(t, ctx, store, []byte("fresh"), nil)

	// 3. Completed, claimed, old -> kept (claim protects it regardless of age).
	_, claimedInfo := createAndWrite(t, ctx, store, []byte("claimed"), nil)
	if _, err := ClaimFile(ctx, store, claimedInfo.ID, "test:keep-me"); err != nil {
		t.Fatalf("ClaimFile: %v", err)
	}
	backdate(t, ctx, pool, claimedInfo.ID, 2*ttl)

	// 4. Never completed, old -> swept as abandoned.
	incompleteUpload, err := store.NewUpload(ctx, handler.FileInfo{Size: 100})
	if err != nil {
		t.Fatalf("NewUpload (incomplete): %v", err)
	}
	incompleteInfo, _ := incompleteUpload.GetInfo(ctx)
	backdate(t, ctx, pool, incompleteInfo.ID, 2*ttl)

	// 5. Never completed, fresh -> kept (still within incompleteTTL).
	freshIncomplete, err := store.NewUpload(ctx, handler.FileInfo{Size: 100})
	if err != nil {
		t.Fatalf("NewUpload (fresh incomplete): %v", err)
	}
	freshIncompleteInfo, _ := freshIncomplete.GetInfo(ctx)

	deleted, err := SweepOrphaned(ctx, store, ttl, ttl)
	if err != nil {
		t.Fatalf("SweepOrphaned: %v", err)
	}

	deletedSet := map[string]bool{}
	for _, id := range deleted {
		deletedSet[id] = true
	}

	if !deletedSet[orphanInfo.ID] {
		t.Errorf("SweepOrphaned did not delete the old unclaimed upload %s", orphanInfo.ID)
	}
	if !deletedSet[incompleteInfo.ID] {
		t.Errorf("SweepOrphaned did not delete the old incomplete upload %s", incompleteInfo.ID)
	}
	if deletedSet[freshInfo.ID] {
		t.Errorf("SweepOrphaned deleted a fresh unclaimed upload %s", freshInfo.ID)
	}
	if deletedSet[claimedInfo.ID] {
		t.Errorf("SweepOrphaned deleted a claimed upload %s", claimedInfo.ID)
	}
	if deletedSet[freshIncompleteInfo.ID] {
		t.Errorf("SweepOrphaned deleted a fresh incomplete upload %s", freshIncompleteInfo.ID)
	}

	// The deleted rows' large objects must be gone too, not just the row -
	// GetFile (a plain row lookup) can't tell us that; GetUpload trying to
	// re-read the row already confirms the row itself is gone.
	if rec, err := GetFile(ctx, store, orphanInfo.ID); err != nil {
		t.Fatalf("GetFile after sweep: %v", err)
	} else if rec != nil {
		t.Error("orphaned upload's row still exists after SweepOrphaned")
	}

	// Cleanup the survivors so they don't linger for the next run (best
	// effort; TestMain will also unlink everything on the next invocation).
	for _, id := range []string{freshInfo.ID, claimedInfo.ID, freshIncompleteInfo.ID} {
		_ = DeleteFile(ctx, store, id)
	}
}
