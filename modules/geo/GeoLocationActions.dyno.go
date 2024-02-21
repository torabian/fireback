package geo

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	reflect "reflect"
	"strings"

	"github.com/gookit/event"
	"github.com/microcosm-cc/bluemonday"
	"github.com/schollz/progressbar/v3"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	jsoniter "github.com/json-iterator/go"

	queries "pixelplux.com/fireback/modules/geo/queries"
	"pixelplux.com/fireback/modules/workspaces"
)

func entityGeoLocationFormatter(dto *GeoLocationEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}

	if dto.Created > 0 {
		dto.CreatedFormatted = workspaces.FormatDateBasedOnQuery(dto.Created, query)
	}

	if dto.Updated > 0 {
		dto.CreatedFormatted = workspaces.FormatDateBasedOnQuery(dto.Updated, query)
	}
}

func GeoLocationMockEntity() *GeoLocationEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &GeoLocationEntity{
		UniqueId:     workspaces.UUID(),
		Name:         &stringHolder,
		Code:         &stringHolder,
		Status:       &stringHolder,
		Flag:         &stringHolder,
		OfficialName: &stringHolder,
	}

	return entity
}

func GeoLocationActionSeeder(query workspaces.QueryDSL, count int) {

	successInsert := 0
	failureInsert := 0

	bar := progressbar.Default(int64(count))

	for i := 1; i <= count; i++ {
		entity := GeoLocationMockEntity()
		_, err := GeoLocationActionCreate(entity, query)
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

/**
* For translation in proto
 */

func (x *GeoLocationEntity) GetNameTranslated(language string) string {
	if x.Translations != nil && len(x.Translations) > 0 {
		for _, item := range x.Translations {
			if item.LanguageId == language {
				return item.Name
			}
		}
	}

	return ""
}

func (x *GeoLocationEntity) GetOfficialNameTranslated(language string) string {
	if x.Translations != nil && len(x.Translations) > 0 {
		for _, item := range x.Translations {
			if item.LanguageId == language {
				return item.OfficialName
			}
		}
	}

	return ""
}

func GeoLocationActionSeederInit(query workspaces.QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*GeoLocationEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &GeoLocationEntity{

		Name: &tildaRef,

		Code: &tildaRef,

		Status: &tildaRef,

		Flag: &tildaRef,

		OfficialName: &tildaRef,
	}

	data = append(data, entity)

	if format == "yml" || format == "yaml" {
		body, err = yaml.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
	}
	if format == "json" {
		body, err = json.MarshalIndent(data, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		file = strings.Replace(file, ".yml", ".json", -1)
	}

	os.WriteFile(file, body, 0644)
}

func GeoLocationAssociationCreate(dto *GeoLocationEntity, query workspaces.QueryDSL) error {

	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func GeoLocationRelationContentCreate(dto *GeoLocationEntity, query workspaces.QueryDSL) error {

	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* This is when we edit them directly inside the parent entity, here we would know about it.
**/
func GeoLocationRelationContentUpdate(dto *GeoLocationEntity, query workspaces.QueryDSL) error {

	return nil
}

func GeoLocationPolyglotCreateHandler(dto *GeoLocationEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}

	workspaces.PolyglotCreateHandler(dto, &GeoLocationEntityPolyglot{}, query)

}

/**
* This will be validating your entity fully. Important note is that, you add validate:* tag
* in your entity, it will automatically work here. For slices inside entity, make sure you add
* extra line of AppendSliceErrors, otherwise they won't be detected
 */
func GeoLocationValidator(dto *GeoLocationEntity, isPatch bool) *workspaces.IError {
	err := workspaces.CommonStructValidatorPointer(dto, isPatch)

	return err
}

func GeoLocationEntityPreSanitize(dto *GeoLocationEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()

	_ = stripPolicy
	_ = ugcPolicy

}

func GeoLocationEntityBeforeCreateAppend(dto *GeoLocationEntity, query workspaces.QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = workspaces.UUID()
	}

	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId

	GeoLocationRecursiveAddUniqueId(dto, query)
}

func GeoLocationRecursiveAddUniqueId(dto *GeoLocationEntity, query workspaces.QueryDSL) {

}

func GeoLocationActionBatchCreateFn(dtos []*GeoLocationEntity, query workspaces.QueryDSL) ([]*GeoLocationEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*GeoLocationEntity{}
		for _, item := range dtos {
			s, err := GeoLocationActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)

		}
		return items, nil
	}

	return dtos, nil
}

func GeoLocationActionCreateFn(dto *GeoLocationEntity, query workspaces.QueryDSL) (*GeoLocationEntity, *workspaces.IError) {

	// 1. Validate always
	if iError := GeoLocationValidator(dto, false); iError != nil {
		return nil, iError
	}

	// 1.5 Sanitize the content coming of the front-end
	GeoLocationEntityPreSanitize(dto, query)

	// 2. Append the necessary information about user, workspace
	GeoLocationEntityBeforeCreateAppend(dto, query)

	// 3. Append the necessary translations, even if english
	GeoLocationPolyglotCreateHandler(dto, query)

	// 3.5. Create other entities if we want select from them
	GeoLocationRelationContentCreate(dto, query)

	// 4. Create the entity
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()
	} else {
		dbref = query.Tx
	}

	query.Tx = dbref
	err := dbref.Create(&dto).Error
	if err != nil {
		err := workspaces.GormErrorToIError(err)
		return dto, err
	}

	// 5. Create sub entities, objects or arrays, association to other entities
	GeoLocationAssociationCreate(dto, query)

	// 6. Fire the event into system
	event.MustFire(GEOLOCATION_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": workspaces.GetTypeString(&GeoLocationEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})

	return dto, nil
}

func GeoLocationActionGetOne(query workspaces.QueryDSL) (*GeoLocationEntity, *workspaces.IError) {
	refl := reflect.ValueOf(&GeoLocationEntity{})
	item, err := workspaces.GetOneEntity[GeoLocationEntity](query, refl)

	entityGeoLocationFormatter(item, query)
	return item, err
}

func GeoLocationActionQuery(query workspaces.QueryDSL) ([]*GeoLocationEntity, *workspaces.QueryResultMeta, error) {
	refl := reflect.ValueOf(&GeoLocationEntity{})
	items, meta, err := workspaces.QueryEntitiesPointer[GeoLocationEntity](query, refl)

	for _, item := range items {
		entityGeoLocationFormatter(item, query)
	}

	return items, meta, err
}

func (dto *GeoLocationEntity) Size() int {
	var size int = len(dto.Children)
	for _, c := range dto.Children {
		size += c.Size()
	}
	return size
}

func (dto *GeoLocationEntity) Add(nodes ...*GeoLocationEntity) bool {
	var size = dto.Size()
	for _, n := range nodes {
		if n.ParentId != nil && *n.ParentId == dto.UniqueId {
			dto.Children = append(dto.Children, n)
		} else {
			for _, c := range dto.Children {
				if c.Add(n) {
					break
				}
			}
		}
	}
	return dto.Size() == size+len(nodes)
}

func GeoLocationActionCommonPivotQuery(query workspaces.QueryDSL) ([]*workspaces.PivotResult, *workspaces.QueryResultMeta, error) {

	items, meta, err := workspaces.UnsafeQuerySqlFromFs[workspaces.PivotResult](
		&queries.QueriesFs, "GeoLocationCommonPivot.sqlite.dyno", query,
	)

	return items, meta, err
}

func GeoLocationActionCteQuery(query workspaces.QueryDSL) ([]*GeoLocationEntity, *workspaces.QueryResultMeta, error) {

	items, meta, err := workspaces.UnsafeQuerySqlFromFs[GeoLocationEntity](
		&queries.QueriesFs, "GeoLocationCTE.sqlite.dyno", query,
	)

	for _, item := range items {
		entityGeoLocationFormatter(item, query)
	}

	var tree []*GeoLocationEntity

	for _, item := range items {
		if item.ParentId == nil {
			root := item
			root.Add(items...)
			tree = append(tree, root)
		}
	}

	return tree, meta, err
}

func GeoLocationActionUpdateFn(query workspaces.QueryDSL, fields *GeoLocationEntity) (*GeoLocationEntity, *workspaces.IError) {
	if fields == nil {
		return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}

	// 1. Validate always
	if iError := GeoLocationValidator(fields, true); iError != nil {
		return nil, iError
	}

	GeoLocationRecursiveAddUniqueId(fields, query)

	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()

		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			_, err := GeoLocationUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}

		})
		return nil, workspaces.CastToIError(vf)
	} else {
		dbref = query.Tx
		return GeoLocationUpdateExec(dbref, query, fields)
	}

}

func GeoLocationUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *GeoLocationEntity) (*GeoLocationEntity, *workspaces.IError) {
	uniqueId := fields.UniqueId

	query.TriggerEventName = GEOLOCATION_EVENT_UPDATED

	GeoLocationEntityPreSanitize(fields, query)
	var item GeoLocationEntity
	q := dbref.
		Where(&GeoLocationEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)

	err := q.UpdateColumns(fields).Error
	if err != nil {

		return nil, workspaces.GormErrorToIError(err)
	}

	query.Tx = dbref
	GeoLocationRelationContentUpdate(fields, query)

	GeoLocationPolyglotCreateHandler(fields, query)

	// @meta(update has many)

	err = dbref.
		Preload(clause.Associations).
		Where(&GeoLocationEntity{UniqueId: uniqueId}).
		First(&item).Error

	event.MustFire(query.TriggerEventName, event.M{
		"entity":   &item,
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})

	if err != nil {
		return &item, workspaces.GormErrorToIError(err)
	}

	return &item, nil
}

func GeoLocationActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&GeoLocationEntity{})
	query.ActionRequires = []string{PERM_ROOT_GEOLOCATION_DELETE}
	return workspaces.RemoveEntity[GeoLocationEntity](query, refl)
}

func GeoLocationActionWipeClean(query workspaces.QueryDSL) (int64, error) {

	var err error
	var count int64 = 0

	{
		// Hi
		subCount, subErr := workspaces.WipeCleanEntity[GeoLocationEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'GeoLocationEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}

	return count, err
}

func GeoLocationActionBulkUpdate(
	query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[GeoLocationEntity]) (
	*workspaces.BulkRecordRequest[GeoLocationEntity], *workspaces.IError,
) {
	result := []*GeoLocationEntity{}
	err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := GeoLocationActionUpdate(query, record)

			if err != nil {
				return err
			} else {
				result = append(result, item)
			}

		}

		return nil
	})

	if err == nil {
		return dto, nil
	}

	return nil, err.(*workspaces.IError)
}

func (x *GeoLocationEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))

	}
	return ""
}

var GeoLocationEntityMeta = workspaces.TableMetaData{
	EntityName:    "GeoLocation",
	ExportKey:     "geoLocations",
	TableNameInDb: "fb_geoLocation_entities",
	EntityObject:  &GeoLocationEntity{},
	ExportStream:  GeoLocationActionExportT,
	ImportQuery:   GeoLocationActionImport,
}

func GeoLocationActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[GeoLocationEntity](query, GeoLocationActionQuery, GeoLocationPreloadRelations)
}

func GeoLocationActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[GeoLocationEntity](query, GeoLocationActionQuery, GeoLocationPreloadRelations)
}

func GeoLocationActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content GeoLocationEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}

	json.Unmarshal(cx, &content)

	_, err := GeoLocationActionCreate(&content, query)

	return err
}
