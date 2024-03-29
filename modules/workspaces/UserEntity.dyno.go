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
	mocks "github.com/torabian/fireback/modules/workspaces/mocks/User"
)
type UserEntity struct {
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
    Person   *  PersonEntity `json:"person" yaml:"person"    gorm:"foreignKey:PersonId;references:UniqueId"     `
    // Datenano also has a text representation
        PersonId *string `json:"personId" yaml:"personId"`
    Children []*UserEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *UserEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var UserPreloadRelations []string = []string{}
var USER_EVENT_CREATED = "user.created"
var USER_EVENT_UPDATED = "user.updated"
var USER_EVENT_DELETED = "user.deleted"
var USER_EVENTS = []string{
	USER_EVENT_CREATED,
	USER_EVENT_UPDATED,
	USER_EVENT_DELETED,
}
type UserFieldMap struct {
		Person TranslatedString `yaml:"person"`
}
var UserEntityMetaConfig map[string]int64 = map[string]int64{
}
var UserEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&UserEntity{}))
func entityUserFormatter(dto *UserEntity, query QueryDSL) {
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
func UserMockEntity() *UserEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &UserEntity{
	}
	return entity
}
func UserActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := UserMockEntity()
		_, err := UserActionCreate(entity, query)
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
  func UserActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*UserEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &UserEntity{
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
  func UserAssociationCreate(dto *UserEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func UserRelationContentCreate(dto *UserEntity, query QueryDSL) error {
  {
    if dto.Person != nil {
      dt, err := PersonActionCreate(dto.Person, query);
      if err != nil {
        return err;
      }
      dto.Person = dt;
    }
  }
return nil
}
func UserRelationContentUpdate(dto *UserEntity, query QueryDSL) error {
		{
			if dto.Person != nil {
				dt, err := PersonActionUpdate(query, dto.Person);
				if err != nil {
					return err;
				}
				dto.Person = dt;
			}
		}
	return nil
}
func UserPolyglotCreateHandler(dto *UserEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func UserValidator(dto *UserEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func UserEntityPreSanitize(dto *UserEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func UserEntityBeforeCreateAppend(dto *UserEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    UserRecursiveAddUniqueId(dto, query)
  }
  func UserRecursiveAddUniqueId(dto *UserEntity, query QueryDSL) {
  }
func UserActionBatchCreateFn(dtos []*UserEntity, query QueryDSL) ([]*UserEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*UserEntity{}
		for _, item := range dtos {
			s, err := UserActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func UserDeleteEntireChildren(query QueryDSL, dto *UserEntity) (*IError) {
  return nil
}
func UserActionCreateFn(dto *UserEntity, query QueryDSL) (*UserEntity, *IError) {
	// 1. Validate always
	if iError := UserValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	UserEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	UserEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	UserPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	UserRelationContentCreate(dto, query)
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
	UserAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(USER_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&UserEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func UserActionGetOne(query QueryDSL) (*UserEntity, *IError) {
    refl := reflect.ValueOf(&UserEntity{})
    item, err := GetOneEntity[UserEntity](query, refl)
    entityUserFormatter(item, query)
    return item, err
  }
  func UserActionQuery(query QueryDSL) ([]*UserEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&UserEntity{})
    items, meta, err := QueryEntitiesPointer[UserEntity](query, refl)
    for _, item := range items {
      entityUserFormatter(item, query)
    }
    return items, meta, err
  }
  func UserUpdateExec(dbref *gorm.DB, query QueryDSL, fields *UserEntity) (*UserEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = USER_EVENT_UPDATED
    UserEntityPreSanitize(fields, query)
    var item UserEntity
    q := dbref.
      Where(&UserEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    UserRelationContentUpdate(fields, query)
    UserPolyglotCreateHandler(fields, query)
    if ero := UserDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&UserEntity{UniqueId: uniqueId}).
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
  func UserActionUpdateFn(query QueryDSL, fields *UserEntity) (*UserEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := UserValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // UserRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      var item *UserEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *IError
        item, err = UserUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, CastToIError(vf)
    } else {
      dbref = query.Tx
      return UserUpdateExec(dbref, query, fields)
    }
  }
var UserWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire users ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
      ActionRequires: []PermissionInfo{PERM_ROOT_USER_DELETE},
    })
		count, _ := UserActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func UserActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&UserEntity{})
	query.ActionRequires = []PermissionInfo{PERM_ROOT_USER_DELETE}
	return RemoveEntity[UserEntity](query, refl)
}
func UserActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[UserEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'UserEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func UserActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[UserEntity]) (
    *BulkRecordRequest[UserEntity], *IError,
  ) {
    result := []*UserEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := UserActionUpdate(query, record)
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
func (x *UserEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var UserEntityMeta = TableMetaData{
	EntityName:    "User",
	ExportKey:    "users",
	TableNameInDb: "fb_user_entities",
	EntityObject:  &UserEntity{},
	ExportStream: UserActionExportT,
	ImportQuery: UserActionImport,
}
func UserActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[UserEntity](query, UserActionQuery, UserPreloadRelations)
}
func UserActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[UserEntity](query, UserActionQuery, UserPreloadRelations)
}
func UserActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content UserEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := UserActionCreate(&content, query)
	return err
}
var UserCommonCliFlags = []cli.Flag{
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
      Name:     "person-id",
      Required: false,
      Usage:    "person",
    },
}
var UserCommonInteractiveCliFlags = []CliInteractiveFlag{
}
var UserCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "person-id",
      Required: false,
      Usage:    "person",
    },
}
  var UserCreateCmd cli.Command = USER_ACTION_POST_ONE.ToCli()
  var UserCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_USER_CREATE},
      })
      entity := &UserEntity{}
      for _, item := range UserCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := UserActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var UserUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: UserCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_USER_UPDATE},
      })
      entity := CastUserFromCli(c)
      if entity, err := UserActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* UserEntity) FromCli(c *cli.Context) *UserEntity {
	return CastUserFromCli(c)
}
func CastUserFromCli (c *cli.Context) *UserEntity {
	template := &UserEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("person-id") {
        value := c.String("person-id")
        template.PersonId = &value
      }
	return template
}
  func UserSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      UserActionCreate,
      reflect.ValueOf(&UserEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func UserImportMocks() {
    SeederFromFSImport(
      QueryDSL{},
      UserActionCreate,
      reflect.ValueOf(&UserEntity{}).Elem(),
      &mocks.ViewsFs,
      []string{},
      false,
    )
  }
  func UserWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := UserActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "User", result)
    }
  }
var UserImportExportCommands = []cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_USER_CREATE},
      })
			UserActionSeeder(query, c.Int("count"))
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
				Value: "user-seeder.yml",
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
        ActionRequires: []PermissionInfo{PERM_ROOT_USER_CREATE},
      })
			UserActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "user-seeder-user.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of users, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]UserEntity{}
			ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
		cli.Command{
			Name:  "mocks",
			Usage: "Prints the list of mocks",
			Action: func(c *cli.Context) error {
				if entity, err := GetSeederFilenames(&mocks.ViewsFs, ""); err != nil {
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
				CommonCliImportEmbedCmd(c,
					UserActionCreate,
					reflect.ValueOf(&UserEntity{}).Elem(),
					&mocks.ViewsFs,
				)
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
			UserCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				UserActionCreate,
				reflect.ValueOf(&UserEntity{}).Elem(),
				c.String("file"),
        &SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_USER_CREATE},
				},
        func() UserEntity {
					v := CastUserFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var UserCliCommands []cli.Command = []cli.Command{
      USER_ACTION_QUERY.ToCli(),
      USER_ACTION_TABLE.ToCli(),
      UserCreateCmd,
      UserUpdateCmd,
      UserCreateInteractiveCmd,
      UserWipeCmd,
      GetCommonRemoveQuery(reflect.ValueOf(&UserEntity{}).Elem(), UserActionRemove),
  }
  func UserCliFn() cli.Command {
    UserCliCommands = append(UserCliCommands, UserImportExportCommands...)
    return cli.Command{
      Name:        "user",
      Description: "Users module actions (sample module to handle complex entities)",
      Usage:       "Manage the users who are in the current app (root only)",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: UserCliCommands,
    }
  }
var USER_ACTION_TABLE = Module2Action{
  Name:    "table",
  ActionAliases: []string{"t"},
  Flags:  CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: UserActionQuery,
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    CommonCliTableCmd2(c,
      UserActionQuery,
      security,
      reflect.ValueOf(&UserEntity{}).Elem(),
    )
    return nil
  },
}
var USER_ACTION_QUERY = Module2Action{
  Method: "GET",
  Url:    "/users",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_USER_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpQueryEntity(c, UserActionQuery)
    },
  },
  Format: "QUERY",
  Action: UserActionQuery,
  ResponseEntity: &[]UserEntity{},
  CliAction: func(c *cli.Context, security *SecurityModel) error {
		CommonCliQueryCmd2(
			c,
			UserActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionAliases: []string{"q"},
	Flags:         CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var USER_ACTION_EXPORT = Module2Action{
  Method: "GET",
  Url:    "/users/export",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_USER_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpStreamFileChannel(c, UserActionExport)
    },
  },
  Format: "QUERY",
  Action: UserActionExport,
  ResponseEntity: &[]UserEntity{},
}
var USER_ACTION_GET_ONE = Module2Action{
  Method: "GET",
  Url:    "/user/:uniqueId",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_USER_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpGetEntity(c, UserActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: UserActionGetOne,
  ResponseEntity: &UserEntity{},
}
var USER_ACTION_POST_ONE = Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new user",
  Flags: UserCommonCliFlags,
  Method: "POST",
  Url:    "/user",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_USER_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpPostEntity(c, UserActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    result, err := CliPostEntity(c, UserActionCreate, security)
    HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: UserActionCreate,
  Format: "POST_ONE",
  RequestEntity: &UserEntity{},
  ResponseEntity: &UserEntity{},
}
var USER_ACTION_PATCH = Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: UserCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/user",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_USER_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntity(c, UserActionUpdate)
    },
  },
  Action: UserActionUpdate,
  RequestEntity: &UserEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &UserEntity{},
}
var USER_ACTION_PATCH_BULK = Module2Action{
  Method: "PATCH",
  Url:    "/users",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_USER_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntities(c, UserActionBulkUpdate)
    },
  },
  Action: UserActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &BulkRecordRequest[UserEntity]{},
  ResponseEntity: &BulkRecordRequest[UserEntity]{},
}
var USER_ACTION_DELETE = Module2Action{
  Method: "DELETE",
  Url:    "/user",
  Format: "DELETE_DSL",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_USER_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpRemoveEntity(c, UserActionRemove)
    },
  },
  Action: UserActionRemove,
  RequestEntity: &DeleteRequest{},
  ResponseEntity: &DeleteResponse{},
  TargetEntity: &UserEntity{},
}
  /**
  *	Override this function on UserEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendUserRouter = func(r *[]Module2Action) {}
  func GetUserModule2Actions() []Module2Action {
    routes := []Module2Action{
      USER_ACTION_QUERY,
      USER_ACTION_EXPORT,
      USER_ACTION_GET_ONE,
      USER_ACTION_POST_ONE,
      USER_ACTION_PATCH,
      USER_ACTION_PATCH_BULK,
      USER_ACTION_DELETE,
    }
    // Append user defined functions
    AppendUserRouter(&routes)
    return routes
  }
  func CreateUserRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetUserModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, UserEntityJsonSchema, "user-http", "workspaces")
    WriteEntitySchema("UserEntity", UserEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_USER_DELETE = PermissionInfo{
  CompleteKey: "root/workspaces/user/delete",
  Name: "Delete user",
}
var PERM_ROOT_USER_CREATE = PermissionInfo{
  CompleteKey: "root/workspaces/user/create",
  Name: "Create user",
}
var PERM_ROOT_USER_UPDATE = PermissionInfo{
  CompleteKey: "root/workspaces/user/update",
  Name: "Update user",
}
var PERM_ROOT_USER_QUERY = PermissionInfo{
  CompleteKey: "root/workspaces/user/query",
  Name: "Query user",
}
var PERM_ROOT_USER = PermissionInfo{
  CompleteKey: "root/workspaces/user/*",
  Name: "Entire user actions (*)",
}
var ALL_USER_PERMISSIONS = []PermissionInfo{
	PERM_ROOT_USER_DELETE,
	PERM_ROOT_USER_CREATE,
	PERM_ROOT_USER_UPDATE,
	PERM_ROOT_USER_QUERY,
	PERM_ROOT_USER,
}