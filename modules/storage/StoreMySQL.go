// StoreMySQL.go implements Store (see Store.go) on top of MySQL/MariaDB, the
// same chunked-pages design StoreSQLite.go uses (see its own package doc for
// why a chunks table exists at all instead of a single blob) - chunkSize is
// shared verbatim from there. Two real differences from the sqlite backend,
// both because MySQL/InnoDB is a proper multi-writer database unlike
// sqlite:
//   - blob concatenation is CONCAT(data, ?), not sqlite's `||` (MySQL's `||`
//     means logical OR unless the non-default PIPES_AS_CONCAT sql_mode is
//     set - not something this module should depend on the server having).
//   - ClaimFile uses a real `SELECT ... FOR UPDATE` row lock (InnoDB
//     supports it, same as Postgres) instead of sqlite's whole-database
//     "immediate transaction" workaround.
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

// mysqlTimeLayout mirrors sqliteTimeLayout's own doc comment: fixed-width
// (always 9 fractional digits) so plain lexical string comparison agrees
// with chronological order - used instead of MySQL's own DATETIME/TIMESTAMP
// types so the reaper's TTL comparisons never depend on the connection's
// session timezone.
const mysqlTimeLayout = "2006-01-02T15:04:05.000000000Z07:00"

func nowMySQLTime() string { return time.Now().UTC().Format(mysqlTimeLayout) }

func formatMySQLTime(t time.Time) string { return t.UTC().Format(mysqlTimeLayout) }

func parseMySQLTime(s string) (time.Time, error) {
	return time.Parse(mysqlTimeLayout, s)
}

// MySQLStore implements Store on top of the schema described in
// migrations_mysql/00001_create_tus_uploads.sql.
type MySQLStore struct {
	DB *sql.DB
}

// NewMySQLStore wraps an existing *sql.DB. It must point at a database
// where the migration in migrations_mysql/00001_create_tus_uploads.sql has
// run (see MigrateMySQL) - typically via OpenMySQLStore, not called
// directly.
func NewMySQLStore(db *sql.DB) Store {
	return &MySQLStore{DB: db}
}

var (
	_ Store                       = (*MySQLStore)(nil)
	_ handler.DataStore           = (*MySQLStore)(nil)
	_ handler.TerminaterDataStore = (*MySQLStore)(nil)
)

func (s *MySQLStore) NewUpload(ctx context.Context, info handler.FileInfo) (handler.Upload, error) {
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

	now := nowMySQLTime()
	_, err = s.DB.ExecContext(ctx, `
		INSERT INTO tus_uploads (id, size, size_is_deferred, upload_offset, metadata, is_partial, is_final, partial_uploads, completed, created_at, updated_at, user_id, workspace_id, access_level)
		VALUES (?, ?, ?, 0, ?, ?, ?, ?, 0, ?, ?, ?, ?, ?)
	`, info.ID, info.Size, boolToInt(info.SizeIsDeferred), metaJSON, boolToInt(info.IsPartial), boolToInt(info.IsFinal), partialJSON, now, now, userId, workspaceId, accessLevel)
	if err != nil {
		return nil, err
	}

	info.Storage = map[string]string{"Type": "mysql-chunks"}
	return &mysqlChunkUpload{store: s, info: info}, nil
}

func (s *MySQLStore) GetUpload(ctx context.Context, id string) (handler.Upload, error) {
	info := handler.FileInfo{ID: id}

	var (
		metaJSON, partialJSON              []byte
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

	info.Storage = map[string]string{"Type": "mysql-chunks"}
	return &mysqlChunkUpload{store: s, info: info}, nil
}

func (s *MySQLStore) AsTerminatableUpload(upload handler.Upload) handler.TerminatableUpload {
	return upload.(*mysqlChunkUpload)
}

// OpenRange opens id for reading, seeked to offset - see pgLoStore.OpenRange
// for the full contract Download.go relies on.
func (s *MySQLStore) OpenRange(ctx context.Context, id string, offset int64) (io.ReadCloser, error) {
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
func (s *MySQLStore) OwnerOf(ctx context.Context, id string) (ownerId *string, found bool, err error) {
	err = s.DB.QueryRowContext(ctx, `SELECT user_id FROM tus_uploads WHERE id = ?`, id).Scan(&ownerId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return ownerId, true, nil
}

type mysqlChunkUpload struct {
	store *MySQLStore
	info  handler.FileInfo
}

var _ handler.Upload = (*mysqlChunkUpload)(nil)

func (u *mysqlChunkUpload) GetInfo(ctx context.Context) (handler.FileInfo, error) {
	return u.info, nil
}

func (u *mysqlChunkUpload) FinishUpload(ctx context.Context) error {
	return nil
}

// WriteChunk mirrors sqliteChunkUpload.WriteChunk's page-by-page loop
// exactly (see its own doc comment for the invariant this relies on) - the
// only difference is the append statement itself: CONCAT(data, ?) instead
// of sqlite's `data || ?`.
func (u *mysqlChunkUpload) WriteChunk(ctx context.Context, offset int64, src io.Reader) (int64, error) {
	if offset != u.info.Offset {
		return 0, fmt.Errorf("fileupload: WriteChunk offset %d does not match upload's current offset %d", offset, u.info.Offset)
	}

	tx, err := u.store.DB.BeginTx(ctx, nil)
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
					UPDATE tus_upload_chunks SET data = CONCAT(data, ?) WHERE upload_id = ? AND chunk_offset = ?
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
	now := nowMySQLTime()

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

func (u *mysqlChunkUpload) GetReader(ctx context.Context) (io.Reader, error) {
	rows, err := u.store.DB.QueryContext(ctx, `
		SELECT data FROM tus_upload_chunks WHERE upload_id = ? ORDER BY chunk_offset ASC
	`, u.info.ID)
	if err != nil {
		return nil, err
	}
	return &sqliteChunkReader{rows: rows}, nil
}

// Terminate removes both the chunks and the bookkeeping row. InnoDB's own
// ON DELETE CASCADE (tus_upload_chunks' foreign key) would do this
// automatically and is always enforced (unlike sqlite, which needs an
// explicit pragma) - deleting explicitly here regardless is still cheap
// belt-and-suspenders, and keeps this method's shape identical to the other
// two Store implementations'.
func (u *mysqlChunkUpload) Terminate(ctx context.Context) error {
	tx, err := u.store.DB.BeginTx(ctx, nil)
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

// --- bookkeeping-only queries (ListFiles/GetFile/ClaimFile/ReleaseFile/
// UsedBytes/WorkspaceUsedBytes/orphanCandidateIDs) - the mysql counterparts
// Queries.go/Claim.go/Quota.go/Reaper.go's public functions dispatch to. ---

func (s *MySQLStore) ListFiles(ctx context.Context, limit, offset int) ([]FileRecord, error) {
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
		rec, err := scanFileRecordMySQL(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, rec)
	}
	return results, rows.Err()
}

func (s *MySQLStore) GetFile(ctx context.Context, id string) (*FileRecord, error) {
	row := s.DB.QueryRowContext(ctx, `
		SELECT id, size, upload_offset, completed, metadata, created_at, updated_at, completed_at, claimed_by, claimed_at, user_id, workspace_id, access_level
		FROM tus_uploads WHERE id = ?
	`, id)

	rec, err := scanFileRecordMySQL(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &rec, nil
}

func scanFileRecordMySQL(row rowScanner) (FileRecord, error) {
	var (
		rec                          FileRecord
		metaJSON                     []byte
		completedInt                 int
		createdAtStr, updatedAtStr   string
		completedAtStr, claimedAtStr sql.NullString
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

	createdAt, err := parseMySQLTime(createdAtStr)
	if err != nil {
		return FileRecord{}, fmt.Errorf("parsing created_at %q: %w", createdAtStr, err)
	}
	rec.CreatedAt = createdAt

	updatedAt, err := parseMySQLTime(updatedAtStr)
	if err != nil {
		return FileRecord{}, fmt.Errorf("parsing updated_at %q: %w", updatedAtStr, err)
	}
	rec.UpdatedAt = updatedAt

	if completedAtStr.Valid {
		t, err := parseMySQLTime(completedAtStr.String)
		if err != nil {
			return FileRecord{}, fmt.Errorf("parsing completed_at %q: %w", completedAtStr.String, err)
		}
		rec.CompletedAt = &t
	}
	if claimedAtStr.Valid {
		t, err := parseMySQLTime(claimedAtStr.String)
		if err != nil {
			return FileRecord{}, fmt.Errorf("parsing claimed_at %q: %w", claimedAtStr.String, err)
		}
		rec.ClaimedAt = &t
	}

	return rec, nil
}

// ClaimFile uses a real `SELECT ... FOR UPDATE` row lock (InnoDB, like
// Postgres) rather than sqlite's whole-database "immediate transaction"
// workaround - MySQL doesn't need it, it has real per-row locking.
func (s *MySQLStore) ClaimFile(ctx context.Context, id, claimedBy string) (*FileRecord, error) {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var (
		completedInt      int
		existingClaimedBy sql.NullString
	)
	err = tx.QueryRowContext(ctx, `SELECT completed, claimed_by FROM tus_uploads WHERE id = ? FOR UPDATE`, id).Scan(&completedInt, &existingClaimedBy)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUploadNotFound
		}
		return nil, err
	}

	if completedInt == 0 {
		return nil, ErrUploadNotCompleted
	}

	// Claiming twice with the same owner is a no-op success, same as the
	// other two Store implementations - claiming with a *different* owner
	// than what's recorded is intentionally not rejected here either,
	// matching that existing behavior exactly.
	if _, err := tx.ExecContext(ctx, `UPDATE tus_uploads SET claimed_by = ?, claimed_at = ? WHERE id = ?`, claimedBy, nowMySQLTime(), id); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return s.GetFile(ctx, id)
}

func (s *MySQLStore) ReleaseFile(ctx context.Context, id, claimedBy string) error {
	_, err := s.DB.ExecContext(ctx, `
		UPDATE tus_uploads SET claimed_by = NULL, claimed_at = NULL
		WHERE id = ? AND claimed_by = ?
	`, id, claimedBy)
	return err
}

func (s *MySQLStore) UsedBytes(ctx context.Context, userId string) (int64, error) {
	var used int64
	err := s.DB.QueryRowContext(ctx, `SELECT COALESCE(SUM(size), 0) FROM tus_uploads WHERE user_id = ?`, userId).Scan(&used)
	return used, err
}

func (s *MySQLStore) WorkspaceUsedBytes(ctx context.Context, workspaceId string) (int64, error) {
	var used int64
	err := s.DB.QueryRowContext(ctx, `SELECT COALESCE(SUM(size), 0) FROM tus_uploads WHERE workspace_id = ?`, workspaceId).Scan(&used)
	return used, err
}

// orphanCandidateIDs mirrors SQLiteStore.orphanCandidateIDs - cutoffs are
// computed here in Go using the same fixed-width mysqlTimeLayout every
// write uses, so the plain "<" string comparison below agrees with
// chronological order.
func (s *MySQLStore) orphanCandidateIDs(ctx context.Context, unclaimedTTL, incompleteTTL time.Duration) ([]string, error) {
	unclaimedCutoff := formatMySQLTime(time.Now().Add(-unclaimedTTL))
	incompleteCutoff := formatMySQLTime(time.Now().Add(-incompleteTTL))

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
