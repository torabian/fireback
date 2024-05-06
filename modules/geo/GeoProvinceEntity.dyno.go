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
	seeders "github.com/torabian/fireback/modules/geo/seeders/GeoProvince"
	metas "github.com/torabian/fireback/modules/geo/metas"
)
var geoProvinceSeedersFs = &seeders.ViewsFs
func ResetGeoProvinceSeeders(fs *embed.FS) {
	geoProvinceSeedersFs = fs
}
type GeoProvinceEntity struct {
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
    Country   *  GeoCountryEntity `json:"country" yaml:"country"    gorm:"foreignKey:CountryId;references:UniqueId"     `
    // Datenano also has a text representation
        CountryId *string `json:"countryId" yaml:"countryId"`
    Translations     []*GeoProvinceEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    Children []*GeoProvinceEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *GeoProvinceEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var GeoProvincePreloadRelations []string = []string{}
var GEO_PROVINCE_EVENT_CREATED = "geoProvince.created"
var GEO_PROVINCE_EVENT_UPDATED = "geoProvince.updated"
var GEO_PROVINCE_EVENT_DELETED = "geoProvince.deleted"
var GEO_PROVINCE_EVENTS = []string{
	GEO_PROVINCE_EVENT_CREATED,
	GEO_PROVINCE_EVENT_UPDATED,
	GEO_PROVINCE_EVENT_DELETED,
}
type GeoProvinceFieldMap struct {
		Name workspaces.TranslatedString `yaml:"name"`
		Country workspaces.TranslatedString `yaml:"country"`
}
var GeoProvinceEntityMetaConfig map[string]int64 = map[string]int64{
}
var GeoProvinceEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&GeoProvinceEntity{}))
  type GeoProvinceEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        Name string `yaml:"name" json:"name"`
  }
func entityGeoProvinceFormatter(dto *GeoProvinceEntity, query workspaces.QueryDSL) {
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
func GeoProvinceMockEntity() *GeoProvinceEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &GeoProvinceEntity{
      Name : &stringHolder,
	}
	return entity
}
func GeoProvinceActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := GeoProvinceMockEntity()
		_, err := GeoProvinceActionCreate(entity, query)
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
    func (x*GeoProvinceEntity) GetNameTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Name
          }
        }
      }
      return ""
    }
  func GeoProvinceActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*GeoProvinceEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &GeoProvinceEntity{
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
  func GeoProvinceAssociationCreate(dto *GeoProvinceEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func GeoProvinceRelationContentCreate(dto *GeoProvinceEntity, query workspaces.QueryDSL) error {
return nil
}
func GeoProvinceRelationContentUpdate(dto *GeoProvinceEntity, query workspaces.QueryDSL) error {
	return nil
}
func GeoProvincePolyglotCreateHandler(dto *GeoProvinceEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &GeoProvinceEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func GeoProvinceValidator(dto *GeoProvinceEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func GeoProvinceEntityPreSanitize(dto *GeoProvinceEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func GeoProvinceEntityBeforeCreateAppend(dto *GeoProvinceEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    GeoProvinceRecursiveAddUniqueId(dto, query)
  }
  func GeoProvinceRecursiveAddUniqueId(dto *GeoProvinceEntity, query workspaces.QueryDSL) {
  }
func GeoProvinceActionBatchCreateFn(dtos []*GeoProvinceEntity, query workspaces.QueryDSL) ([]*GeoProvinceEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*GeoProvinceEntity{}
		for _, item := range dtos {
			s, err := GeoProvinceActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func GeoProvinceDeleteEntireChildren(query workspaces.QueryDSL, dto *GeoProvinceEntity) (*workspaces.IError) {
  return nil
}
func GeoProvinceActionCreateFn(dto *GeoProvinceEntity, query workspaces.QueryDSL) (*GeoProvinceEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := GeoProvinceValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	GeoProvinceEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	GeoProvinceEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	GeoProvincePolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	GeoProvinceRelationContentCreate(dto, query)
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
	GeoProvinceAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(GEO_PROVINCE_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&GeoProvinceEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func GeoProvinceActionGetOne(query workspaces.QueryDSL) (*GeoProvinceEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&GeoProvinceEntity{})
    item, err := workspaces.GetOneEntity[GeoProvinceEntity](query, refl)
    entityGeoProvinceFormatter(item, query)
    return item, err
  }
  func GeoProvinceActionQuery(query workspaces.QueryDSL) ([]*GeoProvinceEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&GeoProvinceEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[GeoProvinceEntity](query, refl)
    for _, item := range items {
      entityGeoProvinceFormatter(item, query)
    }
    return items, meta, err
  }
  func GeoProvinceUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *GeoProvinceEntity) (*GeoProvinceEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = GEO_PROVINCE_EVENT_UPDATED
    GeoProvinceEntityPreSanitize(fields, query)
    var item GeoProvinceEntity
    q := dbref.
      Where(&GeoProvinceEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    GeoProvinceRelationContentUpdate(fields, query)
    GeoProvincePolyglotCreateHandler(fields, query)
    if ero := GeoProvinceDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&GeoProvinceEntity{UniqueId: uniqueId}).
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
  func GeoProvinceActionUpdateFn(query workspaces.QueryDSL, fields *GeoProvinceEntity) (*GeoProvinceEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := GeoProvinceValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // GeoProvinceRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *GeoProvinceEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = GeoProvinceUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return GeoProvinceUpdateExec(dbref, query, fields)
    }
  }
var GeoProvinceWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire geoprovinces ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_PROVINCE_DELETE},
    })
		count, _ := GeoProvinceActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func GeoProvinceActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&GeoProvinceEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_GEO_PROVINCE_DELETE}
	return workspaces.RemoveEntity[GeoProvinceEntity](query, refl)
}
func GeoProvinceActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[GeoProvinceEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'GeoProvinceEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func GeoProvinceActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[GeoProvinceEntity]) (
    *workspaces.BulkRecordRequest[GeoProvinceEntity], *workspaces.IError,
  ) {
    result := []*GeoProvinceEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := GeoProvinceActionUpdate(query, record)
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
func (x *GeoProvinceEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var GeoProvinceEntityMeta = workspaces.TableMetaData{
	EntityName:    "GeoProvince",
	ExportKey:    "geo-provinces",
	TableNameInDb: "fb_geo-province_entities",
	EntityObject:  &GeoProvinceEntity{},
	ExportStream: GeoProvinceActionExportT,
	ImportQuery: GeoProvinceActionImport,
}
func GeoProvinceActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[GeoProvinceEntity](query, GeoProvinceActionQuery, GeoProvincePreloadRelations)
}
func GeoProvinceActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[GeoProvinceEntity](query, GeoProvinceActionQuery, GeoProvincePreloadRelations)
}
func GeoProvinceActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content GeoProvinceEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := GeoProvinceActionCreate(&content, query)
	return err
}
var GeoProvinceCommonCliFlags = []cli.Flag{
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
      Name:     "country-id",
      Required: false,
      Usage:    "country",
    },
}
var GeoProvinceCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "name",
		StructField:     "Name",
		Required: false,
		Usage:    "name",
		Type: "string",
	},
}
var GeoProvinceCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "country-id",
      Required: false,
      Usage:    "country",
    },
}
  var GeoProvinceCreateCmd cli.Command = GEO_PROVINCE_ACTION_POST_ONE.ToCli()
  var GeoProvinceCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_PROVINCE_CREATE},
      })
      entity := &GeoProvinceEntity{}
      for _, item := range GeoProvinceCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := GeoProvinceActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var GeoProvinceUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: GeoProvinceCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_PROVINCE_UPDATE},
      })
      entity := CastGeoProvinceFromCli(c)
      if entity, err := GeoProvinceActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* GeoProvinceEntity) FromCli(c *cli.Context) *GeoProvinceEntity {
	return CastGeoProvinceFromCli(c)
}
func CastGeoProvinceFromCli (c *cli.Context) *GeoProvinceEntity {
	template := &GeoProvinceEntity{}
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
      if c.IsSet("country-id") {
        value := c.String("country-id")
        template.CountryId = &value
      }
	return template
}
  func GeoProvinceSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      GeoProvinceActionCreate,
      reflect.ValueOf(&GeoProvinceEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func GeoProvinceSyncSeeders() {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{WorkspaceId: workspaces.USER_SYSTEM},
      GeoProvinceActionCreate,
      reflect.ValueOf(&GeoProvinceEntity{}).Elem(),
      geoProvinceSeedersFs,
      []string{},
      true,
    )
  }
  func GeoProvinceWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := GeoProvinceActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "GeoProvince", result)
    }
  }
var GeoProvinceImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_PROVINCE_CREATE},
      })
			GeoProvinceActionSeeder(query, c.Int("count"))
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
				Value: "geo-province-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_PROVINCE_CREATE},
      })
			GeoProvinceActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "geo-province-seeder-geo-province.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of geo-provinces, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]GeoProvinceEntity{}
			workspaces.ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
	cli.Command{
		Name:  "list",
		Usage: "Prints the list of files attached to this module for syncing or bootstrapping project",
		Action: func(c *cli.Context) error {
			if entity, err := workspaces.GetSeederFilenames(geoProvinceSeedersFs, ""); err != nil {
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
				GeoProvinceActionCreate,
				reflect.ValueOf(&GeoProvinceEntity{}).Elem(),
				geoProvinceSeedersFs,
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
				GeoProvinceActionQuery,
				reflect.ValueOf(&GeoProvinceEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"GeoProvinceFieldMap.yml",
				GeoProvincePreloadRelations,
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
			GeoProvinceCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				GeoProvinceActionCreate,
				reflect.ValueOf(&GeoProvinceEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_PROVINCE_CREATE},
				},
        func() GeoProvinceEntity {
					v := CastGeoProvinceFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var GeoProvinceCliCommands []cli.Command = []cli.Command{
      GEO_PROVINCE_ACTION_QUERY.ToCli(),
      GEO_PROVINCE_ACTION_TABLE.ToCli(),
      GeoProvinceCreateCmd,
      GeoProvinceUpdateCmd,
      GeoProvinceCreateInteractiveCmd,
      GeoProvinceWipeCmd,
      workspaces.GetCommonRemoveQuery(reflect.ValueOf(&GeoProvinceEntity{}).Elem(), GeoProvinceActionRemove),
  }
  func GeoProvinceCliFn() cli.Command {
    GeoProvinceCliCommands = append(GeoProvinceCliCommands, GeoProvinceImportExportCommands...)
    return cli.Command{
      Name:        "province",
      Description: "GeoProvinces module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: GeoProvinceCliCommands,
    }
  }
var GEO_PROVINCE_ACTION_TABLE = workspaces.Module2Action{
  Name:    "table",
  ActionName: "table",
  ActionAliases: []string{"t"},
  Flags:  workspaces.CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: GeoProvinceActionQuery,
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    workspaces.CommonCliTableCmd2(c,
      GeoProvinceActionQuery,
      security,
      reflect.ValueOf(&GeoProvinceEntity{}).Elem(),
    )
    return nil
  },
}
var GEO_PROVINCE_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/geo-provinces",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_PROVINCE_QUERY},
  },
  Group: "geoProvince",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, GeoProvinceActionQuery)
    },
  },
  Format: "QUERY",
  Action: GeoProvinceActionQuery,
  ResponseEntity: &[]GeoProvinceEntity{},
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			GeoProvinceActionQuery,
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
var GEO_PROVINCE_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/geo-provinces/export",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_PROVINCE_QUERY},
  },
  Group: "geoProvince",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, GeoProvinceActionExport)
    },
  },
  Format: "QUERY",
  Action: GeoProvinceActionExport,
  ResponseEntity: &[]GeoProvinceEntity{},
}
var GEO_PROVINCE_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/geo-province/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_PROVINCE_QUERY},
  },
  Group: "geoProvince",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, GeoProvinceActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: GeoProvinceActionGetOne,
  ResponseEntity: &GeoProvinceEntity{},
}
var GEO_PROVINCE_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new geoProvince",
  Flags: GeoProvinceCommonCliFlags,
  Method: "POST",
  Url:    "/geo-province",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_PROVINCE_CREATE},
  },
  Group: "geoProvince",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, GeoProvinceActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, GeoProvinceActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: GeoProvinceActionCreate,
  Format: "POST_ONE",
  RequestEntity: &GeoProvinceEntity{},
  ResponseEntity: &GeoProvinceEntity{},
}
var GEO_PROVINCE_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: GeoProvinceCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/geo-province",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_PROVINCE_UPDATE},
  },
  Group: "geoProvince",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, GeoProvinceActionUpdate)
    },
  },
  Action: GeoProvinceActionUpdate,
  RequestEntity: &GeoProvinceEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &GeoProvinceEntity{},
}
var GEO_PROVINCE_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/geo-provinces",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_PROVINCE_UPDATE},
  },
  Group: "geoProvince",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, GeoProvinceActionBulkUpdate)
    },
  },
  Action: GeoProvinceActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[GeoProvinceEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[GeoProvinceEntity]{},
}
var GEO_PROVINCE_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/geo-province",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_PROVINCE_DELETE},
  },
  Group: "geoProvince",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, GeoProvinceActionRemove)
    },
  },
  Action: GeoProvinceActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &GeoProvinceEntity{},
}
  /**
  *	Override this function on GeoProvinceEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendGeoProvinceRouter = func(r *[]workspaces.Module2Action) {}
  func GetGeoProvinceModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      GEO_PROVINCE_ACTION_QUERY,
      GEO_PROVINCE_ACTION_EXPORT,
      GEO_PROVINCE_ACTION_GET_ONE,
      GEO_PROVINCE_ACTION_POST_ONE,
      GEO_PROVINCE_ACTION_PATCH,
      GEO_PROVINCE_ACTION_PATCH_BULK,
      GEO_PROVINCE_ACTION_DELETE,
    }
    // Append user defined functions
    AppendGeoProvinceRouter(&routes)
    return routes
  }
  func CreateGeoProvinceRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetGeoProvinceModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, GeoProvinceEntityJsonSchema, "geo-province-http", "geo")
    workspaces.WriteEntitySchema("GeoProvinceEntity", GeoProvinceEntityJsonSchema, "geo")
    return httpRoutes
  }
var PERM_ROOT_GEO_PROVINCE_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-province/delete",
  Name: "Delete geo province",
}
var PERM_ROOT_GEO_PROVINCE_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-province/create",
  Name: "Create geo province",
}
var PERM_ROOT_GEO_PROVINCE_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-province/update",
  Name: "Update geo province",
}
var PERM_ROOT_GEO_PROVINCE_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-province/query",
  Name: "Query geo province",
}
var PERM_ROOT_GEO_PROVINCE = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-province/*",
  Name: "Entire geo province actions (*)",
}
var ALL_GEO_PROVINCE_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_GEO_PROVINCE_DELETE,
	PERM_ROOT_GEO_PROVINCE_CREATE,
	PERM_ROOT_GEO_PROVINCE_UPDATE,
	PERM_ROOT_GEO_PROVINCE_QUERY,
	PERM_ROOT_GEO_PROVINCE,
}