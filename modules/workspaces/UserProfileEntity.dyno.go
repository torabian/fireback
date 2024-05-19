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
var userProfileSeedersFs *embed.FS = nil
func ResetUserProfileSeeders(fs *embed.FS) {
	userProfileSeedersFs = fs
}
type UserProfileEntity struct {
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
    FirstName   *string `json:"firstName" yaml:"firstName"       `
    // Datenano also has a text representation
    LastName   *string `json:"lastName" yaml:"lastName"       `
    // Datenano also has a text representation
    Children []*UserProfileEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *UserProfileEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var UserProfilePreloadRelations []string = []string{}
var USER_PROFILE_EVENT_CREATED = "userProfile.created"
var USER_PROFILE_EVENT_UPDATED = "userProfile.updated"
var USER_PROFILE_EVENT_DELETED = "userProfile.deleted"
var USER_PROFILE_EVENTS = []string{
	USER_PROFILE_EVENT_CREATED,
	USER_PROFILE_EVENT_UPDATED,
	USER_PROFILE_EVENT_DELETED,
}
type UserProfileFieldMap struct {
		FirstName TranslatedString `yaml:"firstName"`
		LastName TranslatedString `yaml:"lastName"`
}
var UserProfileEntityMetaConfig map[string]int64 = map[string]int64{
}
var UserProfileEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&UserProfileEntity{}))
func entityUserProfileFormatter(dto *UserProfileEntity, query QueryDSL) {
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
func UserProfileMockEntity() *UserProfileEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &UserProfileEntity{
      FirstName : &stringHolder,
      LastName : &stringHolder,
	}
	return entity
}
func UserProfileActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := UserProfileMockEntity()
		_, err := UserProfileActionCreate(entity, query)
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
  func UserProfileActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*UserProfileEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &UserProfileEntity{
          FirstName: &tildaRef,
          LastName: &tildaRef,
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
  func UserProfileAssociationCreate(dto *UserProfileEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func UserProfileRelationContentCreate(dto *UserProfileEntity, query QueryDSL) error {
return nil
}
func UserProfileRelationContentUpdate(dto *UserProfileEntity, query QueryDSL) error {
	return nil
}
func UserProfilePolyglotCreateHandler(dto *UserProfileEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func UserProfileValidator(dto *UserProfileEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func UserProfileEntityPreSanitize(dto *UserProfileEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func UserProfileEntityBeforeCreateAppend(dto *UserProfileEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    UserProfileRecursiveAddUniqueId(dto, query)
  }
  func UserProfileRecursiveAddUniqueId(dto *UserProfileEntity, query QueryDSL) {
  }
func UserProfileActionBatchCreateFn(dtos []*UserProfileEntity, query QueryDSL) ([]*UserProfileEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*UserProfileEntity{}
		for _, item := range dtos {
			s, err := UserProfileActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func UserProfileDeleteEntireChildren(query QueryDSL, dto *UserProfileEntity) (*IError) {
  // intentionally removed this. It's hard to implement it, and probably wrong without
  // proper on delete cascade
  return nil
}
func UserProfileActionCreateFn(dto *UserProfileEntity, query QueryDSL) (*UserProfileEntity, *IError) {
	// 1. Validate always
	if iError := UserProfileValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	UserProfileEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	UserProfileEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	UserProfilePolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	UserProfileRelationContentCreate(dto, query)
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
	UserProfileAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(USER_PROFILE_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&UserProfileEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func UserProfileActionGetOne(query QueryDSL) (*UserProfileEntity, *IError) {
    refl := reflect.ValueOf(&UserProfileEntity{})
    item, err := GetOneEntity[UserProfileEntity](query, refl)
    entityUserProfileFormatter(item, query)
    return item, err
  }
  func UserProfileActionQuery(query QueryDSL) ([]*UserProfileEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&UserProfileEntity{})
    items, meta, err := QueryEntitiesPointer[UserProfileEntity](query, refl)
    for _, item := range items {
      entityUserProfileFormatter(item, query)
    }
    return items, meta, err
  }
  func UserProfileUpdateExec(dbref *gorm.DB, query QueryDSL, fields *UserProfileEntity) (*UserProfileEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = USER_PROFILE_EVENT_UPDATED
    UserProfileEntityPreSanitize(fields, query)
    var item UserProfileEntity
    q := dbref.
      Where(&UserProfileEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    UserProfileRelationContentUpdate(fields, query)
    UserProfilePolyglotCreateHandler(fields, query)
    if ero := UserProfileDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&UserProfileEntity{UniqueId: uniqueId}).
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
  func UserProfileActionUpdateFn(query QueryDSL, fields *UserProfileEntity) (*UserProfileEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := UserProfileValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // UserProfileRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      var item *UserProfileEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *IError
        item, err = UserProfileUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, CastToIError(vf)
    } else {
      dbref = query.Tx
      return UserProfileUpdateExec(dbref, query, fields)
    }
  }
var UserProfileWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire userprofiles ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
      ActionRequires: []PermissionInfo{PERM_ROOT_USER_PROFILE_DELETE},
    })
		count, _ := UserProfileActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func UserProfileActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&UserProfileEntity{})
	query.ActionRequires = []PermissionInfo{PERM_ROOT_USER_PROFILE_DELETE}
	return RemoveEntity[UserProfileEntity](query, refl)
}
func UserProfileActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[UserProfileEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'UserProfileEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func UserProfileActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[UserProfileEntity]) (
    *BulkRecordRequest[UserProfileEntity], *IError,
  ) {
    result := []*UserProfileEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := UserProfileActionUpdate(query, record)
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
func (x *UserProfileEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var UserProfileEntityMeta = TableMetaData{
	EntityName:    "UserProfile",
	ExportKey:    "user-profiles",
	TableNameInDb: "fb_user-profile_entities",
	EntityObject:  &UserProfileEntity{},
	ExportStream: UserProfileActionExportT,
	ImportQuery: UserProfileActionImport,
}
func UserProfileActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[UserProfileEntity](query, UserProfileActionQuery, UserProfilePreloadRelations)
}
func UserProfileActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[UserProfileEntity](query, UserProfileActionQuery, UserProfilePreloadRelations)
}
func UserProfileActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content UserProfileEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := UserProfileActionCreate(&content, query)
	return err
}
var UserProfileCommonCliFlags = []cli.Flag{
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
      Name:     "first-name",
      Required: false,
      Usage:    "firstName",
    },
    &cli.StringFlag{
      Name:     "last-name",
      Required: false,
      Usage:    "lastName",
    },
}
var UserProfileCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:     "firstName",
		StructField:     "FirstName",
		Required: false,
		Usage:    "firstName",
		Type: "string",
	},
	{
		Name:     "lastName",
		StructField:     "LastName",
		Required: false,
		Usage:    "lastName",
		Type: "string",
	},
}
var UserProfileCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "first-name",
      Required: false,
      Usage:    "firstName",
    },
    &cli.StringFlag{
      Name:     "last-name",
      Required: false,
      Usage:    "lastName",
    },
}
  var UserProfileCreateCmd cli.Command = USER_PROFILE_ACTION_POST_ONE.ToCli()
  var UserProfileCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_USER_PROFILE_CREATE},
      })
      entity := &UserProfileEntity{}
      for _, item := range UserProfileCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := UserProfileActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var UserProfileUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: UserProfileCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_USER_PROFILE_UPDATE},
      })
      entity := CastUserProfileFromCli(c)
      if entity, err := UserProfileActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* UserProfileEntity) FromCli(c *cli.Context) *UserProfileEntity {
	return CastUserProfileFromCli(c)
}
func CastUserProfileFromCli (c *cli.Context) *UserProfileEntity {
	template := &UserProfileEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("first-name") {
        value := c.String("first-name")
        template.FirstName = &value
      }
      if c.IsSet("last-name") {
        value := c.String("last-name")
        template.LastName = &value
      }
	return template
}
  func UserProfileSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      UserProfileActionCreate,
      reflect.ValueOf(&UserProfileEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func UserProfileWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := UserProfileActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "UserProfile", result)
    }
  }
var UserProfileImportExportCommands = []cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_USER_PROFILE_CREATE},
      })
			UserProfileActionSeeder(query, c.Int("count"))
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
				Value: "user-profile-seeder.yml",
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
        ActionRequires: []PermissionInfo{PERM_ROOT_USER_PROFILE_CREATE},
      })
			UserProfileActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "user-profile-seeder-user-profile.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of user-profiles, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]UserProfileEntity{}
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
			UserProfileCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				UserProfileActionCreate,
				reflect.ValueOf(&UserProfileEntity{}).Elem(),
				c.String("file"),
        &SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_USER_PROFILE_CREATE},
				},
        func() UserProfileEntity {
					v := CastUserProfileFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var UserProfileCliCommands []cli.Command = []cli.Command{
      USER_PROFILE_ACTION_QUERY.ToCli(),
      USER_PROFILE_ACTION_TABLE.ToCli(),
      UserProfileCreateCmd,
      UserProfileUpdateCmd,
      UserProfileCreateInteractiveCmd,
      UserProfileWipeCmd,
      GetCommonRemoveQuery(reflect.ValueOf(&UserProfileEntity{}).Elem(), UserProfileActionRemove),
  }
  func UserProfileCliFn() cli.Command {
    UserProfileCliCommands = append(UserProfileCliCommands, UserProfileImportExportCommands...)
    return cli.Command{
      Name:        "userProfile",
      Description: "UserProfiles module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: UserProfileCliCommands,
    }
  }
var USER_PROFILE_ACTION_TABLE = Module2Action{
  Name:    "table",
  ActionName: "table",
  ActionAliases: []string{"t"},
  Flags:  CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: UserProfileActionQuery,
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    CommonCliTableCmd2(c,
      UserProfileActionQuery,
      security,
      reflect.ValueOf(&UserProfileEntity{}).Elem(),
    )
    return nil
  },
}
var USER_PROFILE_ACTION_QUERY = Module2Action{
  Method: "GET",
  Url:    "/user-profiles",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_USER_PROFILE_QUERY},
  },
  Group: "userProfile",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpQueryEntity(c, UserProfileActionQuery)
    },
  },
  Format: "QUERY",
  Action: UserProfileActionQuery,
  ResponseEntity: &[]UserProfileEntity{},
  Out: Module2ActionBody{
		Entity: "UserProfileEntity",
	},
  CliAction: func(c *cli.Context, security *SecurityModel) error {
		CommonCliQueryCmd2(
			c,
			UserProfileActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionName:    "query",
	ActionAliases: []string{"q"},
	Flags:         CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var USER_PROFILE_ACTION_EXPORT = Module2Action{
  Method: "GET",
  Url:    "/user-profiles/export",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_USER_PROFILE_QUERY},
  },
  Group: "userProfile",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpStreamFileChannel(c, UserProfileActionExport)
    },
  },
  Format: "QUERY",
  Action: UserProfileActionExport,
  ResponseEntity: &[]UserProfileEntity{},
  Out: Module2ActionBody{
		Entity: "UserProfileEntity",
	},
}
var USER_PROFILE_ACTION_GET_ONE = Module2Action{
  Method: "GET",
  Url:    "/user-profile/:uniqueId",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_USER_PROFILE_QUERY},
  },
  Group: "userProfile",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpGetEntity(c, UserProfileActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: UserProfileActionGetOne,
  ResponseEntity: &UserProfileEntity{},
  Out: Module2ActionBody{
		Entity: "UserProfileEntity",
	},
}
var USER_PROFILE_ACTION_POST_ONE = Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new userProfile",
  Flags: UserProfileCommonCliFlags,
  Method: "POST",
  Url:    "/user-profile",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_USER_PROFILE_CREATE},
  },
  Group: "userProfile",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpPostEntity(c, UserProfileActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    result, err := CliPostEntity(c, UserProfileActionCreate, security)
    HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: UserProfileActionCreate,
  Format: "POST_ONE",
  RequestEntity: &UserProfileEntity{},
  ResponseEntity: &UserProfileEntity{},
  Out: Module2ActionBody{
		Entity: "UserProfileEntity",
	},
  In: Module2ActionBody{
		Entity: "UserProfileEntity",
	},
}
var USER_PROFILE_ACTION_PATCH = Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: UserProfileCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/user-profile",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_USER_PROFILE_UPDATE},
  },
  Group: "userProfile",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntity(c, UserProfileActionUpdate)
    },
  },
  Action: UserProfileActionUpdate,
  RequestEntity: &UserProfileEntity{},
  ResponseEntity: &UserProfileEntity{},
  Format: "PATCH_ONE",
  Out: Module2ActionBody{
		Entity: "UserProfileEntity",
	},
  In: Module2ActionBody{
		Entity: "UserProfileEntity",
	},
}
var USER_PROFILE_ACTION_PATCH_BULK = Module2Action{
  Method: "PATCH",
  Url:    "/user-profiles",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_USER_PROFILE_UPDATE},
  },
  Group: "userProfile",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntities(c, UserProfileActionBulkUpdate)
    },
  },
  Action: UserProfileActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &BulkRecordRequest[UserProfileEntity]{},
  ResponseEntity: &BulkRecordRequest[UserProfileEntity]{},
  Out: Module2ActionBody{
		Entity: "UserProfileEntity",
	},
  In: Module2ActionBody{
		Entity: "UserProfileEntity",
	},
}
var USER_PROFILE_ACTION_DELETE = Module2Action{
  Method: "DELETE",
  Url:    "/user-profile",
  Format: "DELETE_DSL",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_USER_PROFILE_DELETE},
  },
  Group: "userProfile",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpRemoveEntity(c, UserProfileActionRemove)
    },
  },
  Action: UserProfileActionRemove,
  RequestEntity: &DeleteRequest{},
  ResponseEntity: &DeleteResponse{},
  TargetEntity: &UserProfileEntity{},
}
  /**
  *	Override this function on UserProfileEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendUserProfileRouter = func(r *[]Module2Action) {}
  func GetUserProfileModule2Actions() []Module2Action {
    routes := []Module2Action{
      USER_PROFILE_ACTION_QUERY,
      USER_PROFILE_ACTION_EXPORT,
      USER_PROFILE_ACTION_GET_ONE,
      USER_PROFILE_ACTION_POST_ONE,
      USER_PROFILE_ACTION_PATCH,
      USER_PROFILE_ACTION_PATCH_BULK,
      USER_PROFILE_ACTION_DELETE,
    }
    // Append user defined functions
    AppendUserProfileRouter(&routes)
    return routes
  }
  func CreateUserProfileRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetUserProfileModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, UserProfileEntityJsonSchema, "user-profile-http", "workspaces")
    WriteEntitySchema("UserProfileEntity", UserProfileEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_USER_PROFILE_DELETE = PermissionInfo{
  CompleteKey: "root/workspaces/user-profile/delete",
  Name: "Delete user profile",
}
var PERM_ROOT_USER_PROFILE_CREATE = PermissionInfo{
  CompleteKey: "root/workspaces/user-profile/create",
  Name: "Create user profile",
}
var PERM_ROOT_USER_PROFILE_UPDATE = PermissionInfo{
  CompleteKey: "root/workspaces/user-profile/update",
  Name: "Update user profile",
}
var PERM_ROOT_USER_PROFILE_QUERY = PermissionInfo{
  CompleteKey: "root/workspaces/user-profile/query",
  Name: "Query user profile",
}
var PERM_ROOT_USER_PROFILE = PermissionInfo{
  CompleteKey: "root/workspaces/user-profile/*",
  Name: "Entire user profile actions (*)",
}
var ALL_USER_PROFILE_PERMISSIONS = []PermissionInfo{
	PERM_ROOT_USER_PROFILE_DELETE,
	PERM_ROOT_USER_PROFILE_CREATE,
	PERM_ROOT_USER_PROFILE_UPDATE,
	PERM_ROOT_USER_PROFILE_QUERY,
	PERM_ROOT_USER_PROFILE,
}
var UserProfileEntityBundle = EntityBundle{
	Permissions: ALL_USER_PROFILE_PERMISSIONS,
	CliCommands: []cli.Command{
		UserProfileCliFn(),
	},
	Actions: GetUserProfileModule2Actions(),
	AutoMigrationEntities: []interface{}{
		&UserProfileEntity{},
  	},
}