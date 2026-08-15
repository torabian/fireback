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

**The convenient way: `backup enable`.** Connected as a Postgres superuser,
it applies `wal_level=replica`, `archive_mode=on`, `archive_command`
(pointed at this same binary's own `backup archive-push`, so no separate
wal-g binary is needed for this path either - see "wal-g is embedded, not a
binary" above), and `archive_timeout` via `ALTER SYSTEM SET`, then reloads:

```
$ backup enable
archive_mode     "off" -> "on"
archive_command  "(disabled)" -> "WALG_FILE_PREFIX=/backups /usr/local/bin/app __walg-exec wal-push %p"
archive_timeout  "0" -> "60"
wal_level/archive_mode need a full Postgres restart (not just a reload) to actually take effect - restart it yourself when ready, or re-run with --restart --pgdata <dir> to have this do it now
```

`wal_level`/`archive_mode` are `PGC_POSTMASTER` parameters - Postgres can
only pick up a change to either on a full restart, which `backup enable`
does not do on its own unless you also pass `--restart --pgdata <dir>`
(shells out to `pg_ctl restart -D <dir> -w`) - restarting a database out
from under you without being asked to isn't something this does silently.
`archive_command`/`archive_timeout` need only the reload `backup enable`
already does. Re-running it is safe and accurately idempotent either way -
already-applied settings are left alone and not reported, and a setting
that's staged but still waiting on a restart is reported as pending rather
than re-issued every time (this needed real care: Postgres's own
`SHOW archive_command` always displays `"(disabled)"` while `archive_mode`
is off, *regardless of archive_command's real configured value* - a
documented quirk of its GUC show-hook, confirmed directly, which would
otherwise make this look unset forever until after the restart).

Doing it by hand instead - or on a deployment `backup enable` can't reach
(no direct superuser connection available) - replicate the same settings
manually. Already wired up in this repo's `docker-compose.yml` /
`Dockerfile.postgres`
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
| `backup enable [--archive-timeout <dur>] [--restart --pgdata <dir>]` | Configure Postgres itself for wal-g (`wal_level`, `archive_mode`, `archive_command`, `archive_timeout`) via `ALTER SYSTEM SET` - the one-shot version of "Postgres-side configuration" below. Requires a superuser connection. |
| `backup push [--full] --pgdata <dir>` | Take a base backup now (cron calls this on schedule). Delta unless `--full`. |
| `backup list` | List known backups with their creation time. |
| `backup restore --at <RFC3339> --target <dir> [--start] [--promote] [--port N]` | Fetch the nearest backup ≤ `--at` into `--target` and configure it to replay WAL to that instant. `--start` also starts Postgres there and waits. |
| `backup verify [--scratch <dir>] [--port N]` | Full restore drill: latest backup → replay everything → confirm Postgres actually promotes. Run on a schedule (e.g. weekly). |
| `backup check [--max-age <dur>]` | Fast, non-destructive health check for cron/alerting: fails if backups are stale or `wal-g wal-verify` finds WAL gaps. |
| `backup prune [--retain N]` | Keep the newest N full backups (+ what's needed to restore from them), delete the rest. Run after every successful push. |
| `backup download <name> <dest>` | Copy a backup + exactly the WAL it needs to `<dest>`, ready to move off-box. |
| `backup archive-push` / `backup archive-get` | Thin wal-push/wal-fetch wrappers with failure logging to `<prefix>/.nima-archive.log`, for use as `archive_command`/`restore_command` wherever `backup-cli` itself is built and installed next to Postgres (see "wal-g is embedded, not a binary" for why this repo's own dev container doesn't do that yet). |
| `backup dump [--database <name>] [--output <path>] [--hash]` | Plain single-database snapshot (`pg_dump`/`mysqldump`/sqlite `VACUUM INTO`, zipped) - postgres/mysql/sqlite, not just postgres. See "Dumping a single database" below - this is a different, simpler tool than everything above, which is wal-g/postgres-only. |
| `backup restore-dump [--database <name>] [--file <path> \| --hash <hash>] [--force]` | Loads a `backup dump` zip back into a database. Postgres/mysql: creates `--database` automatically if it doesn't exist; refuses if it already does, unless `--force`. Sqlite: same rule against the destination file. |

## Dumping a single database (postgres, mysql, sqlite)

`backup dump`/`backup restore-dump` are separate from every command above:
they don't touch wal-g or WAL archiving at all, and they work against
whichever `DB_VENDOR` the app is actually configured for (`postgres`,
`mysql`/`mariadb`, or `sqlite`), not postgres only. Use these when you just
want a portable, single-file copy of one database - the wal-g commands above
are for continuous archiving + true point-in-time restore of a whole
Postgres cluster.

Run `backup dump` with **no flags at all** and it's a short interactive
wizard: disk vs HTTP first, then the database (as below), then it suggests
an output filename, then it shows a rough size estimate and (for the disk
path) available disk space before actually running anything:

```
$ backup dump
? Where should the backup go?
  > Write to a local zip file
    Stream it over HTTP behind a one-time hash
? Select database to dump
  > shop_prod
    shop_staging
? Output file: (./backups/dumps/shop_prod_20260815.zip)
estimated size: ~1.2 GiB (database size on disk - the actual dump is often smaller)
available disk space (./backups/dumps): 84.3 GiB
dumped shop_prod to ./backups/dumps/shop_prod_20260815.zip
```

Passing any flag at all (`--database`, `--output`, or `--hash`) skips the
"which mode"/suggested-filename/size-estimate parts of the wizard and
behaves exactly as before - safe for cron/scripts. The one prompt that can
still happen either way is picking a database when `--database` is omitted
and more than one exists (unchanged, same as always):

```
# postgres/mysql: pick a database interactively if more than one is visible
# on the connection, otherwise just use the only one there is
backup dump --database shop_prod

# or choose where the zip goes too
backup dump --database shop_prod --output ./shop_prod.zip

# restore it back - creates shop_staging automatically since it doesn't
# exist yet; refuses instead if it does, to avoid silently overwriting or
# merging into whatever's already there (pass --force to do it anyway)
backup restore-dump --database shop_staging --file ./shop_prod.zip
```

restore-dump never assumes the target already exists the way a bare
`psql -f dump.sql -d shop_staging` would: for postgres/mysql it checks
first, creates the database itself if it's missing, and refuses outright if
it's already there (sqlite: same check against the destination file, minus
the create step - there's nothing to "create", just don't overwrite it).
`--force` is the escape hatch for "yes, I really do want to restore into
this existing one" - it does not drop or clean anything first, so a real
conflict (e.g. a table that already exists from a previous restore) still
fails with a normal `psql`/`mysql` error rather than silently succeeding
into a mixed state.

### Streaming a dump over HTTP behind a one-time hash

Instead of writing to disk, `--hash` (or picking "Stream it over HTTP" in
the wizard) registers a one-time job and hands back a hash/URL - no auth,
no separate token to configure:

```
$ backup dump --database shop_prod --hash
database:   shop_prod
hash:       f3a1...   (96 hex chars)
url:        http://localhost:4500/backup/dumps/f3a1.../raw
expires at: 2026-08-15T12:35:00+02:00 (sooner if fetched - a hash is disabled the instant it's used)
only works while this app's own HTTP server (`app start`/`app s`) is running - the dump itself starts the moment the URL above is fetched, streamed live, never written to disk anywhere first.
fetch once with: curl -OJ http://localhost:4500/backup/dumps/f3a1.../raw
```

The actual `pg_dump`/`mysqldump`/sqlite work only happens when something
`GET`s that URL - **on the fly**, streamed straight into the HTTP response
as it's produced, never buffered to a file on either end first. Registering
the hash doesn't dump anything by itself, and nothing about it requires
authentication anymore (an earlier version of this needed a
`BACKUP_API_TOKEN` bearer token to register a job - removed): the hash
itself is the only credential the fetch needs, by design, so it can be
handed to whoever needs the file - CI, another operator, a support ticket -
without giving them database credentials or any other access at all.

It's single-use and short-lived: the first fetch (successful or not)
**disables it immediately**, and an unclaimed one expires after
`BACKUP_DUMP_HASH_TTL_SECONDS` (default **1800s / 30 minutes**) regardless.

Job state is a small JSON file (database name + expiry, never the dump
content itself) under this OS's own per-user-private config directory -
`os.UserConfigDir()`, which resolves to `~/.config/fireback/backup-jobs`
(Linux), `~/Library/Application Support/fireback/backup-jobs` (macOS), or
`%AppData%\fireback\backup-jobs` (Windows). That's what makes this both
safe without a token (only whoever already has filesystem access as this OS
user could plant or read a job there) and possible without a network call
to register: `backup dump --hash` writes the job directly, and whatever
process later serves the `GET` (`app start`) just needs to be able to read
that same directory - **the two must run on the same host, as the same OS
user** (true of everything else this module already assumes, e.g. wal-g's
own push/restore - see the caveats at the top of this README). It's lost on
restart and not shared across multiple server replicas behind a load
balancer - this is a short-lived handoff mechanism, not a durable export.

`backup restore-dump --hash <hash> --server <url>` is the read side of the
same thing: it fetches straight from the URL and pipes it into
`psql`/`mysql` (or, for sqlite, spools it to a temp file first) without ever
writing the dump to disk on the machine running the restore. Unlike
registering a job, fetching one genuinely can happen from a different host
(`--server` points wherever the job was registered), since it's a plain
HTTP GET.

### Size estimate and available disk space

Before an interactive `backup dump` run actually dumps anything, it prints
a rough size estimate and (when writing to disk) how much space is free
where the output is going, cross-platform (`DiskSpace_unix.go`/
`DiskSpace_windows.go` - `statfs(2)` on Linux/macOS,
`GetDiskFreeSpaceEx` on Windows):

- Postgres: `pg_database_size()` - the database's actual on-disk size
  (tables + indexes + overhead). The dump itself is usually smaller.
- MySQL/MariaDB: `sum(data_length + index_length)` from
  `information_schema.tables` for that schema - same "usually smaller than
  the real dump" caveat.
- SQLite: the current file size - exact, since `VACUUM INTO` never produces
  something larger than its source.

This is informational only - a too-small disk warns but doesn't block the
dump (the estimate is rough enough that blocking on it would risk false
positives). It's skipped entirely for `--hash`/HTTP dumps, since those never
touch local disk at all.

### Uploaded files (modules/storage) are included - yes, even Postgres large objects

`modules/storage` stores uploaded file *content* differently per vendor -
`StoreMySQL.go`/`StoreSQLite.go` keep it as plain rows in a
`tus_upload_chunks` table (ordinary table data, no special handling needed:
`mysqldump`/`VACUUM INTO` already capture it exactly like every other
table), but `Store.go`'s **Postgres** backend uses real Postgres large
objects (`tx.LargeObjects()`/`pg_largeobject`, referenced by `oid` from
`tus_uploads.oid`) - a separate, database-wide storage mechanism outside
normal tables. `backup dump`/`restore-dump` include these correctly, with no
special flags needed: `pg_dump` includes large objects by default (unless
`--no-blobs` is passed, which this module never does), even in the default
plain-SQL format this module uses, and its restore script re-creates each
one with `lo_create('<original-oid>')` - the *exact original OID*, not a new
one - so `tus_uploads.oid` still points at the right blob after a restore
into an empty/fresh database.

This isn't a "should work" claim - it was verified end to end: created a
large object through the same `pgx` `LargeObjects()` API `Store.go` uses,
referenced its OID from a table row, ran `backup dump`, restored into a
separate fresh database, and confirmed both the OID (`24606` on both sides)
and the exported bytes matched the original exactly.

The one case this doesn't cover: restoring into a database that **already
has** a large object at that same OID (e.g. restoring a dump on top of an
existing, non-empty database rather than into a fresh one) - `lo_create`
with an explicit OID fails on a collision, which surfaces as a `psql` error
during restore rather than silent corruption, but means "restore into a
fresh/empty target" is the only case actually exercised so far.

### Dumping a live database - what happens to concurrent traffic?

Short version: nothing bad. All three vendors give you an atomic, consistent
snapshot as of the instant the dump started - concurrent writes are never
partially captured, and the live database is never blocked, paused, or
rolled back by the dump running against it:

- **Postgres** (`pg_dump`): runs inside one `REPEATABLE READ` transaction, so
  it reads an MVCC snapshot frozen at the moment the dump began. Rows
  committed after that instant simply aren't in the dump; nothing committed
  before it is missing or torn. Only concurrent DDL (`ALTER TABLE`, etc.)
  against tables being dumped can block or be blocked - normal reads/writes
  are unaffected.
- **MySQL/MariaDB** (`mysqldump --single-transaction`): same MVCC-snapshot
  idea, via InnoDB. This module always passes `--single-transaction` -
  without it, a dump of a database under concurrent write load can come out
  *torn* (earlier-dumped tables reflecting an earlier state than
  later-dumped ones), silently breaking cross-table referential consistency
  even though no individual statement failed.
- **SQLite** (`VACUUM INTO`): uses SQLite's own online-backup mechanism
  rather than copying the file's bytes directly (which could race a
  concurrent writer mid-page-write). A concurrent writer just sees the
  `VACUUM INTO` as another reader; the result is a valid, complete database
  file as of a consistent point.

Restoring later doesn't touch the live database it came from either - it's
loaded into a *different* target (a fresh/existing database you name via
`--database`, or a separate file for sqlite), never back into the source.

## Configuration

Every setting `Engine`/`Dump.go`/`RestoreDump.go` read - `WALG_FILE_PREFIX`,
`WALG_COMPRESSION_METHOD`, `WALG_DELTA_MAX_STEPS`, `BACKUP_RETAIN_FULL`,
`BACKUP_DUMP_DIR`, `BACKUP_DUMP_HASH_TTL_SECONDS`,
`PG_DUMP_BIN`, `PSQL_BIN`, `MYSQLDUMP_BIN`, `MYSQL_BIN` - is declared once in
`Backup.emi.yml`'s `config:` block and generated into
`modules/backup/defs/Configuration.go` by the emi compiler (`make defs`,
or directly: `./app emi compile --path modules/backup/Backup.emi.yml`), the
same way every other fireback module's config: block works (see
`Fireback.emi.yml`/`Abac.emi.yml`). Nothing in this module reads
`os.Getenv` for these directly anymore.

`BackupModuleSetup` (`BackupModule.go`) is what makes this actually take
effect: it registers the module's `Config` with fireback's combined
`config list`/`config <field> get`/`config <field> set` CLI (see
`backup config` and `ConfigRegistry.go`), the same way `AbacModule.go` does
for `Abac.emi.yml`. Register it once in your app's `main.go`:

```go
modules := []*application.ModuleProvider{
    // ...
    backup.ModuleSetup(nil),
}
```

instead of hand-assembling a bare `CliHandlers`-only `ModuleProvider` (what
this repo's own `cmd/fireback/main.go` used to do) or setting these env vars
by convention alone.

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
