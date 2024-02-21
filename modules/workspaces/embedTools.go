package workspaces

import (
	"embed"
	"encoding/json"
	"html/template"
	"io/fs"
	"net/http"
	"os"
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

func EmbedFolderForGin(ui *embed.FS, folder string, r *gin.Engine) {
	fs := EmbedFolder(*ui, folder, true)
	staticServer := static.Serve("/", fs)
	r.Use(staticServer)
	r.NoRoute(func(c *gin.Context) {
		if c.Request.Method == http.MethodGet &&
			!strings.ContainsRune(c.Request.URL.Path, '.') &&
			!strings.HasPrefix(c.Request.URL.Path, "/api/") {
			c.Request.URL.Path = "/"
			staticServer(c)
		}
	})
}

func HasChildren(key string, items []string) bool {

	for _, perm := range items {
		if strings.Contains(perm, key+"/") {
			return true
		}
	}

	return false
}

func SyncPermission(modules []*ModuleProvider) {

	data := map[string]string{}
	for _, item := range modules {
		item.PermissionsProvider = append(item.PermissionsProvider, "root/*")

		for _, perm := range item.PermissionsProvider {

			key := strings.ToUpper(strings.ReplaceAll(strings.ReplaceAll(perm, "/", "_"), "_*", ""))

			if HasChildren(perm, item.PermissionsProvider) {
				data[key] = perm + "/*"
			} else {
				data[key] = perm
			}
		}
	}

	dicJson, _ := json.MarshalIndent(data, "", "  ")
	os.Mkdir("./artifacts/intermediate-http", 0777)
	os.WriteFile("./artifacts/intermediate-http/permissions.json", dicJson, 0644)

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
