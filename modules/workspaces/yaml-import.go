package workspaces

import (
	"embed"
	"fmt"

	"github.com/schollz/progressbar/v3"
)

func importYamlFromFileOnDisk[T any](
	importFilePath string,
	fn func(dto *T, query QueryDSL) (*T, *IError),
	f QueryDSL,
) {
	var items []T
	ReadYamlFile(importFilePath, &items)
	importYamlFromArray(items, fn, f, false)
}

func importYamlFromFileEmbed[T any](
	fsRef *embed.FS,
	importFilePath string,
	fn func(dto *T, query QueryDSL) (*T, *IError),
	f QueryDSL,
	silent bool,
) {
	var items []T
	ReadYamlFileEmbed(fsRef, importFilePath, &items)

	importYamlFromArray(items, fn, f, silent)
}

func importYamlFromArray[T any](
	items []T,
	fn func(dto *T, query QueryDSL) (*T, *IError),
	f QueryDSL,
	silent bool,
) {

	successInsert := 0
	failureInsert := 0

	bar := progressbar.Default(int64(len(items)))

	for _, item := range items {

		_, err := fn(&item, f)
		if err == nil {
			successInsert++
		} else {
			failureInsert++
		}

		bar.Add(1)
	}

	fmt.Println("Success", successInsert, "Failure", failureInsert)
}
