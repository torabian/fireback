# Embedding UI folders (`PublicFolderInfo`)

`Application.PublicFolders` (see `modules/fireback/application/Application.go`) is how a
project embeds a built SPA (React, Angular, ...) into the Go binary and serves it. Each
entry is a `gintools.PublicFolderInfo`:

```go
type PublicFolderInfo struct {
    Fs     *embed.FS
    Folder string
    Prefix string

    DisableGzip            bool
    AssetCacheControl      string
    ExtraCacheableSuffixes []string
    IndexCacheControl      string
    InjectHTML             func(c *gin.Context, path string) (HTMLInjection, bool)
}
```

`Fs`/`Folder`/`Prefix` are the original three - what's embedded, the sub-directory of the
`embed.FS` to serve, and the URL prefix to mount it under (`"/"` for the default/root SPA,
`"/manage"` for an admin UI, etc). `EmbedFoldersForGin` mounts every folder and installs one
combined SPA fallback, so a browser refresh on a deep client-side route (e.g.
`/scores/abc123`) still resolves to the right SPA's `index.html` instead of 404ing - see that
function's doc comment for how the longest-matching `Prefix` wins.

The rest are optional and cover three things useful behind a CDN (Cloudflare etc):

## 1. Gzip

Gzip compression happens **entirely inside each `PublicFolderInfo`'s own request handling** -
never as a blanket engine-wide middleware, and deliberately not even as a "compress anything
whose path starts with this folder's `Prefix`" middleware. The reason for that second, less
obvious step: a folder's `Prefix` is very commonly `"/"` (the default/root SPA), and *every*
request path is a string-prefix match for `"/"` - so a prefix-based check would end up
compressing every route in the app, including ones mounted completely outside
`PublicFolders` (a module's own API group, and in particular the storage module's tus upload
endpoints and its manual Range/streaming download handler). Those depend on exact byte counts
and streaming semantics (tus's `Content-Length`/`Upload-Offset` contract, HTTP `Range`/`206`
responses) that gzip rewriting can silently break.

Instead, `mountEmbedFolder` only ever wraps the response in gzip once *this specific folder*
has independently confirmed it's the one about to write it:

- For a real static asset, it checks the embedded filesystem's own `Exists` before deciding to
  compress - the same check `static.Serve` itself uses to decide whether to serve the file at
  all - so a request that doesn't match a file in this folder falls straight through,
  untouched, to whatever handles it next.
- For the SPA/index fallback, it's already precisely scoped for free: that closure is only
  ever invoked by `NoRoute` for a path nothing else - no registered route, no other folder's
  more specific prefix - matched.

Every folder gets gzip by default; set `DisableGzip: true` on a folder to opt that specific one
out too (e.g. content that's already compressed, or mostly tiny files where compression isn't
worth the CPU).

```go
{Fs: &ui, Folder: "public", Prefix: "/"},                     // gzip: on (default, scoped to "/")
{Fs: &adminUi, Folder: "manage", Prefix: "/manage", DisableGzip: true},
// Routes registered elsewhere (an API group, storage.StorageModuleSetup's tus
// endpoints, ...) are never in PublicFolders, so gzip can never reach them.
```

See `EmbedTools_test.go`'s `TestEmbedFoldersForGin_GzipAndInjectHTML` for this in action: a
plain route registered *after* a `"/"` folder (so it inherits that folder's own middleware,
same as a real module's API route would) never gets compressed, while `/item1`/`/item2`
requests that fall through to the SPA both get compressed **and** get their own per-route
`InjectHTML` content spliced in first.

## 2. Cache-Control

A CDN/browser needs to know it can cache a content-hashed `main.a1b2c3.js` forever, but must
*not* cache `index.html` the same way (or a deploy would go invisible until the cache
expires). `EmbedFoldersForGin` sets this for you:

- Static asset requests (`.js .css .svg .png .jpg .jpeg .gif .webp .woff .woff2 .ttf .eot
  .ico` - see `DefaultCacheableSuffixes`) get `DefaultAssetCacheControl`
  (`"public, max-age=604800, immutable"`).
- The index document itself (a literal `/index.html` request, and the SPA fallback used for
  unknown routes) gets `DefaultIndexCacheControl` (`"no-cache"`).

Override either per folder:

```go
{
    Fs: &ui, Folder: "public", Prefix: "/",
    AssetCacheControl:      "public, max-age=31536000, immutable",
    ExtraCacheableSuffixes: []string{".json", ".wasm"},
    IndexCacheControl:      "no-store",
},
```

`AssetCacheControlMiddleware(items)` (what this wires up under the hood) picks the
longest-matching `Prefix` when more than one folder's rule could apply to a path.

## 3. Per-route `<head>` injection (`InjectHTML`)

A single-page app has one static `index.html` for every route, so a chat app/social
network/search crawler unfurling a link to a deep route (e.g. `/scores/abc123`) - which
doesn't run the SPA's JS - only ever sees the app shell's generic `<title>`/OpenGraph tags,
never anything about that specific score/item/whatever.

`InjectHTML`, if set, is called for every request that falls through to that folder's SPA
fallback (i.e. didn't match a real embedded file):

```go
InjectHTML func(c *gin.Context, path string) (HTMLInjection, bool)

type HTMLInjection struct {
    Head string // spliced in just before </head>
    Body string // spliced in just before </body>
}
```

Return `(HTMLInjection{}, false)` to decline - the plain `index.html` is served unchanged,
exactly as before `InjectHTML` existed. Return `(injection, true)` to have `Head`/`Body`
spliced into that request's copy of `index.html` (raw HTML, inserted verbatim - escape any
request- or DB-derived value yourself, e.g. with `html.EscapeString`, before putting it in a
tag). This is not a real SSR pipeline - the SPA still renders on the client the same as
always - it's a small, cheap way to give crawlers/link-unfurlers real per-route metadata.

```go
{
    Fs: &ui, Folder: "public", Prefix: "/",
    InjectHTML: score.InjectScoreOgTags,
},
```

See `modules/score/ScoreOgTagsInjection.go` in the `nima` project for a worked example: it
matches `/scores/:id`, loads that score, and returns `<title>`/`og:title`/`og:description`/
`og:image` tags built from it.
