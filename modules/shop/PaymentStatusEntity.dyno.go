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
	seeders "github.com/torabian/fireback/modules/shop/seeders/PaymentStatus"
	metas "github.com/torabian/fireback/modules/shop/metas"
)
var paymentStatusSeedersFs = &seeders.ViewsFs
func ResetPaymentStatusSeeders(fs *embed.FS) {
	paymentStatusSeedersFs = fs
}
type PaymentStatusEntity struct {
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
    Translations     []*PaymentStatusEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    Children []*PaymentStatusEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *PaymentStatusEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var PaymentStatusPreloadRelations []string = []string{}
var PAYMENT_STATUS_EVENT_CREATED = "paymentStatus.created"
var PAYMENT_STATUS_EVENT_UPDATED = "paymentStatus.updated"
var PAYMENT_STATUS_EVENT_DELETED = "paymentStatus.deleted"
var PAYMENT_STATUS_EVENTS = []string{
	PAYMENT_STATUS_EVENT_CREATED,
	PAYMENT_STATUS_EVENT_UPDATED,
	PAYMENT_STATUS_EVENT_DELETED,
}
type PaymentStatusFieldMap struct {
		Name workspaces.TranslatedString `yaml:"name"`
		Description workspaces.TranslatedString `yaml:"description"`
}
var PaymentStatusEntityMetaConfig map[string]int64 = map[string]int64{
}
var PaymentStatusEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&PaymentStatusEntity{}))
  type PaymentStatusEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        Name string `yaml:"name" json:"name"`
        Description string `yaml:"description" json:"description"`
  }
func entityPaymentStatusFormatter(dto *PaymentStatusEntity, query workspaces.QueryDSL) {
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
func PaymentStatusMockEntity() *PaymentStatusEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &PaymentStatusEntity{
      Name : &stringHolder,
      Description : &stringHolder,
	}
	return entity
}
func PaymentStatusActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := PaymentStatusMockEntity()
		_, err := PaymentStatusActionCreate(entity, query)
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
    func (x*PaymentStatusEntity) GetNameTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Name
          }
        }
      }
      return ""
    }
    func (x*PaymentStatusEntity) GetDescriptionTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Description
          }
        }
      }
      return ""
    }
  func PaymentStatusActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*PaymentStatusEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &PaymentStatusEntity{
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
  func PaymentStatusAssociationCreate(dto *PaymentStatusEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func PaymentStatusRelationContentCreate(dto *PaymentStatusEntity, query workspaces.QueryDSL) error {
return nil
}
func PaymentStatusRelationContentUpdate(dto *PaymentStatusEntity, query workspaces.QueryDSL) error {
	return nil
}
func PaymentStatusPolyglotCreateHandler(dto *PaymentStatusEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &PaymentStatusEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func PaymentStatusValidator(dto *PaymentStatusEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func PaymentStatusEntityPreSanitize(dto *PaymentStatusEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func PaymentStatusEntityBeforeCreateAppend(dto *PaymentStatusEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    PaymentStatusRecursiveAddUniqueId(dto, query)
  }
  func PaymentStatusRecursiveAddUniqueId(dto *PaymentStatusEntity, query workspaces.QueryDSL) {
  }
func PaymentStatusActionBatchCreateFn(dtos []*PaymentStatusEntity, query workspaces.QueryDSL) ([]*PaymentStatusEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*PaymentStatusEntity{}
		for _, item := range dtos {
			s, err := PaymentStatusActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func PaymentStatusDeleteEntireChildren(query workspaces.QueryDSL, dto *PaymentStatusEntity) (*workspaces.IError) {
  // intentionally removed this. It's hard to implement it, and probably wrong without
  // proper on delete cascade
  return nil
}
func PaymentStatusActionCreateFn(dto *PaymentStatusEntity, query workspaces.QueryDSL) (*PaymentStatusEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := PaymentStatusValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	PaymentStatusEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	PaymentStatusEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	PaymentStatusPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	PaymentStatusRelationContentCreate(dto, query)
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
	PaymentStatusAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(PAYMENT_STATUS_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&PaymentStatusEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func PaymentStatusActionGetOne(query workspaces.QueryDSL) (*PaymentStatusEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&PaymentStatusEntity{})
    item, err := workspaces.GetOneEntity[PaymentStatusEntity](query, refl)
    entityPaymentStatusFormatter(item, query)
    return item, err
  }
  func PaymentStatusActionQuery(query workspaces.QueryDSL) ([]*PaymentStatusEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&PaymentStatusEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[PaymentStatusEntity](query, refl)
    for _, item := range items {
      entityPaymentStatusFormatter(item, query)
    }
    return items, meta, err
  }
  func PaymentStatusUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *PaymentStatusEntity) (*PaymentStatusEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = PAYMENT_STATUS_EVENT_UPDATED
    PaymentStatusEntityPreSanitize(fields, query)
    var item PaymentStatusEntity
    q := dbref.
      Where(&PaymentStatusEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    PaymentStatusRelationContentUpdate(fields, query)
    PaymentStatusPolyglotCreateHandler(fields, query)
    if ero := PaymentStatusDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&PaymentStatusEntity{UniqueId: uniqueId}).
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
  func PaymentStatusActionUpdateFn(query workspaces.QueryDSL, fields *PaymentStatusEntity) (*PaymentStatusEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := PaymentStatusValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // PaymentStatusRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *PaymentStatusEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = PaymentStatusUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return PaymentStatusUpdateExec(dbref, query, fields)
    }
  }
var PaymentStatusWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire paymentstatuses ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_STATUS_DELETE},
    })
		count, _ := PaymentStatusActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func PaymentStatusActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&PaymentStatusEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_STATUS_DELETE}
	return workspaces.RemoveEntity[PaymentStatusEntity](query, refl)
}
func PaymentStatusActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[PaymentStatusEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'PaymentStatusEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func PaymentStatusActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[PaymentStatusEntity]) (
    *workspaces.BulkRecordRequest[PaymentStatusEntity], *workspaces.IError,
  ) {
    result := []*PaymentStatusEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := PaymentStatusActionUpdate(query, record)
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
func (x *PaymentStatusEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var PaymentStatusEntityMeta = workspaces.TableMetaData{
	EntityName:    "PaymentStatus",
	ExportKey:    "payment-statuses",
	TableNameInDb: "fb_payment-status_entities",
	EntityObject:  &PaymentStatusEntity{},
	ExportStream: PaymentStatusActionExportT,
	ImportQuery: PaymentStatusActionImport,
}
func PaymentStatusActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[PaymentStatusEntity](query, PaymentStatusActionQuery, PaymentStatusPreloadRelations)
}
func PaymentStatusActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[PaymentStatusEntity](query, PaymentStatusActionQuery, PaymentStatusPreloadRelations)
}
func PaymentStatusActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content PaymentStatusEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := PaymentStatusActionCreate(&content, query)
	return err
}
var PaymentStatusCommonCliFlags = []cli.Flag{
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
var PaymentStatusCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
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
var PaymentStatusCommonCliFlagsOptional = []cli.Flag{
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
  var PaymentStatusCreateCmd cli.Command = PAYMENT_STATUS_ACTION_POST_ONE.ToCli()
  var PaymentStatusCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_STATUS_CREATE},
      })
      entity := &PaymentStatusEntity{}
      for _, item := range PaymentStatusCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := PaymentStatusActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var PaymentStatusUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: PaymentStatusCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_STATUS_UPDATE},
      })
      entity := CastPaymentStatusFromCli(c)
      if entity, err := PaymentStatusActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* PaymentStatusEntity) FromCli(c *cli.Context) *PaymentStatusEntity {
	return CastPaymentStatusFromCli(c)
}
func CastPaymentStatusFromCli (c *cli.Context) *PaymentStatusEntity {
	template := &PaymentStatusEntity{}
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
  func PaymentStatusSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      PaymentStatusActionCreate,
      reflect.ValueOf(&PaymentStatusEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func PaymentStatusSyncSeeders() {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{WorkspaceId: workspaces.USER_SYSTEM},
      PaymentStatusActionCreate,
      reflect.ValueOf(&PaymentStatusEntity{}).Elem(),
      paymentStatusSeedersFs,
      []string{},
      true,
    )
  }
  func PaymentStatusWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := PaymentStatusActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "PaymentStatus", result)
    }
  }
var PaymentStatusImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_STATUS_CREATE},
      })
			PaymentStatusActionSeeder(query, c.Int("count"))
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
				Value: "payment-status-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_STATUS_CREATE},
      })
			PaymentStatusActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "payment-status-seeder-payment-status.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of payment-statuses, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]PaymentStatusEntity{}
			workspaces.ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
	cli.Command{
		Name:  "list",
		Usage: "Prints the list of files attached to this module for syncing or bootstrapping project",
		Action: func(c *cli.Context) error {
			if entity, err := workspaces.GetSeederFilenames(paymentStatusSeedersFs, ""); err != nil {
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
				PaymentStatusActionCreate,
				reflect.ValueOf(&PaymentStatusEntity{}).Elem(),
				paymentStatusSeedersFs,
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
				PaymentStatusActionQuery,
				reflect.ValueOf(&PaymentStatusEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"PaymentStatusFieldMap.yml",
				PaymentStatusPreloadRelations,
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
			PaymentStatusCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				PaymentStatusActionCreate,
				reflect.ValueOf(&PaymentStatusEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_STATUS_CREATE},
				},
        func() PaymentStatusEntity {
					v := CastPaymentStatusFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var PaymentStatusCliCommands []cli.Command = []cli.Command{
      PAYMENT_STATUS_ACTION_QUERY.ToCli(),
      PAYMENT_STATUS_ACTION_TABLE.ToCli(),
      PaymentStatusCreateCmd,
      PaymentStatusUpdateCmd,
      PaymentStatusCreateInteractiveCmd,
      PaymentStatusWipeCmd,
      workspaces.GetCommonRemoveQuery(reflect.ValueOf(&PaymentStatusEntity{}).Elem(), PaymentStatusActionRemove),
  }
  func PaymentStatusCliFn() cli.Command {
    PaymentStatusCliCommands = append(PaymentStatusCliCommands, PaymentStatusImportExportCommands...)
    return cli.Command{
      Name:        "paymentStatus",
      Description: "PaymentStatuss module actions (sample module to handle complex entities)",
      Usage:       "Status of an payment",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: PaymentStatusCliCommands,
    }
  }
var PAYMENT_STATUS_ACTION_TABLE = workspaces.Module2Action{
  Name:    "table",
  ActionName: "table",
  ActionAliases: []string{"t"},
  Flags:  workspaces.CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: PaymentStatusActionQuery,
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    workspaces.CommonCliTableCmd2(c,
      PaymentStatusActionQuery,
      security,
      reflect.ValueOf(&PaymentStatusEntity{}).Elem(),
    )
    return nil
  },
}
var PAYMENT_STATUS_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/payment-statuses",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_STATUS_QUERY},
  },
  Group: "paymentStatus",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, PaymentStatusActionQuery)
    },
  },
  Format: "QUERY",
  Action: PaymentStatusActionQuery,
  ResponseEntity: &[]PaymentStatusEntity{},
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			PaymentStatusActionQuery,
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
var PAYMENT_STATUS_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/payment-statuses/export",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_STATUS_QUERY},
  },
  Group: "paymentStatus",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, PaymentStatusActionExport)
    },
  },
  Format: "QUERY",
  Action: PaymentStatusActionExport,
  ResponseEntity: &[]PaymentStatusEntity{},
}
var PAYMENT_STATUS_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/payment-status/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_STATUS_QUERY},
  },
  Group: "paymentStatus",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, PaymentStatusActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: PaymentStatusActionGetOne,
  ResponseEntity: &PaymentStatusEntity{},
}
var PAYMENT_STATUS_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new paymentStatus",
  Flags: PaymentStatusCommonCliFlags,
  Method: "POST",
  Url:    "/payment-status",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_STATUS_CREATE},
  },
  Group: "paymentStatus",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, PaymentStatusActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, PaymentStatusActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: PaymentStatusActionCreate,
  Format: "POST_ONE",
  RequestEntity: &PaymentStatusEntity{},
  ResponseEntity: &PaymentStatusEntity{},
}
var PAYMENT_STATUS_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: PaymentStatusCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/payment-status",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_STATUS_UPDATE},
  },
  Group: "paymentStatus",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, PaymentStatusActionUpdate)
    },
  },
  Action: PaymentStatusActionUpdate,
  RequestEntity: &PaymentStatusEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &PaymentStatusEntity{},
}
var PAYMENT_STATUS_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/payment-statuses",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_STATUS_UPDATE},
  },
  Group: "paymentStatus",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, PaymentStatusActionBulkUpdate)
    },
  },
  Action: PaymentStatusActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[PaymentStatusEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[PaymentStatusEntity]{},
}
var PAYMENT_STATUS_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/payment-status",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PAYMENT_STATUS_DELETE},
  },
  Group: "paymentStatus",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, PaymentStatusActionRemove)
    },
  },
  Action: PaymentStatusActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &PaymentStatusEntity{},
}
  /**
  *	Override this function on PaymentStatusEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendPaymentStatusRouter = func(r *[]workspaces.Module2Action) {}
  func GetPaymentStatusModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      PAYMENT_STATUS_ACTION_QUERY,
      PAYMENT_STATUS_ACTION_EXPORT,
      PAYMENT_STATUS_ACTION_GET_ONE,
      PAYMENT_STATUS_ACTION_POST_ONE,
      PAYMENT_STATUS_ACTION_PATCH,
      PAYMENT_STATUS_ACTION_PATCH_BULK,
      PAYMENT_STATUS_ACTION_DELETE,
    }
    // Append user defined functions
    AppendPaymentStatusRouter(&routes)
    return routes
  }
  func CreatePaymentStatusRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetPaymentStatusModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, PaymentStatusEntityJsonSchema, "payment-status-http", "shop")
    workspaces.WriteEntitySchema("PaymentStatusEntity", PaymentStatusEntityJsonSchema, "shop")
    return httpRoutes
  }
var PERM_ROOT_PAYMENT_STATUS_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/payment-status/delete",
  Name: "Delete payment status",
}
var PERM_ROOT_PAYMENT_STATUS_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/payment-status/create",
  Name: "Create payment status",
}
var PERM_ROOT_PAYMENT_STATUS_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/payment-status/update",
  Name: "Update payment status",
}
var PERM_ROOT_PAYMENT_STATUS_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/shop/payment-status/query",
  Name: "Query payment status",
}
var PERM_ROOT_PAYMENT_STATUS = workspaces.PermissionInfo{
  CompleteKey: "root/shop/payment-status/*",
  Name: "Entire payment status actions (*)",
}
var ALL_PAYMENT_STATUS_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_PAYMENT_STATUS_DELETE,
	PERM_ROOT_PAYMENT_STATUS_CREATE,
	PERM_ROOT_PAYMENT_STATUS_UPDATE,
	PERM_ROOT_PAYMENT_STATUS_QUERY,
	PERM_ROOT_PAYMENT_STATUS,
}