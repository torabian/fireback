package workspaces

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

type embedFileSystem struct {
	http.FileSystem
	indexes bool
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
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

func EmbedFolderForGin(ui *embed.FS, folder string, r *gin.Engine, prefix string) {

	// I am not sure about this.
	if prefix == "" {
		prefix = "/"
	}

	fs := EmbedFolder(*ui, folder, true)

	if prefix == "/" {
		staticServer := static.Serve("/", fs)

		r.Use(staticServer)
		r.NoRoute(func(c *gin.Context) {
			if c.Request.Method == http.MethodGet &&
				!strings.ContainsRune(c.Request.URL.Path, '.') &&
				!strings.HasPrefix(c.Request.URL.Path, "/api/") {
				c.Request.URL.Path = prefix
				staticServer(c)
			}
		})
	} else {

		fileServer := http.StripPrefix(prefix, http.FileServer(fs))

		r.GET(prefix+"/*filepath", func(c *gin.Context) {
			fileServer.ServeHTTP(c.Writer, c.Request)
		})

		r.NoRoute(func(c *gin.Context) {
			if c.Request.Method == http.MethodGet &&
				!strings.ContainsRune(c.Request.URL.Path, '.') &&
				!strings.HasPrefix(c.Request.URL.Path, "/api/") {
				c.Request.URL.Path = "/"
				fileServer.ServeHTTP(c.Writer, c.Request)
			}
		})

	}
}

func HasChildren(key string, items []string) bool {

	for _, perm := range items {
		if strings.Contains(perm, key+"/") {
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
