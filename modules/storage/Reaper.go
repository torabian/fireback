package storage

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// SweepOrphaned deletes uploads nothing has claimed: either
//   - completed but left unclaimed for longer than unclaimedTTL, or
//   - never completed (abandoned mid-transfer) for longer than incompleteTTL.
//
// It returns the ids it deleted. Each deletion goes through the same
// Store.Terminate path DELETE /files/:id does, so both the tus_uploads row
// and its backing large object are removed - nothing is left dangling.
func SweepOrphaned(ctx context.Context, pool *pgxpool.Pool, store *Store, unclaimedTTL, incompleteTTL time.Duration) ([]string, error) {
	rows, err := pool.Query(ctx, `
		SELECT id FROM tus_uploads
		WHERE (completed = true AND claimed_at IS NULL AND completed_at < now() - ($1::float8 * interval '1 second'))
		   OR (completed = false AND updated_at < now() - ($2::float8 * interval '1 second'))
	`, unclaimedTTL.Seconds(), incompleteTTL.Seconds())
	if err != nil {
		return nil, err
	}

	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			rows.Close()
			return nil, err
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	rows.Close()

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

// StartReaper runs SweepOrphaned once immediately and then on a fixed
// interval until ctx is canceled, logging what it deletes. Wire it up once
// per process next to Mount, e.g.:
//
//	ctx, cancel := context.WithCancel(context.Background())
//	storage.StartReaper(ctx, pool, store, time.Hour, 24*time.Hour, 24*time.Hour)
//	// cancel() on shutdown to stop the background goroutine
func StartReaper(ctx context.Context, pool *pgxpool.Pool, store *Store, interval, unclaimedTTL, incompleteTTL time.Duration) {
	sweep := func() {
		deleted, err := SweepOrphaned(ctx, pool, store, unclaimedTTL, incompleteTTL)
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
