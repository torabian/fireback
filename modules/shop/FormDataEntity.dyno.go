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
	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FormDataValues struct {
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
	FormField        *FormFields `json:"formField" yaml:"formField"    gorm:"foreignKey:FormFieldId;references:UniqueId"     `
	// Datenano also has a text representation
	FormFieldId *string `json:"formFieldId" yaml:"formFieldId"`
	Value       *string `json:"value" yaml:"value"    gorm:"text"     `
	// Datenano also has a text representation
	ValueExcerpt *string         `json:"valueExcerpt" yaml:"valueExcerpt"`
	LinkedTo     *FormDataEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

func (x *FormDataValues) RootObjectName() string {
	return "FormDataEntity"
}

type FormDataEntity struct {
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
	FormId *string           `json:"formId" yaml:"formId" validate:"required" `
	Values []*FormDataValues `json:"values" yaml:"values"    gorm:"foreignKey:LinkerId;references:UniqueId"     `
	// Datenano also has a text representation
	Children []*FormDataEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
	LinkedTo *FormDataEntity   `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

var FormDataPreloadRelations []string = []string{}
var FORMDATA_EVENT_CREATED = "formData.created"
var FORMDATA_EVENT_UPDATED = "formData.updated"
var FORMDATA_EVENT_DELETED = "formData.deleted"
var FORMDATA_EVENTS = []string{
	FORMDATA_EVENT_CREATED,
	FORMDATA_EVENT_UPDATED,
	FORMDATA_EVENT_DELETED,
}

type FormDataFieldMap struct {
	Form   workspaces.TranslatedString `yaml:"form"`
	Values workspaces.TranslatedString `yaml:"values"`
}

var FormDataEntityMetaConfig map[string]int64 = map[string]int64{}
var FormDataEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&FormDataEntity{}))

func FormDataValuesActionCreate(
	dto *FormDataValues,
	query workspaces.QueryDSL,
) (*FormDataValues, *workspaces.IError) {
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
func FormDataValuesActionUpdate(
	query workspaces.QueryDSL,
	dto *FormDataValues,
) (*FormDataValues, *workspaces.IError) {
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
func FormDataValuesActionGetOne(
	query workspaces.QueryDSL,
) (*FormDataValues, *workspaces.IError) {
	refl := reflect.ValueOf(&FormDataValues{})
	item, err := workspaces.GetOneEntity[FormDataValues](query, refl)
	return item, err
}
func entityFormDataFormatter(dto *FormDataEntity, query workspaces.QueryDSL) {
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
func FormDataMockEntity() *FormDataEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &FormDataEntity{}
	return entity
}
func FormDataActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := FormDataMockEntity()
		_, err := FormDataActionCreate(entity, query)
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
func FormDataActionSeederInit(query workspaces.QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*FormDataEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &FormDataEntity{
		Values: []*FormDataValues{{}},
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
func FormDataAssociationCreate(dto *FormDataEntity, query workspaces.QueryDSL) error {
	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func FormDataRelationContentCreate(dto *FormDataEntity, query workspaces.QueryDSL) error {
	return nil
}
func FormDataRelationContentUpdate(dto *FormDataEntity, query workspaces.QueryDSL) error {
	return nil
}
func FormDataPolyglotCreateHandler(dto *FormDataEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
}

/**
 * This will be validating your entity fully. Important note is that, you add validate:* tag
 * in your entity, it will automatically work here. For slices inside entity, make sure you add
 * extra line of AppendSliceErrors, otherwise they won't be detected
 */
func FormDataValidator(dto *FormDataEntity, isPatch bool) *workspaces.IError {
	err := workspaces.CommonStructValidatorPointer(dto, isPatch)
	if dto != nil && dto.Values != nil {
		workspaces.AppendSliceErrors(dto.Values, isPatch, "values", err)
	}
	return err
}
func FormDataEntityPreSanitize(dto *FormDataEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
func FormDataEntityBeforeCreateAppend(dto *FormDataEntity, query workspaces.QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = workspaces.UUID()
	}
	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId
	FormDataRecursiveAddUniqueId(dto, query)
}
func FormDataRecursiveAddUniqueId(dto *FormDataEntity, query workspaces.QueryDSL) {
	if dto.Values != nil && len(dto.Values) > 0 {
		for index0 := range dto.Values {
			if dto.Values[index0].UniqueId == "" {
				dto.Values[index0].UniqueId = workspaces.UUID()
			}
		}
	}
}
func FormDataActionBatchCreateFn(dtos []*FormDataEntity, query workspaces.QueryDSL) ([]*FormDataEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*FormDataEntity{}
		for _, item := range dtos {
			s, err := FormDataActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil
}
func FormDataActionCreateFn(dto *FormDataEntity, query workspaces.QueryDSL) (*FormDataEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := FormDataValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	FormDataEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	FormDataEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	FormDataPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	FormDataRelationContentCreate(dto, query)
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
	FormDataAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(FORMDATA_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": workspaces.GetTypeString(&FormDataEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})
	return dto, nil
}
func FormDataActionGetOne(query workspaces.QueryDSL) (*FormDataEntity, *workspaces.IError) {
	refl := reflect.ValueOf(&FormDataEntity{})
	item, err := workspaces.GetOneEntity[FormDataEntity](query, refl)
	entityFormDataFormatter(item, query)
	return item, err
}
func FormDataActionQuery(query workspaces.QueryDSL) ([]*FormDataEntity, *workspaces.QueryResultMeta, error) {
	refl := reflect.ValueOf(&FormDataEntity{})
	items, meta, err := workspaces.QueryEntitiesPointer[FormDataEntity](query, refl)
	for _, item := range items {
		entityFormDataFormatter(item, query)
	}
	return items, meta, err
}
func FormDataUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *FormDataEntity) (*FormDataEntity, *workspaces.IError) {
	uniqueId := fields.UniqueId
	query.TriggerEventName = FORMDATA_EVENT_UPDATED
	FormDataEntityPreSanitize(fields, query)
	var item FormDataEntity
	q := dbref.
		Where(&FormDataEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)
	err := q.UpdateColumns(fields).Error
	if err != nil {
		return nil, workspaces.GormErrorToIError(err)
	}
	query.Tx = dbref
	FormDataRelationContentUpdate(fields, query)
	FormDataPolyglotCreateHandler(fields, query)
	// @meta(update has many)
	if fields.Values != nil {
		linkerId := uniqueId
		dbref.Debug().
			Where(&FormDataValues{LinkerId: &linkerId}).
			Delete(&FormDataValues{})
		for _, newItem := range fields.Values {
			newItem.UniqueId = workspaces.UUID()
			newItem.LinkerId = &linkerId
			dbref.Create(&newItem)
		}
	}
	err = dbref.
		Preload(clause.Associations).
		Where(&FormDataEntity{UniqueId: uniqueId}).
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
func FormDataActionUpdateFn(query workspaces.QueryDSL, fields *FormDataEntity) (*FormDataEntity, *workspaces.IError) {
	if fields == nil {
		return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}
	// 1. Validate always
	if iError := FormDataValidator(fields, true); iError != nil {
		return nil, iError
	}
	FormDataRecursiveAddUniqueId(fields, query)
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()
		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			_, err := FormDataUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}
		})
		return nil, workspaces.CastToIError(vf)
	} else {
		dbref = query.Tx
		return FormDataUpdateExec(dbref, query, fields)
	}
}

var FormDataWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire formdata ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		count, _ := FormDataActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}

func FormDataActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&FormDataEntity{})
	query.ActionRequires = []string{PERM_ROOT_FORMDATA_DELETE}
	return workspaces.RemoveEntity[FormDataEntity](query, refl)
}
func FormDataActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error
	var count int64 = 0
	{
		subCount, subErr := workspaces.WipeCleanEntity[FormDataValues]()
		if subErr != nil {
			fmt.Println("Error while wiping 'FormDataValues'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	{
		subCount, subErr := workspaces.WipeCleanEntity[FormDataEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'FormDataEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
func FormDataActionBulkUpdate(
	query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[FormDataEntity]) (
	*workspaces.BulkRecordRequest[FormDataEntity], *workspaces.IError,
) {
	result := []*FormDataEntity{}
	err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := FormDataActionUpdate(query, record)
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
func (x *FormDataEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

var FormDataEntityMeta = workspaces.TableMetaData{
	EntityName:    "FormData",
	ExportKey:     "form-data",
	TableNameInDb: "fb_formdata_entities",
	EntityObject:  &FormDataEntity{},
	ExportStream:  FormDataActionExportT,
	ImportQuery:   FormDataActionImport,
}

func FormDataActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[FormDataEntity](query, FormDataActionQuery, FormDataPreloadRelations)
}
func FormDataActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[FormDataEntity](query, FormDataActionQuery, FormDataPreloadRelations)
}
func FormDataActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content FormDataEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := FormDataActionCreate(&content, query)
	return err
}

var FormDataCommonCliFlags = []cli.Flag{
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
		Name:     "form-id",
		Required: true,
		Usage:    "form",
	},
	&cli.StringSliceFlag{
		Name:     "values",
		Required: false,
		Usage:    "values",
	},
}
var FormDataCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{}
var FormDataCommonCliFlagsOptional = []cli.Flag{
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
		Name:     "form-id",
		Required: true,
		Usage:    "form",
	},
	&cli.StringSliceFlag{
		Name:     "values",
		Required: false,
		Usage:    "values",
	},
}
var FormDataCreateCmd cli.Command = cli.Command{
	Name:    "create",
	Aliases: []string{"c"},
	Flags:   FormDataCommonCliFlags,
	Usage:   "Create a new template",
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastFormDataFromCli(c)
		if entity, err := FormDataActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var FormDataCreateInteractiveCmd cli.Command = cli.Command{
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
		entity := &FormDataEntity{}
		for _, item := range FormDataCommonInteractiveCliFlags {
			if !item.Required && c.Bool("all") == false {
				continue
			}
			result := workspaces.AskForInput(item.Name, "")
			workspaces.SetFieldString(entity, item.StructField, result)
		}
		if entity, err := FormDataActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var FormDataUpdateCmd cli.Command = cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Flags:   FormDataCommonCliFlagsOptional,
	Usage:   "Updates a template by passing the parameters",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastFormDataFromCli(c)
		if entity, err := FormDataActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
		return nil
	},
}

func (x FormDataEntity) FromCli(c *cli.Context) *FormDataEntity {
	return CastFormDataFromCli(c)
}
func CastFormDataFromCli(c *cli.Context) *FormDataEntity {
	template := &FormDataEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
	if c.IsSet("form-id") {
		value := c.String("form-id")
		template.FormId = &value
	}
	return template
}
func FormDataSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		FormDataActionCreate,
		reflect.ValueOf(&FormDataEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}
func FormDataWriteQueryMock(ctx workspaces.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := FormDataActionQuery(f)
		result := workspaces.QueryEntitySuccessResult(f, items, count)
		workspaces.WriteMockDataToFile(lang, "", "FormData", result)
	}
}

var FormDataImportExportCommands = []cli.Command{
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
			FormDataActionSeeder(query, c.Int("count"))
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
				Value: "form-data-seeder.yml",
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
			FormDataActionSeederInit(f, c.String("file"), c.String("format"))
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
				Value: "form-data-seeder-form-data.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of form-data, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]FormDataEntity{}
			workspaces.ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
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
				FormDataActionCreate,
				reflect.ValueOf(&FormDataEntity{}).Elem(),
				c.String("file"),
			)
			return nil
		},
	},
}
var FormDataCliCommands []cli.Command = []cli.Command{
	workspaces.GetCommonQuery(FormDataActionQuery),
	workspaces.GetCommonTableQuery(reflect.ValueOf(&FormDataEntity{}).Elem(), FormDataActionQuery),
	FormDataCreateCmd,
	FormDataUpdateCmd,
	FormDataCreateInteractiveCmd,
	FormDataWipeCmd,
	workspaces.GetCommonRemoveQuery(reflect.ValueOf(&FormDataEntity{}).Elem(), FormDataActionRemove),
}

func FormDataCliFn() cli.Command {
	FormDataCliCommands = append(FormDataCliCommands, FormDataImportExportCommands...)
	return cli.Command{
		Name:        "formData",
		Description: "FormDatas module actions (sample module to handle complex entities)",
		Usage:       "",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: FormDataCliCommands,
	}
}

/**
 *	Override this function on FormDataEntityHttp.go,
 *	In order to add your own http
 **/
var AppendFormDataRouter = func(r *[]workspaces.Module2Action) {}

func GetFormDataModule2Actions() []workspaces.Module2Action {
	routes := []workspaces.Module2Action{
		{
			Method: "GET",
			Url:    "/form-data",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FORMDATA_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpQueryEntity(c, FormDataActionQuery)
				},
			},
			Format:         "QUERY",
			Action:         FormDataActionQuery,
			ResponseEntity: &[]FormDataEntity{},
		},
		{
			Method: "GET",
			Url:    "/form-data/export",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FORMDATA_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpStreamFileChannel(c, FormDataActionExport)
				},
			},
			Format:         "QUERY",
			Action:         FormDataActionExport,
			ResponseEntity: &[]FormDataEntity{},
		},
		{
			Method: "GET",
			Url:    "/form-data/:uniqueId",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FORMDATA_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpGetEntity(c, FormDataActionGetOne)
				},
			},
			Format:         "GET_ONE",
			Action:         FormDataActionGetOne,
			ResponseEntity: &FormDataEntity{},
		},
		{
			ActionName:    "create",
			ActionAliases: []string{"c"},
			Flags:         FormDataCommonCliFlags,
			Method:        "POST",
			Url:           "/form-data",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FORMDATA_CREATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpPostEntity(c, FormDataActionCreate)
				},
			},
			Action:         FormDataActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &FormDataEntity{},
			ResponseEntity: &FormDataEntity{},
		},
		{
			ActionName:    "update",
			ActionAliases: []string{"u"},
			Flags:         FormDataCommonCliFlagsOptional,
			Method:        "PATCH",
			Url:           "/form-data",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FORMDATA_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpUpdateEntity(c, FormDataActionUpdate)
				},
			},
			Action:         FormDataActionUpdate,
			RequestEntity:  &FormDataEntity{},
			Format:         "PATCH_ONE",
			ResponseEntity: &FormDataEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/form-data2",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FORMDATA_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpUpdateEntities(c, FormDataActionBulkUpdate)
				},
			},
			Action:         FormDataActionBulkUpdate,
			Format:         "PATCH_BULK",
			RequestEntity:  &workspaces.BulkRecordRequest[FormDataEntity]{},
			ResponseEntity: &workspaces.BulkRecordRequest[FormDataEntity]{},
		},
		{
			Method: "DELETE",
			Url:    "/form-data",
			Format: "DELETE_DSL",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FORMDATA_DELETE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpRemoveEntity(c, FormDataActionRemove)
				},
			},
			Action:         FormDataActionRemove,
			RequestEntity:  &workspaces.DeleteRequest{},
			ResponseEntity: &workspaces.DeleteResponse{},
			TargetEntity:   &FormDataEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/form-data/:linkerId/values/:uniqueId",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FORMDATA_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(
					c *gin.Context,
				) {
					workspaces.HttpUpdateEntity(c, FormDataValuesActionUpdate)
				},
			},
			Action:         FormDataValuesActionUpdate,
			Format:         "PATCH_ONE",
			RequestEntity:  &FormDataValues{},
			ResponseEntity: &FormDataValues{},
		},
		{
			Method: "GET",
			Url:    "/form-data/values/:linkerId/:uniqueId",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FORMDATA_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(
					c *gin.Context,
				) {
					workspaces.HttpGetEntity(c, FormDataValuesActionGetOne)
				},
			},
			Action:         FormDataValuesActionGetOne,
			Format:         "GET_ONE",
			ResponseEntity: &FormDataValues{},
		},
		{
			Method: "POST",
			Url:    "/form-data/:linkerId/values",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FORMDATA_CREATE},
			},
			Handlers: []gin.HandlerFunc{
				func(
					c *gin.Context,
				) {
					workspaces.HttpPostEntity(c, FormDataValuesActionCreate)
				},
			},
			Action:         FormDataValuesActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &FormDataValues{},
			ResponseEntity: &FormDataValues{},
		},
	}
	// Append user defined functions
	AppendFormDataRouter(&routes)
	return routes
}
func CreateFormDataRouter(r *gin.Engine) []workspaces.Module2Action {
	httpRoutes := GetFormDataModule2Actions()
	workspaces.CastRoutes(httpRoutes, r)
	workspaces.WriteHttpInformationToFile(&httpRoutes, FormDataEntityJsonSchema, "form-data-http", "shop")
	workspaces.WriteEntitySchema("FormDataEntity", FormDataEntityJsonSchema, "shop")
	return httpRoutes
}

var PERM_ROOT_FORMDATA_DELETE = "root/formdata/delete"
var PERM_ROOT_FORMDATA_CREATE = "root/formdata/create"
var PERM_ROOT_FORMDATA_UPDATE = "root/formdata/update"
var PERM_ROOT_FORMDATA_QUERY = "root/formdata/query"
var PERM_ROOT_FORMDATA = "root/formdata"
var ALL_FORMDATA_PERMISSIONS = []string{
	PERM_ROOT_FORMDATA_DELETE,
	PERM_ROOT_FORMDATA_CREATE,
	PERM_ROOT_FORMDATA_UPDATE,
	PERM_ROOT_FORMDATA_QUERY,
	PERM_ROOT_FORMDATA,
}
