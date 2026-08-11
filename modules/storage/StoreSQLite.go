// StoreSQLite.go implements Store (see Store.go) on top of sqlite: a
// tus_uploads bookkeeping table (schema mirrors Postgres', see
// migrations_sqlite/00001_create_tus_uploads.sql's own comment for the
// differences) plus tus_upload_chunks, which holds the actual bytes as
// fixed-size pages aligned to absolute offset - sqlite has no large object
// API (lo_*) to hold a whole upload as one growable value, and its
// incremental BLOB I/O (sqlite3_blob_open) requires a pre-sized zeroblob,
// which breaks for tus's deferred-length uploads. A chunks table sidesteps
// both problems and maps 1:1 onto tus's PATCH-based write model: each
// WriteChunk call appends into (or starts) whichever page its offset falls
// in.
package storage

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"
	"github.com/tus/tusd/pkg/handler"
)

// chunkSize is the fixed page size tus_upload_chunks rows are aligned to.
// Chosen as a middle ground: large enough to keep the row count per upload
// reasonable for realistically sized files (a 1 GiB upload is 256 rows, not
// one row per tus PATCH some client happened to send), small enough that a
// page read/write never has to move an unreasonable amount of memory at
// once. A given row's data can be anywhere from 1 byte up to chunkSize bytes
// long - rows aren't zero-padded, they just grow (via SQLite's blob ||
// concatenation) as consecutive WriteChunk calls land in the same page,
// until the page fills up and the next WriteChunk starts a new row.
const chunkSize = 4 << 20 // 4 MiB

// sqliteTimeLayout is a fixed-width (always 9 fractional digits, unlike
// time.RFC3339Nano which trims trailing zeros) timestamp format, so that
// plain lexical string comparison of two formatted values (what
// orphanCandidateIDs' SQL WHERE clause and ORDER BY both rely on) always
// agrees with chronological order. time.RFC3339Nano would NOT be safe to
// compare as plain strings: e.g. "...:00Z" (no fractional part) sorts before
// "...:00.5Z" lexically ('.' < 'Z'), even though .5s is later.
const sqliteTimeLayout = "2006-01-02T15:04:05.000000000Z07:00"

func nowSQLiteTime() string { return time.Now().UTC().Format(sqliteTimeLayout) }

func formatSQLiteTime(t time.Time) string { return t.UTC().Format(sqliteTimeLayout) }

func parseSQLiteTime(s string) (time.Time, error) {
	return time.Parse(sqliteTimeLayout, s)
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// SQLiteStore implements Store on top of the schema described above.
type SQLiteStore struct {
	DB *sql.DB
}

// NewSQLiteStore wraps an existing *sql.DB. It must point at a database
// where the migration in migrations_sqlite/00001_create_tus_uploads.sql has
// run (see MigrateSQLite) - typically via OpenSQLiteStore, not called
// directly.
func NewSQLiteStore(db *sql.DB) Store {
	return &SQLiteStore{DB: db}
}

var (
	_ Store                       = (*SQLiteStore)(nil)
	_ handler.DataStore           = (*SQLiteStore)(nil)
	_ handler.TerminaterDataStore = (*SQLiteStore)(nil)
)

// serializable asks the ncruces sqlite driver for an "immediate" transaction
// (see its package doc: "a serializable transaction is always immediate"),
// which acquires the write lock up front instead of lazily on the first
// write statement - the closest equivalent sqlite has to Postgres' explicit
// `SELECT ... FOR UPDATE` (used by ClaimFile there), and applied uniformly
// to every write transaction here since sqlite serializes writers at the
// whole-database level regardless.
func (s *SQLiteStore) beginWrite(ctx context.Context) (*sql.Tx, error) {
	return s.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
}

func (s *SQLiteStore) NewUpload(ctx context.Context, info handler.FileInfo) (handler.Upload, error) {
	if info.ID == "" {
		info.ID = uuid.NewString()
	}
	if info.MetaData == nil {
		info.MetaData = handler.MetaData{}
	}
	if info.PartialUploads == nil {
		info.PartialUploads = []string{}
	}

	userId := nullableString(info.MetaData[metaKeyUserId])
	workspaceId := nullableString(info.MetaData[metaKeyWorkspaceId])
	accessLevel := nullableString(info.MetaData[metaKeyAccessLevel])
	delete(info.MetaData, metaKeyUserId)
	delete(info.MetaData, metaKeyWorkspaceId)
	delete(info.MetaData, metaKeyAccessLevel)

	metaJSON, err := json.Marshal(info.MetaData)
	if err != nil {
		return nil, err
	}
	partialJSON, err := json.Marshal(info.PartialUploads)
	if err != nil {
		return nil, err
	}

	now := nowSQLiteTime()
	_, err = s.DB.ExecContext(ctx, `
		INSERT INTO tus_uploads (id, size, size_is_deferred, upload_offset, metadata, is_partial, is_final, partial_uploads, completed, created_at, updated_at, user_id, workspace_id, access_level)
		VALUES (?, ?, ?, 0, ?, ?, ?, ?, 0, ?, ?, ?, ?, ?)
	`, info.ID, info.Size, boolToInt(info.SizeIsDeferred), metaJSON, boolToInt(info.IsPartial), boolToInt(info.IsFinal), partialJSON, now, now, userId, workspaceId, accessLevel)
	if err != nil {
		return nil, err
	}

	info.Storage = map[string]string{"Type": "sqlite-chunks"}
	return &sqliteChunkUpload{store: s, info: info}, nil
}

func (s *SQLiteStore) GetUpload(ctx context.Context, id string) (handler.Upload, error) {
	info := handler.FileInfo{ID: id}

	var (
		metaJSON, partialJSON            []byte
		sizeIsDeferred, isPartial, isFinal int
	)

	row := s.DB.QueryRowContext(ctx, `
		SELECT size, size_is_deferred, upload_offset, metadata, is_partial, is_final, partial_uploads
		FROM tus_uploads WHERE id = ?
	`, id)
	if err := row.Scan(&info.Size, &sizeIsDeferred, &info.Offset, &metaJSON, &isPartial, &isFinal, &partialJSON); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, handler.ErrNotFound
		}
		return nil, err
	}
	info.SizeIsDeferred = sizeIsDeferred != 0
	info.IsPartial = isPartial != 0
	info.IsFinal = isFinal != 0

	info.MetaData = handler.MetaData{}
	if len(metaJSON) > 0 {
		if err := json.Unmarshal(metaJSON, &info.MetaData); err != nil {
			return nil, err
		}
	}
	if len(partialJSON) > 0 {
		if err := json.Unmarshal(partialJSON, &info.PartialUploads); err != nil {
			return nil, err
		}
	}

	info.Storage = map[string]string{"Type": "sqlite-chunks"}
	return &sqliteChunkUpload{store: s, info: info}, nil
}

func (s *SQLiteStore) AsTerminatableUpload(upload handler.Upload) handler.TerminatableUpload {
	return upload.(*sqliteChunkUpload)
}

// OpenRange opens id for reading, seeked to offset - see pgLoStore.OpenRange
// for the full contract Download.go relies on.
func (s *SQLiteStore) OpenRange(ctx context.Context, id string, offset int64) (io.ReadCloser, error) {
	var exists int
	if err := s.DB.QueryRowContext(ctx, `SELECT 1 FROM tus_uploads WHERE id = ?`, id).Scan(&exists); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, handler.ErrNotFound
		}
		return nil, err
	}

	pageOffset := (offset / chunkSize) * chunkSize
	skip := offset - pageOffset

	rows, err := s.DB.QueryContext(ctx, `
		SELECT data FROM tus_upload_chunks WHERE upload_id = ? AND chunk_offset >= ? ORDER BY chunk_offset ASC
	`, id, pageOffset)
	if err != nil {
		return nil, err
	}

	r := &sqliteChunkReader{rows: rows}
	if skip > 0 {
		if _, err := io.CopyN(io.Discard, r, skip); err != nil {
			r.Close()
			return nil, err
		}
	}
	return r, nil
}

// OwnerOf returns the user_id stored for id, and whether id exists at all -
// see pgLoStore.OwnerOf for the full contract.
func (s *SQLiteStore) OwnerOf(ctx context.Context, id string) (ownerId *string, found bool, err error) {
	err = s.DB.QueryRowContext(ctx, `SELECT user_id FROM tus_uploads WHERE id = ?`, id).Scan(&ownerId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return ownerId, true, nil
}

type sqliteChunkUpload struct {
	store *SQLiteStore
	info  handler.FileInfo
}

var _ handler.Upload = (*sqliteChunkUpload)(nil)

func (u *sqliteChunkUpload) GetInfo(ctx context.Context) (handler.FileInfo, error) {
	return u.info, nil
}

func (u *sqliteChunkUpload) FinishUpload(ctx context.Context) error {
	return nil
}

// WriteChunk appends src into whichever page(s) offset falls in, page by
// page: each iteration reads at most however many bytes are left in the
// current page (chunkSize - offset%chunkSize), so a single call to
// appendPage always either starts a brand new page (offset%chunkSize == 0)
// or appends to the tail of one that a previous call already started -
// never both within the same statement. This relies on the same invariant
// pgLoStore's lo.Seek(offset) does: tusd always calls WriteChunk at the
// upload's own current offset, so offset%chunkSize always equals exactly
// how many bytes that page already has.
func (u *sqliteChunkUpload) WriteChunk(ctx context.Context, offset int64, src io.Reader) (int64, error) {
	if offset != u.info.Offset {
		return 0, fmt.Errorf("fileupload: WriteChunk offset %d does not match upload's current offset %d", offset, u.info.Offset)
	}

	tx, err := u.store.beginWrite(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var written int64
	cur := offset
	buf := make([]byte, chunkSize)

	for {
		posInPage := cur % chunkSize
		want := chunkSize - posInPage

		n, readErr := io.ReadFull(src, buf[:want])
		if n > 0 {
			pageOffset := cur - posInPage
			if posInPage == 0 {
				if _, err := tx.ExecContext(ctx, `
					INSERT INTO tus_upload_chunks (upload_id, chunk_offset, data) VALUES (?, ?, ?)
				`, u.info.ID, pageOffset, buf[:n]); err != nil {
					return written, err
				}
			} else {
				if _, err := tx.ExecContext(ctx, `
					UPDATE tus_upload_chunks SET data = data || ? WHERE upload_id = ? AND chunk_offset = ?
				`, buf[:n], u.info.ID, pageOffset); err != nil {
					return written, err
				}
			}
			written += int64(n)
			cur += int64(n)
		}

		if readErr == io.EOF || readErr == io.ErrUnexpectedEOF {
			break // src exhausted - completely normal, a PATCH body is usually smaller than a full page
		}
		if readErr != nil {
			return written, readErr
		}
	}

	newOffset := offset + written
	completed := u.info.Size > 0 && newOffset >= u.info.Size
	now := nowSQLiteTime()

	res, err := tx.ExecContext(ctx, `
		UPDATE tus_uploads
		SET upload_offset = ?, updated_at = ?, completed = ?, completed_at = CASE WHEN ? THEN ? ELSE completed_at END
		WHERE id = ?
	`, newOffset, now, boolToInt(completed), boolToInt(completed), now, u.info.ID)
	if err != nil {
		return written, err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return written, handler.ErrNotFound
	}

	if err := tx.Commit(); err != nil {
		return written, err
	}

	u.info.Offset = newOffset
	return written, nil
}

func (u *sqliteChunkUpload) GetReader(ctx context.Context) (io.Reader, error) {
	rows, err := u.store.DB.QueryContext(ctx, `
		SELECT data FROM tus_upload_chunks WHERE upload_id = ? ORDER BY chunk_offset ASC
	`, u.info.ID)
	if err != nil {
		return nil, err
	}
	return &sqliteChunkReader{rows: rows}, nil
}

// Terminate removes both the chunks and the bookkeeping row. tus_upload_chunks'
// own ON DELETE CASCADE would do this automatically, but only if the
// connection has "PRAGMA foreign_keys = ON" (sqlite ignores FK actions
// otherwise, by default) - deleting explicitly here doesn't depend on that
// pragma being set correctly on whatever connection happens to run this.
func (u *sqliteChunkUpload) Terminate(ctx context.Context) error {
	tx, err := u.store.beginWrite(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, `DELETE FROM tus_upload_chunks WHERE upload_id = ?`, u.info.ID); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, `DELETE FROM tus_uploads WHERE id = ?`, u.info.ID); err != nil {
		return err
	}
	return tx.Commit()
}

// sqliteChunkReader streams consecutive tus_upload_chunks rows as one
// contiguous io.Reader, without pulling the whole upload into memory at
// once - used by both GetReader (all pages) and OpenRange (pages from
// wherever the requested offset starts).
type sqliteChunkReader struct {
	rows *sql.Rows
	cur  []byte
}

func (r *sqliteChunkReader) Read(p []byte) (int, error) {
	for len(r.cur) == 0 {
		if !r.rows.Next() {
			if err := r.rows.Err(); err != nil {
				return 0, err
			}
			return 0, io.EOF
		}
		if err := r.rows.Scan(&r.cur); err != nil {
			return 0, err
		}
	}
	n := copy(p, r.cur)
	r.cur = r.cur[n:]
	return n, nil
}

func (r *sqliteChunkReader) Close() error {
	return r.rows.Close()
}

// --- bookkeeping-only queries (ListFiles/GetFile/ClaimFile/ReleaseFile/
// UsedBytes/WorkspaceUsedBytes/orphanCandidateIDs) - the sqlite counterparts
// Queries.go/Claim.go/Quota.go/Reaper.go's public functions dispatch to. ---

func (s *SQLiteStore) ListFiles(ctx context.Context, limit, offset int) ([]FileRecord, error) {
	rows, err := s.DB.QueryContext(ctx, `
		SELECT id, size, upload_offset, completed, metadata, created_at, updated_at, completed_at, claimed_by, claimed_at, user_id, workspace_id, access_level
		FROM tus_uploads
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []FileRecord
	for rows.Next() {
		rec, err := scanFileRecordSQLite(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, rec)
	}
	return results, rows.Err()
}

func (s *SQLiteStore) GetFile(ctx context.Context, id string) (*FileRecord, error) {
	row := s.DB.QueryRowContext(ctx, `
		SELECT id, size, upload_offset, completed, metadata, created_at, updated_at, completed_at, claimed_by, claimed_at, user_id, workspace_id, access_level
		FROM tus_uploads WHERE id = ?
	`, id)

	rec, err := scanFileRecordSQLite(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &rec, nil
}

func scanFileRecordSQLite(row rowScanner) (FileRecord, error) {
	var (
		rec                           FileRecord
		metaJSON                      []byte
		completedInt                  int
		createdAtStr, updatedAtStr    string
		completedAtStr, claimedAtStr  sql.NullString
	)

	if err := row.Scan(&rec.ID, &rec.Size, &rec.Offset, &completedInt, &metaJSON, &createdAtStr, &updatedAtStr, &completedAtStr, &rec.ClaimedBy, &claimedAtStr, &rec.UserId, &rec.WorkspaceId, &rec.AccessLevel); err != nil {
		return FileRecord{}, err
	}

	rec.Completed = completedInt != 0
	rec.MetaData = map[string]string{}
	if len(metaJSON) > 0 {
		if err := json.Unmarshal(metaJSON, &rec.MetaData); err != nil {
			return FileRecord{}, err
		}
	}

	createdAt, err := parseSQLiteTime(createdAtStr)
	if err != nil {
		return FileRecord{}, fmt.Errorf("parsing created_at %q: %w", createdAtStr, err)
	}
	rec.CreatedAt = createdAt

	updatedAt, err := parseSQLiteTime(updatedAtStr)
	if err != nil {
		return FileRecord{}, fmt.Errorf("parsing updated_at %q: %w", updatedAtStr, err)
	}
	rec.UpdatedAt = updatedAt

	if completedAtStr.Valid {
		t, err := parseSQLiteTime(completedAtStr.String)
		if err != nil {
			return FileRecord{}, fmt.Errorf("parsing completed_at %q: %w", completedAtStr.String, err)
		}
		rec.CompletedAt = &t
	}
	if claimedAtStr.Valid {
		t, err := parseSQLiteTime(claimedAtStr.String)
		if err != nil {
			return FileRecord{}, fmt.Errorf("parsing claimed_at %q: %w", claimedAtStr.String, err)
		}
		rec.ClaimedAt = &t
	}

	return rec, nil
}

func (s *SQLiteStore) ClaimFile(ctx context.Context, id, claimedBy string) (*FileRecord, error) {
	tx, err := s.beginWrite(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var (
		completedInt       int
		existingClaimedBy sql.NullString
	)
	err = tx.QueryRowContext(ctx, `SELECT completed, claimed_by FROM tus_uploads WHERE id = ?`, id).Scan(&completedInt, &existingClaimedBy)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUploadNotFound
		}
		return nil, err
	}

	if completedInt == 0 {
		return nil, ErrUploadNotCompleted
	}

	// Claiming twice with the same owner is a no-op success, same as
	// pgLoStore's own ClaimFile (see its commented-out block) - claiming
	// with a *different* owner than what's recorded is intentionally not
	// rejected here either, matching that existing behavior exactly.
	if _, err := tx.ExecContext(ctx, `UPDATE tus_uploads SET claimed_by = ?, claimed_at = ? WHERE id = ?`, claimedBy, nowSQLiteTime(), id); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return s.GetFile(ctx, id)
}

func (s *SQLiteStore) ReleaseFile(ctx context.Context, id, claimedBy string) error {
	_, err := s.DB.ExecContext(ctx, `
		UPDATE tus_uploads SET claimed_by = NULL, claimed_at = NULL
		WHERE id = ? AND claimed_by = ?
	`, id, claimedBy)
	return err
}

func (s *SQLiteStore) UsedBytes(ctx context.Context, userId string) (int64, error) {
	var used int64
	err := s.DB.QueryRowContext(ctx, `SELECT COALESCE(SUM(size), 0) FROM tus_uploads WHERE user_id = ?`, userId).Scan(&used)
	return used, err
}

func (s *SQLiteStore) WorkspaceUsedBytes(ctx context.Context, workspaceId string) (int64, error) {
	var used int64
	err := s.DB.QueryRowContext(ctx, `SELECT COALESCE(SUM(size), 0) FROM tus_uploads WHERE workspace_id = ?`, workspaceId).Scan(&used)
	return used, err
}

// orphanCandidateIDs mirrors Reaper.go's Postgres query, in sqlite dialect -
// cutoffs are computed here in Go (rather than sqlite date functions) using
// the same fixed-width sqliteTimeLayout every write uses, so the plain "<"
// string comparison below agrees with chronological order.
func (s *SQLiteStore) orphanCandidateIDs(ctx context.Context, unclaimedTTL, incompleteTTL time.Duration) ([]string, error) {
	unclaimedCutoff := formatSQLiteTime(time.Now().Add(-unclaimedTTL))
	incompleteCutoff := formatSQLiteTime(time.Now().Add(-incompleteTTL))

	rows, err := s.DB.QueryContext(ctx, `
		SELECT id FROM tus_uploads
		WHERE (completed = 1 AND claimed_at IS NULL AND completed_at < ?)
		   OR (completed = 0 AND updated_at < ?)
	`, unclaimedCutoff, incompleteCutoff)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, rows.Err()
}
