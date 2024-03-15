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
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WorkspaceTypeEntity struct {
	Visibility       *string `json:"visibility,omitempty" yaml:"visibility"`
	WorkspaceId      *string `json:"workspaceId,omitempty" yaml:"workspaceId" gorm:"unique;not null;" `
	LinkerId         *string `json:"linkerId,omitempty" yaml:"linkerId"`
	ParentId         *string `json:"parentId,omitempty" yaml:"parentId"`
	UniqueId         string  `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
	UserId           *string `json:"userId,omitempty" yaml:"userId"`
	Rank             int64   `json:"rank,omitempty" gorm:"type:int;name:rank"`
	Updated          int64   `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
	Created          int64   `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
	CreatedFormatted string  `json:"createdFormatted,omitempty" sql:"-" gorm:"-"`
	UpdatedFormatted string  `json:"updatedFormatted,omitempty" sql:"-" gorm:"-"`
	Title            *string `json:"title" yaml:"title"  validate:"required,omitempty,min=1,max=250"        translate:"true" `
	// Datenano also has a text representation
	Description *string `json:"description" yaml:"description"        translate:"true" `
	// Datenano also has a text representation
	Slug *string `json:"slug" yaml:"slug"  validate:"required,omitempty,min=2,max=50"    gorm:"unique;not null;size:100;"     `
	// Datenano also has a text representation
	Role *RoleEntity `json:"role" yaml:"role"    gorm:"foreignKey:RoleId;references:UniqueId"     `
	// Datenano also has a text representation
	RoleId       *string                        `json:"roleId" yaml:"roleId" validate:"required" `
	Translations []*WorkspaceTypeEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
	Children     []*WorkspaceTypeEntity         `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
	LinkedTo     *WorkspaceTypeEntity           `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

var WorkspaceTypePreloadRelations []string = []string{}
var WORKSPACETYPE_EVENT_CREATED = "workspaceType.created"
var WORKSPACETYPE_EVENT_UPDATED = "workspaceType.updated"
var WORKSPACETYPE_EVENT_DELETED = "workspaceType.deleted"
var WORKSPACETYPE_EVENTS = []string{
	WORKSPACETYPE_EVENT_CREATED,
	WORKSPACETYPE_EVENT_UPDATED,
	WORKSPACETYPE_EVENT_DELETED,
}

type WorkspaceTypeFieldMap struct {
	Title       TranslatedString `yaml:"title"`
	Description TranslatedString `yaml:"description"`
	Slug        TranslatedString `yaml:"slug"`
	Role        TranslatedString `yaml:"role"`
}

var WorkspaceTypeEntityMetaConfig map[string]int64 = map[string]int64{}
var WorkspaceTypeEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&WorkspaceTypeEntity{}))

type WorkspaceTypeEntityPolyglot struct {
	LinkerId    string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
	LanguageId  string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
	Title       string `yaml:"title" json:"title"`
	Description string `yaml:"description" json:"description"`
}

func entityWorkspaceTypeFormatter(dto *WorkspaceTypeEntity, query QueryDSL) {
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
func WorkspaceTypeMockEntity() *WorkspaceTypeEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &WorkspaceTypeEntity{
		Title:       &stringHolder,
		Description: &stringHolder,
		Slug:        &stringHolder,
	}
	return entity
}
func WorkspaceTypeActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := WorkspaceTypeMockEntity()
		_, err := WorkspaceTypeActionCreate(entity, query)
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
func (x *WorkspaceTypeEntity) GetTitleTranslated(language string) string {
	if x.Translations != nil && len(x.Translations) > 0 {
		for _, item := range x.Translations {
			if item.LanguageId == language {
				return item.Title
			}
		}
	}
	return ""
}
func (x *WorkspaceTypeEntity) GetDescriptionTranslated(language string) string {
	if x.Translations != nil && len(x.Translations) > 0 {
		for _, item := range x.Translations {
			if item.LanguageId == language {
				return item.Description
			}
		}
	}
	return ""
}
func WorkspaceTypeActionSeederInit(query QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*WorkspaceTypeEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &WorkspaceTypeEntity{
		Title:       &tildaRef,
		Description: &tildaRef,
		Slug:        &tildaRef,
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
func WorkspaceTypeAssociationCreate(dto *WorkspaceTypeEntity, query QueryDSL) error {
	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func WorkspaceTypeRelationContentCreate(dto *WorkspaceTypeEntity, query QueryDSL) error {
	return nil
}
func WorkspaceTypeRelationContentUpdate(dto *WorkspaceTypeEntity, query QueryDSL) error {
	return nil
}
func WorkspaceTypePolyglotCreateHandler(dto *WorkspaceTypeEntity, query QueryDSL) {
	if dto == nil {
		return
	}
	PolyglotCreateHandler(dto, &WorkspaceTypeEntityPolyglot{}, query)
}

/**
 * This will be validating your entity fully. Important note is that, you add validate:* tag
 * in your entity, it will automatically work here. For slices inside entity, make sure you add
 * extra line of AppendSliceErrors, otherwise they won't be detected
 */
func WorkspaceTypeValidator(dto *WorkspaceTypeEntity, isPatch bool) *IError {
	err := CommonStructValidatorPointer(dto, isPatch)
	return err
}
func WorkspaceTypeEntityPreSanitize(dto *WorkspaceTypeEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
func WorkspaceTypeEntityBeforeCreateAppend(dto *WorkspaceTypeEntity, query QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = UUID()
	}
	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId
	WorkspaceTypeRecursiveAddUniqueId(dto, query)
}
func WorkspaceTypeRecursiveAddUniqueId(dto *WorkspaceTypeEntity, query QueryDSL) {
}
func WorkspaceTypeActionBatchCreateFn(dtos []*WorkspaceTypeEntity, query QueryDSL) ([]*WorkspaceTypeEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*WorkspaceTypeEntity{}
		for _, item := range dtos {
			s, err := WorkspaceTypeActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil
}
func WorkspaceTypeActionCreateFn(dto *WorkspaceTypeEntity, query QueryDSL) (*WorkspaceTypeEntity, *IError) {
	// 1. Validate always
	if iError := WorkspaceTypeValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	WorkspaceTypeEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	WorkspaceTypeEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	WorkspaceTypePolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	WorkspaceTypeRelationContentCreate(dto, query)
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
	WorkspaceTypeAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(WORKSPACETYPE_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": GetTypeString(&WorkspaceTypeEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})
	return dto, nil
}
func WorkspaceTypeActionGetOne(query QueryDSL) (*WorkspaceTypeEntity, *IError) {
	refl := reflect.ValueOf(&WorkspaceTypeEntity{})
	item, err := GetOneEntity[WorkspaceTypeEntity](query, refl)
	entityWorkspaceTypeFormatter(item, query)
	return item, err
}
func WorkspaceTypeActionQuery(query QueryDSL) ([]*WorkspaceTypeEntity, *QueryResultMeta, error) {
	refl := reflect.ValueOf(&WorkspaceTypeEntity{})
	items, meta, err := QueryEntitiesPointer[WorkspaceTypeEntity](query, refl)
	for _, item := range items {
		entityWorkspaceTypeFormatter(item, query)
	}
	return items, meta, err
}
func WorkspaceTypeUpdateExec(dbref *gorm.DB, query QueryDSL, fields *WorkspaceTypeEntity) (*WorkspaceTypeEntity, *IError) {
	uniqueId := fields.UniqueId
	query.TriggerEventName = WORKSPACETYPE_EVENT_UPDATED
	WorkspaceTypeEntityPreSanitize(fields, query)
	var item WorkspaceTypeEntity
	q := dbref.
		Where(&WorkspaceTypeEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)
	err := q.UpdateColumns(fields).Error
	if err != nil {
		return nil, GormErrorToIError(err)
	}
	query.Tx = dbref
	WorkspaceTypeRelationContentUpdate(fields, query)
	WorkspaceTypePolyglotCreateHandler(fields, query)
	// @meta(update has many)
	err = dbref.
		Preload(clause.Associations).
		Where(&WorkspaceTypeEntity{UniqueId: uniqueId}).
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
func WorkspaceTypeActionUpdateFn(query QueryDSL, fields *WorkspaceTypeEntity) (*WorkspaceTypeEntity, *IError) {
	if fields == nil {
		return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}
	// 1. Validate always
	if iError := WorkspaceTypeValidator(fields, true); iError != nil {
		return nil, iError
	}
	WorkspaceTypeRecursiveAddUniqueId(fields, query)
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = GetDbRef()
		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			_, err := WorkspaceTypeUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}
		})
		return nil, CastToIError(vf)
	} else {
		dbref = query.Tx
		return WorkspaceTypeUpdateExec(dbref, query, fields)
	}
}

var WorkspaceTypeWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire workspacetypes ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilder(c)
		count, _ := WorkspaceTypeActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}

func WorkspaceTypeActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&WorkspaceTypeEntity{})
	query.ActionRequires = []string{PERM_ROOT_WORKSPACETYPE_DELETE}
	return RemoveEntity[WorkspaceTypeEntity](query, refl)
}
func WorkspaceTypeActionWipeClean(query QueryDSL) (int64, error) {
	var err error
	var count int64 = 0
	{
		subCount, subErr := WipeCleanEntity[WorkspaceTypeEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'WorkspaceTypeEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
func WorkspaceTypeActionBulkUpdate(
	query QueryDSL, dto *BulkRecordRequest[WorkspaceTypeEntity]) (
	*BulkRecordRequest[WorkspaceTypeEntity], *IError,
) {
	result := []*WorkspaceTypeEntity{}
	err := GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := WorkspaceTypeActionUpdate(query, record)
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
func (x *WorkspaceTypeEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

var WorkspaceTypeEntityMeta = TableMetaData{
	EntityName:    "WorkspaceType",
	ExportKey:     "workspace-types",
	TableNameInDb: "fb_workspacetype_entities",
	EntityObject:  &WorkspaceTypeEntity{},
	ExportStream:  WorkspaceTypeActionExportT,
	ImportQuery:   WorkspaceTypeActionImport,
}

func WorkspaceTypeActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[WorkspaceTypeEntity](query, WorkspaceTypeActionQuery, WorkspaceTypePreloadRelations)
}
func WorkspaceTypeActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[WorkspaceTypeEntity](query, WorkspaceTypeActionQuery, WorkspaceTypePreloadRelations)
}
func WorkspaceTypeActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content WorkspaceTypeEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := WorkspaceTypeActionCreate(&content, query)
	return err
}

var WorkspaceTypeCommonCliFlags = []cli.Flag{
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
		Name:     "title",
		Required: true,
		Usage:    "title",
	},
	&cli.StringFlag{
		Name:     "description",
		Required: false,
		Usage:    "description",
	},
	&cli.StringFlag{
		Name:     "slug",
		Required: true,
		Usage:    "slug",
	},
	&cli.StringFlag{
		Name:     "role-id",
		Required: true,
		Usage:    "role",
	},
}
var WorkspaceTypeCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:        "title",
		StructField: "Title",
		Required:    true,
		Usage:       "title",
		Type:        "string",
	},
	{
		Name:        "description",
		StructField: "Description",
		Required:    false,
		Usage:       "description",
		Type:        "string",
	},
	{
		Name:        "slug",
		StructField: "Slug",
		Required:    true,
		Usage:       "slug",
		Type:        "string",
	},
}
var WorkspaceTypeCommonCliFlagsOptional = []cli.Flag{
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
		Name:     "title",
		Required: true,
		Usage:    "title",
	},
	&cli.StringFlag{
		Name:     "description",
		Required: false,
		Usage:    "description",
	},
	&cli.StringFlag{
		Name:     "slug",
		Required: true,
		Usage:    "slug",
	},
	&cli.StringFlag{
		Name:     "role-id",
		Required: true,
		Usage:    "role",
	},
}
var WorkspaceTypeCreateCmd cli.Command = cli.Command{
	Name:    "create",
	Aliases: []string{"c"},
	Flags:   WorkspaceTypeCommonCliFlags,
	Usage:   "Create a new template",
	Action: func(c *cli.Context) {
		query := CommonCliQueryDSLBuilder(c)
		entity := CastWorkspaceTypeFromCli(c)
		if entity, err := WorkspaceTypeActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var WorkspaceTypeCreateInteractiveCmd cli.Command = cli.Command{
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
		entity := &WorkspaceTypeEntity{}
		for _, item := range WorkspaceTypeCommonInteractiveCliFlags {
			if !item.Required && c.Bool("all") == false {
				continue
			}
			result := AskForInput(item.Name, "")
			SetFieldString(entity, item.StructField, result)
		}
		if entity, err := WorkspaceTypeActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var WorkspaceTypeUpdateCmd cli.Command = cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Flags:   WorkspaceTypeCommonCliFlagsOptional,
	Usage:   "Updates a template by passing the parameters",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilder(c)
		entity := CastWorkspaceTypeFromCli(c)
		if entity, err := WorkspaceTypeActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
		return nil
	},
}

func (x WorkspaceTypeEntity) FromCli(c *cli.Context) *WorkspaceTypeEntity {
	return CastWorkspaceTypeFromCli(c)
}
func CastWorkspaceTypeFromCli(c *cli.Context) *WorkspaceTypeEntity {
	template := &WorkspaceTypeEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
	if c.IsSet("title") {
		value := c.String("title")
		template.Title = &value
	}
	if c.IsSet("description") {
		value := c.String("description")
		template.Description = &value
	}
	if c.IsSet("slug") {
		value := c.String("slug")
		template.Slug = &value
	}
	if c.IsSet("role-id") {
		value := c.String("role-id")
		template.RoleId = &value
	}
	return template
}
func WorkspaceTypeSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	SeederFromFSImport(
		QueryDSL{},
		WorkspaceTypeActionCreate,
		reflect.ValueOf(&WorkspaceTypeEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}
func WorkspaceTypeWriteQueryMock(ctx MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := WorkspaceTypeActionQuery(f)
		result := QueryEntitySuccessResult(f, items, count)
		WriteMockDataToFile(lang, "", "WorkspaceType", result)
	}
}

var WorkspaceTypeImportExportCommands = []cli.Command{
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
			WorkspaceTypeActionSeeder(query, c.Int("count"))
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
				Value: "workspace-type-seeder.yml",
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
			WorkspaceTypeActionSeederInit(f, c.String("file"), c.String("format"))
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
				Value: "workspace-type-seeder-workspace-type.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of workspace-types, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]WorkspaceTypeEntity{}
			ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
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
				WorkspaceTypeActionCreate,
				reflect.ValueOf(&WorkspaceTypeEntity{}).Elem(),
				c.String("file"),
			)
			return nil
		},
	},
}
var WorkspaceTypeCliCommands []cli.Command = []cli.Command{
	GetCommonQuery(WorkspaceTypeActionQuery),
	GetCommonTableQuery(reflect.ValueOf(&WorkspaceTypeEntity{}).Elem(), WorkspaceTypeActionQuery),
	WorkspaceTypeCreateCmd,
	WorkspaceTypeUpdateCmd,
	WorkspaceTypeCreateInteractiveCmd,
	WorkspaceTypeWipeCmd,
	GetCommonRemoveQuery(reflect.ValueOf(&WorkspaceTypeEntity{}).Elem(), WorkspaceTypeActionRemove),
}

func WorkspaceTypeCliFn() cli.Command {
	WorkspaceTypeCliCommands = append(WorkspaceTypeCliCommands, WorkspaceTypeImportExportCommands...)
	return cli.Command{
		Name:        "type",
		Description: "WorkspaceTypes module actions (sample module to handle complex entities)",
		Usage:       "",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: WorkspaceTypeCliCommands,
	}
}

/**
 *	Override this function on WorkspaceTypeEntityHttp.go,
 *	In order to add your own http
 **/
var AppendWorkspaceTypeRouter = func(r *[]Module2Action) {}

func GetWorkspaceTypeModule2Actions() []Module2Action {
	routes := []Module2Action{
		{
			Method: "GET",
			Url:    "/workspace-types",
			SecurityModel: SecurityModel{
				ActionRequires: []string{PERM_ROOT_WORKSPACETYPE_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpQueryEntity(c, WorkspaceTypeActionQuery)
				},
			},
			Format:         "QUERY",
			Action:         WorkspaceTypeActionQuery,
			ResponseEntity: &[]WorkspaceTypeEntity{},
		},
		{
			Method: "GET",
			Url:    "/workspace-types/export",
			SecurityModel: SecurityModel{
				ActionRequires: []string{PERM_ROOT_WORKSPACETYPE_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpStreamFileChannel(c, WorkspaceTypeActionExport)
				},
			},
			Format:         "QUERY",
			Action:         WorkspaceTypeActionExport,
			ResponseEntity: &[]WorkspaceTypeEntity{},
		},
		{
			Method: "GET",
			Url:    "/workspace-type/:uniqueId",
			SecurityModel: SecurityModel{
				ActionRequires: []string{PERM_ROOT_WORKSPACETYPE_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpGetEntity(c, WorkspaceTypeActionGetOne)
				},
			},
			Format:         "GET_ONE",
			Action:         WorkspaceTypeActionGetOne,
			ResponseEntity: &WorkspaceTypeEntity{},
		},
		{
			ActionName:    "create",
			ActionAliases: []string{"c"},
			Flags:         WorkspaceTypeCommonCliFlags,
			Method:        "POST",
			Url:           "/workspace-type",
			SecurityModel: SecurityModel{
				ActionRequires: []string{PERM_ROOT_WORKSPACETYPE_CREATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpPostEntity(c, WorkspaceTypeActionCreate)
				},
			},
			Action:         WorkspaceTypeActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &WorkspaceTypeEntity{},
			ResponseEntity: &WorkspaceTypeEntity{},
		},
		{
			ActionName:    "update",
			ActionAliases: []string{"u"},
			Flags:         WorkspaceTypeCommonCliFlagsOptional,
			Method:        "PATCH",
			Url:           "/workspace-type",
			SecurityModel: SecurityModel{
				ActionRequires: []string{PERM_ROOT_WORKSPACETYPE_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpUpdateEntity(c, WorkspaceTypeActionUpdate)
				},
			},
			Action:         WorkspaceTypeActionUpdate,
			RequestEntity:  &WorkspaceTypeEntity{},
			Format:         "PATCH_ONE",
			ResponseEntity: &WorkspaceTypeEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/workspace-types",
			SecurityModel: SecurityModel{
				ActionRequires: []string{PERM_ROOT_WORKSPACETYPE_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpUpdateEntities(c, WorkspaceTypeActionBulkUpdate)
				},
			},
			Action:         WorkspaceTypeActionBulkUpdate,
			Format:         "PATCH_BULK",
			RequestEntity:  &BulkRecordRequest[WorkspaceTypeEntity]{},
			ResponseEntity: &BulkRecordRequest[WorkspaceTypeEntity]{},
		},
		{
			Method: "DELETE",
			Url:    "/workspace-type",
			Format: "DELETE_DSL",
			SecurityModel: SecurityModel{
				ActionRequires: []string{PERM_ROOT_WORKSPACETYPE_DELETE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpRemoveEntity(c, WorkspaceTypeActionRemove)
				},
			},
			Action:         WorkspaceTypeActionRemove,
			RequestEntity:  &DeleteRequest{},
			ResponseEntity: &DeleteResponse{},
			TargetEntity:   &WorkspaceTypeEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/workspace-type/distinct",
			SecurityModel: SecurityModel{
				ActionRequires: []string{PERM_ROOT_WORKSPACETYPE_UPDATE_DISTINCT_WORKSPACE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpUpdateEntity(c, WorkspaceTypeDistinctActionUpdate)
				},
			},
			Action:         WorkspaceTypeDistinctActionUpdate,
			Format:         "PATCH_ONE",
			RequestEntity:  &WorkspaceTypeEntity{},
			ResponseEntity: &WorkspaceTypeEntity{},
		},
		{
			Method: "GET",
			Url:    "/workspace-type/distinct",
			SecurityModel: SecurityModel{
				ActionRequires: []string{PERM_ROOT_WORKSPACETYPE_GET_DISTINCT_WORKSPACE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpGetEntity(c, WorkspaceTypeDistinctActionGetOne)
				},
			},
			Action:         WorkspaceTypeDistinctActionGetOne,
			Format:         "GET_ONE",
			ResponseEntity: &WorkspaceTypeEntity{},
		},
	}
	// Append user defined functions
	AppendWorkspaceTypeRouter(&routes)
	return routes
}
func CreateWorkspaceTypeRouter(r *gin.Engine) []Module2Action {
	httpRoutes := GetWorkspaceTypeModule2Actions()
	CastRoutes(httpRoutes, r)
	WriteHttpInformationToFile(&httpRoutes, WorkspaceTypeEntityJsonSchema, "workspace-type-http", "workspaces")
	WriteEntitySchema("WorkspaceTypeEntity", WorkspaceTypeEntityJsonSchema, "workspaces")
	return httpRoutes
}

var PERM_ROOT_WORKSPACETYPE_DELETE = "root/workspacetype/delete"
var PERM_ROOT_WORKSPACETYPE_CREATE = "root/workspacetype/create"
var PERM_ROOT_WORKSPACETYPE_UPDATE = "root/workspacetype/update"
var PERM_ROOT_WORKSPACETYPE_QUERY = "root/workspacetype/query"
var PERM_ROOT_WORKSPACETYPE_GET_DISTINCT_WORKSPACE = "root/workspacetype/get-distinct-workspace"
var PERM_ROOT_WORKSPACETYPE_UPDATE_DISTINCT_WORKSPACE = "root/workspacetype/update-distinct-workspace"
var PERM_ROOT_WORKSPACETYPE = "root/workspacetype"
var ALL_WORKSPACETYPE_PERMISSIONS = []string{
	PERM_ROOT_WORKSPACETYPE_DELETE,
	PERM_ROOT_WORKSPACETYPE_CREATE,
	PERM_ROOT_WORKSPACETYPE_UPDATE,
	PERM_ROOT_WORKSPACETYPE_GET_DISTINCT_WORKSPACE,
	PERM_ROOT_WORKSPACETYPE_UPDATE_DISTINCT_WORKSPACE,
	PERM_ROOT_WORKSPACETYPE_QUERY,
	PERM_ROOT_WORKSPACETYPE,
}

func WorkspaceTypeDistinctActionUpdate(
	query QueryDSL,
	fields *WorkspaceTypeEntity,
) (*WorkspaceTypeEntity, *IError) {
	query.UniqueId = query.UserId
	entity, err := WorkspaceTypeActionGetOne(query)
	if err != nil || entity.UniqueId == "" {
		fields.UniqueId = query.UserId
		return WorkspaceTypeActionCreateFn(fields, query)
	} else {
		fields.UniqueId = query.UniqueId
		return WorkspaceTypeActionUpdateFn(query, fields)
	}
}
func WorkspaceTypeDistinctActionGetOne(
	query QueryDSL,
) (*WorkspaceTypeEntity, *IError) {
	query.UniqueId = query.UserId
	entity, err := WorkspaceTypeActionGetOne(query)
	if err != nil && err.HttpCode == 404 {
		return &WorkspaceTypeEntity{}, nil
	}
	return entity, err
}
