// Package tests for storage's mysql backend (StoreMySQL.go) - the same
// scenarios storage_postgres_test.go/storage_sqlite_test.go exercise,
// reusing their generic Store-interface helpers (createAndWrite,
// bytesReader, readAll) plus the sqlite test file's fillPattern/chunkRowCount
// (both already vendor-agnostic - chunkRowCount just needs a *sql.DB, and
// MySQLStore's DB field is the same type).
//
// Every test skips (rather than fails) when MySQL isn't reachable, the same
// convention storage_postgres_test.go/modules/abac/tests use for their own
// server-dependent tests.
package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/tus/tusd/pkg/handler"
)

// mysqlTestDSN is the MySQL/MariaDB connection string these tests run
// against. Override with STORAGE_TEST_MYSQL_DSN to point at a different
// instance; the default matches a local MySQL with no password on root
// (common for a local dev instance) - never point this at a database
// anything else cares about, each test truncates tus_uploads/
// tus_upload_chunks it doesn't need across runs since ids are always fresh
// UUIDs, but the schema itself is shared across the whole test binary run.
func mysqlTestDSN() string {
	if v := os.Getenv("STORAGE_TEST_MYSQL_DSN"); v != "" {
		return v
	}
	return "root@tcp(127.0.0.1:3306)/fireback_storage_test?charset=utf8mb4&parseTime=True&loc=Local"
}

// newTestMySQLStore connects to mysqlTestDSN, having already run migrations
// against it (see MigrateMySQL) - skips the calling test if MySQL isn't
// reachable there.
func newTestMySQLStore(t *testing.T) *MySQLStore {
	t.Helper()

	dsn := mysqlTestDSN()
	if err := MigrateMySQL(dsn); err != nil {
		t.Skipf("storage tests need a reachable MySQL at %s (override with STORAGE_TEST_MYSQL_DSN): migrate failed: %v", dsn, err)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Skipf("storage tests need a reachable MySQL at %s: %v", dsn, err)
	}

	if err := db.PingContext(context.Background()); err != nil {
		db.Close()
		t.Skipf("storage tests need a reachable MySQL at %s: ping failed: %v", dsn, err)
	}

	t.Cleanup(func() { db.Close() })
	return &MySQLStore{DB: db}
}

func mysqlChunkRowCount(t *testing.T, db *sql.DB, uploadId string) int {
	t.Helper()
	var n int
	if err := db.QueryRow(`SELECT COUNT(*) FROM tus_upload_chunks WHERE upload_id = ?`, uploadId).Scan(&n); err != nil {
		t.Fatalf("counting chunk rows: %v", err)
	}
	return n
}

// --- lifecycle ---------------------------------------------------------

func TestMySQL_UploadLifecycle(t *testing.T) {
	store := newTestMySQLStore(t)
	ctx := context.Background()

	content := []byte("hello mysql chunks world")
	upload, info := createAndWrite(t, ctx, store, content, handler.MetaData{"filename": "hello.txt"})

	if info.Offset != int64(len(content)) {
		t.Fatalf("offset = %d, want %d", info.Offset, len(content))
	}

	reader, err := upload.GetReader(ctx)
	if err != nil {
		t.Fatalf("GetReader: %v", err)
	}
	got := readAll(t, reader)
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
	if mysqlChunkRowCount(t, store.DB, info.ID) != 0 {
		t.Fatal("tus_upload_chunks rows still exist after DeleteFile")
	}

	if _, err := store.GetUpload(ctx, info.ID); !errors.Is(err, handler.ErrNotFound) {
		t.Fatalf("GetUpload after delete = %v, want handler.ErrNotFound", err)
	}

	if err := DeleteFile(ctx, store, info.ID); !errors.Is(err, ErrUploadNotFound) {
		t.Fatalf("DeleteFile on already-deleted id = %v, want ErrUploadNotFound", err)
	}
}

// TestMySQL_ResumableWrite_AcrossPageBoundary mirrors the sqlite test of
// the same name - CONCAT(data, ?) instead of `||`, otherwise identical page
// math (chunkSize is shared package-wide).
func TestMySQL_ResumableWrite_AcrossPageBoundary(t *testing.T) {
	store := newTestMySQLStore(t)
	ctx := context.Background()

	firstLen := chunkSize - 100
	secondLen := 300
	full := fillPattern(firstLen + secondLen)

	upload, err := store.NewUpload(ctx, handler.FileInfo{Size: int64(len(full))})
	if err != nil {
		t.Fatalf("NewUpload: %v", err)
	}
	info, _ := upload.GetInfo(ctx)

	if _, err := upload.WriteChunk(ctx, 0, bytesReader(full[:firstLen])); err != nil {
		t.Fatalf("WriteChunk (first): %v", err)
	}

	resumed, err := store.GetUpload(ctx, info.ID)
	if err != nil {
		t.Fatalf("GetUpload (resume): %v", err)
	}
	resumedInfo, _ := resumed.GetInfo(ctx)
	if resumedInfo.Offset != int64(firstLen) {
		t.Fatalf("resumed offset = %d, want %d", resumedInfo.Offset, firstLen)
	}

	if _, err := resumed.WriteChunk(ctx, resumedInfo.Offset, bytesReader(full[firstLen:])); err != nil {
		t.Fatalf("WriteChunk (second): %v", err)
	}

	if got := mysqlChunkRowCount(t, store.DB, info.ID); got != 2 {
		t.Fatalf("chunk row count = %d, want 2 (one full page, one spillover)", got)
	}

	reader, err := resumed.GetReader(ctx)
	if err != nil {
		t.Fatalf("GetReader: %v", err)
	}
	got := readAll(t, reader)
	if string(got) != string(full) {
		t.Fatal("reassembled content across the page boundary doesn't match what was written")
	}
}

func TestMySQL_MultiPageWrite_And_OpenRange(t *testing.T) {
	store := newTestMySQLStore(t)
	ctx := context.Background()

	full := fillPattern(chunkSize*2 + 500)
	_, info := createAndWrite(t, ctx, store, full, nil)

	if got := mysqlChunkRowCount(t, store.DB, info.ID); got != 3 {
		t.Fatalf("chunk row count = %d, want 3 (two full pages + a 500-byte spillover)", got)
	}

	offset := int64(chunkSize + 100)
	r, err := store.OpenRange(ctx, info.ID, offset)
	if err != nil {
		t.Fatalf("OpenRange: %v", err)
	}
	got := readAll(t, r)
	r.Close()

	if string(got) != string(full[offset:]) {
		t.Fatal("OpenRange from mid-page-1 through page-2 doesn't match the source content")
	}

	if _, err := store.OpenRange(ctx, "does-not-exist", 0); !errors.Is(err, handler.ErrNotFound) {
		t.Fatalf("OpenRange on missing id = %v, want handler.ErrNotFound", err)
	}
}

func TestMySQL_WriteChunk_RejectsOffsetMismatch(t *testing.T) {
	store := newTestMySQLStore(t)
	ctx := context.Background()

	upload, err := store.NewUpload(ctx, handler.FileInfo{Size: 100})
	if err != nil {
		t.Fatalf("NewUpload: %v", err)
	}

	if _, err := upload.WriteChunk(ctx, 42, bytesReader([]byte("x"))); err == nil {
		t.Fatal("WriteChunk at a non-zero offset on a brand new upload should be rejected")
	}
}

// --- claiming ------------------------------------------------------------

func TestMySQL_ClaimAndRelease(t *testing.T) {
	store := newTestMySQLStore(t)
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

func TestMySQL_ReleaseFile_WrongOwnerIsNoop(t *testing.T) {
	store := newTestMySQLStore(t)
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

func TestMySQL_UsedBytes(t *testing.T) {
	store := newTestMySQLStore(t)
	ctx := context.Background()

	userId := "user-" + uuid.NewString()
	sizes := []int{100, 250}
	var total int64
	for _, size := range sizes {
		upload, err := store.NewUpload(ctx, handler.FileInfo{
			Size:     int64(size),
			MetaData: handler.MetaData{metaKeyUserId: userId},
		})
		if err != nil {
			t.Fatalf("NewUpload: %v", err)
		}
		if _, err := upload.WriteChunk(ctx, 0, bytesReader(make([]byte, size))); err != nil {
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

func TestMySQL_WorkspaceUsedBytes(t *testing.T) {
	store := newTestMySQLStore(t)
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

func TestMySQL_Admin_UploadFile_DeleteFile(t *testing.T) {
	store := newTestMySQLStore(t)
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

func TestMySQL_Admin_UploadFile_RejectsDirectory(t *testing.T) {
	store := newTestMySQLStore(t)
	ctx := context.Background()

	if _, err := UploadFile(ctx, store, t.TempDir(), UploadFileOptions{}); err == nil {
		t.Fatal("UploadFile on a directory should return an error")
	}
}

// --- listing ---------------------------------------------------------------

func TestMySQL_ListFiles(t *testing.T) {
	store := newTestMySQLStore(t)
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

	page, err := ListFiles(ctx, store, 1, 0)
	if err != nil {
		t.Fatalf("ListFiles (limit 1): %v", err)
	}
	if len(page) != 1 || page[0].ID != ids[len(ids)-1] {
		t.Fatalf("ListFiles(limit=1) = %+v, want the most recently created upload (%s) first", page, ids[len(ids)-1])
	}
}

// --- reaper ------------------------------------------------------------

func backdateMySQL(t *testing.T, db *sql.DB, id string, age time.Duration) {
	t.Helper()
	past := formatMySQLTime(time.Now().Add(-age))
	if _, err := db.Exec(`UPDATE tus_uploads SET completed_at = ?, updated_at = ? WHERE id = ?`, past, past, id); err != nil {
		t.Fatalf("backdating %s: %v", id, err)
	}
}

func TestMySQL_SweepOrphaned(t *testing.T) {
	store := newTestMySQLStore(t)
	ctx := context.Background()

	const ttl = time.Hour

	_, orphanInfo := createAndWrite(t, ctx, store, []byte("orphan"), nil)
	backdateMySQL(t, store.DB, orphanInfo.ID, 2*ttl)

	_, freshInfo := createAndWrite(t, ctx, store, []byte("fresh"), nil)

	_, claimedInfo := createAndWrite(t, ctx, store, []byte("claimed"), nil)
	if _, err := ClaimFile(ctx, store, claimedInfo.ID, "test:keep-me"); err != nil {
		t.Fatalf("ClaimFile: %v", err)
	}
	backdateMySQL(t, store.DB, claimedInfo.ID, 2*ttl)

	incompleteUpload, err := store.NewUpload(ctx, handler.FileInfo{Size: 100})
	if err != nil {
		t.Fatalf("NewUpload (incomplete): %v", err)
	}
	incompleteInfo, _ := incompleteUpload.GetInfo(ctx)
	backdateMySQL(t, store.DB, incompleteInfo.ID, 2*ttl)

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

	if rec, err := GetFile(ctx, store, orphanInfo.ID); err != nil {
		t.Fatalf("GetFile after sweep: %v", err)
	} else if rec != nil {
		t.Error("orphaned upload's row still exists after SweepOrphaned")
	}
	if mysqlChunkRowCount(t, store.DB, orphanInfo.ID) != 0 {
		t.Error("orphaned upload's chunk rows still exist after SweepOrphaned")
	}
}
