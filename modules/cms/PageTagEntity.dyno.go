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
var pageTagSeedersFs *embed.FS = nil
func ResetPageTagSeeders(fs *embed.FS) {
	pageTagSeedersFs = fs
}
type PageTagEntity struct {
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
    Translations     []*PageTagEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    Children []*PageTagEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *PageTagEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var PageTagPreloadRelations []string = []string{}
var PAGE_TAG_EVENT_CREATED = "pageTag.created"
var PAGE_TAG_EVENT_UPDATED = "pageTag.updated"
var PAGE_TAG_EVENT_DELETED = "pageTag.deleted"
var PAGE_TAG_EVENTS = []string{
	PAGE_TAG_EVENT_CREATED,
	PAGE_TAG_EVENT_UPDATED,
	PAGE_TAG_EVENT_DELETED,
}
type PageTagFieldMap struct {
		Name workspaces.TranslatedString `yaml:"name"`
}
var PageTagEntityMetaConfig map[string]int64 = map[string]int64{
}
var PageTagEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&PageTagEntity{}))
  type PageTagEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        Name string `yaml:"name" json:"name"`
  }
func entityPageTagFormatter(dto *PageTagEntity, query workspaces.QueryDSL) {
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
func PageTagMockEntity() *PageTagEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &PageTagEntity{
      Name : &stringHolder,
	}
	return entity
}
func PageTagActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := PageTagMockEntity()
		_, err := PageTagActionCreate(entity, query)
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
    func (x*PageTagEntity) GetNameTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Name
          }
        }
      }
      return ""
    }
  func PageTagActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*PageTagEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &PageTagEntity{
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
  func PageTagAssociationCreate(dto *PageTagEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func PageTagRelationContentCreate(dto *PageTagEntity, query workspaces.QueryDSL) error {
return nil
}
func PageTagRelationContentUpdate(dto *PageTagEntity, query workspaces.QueryDSL) error {
	return nil
}
func PageTagPolyglotCreateHandler(dto *PageTagEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &PageTagEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func PageTagValidator(dto *PageTagEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func PageTagEntityPreSanitize(dto *PageTagEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func PageTagEntityBeforeCreateAppend(dto *PageTagEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    PageTagRecursiveAddUniqueId(dto, query)
  }
  func PageTagRecursiveAddUniqueId(dto *PageTagEntity, query workspaces.QueryDSL) {
  }
func PageTagActionBatchCreateFn(dtos []*PageTagEntity, query workspaces.QueryDSL) ([]*PageTagEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*PageTagEntity{}
		for _, item := range dtos {
			s, err := PageTagActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func PageTagDeleteEntireChildren(query workspaces.QueryDSL, dto *PageTagEntity) (*workspaces.IError) {
  // intentionally removed this. It's hard to implement it, and probably wrong without
  // proper on delete cascade
  return nil
}
func PageTagActionCreateFn(dto *PageTagEntity, query workspaces.QueryDSL) (*PageTagEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := PageTagValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	PageTagEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	PageTagEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	PageTagPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	PageTagRelationContentCreate(dto, query)
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
	PageTagAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(PAGE_TAG_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&PageTagEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func PageTagActionGetOne(query workspaces.QueryDSL) (*PageTagEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&PageTagEntity{})
    item, err := workspaces.GetOneEntity[PageTagEntity](query, refl)
    entityPageTagFormatter(item, query)
    return item, err
  }
  func PageTagActionQuery(query workspaces.QueryDSL) ([]*PageTagEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&PageTagEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[PageTagEntity](query, refl)
    for _, item := range items {
      entityPageTagFormatter(item, query)
    }
    return items, meta, err
  }
  func PageTagUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *PageTagEntity) (*PageTagEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = PAGE_TAG_EVENT_UPDATED
    PageTagEntityPreSanitize(fields, query)
    var item PageTagEntity
    q := dbref.
      Where(&PageTagEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    PageTagRelationContentUpdate(fields, query)
    PageTagPolyglotCreateHandler(fields, query)
    if ero := PageTagDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&PageTagEntity{UniqueId: uniqueId}).
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
  func PageTagActionUpdateFn(query workspaces.QueryDSL, fields *PageTagEntity) (*PageTagEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := PageTagValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // PageTagRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *PageTagEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = PageTagUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return PageTagUpdateExec(dbref, query, fields)
    }
  }
var PageTagWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire pagetags ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_TAG_DELETE},
    })
		count, _ := PageTagActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func PageTagActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&PageTagEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_PAGE_TAG_DELETE}
	return workspaces.RemoveEntity[PageTagEntity](query, refl)
}
func PageTagActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[PageTagEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'PageTagEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func PageTagActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[PageTagEntity]) (
    *workspaces.BulkRecordRequest[PageTagEntity], *workspaces.IError,
  ) {
    result := []*PageTagEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := PageTagActionUpdate(query, record)
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
func (x *PageTagEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var PageTagEntityMeta = workspaces.TableMetaData{
	EntityName:    "PageTag",
	ExportKey:    "page-tags",
	TableNameInDb: "fb_page-tag_entities",
	EntityObject:  &PageTagEntity{},
	ExportStream: PageTagActionExportT,
	ImportQuery: PageTagActionImport,
}
func PageTagActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[PageTagEntity](query, PageTagActionQuery, PageTagPreloadRelations)
}
func PageTagActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[PageTagEntity](query, PageTagActionQuery, PageTagPreloadRelations)
}
func PageTagActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content PageTagEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := PageTagActionCreate(&content, query)
	return err
}
var PageTagCommonCliFlags = []cli.Flag{
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
var PageTagCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "name",
		StructField:     "Name",
		Required: true,
		Usage:    "name",
		Type: "string",
	},
}
var PageTagCommonCliFlagsOptional = []cli.Flag{
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
  var PageTagCreateCmd cli.Command = PAGE_TAG_ACTION_POST_ONE.ToCli()
  var PageTagCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_TAG_CREATE},
      })
      entity := &PageTagEntity{}
      for _, item := range PageTagCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := PageTagActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var PageTagUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: PageTagCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_TAG_UPDATE},
      })
      entity := CastPageTagFromCli(c)
      if entity, err := PageTagActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* PageTagEntity) FromCli(c *cli.Context) *PageTagEntity {
	return CastPageTagFromCli(c)
}
func CastPageTagFromCli (c *cli.Context) *PageTagEntity {
	template := &PageTagEntity{}
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
  func PageTagSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      PageTagActionCreate,
      reflect.ValueOf(&PageTagEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func PageTagWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := PageTagActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "PageTag", result)
    }
  }
var PageTagImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_TAG_CREATE},
      })
			PageTagActionSeeder(query, c.Int("count"))
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
				Value: "page-tag-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_TAG_CREATE},
      })
			PageTagActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "page-tag-seeder-page-tag.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of page-tags, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]PageTagEntity{}
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
			PageTagCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				PageTagActionCreate,
				reflect.ValueOf(&PageTagEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_TAG_CREATE},
				},
        func() PageTagEntity {
					v := CastPageTagFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var PageTagCliCommands []cli.Command = []cli.Command{
      PAGE_TAG_ACTION_QUERY.ToCli(),
      PAGE_TAG_ACTION_TABLE.ToCli(),
      PageTagCreateCmd,
      PageTagUpdateCmd,
      PageTagCreateInteractiveCmd,
      PageTagWipeCmd,
      workspaces.GetCommonRemoveQuery(reflect.ValueOf(&PageTagEntity{}).Elem(), PageTagActionRemove),
  }
  func PageTagCliFn() cli.Command {
    PageTagCliCommands = append(PageTagCliCommands, PageTagImportExportCommands...)
    return cli.Command{
      Name:        "pageTag",
      Description: "PageTags module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: PageTagCliCommands,
    }
  }
var PAGE_TAG_ACTION_TABLE = workspaces.Module2Action{
  Name:    "table",
  ActionName: "table",
  ActionAliases: []string{"t"},
  Flags:  workspaces.CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: PageTagActionQuery,
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    workspaces.CommonCliTableCmd2(c,
      PageTagActionQuery,
      security,
      reflect.ValueOf(&PageTagEntity{}).Elem(),
    )
    return nil
  },
}
var PAGE_TAG_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/page-tags",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_TAG_QUERY},
  },
  Group: "pageTag",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, PageTagActionQuery)
    },
  },
  Format: "QUERY",
  Action: PageTagActionQuery,
  ResponseEntity: &[]PageTagEntity{},
  Out: workspaces.Module2ActionBody{
		Entity: "PageTagEntity",
	},
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			PageTagActionQuery,
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
var PAGE_TAG_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/page-tags/export",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_TAG_QUERY},
  },
  Group: "pageTag",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, PageTagActionExport)
    },
  },
  Format: "QUERY",
  Action: PageTagActionExport,
  ResponseEntity: &[]PageTagEntity{},
  Out: workspaces.Module2ActionBody{
		Entity: "PageTagEntity",
	},
}
var PAGE_TAG_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/page-tag/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_TAG_QUERY},
  },
  Group: "pageTag",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, PageTagActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: PageTagActionGetOne,
  ResponseEntity: &PageTagEntity{},
  Out: workspaces.Module2ActionBody{
		Entity: "PageTagEntity",
	},
}
var PAGE_TAG_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new pageTag",
  Flags: PageTagCommonCliFlags,
  Method: "POST",
  Url:    "/page-tag",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_TAG_CREATE},
  },
  Group: "pageTag",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, PageTagActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, PageTagActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: PageTagActionCreate,
  Format: "POST_ONE",
  RequestEntity: &PageTagEntity{},
  ResponseEntity: &PageTagEntity{},
  Out: workspaces.Module2ActionBody{
		Entity: "PageTagEntity",
	},
  In: workspaces.Module2ActionBody{
		Entity: "PageTagEntity",
	},
}
var PAGE_TAG_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: PageTagCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/page-tag",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_TAG_UPDATE},
  },
  Group: "pageTag",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, PageTagActionUpdate)
    },
  },
  Action: PageTagActionUpdate,
  RequestEntity: &PageTagEntity{},
  ResponseEntity: &PageTagEntity{},
  Format: "PATCH_ONE",
  Out: workspaces.Module2ActionBody{
		Entity: "PageTagEntity",
	},
  In: workspaces.Module2ActionBody{
		Entity: "PageTagEntity",
	},
}
var PAGE_TAG_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/page-tags",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_TAG_UPDATE},
  },
  Group: "pageTag",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, PageTagActionBulkUpdate)
    },
  },
  Action: PageTagActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[PageTagEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[PageTagEntity]{},
  Out: workspaces.Module2ActionBody{
		Entity: "PageTagEntity",
	},
  In: workspaces.Module2ActionBody{
		Entity: "PageTagEntity",
	},
}
var PAGE_TAG_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/page-tag",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAGE_TAG_DELETE},
  },
  Group: "pageTag",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, PageTagActionRemove)
    },
  },
  Action: PageTagActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &PageTagEntity{},
}
  /**
  *	Override this function on PageTagEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendPageTagRouter = func(r *[]workspaces.Module2Action) {}
  func GetPageTagModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      PAGE_TAG_ACTION_QUERY,
      PAGE_TAG_ACTION_EXPORT,
      PAGE_TAG_ACTION_GET_ONE,
      PAGE_TAG_ACTION_POST_ONE,
      PAGE_TAG_ACTION_PATCH,
      PAGE_TAG_ACTION_PATCH_BULK,
      PAGE_TAG_ACTION_DELETE,
    }
    // Append user defined functions
    AppendPageTagRouter(&routes)
    return routes
  }
  func CreatePageTagRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetPageTagModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, PageTagEntityJsonSchema, "page-tag-http", "cms")
    workspaces.WriteEntitySchema("PageTagEntity", PageTagEntityJsonSchema, "cms")
    return httpRoutes
  }
var PERM_ROOT_PAGE_TAG_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/cms/page-tag/delete",
  Name: "Delete page tag",
}
var PERM_ROOT_PAGE_TAG_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/cms/page-tag/create",
  Name: "Create page tag",
}
var PERM_ROOT_PAGE_TAG_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/cms/page-tag/update",
  Name: "Update page tag",
}
var PERM_ROOT_PAGE_TAG_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/cms/page-tag/query",
  Name: "Query page tag",
}
var PERM_ROOT_PAGE_TAG = workspaces.PermissionInfo{
  CompleteKey: "root/cms/page-tag/*",
  Name: "Entire page tag actions (*)",
}
var ALL_PAGE_TAG_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_PAGE_TAG_DELETE,
	PERM_ROOT_PAGE_TAG_CREATE,
	PERM_ROOT_PAGE_TAG_UPDATE,
	PERM_ROOT_PAGE_TAG_QUERY,
	PERM_ROOT_PAGE_TAG,
}