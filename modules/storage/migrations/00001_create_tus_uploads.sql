-- +goose Up
CREATE TABLE IF NOT EXISTS tus_uploads (
    id text PRIMARY KEY,
    oid oid NOT NULL,
    size bigint NOT NULL DEFAULT 0,
    size_is_deferred boolean NOT NULL DEFAULT false,
    upload_offset bigint NOT NULL DEFAULT 0,
    metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
    is_partial boolean NOT NULL DEFAULT false,
    is_final boolean NOT NULL DEFAULT false,
    partial_uploads text[] NOT NULL DEFAULT '{}',
    completed boolean NOT NULL DEFAULT false,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    completed_at timestamptz
);

CREATE UNIQUE INDEX IF NOT EXISTS tus_uploads_oid_idx ON tus_uploads (oid);

-- +goose Down
DROP TABLE IF EXISTS tus_uploads;
