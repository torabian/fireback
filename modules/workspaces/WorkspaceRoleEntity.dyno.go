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
var workspaceRoleSeedersFs *embed.FS = nil
func ResetWorkspaceRoleSeeders(fs *embed.FS) {
	workspaceRoleSeedersFs = fs
}
type WorkspaceRoleEntity struct {
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
    UserWorkspace   *  UserWorkspaceEntity `json:"userWorkspace" yaml:"userWorkspace"    gorm:"foreignKey:UserWorkspaceId;references:UniqueId"     `
    // Datenano also has a text representation
        UserWorkspaceId *string `json:"userWorkspaceId" yaml:"userWorkspaceId" gorm:"index:workspacerole_idx,unique" `
    Role   *  RoleEntity `json:"role" yaml:"role"    gorm:"foreignKey:RoleId;references:UniqueId"     `
    // Datenano also has a text representation
        RoleId *string `json:"roleId" yaml:"roleId" gorm:"index:workspacerole_idx,unique" `
    Children []*WorkspaceRoleEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *WorkspaceRoleEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var WorkspaceRolePreloadRelations []string = []string{}
var WORKSPACE_ROLE_EVENT_CREATED = "workspaceRole.created"
var WORKSPACE_ROLE_EVENT_UPDATED = "workspaceRole.updated"
var WORKSPACE_ROLE_EVENT_DELETED = "workspaceRole.deleted"
var WORKSPACE_ROLE_EVENTS = []string{
	WORKSPACE_ROLE_EVENT_CREATED,
	WORKSPACE_ROLE_EVENT_UPDATED,
	WORKSPACE_ROLE_EVENT_DELETED,
}
type WorkspaceRoleFieldMap struct {
		UserWorkspace TranslatedString `yaml:"userWorkspace"`
		Role TranslatedString `yaml:"role"`
}
var WorkspaceRoleEntityMetaConfig map[string]int64 = map[string]int64{
}
var WorkspaceRoleEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&WorkspaceRoleEntity{}))
func entityWorkspaceRoleFormatter(dto *WorkspaceRoleEntity, query QueryDSL) {
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
func WorkspaceRoleMockEntity() *WorkspaceRoleEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &WorkspaceRoleEntity{
	}
	return entity
}
func WorkspaceRoleActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := WorkspaceRoleMockEntity()
		_, err := WorkspaceRoleActionCreate(entity, query)
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
  func WorkspaceRoleActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*WorkspaceRoleEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &WorkspaceRoleEntity{
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
  func WorkspaceRoleAssociationCreate(dto *WorkspaceRoleEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func WorkspaceRoleRelationContentCreate(dto *WorkspaceRoleEntity, query QueryDSL) error {
return nil
}
func WorkspaceRoleRelationContentUpdate(dto *WorkspaceRoleEntity, query QueryDSL) error {
	return nil
}
func WorkspaceRolePolyglotCreateHandler(dto *WorkspaceRoleEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func WorkspaceRoleValidator(dto *WorkspaceRoleEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func WorkspaceRoleEntityPreSanitize(dto *WorkspaceRoleEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func WorkspaceRoleEntityBeforeCreateAppend(dto *WorkspaceRoleEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    WorkspaceRoleRecursiveAddUniqueId(dto, query)
  }
  func WorkspaceRoleRecursiveAddUniqueId(dto *WorkspaceRoleEntity, query QueryDSL) {
  }
func WorkspaceRoleActionBatchCreateFn(dtos []*WorkspaceRoleEntity, query QueryDSL) ([]*WorkspaceRoleEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*WorkspaceRoleEntity{}
		for _, item := range dtos {
			s, err := WorkspaceRoleActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func WorkspaceRoleDeleteEntireChildren(query QueryDSL, dto *WorkspaceRoleEntity) (*IError) {
  // intentionally removed this. It's hard to implement it, and probably wrong without
  // proper on delete cascade
  return nil
}
func WorkspaceRoleActionCreateFn(dto *WorkspaceRoleEntity, query QueryDSL) (*WorkspaceRoleEntity, *IError) {
	// 1. Validate always
	if iError := WorkspaceRoleValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	WorkspaceRoleEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	WorkspaceRoleEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	WorkspaceRolePolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	WorkspaceRoleRelationContentCreate(dto, query)
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
	WorkspaceRoleAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(WORKSPACE_ROLE_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&WorkspaceRoleEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func WorkspaceRoleActionGetOne(query QueryDSL) (*WorkspaceRoleEntity, *IError) {
    refl := reflect.ValueOf(&WorkspaceRoleEntity{})
    item, err := GetOneEntity[WorkspaceRoleEntity](query, refl)
    entityWorkspaceRoleFormatter(item, query)
    return item, err
  }
  func WorkspaceRoleActionQuery(query QueryDSL) ([]*WorkspaceRoleEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&WorkspaceRoleEntity{})
    items, meta, err := QueryEntitiesPointer[WorkspaceRoleEntity](query, refl)
    for _, item := range items {
      entityWorkspaceRoleFormatter(item, query)
    }
    return items, meta, err
  }
  func WorkspaceRoleUpdateExec(dbref *gorm.DB, query QueryDSL, fields *WorkspaceRoleEntity) (*WorkspaceRoleEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = WORKSPACE_ROLE_EVENT_UPDATED
    WorkspaceRoleEntityPreSanitize(fields, query)
    var item WorkspaceRoleEntity
    q := dbref.
      Where(&WorkspaceRoleEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    WorkspaceRoleRelationContentUpdate(fields, query)
    WorkspaceRolePolyglotCreateHandler(fields, query)
    if ero := WorkspaceRoleDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&WorkspaceRoleEntity{UniqueId: uniqueId}).
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
  func WorkspaceRoleActionUpdateFn(query QueryDSL, fields *WorkspaceRoleEntity) (*WorkspaceRoleEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := WorkspaceRoleValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // WorkspaceRoleRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      var item *WorkspaceRoleEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *IError
        item, err = WorkspaceRoleUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, CastToIError(vf)
    } else {
      dbref = query.Tx
      return WorkspaceRoleUpdateExec(dbref, query, fields)
    }
  }
var WorkspaceRoleWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire workspaceroles ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
      ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_ROLE_DELETE},
    })
		count, _ := WorkspaceRoleActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func WorkspaceRoleActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&WorkspaceRoleEntity{})
	query.ActionRequires = []PermissionInfo{PERM_ROOT_WORKSPACE_ROLE_DELETE}
	return RemoveEntity[WorkspaceRoleEntity](query, refl)
}
func WorkspaceRoleActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[WorkspaceRoleEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'WorkspaceRoleEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func WorkspaceRoleActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[WorkspaceRoleEntity]) (
    *BulkRecordRequest[WorkspaceRoleEntity], *IError,
  ) {
    result := []*WorkspaceRoleEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := WorkspaceRoleActionUpdate(query, record)
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
func (x *WorkspaceRoleEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var WorkspaceRoleEntityMeta = TableMetaData{
	EntityName:    "WorkspaceRole",
	ExportKey:    "workspace-roles",
	TableNameInDb: "fb_workspace-role_entities",
	EntityObject:  &WorkspaceRoleEntity{},
	ExportStream: WorkspaceRoleActionExportT,
	ImportQuery: WorkspaceRoleActionImport,
}
func WorkspaceRoleActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[WorkspaceRoleEntity](query, WorkspaceRoleActionQuery, WorkspaceRolePreloadRelations)
}
func WorkspaceRoleActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[WorkspaceRoleEntity](query, WorkspaceRoleActionQuery, WorkspaceRolePreloadRelations)
}
func WorkspaceRoleActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content WorkspaceRoleEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := WorkspaceRoleActionCreate(&content, query)
	return err
}
var WorkspaceRoleCommonCliFlags = []cli.Flag{
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
      Name:     "user-workspace-id",
      Required: false,
      Usage:    "userWorkspace",
    },
    &cli.StringFlag{
      Name:     "role-id",
      Required: false,
      Usage:    "role",
    },
}
var WorkspaceRoleCommonInteractiveCliFlags = []CliInteractiveFlag{
}
var WorkspaceRoleCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "user-workspace-id",
      Required: false,
      Usage:    "userWorkspace",
    },
    &cli.StringFlag{
      Name:     "role-id",
      Required: false,
      Usage:    "role",
    },
}
  var WorkspaceRoleCreateCmd cli.Command = WORKSPACE_ROLE_ACTION_POST_ONE.ToCli()
  var WorkspaceRoleCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_ROLE_CREATE},
      })
      entity := &WorkspaceRoleEntity{}
      for _, item := range WorkspaceRoleCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := WorkspaceRoleActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var WorkspaceRoleUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: WorkspaceRoleCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_ROLE_UPDATE},
      })
      entity := CastWorkspaceRoleFromCli(c)
      if entity, err := WorkspaceRoleActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* WorkspaceRoleEntity) FromCli(c *cli.Context) *WorkspaceRoleEntity {
	return CastWorkspaceRoleFromCli(c)
}
func CastWorkspaceRoleFromCli (c *cli.Context) *WorkspaceRoleEntity {
	template := &WorkspaceRoleEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("user-workspace-id") {
        value := c.String("user-workspace-id")
        template.UserWorkspaceId = &value
      }
      if c.IsSet("role-id") {
        value := c.String("role-id")
        template.RoleId = &value
      }
	return template
}
  func WorkspaceRoleSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      WorkspaceRoleActionCreate,
      reflect.ValueOf(&WorkspaceRoleEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func WorkspaceRoleWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := WorkspaceRoleActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "WorkspaceRole", result)
    }
  }
var WorkspaceRoleImportExportCommands = []cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_ROLE_CREATE},
      })
			WorkspaceRoleActionSeeder(query, c.Int("count"))
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
				Value: "workspace-role-seeder.yml",
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
        ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_ROLE_CREATE},
      })
			WorkspaceRoleActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "workspace-role-seeder-workspace-role.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of workspace-roles, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]WorkspaceRoleEntity{}
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
			WorkspaceRoleCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				WorkspaceRoleActionCreate,
				reflect.ValueOf(&WorkspaceRoleEntity{}).Elem(),
				c.String("file"),
        &SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_ROLE_CREATE},
				},
        func() WorkspaceRoleEntity {
					v := CastWorkspaceRoleFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var WorkspaceRoleCliCommands []cli.Command = []cli.Command{
      WORKSPACE_ROLE_ACTION_QUERY.ToCli(),
      WORKSPACE_ROLE_ACTION_TABLE.ToCli(),
      WorkspaceRoleCreateCmd,
      WorkspaceRoleUpdateCmd,
      WorkspaceRoleCreateInteractiveCmd,
      WorkspaceRoleWipeCmd,
      GetCommonRemoveQuery(reflect.ValueOf(&WorkspaceRoleEntity{}).Elem(), WorkspaceRoleActionRemove),
  }
  func WorkspaceRoleCliFn() cli.Command {
    WorkspaceRoleCliCommands = append(WorkspaceRoleCliCommands, WorkspaceRoleImportExportCommands...)
    return cli.Command{
      Name:        "workspaceRole",
      ShortName:   "role",
      Description: "WorkspaceRoles module actions (sample module to handle complex entities)",
      Usage:       "Manage roles assigned to an specific workspace or created by the workspace itself",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: WorkspaceRoleCliCommands,
    }
  }
var WORKSPACE_ROLE_ACTION_TABLE = Module2Action{
  Name:    "table",
  ActionName: "table",
  ActionAliases: []string{"t"},
  Flags:  CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: WorkspaceRoleActionQuery,
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    CommonCliTableCmd2(c,
      WorkspaceRoleActionQuery,
      security,
      reflect.ValueOf(&WorkspaceRoleEntity{}).Elem(),
    )
    return nil
  },
}
var WORKSPACE_ROLE_ACTION_QUERY = Module2Action{
  Method: "GET",
  Url:    "/workspace-roles",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_ROLE_QUERY},
  },
  Group: "workspaceRole",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpQueryEntity(c, WorkspaceRoleActionQuery)
    },
  },
  Format: "QUERY",
  Action: WorkspaceRoleActionQuery,
  ResponseEntity: &[]WorkspaceRoleEntity{},
  Out: Module2ActionBody{
		Entity: "WorkspaceRoleEntity",
	},
  CliAction: func(c *cli.Context, security *SecurityModel) error {
		CommonCliQueryCmd2(
			c,
			WorkspaceRoleActionQuery,
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
var WORKSPACE_ROLE_ACTION_EXPORT = Module2Action{
  Method: "GET",
  Url:    "/workspace-roles/export",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_ROLE_QUERY},
  },
  Group: "workspaceRole",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpStreamFileChannel(c, WorkspaceRoleActionExport)
    },
  },
  Format: "QUERY",
  Action: WorkspaceRoleActionExport,
  ResponseEntity: &[]WorkspaceRoleEntity{},
  Out: Module2ActionBody{
		Entity: "WorkspaceRoleEntity",
	},
}
var WORKSPACE_ROLE_ACTION_GET_ONE = Module2Action{
  Method: "GET",
  Url:    "/workspace-role/:uniqueId",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_ROLE_QUERY},
  },
  Group: "workspaceRole",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpGetEntity(c, WorkspaceRoleActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: WorkspaceRoleActionGetOne,
  ResponseEntity: &WorkspaceRoleEntity{},
  Out: Module2ActionBody{
		Entity: "WorkspaceRoleEntity",
	},
}
var WORKSPACE_ROLE_ACTION_POST_ONE = Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new workspaceRole",
  Flags: WorkspaceRoleCommonCliFlags,
  Method: "POST",
  Url:    "/workspace-role",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_ROLE_CREATE},
  },
  Group: "workspaceRole",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpPostEntity(c, WorkspaceRoleActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    result, err := CliPostEntity(c, WorkspaceRoleActionCreate, security)
    HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: WorkspaceRoleActionCreate,
  Format: "POST_ONE",
  RequestEntity: &WorkspaceRoleEntity{},
  ResponseEntity: &WorkspaceRoleEntity{},
  Out: Module2ActionBody{
		Entity: "WorkspaceRoleEntity",
	},
  In: Module2ActionBody{
		Entity: "WorkspaceRoleEntity",
	},
}
var WORKSPACE_ROLE_ACTION_PATCH = Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: WorkspaceRoleCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/workspace-role",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_ROLE_UPDATE},
  },
  Group: "workspaceRole",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntity(c, WorkspaceRoleActionUpdate)
    },
  },
  Action: WorkspaceRoleActionUpdate,
  RequestEntity: &WorkspaceRoleEntity{},
  ResponseEntity: &WorkspaceRoleEntity{},
  Format: "PATCH_ONE",
  Out: Module2ActionBody{
		Entity: "WorkspaceRoleEntity",
	},
  In: Module2ActionBody{
		Entity: "WorkspaceRoleEntity",
	},
}
var WORKSPACE_ROLE_ACTION_PATCH_BULK = Module2Action{
  Method: "PATCH",
  Url:    "/workspace-roles",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_ROLE_UPDATE},
  },
  Group: "workspaceRole",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntities(c, WorkspaceRoleActionBulkUpdate)
    },
  },
  Action: WorkspaceRoleActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &BulkRecordRequest[WorkspaceRoleEntity]{},
  ResponseEntity: &BulkRecordRequest[WorkspaceRoleEntity]{},
  Out: Module2ActionBody{
		Entity: "WorkspaceRoleEntity",
	},
  In: Module2ActionBody{
		Entity: "WorkspaceRoleEntity",
	},
}
var WORKSPACE_ROLE_ACTION_DELETE = Module2Action{
  Method: "DELETE",
  Url:    "/workspace-role",
  Format: "DELETE_DSL",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_ROLE_DELETE},
  },
  Group: "workspaceRole",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpRemoveEntity(c, WorkspaceRoleActionRemove)
    },
  },
  Action: WorkspaceRoleActionRemove,
  RequestEntity: &DeleteRequest{},
  ResponseEntity: &DeleteResponse{},
  TargetEntity: &WorkspaceRoleEntity{},
}
  /**
  *	Override this function on WorkspaceRoleEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendWorkspaceRoleRouter = func(r *[]Module2Action) {}
  func GetWorkspaceRoleModule2Actions() []Module2Action {
    routes := []Module2Action{
      WORKSPACE_ROLE_ACTION_QUERY,
      WORKSPACE_ROLE_ACTION_EXPORT,
      WORKSPACE_ROLE_ACTION_GET_ONE,
      WORKSPACE_ROLE_ACTION_POST_ONE,
      WORKSPACE_ROLE_ACTION_PATCH,
      WORKSPACE_ROLE_ACTION_PATCH_BULK,
      WORKSPACE_ROLE_ACTION_DELETE,
    }
    // Append user defined functions
    AppendWorkspaceRoleRouter(&routes)
    return routes
  }
  func CreateWorkspaceRoleRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetWorkspaceRoleModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, WorkspaceRoleEntityJsonSchema, "workspace-role-http", "workspaces")
    WriteEntitySchema("WorkspaceRoleEntity", WorkspaceRoleEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_WORKSPACE_ROLE_DELETE = PermissionInfo{
  CompleteKey: "root/workspaces/workspace-role/delete",
  Name: "Delete workspace role",
}
var PERM_ROOT_WORKSPACE_ROLE_CREATE = PermissionInfo{
  CompleteKey: "root/workspaces/workspace-role/create",
  Name: "Create workspace role",
}
var PERM_ROOT_WORKSPACE_ROLE_UPDATE = PermissionInfo{
  CompleteKey: "root/workspaces/workspace-role/update",
  Name: "Update workspace role",
}
var PERM_ROOT_WORKSPACE_ROLE_QUERY = PermissionInfo{
  CompleteKey: "root/workspaces/workspace-role/query",
  Name: "Query workspace role",
}
var PERM_ROOT_WORKSPACE_ROLE = PermissionInfo{
  CompleteKey: "root/workspaces/workspace-role/*",
  Name: "Entire workspace role actions (*)",
}
var ALL_WORKSPACE_ROLE_PERMISSIONS = []PermissionInfo{
	PERM_ROOT_WORKSPACE_ROLE_DELETE,
	PERM_ROOT_WORKSPACE_ROLE_CREATE,
	PERM_ROOT_WORKSPACE_ROLE_UPDATE,
	PERM_ROOT_WORKSPACE_ROLE_QUERY,
	PERM_ROOT_WORKSPACE_ROLE,
}