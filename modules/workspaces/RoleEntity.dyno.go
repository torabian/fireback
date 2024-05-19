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
	seeders "github.com/torabian/fireback/modules/workspaces/seeders/Role"
	metas "github.com/torabian/fireback/modules/workspaces/metas"
)
var roleSeedersFs = &seeders.ViewsFs
func ResetRoleSeeders(fs *embed.FS) {
	roleSeedersFs = fs
}
type RoleEntity struct {
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
    Name   *string `json:"name" yaml:"name"  validate:"required,omitempty,min=1,max=200"       `
    // Datenano also has a text representation
    Capabilities   []*  CapabilityEntity `json:"capabilities" yaml:"capabilities"    gorm:"many2many:role_capabilities;foreignKey:UniqueId;references:UniqueId"     `
    // Datenano also has a text representation
    CapabilitiesListId []string `json:"capabilitiesListId" yaml:"capabilitiesListId" gorm:"-" sql:"-"`
    Children []*RoleEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *RoleEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var RolePreloadRelations []string = []string{}
var ROLE_EVENT_CREATED = "role.created"
var ROLE_EVENT_UPDATED = "role.updated"
var ROLE_EVENT_DELETED = "role.deleted"
var ROLE_EVENTS = []string{
	ROLE_EVENT_CREATED,
	ROLE_EVENT_UPDATED,
	ROLE_EVENT_DELETED,
}
type RoleFieldMap struct {
		Name TranslatedString `yaml:"name"`
		Capabilities TranslatedString `yaml:"capabilities"`
}
var RoleEntityMetaConfig map[string]int64 = map[string]int64{
}
var RoleEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&RoleEntity{}))
func entityRoleFormatter(dto *RoleEntity, query QueryDSL) {
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
func RoleMockEntity() *RoleEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &RoleEntity{
      Name : &stringHolder,
	}
	return entity
}
func RoleActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := RoleMockEntity()
		_, err := RoleActionCreate(entity, query)
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
  func RoleActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*RoleEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &RoleEntity{
          Name: &tildaRef,
          CapabilitiesListId: []string{"~"},
          Capabilities: []*CapabilityEntity{{}},
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
  func RoleAssociationCreate(dto *RoleEntity, query QueryDSL) error {
      {
        if dto.CapabilitiesListId != nil && len(dto.CapabilitiesListId) > 0 {
          var items []CapabilityEntity
          err := query.Tx.Where(dto.CapabilitiesListId).Find(&items).Error
          if err != nil {
              return err
          }
          err = query.Tx.Model(dto).Association("Capabilities").Replace(items)
          if err != nil {
              return err
          }
        }
      }
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func RoleRelationContentCreate(dto *RoleEntity, query QueryDSL) error {
return nil
}
func RoleRelationContentUpdate(dto *RoleEntity, query QueryDSL) error {
	return nil
}
func RolePolyglotCreateHandler(dto *RoleEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func RoleValidator(dto *RoleEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func RoleEntityPreSanitize(dto *RoleEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func RoleEntityBeforeCreateAppend(dto *RoleEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    RoleRecursiveAddUniqueId(dto, query)
  }
  func RoleRecursiveAddUniqueId(dto *RoleEntity, query QueryDSL) {
  }
func RoleActionBatchCreateFn(dtos []*RoleEntity, query QueryDSL) ([]*RoleEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*RoleEntity{}
		for _, item := range dtos {
			s, err := RoleActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func RoleDeleteEntireChildren(query QueryDSL, dto *RoleEntity) (*IError) {
  // intentionally removed this. It's hard to implement it, and probably wrong without
  // proper on delete cascade
  return nil
}
func RoleActionCreateFn(dto *RoleEntity, query QueryDSL) (*RoleEntity, *IError) {
	// 1. Validate always
	if iError := RoleValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	RoleEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	RoleEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	RolePolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	RoleRelationContentCreate(dto, query)
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
	RoleAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(ROLE_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&RoleEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func RoleActionGetOne(query QueryDSL) (*RoleEntity, *IError) {
    refl := reflect.ValueOf(&RoleEntity{})
    item, err := GetOneEntity[RoleEntity](query, refl)
    entityRoleFormatter(item, query)
    return item, err
  }
  func RoleActionQuery(query QueryDSL) ([]*RoleEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&RoleEntity{})
    items, meta, err := QueryEntitiesPointer[RoleEntity](query, refl)
    for _, item := range items {
      entityRoleFormatter(item, query)
    }
    return items, meta, err
  }
  func RoleUpdateExec(dbref *gorm.DB, query QueryDSL, fields *RoleEntity) (*RoleEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = ROLE_EVENT_UPDATED
    RoleEntityPreSanitize(fields, query)
    var item RoleEntity
    q := dbref.
      Where(&RoleEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    RoleRelationContentUpdate(fields, query)
    RolePolyglotCreateHandler(fields, query)
    if ero := RoleDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
        if fields.CapabilitiesListId  != nil {
          var items []CapabilityEntity
          if len(fields.CapabilitiesListId ) > 0 {
            dbref.
              Where(&fields.CapabilitiesListId ).
              Find(&items)
          }
          dbref.
            Model(&RoleEntity{UniqueId: uniqueId}).
            Association("Capabilities").
            Replace(&items)
        }
    err = dbref.
      Preload(clause.Associations).
      Where(&RoleEntity{UniqueId: uniqueId}).
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
  func RoleActionUpdateFn(query QueryDSL, fields *RoleEntity) (*RoleEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := RoleValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // RoleRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      var item *RoleEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *IError
        item, err = RoleUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, CastToIError(vf)
    } else {
      dbref = query.Tx
      return RoleUpdateExec(dbref, query, fields)
    }
  }
var RoleWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire roles ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
      ActionRequires: []PermissionInfo{PERM_ROOT_ROLE_DELETE},
    })
		count, _ := RoleActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func RoleActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&RoleEntity{})
	query.ActionRequires = []PermissionInfo{PERM_ROOT_ROLE_DELETE}
	return RemoveEntity[RoleEntity](query, refl)
}
func RoleActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[RoleEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'RoleEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func RoleActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[RoleEntity]) (
    *BulkRecordRequest[RoleEntity], *IError,
  ) {
    result := []*RoleEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := RoleActionUpdate(query, record)
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
func (x *RoleEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var RoleEntityMeta = TableMetaData{
	EntityName:    "Role",
	ExportKey:    "roles",
	TableNameInDb: "fb_role_entities",
	EntityObject:  &RoleEntity{},
	ExportStream: RoleActionExportT,
	ImportQuery: RoleActionImport,
}
func RoleActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[RoleEntity](query, RoleActionQuery, RolePreloadRelations)
}
func RoleActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[RoleEntity](query, RoleActionQuery, RolePreloadRelations)
}
func RoleActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content RoleEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := RoleActionCreate(&content, query)
	return err
}
var RoleCommonCliFlags = []cli.Flag{
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
    &cli.StringSliceFlag{
      Name:     "capabilities",
      Required: false,
      Usage:    "capabilities",
    },
}
var RoleCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:     "name",
		StructField:     "Name",
		Required: true,
		Usage:    "name",
		Type: "string",
	},
}
var RoleCommonCliFlagsOptional = []cli.Flag{
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
    &cli.StringSliceFlag{
      Name:     "capabilities",
      Required: false,
      Usage:    "capabilities",
    },
}
  var RoleCreateCmd cli.Command = ROLE_ACTION_POST_ONE.ToCli()
  var RoleCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_ROLE_CREATE},
      })
      entity := &RoleEntity{}
      for _, item := range RoleCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := RoleActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var RoleUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: RoleCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_ROLE_UPDATE},
      })
      entity := CastRoleFromCli(c)
      if entity, err := RoleActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* RoleEntity) FromCli(c *cli.Context) *RoleEntity {
	return CastRoleFromCli(c)
}
func CastRoleFromCli (c *cli.Context) *RoleEntity {
	template := &RoleEntity{}
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
      if c.IsSet("capabilities") {
        value := c.String("capabilities")
        template.CapabilitiesListId = strings.Split(value, ",")
      }
	return template
}
  func RoleSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      RoleActionCreate,
      reflect.ValueOf(&RoleEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func RoleSyncSeeders() {
    SeederFromFSImport(
      QueryDSL{WorkspaceId: USER_SYSTEM},
      RoleActionCreate,
      reflect.ValueOf(&RoleEntity{}).Elem(),
      roleSeedersFs,
      []string{},
      true,
    )
  }
  func RoleWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := RoleActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "Role", result)
    }
  }
var RoleImportExportCommands = []cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_ROLE_CREATE},
      })
			RoleActionSeeder(query, c.Int("count"))
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
				Value: "role-seeder.yml",
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
        ActionRequires: []PermissionInfo{PERM_ROOT_ROLE_CREATE},
      })
			RoleActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "role-seeder-role.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of roles, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]RoleEntity{}
			ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
	cli.Command{
		Name:  "list",
		Usage: "Prints the list of files attached to this module for syncing or bootstrapping project",
		Action: func(c *cli.Context) error {
			if entity, err := GetSeederFilenames(roleSeedersFs, ""); err != nil {
				fmt.Println(err.Error())
			} else {
				f, _ := json.MarshalIndent(entity, "", "  ")
				fmt.Println(string(f))
			}
			return nil
		},
	},
	cli.Command{
		Name:  "sync",
		Usage: "Tries to sync the embedded content into the database, the list could be seen by 'list' command",
		Action: func(c *cli.Context) error {
			CommonCliImportEmbedCmd(c,
				RoleActionCreate,
				reflect.ValueOf(&RoleEntity{}).Elem(),
				roleSeedersFs,
			)
			return nil
		},
	},
	cli.Command{
		Name:    "export",
		Aliases: []string{"e"},
		Flags: append(CommonQueryFlags,
			&cli.StringFlag{
				Name:     "file",
				Usage:    "The address of file you want the csv/yaml/json be exported to",
				Required: true,
			}),
		Usage: "Exports a query results into the csv/yaml/json format",
		Action: func(c *cli.Context) error {
			CommonCliExportCmd(c,
				RoleActionQuery,
				reflect.ValueOf(&RoleEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"RoleFieldMap.yml",
				RolePreloadRelations,
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
			RoleCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				RoleActionCreate,
				reflect.ValueOf(&RoleEntity{}).Elem(),
				c.String("file"),
        &SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_ROLE_CREATE},
				},
        func() RoleEntity {
					v := CastRoleFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var RoleCliCommands []cli.Command = []cli.Command{
      ROLE_ACTION_QUERY.ToCli(),
      ROLE_ACTION_TABLE.ToCli(),
      RoleCreateCmd,
      RoleUpdateCmd,
      RoleCreateInteractiveCmd,
      RoleWipeCmd,
      GetCommonRemoveQuery(reflect.ValueOf(&RoleEntity{}).Elem(), RoleActionRemove),
  }
  func RoleCliFn() cli.Command {
    RoleCliCommands = append(RoleCliCommands, RoleImportExportCommands...)
    return cli.Command{
      Name:        "role",
      Description: "Roles module actions (sample module to handle complex entities)",
      Usage:       "Manage roles within the workspaces, or root configuration",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: RoleCliCommands,
    }
  }
var ROLE_ACTION_TABLE = Module2Action{
  Name:    "table",
  ActionName: "table",
  ActionAliases: []string{"t"},
  Flags:  CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: RoleActionQuery,
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    CommonCliTableCmd2(c,
      RoleActionQuery,
      security,
      reflect.ValueOf(&RoleEntity{}).Elem(),
    )
    return nil
  },
}
var ROLE_ACTION_QUERY = Module2Action{
  Method: "GET",
  Url:    "/roles",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_ROLE_QUERY},
  },
  Group: "role",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpQueryEntity(c, RoleActionQuery)
    },
  },
  Format: "QUERY",
  Action: RoleActionQuery,
  ResponseEntity: &[]RoleEntity{},
  Out: Module2ActionBody{
		Entity: "RoleEntity",
	},
  CliAction: func(c *cli.Context, security *SecurityModel) error {
		CommonCliQueryCmd2(
			c,
			RoleActionQuery,
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
var ROLE_ACTION_EXPORT = Module2Action{
  Method: "GET",
  Url:    "/roles/export",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_ROLE_QUERY},
  },
  Group: "role",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpStreamFileChannel(c, RoleActionExport)
    },
  },
  Format: "QUERY",
  Action: RoleActionExport,
  ResponseEntity: &[]RoleEntity{},
  Out: Module2ActionBody{
		Entity: "RoleEntity",
	},
}
var ROLE_ACTION_GET_ONE = Module2Action{
  Method: "GET",
  Url:    "/role/:uniqueId",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_ROLE_QUERY},
  },
  Group: "role",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpGetEntity(c, RoleActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: RoleActionGetOne,
  ResponseEntity: &RoleEntity{},
  Out: Module2ActionBody{
		Entity: "RoleEntity",
	},
}
var ROLE_ACTION_POST_ONE = Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new role",
  Flags: RoleCommonCliFlags,
  Method: "POST",
  Url:    "/role",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_ROLE_CREATE},
  },
  Group: "role",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpPostEntity(c, RoleActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    result, err := CliPostEntity(c, RoleActionCreate, security)
    HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: RoleActionCreate,
  Format: "POST_ONE",
  RequestEntity: &RoleEntity{},
  ResponseEntity: &RoleEntity{},
  Out: Module2ActionBody{
		Entity: "RoleEntity",
	},
  In: Module2ActionBody{
		Entity: "RoleEntity",
	},
}
var ROLE_ACTION_PATCH = Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: RoleCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/role",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_ROLE_UPDATE},
  },
  Group: "role",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntity(c, RoleActionUpdate)
    },
  },
  Action: RoleActionUpdate,
  RequestEntity: &RoleEntity{},
  ResponseEntity: &RoleEntity{},
  Format: "PATCH_ONE",
  Out: Module2ActionBody{
		Entity: "RoleEntity",
	},
  In: Module2ActionBody{
		Entity: "RoleEntity",
	},
}
var ROLE_ACTION_PATCH_BULK = Module2Action{
  Method: "PATCH",
  Url:    "/roles",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_ROLE_UPDATE},
  },
  Group: "role",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntities(c, RoleActionBulkUpdate)
    },
  },
  Action: RoleActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &BulkRecordRequest[RoleEntity]{},
  ResponseEntity: &BulkRecordRequest[RoleEntity]{},
  Out: Module2ActionBody{
		Entity: "RoleEntity",
	},
  In: Module2ActionBody{
		Entity: "RoleEntity",
	},
}
var ROLE_ACTION_DELETE = Module2Action{
  Method: "DELETE",
  Url:    "/role",
  Format: "DELETE_DSL",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_ROLE_DELETE},
  },
  Group: "role",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpRemoveEntity(c, RoleActionRemove)
    },
  },
  Action: RoleActionRemove,
  RequestEntity: &DeleteRequest{},
  ResponseEntity: &DeleteResponse{},
  TargetEntity: &RoleEntity{},
}
  /**
  *	Override this function on RoleEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendRoleRouter = func(r *[]Module2Action) {}
  func GetRoleModule2Actions() []Module2Action {
    routes := []Module2Action{
      ROLE_ACTION_QUERY,
      ROLE_ACTION_EXPORT,
      ROLE_ACTION_GET_ONE,
      ROLE_ACTION_POST_ONE,
      ROLE_ACTION_PATCH,
      ROLE_ACTION_PATCH_BULK,
      ROLE_ACTION_DELETE,
    }
    // Append user defined functions
    AppendRoleRouter(&routes)
    return routes
  }
  func CreateRoleRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetRoleModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, RoleEntityJsonSchema, "role-http", "workspaces")
    WriteEntitySchema("RoleEntity", RoleEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_ROLE_DELETE = PermissionInfo{
  CompleteKey: "root/workspaces/role/delete",
  Name: "Delete role",
}
var PERM_ROOT_ROLE_CREATE = PermissionInfo{
  CompleteKey: "root/workspaces/role/create",
  Name: "Create role",
}
var PERM_ROOT_ROLE_UPDATE = PermissionInfo{
  CompleteKey: "root/workspaces/role/update",
  Name: "Update role",
}
var PERM_ROOT_ROLE_QUERY = PermissionInfo{
  CompleteKey: "root/workspaces/role/query",
  Name: "Query role",
}
var PERM_ROOT_ROLE = PermissionInfo{
  CompleteKey: "root/workspaces/role/*",
  Name: "Entire role actions (*)",
}
var ALL_ROLE_PERMISSIONS = []PermissionInfo{
	PERM_ROOT_ROLE_DELETE,
	PERM_ROOT_ROLE_CREATE,
	PERM_ROOT_ROLE_UPDATE,
	PERM_ROOT_ROLE_QUERY,
	PERM_ROOT_ROLE,
}
var RoleEntityBundle = EntityBundle{
	Permissions: ALL_ROLE_PERMISSIONS,
	CliCommands: []cli.Command{
		RoleCliFn(),
	},
	Actions: GetRoleModule2Actions(),
	AutoMigrationEntities: []interface{}{
		&RoleEntity{},
  	},
}