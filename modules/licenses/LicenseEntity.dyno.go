package licenses

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	reflect "reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gookit/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/microcosm-cc/bluemonday"
	"github.com/schollz/progressbar/v3"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"pixelplux.com/fireback/modules/workspaces"
)

type LicensePermissions struct {
	Visibility       *string                      `json:"visibility,omitempty" yaml:"visibility"`
	WorkspaceId      *string                      `json:"workspaceId,omitempty" yaml:"workspaceId"`
	LinkerId         *string                      `json:"linkerId,omitempty" yaml:"linkerId"`
	ParentId         *string                      `json:"parentId,omitempty" yaml:"parentId"`
	UniqueId         string                       `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
	UserId           *string                      `json:"userId,omitempty" yaml:"userId"`
	Rank             int64                        `json:"rank,omitempty" gorm:"type:int;name:rank"`
	Updated          int64                        `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
	Created          int64                        `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
	CreatedFormatted string                       `json:"createdFormatted,omitempty" sql:"-"`
	UpdatedFormatted string                       `json:"updatedFormatted,omitempty" sql:"-"`
	Capability       *workspaces.CapabilityEntity `json:"capability" yaml:"capability"    gorm:"foreignKey:CapabilityId;references:UniqueId"     `
	// Datenano also has a text representation
	CapabilityId *string        `json:"capabilityId" yaml:"capabilityId"`
	LinkedTo     *LicenseEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

func (x *LicensePermissions) RootObjectName() string {
	return "LicenseEntity"
}

type LicenseEntity struct {
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
	Name             *string `json:"name" yaml:"name"       `
	// Datenano also has a text representation
	SignedLicense *string `json:"signedLicense" yaml:"signedLicense"       `
	// Datenano also has a text representation
	ValidityStartDate time.Time `json:"validityStartDate" yaml:"validityStartDate"       `
	// Datenano also has a text representation
	ValidityEndDate time.Time `json:"validityEndDate" yaml:"validityEndDate"       `
	// Datenano also has a text representation
	Permissions []*LicensePermissions `json:"permissions" yaml:"permissions"    gorm:"foreignKey:LinkerId;references:UniqueId"     `
	// Datenano also has a text representation
	Children []*LicenseEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
	LinkedTo *LicenseEntity   `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

var LicensePreloadRelations []string = []string{}
var LICENSE_EVENT_CREATED = "license.created"
var LICENSE_EVENT_UPDATED = "license.updated"
var LICENSE_EVENT_DELETED = "license.deleted"
var LICENSE_EVENTS = []string{
	LICENSE_EVENT_CREATED,
	LICENSE_EVENT_UPDATED,
	LICENSE_EVENT_DELETED,
}

type LicenseFieldMap struct {
	Name              workspaces.TranslatedString `yaml:"name"`
	SignedLicense     workspaces.TranslatedString `yaml:"signedLicense"`
	ValidityStartDate workspaces.TranslatedString `yaml:"validityStartDate"`
	ValidityEndDate   workspaces.TranslatedString `yaml:"validityEndDate"`
	Permissions       workspaces.TranslatedString `yaml:"permissions"`
}

var LicenseEntityMetaConfig map[string]int64 = map[string]int64{}
var LicenseEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&LicenseEntity{}))

func LicensePermissionsActionCreate(
	dto *LicensePermissions,
	query workspaces.QueryDSL,
) (*LicensePermissions, *workspaces.IError) {
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
func LicensePermissionsActionUpdate(
	query workspaces.QueryDSL,
	dto *LicensePermissions,
) (*LicensePermissions, *workspaces.IError) {
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
func LicensePermissionsActionGetOne(
	query workspaces.QueryDSL,
) (*LicensePermissions, *workspaces.IError) {
	refl := reflect.ValueOf(&LicensePermissions{})
	item, err := workspaces.GetOneEntity[LicensePermissions](query, refl)
	return item, err
}
func entityLicenseFormatter(dto *LicenseEntity, query workspaces.QueryDSL) {
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
func LicenseMockEntity() *LicenseEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &LicenseEntity{
		Name:          &stringHolder,
		SignedLicense: &stringHolder,
	}
	return entity
}
func LicenseActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := LicenseMockEntity()
		_, err := LicenseActionCreate(entity, query)
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
func LicenseActionSeederInit(query workspaces.QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*LicenseEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &LicenseEntity{
		Name:          &tildaRef,
		SignedLicense: &tildaRef,
		Permissions:   []*LicensePermissions{{}},
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
func LicenseAssociationCreate(dto *LicenseEntity, query workspaces.QueryDSL) error {
	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func LicenseRelationContentCreate(dto *LicenseEntity, query workspaces.QueryDSL) error {
	return nil
}
func LicenseRelationContentUpdate(dto *LicenseEntity, query workspaces.QueryDSL) error {
	return nil
}
func LicensePolyglotCreateHandler(dto *LicenseEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
}

/**
 * This will be validating your entity fully. Important note is that, you add validate:* tag
 * in your entity, it will automatically work here. For slices inside entity, make sure you add
 * extra line of AppendSliceErrors, otherwise they won't be detected
 */
func LicenseValidator(dto *LicenseEntity, isPatch bool) *workspaces.IError {
	err := workspaces.CommonStructValidatorPointer(dto, isPatch)
	if dto != nil && dto.Permissions != nil {
		workspaces.AppendSliceErrors(dto.Permissions, isPatch, "permissions", err)
	}
	return err
}
func LicenseEntityPreSanitize(dto *LicenseEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
func LicenseEntityBeforeCreateAppend(dto *LicenseEntity, query workspaces.QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = workspaces.UUID()
	}
	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId
	LicenseRecursiveAddUniqueId(dto, query)
}
func LicenseRecursiveAddUniqueId(dto *LicenseEntity, query workspaces.QueryDSL) {
	if dto.Permissions != nil && len(dto.Permissions) > 0 {
		for index0 := range dto.Permissions {
			if dto.Permissions[index0].UniqueId == "" {
				dto.Permissions[index0].UniqueId = workspaces.UUID()
			}
		}
	}
}
func LicenseActionBatchCreateFn(dtos []*LicenseEntity, query workspaces.QueryDSL) ([]*LicenseEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*LicenseEntity{}
		for _, item := range dtos {
			s, err := LicenseActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil
}
func LicenseActionCreateFn(dto *LicenseEntity, query workspaces.QueryDSL) (*LicenseEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := LicenseValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	LicenseEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	LicenseEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	LicensePolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	LicenseRelationContentCreate(dto, query)
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
	LicenseAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(LICENSE_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": workspaces.GetTypeString(&LicenseEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})
	return dto, nil
}
func LicenseActionGetOne(query workspaces.QueryDSL) (*LicenseEntity, *workspaces.IError) {
	refl := reflect.ValueOf(&LicenseEntity{})
	item, err := workspaces.GetOneEntity[LicenseEntity](query, refl)
	entityLicenseFormatter(item, query)
	return item, err
}
func LicenseActionQuery(query workspaces.QueryDSL) ([]*LicenseEntity, *workspaces.QueryResultMeta, error) {
	refl := reflect.ValueOf(&LicenseEntity{})
	items, meta, err := workspaces.QueryEntitiesPointer[LicenseEntity](query, refl)
	for _, item := range items {
		entityLicenseFormatter(item, query)
	}
	return items, meta, err
}
func LicenseUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *LicenseEntity) (*LicenseEntity, *workspaces.IError) {
	uniqueId := fields.UniqueId
	query.TriggerEventName = LICENSE_EVENT_UPDATED
	LicenseEntityPreSanitize(fields, query)
	var item LicenseEntity
	q := dbref.
		Where(&LicenseEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)
	err := q.UpdateColumns(fields).Error
	if err != nil {
		return nil, workspaces.GormErrorToIError(err)
	}
	query.Tx = dbref
	LicenseRelationContentUpdate(fields, query)
	LicensePolyglotCreateHandler(fields, query)
	// @meta(update has many)
	if fields.Permissions != nil {
		linkerId := uniqueId
		dbref.Debug().
			Where(&LicensePermissions{LinkerId: &linkerId}).
			Delete(&LicensePermissions{})
		for _, newItem := range fields.Permissions {
			newItem.UniqueId = workspaces.UUID()
			newItem.LinkerId = &linkerId
			dbref.Create(&newItem)
		}
	}
	err = dbref.
		Preload(clause.Associations).
		Where(&LicenseEntity{UniqueId: uniqueId}).
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
func LicenseActionUpdateFn(query workspaces.QueryDSL, fields *LicenseEntity) (*LicenseEntity, *workspaces.IError) {
	if fields == nil {
		return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}
	// 1. Validate always
	if iError := LicenseValidator(fields, true); iError != nil {
		return nil, iError
	}
	LicenseRecursiveAddUniqueId(fields, query)
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()
		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			_, err := LicenseUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}
		})
		return nil, workspaces.CastToIError(vf)
	} else {
		dbref = query.Tx
		return LicenseUpdateExec(dbref, query, fields)
	}
}

var LicenseWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire licenses ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		count, _ := LicenseActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}

func LicenseActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&LicenseEntity{})
	query.ActionRequires = []string{PERM_ROOT_LICENSE_DELETE}
	return workspaces.RemoveEntity[LicenseEntity](query, refl)
}
func LicenseActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error
	var count int64 = 0
	{
		subCount, subErr := workspaces.WipeCleanEntity[LicensePermissions]()
		if subErr != nil {
			fmt.Println("Error while wiping 'LicensePermissions'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	{
		subCount, subErr := workspaces.WipeCleanEntity[LicenseEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'LicenseEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
func LicenseActionBulkUpdate(
	query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[LicenseEntity]) (
	*workspaces.BulkRecordRequest[LicenseEntity], *workspaces.IError,
) {
	result := []*LicenseEntity{}
	err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := LicenseActionUpdate(query, record)
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
func (x *LicenseEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

var LicenseEntityMeta = workspaces.TableMetaData{
	EntityName:    "License",
	ExportKey:     "licenses",
	TableNameInDb: "fb_license_entities",
	EntityObject:  &LicenseEntity{},
	ExportStream:  LicenseActionExportT,
	ImportQuery:   LicenseActionImport,
}

func LicenseActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[LicenseEntity](query, LicenseActionQuery, LicensePreloadRelations)
}
func LicenseActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[LicenseEntity](query, LicenseActionQuery, LicensePreloadRelations)
}
func LicenseActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content LicenseEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := LicenseActionCreate(&content, query)
	return err
}

var LicenseCommonCliFlags = []cli.Flag{
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
		Name:     "signed-license",
		Required: false,
		Usage:    "signedLicense",
	},
	&cli.StringSliceFlag{
		Name:     "permissions",
		Required: false,
		Usage:    "permissions",
	},
}
var LicenseCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:        "name",
		StructField: "Name",
		Required:    false,
		Usage:       "name",
		Type:        "string",
	},
	{
		Name:        "signedLicense",
		StructField: "SignedLicense",
		Required:    false,
		Usage:       "signedLicense",
		Type:        "string",
	},
}
var LicenseCommonCliFlagsOptional = []cli.Flag{
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
		Name:     "signed-license",
		Required: false,
		Usage:    "signedLicense",
	},
	&cli.StringSliceFlag{
		Name:     "permissions",
		Required: false,
		Usage:    "permissions",
	},
}
var LicenseCreateCmd cli.Command = cli.Command{
	Name:    "create",
	Aliases: []string{"c"},
	Flags:   LicenseCommonCliFlags,
	Usage:   "Create a new template",
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastLicenseFromCli(c)
		if entity, err := LicenseActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var LicenseCreateInteractiveCmd cli.Command = cli.Command{
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
		entity := &LicenseEntity{}
		for _, item := range LicenseCommonInteractiveCliFlags {
			if !item.Required && c.Bool("all") == false {
				continue
			}
			result := workspaces.AskForInput(item.Name, "")
			workspaces.SetFieldString(entity, item.StructField, result)
		}
		if entity, err := LicenseActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var LicenseUpdateCmd cli.Command = cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Flags:   LicenseCommonCliFlagsOptional,
	Usage:   "Updates a template by passing the parameters",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastLicenseFromCli(c)
		if entity, err := LicenseActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
		return nil
	},
}

func CastLicenseFromCli(c *cli.Context) *LicenseEntity {
	template := &LicenseEntity{}
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
	if c.IsSet("signed-license") {
		value := c.String("signed-license")
		template.SignedLicense = &value
	}
	return template
}
func LicenseSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		LicenseActionCreate,
		reflect.ValueOf(&LicenseEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}
func LicenseWriteQueryMock(ctx workspaces.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := LicenseActionQuery(f)
		result := workspaces.QueryEntitySuccessResult(f, items, count)
		workspaces.WriteMockDataToFile(lang, "", "License", result)
	}
}

var LicenseImportExportCommands = []cli.Command{
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
			LicenseActionSeeder(query, c.Int("count"))
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
				Value: "license-seeder.yml",
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
			LicenseActionSeederInit(f, c.String("file"), c.String("format"))
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
				Value: "license-seeder-license.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of licenses, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]LicenseEntity{}
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
				LicenseActionCreate,
				reflect.ValueOf(&LicenseEntity{}).Elem(),
				c.String("file"),
			)
			return nil
		},
	},
}
var LicenseCliCommands []cli.Command = []cli.Command{
	workspaces.GetCommonQuery(LicenseActionQuery),
	workspaces.GetCommonTableQuery(reflect.ValueOf(&LicenseEntity{}).Elem(), LicenseActionQuery),
	LicenseCreateCmd,
	LicenseUpdateCmd,
	LicenseCreateInteractiveCmd,
	LicenseWipeCmd,
	workspaces.GetCommonRemoveQuery(reflect.ValueOf(&LicenseEntity{}).Elem(), LicenseActionRemove),
}

func LicenseCliFn() cli.Command {
	LicenseCliCommands = append(LicenseCliCommands, LicenseImportExportCommands...)
	return cli.Command{
		Name:        "license",
		Description: "Licenses module actions (sample module to handle complex entities)",
		Usage:       "Manage the licenses in the app (either to issue, or to activate current product)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: LicenseCliCommands,
	}
}

/**
 *	Override this function on LicenseEntityHttp.go,
 *	In order to add your own http
 **/
var AppendLicenseRouter = func(r *[]workspaces.Module2Action) {}

func GetLicenseModule2Actions() []workspaces.Module2Action {
	routes := []workspaces.Module2Action{
		{
			Method: "GET",
			Url:    "/licenses",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_LICENSE_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpQueryEntity(c, LicenseActionQuery)
				},
			},
			Format:         "QUERY",
			Action:         LicenseActionQuery,
			ResponseEntity: &[]LicenseEntity{},
		},
		{
			Method: "GET",
			Url:    "/licenses/export",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_LICENSE_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpStreamFileChannel(c, LicenseActionExport)
				},
			},
			Format:         "QUERY",
			Action:         LicenseActionExport,
			ResponseEntity: &[]LicenseEntity{},
		},
		{
			Method: "GET",
			Url:    "/license/:uniqueId",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_LICENSE_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpGetEntity(c, LicenseActionGetOne)
				},
			},
			Format:         "GET_ONE",
			Action:         LicenseActionGetOne,
			ResponseEntity: &LicenseEntity{},
		},
		{
			Method: "POST",
			Url:    "/license",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_LICENSE_CREATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpPostEntity(c, LicenseActionCreate)
				},
			},
			Action:         LicenseActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &LicenseEntity{},
			ResponseEntity: &LicenseEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/license",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_LICENSE_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpUpdateEntity(c, LicenseActionUpdate)
				},
			},
			Action:         LicenseActionUpdate,
			RequestEntity:  &LicenseEntity{},
			Format:         "PATCH_ONE",
			ResponseEntity: &LicenseEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/licenses",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_LICENSE_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpUpdateEntities(c, LicenseActionBulkUpdate)
				},
			},
			Action:         LicenseActionBulkUpdate,
			Format:         "PATCH_BULK",
			RequestEntity:  &workspaces.BulkRecordRequest[LicenseEntity]{},
			ResponseEntity: &workspaces.BulkRecordRequest[LicenseEntity]{},
		},
		{
			Method: "DELETE",
			Url:    "/license",
			Format: "DELETE_DSL",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_LICENSE_DELETE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpRemoveEntity(c, LicenseActionRemove)
				},
			},
			Action:         LicenseActionRemove,
			RequestEntity:  &workspaces.DeleteRequest{},
			ResponseEntity: &workspaces.DeleteResponse{},
			TargetEntity:   &LicenseEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/license/:linkerId/permissions/:uniqueId",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_LICENSE_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(
					c *gin.Context,
				) {
					workspaces.HttpUpdateEntity(c, LicensePermissionsActionUpdate)
				},
			},
			Action:         LicensePermissionsActionUpdate,
			Format:         "PATCH_ONE",
			RequestEntity:  &LicensePermissions{},
			ResponseEntity: &LicensePermissions{},
		},
		{
			Method: "GET",
			Url:    "/license/permissions/:linkerId/:uniqueId",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_LICENSE_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(
					c *gin.Context,
				) {
					workspaces.HttpGetEntity(c, LicensePermissionsActionGetOne)
				},
			},
			Action:         LicensePermissionsActionGetOne,
			Format:         "GET_ONE",
			ResponseEntity: &LicensePermissions{},
		},
		{
			Method: "POST",
			Url:    "/license/:linkerId/permissions",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_LICENSE_CREATE},
			},
			Handlers: []gin.HandlerFunc{
				func(
					c *gin.Context,
				) {
					workspaces.HttpPostEntity(c, LicensePermissionsActionCreate)
				},
			},
			Action:         LicensePermissionsActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &LicensePermissions{},
			ResponseEntity: &LicensePermissions{},
		},
	}
	// Append user defined functions
	AppendLicenseRouter(&routes)
	return routes
}
func CreateLicenseRouter(r *gin.Engine) []workspaces.Module2Action {
	httpRoutes := GetLicenseModule2Actions()
	workspaces.CastRoutes(httpRoutes, r)
	workspaces.WriteHttpInformationToFile(&httpRoutes, LicenseEntityJsonSchema, "license-http", "licenses")
	workspaces.WriteEntitySchema("LicenseEntity", LicenseEntityJsonSchema, "licenses")
	return httpRoutes
}

var PERM_ROOT_LICENSE_DELETE = "root/license/delete"
var PERM_ROOT_LICENSE_CREATE = "root/license/create"
var PERM_ROOT_LICENSE_UPDATE = "root/license/update"
var PERM_ROOT_LICENSE_QUERY = "root/license/query"
var PERM_ROOT_LICENSE = "root/license"
var ALL_LICENSE_PERMISSIONS = []string{
	PERM_ROOT_LICENSE_DELETE,
	PERM_ROOT_LICENSE_CREATE,
	PERM_ROOT_LICENSE_UPDATE,
	PERM_ROOT_LICENSE_QUERY,
	PERM_ROOT_LICENSE,
}
