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
	mocks "github.com/torabian/fireback/modules/shop/mocks/Category"
)
type CategoryEntity struct {
    Visibility       *string                         `json:"visibility,omitempty" yaml:"visibility"`
    WorkspaceId      *string                         `json:"workspaceId,omitempty" yaml:"workspaceId"`
    LinkerId         *string                         `json:"linkerId,omitempty" yaml:"linkerId"`
    ParentId         *string                         `json:"parentId,omitempty" yaml:"parentId"`
    UniqueId         string                          `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
    UserId           *string                         `json:"userId,omitempty" yaml:"userId"`
    Rank             int64                           `json:"rank,omitempty" gorm:"type:int;name:rank"`
    Updated          int64                           `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
    Created          int64                           `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
    CreatedFormatted string                          `json:"createdFormatted,omitempty" sql:"-" gorm:"-"`
    UpdatedFormatted string                          `json:"updatedFormatted,omitempty" sql:"-" gorm:"-"`
    Name   *string `json:"name" yaml:"name"        translate:"true" `
    // Datenano also has a text representation
    Translations     []*CategoryEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    Children []*CategoryEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *CategoryEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var CategoryPreloadRelations []string = []string{}
var CATEGORY_EVENT_CREATED = "category.created"
var CATEGORY_EVENT_UPDATED = "category.updated"
var CATEGORY_EVENT_DELETED = "category.deleted"
var CATEGORY_EVENTS = []string{
	CATEGORY_EVENT_CREATED,
	CATEGORY_EVENT_UPDATED,
	CATEGORY_EVENT_DELETED,
}
type CategoryFieldMap struct {
		Name workspaces.TranslatedString `yaml:"name"`
}
var CategoryEntityMetaConfig map[string]int64 = map[string]int64{
}
var CategoryEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&CategoryEntity{}))
  type CategoryEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        Name string `yaml:"name" json:"name"`
  }
func entityCategoryFormatter(dto *CategoryEntity, query workspaces.QueryDSL) {
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
func CategoryMockEntity() *CategoryEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &CategoryEntity{
      Name : &stringHolder,
	}
	return entity
}
func CategoryActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := CategoryMockEntity()
		_, err := CategoryActionCreate(entity, query)
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
    func (x*CategoryEntity) GetNameTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Name
          }
        }
      }
      return ""
    }
  func CategoryActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*CategoryEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &CategoryEntity{
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
  func CategoryAssociationCreate(dto *CategoryEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func CategoryRelationContentCreate(dto *CategoryEntity, query workspaces.QueryDSL) error {
return nil
}
func CategoryRelationContentUpdate(dto *CategoryEntity, query workspaces.QueryDSL) error {
	return nil
}
func CategoryPolyglotCreateHandler(dto *CategoryEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &CategoryEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func CategoryValidator(dto *CategoryEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func CategoryEntityPreSanitize(dto *CategoryEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func CategoryEntityBeforeCreateAppend(dto *CategoryEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    CategoryRecursiveAddUniqueId(dto, query)
  }
  func CategoryRecursiveAddUniqueId(dto *CategoryEntity, query workspaces.QueryDSL) {
  }
func CategoryActionBatchCreateFn(dtos []*CategoryEntity, query workspaces.QueryDSL) ([]*CategoryEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*CategoryEntity{}
		for _, item := range dtos {
			s, err := CategoryActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func CategoryActionCreateFn(dto *CategoryEntity, query workspaces.QueryDSL) (*CategoryEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := CategoryValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	CategoryEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	CategoryEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	CategoryPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	CategoryRelationContentCreate(dto, query)
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
	CategoryAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(CATEGORY_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&CategoryEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func CategoryActionGetOne(query workspaces.QueryDSL) (*CategoryEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&CategoryEntity{})
    item, err := workspaces.GetOneEntity[CategoryEntity](query, refl)
    entityCategoryFormatter(item, query)
    return item, err
  }
  func CategoryActionQuery(query workspaces.QueryDSL) ([]*CategoryEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&CategoryEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[CategoryEntity](query, refl)
    for _, item := range items {
      entityCategoryFormatter(item, query)
    }
    return items, meta, err
  }
  func CategoryUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *CategoryEntity) (*CategoryEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = CATEGORY_EVENT_UPDATED
    CategoryEntityPreSanitize(fields, query)
    var item CategoryEntity
    q := dbref.
      Where(&CategoryEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    CategoryRelationContentUpdate(fields, query)
    CategoryPolyglotCreateHandler(fields, query)
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&CategoryEntity{UniqueId: uniqueId}).
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
  func CategoryActionUpdateFn(query workspaces.QueryDSL, fields *CategoryEntity) (*CategoryEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := CategoryValidator(fields, true); iError != nil {
      return nil, iError
    }
    CategoryRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *CategoryEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = CategoryUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return CategoryUpdateExec(dbref, query, fields)
    }
  }
var CategoryWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire categories ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []string{PERM_ROOT_CATEGORY_DELETE},
    })
		count, _ := CategoryActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func CategoryActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&CategoryEntity{})
	query.ActionRequires = []string{PERM_ROOT_CATEGORY_DELETE}
	return workspaces.RemoveEntity[CategoryEntity](query, refl)
}
func CategoryActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[CategoryEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'CategoryEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func CategoryActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[CategoryEntity]) (
    *workspaces.BulkRecordRequest[CategoryEntity], *workspaces.IError,
  ) {
    result := []*CategoryEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := CategoryActionUpdate(query, record)
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
func (x *CategoryEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var CategoryEntityMeta = workspaces.TableMetaData{
	EntityName:    "Category",
	ExportKey:    "categories",
	TableNameInDb: "fb_category_entities",
	EntityObject:  &CategoryEntity{},
	ExportStream: CategoryActionExportT,
	ImportQuery: CategoryActionImport,
}
func CategoryActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[CategoryEntity](query, CategoryActionQuery, CategoryPreloadRelations)
}
func CategoryActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[CategoryEntity](query, CategoryActionQuery, CategoryPreloadRelations)
}
func CategoryActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content CategoryEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := CategoryActionCreate(&content, query)
	return err
}
var CategoryCommonCliFlags = []cli.Flag{
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
var CategoryCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "name",
		StructField:     "Name",
		Required: false,
		Usage:    "name",
		Type: "string",
	},
}
var CategoryCommonCliFlagsOptional = []cli.Flag{
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
  var CategoryCreateCmd cli.Command = cli.Command{
    Name:    "create",
    Aliases: []string{"c"},
    Flags: CategoryCommonCliFlags,
    Usage: "Create a new template",
    Action: func(c *cli.Context) {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []string{PERM_ROOT_CATEGORY_CREATE},
      })
      entity := CastCategoryFromCli(c)
      if entity, err := CategoryActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var CategoryCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []string{PERM_ROOT_CATEGORY_CREATE},
      })
      entity := &CategoryEntity{}
      for _, item := range CategoryCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := CategoryActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var CategoryUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: CategoryCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []string{PERM_ROOT_CATEGORY_UPDATE},
      })
      entity := CastCategoryFromCli(c)
      if entity, err := CategoryActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x CategoryEntity) FromCli(c *cli.Context) *CategoryEntity {
	return CastCategoryFromCli(c)
}
func CastCategoryFromCli (c *cli.Context) *CategoryEntity {
	template := &CategoryEntity{}
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
  func CategorySyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      CategoryActionCreate,
      reflect.ValueOf(&CategoryEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func CategoryImportMocks() {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      CategoryActionCreate,
      reflect.ValueOf(&CategoryEntity{}).Elem(),
      &mocks.ViewsFs,
      []string{},
      false,
    )
  }
  func CategoryWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := CategoryActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "Category", result)
    }
  }
var CategoryImportExportCommands = []cli.Command{
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
        ActionRequires: []string{PERM_ROOT_CATEGORY_CREATE},
      })
			CategoryActionSeeder(query, c.Int("count"))
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
				Value: "category-seeder.yml",
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
        ActionRequires: []string{PERM_ROOT_CATEGORY_CREATE},
      })
			CategoryActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "category-seeder-category.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of categories, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]CategoryEntity{}
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
					CategoryActionCreate,
					reflect.ValueOf(&CategoryEntity{}).Elem(),
					&mocks.ViewsFs,
				)
				return nil
			},
		},
	cli.Command{
		Name:    "import",
    Flags: append(
			append(
				workspaces.CommonQueryFlags,
				&cli.StringFlag{
					Name:     "file",
					Usage:    "The address of file you want the csv be imported from",
					Required: true,
				}),
			CategoryCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				CategoryActionCreate,
				reflect.ValueOf(&CategoryEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []string{PERM_ROOT_CATEGORY_CREATE},
				},
        func() CategoryEntity {
					v := CastCategoryFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var CategoryCliCommands []cli.Command = []cli.Command{
      workspaces.GetCommonQuery2(CategoryActionQuery, &workspaces.SecurityModel{
        ActionRequires: []string{PERM_ROOT_CATEGORY_CREATE},
      }),
      workspaces.GetCommonTableQuery(reflect.ValueOf(&CategoryEntity{}).Elem(), CategoryActionQuery),
          CategoryCreateCmd,
          CategoryUpdateCmd,
          CategoryCreateInteractiveCmd,
          CategoryWipeCmd,
          workspaces.GetCommonRemoveQuery(reflect.ValueOf(&CategoryEntity{}).Elem(), CategoryActionRemove),
  }
  func CategoryCliFn() cli.Command {
    CategoryCliCommands = append(CategoryCliCommands, CategoryImportExportCommands...)
    return cli.Command{
      Name:        "category",
      Description: "Categorys module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: CategoryCliCommands,
    }
  }
  /**
  *	Override this function on CategoryEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendCategoryRouter = func(r *[]workspaces.Module2Action) {}
  func GetCategoryModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
       {
        Method: "GET",
        Url:    "/categories",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_CATEGORY_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpQueryEntity(c, CategoryActionQuery)
          },
        },
        Format: "QUERY",
        Action: CategoryActionQuery,
        ResponseEntity: &[]CategoryEntity{},
      },
      {
        Method: "GET",
        Url:    "/categories/export",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_CATEGORY_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpStreamFileChannel(c, CategoryActionExport)
          },
        },
        Format: "QUERY",
        Action: CategoryActionExport,
        ResponseEntity: &[]CategoryEntity{},
      },
      {
        Method: "GET",
        Url:    "/category/:uniqueId",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_CATEGORY_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpGetEntity(c, CategoryActionGetOne)
          },
        },
        Format: "GET_ONE",
        Action: CategoryActionGetOne,
        ResponseEntity: &CategoryEntity{},
      },
      {
        ActionName:    "create",
        ActionAliases: []string{"c"},
        Flags: CategoryCommonCliFlags,
        Method: "POST",
        Url:    "/category",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_CATEGORY_CREATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpPostEntity(c, CategoryActionCreate)
          },
        },
        Action: CategoryActionCreate,
        Format: "POST_ONE",
        RequestEntity: &CategoryEntity{},
        ResponseEntity: &CategoryEntity{},
      },
      {
        ActionName:    "update",
        ActionAliases: []string{"u"},
        Flags: CategoryCommonCliFlagsOptional,
        Method: "PATCH",
        Url:    "/category",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_CATEGORY_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpUpdateEntity(c, CategoryActionUpdate)
          },
        },
        Action: CategoryActionUpdate,
        RequestEntity: &CategoryEntity{},
        Format: "PATCH_ONE",
        ResponseEntity: &CategoryEntity{},
      },
      {
        Method: "PATCH",
        Url:    "/categories",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_CATEGORY_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpUpdateEntities(c, CategoryActionBulkUpdate)
          },
        },
        Action: CategoryActionBulkUpdate,
        Format: "PATCH_BULK",
        RequestEntity:  &workspaces.BulkRecordRequest[CategoryEntity]{},
        ResponseEntity: &workspaces.BulkRecordRequest[CategoryEntity]{},
      },
      {
        Method: "DELETE",
        Url:    "/category",
        Format: "DELETE_DSL",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_CATEGORY_DELETE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpRemoveEntity(c, CategoryActionRemove)
          },
        },
        Action: CategoryActionRemove,
        RequestEntity: &workspaces.DeleteRequest{},
        ResponseEntity: &workspaces.DeleteResponse{},
        TargetEntity: &CategoryEntity{},
      },
    }
    // Append user defined functions
    AppendCategoryRouter(&routes)
    return routes
  }
  func CreateCategoryRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetCategoryModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, CategoryEntityJsonSchema, "category-http", "shop")
    workspaces.WriteEntitySchema("CategoryEntity", CategoryEntityJsonSchema, "shop")
    return httpRoutes
  }
var PERM_ROOT_CATEGORY_DELETE = "root/category/delete"
var PERM_ROOT_CATEGORY_CREATE = "root/category/create"
var PERM_ROOT_CATEGORY_UPDATE = "root/category/update"
var PERM_ROOT_CATEGORY_QUERY = "root/category/query"
var PERM_ROOT_CATEGORY = "root/category"
var ALL_CATEGORY_PERMISSIONS = []string{
	PERM_ROOT_CATEGORY_DELETE,
	PERM_ROOT_CATEGORY_CREATE,
	PERM_ROOT_CATEGORY_UPDATE,
	PERM_ROOT_CATEGORY_QUERY,
	PERM_ROOT_CATEGORY,
}