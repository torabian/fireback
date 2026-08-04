// package storage implements a tus (resumable upload) protocol data store
// that persists uploaded bytes as PostgreSQL large objects (lo_*), with the
// bookkeeping (offset, metadata, completion) kept in the tus_uploads table.
//
// It only depends on tusd's handler package (for the DataStore/Upload
// interfaces) and pgx - nothing here goes through Fireback or Emi.
package storage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tus/tusd/pkg/handler"
)

// Store implements handler.DataStore and handler.TerminaterDataStore on top
// of a single tus_uploads table plus PostgreSQL large objects.
type Store struct {
	Pool *pgxpool.Pool
}

// NewStore wraps an existing pgx pool. The pool must point at a database
// where the migration in migrations/00001_create_tus_uploads.sql has run.
func NewStore(pool *pgxpool.Pool) *Store {
	return &Store{Pool: pool}
}

var (
	_ handler.DataStore           = (*Store)(nil)
	_ handler.TerminaterDataStore = (*Store)(nil)
)

// Reserved Upload-Metadata keys Mount uses to smuggle the caller's resolved
// AuthContext into NewUpload: handler.go's serve rewrites the incoming
// Upload-Metadata header to set these (overwriting anything the client
// itself sent under the same names) before handing the request to tusd, so
// by the time NewUpload sees them here they're trustworthy - never taken
// from the client directly. NewUpload strips them back out of info.MetaData
// before persisting it, so they never leak into the client-facing metadata
// blob (e.g. download.go's JSON info response) - they live only in the
// dedicated user_id/workspace_id/access_level columns.
const (
	metaKeyUserId      = "_fileupload_user_id"
	metaKeyWorkspaceId = "_fileupload_workspace_id"
	metaKeyAccessLevel = "_fileupload_access_level"
)

// NewUpload creates the backing large object and its metadata row inside a
// single transaction, so the two never disagree about a upload's existence.
func (s *Store) NewUpload(ctx context.Context, info handler.FileInfo) (handler.Upload, error) {
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

	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	lo := tx.LargeObjects()
	oid, err := lo.Create(ctx, 0)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO tus_uploads (id, oid, size, size_is_deferred, metadata, is_partial, is_final, partial_uploads, user_id, workspace_id, access_level)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`, info.ID, oid, info.Size, info.SizeIsDeferred, metaJSON, info.IsPartial, info.IsFinal, info.PartialUploads, userId, workspaceId, accessLevel)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	info.Storage = map[string]string{
		"Type": "postgres-lo",
		"Oid":  fmt.Sprintf("%d", oid),
	}

	return &pgLoUpload{store: s, info: info, oid: oid}, nil
}

// GetUpload loads the upload's bookkeeping row. The large object itself is
// only opened lazily by WriteChunk/GetReader, since it must live inside its
// own short transaction.
func (s *Store) GetUpload(ctx context.Context, id string) (handler.Upload, error) {
	var (
		oid      uint32
		metaJSON []byte
	)

	info := handler.FileInfo{ID: id}

	err := s.Pool.QueryRow(ctx, `
		SELECT oid, size, size_is_deferred, upload_offset, metadata, is_partial, is_final, partial_uploads
		FROM tus_uploads WHERE id = $1
	`, id).Scan(&oid, &info.Size, &info.SizeIsDeferred, &info.Offset, &metaJSON, &info.IsPartial, &info.IsFinal, &info.PartialUploads)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, handler.ErrNotFound
		}
		return nil, err
	}

	info.MetaData = handler.MetaData{}
	if len(metaJSON) > 0 {
		if err := json.Unmarshal(metaJSON, &info.MetaData); err != nil {
			return nil, err
		}
	}
	info.Storage = map[string]string{
		"Type": "postgres-lo",
		"Oid":  fmt.Sprintf("%d", oid),
	}

	return &pgLoUpload{store: s, info: info, oid: oid}, nil
}

// AsTerminatableUpload lets the handler package issue DELETE requests.
func (s *Store) AsTerminatableUpload(upload handler.Upload) handler.TerminatableUpload {
	return upload.(*pgLoUpload)
}

type pgLoUpload struct {
	store *Store
	info  handler.FileInfo
	oid   uint32
}

var _ handler.Upload = (*pgLoUpload)(nil)

func (u *pgLoUpload) GetInfo(ctx context.Context) (handler.FileInfo, error) {
	return u.info, nil
}

// WriteChunk opens the large object for the duration of a single PATCH
// request. Large object file descriptors only live for the transaction that
// opened them, so open/seek/write/close all happen inside one tx here.
func (u *pgLoUpload) WriteChunk(ctx context.Context, offset int64, src io.Reader) (int64, error) {
	tx, err := u.store.Pool.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	los := tx.LargeObjects()
	lo, err := los.Open(ctx, u.oid, pgx.LargeObjectModeWrite)
	if err != nil {
		return 0, err
	}
	if _, err := lo.Seek(offset, io.SeekStart); err != nil {
		return 0, err
	}

	n, copyErr := io.Copy(lo, src)
	if copyErr != nil {
		return n, copyErr
	}
	if err := lo.Close(); err != nil {
		return n, err
	}

	newOffset := offset + n
	completed := u.info.Size > 0 && newOffset >= u.info.Size

	tag, err := tx.Exec(ctx, `
		UPDATE tus_uploads
		SET upload_offset = $1,
		    updated_at = now(),
		    completed = $2,
		    completed_at = CASE WHEN $2 THEN now() ELSE completed_at END
		WHERE id = $3
	`, newOffset, completed, u.info.ID)
	if err != nil {
		return n, err
	}
	if tag.RowsAffected() == 0 {
		return n, handler.ErrNotFound
	}

	if err := tx.Commit(ctx); err != nil {
		return n, err
	}

	u.info.Offset = newOffset
	return n, nil
}

// GetReader opens the large object for reading and hands back a reader that
// keeps its own transaction open until Close is called - large objects
// cannot be read outside the transaction that opened them.
func (u *pgLoUpload) GetReader(ctx context.Context) (io.Reader, error) {
	tx, err := u.store.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	los := tx.LargeObjects()
	lo, err := los.Open(ctx, u.oid, pgx.LargeObjectModeRead)
	if err != nil {
		_ = tx.Rollback(ctx)
		return nil, err
	}

	return &loReader{ctx: ctx, tx: tx, lo: lo}, nil
}

// OpenRange opens the large object backing id for reading, seeked to offset,
// so a caller can stream out a byte range without pulling in the tusd Upload
// type. Like GetReader, the returned reader keeps its own transaction open
// until Close is called.
func (s *Store) OpenRange(ctx context.Context, id string, offset int64) (io.ReadCloser, error) {
	var oid uint32
	if err := s.Pool.QueryRow(ctx, `SELECT oid FROM tus_uploads WHERE id = $1`, id).Scan(&oid); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, handler.ErrNotFound
		}
		return nil, err
	}

	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	los := tx.LargeObjects()
	lo, err := los.Open(ctx, oid, pgx.LargeObjectModeRead)
	if err != nil {
		_ = tx.Rollback(ctx)
		return nil, err
	}
	if offset > 0 {
		if _, err := lo.Seek(offset, io.SeekStart); err != nil {
			_ = tx.Rollback(ctx)
			return nil, err
		}
	}

	return &loReader{ctx: ctx, tx: tx, lo: lo}, nil
}

// OwnerOf returns the user_id stored for id, and whether id exists at all.
// A nil ownerId with found=true means the upload has no recorded owner
// (created while Authenticate was nil, or before this mechanism existed) -
// callers should treat that as open to anyone, the same as when
// Authenticate itself is nil.
func (s *Store) OwnerOf(ctx context.Context, id string) (ownerId *string, found bool, err error) {
	err = s.Pool.QueryRow(ctx, `SELECT user_id FROM tus_uploads WHERE id = $1`, id).Scan(&ownerId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return ownerId, true, nil
}

// nullableString turns an empty string into a SQL NULL, so an absent
// user_id/workspace_id/access_level is stored as NULL rather than "".
func nullableString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func (u *pgLoUpload) FinishUpload(ctx context.Context) error {
	return nil
}

// Terminate removes both the large object and its metadata row.
func (u *pgLoUpload) Terminate(ctx context.Context) error {
	tx, err := u.store.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	los := tx.LargeObjects()
	if err := los.Unlink(ctx, u.oid); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx, `DELETE FROM tus_uploads WHERE id = $1`, u.info.ID); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

type loReader struct {
	ctx context.Context
	tx  pgx.Tx
	lo  *pgx.LargeObject
}

func (r *loReader) Read(p []byte) (int, error) {
	return r.lo.Read(p)
}

func (r *loReader) Close() error {
	_ = r.lo.Close()
	return r.tx.Commit(r.ctx)
}
