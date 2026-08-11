package storage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// FileRecord is a plain, tusd-independent view of a stored upload, used to
// list and inspect files without pulling in the tus handler types.
type FileRecord struct {
	ID          string
	Size        int64
	Offset      int64
	Completed   bool
	MetaData    map[string]string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CompletedAt *time.Time
	// ClaimedBy identifies whatever owns this upload once some other feature
	// has attached it to itself (e.g. "score:<uniqueId>"), set via
	// ClaimFile. Uploads left unclaimed past a TTL are removed by
	// SweepOrphaned/StartReaper.
	ClaimedBy *string
	ClaimedAt *time.Time
	// UserId/WorkspaceId/AccessLevel are the AuthContext resolved by
	// StorageModuleConfig.Authenticate at upload creation time - nil for
	// all three if the upload was made anonymously (Authenticate was nil).
	UserId      *string
	WorkspaceId *string
	AccessLevel *string
}

// ListFiles returns the most recently created uploads, newest first.
// Dispatches on which Store implementation it's given - see pgLoStore
// (Postgres) and SQLiteStore.ListFiles (StoreSQLite.go).
func ListFiles(ctx context.Context, store Store, limit, offset int) ([]FileRecord, error) {
	switch s := store.(type) {
	case *pgLoStore:
		return listFilesPostgres(ctx, s.Pool, limit, offset)
	case *SQLiteStore:
		return s.ListFiles(ctx, limit, offset)
	case *MySQLStore:
		return s.ListFiles(ctx, limit, offset)
	default:
		return nil, fmt.Errorf("fileupload: unsupported store implementation %T", store)
	}
}

// GetFile returns a single upload's metadata by id, or nil if it doesn't exist.
func GetFile(ctx context.Context, store Store, id string) (*FileRecord, error) {
	switch s := store.(type) {
	case *pgLoStore:
		return getFilePostgres(ctx, s.Pool, id)
	case *SQLiteStore:
		return s.GetFile(ctx, id)
	case *MySQLStore:
		return s.GetFile(ctx, id)
	default:
		return nil, fmt.Errorf("fileupload: unsupported store implementation %T", store)
	}
}

func listFilesPostgres(ctx context.Context, pool *pgxpool.Pool, limit, offset int) ([]FileRecord, error) {
	rows, err := pool.Query(ctx, `
		SELECT id, size, upload_offset, completed, metadata, created_at, updated_at, completed_at, claimed_by, claimed_at, user_id, workspace_id, access_level
		FROM tus_uploads
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []FileRecord
	for rows.Next() {
		rec, err := scanFileRecord(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, rec)
	}
	return results, rows.Err()
}

func getFilePostgres(ctx context.Context, pool *pgxpool.Pool, id string) (*FileRecord, error) {
	row := pool.QueryRow(ctx, `
		SELECT id, size, upload_offset, completed, metadata, created_at, updated_at, completed_at, claimed_by, claimed_at, user_id, workspace_id, access_level
		FROM tus_uploads WHERE id = $1
	`, id)

	rec, err := scanFileRecord(row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &rec, nil
}

type rowScanner interface {
	Scan(dest ...any) error
}

func scanFileRecord(row rowScanner) (FileRecord, error) {
	var (
		rec         FileRecord
		metaJSON    []byte
		completedAt *time.Time
	)

	if err := row.Scan(&rec.ID, &rec.Size, &rec.Offset, &rec.Completed, &metaJSON, &rec.CreatedAt, &rec.UpdatedAt, &completedAt, &rec.ClaimedBy, &rec.ClaimedAt, &rec.UserId, &rec.WorkspaceId, &rec.AccessLevel); err != nil {
		return FileRecord{}, err
	}

	rec.CompletedAt = completedAt
	rec.MetaData = map[string]string{}
	if len(metaJSON) > 0 {
		if err := json.Unmarshal(metaJSON, &rec.MetaData); err != nil {
			return FileRecord{}, err
		}
	}
	return rec, nil
}
