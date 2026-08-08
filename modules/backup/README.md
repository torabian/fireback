# backup

Point-in-time Postgres backup/restore, built on top of [wal-g](https://github.com/wal-g/wal-g)
(continuous WAL archiving + full/delta base backups) rather than `pg_dump` -
`pg_dump` gives discrete snapshots whose size and restore time both scale
with the full database, which doesn't hold up as the database grows. WAL-G
decouples restore granularity from database size: a periodic base backup
plus every WAL segment archived since lets you replay forward to any
instant, not just to the last dump.

This module is an orchestration layer, not a reimplementation - it drives
wal-g for every actual backup/restore operation, and adds: a local catalog
of verify/prune history, automatic nearest-backup resolution for a given
point in time, restore-drill automation, and monitoring-friendly health
checks.

**wal-g is a real go.mod dependency (`github.com/wal-g/wal-g/cmd/pg`), not
a separately installed binary** - see "wal-g is embedded, not a binary"
below for how that actually works and the one place it still isn't true.

## Two important caveats before you rely on this

1. **Local/mounted storage is only a real disaster-recovery plan if it's
   genuinely separate from the database's own disk.** WAL-G's own docs say
   so explicitly: storing backups on the same disk as the database "is not
   safe... do not use it as a disaster recovery plan." It's fine (and what
   this repo's `docker-compose.yml` does) for protecting against *logical*
   mistakes - a bad migration, an accidental delete, a bug that corrupted
   data - but if the whole host/volume dies, an on-disk backup dies with it.
   For real DR, point `WALG_FILE_PREFIX` at NFS/SAN storage on separate
   physical infrastructure, or reconsider S3-compatible storage instead.

2. **`backup push` needs local filesystem access to Postgres's data
   directory**, not just a network connection to it. WAL-G supports a
   remote/streaming mode (just `PGHOST`, no data directory) for convenience,
   but that mode does not support delta backups - which defeats the "no
   matter how large" requirement this was built for. So `backup push` (and
   `verify`/`prune`, which depend on it having produced real backups) must
   run co-located with Postgres: inside the same container/host as
   `PGDATA`. In this repo's dev compose setup that means
   `docker compose exec postgres wal-g backup-push $PGDATA` directly, or
   running `backup-cli push --pgdata /var/lib/postgresql/data` from inside
   that container. `list`, `restore`, `check`, and `download` only read
   from storage and can run from anywhere with access to `WALG_FILE_PREFIX`.

## wal-g is embedded, not a binary

`modules/backup` imports `github.com/wal-g/wal-g/cmd/pg` directly as a real
go.mod dependency and drives its cobra command tree in-process - there is
no `wal-g` binary anywhere that `backup-cli`/`nima-server` themselves need
installed. This was verified for real, not just assumed:

- Calling wal-g's command tree a *second* time in the same process panics
  with `flag redefined` - its cobra/viper state is global and only ever
  meant to run once per process (wal-g's own `main()` only ever calls it
  once). So `Engine` (Engine.go) re-executes **this same binary** with a
  hidden `__walg-exec <wal-g args>` marker for every single wal-g
  operation, giving each one the fresh process it actually needs -
  `MaybeRunEmbeddedWalg` (Exec.go) is what makes that marker do something,
  and it must be the first thing both `main.go`s call, before urfave/cli
  ever sees `os.Args`.
- wal-g itself re-execs `os.Args[0]` internally for background WAL
  prefetching during `wal-fetch`
  (`internal/databases/postgres/wal_prefetcher.go`), as a bare
  `wal-prefetch ...` command with no notion of our marker convention -
  found by actually running a restore and seeing `No help topic for
  'wal-prefetch'` in the logs. `MaybeRunEmbeddedWalg` recognizes this one
  specifically too, or prefetching (a performance detail, not a
  correctness one) silently breaks the moment wal-g stops being a separate
  binary.
- wal-g's own logs go to stderr and its data payloads (e.g. `backup-list
  --json`) go to stdout even when run this way - confirmed directly, so
  `Engine.RunJSON`'s stdout-only parsing still works exactly like it would
  against a real subprocess.
- Because `github.com/wal-g/wal-g`'s go.mod doesn't use the `/v3` module
  path suffix Go modules requires for v2+ semver tags, it can't be
  `go get`'d by its `v3.0.8` tag directly - go.mod pins it by the tag's
  exact commit hash instead, which Go resolves to a pseudo-version
  (`v1.1.3-rc-with-build...-<hash>`); the source content is identical to
  the `v3.0.8` tag, only the version string looks unusual.
- wal-g's `internal/uploader.go` imports the still-experimental
  `encoding/json/v2` stdlib package directly, so **every** build of this
  repo now needs `GOEXPERIMENT=jsonv2` set (already added to
  `cmd/nima-server/Makefile`) - without it, `go build` fails with "build
  constraints exclude all Go files in .../encoding/json/v2".
- Embedding wal-g pulls in every cloud storage SDK it supports (AWS, GCS,
  Azure...) as transitive dependencies regardless of which one you
  actually use, since nothing here imports them selectively - `nima-server`
  grew from its previous size to ~160MB. That's the real cost of "go.mod
  dependency, not a binary."

**The one place this still isn't fully true:** this repo's own
`Dockerfile.postgres` installs a standalone, from-source-built `wal-g`
binary for Postgres's `archive_command` to call, rather than a from-scratch
container build of `backup-cli` itself. That's not a limitation of the
embedding approach - it's a separate, pre-existing issue: building
`backup-cli` inside a fresh container needs this repo's `fireback`
dependency to resolve from its published `v1.3.2` tag (the checked-out
`go.mod`'s `replace` directive, pointing at a local checkout, obviously
isn't portable into a container build), and that published tag doesn't
currently compile against the `emi` version pinned alongside it - confirmed
by actually trying it and hitting a real compile error in `fireback`'s own
code, not a guess. Fix that upstream incompatibility (or make the local
fireback checkout available in the container build context) and
`Dockerfile.postgres` can drop the standalone wal-g binary entirely in
favor of `backup-cli archive-push %p` / `backup-cli __walg-exec wal-fetch
%f %p` - `restore_command` generated by `backup restore` already points at
`backup-cli` this way (see `restoreCommandString` in Restore.go), since
that only ever runs on a host where `backup-cli` itself was built normally.

## Postgres-side configuration

Already wired up in this repo's `docker-compose.yml` / `Dockerfile.postgres`
(custom image with `wal_level=replica`, `archive_mode=on`,
`archive_command='wal-g wal-push %p'`, `archive_timeout=60`, plus a
`./backups:/backups` bind mount). For a deployment not using that compose
file, replicate the same three settings; `archive_command` needs *some*
executable Postgres can call (see the previous section for why that's a
standalone wal-g binary here specifically, vs. `backup-cli` elsewhere).

`Dockerfile.postgres` builds wal-g **from its own Go source** (pinned to
`WALG_VERSION`, a multi-stage build with a `golang:1.25-bookworm` builder
stage), rather than downloading a prebuilt release binary - nothing
pre-compiled from the internet ends up in the image, only source pinned to
a tag. Its default build (no brotli/libsodium/lzo tags, which we don't need
since `WALG_COMPRESSION_METHOD=lz4` is pure Go) needs no CGO/C libraries,
just plain `go build` (with `GOEXPERIMENT=jsonv2`, same reason as above).

`archive_timeout=60` bounds your worst-case RPO (recovery point objective)
to 60s even on a quiet database that isn't generating WAL fast enough to
trigger archiving on its own - tune it down further if you need tighter
RPO, at the cost of more, smaller WAL files.

## Commands

All commands read Postgres connection details from fireback's own config
when running inside a full app (`nima-server backup ...`), or from the
standard `PGHOST`/`PGPORT`/`PGUSER`/`PGPASSWORD`/`PGDATABASE` env vars when
run standalone (`backup-cli ...`) - the same convention wal-g itself uses.
`WALG_FILE_PREFIX` must always be set explicitly.

| Command | Purpose |
|---|---|
| `backup init` | Sanity check: embedded wal-g runs, storage prefix writable, config resolvable. Run this first. |
| `backup push [--full] --pgdata <dir>` | Take a base backup now (cron calls this on schedule). Delta unless `--full`. |
| `backup list` | List known backups with their creation time. |
| `backup restore --at <RFC3339> --target <dir> [--start] [--promote] [--port N]` | Fetch the nearest backup ≤ `--at` into `--target` and configure it to replay WAL to that instant. `--start` also starts Postgres there and waits. |
| `backup verify [--scratch <dir>] [--port N]` | Full restore drill: latest backup → replay everything → confirm Postgres actually promotes. Run on a schedule (e.g. weekly). |
| `backup check [--max-age <dur>]` | Fast, non-destructive health check for cron/alerting: fails if backups are stale or `wal-g wal-verify` finds WAL gaps. |
| `backup prune [--retain N]` | Keep the newest N full backups (+ what's needed to restore from them), delete the rest. Run after every successful push. |
| `backup download <name> <dest>` | Copy a backup + exactly the WAL it needs to `<dest>`, ready to move off-box. |
| `backup archive-push` / `backup archive-get` | Thin wal-push/wal-fetch wrappers with failure logging to `<prefix>/.nima-archive.log`, for use as `archive_command`/`restore_command` wherever `backup-cli` itself is built and installed next to Postgres (see "wal-g is embedded, not a binary" for why this repo's own dev container doesn't do that yet). |

## Scheduling

No in-process scheduler is needed or included - cron/systemd timers calling
the CLI directly are simpler and keep the CLI itself as the single source
of truth for what a backup run does:

```cron
# every 6 hours: delta backup, then enforce retention
0 */6 * * * docker compose exec -T postgres wal-g backup-push /var/lib/postgresql/data --full 2>&1 | logger -t nima-backup
15 */6 * * * backup-cli prune --retain 4

# weekly restore drill
0 3 * * 0 backup-cli verify

# every 5 minutes: monitoring
*/5 * * * * backup-cli check || alert-on-call "nima backup check failed"
```

(Use `--full` on a slower cadence than plain deltas - e.g. daily full,
6-hourly delta - once you've decided your `WALG_DELTA_MAX_STEPS`.)

## Disaster-recovery runbook

1. `backup-cli list` - confirm a backup exists at or before the point in
   time you need.
2. Pick an empty directory to restore into, e.g. `/var/lib/postgresql/restore`.
3. `backup-cli restore --at 2026-07-30T09:00:00Z --target /var/lib/postgresql/restore --start --promote`
   - fetches the right base backup, configures recovery, starts Postgres on
     port 5433 (won't collide with a live primary on 5432), and waits for
     recovery to finish.
4. Once it reports recovery complete, connect to port 5433 and confirm the
   data looks right *before* pointing anything else at it.
5. If this is a real incident (not a drill), stop the old primary, then
   either promote this instance into its place or `pg_dump`/logical-replicate
   the recovered data back into the original cluster, depending on how much
   else changed since the incident.
6. `backup-cli verify` doing steps 1-4 automatically (against `LATEST`
   instead of a specific time) is exactly this runbook rehearsed on a
   schedule - if it's been passing, this procedure is trustworthy; if it
   hasn't, fix that before you're relying on it live.

## Validated end-to-end

This has been run for real against `docker compose up --build`'s Postgres
16 + wal-g v3.0.8 container, not just built and reviewed:

- `backup-cli init`, `push --full`, `list` - all worked, and the catalog
  file recorded the pushed backup correctly.
- A genuine point-in-time restore: created a row, recorded a precise
  timestamp, created a second row, forced a WAL switch, then
  `backup-cli restore --at <timestamp-between-the-two> --start --promote`.
  Postgres logged `recovery stopping before commit of transaction ...` at
  exactly the requested instant, and the restored database contained the
  first row but correctly not the second.
- `backup-cli verify` (restore-to-latest drill), `backup-cli check`,
  `backup-cli prune`, and `backup-cli download` (via wal-g `copy`) all ran
  successfully against the same backup.
- Fixed two bugs this surfaced that pure code review wouldn't have caught:
  wal-g's actual flag is `--full`, not `--full-backup` as the prose docs
  implied; and the Postgres-vendor guard rejected explicit
  `PGHOST`/`PGDATABASE` env vars because fireback's own config loader
  defaults to `DB_VENDOR=sqlite` when unset.

Also confirmed: a delta backup (`push` without `--full`) correctly chains
off the prior full backup (`base_..._D_...` naming), and `wal-g` fetches
both automatically on restore.

One gotcha this surfaced, worth knowing before you page yourself over it:
`recovery_target_time` needs WAL evidence of a transaction at or after the
target instant. Restoring to "right now" when nothing has written to the
database since the last archived WAL segment fails with `recovery ended
before configured recovery target was reached` - not a bug, just how
Postgres's recovery target works. In practice this never comes up, since
you always restore to a specific *past* moment (e.g. "right before the bad
migration ran"), which necessarily already has WAL evidence.

**Then re-validated after switching wal-g from a shelled-out binary to an
embedded go.mod dependency** (see "wal-g is embedded, not a binary"),
including two more real bugs this second pass caught:

- A genuine sub-second-precision restore (row committed 11ms before the
  target instant, a second row 2s after): the *first* version of
  `recovery_target_time` formatting (`"...05-07"`, no fractional seconds)
  silently truncated the target down to the whole second, which could land
  *before* a transaction that happened earlier in the same second - the
  restore excluded a row it should have included. Fixed by always emitting
  full microsecond precision (`"...05.000000-07"`); re-ran the same
  sub-second scenario afterward and it restored exactly the two rows it
  should have.
- wal-g's own internal WAL-prefetch re-exec (bare `wal-prefetch ...`,
  unaware of this module's `__walg-exec` marker) - see the marker section
  above.

`backup-cli init`/`push`/`list`/`check`/`restore` (with the sub-second PITR
case above) and a full `nima-server` build (`make dev` in
`cmd/nima-server`) all passed with wal-g fully embedded and zero external
`wal-g` binary present in the container. `archive-push`/`archive-get` as
the actual configured `archive_command`/`restore_command` remains
unexercised, for the reason covered in "wal-g is embedded, not a binary" -
this repo's dev container still uses a standalone wal-g binary for that one
purpose.
