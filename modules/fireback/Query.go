package fireback

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"regexp"
	"slices"
	"strings"

	"github.com/schollz/progressbar/v3"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"
	"gopkg.in/yaml.v2"
)

func CliAuth(security *SecurityModel) (*AuthResultDto, *IError) {
	context := &AuthContextDto{
		WorkspaceId:  config.CliWorkspace,
		Token:        config.CliToken,
		Capabilities: []application.PermissionInfo{},
		Security:     security,
	}

	return WithAuthorizationPure(context)
}

func CommonCliQueryDSLBuilderAuthorize(c *cli.Command, security *SecurityModel) QueryDSL {
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
		if result.UserId.IsSet() && result.UserId.OrDefault("") != "" {
			q.UserId = result.UserId.OrDefault("")
		}
		q.UserAccessPerWorkspace = result.UserAccessPerWorkspace

	}

	return q
}

func CommonCliQueryDSLBuilder(c *cli.Command) QueryDSL {

	queryString := c.String("query")
	startIndex := c.Int("offset")
	var cursor *string = nil
	if c.IsSet("cursor") {
		val := c.String("cursor")
		cursor = &val
	}

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
		C:            c,
		Cursor:       cursor,
		WorkspaceId:  workspaceId,
		Language:     lang,
		Region:       strings.ToUpper(region),
		ItemsPerPage: itemsPerPage,
	}

	if c.IsSet("x-select") {
		f.SelectableColumn = SmartSplit(c.String("x-select"))
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

var smartSplitRegex = regexp.MustCompile(`[;\s,]+`)

func SmartSplit(input string) []string {
	if strings.TrimSpace(input) == "" {
		return []string{}
	}

	return smartSplitRegex.Split(input, -1)
}

func GetColumnsFromReflect[T any](v reflect.Value) []string {
	verbose := false

	headers := []string{}
	for j := 0; j < v.NumField(); j++ {
		n := v.Type().Field(j).Name

		if strings.ToUpper(n[0:1]) != n[0:1] {
			continue
		}

		if slices.Contains(FIREBACK_DEFAULT_DB_COLUMNS, n) && !verbose {
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
	"ParentId",
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
			if len(fileNames) > 0 && !slices.Contains(fileNames, path) {
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

type ExportCatalog[T any] struct {
	Writer             *os.File
	ReadSize           int64
	TotalItemsToExport int64
	F                  QueryDSL
	ExportFilePath     string
	QueryResultMeta    *QueryResultMeta
	Fn                 func(query QueryDSL) ([]*T, *QueryResultMeta, *IError)
}

func YamlExporterChannel[T any](
	query QueryDSL,
	fn func(query QueryDSL) ([]*T, *QueryResultMeta, *IError),
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
	fn func(query QueryDSL) ([]*T, *QueryResultMeta, *IError),
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
