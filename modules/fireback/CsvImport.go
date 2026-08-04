package fireback

import (
	"embed"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	reflect "reflect"
	"strings"

	"github.com/schollz/progressbar/v3"
	"github.com/torabian/fireback/modules/fireback/reflectionutil"
)

/**
*	One day will come you need to refactor this
**/

// Source agnostic import, to be used for both file and go embed
func importCsvFromFileReader[T any](
	importFilePath string,
	fn func(dto *T, query QueryDSL) (*T, *IError),
	f QueryDSL,
	initializer func() T,

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

			// fmt.Println(index, err)
			continue
		}

		var item T = initializer()

		/**
		*	This does not handle numbers, take care of the code somehow
		**/
		for index, col := range headerMapping {
			reflectionutil.SetValue(&item, col, &data[index])
		}

		// id := "b1be9b82"
		// SetValue(&item, "ProductId", &id)

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

func SetField(item interface{}, fieldName string, value interface{}) error {
	v := reflect.ValueOf(item).Elem()
	if !v.CanAddr() {
		return fmt.Errorf("cannot assign to the item passed, item must be a pointer in order to assign")
	}
	// It's possible we can cache this, which is why precompute all these ahead of time.
	findJsonName := func(t reflect.StructTag) (string, error) {
		if jt, ok := t.Lookup("json"); ok {
			return strings.Split(jt, ",")[0], nil
		}
		return "", fmt.Errorf("tag provided does not define a json tag %s", fieldName)
	}
	fieldNames := map[string]int{}
	for i := 0; i < v.NumField(); i++ {
		typeField := v.Type().Field(i)
		tag := typeField.Tag
		jname, _ := findJsonName(tag)
		fieldNames[jname] = i
	}

	fieldNum, ok := fieldNames[fieldName]
	if !ok {
		return fmt.Errorf("field %s does not exist within the provided item", fieldName)
	}
	fieldVal := v.Field(fieldNum)
	val := reflect.ValueOf(value)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	fieldVal.Set(val)
	return nil
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
			SetField(&item, col, data[index])
		}

		fmt.Println(item)
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
