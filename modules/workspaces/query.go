package workspaces

import (
	"bufio"
	"bytes"
	"embed"
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

func CliAuth(security *SecurityModel) (*AuthResultDto, *IError) {
	context := &AuthContextDto{
		WorkspaceId:  config.CliWorkspace,
		Token:        config.CliToken,
		Capabilities: []PermissionInfo{},
		Security:     security,
	}

	return WithAuthorizationPure(context)
}

func CommonCliQueryDSLBuilderAuthorize(c *cli.Context, security *SecurityModel) QueryDSL {
	q := CommonCliQueryDSLBuilder(c)

	if security != nil && security.ResolveStrategy != ResolveStrategyPublic {
		result, err := CliAuth(security)
		if err != nil {

			if err.ToPublicEndUser(&q).Message != err.ToPublicEndUser(&q).MessageTranslated {
				log.Fatalf("%s", err.ToPublicEndUser(&q).Message)
			}
			log.Default().Printf("%s", err.ToPublicEndUser(&q).MessageTranslated)
		}

		q.ResolveStrategy = security.ResolveStrategy
		q.InternalQuery = result.SqlContext
		q.UserId = result.User.UniqueId
		q.UserAccessPerWorkspace = result.UserAccessPerWorkspace

	}

	return q
}

func CommonCliQueryDSLBuilder(c *cli.Context) QueryDSL {

	queryString := c.String("query")
	startIndex := c.Int("offset")
	itemsPerPage := c.Int("limit")

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
	workspaceId := config.CliWorkspace

	if config.CliLanguage != "" {
		lang = config.CliLanguage
	}

	if config.CliRegion != "" {
		region = config.CliRegion
	}

	withPreloads := c.String("wp")

	var f QueryDSL = QueryDSL{
		Query:        queryString,
		StartIndex:   startIndex,
		WorkspaceId:  workspaceId,
		Language:     lang,
		Region:       strings.ToUpper(region),
		ItemsPerPage: itemsPerPage,
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
	if c.IsSet("sort") {
		f.Sort = c.String("sort")
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
func CommonCliQueryCmd2[T any](
	c *cli.Context,
	fn func(query QueryDSL) ([]T, *QueryResultMeta, error),
	security *SecurityModel,
) {

	f := CommonCliQueryDSLBuilderAuthorize(c, security)

	if items, count, err := fn(f); err != nil {
		fmt.Println(12)
		fmt.Println(err)
	} else {
		out := gin.H{
			"data": gin.H{
				"startIndex":   f.StartIndex,
				"itemsPerPage": f.ItemsPerPage,
				"items":        items,
				"totalItems":   count.TotalItems,
			},
		}
		if c.Bool("yaml") {
			body, err2 := yaml.Marshal(out)
			if err2 != nil {
				log.Fatal(err2)
			}
			fmt.Println(string(body))
		} else {
			jsonString, _ := json.MarshalIndent(out, "", "  ")

			fmt.Println(string(jsonString))
		}

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

func excludeDefaultFields(items []string) []string {
	excluded := []string{}
	toExclude := []string{
		"CreatedFormatted",
		"UpdatedFormatted",
		"LinkedTo",
		"Children",
		"Rank",
		"IsDeletable",
		"IsUpdatable",
	}

	for _, item := range items {
		if !Contains(toExclude, item) {
			excluded = append(excluded, item)
		}
	}

	fmt.Println("Excluded:", excluded)

	return excluded
}

func CommonCliTableCmd2[T any](
	c *cli.Context,
	fn func(query QueryDSL) ([]*T, *QueryResultMeta, error),
	security *SecurityModel,
	v reflect.Value,
) {

	verbose := false
	if c.IsSet("verbose") && c.Bool("verbose") {
		verbose = true
	}

	f := CommonCliQueryDSLBuilderAuthorize(c, security)
	items, _, err := fn(f)

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

	heads := GetColumnsFromReflect[T](v)

	for _, n := range heads {
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

		tds := ExtractRowStringValues[T](row, v, verbose)

		for _, cellValue := range tds {

			r = append(r, &simpletable.Cell{
				Align: simpletable.AlignRight, Text: cellValue,
			})
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleDefault)
	fmt.Println(table.String())
}

func CommonCliImportCmdAuthorized[T any](
	c *cli.Context,
	fn func(dto *T, query QueryDSL) (*T, *IError),
	v reflect.Value,
	importFilePath string,
	security *SecurityModel,
	initializer func() T,
) {

	f := CommonCliQueryDSLBuilderAuthorize(c, security)
	f.Deep = true

	// fmt.Println(72, f.UniqueId, f.WorkspaceId, f.UserId, f.UserHas)

	if strings.Contains(importFilePath, ".yml") || strings.Contains(importFilePath, ".yaml") {
		importYamlFromFileOnDisk(importFilePath, fn, f)
	}

	if strings.Contains(importFilePath, ".csv") {
		importCsvFromFileReader(importFilePath, fn, f, initializer)
	}

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

	// if strings.Contains(importFilePath, ".csv") {
	// 	importCsvFromFileReader(importFilePath, fn, f)
	// }

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

func CommonCliImportEmbedBatchCmd[T any](
	c *cli.Context,
	fn func(dto []*T, query QueryDSL) ([]*T, *IError),
	v reflect.Value,
	fsRef *embed.FS,
) {
	f := CommonCliQueryDSLBuilder(c)
	f.WorkspaceId = "system"
	SeederFromFSImportBatch(f, fn, v, fsRef, []string{}, false)
}

func SeederFromFSImport[T any](
	f QueryDSL,
	fn func(dto *T, query QueryDSL) (*T, *IError),
	v reflect.Value,
	fsRef *embed.FS,
	fileNames []string,
	silent bool,
) {

	if fsRef == nil {
		return
	}

	f.Deep = true

	if entity, err := GetSeederFilenames(fsRef, ""); err != nil {
		log.Fatalln(err.Error())
	} else {

		for _, path := range entity {
			if len(fileNames) > 0 && !Contains(fileNames, path) {
				continue
			}

			fmt.Println("Importing file:", path)

			if strings.Contains(path, ".yml") || strings.Contains(path, ".yaml") {
				importYamlFromFileEmbed(fsRef, path, fn, f, silent)
			}

			if strings.Contains(path, ".csv") {
				importCsvFromEmbed(fsRef, path, fn, f)
			}
		}

	}

}

func SeederFromFSImportBatch[T any](
	f QueryDSL,
	fn func(dto []*T, query QueryDSL) ([]*T, *IError),
	v reflect.Value,
	fsRef *embed.FS,
	fileNames []string,
	silent bool,
) {

	if fsRef == nil {
		return
	}

	f.Deep = true

	if entity, err := GetSeederFilenames(fsRef, ""); err != nil {
		log.Fatalln(err.Error())
	} else {

		for _, path := range entity {
			if len(fileNames) > 0 && !Contains(fileNames, path) {
				continue
			}

			if strings.Contains(path, ".yml") || strings.Contains(path, ".yaml") {
				importYamlFromFileEmbedBatch(fsRef, path, fn, f, silent)
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

	defer writer.Close()

	translationBox := map[string]interface{}{}
	ReadYamlFileEmbed[map[string]interface{}](translationRef, fsFileName, &translationBox)

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
		FieldsMap:   map[string]string{},
	}

	if strings.Contains(exportFilePath, ".pdf") {
		PdfExporter[T](exportFilePath, f, fn, v, bar, data)
	}
}

func CommonCliExportCmd2[T any](
	c *cli.Context,
	fn func(q QueryDSL) (chan []*T, *QueryResultMeta, error),
	v reflect.Value,
	exportFilePath string,
	translationRef *embed.FS,
	fsFileName string,
	detectedPreloads []string,
) error {

	f := CommonCliQueryDSLBuilder(c)
	f.Deep = true
	f.WithPreloads = append(f.WithPreloads, detectedPreloads...)

	stream, count, err := fn(f)
	bar := progressbar.Default(int64(count.TotalItems))

	if err != nil {
		log.Fatalln(err)
		return err
	}

	translationBox := map[string]interface{}{}
	ReadYamlFileEmbed[map[string]interface{}](translationRef, fsFileName, &translationBox)

	var exporter func(source chan []*T, fp string) (chan ProgressUpdate, error)

	if strings.Contains(exportFilePath, ".csv") {
		exporter = CSV2ExporterWriter
	}

	if strings.Contains(exportFilePath, ".yml") || strings.Contains(exportFilePath, ".yaml") {
		exporter = YamlExporterWriter
	}

	if strings.Contains(exportFilePath, ".json") {
		exporter = JsonExporterWriter
	}

	stats, err := exporter(stream, exportFilePath)
	if err != nil {
		log.Fatalln(err)
	}

	for stat := range stats {
		bar.Add(stat.ItemsProcessed)
	}

	bar.Finish()

	return nil
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

// func SetFieldString[T any](v *T, field string, value string) {
// 	r := reflect.ValueOf(v)
// 	f := reflect.Indirect(r).FieldByName(field)
// 	f.SetString(value)
// }

func GetStructFields(v interface{}) {
	r := reflect.ValueOf(v).Elem()
	// t := r.Type()
	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)
		if field.Kind() == reflect.String {
			// Generate and set random string value
			field.SetString("@@")
		}
	}
}

func PopulateInteractively[T any](entity T, c *cli.Context, flags []CliInteractiveFlag) {
	for _, item := range flags {
		if (!item.Required && !item.Recommended) && !c.Bool("all") {
			continue
		}

		if item.Type == "string" {
			var result string
			if !item.Required {
				result, _, _ = AskForInputOptional(item.Name, "")
			} else {
				result, _, _ = AskForInputOptional(item.Name, "")
			}
			SetField(entity, ToLower(item.StructField), &result)
		}
		if item.Type == "bool" {
			result := AskBoolean(item.Name)
			SetField(entity, ToLower(item.StructField), &result)
		}

	}
}

func SetFieldString[T any](v T, field string, value string) {
	GetStructFields(v)
	r := reflect.ValueOf(v)

	if r.Kind() != reflect.Ptr {
		fmt.Println("Input must be a pointer")
		return
	}

	r = reflect.Indirect(r)
	f := r.FieldByName(field)

	if !f.IsValid() {
		fmt.Printf("Field %s not found\n", field)
		return
	}

	if f.Kind() == reflect.String {
		f.SetString(value)
	} else if f.Kind() == reflect.Ptr && f.Elem().Kind() == reflect.String && f.Elem().CanSet() {
		f.Elem().SetString(value)
	} else {
		fmt.Println(field, "Field is not a string or pointer to string type:", f.Kind())
	}
}

// func SetFieldString[T any](v *T, field string, value string) {
// 	r := reflect.ValueOf(v)
// 	fmt.Println("::", reflect.Indirect(r).FieldByName("Name"))
// 	f := reflect.Indirect(r).FieldByName(field)
// 	f.SetString(value)
// }

// func SetFieldString[T any](v *T, field string, value string) {
// 	r := reflect.ValueOf(v)
// 	if r.Kind() != reflect.Ptr {
// 		fmt.Println("Input must be a pointer")
// 		return
// 	}

// 	r = reflect.Indirect(r)
// 	f := r.FieldByName(field)

// 	if !f.IsValid() {
// 		fmt.Printf("Field %s not found\n", field)
// 		return
// 	}

// 	fmt.Println("0", f.Kind(), f.Elem().Kind())
// 	if f.Kind() == reflect.String {
// 		f.SetString(value)
// 	} else if f.Type().Kind() == reflect.Ptr && f.Elem().Kind() == reflect.String {
// 		f.Elem().SetString(value)
// 	} else {
// 		fmt.Println("Field is not a string or pointer to string type")
// 	}
// }

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
