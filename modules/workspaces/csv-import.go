package workspaces

import (
	"embed"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/schollz/progressbar/v3"
)

/**
*	One day will come you need to refactor this
**/

// Converts a list of string into hireachical module 2 structure

func ToModule2Fields(list []string) []*Module2Field {
	fields := []*Module2Field{}

	for _, item := range list {
		fields = append(fields, &Module2Field{
			Name: ToCamelCaseClean(item),
			Type: FIELD_TYPE_STRING,
		})
	}

	return fields
}

func CastJsonFileToModule2Fields(importFilePath string) []*Module2Field {

	file, err := os.Open(importFilePath)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	file, err = os.Open(importFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(file)

	var index int64 = -1
	for {
		index++

		data, err := csvReader.Read()

		if err != nil {
			log.Fatalln(err)
		}

		if index == 0 {
			return ToModule2Fields(data)

		}

	}

	return []*Module2Field{}
}

// Source agnostic import, to be used for both file and go embed
func importCsvFromFileReader[T any](
	importFilePath string,
	fn func(dto *T, query QueryDSL) (*T, *IError),
	f QueryDSL,
) {

	successInsert := 0
	failureInsert := 0
	file, err := os.Open(importFilePath)
	if err != nil {
		log.Fatal(err)
	}
	lines := lineCounter(file)
	file.Close()

	file, err = os.Open(importFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bar := progressbar.Default(int64(lines))

	// read csv values using csv.Reader
	csvReader := csv.NewReader(file)

	headerMapping := []string{}

	var index int64 = -1
	for {
		index++

		data, err := csvReader.Read()

		if index == 0 {
			headerMapping = data
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			failureInsert++

			fmt.Println(index, err)
			continue
		}

		var item T

		/**
		*	This does not handle numbers, take care of the code somehow
		**/
		for index, col := range headerMapping {
			SetFieldString(&item, col, data[index])
		}

		_, err2 := fn(&item, f)

		if err2 == nil {
			successInsert++
		} else {
			failureInsert++
		}

		bar.Add(1)

	}

	fmt.Println("Success", successInsert, "Failure", failureInsert)
}

func importCsvFromEmbed[T any](
	fsRef *embed.FS,
	importFilePath string,
	fn func(dto *T, query QueryDSL) (*T, *IError),
	f QueryDSL,
) {

	successInsert := 0
	failureInsert := 0
	file, err := fsRef.Open(importFilePath)
	if err != nil {
		log.Fatal(err)
	}
	lines := lineCounter(file)
	file.Close()

	file, err = fsRef.Open(importFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bar := progressbar.Default(int64(lines))

	// read csv values using csv.Reader
	csvReader := csv.NewReader(file)

	headerMapping := []string{}

	var index int64 = -1
	for {
		index++

		data, err := csvReader.Read()

		if index == 0 {
			headerMapping = data
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			failureInsert++

			fmt.Println(index, err)
			continue
		}

		var item T

		/**
		*	This does not handle numbers, take care of the code somehow
		**/
		for index, col := range headerMapping {
			SetFieldString(&item, col, data[index])
		}

		_, err2 := fn(&item, f)

		if err2 == nil {
			successInsert++
		} else {
			failureInsert++
		}

		bar.Add(1)

	}

	fmt.Println("Success", successInsert, "Failure", failureInsert)
}
