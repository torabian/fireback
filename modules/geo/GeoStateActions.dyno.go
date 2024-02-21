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

func entityGeoStateFormatter(dto *GeoStateEntity, query workspaces.QueryDSL) {
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

func GeoStateMockEntity() *GeoStateEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &GeoStateEntity{
		UniqueId: workspaces.UUID(),
		Name:     &stringHolder,
	}

	return entity
}

func GeoStateActionSeeder(query workspaces.QueryDSL, count int) {

	successInsert := 0
	failureInsert := 0

	bar := progressbar.Default(int64(count))

	for i := 1; i <= count; i++ {
		entity := GeoStateMockEntity()
		_, err := GeoStateActionCreate(entity, query)
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

func (x *GeoStateEntity) GetNameTranslated(language string) string {
	if x.Translations != nil && len(x.Translations) > 0 {
		for _, item := range x.Translations {
			if item.LanguageId == language {
				return item.Name
			}
		}
	}

	return ""
}

func GeoStateActionSeederInit(query workspaces.QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*GeoStateEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &GeoStateEntity{

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

func GeoStateAssociationCreate(dto *GeoStateEntity, query workspaces.QueryDSL) error {

	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func GeoStateRelationContentCreate(dto *GeoStateEntity, query workspaces.QueryDSL) error {

	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* This is when we edit them directly inside the parent entity, here we would know about it.
**/
func GeoStateRelationContentUpdate(dto *GeoStateEntity, query workspaces.QueryDSL) error {

	return nil
}

func GeoStatePolyglotCreateHandler(dto *GeoStateEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}

	workspaces.PolyglotCreateHandler(dto, &GeoStateEntityPolyglot{}, query)

}

/**
* This will be validating your entity fully. Important note is that, you add validate:* tag
* in your entity, it will automatically work here. For slices inside entity, make sure you add
* extra line of AppendSliceErrors, otherwise they won't be detected
 */
func GeoStateValidator(dto *GeoStateEntity, isPatch bool) *workspaces.IError {
	err := workspaces.CommonStructValidatorPointer(dto, isPatch)

	return err
}

func GeoStateEntityPreSanitize(dto *GeoStateEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()

	_ = stripPolicy
	_ = ugcPolicy

}

func GeoStateEntityBeforeCreateAppend(dto *GeoStateEntity, query workspaces.QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = workspaces.UUID()
	}

	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId

	GeoStateRecursiveAddUniqueId(dto, query)
}

func GeoStateRecursiveAddUniqueId(dto *GeoStateEntity, query workspaces.QueryDSL) {

}

func GeoStateActionBatchCreateFn(dtos []*GeoStateEntity, query workspaces.QueryDSL) ([]*GeoStateEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*GeoStateEntity{}
		for _, item := range dtos {
			s, err := GeoStateActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)

		}
		return items, nil
	}

	return dtos, nil
}

func GeoStateActionCreateFn(dto *GeoStateEntity, query workspaces.QueryDSL) (*GeoStateEntity, *workspaces.IError) {

	// 1. Validate always
	if iError := GeoStateValidator(dto, false); iError != nil {
		return nil, iError
	}

	// 1.5 Sanitize the content coming of the front-end
	GeoStateEntityPreSanitize(dto, query)

	// 2. Append the necessary information about user, workspace
	GeoStateEntityBeforeCreateAppend(dto, query)

	// 3. Append the necessary translations, even if english
	GeoStatePolyglotCreateHandler(dto, query)

	// 3.5. Create other entities if we want select from them
	GeoStateRelationContentCreate(dto, query)

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
	GeoStateAssociationCreate(dto, query)

	// 6. Fire the event into system
	event.MustFire(GEOSTATE_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": workspaces.GetTypeString(&GeoStateEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})

	return dto, nil
}

func GeoStateActionGetOne(query workspaces.QueryDSL) (*GeoStateEntity, *workspaces.IError) {
	refl := reflect.ValueOf(&GeoStateEntity{})
	item, err := workspaces.GetOneEntity[GeoStateEntity](query, refl)

	entityGeoStateFormatter(item, query)
	return item, err
}

func GeoStateActionQuery(query workspaces.QueryDSL) ([]*GeoStateEntity, *workspaces.QueryResultMeta, error) {
	refl := reflect.ValueOf(&GeoStateEntity{})
	items, meta, err := workspaces.QueryEntitiesPointer[GeoStateEntity](query, refl)

	for _, item := range items {
		entityGeoStateFormatter(item, query)
	}

	return items, meta, err
}

func GeoStateActionUpdateFn(query workspaces.QueryDSL, fields *GeoStateEntity) (*GeoStateEntity, *workspaces.IError) {
	if fields == nil {
		return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}

	// 1. Validate always
	if iError := GeoStateValidator(fields, true); iError != nil {
		return nil, iError
	}

	GeoStateRecursiveAddUniqueId(fields, query)

	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()

		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			_, err := GeoStateUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}

		})
		return nil, workspaces.CastToIError(vf)
	} else {
		dbref = query.Tx
		return GeoStateUpdateExec(dbref, query, fields)
	}

}

func GeoStateUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *GeoStateEntity) (*GeoStateEntity, *workspaces.IError) {
	uniqueId := fields.UniqueId

	query.TriggerEventName = GEOSTATE_EVENT_UPDATED

	GeoStateEntityPreSanitize(fields, query)
	var item GeoStateEntity
	q := dbref.
		Where(&GeoStateEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)

	err := q.UpdateColumns(fields).Error
	if err != nil {

		return nil, workspaces.GormErrorToIError(err)
	}

	query.Tx = dbref
	GeoStateRelationContentUpdate(fields, query)

	GeoStatePolyglotCreateHandler(fields, query)

	// @meta(update has many)

	err = dbref.
		Preload(clause.Associations).
		Where(&GeoStateEntity{UniqueId: uniqueId}).
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

func GeoStateActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&GeoStateEntity{})
	query.ActionRequires = []string{PERM_ROOT_GEOSTATE_DELETE}
	return workspaces.RemoveEntity[GeoStateEntity](query, refl)
}

func GeoStateActionWipeClean(query workspaces.QueryDSL) (int64, error) {

	var err error
	var count int64 = 0

	{
		// Hi
		subCount, subErr := workspaces.WipeCleanEntity[GeoStateEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'GeoStateEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}

	return count, err
}

func GeoStateActionBulkUpdate(
	query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[GeoStateEntity]) (
	*workspaces.BulkRecordRequest[GeoStateEntity], *workspaces.IError,
) {
	result := []*GeoStateEntity{}
	err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := GeoStateActionUpdate(query, record)

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

func (x *GeoStateEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))

	}
	return ""
}

var GeoStateEntityMeta = workspaces.TableMetaData{
	EntityName:    "GeoState",
	ExportKey:     "geoStates",
	TableNameInDb: "fb_geoState_entities",
	EntityObject:  &GeoStateEntity{},
	ExportStream:  GeoStateActionExportT,
	ImportQuery:   GeoStateActionImport,
}

func GeoStateActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[GeoStateEntity](query, GeoStateActionQuery, GeoStatePreloadRelations)
}

func GeoStateActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[GeoStateEntity](query, GeoStateActionQuery, GeoStatePreloadRelations)
}

func GeoStateActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content GeoStateEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}

	json.Unmarshal(cx, &content)

	_, err := GeoStateActionCreate(&content, query)

	return err
}
