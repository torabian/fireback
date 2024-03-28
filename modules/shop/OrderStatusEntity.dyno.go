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
	seeders "github.com/torabian/fireback/modules/shop/seeders/OrderStatus"
	metas "github.com/torabian/fireback/modules/shop/metas"
)
type OrderStatusEntity struct {
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
    Translations     []*OrderStatusEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    Children []*OrderStatusEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *OrderStatusEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var OrderStatusPreloadRelations []string = []string{}
var ORDER_STATUS_EVENT_CREATED = "orderStatus.created"
var ORDER_STATUS_EVENT_UPDATED = "orderStatus.updated"
var ORDER_STATUS_EVENT_DELETED = "orderStatus.deleted"
var ORDER_STATUS_EVENTS = []string{
	ORDER_STATUS_EVENT_CREATED,
	ORDER_STATUS_EVENT_UPDATED,
	ORDER_STATUS_EVENT_DELETED,
}
type OrderStatusFieldMap struct {
		Name workspaces.TranslatedString `yaml:"name"`
		Description workspaces.TranslatedString `yaml:"description"`
}
var OrderStatusEntityMetaConfig map[string]int64 = map[string]int64{
}
var OrderStatusEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&OrderStatusEntity{}))
  type OrderStatusEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        Name string `yaml:"name" json:"name"`
        Description string `yaml:"description" json:"description"`
  }
func entityOrderStatusFormatter(dto *OrderStatusEntity, query workspaces.QueryDSL) {
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
func OrderStatusMockEntity() *OrderStatusEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &OrderStatusEntity{
      Name : &stringHolder,
      Description : &stringHolder,
	}
	return entity
}
func OrderStatusActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := OrderStatusMockEntity()
		_, err := OrderStatusActionCreate(entity, query)
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
    func (x*OrderStatusEntity) GetNameTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Name
          }
        }
      }
      return ""
    }
    func (x*OrderStatusEntity) GetDescriptionTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Description
          }
        }
      }
      return ""
    }
  func OrderStatusActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*OrderStatusEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &OrderStatusEntity{
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
  func OrderStatusAssociationCreate(dto *OrderStatusEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func OrderStatusRelationContentCreate(dto *OrderStatusEntity, query workspaces.QueryDSL) error {
return nil
}
func OrderStatusRelationContentUpdate(dto *OrderStatusEntity, query workspaces.QueryDSL) error {
	return nil
}
func OrderStatusPolyglotCreateHandler(dto *OrderStatusEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &OrderStatusEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func OrderStatusValidator(dto *OrderStatusEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func OrderStatusEntityPreSanitize(dto *OrderStatusEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func OrderStatusEntityBeforeCreateAppend(dto *OrderStatusEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    OrderStatusRecursiveAddUniqueId(dto, query)
  }
  func OrderStatusRecursiveAddUniqueId(dto *OrderStatusEntity, query workspaces.QueryDSL) {
  }
func OrderStatusActionBatchCreateFn(dtos []*OrderStatusEntity, query workspaces.QueryDSL) ([]*OrderStatusEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*OrderStatusEntity{}
		for _, item := range dtos {
			s, err := OrderStatusActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func OrderStatusDeleteEntireChildren(query workspaces.QueryDSL, dto *OrderStatusEntity) (*workspaces.IError) {
  return nil
}
func OrderStatusActionCreateFn(dto *OrderStatusEntity, query workspaces.QueryDSL) (*OrderStatusEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := OrderStatusValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	OrderStatusEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	OrderStatusEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	OrderStatusPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	OrderStatusRelationContentCreate(dto, query)
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
	OrderStatusAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(ORDER_STATUS_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&OrderStatusEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func OrderStatusActionGetOne(query workspaces.QueryDSL) (*OrderStatusEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&OrderStatusEntity{})
    item, err := workspaces.GetOneEntity[OrderStatusEntity](query, refl)
    entityOrderStatusFormatter(item, query)
    return item, err
  }
  func OrderStatusActionQuery(query workspaces.QueryDSL) ([]*OrderStatusEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&OrderStatusEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[OrderStatusEntity](query, refl)
    for _, item := range items {
      entityOrderStatusFormatter(item, query)
    }
    return items, meta, err
  }
  func OrderStatusUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *OrderStatusEntity) (*OrderStatusEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = ORDER_STATUS_EVENT_UPDATED
    OrderStatusEntityPreSanitize(fields, query)
    var item OrderStatusEntity
    q := dbref.
      Where(&OrderStatusEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    OrderStatusRelationContentUpdate(fields, query)
    OrderStatusPolyglotCreateHandler(fields, query)
    if ero := OrderStatusDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&OrderStatusEntity{UniqueId: uniqueId}).
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
  func OrderStatusActionUpdateFn(query workspaces.QueryDSL, fields *OrderStatusEntity) (*OrderStatusEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := OrderStatusValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // OrderStatusRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *OrderStatusEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = OrderStatusUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return OrderStatusUpdateExec(dbref, query, fields)
    }
  }
var OrderStatusWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire orderstatuses ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_STATUS_DELETE},
    })
		count, _ := OrderStatusActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func OrderStatusActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&OrderStatusEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_ORDER_STATUS_DELETE}
	return workspaces.RemoveEntity[OrderStatusEntity](query, refl)
}
func OrderStatusActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[OrderStatusEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'OrderStatusEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func OrderStatusActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[OrderStatusEntity]) (
    *workspaces.BulkRecordRequest[OrderStatusEntity], *workspaces.IError,
  ) {
    result := []*OrderStatusEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := OrderStatusActionUpdate(query, record)
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
func (x *OrderStatusEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var OrderStatusEntityMeta = workspaces.TableMetaData{
	EntityName:    "OrderStatus",
	ExportKey:    "order-statuses",
	TableNameInDb: "fb_order-status_entities",
	EntityObject:  &OrderStatusEntity{},
	ExportStream: OrderStatusActionExportT,
	ImportQuery: OrderStatusActionImport,
}
func OrderStatusActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[OrderStatusEntity](query, OrderStatusActionQuery, OrderStatusPreloadRelations)
}
func OrderStatusActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[OrderStatusEntity](query, OrderStatusActionQuery, OrderStatusPreloadRelations)
}
func OrderStatusActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content OrderStatusEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := OrderStatusActionCreate(&content, query)
	return err
}
var OrderStatusCommonCliFlags = []cli.Flag{
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
var OrderStatusCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
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
var OrderStatusCommonCliFlagsOptional = []cli.Flag{
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
  var OrderStatusCreateCmd cli.Command = ORDER_STATUS_ACTION_POST_ONE.ToCli()
  var OrderStatusCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_STATUS_CREATE},
      })
      entity := &OrderStatusEntity{}
      for _, item := range OrderStatusCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := OrderStatusActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var OrderStatusUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: OrderStatusCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_STATUS_UPDATE},
      })
      entity := CastOrderStatusFromCli(c)
      if entity, err := OrderStatusActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* OrderStatusEntity) FromCli(c *cli.Context) *OrderStatusEntity {
	return CastOrderStatusFromCli(c)
}
func CastOrderStatusFromCli (c *cli.Context) *OrderStatusEntity {
	template := &OrderStatusEntity{}
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
  func OrderStatusSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      OrderStatusActionCreate,
      reflect.ValueOf(&OrderStatusEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func OrderStatusSyncSeeders() {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{WorkspaceId: workspaces.USER_SYSTEM},
      OrderStatusActionCreate,
      reflect.ValueOf(&OrderStatusEntity{}).Elem(),
      &seeders.ViewsFs,
      []string{},
      true,
    )
  }
  func OrderStatusWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := OrderStatusActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "OrderStatus", result)
    }
  }
var OrderStatusImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_STATUS_CREATE},
      })
			OrderStatusActionSeeder(query, c.Int("count"))
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
				Value: "order-status-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_STATUS_CREATE},
      })
			OrderStatusActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "order-status-seeder-order-status.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of order-statuses, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]OrderStatusEntity{}
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
				OrderStatusActionCreate,
				reflect.ValueOf(&OrderStatusEntity{}).Elem(),
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
				OrderStatusActionQuery,
				reflect.ValueOf(&OrderStatusEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"OrderStatusFieldMap.yml",
				OrderStatusPreloadRelations,
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
			OrderStatusCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				OrderStatusActionCreate,
				reflect.ValueOf(&OrderStatusEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_STATUS_CREATE},
				},
        func() OrderStatusEntity {
					v := CastOrderStatusFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var OrderStatusCliCommands []cli.Command = []cli.Command{
      ORDER_STATUS_ACTION_QUERY.ToCli(),
      ORDER_STATUS_ACTION_TABLE.ToCli(),
      OrderStatusCreateCmd,
      OrderStatusUpdateCmd,
      OrderStatusCreateInteractiveCmd,
      OrderStatusWipeCmd,
      workspaces.GetCommonRemoveQuery(reflect.ValueOf(&OrderStatusEntity{}).Elem(), OrderStatusActionRemove),
  }
  func OrderStatusCliFn() cli.Command {
    OrderStatusCliCommands = append(OrderStatusCliCommands, OrderStatusImportExportCommands...)
    return cli.Command{
      Name:        "orderStatus",
      Description: "OrderStatuss module actions (sample module to handle complex entities)",
      Usage:       "Status of an order",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: OrderStatusCliCommands,
    }
  }
var ORDER_STATUS_ACTION_TABLE = workspaces.Module2Action{
  Name:    "table",
  ActionAliases: []string{"t"},
  Flags:  workspaces.CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: OrderStatusActionQuery,
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    workspaces.CommonCliTableCmd2(c,
      OrderStatusActionQuery,
      security,
      reflect.ValueOf(&OrderStatusEntity{}).Elem(),
    )
    return nil
  },
}
var ORDER_STATUS_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/order-statuses",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_STATUS_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, OrderStatusActionQuery)
    },
  },
  Format: "QUERY",
  Action: OrderStatusActionQuery,
  ResponseEntity: &[]OrderStatusEntity{},
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			OrderStatusActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionAliases: []string{"q"},
	Flags:         workspaces.CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var ORDER_STATUS_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/order-statuses/export",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_STATUS_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, OrderStatusActionExport)
    },
  },
  Format: "QUERY",
  Action: OrderStatusActionExport,
  ResponseEntity: &[]OrderStatusEntity{},
}
var ORDER_STATUS_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/order-status/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_STATUS_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, OrderStatusActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: OrderStatusActionGetOne,
  ResponseEntity: &OrderStatusEntity{},
}
var ORDER_STATUS_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new orderStatus",
  Flags: OrderStatusCommonCliFlags,
  Method: "POST",
  Url:    "/order-status",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_STATUS_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, OrderStatusActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, OrderStatusActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: OrderStatusActionCreate,
  Format: "POST_ONE",
  RequestEntity: &OrderStatusEntity{},
  ResponseEntity: &OrderStatusEntity{},
}
var ORDER_STATUS_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: OrderStatusCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/order-status",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_STATUS_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, OrderStatusActionUpdate)
    },
  },
  Action: OrderStatusActionUpdate,
  RequestEntity: &OrderStatusEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &OrderStatusEntity{},
}
var ORDER_STATUS_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/order-statuses",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_STATUS_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, OrderStatusActionBulkUpdate)
    },
  },
  Action: OrderStatusActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[OrderStatusEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[OrderStatusEntity]{},
}
var ORDER_STATUS_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/order-status",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_STATUS_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, OrderStatusActionRemove)
    },
  },
  Action: OrderStatusActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &OrderStatusEntity{},
}
  /**
  *	Override this function on OrderStatusEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendOrderStatusRouter = func(r *[]workspaces.Module2Action) {}
  func GetOrderStatusModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      ORDER_STATUS_ACTION_QUERY,
      ORDER_STATUS_ACTION_EXPORT,
      ORDER_STATUS_ACTION_GET_ONE,
      ORDER_STATUS_ACTION_POST_ONE,
      ORDER_STATUS_ACTION_PATCH,
      ORDER_STATUS_ACTION_PATCH_BULK,
      ORDER_STATUS_ACTION_DELETE,
    }
    // Append user defined functions
    AppendOrderStatusRouter(&routes)
    return routes
  }
  func CreateOrderStatusRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetOrderStatusModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, OrderStatusEntityJsonSchema, "order-status-http", "shop")
    workspaces.WriteEntitySchema("OrderStatusEntity", OrderStatusEntityJsonSchema, "shop")
    return httpRoutes
  }
var PERM_ROOT_ORDER_STATUS_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/order-status/delete",
  Name: "Delete order status",
}
var PERM_ROOT_ORDER_STATUS_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/order-status/create",
  Name: "Create order status",
}
var PERM_ROOT_ORDER_STATUS_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/order-status/update",
  Name: "Update order status",
}
var PERM_ROOT_ORDER_STATUS_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/shop/order-status/query",
  Name: "Query order status",
}
var PERM_ROOT_ORDER_STATUS = workspaces.PermissionInfo{
  CompleteKey: "root/shop/order-status/*",
  Name: "Entire order status actions (*)",
}
var ALL_ORDER_STATUS_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_ORDER_STATUS_DELETE,
	PERM_ROOT_ORDER_STATUS_CREATE,
	PERM_ROOT_ORDER_STATUS_UPDATE,
	PERM_ROOT_ORDER_STATUS_QUERY,
	PERM_ROOT_ORDER_STATUS,
}