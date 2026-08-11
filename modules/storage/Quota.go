package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DefaultQuotaBytes is how much storage a single user gets when
// StorageModuleConfig.Quota is nil: 2 GiB.
const DefaultQuotaBytes int64 = 2 << 30

// ErrQuotaExceeded is the error enforceQuotaCallback reports (as a tusd
// handler.HTTPError, HTTP 403) when a new upload would push its owner over
// their StorageModuleConfig.Quota.
var ErrQuotaExceeded = errors.New("fileupload: storage quota exceeded")

// UsedBytes returns the combined declared size of every upload currently
// stored for userId - completed or still in progress, since an in-progress
// upload has already reserved its declared Size worth of large-object
// space. SweepOrphaned/StartReaper is what reclaims that reservation if the
// upload is abandoned rather than finished. Dispatches on which Store
// implementation it's given - see pgLoStore (Postgres) and
// SQLiteStore.UsedBytes (StoreSQLite.go).
func UsedBytes(ctx context.Context, store Store, userId string) (int64, error) {
	switch s := store.(type) {
	case *pgLoStore:
		return usedBytesPostgres(ctx, s.Pool, userId)
	case *SQLiteStore:
		return s.UsedBytes(ctx, userId)
	case *MySQLStore:
		return s.UsedBytes(ctx, userId)
	default:
		return 0, fmt.Errorf("fileupload: unsupported store implementation %T", store)
	}
}

func usedBytesPostgres(ctx context.Context, pool *pgxpool.Pool, userId string) (int64, error) {
	var used int64
	err := pool.QueryRow(ctx, `
		SELECT COALESCE(SUM(size), 0) FROM tus_uploads WHERE user_id = $1
	`, userId).Scan(&used)
	return used, err
}

// quotaFor returns cfg.Quota(auth), or DefaultQuotaBytes if Quota is nil.
func (c *StorageModuleConfig) quotaFor(auth AuthContext) int64 {
	if c.Quota == nil {
		return DefaultQuotaBytes
	}
	return c.Quota(auth)
}
