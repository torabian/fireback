package workspaces
import (
    "github.com/gin-gonic/gin"
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
type PreferenceEntity struct {
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
    Timezone   *string `json:"timezone" yaml:"timezone"       `
    // Datenano also has a text representation
    Children []*PreferenceEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *PreferenceEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var PreferencePreloadRelations []string = []string{}
var PREFERENCE_EVENT_CREATED = "preference.created"
var PREFERENCE_EVENT_UPDATED = "preference.updated"
var PREFERENCE_EVENT_DELETED = "preference.deleted"
var PREFERENCE_EVENTS = []string{
	PREFERENCE_EVENT_CREATED,
	PREFERENCE_EVENT_UPDATED,
	PREFERENCE_EVENT_DELETED,
}
type PreferenceFieldMap struct {
		Timezone TranslatedString `yaml:"timezone"`
}
var PreferenceEntityMetaConfig map[string]int64 = map[string]int64{
}
var PreferenceEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&PreferenceEntity{}))
func entityPreferenceFormatter(dto *PreferenceEntity, query QueryDSL) {
	if dto == nil {
		return
	}
	if dto.Created > 0 {
		dto.CreatedFormatted = FormatDateBasedOnQuery(dto.Created, query)
	}
	if dto.Updated > 0 {
		dto.CreatedFormatted = FormatDateBasedOnQuery(dto.Updated, query)
	}
}
func PreferenceMockEntity() *PreferenceEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &PreferenceEntity{
      Timezone : &stringHolder,
	}
	return entity
}
func PreferenceActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := PreferenceMockEntity()
		_, err := PreferenceActionCreate(entity, query)
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
  func PreferenceActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*PreferenceEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &PreferenceEntity{
          Timezone: &tildaRef,
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
  func PreferenceAssociationCreate(dto *PreferenceEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func PreferenceRelationContentCreate(dto *PreferenceEntity, query QueryDSL) error {
return nil
}
func PreferenceRelationContentUpdate(dto *PreferenceEntity, query QueryDSL) error {
	return nil
}
func PreferencePolyglotCreateHandler(dto *PreferenceEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func PreferenceValidator(dto *PreferenceEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func PreferenceEntityPreSanitize(dto *PreferenceEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func PreferenceEntityBeforeCreateAppend(dto *PreferenceEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    PreferenceRecursiveAddUniqueId(dto, query)
  }
  func PreferenceRecursiveAddUniqueId(dto *PreferenceEntity, query QueryDSL) {
  }
func PreferenceActionBatchCreateFn(dtos []*PreferenceEntity, query QueryDSL) ([]*PreferenceEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*PreferenceEntity{}
		for _, item := range dtos {
			s, err := PreferenceActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func PreferenceDeleteEntireChildren(query QueryDSL, dto *PreferenceEntity) (*IError) {
  return nil
}
func PreferenceActionCreateFn(dto *PreferenceEntity, query QueryDSL) (*PreferenceEntity, *IError) {
	// 1. Validate always
	if iError := PreferenceValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	PreferenceEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	PreferenceEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	PreferencePolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	PreferenceRelationContentCreate(dto, query)
	// 4. Create the entity
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = GetDbRef()
	} else {
		dbref = query.Tx
	}
	query.Tx = dbref;
	err := dbref.Create(&dto).Error
	if err != nil {
		err := GormErrorToIError(err)
		return dto, err
	}
	// 5. Create sub entities, objects or arrays, association to other entities
	PreferenceAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(PREFERENCE_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&PreferenceEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func PreferenceActionGetOne(query QueryDSL) (*PreferenceEntity, *IError) {
    refl := reflect.ValueOf(&PreferenceEntity{})
    item, err := GetOneEntity[PreferenceEntity](query, refl)
    entityPreferenceFormatter(item, query)
    return item, err
  }
  func PreferenceActionQuery(query QueryDSL) ([]*PreferenceEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&PreferenceEntity{})
    items, meta, err := QueryEntitiesPointer[PreferenceEntity](query, refl)
    for _, item := range items {
      entityPreferenceFormatter(item, query)
    }
    return items, meta, err
  }
  func PreferenceUpdateExec(dbref *gorm.DB, query QueryDSL, fields *PreferenceEntity) (*PreferenceEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = PREFERENCE_EVENT_UPDATED
    PreferenceEntityPreSanitize(fields, query)
    var item PreferenceEntity
    q := dbref.
      Where(&PreferenceEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    PreferenceRelationContentUpdate(fields, query)
    PreferencePolyglotCreateHandler(fields, query)
    if ero := PreferenceDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&PreferenceEntity{UniqueId: uniqueId}).
      First(&item).Error
    event.MustFire(query.TriggerEventName, event.M{
      "entity":   &item,
      "target":   "workspace",
      "unqiueId": query.WorkspaceId,
    })
    if err != nil {
      return &item, GormErrorToIError(err)
    }
    return &item, nil
  }
  func PreferenceActionUpdateFn(query QueryDSL, fields *PreferenceEntity) (*PreferenceEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := PreferenceValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // PreferenceRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      var item *PreferenceEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *IError
        item, err = PreferenceUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, CastToIError(vf)
    } else {
      dbref = query.Tx
      return PreferenceUpdateExec(dbref, query, fields)
    }
  }
var PreferenceWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire preferences ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
      ActionRequires: []string{PERM_ROOT_PREFERENCE_DELETE},
    })
		count, _ := PreferenceActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func PreferenceActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&PreferenceEntity{})
	query.ActionRequires = []string{PERM_ROOT_PREFERENCE_DELETE}
	return RemoveEntity[PreferenceEntity](query, refl)
}
func PreferenceActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[PreferenceEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'PreferenceEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func PreferenceActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[PreferenceEntity]) (
    *BulkRecordRequest[PreferenceEntity], *IError,
  ) {
    result := []*PreferenceEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := PreferenceActionUpdate(query, record)
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
    return nil, err.(*IError)
  }
func (x *PreferenceEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var PreferenceEntityMeta = TableMetaData{
	EntityName:    "Preference",
	ExportKey:    "preferences",
	TableNameInDb: "fb_preference_entities",
	EntityObject:  &PreferenceEntity{},
	ExportStream: PreferenceActionExportT,
	ImportQuery: PreferenceActionImport,
}
func PreferenceActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[PreferenceEntity](query, PreferenceActionQuery, PreferencePreloadRelations)
}
func PreferenceActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[PreferenceEntity](query, PreferenceActionQuery, PreferencePreloadRelations)
}
func PreferenceActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content PreferenceEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := PreferenceActionCreate(&content, query)
	return err
}
var PreferenceCommonCliFlags = []cli.Flag{
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
      Name:     "timezone",
      Required: false,
      Usage:    "timezone",
    },
}
var PreferenceCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:     "timezone",
		StructField:     "Timezone",
		Required: false,
		Usage:    "timezone",
		Type: "string",
	},
}
var PreferenceCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "timezone",
      Required: false,
      Usage:    "timezone",
    },
}
  var PreferenceCreateCmd cli.Command = PREFERENCE_ACTION_POST_ONE.ToCli()
  var PreferenceCreateInteractiveCmd cli.Command = cli.Command{
    Name:  "ic",
    Usage: "Creates a new template, using requied fields in an interactive name",
    Flags: []cli.Flag{
      &cli.BoolFlag{
        Name:  "all",
        Usage: "Interactively asks for all inputs, not only required ones",
      },
    },
    Action: func(c *cli.Context) {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []string{PERM_ROOT_PREFERENCE_CREATE},
      })
      entity := &PreferenceEntity{}
      for _, item := range PreferenceCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := PreferenceActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var PreferenceUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: PreferenceCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []string{PERM_ROOT_PREFERENCE_UPDATE},
      })
      entity := CastPreferenceFromCli(c)
      if entity, err := PreferenceActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* PreferenceEntity) FromCli(c *cli.Context) *PreferenceEntity {
	return CastPreferenceFromCli(c)
}
func CastPreferenceFromCli (c *cli.Context) *PreferenceEntity {
	template := &PreferenceEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("timezone") {
        value := c.String("timezone")
        template.Timezone = &value
      }
	return template
}
  func PreferenceSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      PreferenceActionCreate,
      reflect.ValueOf(&PreferenceEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func PreferenceWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := PreferenceActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "Preference", result)
    }
  }
var PreferenceImportExportCommands = []cli.Command{
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
			query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []string{PERM_ROOT_PREFERENCE_CREATE},
      })
			PreferenceActionSeeder(query, c.Int("count"))
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
				Value: "preference-seeder.yml",
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
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []string{PERM_ROOT_PREFERENCE_CREATE},
      })
			PreferenceActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "preference-seeder-preference.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of preferences, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]PreferenceEntity{}
			ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
	cli.Command{
		Name:    "import",
    Flags: append(
			append(
				CommonQueryFlags,
				&cli.StringFlag{
					Name:     "file",
					Usage:    "The address of file you want the csv be imported from",
					Required: true,
				}),
			PreferenceCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				PreferenceActionCreate,
				reflect.ValueOf(&PreferenceEntity{}).Elem(),
				c.String("file"),
        &SecurityModel{
					ActionRequires: []string{PERM_ROOT_PREFERENCE_CREATE},
				},
        func() PreferenceEntity {
					v := CastPreferenceFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var PreferenceCliCommands []cli.Command = []cli.Command{
      GetCommonQuery2(PreferenceActionQuery, &SecurityModel{
        ActionRequires: []string{PERM_ROOT_PREFERENCE_CREATE},
      }),
      GetCommonTableQuery(reflect.ValueOf(&PreferenceEntity{}).Elem(), PreferenceActionQuery),
          PreferenceCreateCmd,
          PreferenceUpdateCmd,
          PreferenceCreateInteractiveCmd,
          PreferenceWipeCmd,
          GetCommonRemoveQuery(reflect.ValueOf(&PreferenceEntity{}).Elem(), PreferenceActionRemove),
  }
  func PreferenceCliFn() cli.Command {
    PreferenceCliCommands = append(PreferenceCliCommands, PreferenceImportExportCommands...)
    return cli.Command{
      Name:        "preference",
      Description: "Preferences module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: PreferenceCliCommands,
    }
  }
var PREFERENCE_ACTION_POST_ONE = Module2Action{
    ActionName:    "create",
    ActionAliases: []string{"c"},
    Description: "Create new preference",
    Flags: PreferenceCommonCliFlags,
    Method: "POST",
    Url:    "/preference",
    SecurityModel: &SecurityModel{
      ActionRequires: []string{PERM_ROOT_PREFERENCE_CREATE},
    },
    Handlers: []gin.HandlerFunc{
      func (c *gin.Context) {
        HttpPostEntity(c, PreferenceActionCreate)
      },
    },
    CliAction: func(c *cli.Context, security *SecurityModel) error {
      result, err := CliPostEntity(c, PreferenceActionCreate, security)
      HandleActionInCli(c, result, err, map[string]map[string]string{})
      return err
    },
    Action: PreferenceActionCreate,
    Format: "POST_ONE",
    RequestEntity: &PreferenceEntity{},
    ResponseEntity: &PreferenceEntity{},
  }
  /**
  *	Override this function on PreferenceEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendPreferenceRouter = func(r *[]Module2Action) {}
  func GetPreferenceModule2Actions() []Module2Action {
    routes := []Module2Action{
       {
        Method: "GET",
        Url:    "/preferences",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_PREFERENCE_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpQueryEntity(c, PreferenceActionQuery)
          },
        },
        Format: "QUERY",
        Action: PreferenceActionQuery,
        ResponseEntity: &[]PreferenceEntity{},
      },
      {
        Method: "GET",
        Url:    "/preferences/export",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_PREFERENCE_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpStreamFileChannel(c, PreferenceActionExport)
          },
        },
        Format: "QUERY",
        Action: PreferenceActionExport,
        ResponseEntity: &[]PreferenceEntity{},
      },
      {
        Method: "GET",
        Url:    "/preference/:uniqueId",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_PREFERENCE_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpGetEntity(c, PreferenceActionGetOne)
          },
        },
        Format: "GET_ONE",
        Action: PreferenceActionGetOne,
        ResponseEntity: &PreferenceEntity{},
      },
      PREFERENCE_ACTION_POST_ONE,
      {
        ActionName:    "update",
        ActionAliases: []string{"u"},
        Flags: PreferenceCommonCliFlagsOptional,
        Method: "PATCH",
        Url:    "/preference",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_PREFERENCE_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpUpdateEntity(c, PreferenceActionUpdate)
          },
        },
        Action: PreferenceActionUpdate,
        RequestEntity: &PreferenceEntity{},
        Format: "PATCH_ONE",
        ResponseEntity: &PreferenceEntity{},
      },
      {
        Method: "PATCH",
        Url:    "/preferences",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_PREFERENCE_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpUpdateEntities(c, PreferenceActionBulkUpdate)
          },
        },
        Action: PreferenceActionBulkUpdate,
        Format: "PATCH_BULK",
        RequestEntity:  &BulkRecordRequest[PreferenceEntity]{},
        ResponseEntity: &BulkRecordRequest[PreferenceEntity]{},
      },
      {
        Method: "DELETE",
        Url:    "/preference",
        Format: "DELETE_DSL",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_PREFERENCE_DELETE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpRemoveEntity(c, PreferenceActionRemove)
          },
        },
        Action: PreferenceActionRemove,
        RequestEntity: &DeleteRequest{},
        ResponseEntity: &DeleteResponse{},
        TargetEntity: &PreferenceEntity{},
      },
    }
    // Append user defined functions
    AppendPreferenceRouter(&routes)
    return routes
  }
  func CreatePreferenceRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetPreferenceModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, PreferenceEntityJsonSchema, "preference-http", "workspaces")
    WriteEntitySchema("PreferenceEntity", PreferenceEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_PREFERENCE_DELETE = "root/workspaces/preference/delete"
var PERM_ROOT_PREFERENCE_CREATE = "root/workspaces/preference/create"
var PERM_ROOT_PREFERENCE_UPDATE = "root/workspaces/preference/update"
var PERM_ROOT_PREFERENCE_QUERY = "root/workspaces/preference/query"
var PERM_ROOT_PREFERENCE = "root/workspaces/preference/*"
var ALL_PREFERENCE_PERMISSIONS = []string{
	PERM_ROOT_PREFERENCE_DELETE,
	PERM_ROOT_PREFERENCE_CREATE,
	PERM_ROOT_PREFERENCE_UPDATE,
	PERM_ROOT_PREFERENCE_QUERY,
	PERM_ROOT_PREFERENCE,
}