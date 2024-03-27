package cms
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
type PageCategoryEntity struct {
    Visibility       *string                         `json:"visibility,omitempty" yaml:"visibility"`
    WorkspaceId      *string                         `json:"workspaceId,omitempty" yaml:"workspaceId"`
    LinkerId         *string                         `json:"linkerId,omitempty" yaml:"linkerId"`
    ParentId         *string                         `json:"parentId,omitempty" yaml:"parentId"`
    IsDeletable         *bool                         `json:"isDeletable,omitempty" yaml:"isDeletable" gorm:"default:true"`
    IsUpdatable         *bool                         `json:"isUpdatable,omitempty" yaml:"isUpdatable" gorm:"default:true"`
    UniqueId         string                          `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
    UserId           *string                         `json:"userId,omitempty" yaml:"userId"`
    Rank             int64                           `json:"rank,omitempty" gorm:"type:int;name:rank"`
    Updated          int64                           `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
    Created          int64                           `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
    CreatedFormatted string                          `json:"createdFormatted,omitempty" sql:"-" gorm:"-"`
    UpdatedFormatted string                          `json:"updatedFormatted,omitempty" sql:"-" gorm:"-"`
    Name   *string `json:"name" yaml:"name"  validate:"required"        translate:"true" `
    // Datenano also has a text representation
    Translations     []*PageCategoryEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    Children []*PageCategoryEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *PageCategoryEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var PageCategoryPreloadRelations []string = []string{}
var PAGE_CATEGORY_EVENT_CREATED = "pageCategory.created"
var PAGE_CATEGORY_EVENT_UPDATED = "pageCategory.updated"
var PAGE_CATEGORY_EVENT_DELETED = "pageCategory.deleted"
var PAGE_CATEGORY_EVENTS = []string{
	PAGE_CATEGORY_EVENT_CREATED,
	PAGE_CATEGORY_EVENT_UPDATED,
	PAGE_CATEGORY_EVENT_DELETED,
}
type PageCategoryFieldMap struct {
		Name workspaces.TranslatedString `yaml:"name"`
}
var PageCategoryEntityMetaConfig map[string]int64 = map[string]int64{
}
var PageCategoryEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&PageCategoryEntity{}))
  type PageCategoryEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        Name string `yaml:"name" json:"name"`
  }
func entityPageCategoryFormatter(dto *PageCategoryEntity, query workspaces.QueryDSL) {
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
func PageCategoryMockEntity() *PageCategoryEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &PageCategoryEntity{
      Name : &stringHolder,
	}
	return entity
}
func PageCategoryActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := PageCategoryMockEntity()
		_, err := PageCategoryActionCreate(entity, query)
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
    func (x*PageCategoryEntity) GetNameTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Name
          }
        }
      }
      return ""
    }
  func PageCategoryActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*PageCategoryEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &PageCategoryEntity{
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
  func PageCategoryAssociationCreate(dto *PageCategoryEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func PageCategoryRelationContentCreate(dto *PageCategoryEntity, query workspaces.QueryDSL) error {
return nil
}
func PageCategoryRelationContentUpdate(dto *PageCategoryEntity, query workspaces.QueryDSL) error {
	return nil
}
func PageCategoryPolyglotCreateHandler(dto *PageCategoryEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &PageCategoryEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func PageCategoryValidator(dto *PageCategoryEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func PageCategoryEntityPreSanitize(dto *PageCategoryEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func PageCategoryEntityBeforeCreateAppend(dto *PageCategoryEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    PageCategoryRecursiveAddUniqueId(dto, query)
  }
  func PageCategoryRecursiveAddUniqueId(dto *PageCategoryEntity, query workspaces.QueryDSL) {
  }
func PageCategoryActionBatchCreateFn(dtos []*PageCategoryEntity, query workspaces.QueryDSL) ([]*PageCategoryEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*PageCategoryEntity{}
		for _, item := range dtos {
			s, err := PageCategoryActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func PageCategoryDeleteEntireChildren(query workspaces.QueryDSL, dto *PageCategoryEntity) (*workspaces.IError) {
  return nil
}
func PageCategoryActionCreateFn(dto *PageCategoryEntity, query workspaces.QueryDSL) (*PageCategoryEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := PageCategoryValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	PageCategoryEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	PageCategoryEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	PageCategoryPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	PageCategoryRelationContentCreate(dto, query)
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
	PageCategoryAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(PAGE_CATEGORY_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&PageCategoryEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func PageCategoryActionGetOne(query workspaces.QueryDSL) (*PageCategoryEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&PageCategoryEntity{})
    item, err := workspaces.GetOneEntity[PageCategoryEntity](query, refl)
    entityPageCategoryFormatter(item, query)
    return item, err
  }
  func PageCategoryActionQuery(query workspaces.QueryDSL) ([]*PageCategoryEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&PageCategoryEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[PageCategoryEntity](query, refl)
    for _, item := range items {
      entityPageCategoryFormatter(item, query)
    }
    return items, meta, err
  }
  func PageCategoryUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *PageCategoryEntity) (*PageCategoryEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = PAGE_CATEGORY_EVENT_UPDATED
    PageCategoryEntityPreSanitize(fields, query)
    var item PageCategoryEntity
    q := dbref.
      Where(&PageCategoryEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    PageCategoryRelationContentUpdate(fields, query)
    PageCategoryPolyglotCreateHandler(fields, query)
    if ero := PageCategoryDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&PageCategoryEntity{UniqueId: uniqueId}).
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
  func PageCategoryActionUpdateFn(query workspaces.QueryDSL, fields *PageCategoryEntity) (*PageCategoryEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := PageCategoryValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // PageCategoryRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *PageCategoryEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = PageCategoryUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return PageCategoryUpdateExec(dbref, query, fields)
    }
  }
var PageCategoryWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire pagecategories ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_CATEGORY_DELETE},
    })
		count, _ := PageCategoryActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func PageCategoryActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&PageCategoryEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_PAGE_CATEGORY_DELETE}
	return workspaces.RemoveEntity[PageCategoryEntity](query, refl)
}
func PageCategoryActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[PageCategoryEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'PageCategoryEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func PageCategoryActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[PageCategoryEntity]) (
    *workspaces.BulkRecordRequest[PageCategoryEntity], *workspaces.IError,
  ) {
    result := []*PageCategoryEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := PageCategoryActionUpdate(query, record)
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
func (x *PageCategoryEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var PageCategoryEntityMeta = workspaces.TableMetaData{
	EntityName:    "PageCategory",
	ExportKey:    "page-categories",
	TableNameInDb: "fb_page-category_entities",
	EntityObject:  &PageCategoryEntity{},
	ExportStream: PageCategoryActionExportT,
	ImportQuery: PageCategoryActionImport,
}
func PageCategoryActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[PageCategoryEntity](query, PageCategoryActionQuery, PageCategoryPreloadRelations)
}
func PageCategoryActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[PageCategoryEntity](query, PageCategoryActionQuery, PageCategoryPreloadRelations)
}
func PageCategoryActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content PageCategoryEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := PageCategoryActionCreate(&content, query)
	return err
}
var PageCategoryCommonCliFlags = []cli.Flag{
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
}
var PageCategoryCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "name",
		StructField:     "Name",
		Required: true,
		Usage:    "name",
		Type: "string",
	},
}
var PageCategoryCommonCliFlagsOptional = []cli.Flag{
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
}
  var PageCategoryCreateCmd cli.Command = PAGE_CATEGORY_ACTION_POST_ONE.ToCli()
  var PageCategoryCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_CATEGORY_CREATE},
      })
      entity := &PageCategoryEntity{}
      for _, item := range PageCategoryCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := PageCategoryActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var PageCategoryUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: PageCategoryCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_CATEGORY_UPDATE},
      })
      entity := CastPageCategoryFromCli(c)
      if entity, err := PageCategoryActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* PageCategoryEntity) FromCli(c *cli.Context) *PageCategoryEntity {
	return CastPageCategoryFromCli(c)
}
func CastPageCategoryFromCli (c *cli.Context) *PageCategoryEntity {
	template := &PageCategoryEntity{}
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
  func PageCategorySyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      PageCategoryActionCreate,
      reflect.ValueOf(&PageCategoryEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func PageCategoryWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := PageCategoryActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "PageCategory", result)
    }
  }
var PageCategoryImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_CATEGORY_CREATE},
      })
			PageCategoryActionSeeder(query, c.Int("count"))
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
				Value: "page-category-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_CATEGORY_CREATE},
      })
			PageCategoryActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "page-category-seeder-page-category.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of page-categories, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]PageCategoryEntity{}
			workspaces.ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
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
			PageCategoryCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				PageCategoryActionCreate,
				reflect.ValueOf(&PageCategoryEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_CATEGORY_CREATE},
				},
        func() PageCategoryEntity {
					v := CastPageCategoryFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var PageCategoryCliCommands []cli.Command = []cli.Command{
      PAGE_CATEGORY_ACTION_QUERY.ToCli(),
      PAGE_CATEGORY_ACTION_TABLE.ToCli(),
      PageCategoryCreateCmd,
      PageCategoryUpdateCmd,
      PageCategoryCreateInteractiveCmd,
      PageCategoryWipeCmd,
      workspaces.GetCommonRemoveQuery(reflect.ValueOf(&PageCategoryEntity{}).Elem(), PageCategoryActionRemove),
  }
  func PageCategoryCliFn() cli.Command {
    PageCategoryCliCommands = append(PageCategoryCliCommands, PageCategoryImportExportCommands...)
    return cli.Command{
      Name:        "pageCategory",
      Description: "PageCategorys module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: PageCategoryCliCommands,
    }
  }
var PAGE_CATEGORY_ACTION_TABLE = workspaces.Module2Action{
  Name:    "table",
  ActionAliases: []string{"t"},
  Flags:  workspaces.CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: PageCategoryActionQuery,
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    workspaces.CommonCliTableCmd2(c,
      PageCategoryActionQuery,
      security,
      reflect.ValueOf(&PageCategoryEntity{}).Elem(),
    )
    return nil
  },
}
var PAGE_CATEGORY_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/page-categories",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_CATEGORY_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, PageCategoryActionQuery)
    },
  },
  Format: "QUERY",
  Action: PageCategoryActionQuery,
  ResponseEntity: &[]PageCategoryEntity{},
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			PageCategoryActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionAliases: []string{"q"},
	Flags:         workspaces.CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var PAGE_CATEGORY_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/page-categories/export",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_CATEGORY_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, PageCategoryActionExport)
    },
  },
  Format: "QUERY",
  Action: PageCategoryActionExport,
  ResponseEntity: &[]PageCategoryEntity{},
}
var PAGE_CATEGORY_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/page-category/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_CATEGORY_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, PageCategoryActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: PageCategoryActionGetOne,
  ResponseEntity: &PageCategoryEntity{},
}
var PAGE_CATEGORY_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new pageCategory",
  Flags: PageCategoryCommonCliFlags,
  Method: "POST",
  Url:    "/page-category",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_CATEGORY_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, PageCategoryActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, PageCategoryActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: PageCategoryActionCreate,
  Format: "POST_ONE",
  RequestEntity: &PageCategoryEntity{},
  ResponseEntity: &PageCategoryEntity{},
}
var PAGE_CATEGORY_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: PageCategoryCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/page-category",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_CATEGORY_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, PageCategoryActionUpdate)
    },
  },
  Action: PageCategoryActionUpdate,
  RequestEntity: &PageCategoryEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &PageCategoryEntity{},
}
var PAGE_CATEGORY_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/page-categories",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_CATEGORY_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, PageCategoryActionBulkUpdate)
    },
  },
  Action: PageCategoryActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[PageCategoryEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[PageCategoryEntity]{},
}
var PAGE_CATEGORY_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/page-category",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_CATEGORY_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, PageCategoryActionRemove)
    },
  },
  Action: PageCategoryActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &PageCategoryEntity{},
}
  /**
  *	Override this function on PageCategoryEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendPageCategoryRouter = func(r *[]workspaces.Module2Action) {}
  func GetPageCategoryModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      PAGE_CATEGORY_ACTION_QUERY,
      PAGE_CATEGORY_ACTION_EXPORT,
      PAGE_CATEGORY_ACTION_GET_ONE,
      PAGE_CATEGORY_ACTION_POST_ONE,
      PAGE_CATEGORY_ACTION_PATCH,
      PAGE_CATEGORY_ACTION_PATCH_BULK,
      PAGE_CATEGORY_ACTION_DELETE,
    }
    // Append user defined functions
    AppendPageCategoryRouter(&routes)
    return routes
  }
  func CreatePageCategoryRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetPageCategoryModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, PageCategoryEntityJsonSchema, "page-category-http", "cms")
    workspaces.WriteEntitySchema("PageCategoryEntity", PageCategoryEntityJsonSchema, "cms")
    return httpRoutes
  }
var PERM_ROOT_PAGE_CATEGORY_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/cms/page-category/delete",
}
var PERM_ROOT_PAGE_CATEGORY_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/cms/page-category/create",
}
var PERM_ROOT_PAGE_CATEGORY_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/cms/page-category/update",
}
var PERM_ROOT_PAGE_CATEGORY_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/cms/page-category/query",
}
var PERM_ROOT_PAGE_CATEGORY = workspaces.PermissionInfo{
  CompleteKey: "root/cms/page-category/*",
}
var ALL_PAGE_CATEGORY_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_PAGE_CATEGORY_DELETE,
	PERM_ROOT_PAGE_CATEGORY_CREATE,
	PERM_ROOT_PAGE_CATEGORY_UPDATE,
	PERM_ROOT_PAGE_CATEGORY_QUERY,
	PERM_ROOT_PAGE_CATEGORY,
}