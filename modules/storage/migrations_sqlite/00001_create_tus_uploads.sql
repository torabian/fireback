-- +goose Up
-- SQLite equivalent of migrations/*.sql (Postgres), collapsed into one file since
-- there's no existing sqlite deployment to migrate incrementally. Postgres-only
-- types have no sqlite equivalent, so they're represented differently here:
--   oid              -> not needed at all; bytes live in tus_upload_chunks instead
--                        of a large object (see StoreSQLite.go)
--   jsonb            -> text (metadata/partial_uploads are marshaled JSON either way
--                        from Go's side; sqlite has no separate binary JSON type)
--   text[]           -> text (partial_uploads stored as a JSON array string)
--   timestamptz/now() -> text, RFC3339 timestamps written by the Go side
--   boolean          -> integer (0/1) - sqlite has no boolean storage class
CREATE TABLE IF NOT EXISTS tus_uploads (
    id text PRIMARY KEY,
    size integer NOT NULL DEFAULT 0,
    size_is_deferred integer NOT NULL DEFAULT 0,
    upload_offset integer NOT NULL DEFAULT 0,
    metadata text NOT NULL DEFAULT '{}',
    is_partial integer NOT NULL DEFAULT 0,
    is_final integer NOT NULL DEFAULT 0,
    partial_uploads text NOT NULL DEFAULT '[]',
    completed integer NOT NULL DEFAULT 0,
    created_at text NOT NULL,
    updated_at text NOT NULL,
    completed_at text,
    claimed_by text,
    claimed_at text,
    user_id text,
    workspace_id text,
    access_level text
);

-- Sweep query for unclaimed-but-completed uploads (mirrors
-- tus_uploads_unclaimed_idx in migrations/00002_add_claim_columns.sql).
CREATE INDEX IF NOT EXISTS tus_uploads_unclaimed_idx
    ON tus_uploads (completed_at)
    WHERE completed = 1 AND claimed_at IS NULL;

-- Sweep query for abandoned in-progress uploads (mirrors
-- tus_uploads_incomplete_idx).
CREATE INDEX IF NOT EXISTS tus_uploads_incomplete_idx
    ON tus_uploads (updated_at)
    WHERE completed = 0;

-- Used by UsedBytes (SUM(size) WHERE user_id = ?) to enforce per-user quota
-- (mirrors tus_uploads_user_id_idx in migrations/00003_add_owner_columns.sql).
CREATE INDEX IF NOT EXISTS tus_uploads_user_id_idx
    ON tus_uploads (user_id)
    WHERE user_id IS NOT NULL;

-- The actual file bytes, split into fixed-size pages aligned to absolute byte
-- offset (see chunkSize in StoreSQLite.go) - sqlite has no large-object API
-- (lo_*) to hold a whole upload's bytes as one growable value, and incremental
-- BLOB I/O (sqlite3_blob_open) requires knowing the final size upfront
-- (zeroblob), which breaks for tus's deferred-length uploads. A chunks table
-- sidesteps both problems and maps 1:1 onto tus's PATCH-based write model.
CREATE TABLE IF NOT EXISTS tus_upload_chunks (
    upload_id text NOT NULL REFERENCES tus_uploads(id) ON DELETE CASCADE,
    chunk_offset integer NOT NULL,
    data blob NOT NULL,
    PRIMARY KEY (upload_id, chunk_offset)
);

-- +goose Down
DROP TABLE IF EXISTS tus_upload_chunks;
DROP INDEX IF EXISTS tus_uploads_user_id_idx;
DROP INDEX IF EXISTS tus_uploads_incomplete_idx;
DROP INDEX IF EXISTS tus_uploads_unclaimed_idx;
DROP TABLE IF EXISTS tus_uploads;
