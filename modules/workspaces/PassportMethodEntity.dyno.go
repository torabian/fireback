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
	seeders "github.com/torabian/fireback/modules/workspaces/seeders/PassportMethod"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var passportMethodSeedersFs = &seeders.ViewsFs

func ResetPassportMethodSeeders(fs *embed.FS) {
	passportMethodSeedersFs = fs
}

type PassportMethodEntity struct {
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
	Name             *string `json:"name" yaml:"name"  validate:"required"        translate:"true" `
	// Datenano also has a text representation
	Type *string `json:"type" yaml:"type"  validate:"required"       `
	// Datenano also has a text representation
	Region *string `json:"region" yaml:"region"  validate:"required"       `
	// Datenano also has a text representation
	Translations []*PassportMethodEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
	Children     []*PassportMethodEntity         `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
	LinkedTo     *PassportMethodEntity           `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

var PassportMethodPreloadRelations []string = []string{}
var PASSPORT_METHOD_EVENT_CREATED = "passportMethod.created"
var PASSPORT_METHOD_EVENT_UPDATED = "passportMethod.updated"
var PASSPORT_METHOD_EVENT_DELETED = "passportMethod.deleted"
var PASSPORT_METHOD_EVENTS = []string{
	PASSPORT_METHOD_EVENT_CREATED,
	PASSPORT_METHOD_EVENT_UPDATED,
	PASSPORT_METHOD_EVENT_DELETED,
}

type PassportMethodFieldMap struct {
	Name   TranslatedString `yaml:"name"`
	Type   TranslatedString `yaml:"type"`
	Region TranslatedString `yaml:"region"`
}

var PassportMethodEntityMetaConfig map[string]int64 = map[string]int64{}
var PassportMethodEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&PassportMethodEntity{}))

type PassportMethodEntityPolyglot struct {
	LinkerId   string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
	LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
	Name       string `yaml:"name" json:"name"`
}

func entityPassportMethodFormatter(dto *PassportMethodEntity, query QueryDSL) {
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
func PassportMethodMockEntity() *PassportMethodEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &PassportMethodEntity{
		Name:   &stringHolder,
		Type:   &stringHolder,
		Region: &stringHolder,
	}
	return entity
}
func PassportMethodActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := PassportMethodMockEntity()
		_, err := PassportMethodActionCreate(entity, query)
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
func (x *PassportMethodEntity) GetNameTranslated(language string) string {
	if x.Translations != nil && len(x.Translations) > 0 {
		for _, item := range x.Translations {
			if item.LanguageId == language {
				return item.Name
			}
		}
	}
	return ""
}
func PassportMethodActionSeederInit(query QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*PassportMethodEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &PassportMethodEntity{
		Name:   &tildaRef,
		Type:   &tildaRef,
		Region: &tildaRef,
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
func PassportMethodAssociationCreate(dto *PassportMethodEntity, query QueryDSL) error {
	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func PassportMethodRelationContentCreate(dto *PassportMethodEntity, query QueryDSL) error {
	return nil
}
func PassportMethodRelationContentUpdate(dto *PassportMethodEntity, query QueryDSL) error {
	return nil
}
func PassportMethodPolyglotCreateHandler(dto *PassportMethodEntity, query QueryDSL) {
	if dto == nil {
		return
	}
	PolyglotCreateHandler(dto, &PassportMethodEntityPolyglot{}, query)
}

/**
 * This will be validating your entity fully. Important note is that, you add validate:* tag
 * in your entity, it will automatically work here. For slices inside entity, make sure you add
 * extra line of AppendSliceErrors, otherwise they won't be detected
 */
func PassportMethodValidator(dto *PassportMethodEntity, isPatch bool) *IError {
	err := CommonStructValidatorPointer(dto, isPatch)
	return err
}
func PassportMethodEntityPreSanitize(dto *PassportMethodEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
func PassportMethodEntityBeforeCreateAppend(dto *PassportMethodEntity, query QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = UUID()
	}
	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId
	PassportMethodRecursiveAddUniqueId(dto, query)
}
func PassportMethodRecursiveAddUniqueId(dto *PassportMethodEntity, query QueryDSL) {
}
func PassportMethodActionBatchCreateFn(dtos []*PassportMethodEntity, query QueryDSL) ([]*PassportMethodEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*PassportMethodEntity{}
		for _, item := range dtos {
			s, err := PassportMethodActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil
}
func PassportMethodDeleteEntireChildren(query QueryDSL, dto *PassportMethodEntity) *IError {
	// intentionally removed this. It's hard to implement it, and probably wrong without
	// proper on delete cascade
	return nil
}
func PassportMethodActionCreateFn(dto *PassportMethodEntity, query QueryDSL) (*PassportMethodEntity, *IError) {
	// 1. Validate always
	if iError := PassportMethodValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	PassportMethodEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	PassportMethodEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	PassportMethodPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	PassportMethodRelationContentCreate(dto, query)
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
	PassportMethodAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(PASSPORT_METHOD_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": GetTypeString(&PassportMethodEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})
	return dto, nil
}
func PassportMethodActionGetOne(query QueryDSL) (*PassportMethodEntity, *IError) {
	refl := reflect.ValueOf(&PassportMethodEntity{})
	item, err := GetOneEntity[PassportMethodEntity](query, refl)
	entityPassportMethodFormatter(item, query)
	return item, err
}
func PassportMethodActionQuery(query QueryDSL) ([]*PassportMethodEntity, *QueryResultMeta, error) {
	refl := reflect.ValueOf(&PassportMethodEntity{})
	items, meta, err := QueryEntitiesPointer[PassportMethodEntity](query, refl)
	for _, item := range items {
		entityPassportMethodFormatter(item, query)
	}
	return items, meta, err
}
func PassportMethodUpdateExec(dbref *gorm.DB, query QueryDSL, fields *PassportMethodEntity) (*PassportMethodEntity, *IError) {
	uniqueId := fields.UniqueId
	query.TriggerEventName = PASSPORT_METHOD_EVENT_UPDATED
	PassportMethodEntityPreSanitize(fields, query)
	var item PassportMethodEntity
	q := dbref.
		Where(&PassportMethodEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)
	err := q.UpdateColumns(fields).Error
	if err != nil {
		return nil, GormErrorToIError(err)
	}
	query.Tx = dbref
	PassportMethodRelationContentUpdate(fields, query)
	PassportMethodPolyglotCreateHandler(fields, query)
	if ero := PassportMethodDeleteEntireChildren(query, fields); ero != nil {
		return nil, ero
	}
	// @meta(update has many)
	err = dbref.
		Preload(clause.Associations).
		Where(&PassportMethodEntity{UniqueId: uniqueId}).
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
func PassportMethodActionUpdateFn(query QueryDSL, fields *PassportMethodEntity) (*PassportMethodEntity, *IError) {
	if fields == nil {
		return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}
	// 1. Validate always
	if iError := PassportMethodValidator(fields, true); iError != nil {
		return nil, iError
	}
	// Let's not add this. I am not sure of the consequences
	// PassportMethodRecursiveAddUniqueId(fields, query)
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = GetDbRef()
		var item *PassportMethodEntity
		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			var err *IError
			item, err = PassportMethodUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}
		})
		return item, CastToIError(vf)
	} else {
		dbref = query.Tx
		return PassportMethodUpdateExec(dbref, query, fields)
	}
}

var PassportMethodWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire passportmethods ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
			ActionRequires: []PermissionInfo{PERM_ROOT_PASSPORT_METHOD_DELETE},
		})
		count, _ := PassportMethodActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}

func PassportMethodActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&PassportMethodEntity{})
	query.ActionRequires = []PermissionInfo{PERM_ROOT_PASSPORT_METHOD_DELETE}
	return RemoveEntity[PassportMethodEntity](query, refl)
}
func PassportMethodActionWipeClean(query QueryDSL) (int64, error) {
	var err error
	var count int64 = 0
	{
		subCount, subErr := WipeCleanEntity[PassportMethodEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'PassportMethodEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
func PassportMethodActionBulkUpdate(
	query QueryDSL, dto *BulkRecordRequest[PassportMethodEntity]) (
	*BulkRecordRequest[PassportMethodEntity], *IError,
) {
	result := []*PassportMethodEntity{}
	err := GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := PassportMethodActionUpdate(query, record)
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
func (x *PassportMethodEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

var PassportMethodEntityMeta = TableMetaData{
	EntityName:    "PassportMethod",
	ExportKey:     "passport-methods",
	TableNameInDb: "fb_passport-method_entities",
	EntityObject:  &PassportMethodEntity{},
	ExportStream:  PassportMethodActionExportT,
	ImportQuery:   PassportMethodActionImport,
}

func PassportMethodActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[PassportMethodEntity](query, PassportMethodActionQuery, PassportMethodPreloadRelations)
}
func PassportMethodActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[PassportMethodEntity](query, PassportMethodActionQuery, PassportMethodPreloadRelations)
}
func PassportMethodActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content PassportMethodEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := PassportMethodActionCreate(&content, query)
	return err
}

var PassportMethodCommonCliFlags = []cli.Flag{
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
		Required: true,
		Usage:    "name",
	},
	&cli.StringFlag{
		Name:     "type",
		Required: true,
		Usage:    "type",
	},
	&cli.StringFlag{
		Name:     "region",
		Required: true,
		Usage:    "region",
	},
}
var PassportMethodCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:        "name",
		StructField: "Name",
		Required:    true,
		Usage:       "name",
		Type:        "string",
	},
	{
		Name:        "type",
		StructField: "Type",
		Required:    true,
		Usage:       "type",
		Type:        "string",
	},
	{
		Name:        "region",
		StructField: "Region",
		Required:    true,
		Usage:       "region",
		Type:        "string",
	},
}
var PassportMethodCommonCliFlagsOptional = []cli.Flag{
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
		Required: true,
		Usage:    "name",
	},
	&cli.StringFlag{
		Name:     "type",
		Required: true,
		Usage:    "type",
	},
	&cli.StringFlag{
		Name:     "region",
		Required: true,
		Usage:    "region",
	},
}
var PassportMethodCreateCmd cli.Command = PASSPORT_METHOD_ACTION_POST_ONE.ToCli()
var PassportMethodCreateInteractiveCmd cli.Command = cli.Command{
	Name:  "ic",
	Usage: "Creates a new template, using requied fields in an interactive name",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "all",
			Usage: "Interactively asks for all inputs, not only required ones",
		},
	},
	Action: func(c *cli.Context) {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
			ActionRequires: []PermissionInfo{PERM_ROOT_PASSPORT_METHOD_CREATE},
		})
		entity := &PassportMethodEntity{}
		for _, item := range PassportMethodCommonInteractiveCliFlags {
			if !item.Required && c.Bool("all") == false {
				continue
			}
			result := AskForInput(item.Name, "")
			SetFieldString(entity, item.StructField, result)
		}
		if entity, err := PassportMethodActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var PassportMethodUpdateCmd cli.Command = cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Flags:   PassportMethodCommonCliFlagsOptional,
	Usage:   "Updates a template by passing the parameters",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
			ActionRequires: []PermissionInfo{PERM_ROOT_PASSPORT_METHOD_UPDATE},
		})
		entity := CastPassportMethodFromCli(c)
		if entity, err := PassportMethodActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
		return nil
	},
}

func (x *PassportMethodEntity) FromCli(c *cli.Context) *PassportMethodEntity {
	return CastPassportMethodFromCli(c)
}
func CastPassportMethodFromCli(c *cli.Context) *PassportMethodEntity {
	template := &PassportMethodEntity{}
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
	if c.IsSet("type") {
		value := c.String("type")
		template.Type = &value
	}
	if c.IsSet("region") {
		value := c.String("region")
		template.Region = &value
	}
	return template
}
func PassportMethodSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	SeederFromFSImport(
		QueryDSL{},
		PassportMethodActionCreate,
		reflect.ValueOf(&PassportMethodEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}
func PassportMethodSyncSeeders() {
	SeederFromFSImport(
		QueryDSL{WorkspaceId: USER_SYSTEM},
		PassportMethodActionCreate,
		reflect.ValueOf(&PassportMethodEntity{}).Elem(),
		passportMethodSeedersFs,
		[]string{},
		true,
	)
}
func PassportMethodWriteQueryMock(ctx MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := PassportMethodActionQuery(f)
		result := QueryEntitySuccessResult(f, items, count)
		WriteMockDataToFile(lang, "", "PassportMethod", result)
	}
}

var PassportMethodImportExportCommands = []cli.Command{
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
			query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
				ActionRequires: []PermissionInfo{PERM_ROOT_PASSPORT_METHOD_CREATE},
			})
			PassportMethodActionSeeder(query, c.Int("count"))
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
				Value: "passport-method-seeder.yml",
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
			query := CommonCliQueryDSLBuilder(c)
			PassportMethodActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "passport-method-seeder-passport-method.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of passport-methods, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]PassportMethodEntity{}
			ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
	cli.Command{
		Name:  "list",
		Usage: "Prints the list of files attached to this module for syncing or bootstrapping project",
		Action: func(c *cli.Context) error {
			if entity, err := GetSeederFilenames(passportMethodSeedersFs, ""); err != nil {
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
				PassportMethodActionCreate,
				reflect.ValueOf(&PassportMethodEntity{}).Elem(),
				passportMethodSeedersFs,
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
				PassportMethodActionQuery,
				reflect.ValueOf(&PassportMethodEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"PassportMethodFieldMap.yml",
				PassportMethodPreloadRelations,
			)
			return nil
		},
	},
	cli.Command{
		Name: "import",
		Flags: append(
			append(
				CommonQueryFlags,
				&cli.StringFlag{
					Name:     "file",
					Usage:    "The address of file you want the csv be imported from",
					Required: true,
				}),
			PassportMethodCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				PassportMethodActionCreate,
				reflect.ValueOf(&PassportMethodEntity{}).Elem(),
				c.String("file"),
				&SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_PASSPORT_METHOD_CREATE},
				},
				func() PassportMethodEntity {
					v := CastPassportMethodFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
var PassportMethodCliCommands []cli.Command = []cli.Command{
	PASSPORT_METHOD_ACTION_QUERY.ToCli(),
	PASSPORT_METHOD_ACTION_TABLE.ToCli(),
	PassportMethodCreateCmd,
	PassportMethodUpdateCmd,
	PassportMethodCreateInteractiveCmd,
	PassportMethodWipeCmd,
	GetCommonRemoveQuery(reflect.ValueOf(&PassportMethodEntity{}).Elem(), PassportMethodActionRemove),
}

func PassportMethodCliFn() cli.Command {
	PassportMethodCliCommands = append(PassportMethodCliCommands, PassportMethodImportExportCommands...)
	return cli.Command{
		Name:        "passportMethod",
		ShortName:   "method",
		Description: "PassportMethods module actions (sample module to handle complex entities)",
		Usage:       "Login/Signup methods which are available in the app for different regions (Email, Phone Number, Google, etc)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: PassportMethodCliCommands,
	}
}

var PASSPORT_METHOD_ACTION_TABLE = Module2Action{
	Name:          "table",
	ActionName:    "table",
	ActionAliases: []string{"t"},
	Flags:         CommonQueryFlags,
	Description:   "Table formatted queries all of the entities in database based on the standard query format",
	Action:        PassportMethodActionQuery,
	CliAction: func(c *cli.Context, security *SecurityModel) error {
		CommonCliTableCmd2(c,
			PassportMethodActionQuery,
			security,
			reflect.ValueOf(&PassportMethodEntity{}).Elem(),
		)
		return nil
	},
}
var PASSPORT_METHOD_ACTION_QUERY = Module2Action{
	Method:        "GET",
	Url:           "/passport-methods",
	SecurityModel: &SecurityModel{},
	Group:         "passportMethod",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpQueryEntity(c, PassportMethodActionQuery)
		},
	},
	Format:         "QUERY",
	Action:         PassportMethodActionQuery,
	ResponseEntity: &[]PassportMethodEntity{},
	Out: Module2ActionBody{
		Entity: "PassportMethodEntity",
	},
	CliAction: func(c *cli.Context, security *SecurityModel) error {
		CommonCliQueryCmd2(
			c,
			PassportMethodActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionName:    "query",
	ActionAliases: []string{"q"},
	Flags:         CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var PASSPORT_METHOD_ACTION_EXPORT = Module2Action{
	Method:        "GET",
	Url:           "/passport-methods/export",
	SecurityModel: &SecurityModel{},
	Group:         "passportMethod",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpStreamFileChannel(c, PassportMethodActionExport)
		},
	},
	Format:         "QUERY",
	Action:         PassportMethodActionExport,
	ResponseEntity: &[]PassportMethodEntity{},
	Out: Module2ActionBody{
		Entity: "PassportMethodEntity",
	},
}
var PASSPORT_METHOD_ACTION_GET_ONE = Module2Action{
	Method:        "GET",
	Url:           "/passport-method/:uniqueId",
	SecurityModel: &SecurityModel{},
	Group:         "passportMethod",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpGetEntity(c, PassportMethodActionGetOne)
		},
	},
	Format:         "GET_ONE",
	Action:         PassportMethodActionGetOne,
	ResponseEntity: &PassportMethodEntity{},
	Out: Module2ActionBody{
		Entity: "PassportMethodEntity",
	},
}
var PASSPORT_METHOD_ACTION_POST_ONE = Module2Action{
	ActionName:    "create",
	ActionAliases: []string{"c"},
	Description:   "Create new passportMethod",
	Flags:         PassportMethodCommonCliFlags,
	Method:        "POST",
	Url:           "/passport-method",
	SecurityModel: &SecurityModel{},
	Group:         "passportMethod",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpPostEntity(c, PassportMethodActionCreate)
		},
	},
	CliAction: func(c *cli.Context, security *SecurityModel) error {
		result, err := CliPostEntity(c, PassportMethodActionCreate, security)
		HandleActionInCli(c, result, err, map[string]map[string]string{})
		return err
	},
	Action:         PassportMethodActionCreate,
	Format:         "POST_ONE",
	RequestEntity:  &PassportMethodEntity{},
	ResponseEntity: &PassportMethodEntity{},
	Out: Module2ActionBody{
		Entity: "PassportMethodEntity",
	},
	In: Module2ActionBody{
		Entity: "PassportMethodEntity",
	},
}
var PASSPORT_METHOD_ACTION_PATCH = Module2Action{
	ActionName:    "update",
	ActionAliases: []string{"u"},
	Flags:         PassportMethodCommonCliFlagsOptional,
	Method:        "PATCH",
	Url:           "/passport-method",
	SecurityModel: &SecurityModel{},
	Group:         "passportMethod",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpUpdateEntity(c, PassportMethodActionUpdate)
		},
	},
	Action:         PassportMethodActionUpdate,
	RequestEntity:  &PassportMethodEntity{},
	ResponseEntity: &PassportMethodEntity{},
	Format:         "PATCH_ONE",
	Out: Module2ActionBody{
		Entity: "PassportMethodEntity",
	},
	In: Module2ActionBody{
		Entity: "PassportMethodEntity",
	},
}
var PASSPORT_METHOD_ACTION_PATCH_BULK = Module2Action{
	Method:        "PATCH",
	Url:           "/passport-methods",
	SecurityModel: &SecurityModel{},
	Group:         "passportMethod",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpUpdateEntities(c, PassportMethodActionBulkUpdate)
		},
	},
	Action:         PassportMethodActionBulkUpdate,
	Format:         "PATCH_BULK",
	RequestEntity:  &BulkRecordRequest[PassportMethodEntity]{},
	ResponseEntity: &BulkRecordRequest[PassportMethodEntity]{},
	Out: Module2ActionBody{
		Entity: "PassportMethodEntity",
	},
	In: Module2ActionBody{
		Entity: "PassportMethodEntity",
	},
}
var PASSPORT_METHOD_ACTION_DELETE = Module2Action{
	Method:        "DELETE",
	Url:           "/passport-method",
	Format:        "DELETE_DSL",
	SecurityModel: &SecurityModel{},
	Group:         "passportMethod",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpRemoveEntity(c, PassportMethodActionRemove)
		},
	},
	Action:         PassportMethodActionRemove,
	RequestEntity:  &DeleteRequest{},
	ResponseEntity: &DeleteResponse{},
	TargetEntity:   &PassportMethodEntity{},
}

/**
 *	Override this function on PassportMethodEntityHttp.go,
 *	In order to add your own http
 **/
var AppendPassportMethodRouter = func(r *[]Module2Action) {}

func GetPassportMethodModule2Actions() []Module2Action {
	routes := []Module2Action{
		PASSPORT_METHOD_ACTION_QUERY,
		PASSPORT_METHOD_ACTION_EXPORT,
		PASSPORT_METHOD_ACTION_GET_ONE,
		PASSPORT_METHOD_ACTION_POST_ONE,
		PASSPORT_METHOD_ACTION_PATCH,
		PASSPORT_METHOD_ACTION_PATCH_BULK,
		PASSPORT_METHOD_ACTION_DELETE,
	}
	// Append user defined functions
	AppendPassportMethodRouter(&routes)
	return routes
}
func CreatePassportMethodRouter(r *gin.Engine) []Module2Action {
	httpRoutes := GetPassportMethodModule2Actions()
	CastRoutes(httpRoutes, r)
	WriteHttpInformationToFile(&httpRoutes, PassportMethodEntityJsonSchema, "passport-method-http", "workspaces")
	WriteEntitySchema("PassportMethodEntity", PassportMethodEntityJsonSchema, "workspaces")
	return httpRoutes
}

var PERM_ROOT_PASSPORT_METHOD_DELETE = PermissionInfo{
	CompleteKey: "root/workspaces/passport-method/delete",
	Name:        "Delete passport method",
}
var PERM_ROOT_PASSPORT_METHOD_CREATE = PermissionInfo{
	CompleteKey: "root/workspaces/passport-method/create",
	Name:        "Create passport method",
}
var PERM_ROOT_PASSPORT_METHOD_UPDATE = PermissionInfo{
	CompleteKey: "root/workspaces/passport-method/update",
	Name:        "Update passport method",
}
var PERM_ROOT_PASSPORT_METHOD_QUERY = PermissionInfo{
	CompleteKey: "root/workspaces/passport-method/query",
	Name:        "Query passport method",
}
var PERM_ROOT_PASSPORT_METHOD = PermissionInfo{
	CompleteKey: "root/workspaces/passport-method/*",
	Name:        "Entire passport method actions (*)",
}
var ALL_PASSPORT_METHOD_PERMISSIONS = []PermissionInfo{
	PERM_ROOT_PASSPORT_METHOD_DELETE,
	PERM_ROOT_PASSPORT_METHOD_CREATE,
	PERM_ROOT_PASSPORT_METHOD_UPDATE,
	PERM_ROOT_PASSPORT_METHOD_QUERY,
	PERM_ROOT_PASSPORT_METHOD,
}
var PassportMethodEntityBundle = EntityBundle{
	Permissions: ALL_PASSPORT_METHOD_PERMISSIONS,
	CliCommands: []cli.Command{
		PassportMethodCliFn(),
	},
	Actions: GetPassportMethodModule2Actions(),
	AutoMigrationEntities: []interface{}{
		&PassportMethodEntity{},
	},
}
