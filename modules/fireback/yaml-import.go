package fireback

import (
	"embed"
	"fmt"
	"log"
	reflect "reflect"
	"regexp"

	"github.com/schollz/progressbar/v3"
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

func ReplaceResourcesInStruct(input any, resourceList []ResourceMap) {
	v := reflect.ValueOf(input)
	replaceStringsRecursively(v, resourceList)
	addUniqueIdRecursively(v)
}

// When we are working with objects, it's necessary to add a uniqueId to them
// when they are importing. This, do it :)
func AutoUniqueIdItems(input any) {
	v := reflect.ValueOf(input)
	addUniqueIdRecursively(v)
}

func replaceRef(input string, items []ResourceMap) string {
	re := regexp.MustCompile(`\(\$ref:([^\)]+)\)`)

	result := re.ReplaceAllStringFunc(input, func(match string) string {
		key := re.FindStringSubmatch(match)[1]
		for _, item := range items {
			if item.Key == key {
				return item.DriveId
			}
		}
		return match
	})

	return result
}

/*
Iterates through an struct instance, and it will replace the ($ref:key) based on the resourceList
paramters in all of them. It's useful for importing, exporting context to manage the binary files
to be downloaded or uploaded
*/
func replaceStringsRecursively(v reflect.Value, resourceList []ResourceMap) {
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
			replaceStringsRecursively(v.Field(i), resourceList)
		}
	case reflect.String:
		if v.CanSet() {
			value := v.String()
			v.SetString(replaceRef(value, resourceList))

		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			replaceStringsRecursively(v.Index(i), resourceList)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			val := v.MapIndex(key)
			replaceStringsRecursively(val, resourceList)
		}
	}
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
	var content ContentImport[T]
	if err := ReadYamlFileEmbed(fsRef, importFilePath, &content); err != nil {
		log.Fatalln(err)
	}
	resourceMap := ImportYamlFromFsResources(fsRef, importFilePath)

	for _, item := range content.Items {
		ReplaceResourcesInStruct(item, resourceMap)
	}

	importYamlFromArray(content, fn, f, silent)
}

func importYamlFromFileEmbedBatch[T any](
	fsRef *embed.FS,
	importFilePath string,
	fn func(dto []*T, query QueryDSL) ([]*T, *IError),
	f QueryDSL,
	silent bool,
) {
	var content ContentImport[T]
	if err := ReadYamlFileEmbed(fsRef, importFilePath, &content); err != nil {
		log.Fatalln(err)
	}
	resourceMap := ImportYamlFromFsResources(fsRef, importFilePath)

	for _, item := range content.Items {
		ReplaceResourcesInStruct(item, resourceMap)
	}

	importYamlFromArrayBatch(content, fn, f, silent)
}

func importYamlFromArrayBatch[T any](
	content ContentImport[T],
	fn func(dto []*T, query QueryDSL) ([]*T, *IError),
	f QueryDSL,
	silent bool,
) {

	successInsert := 0
	failureInsert := 0

	bar := progressbar.Default(int64(len(content.Items)))

	count := 0
	items := []*T{}

	for _, item := range content.Items {
		items = append(items, &item)
		count++

		if count == 10 {
			_, err := fn(items, f)
			if err == nil {
				successInsert++
			} else {
				failureInsert++
			}

			bar.Add(count)
			count = 0
			items = []*T{}
		}
	}

	_, err := fn(items, f)
	if err == nil {
		successInsert++
	} else {
		failureInsert++
	}

	fmt.Println("Success", successInsert, "Failure", failureInsert)
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
}

// Implement this
var ImportYamlFromFsResources func(fs *embed.FS, filePath string) []ResourceMap = func(fs *embed.FS, filePath string) []ResourceMap {
	fmt.Println("Importing file:", filePath, " is skipped. You need to override this function with a storage module")

	return []ResourceMap{}
}
