package geo

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	reflect "reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gookit/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/microcosm-cc/bluemonday"
	"github.com/schollz/progressbar/v3"
	metas "github.com/torabian/fireback/modules/geo/metas"
	queries "github.com/torabian/fireback/modules/geo/queries"
	seeders "github.com/torabian/fireback/modules/geo/seeders/GeoLocation"
	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GeoLocationEntity struct {
	Visibility       *string `json:"visibility,omitempty" yaml:"visibility"`
	WorkspaceId      *string `json:"workspaceId,omitempty" yaml:"workspaceId"`
	LinkerId         *string `json:"linkerId,omitempty" yaml:"linkerId"`
	ParentId         *string `json:"parentId,omitempty" yaml:"parentId"`
	UniqueId         string  `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
	UserId           *string `json:"userId,omitempty" yaml:"userId"`
	Rank             int64   `json:"rank,omitempty" gorm:"type:int;name:rank"`
	Updated          int64   `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
	Created          int64   `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
	CreatedFormatted string  `json:"createdFormatted,omitempty" sql:"-" gorm:"-"`
	UpdatedFormatted string  `json:"updatedFormatted,omitempty" sql:"-" gorm:"-"`
	Name             *string `json:"name" yaml:"name"        translate:"true" `
	// Datenano also has a text representation
	Code *string `json:"code" yaml:"code"       `
	// Datenano also has a text representation
	Type *GeoLocationTypeEntity `json:"type" yaml:"type"    gorm:"foreignKey:TypeId;references:UniqueId"     `
	// Datenano also has a text representation
	TypeId *string `json:"typeId" yaml:"typeId"`
	Status *string `json:"status" yaml:"status"       `
	// Datenano also has a text representation
	Flag *string `json:"flag" yaml:"flag"       `
	// Datenano also has a text representation
	OfficialName *string `json:"officialName" yaml:"officialName"        translate:"true" `
	// Datenano also has a text representation
	Translations []*GeoLocationEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
	Children     []*GeoLocationEntity         `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
	LinkedTo     *GeoLocationEntity           `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

var GeoLocationPreloadRelations []string = []string{}
var GEOLOCATION_EVENT_CREATED = "geoLocation.created"
var GEOLOCATION_EVENT_UPDATED = "geoLocation.updated"
var GEOLOCATION_EVENT_DELETED = "geoLocation.deleted"
var GEOLOCATION_EVENTS = []string{
	GEOLOCATION_EVENT_CREATED,
	GEOLOCATION_EVENT_UPDATED,
	GEOLOCATION_EVENT_DELETED,
}

type GeoLocationFieldMap struct {
	Name         workspaces.TranslatedString `yaml:"name"`
	Code         workspaces.TranslatedString `yaml:"code"`
	Type         workspaces.TranslatedString `yaml:"type"`
	Status       workspaces.TranslatedString `yaml:"status"`
	Flag         workspaces.TranslatedString `yaml:"flag"`
	OfficialName workspaces.TranslatedString `yaml:"officialName"`
}

var GeoLocationEntityMetaConfig map[string]int64 = map[string]int64{}
var GeoLocationEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&GeoLocationEntity{}))

type GeoLocationEntityPolyglot struct {
	LinkerId     string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
	LanguageId   string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
	Name         string `yaml:"name" json:"name"`
	OfficialName string `yaml:"officialName" json:"officialName"`
}

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
		Name:         &tildaRef,
		Code:         &tildaRef,
		Status:       &tildaRef,
		Flag:         &tildaRef,
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

var GeoLocationWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire geolocations ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		count, _ := GeoLocationActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
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
	ExportKey:     "geo-locations",
	TableNameInDb: "fb_geolocation_entities",
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

var GeoLocationCommonCliFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "wid",
		Required: false,
		Usage:    "Provide workspace id, if you want to change the data workspace",
	},
	&cli.StringFlag{
		Name:     "uid",
		Required: false,
		Usage:    "uniqueId (primary key)",
	},
	&cli.StringFlag{
		Name:     "pid",
		Required: false,
		Usage:    " Parent record id of the same type",
	},
	&cli.StringFlag{
		Name:     "name",
		Required: false,
		Usage:    "name",
	},
	&cli.StringFlag{
		Name:     "code",
		Required: false,
		Usage:    "code",
	},
	&cli.StringFlag{
		Name:     "type-id",
		Required: false,
		Usage:    "type",
	},
	&cli.StringFlag{
		Name:     "status",
		Required: false,
		Usage:    "status",
	},
	&cli.StringFlag{
		Name:     "flag",
		Required: false,
		Usage:    "flag",
	},
	&cli.StringFlag{
		Name:     "official-name",
		Required: false,
		Usage:    "officialName",
	},
}
var GeoLocationCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:        "name",
		StructField: "Name",
		Required:    false,
		Usage:       "name",
		Type:        "string",
	},
	{
		Name:        "code",
		StructField: "Code",
		Required:    false,
		Usage:       "code",
		Type:        "string",
	},
	{
		Name:        "status",
		StructField: "Status",
		Required:    false,
		Usage:       "status",
		Type:        "string",
	},
	{
		Name:        "flag",
		StructField: "Flag",
		Required:    false,
		Usage:       "flag",
		Type:        "string",
	},
	{
		Name:        "officialName",
		StructField: "OfficialName",
		Required:    false,
		Usage:       "officialName",
		Type:        "string",
	},
}
var GeoLocationCommonCliFlagsOptional = []cli.Flag{
	&cli.StringFlag{
		Name:     "wid",
		Required: false,
		Usage:    "Provide workspace id, if you want to change the data workspace",
	},
	&cli.StringFlag{
		Name:     "uid",
		Required: false,
		Usage:    "uniqueId (primary key)",
	},
	&cli.StringFlag{
		Name:     "pid",
		Required: false,
		Usage:    " Parent record id of the same type",
	},
	&cli.StringFlag{
		Name:     "name",
		Required: false,
		Usage:    "name",
	},
	&cli.StringFlag{
		Name:     "code",
		Required: false,
		Usage:    "code",
	},
	&cli.StringFlag{
		Name:     "type-id",
		Required: false,
		Usage:    "type",
	},
	&cli.StringFlag{
		Name:     "status",
		Required: false,
		Usage:    "status",
	},
	&cli.StringFlag{
		Name:     "flag",
		Required: false,
		Usage:    "flag",
	},
	&cli.StringFlag{
		Name:     "official-name",
		Required: false,
		Usage:    "officialName",
	},
}
var GeoLocationCreateCmd cli.Command = cli.Command{
	Name:    "create",
	Aliases: []string{"c"},
	Flags:   GeoLocationCommonCliFlags,
	Usage:   "Create a new template",
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastGeoLocationFromCli(c)
		if entity, err := GeoLocationActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var GeoLocationCreateInteractiveCmd cli.Command = cli.Command{
	Name:  "ic",
	Usage: "Creates a new template, using requied fields in an interactive name",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "all",
			Usage: "Interactively asks for all inputs, not only required ones",
		},
	},
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := &GeoLocationEntity{}
		for _, item := range GeoLocationCommonInteractiveCliFlags {
			if !item.Required && c.Bool("all") == false {
				continue
			}
			result := workspaces.AskForInput(item.Name, "")
			workspaces.SetFieldString(entity, item.StructField, result)
		}
		if entity, err := GeoLocationActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var GeoLocationUpdateCmd cli.Command = cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Flags:   GeoLocationCommonCliFlagsOptional,
	Usage:   "Updates a template by passing the parameters",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastGeoLocationFromCli(c)
		if entity, err := GeoLocationActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
		return nil
	},
}

func CastGeoLocationFromCli(c *cli.Context) *GeoLocationEntity {
	template := &GeoLocationEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
	if c.IsSet("name") {
		value := c.String("name")
		template.Name = &value
	}
	if c.IsSet("code") {
		value := c.String("code")
		template.Code = &value
	}
	if c.IsSet("type-id") {
		value := c.String("type-id")
		template.TypeId = &value
	}
	if c.IsSet("status") {
		value := c.String("status")
		template.Status = &value
	}
	if c.IsSet("flag") {
		value := c.String("flag")
		template.Flag = &value
	}
	if c.IsSet("official-name") {
		value := c.String("official-name")
		template.OfficialName = &value
	}
	return template
}
func GeoLocationSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		GeoLocationActionCreate,
		reflect.ValueOf(&GeoLocationEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}
func GeoLocationSyncSeeders() {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{WorkspaceId: workspaces.USER_SYSTEM},
		GeoLocationActionCreate,
		reflect.ValueOf(&GeoLocationEntity{}).Elem(),
		&seeders.ViewsFs,
		[]string{},
		true,
	)
}
func GeoLocationWriteQueryMock(ctx workspaces.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := GeoLocationActionQuery(f)
		result := workspaces.QueryEntitySuccessResult(f, items, count)
		workspaces.WriteMockDataToFile(lang, "", "GeoLocation", result)
	}
}

var GeoLocationImportExportCommands = []cli.Command{
	{
		Name:  "mock",
		Usage: "Generates mock records based on the entity definition",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "count",
				Usage: "how many activation key do you need to be generated and stored in database",
				Value: 10,
			},
		},
		Action: func(c *cli.Context) error {
			query := workspaces.CommonCliQueryDSLBuilder(c)
			GeoLocationActionSeeder(query, c.Int("count"))
			return nil
		},
	},
	{
		Name:    "init",
		Aliases: []string{"i"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "file",
				Usage: "The address of file you want the csv be exported to",
				Value: "geo-location-seeder.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Creates a basic seeder file for you, based on the definition module we have. You can populate this file as an example",
		Action: func(c *cli.Context) error {
			f := workspaces.CommonCliQueryDSLBuilder(c)
			GeoLocationActionSeederInit(f, c.String("file"), c.String("format"))
			return nil
		},
	},
	{
		Name:    "validate",
		Aliases: []string{"v"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "file",
				Usage: "Validates an import file, such as yaml, json, csv, and gives some insights how the after import it would look like",
				Value: "geo-location-seeder-geo-location.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of geo-locations, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]GeoLocationEntity{}
			workspaces.ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
	cli.Command{
		Name:  "list",
		Usage: "Prints the list of files attached to this module for syncing or bootstrapping project",
		Action: func(c *cli.Context) error {
			if entity, err := workspaces.GetSeederFilenames(&seeders.ViewsFs, ""); err != nil {
				fmt.Println(err.Error())
			} else {
				f, _ := json.MarshalIndent(entity, "", "  ")
				fmt.Println(string(f))
			}
			return nil
		},
	},
	cli.Command{
		Name:  "sync",
		Usage: "Tries to sync the embedded content into the database, the list could be seen by 'list' command",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportEmbedCmd(c,
				GeoLocationActionCreate,
				reflect.ValueOf(&GeoLocationEntity{}).Elem(),
				&seeders.ViewsFs,
			)
			return nil
		},
	},
	cli.Command{
		Name:    "export",
		Aliases: []string{"e"},
		Flags: append(workspaces.CommonQueryFlags,
			&cli.StringFlag{
				Name:     "file",
				Usage:    "The address of file you want the csv/yaml/json be exported to",
				Required: true,
			}),
		Usage: "Exports a query results into the csv/yaml/json format",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliExportCmd(c,
				GeoLocationActionQuery,
				reflect.ValueOf(&GeoLocationEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"GeoLocationFieldMap.yml",
				GeoLocationPreloadRelations,
			)
			return nil
		},
	},
	cli.Command{
		Name: "import",
		Flags: append(workspaces.CommonQueryFlags,
			&cli.StringFlag{
				Name:     "file",
				Usage:    "The address of file you want the csv be imported from",
				Required: true,
			}),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmd(c,
				GeoLocationActionCreate,
				reflect.ValueOf(&GeoLocationEntity{}).Elem(),
				c.String("file"),
			)
			return nil
		},
	},
}
var GeoLocationCliCommands []cli.Command = []cli.Command{
	workspaces.GetCommonQuery(GeoLocationActionQuery),
	workspaces.GetCommonTableQuery(reflect.ValueOf(&GeoLocationEntity{}).Elem(), GeoLocationActionQuery),
	GeoLocationCreateCmd,
	GeoLocationUpdateCmd,
	GeoLocationCreateInteractiveCmd,
	GeoLocationWipeCmd,
	workspaces.GetCommonRemoveQuery(reflect.ValueOf(&GeoLocationEntity{}).Elem(), GeoLocationActionRemove),
	workspaces.GetCommonCteQuery(GeoLocationActionCteQuery),
	workspaces.GetCommonPivotQuery(GeoLocationActionCommonPivotQuery),
}

func GeoLocationCliFn() cli.Command {
	GeoLocationCliCommands = append(GeoLocationCliCommands, GeoLocationImportExportCommands...)
	return cli.Command{
		Name:        "location",
		Description: "GeoLocations module actions (sample module to handle complex entities)",
		Usage:       "",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: GeoLocationCliCommands,
	}
}

/**
 *	Override this function on GeoLocationEntityHttp.go,
 *	In order to add your own http
 **/
var AppendGeoLocationRouter = func(r *[]workspaces.Module2Action) {}

func GetGeoLocationModule2Actions() []workspaces.Module2Action {
	routes := []workspaces.Module2Action{
		{
			Method: "GET",
			Url:    "/cte-geo-locations",
			SecurityModel: &workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_GEOLOCATION_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpQueryEntity(c, GeoLocationActionCteQuery)
				},
			},
			Format:         "QUERY",
			Action:         GeoLocationActionCteQuery,
			ResponseEntity: &[]GeoLocationEntity{},
		},
		{
			Method: "GET",
			Url:    "/geo-locations",
			SecurityModel: &workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_GEOLOCATION_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpQueryEntity(c, GeoLocationActionQuery)
				},
			},
			Format:         "QUERY",
			Action:         GeoLocationActionQuery,
			ResponseEntity: &[]GeoLocationEntity{},
		},
		{
			Method: "GET",
			Url:    "/geo-locations/export",
			SecurityModel: &workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_GEOLOCATION_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpStreamFileChannel(c, GeoLocationActionExport)
				},
			},
			Format:         "QUERY",
			Action:         GeoLocationActionExport,
			ResponseEntity: &[]GeoLocationEntity{},
		},
		{
			Method: "GET",
			Url:    "/geo-location/:uniqueId",
			SecurityModel: &workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_GEOLOCATION_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpGetEntity(c, GeoLocationActionGetOne)
				},
			},
			Format:         "GET_ONE",
			Action:         GeoLocationActionGetOne,
			ResponseEntity: &GeoLocationEntity{},
		},
		{
			Method: "POST",
			Url:    "/geo-location",
			SecurityModel: &workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_GEOLOCATION_CREATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpPostEntity(c, GeoLocationActionCreate)
				},
			},
			Action:         GeoLocationActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &GeoLocationEntity{},
			ResponseEntity: &GeoLocationEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/geo-location",
			SecurityModel: &workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_GEOLOCATION_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpUpdateEntity(c, GeoLocationActionUpdate)
				},
			},
			Action:         GeoLocationActionUpdate,
			RequestEntity:  &GeoLocationEntity{},
			Format:         "PATCH_ONE",
			ResponseEntity: &GeoLocationEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/geo-locations",
			SecurityModel: &workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_GEOLOCATION_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpUpdateEntities(c, GeoLocationActionBulkUpdate)
				},
			},
			Action:         GeoLocationActionBulkUpdate,
			Format:         "PATCH_BULK",
			RequestEntity:  &workspaces.BulkRecordRequest[GeoLocationEntity]{},
			ResponseEntity: &workspaces.BulkRecordRequest[GeoLocationEntity]{},
		},
		{
			Method: "DELETE",
			Url:    "/geo-location",
			Format: "DELETE_DSL",
			SecurityModel: &workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_GEOLOCATION_DELETE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpRemoveEntity(c, GeoLocationActionRemove)
				},
			},
			Action:         GeoLocationActionRemove,
			RequestEntity:  &workspaces.DeleteRequest{},
			ResponseEntity: &workspaces.DeleteResponse{},
			TargetEntity:   &GeoLocationEntity{},
		},
	}
	// Append user defined functions
	AppendGeoLocationRouter(&routes)
	return routes
}
func CreateGeoLocationRouter(r *gin.Engine) []workspaces.Module2Action {
	httpRoutes := GetGeoLocationModule2Actions()
	workspaces.CastRoutes(httpRoutes, r)
	workspaces.WriteHttpInformationToFile(&httpRoutes, GeoLocationEntityJsonSchema, "geo-location-http", "geo")
	workspaces.WriteEntitySchema("GeoLocationEntity", GeoLocationEntityJsonSchema, "geo")
	return httpRoutes
}

var PERM_ROOT_GEOLOCATION_DELETE = "root/geolocation/delete"
var PERM_ROOT_GEOLOCATION_CREATE = "root/geolocation/create"
var PERM_ROOT_GEOLOCATION_UPDATE = "root/geolocation/update"
var PERM_ROOT_GEOLOCATION_QUERY = "root/geolocation/query"
var PERM_ROOT_GEOLOCATION = "root/geolocation"
var ALL_GEOLOCATION_PERMISSIONS = []string{
	PERM_ROOT_GEOLOCATION_DELETE,
	PERM_ROOT_GEOLOCATION_CREATE,
	PERM_ROOT_GEOLOCATION_UPDATE,
	PERM_ROOT_GEOLOCATION_QUERY,
	PERM_ROOT_GEOLOCATION,
}
