package workspaces

import (
	"bufio"
	"bytes"
	"embed"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/gin-gonic/gin"
	"github.com/schollz/progressbar/v3"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

func CommonCliQueryDSLBuilder(c *cli.Context) QueryDSL {
	queryString := c.String("query")
	startIndex := c.Int("startIndex")
	itemsPerPage := c.Int("itemsPerPage")

	if startIndex < 0 {
		startIndex = 0
	}

	switch {
	case itemsPerPage > 1000:
		itemsPerPage = 1000
	case itemsPerPage <= 0:
		itemsPerPage = 20
	}

	lang := "en"
	region := "US"
	workspaceId := "root"
	userId := ""
	cfg := GetAppConfig()

	if cfg.WorkspaceAs != "" {
		workspaceId = cfg.WorkspaceAs
	}
	if cfg.UserAs != "" {
		userId = cfg.UserAs
	}
	if cfg.CliLanguage != "" {
		lang = cfg.CliLanguage
	}
	if cfg.CliRegion != "" {
		region = cfg.CliRegion
	}

	if c.IsSet("user-id") {
		userId = c.String("user-id")
	}

	withPreloads := c.String("wp")

	var f QueryDSL = QueryDSL{
		Query:        queryString,
		StartIndex:   startIndex,
		WorkspaceId:  workspaceId,
		UserId:       userId,
		Language:     lang,
		Region:       strings.ToUpper(region),
		ItemsPerPage: itemsPerPage,
		UserHas:      []string{"root/*"},
	}

	if len(withPreloads) > 0 {
		f.WithPreloads = strings.Split(strings.Trim(withPreloads, " "), ",")
	}

	if c.IsSet("lang") {
		f.Language = c.String("lang")
	}

	if c.IsSet("deep") {
		f.Deep = c.Bool("deep")
	}

	if c.IsSet("workspaceId") {
		f.WorkspaceId = c.String("workspaceId")
	}

	if c.IsSet("userId") {
		f.UserId = c.String("userId")
	}

	if c.IsSet("id") {
		f.UniqueId = c.String("id")
		fmt.Println(f.UniqueId)
	}

	return f
}

func lineCounter(r io.Reader) int {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count

		case err != nil:
			return count
		}
	}
}
func CommonCliQueryCmd[T any](
	c *cli.Context,
	fn func(query QueryDSL) ([]T, *QueryResultMeta, error),
) {

	f := CommonCliQueryDSLBuilder(c)

	if items, count, err := fn(f); err != nil {
		fmt.Println(err)
	} else {
		jsonString, _ := json.MarshalIndent(gin.H{
			"data": gin.H{
				"startIndex":   f.StartIndex,
				"itemsPerPage": f.ItemsPerPage,
				"items":        items,
				"totalItems":   count.TotalItems,
			},
		}, "", "  ")

		fmt.Println(string(jsonString))
	}
}

func GetColumnsFromReflect[T any](v reflect.Value) []string {
	verbose := false

	headers := []string{}
	for j := 0; j < v.NumField(); j++ {
		n := v.Type().Field(j).Name

		if strings.ToUpper(n[0:1]) != n[0:1] {
			continue
		}

		if Contains(FIREBACK_DEFAULT_DB_COLUMNS, n) && !verbose {
			continue
		}

		headers = append(headers, n)
	}

	return headers

}

var FIREBACK_DEFAULT_DB_COLUMNS []string = []string{
	"LinkerId",
	"WorkspaceId",
	"Translations",
	"Updated",
	"Created",
	"Visibility",
	"ParentId",
}

func ExtractRowStringValues[T any](row *T, v reflect.Value, verbose bool) []string {
	data := []string{}
	for j := 0; j < v.NumField(); j++ {

		f := v.Field(j)
		n := v.Type().Field(j).Name
		t := f.Type().String()

		if strings.ToUpper(n[0:1]) != n[0:1] {
			continue
		}

		if Contains(FIREBACK_DEFAULT_DB_COLUMNS, n) && !verbose {
			continue
		}

		value := ExtractStringValueFromReflectCell[T](row, t, n)

		data = append(data, value)
	}

	return data
}
func ExtractStringValueFromReflectCell[T any](row *T, t string, n string) string {
	value := ""

	if t == "string" {
		value = GetFieldString(row, n)
	} else if t == "*string" {
		value = GetFieldStringP(row, n)
	} else if t == "int32" || t == "int64" || t == "int" {
		value = fmt.Sprint(GetFieldInt(row, n))
	} else if t == "bool" {
		value = fmt.Sprint(GetFieldBool(row, n))
	} else if t == "float64" {
		value = fmt.Sprint(GetFieldFloat(row, n))
	} else if t == "*float64" {
		value = fmt.Sprint(GetFieldFloatP(row, n))
	} else if t == "*int64" {
		v0 := GetFieldInt64P(row, n)
		if v0 == nil {
			value = "N/A"
		} else {
			value = fmt.Sprint(*v0)
		}
	} else {
		value = "N/A"
	}
	return value
}
func CommonCliTableCmd[T any](
	c *cli.Context,
	fn func(query QueryDSL) ([]*T, *QueryResultMeta, error),
	v reflect.Value,
) {

	verbose := false
	if c.IsSet("verbose") && c.Bool("verbose") {
		verbose = true
	}

	f := CommonCliQueryDSLBuilder(c)
	items, count, err := fn(f)
	fmt.Println("Count", count)

	if err != nil {
		fmt.Println(err)
		panic("Cannot query")
	}

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
		},
	}

	for _, n := range GetColumnsFromReflect[T](v) {
		table.Header.Cells = append(table.Header.Cells,
			&simpletable.Cell{Align: simpletable.AlignLeft, Text: n},
		)
	}

	var counter = 0
	for _, row := range items {
		counter++
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", counter)},
		}

		for _, cellValue := range ExtractRowStringValues[T](row, v, verbose) {

			r = append(r, &simpletable.Cell{
				Align: simpletable.AlignRight, Text: cellValue,
			})
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleDefault)
	fmt.Println(table.String())
}

func CommonCliImportCmd[T any](
	c *cli.Context,
	fn func(dto *T, query QueryDSL) (*T, *IError),
	v reflect.Value,
	importFilePath string,
) {

	f := CommonCliQueryDSLBuilder(c)
	f.Deep = true

	if strings.Contains(importFilePath, ".yml") || strings.Contains(importFilePath, ".yaml") {
		importYamlFromFileOnDisk(importFilePath, fn, f)
	}

	if strings.Contains(importFilePath, ".csv") {
		importCsvFromFileReader(importFilePath, fn, f)
	}

}

func CommonCliImportEmbedCmd[T any](
	c *cli.Context,
	fn func(dto *T, query QueryDSL) (*T, *IError),
	v reflect.Value,
	fsRef *embed.FS,
) {
	f := CommonCliQueryDSLBuilder(c)
	f.WorkspaceId = "system"
	SeederFromFSImport(f, fn, v, fsRef, []string{}, false)
}

func SeederFromFSImport[T any](
	f QueryDSL,
	fn func(dto *T, query QueryDSL) (*T, *IError),
	v reflect.Value,
	fsRef *embed.FS,
	fileNames []string,
	silent bool,
) {

	f.Deep = true

	if entity, err := GetSeederFilenames(fsRef, ""); err != nil {
		fmt.Println(err.Error())
	} else {

		for _, path := range entity {
			if len(fileNames) > 0 && !Contains(fileNames, path) {
				continue
			}

			if strings.Contains(path, ".yml") || strings.Contains(path, ".yaml") {
				importYamlFromFileEmbed(fsRef, path, fn, f, silent)
			}

			if strings.Contains(path, ".csv") {
				importCsvFromEmbed(fsRef, path, fn, f)
			}
		}

	}

}

// func Unmarshal(reader *csv.Reader, v interface{}) error {
// 	record, err := reader.Read()
// 	if err != nil {
// 		return err
// 	}
// 	s := reflect.ValueOf(v).Elem()
// 	if s.NumField() != len(record) {
// 		return &FieldMismatch{s.NumField(), len(record)}
// 	}
// 	for i := 0; i < s.NumField(); i++ {
// 		f := s.Field(i)
// 		switch f.Type().String() {
// 		case "string":
// 			f.SetString(record[i])
// 		case "int":
// 			ival, err := strconv.ParseInt(record[i], 10, 0)
// 			if err != nil {
// 				return err
// 			}
// 			f.SetInt(ival)
// 		default:
// 			return &UnsupportedType{f.Type().String()}
// 		}
// 	}
// 	return nil
// }

type ExportCatalog[T any] struct {
	Writer             *os.File
	ReadSize           int64
	TotalItemsToExport int64
	F                  QueryDSL
	ExportFilePath     string
	QueryResultMeta    *QueryResultMeta
	Fn                 func(query QueryDSL) ([]*T, *QueryResultMeta, error)
}

func YamlExporterChannel[T any](
	query QueryDSL,
	fn func(query QueryDSL) ([]*T, *QueryResultMeta, error),
	preloads []string,
) (chan []byte, *IError) {

	chanStream := make(chan []byte)

	query.Deep = true
	query.WithPreloads = append(
		query.WithPreloads,
		preloads...,
	)
	_, count, _ := fn(query)

	catalog := &ExportCatalog[T]{
		ReadSize:        10,
		QueryResultMeta: count,
		F:               query,
		Fn:              fn,
	}

	go func() {
		defer close(chanStream)

		var index int64 = 0
		for ; index <= catalog.QueryResultMeta.TotalItems; index += catalog.ReadSize {

			catalog.F.ItemsPerPage = int(catalog.ReadSize)
			catalog.F.StartIndex = int(index)
			items, _, _ := catalog.Fn(catalog.F)

			if len(items) > 0 {
				data, _ := yaml.Marshal(items)
				chanStream <- data
			}
		}
	}()

	return chanStream, nil
}

func YamlExporterChannelT[T any](
	query QueryDSL,
	fn func(query QueryDSL) ([]*T, *QueryResultMeta, error),
	preloads []string,
) (chan []interface{}, *IError) {

	chanStream := make(chan []interface{})

	query.Deep = true
	query.WithPreloads = append(
		query.WithPreloads,
		preloads...,
	)
	_, count, _ := fn(query)

	catalog := &ExportCatalog[T]{
		ReadSize:        10,
		QueryResultMeta: count,
		F:               query,
		Fn:              fn,
	}

	go func() {
		defer close(chanStream)

		var index int64 = 0
		for ; index <= catalog.QueryResultMeta.TotalItems; index += catalog.ReadSize {

			catalog.F.ItemsPerPage = int(catalog.ReadSize)
			catalog.F.StartIndex = int(index)
			items, _, _ := catalog.Fn(catalog.F)

			if len(items) > 0 {
				var m []interface{} = []interface{}{}
				for _, item := range items {
					m = append(m, item)
				}
				chanStream <- m
			}
		}
	}()

	return chanStream, nil
}

func YamlExporter[T any](catalog *ExportCatalog[T], bar *progressbar.ProgressBar) {
	enc := yaml.NewEncoder(catalog.Writer)

	var index int64 = 0
	for ; index <= catalog.QueryResultMeta.TotalItems; index += catalog.ReadSize {

		catalog.F.ItemsPerPage = int(catalog.ReadSize)
		catalog.F.StartIndex = int(index)
		items, _, _ := catalog.Fn(catalog.F)

		if len(items) > 0 {
			err := enc.Encode(items)
			bar.Add(len(items))

			if err != nil {
				fmt.Println(err)
			}
		}
	}

	catalog.Writer.Close()

	// Since yaml package adds a lot of ---, now let's read line by line, and delete them
	inFile, _ := os.Open(catalog.ExportFilePath)
	defer inFile.Close()

	outFile, _ := os.OpenFile(catalog.ExportFilePath, os.O_RDWR, 0644)
	defer outFile.Close()

	reader := bufio.NewReaderSize(inFile, 10*1024)

	for {
		line, err := reader.ReadString('\n')
		if strings.Contains(line, "---") {
			outFile.WriteString("###\n")
		} else {
			outFile.WriteString(line)
		}
		if err != nil {
			if err != io.EOF {
				fmt.Println("error:", err)
			}
			break
		}
	}
}

func CommonCliExportCmd[T any](
	c *cli.Context,
	fn func(query QueryDSL) ([]*T, *QueryResultMeta, error),
	v reflect.Value,
	exportFilePath string,
	translationRef *embed.FS,
	fsFileName string,
	detectedPreloads []string,
) {

	f := CommonCliQueryDSLBuilder(c)
	f.Deep = true
	f.WithPreloads = append(f.WithPreloads, detectedPreloads...)

	_, count, err := fn(f)
	bar := progressbar.Default(int64(count.TotalItems))

	if err != nil {
		fmt.Println(err)
	}
	writer, err := os.Create(exportFilePath)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	translationBox := map[string]interface{}{}
	ReadYamlFileEmbed[map[string]interface{}](translationRef, fsFileName, &translationBox)

	// fmt.Println(72, GetTranslationKey(translationBox, "capabilities", "en"))

	catalog := &ExportCatalog[T]{
		Writer:          writer,
		ReadSize:        2,
		ExportFilePath:  exportFilePath,
		QueryResultMeta: count,
		F:               f,
		Fn:              fn,
	}

	if strings.Contains(exportFilePath, ".yml") || strings.Contains(exportFilePath, ".yaml") {
		YamlExporter[T](catalog, bar)
	}

	data := &PdfExportData{
		Name:        "General Report",
		Description: "General report of the entities",
		FieldsMap:   map[string]string{
			// "UserId": "کد کاربر",
		},
	}

	if strings.Contains(exportFilePath, ".pdf") {
		PdfExporter[T](exportFilePath, f, fn, v, bar, data)
	}

	if strings.Contains(exportFilePath, ".csv") {
		defer writer.Close()
		csvwriter := csv.NewWriter(writer)
		header := []string{}

		for j := 0; j < v.NumField(); j++ {
			n := v.Type().Field(j).Name

			if strings.ToUpper(n[0:1]) != n[0:1] {
				continue
			}

			header = append(header, n)
		}

		err2 := csvwriter.Write(header)

		if err2 != nil {
			log.Fatalln("Failed written header", err2)
		}

		var index int64 = 0
		for ; index <= count.TotalItems; index += catalog.ReadSize {

			f.ItemsPerPage = 50
			f.StartIndex = int(index)
			items, _, err := fn(f)

			if err != nil {
				fmt.Println(err)
			}

			for _, item := range items {

				data := []string{}

				for j := 0; j < v.NumField(); j++ {

					f := v.Field(j)
					n := v.Type().Field(j).Name
					t := f.Type().String()

					if strings.ToUpper(n[0:1]) != n[0:1] {
						continue
					}

					value := ""

					if t == "string" {
						value = GetFieldString(item, n)
					}

					if t == "int32" || t == "int64" || t == "int" {
						value = fmt.Sprint(GetFieldInt(item, n))
					}

					if t == "float64" {
						value = fmt.Sprint(GetFieldFloat(item, n))
					}

					if t == "bool" {
						value = fmt.Sprint(GetFieldBool(item, n))
					}

					data = append(data, value)
				}

				fmt.Println(data)
				csvwriter.Write(data)
				csvwriter.Flush()
				bar.Add(1)
				// time.Sleep(1 * time.Millisecond)
			}

		}
	}
}

func GetFieldString[T any](v *T, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return string(f.String())
}

func GetFieldStringP[T any](v *T, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)

	if f.IsNil() {
		return ""
	}

	str := f.Interface().(*string)

	return *str
}

func GetFieldInt[T any](v *T, field string) int {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return int(f.Int())
}

func GetFieldInt64P[T any](v *T, field string) *int64 {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.Interface().(*int64)
}

func GetFieldFloat[T any](v *T, field string) float64 {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return float64(f.Float())
}

func GetFieldFloatP[T any](v *T, field string) *float64 {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)

	return f.Interface().(*float64)
}

func GetFieldBool[T any](v *T, field string) bool {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return bool(f.Bool())
}

func SetFieldString[T any](v *T, field string, value string) {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	f.SetString(value)
}

func GinStreamFromChannel(c *gin.Context, chanStream chan []byte) {
	rc := http.NewResponseController(c.Writer)
	rc.SetWriteDeadline(time.Time{})

	c.Header("Content-Type", "application/x-yaml")
	c.Header("Connection", "Keep-Alive")
	c.Header("Transfer-Encoding", "chunked")
	c.Header("Content-Disposition", `inline; filename="myfile.txt"`)
	c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

	Stream(c, func(w io.Writer) bool {
		if msg, ok := <-chanStream; ok {
			WriteToStream(c, msg)
			return true
		}
		return false
	})
}
