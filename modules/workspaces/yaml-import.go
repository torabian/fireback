package workspaces

import (
	"embed"
	"fmt"
	"path"
	"path/filepath"
	"regexp"

	"github.com/schollz/progressbar/v3"
)

type ContentImport[T any] struct {
	Items     []T `json:"items" yaml:"items"`
	Resources []struct {
		Path string `yaml:"path"`
		Key  string `yaml:"key"`
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

func importYamlFromFileEmbed[T any](
	fsRef *embed.FS,
	importFilePath string,
	fn func(dto *T, query QueryDSL) (*T, *IError),
	f QueryDSL,
	silent bool,
) {
	var content ContentImport[T]
	ReadYamlFileEmbed(fsRef, importFilePath, &content)
	ImportYamlFromFsResources(fsRef, importFilePath)
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
			fmt.Println(err)
			failureInsert++
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

// func ImportYamlResources(filePath string) []ResourceMap {
// 	result := []ResourceMap{}
// 	var resources ContentImport[any]
// 	ReadYamlFile(filePath, &resources)

// 	for _, resource := range resources.Resources {
// 		actualPath := path.Join(filepath.Dir(filePath), resource.Path)
// 		entity, fileId := UploadFromDisk(actualPath)
// 		result = append(result, ResourceMap{
// 			DriveId:  entity.UniqueId,
// 			FileId:   fileId,
// 			Key:      resource.Key,
// 			DiskPath: actualPath,
// 		})

// 	}

// 	return result
// }

func ImportYamlFromFsResources(fs *embed.FS, filePath string) []ResourceMap {
	result := []ResourceMap{}
	var resources ContentImport[any]
	err := ReadYamlFileEmbed(fs, filePath, &resources)

	if err != nil {
		fmt.Println("Error importing content:", err, filePath)
	}

	for _, resource := range resources.Resources {
		actualPath := path.Join(filepath.Dir(filePath), resource.Path)
		entity, fileId := UploadFromFs(fs, actualPath)
		result = append(result, ResourceMap{
			DriveId:  entity.UniqueId,
			FileId:   fileId,
			Key:      resource.Key,
			DiskPath: actualPath,
		})

	}

	return result
}

func ApplyResourceMap(content string, rm []ResourceMap, mode string) string {

	r, _ := regexp.Compile(`(\{\{resource:(.*?)\}\})`)

	data := ReplaceAllStringSubmatchFunc(r, content, func(str []string) string {

		result := getFromResourceMap(str[2], rm, mode)

		return result

	})

	return data
}

func ReplaceAllStringSubmatchFunc(re *regexp.Regexp, str string, repl func([]string) string) string {
	result := ""
	lastIndex := 0

	for _, v := range re.FindAllSubmatchIndex([]byte(str), -1) {
		groups := []string{}
		for i := 0; i < len(v); i += 2 {
			groups = append(groups, str[v[i]:v[i+1]])
		}

		result += str[lastIndex:v[0]] + repl(groups)
		lastIndex = v[1]
	}

	return result + str[lastIndex:]
}

func getFromResourceMap(key string, rm []ResourceMap, mode string) string {

	for _, v := range rm {
		if v.Key == key {
			if mode == "directasset" {
				return "directasset_____" + v.DiskPath + "_____"
			} else if mode == "drive" {
				return v.DriveId
			} else {
				return "fbtusid_____" + v.FileId + "_____"
			}
		}
	}
	return ""
}
