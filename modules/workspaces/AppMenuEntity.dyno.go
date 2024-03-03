package workspaces

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
	metas "github.com/torabian/fireback/modules/workspaces/metas"
	queries "github.com/torabian/fireback/modules/workspaces/queries"
	seeders "github.com/torabian/fireback/modules/workspaces/seeders/AppMenu"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AppMenuEntity struct {
	Visibility       *string `json:"visibility,omitempty" yaml:"visibility"`
	WorkspaceId      *string `json:"workspaceId,omitempty" yaml:"workspaceId"`
	LinkerId         *string `json:"linkerId,omitempty" yaml:"linkerId"`
	ParentId         *string `json:"parentId,omitempty" yaml:"parentId"`
	UniqueId         string  `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
	UserId           *string `json:"userId,omitempty" yaml:"userId"`
	Rank             int64   `json:"rank,omitempty" gorm:"type:int;name:rank"`
	Updated          int64   `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
	Created          int64   `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
	CreatedFormatted string  `json:"createdFormatted,omitempty" sql:"-"`
	UpdatedFormatted string  `json:"updatedFormatted,omitempty" sql:"-"`
	Href             *string `json:"href" yaml:"href"       `
	// Datenano also has a text representation
	Icon *string `json:"icon" yaml:"icon"       `
	// Datenano also has a text representation
	Label *string `json:"label" yaml:"label"        translate:"true" `
	// Datenano also has a text representation
	ActiveMatcher *string `json:"activeMatcher" yaml:"activeMatcher"       `
	// Datenano also has a text representation
	ApplyType *string `json:"applyType" yaml:"applyType"       `
	// Datenano also has a text representation
	Capability *CapabilityEntity `json:"capability" yaml:"capability"    gorm:"foreignKey:CapabilityId;references:UniqueId"     `
	// Datenano also has a text representation
	CapabilityId *string                  `json:"capabilityId" yaml:"capabilityId"`
	Translations []*AppMenuEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
	Children     []*AppMenuEntity         `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
	LinkedTo     *AppMenuEntity           `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

var AppMenuPreloadRelations []string = []string{}
var APPMENU_EVENT_CREATED = "appMenu.created"
var APPMENU_EVENT_UPDATED = "appMenu.updated"
var APPMENU_EVENT_DELETED = "appMenu.deleted"
var APPMENU_EVENTS = []string{
	APPMENU_EVENT_CREATED,
	APPMENU_EVENT_UPDATED,
	APPMENU_EVENT_DELETED,
}

type AppMenuFieldMap struct {
	Href          TranslatedString `yaml:"href"`
	Icon          TranslatedString `yaml:"icon"`
	Label         TranslatedString `yaml:"label"`
	ActiveMatcher TranslatedString `yaml:"activeMatcher"`
	ApplyType     TranslatedString `yaml:"applyType"`
	Capability    TranslatedString `yaml:"capability"`
}

var AppMenuEntityMetaConfig map[string]int64 = map[string]int64{}
var AppMenuEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&AppMenuEntity{}))

type AppMenuEntityPolyglot struct {
	LinkerId   string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
	LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
	Label      string `yaml:"label" json:"label"`
}

func entityAppMenuFormatter(dto *AppMenuEntity, query QueryDSL) {
	if dto == nil {
		return
	}
	if dto.Created > 0 {
		dto.CreatedFormatted = FormatDateBasedOnQuery(dto.Created, query)
	}
	if dto.Updated > 0 {
		dto.CreatedFormatted = FormatDateBasedOnQuery(dto.Updated, query)
	}
}
func AppMenuMockEntity() *AppMenuEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &AppMenuEntity{
		Href:          &stringHolder,
		Icon:          &stringHolder,
		Label:         &stringHolder,
		ActiveMatcher: &stringHolder,
		ApplyType:     &stringHolder,
	}
	return entity
}
func AppMenuActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := AppMenuMockEntity()
		_, err := AppMenuActionCreate(entity, query)
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
func (x *AppMenuEntity) GetLabelTranslated(language string) string {
	if x.Translations != nil && len(x.Translations) > 0 {
		for _, item := range x.Translations {
			if item.LanguageId == language {
				return item.Label
			}
		}
	}
	return ""
}
func AppMenuActionSeederInit(query QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*AppMenuEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &AppMenuEntity{
		Href:          &tildaRef,
		Icon:          &tildaRef,
		Label:         &tildaRef,
		ActiveMatcher: &tildaRef,
		ApplyType:     &tildaRef,
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
func AppMenuAssociationCreate(dto *AppMenuEntity, query QueryDSL) error {
	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func AppMenuRelationContentCreate(dto *AppMenuEntity, query QueryDSL) error {
	return nil
}
func AppMenuRelationContentUpdate(dto *AppMenuEntity, query QueryDSL) error {
	return nil
}
func AppMenuPolyglotCreateHandler(dto *AppMenuEntity, query QueryDSL) {
	if dto == nil {
		return
	}
	PolyglotCreateHandler(dto, &AppMenuEntityPolyglot{}, query)
}

/**
 * This will be validating your entity fully. Important note is that, you add validate:* tag
 * in your entity, it will automatically work here. For slices inside entity, make sure you add
 * extra line of AppendSliceErrors, otherwise they won't be detected
 */
func AppMenuValidator(dto *AppMenuEntity, isPatch bool) *IError {
	err := CommonStructValidatorPointer(dto, isPatch)
	return err
}
func AppMenuEntityPreSanitize(dto *AppMenuEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
func AppMenuEntityBeforeCreateAppend(dto *AppMenuEntity, query QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = UUID()
	}
	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId
	AppMenuRecursiveAddUniqueId(dto, query)
}
func AppMenuRecursiveAddUniqueId(dto *AppMenuEntity, query QueryDSL) {
}
func AppMenuActionBatchCreateFn(dtos []*AppMenuEntity, query QueryDSL) ([]*AppMenuEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*AppMenuEntity{}
		for _, item := range dtos {
			s, err := AppMenuActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil
}
func AppMenuActionCreateFn(dto *AppMenuEntity, query QueryDSL) (*AppMenuEntity, *IError) {
	// 1. Validate always
	if iError := AppMenuValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	AppMenuEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	AppMenuEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	AppMenuPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	AppMenuRelationContentCreate(dto, query)
	// 4. Create the entity
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = GetDbRef()
	} else {
		dbref = query.Tx
	}
	query.Tx = dbref
	err := dbref.Create(&dto).Error
	if err != nil {
		err := GormErrorToIError(err)
		return dto, err
	}
	// 5. Create sub entities, objects or arrays, association to other entities
	AppMenuAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(APPMENU_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": GetTypeString(&AppMenuEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})
	return dto, nil
}
func AppMenuActionGetOne(query QueryDSL) (*AppMenuEntity, *IError) {
	refl := reflect.ValueOf(&AppMenuEntity{})
	item, err := GetOneEntity[AppMenuEntity](query, refl)
	entityAppMenuFormatter(item, query)
	return item, err
}
func AppMenuActionQuery(query QueryDSL) ([]*AppMenuEntity, *QueryResultMeta, error) {
	refl := reflect.ValueOf(&AppMenuEntity{})
	items, meta, err := QueryEntitiesPointer[AppMenuEntity](query, refl)
	for _, item := range items {
		entityAppMenuFormatter(item, query)
	}
	return items, meta, err
}
func (dto *AppMenuEntity) Size() int {
	var size int = len(dto.Children)
	for _, c := range dto.Children {
		size += c.Size()
	}
	return size
}
func (dto *AppMenuEntity) Add(nodes ...*AppMenuEntity) bool {
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
func AppMenuActionCommonPivotQuery(query QueryDSL) ([]*PivotResult, *QueryResultMeta, error) {
	items, meta, err := UnsafeQuerySqlFromFs[PivotResult](
		&queries.QueriesFs, "AppMenuCommonPivot.sqlite.dyno", query,
	)
	return items, meta, err
}
func AppMenuActionCteQuery(query QueryDSL) ([]*AppMenuEntity, *QueryResultMeta, error) {
	items, meta, err := UnsafeQuerySqlFromFs[AppMenuEntity](
		&queries.QueriesFs, "AppMenuCTE.sqlite.dyno", query,
	)
	for _, item := range items {
		entityAppMenuFormatter(item, query)
	}
	var tree []*AppMenuEntity
	for _, item := range items {
		if item.ParentId == nil {
			root := item
			root.Add(items...)
			tree = append(tree, root)
		}
	}
	return tree, meta, err
}
func AppMenuUpdateExec(dbref *gorm.DB, query QueryDSL, fields *AppMenuEntity) (*AppMenuEntity, *IError) {
	uniqueId := fields.UniqueId
	query.TriggerEventName = APPMENU_EVENT_UPDATED
	AppMenuEntityPreSanitize(fields, query)
	var item AppMenuEntity
	q := dbref.
		Where(&AppMenuEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)
	err := q.UpdateColumns(fields).Error
	if err != nil {
		return nil, GormErrorToIError(err)
	}
	query.Tx = dbref
	AppMenuRelationContentUpdate(fields, query)
	AppMenuPolyglotCreateHandler(fields, query)
	// @meta(update has many)
	err = dbref.
		Preload(clause.Associations).
		Where(&AppMenuEntity{UniqueId: uniqueId}).
		First(&item).Error
	event.MustFire(query.TriggerEventName, event.M{
		"entity":   &item,
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	if err != nil {
		return &item, GormErrorToIError(err)
	}
	return &item, nil
}
func AppMenuActionUpdateFn(query QueryDSL, fields *AppMenuEntity) (*AppMenuEntity, *IError) {
	if fields == nil {
		return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}
	// 1. Validate always
	if iError := AppMenuValidator(fields, true); iError != nil {
		return nil, iError
	}
	AppMenuRecursiveAddUniqueId(fields, query)
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = GetDbRef()
		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			_, err := AppMenuUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}
		})
		return nil, CastToIError(vf)
	} else {
		dbref = query.Tx
		return AppMenuUpdateExec(dbref, query, fields)
	}
}

var AppMenuWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire appmenus ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilder(c)
		count, _ := AppMenuActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}

func AppMenuActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&AppMenuEntity{})
	query.ActionRequires = []string{PERM_ROOT_APPMENU_DELETE}
	return RemoveEntity[AppMenuEntity](query, refl)
}
func AppMenuActionWipeClean(query QueryDSL) (int64, error) {
	var err error
	var count int64 = 0
	{
		subCount, subErr := WipeCleanEntity[AppMenuEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'AppMenuEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
func AppMenuActionBulkUpdate(
	query QueryDSL, dto *BulkRecordRequest[AppMenuEntity]) (
	*BulkRecordRequest[AppMenuEntity], *IError,
) {
	result := []*AppMenuEntity{}
	err := GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := AppMenuActionUpdate(query, record)
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
	return nil, err.(*IError)
}
func (x *AppMenuEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

var AppMenuEntityMeta = TableMetaData{
	EntityName:    "AppMenu",
	ExportKey:     "app-menus",
	TableNameInDb: "fb_appmenu_entities",
	EntityObject:  &AppMenuEntity{},
	ExportStream:  AppMenuActionExportT,
	ImportQuery:   AppMenuActionImport,
}

func AppMenuActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[AppMenuEntity](query, AppMenuActionQuery, AppMenuPreloadRelations)
}
func AppMenuActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[AppMenuEntity](query, AppMenuActionQuery, AppMenuPreloadRelations)
}
func AppMenuActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content AppMenuEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := AppMenuActionCreate(&content, query)
	return err
}

var AppMenuCommonCliFlags = []cli.Flag{
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
		Name:     "href",
		Required: false,
		Usage:    "href",
	},
	&cli.StringFlag{
		Name:     "icon",
		Required: false,
		Usage:    "icon",
	},
	&cli.StringFlag{
		Name:     "label",
		Required: false,
		Usage:    "label",
	},
	&cli.StringFlag{
		Name:     "active-matcher",
		Required: false,
		Usage:    "activeMatcher",
	},
	&cli.StringFlag{
		Name:     "apply-type",
		Required: false,
		Usage:    "applyType",
	},
	&cli.StringFlag{
		Name:     "capability-id",
		Required: false,
		Usage:    "capability",
	},
}
var AppMenuCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:        "href",
		StructField: "Href",
		Required:    false,
		Usage:       "href",
		Type:        "string",
	},
	{
		Name:        "icon",
		StructField: "Icon",
		Required:    false,
		Usage:       "icon",
		Type:        "string",
	},
	{
		Name:        "label",
		StructField: "Label",
		Required:    false,
		Usage:       "label",
		Type:        "string",
	},
	{
		Name:        "activeMatcher",
		StructField: "ActiveMatcher",
		Required:    false,
		Usage:       "activeMatcher",
		Type:        "string",
	},
	{
		Name:        "applyType",
		StructField: "ApplyType",
		Required:    false,
		Usage:       "applyType",
		Type:        "string",
	},
}
var AppMenuCommonCliFlagsOptional = []cli.Flag{
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
		Name:     "href",
		Required: false,
		Usage:    "href",
	},
	&cli.StringFlag{
		Name:     "icon",
		Required: false,
		Usage:    "icon",
	},
	&cli.StringFlag{
		Name:     "label",
		Required: false,
		Usage:    "label",
	},
	&cli.StringFlag{
		Name:     "active-matcher",
		Required: false,
		Usage:    "activeMatcher",
	},
	&cli.StringFlag{
		Name:     "apply-type",
		Required: false,
		Usage:    "applyType",
	},
	&cli.StringFlag{
		Name:     "capability-id",
		Required: false,
		Usage:    "capability",
	},
}
var AppMenuCreateCmd cli.Command = cli.Command{
	Name:    "create",
	Aliases: []string{"c"},
	Flags:   AppMenuCommonCliFlags,
	Usage:   "Create a new template",
	Action: func(c *cli.Context) {
		query := CommonCliQueryDSLBuilder(c)
		entity := CastAppMenuFromCli(c)
		if entity, err := AppMenuActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var AppMenuCreateInteractiveCmd cli.Command = cli.Command{
	Name:  "ic",
	Usage: "Creates a new template, using requied fields in an interactive name",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "all",
			Usage: "Interactively asks for all inputs, not only required ones",
		},
	},
	Action: func(c *cli.Context) {
		query := CommonCliQueryDSLBuilder(c)
		entity := &AppMenuEntity{}
		for _, item := range AppMenuCommonInteractiveCliFlags {
			if !item.Required && c.Bool("all") == false {
				continue
			}
			result := AskForInput(item.Name, "")
			SetFieldString(entity, item.StructField, result)
		}
		if entity, err := AppMenuActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var AppMenuUpdateCmd cli.Command = cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Flags:   AppMenuCommonCliFlagsOptional,
	Usage:   "Updates a template by passing the parameters",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilder(c)
		entity := CastAppMenuFromCli(c)
		if entity, err := AppMenuActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
		return nil
	},
}

func CastAppMenuFromCli(c *cli.Context) *AppMenuEntity {
	template := &AppMenuEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
	if c.IsSet("href") {
		value := c.String("href")
		template.Href = &value
	}
	if c.IsSet("icon") {
		value := c.String("icon")
		template.Icon = &value
	}
	if c.IsSet("label") {
		value := c.String("label")
		template.Label = &value
	}
	if c.IsSet("active-matcher") {
		value := c.String("active-matcher")
		template.ActiveMatcher = &value
	}
	if c.IsSet("apply-type") {
		value := c.String("apply-type")
		template.ApplyType = &value
	}
	if c.IsSet("capability-id") {
		value := c.String("capability-id")
		template.CapabilityId = &value
	}
	return template
}
func AppMenuSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	SeederFromFSImport(
		QueryDSL{},
		AppMenuActionCreate,
		reflect.ValueOf(&AppMenuEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}
func AppMenuSyncSeeders() {
	SeederFromFSImport(
		QueryDSL{WorkspaceId: USER_SYSTEM},
		AppMenuActionCreate,
		reflect.ValueOf(&AppMenuEntity{}).Elem(),
		&seeders.ViewsFs,
		[]string{},
		true,
	)
}
func AppMenuWriteQueryMock(ctx MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := AppMenuActionQuery(f)
		result := QueryEntitySuccessResult(f, items, count)
		WriteMockDataToFile(lang, "", "AppMenu", result)
	}
}

var AppMenuImportExportCommands = []cli.Command{
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
			query := CommonCliQueryDSLBuilder(c)
			AppMenuActionSeeder(query, c.Int("count"))
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
				Value: "app-menu-seeder.yml",
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
			f := CommonCliQueryDSLBuilder(c)
			AppMenuActionSeederInit(f, c.String("file"), c.String("format"))
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
				Value: "app-menu-seeder-app-menu.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of app-menus, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]AppMenuEntity{}
			ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
	cli.Command{
		Name:  "list",
		Usage: "Prints the list of files attached to this module for syncing or bootstrapping project",
		Action: func(c *cli.Context) error {
			if entity, err := GetSeederFilenames(&seeders.ViewsFs, ""); err != nil {
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
			CommonCliImportEmbedCmd(c,
				AppMenuActionCreate,
				reflect.ValueOf(&AppMenuEntity{}).Elem(),
				&seeders.ViewsFs,
			)
			return nil
		},
	},
	cli.Command{
		Name:    "export",
		Aliases: []string{"e"},
		Flags: append(CommonQueryFlags,
			&cli.StringFlag{
				Name:     "file",
				Usage:    "The address of file you want the csv/yaml/json be exported to",
				Required: true,
			}),
		Usage: "Exports a query results into the csv/yaml/json format",
		Action: func(c *cli.Context) error {
			CommonCliExportCmd(c,
				AppMenuActionQuery,
				reflect.ValueOf(&AppMenuEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"AppMenuFieldMap.yml",
				AppMenuPreloadRelations,
			)
			return nil
		},
	},
	cli.Command{
		Name: "import",
		Flags: append(CommonQueryFlags,
			&cli.StringFlag{
				Name:     "file",
				Usage:    "The address of file you want the csv be imported from",
				Required: true,
			}),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmd(c,
				AppMenuActionCreate,
				reflect.ValueOf(&AppMenuEntity{}).Elem(),
				c.String("file"),
			)
			return nil
		},
	},
}
var AppMenuCliCommands []cli.Command = []cli.Command{
	GetCommonQuery(AppMenuActionQuery),
	GetCommonTableQuery(reflect.ValueOf(&AppMenuEntity{}).Elem(), AppMenuActionQuery),
	AppMenuCreateCmd,
	AppMenuUpdateCmd,
	AppMenuCreateInteractiveCmd,
	AppMenuWipeCmd,
	GetCommonRemoveQuery(reflect.ValueOf(&AppMenuEntity{}).Elem(), AppMenuActionRemove),
	GetCommonCteQuery(AppMenuActionCteQuery),
	GetCommonPivotQuery(AppMenuActionCommonPivotQuery),
}

func AppMenuCliFn() cli.Command {
	AppMenuCliCommands = append(AppMenuCliCommands, AppMenuImportExportCommands...)
	return cli.Command{
		Name:        "appMenu",
		Description: "AppMenus module actions (sample module to handle complex entities)",
		Usage:       "Manages the menus in the app, (for example tab views, sidebar items, etc.)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: AppMenuCliCommands,
	}
}

/**
 *	Override this function on AppMenuEntityHttp.go,
 *	In order to add your own http
 **/
var AppendAppMenuRouter = func(r *[]Module2Action) {}

func GetAppMenuModule2Actions() []Module2Action {
	routes := []Module2Action{
		{
			Method: "GET",
			Url:    "/cte-app-menus",
			SecurityModel: SecurityModel{
				ActionRequires: []string{PERM_ROOT_APPMENU_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpQueryEntity(c, AppMenuActionCteQuery)
				},
			},
			Format:         "QUERY",
			Action:         AppMenuActionCteQuery,
			ResponseEntity: &[]AppMenuEntity{},
		},
		{
			Method: "GET",
			Url:    "/app-menus",
			SecurityModel: SecurityModel{
				ActionRequires: []string{PERM_ROOT_APPMENU_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpQueryEntity(c, AppMenuActionQuery)
				},
			},
			Format:         "QUERY",
			Action:         AppMenuActionQuery,
			ResponseEntity: &[]AppMenuEntity{},
		},
		{
			Method: "GET",
			Url:    "/app-menus/export",
			SecurityModel: SecurityModel{
				ActionRequires: []string{PERM_ROOT_APPMENU_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpStreamFileChannel(c, AppMenuActionExport)
				},
			},
			Format:         "QUERY",
			Action:         AppMenuActionExport,
			ResponseEntity: &[]AppMenuEntity{},
		},
		{
			Method: "GET",
			Url:    "/app-menu/:uniqueId",
			SecurityModel: SecurityModel{
				ActionRequires: []string{PERM_ROOT_APPMENU_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpGetEntity(c, AppMenuActionGetOne)
				},
			},
			Format:         "GET_ONE",
			Action:         AppMenuActionGetOne,
			ResponseEntity: &AppMenuEntity{},
		},
		{
			Method: "POST",
			Url:    "/app-menu",
			SecurityModel: SecurityModel{
				ActionRequires: []string{PERM_ROOT_APPMENU_CREATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpPostEntity(c, AppMenuActionCreate)
				},
			},
			Action:         AppMenuActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &AppMenuEntity{},
			ResponseEntity: &AppMenuEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/app-menu",
			SecurityModel: SecurityModel{
				ActionRequires: []string{PERM_ROOT_APPMENU_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpUpdateEntity(c, AppMenuActionUpdate)
				},
			},
			Action:         AppMenuActionUpdate,
			RequestEntity:  &AppMenuEntity{},
			Format:         "PATCH_ONE",
			ResponseEntity: &AppMenuEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/app-menus",
			SecurityModel: SecurityModel{
				ActionRequires: []string{PERM_ROOT_APPMENU_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpUpdateEntities(c, AppMenuActionBulkUpdate)
				},
			},
			Action:         AppMenuActionBulkUpdate,
			Format:         "PATCH_BULK",
			RequestEntity:  &BulkRecordRequest[AppMenuEntity]{},
			ResponseEntity: &BulkRecordRequest[AppMenuEntity]{},
		},
		{
			Method: "DELETE",
			Url:    "/app-menu",
			Format: "DELETE_DSL",
			SecurityModel: SecurityModel{
				ActionRequires: []string{PERM_ROOT_APPMENU_DELETE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpRemoveEntity(c, AppMenuActionRemove)
				},
			},
			Action:         AppMenuActionRemove,
			RequestEntity:  &DeleteRequest{},
			ResponseEntity: &DeleteResponse{},
			TargetEntity:   &AppMenuEntity{},
		},
	}
	// Append user defined functions
	AppendAppMenuRouter(&routes)
	return routes
}
func CreateAppMenuRouter(r *gin.Engine) []Module2Action {
	httpRoutes := GetAppMenuModule2Actions()
	CastRoutes(httpRoutes, r)
	WriteHttpInformationToFile(&httpRoutes, AppMenuEntityJsonSchema, "app-menu-http", "workspaces")
	WriteEntitySchema("AppMenuEntity", AppMenuEntityJsonSchema, "workspaces")
	return httpRoutes
}

var PERM_ROOT_APPMENU_DELETE = "root/appmenu/delete"
var PERM_ROOT_APPMENU_CREATE = "root/appmenu/create"
var PERM_ROOT_APPMENU_UPDATE = "root/appmenu/update"
var PERM_ROOT_APPMENU_QUERY = "root/appmenu/query"
var PERM_ROOT_APPMENU = "root/appmenu"
var ALL_APPMENU_PERMISSIONS = []string{
	PERM_ROOT_APPMENU_DELETE,
	PERM_ROOT_APPMENU_CREATE,
	PERM_ROOT_APPMENU_UPDATE,
	PERM_ROOT_APPMENU_QUERY,
	PERM_ROOT_APPMENU,
}
