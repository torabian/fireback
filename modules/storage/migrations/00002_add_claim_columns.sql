-- +goose Up
ALTER TABLE tus_uploads
    ADD COLUMN if not exists claimed_by text,
    ADD COLUMN if not exists claimed_at timestamptz;

-- Sweep query for unclaimed-but-completed uploads: filters on completed +
-- claimed_at, ordered/bounded by completed_at.
CREATE INDEX IF NOT EXISTS tus_uploads_unclaimed_idx
    ON tus_uploads (completed_at)
    WHERE completed = true AND claimed_at IS NULL;

-- Sweep query for abandoned in-progress uploads: filters on completed,
-- bounded by updated_at.
CREATE INDEX IF NOT EXISTS tus_uploads_incomplete_idx
    ON tus_uploads (updated_at)
    WHERE completed = false;

-- +goose Down
DROP INDEX IF EXISTS tus_uploads_incomplete_idx;
DROP INDEX IF EXISTS tus_uploads_unclaimed_idx;

ALTER TABLE tus_uploads
    DROP COLUMN claimed_by,
    DROP COLUMN claimed_at;
