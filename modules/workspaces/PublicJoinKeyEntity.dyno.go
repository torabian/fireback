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
type PublicJoinKeyEntity struct {
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
    Role   *  RoleEntity `json:"role" yaml:"role"    gorm:"foreignKey:RoleId;references:UniqueId"     `
    // Datenano also has a text representation
        RoleId *string `json:"roleId" yaml:"roleId"`
    Workspace   *  WorkspaceEntity `json:"workspace" yaml:"workspace"    gorm:"foreignKey:WorkspaceId;references:UniqueId"     `
    // Datenano also has a text representation
    Children []*PublicJoinKeyEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *PublicJoinKeyEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var PublicJoinKeyPreloadRelations []string = []string{}
var PUBLIC_JOIN_KEY_EVENT_CREATED = "publicJoinKey.created"
var PUBLIC_JOIN_KEY_EVENT_UPDATED = "publicJoinKey.updated"
var PUBLIC_JOIN_KEY_EVENT_DELETED = "publicJoinKey.deleted"
var PUBLIC_JOIN_KEY_EVENTS = []string{
	PUBLIC_JOIN_KEY_EVENT_CREATED,
	PUBLIC_JOIN_KEY_EVENT_UPDATED,
	PUBLIC_JOIN_KEY_EVENT_DELETED,
}
type PublicJoinKeyFieldMap struct {
		Role TranslatedString `yaml:"role"`
		Workspace TranslatedString `yaml:"workspace"`
}
var PublicJoinKeyEntityMetaConfig map[string]int64 = map[string]int64{
}
var PublicJoinKeyEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&PublicJoinKeyEntity{}))
func entityPublicJoinKeyFormatter(dto *PublicJoinKeyEntity, query QueryDSL) {
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
func PublicJoinKeyMockEntity() *PublicJoinKeyEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &PublicJoinKeyEntity{
	}
	return entity
}
func PublicJoinKeyActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := PublicJoinKeyMockEntity()
		_, err := PublicJoinKeyActionCreate(entity, query)
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
  func PublicJoinKeyActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*PublicJoinKeyEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &PublicJoinKeyEntity{
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
  func PublicJoinKeyAssociationCreate(dto *PublicJoinKeyEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func PublicJoinKeyRelationContentCreate(dto *PublicJoinKeyEntity, query QueryDSL) error {
return nil
}
func PublicJoinKeyRelationContentUpdate(dto *PublicJoinKeyEntity, query QueryDSL) error {
	return nil
}
func PublicJoinKeyPolyglotCreateHandler(dto *PublicJoinKeyEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func PublicJoinKeyValidator(dto *PublicJoinKeyEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func PublicJoinKeyEntityPreSanitize(dto *PublicJoinKeyEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func PublicJoinKeyEntityBeforeCreateAppend(dto *PublicJoinKeyEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    PublicJoinKeyRecursiveAddUniqueId(dto, query)
  }
  func PublicJoinKeyRecursiveAddUniqueId(dto *PublicJoinKeyEntity, query QueryDSL) {
  }
func PublicJoinKeyActionBatchCreateFn(dtos []*PublicJoinKeyEntity, query QueryDSL) ([]*PublicJoinKeyEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*PublicJoinKeyEntity{}
		for _, item := range dtos {
			s, err := PublicJoinKeyActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func PublicJoinKeyDeleteEntireChildren(query QueryDSL, dto *PublicJoinKeyEntity) (*IError) {
  return nil
}
func PublicJoinKeyActionCreateFn(dto *PublicJoinKeyEntity, query QueryDSL) (*PublicJoinKeyEntity, *IError) {
	// 1. Validate always
	if iError := PublicJoinKeyValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	PublicJoinKeyEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	PublicJoinKeyEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	PublicJoinKeyPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	PublicJoinKeyRelationContentCreate(dto, query)
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
	PublicJoinKeyAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(PUBLIC_JOIN_KEY_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&PublicJoinKeyEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func PublicJoinKeyActionGetOne(query QueryDSL) (*PublicJoinKeyEntity, *IError) {
    refl := reflect.ValueOf(&PublicJoinKeyEntity{})
    item, err := GetOneEntity[PublicJoinKeyEntity](query, refl)
    entityPublicJoinKeyFormatter(item, query)
    return item, err
  }
  func PublicJoinKeyActionQuery(query QueryDSL) ([]*PublicJoinKeyEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&PublicJoinKeyEntity{})
    items, meta, err := QueryEntitiesPointer[PublicJoinKeyEntity](query, refl)
    for _, item := range items {
      entityPublicJoinKeyFormatter(item, query)
    }
    return items, meta, err
  }
  func PublicJoinKeyUpdateExec(dbref *gorm.DB, query QueryDSL, fields *PublicJoinKeyEntity) (*PublicJoinKeyEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = PUBLIC_JOIN_KEY_EVENT_UPDATED
    PublicJoinKeyEntityPreSanitize(fields, query)
    var item PublicJoinKeyEntity
    q := dbref.
      Where(&PublicJoinKeyEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    PublicJoinKeyRelationContentUpdate(fields, query)
    PublicJoinKeyPolyglotCreateHandler(fields, query)
    if ero := PublicJoinKeyDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&PublicJoinKeyEntity{UniqueId: uniqueId}).
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
  func PublicJoinKeyActionUpdateFn(query QueryDSL, fields *PublicJoinKeyEntity) (*PublicJoinKeyEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := PublicJoinKeyValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // PublicJoinKeyRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      var item *PublicJoinKeyEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *IError
        item, err = PublicJoinKeyUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, CastToIError(vf)
    } else {
      dbref = query.Tx
      return PublicJoinKeyUpdateExec(dbref, query, fields)
    }
  }
var PublicJoinKeyWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire publicjoinkeys ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
      ActionRequires: []PermissionInfo{PERM_ROOT_PUBLIC_JOIN_KEY_DELETE},
    })
		count, _ := PublicJoinKeyActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func PublicJoinKeyActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&PublicJoinKeyEntity{})
	query.ActionRequires = []PermissionInfo{PERM_ROOT_PUBLIC_JOIN_KEY_DELETE}
	return RemoveEntity[PublicJoinKeyEntity](query, refl)
}
func PublicJoinKeyActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[PublicJoinKeyEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'PublicJoinKeyEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func PublicJoinKeyActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[PublicJoinKeyEntity]) (
    *BulkRecordRequest[PublicJoinKeyEntity], *IError,
  ) {
    result := []*PublicJoinKeyEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := PublicJoinKeyActionUpdate(query, record)
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
func (x *PublicJoinKeyEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var PublicJoinKeyEntityMeta = TableMetaData{
	EntityName:    "PublicJoinKey",
	ExportKey:    "public-join-keys",
	TableNameInDb: "fb_public-join-key_entities",
	EntityObject:  &PublicJoinKeyEntity{},
	ExportStream: PublicJoinKeyActionExportT,
	ImportQuery: PublicJoinKeyActionImport,
}
func PublicJoinKeyActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[PublicJoinKeyEntity](query, PublicJoinKeyActionQuery, PublicJoinKeyPreloadRelations)
}
func PublicJoinKeyActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[PublicJoinKeyEntity](query, PublicJoinKeyActionQuery, PublicJoinKeyPreloadRelations)
}
func PublicJoinKeyActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content PublicJoinKeyEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := PublicJoinKeyActionCreate(&content, query)
	return err
}
var PublicJoinKeyCommonCliFlags = []cli.Flag{
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
      Name:     "role-id",
      Required: false,
      Usage:    "role",
    },
    &cli.StringFlag{
      Name:     "workspace-id",
      Required: false,
      Usage:    "workspace",
    },
}
var PublicJoinKeyCommonInteractiveCliFlags = []CliInteractiveFlag{
}
var PublicJoinKeyCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "role-id",
      Required: false,
      Usage:    "role",
    },
    &cli.StringFlag{
      Name:     "workspace-id",
      Required: false,
      Usage:    "workspace",
    },
}
  var PublicJoinKeyCreateCmd cli.Command = PUBLIC_JOIN_KEY_ACTION_POST_ONE.ToCli()
  var PublicJoinKeyCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_PUBLIC_JOIN_KEY_CREATE},
      })
      entity := &PublicJoinKeyEntity{}
      for _, item := range PublicJoinKeyCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := PublicJoinKeyActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var PublicJoinKeyUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: PublicJoinKeyCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_PUBLIC_JOIN_KEY_UPDATE},
      })
      entity := CastPublicJoinKeyFromCli(c)
      if entity, err := PublicJoinKeyActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* PublicJoinKeyEntity) FromCli(c *cli.Context) *PublicJoinKeyEntity {
	return CastPublicJoinKeyFromCli(c)
}
func CastPublicJoinKeyFromCli (c *cli.Context) *PublicJoinKeyEntity {
	template := &PublicJoinKeyEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("role-id") {
        value := c.String("role-id")
        template.RoleId = &value
      }
      if c.IsSet("workspace-id") {
        value := c.String("workspace-id")
        template.WorkspaceId = &value
      }
	return template
}
  func PublicJoinKeySyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      PublicJoinKeyActionCreate,
      reflect.ValueOf(&PublicJoinKeyEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func PublicJoinKeyWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := PublicJoinKeyActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "PublicJoinKey", result)
    }
  }
var PublicJoinKeyImportExportCommands = []cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_PUBLIC_JOIN_KEY_CREATE},
      })
			PublicJoinKeyActionSeeder(query, c.Int("count"))
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
				Value: "public-join-key-seeder.yml",
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
        ActionRequires: []PermissionInfo{PERM_ROOT_PUBLIC_JOIN_KEY_CREATE},
      })
			PublicJoinKeyActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "public-join-key-seeder-public-join-key.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of public-join-keys, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]PublicJoinKeyEntity{}
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
			PublicJoinKeyCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				PublicJoinKeyActionCreate,
				reflect.ValueOf(&PublicJoinKeyEntity{}).Elem(),
				c.String("file"),
        &SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_PUBLIC_JOIN_KEY_CREATE},
				},
        func() PublicJoinKeyEntity {
					v := CastPublicJoinKeyFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var PublicJoinKeyCliCommands []cli.Command = []cli.Command{
      PUBLIC_JOIN_KEY_ACTION_QUERY.ToCli(),
      GetCommonTableQuery(reflect.ValueOf(&PublicJoinKeyEntity{}).Elem(), PublicJoinKeyActionQuery),
      PublicJoinKeyCreateCmd,
      PublicJoinKeyUpdateCmd,
      PublicJoinKeyCreateInteractiveCmd,
      PublicJoinKeyWipeCmd,
      GetCommonRemoveQuery(reflect.ValueOf(&PublicJoinKeyEntity{}).Elem(), PublicJoinKeyActionRemove),
  }
  func PublicJoinKeyCliFn() cli.Command {
    PublicJoinKeyCliCommands = append(PublicJoinKeyCliCommands, PublicJoinKeyImportExportCommands...)
    return cli.Command{
      Name:        "publicJoinKey",
      Description: "PublicJoinKeys module actions (sample module to handle complex entities)",
      Usage:       "Joining to different workspaces using a public link directly",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: PublicJoinKeyCliCommands,
    }
  }
var PUBLIC_JOIN_KEY_ACTION_QUERY = Module2Action{
  Method: "GET",
  Url:    "/public-join-keys",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PUBLIC_JOIN_KEY_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpQueryEntity(c, PublicJoinKeyActionQuery)
    },
  },
  Format: "QUERY",
  Action: PublicJoinKeyActionQuery,
  ResponseEntity: &[]PublicJoinKeyEntity{},
  CliAction: func(c *cli.Context, security *SecurityModel) error {
		CommonCliQueryCmd2(
			c,
			PublicJoinKeyActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionAliases: []string{"q"},
	Flags:         CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var PUBLIC_JOIN_KEY_ACTION_EXPORT = Module2Action{
  Method: "GET",
  Url:    "/public-join-keys/export",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PUBLIC_JOIN_KEY_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpStreamFileChannel(c, PublicJoinKeyActionExport)
    },
  },
  Format: "QUERY",
  Action: PublicJoinKeyActionExport,
  ResponseEntity: &[]PublicJoinKeyEntity{},
}
var PUBLIC_JOIN_KEY_ACTION_GET_ONE = Module2Action{
  Method: "GET",
  Url:    "/public-join-key/:uniqueId",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PUBLIC_JOIN_KEY_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpGetEntity(c, PublicJoinKeyActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: PublicJoinKeyActionGetOne,
  ResponseEntity: &PublicJoinKeyEntity{},
}
var PUBLIC_JOIN_KEY_ACTION_POST_ONE = Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new publicJoinKey",
  Flags: PublicJoinKeyCommonCliFlags,
  Method: "POST",
  Url:    "/public-join-key",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PUBLIC_JOIN_KEY_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpPostEntity(c, PublicJoinKeyActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    result, err := CliPostEntity(c, PublicJoinKeyActionCreate, security)
    HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: PublicJoinKeyActionCreate,
  Format: "POST_ONE",
  RequestEntity: &PublicJoinKeyEntity{},
  ResponseEntity: &PublicJoinKeyEntity{},
}
var PUBLIC_JOIN_KEY_ACTION_PATCH = Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: PublicJoinKeyCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/public-join-key",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PUBLIC_JOIN_KEY_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntity(c, PublicJoinKeyActionUpdate)
    },
  },
  Action: PublicJoinKeyActionUpdate,
  RequestEntity: &PublicJoinKeyEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &PublicJoinKeyEntity{},
}
var PUBLIC_JOIN_KEY_ACTION_PATCH_BULK = Module2Action{
  Method: "PATCH",
  Url:    "/public-join-keys",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PUBLIC_JOIN_KEY_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntities(c, PublicJoinKeyActionBulkUpdate)
    },
  },
  Action: PublicJoinKeyActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &BulkRecordRequest[PublicJoinKeyEntity]{},
  ResponseEntity: &BulkRecordRequest[PublicJoinKeyEntity]{},
}
var PUBLIC_JOIN_KEY_ACTION_DELETE = Module2Action{
  Method: "DELETE",
  Url:    "/public-join-key",
  Format: "DELETE_DSL",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PUBLIC_JOIN_KEY_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpRemoveEntity(c, PublicJoinKeyActionRemove)
    },
  },
  Action: PublicJoinKeyActionRemove,
  RequestEntity: &DeleteRequest{},
  ResponseEntity: &DeleteResponse{},
  TargetEntity: &PublicJoinKeyEntity{},
}
  /**
  *	Override this function on PublicJoinKeyEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendPublicJoinKeyRouter = func(r *[]Module2Action) {}
  func GetPublicJoinKeyModule2Actions() []Module2Action {
    routes := []Module2Action{
      PUBLIC_JOIN_KEY_ACTION_QUERY,
      PUBLIC_JOIN_KEY_ACTION_EXPORT,
      PUBLIC_JOIN_KEY_ACTION_GET_ONE,
      PUBLIC_JOIN_KEY_ACTION_POST_ONE,
      PUBLIC_JOIN_KEY_ACTION_PATCH,
      PUBLIC_JOIN_KEY_ACTION_PATCH_BULK,
      PUBLIC_JOIN_KEY_ACTION_DELETE,
    }
    // Append user defined functions
    AppendPublicJoinKeyRouter(&routes)
    return routes
  }
  func CreatePublicJoinKeyRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetPublicJoinKeyModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, PublicJoinKeyEntityJsonSchema, "public-join-key-http", "workspaces")
    WriteEntitySchema("PublicJoinKeyEntity", PublicJoinKeyEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_PUBLIC_JOIN_KEY_DELETE = PermissionInfo{
  CompleteKey: "root/workspaces/public-join-key/delete",
}
var PERM_ROOT_PUBLIC_JOIN_KEY_CREATE = PermissionInfo{
  CompleteKey: "root/workspaces/public-join-key/create",
}
var PERM_ROOT_PUBLIC_JOIN_KEY_UPDATE = PermissionInfo{
  CompleteKey: "root/workspaces/public-join-key/update",
}
var PERM_ROOT_PUBLIC_JOIN_KEY_QUERY = PermissionInfo{
  CompleteKey: "root/workspaces/public-join-key/query",
}
var PERM_ROOT_PUBLIC_JOIN_KEY = PermissionInfo{
  CompleteKey: "root/workspaces/public-join-key/*",
}
var ALL_PUBLIC_JOIN_KEY_PERMISSIONS = []PermissionInfo{
	PERM_ROOT_PUBLIC_JOIN_KEY_DELETE,
	PERM_ROOT_PUBLIC_JOIN_KEY_CREATE,
	PERM_ROOT_PUBLIC_JOIN_KEY_UPDATE,
	PERM_ROOT_PUBLIC_JOIN_KEY_QUERY,
	PERM_ROOT_PUBLIC_JOIN_KEY,
}