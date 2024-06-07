package book
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
var bookSeedersFs *embed.FS = nil
func ResetBookSeeders(fs *embed.FS) {
	bookSeedersFs = fs
}
type BookEntity struct {
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
    Title   *string `json:"title" yaml:"title"  validate:"required"        translate:"true" `
    // Datenano also has a text representation
    PageCount   *int64 `json:"pageCount" yaml:"pageCount"       `
    // Datenano also has a text representation
    Isbn   *string `json:"isbn" yaml:"isbn"       `
    // Datenano also has a text representation
    Translations     []*BookEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId;constraint:OnDelete:CASCADE;"`
    Children []*BookEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *BookEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var BookPreloadRelations []string = []string{}
var BOOK_EVENT_CREATED = "book.created"
var BOOK_EVENT_UPDATED = "book.updated"
var BOOK_EVENT_DELETED = "book.deleted"
var BOOK_EVENTS = []string{
	BOOK_EVENT_CREATED,
	BOOK_EVENT_UPDATED,
	BOOK_EVENT_DELETED,
}
type BookFieldMap struct {
		Title workspaces.TranslatedString `yaml:"title"`
		PageCount workspaces.TranslatedString `yaml:"pageCount"`
		Isbn workspaces.TranslatedString `yaml:"isbn"`
}
var BookEntityMetaConfig map[string]int64 = map[string]int64{
}
var BookEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&BookEntity{}))
  type BookEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        Title string `yaml:"title" json:"title"`
  }
func entityBookFormatter(dto *BookEntity, query workspaces.QueryDSL) {
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
func BookMockEntity() *BookEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &BookEntity{
      Title : &stringHolder,
      PageCount : &int64Holder,
      Isbn : &stringHolder,
	}
	return entity
}
func BookActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := BookMockEntity()
		_, err := BookActionCreate(entity, query)
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
    func (x*BookEntity) GetTitleTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Title
          }
        }
      }
      return ""
    }
  func BookActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*BookEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &BookEntity{
          Title: &tildaRef,
          Isbn: &tildaRef,
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
  func BookAssociationCreate(dto *BookEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func BookRelationContentCreate(dto *BookEntity, query workspaces.QueryDSL) error {
return nil
}
func BookRelationContentUpdate(dto *BookEntity, query workspaces.QueryDSL) error {
	return nil
}
func BookPolyglotCreateHandler(dto *BookEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &BookEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func BookValidator(dto *BookEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func BookEntityPreSanitize(dto *BookEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func BookEntityBeforeCreateAppend(dto *BookEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    BookRecursiveAddUniqueId(dto, query)
  }
  func BookRecursiveAddUniqueId(dto *BookEntity, query workspaces.QueryDSL) {
  }
func BookActionBatchCreateFn(dtos []*BookEntity, query workspaces.QueryDSL) ([]*BookEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*BookEntity{}
		for _, item := range dtos {
			s, err := BookActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func BookDeleteEntireChildren(query workspaces.QueryDSL, dto *BookEntity) (*workspaces.IError) {
  // intentionally removed this. It's hard to implement it, and probably wrong without
  // proper on delete cascade
  return nil
}
func BookActionCreateFn(dto *BookEntity, query workspaces.QueryDSL) (*BookEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := BookValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	BookEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	BookEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	BookPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	BookRelationContentCreate(dto, query)
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
	BookAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(BOOK_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&BookEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func BookActionGetOne(query workspaces.QueryDSL) (*BookEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&BookEntity{})
    item, err := workspaces.GetOneEntity[BookEntity](query, refl)
    entityBookFormatter(item, query)
    return item, err
  }
  func BookActionQuery(query workspaces.QueryDSL) ([]*BookEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&BookEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[BookEntity](query, refl)
    for _, item := range items {
      entityBookFormatter(item, query)
    }
    return items, meta, err
  }
  func BookUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *BookEntity) (*BookEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = BOOK_EVENT_UPDATED
    BookEntityPreSanitize(fields, query)
    var item BookEntity
    q := dbref.
      Where(&BookEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    BookRelationContentUpdate(fields, query)
    BookPolyglotCreateHandler(fields, query)
    if ero := BookDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&BookEntity{UniqueId: uniqueId}).
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
  func BookActionUpdateFn(query workspaces.QueryDSL, fields *BookEntity) (*BookEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.Create401Error(&workspaces.WorkspacesMessages.BodyIsMissing, []string{})
    }
    // 1. Validate always
    if iError := BookValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // BookRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *BookEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = BookUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return BookUpdateExec(dbref, query, fields)
    }
  }
var BookWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire books ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_BOOK_DELETE},
    })
		count, _ := BookActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func BookActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&BookEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_BOOK_DELETE}
	return workspaces.RemoveEntity[BookEntity](query, refl)
}
func BookActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[BookEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'BookEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func BookActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[BookEntity]) (
    *workspaces.BulkRecordRequest[BookEntity], *workspaces.IError,
  ) {
    result := []*BookEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := BookActionUpdate(query, record)
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
func (x *BookEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var BookEntityMeta = workspaces.TableMetaData{
	EntityName:    "Book",
	ExportKey:    "books",
	TableNameInDb: "fb_book_entities",
	EntityObject:  &BookEntity{},
	ExportStream: BookActionExportT,
	ImportQuery: BookActionImport,
}
func BookActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[BookEntity](query, BookActionQuery, BookPreloadRelations)
}
func BookActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[BookEntity](query, BookActionQuery, BookPreloadRelations)
}
func BookActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content BookEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.Create401Error(&workspaces.WorkspacesMessages.InvalidContent, []string{})
	}
	json.Unmarshal(cx, &content)
	_, err := BookActionCreate(&content, query)
	return err
}
var BookCommonCliFlags = []cli.Flag{
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
      Required: true,
      Usage:    "title",
    },
    &cli.Int64Flag{
      Name:     "page-count",
      Required: false,
      Usage:    "pageCount",
    },
    &cli.StringFlag{
      Name:     "isbn",
      Required: false,
      Usage:    "isbn",
    },
}
var BookCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "title",
		StructField:     "Title",
		Required: true,
		Usage:    "title",
		Type: "string",
	},
	{
		Name:     "pageCount",
		StructField:     "PageCount",
		Required: false,
		Usage:    "pageCount",
		Type: "int64",
	},
	{
		Name:     "isbn",
		StructField:     "Isbn",
		Required: false,
		Usage:    "isbn",
		Type: "string",
	},
}
var BookCommonCliFlagsOptional = []cli.Flag{
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
      Required: true,
      Usage:    "title",
    },
    &cli.Int64Flag{
      Name:     "page-count",
      Required: false,
      Usage:    "pageCount",
    },
    &cli.StringFlag{
      Name:     "isbn",
      Required: false,
      Usage:    "isbn",
    },
}
  var BookCreateCmd cli.Command = BOOK_ACTION_POST_ONE.ToCli()
  var BookCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_BOOK_CREATE},
      })
      entity := &BookEntity{}
      for _, item := range BookCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := BookActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var BookUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: BookCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_BOOK_UPDATE},
      })
      entity := CastBookFromCli(c)
      if entity, err := BookActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* BookEntity) FromCli(c *cli.Context) *BookEntity {
	return CastBookFromCli(c)
}
func CastBookFromCli (c *cli.Context) *BookEntity {
	template := &BookEntity{}
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
      if c.IsSet("isbn") {
        value := c.String("isbn")
        template.Isbn = &value
      }
	return template
}
  func BookSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      BookActionCreate,
      reflect.ValueOf(&BookEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func BookWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := BookActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "Book", result)
    }
  }
var BookImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_BOOK_CREATE},
      })
			BookActionSeeder(query, c.Int("count"))
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
				Value: "book-seeder.yml",
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
            query := workspaces.CommonCliQueryDSLBuilder(c)
			BookActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "book-seeder-book.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of books, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]BookEntity{}
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
			BookCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				BookActionCreate,
				reflect.ValueOf(&BookEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_BOOK_CREATE},
				},
        func() BookEntity {
					v := CastBookFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var BookCliCommands []cli.Command = []cli.Command{
      BOOK_ACTION_QUERY.ToCli(),
      BOOK_ACTION_TABLE.ToCli(),
      BookCreateCmd,
      BookUpdateCmd,
      BookCreateInteractiveCmd,
      BookWipeCmd,
      workspaces.GetCommonRemoveQuery(reflect.ValueOf(&BookEntity{}).Elem(), BookActionRemove),
  }
  func BookCliFn() cli.Command {
    BookCliCommands = append(BookCliCommands, BookImportExportCommands...)
    return cli.Command{
      Name:        "book",
      Description: "Books module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: BookCliCommands,
    }
  }
var BOOK_ACTION_TABLE = workspaces.Module2Action{
  Name:    "table",
  ActionName: "table",
  ActionAliases: []string{"t"},
  Flags:  workspaces.CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: BookActionQuery,
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    workspaces.CommonCliTableCmd2(c,
      BookActionQuery,
      security,
      reflect.ValueOf(&BookEntity{}).Elem(),
    )
    return nil
  },
}
var BOOK_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/books",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_BOOK_QUERY},
  },
  Group: "book",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, BookActionQuery)
    },
  },
  Format: "QUERY",
  Action: BookActionQuery,
  ResponseEntity: &[]BookEntity{},
  Out: workspaces.Module2ActionBody{
		Entity: "BookEntity",
	},
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			BookActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionName:    "query",
	ActionAliases: []string{"q"},
	Flags:         workspaces.CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var BOOK_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/books/export",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_BOOK_QUERY},
  },
  Group: "book",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, BookActionExport)
    },
  },
  Format: "QUERY",
  Action: BookActionExport,
  ResponseEntity: &[]BookEntity{},
  Out: workspaces.Module2ActionBody{
		Entity: "BookEntity",
	},
}
var BOOK_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/book/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_BOOK_QUERY},
  },
  Group: "book",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, BookActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: BookActionGetOne,
  ResponseEntity: &BookEntity{},
  Out: workspaces.Module2ActionBody{
		Entity: "BookEntity",
	},
}
var BOOK_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new book",
  Flags: BookCommonCliFlags,
  Method: "POST",
  Url:    "/book",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_BOOK_CREATE},
  },
  Group: "book",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, BookActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, BookActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: BookActionCreate,
  Format: "POST_ONE",
  RequestEntity: &BookEntity{},
  ResponseEntity: &BookEntity{},
  Out: workspaces.Module2ActionBody{
		Entity: "BookEntity",
	},
  In: workspaces.Module2ActionBody{
		Entity: "BookEntity",
	},
}
var BOOK_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: BookCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/book",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_BOOK_UPDATE},
  },
  Group: "book",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, BookActionUpdate)
    },
  },
  Action: BookActionUpdate,
  RequestEntity: &BookEntity{},
  ResponseEntity: &BookEntity{},
  Format: "PATCH_ONE",
  Out: workspaces.Module2ActionBody{
		Entity: "BookEntity",
	},
  In: workspaces.Module2ActionBody{
		Entity: "BookEntity",
	},
}
var BOOK_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/books",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_BOOK_UPDATE},
  },
  Group: "book",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, BookActionBulkUpdate)
    },
  },
  Action: BookActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[BookEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[BookEntity]{},
  Out: workspaces.Module2ActionBody{
		Entity: "BookEntity",
	},
  In: workspaces.Module2ActionBody{
		Entity: "BookEntity",
	},
}
var BOOK_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/book",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_BOOK_DELETE},
  },
  Group: "book",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, BookActionRemove)
    },
  },
  Action: BookActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &BookEntity{},
}
  /**
  *	Override this function on BookEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendBookRouter = func(r *[]workspaces.Module2Action) {}
  func GetBookModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      BOOK_ACTION_QUERY,
      BOOK_ACTION_EXPORT,
      BOOK_ACTION_GET_ONE,
      BOOK_ACTION_POST_ONE,
      BOOK_ACTION_PATCH,
      BOOK_ACTION_PATCH_BULK,
      BOOK_ACTION_DELETE,
    }
    // Append user defined functions
    AppendBookRouter(&routes)
    return routes
  }
  func CreateBookRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetBookModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, BookEntityJsonSchema, "book-http", "book")
    workspaces.WriteEntitySchema("BookEntity", BookEntityJsonSchema, "book")
    return httpRoutes
  }
var PERM_ROOT_BOOK_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/book/book/delete",
  Name: "Delete book",
}
var PERM_ROOT_BOOK_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/book/book/create",
  Name: "Create book",
}
var PERM_ROOT_BOOK_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/book/book/update",
  Name: "Update book",
}
var PERM_ROOT_BOOK_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/book/book/query",
  Name: "Query book",
}
var PERM_ROOT_BOOK = workspaces.PermissionInfo{
  CompleteKey: "root/book/book/*",
  Name: "Entire book actions (*)",
}
var ALL_BOOK_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_BOOK_DELETE,
	PERM_ROOT_BOOK_CREATE,
	PERM_ROOT_BOOK_UPDATE,
	PERM_ROOT_BOOK_QUERY,
	PERM_ROOT_BOOK,
}
var BookEntityBundle = workspaces.EntityBundle{
	Permissions: ALL_BOOK_PERMISSIONS,
	CliCommands: []cli.Command{
		BookCliFn(),
	},
	Actions: GetBookModule2Actions(),
	AutoMigrationEntities: []interface{}{
		&BookEntity{},
  	},
}