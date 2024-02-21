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
	"pixelplux.com/fireback/modules/workspaces"

	jsoniter "github.com/json-iterator/go"
)

func entityGeoProvinceFormatter(dto *GeoProvinceEntity, query workspaces.QueryDSL) {
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

func GeoProvinceMockEntity() *GeoProvinceEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &GeoProvinceEntity{
		UniqueId: workspaces.UUID(),
		Name:     &stringHolder,
	}

	return entity
}

func GeoProvinceActionSeeder(query workspaces.QueryDSL, count int) {

	successInsert := 0
	failureInsert := 0

	bar := progressbar.Default(int64(count))

	for i := 1; i <= count; i++ {
		entity := GeoProvinceMockEntity()
		_, err := GeoProvinceActionCreate(entity, query)
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

func (x *GeoProvinceEntity) GetNameTranslated(language string) string {
	if x.Translations != nil && len(x.Translations) > 0 {
		for _, item := range x.Translations {
			if item.LanguageId == language {
				return item.Name
			}
		}
	}

	return ""
}

func GeoProvinceActionSeederInit(query workspaces.QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*GeoProvinceEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &GeoProvinceEntity{

		Name: &tildaRef,
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

func GeoProvinceAssociationCreate(dto *GeoProvinceEntity, query workspaces.QueryDSL) error {

	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func GeoProvinceRelationContentCreate(dto *GeoProvinceEntity, query workspaces.QueryDSL) error {

	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* This is when we edit them directly inside the parent entity, here we would know about it.
**/
func GeoProvinceRelationContentUpdate(dto *GeoProvinceEntity, query workspaces.QueryDSL) error {

	return nil
}

func GeoProvincePolyglotCreateHandler(dto *GeoProvinceEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}

	workspaces.PolyglotCreateHandler(dto, &GeoProvinceEntityPolyglot{}, query)

}

/**
* This will be validating your entity fully. Important note is that, you add validate:* tag
* in your entity, it will automatically work here. For slices inside entity, make sure you add
* extra line of AppendSliceErrors, otherwise they won't be detected
 */
func GeoProvinceValidator(dto *GeoProvinceEntity, isPatch bool) *workspaces.IError {
	err := workspaces.CommonStructValidatorPointer(dto, isPatch)

	return err
}

func GeoProvinceEntityPreSanitize(dto *GeoProvinceEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()

	_ = stripPolicy
	_ = ugcPolicy

}

func GeoProvinceEntityBeforeCreateAppend(dto *GeoProvinceEntity, query workspaces.QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = workspaces.UUID()
	}

	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId

	GeoProvinceRecursiveAddUniqueId(dto, query)
}

func GeoProvinceRecursiveAddUniqueId(dto *GeoProvinceEntity, query workspaces.QueryDSL) {

}

func GeoProvinceActionBatchCreateFn(dtos []*GeoProvinceEntity, query workspaces.QueryDSL) ([]*GeoProvinceEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*GeoProvinceEntity{}
		for _, item := range dtos {
			s, err := GeoProvinceActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)

		}
		return items, nil
	}

	return dtos, nil
}

func GeoProvinceActionCreateFn(dto *GeoProvinceEntity, query workspaces.QueryDSL) (*GeoProvinceEntity, *workspaces.IError) {

	// 1. Validate always
	if iError := GeoProvinceValidator(dto, false); iError != nil {
		return nil, iError
	}

	// 1.5 Sanitize the content coming of the front-end
	GeoProvinceEntityPreSanitize(dto, query)

	// 2. Append the necessary information about user, workspace
	GeoProvinceEntityBeforeCreateAppend(dto, query)

	// 3. Append the necessary translations, even if english
	GeoProvincePolyglotCreateHandler(dto, query)

	// 3.5. Create other entities if we want select from them
	GeoProvinceRelationContentCreate(dto, query)

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
	GeoProvinceAssociationCreate(dto, query)

	// 6. Fire the event into system
	event.MustFire(GEOPROVINCE_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": workspaces.GetTypeString(&GeoProvinceEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})

	return dto, nil
}

func GeoProvinceActionGetOne(query workspaces.QueryDSL) (*GeoProvinceEntity, *workspaces.IError) {
	refl := reflect.ValueOf(&GeoProvinceEntity{})
	item, err := workspaces.GetOneEntity[GeoProvinceEntity](query, refl)

	entityGeoProvinceFormatter(item, query)
	return item, err
}

func GeoProvinceActionQuery(query workspaces.QueryDSL) ([]*GeoProvinceEntity, *workspaces.QueryResultMeta, error) {
	refl := reflect.ValueOf(&GeoProvinceEntity{})
	items, meta, err := workspaces.QueryEntitiesPointer[GeoProvinceEntity](query, refl)

	for _, item := range items {
		entityGeoProvinceFormatter(item, query)
	}

	return items, meta, err
}

func GeoProvinceActionUpdateFn(query workspaces.QueryDSL, fields *GeoProvinceEntity) (*GeoProvinceEntity, *workspaces.IError) {
	if fields == nil {
		return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}

	// 1. Validate always
	if iError := GeoProvinceValidator(fields, true); iError != nil {
		return nil, iError
	}

	GeoProvinceRecursiveAddUniqueId(fields, query)

	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()

		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			_, err := GeoProvinceUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}

		})
		return nil, workspaces.CastToIError(vf)
	} else {
		dbref = query.Tx
		return GeoProvinceUpdateExec(dbref, query, fields)
	}

}

func GeoProvinceUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *GeoProvinceEntity) (*GeoProvinceEntity, *workspaces.IError) {
	uniqueId := fields.UniqueId

	query.TriggerEventName = GEOPROVINCE_EVENT_UPDATED

	GeoProvinceEntityPreSanitize(fields, query)
	var item GeoProvinceEntity
	q := dbref.
		Where(&GeoProvinceEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)

	err := q.UpdateColumns(fields).Error
	if err != nil {

		return nil, workspaces.GormErrorToIError(err)
	}

	query.Tx = dbref
	GeoProvinceRelationContentUpdate(fields, query)

	GeoProvincePolyglotCreateHandler(fields, query)

	// @meta(update has many)

	err = dbref.
		Preload(clause.Associations).
		Where(&GeoProvinceEntity{UniqueId: uniqueId}).
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

func GeoProvinceActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&GeoProvinceEntity{})
	query.ActionRequires = []string{PERM_ROOT_GEOPROVINCE_DELETE}
	return workspaces.RemoveEntity[GeoProvinceEntity](query, refl)
}

func GeoProvinceActionWipeClean(query workspaces.QueryDSL) (int64, error) {

	var err error
	var count int64 = 0

	{
		// Hi
		subCount, subErr := workspaces.WipeCleanEntity[GeoProvinceEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'GeoProvinceEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}

	return count, err
}

func GeoProvinceActionBulkUpdate(
	query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[GeoProvinceEntity]) (
	*workspaces.BulkRecordRequest[GeoProvinceEntity], *workspaces.IError,
) {
	result := []*GeoProvinceEntity{}
	err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := GeoProvinceActionUpdate(query, record)

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

func (x *GeoProvinceEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))

	}
	return ""
}

var GeoProvinceEntityMeta = workspaces.TableMetaData{
	EntityName:    "GeoProvince",
	ExportKey:     "geoProvinces",
	TableNameInDb: "fb_geoProvince_entities",
	EntityObject:  &GeoProvinceEntity{},
	ExportStream:  GeoProvinceActionExportT,
	ImportQuery:   GeoProvinceActionImport,
}

func GeoProvinceActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[GeoProvinceEntity](query, GeoProvinceActionQuery, GeoProvincePreloadRelations)
}

func GeoProvinceActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[GeoProvinceEntity](query, GeoProvinceActionQuery, GeoProvincePreloadRelations)
}

func GeoProvinceActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content GeoProvinceEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}

	json.Unmarshal(cx, &content)

	_, err := GeoProvinceActionCreate(&content, query)

	return err
}
