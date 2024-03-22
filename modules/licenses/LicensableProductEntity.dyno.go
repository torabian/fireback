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
	mocks "github.com/torabian/fireback/modules/licenses/mocks/LicensableProduct"
	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LicensableProductEntity struct {
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
	Name             *string `json:"name" yaml:"name"  validate:"required,omitempty,min=1,max=100"        translate:"true" `
	// Datenano also has a text representation
	PrivateKey *string `json:"privateKey" yaml:"privateKey"       `
	// Datenano also has a text representation
	PublicKey *string `json:"publicKey" yaml:"publicKey"       `
	// Datenano also has a text representation
	Translations []*LicensableProductEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
	Children     []*LicensableProductEntity         `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
	LinkedTo     *LicensableProductEntity           `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

var LicensableProductPreloadRelations []string = []string{}
var LICENSABLEPRODUCT_EVENT_CREATED = "licensableProduct.created"
var LICENSABLEPRODUCT_EVENT_UPDATED = "licensableProduct.updated"
var LICENSABLEPRODUCT_EVENT_DELETED = "licensableProduct.deleted"
var LICENSABLEPRODUCT_EVENTS = []string{
	LICENSABLEPRODUCT_EVENT_CREATED,
	LICENSABLEPRODUCT_EVENT_UPDATED,
	LICENSABLEPRODUCT_EVENT_DELETED,
}

type LicensableProductFieldMap struct {
	Name       workspaces.TranslatedString `yaml:"name"`
	PrivateKey workspaces.TranslatedString `yaml:"privateKey"`
	PublicKey  workspaces.TranslatedString `yaml:"publicKey"`
}

var LicensableProductEntityMetaConfig map[string]int64 = map[string]int64{}
var LicensableProductEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&LicensableProductEntity{}))

type LicensableProductEntityPolyglot struct {
	LinkerId   string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
	LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
	Name       string `yaml:"name" json:"name"`
}

func entityLicensableProductFormatter(dto *LicensableProductEntity, query workspaces.QueryDSL) {
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
func LicensableProductMockEntity() *LicensableProductEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &LicensableProductEntity{
		Name:       &stringHolder,
		PrivateKey: &stringHolder,
		PublicKey:  &stringHolder,
	}
	return entity
}
func LicensableProductActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := LicensableProductMockEntity()
		_, err := LicensableProductActionCreate(entity, query)
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
func (x *LicensableProductEntity) GetNameTranslated(language string) string {
	if x.Translations != nil && len(x.Translations) > 0 {
		for _, item := range x.Translations {
			if item.LanguageId == language {
				return item.Name
			}
		}
	}
	return ""
}
func LicensableProductActionSeederInit(query workspaces.QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*LicensableProductEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &LicensableProductEntity{
		Name:       &tildaRef,
		PrivateKey: &tildaRef,
		PublicKey:  &tildaRef,
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
func LicensableProductAssociationCreate(dto *LicensableProductEntity, query workspaces.QueryDSL) error {
	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func LicensableProductRelationContentCreate(dto *LicensableProductEntity, query workspaces.QueryDSL) error {
	return nil
}
func LicensableProductRelationContentUpdate(dto *LicensableProductEntity, query workspaces.QueryDSL) error {
	return nil
}
func LicensableProductPolyglotCreateHandler(dto *LicensableProductEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
	workspaces.PolyglotCreateHandler(dto, &LicensableProductEntityPolyglot{}, query)
}

/**
 * This will be validating your entity fully. Important note is that, you add validate:* tag
 * in your entity, it will automatically work here. For slices inside entity, make sure you add
 * extra line of AppendSliceErrors, otherwise they won't be detected
 */
func LicensableProductValidator(dto *LicensableProductEntity, isPatch bool) *workspaces.IError {
	err := workspaces.CommonStructValidatorPointer(dto, isPatch)
	return err
}
func LicensableProductEntityPreSanitize(dto *LicensableProductEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
func LicensableProductEntityBeforeCreateAppend(dto *LicensableProductEntity, query workspaces.QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = workspaces.UUID()
	}
	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId
	LicensableProductRecursiveAddUniqueId(dto, query)
}
func LicensableProductRecursiveAddUniqueId(dto *LicensableProductEntity, query workspaces.QueryDSL) {
}
func LicensableProductActionBatchCreateFn(dtos []*LicensableProductEntity, query workspaces.QueryDSL) ([]*LicensableProductEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*LicensableProductEntity{}
		for _, item := range dtos {
			s, err := LicensableProductActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil
}
func LicensableProductActionCreateFn(dto *LicensableProductEntity, query workspaces.QueryDSL) (*LicensableProductEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := LicensableProductValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	LicensableProductEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	LicensableProductEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	LicensableProductPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	LicensableProductRelationContentCreate(dto, query)
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
	LicensableProductAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(LICENSABLEPRODUCT_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": workspaces.GetTypeString(&LicensableProductEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})
	return dto, nil
}
func LicensableProductActionGetOne(query workspaces.QueryDSL) (*LicensableProductEntity, *workspaces.IError) {
	refl := reflect.ValueOf(&LicensableProductEntity{})
	item, err := workspaces.GetOneEntity[LicensableProductEntity](query, refl)
	entityLicensableProductFormatter(item, query)
	return item, err
}
func LicensableProductActionQuery(query workspaces.QueryDSL) ([]*LicensableProductEntity, *workspaces.QueryResultMeta, error) {
	refl := reflect.ValueOf(&LicensableProductEntity{})
	items, meta, err := workspaces.QueryEntitiesPointer[LicensableProductEntity](query, refl)
	for _, item := range items {
		entityLicensableProductFormatter(item, query)
	}
	return items, meta, err
}
func LicensableProductUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *LicensableProductEntity) (*LicensableProductEntity, *workspaces.IError) {
	uniqueId := fields.UniqueId
	query.TriggerEventName = LICENSABLEPRODUCT_EVENT_UPDATED
	LicensableProductEntityPreSanitize(fields, query)
	var item LicensableProductEntity
	q := dbref.
		Where(&LicensableProductEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)
	err := q.UpdateColumns(fields).Error
	if err != nil {
		return nil, workspaces.GormErrorToIError(err)
	}
	query.Tx = dbref
	LicensableProductRelationContentUpdate(fields, query)
	LicensableProductPolyglotCreateHandler(fields, query)
	// @meta(update has many)
	err = dbref.
		Preload(clause.Associations).
		Where(&LicensableProductEntity{UniqueId: uniqueId}).
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
func LicensableProductActionUpdateFn(query workspaces.QueryDSL, fields *LicensableProductEntity) (*LicensableProductEntity, *workspaces.IError) {
	if fields == nil {
		return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}
	// 1. Validate always
	if iError := LicensableProductValidator(fields, true); iError != nil {
		return nil, iError
	}
	LicensableProductRecursiveAddUniqueId(fields, query)
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()
		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			_, err := LicensableProductUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}
		})
		return nil, workspaces.CastToIError(vf)
	} else {
		dbref = query.Tx
		return LicensableProductUpdateExec(dbref, query, fields)
	}
}

var LicensableProductWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire licensableproducts ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		count, _ := LicensableProductActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}

func LicensableProductActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&LicensableProductEntity{})
	query.ActionRequires = []string{PERM_ROOT_LICENSABLEPRODUCT_DELETE}
	return workspaces.RemoveEntity[LicensableProductEntity](query, refl)
}
func LicensableProductActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error
	var count int64 = 0
	{
		subCount, subErr := workspaces.WipeCleanEntity[LicensableProductEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'LicensableProductEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
func LicensableProductActionBulkUpdate(
	query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[LicensableProductEntity]) (
	*workspaces.BulkRecordRequest[LicensableProductEntity], *workspaces.IError,
) {
	result := []*LicensableProductEntity{}
	err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := LicensableProductActionUpdate(query, record)
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
func (x *LicensableProductEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

var LicensableProductEntityMeta = workspaces.TableMetaData{
	EntityName:    "LicensableProduct",
	ExportKey:     "licensable-products",
	TableNameInDb: "fb_licensableproduct_entities",
	EntityObject:  &LicensableProductEntity{},
	ExportStream:  LicensableProductActionExportT,
	ImportQuery:   LicensableProductActionImport,
}

func LicensableProductActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[LicensableProductEntity](query, LicensableProductActionQuery, LicensableProductPreloadRelations)
}
func LicensableProductActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[LicensableProductEntity](query, LicensableProductActionQuery, LicensableProductPreloadRelations)
}
func LicensableProductActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content LicensableProductEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := LicensableProductActionCreate(&content, query)
	return err
}

var LicensableProductCommonCliFlags = []cli.Flag{
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
		Name:     "private-key",
		Required: false,
		Usage:    "privateKey",
	},
	&cli.StringFlag{
		Name:     "public-key",
		Required: false,
		Usage:    "publicKey",
	},
}
var LicensableProductCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:        "name",
		StructField: "Name",
		Required:    true,
		Usage:       "name",
		Type:        "string",
	},
	{
		Name:        "privateKey",
		StructField: "PrivateKey",
		Required:    false,
		Usage:       "privateKey",
		Type:        "string",
	},
	{
		Name:        "publicKey",
		StructField: "PublicKey",
		Required:    false,
		Usage:       "publicKey",
		Type:        "string",
	},
}
var LicensableProductCommonCliFlagsOptional = []cli.Flag{
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
		Name:     "private-key",
		Required: false,
		Usage:    "privateKey",
	},
	&cli.StringFlag{
		Name:     "public-key",
		Required: false,
		Usage:    "publicKey",
	},
}
var LicensableProductCreateCmd cli.Command = cli.Command{
	Name:    "create",
	Aliases: []string{"c"},
	Flags:   LicensableProductCommonCliFlags,
	Usage:   "Create a new template",
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastLicensableProductFromCli(c)
		if entity, err := LicensableProductActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var LicensableProductCreateInteractiveCmd cli.Command = cli.Command{
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
		entity := &LicensableProductEntity{}
		for _, item := range LicensableProductCommonInteractiveCliFlags {
			if !item.Required && c.Bool("all") == false {
				continue
			}
			result := workspaces.AskForInput(item.Name, "")
			workspaces.SetFieldString(entity, item.StructField, result)
		}
		if entity, err := LicensableProductActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var LicensableProductUpdateCmd cli.Command = cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Flags:   LicensableProductCommonCliFlagsOptional,
	Usage:   "Updates a template by passing the parameters",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastLicensableProductFromCli(c)
		if entity, err := LicensableProductActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
		return nil
	},
}

func CastLicensableProductFromCli(c *cli.Context) *LicensableProductEntity {
	template := &LicensableProductEntity{}
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
	if c.IsSet("private-key") {
		value := c.String("private-key")
		template.PrivateKey = &value
	}
	if c.IsSet("public-key") {
		value := c.String("public-key")
		template.PublicKey = &value
	}
	return template
}
func LicensableProductSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		LicensableProductActionCreate,
		reflect.ValueOf(&LicensableProductEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}
func LicensableProductImportMocks() {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		LicensableProductActionCreate,
		reflect.ValueOf(&LicensableProductEntity{}).Elem(),
		&mocks.ViewsFs,
		[]string{},
		false,
	)
}
func LicensableProductWriteQueryMock(ctx workspaces.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := LicensableProductActionQuery(f)
		result := workspaces.QueryEntitySuccessResult(f, items, count)
		workspaces.WriteMockDataToFile(lang, "", "LicensableProduct", result)
	}
}

var LicensableProductImportExportCommands = []cli.Command{
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
			LicensableProductActionSeeder(query, c.Int("count"))
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
				Value: "licensable-product-seeder.yml",
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
			LicensableProductActionSeederInit(f, c.String("file"), c.String("format"))
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
				Value: "licensable-product-seeder-licensable-product.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of licensable-products, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]LicensableProductEntity{}
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
				LicensableProductActionCreate,
				reflect.ValueOf(&LicensableProductEntity{}).Elem(),
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
				LicensableProductActionCreate,
				reflect.ValueOf(&LicensableProductEntity{}).Elem(),
				c.String("file"),
			)
			return nil
		},
	},
}
var LicensableProductCliCommands []cli.Command = []cli.Command{
	workspaces.GetCommonQuery(LicensableProductActionQuery),
	workspaces.GetCommonTableQuery(reflect.ValueOf(&LicensableProductEntity{}).Elem(), LicensableProductActionQuery),
	LicensableProductCreateCmd,
	LicensableProductUpdateCmd,
	LicensableProductCreateInteractiveCmd,
	LicensableProductWipeCmd,
	workspaces.GetCommonRemoveQuery(reflect.ValueOf(&LicensableProductEntity{}).Elem(), LicensableProductActionRemove),
}

func LicensableProductCliFn() cli.Command {
	LicensableProductCliCommands = append(LicensableProductCliCommands, LicensableProductImportExportCommands...)
	return cli.Command{
		Name:        "product",
		Description: "LicensableProducts module actions (sample module to handle complex entities)",
		Usage:       "",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: LicensableProductCliCommands,
	}
}

/**
 *	Override this function on LicensableProductEntityHttp.go,
 *	In order to add your own http
 **/
var AppendLicensableProductRouter = func(r *[]workspaces.Module2Action) {}

func GetLicensableProductModule2Actions() []workspaces.Module2Action {
	routes := []workspaces.Module2Action{
		{
			Method:        "GET",
			Url:           "/licensable-products",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpQueryEntity(c, LicensableProductActionQuery)
				},
			},
			Format:         "QUERY",
			Action:         LicensableProductActionQuery,
			ResponseEntity: &[]LicensableProductEntity{},
		},
		{
			Method:        "GET",
			Url:           "/licensable-products/export",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpStreamFileChannel(c, LicensableProductActionExport)
				},
			},
			Format:         "QUERY",
			Action:         LicensableProductActionExport,
			ResponseEntity: &[]LicensableProductEntity{},
		},
		{
			Method:        "GET",
			Url:           "/licensable-product/:uniqueId",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpGetEntity(c, LicensableProductActionGetOne)
				},
			},
			Format:         "GET_ONE",
			Action:         LicensableProductActionGetOne,
			ResponseEntity: &LicensableProductEntity{},
		},
		{
			Method:        "POST",
			Url:           "/licensable-product",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpPostEntity(c, LicensableProductActionCreate)
				},
			},
			Action:         LicensableProductActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &LicensableProductEntity{},
			ResponseEntity: &LicensableProductEntity{},
		},
		{
			Method:        "PATCH",
			Url:           "/licensable-product",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpUpdateEntity(c, LicensableProductActionUpdate)
				},
			},
			Action:         LicensableProductActionUpdate,
			RequestEntity:  &LicensableProductEntity{},
			Format:         "PATCH_ONE",
			ResponseEntity: &LicensableProductEntity{},
		},
		{
			Method:        "PATCH",
			Url:           "/licensable-products",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpUpdateEntities(c, LicensableProductActionBulkUpdate)
				},
			},
			Action:         LicensableProductActionBulkUpdate,
			Format:         "PATCH_BULK",
			RequestEntity:  &workspaces.BulkRecordRequest[LicensableProductEntity]{},
			ResponseEntity: &workspaces.BulkRecordRequest[LicensableProductEntity]{},
		},
		{
			Method:        "DELETE",
			Url:           "/licensable-product",
			Format:        "DELETE_DSL",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpRemoveEntity(c, LicensableProductActionRemove)
				},
			},
			Action:         LicensableProductActionRemove,
			RequestEntity:  &workspaces.DeleteRequest{},
			ResponseEntity: &workspaces.DeleteResponse{},
			TargetEntity:   &LicensableProductEntity{},
		},
	}
	// Append user defined functions
	AppendLicensableProductRouter(&routes)
	return routes
}
func CreateLicensableProductRouter(r *gin.Engine) []workspaces.Module2Action {
	httpRoutes := GetLicensableProductModule2Actions()
	workspaces.CastRoutes(httpRoutes, r)
	workspaces.WriteHttpInformationToFile(&httpRoutes, LicensableProductEntityJsonSchema, "licensable-product-http", "licenses")
	workspaces.WriteEntitySchema("LicensableProductEntity", LicensableProductEntityJsonSchema, "licenses")
	return httpRoutes
}

var PERM_ROOT_LICENSABLEPRODUCT_DELETE = "root/licensableproduct/delete"
var PERM_ROOT_LICENSABLEPRODUCT_CREATE = "root/licensableproduct/create"
var PERM_ROOT_LICENSABLEPRODUCT_UPDATE = "root/licensableproduct/update"
var PERM_ROOT_LICENSABLEPRODUCT_QUERY = "root/licensableproduct/query"
var PERM_ROOT_LICENSABLEPRODUCT = "root/licensableproduct"
var ALL_LICENSABLEPRODUCT_PERMISSIONS = []string{
	PERM_ROOT_LICENSABLEPRODUCT_DELETE,
	PERM_ROOT_LICENSABLEPRODUCT_CREATE,
	PERM_ROOT_LICENSABLEPRODUCT_UPDATE,
	PERM_ROOT_LICENSABLEPRODUCT_QUERY,
	PERM_ROOT_LICENSABLEPRODUCT,
}
