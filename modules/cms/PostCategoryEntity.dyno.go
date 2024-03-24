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
type PostCategoryEntity struct {
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
    Translations     []*PostCategoryEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    Children []*PostCategoryEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *PostCategoryEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var PostCategoryPreloadRelations []string = []string{}
var POST_CATEGORY_EVENT_CREATED = "postCategory.created"
var POST_CATEGORY_EVENT_UPDATED = "postCategory.updated"
var POST_CATEGORY_EVENT_DELETED = "postCategory.deleted"
var POST_CATEGORY_EVENTS = []string{
	POST_CATEGORY_EVENT_CREATED,
	POST_CATEGORY_EVENT_UPDATED,
	POST_CATEGORY_EVENT_DELETED,
}
type PostCategoryFieldMap struct {
		Name workspaces.TranslatedString `yaml:"name"`
}
var PostCategoryEntityMetaConfig map[string]int64 = map[string]int64{
}
var PostCategoryEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&PostCategoryEntity{}))
  type PostCategoryEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        Name string `yaml:"name" json:"name"`
  }
func entityPostCategoryFormatter(dto *PostCategoryEntity, query workspaces.QueryDSL) {
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
func PostCategoryMockEntity() *PostCategoryEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &PostCategoryEntity{
      Name : &stringHolder,
	}
	return entity
}
func PostCategoryActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := PostCategoryMockEntity()
		_, err := PostCategoryActionCreate(entity, query)
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
    func (x*PostCategoryEntity) GetNameTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Name
          }
        }
      }
      return ""
    }
  func PostCategoryActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*PostCategoryEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &PostCategoryEntity{
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
  func PostCategoryAssociationCreate(dto *PostCategoryEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func PostCategoryRelationContentCreate(dto *PostCategoryEntity, query workspaces.QueryDSL) error {
return nil
}
func PostCategoryRelationContentUpdate(dto *PostCategoryEntity, query workspaces.QueryDSL) error {
	return nil
}
func PostCategoryPolyglotCreateHandler(dto *PostCategoryEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &PostCategoryEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func PostCategoryValidator(dto *PostCategoryEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func PostCategoryEntityPreSanitize(dto *PostCategoryEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func PostCategoryEntityBeforeCreateAppend(dto *PostCategoryEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    PostCategoryRecursiveAddUniqueId(dto, query)
  }
  func PostCategoryRecursiveAddUniqueId(dto *PostCategoryEntity, query workspaces.QueryDSL) {
  }
func PostCategoryActionBatchCreateFn(dtos []*PostCategoryEntity, query workspaces.QueryDSL) ([]*PostCategoryEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*PostCategoryEntity{}
		for _, item := range dtos {
			s, err := PostCategoryActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func PostCategoryDeleteEntireChildren(query workspaces.QueryDSL, dto *PostCategoryEntity) (*workspaces.IError) {
  return nil
}
func PostCategoryActionCreateFn(dto *PostCategoryEntity, query workspaces.QueryDSL) (*PostCategoryEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := PostCategoryValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	PostCategoryEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	PostCategoryEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	PostCategoryPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	PostCategoryRelationContentCreate(dto, query)
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
	PostCategoryAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(POST_CATEGORY_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&PostCategoryEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func PostCategoryActionGetOne(query workspaces.QueryDSL) (*PostCategoryEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&PostCategoryEntity{})
    item, err := workspaces.GetOneEntity[PostCategoryEntity](query, refl)
    entityPostCategoryFormatter(item, query)
    return item, err
  }
  func PostCategoryActionQuery(query workspaces.QueryDSL) ([]*PostCategoryEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&PostCategoryEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[PostCategoryEntity](query, refl)
    for _, item := range items {
      entityPostCategoryFormatter(item, query)
    }
    return items, meta, err
  }
  func PostCategoryUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *PostCategoryEntity) (*PostCategoryEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = POST_CATEGORY_EVENT_UPDATED
    PostCategoryEntityPreSanitize(fields, query)
    var item PostCategoryEntity
    q := dbref.
      Where(&PostCategoryEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    PostCategoryRelationContentUpdate(fields, query)
    PostCategoryPolyglotCreateHandler(fields, query)
    if ero := PostCategoryDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&PostCategoryEntity{UniqueId: uniqueId}).
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
  func PostCategoryActionUpdateFn(query workspaces.QueryDSL, fields *PostCategoryEntity) (*PostCategoryEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := PostCategoryValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // PostCategoryRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *PostCategoryEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = PostCategoryUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return PostCategoryUpdateExec(dbref, query, fields)
    }
  }
var PostCategoryWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire postcategories ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_CATEGORY_DELETE},
    })
		count, _ := PostCategoryActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func PostCategoryActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&PostCategoryEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_POST_CATEGORY_DELETE}
	return workspaces.RemoveEntity[PostCategoryEntity](query, refl)
}
func PostCategoryActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[PostCategoryEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'PostCategoryEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func PostCategoryActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[PostCategoryEntity]) (
    *workspaces.BulkRecordRequest[PostCategoryEntity], *workspaces.IError,
  ) {
    result := []*PostCategoryEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := PostCategoryActionUpdate(query, record)
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
func (x *PostCategoryEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var PostCategoryEntityMeta = workspaces.TableMetaData{
	EntityName:    "PostCategory",
	ExportKey:    "post-categories",
	TableNameInDb: "fb_post-category_entities",
	EntityObject:  &PostCategoryEntity{},
	ExportStream: PostCategoryActionExportT,
	ImportQuery: PostCategoryActionImport,
}
func PostCategoryActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[PostCategoryEntity](query, PostCategoryActionQuery, PostCategoryPreloadRelations)
}
func PostCategoryActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[PostCategoryEntity](query, PostCategoryActionQuery, PostCategoryPreloadRelations)
}
func PostCategoryActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content PostCategoryEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := PostCategoryActionCreate(&content, query)
	return err
}
var PostCategoryCommonCliFlags = []cli.Flag{
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
var PostCategoryCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "name",
		StructField:     "Name",
		Required: true,
		Usage:    "name",
		Type: "string",
	},
}
var PostCategoryCommonCliFlagsOptional = []cli.Flag{
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
  var PostCategoryCreateCmd cli.Command = POST_CATEGORY_ACTION_POST_ONE.ToCli()
  var PostCategoryCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_CATEGORY_CREATE},
      })
      entity := &PostCategoryEntity{}
      for _, item := range PostCategoryCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := PostCategoryActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var PostCategoryUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: PostCategoryCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_CATEGORY_UPDATE},
      })
      entity := CastPostCategoryFromCli(c)
      if entity, err := PostCategoryActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* PostCategoryEntity) FromCli(c *cli.Context) *PostCategoryEntity {
	return CastPostCategoryFromCli(c)
}
func CastPostCategoryFromCli (c *cli.Context) *PostCategoryEntity {
	template := &PostCategoryEntity{}
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
  func PostCategorySyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      PostCategoryActionCreate,
      reflect.ValueOf(&PostCategoryEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func PostCategoryWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := PostCategoryActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "PostCategory", result)
    }
  }
var PostCategoryImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_CATEGORY_CREATE},
      })
			PostCategoryActionSeeder(query, c.Int("count"))
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
				Value: "post-category-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_CATEGORY_CREATE},
      })
			PostCategoryActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "post-category-seeder-post-category.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of post-categories, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]PostCategoryEntity{}
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
			PostCategoryCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				PostCategoryActionCreate,
				reflect.ValueOf(&PostCategoryEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_CATEGORY_CREATE},
				},
        func() PostCategoryEntity {
					v := CastPostCategoryFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var PostCategoryCliCommands []cli.Command = []cli.Command{
      workspaces.GetCommonQuery2(PostCategoryActionQuery, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_CATEGORY_CREATE},
      }),
      workspaces.GetCommonTableQuery(reflect.ValueOf(&PostCategoryEntity{}).Elem(), PostCategoryActionQuery),
          PostCategoryCreateCmd,
          PostCategoryUpdateCmd,
          PostCategoryCreateInteractiveCmd,
          PostCategoryWipeCmd,
          workspaces.GetCommonRemoveQuery(reflect.ValueOf(&PostCategoryEntity{}).Elem(), PostCategoryActionRemove),
  }
  func PostCategoryCliFn() cli.Command {
    PostCategoryCliCommands = append(PostCategoryCliCommands, PostCategoryImportExportCommands...)
    return cli.Command{
      Name:        "postCategory",
      Description: "PostCategorys module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: PostCategoryCliCommands,
    }
  }
var POST_CATEGORY_ACTION_POST_ONE = workspaces.Module2Action{
    ActionName:    "create",
    ActionAliases: []string{"c"},
    Description: "Create new postCategory",
    Flags: PostCategoryCommonCliFlags,
    Method: "POST",
    Url:    "/post-category",
    SecurityModel: &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_CATEGORY_CREATE},
    },
    Handlers: []gin.HandlerFunc{
      func (c *gin.Context) {
        workspaces.HttpPostEntity(c, PostCategoryActionCreate)
      },
    },
    CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
      result, err := workspaces.CliPostEntity(c, PostCategoryActionCreate, security)
      workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
      return err
    },
    Action: PostCategoryActionCreate,
    Format: "POST_ONE",
    RequestEntity: &PostCategoryEntity{},
    ResponseEntity: &PostCategoryEntity{},
  }
  /**
  *	Override this function on PostCategoryEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendPostCategoryRouter = func(r *[]workspaces.Module2Action) {}
  func GetPostCategoryModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
       {
        Method: "GET",
        Url:    "/post-categories",
        SecurityModel: &workspaces.SecurityModel{
          ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_CATEGORY_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpQueryEntity(c, PostCategoryActionQuery)
          },
        },
        Format: "QUERY",
        Action: PostCategoryActionQuery,
        ResponseEntity: &[]PostCategoryEntity{},
      },
      {
        Method: "GET",
        Url:    "/post-categories/export",
        SecurityModel: &workspaces.SecurityModel{
          ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_CATEGORY_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpStreamFileChannel(c, PostCategoryActionExport)
          },
        },
        Format: "QUERY",
        Action: PostCategoryActionExport,
        ResponseEntity: &[]PostCategoryEntity{},
      },
      {
        Method: "GET",
        Url:    "/post-category/:uniqueId",
        SecurityModel: &workspaces.SecurityModel{
          ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_CATEGORY_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpGetEntity(c, PostCategoryActionGetOne)
          },
        },
        Format: "GET_ONE",
        Action: PostCategoryActionGetOne,
        ResponseEntity: &PostCategoryEntity{},
      },
      POST_CATEGORY_ACTION_POST_ONE,
      {
        ActionName:    "update",
        ActionAliases: []string{"u"},
        Flags: PostCategoryCommonCliFlagsOptional,
        Method: "PATCH",
        Url:    "/post-category",
        SecurityModel: &workspaces.SecurityModel{
          ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_CATEGORY_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpUpdateEntity(c, PostCategoryActionUpdate)
          },
        },
        Action: PostCategoryActionUpdate,
        RequestEntity: &PostCategoryEntity{},
        Format: "PATCH_ONE",
        ResponseEntity: &PostCategoryEntity{},
      },
      {
        Method: "PATCH",
        Url:    "/post-categories",
        SecurityModel: &workspaces.SecurityModel{
          ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_CATEGORY_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpUpdateEntities(c, PostCategoryActionBulkUpdate)
          },
        },
        Action: PostCategoryActionBulkUpdate,
        Format: "PATCH_BULK",
        RequestEntity:  &workspaces.BulkRecordRequest[PostCategoryEntity]{},
        ResponseEntity: &workspaces.BulkRecordRequest[PostCategoryEntity]{},
      },
      {
        Method: "DELETE",
        Url:    "/post-category",
        Format: "DELETE_DSL",
        SecurityModel: &workspaces.SecurityModel{
          ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_CATEGORY_DELETE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpRemoveEntity(c, PostCategoryActionRemove)
          },
        },
        Action: PostCategoryActionRemove,
        RequestEntity: &workspaces.DeleteRequest{},
        ResponseEntity: &workspaces.DeleteResponse{},
        TargetEntity: &PostCategoryEntity{},
      },
    }
    // Append user defined functions
    AppendPostCategoryRouter(&routes)
    return routes
  }
  func CreatePostCategoryRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetPostCategoryModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, PostCategoryEntityJsonSchema, "post-category-http", "cms")
    workspaces.WriteEntitySchema("PostCategoryEntity", PostCategoryEntityJsonSchema, "cms")
    return httpRoutes
  }
var PERM_ROOT_POST_CATEGORY_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/cms/post-category/delete",
}
var PERM_ROOT_POST_CATEGORY_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/cms/post-category/create",
}
var PERM_ROOT_POST_CATEGORY_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/cms/post-category/update",
}
var PERM_ROOT_POST_CATEGORY_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/cms/post-category/query",
}
var PERM_ROOT_POST_CATEGORY = workspaces.PermissionInfo{
  CompleteKey: "root/cms/post-category/*",
}
var ALL_POST_CATEGORY_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_POST_CATEGORY_DELETE,
	PERM_ROOT_POST_CATEGORY_CREATE,
	PERM_ROOT_POST_CATEGORY_UPDATE,
	PERM_ROOT_POST_CATEGORY_QUERY,
	PERM_ROOT_POST_CATEGORY,
}