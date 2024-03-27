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
)
type ShoppingCartItems struct {
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
    Quantity   *float64 `json:"quantity" yaml:"quantity"       `
    // Datenano also has a text representation
    Product   *  ProductSubmissionEntity `json:"product" yaml:"product"    gorm:"foreignKey:ProductId;references:UniqueId"     `
    // Datenano also has a text representation
        ProductId *string `json:"productId" yaml:"productId"`
	LinkedTo *ShoppingCartEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
func ( x * ShoppingCartItems) RootObjectName() string {
	return "ShoppingCartEntity"
}
type ShoppingCartEntity struct {
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
    Items   []*  ShoppingCartItems `json:"items" yaml:"items"    gorm:"foreignKey:LinkerId;references:UniqueId"     `
    // Datenano also has a text representation
    Children []*ShoppingCartEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *ShoppingCartEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var ShoppingCartPreloadRelations []string = []string{}
var SHOPPING_CART_EVENT_CREATED = "shoppingCart.created"
var SHOPPING_CART_EVENT_UPDATED = "shoppingCart.updated"
var SHOPPING_CART_EVENT_DELETED = "shoppingCart.deleted"
var SHOPPING_CART_EVENTS = []string{
	SHOPPING_CART_EVENT_CREATED,
	SHOPPING_CART_EVENT_UPDATED,
	SHOPPING_CART_EVENT_DELETED,
}
type ShoppingCartFieldMap struct {
		Items workspaces.TranslatedString `yaml:"items"`
}
var ShoppingCartEntityMetaConfig map[string]int64 = map[string]int64{
}
var ShoppingCartEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&ShoppingCartEntity{}))
func ShoppingCartItemsActionCreate(
  dto *ShoppingCartItems ,
  query workspaces.QueryDSL,
) (*ShoppingCartItems , *workspaces.IError) {
    dto.LinkerId = &query.LinkerId
    var dbref *gorm.DB = nil
    if query.Tx == nil {
        dbref = workspaces.GetDbRef()
    } else {
        dbref = query.Tx
    }
    query.Tx = dbref
    if dto.UniqueId == "" {
        dto.UniqueId = workspaces.UUID()
    }
    err := dbref.Create(&dto).Error
    if err != nil {
        err := workspaces.GormErrorToIError(err)
        return dto, err
    }
    return dto, nil
}
func ShoppingCartItemsActionUpdate(
    query workspaces.QueryDSL,
    dto *ShoppingCartItems,
) (*ShoppingCartItems, *workspaces.IError) {
    dto.LinkerId = &query.LinkerId
    var dbref *gorm.DB = nil
    if query.Tx == nil {
        dbref = workspaces.GetDbRef()
    } else {
        dbref = query.Tx
    }
    query.Tx = dbref
    err := dbref.UpdateColumns(&dto).Error
    if err != nil {
        err := workspaces.GormErrorToIError(err)
        return dto, err
    }
    return dto, nil
}
func ShoppingCartItemsActionGetOne(
    query workspaces.QueryDSL,
) (*ShoppingCartItems , *workspaces.IError) {
    refl := reflect.ValueOf(&ShoppingCartItems {})
    item, err := workspaces.GetOneEntity[ShoppingCartItems ](query, refl)
    return item, err
}
func entityShoppingCartFormatter(dto *ShoppingCartEntity, query workspaces.QueryDSL) {
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
func ShoppingCartMockEntity() *ShoppingCartEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &ShoppingCartEntity{
	}
	return entity
}
func ShoppingCartActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := ShoppingCartMockEntity()
		_, err := ShoppingCartActionCreate(entity, query)
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
  func ShoppingCartActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*ShoppingCartEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &ShoppingCartEntity{
          Items: []*ShoppingCartItems{{}},
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
  func ShoppingCartAssociationCreate(dto *ShoppingCartEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func ShoppingCartRelationContentCreate(dto *ShoppingCartEntity, query workspaces.QueryDSL) error {
return nil
}
func ShoppingCartRelationContentUpdate(dto *ShoppingCartEntity, query workspaces.QueryDSL) error {
	return nil
}
func ShoppingCartPolyglotCreateHandler(dto *ShoppingCartEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func ShoppingCartValidator(dto *ShoppingCartEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
        if dto != nil && dto.Items != nil {
          workspaces.AppendSliceErrors(dto.Items, isPatch, "items", err)
        }
    return err
  }
func ShoppingCartEntityPreSanitize(dto *ShoppingCartEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func ShoppingCartEntityBeforeCreateAppend(dto *ShoppingCartEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    ShoppingCartRecursiveAddUniqueId(dto, query)
  }
  func ShoppingCartRecursiveAddUniqueId(dto *ShoppingCartEntity, query workspaces.QueryDSL) {
      if dto.Items != nil && len(dto.Items) > 0 {
        for index0 := range dto.Items {
          if (dto.Items[index0].UniqueId == "") {
            dto.Items[index0].UniqueId = workspaces.UUID()
          }
        }
    }
  }
func ShoppingCartActionBatchCreateFn(dtos []*ShoppingCartEntity, query workspaces.QueryDSL) ([]*ShoppingCartEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*ShoppingCartEntity{}
		for _, item := range dtos {
			s, err := ShoppingCartActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func ShoppingCartDeleteEntireChildren(query workspaces.QueryDSL, dto *ShoppingCartEntity) (*workspaces.IError) {
  if dto.Items != nil {
    q := query.Tx.
      Model(&dto.Items).
      Where(&ShoppingCartItems{LinkerId: &dto.UniqueId }).
      Delete(&ShoppingCartItems{})
    err := q.Error
    if err != nil {
      return workspaces.GormErrorToIError(err)
    }
  }
  return nil
}
func ShoppingCartActionCreateFn(dto *ShoppingCartEntity, query workspaces.QueryDSL) (*ShoppingCartEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := ShoppingCartValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	ShoppingCartEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	ShoppingCartEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	ShoppingCartPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	ShoppingCartRelationContentCreate(dto, query)
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
	ShoppingCartAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(SHOPPING_CART_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&ShoppingCartEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func ShoppingCartActionGetOne(query workspaces.QueryDSL) (*ShoppingCartEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&ShoppingCartEntity{})
    item, err := workspaces.GetOneEntity[ShoppingCartEntity](query, refl)
    entityShoppingCartFormatter(item, query)
    return item, err
  }
  func ShoppingCartActionQuery(query workspaces.QueryDSL) ([]*ShoppingCartEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&ShoppingCartEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[ShoppingCartEntity](query, refl)
    for _, item := range items {
      entityShoppingCartFormatter(item, query)
    }
    return items, meta, err
  }
  func ShoppingCartUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *ShoppingCartEntity) (*ShoppingCartEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = SHOPPING_CART_EVENT_UPDATED
    ShoppingCartEntityPreSanitize(fields, query)
    var item ShoppingCartEntity
    q := dbref.
      Where(&ShoppingCartEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    ShoppingCartRelationContentUpdate(fields, query)
    ShoppingCartPolyglotCreateHandler(fields, query)
    if ero := ShoppingCartDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
      if fields.Items != nil {
       linkerId := uniqueId;
        dbref.
          Where(&ShoppingCartItems {LinkerId: &linkerId}).
          Delete(&ShoppingCartItems {})
        for _, newItem := range fields.Items {
          newItem.UniqueId = workspaces.UUID()
          newItem.LinkerId = &linkerId
          dbref.Create(&newItem)
        }
      }
    err = dbref.
      Preload(clause.Associations).
      Where(&ShoppingCartEntity{UniqueId: uniqueId}).
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
  func ShoppingCartActionUpdateFn(query workspaces.QueryDSL, fields *ShoppingCartEntity) (*ShoppingCartEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := ShoppingCartValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // ShoppingCartRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *ShoppingCartEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = ShoppingCartUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return ShoppingCartUpdateExec(dbref, query, fields)
    }
  }
var ShoppingCartWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire shoppingcarts ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_SHOPPING_CART_DELETE},
    })
		count, _ := ShoppingCartActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func ShoppingCartActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&ShoppingCartEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_SHOPPING_CART_DELETE}
	return workspaces.RemoveEntity[ShoppingCartEntity](query, refl)
}
func ShoppingCartActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
			{
				subCount, subErr := workspaces.WipeCleanEntity[ShoppingCartItems]()
				if (subErr != nil) {
					fmt.Println("Error while wiping 'ShoppingCartItems'", subErr)
					return count, subErr
				} else {
					count += subCount
				}
			}
	{
		subCount, subErr := workspaces.WipeCleanEntity[ShoppingCartEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'ShoppingCartEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func ShoppingCartActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[ShoppingCartEntity]) (
    *workspaces.BulkRecordRequest[ShoppingCartEntity], *workspaces.IError,
  ) {
    result := []*ShoppingCartEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := ShoppingCartActionUpdate(query, record)
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
func (x *ShoppingCartEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var ShoppingCartEntityMeta = workspaces.TableMetaData{
	EntityName:    "ShoppingCart",
	ExportKey:    "shopping-carts",
	TableNameInDb: "fb_shopping-cart_entities",
	EntityObject:  &ShoppingCartEntity{},
	ExportStream: ShoppingCartActionExportT,
	ImportQuery: ShoppingCartActionImport,
}
func ShoppingCartActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[ShoppingCartEntity](query, ShoppingCartActionQuery, ShoppingCartPreloadRelations)
}
func ShoppingCartActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[ShoppingCartEntity](query, ShoppingCartActionQuery, ShoppingCartPreloadRelations)
}
func ShoppingCartActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content ShoppingCartEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := ShoppingCartActionCreate(&content, query)
	return err
}
var ShoppingCartCommonCliFlags = []cli.Flag{
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
    &cli.StringSliceFlag{
      Name:     "items",
      Required: false,
      Usage:    "items",
    },
}
var ShoppingCartCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
}
var ShoppingCartCommonCliFlagsOptional = []cli.Flag{
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
    &cli.StringSliceFlag{
      Name:     "items",
      Required: false,
      Usage:    "items",
    },
}
  var ShoppingCartCreateCmd cli.Command = SHOPPING_CART_ACTION_POST_ONE.ToCli()
  var ShoppingCartCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_SHOPPING_CART_CREATE},
      })
      entity := &ShoppingCartEntity{}
      for _, item := range ShoppingCartCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := ShoppingCartActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var ShoppingCartUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: ShoppingCartCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_SHOPPING_CART_UPDATE},
      })
      entity := CastShoppingCartFromCli(c)
      if entity, err := ShoppingCartActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* ShoppingCartEntity) FromCli(c *cli.Context) *ShoppingCartEntity {
	return CastShoppingCartFromCli(c)
}
func CastShoppingCartFromCli (c *cli.Context) *ShoppingCartEntity {
	template := &ShoppingCartEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
	return template
}
  func ShoppingCartSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      ShoppingCartActionCreate,
      reflect.ValueOf(&ShoppingCartEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func ShoppingCartWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := ShoppingCartActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "ShoppingCart", result)
    }
  }
var ShoppingCartImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_SHOPPING_CART_CREATE},
      })
			ShoppingCartActionSeeder(query, c.Int("count"))
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
				Value: "shopping-cart-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_SHOPPING_CART_CREATE},
      })
			ShoppingCartActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "shopping-cart-seeder-shopping-cart.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of shopping-carts, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]ShoppingCartEntity{}
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
			ShoppingCartCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				ShoppingCartActionCreate,
				reflect.ValueOf(&ShoppingCartEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_SHOPPING_CART_CREATE},
				},
        func() ShoppingCartEntity {
					v := CastShoppingCartFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var ShoppingCartCliCommands []cli.Command = []cli.Command{
      SHOPPING_CART_ACTION_QUERY.ToCli(),
      SHOPPING_CART_ACTION_TABLE.ToCli(),
      ShoppingCartCreateCmd,
      ShoppingCartUpdateCmd,
      ShoppingCartCreateInteractiveCmd,
      ShoppingCartWipeCmd,
      workspaces.GetCommonRemoveQuery(reflect.ValueOf(&ShoppingCartEntity{}).Elem(), ShoppingCartActionRemove),
  }
  func ShoppingCartCliFn() cli.Command {
    ShoppingCartCliCommands = append(ShoppingCartCliCommands, ShoppingCartImportExportCommands...)
    return cli.Command{
      Name:        "shoppingCart",
      Description: "ShoppingCarts module actions (sample module to handle complex entities)",
      Usage:       "Manage the active shopping carts (not ordered yet of the store)",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: ShoppingCartCliCommands,
    }
  }
var SHOPPING_CART_ACTION_TABLE = workspaces.Module2Action{
  Name:    "table",
  ActionAliases: []string{"t"},
  Flags:  workspaces.CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: ShoppingCartActionQuery,
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    workspaces.CommonCliTableCmd2(c,
      ShoppingCartActionQuery,
      security,
      reflect.ValueOf(&ShoppingCartEntity{}).Elem(),
    )
    return nil
  },
}
var SHOPPING_CART_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/shopping-carts",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_SHOPPING_CART_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, ShoppingCartActionQuery)
    },
  },
  Format: "QUERY",
  Action: ShoppingCartActionQuery,
  ResponseEntity: &[]ShoppingCartEntity{},
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			ShoppingCartActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionAliases: []string{"q"},
	Flags:         workspaces.CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var SHOPPING_CART_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/shopping-carts/export",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_SHOPPING_CART_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, ShoppingCartActionExport)
    },
  },
  Format: "QUERY",
  Action: ShoppingCartActionExport,
  ResponseEntity: &[]ShoppingCartEntity{},
}
var SHOPPING_CART_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/shopping-cart/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_SHOPPING_CART_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, ShoppingCartActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: ShoppingCartActionGetOne,
  ResponseEntity: &ShoppingCartEntity{},
}
var SHOPPING_CART_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new shoppingCart",
  Flags: ShoppingCartCommonCliFlags,
  Method: "POST",
  Url:    "/shopping-cart",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_SHOPPING_CART_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, ShoppingCartActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, ShoppingCartActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: ShoppingCartActionCreate,
  Format: "POST_ONE",
  RequestEntity: &ShoppingCartEntity{},
  ResponseEntity: &ShoppingCartEntity{},
}
var SHOPPING_CART_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: ShoppingCartCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/shopping-cart",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_SHOPPING_CART_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, ShoppingCartActionUpdate)
    },
  },
  Action: ShoppingCartActionUpdate,
  RequestEntity: &ShoppingCartEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &ShoppingCartEntity{},
}
var SHOPPING_CART_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/shopping-carts",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_SHOPPING_CART_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, ShoppingCartActionBulkUpdate)
    },
  },
  Action: ShoppingCartActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[ShoppingCartEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[ShoppingCartEntity]{},
}
var SHOPPING_CART_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/shopping-cart",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_SHOPPING_CART_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, ShoppingCartActionRemove)
    },
  },
  Action: ShoppingCartActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &ShoppingCartEntity{},
}
    var SHOPPING_CART_ITEMS_ACTION_PATCH = workspaces.Module2Action{
      Method: "PATCH",
      Url:    "/shopping-cart/:linkerId/items/:uniqueId",
      SecurityModel: &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_SHOPPING_CART_UPDATE},
      },
      Handlers: []gin.HandlerFunc{
        func (
          c *gin.Context,
        ) {
          workspaces.HttpUpdateEntity(c, ShoppingCartItemsActionUpdate)
        },
      },
      Action: ShoppingCartItemsActionUpdate,
      Format: "PATCH_ONE",
      RequestEntity: &ShoppingCartItems{},
      ResponseEntity: &ShoppingCartItems{},
    }
    var SHOPPING_CART_ITEMS_ACTION_GET = workspaces.Module2Action {
      Method: "GET",
      Url:    "/shopping-cart/items/:linkerId/:uniqueId",
      SecurityModel: &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_SHOPPING_CART_QUERY},
      },
      Handlers: []gin.HandlerFunc{
        func (
          c *gin.Context,
        ) {
          workspaces.HttpGetEntity(c, ShoppingCartItemsActionGetOne)
        },
      },
      Action: ShoppingCartItemsActionGetOne,
      Format: "GET_ONE",
      ResponseEntity: &ShoppingCartItems{},
    }
    var SHOPPING_CART_ITEMS_ACTION_POST = workspaces.Module2Action{
      Method: "POST",
      Url:    "/shopping-cart/:linkerId/items",
      SecurityModel: &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_SHOPPING_CART_CREATE},
      },
      Handlers: []gin.HandlerFunc{
        func (
          c *gin.Context,
        ) {
          workspaces.HttpPostEntity(c, ShoppingCartItemsActionCreate)
        },
      },
      Action: ShoppingCartItemsActionCreate,
      Format: "POST_ONE",
      RequestEntity: &ShoppingCartItems{},
      ResponseEntity: &ShoppingCartItems{},
    }
  /**
  *	Override this function on ShoppingCartEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendShoppingCartRouter = func(r *[]workspaces.Module2Action) {}
  func GetShoppingCartModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      SHOPPING_CART_ACTION_QUERY,
      SHOPPING_CART_ACTION_EXPORT,
      SHOPPING_CART_ACTION_GET_ONE,
      SHOPPING_CART_ACTION_POST_ONE,
      SHOPPING_CART_ACTION_PATCH,
      SHOPPING_CART_ACTION_PATCH_BULK,
      SHOPPING_CART_ACTION_DELETE,
          SHOPPING_CART_ITEMS_ACTION_PATCH,
          SHOPPING_CART_ITEMS_ACTION_GET,
          SHOPPING_CART_ITEMS_ACTION_POST,
    }
    // Append user defined functions
    AppendShoppingCartRouter(&routes)
    return routes
  }
  func CreateShoppingCartRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetShoppingCartModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, ShoppingCartEntityJsonSchema, "shopping-cart-http", "shop")
    workspaces.WriteEntitySchema("ShoppingCartEntity", ShoppingCartEntityJsonSchema, "shop")
    return httpRoutes
  }
var PERM_ROOT_SHOPPING_CART_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/shopping-cart/delete",
}
var PERM_ROOT_SHOPPING_CART_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/shopping-cart/create",
}
var PERM_ROOT_SHOPPING_CART_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/shopping-cart/update",
}
var PERM_ROOT_SHOPPING_CART_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/shop/shopping-cart/query",
}
var PERM_ROOT_SHOPPING_CART = workspaces.PermissionInfo{
  CompleteKey: "root/shop/shopping-cart/*",
}
var ALL_SHOPPING_CART_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_SHOPPING_CART_DELETE,
	PERM_ROOT_SHOPPING_CART_CREATE,
	PERM_ROOT_SHOPPING_CART_UPDATE,
	PERM_ROOT_SHOPPING_CART_QUERY,
	PERM_ROOT_SHOPPING_CART,
}