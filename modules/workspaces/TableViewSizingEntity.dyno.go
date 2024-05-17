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
var tableViewSizingSeedersFs *embed.FS = nil
func ResetTableViewSizingSeeders(fs *embed.FS) {
	tableViewSizingSeedersFs = fs
}
type TableViewSizingEntity struct {
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
    TableName   *string `json:"tableName" yaml:"tableName"  validate:"required"       `
    // Datenano also has a text representation
    Sizes   *string `json:"sizes" yaml:"sizes"       `
    // Datenano also has a text representation
    Children []*TableViewSizingEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *TableViewSizingEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var TableViewSizingPreloadRelations []string = []string{}
var TABLE_VIEW_SIZING_EVENT_CREATED = "tableViewSizing.created"
var TABLE_VIEW_SIZING_EVENT_UPDATED = "tableViewSizing.updated"
var TABLE_VIEW_SIZING_EVENT_DELETED = "tableViewSizing.deleted"
var TABLE_VIEW_SIZING_EVENTS = []string{
	TABLE_VIEW_SIZING_EVENT_CREATED,
	TABLE_VIEW_SIZING_EVENT_UPDATED,
	TABLE_VIEW_SIZING_EVENT_DELETED,
}
type TableViewSizingFieldMap struct {
		TableName TranslatedString `yaml:"tableName"`
		Sizes TranslatedString `yaml:"sizes"`
}
var TableViewSizingEntityMetaConfig map[string]int64 = map[string]int64{
}
var TableViewSizingEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&TableViewSizingEntity{}))
func entityTableViewSizingFormatter(dto *TableViewSizingEntity, query QueryDSL) {
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
func TableViewSizingMockEntity() *TableViewSizingEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &TableViewSizingEntity{
      TableName : &stringHolder,
      Sizes : &stringHolder,
	}
	return entity
}
func TableViewSizingActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := TableViewSizingMockEntity()
		_, err := TableViewSizingActionCreate(entity, query)
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
  func TableViewSizingActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*TableViewSizingEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &TableViewSizingEntity{
          TableName: &tildaRef,
          Sizes: &tildaRef,
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
  func TableViewSizingAssociationCreate(dto *TableViewSizingEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func TableViewSizingRelationContentCreate(dto *TableViewSizingEntity, query QueryDSL) error {
return nil
}
func TableViewSizingRelationContentUpdate(dto *TableViewSizingEntity, query QueryDSL) error {
	return nil
}
func TableViewSizingPolyglotCreateHandler(dto *TableViewSizingEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func TableViewSizingValidator(dto *TableViewSizingEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func TableViewSizingEntityPreSanitize(dto *TableViewSizingEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func TableViewSizingEntityBeforeCreateAppend(dto *TableViewSizingEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    TableViewSizingRecursiveAddUniqueId(dto, query)
  }
  func TableViewSizingRecursiveAddUniqueId(dto *TableViewSizingEntity, query QueryDSL) {
  }
func TableViewSizingActionBatchCreateFn(dtos []*TableViewSizingEntity, query QueryDSL) ([]*TableViewSizingEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*TableViewSizingEntity{}
		for _, item := range dtos {
			s, err := TableViewSizingActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func TableViewSizingDeleteEntireChildren(query QueryDSL, dto *TableViewSizingEntity) (*IError) {
  // intentionally removed this. It's hard to implement it, and probably wrong without
  // proper on delete cascade
  return nil
}
func TableViewSizingActionCreateFn(dto *TableViewSizingEntity, query QueryDSL) (*TableViewSizingEntity, *IError) {
	// 1. Validate always
	if iError := TableViewSizingValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	TableViewSizingEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	TableViewSizingEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	TableViewSizingPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	TableViewSizingRelationContentCreate(dto, query)
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
	TableViewSizingAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(TABLE_VIEW_SIZING_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&TableViewSizingEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func TableViewSizingActionGetOne(query QueryDSL) (*TableViewSizingEntity, *IError) {
    refl := reflect.ValueOf(&TableViewSizingEntity{})
    item, err := GetOneEntity[TableViewSizingEntity](query, refl)
    entityTableViewSizingFormatter(item, query)
    return item, err
  }
  func TableViewSizingActionQuery(query QueryDSL) ([]*TableViewSizingEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&TableViewSizingEntity{})
    items, meta, err := QueryEntitiesPointer[TableViewSizingEntity](query, refl)
    for _, item := range items {
      entityTableViewSizingFormatter(item, query)
    }
    return items, meta, err
  }
  func TableViewSizingUpdateExec(dbref *gorm.DB, query QueryDSL, fields *TableViewSizingEntity) (*TableViewSizingEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = TABLE_VIEW_SIZING_EVENT_UPDATED
    TableViewSizingEntityPreSanitize(fields, query)
    var item TableViewSizingEntity
    q := dbref.
      Where(&TableViewSizingEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    TableViewSizingRelationContentUpdate(fields, query)
    TableViewSizingPolyglotCreateHandler(fields, query)
    if ero := TableViewSizingDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&TableViewSizingEntity{UniqueId: uniqueId}).
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
  func TableViewSizingActionUpdateFn(query QueryDSL, fields *TableViewSizingEntity) (*TableViewSizingEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := TableViewSizingValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // TableViewSizingRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      var item *TableViewSizingEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *IError
        item, err = TableViewSizingUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, CastToIError(vf)
    } else {
      dbref = query.Tx
      return TableViewSizingUpdateExec(dbref, query, fields)
    }
  }
var TableViewSizingWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire tableviewsizings ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
      ActionRequires: []PermissionInfo{PERM_ROOT_TABLE_VIEW_SIZING_DELETE},
    })
		count, _ := TableViewSizingActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func TableViewSizingActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&TableViewSizingEntity{})
	query.ActionRequires = []PermissionInfo{PERM_ROOT_TABLE_VIEW_SIZING_DELETE}
	return RemoveEntity[TableViewSizingEntity](query, refl)
}
func TableViewSizingActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[TableViewSizingEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'TableViewSizingEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func TableViewSizingActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[TableViewSizingEntity]) (
    *BulkRecordRequest[TableViewSizingEntity], *IError,
  ) {
    result := []*TableViewSizingEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := TableViewSizingActionUpdate(query, record)
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
func (x *TableViewSizingEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var TableViewSizingEntityMeta = TableMetaData{
	EntityName:    "TableViewSizing",
	ExportKey:    "table-view-sizings",
	TableNameInDb: "fb_table-view-sizing_entities",
	EntityObject:  &TableViewSizingEntity{},
	ExportStream: TableViewSizingActionExportT,
	ImportQuery: TableViewSizingActionImport,
}
func TableViewSizingActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[TableViewSizingEntity](query, TableViewSizingActionQuery, TableViewSizingPreloadRelations)
}
func TableViewSizingActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[TableViewSizingEntity](query, TableViewSizingActionQuery, TableViewSizingPreloadRelations)
}
func TableViewSizingActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content TableViewSizingEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := TableViewSizingActionCreate(&content, query)
	return err
}
var TableViewSizingCommonCliFlags = []cli.Flag{
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
      Name:     "table-name",
      Required: true,
      Usage:    "tableName",
    },
    &cli.StringFlag{
      Name:     "sizes",
      Required: false,
      Usage:    "sizes",
    },
}
var TableViewSizingCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:     "tableName",
		StructField:     "TableName",
		Required: true,
		Usage:    "tableName",
		Type: "string",
	},
	{
		Name:     "sizes",
		StructField:     "Sizes",
		Required: false,
		Usage:    "sizes",
		Type: "string",
	},
}
var TableViewSizingCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "table-name",
      Required: true,
      Usage:    "tableName",
    },
    &cli.StringFlag{
      Name:     "sizes",
      Required: false,
      Usage:    "sizes",
    },
}
  var TableViewSizingCreateCmd cli.Command = TABLE_VIEW_SIZING_ACTION_POST_ONE.ToCli()
  var TableViewSizingCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_TABLE_VIEW_SIZING_CREATE},
      })
      entity := &TableViewSizingEntity{}
      for _, item := range TableViewSizingCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := TableViewSizingActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var TableViewSizingUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: TableViewSizingCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_TABLE_VIEW_SIZING_UPDATE},
      })
      entity := CastTableViewSizingFromCli(c)
      if entity, err := TableViewSizingActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* TableViewSizingEntity) FromCli(c *cli.Context) *TableViewSizingEntity {
	return CastTableViewSizingFromCli(c)
}
func CastTableViewSizingFromCli (c *cli.Context) *TableViewSizingEntity {
	template := &TableViewSizingEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("table-name") {
        value := c.String("table-name")
        template.TableName = &value
      }
      if c.IsSet("sizes") {
        value := c.String("sizes")
        template.Sizes = &value
      }
	return template
}
  func TableViewSizingSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      TableViewSizingActionCreate,
      reflect.ValueOf(&TableViewSizingEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func TableViewSizingWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := TableViewSizingActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "TableViewSizing", result)
    }
  }
var TableViewSizingImportExportCommands = []cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_TABLE_VIEW_SIZING_CREATE},
      })
			TableViewSizingActionSeeder(query, c.Int("count"))
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
				Value: "table-view-sizing-seeder.yml",
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
        ActionRequires: []PermissionInfo{PERM_ROOT_TABLE_VIEW_SIZING_CREATE},
      })
			TableViewSizingActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "table-view-sizing-seeder-table-view-sizing.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of table-view-sizings, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]TableViewSizingEntity{}
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
			TableViewSizingCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				TableViewSizingActionCreate,
				reflect.ValueOf(&TableViewSizingEntity{}).Elem(),
				c.String("file"),
        &SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_TABLE_VIEW_SIZING_CREATE},
				},
        func() TableViewSizingEntity {
					v := CastTableViewSizingFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var TableViewSizingCliCommands []cli.Command = []cli.Command{
      TABLE_VIEW_SIZING_ACTION_QUERY.ToCli(),
      TABLE_VIEW_SIZING_ACTION_TABLE.ToCli(),
      TableViewSizingCreateCmd,
      TableViewSizingUpdateCmd,
      TableViewSizingCreateInteractiveCmd,
      TableViewSizingWipeCmd,
      GetCommonRemoveQuery(reflect.ValueOf(&TableViewSizingEntity{}).Elem(), TableViewSizingActionRemove),
  }
  func TableViewSizingCliFn() cli.Command {
    TableViewSizingCliCommands = append(TableViewSizingCliCommands, TableViewSizingImportExportCommands...)
    return cli.Command{
      Name:        "tableViewSizing",
      ShortName:   "tvs",
      Description: "TableViewSizings module actions (sample module to handle complex entities)",
      Usage:       "Used to store meta data about user tables (in front-end, or apps for example) about the size of the columns",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: TableViewSizingCliCommands,
    }
  }
var TABLE_VIEW_SIZING_ACTION_TABLE = Module2Action{
  Name:    "table",
  ActionName: "table",
  ActionAliases: []string{"t"},
  Flags:  CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: TableViewSizingActionQuery,
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    CommonCliTableCmd2(c,
      TableViewSizingActionQuery,
      security,
      reflect.ValueOf(&TableViewSizingEntity{}).Elem(),
    )
    return nil
  },
}
var TABLE_VIEW_SIZING_ACTION_QUERY = Module2Action{
  Method: "GET",
  Url:    "/table-view-sizings",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_TABLE_VIEW_SIZING_QUERY},
  },
  Group: "tableViewSizing",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpQueryEntity(c, TableViewSizingActionQuery)
    },
  },
  Format: "QUERY",
  Action: TableViewSizingActionQuery,
  ResponseEntity: &[]TableViewSizingEntity{},
  Out: Module2ActionBody{
		Entity: "TableViewSizingEntity",
	},
  CliAction: func(c *cli.Context, security *SecurityModel) error {
		CommonCliQueryCmd2(
			c,
			TableViewSizingActionQuery,
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
var TABLE_VIEW_SIZING_ACTION_EXPORT = Module2Action{
  Method: "GET",
  Url:    "/table-view-sizings/export",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_TABLE_VIEW_SIZING_QUERY},
  },
  Group: "tableViewSizing",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpStreamFileChannel(c, TableViewSizingActionExport)
    },
  },
  Format: "QUERY",
  Action: TableViewSizingActionExport,
  ResponseEntity: &[]TableViewSizingEntity{},
  Out: Module2ActionBody{
		Entity: "TableViewSizingEntity",
	},
}
var TABLE_VIEW_SIZING_ACTION_GET_ONE = Module2Action{
  Method: "GET",
  Url:    "/table-view-sizing/:uniqueId",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_TABLE_VIEW_SIZING_QUERY},
  },
  Group: "tableViewSizing",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpGetEntity(c, TableViewSizingActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: TableViewSizingActionGetOne,
  ResponseEntity: &TableViewSizingEntity{},
  Out: Module2ActionBody{
		Entity: "TableViewSizingEntity",
	},
}
var TABLE_VIEW_SIZING_ACTION_POST_ONE = Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new tableViewSizing",
  Flags: TableViewSizingCommonCliFlags,
  Method: "POST",
  Url:    "/table-view-sizing",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_TABLE_VIEW_SIZING_CREATE},
  },
  Group: "tableViewSizing",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpPostEntity(c, TableViewSizingActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    result, err := CliPostEntity(c, TableViewSizingActionCreate, security)
    HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: TableViewSizingActionCreate,
  Format: "POST_ONE",
  RequestEntity: &TableViewSizingEntity{},
  ResponseEntity: &TableViewSizingEntity{},
  Out: Module2ActionBody{
		Entity: "TableViewSizingEntity",
	},
  In: Module2ActionBody{
		Entity: "TableViewSizingEntity",
	},
}
var TABLE_VIEW_SIZING_ACTION_PATCH = Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: TableViewSizingCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/table-view-sizing",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_TABLE_VIEW_SIZING_UPDATE},
  },
  Group: "tableViewSizing",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntity(c, TableViewSizingActionUpdate)
    },
  },
  Action: TableViewSizingActionUpdate,
  RequestEntity: &TableViewSizingEntity{},
  ResponseEntity: &TableViewSizingEntity{},
  Format: "PATCH_ONE",
  Out: Module2ActionBody{
		Entity: "TableViewSizingEntity",
	},
  In: Module2ActionBody{
		Entity: "TableViewSizingEntity",
	},
}
var TABLE_VIEW_SIZING_ACTION_PATCH_BULK = Module2Action{
  Method: "PATCH",
  Url:    "/table-view-sizings",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_TABLE_VIEW_SIZING_UPDATE},
  },
  Group: "tableViewSizing",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntities(c, TableViewSizingActionBulkUpdate)
    },
  },
  Action: TableViewSizingActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &BulkRecordRequest[TableViewSizingEntity]{},
  ResponseEntity: &BulkRecordRequest[TableViewSizingEntity]{},
  Out: Module2ActionBody{
		Entity: "TableViewSizingEntity",
	},
  In: Module2ActionBody{
		Entity: "TableViewSizingEntity",
	},
}
var TABLE_VIEW_SIZING_ACTION_DELETE = Module2Action{
  Method: "DELETE",
  Url:    "/table-view-sizing",
  Format: "DELETE_DSL",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_TABLE_VIEW_SIZING_DELETE},
  },
  Group: "tableViewSizing",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpRemoveEntity(c, TableViewSizingActionRemove)
    },
  },
  Action: TableViewSizingActionRemove,
  RequestEntity: &DeleteRequest{},
  ResponseEntity: &DeleteResponse{},
  TargetEntity: &TableViewSizingEntity{},
}
  /**
  *	Override this function on TableViewSizingEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendTableViewSizingRouter = func(r *[]Module2Action) {}
  func GetTableViewSizingModule2Actions() []Module2Action {
    routes := []Module2Action{
      TABLE_VIEW_SIZING_ACTION_QUERY,
      TABLE_VIEW_SIZING_ACTION_EXPORT,
      TABLE_VIEW_SIZING_ACTION_GET_ONE,
      TABLE_VIEW_SIZING_ACTION_POST_ONE,
      TABLE_VIEW_SIZING_ACTION_PATCH,
      TABLE_VIEW_SIZING_ACTION_PATCH_BULK,
      TABLE_VIEW_SIZING_ACTION_DELETE,
    }
    // Append user defined functions
    AppendTableViewSizingRouter(&routes)
    return routes
  }
  func CreateTableViewSizingRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetTableViewSizingModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, TableViewSizingEntityJsonSchema, "table-view-sizing-http", "workspaces")
    WriteEntitySchema("TableViewSizingEntity", TableViewSizingEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_TABLE_VIEW_SIZING_DELETE = PermissionInfo{
  CompleteKey: "root/workspaces/table-view-sizing/delete",
  Name: "Delete table view sizing",
}
var PERM_ROOT_TABLE_VIEW_SIZING_CREATE = PermissionInfo{
  CompleteKey: "root/workspaces/table-view-sizing/create",
  Name: "Create table view sizing",
}
var PERM_ROOT_TABLE_VIEW_SIZING_UPDATE = PermissionInfo{
  CompleteKey: "root/workspaces/table-view-sizing/update",
  Name: "Update table view sizing",
}
var PERM_ROOT_TABLE_VIEW_SIZING_QUERY = PermissionInfo{
  CompleteKey: "root/workspaces/table-view-sizing/query",
  Name: "Query table view sizing",
}
var PERM_ROOT_TABLE_VIEW_SIZING = PermissionInfo{
  CompleteKey: "root/workspaces/table-view-sizing/*",
  Name: "Entire table view sizing actions (*)",
}
var ALL_TABLE_VIEW_SIZING_PERMISSIONS = []PermissionInfo{
	PERM_ROOT_TABLE_VIEW_SIZING_DELETE,
	PERM_ROOT_TABLE_VIEW_SIZING_CREATE,
	PERM_ROOT_TABLE_VIEW_SIZING_UPDATE,
	PERM_ROOT_TABLE_VIEW_SIZING_QUERY,
	PERM_ROOT_TABLE_VIEW_SIZING,
}