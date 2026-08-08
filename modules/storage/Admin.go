// admin.go collects direct, non-HTTP operations against the Store - usage
// accounting, deleting an upload, and pushing a local file in - so they can
// be driven from a CLI (see cmd/fileupload-cli) or any other Go caller
// without going through the tus HTTP handler at all.
package storage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"maps"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tus/tusd/pkg/handler"
)

// WorkspaceUsedBytes mirrors UsedBytes, summed by workspace_id instead of
// user_id.
func WorkspaceUsedBytes(ctx context.Context, pool *pgxpool.Pool, workspaceId string) (int64, error) {
	var used int64
	err := pool.QueryRow(ctx, `
		SELECT COALESCE(SUM(size), 0) FROM tus_uploads WHERE workspace_id = $1
	`, workspaceId).Scan(&used)
	return used, err
}

// DeleteFile removes upload id - both its tus_uploads row and backing large
// object, the same way DELETE /files/:id does - but callable directly
// without an HTTP round trip. Returns ErrUploadNotFound if id doesn't exist.
func DeleteFile(ctx context.Context, store *Store, id string) error {
	upload, err := store.GetUpload(ctx, id)
	if err != nil {
		if errors.Is(err, handler.ErrNotFound) {
			return ErrUploadNotFound
		}
		return err
	}
	return store.AsTerminatableUpload(upload).Terminate(ctx)
}

// UploadFileOptions configures UploadFile.
type UploadFileOptions struct {
	// UserId/WorkspaceId/AccessLevel are recorded as the upload's owner,
	// same as if an authenticated HTTP client had created it - see
	// AuthContext. Leave empty for an anonymous/unowned upload.
	UserId      string
	WorkspaceId string
	AccessLevel string

	// MetaData is stored verbatim as the upload's tus metadata (the same
	// blob GetFile and the JSON download endpoint expose). "filename" is
	// set to localPath's base name if not already present.
	MetaData map[string]string

	// ClaimedBy, if set, immediately claims the upload (see ClaimFile) so
	// the background reaper never deletes it for sitting unclaimed.
	ClaimedBy string
}

// UploadFile reads localPath off disk and stores it through store, the same
// end result as a client's POST+PATCH over the tus HTTP endpoint, but in a
// single call with no HTTP involved and no chunking - meant for CLI/admin
// use. It does not consult StorageModuleConfig.Quota; that check only
// runs on the HTTP creation path.
func UploadFile(ctx context.Context, pool *pgxpool.Pool, store *Store, localPath string, opts UploadFileOptions) (*FileRecord, error) {
	f, err := os.Open(localPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return nil, fmt.Errorf("fileupload: %s is a directory", localPath)
	}

	meta := handler.MetaData{}
	maps.Copy(meta, opts.MetaData)
	if _, ok := meta["filename"]; !ok {
		meta["filename"] = filepath.Base(localPath)
	}
	if opts.UserId != "" {
		meta[metaKeyUserId] = opts.UserId
	}
	if opts.WorkspaceId != "" {
		meta[metaKeyWorkspaceId] = opts.WorkspaceId
	}
	if opts.AccessLevel != "" {
		meta[metaKeyAccessLevel] = opts.AccessLevel
	}

	upload, err := store.NewUpload(ctx, handler.FileInfo{
		Size:     stat.Size(),
		MetaData: meta,
	})
	if err != nil {
		return nil, err
	}

	info, err := upload.GetInfo(ctx)
	if err != nil {
		return nil, err
	}

	if stat.Size() > 0 {
		written, err := upload.WriteChunk(ctx, 0, io.LimitReader(f, stat.Size()))
		if err != nil {
			return nil, err
		}
		if written != stat.Size() {
			return nil, fmt.Errorf("fileupload: wrote %d of %d bytes for %s", written, stat.Size(), localPath)
		}
	}

	if opts.ClaimedBy != "" {
		if _, err := ClaimFile(ctx, pool, info.ID, opts.ClaimedBy); err != nil {
			return nil, err
		}
	}

	return GetFile(ctx, pool, info.ID)
}
