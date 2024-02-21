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
	"github.com/torabian/fireback/modules/workspaces"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	jsoniter "github.com/json-iterator/go"
)

func entityGeoCityFormatter(dto *GeoCityEntity, query workspaces.QueryDSL) {
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

func GeoCityMockEntity() *GeoCityEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &GeoCityEntity{
		UniqueId: workspaces.UUID(),
		Name:     &stringHolder,
	}

	return entity
}

func GeoCityActionSeeder(query workspaces.QueryDSL, count int) {

	successInsert := 0
	failureInsert := 0

	bar := progressbar.Default(int64(count))

	for i := 1; i <= count; i++ {
		entity := GeoCityMockEntity()
		_, err := GeoCityActionCreate(entity, query)
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

func GeoCityActionSeederInit(query workspaces.QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*GeoCityEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &GeoCityEntity{

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

func GeoCityAssociationCreate(dto *GeoCityEntity, query workspaces.QueryDSL) error {

	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func GeoCityRelationContentCreate(dto *GeoCityEntity, query workspaces.QueryDSL) error {

	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* This is when we edit them directly inside the parent entity, here we would know about it.
**/
func GeoCityRelationContentUpdate(dto *GeoCityEntity, query workspaces.QueryDSL) error {

	return nil
}

func GeoCityPolyglotCreateHandler(dto *GeoCityEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}

}

/**
* This will be validating your entity fully. Important note is that, you add validate:* tag
* in your entity, it will automatically work here. For slices inside entity, make sure you add
* extra line of AppendSliceErrors, otherwise they won't be detected
 */
func GeoCityValidator(dto *GeoCityEntity, isPatch bool) *workspaces.IError {
	err := workspaces.CommonStructValidatorPointer(dto, isPatch)

	return err
}

func GeoCityEntityPreSanitize(dto *GeoCityEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()

	_ = stripPolicy
	_ = ugcPolicy

}

func GeoCityEntityBeforeCreateAppend(dto *GeoCityEntity, query workspaces.QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = workspaces.UUID()
	}

	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId

	GeoCityRecursiveAddUniqueId(dto, query)
}

func GeoCityRecursiveAddUniqueId(dto *GeoCityEntity, query workspaces.QueryDSL) {

}

func GeoCityActionBatchCreateFn(dtos []*GeoCityEntity, query workspaces.QueryDSL) ([]*GeoCityEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*GeoCityEntity{}
		for _, item := range dtos {
			s, err := GeoCityActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)

		}
		return items, nil
	}

	return dtos, nil
}

func GeoCityActionCreateFn(dto *GeoCityEntity, query workspaces.QueryDSL) (*GeoCityEntity, *workspaces.IError) {

	// 1. Validate always
	if iError := GeoCityValidator(dto, false); iError != nil {
		return nil, iError
	}

	// 1.5 Sanitize the content coming of the front-end
	GeoCityEntityPreSanitize(dto, query)

	// 2. Append the necessary information about user, workspace
	GeoCityEntityBeforeCreateAppend(dto, query)

	// 3. Append the necessary translations, even if english
	GeoCityPolyglotCreateHandler(dto, query)

	// 3.5. Create other entities if we want select from them
	GeoCityRelationContentCreate(dto, query)

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
	GeoCityAssociationCreate(dto, query)

	// 6. Fire the event into system
	event.MustFire(GEOCITY_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": workspaces.GetTypeString(&GeoCityEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})

	return dto, nil
}

func GeoCityActionGetOne(query workspaces.QueryDSL) (*GeoCityEntity, *workspaces.IError) {
	refl := reflect.ValueOf(&GeoCityEntity{})
	item, err := workspaces.GetOneEntity[GeoCityEntity](query, refl)

	entityGeoCityFormatter(item, query)
	return item, err
}

func GeoCityActionQuery(query workspaces.QueryDSL) ([]*GeoCityEntity, *workspaces.QueryResultMeta, error) {
	refl := reflect.ValueOf(&GeoCityEntity{})
	items, meta, err := workspaces.QueryEntitiesPointer[GeoCityEntity](query, refl)

	for _, item := range items {
		entityGeoCityFormatter(item, query)
	}

	return items, meta, err
}

func GeoCityActionUpdateFn(query workspaces.QueryDSL, fields *GeoCityEntity) (*GeoCityEntity, *workspaces.IError) {
	if fields == nil {
		return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}

	// 1. Validate always
	if iError := GeoCityValidator(fields, true); iError != nil {
		return nil, iError
	}

	GeoCityRecursiveAddUniqueId(fields, query)

	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()

		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			_, err := GeoCityUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}

		})
		return nil, workspaces.CastToIError(vf)
	} else {
		dbref = query.Tx
		return GeoCityUpdateExec(dbref, query, fields)
	}

}

func GeoCityUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *GeoCityEntity) (*GeoCityEntity, *workspaces.IError) {
	uniqueId := fields.UniqueId

	query.TriggerEventName = GEOCITY_EVENT_UPDATED

	GeoCityEntityPreSanitize(fields, query)
	var item GeoCityEntity
	q := dbref.
		Where(&GeoCityEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)

	err := q.UpdateColumns(fields).Error
	if err != nil {

		return nil, workspaces.GormErrorToIError(err)
	}

	query.Tx = dbref
	GeoCityRelationContentUpdate(fields, query)

	GeoCityPolyglotCreateHandler(fields, query)

	// @meta(update has many)

	err = dbref.
		Preload(clause.Associations).
		Where(&GeoCityEntity{UniqueId: uniqueId}).
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

func GeoCityActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&GeoCityEntity{})
	query.ActionRequires = []string{PERM_ROOT_GEOCITY_DELETE}
	return workspaces.RemoveEntity[GeoCityEntity](query, refl)
}

func GeoCityActionWipeClean(query workspaces.QueryDSL) (int64, error) {

	var err error
	var count int64 = 0

	{
		// Hi
		subCount, subErr := workspaces.WipeCleanEntity[GeoCityEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'GeoCityEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}

	return count, err
}

func GeoCityActionBulkUpdate(
	query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[GeoCityEntity]) (
	*workspaces.BulkRecordRequest[GeoCityEntity], *workspaces.IError,
) {
	result := []*GeoCityEntity{}
	err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := GeoCityActionUpdate(query, record)

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

func (x *GeoCityEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))

	}
	return ""
}

var GeoCityEntityMeta = workspaces.TableMetaData{
	EntityName:    "GeoCity",
	ExportKey:     "geoCitys",
	TableNameInDb: "fb_geoCity_entities",
	EntityObject:  &GeoCityEntity{},
	ExportStream:  GeoCityActionExportT,
	ImportQuery:   GeoCityActionImport,
}

func GeoCityActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[GeoCityEntity](query, GeoCityActionQuery, GeoCityPreloadRelations)
}

func GeoCityActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[GeoCityEntity](query, GeoCityActionQuery, GeoCityPreloadRelations)
}

func GeoCityActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content GeoCityEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}

	json.Unmarshal(cx, &content)

	_, err := GeoCityActionCreate(&content, query)

	return err
}
