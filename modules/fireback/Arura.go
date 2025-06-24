package fireback

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"

	"log"

	"github.com/gin-gonic/gin"
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

func RenderPage(fsx fs.FS, c *gin.Context, page string, params any) {
	// Do not use FS until find out a way to embed properly

	renderPageFromEmbed(fsx, c, page, params)

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

		fmt.Println(path, d.Name(), d.Type())
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

// Do not use FS until find out a way to embed properly
func renderPageFromEmbed(fsx fs.FS, c *gin.Context, page string, params any) {

	address := "screens/" + page

	SpfNavigate := c.Query("spf") == "navigate"

	if SpfNavigate {
		c.Header("content-type", "application/json")
	} else {
		c.Header("content-type", "text/html")
	}

	sharedTemplates, err := loadSharedTemplatesEmbed(fsx, "screens/shared")
	if err != nil {
		fmt.Println(3, err)
		c.String(http.StatusInternalServerError, "Error loading shared templates")
		return
	}

	pages := append([]string{
		address,
	}, sharedTemplates...)

	tmpl, err := template.New(filepath.Base(address)).Funcs(funcMap).ParseFS(fsx, pages...)

	if err != nil {
		log.Default().Println(err)

		c.String(http.StatusInternalServerError, "Error loading page")
		return
	}

	var buf bytes.Buffer
	c.Writer.WriteHeader(http.StatusOK)

	err = tmpl.ExecuteTemplate(&buf, filepath.Base(address), params)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error rendering page")
	}

	if SpfNavigate {
		var jsonResponse map[string]interface{} = map[string]interface{}{}
		jsonResponse["body"] = gin.H{
			"content": buf.String(),
		}
		c.JSON(http.StatusOK, jsonResponse)

	} else {
		c.Writer.Write(buf.Bytes())
	}
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
