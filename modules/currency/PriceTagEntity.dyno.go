package currency

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
	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var priceTagSeedersFs *embed.FS = nil

func ResetPriceTagSeeders(fs *embed.FS) {
	priceTagSeedersFs = fs
}

type PriceTagVariations struct {
	Visibility       *string         `json:"visibility,omitempty" yaml:"visibility"`
	WorkspaceId      *string         `json:"workspaceId,omitempty" yaml:"workspaceId"`
	LinkerId         *string         `json:"linkerId,omitempty" yaml:"linkerId"`
	ParentId         *string         `json:"parentId,omitempty" yaml:"parentId"`
	IsDeletable      *bool           `json:"isDeletable,omitempty" yaml:"isDeletable" gorm:"default:true"`
	IsUpdatable      *bool           `json:"isUpdatable,omitempty" yaml:"isUpdatable" gorm:"default:true"`
	UniqueId         string          `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
	UserId           *string         `json:"userId,omitempty" yaml:"userId"`
	Rank             int64           `json:"rank,omitempty" gorm:"type:int;name:rank"`
	Updated          int64           `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
	Created          int64           `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
	CreatedFormatted string          `json:"createdFormatted,omitempty" sql:"-" gorm:"-"`
	UpdatedFormatted string          `json:"updatedFormatted,omitempty" sql:"-" gorm:"-"`
	Currency         *CurrencyEntity `json:"currency" yaml:"currency"    gorm:"foreignKey:CurrencyId;references:UniqueId"     `
	// Datenano also has a text representation
	CurrencyId *string  `json:"currencyId" yaml:"currencyId"`
	Amount     *float64 `json:"amount" yaml:"amount"       `
	// Datenano also has a text representation
	LinkedTo *PriceTagEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

func (x *PriceTagVariations) RootObjectName() string {
	return "PriceTagEntity"
}

type PriceTagEntity struct {
	Visibility       *string               `json:"visibility,omitempty" yaml:"visibility"`
	WorkspaceId      *string               `json:"workspaceId,omitempty" yaml:"workspaceId"`
	LinkerId         *string               `json:"linkerId,omitempty" yaml:"linkerId"`
	ParentId         *string               `json:"parentId,omitempty" yaml:"parentId"`
	IsDeletable      *bool                 `json:"isDeletable,omitempty" yaml:"isDeletable" gorm:"default:true"`
	IsUpdatable      *bool                 `json:"isUpdatable,omitempty" yaml:"isUpdatable" gorm:"default:true"`
	UniqueId         string                `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
	UserId           *string               `json:"userId,omitempty" yaml:"userId"`
	Rank             int64                 `json:"rank,omitempty" gorm:"type:int;name:rank"`
	Updated          int64                 `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
	Created          int64                 `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
	CreatedFormatted string                `json:"createdFormatted,omitempty" sql:"-" gorm:"-"`
	UpdatedFormatted string                `json:"updatedFormatted,omitempty" sql:"-" gorm:"-"`
	Variations       []*PriceTagVariations `json:"variations" yaml:"variations"    gorm:"foreignKey:LinkerId;references:UniqueId"     `
	// Datenano also has a text representation
	Children []*PriceTagEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
	LinkedTo *PriceTagEntity   `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

var PriceTagPreloadRelations []string = []string{}
var PRICE_TAG_EVENT_CREATED = "priceTag.created"
var PRICE_TAG_EVENT_UPDATED = "priceTag.updated"
var PRICE_TAG_EVENT_DELETED = "priceTag.deleted"
var PRICE_TAG_EVENTS = []string{
	PRICE_TAG_EVENT_CREATED,
	PRICE_TAG_EVENT_UPDATED,
	PRICE_TAG_EVENT_DELETED,
}

type PriceTagFieldMap struct {
	Variations workspaces.TranslatedString `yaml:"variations"`
}

var PriceTagEntityMetaConfig map[string]int64 = map[string]int64{}
var PriceTagEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&PriceTagEntity{}))

func PriceTagVariationsActionCreate(
	dto *PriceTagVariations,
	query workspaces.QueryDSL,
) (*PriceTagVariations, *workspaces.IError) {
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
func PriceTagVariationsActionUpdate(
	query workspaces.QueryDSL,
	dto *PriceTagVariations,
) (*PriceTagVariations, *workspaces.IError) {
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
func PriceTagVariationsActionGetOne(
	query workspaces.QueryDSL,
) (*PriceTagVariations, *workspaces.IError) {
	refl := reflect.ValueOf(&PriceTagVariations{})
	item, err := workspaces.GetOneEntity[PriceTagVariations](query, refl)
	return item, err
}
func entityPriceTagFormatter(dto *PriceTagEntity, query workspaces.QueryDSL) {
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
func PriceTagMockEntity() *PriceTagEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &PriceTagEntity{}
	return entity
}
func PriceTagActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := PriceTagMockEntity()
		_, err := PriceTagActionCreate(entity, query)
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
func PriceTagActionSeederInit(query workspaces.QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*PriceTagEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &PriceTagEntity{
		Variations: []*PriceTagVariations{{}},
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
func PriceTagAssociationCreate(dto *PriceTagEntity, query workspaces.QueryDSL) error {
	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func PriceTagRelationContentCreate(dto *PriceTagEntity, query workspaces.QueryDSL) error {
	return nil
}
func PriceTagRelationContentUpdate(dto *PriceTagEntity, query workspaces.QueryDSL) error {
	return nil
}
func PriceTagPolyglotCreateHandler(dto *PriceTagEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
}

/**
 * This will be validating your entity fully. Important note is that, you add validate:* tag
 * in your entity, it will automatically work here. For slices inside entity, make sure you add
 * extra line of AppendSliceErrors, otherwise they won't be detected
 */
func PriceTagValidator(dto *PriceTagEntity, isPatch bool) *workspaces.IError {
	err := workspaces.CommonStructValidatorPointer(dto, isPatch)
	if dto != nil && dto.Variations != nil {
		workspaces.AppendSliceErrors(dto.Variations, isPatch, "variations", err)
	}
	return err
}
func PriceTagEntityPreSanitize(dto *PriceTagEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
func PriceTagEntityBeforeCreateAppend(dto *PriceTagEntity, query workspaces.QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = workspaces.UUID()
	}
	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId
	PriceTagRecursiveAddUniqueId(dto, query)
}
func PriceTagRecursiveAddUniqueId(dto *PriceTagEntity, query workspaces.QueryDSL) {
	if dto.Variations != nil && len(dto.Variations) > 0 {
		for index0 := range dto.Variations {
			if dto.Variations[index0].UniqueId == "" {
				dto.Variations[index0].UniqueId = workspaces.UUID()
			}
		}
	}
}
func PriceTagActionBatchCreateFn(dtos []*PriceTagEntity, query workspaces.QueryDSL) ([]*PriceTagEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*PriceTagEntity{}
		for _, item := range dtos {
			s, err := PriceTagActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil
}
func PriceTagDeleteEntireChildren(query workspaces.QueryDSL, dto *PriceTagEntity) *workspaces.IError {
	// intentionally removed this. It's hard to implement it, and probably wrong without
	// proper on delete cascade
	return nil
}
func PriceTagActionCreateFn(dto *PriceTagEntity, query workspaces.QueryDSL) (*PriceTagEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := PriceTagValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	PriceTagEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	PriceTagEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	PriceTagPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	PriceTagRelationContentCreate(dto, query)
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
	PriceTagAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(PRICE_TAG_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": workspaces.GetTypeString(&PriceTagEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})
	return dto, nil
}
func PriceTagActionGetOne(query workspaces.QueryDSL) (*PriceTagEntity, *workspaces.IError) {
	refl := reflect.ValueOf(&PriceTagEntity{})
	item, err := workspaces.GetOneEntity[PriceTagEntity](query, refl)
	entityPriceTagFormatter(item, query)
	return item, err
}
func PriceTagActionQuery(query workspaces.QueryDSL) ([]*PriceTagEntity, *workspaces.QueryResultMeta, error) {
	refl := reflect.ValueOf(&PriceTagEntity{})
	items, meta, err := workspaces.QueryEntitiesPointer[PriceTagEntity](query, refl)
	for _, item := range items {
		entityPriceTagFormatter(item, query)
	}
	return items, meta, err
}
func PriceTagUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *PriceTagEntity) (*PriceTagEntity, *workspaces.IError) {
	uniqueId := fields.UniqueId
	query.TriggerEventName = PRICE_TAG_EVENT_UPDATED
	PriceTagEntityPreSanitize(fields, query)
	var item PriceTagEntity
	q := dbref.
		Where(&PriceTagEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)
	err := q.UpdateColumns(fields).Error
	if err != nil {
		return nil, workspaces.GormErrorToIError(err)
	}
	query.Tx = dbref
	PriceTagRelationContentUpdate(fields, query)
	PriceTagPolyglotCreateHandler(fields, query)
	if ero := PriceTagDeleteEntireChildren(query, fields); ero != nil {
		return nil, ero
	}
	// @meta(update has many)
	if fields.Variations != nil {
		linkerId := uniqueId
		dbref.
			Where(&PriceTagVariations{LinkerId: &linkerId}).
			Delete(&PriceTagVariations{})
		for _, newItem := range fields.Variations {
			newItem.UniqueId = workspaces.UUID()
			newItem.LinkerId = &linkerId
			dbref.Create(&newItem)
		}
	}
	err = dbref.
		Preload(clause.Associations).
		Where(&PriceTagEntity{UniqueId: uniqueId}).
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
func PriceTagActionUpdateFn(query workspaces.QueryDSL, fields *PriceTagEntity) (*PriceTagEntity, *workspaces.IError) {
	if fields == nil {
		return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}
	// 1. Validate always
	if iError := PriceTagValidator(fields, true); iError != nil {
		return nil, iError
	}
	// Let's not add this. I am not sure of the consequences
	// PriceTagRecursiveAddUniqueId(fields, query)
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()
		var item *PriceTagEntity
		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			var err *workspaces.IError
			item, err = PriceTagUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}
		})
		return item, workspaces.CastToIError(vf)
	} else {
		dbref = query.Tx
		return PriceTagUpdateExec(dbref, query, fields)
	}
}

var PriceTagWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire pricetags ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
			ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRICE_TAG_DELETE},
		})
		count, _ := PriceTagActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}

func PriceTagActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&PriceTagEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_PRICE_TAG_DELETE}
	return workspaces.RemoveEntity[PriceTagEntity](query, refl)
}
func PriceTagActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error
	var count int64 = 0
	{
		subCount, subErr := workspaces.WipeCleanEntity[PriceTagVariations]()
		if subErr != nil {
			fmt.Println("Error while wiping 'PriceTagVariations'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	{
		subCount, subErr := workspaces.WipeCleanEntity[PriceTagEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'PriceTagEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
func PriceTagActionBulkUpdate(
	query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[PriceTagEntity]) (
	*workspaces.BulkRecordRequest[PriceTagEntity], *workspaces.IError,
) {
	result := []*PriceTagEntity{}
	err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := PriceTagActionUpdate(query, record)
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
func (x *PriceTagEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

var PriceTagEntityMeta = workspaces.TableMetaData{
	EntityName:    "PriceTag",
	ExportKey:     "price-tags",
	TableNameInDb: "fb_price-tag_entities",
	EntityObject:  &PriceTagEntity{},
	ExportStream:  PriceTagActionExportT,
	ImportQuery:   PriceTagActionImport,
}

func PriceTagActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[PriceTagEntity](query, PriceTagActionQuery, PriceTagPreloadRelations)
}
func PriceTagActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[PriceTagEntity](query, PriceTagActionQuery, PriceTagPreloadRelations)
}
func PriceTagActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content PriceTagEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := PriceTagActionCreate(&content, query)
	return err
}

var PriceTagCommonCliFlags = []cli.Flag{
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
	&cli.StringSliceFlag{
		Name:     "variations",
		Required: false,
		Usage:    "variations",
	},
}
var PriceTagCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{}
var PriceTagCommonCliFlagsOptional = []cli.Flag{
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
	&cli.StringSliceFlag{
		Name:     "variations",
		Required: false,
		Usage:    "variations",
	},
}
var PriceTagCreateCmd cli.Command = PRICE_TAG_ACTION_POST_ONE.ToCli()
var PriceTagCreateInteractiveCmd cli.Command = cli.Command{
	Name:  "ic",
	Usage: "Creates a new template, using requied fields in an interactive name",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "all",
			Usage: "Interactively asks for all inputs, not only required ones",
		},
	},
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
			ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRICE_TAG_CREATE},
		})
		entity := &PriceTagEntity{}
		for _, item := range PriceTagCommonInteractiveCliFlags {
			if !item.Required && c.Bool("all") == false {
				continue
			}
			result := workspaces.AskForInput(item.Name, "")
			workspaces.SetFieldString(entity, item.StructField, result)
		}
		if entity, err := PriceTagActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var PriceTagUpdateCmd cli.Command = cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Flags:   PriceTagCommonCliFlagsOptional,
	Usage:   "Updates a template by passing the parameters",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
			ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRICE_TAG_UPDATE},
		})
		entity := CastPriceTagFromCli(c)
		if entity, err := PriceTagActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
		return nil
	},
}

func (x *PriceTagEntity) FromCli(c *cli.Context) *PriceTagEntity {
	return CastPriceTagFromCli(c)
}
func CastPriceTagFromCli(c *cli.Context) *PriceTagEntity {
	template := &PriceTagEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
	return template
}
func PriceTagSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		PriceTagActionCreate,
		reflect.ValueOf(&PriceTagEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}
func PriceTagWriteQueryMock(ctx workspaces.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := PriceTagActionQuery(f)
		result := workspaces.QueryEntitySuccessResult(f, items, count)
		workspaces.WriteMockDataToFile(lang, "", "PriceTag", result)
	}
}

var PriceTagImportExportCommands = []cli.Command{
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
			query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
				ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRICE_TAG_CREATE},
			})
			PriceTagActionSeeder(query, c.Int("count"))
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
				Value: "price-tag-seeder.yml",
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
			query := workspaces.CommonCliQueryDSLBuilder(c)
			PriceTagActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "price-tag-seeder-price-tag.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of price-tags, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]PriceTagEntity{}
			workspaces.ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
	cli.Command{
		Name: "import",
		Flags: append(
			append(
				workspaces.CommonQueryFlags,
				&cli.StringFlag{
					Name:     "file",
					Usage:    "The address of file you want the csv be imported from",
					Required: true,
				}),
			PriceTagCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				PriceTagActionCreate,
				reflect.ValueOf(&PriceTagEntity{}).Elem(),
				c.String("file"),
				&workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRICE_TAG_CREATE},
				},
				func() PriceTagEntity {
					v := CastPriceTagFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
var PriceTagCliCommands []cli.Command = []cli.Command{
	PRICE_TAG_ACTION_QUERY.ToCli(),
	PRICE_TAG_ACTION_TABLE.ToCli(),
	PriceTagCreateCmd,
	PriceTagUpdateCmd,
	PriceTagCreateInteractiveCmd,
	PriceTagWipeCmd,
	workspaces.GetCommonRemoveQuery(reflect.ValueOf(&PriceTagEntity{}).Elem(), PriceTagActionRemove),
}

func PriceTagCliFn() cli.Command {
	PriceTagCliCommands = append(PriceTagCliCommands, PriceTagImportExportCommands...)
	return cli.Command{
		Name:        "priceTag",
		Description: "PriceTags module actions (sample module to handle complex entities)",
		Usage:       "Price tag is a definition of a price, in different currencies or regions",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: PriceTagCliCommands,
	}
}

var PRICE_TAG_ACTION_TABLE = workspaces.Module2Action{
	Name:          "table",
	ActionName:    "table",
	ActionAliases: []string{"t"},
	Flags:         workspaces.CommonQueryFlags,
	Description:   "Table formatted queries all of the entities in database based on the standard query format",
	Action:        PriceTagActionQuery,
	CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliTableCmd2(c,
			PriceTagActionQuery,
			security,
			reflect.ValueOf(&PriceTagEntity{}).Elem(),
		)
		return nil
	},
}
var PRICE_TAG_ACTION_QUERY = workspaces.Module2Action{
	Method: "GET",
	Url:    "/price-tags",
	SecurityModel: &workspaces.SecurityModel{
		ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRICE_TAG_QUERY},
	},
	Group: "priceTag",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			workspaces.HttpQueryEntity(c, PriceTagActionQuery)
		},
	},
	Format:         "QUERY",
	Action:         PriceTagActionQuery,
	ResponseEntity: &[]PriceTagEntity{},
	Out: workspaces.Module2ActionBody{
		Entity: "PriceTagEntity",
	},
	CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			PriceTagActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionName:    "query",
	ActionAliases: []string{"q"},
	Flags:         workspaces.CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var PRICE_TAG_ACTION_EXPORT = workspaces.Module2Action{
	Method: "GET",
	Url:    "/price-tags/export",
	SecurityModel: &workspaces.SecurityModel{
		ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRICE_TAG_QUERY},
	},
	Group: "priceTag",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			workspaces.HttpStreamFileChannel(c, PriceTagActionExport)
		},
	},
	Format:         "QUERY",
	Action:         PriceTagActionExport,
	ResponseEntity: &[]PriceTagEntity{},
	Out: workspaces.Module2ActionBody{
		Entity: "PriceTagEntity",
	},
}
var PRICE_TAG_ACTION_GET_ONE = workspaces.Module2Action{
	Method: "GET",
	Url:    "/price-tag/:uniqueId",
	SecurityModel: &workspaces.SecurityModel{
		ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRICE_TAG_QUERY},
	},
	Group: "priceTag",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			workspaces.HttpGetEntity(c, PriceTagActionGetOne)
		},
	},
	Format:         "GET_ONE",
	Action:         PriceTagActionGetOne,
	ResponseEntity: &PriceTagEntity{},
	Out: workspaces.Module2ActionBody{
		Entity: "PriceTagEntity",
	},
}
var PRICE_TAG_ACTION_POST_ONE = workspaces.Module2Action{
	ActionName:    "create",
	ActionAliases: []string{"c"},
	Description:   "Create new priceTag",
	Flags:         PriceTagCommonCliFlags,
	Method:        "POST",
	Url:           "/price-tag",
	SecurityModel: &workspaces.SecurityModel{
		ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRICE_TAG_CREATE},
	},
	Group: "priceTag",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			workspaces.HttpPostEntity(c, PriceTagActionCreate)
		},
	},
	CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		result, err := workspaces.CliPostEntity(c, PriceTagActionCreate, security)
		workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
		return err
	},
	Action:         PriceTagActionCreate,
	Format:         "POST_ONE",
	RequestEntity:  &PriceTagEntity{},
	ResponseEntity: &PriceTagEntity{},
	Out: workspaces.Module2ActionBody{
		Entity: "PriceTagEntity",
	},
	In: workspaces.Module2ActionBody{
		Entity: "PriceTagEntity",
	},
}
var PRICE_TAG_ACTION_PATCH = workspaces.Module2Action{
	ActionName:    "update",
	ActionAliases: []string{"u"},
	Flags:         PriceTagCommonCliFlagsOptional,
	Method:        "PATCH",
	Url:           "/price-tag",
	SecurityModel: &workspaces.SecurityModel{
		ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRICE_TAG_UPDATE},
	},
	Group: "priceTag",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			workspaces.HttpUpdateEntity(c, PriceTagActionUpdate)
		},
	},
	Action:         PriceTagActionUpdate,
	RequestEntity:  &PriceTagEntity{},
	ResponseEntity: &PriceTagEntity{},
	Format:         "PATCH_ONE",
	Out: workspaces.Module2ActionBody{
		Entity: "PriceTagEntity",
	},
	In: workspaces.Module2ActionBody{
		Entity: "PriceTagEntity",
	},
}
var PRICE_TAG_ACTION_PATCH_BULK = workspaces.Module2Action{
	Method: "PATCH",
	Url:    "/price-tags",
	SecurityModel: &workspaces.SecurityModel{
		ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRICE_TAG_UPDATE},
	},
	Group: "priceTag",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			workspaces.HttpUpdateEntities(c, PriceTagActionBulkUpdate)
		},
	},
	Action:         PriceTagActionBulkUpdate,
	Format:         "PATCH_BULK",
	RequestEntity:  &workspaces.BulkRecordRequest[PriceTagEntity]{},
	ResponseEntity: &workspaces.BulkRecordRequest[PriceTagEntity]{},
	Out: workspaces.Module2ActionBody{
		Entity: "PriceTagEntity",
	},
	In: workspaces.Module2ActionBody{
		Entity: "PriceTagEntity",
	},
}
var PRICE_TAG_ACTION_DELETE = workspaces.Module2Action{
	Method: "DELETE",
	Url:    "/price-tag",
	Format: "DELETE_DSL",
	SecurityModel: &workspaces.SecurityModel{
		ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRICE_TAG_DELETE},
	},
	Group: "priceTag",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			workspaces.HttpRemoveEntity(c, PriceTagActionRemove)
		},
	},
	Action:         PriceTagActionRemove,
	RequestEntity:  &workspaces.DeleteRequest{},
	ResponseEntity: &workspaces.DeleteResponse{},
	TargetEntity:   &PriceTagEntity{},
}
var PRICE_TAG_VARIATIONS_ACTION_PATCH = workspaces.Module2Action{
	Method: "PATCH",
	Url:    "/price-tag/:linkerId/variations/:uniqueId",
	SecurityModel: &workspaces.SecurityModel{
		ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRICE_TAG_UPDATE},
	},
	Group: "priceTag",
	Handlers: []gin.HandlerFunc{
		func(
			c *gin.Context,
		) {
			workspaces.HttpUpdateEntity(c, PriceTagVariationsActionUpdate)
		},
	},
	Action:         PriceTagVariationsActionUpdate,
	Format:         "PATCH_ONE",
	RequestEntity:  &PriceTagVariations{},
	ResponseEntity: &PriceTagVariations{},
	Out: workspaces.Module2ActionBody{
		Entity: "PriceTagVariations",
	},
	In: workspaces.Module2ActionBody{
		Entity: "PriceTagVariations",
	},
}
var PRICE_TAG_VARIATIONS_ACTION_GET = workspaces.Module2Action{
	Method: "GET",
	Url:    "/price-tag/variations/:linkerId/:uniqueId",
	SecurityModel: &workspaces.SecurityModel{
		ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRICE_TAG_QUERY},
	},
	Group: "priceTag",
	Handlers: []gin.HandlerFunc{
		func(
			c *gin.Context,
		) {
			workspaces.HttpGetEntity(c, PriceTagVariationsActionGetOne)
		},
	},
	Action:         PriceTagVariationsActionGetOne,
	Format:         "GET_ONE",
	ResponseEntity: &PriceTagVariations{},
	Out: workspaces.Module2ActionBody{
		Entity: "PriceTagVariations",
	},
}
var PRICE_TAG_VARIATIONS_ACTION_POST = workspaces.Module2Action{
	Method: "POST",
	Url:    "/price-tag/:linkerId/variations",
	SecurityModel: &workspaces.SecurityModel{
		ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRICE_TAG_CREATE},
	},
	Group: "priceTag",
	Handlers: []gin.HandlerFunc{
		func(
			c *gin.Context,
		) {
			workspaces.HttpPostEntity(c, PriceTagVariationsActionCreate)
		},
	},
	Action:         PriceTagVariationsActionCreate,
	Format:         "POST_ONE",
	RequestEntity:  &PriceTagVariations{},
	ResponseEntity: &PriceTagVariations{},
	Out: workspaces.Module2ActionBody{
		Entity: "PriceTagVariations",
	},
	In: workspaces.Module2ActionBody{
		Entity: "PriceTagVariations",
	},
}

/**
 *	Override this function on PriceTagEntityHttp.go,
 *	In order to add your own http
 **/
var AppendPriceTagRouter = func(r *[]workspaces.Module2Action) {}

func GetPriceTagModule2Actions() []workspaces.Module2Action {
	routes := []workspaces.Module2Action{
		PRICE_TAG_ACTION_QUERY,
		PRICE_TAG_ACTION_EXPORT,
		PRICE_TAG_ACTION_GET_ONE,
		PRICE_TAG_ACTION_POST_ONE,
		PRICE_TAG_ACTION_PATCH,
		PRICE_TAG_ACTION_PATCH_BULK,
		PRICE_TAG_ACTION_DELETE,
		PRICE_TAG_VARIATIONS_ACTION_PATCH,
		PRICE_TAG_VARIATIONS_ACTION_GET,
		PRICE_TAG_VARIATIONS_ACTION_POST,
	}
	// Append user defined functions
	AppendPriceTagRouter(&routes)
	return routes
}
func CreatePriceTagRouter(r *gin.Engine) []workspaces.Module2Action {
	httpRoutes := GetPriceTagModule2Actions()
	workspaces.CastRoutes(httpRoutes, r)
	workspaces.WriteHttpInformationToFile(&httpRoutes, PriceTagEntityJsonSchema, "price-tag-http", "currency")
	workspaces.WriteEntitySchema("PriceTagEntity", PriceTagEntityJsonSchema, "currency")
	return httpRoutes
}

var PERM_ROOT_PRICE_TAG_DELETE = workspaces.PermissionInfo{
	CompleteKey: "root/currency/price-tag/delete",
	Name:        "Delete price tag",
}
var PERM_ROOT_PRICE_TAG_CREATE = workspaces.PermissionInfo{
	CompleteKey: "root/currency/price-tag/create",
	Name:        "Create price tag",
}
var PERM_ROOT_PRICE_TAG_UPDATE = workspaces.PermissionInfo{
	CompleteKey: "root/currency/price-tag/update",
	Name:        "Update price tag",
}
var PERM_ROOT_PRICE_TAG_QUERY = workspaces.PermissionInfo{
	CompleteKey: "root/currency/price-tag/query",
	Name:        "Query price tag",
}
var PERM_ROOT_PRICE_TAG = workspaces.PermissionInfo{
	CompleteKey: "root/currency/price-tag/*",
	Name:        "Entire price tag actions (*)",
}
var ALL_PRICE_TAG_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_PRICE_TAG_DELETE,
	PERM_ROOT_PRICE_TAG_CREATE,
	PERM_ROOT_PRICE_TAG_UPDATE,
	PERM_ROOT_PRICE_TAG_QUERY,
	PERM_ROOT_PRICE_TAG,
}
var PriceTagEntityBundle = workspaces.EntityBundle{
	Permissions: ALL_PRICE_TAG_PERMISSIONS,
	CliCommands: []cli.Command{
		PriceTagCliFn(),
	},
	Actions: GetPriceTagModule2Actions(),
	AutoMigrationEntities: []interface{}{
		&PriceTagEntity{},
		&PriceTagVariations{},
	},
}
