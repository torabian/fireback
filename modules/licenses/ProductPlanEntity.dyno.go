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
	"github.com/torabian/fireback/modules/currency"
	mocks "github.com/torabian/fireback/modules/licenses/mocks/ProductPlan"
	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductPlanPermissions struct {
	Visibility       *string                      `json:"visibility,omitempty" yaml:"visibility"`
	WorkspaceId      *string                      `json:"workspaceId,omitempty" yaml:"workspaceId"`
	LinkerId         *string                      `json:"linkerId,omitempty" yaml:"linkerId"`
	ParentId         *string                      `json:"parentId,omitempty" yaml:"parentId"`
	UniqueId         string                       `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
	UserId           *string                      `json:"userId,omitempty" yaml:"userId"`
	Rank             int64                        `json:"rank,omitempty" gorm:"type:int;name:rank"`
	Updated          int64                        `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
	Created          int64                        `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
	CreatedFormatted string                       `json:"createdFormatted,omitempty" sql:"-" gorm:"-"`
	UpdatedFormatted string                       `json:"updatedFormatted,omitempty" sql:"-" gorm:"-"`
	Capability       *workspaces.CapabilityEntity `json:"capability" yaml:"capability"    gorm:"foreignKey:CapabilityId;references:UniqueId"     `
	// Datenano also has a text representation
	CapabilityId *string            `json:"capabilityId" yaml:"capabilityId"`
	LinkedTo     *ProductPlanEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

func (x *ProductPlanPermissions) RootObjectName() string {
	return "ProductPlanEntity"
}

type ProductPlanEntity struct {
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
	Duration *int64 `json:"duration" yaml:"duration"  validate:"required"       `
	// Datenano also has a text representation
	Product *LicensableProductEntity `json:"product" yaml:"product"    gorm:"foreignKey:ProductId;references:UniqueId"     `
	// Datenano also has a text representation
	ProductId *string                  `json:"productId" yaml:"productId" validate:"required" `
	PriceTag  *currency.PriceTagEntity `json:"priceTag" yaml:"priceTag"    gorm:"foreignKey:PriceTagId;references:UniqueId"     `
	// Datenano also has a text representation
	PriceTagId  *string                   `json:"priceTagId" yaml:"priceTagId"`
	Permissions []*ProductPlanPermissions `json:"permissions" yaml:"permissions"    gorm:"foreignKey:LinkerId;references:UniqueId"     `
	// Datenano also has a text representation
	Translations []*ProductPlanEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
	Children     []*ProductPlanEntity         `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
	LinkedTo     *ProductPlanEntity           `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

var ProductPlanPreloadRelations []string = []string{}
var PRODUCT_PLAN_EVENT_CREATED = "productPlan.created"
var PRODUCT_PLAN_EVENT_UPDATED = "productPlan.updated"
var PRODUCT_PLAN_EVENT_DELETED = "productPlan.deleted"
var PRODUCT_PLAN_EVENTS = []string{
	PRODUCT_PLAN_EVENT_CREATED,
	PRODUCT_PLAN_EVENT_UPDATED,
	PRODUCT_PLAN_EVENT_DELETED,
}

type ProductPlanFieldMap struct {
	Name        workspaces.TranslatedString `yaml:"name"`
	Duration    workspaces.TranslatedString `yaml:"duration"`
	Product     workspaces.TranslatedString `yaml:"product"`
	PriceTag    workspaces.TranslatedString `yaml:"priceTag"`
	Permissions workspaces.TranslatedString `yaml:"permissions"`
}

var ProductPlanEntityMetaConfig map[string]int64 = map[string]int64{}
var ProductPlanEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&ProductPlanEntity{}))

type ProductPlanEntityPolyglot struct {
	LinkerId   string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
	LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
	Name       string `yaml:"name" json:"name"`
}

func ProductPlanPermissionsActionCreate(
	dto *ProductPlanPermissions,
	query workspaces.QueryDSL,
) (*ProductPlanPermissions, *workspaces.IError) {
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
func ProductPlanPermissionsActionUpdate(
	query workspaces.QueryDSL,
	dto *ProductPlanPermissions,
) (*ProductPlanPermissions, *workspaces.IError) {
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
func ProductPlanPermissionsActionGetOne(
	query workspaces.QueryDSL,
) (*ProductPlanPermissions, *workspaces.IError) {
	refl := reflect.ValueOf(&ProductPlanPermissions{})
	item, err := workspaces.GetOneEntity[ProductPlanPermissions](query, refl)
	return item, err
}
func entityProductPlanFormatter(dto *ProductPlanEntity, query workspaces.QueryDSL) {
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
func ProductPlanMockEntity() *ProductPlanEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &ProductPlanEntity{
		Name:     &stringHolder,
		Duration: &int64Holder,
	}
	return entity
}
func ProductPlanActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := ProductPlanMockEntity()
		_, err := ProductPlanActionCreate(entity, query)
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
func (x *ProductPlanEntity) GetNameTranslated(language string) string {
	if x.Translations != nil && len(x.Translations) > 0 {
		for _, item := range x.Translations {
			if item.LanguageId == language {
				return item.Name
			}
		}
	}
	return ""
}
func ProductPlanActionSeederInit(query workspaces.QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*ProductPlanEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &ProductPlanEntity{
		Name:        &tildaRef,
		Permissions: []*ProductPlanPermissions{{}},
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
func ProductPlanAssociationCreate(dto *ProductPlanEntity, query workspaces.QueryDSL) error {
	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func ProductPlanRelationContentCreate(dto *ProductPlanEntity, query workspaces.QueryDSL) error {
	return nil
}
func ProductPlanRelationContentUpdate(dto *ProductPlanEntity, query workspaces.QueryDSL) error {
	return nil
}
func ProductPlanPolyglotCreateHandler(dto *ProductPlanEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
	workspaces.PolyglotCreateHandler(dto, &ProductPlanEntityPolyglot{}, query)
}

/**
 * This will be validating your entity fully. Important note is that, you add validate:* tag
 * in your entity, it will automatically work here. For slices inside entity, make sure you add
 * extra line of AppendSliceErrors, otherwise they won't be detected
 */
func ProductPlanValidator(dto *ProductPlanEntity, isPatch bool) *workspaces.IError {
	err := workspaces.CommonStructValidatorPointer(dto, isPatch)
	if dto != nil && dto.Permissions != nil {
		workspaces.AppendSliceErrors(dto.Permissions, isPatch, "permissions", err)
	}
	return err
}
func ProductPlanEntityPreSanitize(dto *ProductPlanEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
func ProductPlanEntityBeforeCreateAppend(dto *ProductPlanEntity, query workspaces.QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = workspaces.UUID()
	}
	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId
	ProductPlanRecursiveAddUniqueId(dto, query)
}
func ProductPlanRecursiveAddUniqueId(dto *ProductPlanEntity, query workspaces.QueryDSL) {
	if dto.Permissions != nil && len(dto.Permissions) > 0 {
		for index0 := range dto.Permissions {
			if dto.Permissions[index0].UniqueId == "" {
				dto.Permissions[index0].UniqueId = workspaces.UUID()
			}
		}
	}
}
func ProductPlanActionBatchCreateFn(dtos []*ProductPlanEntity, query workspaces.QueryDSL) ([]*ProductPlanEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*ProductPlanEntity{}
		for _, item := range dtos {
			s, err := ProductPlanActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil
}
func ProductPlanDeleteEntireChildren(query workspaces.QueryDSL, dto *ProductPlanEntity) *workspaces.IError {
	if dto.Permissions != nil {
		q := query.Tx.
			Model(&dto.Permissions).
			Where(&ProductPlanPermissions{LinkerId: &dto.UniqueId}).
			Delete(&ProductPlanPermissions{})
		err := q.Error
		if err != nil {
			return workspaces.GormErrorToIError(err)
		}
	}
	return nil
}
func ProductPlanActionCreateFn(dto *ProductPlanEntity, query workspaces.QueryDSL) (*ProductPlanEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := ProductPlanValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	ProductPlanEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	ProductPlanEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	ProductPlanPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	ProductPlanRelationContentCreate(dto, query)
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
	ProductPlanAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(PRODUCT_PLAN_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": workspaces.GetTypeString(&ProductPlanEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})
	return dto, nil
}
func ProductPlanActionGetOne(query workspaces.QueryDSL) (*ProductPlanEntity, *workspaces.IError) {
	refl := reflect.ValueOf(&ProductPlanEntity{})
	item, err := workspaces.GetOneEntity[ProductPlanEntity](query, refl)
	entityProductPlanFormatter(item, query)
	return item, err
}
func ProductPlanActionQuery(query workspaces.QueryDSL) ([]*ProductPlanEntity, *workspaces.QueryResultMeta, error) {
	refl := reflect.ValueOf(&ProductPlanEntity{})
	items, meta, err := workspaces.QueryEntitiesPointer[ProductPlanEntity](query, refl)
	for _, item := range items {
		entityProductPlanFormatter(item, query)
	}
	return items, meta, err
}
func ProductPlanUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *ProductPlanEntity) (*ProductPlanEntity, *workspaces.IError) {
	uniqueId := fields.UniqueId
	query.TriggerEventName = PRODUCT_PLAN_EVENT_UPDATED
	ProductPlanEntityPreSanitize(fields, query)
	var item ProductPlanEntity
	q := dbref.
		Where(&ProductPlanEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)
	err := q.UpdateColumns(fields).Error
	if err != nil {
		return nil, workspaces.GormErrorToIError(err)
	}
	query.Tx = dbref
	ProductPlanRelationContentUpdate(fields, query)
	ProductPlanPolyglotCreateHandler(fields, query)
	if ero := ProductPlanDeleteEntireChildren(query, fields); ero != nil {
		return nil, ero
	}
	// @meta(update has many)
	if fields.Permissions != nil {
		linkerId := uniqueId
		dbref.
			Where(&ProductPlanPermissions{LinkerId: &linkerId}).
			Delete(&ProductPlanPermissions{})
		for _, newItem := range fields.Permissions {
			newItem.UniqueId = workspaces.UUID()
			newItem.LinkerId = &linkerId
			dbref.Create(&newItem)
		}
	}
	err = dbref.
		Preload(clause.Associations).
		Where(&ProductPlanEntity{UniqueId: uniqueId}).
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
func ProductPlanActionUpdateFn(query workspaces.QueryDSL, fields *ProductPlanEntity) (*ProductPlanEntity, *workspaces.IError) {
	if fields == nil {
		return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}
	// 1. Validate always
	if iError := ProductPlanValidator(fields, true); iError != nil {
		return nil, iError
	}
	// Let's not add this. I am not sure of the consequences
	// ProductPlanRecursiveAddUniqueId(fields, query)
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()
		var item *ProductPlanEntity
		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			var err *workspaces.IError
			item, err = ProductPlanUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}
		})
		return item, workspaces.CastToIError(vf)
	} else {
		dbref = query.Tx
		return ProductPlanUpdateExec(dbref, query, fields)
	}
}

var ProductPlanWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire productplans ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
			ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_PLAN_DELETE},
		})
		count, _ := ProductPlanActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}

func ProductPlanActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&ProductPlanEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_PLAN_DELETE}
	return workspaces.RemoveEntity[ProductPlanEntity](query, refl)
}
func ProductPlanActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error
	var count int64 = 0
	{
		subCount, subErr := workspaces.WipeCleanEntity[ProductPlanPermissions]()
		if subErr != nil {
			fmt.Println("Error while wiping 'ProductPlanPermissions'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	{
		subCount, subErr := workspaces.WipeCleanEntity[ProductPlanEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'ProductPlanEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
func ProductPlanActionBulkUpdate(
	query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[ProductPlanEntity]) (
	*workspaces.BulkRecordRequest[ProductPlanEntity], *workspaces.IError,
) {
	result := []*ProductPlanEntity{}
	err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := ProductPlanActionUpdate(query, record)
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
func (x *ProductPlanEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

var ProductPlanEntityMeta = workspaces.TableMetaData{
	EntityName:    "ProductPlan",
	ExportKey:     "product-plans",
	TableNameInDb: "fb_product-plan_entities",
	EntityObject:  &ProductPlanEntity{},
	ExportStream:  ProductPlanActionExportT,
	ImportQuery:   ProductPlanActionImport,
}

func ProductPlanActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[ProductPlanEntity](query, ProductPlanActionQuery, ProductPlanPreloadRelations)
}
func ProductPlanActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[ProductPlanEntity](query, ProductPlanActionQuery, ProductPlanPreloadRelations)
}
func ProductPlanActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content ProductPlanEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := ProductPlanActionCreate(&content, query)
	return err
}

var ProductPlanCommonCliFlags = []cli.Flag{
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
	&cli.Int64Flag{
		Name:     "duration",
		Required: true,
		Usage:    "duration",
	},
	&cli.StringFlag{
		Name:     "product-id",
		Required: true,
		Usage:    "product",
	},
	&cli.StringFlag{
		Name:     "price-tag-id",
		Required: false,
		Usage:    "priceTag",
	},
	&cli.StringSliceFlag{
		Name:     "permissions",
		Required: false,
		Usage:    "permissions",
	},
}
var ProductPlanCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:        "name",
		StructField: "Name",
		Required:    true,
		Usage:       "name",
		Type:        "string",
	},
	{
		Name:        "duration",
		StructField: "Duration",
		Required:    true,
		Usage:       "duration",
		Type:        "int64",
	},
}
var ProductPlanCommonCliFlagsOptional = []cli.Flag{
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
	&cli.Int64Flag{
		Name:     "duration",
		Required: true,
		Usage:    "duration",
	},
	&cli.StringFlag{
		Name:     "product-id",
		Required: true,
		Usage:    "product",
	},
	&cli.StringFlag{
		Name:     "price-tag-id",
		Required: false,
		Usage:    "priceTag",
	},
	&cli.StringSliceFlag{
		Name:     "permissions",
		Required: false,
		Usage:    "permissions",
	},
}
var ProductPlanCreateCmd cli.Command = PRODUCT_PLAN_ACTION_POST_ONE.ToCli()
var ProductPlanCreateInteractiveCmd cli.Command = cli.Command{
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
			ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_PLAN_CREATE},
		})
		entity := &ProductPlanEntity{}
		for _, item := range ProductPlanCommonInteractiveCliFlags {
			if !item.Required && c.Bool("all") == false {
				continue
			}
			result := workspaces.AskForInput(item.Name, "")
			workspaces.SetFieldString(entity, item.StructField, result)
		}
		if entity, err := ProductPlanActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var ProductPlanUpdateCmd cli.Command = cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Flags:   ProductPlanCommonCliFlagsOptional,
	Usage:   "Updates a template by passing the parameters",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
			ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_PLAN_UPDATE},
		})
		entity := CastProductPlanFromCli(c)
		if entity, err := ProductPlanActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
		return nil
	},
}

func (x *ProductPlanEntity) FromCli(c *cli.Context) *ProductPlanEntity {
	return CastProductPlanFromCli(c)
}
func CastProductPlanFromCli(c *cli.Context) *ProductPlanEntity {
	template := &ProductPlanEntity{}
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
	if c.IsSet("product-id") {
		value := c.String("product-id")
		template.ProductId = &value
	}
	if c.IsSet("price-tag-id") {
		value := c.String("price-tag-id")
		template.PriceTagId = &value
	}
	return template
}
func ProductPlanSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		ProductPlanActionCreate,
		reflect.ValueOf(&ProductPlanEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}
func ProductPlanImportMocks() {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		ProductPlanActionCreate,
		reflect.ValueOf(&ProductPlanEntity{}).Elem(),
		&mocks.ViewsFs,
		[]string{},
		false,
	)
}
func ProductPlanWriteQueryMock(ctx workspaces.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := ProductPlanActionQuery(f)
		result := workspaces.QueryEntitySuccessResult(f, items, count)
		workspaces.WriteMockDataToFile(lang, "", "ProductPlan", result)
	}
}

var ProductPlanImportExportCommands = []cli.Command{
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
				ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_PLAN_CREATE},
			})
			ProductPlanActionSeeder(query, c.Int("count"))
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
				Value: "product-plan-seeder.yml",
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
			query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
				ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_PLAN_CREATE},
			})
			ProductPlanActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "product-plan-seeder-product-plan.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of product-plans, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]ProductPlanEntity{}
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
				ProductPlanActionCreate,
				reflect.ValueOf(&ProductPlanEntity{}).Elem(),
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
			ProductPlanCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				ProductPlanActionCreate,
				reflect.ValueOf(&ProductPlanEntity{}).Elem(),
				c.String("file"),
				&workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_PLAN_CREATE},
				},
				func() ProductPlanEntity {
					v := CastProductPlanFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
var ProductPlanCliCommands []cli.Command = []cli.Command{
	workspaces.GetCommonQuery2(ProductPlanActionQuery, &workspaces.SecurityModel{
		ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_PLAN_CREATE},
	}),
	workspaces.GetCommonTableQuery(reflect.ValueOf(&ProductPlanEntity{}).Elem(), ProductPlanActionQuery),
	ProductPlanCreateCmd,
	ProductPlanUpdateCmd,
	ProductPlanCreateInteractiveCmd,
	ProductPlanWipeCmd,
	workspaces.GetCommonRemoveQuery(reflect.ValueOf(&ProductPlanEntity{}).Elem(), ProductPlanActionRemove),
}

func ProductPlanCliFn() cli.Command {
	ProductPlanCliCommands = append(ProductPlanCliCommands, ProductPlanImportExportCommands...)
	return cli.Command{
		Name:        "plan",
		Description: "ProductPlans module actions (sample module to handle complex entities)",
		Usage:       "",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: ProductPlanCliCommands,
	}
}

var PRODUCT_PLAN_ACTION_POST_ONE = workspaces.Module2Action{
	ActionName:    "create",
	ActionAliases: []string{"c"},
	Description:   "Create new productPlan",
	Flags:         ProductPlanCommonCliFlags,
	Method:        "POST",
	Url:           "/product-plan",
	SecurityModel: &workspaces.SecurityModel{},
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			workspaces.HttpPostEntity(c, ProductPlanActionCreate)
		},
	},
	CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		result, err := workspaces.CliPostEntity(c, ProductPlanActionCreate, security)
		workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
		return err
	},
	Action:         ProductPlanActionCreate,
	Format:         "POST_ONE",
	RequestEntity:  &ProductPlanEntity{},
	ResponseEntity: &ProductPlanEntity{},
}

/**
 *	Override this function on ProductPlanEntityHttp.go,
 *	In order to add your own http
 **/
var AppendProductPlanRouter = func(r *[]workspaces.Module2Action) {}

func GetProductPlanModule2Actions() []workspaces.Module2Action {
	routes := []workspaces.Module2Action{
		{
			Method:        "GET",
			Url:           "/product-plans",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpQueryEntity(c, ProductPlanActionQuery)
				},
			},
			Format:         "QUERY",
			Action:         ProductPlanActionQuery,
			ResponseEntity: &[]ProductPlanEntity{},
		},
		{
			Method:        "GET",
			Url:           "/product-plans/export",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpStreamFileChannel(c, ProductPlanActionExport)
				},
			},
			Format:         "QUERY",
			Action:         ProductPlanActionExport,
			ResponseEntity: &[]ProductPlanEntity{},
		},
		{
			Method:        "GET",
			Url:           "/product-plan/:uniqueId",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpGetEntity(c, ProductPlanActionGetOne)
				},
			},
			Format:         "GET_ONE",
			Action:         ProductPlanActionGetOne,
			ResponseEntity: &ProductPlanEntity{},
		},
		PRODUCT_PLAN_ACTION_POST_ONE,
		{
			ActionName:    "update",
			ActionAliases: []string{"u"},
			Flags:         ProductPlanCommonCliFlagsOptional,
			Method:        "PATCH",
			Url:           "/product-plan",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpUpdateEntity(c, ProductPlanActionUpdate)
				},
			},
			Action:         ProductPlanActionUpdate,
			RequestEntity:  &ProductPlanEntity{},
			Format:         "PATCH_ONE",
			ResponseEntity: &ProductPlanEntity{},
		},
		{
			Method:        "PATCH",
			Url:           "/product-plans",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpUpdateEntities(c, ProductPlanActionBulkUpdate)
				},
			},
			Action:         ProductPlanActionBulkUpdate,
			Format:         "PATCH_BULK",
			RequestEntity:  &workspaces.BulkRecordRequest[ProductPlanEntity]{},
			ResponseEntity: &workspaces.BulkRecordRequest[ProductPlanEntity]{},
		},
		{
			Method:        "DELETE",
			Url:           "/product-plan",
			Format:        "DELETE_DSL",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpRemoveEntity(c, ProductPlanActionRemove)
				},
			},
			Action:         ProductPlanActionRemove,
			RequestEntity:  &workspaces.DeleteRequest{},
			ResponseEntity: &workspaces.DeleteResponse{},
			TargetEntity:   &ProductPlanEntity{},
		},
		{
			Method:        "PATCH",
			Url:           "/product-plan/:linkerId/permissions/:uniqueId",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(
					c *gin.Context,
				) {
					workspaces.HttpUpdateEntity(c, ProductPlanPermissionsActionUpdate)
				},
			},
			Action:         ProductPlanPermissionsActionUpdate,
			Format:         "PATCH_ONE",
			RequestEntity:  &ProductPlanPermissions{},
			ResponseEntity: &ProductPlanPermissions{},
		},
		{
			Method:        "GET",
			Url:           "/product-plan/permissions/:linkerId/:uniqueId",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(
					c *gin.Context,
				) {
					workspaces.HttpGetEntity(c, ProductPlanPermissionsActionGetOne)
				},
			},
			Action:         ProductPlanPermissionsActionGetOne,
			Format:         "GET_ONE",
			ResponseEntity: &ProductPlanPermissions{},
		},
		{
			Method:        "POST",
			Url:           "/product-plan/:linkerId/permissions",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(
					c *gin.Context,
				) {
					workspaces.HttpPostEntity(c, ProductPlanPermissionsActionCreate)
				},
			},
			Action:         ProductPlanPermissionsActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &ProductPlanPermissions{},
			ResponseEntity: &ProductPlanPermissions{},
		},
	}
	// Append user defined functions
	AppendProductPlanRouter(&routes)
	return routes
}
func CreateProductPlanRouter(r *gin.Engine) []workspaces.Module2Action {
	httpRoutes := GetProductPlanModule2Actions()
	workspaces.CastRoutes(httpRoutes, r)
	workspaces.WriteHttpInformationToFile(&httpRoutes, ProductPlanEntityJsonSchema, "product-plan-http", "licenses")
	workspaces.WriteEntitySchema("ProductPlanEntity", ProductPlanEntityJsonSchema, "licenses")
	return httpRoutes
}

var PERM_ROOT_PRODUCT_PLAN_DELETE = workspaces.PermissionInfo{
	CompleteKey: "root/licenses/product-plan/delete",
}
var PERM_ROOT_PRODUCT_PLAN_CREATE = workspaces.PermissionInfo{
	CompleteKey: "root/licenses/product-plan/create",
}
var PERM_ROOT_PRODUCT_PLAN_UPDATE = workspaces.PermissionInfo{
	CompleteKey: "root/licenses/product-plan/update",
}
var PERM_ROOT_PRODUCT_PLAN_QUERY = workspaces.PermissionInfo{
	CompleteKey: "root/licenses/product-plan/query",
}
var PERM_ROOT_PRODUCT_PLAN = workspaces.PermissionInfo{
	CompleteKey: "root/licenses/product-plan/*",
}
var ALL_PRODUCT_PLAN_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_PRODUCT_PLAN_DELETE,
	PERM_ROOT_PRODUCT_PLAN_CREATE,
	PERM_ROOT_PRODUCT_PLAN_UPDATE,
	PERM_ROOT_PRODUCT_PLAN_QUERY,
	PERM_ROOT_PRODUCT_PLAN,
}
