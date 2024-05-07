package shop
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
	seeders "github.com/torabian/fireback/modules/shop/seeders/PaymentMethod"
	metas "github.com/torabian/fireback/modules/shop/metas"
)
var paymentMethodSeedersFs = &seeders.ViewsFs
func ResetPaymentMethodSeeders(fs *embed.FS) {
	paymentMethodSeedersFs = fs
}
type PaymentMethodEntity struct {
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
    Description   *string `json:"description" yaml:"description"        translate:"true" `
    // Datenano also has a text representation
    Translations     []*PaymentMethodEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    Children []*PaymentMethodEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *PaymentMethodEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var PaymentMethodPreloadRelations []string = []string{}
var PAYMENT_METHOD_EVENT_CREATED = "paymentMethod.created"
var PAYMENT_METHOD_EVENT_UPDATED = "paymentMethod.updated"
var PAYMENT_METHOD_EVENT_DELETED = "paymentMethod.deleted"
var PAYMENT_METHOD_EVENTS = []string{
	PAYMENT_METHOD_EVENT_CREATED,
	PAYMENT_METHOD_EVENT_UPDATED,
	PAYMENT_METHOD_EVENT_DELETED,
}
type PaymentMethodFieldMap struct {
		Name workspaces.TranslatedString `yaml:"name"`
		Description workspaces.TranslatedString `yaml:"description"`
}
var PaymentMethodEntityMetaConfig map[string]int64 = map[string]int64{
}
var PaymentMethodEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&PaymentMethodEntity{}))
  type PaymentMethodEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        Name string `yaml:"name" json:"name"`
        Description string `yaml:"description" json:"description"`
  }
func entityPaymentMethodFormatter(dto *PaymentMethodEntity, query workspaces.QueryDSL) {
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
func PaymentMethodMockEntity() *PaymentMethodEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &PaymentMethodEntity{
      Name : &stringHolder,
      Description : &stringHolder,
	}
	return entity
}
func PaymentMethodActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := PaymentMethodMockEntity()
		_, err := PaymentMethodActionCreate(entity, query)
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
    func (x*PaymentMethodEntity) GetNameTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Name
          }
        }
      }
      return ""
    }
    func (x*PaymentMethodEntity) GetDescriptionTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Description
          }
        }
      }
      return ""
    }
  func PaymentMethodActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*PaymentMethodEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &PaymentMethodEntity{
          Name: &tildaRef,
          Description: &tildaRef,
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
  func PaymentMethodAssociationCreate(dto *PaymentMethodEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func PaymentMethodRelationContentCreate(dto *PaymentMethodEntity, query workspaces.QueryDSL) error {
return nil
}
func PaymentMethodRelationContentUpdate(dto *PaymentMethodEntity, query workspaces.QueryDSL) error {
	return nil
}
func PaymentMethodPolyglotCreateHandler(dto *PaymentMethodEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &PaymentMethodEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func PaymentMethodValidator(dto *PaymentMethodEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func PaymentMethodEntityPreSanitize(dto *PaymentMethodEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func PaymentMethodEntityBeforeCreateAppend(dto *PaymentMethodEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    PaymentMethodRecursiveAddUniqueId(dto, query)
  }
  func PaymentMethodRecursiveAddUniqueId(dto *PaymentMethodEntity, query workspaces.QueryDSL) {
  }
func PaymentMethodActionBatchCreateFn(dtos []*PaymentMethodEntity, query workspaces.QueryDSL) ([]*PaymentMethodEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*PaymentMethodEntity{}
		for _, item := range dtos {
			s, err := PaymentMethodActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func PaymentMethodDeleteEntireChildren(query workspaces.QueryDSL, dto *PaymentMethodEntity) (*workspaces.IError) {
  // intentionally removed this. It's hard to implement it, and probably wrong without
  // proper on delete cascade
  return nil
}
func PaymentMethodActionCreateFn(dto *PaymentMethodEntity, query workspaces.QueryDSL) (*PaymentMethodEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := PaymentMethodValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	PaymentMethodEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	PaymentMethodEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	PaymentMethodPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	PaymentMethodRelationContentCreate(dto, query)
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
	PaymentMethodAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(PAYMENT_METHOD_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&PaymentMethodEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func PaymentMethodActionGetOne(query workspaces.QueryDSL) (*PaymentMethodEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&PaymentMethodEntity{})
    item, err := workspaces.GetOneEntity[PaymentMethodEntity](query, refl)
    entityPaymentMethodFormatter(item, query)
    return item, err
  }
  func PaymentMethodActionQuery(query workspaces.QueryDSL) ([]*PaymentMethodEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&PaymentMethodEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[PaymentMethodEntity](query, refl)
    for _, item := range items {
      entityPaymentMethodFormatter(item, query)
    }
    return items, meta, err
  }
  func PaymentMethodUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *PaymentMethodEntity) (*PaymentMethodEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = PAYMENT_METHOD_EVENT_UPDATED
    PaymentMethodEntityPreSanitize(fields, query)
    var item PaymentMethodEntity
    q := dbref.
      Where(&PaymentMethodEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    PaymentMethodRelationContentUpdate(fields, query)
    PaymentMethodPolyglotCreateHandler(fields, query)
    if ero := PaymentMethodDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&PaymentMethodEntity{UniqueId: uniqueId}).
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
  func PaymentMethodActionUpdateFn(query workspaces.QueryDSL, fields *PaymentMethodEntity) (*PaymentMethodEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := PaymentMethodValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // PaymentMethodRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *PaymentMethodEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = PaymentMethodUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return PaymentMethodUpdateExec(dbref, query, fields)
    }
  }
var PaymentMethodWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire paymentmethods ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_METHOD_DELETE},
    })
		count, _ := PaymentMethodActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func PaymentMethodActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&PaymentMethodEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_METHOD_DELETE}
	return workspaces.RemoveEntity[PaymentMethodEntity](query, refl)
}
func PaymentMethodActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[PaymentMethodEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'PaymentMethodEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func PaymentMethodActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[PaymentMethodEntity]) (
    *workspaces.BulkRecordRequest[PaymentMethodEntity], *workspaces.IError,
  ) {
    result := []*PaymentMethodEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := PaymentMethodActionUpdate(query, record)
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
func (x *PaymentMethodEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var PaymentMethodEntityMeta = workspaces.TableMetaData{
	EntityName:    "PaymentMethod",
	ExportKey:    "payment-methods",
	TableNameInDb: "fb_payment-method_entities",
	EntityObject:  &PaymentMethodEntity{},
	ExportStream: PaymentMethodActionExportT,
	ImportQuery: PaymentMethodActionImport,
}
func PaymentMethodActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[PaymentMethodEntity](query, PaymentMethodActionQuery, PaymentMethodPreloadRelations)
}
func PaymentMethodActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[PaymentMethodEntity](query, PaymentMethodActionQuery, PaymentMethodPreloadRelations)
}
func PaymentMethodActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content PaymentMethodEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := PaymentMethodActionCreate(&content, query)
	return err
}
var PaymentMethodCommonCliFlags = []cli.Flag{
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
      Name:     "description",
      Required: false,
      Usage:    "description",
    },
}
var PaymentMethodCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "name",
		StructField:     "Name",
		Required: false,
		Usage:    "name",
		Type: "string",
	},
	{
		Name:     "description",
		StructField:     "Description",
		Required: false,
		Usage:    "description",
		Type: "string",
	},
}
var PaymentMethodCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "description",
      Required: false,
      Usage:    "description",
    },
}
  var PaymentMethodCreateCmd cli.Command = PAYMENT_METHOD_ACTION_POST_ONE.ToCli()
  var PaymentMethodCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_METHOD_CREATE},
      })
      entity := &PaymentMethodEntity{}
      for _, item := range PaymentMethodCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := PaymentMethodActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var PaymentMethodUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: PaymentMethodCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_METHOD_UPDATE},
      })
      entity := CastPaymentMethodFromCli(c)
      if entity, err := PaymentMethodActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* PaymentMethodEntity) FromCli(c *cli.Context) *PaymentMethodEntity {
	return CastPaymentMethodFromCli(c)
}
func CastPaymentMethodFromCli (c *cli.Context) *PaymentMethodEntity {
	template := &PaymentMethodEntity{}
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
      if c.IsSet("description") {
        value := c.String("description")
        template.Description = &value
      }
	return template
}
  func PaymentMethodSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      PaymentMethodActionCreate,
      reflect.ValueOf(&PaymentMethodEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func PaymentMethodSyncSeeders() {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{WorkspaceId: workspaces.USER_SYSTEM},
      PaymentMethodActionCreate,
      reflect.ValueOf(&PaymentMethodEntity{}).Elem(),
      paymentMethodSeedersFs,
      []string{},
      true,
    )
  }
  func PaymentMethodWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := PaymentMethodActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "PaymentMethod", result)
    }
  }
var PaymentMethodImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_METHOD_CREATE},
      })
			PaymentMethodActionSeeder(query, c.Int("count"))
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
				Value: "payment-method-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_METHOD_CREATE},
      })
			PaymentMethodActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "payment-method-seeder-payment-method.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of payment-methods, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]PaymentMethodEntity{}
			workspaces.ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
	cli.Command{
		Name:  "list",
		Usage: "Prints the list of files attached to this module for syncing or bootstrapping project",
		Action: func(c *cli.Context) error {
			if entity, err := workspaces.GetSeederFilenames(paymentMethodSeedersFs, ""); err != nil {
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
				PaymentMethodActionCreate,
				reflect.ValueOf(&PaymentMethodEntity{}).Elem(),
				paymentMethodSeedersFs,
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
				PaymentMethodActionQuery,
				reflect.ValueOf(&PaymentMethodEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"PaymentMethodFieldMap.yml",
				PaymentMethodPreloadRelations,
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
			PaymentMethodCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				PaymentMethodActionCreate,
				reflect.ValueOf(&PaymentMethodEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_METHOD_CREATE},
				},
        func() PaymentMethodEntity {
					v := CastPaymentMethodFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var PaymentMethodCliCommands []cli.Command = []cli.Command{
      PAYMENT_METHOD_ACTION_QUERY.ToCli(),
      PAYMENT_METHOD_ACTION_TABLE.ToCli(),
      PaymentMethodCreateCmd,
      PaymentMethodUpdateCmd,
      PaymentMethodCreateInteractiveCmd,
      PaymentMethodWipeCmd,
      workspaces.GetCommonRemoveQuery(reflect.ValueOf(&PaymentMethodEntity{}).Elem(), PaymentMethodActionRemove),
  }
  func PaymentMethodCliFn() cli.Command {
    PaymentMethodCliCommands = append(PaymentMethodCliCommands, PaymentMethodImportExportCommands...)
    return cli.Command{
      Name:        "paymentMethod",
      Description: "PaymentMethods module actions (sample module to handle complex entities)",
      Usage:       "Method of payment",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: PaymentMethodCliCommands,
    }
  }
var PAYMENT_METHOD_ACTION_TABLE = workspaces.Module2Action{
  Name:    "table",
  ActionName: "table",
  ActionAliases: []string{"t"},
  Flags:  workspaces.CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: PaymentMethodActionQuery,
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    workspaces.CommonCliTableCmd2(c,
      PaymentMethodActionQuery,
      security,
      reflect.ValueOf(&PaymentMethodEntity{}).Elem(),
    )
    return nil
  },
}
var PAYMENT_METHOD_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/payment-methods",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_METHOD_QUERY},
  },
  Group: "paymentMethod",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, PaymentMethodActionQuery)
    },
  },
  Format: "QUERY",
  Action: PaymentMethodActionQuery,
  ResponseEntity: &[]PaymentMethodEntity{},
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			PaymentMethodActionQuery,
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
var PAYMENT_METHOD_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/payment-methods/export",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_METHOD_QUERY},
  },
  Group: "paymentMethod",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, PaymentMethodActionExport)
    },
  },
  Format: "QUERY",
  Action: PaymentMethodActionExport,
  ResponseEntity: &[]PaymentMethodEntity{},
}
var PAYMENT_METHOD_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/payment-method/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_METHOD_QUERY},
  },
  Group: "paymentMethod",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, PaymentMethodActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: PaymentMethodActionGetOne,
  ResponseEntity: &PaymentMethodEntity{},
}
var PAYMENT_METHOD_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new paymentMethod",
  Flags: PaymentMethodCommonCliFlags,
  Method: "POST",
  Url:    "/payment-method",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_METHOD_CREATE},
  },
  Group: "paymentMethod",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, PaymentMethodActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, PaymentMethodActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: PaymentMethodActionCreate,
  Format: "POST_ONE",
  RequestEntity: &PaymentMethodEntity{},
  ResponseEntity: &PaymentMethodEntity{},
}
var PAYMENT_METHOD_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: PaymentMethodCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/payment-method",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_METHOD_UPDATE},
  },
  Group: "paymentMethod",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, PaymentMethodActionUpdate)
    },
  },
  Action: PaymentMethodActionUpdate,
  RequestEntity: &PaymentMethodEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &PaymentMethodEntity{},
}
var PAYMENT_METHOD_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/payment-methods",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_METHOD_UPDATE},
  },
  Group: "paymentMethod",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, PaymentMethodActionBulkUpdate)
    },
  },
  Action: PaymentMethodActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[PaymentMethodEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[PaymentMethodEntity]{},
}
var PAYMENT_METHOD_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/payment-method",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_METHOD_DELETE},
  },
  Group: "paymentMethod",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, PaymentMethodActionRemove)
    },
  },
  Action: PaymentMethodActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &PaymentMethodEntity{},
}
  /**
  *	Override this function on PaymentMethodEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendPaymentMethodRouter = func(r *[]workspaces.Module2Action) {}
  func GetPaymentMethodModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      PAYMENT_METHOD_ACTION_QUERY,
      PAYMENT_METHOD_ACTION_EXPORT,
      PAYMENT_METHOD_ACTION_GET_ONE,
      PAYMENT_METHOD_ACTION_POST_ONE,
      PAYMENT_METHOD_ACTION_PATCH,
      PAYMENT_METHOD_ACTION_PATCH_BULK,
      PAYMENT_METHOD_ACTION_DELETE,
    }
    // Append user defined functions
    AppendPaymentMethodRouter(&routes)
    return routes
  }
  func CreatePaymentMethodRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetPaymentMethodModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, PaymentMethodEntityJsonSchema, "payment-method-http", "shop")
    workspaces.WriteEntitySchema("PaymentMethodEntity", PaymentMethodEntityJsonSchema, "shop")
    return httpRoutes
  }
var PERM_ROOT_PAYMENT_METHOD_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/payment-method/delete",
  Name: "Delete payment method",
}
var PERM_ROOT_PAYMENT_METHOD_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/payment-method/create",
  Name: "Create payment method",
}
var PERM_ROOT_PAYMENT_METHOD_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/payment-method/update",
  Name: "Update payment method",
}
var PERM_ROOT_PAYMENT_METHOD_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/shop/payment-method/query",
  Name: "Query payment method",
}
var PERM_ROOT_PAYMENT_METHOD = workspaces.PermissionInfo{
  CompleteKey: "root/shop/payment-method/*",
  Name: "Entire payment method actions (*)",
}
var ALL_PAYMENT_METHOD_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_PAYMENT_METHOD_DELETE,
	PERM_ROOT_PAYMENT_METHOD_CREATE,
	PERM_ROOT_PAYMENT_METHOD_UPDATE,
	PERM_ROOT_PAYMENT_METHOD_QUERY,
	PERM_ROOT_PAYMENT_METHOD,
}