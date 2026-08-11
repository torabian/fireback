package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// SweepOrphaned deletes uploads nothing has claimed: either
//   - completed but left unclaimed for longer than unclaimedTTL, or
//   - never completed (abandoned mid-transfer) for longer than incompleteTTL.
//
// It returns the ids it deleted. Each deletion goes through the same
// store.AsTerminatableUpload(...).Terminate path DELETE /files/:id does, so
// both the tus_uploads row and its backing bytes (a Postgres large object,
// or sqlite's tus_upload_chunks rows) are removed - nothing is left
// dangling, regardless of which Store implementation this is given.
func SweepOrphaned(ctx context.Context, store Store, unclaimedTTL, incompleteTTL time.Duration) ([]string, error) {
	var (
		ids []string
		err error
	)

	switch s := store.(type) {
	case *pgLoStore:
		ids, err = orphanCandidateIDsPostgres(ctx, s.Pool, unclaimedTTL, incompleteTTL)
	case *SQLiteStore:
		ids, err = s.orphanCandidateIDs(ctx, unclaimedTTL, incompleteTTL)
	case *MySQLStore:
		ids, err = s.orphanCandidateIDs(ctx, unclaimedTTL, incompleteTTL)
	default:
		return nil, fmt.Errorf("fileupload: unsupported store implementation %T", store)
	}
	if err != nil {
		return nil, err
	}

	var deleted []string
	for _, id := range ids {
		upload, err := store.GetUpload(ctx, id)
		if err != nil {
			// Gone already, e.g. a concurrent sweep or a claim that raced
			// this one - nothing left to clean up either way.
			continue
		}
		if err := store.AsTerminatableUpload(upload).Terminate(ctx); err != nil {
			return deleted, err
		}
		deleted = append(deleted, id)
	}

	return deleted, nil
}

func orphanCandidateIDsPostgres(ctx context.Context, pool *pgxpool.Pool, unclaimedTTL, incompleteTTL time.Duration) ([]string, error) {
	rows, err := pool.Query(ctx, `
		SELECT id FROM tus_uploads
		WHERE (completed = true AND claimed_at IS NULL AND completed_at < now() - ($1::float8 * interval '1 second'))
		   OR (completed = false AND updated_at < now() - ($2::float8 * interval '1 second'))
	`, unclaimedTTL.Seconds(), incompleteTTL.Seconds())
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

// StartReaper runs SweepOrphaned once immediately and then on a fixed
// interval until ctx is canceled, logging what it deletes. Wire it up once
// per process next to Mount, e.g.:
//
//	ctx, cancel := context.WithCancel(context.Background())
//	storage.StartReaper(ctx, store, time.Hour, 24*time.Hour, 24*time.Hour)
//	// cancel() on shutdown to stop the background goroutine
func StartReaper(ctx context.Context, store Store, interval, unclaimedTTL, incompleteTTL time.Duration) {
	sweep := func() {
		deleted, err := SweepOrphaned(ctx, store, unclaimedTTL, incompleteTTL)
		if err != nil {
			log.Printf("fileupload: sweep orphaned uploads: %v", err)
			return
		}
		if len(deleted) > 0 {
			log.Printf("fileupload: reaped %d orphaned upload(s): %v", len(deleted), deleted)
		}
	}

	go func() {
		sweep()

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				sweep()
			}
		}
	}()
}
