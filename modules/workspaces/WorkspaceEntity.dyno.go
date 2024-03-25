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
    queries "github.com/torabian/fireback/modules/workspaces/queries"
	"embed"
	reflect "reflect"
	"github.com/urfave/cli"
)
type WorkspaceEntity struct {
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
    Description   *string `json:"description" yaml:"description"       `
    // Datenano also has a text representation
    Name   *string `json:"name" yaml:"name"  validate:"required"       `
    // Datenano also has a text representation
    Children []*WorkspaceEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *WorkspaceEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var WorkspacePreloadRelations []string = []string{}
var WORKSPACE_EVENT_CREATED = "workspace.created"
var WORKSPACE_EVENT_UPDATED = "workspace.updated"
var WORKSPACE_EVENT_DELETED = "workspace.deleted"
var WORKSPACE_EVENTS = []string{
	WORKSPACE_EVENT_CREATED,
	WORKSPACE_EVENT_UPDATED,
	WORKSPACE_EVENT_DELETED,
}
type WorkspaceFieldMap struct {
		Description TranslatedString `yaml:"description"`
		Name TranslatedString `yaml:"name"`
}
var WorkspaceEntityMetaConfig map[string]int64 = map[string]int64{
}
var WorkspaceEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&WorkspaceEntity{}))
func entityWorkspaceFormatter(dto *WorkspaceEntity, query QueryDSL) {
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
func WorkspaceMockEntity() *WorkspaceEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &WorkspaceEntity{
      Description : &stringHolder,
      Name : &stringHolder,
	}
	return entity
}
func WorkspaceActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := WorkspaceMockEntity()
		_, err := WorkspaceActionCreate(entity, query)
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
  func WorkspaceActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*WorkspaceEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &WorkspaceEntity{
          Description: &tildaRef,
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
  func WorkspaceAssociationCreate(dto *WorkspaceEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func WorkspaceRelationContentCreate(dto *WorkspaceEntity, query QueryDSL) error {
return nil
}
func WorkspaceRelationContentUpdate(dto *WorkspaceEntity, query QueryDSL) error {
	return nil
}
func WorkspacePolyglotCreateHandler(dto *WorkspaceEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func WorkspaceValidator(dto *WorkspaceEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func WorkspaceEntityPreSanitize(dto *WorkspaceEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func WorkspaceEntityBeforeCreateAppend(dto *WorkspaceEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    WorkspaceRecursiveAddUniqueId(dto, query)
  }
  func WorkspaceRecursiveAddUniqueId(dto *WorkspaceEntity, query QueryDSL) {
  }
func WorkspaceActionBatchCreateFn(dtos []*WorkspaceEntity, query QueryDSL) ([]*WorkspaceEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*WorkspaceEntity{}
		for _, item := range dtos {
			s, err := WorkspaceActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func WorkspaceDeleteEntireChildren(query QueryDSL, dto *WorkspaceEntity) (*IError) {
  return nil
}
func WorkspaceActionCreateFn(dto *WorkspaceEntity, query QueryDSL) (*WorkspaceEntity, *IError) {
	// 1. Validate always
	if iError := WorkspaceValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	WorkspaceEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	WorkspaceEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	WorkspacePolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	WorkspaceRelationContentCreate(dto, query)
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
	WorkspaceAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(WORKSPACE_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&WorkspaceEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func WorkspaceActionGetOne(query QueryDSL) (*WorkspaceEntity, *IError) {
    refl := reflect.ValueOf(&WorkspaceEntity{})
    item, err := GetOneEntity[WorkspaceEntity](query, refl)
    entityWorkspaceFormatter(item, query)
    return item, err
  }
  func WorkspaceActionQuery(query QueryDSL) ([]*WorkspaceEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&WorkspaceEntity{})
    items, meta, err := QueryEntitiesPointer[WorkspaceEntity](query, refl)
    for _, item := range items {
      entityWorkspaceFormatter(item, query)
    }
    return items, meta, err
  }
  func (dto *WorkspaceEntity) Size() int {
    var size int = len(dto.Children)
    for _, c := range dto.Children {
      size += c.Size()
    }
    return size
  }
  func (dto *WorkspaceEntity) Add(nodes ...*WorkspaceEntity) bool {
    var size = dto.Size()
    for _, n := range nodes {
      if n.ParentId != nil && *n.ParentId == dto.UniqueId {
        dto.Children = append(dto.Children, n)
      } else {
        for _, c := range dto.Children {
          if c.Add(n) {
            break
          }
        }
      }
    }
    return dto.Size() == size+len(nodes)
  }
  func WorkspaceActionCommonPivotQuery(query QueryDSL) ([]*PivotResult, *QueryResultMeta, error) {
    items, meta, err := UnsafeQuerySqlFromFs[PivotResult](
      &queries.QueriesFs, "WorkspaceCommonPivot.sqlite.dyno", query,
    )
    return items, meta, err
  }
  func WorkspaceActionCteQuery(query QueryDSL) ([]*WorkspaceEntity, *QueryResultMeta, error) {
    items, meta, err := UnsafeQuerySqlFromFs[WorkspaceEntity](
      &queries.QueriesFs, "WorkspaceCTE.sqlite.dyno", query,
    )
    for _, item := range items {
      entityWorkspaceFormatter(item, query)
    }
    var tree []*WorkspaceEntity
    for _, item := range items {
      if item.ParentId == nil {
        root := item
        root.Add(items...)
        tree = append(tree, root)
      }
    }
    return tree, meta, err
  }
  func WorkspaceUpdateExec(dbref *gorm.DB, query QueryDSL, fields *WorkspaceEntity) (*WorkspaceEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = WORKSPACE_EVENT_UPDATED
    WorkspaceEntityPreSanitize(fields, query)
    var item WorkspaceEntity
    q := dbref.
      Where(&WorkspaceEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    WorkspaceRelationContentUpdate(fields, query)
    WorkspacePolyglotCreateHandler(fields, query)
    if ero := WorkspaceDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&WorkspaceEntity{UniqueId: uniqueId}).
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
  func WorkspaceActionUpdateFn(query QueryDSL, fields *WorkspaceEntity) (*WorkspaceEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := WorkspaceValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // WorkspaceRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      var item *WorkspaceEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *IError
        item, err = WorkspaceUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, CastToIError(vf)
    } else {
      dbref = query.Tx
      return WorkspaceUpdateExec(dbref, query, fields)
    }
  }
var WorkspaceWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire workspaces ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
      ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_DELETE},
    })
		count, _ := WorkspaceActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func WorkspaceActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&WorkspaceEntity{})
	query.ActionRequires = []PermissionInfo{PERM_ROOT_WORKSPACE_DELETE}
	return RemoveEntity[WorkspaceEntity](query, refl)
}
func WorkspaceActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[WorkspaceEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'WorkspaceEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func WorkspaceActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[WorkspaceEntity]) (
    *BulkRecordRequest[WorkspaceEntity], *IError,
  ) {
    result := []*WorkspaceEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := WorkspaceActionUpdate(query, record)
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
func (x *WorkspaceEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var WorkspaceEntityMeta = TableMetaData{
	EntityName:    "Workspace",
	ExportKey:    "workspaces",
	TableNameInDb: "fb_workspace_entities",
	EntityObject:  &WorkspaceEntity{},
	ExportStream: WorkspaceActionExportT,
	ImportQuery: WorkspaceActionImport,
}
func WorkspaceActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[WorkspaceEntity](query, WorkspaceActionQuery, WorkspacePreloadRelations)
}
func WorkspaceActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[WorkspaceEntity](query, WorkspaceActionQuery, WorkspacePreloadRelations)
}
func WorkspaceActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content WorkspaceEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := WorkspaceActionCreate(&content, query)
	return err
}
var WorkspaceCommonCliFlags = []cli.Flag{
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
      Name:     "description",
      Required: false,
      Usage:    "description",
    },
    &cli.StringFlag{
      Name:     "name",
      Required: true,
      Usage:    "name",
    },
}
var WorkspaceCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:     "description",
		StructField:     "Description",
		Required: false,
		Usage:    "description",
		Type: "string",
	},
	{
		Name:     "name",
		StructField:     "Name",
		Required: true,
		Usage:    "name",
		Type: "string",
	},
}
var WorkspaceCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "description",
      Required: false,
      Usage:    "description",
    },
    &cli.StringFlag{
      Name:     "name",
      Required: true,
      Usage:    "name",
    },
}
  var WorkspaceCreateCmd cli.Command = WORKSPACE_ACTION_POST_ONE.ToCli()
  var WorkspaceCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_CREATE},
      })
      entity := &WorkspaceEntity{}
      for _, item := range WorkspaceCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := WorkspaceActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var WorkspaceUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: WorkspaceCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_UPDATE},
      })
      entity := CastWorkspaceFromCli(c)
      if entity, err := WorkspaceActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* WorkspaceEntity) FromCli(c *cli.Context) *WorkspaceEntity {
	return CastWorkspaceFromCli(c)
}
func CastWorkspaceFromCli (c *cli.Context) *WorkspaceEntity {
	template := &WorkspaceEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("description") {
        value := c.String("description")
        template.Description = &value
      }
      if c.IsSet("name") {
        value := c.String("name")
        template.Name = &value
      }
	return template
}
  func WorkspaceSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      WorkspaceActionCreate,
      reflect.ValueOf(&WorkspaceEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func WorkspaceWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := WorkspaceActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "Workspace", result)
    }
  }
var WorkspaceImportExportCommands = []cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_CREATE},
      })
			WorkspaceActionSeeder(query, c.Int("count"))
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
				Value: "workspace-seeder.yml",
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
        ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_CREATE},
      })
			WorkspaceActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "workspace-seeder-workspace.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of workspaces, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]WorkspaceEntity{}
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
			WorkspaceCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				WorkspaceActionCreate,
				reflect.ValueOf(&WorkspaceEntity{}).Elem(),
				c.String("file"),
        &SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_CREATE},
				},
        func() WorkspaceEntity {
					v := CastWorkspaceFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var WorkspaceCliCommands []cli.Command = []cli.Command{
      WORKSPACE_ACTION_QUERY.ToCli(),
      WORKSPACE_ACTION_TABLE.ToCli(),
      GetCommonTableQuery(reflect.ValueOf(&WorkspaceEntity{}).Elem(), WorkspaceActionQuery),
      WorkspaceCreateCmd,
      WorkspaceUpdateCmd,
      WorkspaceCreateInteractiveCmd,
      WorkspaceWipeCmd,
      GetCommonRemoveQuery(reflect.ValueOf(&WorkspaceEntity{}).Elem(), WorkspaceActionRemove),
          GetCommonCteQuery(WorkspaceActionCteQuery),
          GetCommonPivotQuery(WorkspaceActionCommonPivotQuery),
  }
  func WorkspaceCliFn() cli.Command {
    WorkspaceCliCommands = append(WorkspaceCliCommands, WorkspaceImportExportCommands...)
    return cli.Command{
      Name:        "ws",
      Description: "Workspaces module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: WorkspaceCliCommands,
    }
  }
var WORKSPACE_ACTION_TABLE = Module2Action{
  Name:    "table",
  ActionAliases: []string{"t"},
  Flags:  CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: WorkspaceActionQuery,
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    CommonCliTableCmd2(c,
      WorkspaceActionQuery,
      security,
      reflect.ValueOf(&WorkspaceEntity{}).Elem(),
    )
    return nil
  },
}
var WORKSPACE_ACTION_QUERY = Module2Action{
  Method: "GET",
  Url:    "/workspaces",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpQueryEntity(c, WorkspaceActionQuery)
    },
  },
  Format: "QUERY",
  Action: WorkspaceActionQuery,
  ResponseEntity: &[]WorkspaceEntity{},
  CliAction: func(c *cli.Context, security *SecurityModel) error {
		CommonCliQueryCmd2(
			c,
			WorkspaceActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionAliases: []string{"q"},
	Flags:         CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var WORKSPACE_ACTION_QUERY_CTE = Module2Action{
  Method: "GET",
  Url:    "/cte-workspaces",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpQueryEntity(c, WorkspaceActionCteQuery)
    },
  },
  Format: "QUERY",
  Action: WorkspaceActionCteQuery,
  ResponseEntity: &[]WorkspaceEntity{},
}
var WORKSPACE_ACTION_EXPORT = Module2Action{
  Method: "GET",
  Url:    "/workspaces/export",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpStreamFileChannel(c, WorkspaceActionExport)
    },
  },
  Format: "QUERY",
  Action: WorkspaceActionExport,
  ResponseEntity: &[]WorkspaceEntity{},
}
var WORKSPACE_ACTION_GET_ONE = Module2Action{
  Method: "GET",
  Url:    "/workspace/:uniqueId",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpGetEntity(c, WorkspaceActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: WorkspaceActionGetOne,
  ResponseEntity: &WorkspaceEntity{},
}
var WORKSPACE_ACTION_POST_ONE = Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new workspace",
  Flags: WorkspaceCommonCliFlags,
  Method: "POST",
  Url:    "/workspace",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpPostEntity(c, WorkspaceActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    result, err := CliPostEntity(c, WorkspaceActionCreate, security)
    HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: WorkspaceActionCreate,
  Format: "POST_ONE",
  RequestEntity: &WorkspaceEntity{},
  ResponseEntity: &WorkspaceEntity{},
}
var WORKSPACE_ACTION_PATCH = Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: WorkspaceCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/workspace",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntity(c, WorkspaceActionUpdate)
    },
  },
  Action: WorkspaceActionUpdate,
  RequestEntity: &WorkspaceEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &WorkspaceEntity{},
}
var WORKSPACE_ACTION_PATCH_BULK = Module2Action{
  Method: "PATCH",
  Url:    "/workspaces",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntities(c, WorkspaceActionBulkUpdate)
    },
  },
  Action: WorkspaceActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &BulkRecordRequest[WorkspaceEntity]{},
  ResponseEntity: &BulkRecordRequest[WorkspaceEntity]{},
}
var WORKSPACE_ACTION_DELETE = Module2Action{
  Method: "DELETE",
  Url:    "/workspace",
  Format: "DELETE_DSL",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpRemoveEntity(c, WorkspaceActionRemove)
    },
  },
  Action: WorkspaceActionRemove,
  RequestEntity: &DeleteRequest{},
  ResponseEntity: &DeleteResponse{},
  TargetEntity: &WorkspaceEntity{},
}
  /**
  *	Override this function on WorkspaceEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendWorkspaceRouter = func(r *[]Module2Action) {}
  func GetWorkspaceModule2Actions() []Module2Action {
    routes := []Module2Action{
        WORKSPACE_ACTION_QUERY_CTE,
      WORKSPACE_ACTION_QUERY,
      WORKSPACE_ACTION_EXPORT,
      WORKSPACE_ACTION_GET_ONE,
      WORKSPACE_ACTION_POST_ONE,
      WORKSPACE_ACTION_PATCH,
      WORKSPACE_ACTION_PATCH_BULK,
      WORKSPACE_ACTION_DELETE,
    }
    // Append user defined functions
    AppendWorkspaceRouter(&routes)
    return routes
  }
  func CreateWorkspaceRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetWorkspaceModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, WorkspaceEntityJsonSchema, "workspace-http", "workspaces")
    WriteEntitySchema("WorkspaceEntity", WorkspaceEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_WORKSPACE_DELETE = PermissionInfo{
  CompleteKey: "root/workspaces/workspace/delete",
}
var PERM_ROOT_WORKSPACE_CREATE = PermissionInfo{
  CompleteKey: "root/workspaces/workspace/create",
}
var PERM_ROOT_WORKSPACE_UPDATE = PermissionInfo{
  CompleteKey: "root/workspaces/workspace/update",
}
var PERM_ROOT_WORKSPACE_QUERY = PermissionInfo{
  CompleteKey: "root/workspaces/workspace/query",
}
var PERM_ROOT_WORKSPACE = PermissionInfo{
  CompleteKey: "root/workspaces/workspace/*",
}
var ALL_WORKSPACE_PERMISSIONS = []PermissionInfo{
	PERM_ROOT_WORKSPACE_DELETE,
	PERM_ROOT_WORKSPACE_CREATE,
	PERM_ROOT_WORKSPACE_UPDATE,
	PERM_ROOT_WORKSPACE_QUERY,
	PERM_ROOT_WORKSPACE,
}