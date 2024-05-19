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
	mocks "github.com/torabian/fireback/modules/workspaces/mocks/GsmProvider"
)
var gsmProviderSeedersFs *embed.FS = nil
func ResetGsmProviderSeeders(fs *embed.FS) {
	gsmProviderSeedersFs = fs
}
type GsmProviderEntity struct {
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
    ApiKey   *string `json:"apiKey" yaml:"apiKey"       `
    // Datenano also has a text representation
    MainSenderNumber   *string `json:"mainSenderNumber" yaml:"mainSenderNumber"  validate:"required"       `
    // Datenano also has a text representation
    Type   *string `json:"type" yaml:"type"  validate:"required"       `
    // Datenano also has a text representation
    InvokeUrl   *string `json:"invokeUrl" yaml:"invokeUrl"       `
    // Datenano also has a text representation
    InvokeBody   *string `json:"invokeBody" yaml:"invokeBody"       `
    // Datenano also has a text representation
    Children []*GsmProviderEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *GsmProviderEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var GsmProviderPreloadRelations []string = []string{}
var GSM_PROVIDER_EVENT_CREATED = "gsmProvider.created"
var GSM_PROVIDER_EVENT_UPDATED = "gsmProvider.updated"
var GSM_PROVIDER_EVENT_DELETED = "gsmProvider.deleted"
var GSM_PROVIDER_EVENTS = []string{
	GSM_PROVIDER_EVENT_CREATED,
	GSM_PROVIDER_EVENT_UPDATED,
	GSM_PROVIDER_EVENT_DELETED,
}
type GsmProviderFieldMap struct {
		ApiKey TranslatedString `yaml:"apiKey"`
		MainSenderNumber TranslatedString `yaml:"mainSenderNumber"`
		Type TranslatedString `yaml:"type"`
		InvokeUrl TranslatedString `yaml:"invokeUrl"`
		InvokeBody TranslatedString `yaml:"invokeBody"`
}
var GsmProviderEntityMetaConfig map[string]int64 = map[string]int64{
}
var GsmProviderEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&GsmProviderEntity{}))
func entityGsmProviderFormatter(dto *GsmProviderEntity, query QueryDSL) {
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
func GsmProviderMockEntity() *GsmProviderEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &GsmProviderEntity{
      ApiKey : &stringHolder,
      MainSenderNumber : &stringHolder,
      Type : &stringHolder,
      InvokeUrl : &stringHolder,
      InvokeBody : &stringHolder,
	}
	return entity
}
func GsmProviderActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := GsmProviderMockEntity()
		_, err := GsmProviderActionCreate(entity, query)
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
  func GsmProviderActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*GsmProviderEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &GsmProviderEntity{
          ApiKey: &tildaRef,
          MainSenderNumber: &tildaRef,
          Type: &tildaRef,
          InvokeUrl: &tildaRef,
          InvokeBody: &tildaRef,
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
  func GsmProviderAssociationCreate(dto *GsmProviderEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func GsmProviderRelationContentCreate(dto *GsmProviderEntity, query QueryDSL) error {
return nil
}
func GsmProviderRelationContentUpdate(dto *GsmProviderEntity, query QueryDSL) error {
	return nil
}
func GsmProviderPolyglotCreateHandler(dto *GsmProviderEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func GsmProviderValidator(dto *GsmProviderEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func GsmProviderEntityPreSanitize(dto *GsmProviderEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func GsmProviderEntityBeforeCreateAppend(dto *GsmProviderEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    GsmProviderRecursiveAddUniqueId(dto, query)
  }
  func GsmProviderRecursiveAddUniqueId(dto *GsmProviderEntity, query QueryDSL) {
  }
func GsmProviderActionBatchCreateFn(dtos []*GsmProviderEntity, query QueryDSL) ([]*GsmProviderEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*GsmProviderEntity{}
		for _, item := range dtos {
			s, err := GsmProviderActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func GsmProviderDeleteEntireChildren(query QueryDSL, dto *GsmProviderEntity) (*IError) {
  // intentionally removed this. It's hard to implement it, and probably wrong without
  // proper on delete cascade
  return nil
}
func GsmProviderActionCreateFn(dto *GsmProviderEntity, query QueryDSL) (*GsmProviderEntity, *IError) {
	// 1. Validate always
	if iError := GsmProviderValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	GsmProviderEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	GsmProviderEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	GsmProviderPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	GsmProviderRelationContentCreate(dto, query)
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
	GsmProviderAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(GSM_PROVIDER_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&GsmProviderEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func GsmProviderActionGetOne(query QueryDSL) (*GsmProviderEntity, *IError) {
    refl := reflect.ValueOf(&GsmProviderEntity{})
    item, err := GetOneEntity[GsmProviderEntity](query, refl)
    entityGsmProviderFormatter(item, query)
    return item, err
  }
  func GsmProviderActionQuery(query QueryDSL) ([]*GsmProviderEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&GsmProviderEntity{})
    items, meta, err := QueryEntitiesPointer[GsmProviderEntity](query, refl)
    for _, item := range items {
      entityGsmProviderFormatter(item, query)
    }
    return items, meta, err
  }
  func GsmProviderUpdateExec(dbref *gorm.DB, query QueryDSL, fields *GsmProviderEntity) (*GsmProviderEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = GSM_PROVIDER_EVENT_UPDATED
    GsmProviderEntityPreSanitize(fields, query)
    var item GsmProviderEntity
    q := dbref.
      Where(&GsmProviderEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    GsmProviderRelationContentUpdate(fields, query)
    GsmProviderPolyglotCreateHandler(fields, query)
    if ero := GsmProviderDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&GsmProviderEntity{UniqueId: uniqueId}).
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
  func GsmProviderActionUpdateFn(query QueryDSL, fields *GsmProviderEntity) (*GsmProviderEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := GsmProviderValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // GsmProviderRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      var item *GsmProviderEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *IError
        item, err = GsmProviderUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, CastToIError(vf)
    } else {
      dbref = query.Tx
      return GsmProviderUpdateExec(dbref, query, fields)
    }
  }
var GsmProviderWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire gsmproviders ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
      ActionRequires: []PermissionInfo{PERM_ROOT_GSM_PROVIDER_DELETE},
    })
		count, _ := GsmProviderActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func GsmProviderActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&GsmProviderEntity{})
	query.ActionRequires = []PermissionInfo{PERM_ROOT_GSM_PROVIDER_DELETE}
	return RemoveEntity[GsmProviderEntity](query, refl)
}
func GsmProviderActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[GsmProviderEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'GsmProviderEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func GsmProviderActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[GsmProviderEntity]) (
    *BulkRecordRequest[GsmProviderEntity], *IError,
  ) {
    result := []*GsmProviderEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := GsmProviderActionUpdate(query, record)
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
func (x *GsmProviderEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var GsmProviderEntityMeta = TableMetaData{
	EntityName:    "GsmProvider",
	ExportKey:    "gsm-providers",
	TableNameInDb: "fb_gsm-provider_entities",
	EntityObject:  &GsmProviderEntity{},
	ExportStream: GsmProviderActionExportT,
	ImportQuery: GsmProviderActionImport,
}
func GsmProviderActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[GsmProviderEntity](query, GsmProviderActionQuery, GsmProviderPreloadRelations)
}
func GsmProviderActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[GsmProviderEntity](query, GsmProviderActionQuery, GsmProviderPreloadRelations)
}
func GsmProviderActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content GsmProviderEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := GsmProviderActionCreate(&content, query)
	return err
}
var GsmProviderCommonCliFlags = []cli.Flag{
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
      Name:     "api-key",
      Required: false,
      Usage:    "apiKey",
    },
    &cli.StringFlag{
      Name:     "main-sender-number",
      Required: true,
      Usage:    "mainSenderNumber",
    },
    &cli.StringFlag{
      Name:     "type",
      Required: true,
      Usage:    "One of: 'url', 'terminal', 'mediana'",
    },
    &cli.StringFlag{
      Name:     "invoke-url",
      Required: false,
      Usage:    "invokeUrl",
    },
    &cli.StringFlag{
      Name:     "invoke-body",
      Required: false,
      Usage:    "invokeBody",
    },
}
var GsmProviderCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:     "apiKey",
		StructField:     "ApiKey",
		Required: false,
		Usage:    "apiKey",
		Type: "string",
	},
	{
		Name:     "mainSenderNumber",
		StructField:     "MainSenderNumber",
		Required: true,
		Usage:    "mainSenderNumber",
		Type: "string",
	},
	{
		Name:     "type",
		StructField:     "Type",
		Required: true,
		Usage:    "One of: 'url', 'terminal', 'mediana'",
		Type: "string",
	},
	{
		Name:     "invokeUrl",
		StructField:     "InvokeUrl",
		Required: false,
		Usage:    "invokeUrl",
		Type: "string",
	},
	{
		Name:     "invokeBody",
		StructField:     "InvokeBody",
		Required: false,
		Usage:    "invokeBody",
		Type: "string",
	},
}
var GsmProviderCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "api-key",
      Required: false,
      Usage:    "apiKey",
    },
    &cli.StringFlag{
      Name:     "main-sender-number",
      Required: true,
      Usage:    "mainSenderNumber",
    },
    &cli.StringFlag{
      Name:     "type",
      Required: true,
      Usage:    "One of: 'url', 'terminal', 'mediana'",
    },
    &cli.StringFlag{
      Name:     "invoke-url",
      Required: false,
      Usage:    "invokeUrl",
    },
    &cli.StringFlag{
      Name:     "invoke-body",
      Required: false,
      Usage:    "invokeBody",
    },
}
  var GsmProviderCreateCmd cli.Command = GSM_PROVIDER_ACTION_POST_ONE.ToCli()
  var GsmProviderCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_GSM_PROVIDER_CREATE},
      })
      entity := &GsmProviderEntity{}
      for _, item := range GsmProviderCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := GsmProviderActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var GsmProviderUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: GsmProviderCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_GSM_PROVIDER_UPDATE},
      })
      entity := CastGsmProviderFromCli(c)
      if entity, err := GsmProviderActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* GsmProviderEntity) FromCli(c *cli.Context) *GsmProviderEntity {
	return CastGsmProviderFromCli(c)
}
func CastGsmProviderFromCli (c *cli.Context) *GsmProviderEntity {
	template := &GsmProviderEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("api-key") {
        value := c.String("api-key")
        template.ApiKey = &value
      }
      if c.IsSet("main-sender-number") {
        value := c.String("main-sender-number")
        template.MainSenderNumber = &value
      }
      if c.IsSet("type") {
        value := c.String("type")
        template.Type = &value
      }
      if c.IsSet("invoke-url") {
        value := c.String("invoke-url")
        template.InvokeUrl = &value
      }
      if c.IsSet("invoke-body") {
        value := c.String("invoke-body")
        template.InvokeBody = &value
      }
	return template
}
  func GsmProviderSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      GsmProviderActionCreate,
      reflect.ValueOf(&GsmProviderEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func GsmProviderImportMocks() {
    SeederFromFSImport(
      QueryDSL{},
      GsmProviderActionCreate,
      reflect.ValueOf(&GsmProviderEntity{}).Elem(),
      &mocks.ViewsFs,
      []string{},
      false,
    )
  }
  func GsmProviderWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := GsmProviderActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "GsmProvider", result)
    }
  }
var GsmProviderImportExportCommands = []cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_GSM_PROVIDER_CREATE},
      })
			GsmProviderActionSeeder(query, c.Int("count"))
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
				Value: "gsm-provider-seeder.yml",
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
        ActionRequires: []PermissionInfo{PERM_ROOT_GSM_PROVIDER_CREATE},
      })
			GsmProviderActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "gsm-provider-seeder-gsm-provider.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of gsm-providers, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]GsmProviderEntity{}
			ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
		cli.Command{
			Name:  "mocks",
			Usage: "Prints the list of mocks",
			Action: func(c *cli.Context) error {
				if entity, err := GetSeederFilenames(&mocks.ViewsFs, ""); err != nil {
					fmt.Println(err.Error())
				} else {
					f, _ := json.MarshalIndent(entity, "", "  ")
					fmt.Println(string(f))
				}
				return nil
			},
		},
		cli.Command{
			Name:  "msync",
			Usage: "Tries to sync mocks into the system",
			Action: func(c *cli.Context) error {
				CommonCliImportEmbedCmd(c,
					GsmProviderActionCreate,
					reflect.ValueOf(&GsmProviderEntity{}).Elem(),
					&mocks.ViewsFs,
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
			GsmProviderCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				GsmProviderActionCreate,
				reflect.ValueOf(&GsmProviderEntity{}).Elem(),
				c.String("file"),
        &SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_GSM_PROVIDER_CREATE},
				},
        func() GsmProviderEntity {
					v := CastGsmProviderFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var GsmProviderCliCommands []cli.Command = []cli.Command{
      GSM_PROVIDER_ACTION_QUERY.ToCli(),
      GSM_PROVIDER_ACTION_TABLE.ToCli(),
      GsmProviderCreateCmd,
      GsmProviderUpdateCmd,
      GsmProviderCreateInteractiveCmd,
      GsmProviderWipeCmd,
      GetCommonRemoveQuery(reflect.ValueOf(&GsmProviderEntity{}).Elem(), GsmProviderActionRemove),
  }
  func GsmProviderCliFn() cli.Command {
    GsmProviderCliCommands = append(GsmProviderCliCommands, GsmProviderImportExportCommands...)
    return cli.Command{
      Name:        "gsmProvider",
      Description: "GsmProviders module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: GsmProviderCliCommands,
    }
  }
var GSM_PROVIDER_ACTION_TABLE = Module2Action{
  Name:    "table",
  ActionName: "table",
  ActionAliases: []string{"t"},
  Flags:  CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: GsmProviderActionQuery,
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    CommonCliTableCmd2(c,
      GsmProviderActionQuery,
      security,
      reflect.ValueOf(&GsmProviderEntity{}).Elem(),
    )
    return nil
  },
}
var GSM_PROVIDER_ACTION_QUERY = Module2Action{
  Method: "GET",
  Url:    "/gsm-providers",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_GSM_PROVIDER_QUERY},
  },
  Group: "gsmProvider",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpQueryEntity(c, GsmProviderActionQuery)
    },
  },
  Format: "QUERY",
  Action: GsmProviderActionQuery,
  ResponseEntity: &[]GsmProviderEntity{},
  Out: Module2ActionBody{
		Entity: "GsmProviderEntity",
	},
  CliAction: func(c *cli.Context, security *SecurityModel) error {
		CommonCliQueryCmd2(
			c,
			GsmProviderActionQuery,
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
var GSM_PROVIDER_ACTION_EXPORT = Module2Action{
  Method: "GET",
  Url:    "/gsm-providers/export",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_GSM_PROVIDER_QUERY},
  },
  Group: "gsmProvider",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpStreamFileChannel(c, GsmProviderActionExport)
    },
  },
  Format: "QUERY",
  Action: GsmProviderActionExport,
  ResponseEntity: &[]GsmProviderEntity{},
  Out: Module2ActionBody{
		Entity: "GsmProviderEntity",
	},
}
var GSM_PROVIDER_ACTION_GET_ONE = Module2Action{
  Method: "GET",
  Url:    "/gsm-provider/:uniqueId",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_GSM_PROVIDER_QUERY},
  },
  Group: "gsmProvider",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpGetEntity(c, GsmProviderActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: GsmProviderActionGetOne,
  ResponseEntity: &GsmProviderEntity{},
  Out: Module2ActionBody{
		Entity: "GsmProviderEntity",
	},
}
var GSM_PROVIDER_ACTION_POST_ONE = Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new gsmProvider",
  Flags: GsmProviderCommonCliFlags,
  Method: "POST",
  Url:    "/gsm-provider",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_GSM_PROVIDER_CREATE},
  },
  Group: "gsmProvider",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpPostEntity(c, GsmProviderActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    result, err := CliPostEntity(c, GsmProviderActionCreate, security)
    HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: GsmProviderActionCreate,
  Format: "POST_ONE",
  RequestEntity: &GsmProviderEntity{},
  ResponseEntity: &GsmProviderEntity{},
  Out: Module2ActionBody{
		Entity: "GsmProviderEntity",
	},
  In: Module2ActionBody{
		Entity: "GsmProviderEntity",
	},
}
var GSM_PROVIDER_ACTION_PATCH = Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: GsmProviderCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/gsm-provider",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_GSM_PROVIDER_UPDATE},
  },
  Group: "gsmProvider",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntity(c, GsmProviderActionUpdate)
    },
  },
  Action: GsmProviderActionUpdate,
  RequestEntity: &GsmProviderEntity{},
  ResponseEntity: &GsmProviderEntity{},
  Format: "PATCH_ONE",
  Out: Module2ActionBody{
		Entity: "GsmProviderEntity",
	},
  In: Module2ActionBody{
		Entity: "GsmProviderEntity",
	},
}
var GSM_PROVIDER_ACTION_PATCH_BULK = Module2Action{
  Method: "PATCH",
  Url:    "/gsm-providers",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_GSM_PROVIDER_UPDATE},
  },
  Group: "gsmProvider",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntities(c, GsmProviderActionBulkUpdate)
    },
  },
  Action: GsmProviderActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &BulkRecordRequest[GsmProviderEntity]{},
  ResponseEntity: &BulkRecordRequest[GsmProviderEntity]{},
  Out: Module2ActionBody{
		Entity: "GsmProviderEntity",
	},
  In: Module2ActionBody{
		Entity: "GsmProviderEntity",
	},
}
var GSM_PROVIDER_ACTION_DELETE = Module2Action{
  Method: "DELETE",
  Url:    "/gsm-provider",
  Format: "DELETE_DSL",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_GSM_PROVIDER_DELETE},
  },
  Group: "gsmProvider",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpRemoveEntity(c, GsmProviderActionRemove)
    },
  },
  Action: GsmProviderActionRemove,
  RequestEntity: &DeleteRequest{},
  ResponseEntity: &DeleteResponse{},
  TargetEntity: &GsmProviderEntity{},
}
  /**
  *	Override this function on GsmProviderEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendGsmProviderRouter = func(r *[]Module2Action) {}
  func GetGsmProviderModule2Actions() []Module2Action {
    routes := []Module2Action{
      GSM_PROVIDER_ACTION_QUERY,
      GSM_PROVIDER_ACTION_EXPORT,
      GSM_PROVIDER_ACTION_GET_ONE,
      GSM_PROVIDER_ACTION_POST_ONE,
      GSM_PROVIDER_ACTION_PATCH,
      GSM_PROVIDER_ACTION_PATCH_BULK,
      GSM_PROVIDER_ACTION_DELETE,
    }
    // Append user defined functions
    AppendGsmProviderRouter(&routes)
    return routes
  }
  func CreateGsmProviderRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetGsmProviderModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, GsmProviderEntityJsonSchema, "gsm-provider-http", "workspaces")
    WriteEntitySchema("GsmProviderEntity", GsmProviderEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_GSM_PROVIDER_DELETE = PermissionInfo{
  CompleteKey: "root/workspaces/gsm-provider/delete",
  Name: "Delete gsm provider",
}
var PERM_ROOT_GSM_PROVIDER_CREATE = PermissionInfo{
  CompleteKey: "root/workspaces/gsm-provider/create",
  Name: "Create gsm provider",
}
var PERM_ROOT_GSM_PROVIDER_UPDATE = PermissionInfo{
  CompleteKey: "root/workspaces/gsm-provider/update",
  Name: "Update gsm provider",
}
var PERM_ROOT_GSM_PROVIDER_QUERY = PermissionInfo{
  CompleteKey: "root/workspaces/gsm-provider/query",
  Name: "Query gsm provider",
}
var PERM_ROOT_GSM_PROVIDER = PermissionInfo{
  CompleteKey: "root/workspaces/gsm-provider/*",
  Name: "Entire gsm provider actions (*)",
}
var ALL_GSM_PROVIDER_PERMISSIONS = []PermissionInfo{
	PERM_ROOT_GSM_PROVIDER_DELETE,
	PERM_ROOT_GSM_PROVIDER_CREATE,
	PERM_ROOT_GSM_PROVIDER_UPDATE,
	PERM_ROOT_GSM_PROVIDER_QUERY,
	PERM_ROOT_GSM_PROVIDER,
}
var GsmProviderType = newGsmProviderType()
func newGsmProviderType() *xGsmProviderType {
	return &xGsmProviderType{
      Url: "url",
      Terminal: "terminal",
      Mediana: "mediana",
	}
}
type xGsmProviderType struct {
    Url string
    Terminal string
    Mediana string
}
var GsmProviderEntityBundle = EntityBundle{
	Permissions: ALL_GSM_PROVIDER_PERMISSIONS,
	CliCommands: []cli.Command{
		GsmProviderCliFn(),
	},
	Actions: GetGsmProviderModule2Actions(),
	AutoMigrationEntities: []interface{}{
		&GsmProviderEntity{},
  	},
}