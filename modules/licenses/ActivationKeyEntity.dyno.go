package licenses

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

type ActivationKeyEntity struct {
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
	Series           *string `json:"series" yaml:"series"       `
	// Datenano also has a text representation
	Used *int64 `json:"used" yaml:"used"       `
	// Datenano also has a text representation
	Plan *ProductPlanEntity `json:"plan" yaml:"plan"    gorm:"foreignKey:PlanId;references:UniqueId"     `
	// Datenano also has a text representation
	PlanId   *string                `json:"planId" yaml:"planId"`
	Children []*ActivationKeyEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
	LinkedTo *ActivationKeyEntity   `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

var ActivationKeyPreloadRelations []string = []string{}
var ACTIVATIONKEY_EVENT_CREATED = "activationKey.created"
var ACTIVATIONKEY_EVENT_UPDATED = "activationKey.updated"
var ACTIVATIONKEY_EVENT_DELETED = "activationKey.deleted"
var ACTIVATIONKEY_EVENTS = []string{
	ACTIVATIONKEY_EVENT_CREATED,
	ACTIVATIONKEY_EVENT_UPDATED,
	ACTIVATIONKEY_EVENT_DELETED,
}

type ActivationKeyFieldMap struct {
	Series workspaces.TranslatedString `yaml:"series"`
	Used   workspaces.TranslatedString `yaml:"used"`
	Plan   workspaces.TranslatedString `yaml:"plan"`
}

var ActivationKeyEntityMetaConfig map[string]int64 = map[string]int64{}
var ActivationKeyEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&ActivationKeyEntity{}))

func entityActivationKeyFormatter(dto *ActivationKeyEntity, query workspaces.QueryDSL) {
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
func ActivationKeyMockEntity() *ActivationKeyEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &ActivationKeyEntity{
		Series: &stringHolder,
		Used:   &int64Holder,
	}
	return entity
}
func ActivationKeyActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := ActivationKeyMockEntity()
		_, err := ActivationKeyActionCreate(entity, query)
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
func ActivationKeyActionSeederInit(query workspaces.QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*ActivationKeyEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &ActivationKeyEntity{
		Series: &tildaRef,
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
func ActivationKeyAssociationCreate(dto *ActivationKeyEntity, query workspaces.QueryDSL) error {
	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func ActivationKeyRelationContentCreate(dto *ActivationKeyEntity, query workspaces.QueryDSL) error {
	return nil
}
func ActivationKeyRelationContentUpdate(dto *ActivationKeyEntity, query workspaces.QueryDSL) error {
	return nil
}
func ActivationKeyPolyglotCreateHandler(dto *ActivationKeyEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
}

/**
 * This will be validating your entity fully. Important note is that, you add validate:* tag
 * in your entity, it will automatically work here. For slices inside entity, make sure you add
 * extra line of AppendSliceErrors, otherwise they won't be detected
 */
func ActivationKeyValidator(dto *ActivationKeyEntity, isPatch bool) *workspaces.IError {
	err := workspaces.CommonStructValidatorPointer(dto, isPatch)
	return err
}
func ActivationKeyEntityPreSanitize(dto *ActivationKeyEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
func ActivationKeyEntityBeforeCreateAppend(dto *ActivationKeyEntity, query workspaces.QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = workspaces.UUID()
	}
	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId
	ActivationKeyRecursiveAddUniqueId(dto, query)
}
func ActivationKeyRecursiveAddUniqueId(dto *ActivationKeyEntity, query workspaces.QueryDSL) {
}
func ActivationKeyActionBatchCreateFn(dtos []*ActivationKeyEntity, query workspaces.QueryDSL) ([]*ActivationKeyEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*ActivationKeyEntity{}
		for _, item := range dtos {
			s, err := ActivationKeyActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil
}
func ActivationKeyActionCreateFn(dto *ActivationKeyEntity, query workspaces.QueryDSL) (*ActivationKeyEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := ActivationKeyValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	ActivationKeyEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	ActivationKeyEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	ActivationKeyPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	ActivationKeyRelationContentCreate(dto, query)
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
	ActivationKeyAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(ACTIVATIONKEY_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": workspaces.GetTypeString(&ActivationKeyEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})
	return dto, nil
}
func ActivationKeyActionGetOne(query workspaces.QueryDSL) (*ActivationKeyEntity, *workspaces.IError) {
	refl := reflect.ValueOf(&ActivationKeyEntity{})
	item, err := workspaces.GetOneEntity[ActivationKeyEntity](query, refl)
	entityActivationKeyFormatter(item, query)
	return item, err
}
func ActivationKeyActionQuery(query workspaces.QueryDSL) ([]*ActivationKeyEntity, *workspaces.QueryResultMeta, error) {
	refl := reflect.ValueOf(&ActivationKeyEntity{})
	items, meta, err := workspaces.QueryEntitiesPointer[ActivationKeyEntity](query, refl)
	for _, item := range items {
		entityActivationKeyFormatter(item, query)
	}
	return items, meta, err
}
func ActivationKeyUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *ActivationKeyEntity) (*ActivationKeyEntity, *workspaces.IError) {
	uniqueId := fields.UniqueId
	query.TriggerEventName = ACTIVATIONKEY_EVENT_UPDATED
	ActivationKeyEntityPreSanitize(fields, query)
	var item ActivationKeyEntity
	q := dbref.
		Where(&ActivationKeyEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)
	err := q.UpdateColumns(fields).Error
	if err != nil {
		return nil, workspaces.GormErrorToIError(err)
	}
	query.Tx = dbref
	ActivationKeyRelationContentUpdate(fields, query)
	ActivationKeyPolyglotCreateHandler(fields, query)
	// @meta(update has many)
	err = dbref.
		Preload(clause.Associations).
		Where(&ActivationKeyEntity{UniqueId: uniqueId}).
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
func ActivationKeyActionUpdateFn(query workspaces.QueryDSL, fields *ActivationKeyEntity) (*ActivationKeyEntity, *workspaces.IError) {
	if fields == nil {
		return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}
	// 1. Validate always
	if iError := ActivationKeyValidator(fields, true); iError != nil {
		return nil, iError
	}
	ActivationKeyRecursiveAddUniqueId(fields, query)
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()
		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			_, err := ActivationKeyUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}
		})
		return nil, workspaces.CastToIError(vf)
	} else {
		dbref = query.Tx
		return ActivationKeyUpdateExec(dbref, query, fields)
	}
}

var ActivationKeyWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire activationkeys ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		count, _ := ActivationKeyActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}

func ActivationKeyActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&ActivationKeyEntity{})
	query.ActionRequires = []string{PERM_ROOT_ACTIVATIONKEY_DELETE}
	return workspaces.RemoveEntity[ActivationKeyEntity](query, refl)
}
func ActivationKeyActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error
	var count int64 = 0
	{
		subCount, subErr := workspaces.WipeCleanEntity[ActivationKeyEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'ActivationKeyEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
func ActivationKeyActionBulkUpdate(
	query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[ActivationKeyEntity]) (
	*workspaces.BulkRecordRequest[ActivationKeyEntity], *workspaces.IError,
) {
	result := []*ActivationKeyEntity{}
	err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := ActivationKeyActionUpdate(query, record)
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
func (x *ActivationKeyEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

var ActivationKeyEntityMeta = workspaces.TableMetaData{
	EntityName:    "ActivationKey",
	ExportKey:     "activation-keys",
	TableNameInDb: "fb_activationkey_entities",
	EntityObject:  &ActivationKeyEntity{},
	ExportStream:  ActivationKeyActionExportT,
	ImportQuery:   ActivationKeyActionImport,
}

func ActivationKeyActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[ActivationKeyEntity](query, ActivationKeyActionQuery, ActivationKeyPreloadRelations)
}
func ActivationKeyActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[ActivationKeyEntity](query, ActivationKeyActionQuery, ActivationKeyPreloadRelations)
}
func ActivationKeyActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content ActivationKeyEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := ActivationKeyActionCreate(&content, query)
	return err
}

var ActivationKeyCommonCliFlags = []cli.Flag{
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
		Name:     "series",
		Required: false,
		Usage:    "series",
	},
	&cli.Int64Flag{
		Name:     "used",
		Required: false,
		Usage:    "used",
	},
	&cli.StringFlag{
		Name:     "plan-id",
		Required: false,
		Usage:    "plan",
	},
}
var ActivationKeyCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:        "series",
		StructField: "Series",
		Required:    false,
		Usage:       "series",
		Type:        "string",
	},
	{
		Name:        "used",
		StructField: "Used",
		Required:    false,
		Usage:       "used",
		Type:        "int64",
	},
}
var ActivationKeyCommonCliFlagsOptional = []cli.Flag{
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
		Name:     "series",
		Required: false,
		Usage:    "series",
	},
	&cli.Int64Flag{
		Name:     "used",
		Required: false,
		Usage:    "used",
	},
	&cli.StringFlag{
		Name:     "plan-id",
		Required: false,
		Usage:    "plan",
	},
}
var ActivationKeyCreateCmd cli.Command = cli.Command{
	Name:    "create",
	Aliases: []string{"c"},
	Flags:   ActivationKeyCommonCliFlags,
	Usage:   "Create a new template",
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastActivationKeyFromCli(c)
		if entity, err := ActivationKeyActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var ActivationKeyCreateInteractiveCmd cli.Command = cli.Command{
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
		entity := &ActivationKeyEntity{}
		for _, item := range ActivationKeyCommonInteractiveCliFlags {
			if !item.Required && c.Bool("all") == false {
				continue
			}
			result := workspaces.AskForInput(item.Name, "")
			workspaces.SetFieldString(entity, item.StructField, result)
		}
		if entity, err := ActivationKeyActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var ActivationKeyUpdateCmd cli.Command = cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Flags:   ActivationKeyCommonCliFlagsOptional,
	Usage:   "Updates a template by passing the parameters",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastActivationKeyFromCli(c)
		if entity, err := ActivationKeyActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
		return nil
	},
}

func CastActivationKeyFromCli(c *cli.Context) *ActivationKeyEntity {
	template := &ActivationKeyEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
	if c.IsSet("series") {
		value := c.String("series")
		template.Series = &value
	}
	if c.IsSet("plan-id") {
		value := c.String("plan-id")
		template.PlanId = &value
	}
	return template
}
func ActivationKeySyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		ActivationKeyActionCreate,
		reflect.ValueOf(&ActivationKeyEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}
func ActivationKeyWriteQueryMock(ctx workspaces.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := ActivationKeyActionQuery(f)
		result := workspaces.QueryEntitySuccessResult(f, items, count)
		workspaces.WriteMockDataToFile(lang, "", "ActivationKey", result)
	}
}

var ActivationKeyImportExportCommands = []cli.Command{
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
			ActivationKeyActionSeeder(query, c.Int("count"))
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
				Value: "activation-key-seeder.yml",
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
			ActivationKeyActionSeederInit(f, c.String("file"), c.String("format"))
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
				Value: "activation-key-seeder-activation-key.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of activation-keys, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]ActivationKeyEntity{}
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
				ActivationKeyActionCreate,
				reflect.ValueOf(&ActivationKeyEntity{}).Elem(),
				c.String("file"),
			)
			return nil
		},
	},
}
var ActivationKeyCliCommands []cli.Command = []cli.Command{
	workspaces.GetCommonQuery(ActivationKeyActionQuery),
	workspaces.GetCommonTableQuery(reflect.ValueOf(&ActivationKeyEntity{}).Elem(), ActivationKeyActionQuery),
	ActivationKeyCreateCmd,
	ActivationKeyUpdateCmd,
	ActivationKeyCreateInteractiveCmd,
	ActivationKeyWipeCmd,
	workspaces.GetCommonRemoveQuery(reflect.ValueOf(&ActivationKeyEntity{}).Elem(), ActivationKeyActionRemove),
}

func ActivationKeyCliFn() cli.Command {
	ActivationKeyCliCommands = append(ActivationKeyCliCommands, ActivationKeyImportExportCommands...)
	return cli.Command{
		Name:        "key",
		Description: "ActivationKeys module actions (sample module to handle complex entities)",
		Usage:       "",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: ActivationKeyCliCommands,
	}
}

/**
 *	Override this function on ActivationKeyEntityHttp.go,
 *	In order to add your own http
 **/
var AppendActivationKeyRouter = func(r *[]workspaces.Module2Action) {}

func GetActivationKeyModule2Actions() []workspaces.Module2Action {
	routes := []workspaces.Module2Action{
		{
			Method: "GET",
			Url:    "/activation-keys",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_ACTIVATIONKEY_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpQueryEntity(c, ActivationKeyActionQuery)
				},
			},
			Format:         "QUERY",
			Action:         ActivationKeyActionQuery,
			ResponseEntity: &[]ActivationKeyEntity{},
		},
		{
			Method: "GET",
			Url:    "/activation-keys/export",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_ACTIVATIONKEY_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpStreamFileChannel(c, ActivationKeyActionExport)
				},
			},
			Format:         "QUERY",
			Action:         ActivationKeyActionExport,
			ResponseEntity: &[]ActivationKeyEntity{},
		},
		{
			Method: "GET",
			Url:    "/activation-key/:uniqueId",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_ACTIVATIONKEY_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpGetEntity(c, ActivationKeyActionGetOne)
				},
			},
			Format:         "GET_ONE",
			Action:         ActivationKeyActionGetOne,
			ResponseEntity: &ActivationKeyEntity{},
		},
		{
			Method: "POST",
			Url:    "/activation-key",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_ACTIVATIONKEY_CREATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpPostEntity(c, ActivationKeyActionCreate)
				},
			},
			Action:         ActivationKeyActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &ActivationKeyEntity{},
			ResponseEntity: &ActivationKeyEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/activation-key",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_ACTIVATIONKEY_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpUpdateEntity(c, ActivationKeyActionUpdate)
				},
			},
			Action:         ActivationKeyActionUpdate,
			RequestEntity:  &ActivationKeyEntity{},
			Format:         "PATCH_ONE",
			ResponseEntity: &ActivationKeyEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/activation-keys",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_ACTIVATIONKEY_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpUpdateEntities(c, ActivationKeyActionBulkUpdate)
				},
			},
			Action:         ActivationKeyActionBulkUpdate,
			Format:         "PATCH_BULK",
			RequestEntity:  &workspaces.BulkRecordRequest[ActivationKeyEntity]{},
			ResponseEntity: &workspaces.BulkRecordRequest[ActivationKeyEntity]{},
		},
		{
			Method: "DELETE",
			Url:    "/activation-key",
			Format: "DELETE_DSL",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_ACTIVATIONKEY_DELETE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpRemoveEntity(c, ActivationKeyActionRemove)
				},
			},
			Action:         ActivationKeyActionRemove,
			RequestEntity:  &workspaces.DeleteRequest{},
			ResponseEntity: &workspaces.DeleteResponse{},
			TargetEntity:   &ActivationKeyEntity{},
		},
	}
	// Append user defined functions
	AppendActivationKeyRouter(&routes)
	return routes
}
func CreateActivationKeyRouter(r *gin.Engine) []workspaces.Module2Action {
	httpRoutes := GetActivationKeyModule2Actions()
	workspaces.CastRoutes(httpRoutes, r)
	workspaces.WriteHttpInformationToFile(&httpRoutes, ActivationKeyEntityJsonSchema, "activation-key-http", "licenses")
	workspaces.WriteEntitySchema("ActivationKeyEntity", ActivationKeyEntityJsonSchema, "licenses")
	return httpRoutes
}

var PERM_ROOT_ACTIVATIONKEY_DELETE = "root/activationkey/delete"
var PERM_ROOT_ACTIVATIONKEY_CREATE = "root/activationkey/create"
var PERM_ROOT_ACTIVATIONKEY_UPDATE = "root/activationkey/update"
var PERM_ROOT_ACTIVATIONKEY_QUERY = "root/activationkey/query"
var PERM_ROOT_ACTIVATIONKEY = "root/activationkey"
var ALL_ACTIVATIONKEY_PERMISSIONS = []string{
	PERM_ROOT_ACTIVATIONKEY_DELETE,
	PERM_ROOT_ACTIVATIONKEY_CREATE,
	PERM_ROOT_ACTIVATIONKEY_UPDATE,
	PERM_ROOT_ACTIVATIONKEY_QUERY,
	PERM_ROOT_ACTIVATIONKEY,
}
