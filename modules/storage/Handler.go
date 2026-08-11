package storage

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tus/tusd/pkg/handler"
)

// NewTusHandler builds the tus protocol http.Handler backed by store.
// basePath must match wherever the handler ends up mounted (e.g. "/files").
// cfg's Quota is enforced via PreUploadCreateCallback, using the owner
// identity Mount's serve stashes into the creation request's
// Upload-Metadata header.
func NewTusHandler(store Store, basePath string, cfg *StorageModuleConfig) (*handler.Handler, error) {
	composer := handler.NewStoreComposer()
	composer.UseCore(store)
	composer.UseTerminater(store)

	return handler.NewHandler(handler.Config{
		BasePath:                basePath,
		StoreComposer:           composer,
		PreUploadCreateCallback: enforceQuotaCallback(store, cfg),
	})
}

// enforceQuotaCallback rejects a new upload before any bytes are written if
// it would push its owner over their quota. It reads the owner identity
// Mount's serve already stashed into the Upload-Metadata header (under the
// reserved metaKeyUserId/... keys, stripped back out again by
// Store.NewUpload) rather than re-resolving auth itself - by the time this
// runs, the request has already passed Mount's own auth/ownership gate.
func enforceQuotaCallback(store Store, cfg *StorageModuleConfig) func(handler.HookEvent) error {
	return func(hook handler.HookEvent) error {
		userId := hook.Upload.MetaData[metaKeyUserId]
		if userId == "" {
			// Anonymous upload (Authenticate is nil, or resolved no
			// UserId) - nothing to meter against.
			return nil
		}

		auth := AuthContext{
			UserId:      userId,
			WorkspaceId: hook.Upload.MetaData[metaKeyWorkspaceId],
			AccessLevel: hook.Upload.MetaData[metaKeyAccessLevel],
		}

		quota := cfg.quotaFor(auth)
		if quota < 0 {
			return nil // negative means unlimited for this identity
		}

		used, err := UsedBytes(context.Background(), store, userId)
		if err != nil {
			return err
		}
		if used+hook.Upload.Size > quota {
			return handler.NewHTTPError(ErrQuotaExceeded, http.StatusForbidden)
		}
		return nil
	}
}

// injectOwnerMetadata rewrites a creation request's Upload-Metadata header
// to record auth's identity under reserved keys, overwriting anything the
// client itself sent under the same names so it can never be spoofed.
func injectOwnerMetadata(existingHeader string, auth AuthContext) string {
	meta := handler.ParseMetadataHeader(existingHeader)
	if meta == nil {
		meta = map[string]string{}
	}
	delete(meta, metaKeyUserId)
	delete(meta, metaKeyWorkspaceId)
	delete(meta, metaKeyAccessLevel)

	if auth.UserId != "" {
		meta[metaKeyUserId] = auth.UserId
	}
	if auth.WorkspaceId != "" {
		meta[metaKeyWorkspaceId] = auth.WorkspaceId
	}
	if auth.AccessLevel != "" {
		meta[metaKeyAccessLevel] = auth.AccessLevel
	}

	return handler.SerializeMetadataHeader(meta)
}

// Mount wires the tus upload endpoints (POST/HEAD/PATCH/GET/DELETE) onto a
// gin router at basePath, e.g. Mount(engine, "/files", store, cfg).
//
// Every request is first run through cfg.Authenticate (see
// StorageModuleConfig): a nil Authenticate disables authentication
// entirely (every request is Anonymous, ownership is never checked, and
// Quota is never consulted - this module's original behavior). When
// Authenticate is set, requests against an existing upload (anything but a
// bare POST to basePath) are rejected with 403 unless the resolved identity
// matches the upload's recorded owner (or the upload has no recorded owner
// at all, e.g. it predates this mechanism). A creation request
// (POST to basePath) has the caller's identity stashed into the
// Upload-Metadata header for Store.NewUpload/enforceQuotaCallback to pick
// up.
//
// tusd's routed Handler uses bmizerany/pat internally, whose patterns ("" for
// creation, ":id" for everything else) are matched against the path with the
// mount prefix already removed and, for ":id", without a leading slash
// either. gin.WrapH alone doesn't do that rewriting, so this rewrites
// req.URL.Path itself before handing the request to tusd.
func Mount(r gin.IRouter, basePath string, store Store, cfg *StorageModuleConfig) error {
	cfg = cfg.withDefaults()

	h, err := NewTusHandler(store, basePath, cfg)
	if err != nil {
		return err
	}

	trimmed := strings.TrimSuffix(basePath, "/")

	serve := func(c *gin.Context) {
		ctx := c.Request.Context()

		// resolveAuth/OwnerOf rejections below respond directly via c.JSON
		// instead of reaching h.ServeHTTP, so tusd's own CORS-handling
		// Middleware never runs for them. Set this unconditionally, up
		// front, so a 401/403 doesn't additionally look like a CORS
		// failure in the browser. h.ServeHTTP (success path) sets the same
		// header via Set(), not Add(), so this doesn't produce a duplicate.
		if origin := c.GetHeader("Origin"); origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
		}

		auth, err := resolveAuth(c, cfg)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		id := strings.TrimPrefix(strings.TrimPrefix(c.Request.URL.Path, trimmed), "/")

		if cfg.Authenticate != nil && id != "" {
			ownerId, found, err := store.OwnerOf(ctx, id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			if found && !authorizeOwner(auth, ownerId) {
				c.JSON(http.StatusForbidden, gin.H{"error": "not your upload"})
				return
			}
		}

		req := c.Request.Clone(ctx)
		req.URL.Path = id
		req.URL.RawPath = ""

		if id == "" && c.Request.Method == http.MethodPost && cfg.Authenticate != nil && !auth.isAnonymous() {
			req.Header.Set("Upload-Metadata", injectOwnerMetadata(req.Header.Get("Upload-Metadata"), auth))
		}

		h.ServeHTTP(c.Writer, req)
	}

	r.Any(trimmed, serve)
	r.Any(trimmed+"/*id", serve)

	return nil
}
