package currency
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
	seeders "github.com/torabian/fireback/modules/currency/seeders/Currency"
	metas "github.com/torabian/fireback/modules/currency/metas"
)
type CurrencyEntity struct {
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
    Symbol   *string `json:"symbol" yaml:"symbol"       `
    // Datenano also has a text representation
    Name   *string `json:"name" yaml:"name"        translate:"true" `
    // Datenano also has a text representation
    SymbolNative   *string `json:"symbolNative" yaml:"symbolNative"       `
    // Datenano also has a text representation
    DecimalDigits   *int64 `json:"decimalDigits" yaml:"decimalDigits"       `
    // Datenano also has a text representation
    Rounding   *int64 `json:"rounding" yaml:"rounding"       `
    // Datenano also has a text representation
    Code   *string `json:"code" yaml:"code"       `
    // Datenano also has a text representation
    NamePlural   *string `json:"namePlural" yaml:"namePlural"       `
    // Datenano also has a text representation
    Translations     []*CurrencyEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    Children []*CurrencyEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *CurrencyEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var CurrencyPreloadRelations []string = []string{}
var CURRENCY_EVENT_CREATED = "currency.created"
var CURRENCY_EVENT_UPDATED = "currency.updated"
var CURRENCY_EVENT_DELETED = "currency.deleted"
var CURRENCY_EVENTS = []string{
	CURRENCY_EVENT_CREATED,
	CURRENCY_EVENT_UPDATED,
	CURRENCY_EVENT_DELETED,
}
type CurrencyFieldMap struct {
		Symbol workspaces.TranslatedString `yaml:"symbol"`
		Name workspaces.TranslatedString `yaml:"name"`
		SymbolNative workspaces.TranslatedString `yaml:"symbolNative"`
		DecimalDigits workspaces.TranslatedString `yaml:"decimalDigits"`
		Rounding workspaces.TranslatedString `yaml:"rounding"`
		Code workspaces.TranslatedString `yaml:"code"`
		NamePlural workspaces.TranslatedString `yaml:"namePlural"`
}
var CurrencyEntityMetaConfig map[string]int64 = map[string]int64{
}
var CurrencyEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&CurrencyEntity{}))
  type CurrencyEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        Name string `yaml:"name" json:"name"`
  }
func entityCurrencyFormatter(dto *CurrencyEntity, query workspaces.QueryDSL) {
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
func CurrencyMockEntity() *CurrencyEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &CurrencyEntity{
      Symbol : &stringHolder,
      Name : &stringHolder,
      SymbolNative : &stringHolder,
      DecimalDigits : &int64Holder,
      Rounding : &int64Holder,
      Code : &stringHolder,
      NamePlural : &stringHolder,
	}
	return entity
}
func CurrencyActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := CurrencyMockEntity()
		_, err := CurrencyActionCreate(entity, query)
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
    func (x*CurrencyEntity) GetNameTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Name
          }
        }
      }
      return ""
    }
  func CurrencyActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*CurrencyEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &CurrencyEntity{
          Symbol: &tildaRef,
          Name: &tildaRef,
          SymbolNative: &tildaRef,
          Code: &tildaRef,
          NamePlural: &tildaRef,
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
  func CurrencyAssociationCreate(dto *CurrencyEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func CurrencyRelationContentCreate(dto *CurrencyEntity, query workspaces.QueryDSL) error {
return nil
}
func CurrencyRelationContentUpdate(dto *CurrencyEntity, query workspaces.QueryDSL) error {
	return nil
}
func CurrencyPolyglotCreateHandler(dto *CurrencyEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &CurrencyEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func CurrencyValidator(dto *CurrencyEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func CurrencyEntityPreSanitize(dto *CurrencyEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func CurrencyEntityBeforeCreateAppend(dto *CurrencyEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    CurrencyRecursiveAddUniqueId(dto, query)
  }
  func CurrencyRecursiveAddUniqueId(dto *CurrencyEntity, query workspaces.QueryDSL) {
  }
func CurrencyActionBatchCreateFn(dtos []*CurrencyEntity, query workspaces.QueryDSL) ([]*CurrencyEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*CurrencyEntity{}
		for _, item := range dtos {
			s, err := CurrencyActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func CurrencyDeleteEntireChildren(query workspaces.QueryDSL, dto *CurrencyEntity) (*workspaces.IError) {
  return nil
}
func CurrencyActionCreateFn(dto *CurrencyEntity, query workspaces.QueryDSL) (*CurrencyEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := CurrencyValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	CurrencyEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	CurrencyEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	CurrencyPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	CurrencyRelationContentCreate(dto, query)
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
	CurrencyAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(CURRENCY_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&CurrencyEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func CurrencyActionGetOne(query workspaces.QueryDSL) (*CurrencyEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&CurrencyEntity{})
    item, err := workspaces.GetOneEntity[CurrencyEntity](query, refl)
    entityCurrencyFormatter(item, query)
    return item, err
  }
  func CurrencyActionQuery(query workspaces.QueryDSL) ([]*CurrencyEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&CurrencyEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[CurrencyEntity](query, refl)
    for _, item := range items {
      entityCurrencyFormatter(item, query)
    }
    return items, meta, err
  }
  func CurrencyUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *CurrencyEntity) (*CurrencyEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = CURRENCY_EVENT_UPDATED
    CurrencyEntityPreSanitize(fields, query)
    var item CurrencyEntity
    q := dbref.
      Where(&CurrencyEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    CurrencyRelationContentUpdate(fields, query)
    CurrencyPolyglotCreateHandler(fields, query)
    if ero := CurrencyDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&CurrencyEntity{UniqueId: uniqueId}).
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
  func CurrencyActionUpdateFn(query workspaces.QueryDSL, fields *CurrencyEntity) (*CurrencyEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := CurrencyValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // CurrencyRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *CurrencyEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = CurrencyUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return CurrencyUpdateExec(dbref, query, fields)
    }
  }
var CurrencyWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire currencies ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CURRENCY_DELETE},
    })
		count, _ := CurrencyActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func CurrencyActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&CurrencyEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_CURRENCY_DELETE}
	return workspaces.RemoveEntity[CurrencyEntity](query, refl)
}
func CurrencyActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[CurrencyEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'CurrencyEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func CurrencyActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[CurrencyEntity]) (
    *workspaces.BulkRecordRequest[CurrencyEntity], *workspaces.IError,
  ) {
    result := []*CurrencyEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := CurrencyActionUpdate(query, record)
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
func (x *CurrencyEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var CurrencyEntityMeta = workspaces.TableMetaData{
	EntityName:    "Currency",
	ExportKey:    "currencies",
	TableNameInDb: "fb_currency_entities",
	EntityObject:  &CurrencyEntity{},
	ExportStream: CurrencyActionExportT,
	ImportQuery: CurrencyActionImport,
}
func CurrencyActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[CurrencyEntity](query, CurrencyActionQuery, CurrencyPreloadRelations)
}
func CurrencyActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[CurrencyEntity](query, CurrencyActionQuery, CurrencyPreloadRelations)
}
func CurrencyActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content CurrencyEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := CurrencyActionCreate(&content, query)
	return err
}
var CurrencyCommonCliFlags = []cli.Flag{
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
      Name:     "symbol",
      Required: false,
      Usage:    "symbol",
    },
    &cli.StringFlag{
      Name:     "name",
      Required: false,
      Usage:    "name",
    },
    &cli.StringFlag{
      Name:     "symbol-native",
      Required: false,
      Usage:    "symbolNative",
    },
    &cli.Int64Flag{
      Name:     "decimal-digits",
      Required: false,
      Usage:    "decimalDigits",
    },
    &cli.Int64Flag{
      Name:     "rounding",
      Required: false,
      Usage:    "rounding",
    },
    &cli.StringFlag{
      Name:     "code",
      Required: false,
      Usage:    "code",
    },
    &cli.StringFlag{
      Name:     "name-plural",
      Required: false,
      Usage:    "namePlural",
    },
}
var CurrencyCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "symbol",
		StructField:     "Symbol",
		Required: false,
		Usage:    "symbol",
		Type: "string",
	},
	{
		Name:     "name",
		StructField:     "Name",
		Required: false,
		Usage:    "name",
		Type: "string",
	},
	{
		Name:     "symbolNative",
		StructField:     "SymbolNative",
		Required: false,
		Usage:    "symbolNative",
		Type: "string",
	},
	{
		Name:     "decimalDigits",
		StructField:     "DecimalDigits",
		Required: false,
		Usage:    "decimalDigits",
		Type: "int64",
	},
	{
		Name:     "rounding",
		StructField:     "Rounding",
		Required: false,
		Usage:    "rounding",
		Type: "int64",
	},
	{
		Name:     "code",
		StructField:     "Code",
		Required: false,
		Usage:    "code",
		Type: "string",
	},
	{
		Name:     "namePlural",
		StructField:     "NamePlural",
		Required: false,
		Usage:    "namePlural",
		Type: "string",
	},
}
var CurrencyCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "symbol",
      Required: false,
      Usage:    "symbol",
    },
    &cli.StringFlag{
      Name:     "name",
      Required: false,
      Usage:    "name",
    },
    &cli.StringFlag{
      Name:     "symbol-native",
      Required: false,
      Usage:    "symbolNative",
    },
    &cli.Int64Flag{
      Name:     "decimal-digits",
      Required: false,
      Usage:    "decimalDigits",
    },
    &cli.Int64Flag{
      Name:     "rounding",
      Required: false,
      Usage:    "rounding",
    },
    &cli.StringFlag{
      Name:     "code",
      Required: false,
      Usage:    "code",
    },
    &cli.StringFlag{
      Name:     "name-plural",
      Required: false,
      Usage:    "namePlural",
    },
}
  var CurrencyCreateCmd cli.Command = CURRENCY_ACTION_POST_ONE.ToCli()
  var CurrencyCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CURRENCY_CREATE},
      })
      entity := &CurrencyEntity{}
      for _, item := range CurrencyCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := CurrencyActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var CurrencyUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: CurrencyCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CURRENCY_UPDATE},
      })
      entity := CastCurrencyFromCli(c)
      if entity, err := CurrencyActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* CurrencyEntity) FromCli(c *cli.Context) *CurrencyEntity {
	return CastCurrencyFromCli(c)
}
func CastCurrencyFromCli (c *cli.Context) *CurrencyEntity {
	template := &CurrencyEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("symbol") {
        value := c.String("symbol")
        template.Symbol = &value
      }
      if c.IsSet("name") {
        value := c.String("name")
        template.Name = &value
      }
      if c.IsSet("symbol-native") {
        value := c.String("symbol-native")
        template.SymbolNative = &value
      }
      if c.IsSet("code") {
        value := c.String("code")
        template.Code = &value
      }
      if c.IsSet("name-plural") {
        value := c.String("name-plural")
        template.NamePlural = &value
      }
	return template
}
  func CurrencySyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      CurrencyActionCreate,
      reflect.ValueOf(&CurrencyEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func CurrencySyncSeeders() {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{WorkspaceId: workspaces.USER_SYSTEM},
      CurrencyActionCreate,
      reflect.ValueOf(&CurrencyEntity{}).Elem(),
      &seeders.ViewsFs,
      []string{},
      true,
    )
  }
  func CurrencyWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := CurrencyActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "Currency", result)
    }
  }
var CurrencyImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CURRENCY_CREATE},
      })
			CurrencyActionSeeder(query, c.Int("count"))
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
				Value: "currency-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CURRENCY_CREATE},
      })
			CurrencyActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "currency-seeder-currency.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of currencies, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]CurrencyEntity{}
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
				CurrencyActionCreate,
				reflect.ValueOf(&CurrencyEntity{}).Elem(),
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
				CurrencyActionQuery,
				reflect.ValueOf(&CurrencyEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"CurrencyFieldMap.yml",
				CurrencyPreloadRelations,
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
			CurrencyCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				CurrencyActionCreate,
				reflect.ValueOf(&CurrencyEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CURRENCY_CREATE},
				},
        func() CurrencyEntity {
					v := CastCurrencyFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var CurrencyCliCommands []cli.Command = []cli.Command{
      CURRENCY_ACTION_QUERY.ToCli(),
      CURRENCY_ACTION_TABLE.ToCli(),
      CurrencyCreateCmd,
      CurrencyUpdateCmd,
      CurrencyCreateInteractiveCmd,
      CurrencyWipeCmd,
      workspaces.GetCommonRemoveQuery(reflect.ValueOf(&CurrencyEntity{}).Elem(), CurrencyActionRemove),
  }
  func CurrencyCliFn() cli.Command {
    CurrencyCliCommands = append(CurrencyCliCommands, CurrencyImportExportCommands...)
    return cli.Command{
      Name:        "currency",
      ShortName:   "curr",
      Description: "Currencys module actions (sample module to handle complex entities)",
      Usage:       "List of all famous currencies, both internal and user defined ones",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: CurrencyCliCommands,
    }
  }
var CURRENCY_ACTION_TABLE = workspaces.Module2Action{
  Name:    "table",
  ActionAliases: []string{"t"},
  Flags:  workspaces.CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: CurrencyActionQuery,
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    workspaces.CommonCliTableCmd2(c,
      CurrencyActionQuery,
      security,
      reflect.ValueOf(&CurrencyEntity{}).Elem(),
    )
    return nil
  },
}
var CURRENCY_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/currencies",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CURRENCY_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, CurrencyActionQuery)
    },
  },
  Format: "QUERY",
  Action: CurrencyActionQuery,
  ResponseEntity: &[]CurrencyEntity{},
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			CurrencyActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionAliases: []string{"q"},
	Flags:         workspaces.CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var CURRENCY_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/currencies/export",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CURRENCY_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, CurrencyActionExport)
    },
  },
  Format: "QUERY",
  Action: CurrencyActionExport,
  ResponseEntity: &[]CurrencyEntity{},
}
var CURRENCY_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/currency/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CURRENCY_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, CurrencyActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: CurrencyActionGetOne,
  ResponseEntity: &CurrencyEntity{},
}
var CURRENCY_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new currency",
  Flags: CurrencyCommonCliFlags,
  Method: "POST",
  Url:    "/currency",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CURRENCY_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, CurrencyActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, CurrencyActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: CurrencyActionCreate,
  Format: "POST_ONE",
  RequestEntity: &CurrencyEntity{},
  ResponseEntity: &CurrencyEntity{},
}
var CURRENCY_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: CurrencyCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/currency",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CURRENCY_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, CurrencyActionUpdate)
    },
  },
  Action: CurrencyActionUpdate,
  RequestEntity: &CurrencyEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &CurrencyEntity{},
}
var CURRENCY_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/currencies",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CURRENCY_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, CurrencyActionBulkUpdate)
    },
  },
  Action: CurrencyActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[CurrencyEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[CurrencyEntity]{},
}
var CURRENCY_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/currency",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CURRENCY_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, CurrencyActionRemove)
    },
  },
  Action: CurrencyActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &CurrencyEntity{},
}
  /**
  *	Override this function on CurrencyEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendCurrencyRouter = func(r *[]workspaces.Module2Action) {}
  func GetCurrencyModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      CURRENCY_ACTION_QUERY,
      CURRENCY_ACTION_EXPORT,
      CURRENCY_ACTION_GET_ONE,
      CURRENCY_ACTION_POST_ONE,
      CURRENCY_ACTION_PATCH,
      CURRENCY_ACTION_PATCH_BULK,
      CURRENCY_ACTION_DELETE,
    }
    // Append user defined functions
    AppendCurrencyRouter(&routes)
    return routes
  }
  func CreateCurrencyRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetCurrencyModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, CurrencyEntityJsonSchema, "currency-http", "currency")
    workspaces.WriteEntitySchema("CurrencyEntity", CurrencyEntityJsonSchema, "currency")
    return httpRoutes
  }
var PERM_ROOT_CURRENCY_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/currency/currency/delete",
  Name: "Delete currency",
}
var PERM_ROOT_CURRENCY_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/currency/currency/create",
  Name: "Create currency",
}
var PERM_ROOT_CURRENCY_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/currency/currency/update",
  Name: "Update currency",
}
var PERM_ROOT_CURRENCY_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/currency/currency/query",
  Name: "Query currency",
}
var PERM_ROOT_CURRENCY = workspaces.PermissionInfo{
  CompleteKey: "root/currency/currency/*",
  Name: "Entire currency actions (*)",
}
var ALL_CURRENCY_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_CURRENCY_DELETE,
	PERM_ROOT_CURRENCY_CREATE,
	PERM_ROOT_CURRENCY_UPDATE,
	PERM_ROOT_CURRENCY_QUERY,
	PERM_ROOT_CURRENCY,
}