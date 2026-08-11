-- +goose Up
-- MySQL/MariaDB equivalent of migrations/*.sql (Postgres) and
-- migrations_sqlite/00001_create_tus_uploads.sql - collapsed into one file
-- like the sqlite migration, since there's no existing MySQL deployment to
-- migrate incrementally. Differences from Postgres' types:
--   oid              -> not needed at all; bytes live in tus_upload_chunks
--                        instead of a large object (see StoreMySQL.go)
--   jsonb            -> longtext - metadata/partial_uploads are marshaled
--                        JSON either way from Go's side; MySQL's native JSON
--                        column type would just add validation overhead for
--                        no benefit here
--   text[]           -> longtext (partial_uploads stored as a JSON array
--                        string, same as the sqlite schema)
--   timestamptz/now() -> varchar(40), RFC3339-ish fixed-width text written
--                        by the Go side (see chunkTimeLayout in
--                        StoreMySQL.go) - kept as text rather than
--                        DATETIME(6) so the reaper's TTL comparisons don't
--                        depend on the connection's session timezone at all
--   boolean          -> tinyint(1) - same convention as sqlite's schema
--
-- Unlike sqlite, MySQL/InnoDB enforces foreign keys (including ON DELETE
-- CASCADE) by default - no equivalent of sqlite's "PRAGMA foreign_keys=ON"
-- caveat here.
--
-- MySQL has no partial/filtered index support (no WHERE clause on CREATE
-- INDEX), unlike Postgres/sqlite - the two sweep indexes below lead with
-- `completed` instead, to get comparable selectivity for SweepOrphaned's
-- two queries.
CREATE TABLE IF NOT EXISTS tus_uploads (
    id varchar(191) NOT NULL PRIMARY KEY,
    size bigint NOT NULL DEFAULT 0,
    size_is_deferred tinyint(1) NOT NULL DEFAULT 0,
    upload_offset bigint NOT NULL DEFAULT 0,
    metadata longtext NOT NULL,
    is_partial tinyint(1) NOT NULL DEFAULT 0,
    is_final tinyint(1) NOT NULL DEFAULT 0,
    partial_uploads longtext NOT NULL,
    completed tinyint(1) NOT NULL DEFAULT 0,
    created_at varchar(40) NOT NULL,
    updated_at varchar(40) NOT NULL,
    completed_at varchar(40) NULL,
    claimed_by varchar(191) NULL,
    claimed_at varchar(40) NULL,
    user_id varchar(191) NULL,
    workspace_id varchar(191) NULL,
    access_level varchar(191) NULL
) ENGINE=InnoDB;

CREATE INDEX tus_uploads_unclaimed_idx ON tus_uploads (completed, claimed_at, completed_at);
CREATE INDEX tus_uploads_incomplete_idx ON tus_uploads (completed, updated_at);
CREATE INDEX tus_uploads_user_id_idx ON tus_uploads (user_id);
CREATE INDEX tus_uploads_workspace_id_idx ON tus_uploads (workspace_id);

-- The actual file bytes, split into fixed-size pages aligned to absolute
-- byte offset (see chunkSize in StoreSQLite.go, reused as-is by
-- StoreMySQL.go) - same rationale as sqlite: no large-object API to hold a
-- whole upload as one growable value.
CREATE TABLE IF NOT EXISTS tus_upload_chunks (
    upload_id varchar(191) NOT NULL,
    chunk_offset bigint NOT NULL,
    data longblob NOT NULL,
    PRIMARY KEY (upload_id, chunk_offset),
    CONSTRAINT fk_tus_upload_chunks_upload
        FOREIGN KEY (upload_id) REFERENCES tus_uploads(id) ON DELETE CASCADE
) ENGINE=InnoDB;

-- +goose Down
DROP TABLE IF EXISTS tus_upload_chunks;
DROP TABLE IF EXISTS tus_uploads;
