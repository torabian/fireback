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
type PostTagEntity struct {
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
    Name   *string `json:"name" yaml:"name"  validate:"required"        translate:"true" `
    // Datenano also has a text representation
    Translations     []*PostTagEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    Children []*PostTagEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *PostTagEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var PostTagPreloadRelations []string = []string{}
var POST_TAG_EVENT_CREATED = "postTag.created"
var POST_TAG_EVENT_UPDATED = "postTag.updated"
var POST_TAG_EVENT_DELETED = "postTag.deleted"
var POST_TAG_EVENTS = []string{
	POST_TAG_EVENT_CREATED,
	POST_TAG_EVENT_UPDATED,
	POST_TAG_EVENT_DELETED,
}
type PostTagFieldMap struct {
		Name workspaces.TranslatedString `yaml:"name"`
}
var PostTagEntityMetaConfig map[string]int64 = map[string]int64{
}
var PostTagEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&PostTagEntity{}))
  type PostTagEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        Name string `yaml:"name" json:"name"`
  }
func entityPostTagFormatter(dto *PostTagEntity, query workspaces.QueryDSL) {
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
func PostTagMockEntity() *PostTagEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &PostTagEntity{
      Name : &stringHolder,
	}
	return entity
}
func PostTagActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := PostTagMockEntity()
		_, err := PostTagActionCreate(entity, query)
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
    func (x*PostTagEntity) GetNameTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Name
          }
        }
      }
      return ""
    }
  func PostTagActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*PostTagEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &PostTagEntity{
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
  func PostTagAssociationCreate(dto *PostTagEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func PostTagRelationContentCreate(dto *PostTagEntity, query workspaces.QueryDSL) error {
return nil
}
func PostTagRelationContentUpdate(dto *PostTagEntity, query workspaces.QueryDSL) error {
	return nil
}
func PostTagPolyglotCreateHandler(dto *PostTagEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &PostTagEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func PostTagValidator(dto *PostTagEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func PostTagEntityPreSanitize(dto *PostTagEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func PostTagEntityBeforeCreateAppend(dto *PostTagEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    PostTagRecursiveAddUniqueId(dto, query)
  }
  func PostTagRecursiveAddUniqueId(dto *PostTagEntity, query workspaces.QueryDSL) {
  }
func PostTagActionBatchCreateFn(dtos []*PostTagEntity, query workspaces.QueryDSL) ([]*PostTagEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*PostTagEntity{}
		for _, item := range dtos {
			s, err := PostTagActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func PostTagDeleteEntireChildren(query workspaces.QueryDSL, dto *PostTagEntity) (*workspaces.IError) {
  return nil
}
func PostTagActionCreateFn(dto *PostTagEntity, query workspaces.QueryDSL) (*PostTagEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := PostTagValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	PostTagEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	PostTagEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	PostTagPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	PostTagRelationContentCreate(dto, query)
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
	PostTagAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(POST_TAG_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&PostTagEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func PostTagActionGetOne(query workspaces.QueryDSL) (*PostTagEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&PostTagEntity{})
    item, err := workspaces.GetOneEntity[PostTagEntity](query, refl)
    entityPostTagFormatter(item, query)
    return item, err
  }
  func PostTagActionQuery(query workspaces.QueryDSL) ([]*PostTagEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&PostTagEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[PostTagEntity](query, refl)
    for _, item := range items {
      entityPostTagFormatter(item, query)
    }
    return items, meta, err
  }
  func PostTagUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *PostTagEntity) (*PostTagEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = POST_TAG_EVENT_UPDATED
    PostTagEntityPreSanitize(fields, query)
    var item PostTagEntity
    q := dbref.
      Where(&PostTagEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    PostTagRelationContentUpdate(fields, query)
    PostTagPolyglotCreateHandler(fields, query)
    if ero := PostTagDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&PostTagEntity{UniqueId: uniqueId}).
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
  func PostTagActionUpdateFn(query workspaces.QueryDSL, fields *PostTagEntity) (*PostTagEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := PostTagValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // PostTagRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *PostTagEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = PostTagUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return PostTagUpdateExec(dbref, query, fields)
    }
  }
var PostTagWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire posttags ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_TAG_DELETE},
    })
		count, _ := PostTagActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func PostTagActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&PostTagEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_POST_TAG_DELETE}
	return workspaces.RemoveEntity[PostTagEntity](query, refl)
}
func PostTagActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[PostTagEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'PostTagEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func PostTagActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[PostTagEntity]) (
    *workspaces.BulkRecordRequest[PostTagEntity], *workspaces.IError,
  ) {
    result := []*PostTagEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := PostTagActionUpdate(query, record)
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
func (x *PostTagEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var PostTagEntityMeta = workspaces.TableMetaData{
	EntityName:    "PostTag",
	ExportKey:    "post-tags",
	TableNameInDb: "fb_post-tag_entities",
	EntityObject:  &PostTagEntity{},
	ExportStream: PostTagActionExportT,
	ImportQuery: PostTagActionImport,
}
func PostTagActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[PostTagEntity](query, PostTagActionQuery, PostTagPreloadRelations)
}
func PostTagActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[PostTagEntity](query, PostTagActionQuery, PostTagPreloadRelations)
}
func PostTagActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content PostTagEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := PostTagActionCreate(&content, query)
	return err
}
var PostTagCommonCliFlags = []cli.Flag{
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
var PostTagCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "name",
		StructField:     "Name",
		Required: true,
		Usage:    "name",
		Type: "string",
	},
}
var PostTagCommonCliFlagsOptional = []cli.Flag{
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
  var PostTagCreateCmd cli.Command = POST_TAG_ACTION_POST_ONE.ToCli()
  var PostTagCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_TAG_CREATE},
      })
      entity := &PostTagEntity{}
      for _, item := range PostTagCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := PostTagActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var PostTagUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: PostTagCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_TAG_UPDATE},
      })
      entity := CastPostTagFromCli(c)
      if entity, err := PostTagActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* PostTagEntity) FromCli(c *cli.Context) *PostTagEntity {
	return CastPostTagFromCli(c)
}
func CastPostTagFromCli (c *cli.Context) *PostTagEntity {
	template := &PostTagEntity{}
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
  func PostTagSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      PostTagActionCreate,
      reflect.ValueOf(&PostTagEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func PostTagWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := PostTagActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "PostTag", result)
    }
  }
var PostTagImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_TAG_CREATE},
      })
			PostTagActionSeeder(query, c.Int("count"))
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
				Value: "post-tag-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_TAG_CREATE},
      })
			PostTagActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "post-tag-seeder-post-tag.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of post-tags, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]PostTagEntity{}
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
			PostTagCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				PostTagActionCreate,
				reflect.ValueOf(&PostTagEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_TAG_CREATE},
				},
        func() PostTagEntity {
					v := CastPostTagFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var PostTagCliCommands []cli.Command = []cli.Command{
      POST_TAG_ACTION_QUERY.ToCli(),
      POST_TAG_ACTION_TABLE.ToCli(),
      PostTagCreateCmd,
      PostTagUpdateCmd,
      PostTagCreateInteractiveCmd,
      PostTagWipeCmd,
      workspaces.GetCommonRemoveQuery(reflect.ValueOf(&PostTagEntity{}).Elem(), PostTagActionRemove),
  }
  func PostTagCliFn() cli.Command {
    PostTagCliCommands = append(PostTagCliCommands, PostTagImportExportCommands...)
    return cli.Command{
      Name:        "postTag",
      Description: "PostTags module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: PostTagCliCommands,
    }
  }
var POST_TAG_ACTION_TABLE = workspaces.Module2Action{
  Name:    "table",
  ActionAliases: []string{"t"},
  Flags:  workspaces.CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: PostTagActionQuery,
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    workspaces.CommonCliTableCmd2(c,
      PostTagActionQuery,
      security,
      reflect.ValueOf(&PostTagEntity{}).Elem(),
    )
    return nil
  },
}
var POST_TAG_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/post-tags",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_TAG_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, PostTagActionQuery)
    },
  },
  Format: "QUERY",
  Action: PostTagActionQuery,
  ResponseEntity: &[]PostTagEntity{},
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			PostTagActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionAliases: []string{"q"},
	Flags:         workspaces.CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var POST_TAG_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/post-tags/export",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_TAG_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, PostTagActionExport)
    },
  },
  Format: "QUERY",
  Action: PostTagActionExport,
  ResponseEntity: &[]PostTagEntity{},
}
var POST_TAG_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/post-tag/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_TAG_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, PostTagActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: PostTagActionGetOne,
  ResponseEntity: &PostTagEntity{},
}
var POST_TAG_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new postTag",
  Flags: PostTagCommonCliFlags,
  Method: "POST",
  Url:    "/post-tag",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_TAG_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, PostTagActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, PostTagActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: PostTagActionCreate,
  Format: "POST_ONE",
  RequestEntity: &PostTagEntity{},
  ResponseEntity: &PostTagEntity{},
}
var POST_TAG_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: PostTagCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/post-tag",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_TAG_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, PostTagActionUpdate)
    },
  },
  Action: PostTagActionUpdate,
  RequestEntity: &PostTagEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &PostTagEntity{},
}
var POST_TAG_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/post-tags",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_TAG_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, PostTagActionBulkUpdate)
    },
  },
  Action: PostTagActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[PostTagEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[PostTagEntity]{},
}
var POST_TAG_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/post-tag",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_TAG_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, PostTagActionRemove)
    },
  },
  Action: PostTagActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &PostTagEntity{},
}
  /**
  *	Override this function on PostTagEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendPostTagRouter = func(r *[]workspaces.Module2Action) {}
  func GetPostTagModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      POST_TAG_ACTION_QUERY,
      POST_TAG_ACTION_EXPORT,
      POST_TAG_ACTION_GET_ONE,
      POST_TAG_ACTION_POST_ONE,
      POST_TAG_ACTION_PATCH,
      POST_TAG_ACTION_PATCH_BULK,
      POST_TAG_ACTION_DELETE,
    }
    // Append user defined functions
    AppendPostTagRouter(&routes)
    return routes
  }
  func CreatePostTagRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetPostTagModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, PostTagEntityJsonSchema, "post-tag-http", "cms")
    workspaces.WriteEntitySchema("PostTagEntity", PostTagEntityJsonSchema, "cms")
    return httpRoutes
  }
var PERM_ROOT_POST_TAG_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/cms/post-tag/delete",
}
var PERM_ROOT_POST_TAG_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/cms/post-tag/create",
}
var PERM_ROOT_POST_TAG_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/cms/post-tag/update",
}
var PERM_ROOT_POST_TAG_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/cms/post-tag/query",
}
var PERM_ROOT_POST_TAG = workspaces.PermissionInfo{
  CompleteKey: "root/cms/post-tag/*",
}
var ALL_POST_TAG_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_POST_TAG_DELETE,
	PERM_ROOT_POST_TAG_CREATE,
	PERM_ROOT_POST_TAG_UPDATE,
	PERM_ROOT_POST_TAG_QUERY,
	PERM_ROOT_POST_TAG,
}