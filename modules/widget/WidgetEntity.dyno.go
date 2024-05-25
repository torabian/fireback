package widget

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
	mocks "github.com/torabian/fireback/modules/widget/mocks/Widget"
	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var widgetSeedersFs *embed.FS = nil

func ResetWidgetSeeders(fs *embed.FS) {
	widgetSeedersFs = fs
}

type WidgetEntity struct {
	Visibility       *string `json:"visibility,omitempty" yaml:"visibility"`
	WorkspaceId      *string `json:"workspaceId,omitempty" yaml:"workspaceId"`
	LinkerId         *string `json:"linkerId,omitempty" yaml:"linkerId"`
	ParentId         *string `json:"parentId,omitempty" yaml:"parentId"`
	IsDeletable      *bool   `json:"isDeletable,omitempty" yaml:"isDeletable" gorm:"default:true"`
	IsUpdatable      *bool   `json:"isUpdatable,omitempty" yaml:"isUpdatable" gorm:"default:true"`
	UniqueId         string  `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
	UserId           *string `json:"userId,omitempty" yaml:"userId"`
	Rank             int64   `json:"rank,omitempty" gorm:"type:int;name:rank"`
	Updated          int64   `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
	Created          int64   `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
	CreatedFormatted string  `json:"createdFormatted,omitempty" sql:"-" gorm:"-"`
	UpdatedFormatted string  `json:"updatedFormatted,omitempty" sql:"-" gorm:"-"`
	Name             *string `json:"name" yaml:"name"        translate:"true" `
	// Datenano also has a text representation
	Family *string `json:"family" yaml:"family"       `
	// Datenano also has a text representation
	ProviderKey *string `json:"providerKey" yaml:"providerKey"       `
	// Datenano also has a text representation
	Translations []*WidgetEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
	Children     []*WidgetEntity         `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
	LinkedTo     *WidgetEntity           `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

var WidgetPreloadRelations []string = []string{}
var WIDGET_EVENT_CREATED = "widget.created"
var WIDGET_EVENT_UPDATED = "widget.updated"
var WIDGET_EVENT_DELETED = "widget.deleted"
var WIDGET_EVENTS = []string{
	WIDGET_EVENT_CREATED,
	WIDGET_EVENT_UPDATED,
	WIDGET_EVENT_DELETED,
}

type WidgetFieldMap struct {
	Name        workspaces.TranslatedString `yaml:"name"`
	Family      workspaces.TranslatedString `yaml:"family"`
	ProviderKey workspaces.TranslatedString `yaml:"providerKey"`
}

var WidgetEntityMetaConfig map[string]int64 = map[string]int64{}
var WidgetEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&WidgetEntity{}))

type WidgetEntityPolyglot struct {
	LinkerId   string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
	LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
	Name       string `yaml:"name" json:"name"`
}

func entityWidgetFormatter(dto *WidgetEntity, query workspaces.QueryDSL) {
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
func WidgetMockEntity() *WidgetEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &WidgetEntity{
		Name:        &stringHolder,
		Family:      &stringHolder,
		ProviderKey: &stringHolder,
	}
	return entity
}
func WidgetActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := WidgetMockEntity()
		_, err := WidgetActionCreate(entity, query)
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
func (x *WidgetEntity) GetNameTranslated(language string) string {
	if x.Translations != nil && len(x.Translations) > 0 {
		for _, item := range x.Translations {
			if item.LanguageId == language {
				return item.Name
			}
		}
	}
	return ""
}
func WidgetActionSeederInit(query workspaces.QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*WidgetEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &WidgetEntity{
		Name:        &tildaRef,
		Family:      &tildaRef,
		ProviderKey: &tildaRef,
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
func WidgetAssociationCreate(dto *WidgetEntity, query workspaces.QueryDSL) error {
	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func WidgetRelationContentCreate(dto *WidgetEntity, query workspaces.QueryDSL) error {
	return nil
}
func WidgetRelationContentUpdate(dto *WidgetEntity, query workspaces.QueryDSL) error {
	return nil
}
func WidgetPolyglotCreateHandler(dto *WidgetEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
	workspaces.PolyglotCreateHandler(dto, &WidgetEntityPolyglot{}, query)
}

/**
 * This will be validating your entity fully. Important note is that, you add validate:* tag
 * in your entity, it will automatically work here. For slices inside entity, make sure you add
 * extra line of AppendSliceErrors, otherwise they won't be detected
 */
func WidgetValidator(dto *WidgetEntity, isPatch bool) *workspaces.IError {
	err := workspaces.CommonStructValidatorPointer(dto, isPatch)
	return err
}
func WidgetEntityPreSanitize(dto *WidgetEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
func WidgetEntityBeforeCreateAppend(dto *WidgetEntity, query workspaces.QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = workspaces.UUID()
	}
	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId
	WidgetRecursiveAddUniqueId(dto, query)
}
func WidgetRecursiveAddUniqueId(dto *WidgetEntity, query workspaces.QueryDSL) {
}
func WidgetActionBatchCreateFn(dtos []*WidgetEntity, query workspaces.QueryDSL) ([]*WidgetEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*WidgetEntity{}
		for _, item := range dtos {
			s, err := WidgetActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil
}
func WidgetDeleteEntireChildren(query workspaces.QueryDSL, dto *WidgetEntity) *workspaces.IError {
	// intentionally removed this. It's hard to implement it, and probably wrong without
	// proper on delete cascade
	return nil
}
func WidgetActionCreateFn(dto *WidgetEntity, query workspaces.QueryDSL) (*WidgetEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := WidgetValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	WidgetEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	WidgetEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	WidgetPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	WidgetRelationContentCreate(dto, query)
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
	WidgetAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(WIDGET_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": workspaces.GetTypeString(&WidgetEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})
	return dto, nil
}
func WidgetActionGetOne(query workspaces.QueryDSL) (*WidgetEntity, *workspaces.IError) {
	refl := reflect.ValueOf(&WidgetEntity{})
	item, err := workspaces.GetOneEntity[WidgetEntity](query, refl)
	entityWidgetFormatter(item, query)
	return item, err
}
func WidgetActionQuery(query workspaces.QueryDSL) ([]*WidgetEntity, *workspaces.QueryResultMeta, error) {
	refl := reflect.ValueOf(&WidgetEntity{})
	items, meta, err := workspaces.QueryEntitiesPointer[WidgetEntity](query, refl)
	for _, item := range items {
		entityWidgetFormatter(item, query)
	}
	return items, meta, err
}
func WidgetUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *WidgetEntity) (*WidgetEntity, *workspaces.IError) {
	uniqueId := fields.UniqueId
	query.TriggerEventName = WIDGET_EVENT_UPDATED
	WidgetEntityPreSanitize(fields, query)
	var item WidgetEntity
	q := dbref.
		Where(&WidgetEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)
	err := q.UpdateColumns(fields).Error
	if err != nil {
		return nil, workspaces.GormErrorToIError(err)
	}
	query.Tx = dbref
	WidgetRelationContentUpdate(fields, query)
	WidgetPolyglotCreateHandler(fields, query)
	if ero := WidgetDeleteEntireChildren(query, fields); ero != nil {
		return nil, ero
	}
	// @meta(update has many)
	err = dbref.
		Preload(clause.Associations).
		Where(&WidgetEntity{UniqueId: uniqueId}).
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
func WidgetActionUpdateFn(query workspaces.QueryDSL, fields *WidgetEntity) (*WidgetEntity, *workspaces.IError) {
	if fields == nil {
		return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}
	// 1. Validate always
	if iError := WidgetValidator(fields, true); iError != nil {
		return nil, iError
	}
	// Let's not add this. I am not sure of the consequences
	// WidgetRecursiveAddUniqueId(fields, query)
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()
		var item *WidgetEntity
		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			var err *workspaces.IError
			item, err = WidgetUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}
		})
		return item, workspaces.CastToIError(vf)
	} else {
		dbref = query.Tx
		return WidgetUpdateExec(dbref, query, fields)
	}
}

var WidgetWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire widgets ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
			ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_WIDGET_DELETE},
		})
		count, _ := WidgetActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}

func WidgetActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&WidgetEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_WIDGET_DELETE}
	return workspaces.RemoveEntity[WidgetEntity](query, refl)
}
func WidgetActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error
	var count int64 = 0
	{
		subCount, subErr := workspaces.WipeCleanEntity[WidgetEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'WidgetEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
func WidgetActionBulkUpdate(
	query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[WidgetEntity]) (
	*workspaces.BulkRecordRequest[WidgetEntity], *workspaces.IError,
) {
	result := []*WidgetEntity{}
	err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := WidgetActionUpdate(query, record)
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
func (x *WidgetEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

var WidgetEntityMeta = workspaces.TableMetaData{
	EntityName:    "Widget",
	ExportKey:     "widgets",
	TableNameInDb: "fb_widget_entities",
	EntityObject:  &WidgetEntity{},
	ExportStream:  WidgetActionExportT,
	ImportQuery:   WidgetActionImport,
}

func WidgetActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[WidgetEntity](query, WidgetActionQuery, WidgetPreloadRelations)
}
func WidgetActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[WidgetEntity](query, WidgetActionQuery, WidgetPreloadRelations)
}
func WidgetActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content WidgetEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := WidgetActionCreate(&content, query)
	return err
}

var WidgetCommonCliFlags = []cli.Flag{
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
		Name:     "family",
		Required: false,
		Usage:    "family",
	},
	&cli.StringFlag{
		Name:     "provider-key",
		Required: false,
		Usage:    "providerKey",
	},
}
var WidgetCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:        "name",
		StructField: "Name",
		Required:    false,
		Usage:       "name",
		Type:        "string",
	},
	{
		Name:        "family",
		StructField: "Family",
		Required:    false,
		Usage:       "family",
		Type:        "string",
	},
	{
		Name:        "providerKey",
		StructField: "ProviderKey",
		Required:    false,
		Usage:       "providerKey",
		Type:        "string",
	},
}
var WidgetCommonCliFlagsOptional = []cli.Flag{
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
		Name:     "family",
		Required: false,
		Usage:    "family",
	},
	&cli.StringFlag{
		Name:     "provider-key",
		Required: false,
		Usage:    "providerKey",
	},
}
var WidgetCreateCmd cli.Command = WIDGET_ACTION_POST_ONE.ToCli()
var WidgetCreateInteractiveCmd cli.Command = cli.Command{
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
			ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_WIDGET_CREATE},
		})
		entity := &WidgetEntity{}
		for _, item := range WidgetCommonInteractiveCliFlags {
			if !item.Required && c.Bool("all") == false {
				continue
			}
			result := workspaces.AskForInput(item.Name, "")
			workspaces.SetFieldString(entity, item.StructField, result)
		}
		if entity, err := WidgetActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var WidgetUpdateCmd cli.Command = cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Flags:   WidgetCommonCliFlagsOptional,
	Usage:   "Updates a template by passing the parameters",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
			ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_WIDGET_UPDATE},
		})
		entity := CastWidgetFromCli(c)
		if entity, err := WidgetActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
		return nil
	},
}

func (x *WidgetEntity) FromCli(c *cli.Context) *WidgetEntity {
	return CastWidgetFromCli(c)
}
func CastWidgetFromCli(c *cli.Context) *WidgetEntity {
	template := &WidgetEntity{}
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
	if c.IsSet("family") {
		value := c.String("family")
		template.Family = &value
	}
	if c.IsSet("provider-key") {
		value := c.String("provider-key")
		template.ProviderKey = &value
	}
	return template
}
func WidgetSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		WidgetActionCreate,
		reflect.ValueOf(&WidgetEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}
func WidgetImportMocks() {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		WidgetActionCreate,
		reflect.ValueOf(&WidgetEntity{}).Elem(),
		&mocks.ViewsFs,
		[]string{},
		false,
	)
}
func WidgetWriteQueryMock(ctx workspaces.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := WidgetActionQuery(f)
		result := workspaces.QueryEntitySuccessResult(f, items, count)
		workspaces.WriteMockDataToFile(lang, "", "Widget", result)
	}
}

var WidgetImportExportCommands = []cli.Command{
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
				ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_WIDGET_CREATE},
			})
			WidgetActionSeeder(query, c.Int("count"))
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
				Value: "widget-seeder.yml",
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
			WidgetActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "widget-seeder-widget.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of widgets, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]WidgetEntity{}
			workspaces.ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
	cli.Command{
		Name:  "mocks",
		Usage: "Prints the list of mocks",
		Action: func(c *cli.Context) error {
			if entity, err := workspaces.GetSeederFilenames(&mocks.ViewsFs, ""); err != nil {
				fmt.Println(err.Error())
			} else {
				f, _ := json.MarshalIndent(entity, "", "  ")
				fmt.Println(string(f))
			}
			return nil
		},
	},
	cli.Command{
		Name:  "msync",
		Usage: "Tries to sync mocks into the system",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportEmbedCmd(c,
				WidgetActionCreate,
				reflect.ValueOf(&WidgetEntity{}).Elem(),
				&mocks.ViewsFs,
			)
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
			WidgetCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				WidgetActionCreate,
				reflect.ValueOf(&WidgetEntity{}).Elem(),
				c.String("file"),
				&workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_WIDGET_CREATE},
				},
				func() WidgetEntity {
					v := CastWidgetFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
var WidgetCliCommands []cli.Command = []cli.Command{
	WIDGET_ACTION_QUERY.ToCli(),
	WIDGET_ACTION_TABLE.ToCli(),
	WidgetCreateCmd,
	WidgetUpdateCmd,
	WidgetCreateInteractiveCmd,
	WidgetWipeCmd,
	workspaces.GetCommonRemoveQuery(reflect.ValueOf(&WidgetEntity{}).Elem(), WidgetActionRemove),
}

func WidgetCliFn() cli.Command {
	WidgetCliCommands = append(WidgetCliCommands, WidgetImportExportCommands...)
	return cli.Command{
		Name:        "widget",
		Description: "Widgets module actions (sample module to handle complex entities)",
		Usage:       "Widget is an item which can be placed on a widget area, such as weather widget",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: WidgetCliCommands,
	}
}

var WIDGET_ACTION_TABLE = workspaces.Module2Action{
	Name:          "table",
	ActionName:    "table",
	ActionAliases: []string{"t"},
	Flags:         workspaces.CommonQueryFlags,
	Description:   "Table formatted queries all of the entities in database based on the standard query format",
	Action:        WidgetActionQuery,
	CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliTableCmd2(c,
			WidgetActionQuery,
			security,
			reflect.ValueOf(&WidgetEntity{}).Elem(),
		)
		return nil
	},
}
var WIDGET_ACTION_QUERY = workspaces.Module2Action{
	Method: "GET",
	Url:    "/widgets",
	SecurityModel: &workspaces.SecurityModel{
		ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_WIDGET_QUERY},
	},
	Group: "widget",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			workspaces.HttpQueryEntity(c, WidgetActionQuery)
		},
	},
	Format:         "QUERY",
	Action:         WidgetActionQuery,
	ResponseEntity: &[]WidgetEntity{},
	Out: workspaces.Module2ActionBody{
		Entity: "WidgetEntity",
	},
	CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			WidgetActionQuery,
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
var WIDGET_ACTION_EXPORT = workspaces.Module2Action{
	Method: "GET",
	Url:    "/widgets/export",
	SecurityModel: &workspaces.SecurityModel{
		ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_WIDGET_QUERY},
	},
	Group: "widget",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			workspaces.HttpStreamFileChannel(c, WidgetActionExport)
		},
	},
	Format:         "QUERY",
	Action:         WidgetActionExport,
	ResponseEntity: &[]WidgetEntity{},
	Out: workspaces.Module2ActionBody{
		Entity: "WidgetEntity",
	},
}
var WIDGET_ACTION_GET_ONE = workspaces.Module2Action{
	Method: "GET",
	Url:    "/widget/:uniqueId",
	SecurityModel: &workspaces.SecurityModel{
		ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_WIDGET_QUERY},
	},
	Group: "widget",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			workspaces.HttpGetEntity(c, WidgetActionGetOne)
		},
	},
	Format:         "GET_ONE",
	Action:         WidgetActionGetOne,
	ResponseEntity: &WidgetEntity{},
	Out: workspaces.Module2ActionBody{
		Entity: "WidgetEntity",
	},
}
var WIDGET_ACTION_POST_ONE = workspaces.Module2Action{
	ActionName:    "create",
	ActionAliases: []string{"c"},
	Description:   "Create new widget",
	Flags:         WidgetCommonCliFlags,
	Method:        "POST",
	Url:           "/widget",
	SecurityModel: &workspaces.SecurityModel{
		ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_WIDGET_CREATE},
	},
	Group: "widget",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			workspaces.HttpPostEntity(c, WidgetActionCreate)
		},
	},
	CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		result, err := workspaces.CliPostEntity(c, WidgetActionCreate, security)
		workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
		return err
	},
	Action:         WidgetActionCreate,
	Format:         "POST_ONE",
	RequestEntity:  &WidgetEntity{},
	ResponseEntity: &WidgetEntity{},
	Out: workspaces.Module2ActionBody{
		Entity: "WidgetEntity",
	},
	In: workspaces.Module2ActionBody{
		Entity: "WidgetEntity",
	},
}
var WIDGET_ACTION_PATCH = workspaces.Module2Action{
	ActionName:    "update",
	ActionAliases: []string{"u"},
	Flags:         WidgetCommonCliFlagsOptional,
	Method:        "PATCH",
	Url:           "/widget",
	SecurityModel: &workspaces.SecurityModel{
		ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_WIDGET_UPDATE},
	},
	Group: "widget",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			workspaces.HttpUpdateEntity(c, WidgetActionUpdate)
		},
	},
	Action:         WidgetActionUpdate,
	RequestEntity:  &WidgetEntity{},
	ResponseEntity: &WidgetEntity{},
	Format:         "PATCH_ONE",
	Out: workspaces.Module2ActionBody{
		Entity: "WidgetEntity",
	},
	In: workspaces.Module2ActionBody{
		Entity: "WidgetEntity",
	},
}
var WIDGET_ACTION_PATCH_BULK = workspaces.Module2Action{
	Method: "PATCH",
	Url:    "/widgets",
	SecurityModel: &workspaces.SecurityModel{
		ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_WIDGET_UPDATE},
	},
	Group: "widget",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			workspaces.HttpUpdateEntities(c, WidgetActionBulkUpdate)
		},
	},
	Action:         WidgetActionBulkUpdate,
	Format:         "PATCH_BULK",
	RequestEntity:  &workspaces.BulkRecordRequest[WidgetEntity]{},
	ResponseEntity: &workspaces.BulkRecordRequest[WidgetEntity]{},
	Out: workspaces.Module2ActionBody{
		Entity: "WidgetEntity",
	},
	In: workspaces.Module2ActionBody{
		Entity: "WidgetEntity",
	},
}
var WIDGET_ACTION_DELETE = workspaces.Module2Action{
	Method: "DELETE",
	Url:    "/widget",
	Format: "DELETE_DSL",
	SecurityModel: &workspaces.SecurityModel{
		ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_WIDGET_DELETE},
	},
	Group: "widget",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			workspaces.HttpRemoveEntity(c, WidgetActionRemove)
		},
	},
	Action:         WidgetActionRemove,
	RequestEntity:  &workspaces.DeleteRequest{},
	ResponseEntity: &workspaces.DeleteResponse{},
	TargetEntity:   &WidgetEntity{},
}

/**
 *	Override this function on WidgetEntityHttp.go,
 *	In order to add your own http
 **/
var AppendWidgetRouter = func(r *[]workspaces.Module2Action) {}

func GetWidgetModule2Actions() []workspaces.Module2Action {
	routes := []workspaces.Module2Action{
		WIDGET_ACTION_QUERY,
		WIDGET_ACTION_EXPORT,
		WIDGET_ACTION_GET_ONE,
		WIDGET_ACTION_POST_ONE,
		WIDGET_ACTION_PATCH,
		WIDGET_ACTION_PATCH_BULK,
		WIDGET_ACTION_DELETE,
	}
	// Append user defined functions
	AppendWidgetRouter(&routes)
	return routes
}
func CreateWidgetRouter(r *gin.Engine) []workspaces.Module2Action {
	httpRoutes := GetWidgetModule2Actions()
	workspaces.CastRoutes(httpRoutes, r)
	workspaces.WriteHttpInformationToFile(&httpRoutes, WidgetEntityJsonSchema, "widget-http", "widget")
	workspaces.WriteEntitySchema("WidgetEntity", WidgetEntityJsonSchema, "widget")
	return httpRoutes
}

var PERM_ROOT_WIDGET_DELETE = workspaces.PermissionInfo{
	CompleteKey: "root/widget/widget/delete",
	Name:        "Delete widget",
}
var PERM_ROOT_WIDGET_CREATE = workspaces.PermissionInfo{
	CompleteKey: "root/widget/widget/create",
	Name:        "Create widget",
}
var PERM_ROOT_WIDGET_UPDATE = workspaces.PermissionInfo{
	CompleteKey: "root/widget/widget/update",
	Name:        "Update widget",
}
var PERM_ROOT_WIDGET_QUERY = workspaces.PermissionInfo{
	CompleteKey: "root/widget/widget/query",
	Name:        "Query widget",
}
var PERM_ROOT_WIDGET = workspaces.PermissionInfo{
	CompleteKey: "root/widget/widget/*",
	Name:        "Entire widget actions (*)",
}
var ALL_WIDGET_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_WIDGET_DELETE,
	PERM_ROOT_WIDGET_CREATE,
	PERM_ROOT_WIDGET_UPDATE,
	PERM_ROOT_WIDGET_QUERY,
	PERM_ROOT_WIDGET,
}
var WidgetEntityBundle = workspaces.EntityBundle{
	Permissions: ALL_WIDGET_PERMISSIONS,
	CliCommands: []cli.Command{
		WidgetCliFn(),
	},
	Actions: GetWidgetModule2Actions(),
	AutoMigrationEntities: []interface{}{
		&WidgetEntity{},
	},
}
