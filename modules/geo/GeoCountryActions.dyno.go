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

func entityGeoCountryFormatter(dto *GeoCountryEntity, query workspaces.QueryDSL) {
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

func GeoCountryMockEntity() *GeoCountryEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &GeoCountryEntity{
		UniqueId:     workspaces.UUID(),
		Status:       &stringHolder,
		Flag:         &stringHolder,
		CommonName:   &stringHolder,
		OfficialName: &stringHolder,
	}

	return entity
}

func GeoCountryActionSeeder(query workspaces.QueryDSL, count int) {

	successInsert := 0
	failureInsert := 0

	bar := progressbar.Default(int64(count))

	for i := 1; i <= count; i++ {
		entity := GeoCountryMockEntity()
		_, err := GeoCountryActionCreate(entity, query)
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

func (x *GeoCountryEntity) GetCommonNameTranslated(language string) string {
	if x.Translations != nil && len(x.Translations) > 0 {
		for _, item := range x.Translations {
			if item.LanguageId == language {
				return item.CommonName
			}
		}
	}

	return ""
}

func (x *GeoCountryEntity) GetOfficialNameTranslated(language string) string {
	if x.Translations != nil && len(x.Translations) > 0 {
		for _, item := range x.Translations {
			if item.LanguageId == language {
				return item.OfficialName
			}
		}
	}

	return ""
}

func GeoCountryActionSeederInit(query workspaces.QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*GeoCountryEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &GeoCountryEntity{

		Status: &tildaRef,

		Flag: &tildaRef,

		CommonName: &tildaRef,

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

func GeoCountryAssociationCreate(dto *GeoCountryEntity, query workspaces.QueryDSL) error {

	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func GeoCountryRelationContentCreate(dto *GeoCountryEntity, query workspaces.QueryDSL) error {

	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* This is when we edit them directly inside the parent entity, here we would know about it.
**/
func GeoCountryRelationContentUpdate(dto *GeoCountryEntity, query workspaces.QueryDSL) error {

	return nil
}

func GeoCountryPolyglotCreateHandler(dto *GeoCountryEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}

	workspaces.PolyglotCreateHandler(dto, &GeoCountryEntityPolyglot{}, query)

}

/**
* This will be validating your entity fully. Important note is that, you add validate:* tag
* in your entity, it will automatically work here. For slices inside entity, make sure you add
* extra line of AppendSliceErrors, otherwise they won't be detected
 */
func GeoCountryValidator(dto *GeoCountryEntity, isPatch bool) *workspaces.IError {
	err := workspaces.CommonStructValidatorPointer(dto, isPatch)

	return err
}

func GeoCountryEntityPreSanitize(dto *GeoCountryEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()

	_ = stripPolicy
	_ = ugcPolicy

}

func GeoCountryEntityBeforeCreateAppend(dto *GeoCountryEntity, query workspaces.QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = workspaces.UUID()
	}

	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId

	GeoCountryRecursiveAddUniqueId(dto, query)
}

func GeoCountryRecursiveAddUniqueId(dto *GeoCountryEntity, query workspaces.QueryDSL) {

}

func GeoCountryActionBatchCreateFn(dtos []*GeoCountryEntity, query workspaces.QueryDSL) ([]*GeoCountryEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*GeoCountryEntity{}
		for _, item := range dtos {
			s, err := GeoCountryActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)

		}
		return items, nil
	}

	return dtos, nil
}

func GeoCountryActionCreateFn(dto *GeoCountryEntity, query workspaces.QueryDSL) (*GeoCountryEntity, *workspaces.IError) {

	// 1. Validate always
	if iError := GeoCountryValidator(dto, false); iError != nil {
		return nil, iError
	}

	// 1.5 Sanitize the content coming of the front-end
	GeoCountryEntityPreSanitize(dto, query)

	// 2. Append the necessary information about user, workspace
	GeoCountryEntityBeforeCreateAppend(dto, query)

	// 3. Append the necessary translations, even if english
	GeoCountryPolyglotCreateHandler(dto, query)

	// 3.5. Create other entities if we want select from them
	GeoCountryRelationContentCreate(dto, query)

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
	GeoCountryAssociationCreate(dto, query)

	// 6. Fire the event into system
	event.MustFire(GEOCOUNTRY_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": workspaces.GetTypeString(&GeoCountryEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})

	return dto, nil
}

func GeoCountryActionGetOne(query workspaces.QueryDSL) (*GeoCountryEntity, *workspaces.IError) {
	refl := reflect.ValueOf(&GeoCountryEntity{})
	item, err := workspaces.GetOneEntity[GeoCountryEntity](query, refl)

	entityGeoCountryFormatter(item, query)
	return item, err
}

func GeoCountryActionQuery(query workspaces.QueryDSL) ([]*GeoCountryEntity, *workspaces.QueryResultMeta, error) {
	refl := reflect.ValueOf(&GeoCountryEntity{})
	items, meta, err := workspaces.QueryEntitiesPointer[GeoCountryEntity](query, refl)

	for _, item := range items {
		entityGeoCountryFormatter(item, query)
	}

	return items, meta, err
}

func GeoCountryActionUpdateFn(query workspaces.QueryDSL, fields *GeoCountryEntity) (*GeoCountryEntity, *workspaces.IError) {
	if fields == nil {
		return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}

	// 1. Validate always
	if iError := GeoCountryValidator(fields, true); iError != nil {
		return nil, iError
	}

	GeoCountryRecursiveAddUniqueId(fields, query)

	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()

		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			_, err := GeoCountryUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}

		})
		return nil, workspaces.CastToIError(vf)
	} else {
		dbref = query.Tx
		return GeoCountryUpdateExec(dbref, query, fields)
	}

}

func GeoCountryUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *GeoCountryEntity) (*GeoCountryEntity, *workspaces.IError) {
	uniqueId := fields.UniqueId

	query.TriggerEventName = GEOCOUNTRY_EVENT_UPDATED

	GeoCountryEntityPreSanitize(fields, query)
	var item GeoCountryEntity
	q := dbref.
		Where(&GeoCountryEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)

	err := q.UpdateColumns(fields).Error
	if err != nil {

		return nil, workspaces.GormErrorToIError(err)
	}

	query.Tx = dbref
	GeoCountryRelationContentUpdate(fields, query)

	GeoCountryPolyglotCreateHandler(fields, query)

	// @meta(update has many)

	err = dbref.
		Preload(clause.Associations).
		Where(&GeoCountryEntity{UniqueId: uniqueId}).
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

func GeoCountryActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&GeoCountryEntity{})
	query.ActionRequires = []string{PERM_ROOT_GEOCOUNTRY_DELETE}
	return workspaces.RemoveEntity[GeoCountryEntity](query, refl)
}

func GeoCountryActionWipeClean(query workspaces.QueryDSL) (int64, error) {

	var err error
	var count int64 = 0

	{
		// Hi
		subCount, subErr := workspaces.WipeCleanEntity[GeoCountryEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'GeoCountryEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}

	return count, err
}

func GeoCountryActionBulkUpdate(
	query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[GeoCountryEntity]) (
	*workspaces.BulkRecordRequest[GeoCountryEntity], *workspaces.IError,
) {
	result := []*GeoCountryEntity{}
	err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := GeoCountryActionUpdate(query, record)

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

func (x *GeoCountryEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))

	}
	return ""
}

var GeoCountryEntityMeta = workspaces.TableMetaData{
	EntityName:    "GeoCountry",
	ExportKey:     "geoCountrys",
	TableNameInDb: "fb_geoCountry_entities",
	EntityObject:  &GeoCountryEntity{},
	ExportStream:  GeoCountryActionExportT,
	ImportQuery:   GeoCountryActionImport,
}

func GeoCountryActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[GeoCountryEntity](query, GeoCountryActionQuery, GeoCountryPreloadRelations)
}

func GeoCountryActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[GeoCountryEntity](query, GeoCountryActionQuery, GeoCountryPreloadRelations)
}

func GeoCountryActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content GeoCountryEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}

	json.Unmarshal(cx, &content)

	_, err := GeoCountryActionCreate(&content, query)

	return err
}
