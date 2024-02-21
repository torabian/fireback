package worldtimezone

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

func TimezoneGroupUtcItemsActionCreate(
	dto *TimezoneGroupUtcItemsEntity,
	query workspaces.QueryDSL,
) (*TimezoneGroupUtcItemsEntity, *workspaces.IError) {

	dto.LinkerId = &query.LinkerId

	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()
	} else {
		dbref = query.Tx
	}

	query.Tx = dbref
	if dto.UniqueId == "" {
		dto.UniqueId = workspaces.UUID()
	}
	err := dbref.Create(&dto).Error
	if err != nil {
		err := workspaces.GormErrorToIError(err)
		return dto, err
	}

	return dto, nil
}

func TimezoneGroupUtcItemsActionUpdate(
	query workspaces.QueryDSL,
	dto *TimezoneGroupUtcItemsEntity,
) (*TimezoneGroupUtcItemsEntity, *workspaces.IError) {

	dto.LinkerId = &query.LinkerId

	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()
	} else {
		dbref = query.Tx
	}

	query.Tx = dbref
	err := dbref.UpdateColumns(&dto).Error
	if err != nil {
		err := workspaces.GormErrorToIError(err)
		return dto, err
	}

	return dto, nil
}

func TimezoneGroupUtcItemsActionGetOne(
	query workspaces.QueryDSL,
) (*TimezoneGroupUtcItemsEntity, *workspaces.IError) {

	refl := reflect.ValueOf(&TimezoneGroupUtcItemsEntity{})
	item, err := workspaces.GetOneEntity[TimezoneGroupUtcItemsEntity](query, refl)
	return item, err
}

func entityTimezoneGroupFormatter(dto *TimezoneGroupEntity, query workspaces.QueryDSL) {
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

func TimezoneGroupMockEntity() *TimezoneGroupEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &TimezoneGroupEntity{
		UniqueId: workspaces.UUID(),
		Value:    &stringHolder,
		Abbr:     &stringHolder,
		Offset:   &int64Holder,
		Text:     &stringHolder,
		UtcItems: []*TimezoneGroupUtcItemsEntity{{
			Name: &stringHolder,
		}},
	}

	return entity
}

func TimezoneGroupActionSeeder(query workspaces.QueryDSL, count int) {

	successInsert := 0
	failureInsert := 0

	bar := progressbar.Default(int64(count))

	for i := 1; i <= count; i++ {
		entity := TimezoneGroupMockEntity()
		_, err := TimezoneGroupActionCreate(entity, query)
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

func (x *TimezoneGroupEntity) GetValueTranslated(language string) string {
	if x.Translations != nil && len(x.Translations) > 0 {
		for _, item := range x.Translations {
			if item.LanguageId == language {
				return item.Value
			}
		}
	}

	return ""
}

func (x *TimezoneGroupEntity) GetTextTranslated(language string) string {
	if x.Translations != nil && len(x.Translations) > 0 {
		for _, item := range x.Translations {
			if item.LanguageId == language {
				return item.Text
			}
		}
	}

	return ""
}

func TimezoneGroupActionSeederInit(query workspaces.QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*TimezoneGroupEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &TimezoneGroupEntity{

		Value: &tildaRef,

		Abbr: &tildaRef,

		Text: &tildaRef,

		UtcItems: []*TimezoneGroupUtcItemsEntity{{}},
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

func TimezoneGroupAssociationCreate(dto *TimezoneGroupEntity, query workspaces.QueryDSL) error {

	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func TimezoneGroupRelationContentCreate(dto *TimezoneGroupEntity, query workspaces.QueryDSL) error {

	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* This is when we edit them directly inside the parent entity, here we would know about it.
**/
func TimezoneGroupRelationContentUpdate(dto *TimezoneGroupEntity, query workspaces.QueryDSL) error {

	return nil
}

func TimezoneGroupPolyglotCreateHandler(dto *TimezoneGroupEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}

	workspaces.PolyglotCreateHandler(dto, &TimezoneGroupEntityPolyglot{}, query)

}

/**
* This will be validating your entity fully. Important note is that, you add validate:* tag
* in your entity, it will automatically work here. For slices inside entity, make sure you add
* extra line of AppendSliceErrors, otherwise they won't be detected
 */
func TimezoneGroupValidator(dto *TimezoneGroupEntity, isPatch bool) *workspaces.IError {
	err := workspaces.CommonStructValidatorPointer(dto, isPatch)

	if dto != nil && dto.UtcItems != nil {
		workspaces.AppendSliceErrors(dto.UtcItems, isPatch, "utcitems", err)
	}

	return err
}

func TimezoneGroupEntityPreSanitize(dto *TimezoneGroupEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()

	_ = stripPolicy
	_ = ugcPolicy

}

func TimezoneGroupEntityBeforeCreateAppend(dto *TimezoneGroupEntity, query workspaces.QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = workspaces.UUID()
	}

	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId

	TimezoneGroupRecursiveAddUniqueId(dto, query)
}

func TimezoneGroupRecursiveAddUniqueId(dto *TimezoneGroupEntity, query workspaces.QueryDSL) {

	if dto.UtcItems != nil && len(dto.UtcItems) > 0 {
		for index0 := range dto.UtcItems {
			if dto.UtcItems[index0].UniqueId == "" {
				dto.UtcItems[index0].UniqueId = workspaces.UUID()
			}

		}
	}

}

func TimezoneGroupActionBatchCreateFn(dtos []*TimezoneGroupEntity, query workspaces.QueryDSL) ([]*TimezoneGroupEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*TimezoneGroupEntity{}
		for _, item := range dtos {
			s, err := TimezoneGroupActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)

		}
		return items, nil
	}

	return dtos, nil
}

func TimezoneGroupActionCreateFn(dto *TimezoneGroupEntity, query workspaces.QueryDSL) (*TimezoneGroupEntity, *workspaces.IError) {

	// 1. Validate always
	if iError := TimezoneGroupValidator(dto, false); iError != nil {
		return nil, iError
	}

	// 1.5 Sanitize the content coming of the front-end
	TimezoneGroupEntityPreSanitize(dto, query)

	// 2. Append the necessary information about user, workspace
	TimezoneGroupEntityBeforeCreateAppend(dto, query)

	// 3. Append the necessary translations, even if english
	TimezoneGroupPolyglotCreateHandler(dto, query)

	// 3.5. Create other entities if we want select from them
	TimezoneGroupRelationContentCreate(dto, query)

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
	TimezoneGroupAssociationCreate(dto, query)

	// 6. Fire the event into system
	event.MustFire(TIMEZONEGROUP_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": workspaces.GetTypeString(&TimezoneGroupEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})

	return dto, nil
}

func TimezoneGroupActionGetOne(query workspaces.QueryDSL) (*TimezoneGroupEntity, *workspaces.IError) {
	refl := reflect.ValueOf(&TimezoneGroupEntity{})
	item, err := workspaces.GetOneEntity[TimezoneGroupEntity](query, refl)

	entityTimezoneGroupFormatter(item, query)
	return item, err
}

func TimezoneGroupActionQuery(query workspaces.QueryDSL) ([]*TimezoneGroupEntity, *workspaces.QueryResultMeta, error) {
	refl := reflect.ValueOf(&TimezoneGroupEntity{})
	items, meta, err := workspaces.QueryEntitiesPointer[TimezoneGroupEntity](query, refl)

	for _, item := range items {
		entityTimezoneGroupFormatter(item, query)
	}

	return items, meta, err
}

func TimezoneGroupActionUpdateFn(query workspaces.QueryDSL, fields *TimezoneGroupEntity) (*TimezoneGroupEntity, *workspaces.IError) {
	if fields == nil {
		return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}

	// 1. Validate always
	if iError := TimezoneGroupValidator(fields, true); iError != nil {
		return nil, iError
	}

	TimezoneGroupRecursiveAddUniqueId(fields, query)

	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()

		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			_, err := TimezoneGroupUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}

		})
		return nil, workspaces.CastToIError(vf)
	} else {
		dbref = query.Tx
		return TimezoneGroupUpdateExec(dbref, query, fields)
	}

}

func TimezoneGroupUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *TimezoneGroupEntity) (*TimezoneGroupEntity, *workspaces.IError) {
	uniqueId := fields.UniqueId

	query.TriggerEventName = TIMEZONEGROUP_EVENT_UPDATED

	TimezoneGroupEntityPreSanitize(fields, query)
	var item TimezoneGroupEntity
	q := dbref.
		Where(&TimezoneGroupEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)

	err := q.UpdateColumns(fields).Error
	if err != nil {

		return nil, workspaces.GormErrorToIError(err)
	}

	query.Tx = dbref
	TimezoneGroupRelationContentUpdate(fields, query)

	TimezoneGroupPolyglotCreateHandler(fields, query)

	// @meta(update has many)

	// Inner array xobjects
	if fields.UtcItems != nil {
		// var items []TimezoneGroupUtcItemsEntity

		// if len(fields.UtcItems) > 0 {
		// 	dbref.
		// 		Where(&fields.UtcItems).
		// 		Find(&items)
		// }

		// dbref.
		// 	Model(&TimezoneGroupEntity{UniqueId: uniqueId}).
		// 	Association("UtcItems").
		// 	Replace(&items)

		/**
		*	Strategy is another hell of work to be handled.
		*   We might have strategy to delete only a set, append, delete all.
		*   This requries Every entity with an array, to have a DTO as well
		*   So let's just send entire items each time
		 */
		strategy := "replace_all"

		if strategy == "replace_all" {

			linkerId := uniqueId

			dbref.Debug().
				Where(&TimezoneGroupUtcItemsEntity{LinkerId: &linkerId}).
				Delete(&TimezoneGroupUtcItemsEntity{})

			for _, newItem := range fields.UtcItems {

				newItem.UniqueId = workspaces.UUID()

				newItem.LinkerId = &linkerId
				dbref.Create(&newItem)
			}
		}
	}

	err = dbref.
		Preload(clause.Associations).
		Where(&TimezoneGroupEntity{UniqueId: uniqueId}).
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

func TimezoneGroupActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&TimezoneGroupEntity{})
	query.ActionRequires = []string{PERM_ROOT_TIMEZONEGROUP_DELETE}
	return workspaces.RemoveEntity[TimezoneGroupEntity](query, refl)
}

func TimezoneGroupActionWipeClean(query workspaces.QueryDSL) (int64, error) {

	var err error
	var count int64 = 0

	{
		subCount, subErr := workspaces.WipeCleanEntity[TimezoneGroupUtcItemsEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'TimezoneGroupUtcItemsEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}

	{
		// Hi
		subCount, subErr := workspaces.WipeCleanEntity[TimezoneGroupEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'TimezoneGroupEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}

	return count, err
}

func TimezoneGroupActionBulkUpdate(
	query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[TimezoneGroupEntity]) (
	*workspaces.BulkRecordRequest[TimezoneGroupEntity], *workspaces.IError,
) {
	result := []*TimezoneGroupEntity{}
	err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := TimezoneGroupActionUpdate(query, record)

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

func (x *TimezoneGroupEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))

	}
	return ""
}

var TimezoneGroupEntityMeta = workspaces.TableMetaData{
	EntityName:    "TimezoneGroup",
	ExportKey:     "timezoneGroups",
	TableNameInDb: "fb_timezoneGroup_entities",
	EntityObject:  &TimezoneGroupEntity{},
	ExportStream:  TimezoneGroupActionExportT,
	ImportQuery:   TimezoneGroupActionImport,
}

func TimezoneGroupActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[TimezoneGroupEntity](query, TimezoneGroupActionQuery, TimezoneGroupPreloadRelations)
}

func TimezoneGroupActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[TimezoneGroupEntity](query, TimezoneGroupActionQuery, TimezoneGroupPreloadRelations)
}

func TimezoneGroupActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content TimezoneGroupEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}

	json.Unmarshal(cx, &content)

	_, err := TimezoneGroupActionCreate(&content, query)

	return err
}
