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

func entityLocationDataFormatter(dto *LocationDataEntity, query workspaces.QueryDSL) {
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

func LocationDataMockEntity() *LocationDataEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &LocationDataEntity{
		UniqueId:        workspaces.UUID(),
		Lat:             &float64Holder,
		Lng:             &float64Holder,
		PhysicalAddress: &stringHolder,
	}

	return entity
}

func LocationDataActionSeeder(query workspaces.QueryDSL, count int) {

	successInsert := 0
	failureInsert := 0

	bar := progressbar.Default(int64(count))

	for i := 1; i <= count; i++ {
		entity := LocationDataMockEntity()
		_, err := LocationDataActionCreate(entity, query)
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

func LocationDataActionSeederInit(query workspaces.QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*LocationDataEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &LocationDataEntity{

		PhysicalAddress: &tildaRef,
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

func LocationDataAssociationCreate(dto *LocationDataEntity, query workspaces.QueryDSL) error {

	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func LocationDataRelationContentCreate(dto *LocationDataEntity, query workspaces.QueryDSL) error {

	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* This is when we edit them directly inside the parent entity, here we would know about it.
**/
func LocationDataRelationContentUpdate(dto *LocationDataEntity, query workspaces.QueryDSL) error {

	return nil
}

func LocationDataPolyglotCreateHandler(dto *LocationDataEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}

}

/**
* This will be validating your entity fully. Important note is that, you add validate:* tag
* in your entity, it will automatically work here. For slices inside entity, make sure you add
* extra line of AppendSliceErrors, otherwise they won't be detected
 */
func LocationDataValidator(dto *LocationDataEntity, isPatch bool) *workspaces.IError {
	err := workspaces.CommonStructValidatorPointer(dto, isPatch)

	return err
}

func LocationDataEntityPreSanitize(dto *LocationDataEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()

	_ = stripPolicy
	_ = ugcPolicy

}

func LocationDataEntityBeforeCreateAppend(dto *LocationDataEntity, query workspaces.QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = workspaces.UUID()
	}

	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId

	LocationDataRecursiveAddUniqueId(dto, query)
}

func LocationDataRecursiveAddUniqueId(dto *LocationDataEntity, query workspaces.QueryDSL) {

}

func LocationDataActionBatchCreateFn(dtos []*LocationDataEntity, query workspaces.QueryDSL) ([]*LocationDataEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*LocationDataEntity{}
		for _, item := range dtos {
			s, err := LocationDataActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)

		}
		return items, nil
	}

	return dtos, nil
}

func LocationDataActionCreateFn(dto *LocationDataEntity, query workspaces.QueryDSL) (*LocationDataEntity, *workspaces.IError) {

	// 1. Validate always
	if iError := LocationDataValidator(dto, false); iError != nil {
		return nil, iError
	}

	// 1.5 Sanitize the content coming of the front-end
	LocationDataEntityPreSanitize(dto, query)

	// 2. Append the necessary information about user, workspace
	LocationDataEntityBeforeCreateAppend(dto, query)

	// 3. Append the necessary translations, even if english
	LocationDataPolyglotCreateHandler(dto, query)

	// 3.5. Create other entities if we want select from them
	LocationDataRelationContentCreate(dto, query)

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
	LocationDataAssociationCreate(dto, query)

	// 6. Fire the event into system
	event.MustFire(LOCATIONDATA_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": workspaces.GetTypeString(&LocationDataEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})

	return dto, nil
}

func LocationDataActionGetOne(query workspaces.QueryDSL) (*LocationDataEntity, *workspaces.IError) {
	refl := reflect.ValueOf(&LocationDataEntity{})
	item, err := workspaces.GetOneEntity[LocationDataEntity](query, refl)

	entityLocationDataFormatter(item, query)
	return item, err
}

func LocationDataActionQuery(query workspaces.QueryDSL) ([]*LocationDataEntity, *workspaces.QueryResultMeta, error) {
	refl := reflect.ValueOf(&LocationDataEntity{})
	items, meta, err := workspaces.QueryEntitiesPointer[LocationDataEntity](query, refl)

	for _, item := range items {
		entityLocationDataFormatter(item, query)
	}

	return items, meta, err
}

func LocationDataActionUpdateFn(query workspaces.QueryDSL, fields *LocationDataEntity) (*LocationDataEntity, *workspaces.IError) {
	if fields == nil {
		return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}

	// 1. Validate always
	if iError := LocationDataValidator(fields, true); iError != nil {
		return nil, iError
	}

	LocationDataRecursiveAddUniqueId(fields, query)

	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()

		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			_, err := LocationDataUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}

		})
		return nil, workspaces.CastToIError(vf)
	} else {
		dbref = query.Tx
		return LocationDataUpdateExec(dbref, query, fields)
	}

}

func LocationDataUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *LocationDataEntity) (*LocationDataEntity, *workspaces.IError) {
	uniqueId := fields.UniqueId

	query.TriggerEventName = LOCATIONDATA_EVENT_UPDATED

	LocationDataEntityPreSanitize(fields, query)
	var item LocationDataEntity
	q := dbref.
		Where(&LocationDataEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)

	err := q.UpdateColumns(fields).Error
	if err != nil {

		return nil, workspaces.GormErrorToIError(err)
	}

	query.Tx = dbref
	LocationDataRelationContentUpdate(fields, query)

	LocationDataPolyglotCreateHandler(fields, query)

	// @meta(update has many)

	err = dbref.
		Preload(clause.Associations).
		Where(&LocationDataEntity{UniqueId: uniqueId}).
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

func LocationDataActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&LocationDataEntity{})
	query.ActionRequires = []string{PERM_ROOT_LOCATIONDATA_DELETE}
	return workspaces.RemoveEntity[LocationDataEntity](query, refl)
}

func LocationDataActionWipeClean(query workspaces.QueryDSL) (int64, error) {

	var err error
	var count int64 = 0

	{
		// Hi
		subCount, subErr := workspaces.WipeCleanEntity[LocationDataEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'LocationDataEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}

	return count, err
}

func LocationDataActionBulkUpdate(
	query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[LocationDataEntity]) (
	*workspaces.BulkRecordRequest[LocationDataEntity], *workspaces.IError,
) {
	result := []*LocationDataEntity{}
	err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := LocationDataActionUpdate(query, record)

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

func (x *LocationDataEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))

	}
	return ""
}

var LocationDataEntityMeta = workspaces.TableMetaData{
	EntityName:    "LocationData",
	ExportKey:     "locationDatas",
	TableNameInDb: "fb_locationData_entities",
	EntityObject:  &LocationDataEntity{},
	ExportStream:  LocationDataActionExportT,
	ImportQuery:   LocationDataActionImport,
}

func LocationDataActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[LocationDataEntity](query, LocationDataActionQuery, LocationDataPreloadRelations)
}

func LocationDataActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[LocationDataEntity](query, LocationDataActionQuery, LocationDataPreloadRelations)
}

func LocationDataActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content LocationDataEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}

	json.Unmarshal(cx, &content)

	_, err := LocationDataActionCreate(&content, query)

	return err
}
