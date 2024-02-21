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

func entityGeoLocationTypeFormatter(dto *GeoLocationTypeEntity, query workspaces.QueryDSL) {
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

func GeoLocationTypeMockEntity() *GeoLocationTypeEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &GeoLocationTypeEntity{
		UniqueId: workspaces.UUID(),
		Name:     &stringHolder,
	}

	return entity
}

func GeoLocationTypeActionSeeder(query workspaces.QueryDSL, count int) {

	successInsert := 0
	failureInsert := 0

	bar := progressbar.Default(int64(count))

	for i := 1; i <= count; i++ {
		entity := GeoLocationTypeMockEntity()
		_, err := GeoLocationTypeActionCreate(entity, query)
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

func (x *GeoLocationTypeEntity) GetNameTranslated(language string) string {
	if x.Translations != nil && len(x.Translations) > 0 {
		for _, item := range x.Translations {
			if item.LanguageId == language {
				return item.Name
			}
		}
	}

	return ""
}

func GeoLocationTypeActionSeederInit(query workspaces.QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*GeoLocationTypeEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &GeoLocationTypeEntity{

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

func GeoLocationTypeAssociationCreate(dto *GeoLocationTypeEntity, query workspaces.QueryDSL) error {

	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func GeoLocationTypeRelationContentCreate(dto *GeoLocationTypeEntity, query workspaces.QueryDSL) error {

	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* This is when we edit them directly inside the parent entity, here we would know about it.
**/
func GeoLocationTypeRelationContentUpdate(dto *GeoLocationTypeEntity, query workspaces.QueryDSL) error {

	return nil
}

func GeoLocationTypePolyglotCreateHandler(dto *GeoLocationTypeEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}

	workspaces.PolyglotCreateHandler(dto, &GeoLocationTypeEntityPolyglot{}, query)

}

/**
* This will be validating your entity fully. Important note is that, you add validate:* tag
* in your entity, it will automatically work here. For slices inside entity, make sure you add
* extra line of AppendSliceErrors, otherwise they won't be detected
 */
func GeoLocationTypeValidator(dto *GeoLocationTypeEntity, isPatch bool) *workspaces.IError {
	err := workspaces.CommonStructValidatorPointer(dto, isPatch)

	return err
}

func GeoLocationTypeEntityPreSanitize(dto *GeoLocationTypeEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()

	_ = stripPolicy
	_ = ugcPolicy

}

func GeoLocationTypeEntityBeforeCreateAppend(dto *GeoLocationTypeEntity, query workspaces.QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = workspaces.UUID()
	}

	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId

	GeoLocationTypeRecursiveAddUniqueId(dto, query)
}

func GeoLocationTypeRecursiveAddUniqueId(dto *GeoLocationTypeEntity, query workspaces.QueryDSL) {

}

func GeoLocationTypeActionBatchCreateFn(dtos []*GeoLocationTypeEntity, query workspaces.QueryDSL) ([]*GeoLocationTypeEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*GeoLocationTypeEntity{}
		for _, item := range dtos {
			s, err := GeoLocationTypeActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)

		}
		return items, nil
	}

	return dtos, nil
}

func GeoLocationTypeActionCreateFn(dto *GeoLocationTypeEntity, query workspaces.QueryDSL) (*GeoLocationTypeEntity, *workspaces.IError) {

	// 1. Validate always
	if iError := GeoLocationTypeValidator(dto, false); iError != nil {
		return nil, iError
	}

	// 1.5 Sanitize the content coming of the front-end
	GeoLocationTypeEntityPreSanitize(dto, query)

	// 2. Append the necessary information about user, workspace
	GeoLocationTypeEntityBeforeCreateAppend(dto, query)

	// 3. Append the necessary translations, even if english
	GeoLocationTypePolyglotCreateHandler(dto, query)

	// 3.5. Create other entities if we want select from them
	GeoLocationTypeRelationContentCreate(dto, query)

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
	GeoLocationTypeAssociationCreate(dto, query)

	// 6. Fire the event into system
	event.MustFire(GEOLOCATIONTYPE_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": workspaces.GetTypeString(&GeoLocationTypeEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})

	return dto, nil
}

func GeoLocationTypeActionGetOne(query workspaces.QueryDSL) (*GeoLocationTypeEntity, *workspaces.IError) {
	refl := reflect.ValueOf(&GeoLocationTypeEntity{})
	item, err := workspaces.GetOneEntity[GeoLocationTypeEntity](query, refl)

	entityGeoLocationTypeFormatter(item, query)
	return item, err
}

func GeoLocationTypeActionQuery(query workspaces.QueryDSL) ([]*GeoLocationTypeEntity, *workspaces.QueryResultMeta, error) {
	refl := reflect.ValueOf(&GeoLocationTypeEntity{})
	items, meta, err := workspaces.QueryEntitiesPointer[GeoLocationTypeEntity](query, refl)

	for _, item := range items {
		entityGeoLocationTypeFormatter(item, query)
	}

	return items, meta, err
}

func GeoLocationTypeActionUpdateFn(query workspaces.QueryDSL, fields *GeoLocationTypeEntity) (*GeoLocationTypeEntity, *workspaces.IError) {
	if fields == nil {
		return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}

	// 1. Validate always
	if iError := GeoLocationTypeValidator(fields, true); iError != nil {
		return nil, iError
	}

	GeoLocationTypeRecursiveAddUniqueId(fields, query)

	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()

		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			_, err := GeoLocationTypeUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}

		})
		return nil, workspaces.CastToIError(vf)
	} else {
		dbref = query.Tx
		return GeoLocationTypeUpdateExec(dbref, query, fields)
	}

}

func GeoLocationTypeUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *GeoLocationTypeEntity) (*GeoLocationTypeEntity, *workspaces.IError) {
	uniqueId := fields.UniqueId

	query.TriggerEventName = GEOLOCATIONTYPE_EVENT_UPDATED

	GeoLocationTypeEntityPreSanitize(fields, query)
	var item GeoLocationTypeEntity
	q := dbref.
		Where(&GeoLocationTypeEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)

	err := q.UpdateColumns(fields).Error
	if err != nil {

		return nil, workspaces.GormErrorToIError(err)
	}

	query.Tx = dbref
	GeoLocationTypeRelationContentUpdate(fields, query)

	GeoLocationTypePolyglotCreateHandler(fields, query)

	// @meta(update has many)

	err = dbref.
		Preload(clause.Associations).
		Where(&GeoLocationTypeEntity{UniqueId: uniqueId}).
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

func GeoLocationTypeActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&GeoLocationTypeEntity{})
	query.ActionRequires = []string{PERM_ROOT_GEOLOCATIONTYPE_DELETE}
	return workspaces.RemoveEntity[GeoLocationTypeEntity](query, refl)
}

func GeoLocationTypeActionWipeClean(query workspaces.QueryDSL) (int64, error) {

	var err error
	var count int64 = 0

	{
		// Hi
		subCount, subErr := workspaces.WipeCleanEntity[GeoLocationTypeEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'GeoLocationTypeEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}

	return count, err
}

func GeoLocationTypeActionBulkUpdate(
	query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[GeoLocationTypeEntity]) (
	*workspaces.BulkRecordRequest[GeoLocationTypeEntity], *workspaces.IError,
) {
	result := []*GeoLocationTypeEntity{}
	err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := GeoLocationTypeActionUpdate(query, record)

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

func (x *GeoLocationTypeEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))

	}
	return ""
}

var GeoLocationTypeEntityMeta = workspaces.TableMetaData{
	EntityName:    "GeoLocationType",
	ExportKey:     "geoLocationTypes",
	TableNameInDb: "fb_geoLocationType_entities",
	EntityObject:  &GeoLocationTypeEntity{},
	ExportStream:  GeoLocationTypeActionExportT,
	ImportQuery:   GeoLocationTypeActionImport,
}

func GeoLocationTypeActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[GeoLocationTypeEntity](query, GeoLocationTypeActionQuery, GeoLocationTypePreloadRelations)
}

func GeoLocationTypeActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[GeoLocationTypeEntity](query, GeoLocationTypeActionQuery, GeoLocationTypePreloadRelations)
}

func GeoLocationTypeActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content GeoLocationTypeEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}

	json.Unmarshal(cx, &content)

	_, err := GeoLocationTypeActionCreate(&content, query)

	return err
}
