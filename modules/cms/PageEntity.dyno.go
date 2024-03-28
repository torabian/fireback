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
type PageEntity struct {
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
    Title   *string `json:"title" yaml:"title"       `
    // Datenano also has a text representation
    Content   *string `json:"content" yaml:"content"       `
    // Datenano also has a text representation
    ContentExcerpt * string `json:"contentExcerpt" yaml:"contentExcerpt"`
    Category   *  PageCategoryEntity `json:"category" yaml:"category"    gorm:"foreignKey:CategoryId;references:UniqueId"     `
    // Datenano also has a text representation
        CategoryId *string `json:"categoryId" yaml:"categoryId"`
    Tags   []*  PageTagEntity `json:"tags" yaml:"tags"    gorm:"many2many:page_tags;foreignKey:UniqueId;references:UniqueId"     `
    // Datenano also has a text representation
    TagsListId []string `json:"tagsListId" yaml:"tagsListId" gorm:"-" sql:"-"`
    Children []*PageEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *PageEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var PagePreloadRelations []string = []string{}
var PAGE_EVENT_CREATED = "page.created"
var PAGE_EVENT_UPDATED = "page.updated"
var PAGE_EVENT_DELETED = "page.deleted"
var PAGE_EVENTS = []string{
	PAGE_EVENT_CREATED,
	PAGE_EVENT_UPDATED,
	PAGE_EVENT_DELETED,
}
type PageFieldMap struct {
		Title workspaces.TranslatedString `yaml:"title"`
		Content workspaces.TranslatedString `yaml:"content"`
		Category workspaces.TranslatedString `yaml:"category"`
		Tags workspaces.TranslatedString `yaml:"tags"`
}
var PageEntityMetaConfig map[string]int64 = map[string]int64{
            "ContentExcerptSize": 100,
}
var PageEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&PageEntity{}))
func entityPageFormatter(dto *PageEntity, query workspaces.QueryDSL) {
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
func PageMockEntity() *PageEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &PageEntity{
      Title : &stringHolder,
	}
	return entity
}
func PageActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := PageMockEntity()
		_, err := PageActionCreate(entity, query)
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
  func PageActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*PageEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &PageEntity{
          Title: &tildaRef,
          TagsListId: []string{"~"},
          Tags: []*PageTagEntity{{}},
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
  func PageAssociationCreate(dto *PageEntity, query workspaces.QueryDSL) error {
      {
        if dto.TagsListId != nil && len(dto.TagsListId) > 0 {
          var items []PageTagEntity
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
func PageRelationContentCreate(dto *PageEntity, query workspaces.QueryDSL) error {
return nil
}
func PageRelationContentUpdate(dto *PageEntity, query workspaces.QueryDSL) error {
	return nil
}
func PagePolyglotCreateHandler(dto *PageEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func PageValidator(dto *PageEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func PageEntityPreSanitize(dto *PageEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
			if (dto.Content != nil ) {
          Content := *dto.Content
          ContentExcerpt := stripPolicy.Sanitize(*dto.Content)
            Content = ugcPolicy.Sanitize(Content)
            ContentExcerpt = stripPolicy.Sanitize(ContentExcerpt)
        ContentExcerptSize, ContentExcerptSizeExists := PageEntityMetaConfig["ContentExcerptSize"]
        if ContentExcerptSizeExists {
          ContentExcerpt = workspaces.PickFirstNWords(ContentExcerpt, int(ContentExcerptSize))
        } else {
          ContentExcerpt = workspaces.PickFirstNWords(ContentExcerpt, 30)
        }
        dto.ContentExcerpt = &ContentExcerpt
        dto.Content = &Content
      }
}
  func PageEntityBeforeCreateAppend(dto *PageEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    PageRecursiveAddUniqueId(dto, query)
  }
  func PageRecursiveAddUniqueId(dto *PageEntity, query workspaces.QueryDSL) {
  }
func PageActionBatchCreateFn(dtos []*PageEntity, query workspaces.QueryDSL) ([]*PageEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*PageEntity{}
		for _, item := range dtos {
			s, err := PageActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func PageDeleteEntireChildren(query workspaces.QueryDSL, dto *PageEntity) (*workspaces.IError) {
  return nil
}
func PageActionCreateFn(dto *PageEntity, query workspaces.QueryDSL) (*PageEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := PageValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	PageEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	PageEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	PagePolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	PageRelationContentCreate(dto, query)
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
	PageAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(PAGE_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&PageEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func PageActionGetOne(query workspaces.QueryDSL) (*PageEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&PageEntity{})
    item, err := workspaces.GetOneEntity[PageEntity](query, refl)
    entityPageFormatter(item, query)
    return item, err
  }
  func PageActionQuery(query workspaces.QueryDSL) ([]*PageEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&PageEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[PageEntity](query, refl)
    for _, item := range items {
      entityPageFormatter(item, query)
    }
    return items, meta, err
  }
  func PageUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *PageEntity) (*PageEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = PAGE_EVENT_UPDATED
    PageEntityPreSanitize(fields, query)
    var item PageEntity
    q := dbref.
      Where(&PageEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    PageRelationContentUpdate(fields, query)
    PagePolyglotCreateHandler(fields, query)
    if ero := PageDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
        if fields.TagsListId  != nil {
          var items []PageTagEntity
          if len(fields.TagsListId ) > 0 {
            dbref.
              Where(&fields.TagsListId ).
              Find(&items)
          }
          dbref.
            Model(&PageEntity{UniqueId: uniqueId}).
            Association("Tags").
            Replace(&items)
        }
    err = dbref.
      Preload(clause.Associations).
      Where(&PageEntity{UniqueId: uniqueId}).
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
  func PageActionUpdateFn(query workspaces.QueryDSL, fields *PageEntity) (*PageEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := PageValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // PageRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *PageEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = PageUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return PageUpdateExec(dbref, query, fields)
    }
  }
var PageWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire pages ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_DELETE},
    })
		count, _ := PageActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func PageActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&PageEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_PAGE_DELETE}
	return workspaces.RemoveEntity[PageEntity](query, refl)
}
func PageActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[PageEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'PageEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func PageActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[PageEntity]) (
    *workspaces.BulkRecordRequest[PageEntity], *workspaces.IError,
  ) {
    result := []*PageEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := PageActionUpdate(query, record)
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
func (x *PageEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var PageEntityMeta = workspaces.TableMetaData{
	EntityName:    "Page",
	ExportKey:    "pages",
	TableNameInDb: "fb_page_entities",
	EntityObject:  &PageEntity{},
	ExportStream: PageActionExportT,
	ImportQuery: PageActionImport,
}
func PageActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[PageEntity](query, PageActionQuery, PagePreloadRelations)
}
func PageActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[PageEntity](query, PageActionQuery, PagePreloadRelations)
}
func PageActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content PageEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := PageActionCreate(&content, query)
	return err
}
var PageCommonCliFlags = []cli.Flag{
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
      Name:     "title",
      Required: false,
      Usage:    "title",
    },
    &cli.StringFlag{
      Name:     "content",
      Required: false,
      Usage:    "content",
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
var PageCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "title",
		StructField:     "Title",
		Required: false,
		Usage:    "title",
		Type: "string",
	},
}
var PageCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "title",
      Required: false,
      Usage:    "title",
    },
    &cli.StringFlag{
      Name:     "content",
      Required: false,
      Usage:    "content",
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
  var PageCreateCmd cli.Command = PAGE_ACTION_POST_ONE.ToCli()
  var PageCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_CREATE},
      })
      entity := &PageEntity{}
      for _, item := range PageCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := PageActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var PageUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: PageCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_UPDATE},
      })
      entity := CastPageFromCli(c)
      if entity, err := PageActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* PageEntity) FromCli(c *cli.Context) *PageEntity {
	return CastPageFromCli(c)
}
func CastPageFromCli (c *cli.Context) *PageEntity {
	template := &PageEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("title") {
        value := c.String("title")
        template.Title = &value
      }
      if c.IsSet("content") {
        value := c.String("content")
        template.Content = &value
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
  func PageSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      PageActionCreate,
      reflect.ValueOf(&PageEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func PageWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := PageActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "Page", result)
    }
  }
var PageImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_CREATE},
      })
			PageActionSeeder(query, c.Int("count"))
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
				Value: "page-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_CREATE},
      })
			PageActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "page-seeder-page.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of pages, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]PageEntity{}
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
			PageCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				PageActionCreate,
				reflect.ValueOf(&PageEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_CREATE},
				},
        func() PageEntity {
					v := CastPageFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var PageCliCommands []cli.Command = []cli.Command{
      PAGE_ACTION_QUERY.ToCli(),
      PAGE_ACTION_TABLE.ToCli(),
      PageCreateCmd,
      PageUpdateCmd,
      PageCreateInteractiveCmd,
      PageWipeCmd,
      workspaces.GetCommonRemoveQuery(reflect.ValueOf(&PageEntity{}).Elem(), PageActionRemove),
  }
  func PageCliFn() cli.Command {
    PageCliCommands = append(PageCliCommands, PageImportExportCommands...)
    return cli.Command{
      Name:        "page",
      Description: "Pages module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: PageCliCommands,
    }
  }
var PAGE_ACTION_TABLE = workspaces.Module2Action{
  Name:    "table",
  ActionAliases: []string{"t"},
  Flags:  workspaces.CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: PageActionQuery,
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    workspaces.CommonCliTableCmd2(c,
      PageActionQuery,
      security,
      reflect.ValueOf(&PageEntity{}).Elem(),
    )
    return nil
  },
}
var PAGE_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/pages",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, PageActionQuery)
    },
  },
  Format: "QUERY",
  Action: PageActionQuery,
  ResponseEntity: &[]PageEntity{},
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			PageActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionAliases: []string{"q"},
	Flags:         workspaces.CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var PAGE_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/pages/export",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, PageActionExport)
    },
  },
  Format: "QUERY",
  Action: PageActionExport,
  ResponseEntity: &[]PageEntity{},
}
var PAGE_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/page/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, PageActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: PageActionGetOne,
  ResponseEntity: &PageEntity{},
}
var PAGE_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new page",
  Flags: PageCommonCliFlags,
  Method: "POST",
  Url:    "/page",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, PageActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, PageActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: PageActionCreate,
  Format: "POST_ONE",
  RequestEntity: &PageEntity{},
  ResponseEntity: &PageEntity{},
}
var PAGE_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: PageCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/page",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, PageActionUpdate)
    },
  },
  Action: PageActionUpdate,
  RequestEntity: &PageEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &PageEntity{},
}
var PAGE_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/pages",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, PageActionBulkUpdate)
    },
  },
  Action: PageActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[PageEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[PageEntity]{},
}
var PAGE_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/page",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, PageActionRemove)
    },
  },
  Action: PageActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &PageEntity{},
}
  /**
  *	Override this function on PageEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendPageRouter = func(r *[]workspaces.Module2Action) {}
  func GetPageModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      PAGE_ACTION_QUERY,
      PAGE_ACTION_EXPORT,
      PAGE_ACTION_GET_ONE,
      PAGE_ACTION_POST_ONE,
      PAGE_ACTION_PATCH,
      PAGE_ACTION_PATCH_BULK,
      PAGE_ACTION_DELETE,
    }
    // Append user defined functions
    AppendPageRouter(&routes)
    return routes
  }
  func CreatePageRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetPageModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, PageEntityJsonSchema, "page-http", "cms")
    workspaces.WriteEntitySchema("PageEntity", PageEntityJsonSchema, "cms")
    return httpRoutes
  }
var PERM_ROOT_PAGE_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/cms/page/delete",
  Name: "Delete page",
}
var PERM_ROOT_PAGE_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/cms/page/create",
  Name: "Create page",
}
var PERM_ROOT_PAGE_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/cms/page/update",
  Name: "Update page",
}
var PERM_ROOT_PAGE_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/cms/page/query",
  Name: "Query page",
}
var PERM_ROOT_PAGE = workspaces.PermissionInfo{
  CompleteKey: "root/cms/page/*",
  Name: "Entire page actions (*)",
}
var ALL_PAGE_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_PAGE_DELETE,
	PERM_ROOT_PAGE_CREATE,
	PERM_ROOT_PAGE_UPDATE,
	PERM_ROOT_PAGE_QUERY,
	PERM_ROOT_PAGE,
}