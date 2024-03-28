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
	seeders "github.com/torabian/fireback/modules/geo/seeders/GeoCity"
	metas "github.com/torabian/fireback/modules/geo/metas"
)
type GeoCityEntity struct {
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
    Name   *string `json:"name" yaml:"name"       `
    // Datenano also has a text representation
    Province   *  GeoProvinceEntity `json:"province" yaml:"province"    gorm:"foreignKey:ProvinceId;references:UniqueId"     `
    // Datenano also has a text representation
        ProvinceId *string `json:"provinceId" yaml:"provinceId"`
    State   *  GeoStateEntity `json:"state" yaml:"state"    gorm:"foreignKey:StateId;references:UniqueId"     `
    // Datenano also has a text representation
        StateId *string `json:"stateId" yaml:"stateId"`
    Country   *  GeoCountryEntity `json:"country" yaml:"country"    gorm:"foreignKey:CountryId;references:UniqueId"     `
    // Datenano also has a text representation
        CountryId *string `json:"countryId" yaml:"countryId"`
    Children []*GeoCityEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *GeoCityEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var GeoCityPreloadRelations []string = []string{}
var GEO_CITY_EVENT_CREATED = "geoCity.created"
var GEO_CITY_EVENT_UPDATED = "geoCity.updated"
var GEO_CITY_EVENT_DELETED = "geoCity.deleted"
var GEO_CITY_EVENTS = []string{
	GEO_CITY_EVENT_CREATED,
	GEO_CITY_EVENT_UPDATED,
	GEO_CITY_EVENT_DELETED,
}
type GeoCityFieldMap struct {
		Name workspaces.TranslatedString `yaml:"name"`
		Province workspaces.TranslatedString `yaml:"province"`
		State workspaces.TranslatedString `yaml:"state"`
		Country workspaces.TranslatedString `yaml:"country"`
}
var GeoCityEntityMetaConfig map[string]int64 = map[string]int64{
}
var GeoCityEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&GeoCityEntity{}))
func entityGeoCityFormatter(dto *GeoCityEntity, query workspaces.QueryDSL) {
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
func GeoCityMockEntity() *GeoCityEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &GeoCityEntity{
      Name : &stringHolder,
	}
	return entity
}
func GeoCityActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := GeoCityMockEntity()
		_, err := GeoCityActionCreate(entity, query)
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
  func GeoCityActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*GeoCityEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &GeoCityEntity{
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
  func GeoCityAssociationCreate(dto *GeoCityEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func GeoCityRelationContentCreate(dto *GeoCityEntity, query workspaces.QueryDSL) error {
return nil
}
func GeoCityRelationContentUpdate(dto *GeoCityEntity, query workspaces.QueryDSL) error {
	return nil
}
func GeoCityPolyglotCreateHandler(dto *GeoCityEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func GeoCityValidator(dto *GeoCityEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func GeoCityEntityPreSanitize(dto *GeoCityEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func GeoCityEntityBeforeCreateAppend(dto *GeoCityEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    GeoCityRecursiveAddUniqueId(dto, query)
  }
  func GeoCityRecursiveAddUniqueId(dto *GeoCityEntity, query workspaces.QueryDSL) {
  }
func GeoCityActionBatchCreateFn(dtos []*GeoCityEntity, query workspaces.QueryDSL) ([]*GeoCityEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*GeoCityEntity{}
		for _, item := range dtos {
			s, err := GeoCityActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func GeoCityDeleteEntireChildren(query workspaces.QueryDSL, dto *GeoCityEntity) (*workspaces.IError) {
  return nil
}
func GeoCityActionCreateFn(dto *GeoCityEntity, query workspaces.QueryDSL) (*GeoCityEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := GeoCityValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	GeoCityEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	GeoCityEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	GeoCityPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	GeoCityRelationContentCreate(dto, query)
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
	GeoCityAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(GEO_CITY_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&GeoCityEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func GeoCityActionGetOne(query workspaces.QueryDSL) (*GeoCityEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&GeoCityEntity{})
    item, err := workspaces.GetOneEntity[GeoCityEntity](query, refl)
    entityGeoCityFormatter(item, query)
    return item, err
  }
  func GeoCityActionQuery(query workspaces.QueryDSL) ([]*GeoCityEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&GeoCityEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[GeoCityEntity](query, refl)
    for _, item := range items {
      entityGeoCityFormatter(item, query)
    }
    return items, meta, err
  }
  func GeoCityUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *GeoCityEntity) (*GeoCityEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = GEO_CITY_EVENT_UPDATED
    GeoCityEntityPreSanitize(fields, query)
    var item GeoCityEntity
    q := dbref.
      Where(&GeoCityEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    GeoCityRelationContentUpdate(fields, query)
    GeoCityPolyglotCreateHandler(fields, query)
    if ero := GeoCityDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&GeoCityEntity{UniqueId: uniqueId}).
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
  func GeoCityActionUpdateFn(query workspaces.QueryDSL, fields *GeoCityEntity) (*GeoCityEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := GeoCityValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // GeoCityRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *GeoCityEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = GeoCityUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return GeoCityUpdateExec(dbref, query, fields)
    }
  }
var GeoCityWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire geocities ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_CITY_DELETE},
    })
		count, _ := GeoCityActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func GeoCityActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&GeoCityEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_GEO_CITY_DELETE}
	return workspaces.RemoveEntity[GeoCityEntity](query, refl)
}
func GeoCityActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[GeoCityEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'GeoCityEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func GeoCityActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[GeoCityEntity]) (
    *workspaces.BulkRecordRequest[GeoCityEntity], *workspaces.IError,
  ) {
    result := []*GeoCityEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := GeoCityActionUpdate(query, record)
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
func (x *GeoCityEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var GeoCityEntityMeta = workspaces.TableMetaData{
	EntityName:    "GeoCity",
	ExportKey:    "geo-cities",
	TableNameInDb: "fb_geo-city_entities",
	EntityObject:  &GeoCityEntity{},
	ExportStream: GeoCityActionExportT,
	ImportQuery: GeoCityActionImport,
}
func GeoCityActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[GeoCityEntity](query, GeoCityActionQuery, GeoCityPreloadRelations)
}
func GeoCityActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[GeoCityEntity](query, GeoCityActionQuery, GeoCityPreloadRelations)
}
func GeoCityActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content GeoCityEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := GeoCityActionCreate(&content, query)
	return err
}
var GeoCityCommonCliFlags = []cli.Flag{
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
    &cli.StringFlag{
      Name:     "province-id",
      Required: false,
      Usage:    "province",
    },
    &cli.StringFlag{
      Name:     "state-id",
      Required: false,
      Usage:    "state",
    },
    &cli.StringFlag{
      Name:     "country-id",
      Required: false,
      Usage:    "country",
    },
}
var GeoCityCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "name",
		StructField:     "Name",
		Required: false,
		Usage:    "name",
		Type: "string",
	},
}
var GeoCityCommonCliFlagsOptional = []cli.Flag{
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
    &cli.StringFlag{
      Name:     "province-id",
      Required: false,
      Usage:    "province",
    },
    &cli.StringFlag{
      Name:     "state-id",
      Required: false,
      Usage:    "state",
    },
    &cli.StringFlag{
      Name:     "country-id",
      Required: false,
      Usage:    "country",
    },
}
  var GeoCityCreateCmd cli.Command = GEO_CITY_ACTION_POST_ONE.ToCli()
  var GeoCityCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_CITY_CREATE},
      })
      entity := &GeoCityEntity{}
      for _, item := range GeoCityCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := GeoCityActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var GeoCityUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: GeoCityCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_CITY_UPDATE},
      })
      entity := CastGeoCityFromCli(c)
      if entity, err := GeoCityActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* GeoCityEntity) FromCli(c *cli.Context) *GeoCityEntity {
	return CastGeoCityFromCli(c)
}
func CastGeoCityFromCli (c *cli.Context) *GeoCityEntity {
	template := &GeoCityEntity{}
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
      if c.IsSet("province-id") {
        value := c.String("province-id")
        template.ProvinceId = &value
      }
      if c.IsSet("state-id") {
        value := c.String("state-id")
        template.StateId = &value
      }
      if c.IsSet("country-id") {
        value := c.String("country-id")
        template.CountryId = &value
      }
	return template
}
  func GeoCitySyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      GeoCityActionCreate,
      reflect.ValueOf(&GeoCityEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func GeoCitySyncSeeders() {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{WorkspaceId: workspaces.USER_SYSTEM},
      GeoCityActionCreate,
      reflect.ValueOf(&GeoCityEntity{}).Elem(),
      &seeders.ViewsFs,
      []string{},
      true,
    )
  }
  func GeoCityWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := GeoCityActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "GeoCity", result)
    }
  }
var GeoCityImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_CITY_CREATE},
      })
			GeoCityActionSeeder(query, c.Int("count"))
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
				Value: "geo-city-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_CITY_CREATE},
      })
			GeoCityActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "geo-city-seeder-geo-city.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of geo-cities, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]GeoCityEntity{}
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
				GeoCityActionCreate,
				reflect.ValueOf(&GeoCityEntity{}).Elem(),
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
				GeoCityActionQuery,
				reflect.ValueOf(&GeoCityEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"GeoCityFieldMap.yml",
				GeoCityPreloadRelations,
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
			GeoCityCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				GeoCityActionCreate,
				reflect.ValueOf(&GeoCityEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_CITY_CREATE},
				},
        func() GeoCityEntity {
					v := CastGeoCityFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var GeoCityCliCommands []cli.Command = []cli.Command{
      GEO_CITY_ACTION_QUERY.ToCli(),
      GEO_CITY_ACTION_TABLE.ToCli(),
      GeoCityCreateCmd,
      GeoCityUpdateCmd,
      GeoCityCreateInteractiveCmd,
      GeoCityWipeCmd,
      workspaces.GetCommonRemoveQuery(reflect.ValueOf(&GeoCityEntity{}).Elem(), GeoCityActionRemove),
  }
  func GeoCityCliFn() cli.Command {
    GeoCityCliCommands = append(GeoCityCliCommands, GeoCityImportExportCommands...)
    return cli.Command{
      Name:        "city",
      Description: "GeoCitys module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: GeoCityCliCommands,
    }
  }
var GEO_CITY_ACTION_TABLE = workspaces.Module2Action{
  Name:    "table",
  ActionAliases: []string{"t"},
  Flags:  workspaces.CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: GeoCityActionQuery,
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    workspaces.CommonCliTableCmd2(c,
      GeoCityActionQuery,
      security,
      reflect.ValueOf(&GeoCityEntity{}).Elem(),
    )
    return nil
  },
}
var GEO_CITY_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/geo-cities",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_CITY_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, GeoCityActionQuery)
    },
  },
  Format: "QUERY",
  Action: GeoCityActionQuery,
  ResponseEntity: &[]GeoCityEntity{},
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			GeoCityActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionAliases: []string{"q"},
	Flags:         workspaces.CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var GEO_CITY_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/geo-cities/export",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_CITY_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, GeoCityActionExport)
    },
  },
  Format: "QUERY",
  Action: GeoCityActionExport,
  ResponseEntity: &[]GeoCityEntity{},
}
var GEO_CITY_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/geo-city/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_CITY_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, GeoCityActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: GeoCityActionGetOne,
  ResponseEntity: &GeoCityEntity{},
}
var GEO_CITY_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new geoCity",
  Flags: GeoCityCommonCliFlags,
  Method: "POST",
  Url:    "/geo-city",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_CITY_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, GeoCityActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, GeoCityActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: GeoCityActionCreate,
  Format: "POST_ONE",
  RequestEntity: &GeoCityEntity{},
  ResponseEntity: &GeoCityEntity{},
}
var GEO_CITY_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: GeoCityCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/geo-city",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_CITY_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, GeoCityActionUpdate)
    },
  },
  Action: GeoCityActionUpdate,
  RequestEntity: &GeoCityEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &GeoCityEntity{},
}
var GEO_CITY_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/geo-cities",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_CITY_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, GeoCityActionBulkUpdate)
    },
  },
  Action: GeoCityActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[GeoCityEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[GeoCityEntity]{},
}
var GEO_CITY_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/geo-city",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_CITY_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, GeoCityActionRemove)
    },
  },
  Action: GeoCityActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &GeoCityEntity{},
}
  /**
  *	Override this function on GeoCityEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendGeoCityRouter = func(r *[]workspaces.Module2Action) {}
  func GetGeoCityModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      GEO_CITY_ACTION_QUERY,
      GEO_CITY_ACTION_EXPORT,
      GEO_CITY_ACTION_GET_ONE,
      GEO_CITY_ACTION_POST_ONE,
      GEO_CITY_ACTION_PATCH,
      GEO_CITY_ACTION_PATCH_BULK,
      GEO_CITY_ACTION_DELETE,
    }
    // Append user defined functions
    AppendGeoCityRouter(&routes)
    return routes
  }
  func CreateGeoCityRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetGeoCityModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, GeoCityEntityJsonSchema, "geo-city-http", "geo")
    workspaces.WriteEntitySchema("GeoCityEntity", GeoCityEntityJsonSchema, "geo")
    return httpRoutes
  }
var PERM_ROOT_GEO_CITY_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-city/delete",
  Name: "Delete geo city",
}
var PERM_ROOT_GEO_CITY_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-city/create",
  Name: "Create geo city",
}
var PERM_ROOT_GEO_CITY_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-city/update",
  Name: "Update geo city",
}
var PERM_ROOT_GEO_CITY_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-city/query",
  Name: "Query geo city",
}
var PERM_ROOT_GEO_CITY = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-city/*",
  Name: "Entire geo city actions (*)",
}
var ALL_GEO_CITY_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_GEO_CITY_DELETE,
	PERM_ROOT_GEO_CITY_CREATE,
	PERM_ROOT_GEO_CITY_UPDATE,
	PERM_ROOT_GEO_CITY_QUERY,
	PERM_ROOT_GEO_CITY,
}