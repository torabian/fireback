// DumpHttp.go mounts the HTTP half of `backup dump --hash`: registering a
// job (POST) and streaming it (GET .../raw) - see README.md's "Streaming a
// dump over HTTP" section. Neither endpoint requires its own auth: POST's
// access control is jobStoreDir()'s OS-level per-user privacy (see
// HashRegistry.go), and GET's is the hash itself being the sole credential,
// by design. Wired in via ModuleSetup's GinWebServerInitHooks, the same
// convention every other module's HTTP routes use (see eventbus/storage).
package backup

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type createDumpRequest struct {
	Database string `json:"database"`
}

type createDumpResponse struct {
	Hash      string    `json:"hash"`
	Url       string    `json:"url"`
	Database  string    `json:"database"`
	ExpiresAt time.Time `json:"expiresAt"`
}

// MountDumpHttp wires POST {prefix}/backup/dumps and
// GET {prefix}/backup/dumps/:hash/raw onto g.
func MountDumpHttp(g gin.IRouter) {
	g.POST("/backup/dumps", createDumpHandler)
	g.GET("/backup/dumps/:hash/raw", fetchDumpHandler)
}

func createDumpHandler(c *gin.Context) {
	cfg, err := LoadDumpConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var body createDumpRequest
	_ = c.ShouldBindJSON(&body) // an empty/absent body is valid - it just means "let me pick"

	ctx := c.Request.Context()

	if body.Database == "" {
		names, err := ListDatabases(ctx, cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if len(names) == 1 {
			body.Database = names[0]
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":     "database is required: more than one database is visible on this connection",
				"databases": names,
			})
			return
		}
	} else if cfg.Vendor != VendorSqlite {
		// Validate the requested name is actually one this connection can
		// see, so a typo fails fast here with a clear message instead of
		// surfacing as a confusing pg_dump/mysqldump error later on GET.
		names, err := ListDatabases(ctx, cfg)
		if err == nil && !contains(names, body.Database) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":     fmt.Sprintf("database %q not found on this connection", body.Database),
				"databases": names,
			})
			return
		}
	}

	ttl := time.Duration(cfg.HashTTLSeconds) * time.Second
	if ttl <= 0 {
		ttl = 30 * time.Minute
	}
	hash, err := registerDumpJob(body.Database, ttl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createDumpResponse{
		Hash:      hash,
		Url:       requestOrigin(c) + c.Request.URL.Path + "/" + hash + "/raw",
		Database:  body.Database,
		ExpiresAt: time.Now().Add(ttl),
	})
}

func fetchDumpHandler(c *gin.Context) {
	job, ok := claimDumpJob(c.Param("hash"))
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "unknown, already-claimed, or expired hash"})
		return
	}

	cfg, err := LoadDumpConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s.zip"`, job.Database))

	if err := DumpDatabase(c.Request.Context(), cfg, job.Database, c.Writer); err != nil {
		// Headers/partial body may already be on the wire - there's no
		// clean way to turn this into a JSON error response at this point,
		// only to stop writing and let the client see a truncated
		// download, which is at least detectable (short Content-Length-
		// less response) rather than silently "succeeding".
		c.Error(err)
	}
}

// requestOrigin best-effort reconstructs scheme://host from the incoming
// request, so createDumpResponse.Url is a fully-qualified link a caller on a
// different machine can curl directly, not just a path. Honors
// X-Forwarded-Proto since this commonly sits behind a reverse proxy
// terminating TLS.
func requestOrigin(c *gin.Context) string {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	if proto := c.GetHeader("X-Forwarded-Proto"); proto != "" {
		scheme = proto
	}
	return scheme + "://" + c.Request.Host
}

func contains(items []string, item string) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}
