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
	mocks "github.com/torabian/fireback/modules/shop/mocks/Tag"
)
var tagSeedersFs *embed.FS = nil
func ResetTagSeeders(fs *embed.FS) {
	tagSeedersFs = fs
}
type TagEntity struct {
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
    Translations     []*TagEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    Children []*TagEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *TagEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var TagPreloadRelations []string = []string{}
var TAG_EVENT_CREATED = "tag.created"
var TAG_EVENT_UPDATED = "tag.updated"
var TAG_EVENT_DELETED = "tag.deleted"
var TAG_EVENTS = []string{
	TAG_EVENT_CREATED,
	TAG_EVENT_UPDATED,
	TAG_EVENT_DELETED,
}
type TagFieldMap struct {
		Name workspaces.TranslatedString `yaml:"name"`
}
var TagEntityMetaConfig map[string]int64 = map[string]int64{
}
var TagEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&TagEntity{}))
  type TagEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        Name string `yaml:"name" json:"name"`
  }
func entityTagFormatter(dto *TagEntity, query workspaces.QueryDSL) {
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
func TagMockEntity() *TagEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &TagEntity{
      Name : &stringHolder,
	}
	return entity
}
func TagActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := TagMockEntity()
		_, err := TagActionCreate(entity, query)
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
    func (x*TagEntity) GetNameTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Name
          }
        }
      }
      return ""
    }
  func TagActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*TagEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &TagEntity{
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
  func TagAssociationCreate(dto *TagEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func TagRelationContentCreate(dto *TagEntity, query workspaces.QueryDSL) error {
return nil
}
func TagRelationContentUpdate(dto *TagEntity, query workspaces.QueryDSL) error {
	return nil
}
func TagPolyglotCreateHandler(dto *TagEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &TagEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func TagValidator(dto *TagEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func TagEntityPreSanitize(dto *TagEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func TagEntityBeforeCreateAppend(dto *TagEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    TagRecursiveAddUniqueId(dto, query)
  }
  func TagRecursiveAddUniqueId(dto *TagEntity, query workspaces.QueryDSL) {
  }
func TagActionBatchCreateFn(dtos []*TagEntity, query workspaces.QueryDSL) ([]*TagEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*TagEntity{}
		for _, item := range dtos {
			s, err := TagActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func TagDeleteEntireChildren(query workspaces.QueryDSL, dto *TagEntity) (*workspaces.IError) {
  return nil
}
func TagActionCreateFn(dto *TagEntity, query workspaces.QueryDSL) (*TagEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := TagValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	TagEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	TagEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	TagPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	TagRelationContentCreate(dto, query)
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
	TagAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(TAG_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&TagEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func TagActionGetOne(query workspaces.QueryDSL) (*TagEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&TagEntity{})
    item, err := workspaces.GetOneEntity[TagEntity](query, refl)
    entityTagFormatter(item, query)
    return item, err
  }
  func TagActionQuery(query workspaces.QueryDSL) ([]*TagEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&TagEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[TagEntity](query, refl)
    for _, item := range items {
      entityTagFormatter(item, query)
    }
    return items, meta, err
  }
  func TagUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *TagEntity) (*TagEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = TAG_EVENT_UPDATED
    TagEntityPreSanitize(fields, query)
    var item TagEntity
    q := dbref.
      Where(&TagEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    TagRelationContentUpdate(fields, query)
    TagPolyglotCreateHandler(fields, query)
    if ero := TagDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&TagEntity{UniqueId: uniqueId}).
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
  func TagActionUpdateFn(query workspaces.QueryDSL, fields *TagEntity) (*TagEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := TagValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // TagRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *TagEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = TagUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return TagUpdateExec(dbref, query, fields)
    }
  }
var TagWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire tags ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_TAG_DELETE},
    })
		count, _ := TagActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func TagActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&TagEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_TAG_DELETE}
	return workspaces.RemoveEntity[TagEntity](query, refl)
}
func TagActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[TagEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'TagEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func TagActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[TagEntity]) (
    *workspaces.BulkRecordRequest[TagEntity], *workspaces.IError,
  ) {
    result := []*TagEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := TagActionUpdate(query, record)
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
func (x *TagEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var TagEntityMeta = workspaces.TableMetaData{
	EntityName:    "Tag",
	ExportKey:    "tags",
	TableNameInDb: "fb_tag_entities",
	EntityObject:  &TagEntity{},
	ExportStream: TagActionExportT,
	ImportQuery: TagActionImport,
}
func TagActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[TagEntity](query, TagActionQuery, TagPreloadRelations)
}
func TagActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[TagEntity](query, TagActionQuery, TagPreloadRelations)
}
func TagActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content TagEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := TagActionCreate(&content, query)
	return err
}
var TagCommonCliFlags = []cli.Flag{
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
var TagCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "name",
		StructField:     "Name",
		Required: true,
		Usage:    "name",
		Type: "string",
	},
}
var TagCommonCliFlagsOptional = []cli.Flag{
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
  var TagCreateCmd cli.Command = TAG_ACTION_POST_ONE.ToCli()
  var TagCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_TAG_CREATE},
      })
      entity := &TagEntity{}
      for _, item := range TagCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := TagActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var TagUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: TagCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_TAG_UPDATE},
      })
      entity := CastTagFromCli(c)
      if entity, err := TagActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* TagEntity) FromCli(c *cli.Context) *TagEntity {
	return CastTagFromCli(c)
}
func CastTagFromCli (c *cli.Context) *TagEntity {
	template := &TagEntity{}
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
  func TagSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      TagActionCreate,
      reflect.ValueOf(&TagEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func TagImportMocks() {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      TagActionCreate,
      reflect.ValueOf(&TagEntity{}).Elem(),
      &mocks.ViewsFs,
      []string{},
      false,
    )
  }
  func TagWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := TagActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "Tag", result)
    }
  }
var TagImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_TAG_CREATE},
      })
			TagActionSeeder(query, c.Int("count"))
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
				Value: "tag-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_TAG_CREATE},
      })
			TagActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "tag-seeder-tag.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of tags, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]TagEntity{}
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
					TagActionCreate,
					reflect.ValueOf(&TagEntity{}).Elem(),
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
			TagCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				TagActionCreate,
				reflect.ValueOf(&TagEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_TAG_CREATE},
				},
        func() TagEntity {
					v := CastTagFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var TagCliCommands []cli.Command = []cli.Command{
      TAG_ACTION_QUERY.ToCli(),
      TAG_ACTION_TABLE.ToCli(),
      TagCreateCmd,
      TagUpdateCmd,
      TagCreateInteractiveCmd,
      TagWipeCmd,
      workspaces.GetCommonRemoveQuery(reflect.ValueOf(&TagEntity{}).Elem(), TagActionRemove),
  }
  func TagCliFn() cli.Command {
    TagCliCommands = append(TagCliCommands, TagImportExportCommands...)
    return cli.Command{
      Name:        "tag",
      Description: "Tags module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: TagCliCommands,
    }
  }
var TAG_ACTION_TABLE = workspaces.Module2Action{
  Name:    "table",
  ActionName: "table",
  ActionAliases: []string{"t"},
  Flags:  workspaces.CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: TagActionQuery,
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    workspaces.CommonCliTableCmd2(c,
      TagActionQuery,
      security,
      reflect.ValueOf(&TagEntity{}).Elem(),
    )
    return nil
  },
}
var TAG_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/tags",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_TAG_QUERY},
  },
  Group: "tag",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, TagActionQuery)
    },
  },
  Format: "QUERY",
  Action: TagActionQuery,
  ResponseEntity: &[]TagEntity{},
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			TagActionQuery,
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
var TAG_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/tags/export",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_TAG_QUERY},
  },
  Group: "tag",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, TagActionExport)
    },
  },
  Format: "QUERY",
  Action: TagActionExport,
  ResponseEntity: &[]TagEntity{},
}
var TAG_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/tag/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_TAG_QUERY},
  },
  Group: "tag",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, TagActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: TagActionGetOne,
  ResponseEntity: &TagEntity{},
}
var TAG_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new tag",
  Flags: TagCommonCliFlags,
  Method: "POST",
  Url:    "/tag",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_TAG_CREATE},
  },
  Group: "tag",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, TagActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, TagActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: TagActionCreate,
  Format: "POST_ONE",
  RequestEntity: &TagEntity{},
  ResponseEntity: &TagEntity{},
}
var TAG_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: TagCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/tag",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_TAG_UPDATE},
  },
  Group: "tag",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, TagActionUpdate)
    },
  },
  Action: TagActionUpdate,
  RequestEntity: &TagEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &TagEntity{},
}
var TAG_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/tags",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_TAG_UPDATE},
  },
  Group: "tag",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, TagActionBulkUpdate)
    },
  },
  Action: TagActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[TagEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[TagEntity]{},
}
var TAG_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/tag",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_TAG_DELETE},
  },
  Group: "tag",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, TagActionRemove)
    },
  },
  Action: TagActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &TagEntity{},
}
  /**
  *	Override this function on TagEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendTagRouter = func(r *[]workspaces.Module2Action) {}
  func GetTagModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      TAG_ACTION_QUERY,
      TAG_ACTION_EXPORT,
      TAG_ACTION_GET_ONE,
      TAG_ACTION_POST_ONE,
      TAG_ACTION_PATCH,
      TAG_ACTION_PATCH_BULK,
      TAG_ACTION_DELETE,
    }
    // Append user defined functions
    AppendTagRouter(&routes)
    return routes
  }
  func CreateTagRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetTagModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, TagEntityJsonSchema, "tag-http", "shop")
    workspaces.WriteEntitySchema("TagEntity", TagEntityJsonSchema, "shop")
    return httpRoutes
  }
var PERM_ROOT_TAG_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/tag/delete",
  Name: "Delete tag",
}
var PERM_ROOT_TAG_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/tag/create",
  Name: "Create tag",
}
var PERM_ROOT_TAG_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/tag/update",
  Name: "Update tag",
}
var PERM_ROOT_TAG_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/shop/tag/query",
  Name: "Query tag",
}
var PERM_ROOT_TAG = workspaces.PermissionInfo{
  CompleteKey: "root/shop/tag/*",
  Name: "Entire tag actions (*)",
}
var ALL_TAG_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_TAG_DELETE,
	PERM_ROOT_TAG_CREATE,
	PERM_ROOT_TAG_UPDATE,
	PERM_ROOT_TAG_QUERY,
	PERM_ROOT_TAG,
}