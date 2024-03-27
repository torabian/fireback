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
	seeders "github.com/torabian/fireback/modules/geo/seeders/GeoCountry"
	metas "github.com/torabian/fireback/modules/geo/metas"
)
type GeoCountryEntity struct {
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
    Status   *string `json:"status" yaml:"status"       `
    // Datenano also has a text representation
    Flag   *string `json:"flag" yaml:"flag"       `
    // Datenano also has a text representation
    CommonName   *string `json:"commonName" yaml:"commonName"        translate:"true" `
    // Datenano also has a text representation
    OfficialName   *string `json:"officialName" yaml:"officialName"        translate:"true" `
    // Datenano also has a text representation
    Translations     []*GeoCountryEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    Children []*GeoCountryEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *GeoCountryEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var GeoCountryPreloadRelations []string = []string{}
var GEO_COUNTRY_EVENT_CREATED = "geoCountry.created"
var GEO_COUNTRY_EVENT_UPDATED = "geoCountry.updated"
var GEO_COUNTRY_EVENT_DELETED = "geoCountry.deleted"
var GEO_COUNTRY_EVENTS = []string{
	GEO_COUNTRY_EVENT_CREATED,
	GEO_COUNTRY_EVENT_UPDATED,
	GEO_COUNTRY_EVENT_DELETED,
}
type GeoCountryFieldMap struct {
		Status workspaces.TranslatedString `yaml:"status"`
		Flag workspaces.TranslatedString `yaml:"flag"`
		CommonName workspaces.TranslatedString `yaml:"commonName"`
		OfficialName workspaces.TranslatedString `yaml:"officialName"`
}
var GeoCountryEntityMetaConfig map[string]int64 = map[string]int64{
}
var GeoCountryEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&GeoCountryEntity{}))
  type GeoCountryEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        CommonName string `yaml:"commonName" json:"commonName"`
        OfficialName string `yaml:"officialName" json:"officialName"`
  }
func entityGeoCountryFormatter(dto *GeoCountryEntity, query workspaces.QueryDSL) {
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
func GeoCountryMockEntity() *GeoCountryEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &GeoCountryEntity{
      Status : &stringHolder,
      Flag : &stringHolder,
      CommonName : &stringHolder,
      OfficialName : &stringHolder,
	}
	return entity
}
func GeoCountryActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := GeoCountryMockEntity()
		_, err := GeoCountryActionCreate(entity, query)
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
    func (x*GeoCountryEntity) GetCommonNameTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.CommonName
          }
        }
      }
      return ""
    }
    func (x*GeoCountryEntity) GetOfficialNameTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.OfficialName
          }
        }
      }
      return ""
    }
  func GeoCountryActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*GeoCountryEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &GeoCountryEntity{
          Status: &tildaRef,
          Flag: &tildaRef,
          CommonName: &tildaRef,
          OfficialName: &tildaRef,
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
  func GeoCountryAssociationCreate(dto *GeoCountryEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func GeoCountryRelationContentCreate(dto *GeoCountryEntity, query workspaces.QueryDSL) error {
return nil
}
func GeoCountryRelationContentUpdate(dto *GeoCountryEntity, query workspaces.QueryDSL) error {
	return nil
}
func GeoCountryPolyglotCreateHandler(dto *GeoCountryEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &GeoCountryEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func GeoCountryValidator(dto *GeoCountryEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func GeoCountryEntityPreSanitize(dto *GeoCountryEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func GeoCountryEntityBeforeCreateAppend(dto *GeoCountryEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    GeoCountryRecursiveAddUniqueId(dto, query)
  }
  func GeoCountryRecursiveAddUniqueId(dto *GeoCountryEntity, query workspaces.QueryDSL) {
  }
func GeoCountryActionBatchCreateFn(dtos []*GeoCountryEntity, query workspaces.QueryDSL) ([]*GeoCountryEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*GeoCountryEntity{}
		for _, item := range dtos {
			s, err := GeoCountryActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func GeoCountryDeleteEntireChildren(query workspaces.QueryDSL, dto *GeoCountryEntity) (*workspaces.IError) {
  return nil
}
func GeoCountryActionCreateFn(dto *GeoCountryEntity, query workspaces.QueryDSL) (*GeoCountryEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := GeoCountryValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	GeoCountryEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	GeoCountryEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	GeoCountryPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	GeoCountryRelationContentCreate(dto, query)
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
	GeoCountryAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(GEO_COUNTRY_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&GeoCountryEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func GeoCountryActionGetOne(query workspaces.QueryDSL) (*GeoCountryEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&GeoCountryEntity{})
    item, err := workspaces.GetOneEntity[GeoCountryEntity](query, refl)
    entityGeoCountryFormatter(item, query)
    return item, err
  }
  func GeoCountryActionQuery(query workspaces.QueryDSL) ([]*GeoCountryEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&GeoCountryEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[GeoCountryEntity](query, refl)
    for _, item := range items {
      entityGeoCountryFormatter(item, query)
    }
    return items, meta, err
  }
  func GeoCountryUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *GeoCountryEntity) (*GeoCountryEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = GEO_COUNTRY_EVENT_UPDATED
    GeoCountryEntityPreSanitize(fields, query)
    var item GeoCountryEntity
    q := dbref.
      Where(&GeoCountryEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    GeoCountryRelationContentUpdate(fields, query)
    GeoCountryPolyglotCreateHandler(fields, query)
    if ero := GeoCountryDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&GeoCountryEntity{UniqueId: uniqueId}).
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
  func GeoCountryActionUpdateFn(query workspaces.QueryDSL, fields *GeoCountryEntity) (*GeoCountryEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := GeoCountryValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // GeoCountryRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *GeoCountryEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = GeoCountryUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return GeoCountryUpdateExec(dbref, query, fields)
    }
  }
var GeoCountryWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire geocountries ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_COUNTRY_DELETE},
    })
		count, _ := GeoCountryActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func GeoCountryActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&GeoCountryEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_GEO_COUNTRY_DELETE}
	return workspaces.RemoveEntity[GeoCountryEntity](query, refl)
}
func GeoCountryActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[GeoCountryEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'GeoCountryEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func GeoCountryActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[GeoCountryEntity]) (
    *workspaces.BulkRecordRequest[GeoCountryEntity], *workspaces.IError,
  ) {
    result := []*GeoCountryEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := GeoCountryActionUpdate(query, record)
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
func (x *GeoCountryEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var GeoCountryEntityMeta = workspaces.TableMetaData{
	EntityName:    "GeoCountry",
	ExportKey:    "geo-countries",
	TableNameInDb: "fb_geo-country_entities",
	EntityObject:  &GeoCountryEntity{},
	ExportStream: GeoCountryActionExportT,
	ImportQuery: GeoCountryActionImport,
}
func GeoCountryActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[GeoCountryEntity](query, GeoCountryActionQuery, GeoCountryPreloadRelations)
}
func GeoCountryActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[GeoCountryEntity](query, GeoCountryActionQuery, GeoCountryPreloadRelations)
}
func GeoCountryActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content GeoCountryEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := GeoCountryActionCreate(&content, query)
	return err
}
var GeoCountryCommonCliFlags = []cli.Flag{
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
      Name:     "status",
      Required: false,
      Usage:    "status",
    },
    &cli.StringFlag{
      Name:     "flag",
      Required: false,
      Usage:    "flag",
    },
    &cli.StringFlag{
      Name:     "common-name",
      Required: false,
      Usage:    "commonName",
    },
    &cli.StringFlag{
      Name:     "official-name",
      Required: false,
      Usage:    "officialName",
    },
}
var GeoCountryCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "status",
		StructField:     "Status",
		Required: false,
		Usage:    "status",
		Type: "string",
	},
	{
		Name:     "flag",
		StructField:     "Flag",
		Required: false,
		Usage:    "flag",
		Type: "string",
	},
	{
		Name:     "commonName",
		StructField:     "CommonName",
		Required: false,
		Usage:    "commonName",
		Type: "string",
	},
	{
		Name:     "officialName",
		StructField:     "OfficialName",
		Required: false,
		Usage:    "officialName",
		Type: "string",
	},
}
var GeoCountryCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "status",
      Required: false,
      Usage:    "status",
    },
    &cli.StringFlag{
      Name:     "flag",
      Required: false,
      Usage:    "flag",
    },
    &cli.StringFlag{
      Name:     "common-name",
      Required: false,
      Usage:    "commonName",
    },
    &cli.StringFlag{
      Name:     "official-name",
      Required: false,
      Usage:    "officialName",
    },
}
  var GeoCountryCreateCmd cli.Command = GEO_COUNTRY_ACTION_POST_ONE.ToCli()
  var GeoCountryCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_COUNTRY_CREATE},
      })
      entity := &GeoCountryEntity{}
      for _, item := range GeoCountryCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := GeoCountryActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var GeoCountryUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: GeoCountryCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_COUNTRY_UPDATE},
      })
      entity := CastGeoCountryFromCli(c)
      if entity, err := GeoCountryActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* GeoCountryEntity) FromCli(c *cli.Context) *GeoCountryEntity {
	return CastGeoCountryFromCli(c)
}
func CastGeoCountryFromCli (c *cli.Context) *GeoCountryEntity {
	template := &GeoCountryEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("status") {
        value := c.String("status")
        template.Status = &value
      }
      if c.IsSet("flag") {
        value := c.String("flag")
        template.Flag = &value
      }
      if c.IsSet("common-name") {
        value := c.String("common-name")
        template.CommonName = &value
      }
      if c.IsSet("official-name") {
        value := c.String("official-name")
        template.OfficialName = &value
      }
	return template
}
  func GeoCountrySyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      GeoCountryActionCreate,
      reflect.ValueOf(&GeoCountryEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func GeoCountrySyncSeeders() {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{WorkspaceId: workspaces.USER_SYSTEM},
      GeoCountryActionCreate,
      reflect.ValueOf(&GeoCountryEntity{}).Elem(),
      &seeders.ViewsFs,
      []string{},
      true,
    )
  }
  func GeoCountryWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := GeoCountryActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "GeoCountry", result)
    }
  }
var GeoCountryImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_COUNTRY_CREATE},
      })
			GeoCountryActionSeeder(query, c.Int("count"))
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
				Value: "geo-country-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_COUNTRY_CREATE},
      })
			GeoCountryActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "geo-country-seeder-geo-country.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of geo-countries, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]GeoCountryEntity{}
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
				GeoCountryActionCreate,
				reflect.ValueOf(&GeoCountryEntity{}).Elem(),
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
				GeoCountryActionQuery,
				reflect.ValueOf(&GeoCountryEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"GeoCountryFieldMap.yml",
				GeoCountryPreloadRelations,
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
			GeoCountryCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				GeoCountryActionCreate,
				reflect.ValueOf(&GeoCountryEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_COUNTRY_CREATE},
				},
        func() GeoCountryEntity {
					v := CastGeoCountryFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var GeoCountryCliCommands []cli.Command = []cli.Command{
      GEO_COUNTRY_ACTION_QUERY.ToCli(),
      GEO_COUNTRY_ACTION_TABLE.ToCli(),
      GeoCountryCreateCmd,
      GeoCountryUpdateCmd,
      GeoCountryCreateInteractiveCmd,
      GeoCountryWipeCmd,
      workspaces.GetCommonRemoveQuery(reflect.ValueOf(&GeoCountryEntity{}).Elem(), GeoCountryActionRemove),
  }
  func GeoCountryCliFn() cli.Command {
    GeoCountryCliCommands = append(GeoCountryCliCommands, GeoCountryImportExportCommands...)
    return cli.Command{
      Name:        "country",
      Description: "GeoCountrys module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: GeoCountryCliCommands,
    }
  }
var GEO_COUNTRY_ACTION_TABLE = workspaces.Module2Action{
  Name:    "table",
  ActionAliases: []string{"t"},
  Flags:  workspaces.CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: GeoCountryActionQuery,
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    workspaces.CommonCliTableCmd2(c,
      GeoCountryActionQuery,
      security,
      reflect.ValueOf(&GeoCountryEntity{}).Elem(),
    )
    return nil
  },
}
var GEO_COUNTRY_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/geo-countries",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_COUNTRY_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, GeoCountryActionQuery)
    },
  },
  Format: "QUERY",
  Action: GeoCountryActionQuery,
  ResponseEntity: &[]GeoCountryEntity{},
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			GeoCountryActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionAliases: []string{"q"},
	Flags:         workspaces.CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var GEO_COUNTRY_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/geo-countries/export",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_COUNTRY_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, GeoCountryActionExport)
    },
  },
  Format: "QUERY",
  Action: GeoCountryActionExport,
  ResponseEntity: &[]GeoCountryEntity{},
}
var GEO_COUNTRY_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/geo-country/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_COUNTRY_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, GeoCountryActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: GeoCountryActionGetOne,
  ResponseEntity: &GeoCountryEntity{},
}
var GEO_COUNTRY_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new geoCountry",
  Flags: GeoCountryCommonCliFlags,
  Method: "POST",
  Url:    "/geo-country",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_COUNTRY_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, GeoCountryActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, GeoCountryActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: GeoCountryActionCreate,
  Format: "POST_ONE",
  RequestEntity: &GeoCountryEntity{},
  ResponseEntity: &GeoCountryEntity{},
}
var GEO_COUNTRY_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: GeoCountryCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/geo-country",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_COUNTRY_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, GeoCountryActionUpdate)
    },
  },
  Action: GeoCountryActionUpdate,
  RequestEntity: &GeoCountryEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &GeoCountryEntity{},
}
var GEO_COUNTRY_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/geo-countries",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_COUNTRY_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, GeoCountryActionBulkUpdate)
    },
  },
  Action: GeoCountryActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[GeoCountryEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[GeoCountryEntity]{},
}
var GEO_COUNTRY_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/geo-country",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_GEO_COUNTRY_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, GeoCountryActionRemove)
    },
  },
  Action: GeoCountryActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &GeoCountryEntity{},
}
  /**
  *	Override this function on GeoCountryEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendGeoCountryRouter = func(r *[]workspaces.Module2Action) {}
  func GetGeoCountryModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      GEO_COUNTRY_ACTION_QUERY,
      GEO_COUNTRY_ACTION_EXPORT,
      GEO_COUNTRY_ACTION_GET_ONE,
      GEO_COUNTRY_ACTION_POST_ONE,
      GEO_COUNTRY_ACTION_PATCH,
      GEO_COUNTRY_ACTION_PATCH_BULK,
      GEO_COUNTRY_ACTION_DELETE,
    }
    // Append user defined functions
    AppendGeoCountryRouter(&routes)
    return routes
  }
  func CreateGeoCountryRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetGeoCountryModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, GeoCountryEntityJsonSchema, "geo-country-http", "geo")
    workspaces.WriteEntitySchema("GeoCountryEntity", GeoCountryEntityJsonSchema, "geo")
    return httpRoutes
  }
var PERM_ROOT_GEO_COUNTRY_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-country/delete",
}
var PERM_ROOT_GEO_COUNTRY_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-country/create",
}
var PERM_ROOT_GEO_COUNTRY_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-country/update",
}
var PERM_ROOT_GEO_COUNTRY_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-country/query",
}
var PERM_ROOT_GEO_COUNTRY = workspaces.PermissionInfo{
  CompleteKey: "root/geo/geo-country/*",
}
var ALL_GEO_COUNTRY_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_GEO_COUNTRY_DELETE,
	PERM_ROOT_GEO_COUNTRY_CREATE,
	PERM_ROOT_GEO_COUNTRY_UPDATE,
	PERM_ROOT_GEO_COUNTRY_QUERY,
	PERM_ROOT_GEO_COUNTRY,
}