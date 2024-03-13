package shop
import (
    "github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/workspaces"
	"log"
	"os"
	"fmt"
	"encoding/json"
	"strings"
	"github.com/schollz/progressbar/v3"
	"github.com/gookit/event"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	jsoniter "github.com/json-iterator/go"
	"embed"
	reflect "reflect"
	"github.com/urfave/cli"
	mocks "github.com/torabian/fireback/modules/shop/mocks/Product"
)
import  "github.com/torabian/fireback/modules/currency"
type ProductEntity struct {
    Visibility       *string                         `json:"visibility,omitempty" yaml:"visibility"`
    WorkspaceId      *string                         `json:"workspaceId,omitempty" yaml:"workspaceId"`
    LinkerId         *string                         `json:"linkerId,omitempty" yaml:"linkerId"`
    ParentId         *string                         `json:"parentId,omitempty" yaml:"parentId"`
    UniqueId         string                          `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
    UserId           *string                         `json:"userId,omitempty" yaml:"userId"`
    Rank             int64                           `json:"rank,omitempty" gorm:"type:int;name:rank"`
    Updated          int64                           `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
    Created          int64                           `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
    CreatedFormatted string                          `json:"createdFormatted,omitempty" sql:"-"`
    UpdatedFormatted string                          `json:"updatedFormatted,omitempty" sql:"-"`
    Name   *string `json:"name" yaml:"name"        translate:"true" `
    // Datenano also has a text representation
    Price   *  currency.PriceTagEntity `json:"price" yaml:"price"    gorm:"foreignKey:PriceId;references:UniqueId"     `
    // Datenano also has a text representation
        PriceId *string `json:"priceId" yaml:"priceId"`
    ProductId   *int64 `json:"productId" yaml:"productId"       `
    // Datenano also has a text representation
    Description   *string `json:"description" yaml:"description"       `
    // Datenano also has a text representation
    Sku   *string `json:"sku" yaml:"sku"       `
    // Datenano also has a text representation
    Brand   *string `json:"brand" yaml:"brand"       `
    // Datenano also has a text representation
    Category   *  CategoryEntity `json:"category" yaml:"category"    gorm:"foreignKey:CategoryId;references:UniqueId"     `
    // Datenano also has a text representation
        CategoryId *string `json:"categoryId" yaml:"categoryId"`
    Tags   []*  TagEntity `json:"tags" yaml:"tags"    gorm:"many2many:product_tags;foreignKey:UniqueId;references:UniqueId"     `
    // Datenano also has a text representation
    TagsListId []string `json:"tagsListId" yaml:"tagsListId" gorm:"-" sql:"-"`
    Translations     []*ProductEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    Children []*ProductEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *ProductEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var ProductPreloadRelations []string = []string{}
var PRODUCT_EVENT_CREATED = "product.created"
var PRODUCT_EVENT_UPDATED = "product.updated"
var PRODUCT_EVENT_DELETED = "product.deleted"
var PRODUCT_EVENTS = []string{
	PRODUCT_EVENT_CREATED,
	PRODUCT_EVENT_UPDATED,
	PRODUCT_EVENT_DELETED,
}
type ProductFieldMap struct {
		Name workspaces.TranslatedString `yaml:"name"`
		Price workspaces.TranslatedString `yaml:"price"`
		ProductId workspaces.TranslatedString `yaml:"productId"`
		Description workspaces.TranslatedString `yaml:"description"`
		Sku workspaces.TranslatedString `yaml:"sku"`
		Brand workspaces.TranslatedString `yaml:"brand"`
		Category workspaces.TranslatedString `yaml:"category"`
		Tags workspaces.TranslatedString `yaml:"tags"`
}
var ProductEntityMetaConfig map[string]int64 = map[string]int64{
}
var ProductEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&ProductEntity{}))
  type ProductEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        Name string `yaml:"name" json:"name"`
  }
func entityProductFormatter(dto *ProductEntity, query workspaces.QueryDSL) {
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
func ProductMockEntity() *ProductEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &ProductEntity{
      Name : &stringHolder,
      ProductId : &int64Holder,
      Description : &stringHolder,
      Sku : &stringHolder,
      Brand : &stringHolder,
	}
	return entity
}
func ProductActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := ProductMockEntity()
		_, err := ProductActionCreate(entity, query)
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
    func (x*ProductEntity) GetNameTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Name
          }
        }
      }
      return ""
    }
  func ProductActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*ProductEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &ProductEntity{
          Name: &tildaRef,
          Description: &tildaRef,
          Sku: &tildaRef,
          Brand: &tildaRef,
          TagsListId: []string{"~"},
          Tags: []*TagEntity{{}},
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
  func ProductAssociationCreate(dto *ProductEntity, query workspaces.QueryDSL) error {
      {
        if dto.TagsListId != nil && len(dto.TagsListId) > 0 {
          var items []TagEntity
          err := query.Tx.Where(dto.TagsListId).Find(&items).Error
          if err != nil {
              return err
          }
          err = query.Tx.Model(dto).Association("Tags").Replace(items)
          if err != nil {
              return err
          }
        }
      }
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func ProductRelationContentCreate(dto *ProductEntity, query workspaces.QueryDSL) error {
return nil
}
func ProductRelationContentUpdate(dto *ProductEntity, query workspaces.QueryDSL) error {
	return nil
}
func ProductPolyglotCreateHandler(dto *ProductEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &ProductEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func ProductValidator(dto *ProductEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func ProductEntityPreSanitize(dto *ProductEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func ProductEntityBeforeCreateAppend(dto *ProductEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    ProductRecursiveAddUniqueId(dto, query)
  }
  func ProductRecursiveAddUniqueId(dto *ProductEntity, query workspaces.QueryDSL) {
  }
func ProductActionBatchCreateFn(dtos []*ProductEntity, query workspaces.QueryDSL) ([]*ProductEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*ProductEntity{}
		for _, item := range dtos {
			s, err := ProductActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func ProductActionCreateFn(dto *ProductEntity, query workspaces.QueryDSL) (*ProductEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := ProductValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	ProductEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	ProductEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	ProductPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	ProductRelationContentCreate(dto, query)
	// 4. Create the entity
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()
	} else {
		dbref = query.Tx
	}
	query.Tx = dbref;
	err := dbref.Create(&dto).Error
	if err != nil {
		err := workspaces.GormErrorToIError(err)
		return dto, err
	}
	// 5. Create sub entities, objects or arrays, association to other entities
	ProductAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(PRODUCT_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&ProductEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func ProductActionGetOne(query workspaces.QueryDSL) (*ProductEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&ProductEntity{})
    item, err := workspaces.GetOneEntity[ProductEntity](query, refl)
    entityProductFormatter(item, query)
    return item, err
  }
  func ProductActionQuery(query workspaces.QueryDSL) ([]*ProductEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&ProductEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[ProductEntity](query, refl)
    for _, item := range items {
      entityProductFormatter(item, query)
    }
    return items, meta, err
  }
  func ProductUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *ProductEntity) (*ProductEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = PRODUCT_EVENT_UPDATED
    ProductEntityPreSanitize(fields, query)
    var item ProductEntity
    q := dbref.
      Where(&ProductEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    ProductRelationContentUpdate(fields, query)
    ProductPolyglotCreateHandler(fields, query)
    // @meta(update has many)
        if fields.TagsListId  != nil {
          var items []TagEntity
          if len(fields.TagsListId ) > 0 {
            dbref.
              Where(&fields.TagsListId ).
              Find(&items)
          }
          dbref.
            Model(&ProductEntity{UniqueId: uniqueId}).
            Association("Tags").
            Replace(&items)
        }
    err = dbref.
      Preload(clause.Associations).
      Where(&ProductEntity{UniqueId: uniqueId}).
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
  func ProductActionUpdateFn(query workspaces.QueryDSL, fields *ProductEntity) (*ProductEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := ProductValidator(fields, true); iError != nil {
      return nil, iError
    }
    ProductRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        _, err := ProductUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return nil, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return ProductUpdateExec(dbref, query, fields)
    }
  }
var ProductWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire products ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		count, _ := ProductActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func ProductActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&ProductEntity{})
	query.ActionRequires = []string{PERM_ROOT_PRODUCT_DELETE}
	return workspaces.RemoveEntity[ProductEntity](query, refl)
}
func ProductActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[ProductEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'ProductEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func ProductActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[ProductEntity]) (
    *workspaces.BulkRecordRequest[ProductEntity], *workspaces.IError,
  ) {
    result := []*ProductEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := ProductActionUpdate(query, record)
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
func (x *ProductEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var ProductEntityMeta = workspaces.TableMetaData{
	EntityName:    "Product",
	ExportKey:    "products",
	TableNameInDb: "fb_product_entities",
	EntityObject:  &ProductEntity{},
	ExportStream: ProductActionExportT,
	ImportQuery: ProductActionImport,
}
func ProductActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[ProductEntity](query, ProductActionQuery, ProductPreloadRelations)
}
func ProductActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[ProductEntity](query, ProductActionQuery, ProductPreloadRelations)
}
func ProductActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content ProductEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := ProductActionCreate(&content, query)
	return err
}
var ProductCommonCliFlags = []cli.Flag{
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
      Name:     "price-id",
      Required: false,
      Usage:    "price",
    },
    &cli.Int64Flag{
      Name:     "product-id",
      Required: false,
      Usage:    "productId",
    },
    &cli.StringFlag{
      Name:     "description",
      Required: false,
      Usage:    "Detailed description of the product",
    },
    &cli.StringFlag{
      Name:     "sku",
      Required: false,
      Usage:    "Stock Keeping Unit code for the product",
    },
    &cli.StringFlag{
      Name:     "brand",
      Required: false,
      Usage:    "Brand of the product",
    },
    &cli.StringFlag{
      Name:     "category-id",
      Required: false,
      Usage:    "category",
    },
    &cli.StringSliceFlag{
      Name:     "tags",
      Required: false,
      Usage:    "tags",
    },
}
var ProductCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "name",
		StructField:     "Name",
		Required: false,
		Usage:    "name",
		Type: "string",
	},
	{
		Name:     "productId",
		StructField:     "ProductId",
		Required: false,
		Usage:    "productId",
		Type: "int64",
	},
	{
		Name:     "description",
		StructField:     "Description",
		Required: false,
		Usage:    "Detailed description of the product",
		Type: "string",
	},
	{
		Name:     "sku",
		StructField:     "Sku",
		Required: false,
		Usage:    "Stock Keeping Unit code for the product",
		Type: "string",
	},
	{
		Name:     "brand",
		StructField:     "Brand",
		Required: false,
		Usage:    "Brand of the product",
		Type: "string",
	},
}
var ProductCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "price-id",
      Required: false,
      Usage:    "price",
    },
    &cli.Int64Flag{
      Name:     "product-id",
      Required: false,
      Usage:    "productId",
    },
    &cli.StringFlag{
      Name:     "description",
      Required: false,
      Usage:    "Detailed description of the product",
    },
    &cli.StringFlag{
      Name:     "sku",
      Required: false,
      Usage:    "Stock Keeping Unit code for the product",
    },
    &cli.StringFlag{
      Name:     "brand",
      Required: false,
      Usage:    "Brand of the product",
    },
    &cli.StringFlag{
      Name:     "category-id",
      Required: false,
      Usage:    "category",
    },
    &cli.StringSliceFlag{
      Name:     "tags",
      Required: false,
      Usage:    "tags",
    },
}
  var ProductCreateCmd cli.Command = cli.Command{
    Name:    "create",
    Aliases: []string{"c"},
    Flags: ProductCommonCliFlags,
    Usage: "Create a new template",
    Action: func(c *cli.Context) {
      query := workspaces.CommonCliQueryDSLBuilder(c)
      entity := CastProductFromCli(c)
      if entity, err := ProductActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var ProductCreateInteractiveCmd cli.Command = cli.Command{
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
      entity := &ProductEntity{}
      for _, item := range ProductCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := ProductActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var ProductUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: ProductCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilder(c)
      entity := CastProductFromCli(c)
      if entity, err := ProductActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x ProductEntity) FromCli(c *cli.Context) *ProductEntity {
	return CastProductFromCli(c)
}
func CastProductFromCli (c *cli.Context) *ProductEntity {
	template := &ProductEntity{}
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
      if c.IsSet("price-id") {
        value := c.String("price-id")
        template.PriceId = &value
      }
      if c.IsSet("description") {
        value := c.String("description")
        template.Description = &value
      }
      if c.IsSet("sku") {
        value := c.String("sku")
        template.Sku = &value
      }
      if c.IsSet("brand") {
        value := c.String("brand")
        template.Brand = &value
      }
      if c.IsSet("category-id") {
        value := c.String("category-id")
        template.CategoryId = &value
      }
      if c.IsSet("tags") {
        value := c.String("tags")
        template.TagsListId = strings.Split(value, ",")
      }
	return template
}
  func ProductSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      ProductActionCreate,
      reflect.ValueOf(&ProductEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func ProductImportMocks() {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      ProductActionCreate,
      reflect.ValueOf(&ProductEntity{}).Elem(),
      &mocks.ViewsFs,
      []string{},
      false,
    )
  }
  func ProductWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := ProductActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "Product", result)
    }
  }
var ProductImportExportCommands = []cli.Command{
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
			ProductActionSeeder(query, c.Int("count"))
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
				Value: "product-seeder.yml",
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
			ProductActionSeederInit(f, c.String("file"), c.String("format"))
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
				Value: "product-seeder-product.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of products, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]ProductEntity{}
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
					ProductActionCreate,
					reflect.ValueOf(&ProductEntity{}).Elem(),
					&mocks.ViewsFs,
				)
				return nil
			},
		},
	cli.Command{
		Name:    "import",
		Flags: append(workspaces.CommonQueryFlags,
			&cli.StringFlag{
				Name:     "file",
				Usage:    "The address of file you want the csv be imported from",
				Required: true,
			}),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmd(c,
				ProductActionCreate,
				reflect.ValueOf(&ProductEntity{}).Elem(),
				c.String("file"),
			)
			return nil
		},
	},
}
    var ProductCliCommands []cli.Command = []cli.Command{
      workspaces.GetCommonQuery(ProductActionQuery),
      workspaces.GetCommonTableQuery(reflect.ValueOf(&ProductEntity{}).Elem(), ProductActionQuery),
          ProductCreateCmd,
          ProductUpdateCmd,
          ProductCreateInteractiveCmd,
          ProductWipeCmd,
          workspaces.GetCommonRemoveQuery(reflect.ValueOf(&ProductEntity{}).Elem(), ProductActionRemove),
  }
  func ProductCliFn() cli.Command {
    ProductCliCommands = append(ProductCliCommands, ProductImportExportCommands...)
    return cli.Command{
      Name:        "product",
      Description: "Products module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: ProductCliCommands,
    }
  }
  /**
  *	Override this function on ProductEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendProductRouter = func(r *[]workspaces.Module2Action) {}
  func GetProductModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
       {
        Method: "GET",
        Url:    "/products",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_PRODUCT_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpQueryEntity(c, ProductActionQuery)
          },
        },
        Format: "QUERY",
        Action: ProductActionQuery,
        ResponseEntity: &[]ProductEntity{},
      },
      {
        Method: "GET",
        Url:    "/products/export",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_PRODUCT_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpStreamFileChannel(c, ProductActionExport)
          },
        },
        Format: "QUERY",
        Action: ProductActionExport,
        ResponseEntity: &[]ProductEntity{},
      },
      {
        Method: "GET",
        Url:    "/product/:uniqueId",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_PRODUCT_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpGetEntity(c, ProductActionGetOne)
          },
        },
        Format: "GET_ONE",
        Action: ProductActionGetOne,
        ResponseEntity: &ProductEntity{},
      },
      {
        ActionName:    "create",
        ActionAliases: []string{"c"},
        Flags: ProductCommonCliFlags,
        Method: "POST",
        Url:    "/product",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_PRODUCT_CREATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpPostEntity(c, ProductActionCreate)
          },
        },
        Action: ProductActionCreate,
        Format: "POST_ONE",
        RequestEntity: &ProductEntity{},
        ResponseEntity: &ProductEntity{},
      },
      {
        ActionName:    "update",
        ActionAliases: []string{"u"},
        Flags: ProductCommonCliFlagsOptional,
        Method: "PATCH",
        Url:    "/product",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_PRODUCT_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpUpdateEntity(c, ProductActionUpdate)
          },
        },
        Action: ProductActionUpdate,
        RequestEntity: &ProductEntity{},
        Format: "PATCH_ONE",
        ResponseEntity: &ProductEntity{},
      },
      {
        Method: "PATCH",
        Url:    "/products",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_PRODUCT_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpUpdateEntities(c, ProductActionBulkUpdate)
          },
        },
        Action: ProductActionBulkUpdate,
        Format: "PATCH_BULK",
        RequestEntity:  &workspaces.BulkRecordRequest[ProductEntity]{},
        ResponseEntity: &workspaces.BulkRecordRequest[ProductEntity]{},
      },
      {
        Method: "DELETE",
        Url:    "/product",
        Format: "DELETE_DSL",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_PRODUCT_DELETE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpRemoveEntity(c, ProductActionRemove)
          },
        },
        Action: ProductActionRemove,
        RequestEntity: &workspaces.DeleteRequest{},
        ResponseEntity: &workspaces.DeleteResponse{},
        TargetEntity: &ProductEntity{},
      },
    }
    // Append user defined functions
    AppendProductRouter(&routes)
    return routes
  }
  func CreateProductRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetProductModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, ProductEntityJsonSchema, "product-http", "shop")
    workspaces.WriteEntitySchema("ProductEntity", ProductEntityJsonSchema, "shop")
    return httpRoutes
  }
var PERM_ROOT_PRODUCT_DELETE = "root/product/delete"
var PERM_ROOT_PRODUCT_CREATE = "root/product/create"
var PERM_ROOT_PRODUCT_UPDATE = "root/product/update"
var PERM_ROOT_PRODUCT_QUERY = "root/product/query"
var PERM_ROOT_PRODUCT = "root/product"
var ALL_PRODUCT_PERMISSIONS = []string{
	PERM_ROOT_PRODUCT_DELETE,
	PERM_ROOT_PRODUCT_CREATE,
	PERM_ROOT_PRODUCT_UPDATE,
	PERM_ROOT_PRODUCT_QUERY,
	PERM_ROOT_PRODUCT,
}