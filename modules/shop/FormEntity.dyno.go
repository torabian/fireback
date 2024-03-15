package shop

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
	mocks "github.com/torabian/fireback/modules/shop/mocks/Form"
	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FormFields struct {
	Visibility       *string     `json:"visibility,omitempty" yaml:"visibility"`
	WorkspaceId      *string     `json:"workspaceId,omitempty" yaml:"workspaceId"`
	LinkerId         *string     `json:"linkerId,omitempty" yaml:"linkerId"`
	ParentId         *string     `json:"parentId,omitempty" yaml:"parentId"`
	UniqueId         string      `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
	UserId           *string     `json:"userId,omitempty" yaml:"userId"`
	Rank             int64       `json:"rank,omitempty" gorm:"type:int;name:rank"`
	Updated          int64       `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
	Created          int64       `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
	CreatedFormatted string      `json:"createdFormatted,omitempty" sql:"-" gorm:"-"`
	UpdatedFormatted string      `json:"updatedFormatted,omitempty" sql:"-" gorm:"-"`
	Form             *FormEntity `json:"form" yaml:"form"    gorm:"foreignKey:FormId;references:UniqueId"     `
	// Datenano also has a text representation
	FormId *string `json:"formId" yaml:"formId"`
	Name   *string `json:"name" yaml:"name"       `
	// Datenano also has a text representation
	Type *string `json:"type" yaml:"type"       `
	// Datenano also has a text representation
	LinkedTo *FormEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

func (x *FormFields) RootObjectName() string {
	return "FormEntity"
}

type FormEntity struct {
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
	Name             *string `json:"name" yaml:"name"       `
	// Datenano also has a text representation
	Description *string `json:"description" yaml:"description"       `
	// Datenano also has a text representation
	UiSchema *workspaces.JSON `json:"uiSchema" yaml:"uiSchema"       `
	// Datenano also has a text representation
	JsonSchema *workspaces.JSON `json:"jsonSchema" yaml:"jsonSchema"       `
	// Datenano also has a text representation
	Fields []*FormFields `json:"fields" yaml:"fields"    gorm:"foreignKey:LinkerId;references:UniqueId"     `
	// Datenano also has a text representation
	Children []*FormEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
	LinkedTo *FormEntity   `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

var FormPreloadRelations []string = []string{}
var FORM_EVENT_CREATED = "form.created"
var FORM_EVENT_UPDATED = "form.updated"
var FORM_EVENT_DELETED = "form.deleted"
var FORM_EVENTS = []string{
	FORM_EVENT_CREATED,
	FORM_EVENT_UPDATED,
	FORM_EVENT_DELETED,
}

type FormFieldMap struct {
	Name        workspaces.TranslatedString `yaml:"name"`
	Description workspaces.TranslatedString `yaml:"description"`
	UiSchema    workspaces.TranslatedString `yaml:"uiSchema"`
	JsonSchema  workspaces.TranslatedString `yaml:"jsonSchema"`
	Fields      workspaces.TranslatedString `yaml:"fields"`
}

var FormEntityMetaConfig map[string]int64 = map[string]int64{}
var FormEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&FormEntity{}))

func FormFieldsActionCreate(
	dto *FormFields,
	query workspaces.QueryDSL,
) (*FormFields, *workspaces.IError) {
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
func FormFieldsActionUpdate(
	query workspaces.QueryDSL,
	dto *FormFields,
) (*FormFields, *workspaces.IError) {
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
func FormFieldsActionGetOne(
	query workspaces.QueryDSL,
) (*FormFields, *workspaces.IError) {
	refl := reflect.ValueOf(&FormFields{})
	item, err := workspaces.GetOneEntity[FormFields](query, refl)
	return item, err
}
func entityFormFormatter(dto *FormEntity, query workspaces.QueryDSL) {
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
func FormMockEntity() *FormEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &FormEntity{
		Name:        &stringHolder,
		Description: &stringHolder,
	}
	return entity
}
func FormActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := FormMockEntity()
		_, err := FormActionCreate(entity, query)
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
func FormActionSeederInit(query workspaces.QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*FormEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &FormEntity{
		Name:        &tildaRef,
		Description: &tildaRef,
		Fields:      []*FormFields{{}},
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
func FormAssociationCreate(dto *FormEntity, query workspaces.QueryDSL) error {
	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func FormRelationContentCreate(dto *FormEntity, query workspaces.QueryDSL) error {
	return nil
}
func FormRelationContentUpdate(dto *FormEntity, query workspaces.QueryDSL) error {
	return nil
}
func FormPolyglotCreateHandler(dto *FormEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
}

/**
 * This will be validating your entity fully. Important note is that, you add validate:* tag
 * in your entity, it will automatically work here. For slices inside entity, make sure you add
 * extra line of AppendSliceErrors, otherwise they won't be detected
 */
func FormValidator(dto *FormEntity, isPatch bool) *workspaces.IError {
	err := workspaces.CommonStructValidatorPointer(dto, isPatch)
	if dto != nil && dto.Fields != nil {
		workspaces.AppendSliceErrors(dto.Fields, isPatch, "fields", err)
	}
	return err
}
func FormEntityPreSanitize(dto *FormEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
func FormEntityBeforeCreateAppend(dto *FormEntity, query workspaces.QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = workspaces.UUID()
	}
	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId
	FormRecursiveAddUniqueId(dto, query)
}
func FormRecursiveAddUniqueId(dto *FormEntity, query workspaces.QueryDSL) {
	if dto.Fields != nil && len(dto.Fields) > 0 {
		for index0 := range dto.Fields {
			if dto.Fields[index0].UniqueId == "" {
				dto.Fields[index0].UniqueId = workspaces.UUID()
			}
		}
	}
}
func FormActionBatchCreateFn(dtos []*FormEntity, query workspaces.QueryDSL) ([]*FormEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*FormEntity{}
		for _, item := range dtos {
			s, err := FormActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil
}
func FormActionCreateFn(dto *FormEntity, query workspaces.QueryDSL) (*FormEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := FormValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	FormEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	FormEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	FormPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	FormRelationContentCreate(dto, query)
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
	FormAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(FORM_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": workspaces.GetTypeString(&FormEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})
	return dto, nil
}
func FormActionGetOne(query workspaces.QueryDSL) (*FormEntity, *workspaces.IError) {
	refl := reflect.ValueOf(&FormEntity{})
	item, err := workspaces.GetOneEntity[FormEntity](query, refl)
	entityFormFormatter(item, query)
	return item, err
}
func FormActionQuery(query workspaces.QueryDSL) ([]*FormEntity, *workspaces.QueryResultMeta, error) {
	refl := reflect.ValueOf(&FormEntity{})
	items, meta, err := workspaces.QueryEntitiesPointer[FormEntity](query, refl)
	for _, item := range items {
		entityFormFormatter(item, query)
	}
	return items, meta, err
}
func FormUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *FormEntity) (*FormEntity, *workspaces.IError) {
	uniqueId := fields.UniqueId
	query.TriggerEventName = FORM_EVENT_UPDATED
	FormEntityPreSanitize(fields, query)
	var item FormEntity
	q := dbref.
		Where(&FormEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)
	err := q.UpdateColumns(fields).Error
	if err != nil {
		return nil, workspaces.GormErrorToIError(err)
	}
	query.Tx = dbref
	FormRelationContentUpdate(fields, query)
	FormPolyglotCreateHandler(fields, query)
	// @meta(update has many)
	if fields.Fields != nil {
		linkerId := uniqueId
		dbref.Debug().
			Where(&FormFields{LinkerId: &linkerId}).
			Delete(&FormFields{})
		for _, newItem := range fields.Fields {
			newItem.UniqueId = workspaces.UUID()
			newItem.LinkerId = &linkerId
			dbref.Create(&newItem)
		}
	}
	err = dbref.
		Preload(clause.Associations).
		Where(&FormEntity{UniqueId: uniqueId}).
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
func FormActionUpdateFn(query workspaces.QueryDSL, fields *FormEntity) (*FormEntity, *workspaces.IError) {
	if fields == nil {
		return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}
	// 1. Validate always
	if iError := FormValidator(fields, true); iError != nil {
		return nil, iError
	}
	FormRecursiveAddUniqueId(fields, query)
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()
		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			_, err := FormUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}
		})
		return nil, workspaces.CastToIError(vf)
	} else {
		dbref = query.Tx
		return FormUpdateExec(dbref, query, fields)
	}
}

var FormWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire forms ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		count, _ := FormActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}

func FormActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&FormEntity{})
	query.ActionRequires = []string{PERM_ROOT_FORM_DELETE}
	return workspaces.RemoveEntity[FormEntity](query, refl)
}
func FormActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error
	var count int64 = 0
	{
		subCount, subErr := workspaces.WipeCleanEntity[FormFields]()
		if subErr != nil {
			fmt.Println("Error while wiping 'FormFields'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	{
		subCount, subErr := workspaces.WipeCleanEntity[FormEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'FormEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
func FormActionBulkUpdate(
	query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[FormEntity]) (
	*workspaces.BulkRecordRequest[FormEntity], *workspaces.IError,
) {
	result := []*FormEntity{}
	err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := FormActionUpdate(query, record)
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
func (x *FormEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

var FormEntityMeta = workspaces.TableMetaData{
	EntityName:    "Form",
	ExportKey:     "forms",
	TableNameInDb: "fb_form_entities",
	EntityObject:  &FormEntity{},
	ExportStream:  FormActionExportT,
	ImportQuery:   FormActionImport,
}

func FormActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[FormEntity](query, FormActionQuery, FormPreloadRelations)
}
func FormActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[FormEntity](query, FormActionQuery, FormPreloadRelations)
}
func FormActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content FormEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := FormActionCreate(&content, query)
	return err
}

var FormCommonCliFlags = []cli.Flag{
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
		Name:     "description",
		Required: false,
		Usage:    "description",
	},
	&cli.StringSliceFlag{
		Name:     "fields",
		Required: false,
		Usage:    "fields",
	},
}
var FormCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:        "name",
		StructField: "Name",
		Required:    false,
		Usage:       "name",
		Type:        "string",
	},
	{
		Name:        "description",
		StructField: "Description",
		Required:    false,
		Usage:       "description",
		Type:        "string",
	},
}
var FormCommonCliFlagsOptional = []cli.Flag{
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
		Name:     "description",
		Required: false,
		Usage:    "description",
	},
	&cli.StringSliceFlag{
		Name:     "fields",
		Required: false,
		Usage:    "fields",
	},
}
var FormCreateCmd cli.Command = cli.Command{
	Name:    "create",
	Aliases: []string{"c"},
	Flags:   FormCommonCliFlags,
	Usage:   "Create a new template",
	Action: func(c *cli.Context) {

		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
			ActionRequires: []string{PERM_ROOT_FORM_CREATE},
		})

		entity := CastFormFromCli(c)
		if entity, err := FormActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var FormCreateInteractiveCmd cli.Command = cli.Command{
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
		entity := &FormEntity{}
		for _, item := range FormCommonInteractiveCliFlags {
			if !item.Required && c.Bool("all") == false {
				continue
			}
			result := workspaces.AskForInput(item.Name, "")
			workspaces.SetFieldString(entity, item.StructField, result)
		}
		if entity, err := FormActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var FormUpdateCmd cli.Command = cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Flags:   FormCommonCliFlagsOptional,
	Usage:   "Updates a template by passing the parameters",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastFormFromCli(c)
		if entity, err := FormActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
		return nil
	},
}

func (x FormEntity) FromCli(c *cli.Context) *FormEntity {
	return CastFormFromCli(c)
}
func CastFormFromCli(c *cli.Context) *FormEntity {
	template := &FormEntity{}
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
	if c.IsSet("description") {
		value := c.String("description")
		template.Description = &value
	}
	return template
}
func FormSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		FormActionCreate,
		reflect.ValueOf(&FormEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}
func FormImportMocks() {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		FormActionCreate,
		reflect.ValueOf(&FormEntity{}).Elem(),
		&mocks.ViewsFs,
		[]string{},
		false,
	)
}
func FormWriteQueryMock(ctx workspaces.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := FormActionQuery(f)
		result := workspaces.QueryEntitySuccessResult(f, items, count)
		workspaces.WriteMockDataToFile(lang, "", "Form", result)
	}
}

var FormImportExportCommands = []cli.Command{
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
			FormActionSeeder(query, c.Int("count"))
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
				Value: "form-seeder.yml",
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
			FormActionSeederInit(f, c.String("file"), c.String("format"))
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
				Value: "form-seeder-form.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of forms, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]FormEntity{}
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
				FormActionCreate,
				reflect.ValueOf(&FormEntity{}).Elem(),
				&mocks.ViewsFs,
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
				FormActionCreate,
				reflect.ValueOf(&FormEntity{}).Elem(),
				c.String("file"),
			)
			return nil
		},
	},
}
var FormCliCommands []cli.Command = []cli.Command{
	workspaces.GetCommonQuery2(FormActionQuery, &workspaces.SecurityModel{
		ActionRequires: []string{PERM_ROOT_FORMDATA_QUERY},
	}),
	workspaces.GetCommonTableQuery(reflect.ValueOf(&FormEntity{}).Elem(), FormActionQuery),
	FormCreateCmd,
	FormUpdateCmd,
	FormCreateInteractiveCmd,
	FormWipeCmd,
	workspaces.GetCommonRemoveQuery(reflect.ValueOf(&FormEntity{}).Elem(), FormActionRemove),
}

func FormCliFn() cli.Command {
	FormCliCommands = append(FormCliCommands, FormImportExportCommands...)
	return cli.Command{
		Name:        "form",
		Description: "Forms module actions (sample module to handle complex entities)",
		Usage:       "",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: FormCliCommands,
	}
}

/**
 *	Override this function on FormEntityHttp.go,
 *	In order to add your own http
 **/
var AppendFormRouter = func(r *[]workspaces.Module2Action) {}

func GetFormModule2Actions() []workspaces.Module2Action {
	routes := []workspaces.Module2Action{
		{
			Method: "GET",
			Url:    "/forms",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FORM_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpQueryEntity(c, FormActionQuery)
				},
			},
			Format:         "QUERY",
			Action:         FormActionQuery,
			ResponseEntity: &[]FormEntity{},
		},
		{
			Method: "GET",
			Url:    "/forms/export",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FORM_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpStreamFileChannel(c, FormActionExport)
				},
			},
			Format:         "QUERY",
			Action:         FormActionExport,
			ResponseEntity: &[]FormEntity{},
		},
		{
			Method: "GET",
			Url:    "/form/:uniqueId",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FORM_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpGetEntity(c, FormActionGetOne)
				},
			},
			Format:         "GET_ONE",
			Action:         FormActionGetOne,
			ResponseEntity: &FormEntity{},
		},
		{
			ActionName:    "create",
			ActionAliases: []string{"c"},
			Flags:         FormCommonCliFlags,
			Method:        "POST",
			Url:           "/form",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FORM_CREATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpPostEntity(c, FormActionCreate)
				},
			},
			Action:         FormActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &FormEntity{},
			ResponseEntity: &FormEntity{},
		},
		{
			ActionName:    "update",
			ActionAliases: []string{"u"},
			Flags:         FormCommonCliFlagsOptional,
			Method:        "PATCH",
			Url:           "/form",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FORM_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpUpdateEntity(c, FormActionUpdate)
				},
			},
			Action:         FormActionUpdate,
			RequestEntity:  &FormEntity{},
			Format:         "PATCH_ONE",
			ResponseEntity: &FormEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/forms",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FORM_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpUpdateEntities(c, FormActionBulkUpdate)
				},
			},
			Action:         FormActionBulkUpdate,
			Format:         "PATCH_BULK",
			RequestEntity:  &workspaces.BulkRecordRequest[FormEntity]{},
			ResponseEntity: &workspaces.BulkRecordRequest[FormEntity]{},
		},
		{
			Method: "DELETE",
			Url:    "/form",
			Format: "DELETE_DSL",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FORM_DELETE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpRemoveEntity(c, FormActionRemove)
				},
			},
			Action:         FormActionRemove,
			RequestEntity:  &workspaces.DeleteRequest{},
			ResponseEntity: &workspaces.DeleteResponse{},
			TargetEntity:   &FormEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/form/:linkerId/fields/:uniqueId",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FORM_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(
					c *gin.Context,
				) {
					workspaces.HttpUpdateEntity(c, FormFieldsActionUpdate)
				},
			},
			Action:         FormFieldsActionUpdate,
			Format:         "PATCH_ONE",
			RequestEntity:  &FormFields{},
			ResponseEntity: &FormFields{},
		},
		{
			Method: "GET",
			Url:    "/form/fields/:linkerId/:uniqueId",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FORM_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(
					c *gin.Context,
				) {
					workspaces.HttpGetEntity(c, FormFieldsActionGetOne)
				},
			},
			Action:         FormFieldsActionGetOne,
			Format:         "GET_ONE",
			ResponseEntity: &FormFields{},
		},
		{
			Method: "POST",
			Url:    "/form/:linkerId/fields",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FORM_CREATE},
			},
			Handlers: []gin.HandlerFunc{
				func(
					c *gin.Context,
				) {
					workspaces.HttpPostEntity(c, FormFieldsActionCreate)
				},
			},
			Action:         FormFieldsActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &FormFields{},
			ResponseEntity: &FormFields{},
		},
	}
	// Append user defined functions
	AppendFormRouter(&routes)
	return routes
}
func CreateFormRouter(r *gin.Engine) []workspaces.Module2Action {
	httpRoutes := GetFormModule2Actions()
	workspaces.CastRoutes(httpRoutes, r)
	workspaces.WriteHttpInformationToFile(&httpRoutes, FormEntityJsonSchema, "form-http", "shop")
	workspaces.WriteEntitySchema("FormEntity", FormEntityJsonSchema, "shop")
	return httpRoutes
}

var PERM_ROOT_FORM_DELETE = "root/form/delete"
var PERM_ROOT_FORM_CREATE = "root/form/create"
var PERM_ROOT_FORM_UPDATE = "root/form/update"
var PERM_ROOT_FORM_QUERY = "root/form/query"
var PERM_ROOT_FORM = "root/form"
var ALL_FORM_PERMISSIONS = []string{
	PERM_ROOT_FORM_DELETE,
	PERM_ROOT_FORM_CREATE,
	PERM_ROOT_FORM_UPDATE,
	PERM_ROOT_FORM_QUERY,
	PERM_ROOT_FORM,
}
