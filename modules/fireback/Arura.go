package fireback

import (
	"bytes"
	"embed"
	"encoding/json"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
)

type Screen[T any] struct {
	Url          string
	DataProvider func(c *gin.Context) T
}

type XHtml struct {
	TemplateName string
	ScreensFs    fs.FS
	Params       any
}

type BaseScreenData struct {
	Error       error
	SpfNavigate bool
}

func ServeFileInliner(c *gin.Context, page string, file string, fsx fs.FS) {
	// get the directory of the page, e.g. for "foo/bar.html" → "foo/"
	pageDir := filepath.Dir(page)
	// join pageDir + file to get relative path
	relativePath := filepath.ToSlash(filepath.Join(pageDir, file))

	data, err := fs.ReadFile(fsx, "screens/"+relativePath)
	if err != nil {
		c.String(http.StatusNotFound, "File not found")
		return
	}
	// Detect content-type by file extension
	ext := strings.ToLower(filepath.Ext(file))
	var contentType string
	switch ext {
	case ".css":
		contentType = "text/css"
		cont, _ := resolveCssImportsFS(fsx, "screens/"+relativePath, map[string]bool{})

		m := minify.New()
		m.AddFunc("text/css", css.Minify)

		minified, err := m.String("text/css", cont)
		if err != nil {
			c.Data(http.StatusBadRequest, contentType, data)
			return
		}

		data = []byte(minified)
	case ".js":
		contentType = "application/javascript"
		data = serveJavascript(data, c.Request.URL.Path)

	case ".png":
		contentType = "image/png"
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".gif":
		contentType = "image/gif"
	case ".svg":
		contentType = "image/svg+xml"
	default:
		contentType = "application/octet-stream"
	}

	c.Data(http.StatusOK, contentType, data)
}

func RenderPage(fsx fs.FS, c *gin.Context, page string, params any) {

	file := c.Query("file")
	if file != "" {
		ServeFileInliner(c, page, file, fsx)
		return
	}

	address := "screens/" + page

	SpfNavigate := c.Query("spf") == "navigate"

	if SpfNavigate {
		c.Header("content-type", "application/json")
	} else {
		c.Header("content-type", "text/html")
	}

	sharedTemplates, err := loadSharedTemplatesEmbed(fsx, "screens/shared")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error loading shared templates")
		return
	}

	pages := append([]string{
		address,
	}, sharedTemplates...)

	tmpl, err := template.New(filepath.Base(address)).Funcs(funcMap).ParseFS(fsx, pages...)

	if err != nil {

		c.String(http.StatusInternalServerError, "Error loading page")
		return
	}

	var buf bytes.Buffer
	c.Writer.WriteHeader(http.StatusOK)

	err = tmpl.ExecuteTemplate(&buf, filepath.Base(address), params)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error rendering page:"+err.Error())
	}

	if SpfNavigate {
		var jsonResponse map[string]interface{} = map[string]interface{}{}
		jsonResponse["body"] = gin.H{
			"content": buf.String(),
		}
		c.JSON(http.StatusOK, jsonResponse)

	} else {
		htmlWithPrependedFiles := prependFileQuery(buf.String())
		c.Writer.Write([]byte(htmlWithPrependedFiles))
	}

}

func intRange(start, end int) []int {
	slice := make([]int, end-start+1)
	for i := range slice {
		slice[i] = start + i
	}
	return slice
}
func link(items []interface{}) string {
	var strItems []string
	for _, item := range items {
		strItems = append(strItems, item.(string)) // Convert each item to string and append to slice
	}
	return strings.Join(strItems, "/") // Join slice elements with "/"
}

func jsonx(u any) string {
	j, _ := json.MarshalIndent(u, "", "  ")

	return string(j)
}

var funcMap = template.FuncMap{
	"link": link,
	"json": jsonx,
	"arr":  func(els ...any) []any { return els },
}

type RenderContext struct {
	Fs   *embed.FS
	Path string
}

func loadSharedTemplatesEmbed(efs fs.FS, root string) ([]string, error) {
	var files []string

	err := fs.WalkDir(efs, root, func(path string, d fs.DirEntry, err error) error {
		if d == nil {
			return nil
		}

		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(d.Name()) == ".html" {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func loadSharedTemplates(sharedDir string) ([]string, error) {
	// Find all .html files in the shared directory
	files, err := filepath.Glob(filepath.Join(sharedDir, "*.html"))
	if err != nil {
		return nil, err
	}

	return files, nil
}

func prependFileQuery(htmlContent string) string {
	re := regexp.MustCompile(`(?i)<(img|script|link)[^>]*\s(href|src)=["'](\.\/)?([^"']+)["']`)

	return re.ReplaceAllStringFunc(htmlContent, func(m string) string {
		// Find href/src attr inside the matched tag string
		attrRe := regexp.MustCompile(`(?i)(href|src)=["'](\.\/)?([^"']+)["']`)
		return attrRe.ReplaceAllString(m, `$1="?file=$3"`)
	})
}

func resolveCssImportsFS(fsx fs.FS, path string, visited map[string]bool) (string, error) {
	if visited[path] {
		return "", nil // avoid circular imports
	}
	visited[path] = true

	data, err := fs.ReadFile(fsx, path)
	if err != nil {
		return "", err
	}

	dir := filepath.Dir(path)
	re := regexp.MustCompile(`@import\s+(url\()?["']?([^"')]+)["']?\)?[^;]*;`)
	content := re.ReplaceAllStringFunc(string(data), func(match string) string {
		matches := re.FindStringSubmatch(match)
		if len(matches) < 3 {
			return match // malformed
		}

		importPath := filepath.Join(dir, matches[2])
		imported, err := resolveCssImportsFS(fsx, importPath, visited)
		if err != nil {
			return "\n/* failed to import: " + matches[2] + " */\n"
		}

		return "\n/* " + matches[2] + " */\n" + imported
	})

	return content, nil
}

// Do not use FS until find out a way to embed properly
func renderPageFromEmbed(fsx fs.FS, c *gin.Context, page string, params any) {

}

func serveJavascript(data []byte, ginPath string) []byte {
	// Rewrite import statements only for JS files
	content := string(data)

	// Simple regex to rewrite imports of form: import ... from './something.js'
	importRe := regexp.MustCompile(`(?m)(import\s+[^'"]+['"])(\.\/[^'"]+)(['"])`)
	content = importRe.ReplaceAllStringFunc(content, func(s string) string {
		parts := importRe.FindStringSubmatch(s)
		if len(parts) < 4 {
			return s
		}
		prefix, path, suffix := parts[1], parts[2], parts[3]
		// Remove leading "./"
		cleanPath := strings.TrimPrefix(path, "./")
		// Rewrite to ?file=...

		if !strings.HasSuffix(cleanPath, ".js") {
			cleanPath = cleanPath + ".js"
		}

		lastSegment := filepath.Base(ginPath)

		// lastSegment := filepath.Base(c.Request.URL.Path)

		return prefix + "./" + lastSegment + "?file=" + cleanPath + suffix
	})

	return []byte(content)
}

func ResolveScreens(embedscreesn embed.FS) fs.FS {
	if !GetConfig().Production {
		_, filename, _, _ := runtime.Caller(1)
		dir := filepath.Join(filepath.Dir(filename))
		return os.DirFS(dir) // OK because DirFS returns fs.FS
	} else {
		return embedscreesn
	}
}
