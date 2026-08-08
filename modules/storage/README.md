# storage

tus-protocol file upload module (Go package `github.com/torabian/nima/modules/storage`,
fireback module name `"fileupload"`). Uploaded bytes are stored as PostgreSQL
large objects (`lo_*`); upload bookkeeping (offset, metadata, completion
state, owner, claim state) lives in the `tus_uploads` table.

The storage/HTTP layer (`Store.go`, `Handler.go`, `Download.go`, `Claim.go`,
`Quota.go`, `Queries.go`, `Reaper.go`, `Admin.go`) is plain Go/pgx — no
Fireback/Emi entities or codegen involved, per `StorageModule3.yml`.
`StorageModule.go`/`Webserver.go` are the one place this plugs into Fireback:
they register a `GinWebServerInitHooks` callback (mounts the HTTP routes) and
a `CliHandlers` entry (mounts the admin CLI, see §4) so any app that lists
`storage.StorageModuleSetup(...)` in its `Modules` gets both for free, with
no manual wiring in that app's `main.go`.

> **Naming note**: the Go package, directory, and config/type names (`storage`,
> `StorageModuleConfig`, ...) all say "storage", but `StorageModuleSetup`
> registers the underlying `application.ModuleProvider` under the name
> `"fileupload"` (`StorageModule.go:87`), and several error strings/log lines
> still say `"fileupload: ..."`. Both names refer to the same module — this
> is a leftover from an in-progress rename, not two different things.

| File                          | Purpose                                                                                                                                                                                                                                        |
| ----------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `migrations/*.sql`            | `tus_uploads` schema: create table → add `claimed_by`/`claimed_at` → add `user_id`/`workspace_id`/`access_level`                                                                                                                               |
| `Migrate.go`                  | `Migrate(connString)` — runs the migrations via goose                                                                                                                                                                                          |
| `Store.go`                    | `Store` — implements tus's `handler.DataStore`/`TerminaterDataStore` on top of Postgres large objects                                                                                                                                          |
| `Handler.go`                  | `Mount(router, basePath, store, cfg)` — wires the tus HTTP endpoints into gin, enforces auth/ownership/quota                                                                                                                                   |
| `Auth.go`                     | `AuthContext`, `Anonymous` — the identity type every request/upload/quota check is expressed in terms of                                                                                                                                       |
| `Quota.go`                    | `UsedBytes` — per-user byte total; `DefaultQuotaBytes`, `ErrQuotaExceeded`                                                                                                                                                                     |
| `Queries.go`                  | `ListFiles` / `GetFile` — read upload metadata without touching tusd types                                                                                                                                                                     |
| `Download.go`                 | `MountDownloads(router, basePath, pool, store, cfg)` — JSON metadata + range-capable download endpoints                                                                                                                                        |
| `Claim.go`                    | `ClaimFile` / `ReleaseFile` — attach/detach an upload to whatever feature ends up using it                                                                                                                                                     |
| `Reaper.go`                   | `SweepOrphaned` / `StartReaper` — deletes uploads nothing ever claimed                                                                                                                                                                         |
| `Admin.go`                    | `WorkspaceUsedBytes`, `DeleteFile`, `UploadFile` — the same operations as HTTP, callable directly with no HTTP round trip (see §4)                                                                                                             |
| `Cli.go`                      | `Commands()` — urfave/cli/v3 wrapper (`usage`/`upload`/`delete`) around `Admin.go`, resolving its own Postgres pool (see §4.3)                                                                                                                 |
| `Webserver.go`                | `MountAll` — wires everything above onto a gin router in one call; `Pool()` — shared pgxpool other modules can reuse; `NewPgxPool()` — opens a pool from fireback's own DB config; `mountOnFirebackApp` — the `GinWebServerInitHooks` callback |
| `StorageModule.go`            | `StorageModuleSetup(cfg)` — the `application.ModuleProvider`; `cfg` controls base paths, reaper TTLs, auth, quota                                                                                                                              |
| `StorageModule3.yml`          | Documentation-only manifest — explicitly **not** a real Module3 definition; nothing here is code-generated from it                                                                                                                             |
| `StorageModule.dyno.go`       | Fireback-generated scaffolding (permission constant `PERM_ROOT_STORAGE_EVERYTHING`); regenerated by the fireback CLI, don't hand-edit                                                                                                          |
| `Manifest.go`                 | Unused Fireback-generated scaffold type (`Manifest{DB, FilterResolver}`) — nothing in this module references it                                                                                                                                |
| `metas/`, `queries/`          | Empty `embed.FS` scaffolds Fireback generates for every Module3-style module; unused here since this module keeps its SQL in `migrations/` instead                                                                                             |
| `upload-test/`                | Node.js resumable-upload test client (`tus-js-client`) exercising the tus endpoint end-to-end, including resuming after a killed process                                                                                                       |
| `../../cmd/fileupload-server` | Standalone gin server that mounts this module on its own, for exercising it without booting all of nima-server                                                                                                                                 |
| `../../cmd/fileupload-cli`    | Standalone CLI binary wrapping `storage.Commands()` — usage/upload/delete without any server running                                                                                                                                           |

## 1. Data model

Everything lives in one table, `tus_uploads` (see `migrations/*.sql`):

| Column                                      | Meaning                                                                                              |
| ------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `id` (text, PK)                             | The tus upload id (UUID), also the public id used in URLs and by CLI commands                        |
| `oid` (oid, unique)                         | Postgres large object id holding the actual bytes                                                    |
| `size`, `size_is_deferred`, `upload_offset` | Declared size / whether it was unknown up front / bytes written so far                               |
| `metadata` (jsonb)                          | Client-supplied tus metadata (e.g. `filename`, `filetype`) — **not** the reserved owner keys below   |
| `is_partial`, `is_final`, `partial_uploads` | tus concatenation-extension bookkeeping (unused by this app's clients today, but supported)          |
| `completed`, `completed_at`                 | Whether the upload has received all its declared bytes                                               |
| `claimed_by`, `claimed_at`                  | Set by `ClaimFile` — see §5                                                                          |
| `user_id`, `workspace_id`, `access_level`   | The `AuthContext` resolved for whoever created the upload — `NULL` for all three if made anonymously |

The actual file bytes are **not** in this table — they're in Postgres's
`pg_largeobject` catalog, referenced by `oid`. Deleting the row without also
calling `lo_unlink(oid)` leaks the bytes; always delete through
`Store.Terminate` (i.e. the `DELETE` HTTP endpoint, `Admin.DeleteFile`, or
`SweepOrphaned`) — never a bare `DELETE FROM tus_uploads`.

## 2. Using it inside a fireback app (e.g. nima-server)

Nothing to wire up beyond listing the module, same as any other fireback
module (`cmd/nima-server/main.go`):

```go
Modules: []*application.ModuleProvider{
	// ...
	storage.StorageModuleSetup(&storage.StorageModuleConfig{
		MountPoint: "/storage/",
		Authenticate: func(c *gin.Context) (storage.AuthContext, error) {
			// resolve real identity here; nima-server's example hardcodes
			// one for demo purposes
			return storage.AuthContext{UserId: "12", WorkspaceId: "233"}, nil
		},
	}),
},
```

When the app's gin server starts, `mountOnFirebackApp` (`Webserver.go:136`)
runs automatically: it opens a `*pgxpool.Pool` against the same Postgres
connection fireback's own GORM layer uses (via `NewPgxPool`, see §4.3),
running migrations first, and calls `MountAll` — mounting the tus upload
endpoint, the download endpoints (§6), and starting the orphan reaper (§5).
If the app isn't on Postgres, this is a silent no-op (large objects are
Postgres-only). Its `CliHandlers` entry also adds the `usage`/`upload`/
`delete` admin commands to the app's own CLI under `storage` (§4.3).

If `cfg.Authenticate` is left `nil`, every request is anonymous — no owner
is recorded, `Quota` is never consulted, and ownership checks let any
request through. Leave it unset only for trusted/internal deployments.

### Mount path defaults

`StorageModuleConfig.withDefaults()` (`Webserver.go:23`) always prefixes
both base paths with a mount point — `MountPoint` if set, otherwise the
literal `"storage"`. So with `cfg == nil` (or `MountPoint` unset), routes end
up at `/storage/files` and `/storage/downloads`, not `/files`/`/downloads`.
Set `MountPoint` to change the prefix; there is currently no way to mount
with _no_ prefix at all.

### Disabling the upload or download routes

`DisableUploadsMount`/`DisableDownloadsMount` (`StorageModule.go`) skip
mounting `Mount`/`MountDownloads` respectively in `MountAll`
(`Webserver.go:88`) — either can be omitted independently:

```go
storage.StorageModuleSetup(&storage.StorageModuleConfig{
	// No public write endpoint - files only ever arrive via
	// Admin.UploadFile / the CLI (§4.2/§4.3).
	DisableUploadsMount: true,
})

storage.StorageModuleSetup(&storage.StorageModuleConfig{
	// No built-in download route - app serves files through its own
	// signed-URL/CDN layer instead.
	DisableDownloadsMount: true,
})
```

The reaper (§5) and the in-process/CLI paths (§4.2/§4.3) are unaffected by
either flag — they only control what gets mounted onto the gin router.

## 3. Using it standalone (no fireback app)

You need a `*pgxpool.Pool` (separate from whatever connection GORM/Fireback
uses elsewhere) and to run the migrations yourself first — see
`cmd/fileupload-server/main.go` for the full example:

```go
if err := storage.Migrate(connString); err != nil {
	log.Fatal(err)
}

pool, err := pgxpool.New(context.Background(), connString)
if err != nil {
	log.Fatal(err)
}

r := gin.Default()
if _, err := storage.MountAll(context.Background(), r, pool, nil); err != nil {
	log.Fatal(err)
}
r.Run(":8099")
```

`MountAll` (`Webserver.go:88`) is the same entry point `mountOnFirebackApp`
uses, so both paths behave identically once you have a pool.

## 4. Three ways to call this module

Every operation this module offers — upload, delete, read metadata, check
usage — is reachable through three different paths, all ultimately calling
the same `Store`/`Admin.go` functions:

### 4.1 External: HTTP (tus protocol + downloads)

The only path a browser, `curl`, or a tus client (e.g. `upload-test/index.js`)
can use. Goes through `Mount`/`MountDownloads` (`Handler.go`, `Download.go`),
which run every request through `cfg.Authenticate` and enforce ownership
(`authorizeOwner`, `Auth.go:43`) and quota (`enforceQuotaCallback`,
`Handler.go:35`) before touching `Store`. See §6/§7 for the full route table
and semantics.

### 4.2 Internal: direct Go calls from other modules, in-process

Any other module in the same process (e.g. `score`) can call this module's
exported functions directly — no HTTP, no `cfg.Authenticate` involved. This
is how a feature attaches an upload to its own record:

```go
score, err := ScoreActions.Create(dto, query)
if err == nil {
	_, claimErr := storage.ClaimFile(ctx, storage.Pool(), dto.SheetFileId, "score:"+score.UniqueId)
}
```

`storage.Pool()` (`Webserver.go:63`) returns the same `*pgxpool.Pool`
`MountAll`/`mountOnFirebackApp` already opened, so callers don't need their
own connection. `Admin.go`'s `DeleteFile`/`UploadFile`/`WorkspaceUsedBytes`
and `Quota.go`'s `UsedBytes` are equally callable this way for any in-process
administrative logic (a scheduled job, an admin-only API route elsewhere in
the app, etc.) — they bypass `cfg.Authenticate`/ownership/quota entirely,
since the caller is already trusted code running inside the app.

### 4.3 Administrative: the CLI (`storage.Commands()`)

`Cli.go` wraps `Admin.go`'s operations as three urfave/cli/v3 subcommands —
`usage`, `upload`, `delete` — callable two ways:

**As the standalone `fileupload-cli` binary** (`cmd/fileupload-cli/main.go`),
pointed at any database with no app or server running:

```sh
go run ./cmd/fileupload-cli usage --user 12
go run ./cmd/fileupload-cli usage --workspace 233
go run ./cmd/fileupload-cli upload ./somefile.zip --user 12 --claim "score:abc123"
go run ./cmd/fileupload-cli delete <upload-id>
```

**Mounted inside `nima-server`'s own CLI**, via the `CliHandlers` entry
`StorageModuleSetup` registers (`StorageModule.go:101`) — same commands,
under a `storage` group:

```sh
go run ./cmd/nima-server storage usage --user 12
go run ./cmd/nima-server storage upload ./somefile.zip
go run ./cmd/nima-server storage delete <upload-id>
```

Both forms share the exact same `Commands()`/`withPool` (`Cli.go:19,49`);
what differs is how the Postgres connection gets resolved, tried in order:

1. **`Pool()`** — if this process already has the module mounted (e.g. you
   ran a command inside a long-lived app that already called `MountAll`),
   reuse that pool as-is.
2. **`--database-url` flag / `DATABASE_URL` env var** — what the standalone
   binary relies on, since it has no app config to fall back to.
3. **`NewPgxPool()`** (`Webserver.go:117`) — derives the DSN from fireback's
   own config (`fireback.LoadConfiguration` + `GetDatabaseDsn`), the same
   source `mountOnFirebackApp` uses. This is what lets `nima-server storage
usage ...` work with **no flag at all**: it resolves the same database
   the running app's web server would.

If none of the three resolve to a postgres connection, the command fails
with an explicit error rather than silently doing nothing.

`upload`/`delete`/`usage` bypass `cfg.Authenticate`/ownership/quota the same
way §4.2 does — there is no concept of "logged in as" for a CLI invocation,
so `--user`/`--workspace`/`--access-level` on `upload` just set the owner
columns directly, as if the given identity had uploaded it over HTTP.

## 5. Claiming an upload (avoiding orphaned files)

The tus upload endpoint needs no prior context — a client uploads a file
_before_ whatever record is meant to reference it exists. If that record
never gets created (the user picks a file for a new score, then abandons the
form), the upload has nothing tying its lifetime to anything else, and would
otherwise sit there forever: wasted storage, and free anonymous file hosting
for anyone who can reach the upload endpoint.

The fix is a two-step contract:

1. Whichever feature ends up using an upload calls `storage.ClaimFile`
   (`Claim.go:52`) once it commits to that — typically right after it saves
   its own record, in the same handler. `claimedBy` is just an
   application-defined string (e.g. `"<module>:<uniqueId>"`) — this module
   doesn't need to know about scores, videos, or anything else that might
   reference an upload. Claiming twice with the same `claimedBy` is a no-op
   success (safe to retry); claiming with a different `claimedBy` than
   what's already recorded returns `ErrUploadAlreadyClaimed`.

   `storage.ReleaseFile(ctx, pool, id, claimedBy)` (`Claim.go:105`) undoes a
   claim (e.g. when the owning record is deleted), making the upload
   eligible for reaping again.

2. `SweepOrphaned` (`Reaper.go:18`) deletes whatever is still unclaimed past
   a grace period — completed uploads that sat with `claimed_at IS NULL` for
   longer than `unclaimedTTL`, and uploads that never even finished
   (abandoned mid-transfer) for longer than `incompleteTTL`. `StartReaper`
   runs it on a fixed interval in its own goroutine — this is what
   `MountAll` starts automatically (configurable via
   `StorageModuleConfig.ReaperInterval`/`UnclaimedTTL`/`IncompleteTTL`, or
   disable it by setting `ReaperInterval` to a negative value).

Deletion always goes through `Store.Terminate` (`Store.go:300`), so both the
`tus_uploads` row and its backing large object are removed together —
whether triggered by `DELETE /storage/files/:id`, `storage.DeleteFile`, or
the reaper.

## 6. HTTP route reference

Mounted by `MountAll`/`mountOnFirebackApp` at `<MountPoint>/files` (tus
protocol, default mount point `"storage"` — see §2) and
`<MountPoint>/downloads` (read-only):

| Method | Path                         | Purpose                                                                 |
| ------ | ---------------------------- | ----------------------------------------------------------------------- |
| POST   | `/storage/files`             | Create an upload (tus `Upload-Length`/`Upload-Metadata` headers)        |
| HEAD   | `/storage/files/:id`         | Current offset (`Upload-Offset` header) — for resuming                  |
| PATCH  | `/storage/files/:id`         | Write the next chunk at the client's current offset                     |
| GET    | `/storage/files/:id`         | tus's own read (no `Range` support — use `/downloads/:id/raw` instead)  |
| DELETE | `/storage/files/:id`         | Terminate: removes both the large object and the `tus_uploads` row      |
| GET    | `/storage/downloads/:id`     | JSON metadata: `size`, `metadata`, `completed`, timestamps, claim state |
| HEAD   | `/storage/downloads/:id/raw` | Same headers a GET would send, no body                                  |
| GET    | `/storage/downloads/:id/raw` | The file bytes, with `Range`/`If-Range`/`If-None-Match` support         |

All of the above run through `cfg.Authenticate` (when set) and
`authorizeOwner`: an upload with a recorded owner (`user_id` not `NULL`) is
only readable/writable by the matching, non-anonymous identity; an upload
with no recorded owner is open to anyone (`Auth.go:43`).

### `GET /downloads/:id/raw` details (`Download.go:102`)

- `Content-Type`/`Content-Disposition` come from the upload's `filetype`/
  `filename` metadata, sanitized: an unrecognized/missing `filetype` always
  falls back to `application/octet-stream` + `attachment`, so a crafted
  metadata value can't turn this into a stored-XSS vector when opened
  directly in a browser (`contentTypeAndDisposition`, `Download.go:89`).
- `Range: bytes=START-END` / `bytes=START-` / `bytes=-N` → `206 Partial
Content`. Lets a resumed download continue from an offset, or a download
  manager fetch several ranges concurrently over separate connections.
- `If-Range`/`If-None-Match` are checked against the file's `ETag`
  (`"<id>-<size>"`).
- An out-of-bounds range → `416 Range Not Satisfiable`.
- Multiple comma-separated ranges in one `Range` header
  (`multipart/byteranges`) are not implemented — falls back to sending the
  whole file (`200`), which the RFC explicitly permits.
- Only completed uploads (`rec.Completed`) are served — a link handed out
  mid-upload 404s rather than racing the writer.

### Caveat: embedding an owned file in `<img src="...">`

If the upload has a recorded owner, the download route requires
`cfg.Authenticate` to succeed — normally the `Authorization` header
(`defaultAuthenticate`, `Webserver.go:175`). A plain HTML `<img src="...">`
(or CSS `background-image`, etc.) cannot attach a custom header, so the
browser's request resolves to `Anonymous` and gets `403`. Anonymous/unowned
uploads work fine directly in `<img>`; for owned uploads, either:

- fetch the bytes via JS `fetch()` with the header set, then
  `URL.createObjectURL(blob)` into the `<img>`'s `src`, or
- supply a custom `Authenticate` that also accepts a short-lived signed
  token in the query string, since `Authenticate` is a plain
  `func(*gin.Context) (AuthContext, error)` — nothing here currently
  implements that, but nothing stops an app from doing so.

## 7. Auth, ownership, and quota model

- `AuthContext{UserId, WorkspaceId, AccessLevel}` (`Auth.go:11`) is the one
  identity type every check below is expressed in terms of. `Anonymous` is
  its zero value.
- `cfg.Authenticate` resolves it per-request; `nil` disables the whole
  mechanism (§2). Inside a fireback app, `StorageModuleSetup` defaults it to
  `defaultAuthenticate` (`Webserver.go:175`), which reads the `Authorization`
  header the same way fireback's own middleware does elsewhere, plus
  `Workspace-id`/`Role-id` headers for the other two fields — **unverified**:
  fireback doesn't check the resolved user actually belongs to that
  workspace at this layer.
- Ownership (`authorizeOwner`, `Auth.go:43`): a `NULL` owner is open to
  anyone; a non-`NULL` owner requires an exact, non-anonymous `UserId` match.
- Quota (`Quota.go`, `Handler.go:35`): checked once per upload creation,
  before any bytes are written, against `UsedBytes` (sum of `size` across
  all of that user's uploads, completed or in-progress) plus the new
  upload's declared size. Defaults to `DefaultQuotaBytes` (2 GiB) per user;
  `cfg.Quota` overrides it per-identity, and a negative return means
  unlimited. Never consulted for anonymous uploads. **Only enforced on the
  HTTP creation path** — `Admin.UploadFile` (CLI/internal, §4.2/§4.3) does
  not call it.

## Notes

- Large objects are transaction-scoped in Postgres, so `WriteChunk` and
  `GetReader` each open their own short-lived transaction — don't try to
  reuse a `pgx.LargeObject` outside the call that opened it.
- `DELETE` (HTTP), `Admin.DeleteFile`, `fileupload-cli delete`, and the
  reaper all funnel through the same `Store.Terminate` — never delete a
  `tus_uploads` row with raw SQL, or you leak its large object (see §1).
- Regular Postgres `VACUUM`/autovacuum applies here like any other table:
  deleting an upload doesn't shrink Postgres's on-disk files immediately —
  see `pg_stat_all_tables`/`VACUUM` if you need to reclaim space sooner than
  autovacuum's own schedule.
- No resumable-upload support on the Go CLI side (`fileupload-cli upload`
  always writes the whole file in one call, §4.3); the HTTP endpoint itself
  is fully tus-resumable (that's `tusd`'s handler), which is what
  `upload-test/index.js` exercises end-to-end, including resuming after a
  killed process.
