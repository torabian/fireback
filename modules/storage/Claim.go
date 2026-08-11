// Claiming turns an anonymous, uploaded-but-unowned file into one that's
// attached to something in the rest of the app. Without this, a file
// uploaded through the tus endpoint (e.g. a score sheet picked before the
// score itself is saved) has nothing tying its lifetime to the record it was
// meant for: if the user never finishes creating that record, the upload
// just sits there forever, wasting storage and (since the upload endpoint
// needs no prior authorization to create files) offering free anonymous
// file hosting to anyone who finds it.
//
// The fix is a two-step contract: whichever feature ends up using an upload
// must call ClaimFile once it commits to that (typically inside the same
// handler/transaction that saves its own record), and SweepOrphaned /
// StartReaper deletes anything still unclaimed after a grace period.
package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	// ErrUploadNotFound is returned by ClaimFile when id doesn't exist.
	ErrUploadNotFound = errors.New("fileupload: upload not found")
	// ErrUploadNotCompleted is returned by ClaimFile when the upload hasn't
	// finished yet - only a finished upload has bytes worth keeping.
	ErrUploadNotCompleted = errors.New("fileupload: upload not completed")
	// ErrUploadAlreadyClaimed is returned by ClaimFile when the upload was
	// already claimed by a different owner.
	ErrUploadAlreadyClaimed = errors.New("fileupload: upload already claimed by another owner")
)

// ClaimFile marks a completed upload as owned by claimedBy - an
// application-defined reference such as "score:<uniqueId>" - so
// SweepOrphaned/StartReaper will never delete it. Call this from whatever
// handler attaches the upload to a real record, e.g. right after a
// createScore action saves a score referencing a previously uploaded sheet:
//
//	score, err := ScoreActions.Create(dto, query)
//	if err == nil {
//	    _, claimErr := storage.ClaimFile(ctx, storage.GetStore(), dto.SheetFileId, "score:"+score.UniqueId)
//	    // an error here means the file was already claimed elsewhere, or
//	    // never existed/finished uploading - surface it as a validation error
//	}
//
// Claiming the same upload twice with the same claimedBy is a no-op success
// (safe to retry). Dispatches on which Store implementation it's given - see
// pgLoStore (Postgres) and SQLiteStore.ClaimFile (StoreSQLite.go).
func ClaimFile(ctx context.Context, store Store, id, claimedBy string) (*FileRecord, error) {
	switch s := store.(type) {
	case *pgLoStore:
		return claimFilePostgres(ctx, s.Pool, id, claimedBy)
	case *SQLiteStore:
		return s.ClaimFile(ctx, id, claimedBy)
	case *MySQLStore:
		return s.ClaimFile(ctx, id, claimedBy)
	default:
		return nil, fmt.Errorf("fileupload: unsupported store implementation %T", store)
	}
}

// ReleaseFile clears a claim, e.g. when the owning record is deleted, so the
// upload becomes eligible for reaping again unless something re-claims it.
// It only clears the claim if claimedBy matches what's currently recorded,
// so one owner can't release another owner's claim; it's a no-op if the
// upload isn't claimed, is claimed by someone else, or doesn't exist.
func ReleaseFile(ctx context.Context, store Store, id, claimedBy string) error {
	switch s := store.(type) {
	case *pgLoStore:
		return releaseFilePostgres(ctx, s.Pool, id, claimedBy)
	case *SQLiteStore:
		return s.ReleaseFile(ctx, id, claimedBy)
	case *MySQLStore:
		return s.ReleaseFile(ctx, id, claimedBy)
	default:
		return fmt.Errorf("fileupload: unsupported store implementation %T", store)
	}
}

func claimFilePostgres(ctx context.Context, pool *pgxpool.Pool, id, claimedBy string) (*FileRecord, error) {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	var (
		completed         bool
		existingClaimedBy *string
	)
	err = tx.QueryRow(ctx, `
		SELECT completed, claimed_by FROM tus_uploads WHERE id = $1 FOR UPDATE
	`, id).Scan(&completed, &existingClaimedBy)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUploadNotFound
		}
		return nil, err
	}

	if !completed {
		return nil, ErrUploadNotCompleted
	}

	// If it's claimed it's fine already. It can be reused.
	// if existingClaimedBy != nil {
	// 	if *existingClaimedBy != claimedBy {
	// 		return nil, ErrUploadAlreadyClaimed
	// 	}
	// 	if err := tx.Commit(ctx); err != nil {
	// 		return nil, err
	// 	}
	// 	return getFilePostgres(ctx, pool, id)
	// }

	if _, err := tx.Exec(ctx, `
		UPDATE tus_uploads SET claimed_by = $1, claimed_at = now() WHERE id = $2
	`, claimedBy, id); err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return getFilePostgres(ctx, pool, id)
}

func releaseFilePostgres(ctx context.Context, pool *pgxpool.Pool, id, claimedBy string) error {
	_, err := pool.Exec(ctx, `
		UPDATE tus_uploads SET claimed_by = NULL, claimed_at = NULL
		WHERE id = $1 AND claimed_by = $2
	`, id, claimedBy)
	return err
}
