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
)
type ProductCategoryEntity struct {
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
    Translations     []*ProductCategoryEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    Children []*ProductCategoryEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *ProductCategoryEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var ProductCategoryPreloadRelations []string = []string{}
var PRODUCTCATEGORY_EVENT_CREATED = "productCategory.created"
var PRODUCTCATEGORY_EVENT_UPDATED = "productCategory.updated"
var PRODUCTCATEGORY_EVENT_DELETED = "productCategory.deleted"
var PRODUCTCATEGORY_EVENTS = []string{
	PRODUCTCATEGORY_EVENT_CREATED,
	PRODUCTCATEGORY_EVENT_UPDATED,
	PRODUCTCATEGORY_EVENT_DELETED,
}
type ProductCategoryFieldMap struct {
		Name workspaces.TranslatedString `yaml:"name"`
}
var ProductCategoryEntityMetaConfig map[string]int64 = map[string]int64{
}
var ProductCategoryEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&ProductCategoryEntity{}))
  type ProductCategoryEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        Name string `yaml:"name" json:"name"`
  }
func entityProductCategoryFormatter(dto *ProductCategoryEntity, query workspaces.QueryDSL) {
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
func ProductCategoryMockEntity() *ProductCategoryEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &ProductCategoryEntity{
      Name : &stringHolder,
	}
	return entity
}
func ProductCategoryActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := ProductCategoryMockEntity()
		_, err := ProductCategoryActionCreate(entity, query)
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
    func (x*ProductCategoryEntity) GetNameTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Name
          }
        }
      }
      return ""
    }
  func ProductCategoryActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*ProductCategoryEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &ProductCategoryEntity{
          Name: &tildaRef,
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
  func ProductCategoryAssociationCreate(dto *ProductCategoryEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func ProductCategoryRelationContentCreate(dto *ProductCategoryEntity, query workspaces.QueryDSL) error {
return nil
}
func ProductCategoryRelationContentUpdate(dto *ProductCategoryEntity, query workspaces.QueryDSL) error {
	return nil
}
func ProductCategoryPolyglotCreateHandler(dto *ProductCategoryEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &ProductCategoryEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func ProductCategoryValidator(dto *ProductCategoryEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func ProductCategoryEntityPreSanitize(dto *ProductCategoryEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func ProductCategoryEntityBeforeCreateAppend(dto *ProductCategoryEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    ProductCategoryRecursiveAddUniqueId(dto, query)
  }
  func ProductCategoryRecursiveAddUniqueId(dto *ProductCategoryEntity, query workspaces.QueryDSL) {
  }
func ProductCategoryActionBatchCreateFn(dtos []*ProductCategoryEntity, query workspaces.QueryDSL) ([]*ProductCategoryEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*ProductCategoryEntity{}
		for _, item := range dtos {
			s, err := ProductCategoryActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func ProductCategoryActionCreateFn(dto *ProductCategoryEntity, query workspaces.QueryDSL) (*ProductCategoryEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := ProductCategoryValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	ProductCategoryEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	ProductCategoryEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	ProductCategoryPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	ProductCategoryRelationContentCreate(dto, query)
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
	ProductCategoryAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(PRODUCTCATEGORY_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&ProductCategoryEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func ProductCategoryActionGetOne(query workspaces.QueryDSL) (*ProductCategoryEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&ProductCategoryEntity{})
    item, err := workspaces.GetOneEntity[ProductCategoryEntity](query, refl)
    entityProductCategoryFormatter(item, query)
    return item, err
  }
  func ProductCategoryActionQuery(query workspaces.QueryDSL) ([]*ProductCategoryEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&ProductCategoryEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[ProductCategoryEntity](query, refl)
    for _, item := range items {
      entityProductCategoryFormatter(item, query)
    }
    return items, meta, err
  }
  func ProductCategoryUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *ProductCategoryEntity) (*ProductCategoryEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = PRODUCTCATEGORY_EVENT_UPDATED
    ProductCategoryEntityPreSanitize(fields, query)
    var item ProductCategoryEntity
    q := dbref.
      Where(&ProductCategoryEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    ProductCategoryRelationContentUpdate(fields, query)
    ProductCategoryPolyglotCreateHandler(fields, query)
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&ProductCategoryEntity{UniqueId: uniqueId}).
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
  func ProductCategoryActionUpdateFn(query workspaces.QueryDSL, fields *ProductCategoryEntity) (*ProductCategoryEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := ProductCategoryValidator(fields, true); iError != nil {
      return nil, iError
    }
    ProductCategoryRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        _, err := ProductCategoryUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return nil, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return ProductCategoryUpdateExec(dbref, query, fields)
    }
  }
var ProductCategoryWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire productcategories ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		count, _ := ProductCategoryActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func ProductCategoryActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&ProductCategoryEntity{})
	query.ActionRequires = []string{PERM_ROOT_PRODUCTCATEGORY_DELETE}
	return workspaces.RemoveEntity[ProductCategoryEntity](query, refl)
}
func ProductCategoryActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[ProductCategoryEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'ProductCategoryEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func ProductCategoryActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[ProductCategoryEntity]) (
    *workspaces.BulkRecordRequest[ProductCategoryEntity], *workspaces.IError,
  ) {
    result := []*ProductCategoryEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := ProductCategoryActionUpdate(query, record)
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
func (x *ProductCategoryEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var ProductCategoryEntityMeta = workspaces.TableMetaData{
	EntityName:    "ProductCategory",
	ExportKey:    "product-categories",
	TableNameInDb: "fb_productcategory_entities",
	EntityObject:  &ProductCategoryEntity{},
	ExportStream: ProductCategoryActionExportT,
	ImportQuery: ProductCategoryActionImport,
}
func ProductCategoryActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[ProductCategoryEntity](query, ProductCategoryActionQuery, ProductCategoryPreloadRelations)
}
func ProductCategoryActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[ProductCategoryEntity](query, ProductCategoryActionQuery, ProductCategoryPreloadRelations)
}
func ProductCategoryActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content ProductCategoryEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := ProductCategoryActionCreate(&content, query)
	return err
}
var ProductCategoryCommonCliFlags = []cli.Flag{
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
}
var ProductCategoryCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "name",
		StructField:     "Name",
		Required: false,
		Usage:    "name",
		Type: "string",
	},
}
var ProductCategoryCommonCliFlagsOptional = []cli.Flag{
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
}
  var ProductCategoryCreateCmd cli.Command = cli.Command{
    Name:    "create",
    Aliases: []string{"c"},
    Flags: ProductCategoryCommonCliFlags,
    Usage: "Create a new template",
    Action: func(c *cli.Context) {
      query := workspaces.CommonCliQueryDSLBuilder(c)
      entity := CastProductCategoryFromCli(c)
      if entity, err := ProductCategoryActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var ProductCategoryCreateInteractiveCmd cli.Command = cli.Command{
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
      entity := &ProductCategoryEntity{}
      for _, item := range ProductCategoryCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := ProductCategoryActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var ProductCategoryUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: ProductCategoryCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilder(c)
      entity := CastProductCategoryFromCli(c)
      if entity, err := ProductCategoryActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x ProductCategoryEntity) FromCli(c *cli.Context) *ProductCategoryEntity {
	return CastProductCategoryFromCli(c)
}
func CastProductCategoryFromCli (c *cli.Context) *ProductCategoryEntity {
	template := &ProductCategoryEntity{}
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
	return template
}
  func ProductCategorySyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      ProductCategoryActionCreate,
      reflect.ValueOf(&ProductCategoryEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func ProductCategoryWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := ProductCategoryActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "ProductCategory", result)
    }
  }
var ProductCategoryImportExportCommands = []cli.Command{
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
			ProductCategoryActionSeeder(query, c.Int("count"))
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
				Value: "product-category-seeder.yml",
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
			ProductCategoryActionSeederInit(f, c.String("file"), c.String("format"))
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
				Value: "product-category-seeder-product-category.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of product-categories, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]ProductCategoryEntity{}
			workspaces.ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
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
				ProductCategoryActionCreate,
				reflect.ValueOf(&ProductCategoryEntity{}).Elem(),
				c.String("file"),
			)
			return nil
		},
	},
}
    var ProductCategoryCliCommands []cli.Command = []cli.Command{
      workspaces.GetCommonQuery(ProductCategoryActionQuery),
      workspaces.GetCommonTableQuery(reflect.ValueOf(&ProductCategoryEntity{}).Elem(), ProductCategoryActionQuery),
          ProductCategoryCreateCmd,
          ProductCategoryUpdateCmd,
          ProductCategoryCreateInteractiveCmd,
          ProductCategoryWipeCmd,
          workspaces.GetCommonRemoveQuery(reflect.ValueOf(&ProductCategoryEntity{}).Elem(), ProductCategoryActionRemove),
  }
  func ProductCategoryCliFn() cli.Command {
    ProductCategoryCliCommands = append(ProductCategoryCliCommands, ProductCategoryImportExportCommands...)
    return cli.Command{
      Name:        "productCategory",
      Description: "ProductCategorys module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: ProductCategoryCliCommands,
    }
  }
  /**
  *	Override this function on ProductCategoryEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendProductCategoryRouter = func(r *[]workspaces.Module2Action) {}
  func GetProductCategoryModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
       {
        Method: "GET",
        Url:    "/product-categories",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_PRODUCTCATEGORY_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpQueryEntity(c, ProductCategoryActionQuery)
          },
        },
        Format: "QUERY",
        Action: ProductCategoryActionQuery,
        ResponseEntity: &[]ProductCategoryEntity{},
      },
      {
        Method: "GET",
        Url:    "/product-categories/export",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_PRODUCTCATEGORY_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpStreamFileChannel(c, ProductCategoryActionExport)
          },
        },
        Format: "QUERY",
        Action: ProductCategoryActionExport,
        ResponseEntity: &[]ProductCategoryEntity{},
      },
      {
        Method: "GET",
        Url:    "/product-category/:uniqueId",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_PRODUCTCATEGORY_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpGetEntity(c, ProductCategoryActionGetOne)
          },
        },
        Format: "GET_ONE",
        Action: ProductCategoryActionGetOne,
        ResponseEntity: &ProductCategoryEntity{},
      },
      {
        ActionName:    "create",
        ActionAliases: []string{"c"},
        Flags: ProductCategoryCommonCliFlags,
        Method: "POST",
        Url:    "/product-category",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_PRODUCTCATEGORY_CREATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpPostEntity(c, ProductCategoryActionCreate)
          },
        },
        Action: ProductCategoryActionCreate,
        Format: "POST_ONE",
        RequestEntity: &ProductCategoryEntity{},
        ResponseEntity: &ProductCategoryEntity{},
      },
      {
        ActionName:    "update",
        ActionAliases: []string{"u"},
        Flags: ProductCategoryCommonCliFlagsOptional,
        Method: "PATCH",
        Url:    "/product-category",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_PRODUCTCATEGORY_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpUpdateEntity(c, ProductCategoryActionUpdate)
          },
        },
        Action: ProductCategoryActionUpdate,
        RequestEntity: &ProductCategoryEntity{},
        Format: "PATCH_ONE",
        ResponseEntity: &ProductCategoryEntity{},
      },
      {
        Method: "PATCH",
        Url:    "/product-categories",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_PRODUCTCATEGORY_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpUpdateEntities(c, ProductCategoryActionBulkUpdate)
          },
        },
        Action: ProductCategoryActionBulkUpdate,
        Format: "PATCH_BULK",
        RequestEntity:  &workspaces.BulkRecordRequest[ProductCategoryEntity]{},
        ResponseEntity: &workspaces.BulkRecordRequest[ProductCategoryEntity]{},
      },
      {
        Method: "DELETE",
        Url:    "/product-category",
        Format: "DELETE_DSL",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_PRODUCTCATEGORY_DELETE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpRemoveEntity(c, ProductCategoryActionRemove)
          },
        },
        Action: ProductCategoryActionRemove,
        RequestEntity: &workspaces.DeleteRequest{},
        ResponseEntity: &workspaces.DeleteResponse{},
        TargetEntity: &ProductCategoryEntity{},
      },
    }
    // Append user defined functions
    AppendProductCategoryRouter(&routes)
    return routes
  }
  func CreateProductCategoryRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetProductCategoryModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, ProductCategoryEntityJsonSchema, "product-category-http", "shop")
    workspaces.WriteEntitySchema("ProductCategoryEntity", ProductCategoryEntityJsonSchema, "shop")
    return httpRoutes
  }
var PERM_ROOT_PRODUCTCATEGORY_DELETE = "root/productcategory/delete"
var PERM_ROOT_PRODUCTCATEGORY_CREATE = "root/productcategory/create"
var PERM_ROOT_PRODUCTCATEGORY_UPDATE = "root/productcategory/update"
var PERM_ROOT_PRODUCTCATEGORY_QUERY = "root/productcategory/query"
var PERM_ROOT_PRODUCTCATEGORY = "root/productcategory"
var ALL_PRODUCTCATEGORY_PERMISSIONS = []string{
	PERM_ROOT_PRODUCTCATEGORY_DELETE,
	PERM_ROOT_PRODUCTCATEGORY_CREATE,
	PERM_ROOT_PRODUCTCATEGORY_UPDATE,
	PERM_ROOT_PRODUCTCATEGORY_QUERY,
	PERM_ROOT_PRODUCTCATEGORY,
}