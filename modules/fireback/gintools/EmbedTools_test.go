package gintools

import (
	"bytes"
	"compress/gzip"
	"embed"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

//go:embed testdata/site
var testSiteFS embed.FS

//go:embed testdata/manage-site
var testManageSiteFS embed.FS

func TestInjectIntoHTML(t *testing.T) {
	doc := "<html><head><title>App</title></head><body><div id=\"root\"></div></body></html>"

	got := injectIntoHTML(doc, `<meta property="og:title" content="Item">`, `<script>window.x=1</script>`)

	wantHead := `<head><title>App</title><meta property="og:title" content="Item">
</head>`
	wantBody := `<script>window.x=1</script>
</body>`

	if !strings.Contains(got, wantHead) {
		t.Fatalf("expected head injection before </head>, got: %s", got)
	}
	if !strings.Contains(got, wantBody) {
		t.Fatalf("expected body injection before </body>, got: %s", got)
	}
}

// TestInjectIntoHTML_ReplacesExistingTitle is the regression test for a real
// bug caught while building the nima score-detail OG tags feature: appending
// a second <title> before </head> (alongside the app shell's own static one)
// looks fine by eye, but per the HTML spec document.title - and what the
// browser tab shows - comes from the *first* <title> in the document, so the
// injected one was silently never used.
func TestInjectIntoHTML_ReplacesExistingTitle(t *testing.T) {
	doc := `<html><head><title>App Shell</title><meta charset="utf-8"></head><body></body></html>`

	got := injectIntoHTML(doc, `<title>Item One</title><meta property="og:title" content="Item One">`, "")

	if strings.Contains(got, "App Shell") {
		t.Fatalf("expected the original <title> to be removed, got: %s", got)
	}
	if strings.Count(got, "<title>") != 1 {
		t.Fatalf("expected exactly one <title> tag, got: %s", got)
	}
	if !strings.Contains(got, "<title>Item One</title>") {
		t.Fatalf("expected the injected title to be present, got: %s", got)
	}
	// Everything else already in <head> must survive untouched.
	if !strings.Contains(got, `<meta charset="utf-8">`) {
		t.Fatalf("expected unrelated existing head content to survive, got: %s", got)
	}
}

// TestInjectIntoHTML_NoTitleInjection_KeepsExistingTitle covers the other
// half: InjectHTML implementations that only add meta tags (no <title> of
// their own) must leave whatever title the document already had alone.
func TestInjectIntoHTML_NoTitleInjection_KeepsExistingTitle(t *testing.T) {
	doc := `<html><head><title>App Shell</title></head><body></body></html>`

	got := injectIntoHTML(doc, `<meta property="og:title" content="Item One">`, "")

	if !strings.Contains(got, "<title>App Shell</title>") {
		t.Fatalf("expected the original title to survive when nothing replaces it, got: %s", got)
	}
}

func TestInjectIntoHTML_NoOp(t *testing.T) {
	doc := "<html><head></head><body></body></html>"
	if got := injectIntoHTML(doc, "", ""); got != doc {
		t.Fatalf("expected doc unchanged when both injections are empty, got: %s", got)
	}
}

// TestEmbedFoldersForGin_GzipAndInjectHTML exercises the full, real wiring
// (EmbedFoldersForGin -> mountEmbedFolder -> NoRoute) rather than any single
// helper in isolation, since the bug worth guarding against only shows up in
// how the pieces compose: a "/" folder's asset middleware sits in front of
// every route registered after it (see mountEmbedFolder's doc comment), so
// it's not enough for gzip to work on the embedded SPA - it has to *also*
// leave a perfectly ordinary route (like a module's own API handler, or
// storage's tus endpoints in the real app) completely untouched.
func TestEmbedFoldersForGin_GzipAndInjectHTML(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// InjectHTML deciding, per path, whether this folder should handle the
	// request at all and - if so - what to splice into index.html for it.
	// Each case injects its own <title> too (the way nima's real
	// InjectScoreOgTags does), not just a <meta> tag, so this exercises the
	// same replace-the-app-shell's-existing-title path end to end - not just
	// injectIntoHTML in isolation - against testdata/site/index.html's real
	// shape (its own pre-existing <title>, script/link tags, and an inline
	// <script> after </body>, all mirroring the real nima app shell).
	itemInjectHTML := func(c *gin.Context, path string) (HTMLInjection, bool) {
		switch path {
		case "/item1":
			return HTMLInjection{Head: `<title>Item One</title><meta name="item" content="item1">`}, true
		case "/item2":
			return HTMLInjection{Head: `<title>Item Two</title><meta name="item" content="item2">`}, true
		default:
			return HTMLInjection{}, false
		}
	}

	r := gin.New()
	EmbedFoldersForGin([]PublicFolderInfo{
		{Fs: &testSiteFS, Folder: "testdata/site", Prefix: "/", InjectHTML: itemInjectHTML},
	}, r)

	// A "standard" endpoint, registered the same way a real module's API
	// route is in the actual app: *after* EmbedFoldersForGin, so it inherits
	// whatever the "/" folder's own middleware left sitting in the chain.
	helloBody := strings.Repeat("hello world, plain api response. ", 50)
	r.GET("/api/hello", func(c *gin.Context) {
		c.String(http.StatusOK, helloBody)
	})

	t.Run("a standard endpoint is never gzipped", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/hello", nil)
		req.Header.Set("Accept-Encoding", "gzip")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if enc := w.Header().Get("Content-Encoding"); enc != "" {
			t.Fatalf("expected /api/hello to never be gzipped (a folder mounted at \"/\" "+
				"string-prefixes every path), got Content-Encoding: %q", enc)
		}
		if w.Body.String() != helloBody {
			t.Fatalf("expected the plain, uncompressed body, got: %q", w.Body.String())
		}
	})

	t.Run("the SPA fallback is gzipped and injected per route", func(t *testing.T) {
		cases := []struct {
			path         string
			wantTitle    string
			wantInjected string
			wantAbsent   string
		}{
			{"/item1", "Item One", `<meta name="item" content="item1">`, "item2"},
			{"/item2", "Item Two", `<meta name="item" content="item2">`, "item1"},
		}

		for _, tc := range cases {
			req := httptest.NewRequest(http.MethodGet, tc.path, nil)
			req.Header.Set("Accept-Encoding", "gzip")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if enc := w.Header().Get("Content-Encoding"); enc != "gzip" {
				t.Fatalf("%s: expected a gzip Content-Encoding, got %q", tc.path, enc)
			}

			// Regression guard: this is a *custom InjectHTML route*, not the
			// plain SPA fallback - it must still get the same
			// never-cache-me Cache-Control as index.html itself, so a CDN
			// never keeps serving a stale injected page (or a stale build)
			// after a redeploy.
			if got := w.Header().Get("Cache-Control"); got != DefaultIndexCacheControl {
				t.Fatalf("%s: expected Cache-Control %q on an injected route, got %q", tc.path, DefaultIndexCacheControl, got)
			}

			body := mustGunzip(t, w.Body.Bytes())
			if !strings.Contains(body, tc.wantInjected) {
				t.Fatalf("%s: expected injected marker %q in the decompressed body, got: %s", tc.path, tc.wantInjected, body)
			}
			if strings.Contains(body, tc.wantAbsent) {
				t.Fatalf("%s: leaked the other route's injected content: %s", tc.path, body)
			}

			// The fixture's own <title>Test App</title> must be gone,
			// replaced by exactly this route's injected one - not left
			// alongside it (see TestInjectIntoHTML_ReplacesExistingTitle;
			// this is the same guarantee, now checked through the real
			// mountEmbedFolder/NoRoute path against a realistic document).
			if strings.Contains(body, "Test App") {
				t.Fatalf("%s: the fixture's original <title> leaked through, got: %s", tc.path, body)
			}
			if got := strings.Count(body, "<title>"); got != 1 {
				t.Fatalf("%s: expected exactly one <title> tag, found %d: %s", tc.path, got, body)
			}
			if !strings.Contains(body, "<title>"+tc.wantTitle+"</title>") {
				t.Fatalf("%s: expected <title>%s</title>, got: %s", tc.path, tc.wantTitle, body)
			}

			// Everything the "react app" needs to actually boot - its script
			// bundle, stylesheet, root div, and even the app shell's own
			// trailing inline bootstrap script - must survive the injection
			// untouched, the same thing the nima e2e spec checks against the
			// real build (see score-og-tags.cy.js).
			if !strings.Contains(body, `<script type="module" crossorigin src="/assets/index-DQGZjcUQ.js"></script>`) {
				t.Fatalf("%s: expected the app's JS bundle tag to survive, got: %s", tc.path, body)
			}
			if !strings.Contains(body, `<link rel="stylesheet" crossorigin href="/assets/index-76XdJD4x.css">`) {
				t.Fatalf("%s: expected the app's stylesheet tag to survive, got: %s", tc.path, body)
			}
			if !strings.Contains(body, `<div id="root"></div>`) {
				t.Fatalf("%s: expected the SPA's mount point to survive, got: %s", tc.path, body)
			}
			if !strings.Contains(body, `console.log("app shell inline bootstrap");`) {
				t.Fatalf("%s: expected the app shell's own trailing inline script to survive, got: %s", tc.path, body)
			}
		}
	})

	t.Run("a plain SPA fallback (no InjectHTML match) and a literal index.html request both get the never-cache header", func(t *testing.T) {
		for _, path := range []string{"/some/unmatched/route", "/index.html"} {
			req := httptest.NewRequest(http.MethodGet, path, nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if got := w.Header().Get("Cache-Control"); got != DefaultIndexCacheControl {
				t.Fatalf("%s: expected Cache-Control %q, got %q", path, DefaultIndexCacheControl, got)
			}
		}
	})

	t.Run("without a gzip-accepting client, the same route is served plain", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/item1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if enc := w.Header().Get("Content-Encoding"); enc != "" {
			t.Fatalf("expected no compression without an Accept-Encoding request header, got %q", enc)
		}
		if !strings.Contains(w.Body.String(), `content="item1"`) {
			t.Fatalf("expected injection to still apply even when not gzipped, got: %s", w.Body.String())
		}
	})
}

func mustGunzip(t *testing.T, data []byte) string {
	t.Helper()

	zr, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		t.Fatalf("response was labeled gzip but failed to decode: %v", err)
	}
	defer zr.Close()

	out, err := io.ReadAll(zr)
	if err != nil {
		t.Fatalf("failed reading decompressed body: %v", err)
	}
	return string(out)
}

func TestAssetCacheControlMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	items := []PublicFolderInfo{
		{Fs: &testSiteFS, Folder: "testdata/site", Prefix: "/"},
		{Fs: &testManageSiteFS, Folder: "testdata/manage-site", Prefix: "/manage", AssetCacheControl: "public, max-age=60"},
	}

	r := gin.New()
	r.Use(AssetCacheControlMiddleware(items))
	r.GET("/*any", func(c *gin.Context) { c.String(200, "ok") })

	cases := []struct {
		path string
		want string
	}{
		{"/assets/main.js", DefaultAssetCacheControl},
		{"/manage/assets/main.js", "public, max-age=60"},
		{"/index.html", ""},
		// Regression guard: a suffix match alone used to be enough to stamp the
		// aggressive, immutable Cache-Control header - including onto a 404 for a
		// path that doesn't correspond to a real file (a typo, a renamed/removed
		// asset, a route hit before a rebuild landed it). A browser that ever saw
		// that 404 would then refuse to ask again for a week even after the file
		// showed up. The header must only ride along with a response this folder
		// is actually about to serve.
		{"/assets/does-not-exist.js", ""},
	}

	for _, tc := range cases {
		req := httptest.NewRequest("GET", tc.path, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if got := w.Header().Get("Cache-Control"); got != tc.want {
			t.Fatalf("path %s: expected Cache-Control %q, got %q", tc.path, tc.want, got)
		}
	}
}

// TestEmbedFoldersForGin_CacheControl is the real end-to-end regression test for
// the bug TestAssetCacheControlMiddleware's own isolated setup couldn't catch:
// through the actual EmbedFoldersForGin/mountEmbedFolder wiring, gin-contrib/
// static's Serve aborts the middleware chain the instant it writes a response
// for a request fileSystem.Exists confirmed - which used to silently skip
// AssetCacheControlMiddleware (registered afterwards) for every real,
// successfully-served file. Confirmed via a throwaway probe before the fix: a
// request for an existing asset came back with an empty Cache-Control header
// end to end, despite the isolated unit test above passing. This test exercises
// the real path both folders (EmbedFoldersForGin, not a bare
// AssetCacheControlMiddleware) actually run through in production.
func TestEmbedFoldersForGin_CacheControl(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	EmbedFoldersForGin([]PublicFolderInfo{
		{Fs: &testSiteFS, Folder: "testdata/site", Prefix: "/"},
	}, r)

	cases := []struct {
		name string
		path string
		want string
	}{
		{"a real cacheable asset", "/assets/main.js", DefaultAssetCacheControl},
		{"the bare root path (net/http resolves this to index.html)", "/", DefaultIndexCacheControl},
		{"a literal /index.html request", "/index.html", DefaultIndexCacheControl},
		{"an unmatched path falling through to the SPA fallback", "/some/deep/route", DefaultIndexCacheControl},
	}

	for _, tc := range cases {
		req := httptest.NewRequest(http.MethodGet, tc.path, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if got := w.Header().Get("Cache-Control"); got != tc.want {
			t.Fatalf("%s (%s): expected Cache-Control %q, got %q", tc.name, tc.path, tc.want, got)
		}
	}
}

// TestDefaultCacheableSuffixes_OnlyJSCSSAndImages locks in the "cache js/css/
// images, nothing else by default" contract at the value level, independent of
// any fixture files - so it can't accidentally pass just because a testdata
// folder happens not to contain a font/icon file to exercise.
func TestDefaultCacheableSuffixes_OnlyJSCSSAndImages(t *testing.T) {
	allowed := map[string]bool{
		".js": true, ".css": true,
		".svg": true, ".png": true, ".jpg": true, ".jpeg": true, ".gif": true, ".webp": true,
	}
	for _, suffix := range DefaultCacheableSuffixes {
		if !allowed[suffix] {
			t.Fatalf("unexpected cacheable suffix %q in DefaultCacheableSuffixes - only "+
				"js/css/image types should be long-cached by default", suffix)
		}
	}
}

// TestDefaultIndexCacheControl_PreventsCaching guards against DefaultIndexCacheControl
// ever regressing to a directive (like a bare "no-cache") that still permits a
// CDN/browser to store a copy of the document, just with revalidation - the
// actual requirement is that a redeploy is visible on the very next request, which
// only "no-store" (never keep a copy at all) unambiguously guarantees.
func TestDefaultIndexCacheControl_PreventsCaching(t *testing.T) {
	if !strings.Contains(DefaultIndexCacheControl, "no-store") {
		t.Fatalf("expected DefaultIndexCacheControl to include \"no-store\" so a CDN/browser "+
			"never keeps a copy of the index document, got %q", DefaultIndexCacheControl)
	}
}
