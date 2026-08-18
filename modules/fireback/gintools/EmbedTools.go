package gintools

import (
	"compress/gzip"
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"regexp"
	"sort"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// DefaultCacheableSuffixes are the file extensions treated as long-lived static
// assets by AssetCacheControlMiddleware when a PublicFolderInfo doesn't add its
// own via ExtraCacheableSuffixes.
var DefaultCacheableSuffixes = []string{
	".js", ".css", ".svg", ".png", ".jpg", ".jpeg", ".gif", ".webp",
	".woff", ".woff2", ".ttf", ".eot", ".ico",
}

// DefaultAssetCacheControl is the Cache-Control applied to static asset requests
// (see DefaultCacheableSuffixes) when a PublicFolderInfo doesn't set
// AssetCacheControl. A CDN or browser is told it can cache these aggressively -
// safe because a production JS/CSS build names its bundles with a content hash,
// so a new deploy is served under a new URL rather than overwriting a cached one.
const DefaultAssetCacheControl = "public, max-age=604800, immutable"

// DefaultIndexCacheControl is the Cache-Control applied to a folder's index
// document (both a literal "/index.html" request and the SPA fallback used for
// unknown client-side routes) when a PublicFolderInfo doesn't set
// IndexCacheControl. Kept to "no revalidation-free caching" by default so a
// redeploy is visible on the very next request - the document itself is what
// points at the (aggressively cached) hashed asset bundle, so it can't be
// cached the same way without risking a stale app being served indefinitely.
const DefaultIndexCacheControl = "no-cache"

/*
Public folders are used when you want to make an embed folder available through
the web server publicly. Quite useful for serving static content, and you can
prefix them.

Beyond just serving files, a PublicFolderInfo can opt into three things useful
for a CDN-fronted single-page app deployment - see gintools/README.md for a
worked example of all three:

  - DisableGzip: every folder registered here gets gzip compression for its own
    responses by default - never engine-wide. The decision is made only once
    this specific folder has confirmed *it* is about to write the response (a
    real embedded file existed, or this is its own SPA/index fallback), so it
    can't reach routes mounted outside PublicFolders (a module's API group, or
    the storage module's tus/streaming endpoints, whose protocols depend on
    exact byte counts) - not even when a folder's own Prefix is "/", under
    which every path is a string-prefix match. Set DisableGzip to opt this
    specific folder out too (e.g. content that's already compressed, or where
    compression cost isn't worth it for mostly-tiny files).

  - AssetCacheControl / ExtraCacheableSuffixes: override the Cache-Control
    applied to static asset requests (js/css/images/fonts/...) under this
    folder's Prefix, and/or add extra file extensions treated as cacheable
    assets. Left unset, DefaultAssetCacheControl applies to
    DefaultCacheableSuffixes only.

  - IndexCacheControl: override the Cache-Control applied to the index document
    itself. Defaults to DefaultIndexCacheControl.

  - InjectHTML: called for every request that falls through to this folder's
    SPA fallback (i.e. a client-side route with no matching file, such as
    /score/item1). It receives the gin context and the request path, and
    returns (HTMLInjection, true) to splice extra markup into index.html for
    that specific route - typically OpenGraph/meta tags so a social/crawler
    preview of a deep link has real content - or (HTMLInjection{}, false) to
    serve the plain index.html unchanged. This gives a single-page app a small,
    server-rendered slice of per-route <head> content without a real SSR
    pipeline.
*/
type PublicFolderInfo struct {
	Fs     *embed.FS
	Folder string
	Prefix string

	// DisableGzip opts this folder's own responses out of gzip compression
	// (see mountEmbedFolder). Gzip is on by default per folder - set this to
	// exclude a specific one, e.g. content that's already compressed.
	DisableGzip bool

	// AssetCacheControl overrides DefaultAssetCacheControl for static asset
	// requests under this folder's Prefix.
	AssetCacheControl string

	// ExtraCacheableSuffixes adds file extensions (in addition to
	// DefaultCacheableSuffixes) that should receive AssetCacheControl.
	ExtraCacheableSuffixes []string

	// IndexCacheControl overrides DefaultIndexCacheControl for this folder's
	// index document (index.html itself, and the SPA fallback).
	IndexCacheControl string

	// InjectHTML, if set, is consulted for every request that falls through to
	// this folder's SPA fallback. See the PublicFolderInfo doc comment above.
	InjectHTML func(c *gin.Context, path string) (HTMLInjection, bool)
}

// HTMLInjection is returned by PublicFolderInfo.InjectHTML to splice extra
// markup into a folder's index.html before it's served for a request that
// fell through to the SPA fallback. Both fields are raw HTML strings inserted
// verbatim - the caller owns escaping of any request-derived value (e.g. use
// html.EscapeString or html/template on a DB-sourced title before putting it
// in a <meta> tag).
type HTMLInjection struct {
	// Head is inserted immediately before </head> - the natural home for
	// <title>, <meta property="og:*">, <meta name="description">, a canonical
	// <link>, a JSON-LD <script type="application/ld+json">, etc.
	Head string

	// Body is inserted immediately before </body> - useful for bootstrapping
	// data into the page (e.g. <script>window.__PRELOADED__={...}</script>) so
	// the SPA can hydrate without a second fetch.
	Body string
}

var (
	headCloseTag  = regexp.MustCompile(`(?i)</head>`)
	bodyCloseTag  = regexp.MustCompile(`(?i)</body>`)
	existingTitle = regexp.MustCompile(`(?is)<title[^>]*>.*?</title>\s*`)
	injectedTitle = regexp.MustCompile(`(?is)<title[^>]*>`)
)

// injectIntoHTML splices headExtra/bodyExtra into doc just before the closing
// </head>/</body> tags (case-insensitively). If a closing tag isn't found, the
// content is prepended/appended instead, so a malformed template never
// silently drops the injection.
//
// When headExtra itself defines a <title>, the document's own pre-existing
// <title> (the app shell's static/default one) is stripped first rather than
// left in place alongside the new one. Per the HTML spec, document.title (and
// the value a browser shows on the tab) is taken from the *first* <title> in
// tree order - simply appending a second one before </head> would leave it
// permanently shadowed by the original, so the page title would never
// actually reflect the injected route-specific one.
func injectIntoHTML(doc string, headExtra string, bodyExtra string) string {
	if headExtra != "" {
		if injectedTitle.MatchString(headExtra) {
			doc = existingTitle.ReplaceAllString(doc, "")
		}
		if loc := headCloseTag.FindStringIndex(doc); loc != nil {
			doc = doc[:loc[0]] + headExtra + "\n" + doc[loc[0]:]
		} else {
			doc = headExtra + doc
		}
	}
	if bodyExtra != "" {
		if loc := bodyCloseTag.FindStringIndex(doc); loc != nil {
			doc = doc[:loc[0]] + bodyExtra + "\n" + doc[loc[0]:]
		} else {
			doc = doc + bodyExtra
		}
	}
	return doc
}

type embedFileSystem struct {
	http.FileSystem
	indexes bool
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	// Folders mounted below root need the mount prefix stripped before
	// looking the file up inside their own embedded sub-filesystem.
	if prefix != "/" {
		p := strings.TrimPrefix(path, prefix)
		if len(p) == len(path) {
			return false
		}
		path = p
		if path == "" {
			path = "/"
		}
	}

	f, err := e.Open(path)
	if err != nil {
		return false
	}

	// check if indexing is allowed
	s, _ := f.Stat()
	if s.IsDir() && !e.indexes {
		return false
	}

	return true
}

func EmbedFolder(fsEmbed embed.FS, targetPath string, index bool) static.ServeFileSystem {
	subFS, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(subFS),
		indexes:    index,
	}
}

// readIndexHTML reads "index.html" out of a PublicFolderInfo's embedded
// sub-filesystem. It's read once at mount time and cached (embed.FS content is
// baked into the binary, so it never changes at runtime) - InjectHTML then
// only pays for the cheap per-request string splice, not another FS read.
func readIndexHTML(item PublicFolderInfo) ([]byte, error) {
	subFS, err := fs.Sub(*item.Fs, item.Folder)
	if err != nil {
		return nil, err
	}
	return fs.ReadFile(subFS, "index.html")
}

// gzipResponseWriter mirrors gin-contrib/gzip's own (unexported) wrapper: it
// deletes any Content-Length set by the handler it wraps (which would
// otherwise describe the *uncompressed* size) before writing compressed bytes
// through - net/http then falls back to chunked transfer-encoding, which is
// what a gzip response needs since the final compressed size isn't known
// until the stream is closed.
type gzipResponseWriter struct {
	gin.ResponseWriter
	writer *gzip.Writer
}

func (g *gzipResponseWriter) WriteString(s string) (int, error) {
	g.Header().Del("Content-Length")
	return g.writer.Write([]byte(s))
}

func (g *gzipResponseWriter) Write(data []byte) (int, error) {
	g.Header().Del("Content-Length")
	return g.writer.Write(data)
}

func (g *gzipResponseWriter) WriteHeader(code int) {
	g.Header().Del("Content-Length")
	g.ResponseWriter.WriteHeader(code)
}

// acceptsGzip reports whether req negotiated gzip and isn't a byte-range
// request. Range requests (used for resumable downloads and media seeking,
// e.g. an HTML5 <video> tag) return a slice of the underlying resource picked
// after the byte length is known; gzip-encoding just that slice produces a
// self-contained stream whose Content-Range still describes offsets in the
// *original* file, a combination browsers' media/range handling doesn't
// tolerate (Chrome fails such responses with ERR_CONTENT_DECODING_FAILED) -
// so skip compression whenever a Range header is present, the same way
// nginx/Apache do.
func acceptsGzip(req *http.Request) bool {
	return req.Header.Get("Range") == "" &&
		strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") &&
		!strings.Contains(req.Header.Get("Connection"), "Upgrade")
}

// serveGzipped runs fn with c.Writer swapped for a gzip-compressing wrapper
// for its duration, then restores the original writer. Deliberately *not* a
// standalone middleware some route might happen to sit behind: it's called
// directly, inline, only once the caller has already confirmed it is itself
// about to write the response (see mountEmbedFolder below) - so it can never
// end up compressing a response some other route (an API handler, the
// storage module's tus endpoints, ...) produces instead.
func serveGzipped(c *gin.Context, fn func()) {
	original := c.Writer
	gz := gzip.NewWriter(original)
	c.Header("Content-Encoding", "gzip")
	c.Header("Vary", "Accept-Encoding")
	c.Writer = &gzipResponseWriter{ResponseWriter: original, writer: gz}
	defer func() {
		gz.Close()
		c.Writer = original
	}()
	fn()
}

// mountEmbedFolder mounts a public folder as static-file middleware (rather
// than a route) so that requests for paths it doesn't have a file for simply
// fall through instead of terminating in a raw 404 - letting the combined
// NoRoute handler in EmbedFoldersForGin decide on a SPA fallback. It returns
// that fallback handler, which re-serves the folder's index.html for
// SPA-style client-side routing (e.g. a React Router deep link on refresh),
// optionally passed through item.InjectHTML first (see PublicFolderInfo).
//
// Gzip (unless DisableGzip) is applied to both paths this function serves -
// but only once each has independently confirmed it's the one producing the
// response: the asset middleware checks fileSystem.Exists itself before
// wrapping, and the SPA fallback below is only ever invoked by NoRoute for a
// path nothing else matched. Neither depends on this folder's Prefix, so a
// folder mounted at "/" (matching every path as a string prefix) still can't
// leak compression onto some other route - see the longer explanation on
// SetupHttpServer in FirebackApp.go.
func mountEmbedFolder(item PublicFolderInfo, r *gin.Engine) func(c *gin.Context) {
	prefix := item.Prefix
	if prefix == "" {
		prefix = "/"
	}

	fileSystem := EmbedFolder(*item.Fs, item.Folder, true)
	staticServer := static.Serve(prefix, fileSystem)

	r.Use(func(c *gin.Context) {
		if !item.DisableGzip &&
			c.Request.Method == http.MethodGet &&
			acceptsGzip(c.Request) &&
			fileSystem.Exists(prefix, c.Request.URL.Path) {
			serveGzipped(c, func() { staticServer(c) })
			return
		}
		staticServer(c)
	})

	indexCacheControl := item.IndexCacheControl
	if indexCacheControl == "" {
		indexCacheControl = DefaultIndexCacheControl
	}

	// Best-effort: a folder without an index.html (unusual for a SPA mount)
	// just never gets injection - it still falls back to staticServer below.
	indexHTML, indexErr := readIndexHTML(item)

	return func(c *gin.Context) {
		serve := func() {
			if item.InjectHTML != nil && indexErr == nil {
				if injection, ok := item.InjectHTML(c, c.Request.URL.Path); ok {
					html := injectIntoHTML(string(indexHTML), injection.Head, injection.Body)
					c.Header("Cache-Control", indexCacheControl)
					c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
					return
				}
			}

			c.Header("Cache-Control", indexCacheControl)
			c.Request.URL.Path = prefix
			staticServer(c)
		}

		if !item.DisableGzip && c.Request.Method == http.MethodGet && acceptsGzip(c.Request) {
			serveGzipped(c, serve)
			return
		}
		serve()
	}
}

// EmbedFolderForGin mounts a single embedded folder and wires up its own
// SPA fallback via NoRoute. Kept for backward compatibility with callers
// mounting exactly one public folder; when mounting several folders (as
// FirebackApp.PublicFolders does), prefer EmbedFoldersForGin so their
// fallbacks don't clobber each other - gin.Engine.NoRoute keeps only the
// handlers from the most recent call, so calling it once per folder means
// only the last-registered folder's fallback is ever reachable.
func EmbedFolderForGin(ui *embed.FS, folder string, r *gin.Engine, prefix string) {
	EmbedFoldersForGin([]PublicFolderInfo{{Fs: ui, Folder: folder, Prefix: prefix}}, r)
}

// EmbedFoldersForGin mounts every given public folder and registers a single
// combined NoRoute fallback, so a browser refresh on a deep client-side route
// (e.g. /scores/abc123) still resolves to the right SPA's index.html instead
// of 404ing. The most specific (longest) matching prefix wins; folders
// mounted at "/" act as the default fallback when no other prefix matches.
//
// It also installs AssetCacheControlMiddleware for the given folders. Gzip
// compression is handled per-folder inside mountEmbedFolder itself, not here.
func EmbedFoldersForGin(items []PublicFolderInfo, r *gin.Engine) {
	type fallback struct {
		prefix string
		handle func(c *gin.Context)
	}

	fallbacks := make([]fallback, 0, len(items))
	for _, item := range items {
		prefix := item.Prefix
		if prefix == "" {
			prefix = "/"
		}
		fallbacks = append(fallbacks, fallback{
			prefix: prefix,
			handle: mountEmbedFolder(item, r),
		})
	}

	r.Use(AssetCacheControlMiddleware(items))

	r.NoRoute(func(c *gin.Context) {
		if c.Request.Method != http.MethodGet ||
			strings.ContainsRune(c.Request.URL.Path, '.') ||
			strings.HasPrefix(c.Request.URL.Path, "/api/") {
			return
		}

		path := c.Request.URL.Path
		best := -1
		for i, f := range fallbacks {
			if f.prefix == "/" {
				continue
			}
			if strings.HasPrefix(path, f.prefix) && (best == -1 || len(f.prefix) > len(fallbacks[best].prefix)) {
				best = i
			}
		}
		if best != -1 {
			fallbacks[best].handle(c)
			return
		}

		for _, f := range fallbacks {
			if f.prefix == "/" {
				f.handle(c)
				return
			}
		}
	})
}

// AssetCacheControlMiddleware sets a Cache-Control header on GET requests for
// static asset file types (see DefaultCacheableSuffixes), so a CDN or browser
// in front of the app (Cloudflare, etc) knows it's safe to cache hashed
// js/css/font/image bundles aggressively - while leaving index.html and API
// responses alone (those get their own Cache-Control from mountEmbedFolder,
// or none at all).
//
// Folders in items may set AssetCacheControl and ExtraCacheableSuffixes to
// override the default for requests under their Prefix; the longest matching
// prefix wins. This is called for you by EmbedFoldersForGin/EmbedFolderForGin.
func AssetCacheControlMiddleware(items []PublicFolderInfo) gin.HandlerFunc {
	type rule struct {
		prefix       string
		cacheControl string
		suffixes     []string
		fs           static.ServeFileSystem
	}

	rules := make([]rule, 0, len(items)+1)
	for _, item := range items {
		prefix := item.Prefix
		if prefix == "" {
			prefix = "/"
		}
		cacheControl := item.AssetCacheControl
		if cacheControl == "" {
			cacheControl = DefaultAssetCacheControl
		}
		suffixes := DefaultCacheableSuffixes
		if len(item.ExtraCacheableSuffixes) > 0 {
			suffixes = append(append([]string{}, DefaultCacheableSuffixes...), item.ExtraCacheableSuffixes...)
		}
		rules = append(rules, rule{
			prefix:       prefix,
			cacheControl: cacheControl,
			suffixes:     suffixes,
			// Its own embedFileSystem, independent of the one mountEmbedFolder built for
			// the same item - cheap (just wraps the already-in-binary embed.FS again),
			// and lets this middleware check existence itself instead of trusting the
			// suffix match alone (see the Exists call below).
			fs: EmbedFolder(*item.Fs, item.Folder, true),
		})
	}

	// Longest (most specific) prefix first, so e.g. "/manage" isn't shadowed by
	// a "/" folder that comes earlier in items.
	sort.Slice(rules, func(i, j int) bool { return len(rules[i].prefix) > len(rules[j].prefix) })

	return func(c *gin.Context) {
		if c.Request.Method == http.MethodGet {
			path := c.Request.URL.Path
			for _, rule := range rules {
				if rule.prefix != "/" && !strings.HasPrefix(path, rule.prefix) {
					continue
				}
				if !hasSuffix(path, rule.suffixes) {
					continue
				}
				// A suffix match alone used to be enough to stamp the aggressive
				// "immutable, max-age=604800" header - including onto a 404 for an
				// asset that doesn't actually exist (a typo'd path, a renamed/removed
				// file, a route hit before a rebuild landed it). A browser that ever
				// received that 404 then refuses to ask again for a week, even after
				// the file shows up - indistinguishable from "it's just missing"
				// without inspecting response headers. Requiring Exists here means the
				// header only ever rides along with a response this folder is actually
				// about to serve successfully.
				if !rule.fs.Exists(rule.prefix, path) {
					continue
				}
				c.Header("Cache-Control", rule.cacheControl)
				break
			}
		}
		c.Next()
	}
}

func hasSuffix(path string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}
	return false
}

func LoadHTMLFromEmbedFS(engine *gin.Engine, embedFS embed.FS, pattern string) {
	root := template.New("")
	tmpl := template.Must(root, LoadAndAddToRoot(engine.FuncMap, root, embedFS, pattern))
	engine.SetHTMLTemplate(tmpl)
}

func LoadAndAddToRoot(funcMap template.FuncMap, rootTemplate *template.Template, embedFS embed.FS, pattern string) error {
	pattern = strings.ReplaceAll(pattern, ".", "\\.")
	pattern = strings.ReplaceAll(pattern, "*", ".*")

	err := fs.WalkDir(embedFS, ".", func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		if matched, _ := regexp.MatchString(pattern, path); !d.IsDir() && matched {
			data, readErr := embedFS.ReadFile(path)
			if readErr != nil {
				return readErr
			}
			t := rootTemplate.New(path).Funcs(funcMap)
			if _, parseErr := t.Parse(string(data)); parseErr != nil {
				return parseErr
			}
		}
		return nil
	})
	return err
}
