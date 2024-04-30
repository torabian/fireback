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
	seeders "github.com/torabian/fireback/modules/geo/seeders/GeoLocationType"
	metas "github.com/torabian/fireback/modules/geo/metas"
)
type GeoLocationTypeEntity struct {
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
    Translations     []*GeoLocationTypeEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    Children []*GeoLocationTypeEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *GeoLocationTypeEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var GeoLocationTypePreloadRelations []string = []string{}
var GEO_LOCATION_TYPE_EVENT_CREATED = "geoLocationType.created"
var GEO_LOCATION_TYPE_EVENT_UPDATED = "geoLocationType.updated"
var GEO_LOCATION_TYPE_EVENT_DELETED = "geoLocationType.deleted"
var GEO_LOCATION_TYPE_EVENTS = []string{
	GEO_LOCATION_TYPE_EVENT_CREATED,
	GEO_LOCATION_TYPE_EVENT_UPDATED,
	GEO_LOCATION_TYPE_EVENT_DELETED,
}
type GeoLocationTypeFieldMap struct {
		Name workspaces.TranslatedString `yaml:"name"`
}
var GeoLocationTypeEntityMetaConfig map[string]int64 = map[string]int64{
}
var GeoLocationTypeEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&GeoLocationTypeEntity{}))
  type GeoLocationTypeEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        Name string `yaml:"name" json:"name"`
  }
func entityGeoLocationTypeFormatter(dto *GeoLocationTypeEntity, query workspaces.QueryDSL) {
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
func GeoLocationTypeMockEntity() *GeoLocationTypeEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &GeoLocationTypeEntity{
      Name : &stringHolder,
	}
	return entity
}
func GeoLocationTypeActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := GeoLocationTypeMockEntity()
		_, err := GeoLocationTypeActionCreate(entity, query)
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
    func (x*GeoLocationTypeEntity) GetNameTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Name
          }
        }
      }
      return ""
    }
  func GeoLocationTypeActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*GeoLocationTypeEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &GeoLocationTypeEntity{
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
  func GeoLocationTypeAssociationCreate(dto *GeoLocationTypeEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func GeoLocationTypeRelationContentCreate(dto *GeoLocationTypeEntity, query workspaces.QueryDSL) error {
return nil
}
func GeoLocationTypeRelationContentUpdate(dto *GeoLocationTypeEntity, query workspaces.QueryDSL) error {
	return nil
}
func GeoLocationTypePolyglotCreateHandler(dto *GeoLocationTypeEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &GeoLocationTypeEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func GeoLocationTypeValidator(dto *GeoLocationTypeEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func GeoLocationTypeEntityPreSanitize(dto *GeoLocationTypeEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func GeoLocationTypeEntityBeforeCreateAppend(dto *GeoLocationTypeEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    GeoLocationTypeRecursiveAddUniqueId(dto, query)
  }
  func GeoLocationTypeRecursiveAddUniqueId(dto *GeoLocationTypeEntity, query workspaces.QueryDSL) {
  }
func GeoLocationTypeActionBatchCreateFn(dtos []*GeoLocationTypeEntity, query workspaces.QueryDSL) ([]*GeoLocationTypeEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*GeoLocationTypeEntity{}
		for _, item := range dtos {
			s, err := GeoLocationTypeActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func GeoLocationTypeDeleteEntireChildren(query workspaces.QueryDSL, dto *GeoLocationTypeEntity) (*workspaces.IError) {
  return nil
}
func GeoLocationTypeActionCreateFn(dto *GeoLocationTypeEntity, query workspaces.QueryDSL) (*GeoLocationTypeEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := GeoLocationTypeValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	GeoLocationTypeEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	GeoLocationTypeEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	GeoLocationTypePolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	GeoLocationTypeRelationContentCreate(dto, query)
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
	GeoLocationTypeAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(GEO_LOCATION_TYPE_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&GeoLocationTypeEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func GeoLocationTypeActionGetOne(query workspaces.QueryDSL) (*GeoLocationTypeEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&GeoLocationTypeEntity{})
    item, err := workspaces.GetOneEntity[GeoLocationTypeEntity](query, refl)
    entityGeoLocationTypeFormatter(item, query)
    return item, err
  }
  func GeoLocationTypeActionQuery(query workspaces.QueryDSL) ([]*GeoLocationTypeEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&GeoLocationTypeEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[GeoLocationTypeEntity](query, refl)
    for _, item := range items {
      entityGeoLocationTypeFormatter(item, query)
    }
    return items, meta, err
  }
  func GeoLocationTypeUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *GeoLocationTypeEntity) (*GeoLocationTypeEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = GEO_LOCATION_TYPE_EVENT_UPDATED
    GeoLocationTypeEntityPreSanitize(fields, query)
    var item GeoLocationTypeEntity
    q := dbref.
      Where(&GeoLocationTypeEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    GeoLocationTypeRelationContentUpdate(fields, query)
    GeoLocationTypePolyglotCreateHandler(fields, query)
    if ero := GeoLocationTypeDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&GeoLocationTypeEntity{UniqueId: uniqueId}).
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
  func GeoLocationTypeActionUpdateFn(query workspaces.QueryDSL, fields *GeoLocationTypeEntity) (*GeoLocationTypeEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := GeoLocationTypeValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // GeoLocationTypeRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *GeoLocationTypeEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = GeoLocationTypeUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return GeoLocationTypeUpdateExec(dbref, query, fields)
    }
  }
var GeoLocationTypeWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire geolocationtypes ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_LOCATION_TYPE_DELETE},
    })
		count, _ := GeoLocationTypeActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func GeoLocationTypeActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&GeoLocationTypeEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_GEO_LOCATION_TYPE_DELETE}
	return workspaces.RemoveEntity[GeoLocationTypeEntity](query, refl)
}
func GeoLocationTypeActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[GeoLocationTypeEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'GeoLocationTypeEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func GeoLocationTypeActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[GeoLocationTypeEntity]) (
    *workspaces.BulkRecordRequest[GeoLocationTypeEntity], *workspaces.IError,
  ) {
    result := []*GeoLocationTypeEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := GeoLocationTypeActionUpdate(query, record)
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
func (x *GeoLocationTypeEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var GeoLocationTypeEntityMeta = workspaces.TableMetaData{
	EntityName:    "GeoLocationType",
	ExportKey:    "geo-location-types",
	TableNameInDb: "fb_geo-location-type_entities",
	EntityObject:  &GeoLocationTypeEntity{},
	ExportStream: GeoLocationTypeActionExportT,
	ImportQuery: GeoLocationTypeActionImport,
}
func GeoLocationTypeActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[GeoLocationTypeEntity](query, GeoLocationTypeActionQuery, GeoLocationTypePreloadRelations)
}
func GeoLocationTypeActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[GeoLocationTypeEntity](query, GeoLocationTypeActionQuery, GeoLocationTypePreloadRelations)
}
func GeoLocationTypeActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content GeoLocationTypeEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := GeoLocationTypeActionCreate(&content, query)
	return err
}
var GeoLocationTypeCommonCliFlags = []cli.Flag{
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
var GeoLocationTypeCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "name",
		StructField:     "Name",
		Required: false,
		Usage:    "name",
		Type: "string",
	},
}
var GeoLocationTypeCommonCliFlagsOptional = []cli.Flag{
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
  var GeoLocationTypeCreateCmd cli.Command = GEO_LOCATION_TYPE_ACTION_POST_ONE.ToCli()
  var GeoLocationTypeCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_LOCATION_TYPE_CREATE},
      })
      entity := &GeoLocationTypeEntity{}
      for _, item := range GeoLocationTypeCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := GeoLocationTypeActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var GeoLocationTypeUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: GeoLocationTypeCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_LOCATION_TYPE_UPDATE},
      })
      entity := CastGeoLocationTypeFromCli(c)
      if entity, err := GeoLocationTypeActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* GeoLocationTypeEntity) FromCli(c *cli.Context) *GeoLocationTypeEntity {
	return CastGeoLocationTypeFromCli(c)
}
func CastGeoLocationTypeFromCli (c *cli.Context) *GeoLocationTypeEntity {
	template := &GeoLocationTypeEntity{}
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
  func GeoLocationTypeSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      GeoLocationTypeActionCreate,
      reflect.ValueOf(&GeoLocationTypeEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func GeoLocationTypeSyncSeeders() {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{WorkspaceId: workspaces.USER_SYSTEM},
      GeoLocationTypeActionCreate,
      reflect.ValueOf(&GeoLocationTypeEntity{}).Elem(),
      &seeders.ViewsFs,
      []string{},
      true,
    )
  }
  func GeoLocationTypeWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := GeoLocationTypeActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "GeoLocationType", result)
    }
  }
var GeoLocationTypeImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_LOCATION_TYPE_CREATE},
      })
			GeoLocationTypeActionSeeder(query, c.Int("count"))
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
				Value: "geo-location-type-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_LOCATION_TYPE_CREATE},
      })
			GeoLocationTypeActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "geo-location-type-seeder-geo-location-type.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of geo-location-types, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]GeoLocationTypeEntity{}
			workspaces.ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
	cli.Command{
		Name:  "list",
		Usage: "Prints the list of files attached to this module for syncing or bootstrapping project",
		Action: func(c *cli.Context) error {
			if entity, err := workspaces.GetSeederFilenames(&seeders.ViewsFs, ""); err != nil {
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
			workspaces.CommonCliImportEmbedCmd(c,
				GeoLocationTypeActionCreate,
				reflect.ValueOf(&GeoLocationTypeEntity{}).Elem(),
				&seeders.ViewsFs,
			)
			return nil
		},
	},
	cli.Command{
		Name:    "export",
		Aliases: []string{"e"},
		Flags: append(workspaces.CommonQueryFlags,
			&cli.StringFlag{
				Name:     "file",
				Usage:    "The address of file you want the csv/yaml/json be exported to",
				Required: true,
			}),
		Usage: "Exports a query results into the csv/yaml/json format",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliExportCmd(c,
				GeoLocationTypeActionQuery,
				reflect.ValueOf(&GeoLocationTypeEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"GeoLocationTypeFieldMap.yml",
				GeoLocationTypePreloadRelations,
			)
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
			GeoLocationTypeCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				GeoLocationTypeActionCreate,
				reflect.ValueOf(&GeoLocationTypeEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_LOCATION_TYPE_CREATE},
				},
        func() GeoLocationTypeEntity {
					v := CastGeoLocationTypeFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var GeoLocationTypeCliCommands []cli.Command = []cli.Command{
      GEO_LOCATION_TYPE_ACTION_QUERY.ToCli(),
      GEO_LOCATION_TYPE_ACTION_TABLE.ToCli(),
      GeoLocationTypeCreateCmd,
      GeoLocationTypeUpdateCmd,
      GeoLocationTypeCreateInteractiveCmd,
      GeoLocationTypeWipeCmd,
      workspaces.GetCommonRemoveQuery(reflect.ValueOf(&GeoLocationTypeEntity{}).Elem(), GeoLocationTypeActionRemove),
  }
  func GeoLocationTypeCliFn() cli.Command {
    GeoLocationTypeCliCommands = append(GeoLocationTypeCliCommands, GeoLocationTypeImportExportCommands...)
    return cli.Command{
      Name:        "type",
      Description: "GeoLocationTypes module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: GeoLocationTypeCliCommands,
    }
  }
var GEO_LOCATION_TYPE_ACTION_TABLE = workspaces.Module2Action{
  Name:    "table",
  ActionName: "table",
  ActionAliases: []string{"t"},
  Flags:  workspaces.CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: GeoLocationTypeActionQuery,
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    workspaces.CommonCliTableCmd2(c,
      GeoLocationTypeActionQuery,
      security,
      reflect.ValueOf(&GeoLocationTypeEntity{}).Elem(),
    )
    return nil
  },
}
var GEO_LOCATION_TYPE_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/geo-location-types",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_LOCATION_TYPE_QUERY},
  },
  Group: "geoLocationType",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, GeoLocationTypeActionQuery)
    },
  },
  Format: "QUERY",
  Action: GeoLocationTypeActionQuery,
  ResponseEntity: &[]GeoLocationTypeEntity{},
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			GeoLocationTypeActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionName:    "query",
	ActionAliases: []string{"q"},
	Flags:         workspaces.CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var GEO_LOCATION_TYPE_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/geo-location-types/export",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_LOCATION_TYPE_QUERY},
  },
  Group: "geoLocationType",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, GeoLocationTypeActionExport)
    },
  },
  Format: "QUERY",
  Action: GeoLocationTypeActionExport,
  ResponseEntity: &[]GeoLocationTypeEntity{},
}
var GEO_LOCATION_TYPE_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/geo-location-type/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_LOCATION_TYPE_QUERY},
  },
  Group: "geoLocationType",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, GeoLocationTypeActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: GeoLocationTypeActionGetOne,
  ResponseEntity: &GeoLocationTypeEntity{},
}
var GEO_LOCATION_TYPE_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new geoLocationType",
  Flags: GeoLocationTypeCommonCliFlags,
  Method: "POST",
  Url:    "/geo-location-type",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_LOCATION_TYPE_CREATE},
  },
  Group: "geoLocationType",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, GeoLocationTypeActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, GeoLocationTypeActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: GeoLocationTypeActionCreate,
  Format: "POST_ONE",
  RequestEntity: &GeoLocationTypeEntity{},
  ResponseEntity: &GeoLocationTypeEntity{},
}
var GEO_LOCATION_TYPE_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: GeoLocationTypeCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/geo-location-type",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_LOCATION_TYPE_UPDATE},
  },
  Group: "geoLocationType",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, GeoLocationTypeActionUpdate)
    },
  },
  Action: GeoLocationTypeActionUpdate,
  RequestEntity: &GeoLocationTypeEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &GeoLocationTypeEntity{},
}
var GEO_LOCATION_TYPE_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/geo-location-types",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_LOCATION_TYPE_UPDATE},
  },
  Group: "geoLocationType",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, GeoLocationTypeActionBulkUpdate)
    },
  },
  Action: GeoLocationTypeActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[GeoLocationTypeEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[GeoLocationTypeEntity]{},
}
var GEO_LOCATION_TYPE_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/geo-location-type",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_LOCATION_TYPE_DELETE},
  },
  Group: "geoLocationType",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, GeoLocationTypeActionRemove)
    },
  },
  Action: GeoLocationTypeActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &GeoLocationTypeEntity{},
}
  /**
  *	Override this function on GeoLocationTypeEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendGeoLocationTypeRouter = func(r *[]workspaces.Module2Action) {}
  func GetGeoLocationTypeModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      GEO_LOCATION_TYPE_ACTION_QUERY,
      GEO_LOCATION_TYPE_ACTION_EXPORT,
      GEO_LOCATION_TYPE_ACTION_GET_ONE,
      GEO_LOCATION_TYPE_ACTION_POST_ONE,
      GEO_LOCATION_TYPE_ACTION_PATCH,
      GEO_LOCATION_TYPE_ACTION_PATCH_BULK,
      GEO_LOCATION_TYPE_ACTION_DELETE,
    }
    // Append user defined functions
    AppendGeoLocationTypeRouter(&routes)
    return routes
  }
  func CreateGeoLocationTypeRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetGeoLocationTypeModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, GeoLocationTypeEntityJsonSchema, "geo-location-type-http", "geo")
    workspaces.WriteEntitySchema("GeoLocationTypeEntity", GeoLocationTypeEntityJsonSchema, "geo")
    return httpRoutes
  }
var PERM_ROOT_GEO_LOCATION_TYPE_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-location-type/delete",
  Name: "Delete geo location type",
}
var PERM_ROOT_GEO_LOCATION_TYPE_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-location-type/create",
  Name: "Create geo location type",
}
var PERM_ROOT_GEO_LOCATION_TYPE_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-location-type/update",
  Name: "Update geo location type",
}
var PERM_ROOT_GEO_LOCATION_TYPE_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-location-type/query",
  Name: "Query geo location type",
}
var PERM_ROOT_GEO_LOCATION_TYPE = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-location-type/*",
  Name: "Entire geo location type actions (*)",
}
var ALL_GEO_LOCATION_TYPE_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_GEO_LOCATION_TYPE_DELETE,
	PERM_ROOT_GEO_LOCATION_TYPE_CREATE,
	PERM_ROOT_GEO_LOCATION_TYPE_UPDATE,
	PERM_ROOT_GEO_LOCATION_TYPE_QUERY,
	PERM_ROOT_GEO_LOCATION_TYPE,
}