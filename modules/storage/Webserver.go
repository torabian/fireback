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
	poolMu     sync.RWMutex
	sharedPool *pgxpool.Pool
)

// Pool returns the pgxpool.Pool this module was wired up with, once MountAll
// (directly, or via the GinWebServerInitHooks hook StorageModuleSetup
// registers) has run - nil before that. Other modules that need to call
// ClaimFile/ReleaseFile can reuse this instead of opening their own
// connection to the same database.
func Pool() *pgxpool.Pool {
	poolMu.RLock()
	defer poolMu.RUnlock()
	return sharedPool
}

func setPool(p *pgxpool.Pool) {
	poolMu.Lock()
	sharedPool = p
	poolMu.Unlock()
}

// MountAll wires the tus upload endpoints (Mount), the read/download
// endpoints (MountDownloads), and the background orphan reaper (StartReaper)
// onto r, all backed by pool. It's the single place both
// cmd/fileupload-server and the GinWebServerInitHooks hook below call into,
// so there's one code path for standing this module up on a gin router.
//
// cfg.DisableUploadsMount/DisableDownloadsMount skip Mount/MountDownloads
// respectively (the reaper still runs regardless, controlled separately by
// cfg.ReaperInterval).
//
// It does not run migrations - callers running inside a full fireback app
// get the tus_uploads table for free via ModuleProvider.GoMigrateDirectory;
// standalone use should call Migrate first.
//
// The reaper runs in its own goroutine (started by StartReaper), so this
// call itself returns immediately - it never blocks server startup. Cancel
// ctx to stop the reaper, e.g. on shutdown.
func MountAll(ctx context.Context, r gin.IRouter, pool *pgxpool.Pool, cfg *StorageModuleConfig) (*Store, error) {
	cfg = cfg.withDefaults()

	store := NewStore(pool)
	setPool(pool)

	if !cfg.DisableUploadsMount {
		if err := Mount(r, cfg.UploadsBasePath, store, cfg); err != nil {
			return nil, err
		}
	}
	if !cfg.DisableDownloadsMount {
		MountDownloads(r, cfg.DownloadsBasePath, pool, store, cfg)
	}

	if cfg.ReaperInterval > 0 {
		StartReaper(ctx, pool, store, cfg.ReaperInterval, cfg.UnclaimedTTL, cfg.IncompleteTTL)
	}

	return store, nil
}

// NewPgxPool derives fireback's own database configuration (the same DSN
// its GORM layer connects to) and opens a fresh pgxpool against it, running
// migrations first. It calls fireback.LoadConfiguration itself rather than
// GetConfig, so it's self-sufficient whether it's called mid-bootstrap (from
// mountOnFirebackApp, where config is already loaded) or standalone (e.g.
// from cli.go's withPool, when Commands is invoked as a bare CLI with no
// app around it).
//
// Returns (nil, nil) - not an error - if fireback isn't configured for
// postgres: large objects are a postgres-only feature, so there's nothing
// to connect to.
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

// mountOnFirebackApp is the GinWebServerInitHooks callback StorageModuleSetup
// registers: it derives the same Postgres connection fireback's own GORM
// layer uses, opens a separate pgxpool for the large-object API (large
// objects need a native pgx transaction, which GORM's *sql.DB can't give
// us), defaults cfg.Authenticate to defaultAuthenticate if the caller didn't
// supply their own, and calls MountAll against it.
func mountOnFirebackApp(g *gin.RouterGroup, cfg *StorageModuleConfig) error {
	pool, err := NewPgxPool()
	if err != nil {
		return err
	}
	if pool == nil {
		// Large objects are a Postgres-only feature - nothing to mount
		// against any other database vendor.
		return nil
	}

	firebackCfg := StorageModuleConfig{}
	if cfg != nil {
		firebackCfg = *cfg
	}
	if firebackCfg.Authenticate == nil {
		firebackCfg.Authenticate = defaultAuthenticate
	}

	_, err = MountAll(context.Background(), g, pool, &firebackCfg)
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
		Capabilities: []fireback.PermissionInfo{},
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
