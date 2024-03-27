package geo
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
type GeoStateEntity struct {
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
    Name   *string `json:"name" yaml:"name"        translate:"true" `
    // Datenano also has a text representation
    Translations     []*GeoStateEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    Children []*GeoStateEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *GeoStateEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var GeoStatePreloadRelations []string = []string{}
var GEO_STATE_EVENT_CREATED = "geoState.created"
var GEO_STATE_EVENT_UPDATED = "geoState.updated"
var GEO_STATE_EVENT_DELETED = "geoState.deleted"
var GEO_STATE_EVENTS = []string{
	GEO_STATE_EVENT_CREATED,
	GEO_STATE_EVENT_UPDATED,
	GEO_STATE_EVENT_DELETED,
}
type GeoStateFieldMap struct {
		Name workspaces.TranslatedString `yaml:"name"`
}
var GeoStateEntityMetaConfig map[string]int64 = map[string]int64{
}
var GeoStateEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&GeoStateEntity{}))
  type GeoStateEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        Name string `yaml:"name" json:"name"`
  }
func entityGeoStateFormatter(dto *GeoStateEntity, query workspaces.QueryDSL) {
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
func GeoStateMockEntity() *GeoStateEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &GeoStateEntity{
      Name : &stringHolder,
	}
	return entity
}
func GeoStateActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := GeoStateMockEntity()
		_, err := GeoStateActionCreate(entity, query)
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
    func (x*GeoStateEntity) GetNameTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Name
          }
        }
      }
      return ""
    }
  func GeoStateActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*GeoStateEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &GeoStateEntity{
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
  func GeoStateAssociationCreate(dto *GeoStateEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func GeoStateRelationContentCreate(dto *GeoStateEntity, query workspaces.QueryDSL) error {
return nil
}
func GeoStateRelationContentUpdate(dto *GeoStateEntity, query workspaces.QueryDSL) error {
	return nil
}
func GeoStatePolyglotCreateHandler(dto *GeoStateEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &GeoStateEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func GeoStateValidator(dto *GeoStateEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func GeoStateEntityPreSanitize(dto *GeoStateEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func GeoStateEntityBeforeCreateAppend(dto *GeoStateEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    GeoStateRecursiveAddUniqueId(dto, query)
  }
  func GeoStateRecursiveAddUniqueId(dto *GeoStateEntity, query workspaces.QueryDSL) {
  }
func GeoStateActionBatchCreateFn(dtos []*GeoStateEntity, query workspaces.QueryDSL) ([]*GeoStateEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*GeoStateEntity{}
		for _, item := range dtos {
			s, err := GeoStateActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func GeoStateDeleteEntireChildren(query workspaces.QueryDSL, dto *GeoStateEntity) (*workspaces.IError) {
  return nil
}
func GeoStateActionCreateFn(dto *GeoStateEntity, query workspaces.QueryDSL) (*GeoStateEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := GeoStateValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	GeoStateEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	GeoStateEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	GeoStatePolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	GeoStateRelationContentCreate(dto, query)
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
	GeoStateAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(GEO_STATE_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&GeoStateEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func GeoStateActionGetOne(query workspaces.QueryDSL) (*GeoStateEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&GeoStateEntity{})
    item, err := workspaces.GetOneEntity[GeoStateEntity](query, refl)
    entityGeoStateFormatter(item, query)
    return item, err
  }
  func GeoStateActionQuery(query workspaces.QueryDSL) ([]*GeoStateEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&GeoStateEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[GeoStateEntity](query, refl)
    for _, item := range items {
      entityGeoStateFormatter(item, query)
    }
    return items, meta, err
  }
  func GeoStateUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *GeoStateEntity) (*GeoStateEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = GEO_STATE_EVENT_UPDATED
    GeoStateEntityPreSanitize(fields, query)
    var item GeoStateEntity
    q := dbref.
      Where(&GeoStateEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    GeoStateRelationContentUpdate(fields, query)
    GeoStatePolyglotCreateHandler(fields, query)
    if ero := GeoStateDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&GeoStateEntity{UniqueId: uniqueId}).
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
  func GeoStateActionUpdateFn(query workspaces.QueryDSL, fields *GeoStateEntity) (*GeoStateEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := GeoStateValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // GeoStateRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *GeoStateEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = GeoStateUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return GeoStateUpdateExec(dbref, query, fields)
    }
  }
var GeoStateWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire geostates ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_STATE_DELETE},
    })
		count, _ := GeoStateActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func GeoStateActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&GeoStateEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_GEO_STATE_DELETE}
	return workspaces.RemoveEntity[GeoStateEntity](query, refl)
}
func GeoStateActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[GeoStateEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'GeoStateEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func GeoStateActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[GeoStateEntity]) (
    *workspaces.BulkRecordRequest[GeoStateEntity], *workspaces.IError,
  ) {
    result := []*GeoStateEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := GeoStateActionUpdate(query, record)
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
func (x *GeoStateEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var GeoStateEntityMeta = workspaces.TableMetaData{
	EntityName:    "GeoState",
	ExportKey:    "geo-states",
	TableNameInDb: "fb_geo-state_entities",
	EntityObject:  &GeoStateEntity{},
	ExportStream: GeoStateActionExportT,
	ImportQuery: GeoStateActionImport,
}
func GeoStateActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[GeoStateEntity](query, GeoStateActionQuery, GeoStatePreloadRelations)
}
func GeoStateActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[GeoStateEntity](query, GeoStateActionQuery, GeoStatePreloadRelations)
}
func GeoStateActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content GeoStateEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := GeoStateActionCreate(&content, query)
	return err
}
var GeoStateCommonCliFlags = []cli.Flag{
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
      Required: false,
      Usage:    "name",
    },
}
var GeoStateCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "name",
		StructField:     "Name",
		Required: false,
		Usage:    "name",
		Type: "string",
	},
}
var GeoStateCommonCliFlagsOptional = []cli.Flag{
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
      Required: false,
      Usage:    "name",
    },
}
  var GeoStateCreateCmd cli.Command = GEO_STATE_ACTION_POST_ONE.ToCli()
  var GeoStateCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_STATE_CREATE},
      })
      entity := &GeoStateEntity{}
      for _, item := range GeoStateCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := GeoStateActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var GeoStateUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: GeoStateCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_STATE_UPDATE},
      })
      entity := CastGeoStateFromCli(c)
      if entity, err := GeoStateActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* GeoStateEntity) FromCli(c *cli.Context) *GeoStateEntity {
	return CastGeoStateFromCli(c)
}
func CastGeoStateFromCli (c *cli.Context) *GeoStateEntity {
	template := &GeoStateEntity{}
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
  func GeoStateSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      GeoStateActionCreate,
      reflect.ValueOf(&GeoStateEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func GeoStateWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := GeoStateActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "GeoState", result)
    }
  }
var GeoStateImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_STATE_CREATE},
      })
			GeoStateActionSeeder(query, c.Int("count"))
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
				Value: "geo-state-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_STATE_CREATE},
      })
			GeoStateActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "geo-state-seeder-geo-state.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of geo-states, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]GeoStateEntity{}
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
			GeoStateCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				GeoStateActionCreate,
				reflect.ValueOf(&GeoStateEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_STATE_CREATE},
				},
        func() GeoStateEntity {
					v := CastGeoStateFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var GeoStateCliCommands []cli.Command = []cli.Command{
      GEO_STATE_ACTION_QUERY.ToCli(),
      GEO_STATE_ACTION_TABLE.ToCli(),
      GeoStateCreateCmd,
      GeoStateUpdateCmd,
      GeoStateCreateInteractiveCmd,
      GeoStateWipeCmd,
      workspaces.GetCommonRemoveQuery(reflect.ValueOf(&GeoStateEntity{}).Elem(), GeoStateActionRemove),
  }
  func GeoStateCliFn() cli.Command {
    GeoStateCliCommands = append(GeoStateCliCommands, GeoStateImportExportCommands...)
    return cli.Command{
      Name:        "state",
      Description: "GeoStates module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: GeoStateCliCommands,
    }
  }
var GEO_STATE_ACTION_TABLE = workspaces.Module2Action{
  Name:    "table",
  ActionAliases: []string{"t"},
  Flags:  workspaces.CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: GeoStateActionQuery,
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    workspaces.CommonCliTableCmd2(c,
      GeoStateActionQuery,
      security,
      reflect.ValueOf(&GeoStateEntity{}).Elem(),
    )
    return nil
  },
}
var GEO_STATE_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/geo-states",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_STATE_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, GeoStateActionQuery)
    },
  },
  Format: "QUERY",
  Action: GeoStateActionQuery,
  ResponseEntity: &[]GeoStateEntity{},
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			GeoStateActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionAliases: []string{"q"},
	Flags:         workspaces.CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var GEO_STATE_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/geo-states/export",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_STATE_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, GeoStateActionExport)
    },
  },
  Format: "QUERY",
  Action: GeoStateActionExport,
  ResponseEntity: &[]GeoStateEntity{},
}
var GEO_STATE_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/geo-state/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_STATE_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, GeoStateActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: GeoStateActionGetOne,
  ResponseEntity: &GeoStateEntity{},
}
var GEO_STATE_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new geoState",
  Flags: GeoStateCommonCliFlags,
  Method: "POST",
  Url:    "/geo-state",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_STATE_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, GeoStateActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, GeoStateActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: GeoStateActionCreate,
  Format: "POST_ONE",
  RequestEntity: &GeoStateEntity{},
  ResponseEntity: &GeoStateEntity{},
}
var GEO_STATE_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: GeoStateCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/geo-state",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_STATE_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, GeoStateActionUpdate)
    },
  },
  Action: GeoStateActionUpdate,
  RequestEntity: &GeoStateEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &GeoStateEntity{},
}
var GEO_STATE_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/geo-states",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_STATE_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, GeoStateActionBulkUpdate)
    },
  },
  Action: GeoStateActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[GeoStateEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[GeoStateEntity]{},
}
var GEO_STATE_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/geo-state",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_STATE_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, GeoStateActionRemove)
    },
  },
  Action: GeoStateActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &GeoStateEntity{},
}
  /**
  *	Override this function on GeoStateEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendGeoStateRouter = func(r *[]workspaces.Module2Action) {}
  func GetGeoStateModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      GEO_STATE_ACTION_QUERY,
      GEO_STATE_ACTION_EXPORT,
      GEO_STATE_ACTION_GET_ONE,
      GEO_STATE_ACTION_POST_ONE,
      GEO_STATE_ACTION_PATCH,
      GEO_STATE_ACTION_PATCH_BULK,
      GEO_STATE_ACTION_DELETE,
    }
    // Append user defined functions
    AppendGeoStateRouter(&routes)
    return routes
  }
  func CreateGeoStateRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetGeoStateModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, GeoStateEntityJsonSchema, "geo-state-http", "geo")
    workspaces.WriteEntitySchema("GeoStateEntity", GeoStateEntityJsonSchema, "geo")
    return httpRoutes
  }
var PERM_ROOT_GEO_STATE_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-state/delete",
}
var PERM_ROOT_GEO_STATE_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-state/create",
}
var PERM_ROOT_GEO_STATE_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-state/update",
}
var PERM_ROOT_GEO_STATE_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-state/query",
}
var PERM_ROOT_GEO_STATE = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-state/*",
}
var ALL_GEO_STATE_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_GEO_STATE_DELETE,
	PERM_ROOT_GEO_STATE_CREATE,
	PERM_ROOT_GEO_STATE_UPDATE,
	PERM_ROOT_GEO_STATE_QUERY,
	PERM_ROOT_GEO_STATE,
}