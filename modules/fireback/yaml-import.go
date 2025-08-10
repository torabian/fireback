package fireback

import (
	"embed"
	"encoding/base64"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	reflect "reflect"
	"regexp"
	"strings"

	"github.com/schollz/progressbar/v3"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

/*
Fireback comes with a custom yaml format to import data into the entities on database.
It's basically similar signature of the entity definition, and you can define
the entities as an array under 'items: section.

Also if the entities are having binary files (image, audio, etc...) can define an array
of those files on disk, (go embed) under 'resources' field it will upload them,
add record to FileEntity table, and they will be available in every field text of the import.

For example, consider list of books to be imported and their thumbnail:

items:
  - name: Book1
    thumbnailId: ($ref:book_thumbnail)

resources:
  - path: ./book.png
    key: book_thumbnail

Magic string placeholder ($ref:book_thumbnail) will be available on all fields which are strings,
so they will be match currectly to the uploaded file.
*/
type ContentImport[T any] struct {
	Items []T `json:"items" yaml:"items"`

	// list of files which will be uploaded and make available in the text fields of an entity
	Resources []struct {
		// path of the file relative to the import yaml file. for seeders and mocks,
		// they simply will be embed in go binary and it makes sense to use ./filename...
		// to access them
		Path string `yaml:"path"`

		// unique key in the importing file context, which will be available as ($ref:key) in the document
		Key string `yaml:"key"`

		// If true, it would upload it, rather reads it as []bytes and will replace it directly in the field.
		Blob bool `yaml:"blob"`
	} `yaml:"resources"`
}

func importYamlFromFileOnDisk[T any](
	importFilePath string,
	fn func(dto *T, query QueryDSL) (*T, *IError),
	f QueryDSL,
) {
	var content ContentImport[T]
	ReadYamlFile(importFilePath, &content)
	importYamlFromArray(content, fn, f, false)
}

// When we are working with objects, it's necessary to add a uniqueId to them
// when they are importing. This, do it :)
func AutoUniqueIdItems(input any) {
	v := reflect.ValueOf(input)
	addUniqueIdRecursively(v)
}

func detectMimeType(blob []byte, filename string) string {
	// Check extension
	if strings.HasSuffix(strings.ToLower(filename), ".svg") {
		return "image/svg+xml"
	}

	// Check content
	if len(blob) > 10 && strings.Contains(string(blob[:min(len(blob), 512)]), "<svg") {
		return "image/svg+xml"
	}

	// Fallback to built-in detection
	return http.DetectContentType(blob)
}

func ConvertBlobToDataURI(blob []byte, filename string) (string, error) {
	mimeType := detectMimeType(blob, filename) // detects from first 512 bytes
	b64 := base64.StdEncoding.EncodeToString(blob)
	dataURI := fmt.Sprintf("data:%s;base64,%s", mimeType, b64)
	return dataURI, nil
}

func replaceRef(input string, items []ResourceMap) string {
	var result string

	{
		re := regexp.MustCompile(`\(\$ref:([^\)]+)\)`)

		result = re.ReplaceAllStringFunc(input, func(match string) string {
			key := re.FindStringSubmatch(match)[1]
			for _, item := range items {
				if item.Key == key {
					return item.DriveId
				}
			}
			return match
		})
	}

	return result
}

func replaceRefBlob(input string, items []ResourceMap) string {
	var result string

	re := regexp.MustCompile(`\(\$refblob:([^\)]+)\)`)

	result = re.ReplaceAllStringFunc(input, func(match string) string {
		key := re.FindStringSubmatch(match)[1]
		for _, item := range items {
			if item.Key == key {

				// Temporarily we convert all files into the base64, and then
				// assign it. Might be more efficient if directly assigned as bytes,
				// but since it's an string replacement, it's working everywhere not only
				// on xfile? data type
				base64data, err := ConvertBlobToDataURI(item.Blob, "")
				if err != nil {
					LOG.Error("Converting a blob into base64 failed, %w", zap.Error(err))
					continue
				}

				return base64data

			}
		}
		return match
	})

	return result
}

func addUniqueIdRecursively(v reflect.Value) {
	// We need to handle the case where v is a pointer
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if v.Type().Field(i).Name == "UniqueId" {
				addUniqueIdRecursively(v.Field(i))
			}
		}
	case reflect.String:
		if v.CanSet() {
			value := v.String()
			if value == "" {
				fmt.Println("Setting a uniqueId")
				v.SetString(UUID())
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			addUniqueIdRecursively(v.Index(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			val := v.MapIndex(key)
			addUniqueIdRecursively(val)
		}
	}
}

func importYamlFromFileEmbed[T any](
	fsRef *embed.FS,
	importFilePath string,
	fn func(dto *T, query QueryDSL) (*T, *IError),
	f QueryDSL,
	silent bool,
) {
	resourceMap := ImportYamlFromFsResources(fsRef, importFilePath)

	rawContent, err := fs.ReadFile(fsRef, importFilePath)
	if err != nil {
		log.Fatalln(err)
	}

	// Let's replace the text level content, blobs will become strings
	// inefficient for very large files, but easier and reliable than reflect.
	rawContent = []byte(replaceRefBlob(string(rawContent), resourceMap))

	// Old school refereces
	rawContent = []byte(replaceRef(string(rawContent), resourceMap))

	var content ContentImport[T]
	if err := yaml.Unmarshal([]byte(rawContent), &content); err != nil {
		log.Default().Println("Yaml file is broken:", importFilePath)
	}

	importYamlFromArray(content, fn, f, silent)
}

func importYamlFromArray[T any](
	content ContentImport[T],
	fn func(dto *T, query QueryDSL) (*T, *IError),
	f QueryDSL,
	silent bool,
) {

	successInsert := 0
	failureInsert := 0

	bar := progressbar.Default(int64(len(content.Items)))

	for _, item := range content.Items {

		_, err := fn(&item, f)
		if err == nil {
			successInsert++
		} else {
			failureInsert++
			log.Default().Printf("error on insert: %v", err)
			fmt.Println(err.Error())
			fmt.Println(err.ToPublicEndUser(f).MessageTranslated)
		}

		bar.Add(1)
	}

	fmt.Println("Success", successInsert, "Failure", failureInsert)
}

type ResourceMap struct {
	FileId   string
	DriveId  string
	Key      string
	DiskPath string
	Blob     []byte
}

// Implement this
var ImportYamlFromFsResources func(fs *embed.FS, filePath string) []ResourceMap = func(fs *embed.FS, filePath string) []ResourceMap {
	fmt.Println("Importing file:", filePath, " is skipped. You need to override this function with a storage module")

	return []ResourceMap{}
}
