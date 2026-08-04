-- +goose Up
ALTER TABLE tus_uploads
    ADD COLUMN if not exists user_id text,
    ADD COLUMN if not exists workspace_id text,
    ADD COLUMN if not exists access_level text;

-- Used by UsedBytes (SUM(size) WHERE user_id = $1) to enforce per-user quota.
CREATE INDEX IF NOT EXISTS tus_uploads_user_id_idx
    ON tus_uploads (user_id)
    WHERE user_id IS NOT NULL;

-- +goose Down
DROP INDEX IF EXISTS tus_uploads_user_id_idx;

ALTER TABLE tus_uploads
    DROP COLUMN user_id,
    DROP COLUMN workspace_id,
    DROP COLUMN access_level;
