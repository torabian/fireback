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
import  "github.com/torabian/fireback/modules/currency"
type OrderTotalPrice struct {
    Visibility       *string                         `json:"visibility,omitempty" yaml:"visibility"`
    WorkspaceId      *string                         `json:"workspaceId,omitempty" yaml:"workspaceId"`
    LinkerId         *string                         `json:"linkerId,omitempty" yaml:"linkerId"`
    ParentId         *string                         `json:"parentId,omitempty" yaml:"parentId"`
    UniqueId         string                          `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
    UserId           *string                         `json:"userId,omitempty" yaml:"userId"`
    Rank             int64                           `json:"rank,omitempty" gorm:"type:int;name:rank"`
    Updated          int64                           `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
    Created          int64                           `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
    CreatedFormatted string                          `json:"createdFormatted,omitempty" sql:"-" gorm:"-"`
    UpdatedFormatted string                          `json:"updatedFormatted,omitempty" sql:"-" gorm:"-"`
    Amount   *float64 `json:"amount" yaml:"amount"  validate:"required"       `
    // Datenano also has a text representation
    Currency   *  currency.CurrencyEntity `json:"currency" yaml:"currency"    gorm:"foreignKey:CurrencyId;references:UniqueId"     `
    // Datenano also has a text representation
        CurrencyId *string `json:"currencyId" yaml:"currencyId"`
	LinkedTo *OrderEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
func ( x * OrderTotalPrice) RootObjectName() string {
	return "OrderEntity"
}
type OrderItems struct {
    Visibility       *string                         `json:"visibility,omitempty" yaml:"visibility"`
    WorkspaceId      *string                         `json:"workspaceId,omitempty" yaml:"workspaceId"`
    LinkerId         *string                         `json:"linkerId,omitempty" yaml:"linkerId"`
    ParentId         *string                         `json:"parentId,omitempty" yaml:"parentId"`
    UniqueId         string                          `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
    UserId           *string                         `json:"userId,omitempty" yaml:"userId"`
    Rank             int64                           `json:"rank,omitempty" gorm:"type:int;name:rank"`
    Updated          int64                           `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
    Created          int64                           `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
    CreatedFormatted string                          `json:"createdFormatted,omitempty" sql:"-" gorm:"-"`
    UpdatedFormatted string                          `json:"updatedFormatted,omitempty" sql:"-" gorm:"-"`
    Quantity   *float64 `json:"quantity" yaml:"quantity"       `
    // Datenano also has a text representation
    Price   *float64 `json:"price" yaml:"price"       `
    // Datenano also has a text representation
    Product   *  ProductSubmissionEntity `json:"product" yaml:"product"    gorm:"foreignKey:ProductId;references:UniqueId"     `
    // Datenano also has a text representation
        ProductId *string `json:"productId" yaml:"productId"`
    ProductSnapshot  *workspaces.   JSON `json:"productSnapshot" yaml:"productSnapshot"       `
    // Datenano also has a text representation
	LinkedTo *OrderEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
func ( x * OrderItems) RootObjectName() string {
	return "OrderEntity"
}
type OrderEntity struct {
    Visibility       *string                         `json:"visibility,omitempty" yaml:"visibility"`
    WorkspaceId      *string                         `json:"workspaceId,omitempty" yaml:"workspaceId"`
    LinkerId         *string                         `json:"linkerId,omitempty" yaml:"linkerId"`
    ParentId         *string                         `json:"parentId,omitempty" yaml:"parentId"`
    UniqueId         string                          `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
    UserId           *string                         `json:"userId,omitempty" yaml:"userId"`
    Rank             int64                           `json:"rank,omitempty" gorm:"type:int;name:rank"`
    Updated          int64                           `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
    Created          int64                           `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
    CreatedFormatted string                          `json:"createdFormatted,omitempty" sql:"-" gorm:"-"`
    UpdatedFormatted string                          `json:"updatedFormatted,omitempty" sql:"-" gorm:"-"`
    TotalPrice   *  OrderTotalPrice `json:"totalPrice" yaml:"totalPrice"    gorm:"foreignKey:LinkerId;references:UniqueId"     `
    // Datenano also has a text representation
    ShippingAddress   *string `json:"shippingAddress" yaml:"shippingAddress"       `
    // Datenano also has a text representation
    PaymentStatus   *  PaymentStatusEntity `json:"paymentStatus" yaml:"paymentStatus"    gorm:"foreignKey:PaymentStatusId;references:UniqueId"     `
    // Datenano also has a text representation
        PaymentStatusId *string `json:"paymentStatusId" yaml:"paymentStatusId" validate:"required" `
    OrderStatus   *  OrderStatusEntity `json:"orderStatus" yaml:"orderStatus"    gorm:"foreignKey:OrderStatusId;references:UniqueId"     `
    // Datenano also has a text representation
        OrderStatusId *string `json:"orderStatusId" yaml:"orderStatusId" validate:"required" `
    InvoiceNumber   *string `json:"invoiceNumber" yaml:"invoiceNumber"       `
    // Datenano also has a text representation
    DiscountCode   *  DiscountCodeEntity `json:"discountCode" yaml:"discountCode"    gorm:"foreignKey:DiscountCodeId;references:UniqueId"     `
    // Datenano also has a text representation
        DiscountCodeId *string `json:"discountCodeId" yaml:"discountCodeId"`
    Items   []*  OrderItems `json:"items" yaml:"items"    gorm:"foreignKey:LinkerId;references:UniqueId"     `
    // Datenano also has a text representation
    Children []*OrderEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *OrderEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var OrderPreloadRelations []string = []string{}
var ORDER_EVENT_CREATED = "order.created"
var ORDER_EVENT_UPDATED = "order.updated"
var ORDER_EVENT_DELETED = "order.deleted"
var ORDER_EVENTS = []string{
	ORDER_EVENT_CREATED,
	ORDER_EVENT_UPDATED,
	ORDER_EVENT_DELETED,
}
type OrderFieldMap struct {
		TotalPrice workspaces.TranslatedString `yaml:"totalPrice"`
		ShippingAddress workspaces.TranslatedString `yaml:"shippingAddress"`
		PaymentStatus workspaces.TranslatedString `yaml:"paymentStatus"`
		OrderStatus workspaces.TranslatedString `yaml:"orderStatus"`
		InvoiceNumber workspaces.TranslatedString `yaml:"invoiceNumber"`
		DiscountCode workspaces.TranslatedString `yaml:"discountCode"`
		Items workspaces.TranslatedString `yaml:"items"`
}
var OrderEntityMetaConfig map[string]int64 = map[string]int64{
}
var OrderEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&OrderEntity{}))
func OrderTotalPriceActionCreate(
  dto *OrderTotalPrice ,
  query workspaces.QueryDSL,
) (*OrderTotalPrice , *workspaces.IError) {
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
func OrderTotalPriceActionUpdate(
    query workspaces.QueryDSL,
    dto *OrderTotalPrice,
) (*OrderTotalPrice, *workspaces.IError) {
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
func OrderTotalPriceActionGetOne(
    query workspaces.QueryDSL,
) (*OrderTotalPrice , *workspaces.IError) {
    refl := reflect.ValueOf(&OrderTotalPrice {})
    item, err := workspaces.GetOneEntity[OrderTotalPrice ](query, refl)
    return item, err
}
func OrderItemsActionCreate(
  dto *OrderItems ,
  query workspaces.QueryDSL,
) (*OrderItems , *workspaces.IError) {
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
func OrderItemsActionUpdate(
    query workspaces.QueryDSL,
    dto *OrderItems,
) (*OrderItems, *workspaces.IError) {
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
func OrderItemsActionGetOne(
    query workspaces.QueryDSL,
) (*OrderItems , *workspaces.IError) {
    refl := reflect.ValueOf(&OrderItems {})
    item, err := workspaces.GetOneEntity[OrderItems ](query, refl)
    return item, err
}
func entityOrderFormatter(dto *OrderEntity, query workspaces.QueryDSL) {
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
func OrderMockEntity() *OrderEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &OrderEntity{
      ShippingAddress : &stringHolder,
      InvoiceNumber : &stringHolder,
	}
	return entity
}
func OrderActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := OrderMockEntity()
		_, err := OrderActionCreate(entity, query)
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
  func OrderActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*OrderEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &OrderEntity{
          TotalPrice: &OrderTotalPrice{},
          ShippingAddress: &tildaRef,
          InvoiceNumber: &tildaRef,
          Items: []*OrderItems{{}},
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
  func OrderAssociationCreate(dto *OrderEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func OrderRelationContentCreate(dto *OrderEntity, query workspaces.QueryDSL) error {
return nil
}
func OrderRelationContentUpdate(dto *OrderEntity, query workspaces.QueryDSL) error {
	return nil
}
func OrderPolyglotCreateHandler(dto *OrderEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func OrderValidator(dto *OrderEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
        if dto != nil && dto.Items != nil {
          workspaces.AppendSliceErrors(dto.Items, isPatch, "items", err)
        }
    return err
  }
func OrderEntityPreSanitize(dto *OrderEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func OrderEntityBeforeCreateAppend(dto *OrderEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    OrderRecursiveAddUniqueId(dto, query)
  }
  func OrderRecursiveAddUniqueId(dto *OrderEntity, query workspaces.QueryDSL) {
        if dto.TotalPrice != nil {
          dto.TotalPrice.UniqueId = workspaces.UUID()
        }
      if dto.Items != nil && len(dto.Items) > 0 {
        for index0 := range dto.Items {
          if (dto.Items[index0].UniqueId == "") {
            dto.Items[index0].UniqueId = workspaces.UUID()
          }
        }
    }
  }
func OrderActionBatchCreateFn(dtos []*OrderEntity, query workspaces.QueryDSL) ([]*OrderEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*OrderEntity{}
		for _, item := range dtos {
			s, err := OrderActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func OrderDeleteEntireChildren(query workspaces.QueryDSL, dto *OrderEntity) (*workspaces.IError) {
  if dto.TotalPrice != nil {
    q := query.Tx.
      Model(&dto.TotalPrice).
      Where(&OrderTotalPrice{LinkerId: &dto.UniqueId }).
      Delete(&OrderTotalPrice{})
    err := q.Error
    if err != nil {
      return workspaces.GormErrorToIError(err)
    }
  }
  if dto.Items != nil {
    q := query.Tx.
      Model(&dto.Items).
      Where(&OrderItems{LinkerId: &dto.UniqueId }).
      Delete(&OrderItems{})
    err := q.Error
    if err != nil {
      return workspaces.GormErrorToIError(err)
    }
  }
  return nil
}
func OrderActionCreateFn(dto *OrderEntity, query workspaces.QueryDSL) (*OrderEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := OrderValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	OrderEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	OrderEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	OrderPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	OrderRelationContentCreate(dto, query)
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
	OrderAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(ORDER_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&OrderEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func OrderActionGetOne(query workspaces.QueryDSL) (*OrderEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&OrderEntity{})
    item, err := workspaces.GetOneEntity[OrderEntity](query, refl)
    entityOrderFormatter(item, query)
    return item, err
  }
  func OrderActionQuery(query workspaces.QueryDSL) ([]*OrderEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&OrderEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[OrderEntity](query, refl)
    for _, item := range items {
      entityOrderFormatter(item, query)
    }
    return items, meta, err
  }
  func OrderUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *OrderEntity) (*OrderEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = ORDER_EVENT_UPDATED
    OrderEntityPreSanitize(fields, query)
    var item OrderEntity
    q := dbref.
      Where(&OrderEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    OrderRelationContentUpdate(fields, query)
    OrderPolyglotCreateHandler(fields, query)
    if ero := OrderDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
        if fields.TotalPrice != nil {
          linkerId := uniqueId
          q := dbref.
            Model(&item.TotalPrice).
            Where(&OrderTotalPrice{LinkerId: &linkerId}).
            UpdateColumns(fields.TotalPrice)
          err := q.Error
          if err != nil {
            return &item, workspaces.GormErrorToIError(err)
          }
          if q.RowsAffected == 0 {
            fields.TotalPrice.UniqueId = workspaces.UUID()
            fields.TotalPrice.LinkerId = &linkerId
            err := dbref.
              Model(&item.TotalPrice).Create(fields.TotalPrice).Error
            if err != nil {
              return &item, workspaces.GormErrorToIError(err)
            }
          }
        }
    // @meta(update has many)
      if fields.Items != nil {
       linkerId := uniqueId;
        dbref.
          Where(&OrderItems {LinkerId: &linkerId}).
          Delete(&OrderItems {})
        for _, newItem := range fields.Items {
          newItem.UniqueId = workspaces.UUID()
          newItem.LinkerId = &linkerId
          dbref.Create(&newItem)
        }
      }
    err = dbref.
      Preload(clause.Associations).
      Where(&OrderEntity{UniqueId: uniqueId}).
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
  func OrderActionUpdateFn(query workspaces.QueryDSL, fields *OrderEntity) (*OrderEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := OrderValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // OrderRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *OrderEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = OrderUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return OrderUpdateExec(dbref, query, fields)
    }
  }
var OrderWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire orders ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_DELETE},
    })
		count, _ := OrderActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func OrderActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&OrderEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_ORDER_DELETE}
	return workspaces.RemoveEntity[OrderEntity](query, refl)
}
func OrderActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
			{
				subCount, subErr := workspaces.WipeCleanEntity[OrderTotalPrice]()
				if (subErr != nil) {
					fmt.Println("Error while wiping 'OrderTotalPrice'", subErr)
					return count, subErr
				} else {
					count += subCount
				}
			}
			{
				subCount, subErr := workspaces.WipeCleanEntity[OrderItems]()
				if (subErr != nil) {
					fmt.Println("Error while wiping 'OrderItems'", subErr)
					return count, subErr
				} else {
					count += subCount
				}
			}
	{
		subCount, subErr := workspaces.WipeCleanEntity[OrderEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'OrderEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func OrderActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[OrderEntity]) (
    *workspaces.BulkRecordRequest[OrderEntity], *workspaces.IError,
  ) {
    result := []*OrderEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := OrderActionUpdate(query, record)
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
func (x *OrderEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var OrderEntityMeta = workspaces.TableMetaData{
	EntityName:    "Order",
	ExportKey:    "orders",
	TableNameInDb: "fb_order_entities",
	EntityObject:  &OrderEntity{},
	ExportStream: OrderActionExportT,
	ImportQuery: OrderActionImport,
}
func OrderActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[OrderEntity](query, OrderActionQuery, OrderPreloadRelations)
}
func OrderActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[OrderEntity](query, OrderActionQuery, OrderPreloadRelations)
}
func OrderActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content OrderEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := OrderActionCreate(&content, query)
	return err
}
var OrderCommonCliFlags = []cli.Flag{
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
    &cli.Float64Flag{
      Name:     "amount",
      Required: true,
      Usage:    "amount",
    },
    &cli.StringFlag{
      Name:     "currency-id",
      Required: false,
      Usage:    "currency",
    },
    &cli.StringFlag{
      Name:     "shipping-address",
      Required: false,
      Usage:    "Final computed shipping address which will be print on the product",
    },
    &cli.StringFlag{
      Name:     "payment-status-id",
      Required: true,
      Usage:    "paymentStatus",
    },
    &cli.StringFlag{
      Name:     "order-status-id",
      Required: true,
      Usage:    "orderStatus",
    },
    &cli.StringFlag{
      Name:     "invoice-number",
      Required: false,
      Usage:    "invoiceNumber",
    },
    &cli.StringFlag{
      Name:     "discount-code-id",
      Required: false,
      Usage:    "discountCode",
    },
    &cli.StringSliceFlag{
      Name:     "items",
      Required: false,
      Usage:    "items",
    },
}
var OrderCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "shippingAddress",
		StructField:     "ShippingAddress",
		Required: false,
		Usage:    "Final computed shipping address which will be print on the product",
		Type: "string",
	},
	{
		Name:     "invoiceNumber",
		StructField:     "InvoiceNumber",
		Required: false,
		Usage:    "invoiceNumber",
		Type: "string",
	},
}
var OrderCommonCliFlagsOptional = []cli.Flag{
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
    &cli.Float64Flag{
      Name:     "amount",
      Required: true,
      Usage:    "amount",
    },
    &cli.StringFlag{
      Name:     "currency-id",
      Required: false,
      Usage:    "currency",
    },
    &cli.StringFlag{
      Name:     "shipping-address",
      Required: false,
      Usage:    "Final computed shipping address which will be print on the product",
    },
    &cli.StringFlag{
      Name:     "payment-status-id",
      Required: true,
      Usage:    "paymentStatus",
    },
    &cli.StringFlag{
      Name:     "order-status-id",
      Required: true,
      Usage:    "orderStatus",
    },
    &cli.StringFlag{
      Name:     "invoice-number",
      Required: false,
      Usage:    "invoiceNumber",
    },
    &cli.StringFlag{
      Name:     "discount-code-id",
      Required: false,
      Usage:    "discountCode",
    },
    &cli.StringSliceFlag{
      Name:     "items",
      Required: false,
      Usage:    "items",
    },
}
  var OrderCreateCmd cli.Command = ORDER_ACTION_POST_ONE.ToCli()
  var OrderCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_CREATE},
      })
      entity := &OrderEntity{}
      for _, item := range OrderCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := OrderActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var OrderUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: OrderCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_UPDATE},
      })
      entity := CastOrderFromCli(c)
      if entity, err := OrderActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* OrderEntity) FromCli(c *cli.Context) *OrderEntity {
	return CastOrderFromCli(c)
}
func CastOrderFromCli (c *cli.Context) *OrderEntity {
	template := &OrderEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("shipping-address") {
        value := c.String("shipping-address")
        template.ShippingAddress = &value
      }
      if c.IsSet("payment-status-id") {
        value := c.String("payment-status-id")
        template.PaymentStatusId = &value
      }
      if c.IsSet("order-status-id") {
        value := c.String("order-status-id")
        template.OrderStatusId = &value
      }
      if c.IsSet("invoice-number") {
        value := c.String("invoice-number")
        template.InvoiceNumber = &value
      }
      if c.IsSet("discount-code-id") {
        value := c.String("discount-code-id")
        template.DiscountCodeId = &value
      }
	return template
}
  func OrderSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      OrderActionCreate,
      reflect.ValueOf(&OrderEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func OrderWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := OrderActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "Order", result)
    }
  }
var OrderImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_CREATE},
      })
			OrderActionSeeder(query, c.Int("count"))
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
				Value: "order-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_CREATE},
      })
			OrderActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "order-seeder-order.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of orders, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]OrderEntity{}
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
			OrderCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				OrderActionCreate,
				reflect.ValueOf(&OrderEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_CREATE},
				},
        func() OrderEntity {
					v := CastOrderFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var OrderCliCommands []cli.Command = []cli.Command{
      workspaces.GetCommonQuery2(OrderActionQuery, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_QUERY},
      }),
      workspaces.GetCommonTableQuery(reflect.ValueOf(&OrderEntity{}).Elem(), OrderActionQuery),
          OrderCreateCmd,
          OrderUpdateCmd,
          OrderCreateInteractiveCmd,
          OrderWipeCmd,
          workspaces.GetCommonRemoveQuery(reflect.ValueOf(&OrderEntity{}).Elem(), OrderActionRemove),
  }
  func OrderCliFn() cli.Command {
    OrderCliCommands = append(OrderCliCommands, OrderImportExportCommands...)
    return cli.Command{
      Name:        "order",
      Description: "Orders module actions (sample module to handle complex entities)",
      Usage:       "Placed orders by users, history, and their status",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: OrderCliCommands,
    }
  }
var ORDER_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/orders",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, OrderActionQuery)
    },
  },
  Format: "QUERY",
  Action: OrderActionQuery,
  ResponseEntity: &[]OrderEntity{},
}
var ORDER_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/orders/export",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, OrderActionExport)
    },
  },
  Format: "QUERY",
  Action: OrderActionExport,
  ResponseEntity: &[]OrderEntity{},
}
var ORDER_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/order/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, OrderActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: OrderActionGetOne,
  ResponseEntity: &OrderEntity{},
}
var ORDER_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new order",
  Flags: OrderCommonCliFlags,
  Method: "POST",
  Url:    "/order",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, OrderActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, OrderActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: OrderActionCreate,
  Format: "POST_ONE",
  RequestEntity: &OrderEntity{},
  ResponseEntity: &OrderEntity{},
}
var ORDER_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: OrderCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/order",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, OrderActionUpdate)
    },
  },
  Action: OrderActionUpdate,
  RequestEntity: &OrderEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &OrderEntity{},
}
var ORDER_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/orders",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, OrderActionBulkUpdate)
    },
  },
  Action: OrderActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[OrderEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[OrderEntity]{},
}
var ORDER_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/order",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, OrderActionRemove)
    },
  },
  Action: OrderActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &OrderEntity{},
}
    var ORDER_TOTAL_PRICE_ACTION_PATCH = workspaces.Module2Action{
      Method: "PATCH",
      Url:    "/order/:linkerId/total_price/:uniqueId",
      SecurityModel: &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_UPDATE},
      },
      Handlers: []gin.HandlerFunc{
        func (
          c *gin.Context,
        ) {
          workspaces.HttpUpdateEntity(c, OrderTotalPriceActionUpdate)
        },
      },
      Action: OrderTotalPriceActionUpdate,
      Format: "PATCH_ONE",
      RequestEntity: &OrderTotalPrice{},
      ResponseEntity: &OrderTotalPrice{},
    }
    var ORDER_TOTAL_PRICE_ACTION_GET = workspaces.Module2Action {
      Method: "GET",
      Url:    "/order/total_price/:linkerId/:uniqueId",
      SecurityModel: &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_QUERY},
      },
      Handlers: []gin.HandlerFunc{
        func (
          c *gin.Context,
        ) {
          workspaces.HttpGetEntity(c, OrderTotalPriceActionGetOne)
        },
      },
      Action: OrderTotalPriceActionGetOne,
      Format: "GET_ONE",
      ResponseEntity: &OrderTotalPrice{},
    }
    var ORDER_TOTAL_PRICE_ACTION_POST = workspaces.Module2Action{
      Method: "POST",
      Url:    "/order/:linkerId/total_price",
      SecurityModel: &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_CREATE},
      },
      Handlers: []gin.HandlerFunc{
        func (
          c *gin.Context,
        ) {
          workspaces.HttpPostEntity(c, OrderTotalPriceActionCreate)
        },
      },
      Action: OrderTotalPriceActionCreate,
      Format: "POST_ONE",
      RequestEntity: &OrderTotalPrice{},
      ResponseEntity: &OrderTotalPrice{},
    }
    var ORDER_ITEMS_ACTION_PATCH = workspaces.Module2Action{
      Method: "PATCH",
      Url:    "/order/:linkerId/items/:uniqueId",
      SecurityModel: &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_UPDATE},
      },
      Handlers: []gin.HandlerFunc{
        func (
          c *gin.Context,
        ) {
          workspaces.HttpUpdateEntity(c, OrderItemsActionUpdate)
        },
      },
      Action: OrderItemsActionUpdate,
      Format: "PATCH_ONE",
      RequestEntity: &OrderItems{},
      ResponseEntity: &OrderItems{},
    }
    var ORDER_ITEMS_ACTION_GET = workspaces.Module2Action {
      Method: "GET",
      Url:    "/order/items/:linkerId/:uniqueId",
      SecurityModel: &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_QUERY},
      },
      Handlers: []gin.HandlerFunc{
        func (
          c *gin.Context,
        ) {
          workspaces.HttpGetEntity(c, OrderItemsActionGetOne)
        },
      },
      Action: OrderItemsActionGetOne,
      Format: "GET_ONE",
      ResponseEntity: &OrderItems{},
    }
    var ORDER_ITEMS_ACTION_POST = workspaces.Module2Action{
      Method: "POST",
      Url:    "/order/:linkerId/items",
      SecurityModel: &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_ORDER_CREATE},
      },
      Handlers: []gin.HandlerFunc{
        func (
          c *gin.Context,
        ) {
          workspaces.HttpPostEntity(c, OrderItemsActionCreate)
        },
      },
      Action: OrderItemsActionCreate,
      Format: "POST_ONE",
      RequestEntity: &OrderItems{},
      ResponseEntity: &OrderItems{},
    }
  /**
  *	Override this function on OrderEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendOrderRouter = func(r *[]workspaces.Module2Action) {}
  func GetOrderModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      ORDER_ACTION_QUERY,
      ORDER_ACTION_EXPORT,
      ORDER_ACTION_GET_ONE,
      ORDER_ACTION_POST_ONE,
      ORDER_ACTION_PATCH,
      ORDER_ACTION_PATCH_BULK,
      ORDER_ACTION_DELETE,
          ORDER_TOTAL_PRICE_ACTION_PATCH,
          ORDER_TOTAL_PRICE_ACTION_GET,
          ORDER_TOTAL_PRICE_ACTION_POST,
          ORDER_ITEMS_ACTION_PATCH,
          ORDER_ITEMS_ACTION_GET,
          ORDER_ITEMS_ACTION_POST,
    }
    // Append user defined functions
    AppendOrderRouter(&routes)
    return routes
  }
  func CreateOrderRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetOrderModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, OrderEntityJsonSchema, "order-http", "shop")
    workspaces.WriteEntitySchema("OrderEntity", OrderEntityJsonSchema, "shop")
    return httpRoutes
  }
var PERM_ROOT_ORDER_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/order/delete",
}
var PERM_ROOT_ORDER_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/order/create",
}
var PERM_ROOT_ORDER_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/order/update",
}
var PERM_ROOT_ORDER_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/shop/order/query",
}
var PERM_ROOT_ORDER = workspaces.PermissionInfo{
  CompleteKey: "root/shop/order/*",
}
var PERM_ROOT_ORDER_CONFIRM = workspaces.PermissionInfo{
  CompleteKey: "root/shop/order/confirm",
}
var ALL_ORDER_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_ORDER_DELETE,
	PERM_ROOT_ORDER_CREATE,
	PERM_ROOT_ORDER_UPDATE,
	PERM_ROOT_ORDER_QUERY,
	PERM_ROOT_ORDER,
  PERM_ROOT_ORDER_CONFIRM,
}