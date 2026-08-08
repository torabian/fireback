package gintools

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

/*
Public folders are used when you want to make a embed folder available though the web
server publicly. Quite useful on serving static content, also you can prefix them
*/
type PublicFolderInfo struct {
	Fs     *embed.FS
	Folder string
	Prefix string
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

// mountEmbedFolder mounts a public folder as static-file middleware (rather
// than a route) so that requests for paths it doesn't have a file for simply
// fall through instead of terminating in a raw 404 - letting the combined
// NoRoute handler in EmbedFoldersForGin decide on a SPA fallback. It returns
// that fallback handler, which re-serves the folder's index.html for
// SPA-style client-side routing (e.g. a React Router deep link on refresh).
func mountEmbedFolder(ui *embed.FS, folder string, r *gin.Engine, prefix string) func(c *gin.Context) {
	if prefix == "" {
		prefix = "/"
	}

	fs := EmbedFolder(*ui, folder, true)
	staticServer := static.Serve(prefix, fs)
	r.Use(staticServer)

	return func(c *gin.Context) {
		c.Request.URL.Path = prefix
		staticServer(c)
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
			handle: mountEmbedFolder(item.Fs, item.Folder, r, item.Prefix),
		})
	}

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
