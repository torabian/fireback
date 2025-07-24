package fireback

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
	"text/template"

	"github.com/gookit/event"
	"github.com/seldonio/goven/sql_adaptor"
	"gorm.io/gorm"
)

type ActionDeleteSignature = func(query QueryDSL) (int64, *IError)

func CreateEntity[T any](dto T) (T, *IError) {
	// Do not forget to create unique key:
	// u := UUID()
	// dto.UniqueId = u.String()
	err := GetDbRef().Create(&dto).Error

	if err != nil {
		return dto, GormErrorToIError(err)
	}

	return dto, nil
}

func GetTypeArray(arr interface{}) reflect.Type {
	return reflect.TypeOf(arr).Elem()
}

// This function guesses the entities within a gorm model.
// It's useful when querying to get all of the changes
func ListGormSubEntities(entity reflect.Value) []string {

	var subEntities []string = []string{}
	v := entity.Elem()

	for j := 0; j < v.NumField(); j++ {

		f := v.Field(j)
		field := v.Type().Field(j)
		n := v.Type().Field(j).Name
		t := f.Type().String()
		kind := f.Kind()

		if strings.Contains(field.Tag.Get("gorm"), "embedded") {
			continue
		}

		if n == "Children" || n == "Parent" || n == "LinkedTo" {
			continue
		}

		if kind != reflect.Struct && kind != reflect.Slice && kind != reflect.Ptr {
			continue
		}

		if t == "fireback.XDateComputed" || t == "XDateComputed" || t == "fireback.XDateMetaData" || t == "XDateMetaData" || t == "*fireback.XDateMetaData" || t == "*XDateMetaData" || t == "*JSON" || t == "*fireback.JSON" || t == "Model" || strings.HasSuffix(n, "ListId") || t == "*float64" || t == "*int64" || t == "*bool" {
			continue
		}

		if t == "fireback.String" || t == "String" {
			continue
		}

		if t == "fireback.Duration" || t == "Duration" {
			continue
		}

		if t == "fireback.Int" || t == "Int" {
			continue
		}

		if t == "fireback.XDateTime" || t == "XDateTime" {
			continue
		}

		if t == "*fireback.XDateTime" || t == "*XDateTime" {
			continue
		}
		if t == "fireback.Int64" || t == "Int64" {
			continue
		}
		if t == "fireback.Int32" || t == "Int32" {
			continue
		}
		if t == "fireback.Float64" || t == "Float64" {
			continue
		}
		if t == "fireback.Float32" || t == "Float32" {
			continue
		}
		if t == "fireback.Bool" || t == "Bool" {
			continue
		}
		if t == "fireback.Byte" || t == "Byte" {
			continue
		}

		if strings.Contains(t, "impl.MessageState") || n == "unknownFields" {
			continue
		}

		if kind == reflect.Slice && n != "Translations" {

			hasSubTranslations := false
			s := f.Type().Elem().Elem()
			for rj := 0; rj < s.NumField(); rj++ {
				n1 := s.Field(rj).Name
				if n1 == "Translations" {
					hasSubTranslations = true
				}
			}

			if hasSubTranslations {
				subEntities = append(subEntities, n+".Translations")
			}
		}

		if strings.Contains(t, "time.Time") || strings.Contains(t, "*string") {
			continue
		}
		if strings.Contains(t, "gorm") {
			continue
		}

		subEntities = append(subEntities, n)
	}

	return subEntities
}

type EntityJsonField struct {
	Name         string            `json:"name"`
	JsonField    string            `json:"jsonField"`
	FirebackType string            `json:"fbType"`
	Type         string            `json:"type"`
	Children     []EntityJsonField `json:"children"`
}

type EntityJsonTree struct {
	Fields []EntityJsonField
}

func ExtractEntityFields(entity reflect.Value) []EntityJsonField {

	data := []EntityJsonField{}
	v := entity.Elem()

	for j := 0; j < v.NumField(); j++ {

		f := v.Field(j)
		n := v.Type().Field(j).Name
		t := f.Type().String()

		kind := f.Kind()
		jsonField := v.Type().Field(j).Tag.Get("json")
		firebackType := v.Type().Field(j).Tag.Get("fbtype")
		field := EntityJsonField{Name: n, Type: t, JsonField: jsonField, FirebackType: firebackType}
		// fmt.Println(f, n, t, kind)
		// fmt.Println(f, n, t, kind, f.Type().Elem(), f.Type().Elem().Kind())

		if kind == reflect.Ptr {
			if f.Type().Elem().Kind() == reflect.Struct {
				// fmt.Println(f, n, t, kind, f.Type().Elem(), f.Type().Elem().Kind())
				// field.Children = ExtractEntityFields(reflect.Indirect(reflect.ValueOf(v.Type().Field(j).Type.Elem())))
				// os.Exit(3)
			}

		}

		if strings.ToUpper(field.Name[0:1]) == field.Name[0:1] {
			data = append(data, field)
		}
	}

	return data
}

func ListTranslatableFields(entity reflect.Value) []string {
	var subEntities []string = []string{}

	v := entity.Elem()
	for j := 0; j < v.NumField(); j++ {

		n := v.Type().Field(j).Name
		tag := v.Type().Field(j).Tag.Get("translate")

		if tag == "true" {
			subEntities = append(subEntities, n)
		}
	}

	return subEntities
}

/**
*	This function does not enforce the user internal level query, use with caution,
*	Do not expose this directly to the public
**/
func UnsafeQuerySqlStatement[T any](sql string, values ...interface{}) ([]*T, error) {

	var items []*T

	err := GetDbRef().Raw(sql, values...).Scan(&items).Error

	if err != nil {
		return items, err
	}

	return items, nil

}

type CommonCountSqlResult struct {
	TotalItems int64 `gorm:"totalItems"`
}

func UnsafeQuerySqlFromFs[T any](fsRef *embed.FS, queryName string, query QueryDSL, values ...interface{}) ([]*T, *QueryResultMeta, error) {
	qrm := &QueryResultMeta{
		TotalItems:          -1,
		TotalAvailableItems: -1,
	}

	sqlQuery, err := ReadEmbedFileContent(fsRef, queryName+".sql")

	if err != nil {
		return nil, qrm, GormErrorToIError(err)
	}

	sqlQueryCounter, counterError := ReadEmbedFileContent(fsRef, queryName+"Counter.sql")

	if counterError != nil {

		return nil, qrm, GormErrorToIError(counterError)

	}

	return UnsafeQuerySql[T](sqlQuery, sqlQueryCounter, query, "", values...)
}

type VSqlContext struct {
	IsMysql   bool
	IsSqlite  bool
	IsCounter bool
}

func ContextAwareVSqlOperation[T any](refl reflect.Value, fsRef *embed.FS, queryName string, query QueryDSL, values ...interface{}) ([]*T, *QueryResultMeta, error) {
	extraCondition := ""
	qrm := &QueryResultMeta{
		TotalItems:          -1,
		TotalAvailableItems: -1,
	}

	if query.Query != "" {
		queryAdaptor := sql_adaptor.NewDefaultAdaptorFromStruct(refl)
		parsedQuery, dslError := queryAdaptor.Parse(query.Query)
		if dslError == nil {
			extraCondition = RealEscape(parsedQuery.Raw, parsedQuery.Values...)
		}
	}

	content, err := ReadEmbedFileContent(fsRef, queryName)

	if err != nil {
		return nil, qrm, GormErrorToIError(err)
	}

	sqlQuery := ""
	sqlQueryCounter := ""

	vendor := config.DbVendor
	{
		ctx := VSqlContext{
			IsCounter: false,
		}

		if vendor == "mysql" {
			ctx.IsMysql = true
		}

		if vendor == "sqlite" {
			ctx.IsSqlite = true
		}

		var output bytes.Buffer
		tmpl, err := template.New("example").Parse(content)
		if err != nil {
			return nil, nil, err
		}
		err = tmpl.Execute(&output, ctx)

		if err != nil {
			fmt.Println("Error executing template:", err)
			return nil, nil, err
		}

		sqlQuery = output.String()
	}

	// If the is counter is provided we will check, if not we will not add it.
	if strings.Contains(content, ".IsCounter") {
		ctx := VSqlContext{
			IsCounter: true,
		}

		if vendor == "mysql" {
			ctx.IsMysql = true
		}

		if vendor == "sqlite" {
			ctx.IsSqlite = true
		}

		var output bytes.Buffer
		tmpl, err := template.New("example").Parse(content)
		if err != nil {
			return nil, nil, err
		}
		err = tmpl.Execute(&output, ctx)

		if err != nil {
			fmt.Println("Error executing template:", err)
			return nil, nil, err
		}

		sqlQueryCounter = output.String()

	}

	return UnsafeQuerySql[T](sqlQuery, sqlQueryCounter, query, extraCondition, values...)
}

func UnsafeQuerySql[T any](sqlQuery string, sqlQueryCounter string, query QueryDSL, extraCondition string, values ...interface{}) ([]*T, *QueryResultMeta, error) {
	qrm := &QueryResultMeta{
		TotalItems:          -1,
		TotalAvailableItems: -1,
	}

	sqlCondition := query.InternalQuery
	if sqlCondition == "" {
		sqlCondition = "1"
	}

	if extraCondition == "" {
		extraCondition = "1"
	}

	sqlQuery = strings.ReplaceAll(sqlQuery, "@internalCondition", sqlCondition)
	sqlQuery = strings.ReplaceAll(sqlQuery, "(internalCondition)", " and ("+sqlCondition+")")
	sqlQuery = strings.ReplaceAll(sqlQuery, "(extraCondition)", " and ("+extraCondition+")")
	sqlQuery = strings.ReplaceAll(sqlQuery, "(language)", query.Language)
	sqlQuery = strings.ReplaceAll(sqlQuery, "(workspaceId)", query.WorkspaceId)
	sqlQuery = strings.ReplaceAll(sqlQuery, "(userId)", query.UserId)
	sqlQuery = strings.ReplaceAll(sqlQuery, "(id)", query.UniqueId)
	sqlQuery = strings.ReplaceAll(sqlQuery, "@id", query.UniqueId)
	sqlQuery = strings.ReplaceAll(sqlQuery, "@offset", fmt.Sprintf("%v", query.StartIndex))
	sqlQuery = strings.ReplaceAll(sqlQuery, "(offset)", fmt.Sprintf("%v", query.StartIndex))
	sqlQuery = strings.ReplaceAll(sqlQuery, "(limit)", fmt.Sprintf("%v", query.ItemsPerPage))
	sqlQuery = strings.ReplaceAll(sqlQuery, "@limit", fmt.Sprintf("%v", query.ItemsPerPage))

	if sqlQueryCounter != "" {

		resultCount, err := UnsafeQuerySqlStatement[CommonCountSqlResult](sqlQueryCounter)

		if err != nil {
			return nil, qrm, GormErrorToIError(err)
		}

		if len(resultCount) > 0 {
			qrm.TotalItems = resultCount[0].TotalItems

		}

	}

	result, err := UnsafeQuerySqlStatement[T](sqlQuery, values...)

	if err != nil {
		return nil, qrm, GormErrorToIError(err)
	}

	return result, qrm, err
}

// Returns the connection to database, if it has a transaction
// then returns that transaction one
func GetRef(query QueryDSL) *gorm.DB {
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = GetDbRef()
	} else {
		dbref = query.Tx
	}

	return dbref
}

type CursorInfo struct {
	Field    string
	Value    string
	Operator string
}

func parseCursor(cursor string) CursorInfo {
	re := regexp.MustCompile(`(\w+)\((\d+)\)`)

	matches := re.FindStringSubmatch(cursor)
	if len(matches) > 0 {
		return CursorInfo{
			Field: matches[1],
			Value: matches[2],

			// This needs to be based on the sorting I believe
			Operator: ">",
		}
	}

	return CursorInfo{}
}

func QueryEntitiesPointer[T any](query QueryDSL, refl reflect.Value) ([]*T, *QueryResultMeta, error) {
	ref := GetRef(query)
	var items []*T
	var item *T
	var count int64 = 0

	q := ref.
		Offset(query.StartIndex).
		Limit(query.ItemsPerPage)

	if query.Cursor != nil {
		info := parseCursor(*query.Cursor)
		if info.Field != "" {
			q = q.Where(fmt.Sprintf("%v %v %v", info.Field, info.Operator, info.Value))
		}
	}

	if os.Getenv("DISABLE_FIREBACK_DATA_MANAGEMENT") != "true" {

		// We do not want to show the workspce system anywhere, but we want data belongs to it everywhere
		q = q.Where("unique_id <> \"system\"")

		if query.ResolveStrategy == ResolveStrategyUser {
			q = q.Where(`user_id = "` + query.UserId + `"`)
		} else if query.WorkspaceId != "" {
			q = q.Where(`workspace_id = "` + query.WorkspaceId + `" or workspace_id = "system"`)
		}

	}

	if len(query.SelectableColumnSql) > 0 {
		selectall := []string(query.SelectableColumnSql)

		// These are default field which are needed.
		if !Contains(selectall, "uniqueId") {
			selectall = append(selectall, "unique_id")
		}

		if !Contains(selectall, "linkerId") {
			selectall = append(selectall, "linker_id")
		}

		if !Contains(selectall, "ID") {
			selectall = append(selectall, "id")
		}

		for index, w := range selectall {
			selectall[index] = ToSnakeCase(w)
		}

		q = q.Select(selectall)
	}

	q = q.Where(query.InternalQuery).
		Order(ToSnakeCase(query.Sort))

	countQ := ref

	if os.Getenv("DISABLE_FIREBACK_DATA_MANAGEMENT") != "true" {
		// Counter query should not have the limit, and offset, only the where condition is enough
		// We do not want to show the workspce system anywhere, but we want data belongs to it everywhere
		countQ = countQ.Where("unique_id <> \"system\"")
	}

	countQ = countQ.Where(query.InternalQuery).Model(item)

	// Total availble means all records, which user could possiblty see,
	// But the Query (filters, search) won't affect them.
	// countQ shows total options considering those filters
	var countTotalAvailable int64 = 0

	v := ref.Where(query.InternalQuery)

	if query.ResolveStrategy == ResolveStrategyUser {
		q = q.Where(`user_id = "` + query.UserId + `"`)
	} else if query.WorkspaceId != "" {
		q = q.Where(`workspace_id = "` + query.WorkspaceId + `" or workspace_id = "system"`)
	}

	v.Model(item).Count(&countTotalAvailable)

	if query.Deep {
		preloads := ListGormSubEntities(refl)

		for _, f := range preloads {
			q = q.Preload(f)
		}

		if len(query.Preloads) > 0 {
			for _, f := range preloads {
				q = q.Preload(f)
			}
		}

		if len(query.WithPreloads) > 0 {
			for _, f := range query.WithPreloads {
				q = q.Preload(f)
			}
		}
	}

	if query.Query != "" {
		queryAdaptor := sql_adaptor.NewDefaultAdaptorFromStruct(refl)
		parsedQuery, dslError := queryAdaptor.Parse(query.Query)
		if dslError == nil {
			countQ = countQ.Where(parsedQuery.Raw, sql_adaptor.StringSliceToInterfaceSlice(parsedQuery.Values)...)
			q = q.Where(parsedQuery.Raw, sql_adaptor.StringSliceToInterfaceSlice(parsedQuery.Values)...)
		} else {
			fmt.Println("Query error", dslError)
		}
	}

	countQ.Count(&count)
	q = q.Order(ToSnakeCase(query.Sort)).Find(&items)
	err := q.Error

	var cursor *string

	if len(items) > 0 {
		v2 := reflect.ValueOf(items[len(items)-1])
		idField := v2.Elem().FieldByName("ID")
		if idField.IsValid() {
			val := idField.Uint()
			cur := fmt.Sprintf("id(%v)", val)
			if query.Sort != "" {
				cur += fmt.Sprintf("+sort(%v)", query.Sort)
			}
			cursor = &cur
		}
	}

	meta := &QueryResultMeta{
		TotalItems:          count,
		TotalAvailableItems: countTotalAvailable,
		Cursor:              cursor,
	}

	if err != nil {
		return items, meta, err
	}

	return items, meta, nil

}

func GetOneEntity[T any](query QueryDSL, reflectVal reflect.Value) (*T, *IError) {

	var item T

	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = GetDbRef()
	} else {
		dbref = query.Tx
	}
	dbref = dbref
	preloads := ListGormSubEntities(reflectVal)

	for _, f := range preloads {
		if f != "" {
			dbref = dbref.Preload(f)

		}
	}

	if len(query.WithPreloads) > 0 {
		for _, f := range query.WithPreloads {
			dbref = dbref.Preload(f)
		}
	}

	err := dbref.Where(RealEscape("unique_id = ?", query.UniqueId)).First(&item).Error

	if err != nil {
		return &item, GormErrorToIError(err)
	}

	return &item, nil
}

func GetOneByWorkspaceEntity[T any](query QueryDSL, reflectVal reflect.Value) (*T, *IError) {

	var item T

	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = GetDbRef()
	} else {
		dbref = query.Tx
	}
	preloads := ListGormSubEntities(reflectVal)

	for _, f := range preloads {
		if f != "" {
			dbref = dbref.Preload(f)

		}
	}

	if len(query.WithPreloads) > 0 {
		for _, f := range query.WithPreloads {
			dbref = dbref.Preload(f)
		}
	}

	err := dbref.Where(RealEscape("workspace_id = ?", query.WorkspaceId)).First(&item).Error

	if err != nil {
		return nil, GormErrorToIError(err)
	}

	return &item, nil
}

func RealEscape(portion string, values ...string) string {

	for _, item := range values {
		escapedItem := `"` + strings.ReplaceAll(item, "\"", "\\\"") + `"`
		portion = strings.Replace(portion, "?", escapedItem, 1)
	}

	return portion
}

func GetOneEntityByWorkspace[T any](query QueryDSL, reflectVal reflect.Value) (*T, *IError) {

	var item T

	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = GetDbRef()
	} else {
		dbref = query.Tx
	}
	preloads := ListGormSubEntities(reflectVal)

	for _, f := range preloads {
		if f != "" {
			dbref = dbref.Preload(f)

		}
	}

	if len(query.WithPreloads) > 0 {
		for _, f := range query.WithPreloads {
			dbref = dbref.Preload(f)
		}
	}

	err := dbref.Where(RealEscape("workspace_id = ?", query.WorkspaceId)).First(&item).Error

	if err != nil {
		return &item, GormErrorToIError(err)
	}

	return &item, nil
}

func UpdateEntity[T any](query QueryDSL, fields *T) (*T, *IError) {

	var item T
	err := GetDbRef().Where(RealEscape("unique_id = ?", GetFieldString(fields, "UniqueId"))).First(&item).UpdateColumns(fields).Error
	if err != nil {
		return &item, GormErrorToIError(err)
	}

	return &item, nil
}

func AppendQueryDslWhereToDb(db *gorm.DB, query QueryDSL, reflectVal reflect.Value) *gorm.DB {
	queryAdaptor := sql_adaptor.NewDefaultAdaptorFromStruct(reflectVal)
	parsedQuery, dslError := queryAdaptor.Parse(query.Query)

	if dslError == nil {
		db = db.Where(parsedQuery.Raw, sql_adaptor.StringSliceToInterfaceSlice(parsedQuery.Values)...)
	}
	return db
}

func RemoveEntity[T any](query QueryDSL, reflectVal reflect.Value) (int64, *IError) {

	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = GetDbRef()
	} else {
		dbref = query.Tx
	}

	var dto T
	operation := AppendQueryDslWhereToDb(dbref, query, reflectVal).Delete(&dto)

	if operation.Error != nil {

		return 0, GormErrorToIError(operation.Error)
	} else {

		if query.TriggerEventName != "" {
			event.MustFire(query.TriggerEventName, event.M{
				"entity":   dto,
				"target":   "workspace",
				"unqiueId": query.WorkspaceId,
			})
		}

		return operation.RowsAffected, nil
	}
}

/**
* Do not expose this on public apis. Dangerous usage might happen.
**/
func WipeCleanEntity[T any]() (int64, error) {

	// Wipe the main entities
	var item T
	operation := GetDbRef().Where("unique_id <> \"\"").Delete(&item)
	if operation.Error != nil {
		return 0, operation.Error
	} else {
		return operation.RowsAffected, nil
	}
}
