package storage

import (
	"context"
	"fmt"
	"path"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

// joinMountPath joins a base URL path onto an optional mount point,
// normalizing slashes: the result always has exactly one leading slash and
// no doubled-up or missing slashes between segments, regardless of whether
// mountPoint/base individually have leading/trailing slashes or none.
func joinMountPath(mountPoint, base string) string {
	return path.Join("/", mountPoint, base)
}

func (c *StorageModuleConfig) withDefaults() *StorageModuleConfig {
	out := StorageModuleConfig{}
	if c != nil {
		out = *c
	}
	if out.UploadsBasePath == "" {
		out.UploadsBasePath = "/files"
	}
	if out.DownloadsBasePath == "" {
		out.DownloadsBasePath = "/downloads"
	}

	mountPoint := "storage"
	if out.MountPoint != "" {
		mountPoint = out.MountPoint
	}
	out.UploadsBasePath = joinMountPath(mountPoint, out.UploadsBasePath)
	out.DownloadsBasePath = joinMountPath(mountPoint, out.DownloadsBasePath)
	if out.ReaperInterval == 0 {
		out.ReaperInterval = time.Hour
	}
	if out.UnclaimedTTL == 0 {
		out.UnclaimedTTL = 24 * time.Hour
	}
	if out.IncompleteTTL == 0 {
		out.IncompleteTTL = 24 * time.Hour
	}
	return &out
}

var (
	storeMu     sync.RWMutex
	sharedStore Store
)

// GetStore returns the Store this module was wired up with, once MountAll
// (directly, or via the GinWebServerInitHooks hook StorageModuleSetup
// registers) has run - nil before that. Other modules that need to call
// ClaimFile/ReleaseFile can reuse this instead of opening their own
// connection to the same database. Vendor-agnostic - works the same whether
// the app is on postgres or sqlite.
func GetStore() Store {
	storeMu.RLock()
	defer storeMu.RUnlock()
	return sharedStore
}

// Pool returns the pgxpool.Pool backing GetStore(), or nil if either
// nothing has been mounted yet or the current Store isn't Postgres-backed
// (e.g. the app is on sqlite). Kept for existing callers written against
// the Postgres-only pool directly; GetStore/the Store-taking functions
// (ClaimFile, ReleaseFile, UsedBytes, ...) are the vendor-agnostic way to
// reach the same operations and work regardless of vendor.
func Pool() *pgxpool.Pool {
	if s, ok := GetStore().(*pgLoStore); ok {
		return s.Pool
	}
	return nil
}

func setStore(s Store) {
	storeMu.Lock()
	sharedStore = s
	storeMu.Unlock()
}

// MountAll wires the tus upload endpoints (Mount), the read/download
// endpoints (MountDownloads), and the background orphan reaper (StartReaper)
// onto r, all backed by store. It's the single place both
// cmd/fileupload-server and the GinWebServerInitHooks hook below call into,
// so there's one code path for standing this module up on a gin router,
// regardless of which Store implementation (Postgres or sqlite) backs it.
//
// cfg.DisableUploadsMount/DisableDownloadsMount skip Mount/MountDownloads
// respectively (the reaper still runs regardless, controlled separately by
// cfg.ReaperInterval).
//
// It does not run migrations - callers running inside a full fireback app
// get the tus_uploads table for free via ModuleProvider.GoMigrateDirectory;
// standalone use should call Migrate/MigrateSQLite first (or use
// OpenStoreForFireback, which already does).
//
// The reaper runs in its own goroutine (started by StartReaper), so this
// call itself returns immediately - it never blocks server startup. Cancel
// ctx to stop the reaper, e.g. on shutdown.
func MountAll(ctx context.Context, r gin.IRouter, store Store, cfg *StorageModuleConfig) (Store, error) {
	cfg = cfg.withDefaults()

	setStore(store)

	if !cfg.DisableUploadsMount {
		if err := Mount(r, cfg.UploadsBasePath, store, cfg); err != nil {
			return nil, err
		}
	}
	if !cfg.DisableDownloadsMount {
		MountDownloads(r, cfg.DownloadsBasePath, store, cfg)
	}

	if cfg.ReaperInterval > 0 {
		StartReaper(ctx, store, cfg.ReaperInterval, cfg.UnclaimedTTL, cfg.IncompleteTTL)
	}

	return store, nil
}

// NewPgxPool derives fireback's own database configuration (the same DSN
// its GORM layer connects to) and opens a fresh pgxpool against it, running
// migrations first. It calls fireback.LoadConfiguration itself rather than
// GetConfig, so it's self-sufficient whether it's called mid-bootstrap (from
// mountOnFirebackApp, where config is already loaded) or standalone.
//
// Returns (nil, nil) - not an error - if fireback isn't configured for
// postgres: large objects are a postgres-only feature, so there's nothing
// to connect to. Kept as its own function (rather than folded into
// OpenStoreForFireback) for existing callers that specifically want a raw
// pgxpool.Pool rather than a Store.
func NewPgxPool() (*pgxpool.Pool, error) {
	vendor, dsn := fireback.GetDatabaseDsn(fireback.LoadConfiguration())
	if vendor != "postgres" {
		return nil, nil
	}

	if err := Migrate(dsn); err != nil {
		return nil, err
	}

	return pgxpool.New(context.Background(), dsn)
}

// OpenStoreForFireback derives fireback's own database configuration (the
// same one its GORM layer connects to) and opens whichever Store
// implementation matches its vendor, running that vendor's migrations
// first - the vendor-agnostic counterpart to NewPgxPool, and what
// mountOnFirebackApp/withStore actually use.
//
// Returns (nil, nil) - not an error - for any vendor this module doesn't
// support (anything but postgres, sqlite, and mysql/mariadb), and also for
// sqlite specifically configured as ":memory:" or empty: a bare ":memory:"
// sqlite database is private to whichever single connection opened it, so a
// second, separate connection (this module always opens its own, mirroring
// the postgres path - see OpenSQLiteStore) would just be an entirely
// different, empty in-memory database, not sharing anything with the app's
// own GORM connection. There's no meaningful "storage on sqlite :memory:"
// story without a real file to point both connections at, so this silently
// does nothing instead - the same fallback this module has always had for
// any database vendor it doesn't know how to back itself with.
func OpenStoreForFireback() (Store, error) {
	vendor, dsn := fireback.GetDatabaseDsn(fireback.LoadConfiguration())

	switch vendor {
	case "postgres":
		pool, err := NewPgxPool()
		if err != nil || pool == nil {
			return nil, err
		}
		return NewStore(pool), nil
	case "sqlite":
		if dsn == "" || dsn == ":memory:" {
			return nil, nil
		}
		return OpenSQLiteStore(dsn)
	case "mysql", "mariadb":
		return OpenMySQLStore(dsn)
	default:
		return nil, nil
	}
}

// mountOnFirebackApp is the GinWebServerInitHooks callback StorageModuleSetup
// registers: it opens whichever Store implementation matches the app's
// configured database vendor (a separate connection from whatever GORM
// itself uses - Postgres large objects need a native pgx transaction, which
// GORM's *sql.DB can't give us; the sqlite path mirrors that same shape for
// consistency, see OpenSQLiteStore), defaults cfg.Authenticate to
// defaultAuthenticate if the caller didn't supply their own, and calls
// MountAll against it.
func mountOnFirebackApp(g *gin.RouterGroup, cfg *StorageModuleConfig) error {
	store, err := OpenStoreForFireback()
	if err != nil {
		return err
	}
	if store == nil {
		// Nothing this module knows how to back itself with for the
		// configured database vendor - see OpenStoreForFireback's own doc
		// comment for exactly which cases silently no-op here.
		return nil
	}

	firebackCfg := StorageModuleConfig{}
	if cfg != nil {
		firebackCfg = *cfg
	}
	if firebackCfg.Authenticate == nil {
		firebackCfg.Authenticate = defaultAuthenticate
	}

	_, err = MountAll(context.Background(), g, store, &firebackCfg)
	return err
}

// Inject this in your gin server, when using this
var CorsConfig = cors.New(cors.Config{
	AllowAllOrigins: true,
	AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
	AllowHeaders:    []string{"Authorization", "X-Requested-With", "X-Request-ID", "X-HTTP-Method-Override", "Upload-Length", "Upload-Offset", "Tus-Resumable", "Upload-Metadata", "Upload-Defer-Length", "Upload-Concat", "User-Agent", "Referrer", "Origin", "Content-Type", "Content-Length", "Upload-Length"},
	ExposeHeaders:   []string{"Upload-Offset", "Location", "Upload-Length", "Tus-Version", "Tus-Resumable", "Tus-Max-Size", "Tus-Extension", "Upload-Metadata", "Upload-Defer-Length", "Upload-Concat", "Location", "Upload-Offset", "Upload-Length"},
})

// defaultAuthenticate resolves an incoming request's identity the same way
// fireback's own SecurityModel/WithAuthorization middleware does elsewhere
// in the app: via fireback.WithAuthorizationPure, which dispatches to
// abac's real token lookup once abac.WorkspaceModuleSetup() has run (as it
// does in any app that lists it in Modules). No specific permission is
// required here - only that the Authorization header resolves to a real,
// non-deleted user; wrap this with your own Authenticate if an endpoint
// should also require a permission.
//
// WorkspaceId is taken directly from the Workspace-id header, unverified -
// this mirrors AuthorizeRequest's own behavior; fireback doesn't check at
// this layer that the resolved user actually belongs to that workspace.
// AccessLevel is taken directly from the Role-id header for a similar
// reason: fireback's auth primitive doesn't resolve a single canonical
// "access level" value - access there is a per-workspace set of capability
// strings, not a rank - see AuthContext's doc comment.
func defaultAuthenticate(c *gin.Context) (AuthContext, error) {
	workspaceId := c.GetHeader("Workspace-id")

	result, ierr := fireback.WithAuthorizationPure(&fireback.AuthContextDto{
		WorkspaceId:  workspaceId,
		Token:        c.GetHeader("Authorization"),
		Capabilities: []application.PermissionInfo{},
	})
	if ierr != nil {
		return Anonymous, fmt.Errorf("fileupload: authentication failed: %s", ierr.Json())
	}

	return AuthContext{
		UserId:      result.UserId.OrDefault(""),
		WorkspaceId: workspaceId,
		AccessLevel: c.GetHeader("Role-id"),
	}, nil
}
