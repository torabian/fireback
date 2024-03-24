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
type BackupTableMetaEntity struct {
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
    TableNameInDb   *string `json:"tableNameInDb" yaml:"tableNameInDb"       `
    // Datenano also has a text representation
    Children []*BackupTableMetaEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *BackupTableMetaEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var BackupTableMetaPreloadRelations []string = []string{}
var BACKUP_TABLE_META_EVENT_CREATED = "backupTableMeta.created"
var BACKUP_TABLE_META_EVENT_UPDATED = "backupTableMeta.updated"
var BACKUP_TABLE_META_EVENT_DELETED = "backupTableMeta.deleted"
var BACKUP_TABLE_META_EVENTS = []string{
	BACKUP_TABLE_META_EVENT_CREATED,
	BACKUP_TABLE_META_EVENT_UPDATED,
	BACKUP_TABLE_META_EVENT_DELETED,
}
type BackupTableMetaFieldMap struct {
		TableNameInDb TranslatedString `yaml:"tableNameInDb"`
}
var BackupTableMetaEntityMetaConfig map[string]int64 = map[string]int64{
}
var BackupTableMetaEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&BackupTableMetaEntity{}))
func entityBackupTableMetaFormatter(dto *BackupTableMetaEntity, query QueryDSL) {
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
func BackupTableMetaMockEntity() *BackupTableMetaEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &BackupTableMetaEntity{
      TableNameInDb : &stringHolder,
	}
	return entity
}
func BackupTableMetaActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := BackupTableMetaMockEntity()
		_, err := BackupTableMetaActionCreate(entity, query)
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
  func BackupTableMetaActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*BackupTableMetaEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &BackupTableMetaEntity{
          TableNameInDb: &tildaRef,
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
  func BackupTableMetaAssociationCreate(dto *BackupTableMetaEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func BackupTableMetaRelationContentCreate(dto *BackupTableMetaEntity, query QueryDSL) error {
return nil
}
func BackupTableMetaRelationContentUpdate(dto *BackupTableMetaEntity, query QueryDSL) error {
	return nil
}
func BackupTableMetaPolyglotCreateHandler(dto *BackupTableMetaEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func BackupTableMetaValidator(dto *BackupTableMetaEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func BackupTableMetaEntityPreSanitize(dto *BackupTableMetaEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func BackupTableMetaEntityBeforeCreateAppend(dto *BackupTableMetaEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    BackupTableMetaRecursiveAddUniqueId(dto, query)
  }
  func BackupTableMetaRecursiveAddUniqueId(dto *BackupTableMetaEntity, query QueryDSL) {
  }
func BackupTableMetaActionBatchCreateFn(dtos []*BackupTableMetaEntity, query QueryDSL) ([]*BackupTableMetaEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*BackupTableMetaEntity{}
		for _, item := range dtos {
			s, err := BackupTableMetaActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func BackupTableMetaDeleteEntireChildren(query QueryDSL, dto *BackupTableMetaEntity) (*IError) {
  return nil
}
func BackupTableMetaActionCreateFn(dto *BackupTableMetaEntity, query QueryDSL) (*BackupTableMetaEntity, *IError) {
	// 1. Validate always
	if iError := BackupTableMetaValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	BackupTableMetaEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	BackupTableMetaEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	BackupTableMetaPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	BackupTableMetaRelationContentCreate(dto, query)
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
	BackupTableMetaAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(BACKUP_TABLE_META_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&BackupTableMetaEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func BackupTableMetaActionGetOne(query QueryDSL) (*BackupTableMetaEntity, *IError) {
    refl := reflect.ValueOf(&BackupTableMetaEntity{})
    item, err := GetOneEntity[BackupTableMetaEntity](query, refl)
    entityBackupTableMetaFormatter(item, query)
    return item, err
  }
  func BackupTableMetaActionQuery(query QueryDSL) ([]*BackupTableMetaEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&BackupTableMetaEntity{})
    items, meta, err := QueryEntitiesPointer[BackupTableMetaEntity](query, refl)
    for _, item := range items {
      entityBackupTableMetaFormatter(item, query)
    }
    return items, meta, err
  }
  func BackupTableMetaUpdateExec(dbref *gorm.DB, query QueryDSL, fields *BackupTableMetaEntity) (*BackupTableMetaEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = BACKUP_TABLE_META_EVENT_UPDATED
    BackupTableMetaEntityPreSanitize(fields, query)
    var item BackupTableMetaEntity
    q := dbref.
      Where(&BackupTableMetaEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    BackupTableMetaRelationContentUpdate(fields, query)
    BackupTableMetaPolyglotCreateHandler(fields, query)
    if ero := BackupTableMetaDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&BackupTableMetaEntity{UniqueId: uniqueId}).
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
  func BackupTableMetaActionUpdateFn(query QueryDSL, fields *BackupTableMetaEntity) (*BackupTableMetaEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := BackupTableMetaValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // BackupTableMetaRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      var item *BackupTableMetaEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *IError
        item, err = BackupTableMetaUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, CastToIError(vf)
    } else {
      dbref = query.Tx
      return BackupTableMetaUpdateExec(dbref, query, fields)
    }
  }
var BackupTableMetaWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire backuptablemetas ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
      ActionRequires: []PermissionInfo{PERM_ROOT_BACKUP_TABLE_META_DELETE},
    })
		count, _ := BackupTableMetaActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func BackupTableMetaActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&BackupTableMetaEntity{})
	query.ActionRequires = []PermissionInfo{PERM_ROOT_BACKUP_TABLE_META_DELETE}
	return RemoveEntity[BackupTableMetaEntity](query, refl)
}
func BackupTableMetaActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[BackupTableMetaEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'BackupTableMetaEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func BackupTableMetaActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[BackupTableMetaEntity]) (
    *BulkRecordRequest[BackupTableMetaEntity], *IError,
  ) {
    result := []*BackupTableMetaEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := BackupTableMetaActionUpdate(query, record)
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
func (x *BackupTableMetaEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var BackupTableMetaEntityMeta = TableMetaData{
	EntityName:    "BackupTableMeta",
	ExportKey:    "backup-table-metas",
	TableNameInDb: "fb_backup-table-meta_entities",
	EntityObject:  &BackupTableMetaEntity{},
	ExportStream: BackupTableMetaActionExportT,
	ImportQuery: BackupTableMetaActionImport,
}
func BackupTableMetaActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[BackupTableMetaEntity](query, BackupTableMetaActionQuery, BackupTableMetaPreloadRelations)
}
func BackupTableMetaActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[BackupTableMetaEntity](query, BackupTableMetaActionQuery, BackupTableMetaPreloadRelations)
}
func BackupTableMetaActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content BackupTableMetaEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := BackupTableMetaActionCreate(&content, query)
	return err
}
var BackupTableMetaCommonCliFlags = []cli.Flag{
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
      Name:     "table-name-in-db",
      Required: false,
      Usage:    "tableNameInDb",
    },
}
var BackupTableMetaCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:     "tableNameInDb",
		StructField:     "TableNameInDb",
		Required: false,
		Usage:    "tableNameInDb",
		Type: "string",
	},
}
var BackupTableMetaCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "table-name-in-db",
      Required: false,
      Usage:    "tableNameInDb",
    },
}
  var BackupTableMetaCreateCmd cli.Command = BACKUP_TABLE_META_ACTION_POST_ONE.ToCli()
  var BackupTableMetaCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_BACKUP_TABLE_META_CREATE},
      })
      entity := &BackupTableMetaEntity{}
      for _, item := range BackupTableMetaCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := BackupTableMetaActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var BackupTableMetaUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: BackupTableMetaCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_BACKUP_TABLE_META_UPDATE},
      })
      entity := CastBackupTableMetaFromCli(c)
      if entity, err := BackupTableMetaActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* BackupTableMetaEntity) FromCli(c *cli.Context) *BackupTableMetaEntity {
	return CastBackupTableMetaFromCli(c)
}
func CastBackupTableMetaFromCli (c *cli.Context) *BackupTableMetaEntity {
	template := &BackupTableMetaEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("table-name-in-db") {
        value := c.String("table-name-in-db")
        template.TableNameInDb = &value
      }
	return template
}
  func BackupTableMetaSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      BackupTableMetaActionCreate,
      reflect.ValueOf(&BackupTableMetaEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func BackupTableMetaWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := BackupTableMetaActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "BackupTableMeta", result)
    }
  }
var BackupTableMetaImportExportCommands = []cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_BACKUP_TABLE_META_CREATE},
      })
			BackupTableMetaActionSeeder(query, c.Int("count"))
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
				Value: "backup-table-meta-seeder.yml",
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
        ActionRequires: []PermissionInfo{PERM_ROOT_BACKUP_TABLE_META_CREATE},
      })
			BackupTableMetaActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "backup-table-meta-seeder-backup-table-meta.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of backup-table-metas, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]BackupTableMetaEntity{}
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
			BackupTableMetaCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				BackupTableMetaActionCreate,
				reflect.ValueOf(&BackupTableMetaEntity{}).Elem(),
				c.String("file"),
        &SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_BACKUP_TABLE_META_CREATE},
				},
        func() BackupTableMetaEntity {
					v := CastBackupTableMetaFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var BackupTableMetaCliCommands []cli.Command = []cli.Command{
      GetCommonQuery2(BackupTableMetaActionQuery, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_BACKUP_TABLE_META_CREATE},
      }),
      GetCommonTableQuery(reflect.ValueOf(&BackupTableMetaEntity{}).Elem(), BackupTableMetaActionQuery),
          BackupTableMetaCreateCmd,
          BackupTableMetaUpdateCmd,
          BackupTableMetaCreateInteractiveCmd,
          BackupTableMetaWipeCmd,
          GetCommonRemoveQuery(reflect.ValueOf(&BackupTableMetaEntity{}).Elem(), BackupTableMetaActionRemove),
  }
  func BackupTableMetaCliFn() cli.Command {
    BackupTableMetaCliCommands = append(BackupTableMetaCliCommands, BackupTableMetaImportExportCommands...)
    return cli.Command{
      Name:        "backup",
      Description: "BackupTableMetas module actions (sample module to handle complex entities)",
      Usage:       "Keeps information about which tables to be used during backup (mostly internal)",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: BackupTableMetaCliCommands,
    }
  }
var BACKUP_TABLE_META_ACTION_POST_ONE = Module2Action{
    ActionName:    "create",
    ActionAliases: []string{"c"},
    Description: "Create new backupTableMeta",
    Flags: BackupTableMetaCommonCliFlags,
    Method: "POST",
    Url:    "/backup-table-meta",
    SecurityModel: &SecurityModel{
      ActionRequires: []PermissionInfo{PERM_ROOT_BACKUP_TABLE_META_CREATE},
    },
    Handlers: []gin.HandlerFunc{
      func (c *gin.Context) {
        HttpPostEntity(c, BackupTableMetaActionCreate)
      },
    },
    CliAction: func(c *cli.Context, security *SecurityModel) error {
      result, err := CliPostEntity(c, BackupTableMetaActionCreate, security)
      HandleActionInCli(c, result, err, map[string]map[string]string{})
      return err
    },
    Action: BackupTableMetaActionCreate,
    Format: "POST_ONE",
    RequestEntity: &BackupTableMetaEntity{},
    ResponseEntity: &BackupTableMetaEntity{},
  }
  /**
  *	Override this function on BackupTableMetaEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendBackupTableMetaRouter = func(r *[]Module2Action) {}
  func GetBackupTableMetaModule2Actions() []Module2Action {
    routes := []Module2Action{
       {
        Method: "GET",
        Url:    "/backup-table-metas",
        SecurityModel: &SecurityModel{
          ActionRequires: []PermissionInfo{PERM_ROOT_BACKUP_TABLE_META_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpQueryEntity(c, BackupTableMetaActionQuery)
          },
        },
        Format: "QUERY",
        Action: BackupTableMetaActionQuery,
        ResponseEntity: &[]BackupTableMetaEntity{},
      },
      {
        Method: "GET",
        Url:    "/backup-table-metas/export",
        SecurityModel: &SecurityModel{
          ActionRequires: []PermissionInfo{PERM_ROOT_BACKUP_TABLE_META_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpStreamFileChannel(c, BackupTableMetaActionExport)
          },
        },
        Format: "QUERY",
        Action: BackupTableMetaActionExport,
        ResponseEntity: &[]BackupTableMetaEntity{},
      },
      {
        Method: "GET",
        Url:    "/backup-table-meta/:uniqueId",
        SecurityModel: &SecurityModel{
          ActionRequires: []PermissionInfo{PERM_ROOT_BACKUP_TABLE_META_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpGetEntity(c, BackupTableMetaActionGetOne)
          },
        },
        Format: "GET_ONE",
        Action: BackupTableMetaActionGetOne,
        ResponseEntity: &BackupTableMetaEntity{},
      },
      BACKUP_TABLE_META_ACTION_POST_ONE,
      {
        ActionName:    "update",
        ActionAliases: []string{"u"},
        Flags: BackupTableMetaCommonCliFlagsOptional,
        Method: "PATCH",
        Url:    "/backup-table-meta",
        SecurityModel: &SecurityModel{
          ActionRequires: []PermissionInfo{PERM_ROOT_BACKUP_TABLE_META_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpUpdateEntity(c, BackupTableMetaActionUpdate)
          },
        },
        Action: BackupTableMetaActionUpdate,
        RequestEntity: &BackupTableMetaEntity{},
        Format: "PATCH_ONE",
        ResponseEntity: &BackupTableMetaEntity{},
      },
      {
        Method: "PATCH",
        Url:    "/backup-table-metas",
        SecurityModel: &SecurityModel{
          ActionRequires: []PermissionInfo{PERM_ROOT_BACKUP_TABLE_META_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpUpdateEntities(c, BackupTableMetaActionBulkUpdate)
          },
        },
        Action: BackupTableMetaActionBulkUpdate,
        Format: "PATCH_BULK",
        RequestEntity:  &BulkRecordRequest[BackupTableMetaEntity]{},
        ResponseEntity: &BulkRecordRequest[BackupTableMetaEntity]{},
      },
      {
        Method: "DELETE",
        Url:    "/backup-table-meta",
        Format: "DELETE_DSL",
        SecurityModel: &SecurityModel{
          ActionRequires: []PermissionInfo{PERM_ROOT_BACKUP_TABLE_META_DELETE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpRemoveEntity(c, BackupTableMetaActionRemove)
          },
        },
        Action: BackupTableMetaActionRemove,
        RequestEntity: &DeleteRequest{},
        ResponseEntity: &DeleteResponse{},
        TargetEntity: &BackupTableMetaEntity{},
      },
    }
    // Append user defined functions
    AppendBackupTableMetaRouter(&routes)
    return routes
  }
  func CreateBackupTableMetaRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetBackupTableMetaModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, BackupTableMetaEntityJsonSchema, "backup-table-meta-http", "workspaces")
    WriteEntitySchema("BackupTableMetaEntity", BackupTableMetaEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_BACKUP_TABLE_META_DELETE = PermissionInfo{
  CompleteKey: "root/workspaces/backup-table-meta/delete",
}
var PERM_ROOT_BACKUP_TABLE_META_CREATE = PermissionInfo{
  CompleteKey: "root/workspaces/backup-table-meta/create",
}
var PERM_ROOT_BACKUP_TABLE_META_UPDATE = PermissionInfo{
  CompleteKey: "root/workspaces/backup-table-meta/update",
}
var PERM_ROOT_BACKUP_TABLE_META_QUERY = PermissionInfo{
  CompleteKey: "root/workspaces/backup-table-meta/query",
}
var PERM_ROOT_BACKUP_TABLE_META = PermissionInfo{
  CompleteKey: "root/workspaces/backup-table-meta/*",
}
var ALL_BACKUP_TABLE_META_PERMISSIONS = []PermissionInfo{
	PERM_ROOT_BACKUP_TABLE_META_DELETE,
	PERM_ROOT_BACKUP_TABLE_META_CREATE,
	PERM_ROOT_BACKUP_TABLE_META_UPDATE,
	PERM_ROOT_BACKUP_TABLE_META_QUERY,
	PERM_ROOT_BACKUP_TABLE_META,
}