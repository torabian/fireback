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
	mocks "github.com/torabian/fireback/modules/shop/mocks/ProductSubmission"
)
import  "github.com/torabian/fireback/modules/currency"
var productSubmissionSeedersFs *embed.FS = nil
func ResetProductSubmissionSeeders(fs *embed.FS) {
	productSubmissionSeedersFs = fs
}
type ProductSubmissionValues struct {
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
    ProductField   *  ProductFields `json:"productField" yaml:"productField"    gorm:"foreignKey:ProductFieldId;references:UniqueId"     `
    // Datenano also has a text representation
        ProductFieldId *string `json:"productFieldId" yaml:"productFieldId"`
    ValueInt64   *int64 `json:"valueInt64" yaml:"valueInt64"       `
    // Datenano also has a text representation
    ValueFloat64   *float64 `json:"valueFloat64" yaml:"valueFloat64"       `
    // Datenano also has a text representation
    ValueString   *string `json:"valueString" yaml:"valueString"       `
    // Datenano also has a text representation
    ValueBoolean   *bool `json:"valueBoolean" yaml:"valueBoolean"       `
    // Datenano also has a text representation
	LinkedTo *ProductSubmissionEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
func ( x * ProductSubmissionValues) RootObjectName() string {
	return "ProductSubmissionEntity"
}
type ProductSubmissionPrice struct {
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
    StringRepresentationValue   *string `json:"stringRepresentationValue" yaml:"stringRepresentationValue"    gorm:"-"     sql:"-"  `
    // Datenano also has a text representation
    Variations   []*  ProductSubmissionPriceVariations `json:"variations" yaml:"variations"    gorm:"foreignKey:LinkerId;references:UniqueId"     `
    // Datenano also has a text representation
	LinkedTo *ProductSubmissionEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
func ( x * ProductSubmissionPrice) RootObjectName() string {
	return "ProductSubmissionEntity"
}
type ProductSubmissionPriceVariations struct {
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
    Currency   *  currency.CurrencyEntity `json:"currency" yaml:"currency"    gorm:"foreignKey:CurrencyId;references:UniqueId"     `
    // Datenano also has a text representation
        CurrencyId *string `json:"currencyId" yaml:"currencyId"`
    Amount   *float64 `json:"amount" yaml:"amount"       `
    // Datenano also has a text representation
	LinkedTo *ProductSubmissionPrice `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
func ( x * ProductSubmissionPriceVariations) RootObjectName() string {
	return "ProductSubmissionEntity"
}
type ProductSubmissionEntity struct {
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
    Product   *  ProductEntity `json:"product" yaml:"product"    gorm:"foreignKey:ProductId;references:UniqueId"     `
    // Datenano also has a text representation
        ProductId *string `json:"productId" yaml:"productId" validate:"required" `
    Data  *workspaces.   JSON `json:"data" yaml:"data"       `
    // Datenano also has a text representation
    Values   []*  ProductSubmissionValues `json:"values" yaml:"values"    gorm:"foreignKey:LinkerId;references:UniqueId"     `
    // Datenano also has a text representation
    Name   *string `json:"name" yaml:"name"       `
    // Datenano also has a text representation
    Price   *  ProductSubmissionPrice `json:"price" yaml:"price"    gorm:"foreignKey:LinkerId;references:UniqueId"     `
    // Datenano also has a text representation
    Image   []*  workspaces.FileEntity `json:"image" yaml:"image"    gorm:"many2many:productSubmission_image;foreignKey:UniqueId;references:UniqueId"     `
    // Datenano also has a text representation
    ImageListId []string `json:"imageListId" yaml:"imageListId" gorm:"-" sql:"-"`
    Description   *string `json:"description" yaml:"description"       `
    // Datenano also has a text representation
    DescriptionExcerpt * string `json:"descriptionExcerpt" yaml:"descriptionExcerpt"`
    Sku   *string `json:"sku" yaml:"sku"       `
    // Datenano also has a text representation
    Brand   *  BrandEntity `json:"brand" yaml:"brand"    gorm:"foreignKey:BrandId;references:UniqueId"     `
    // Datenano also has a text representation
        BrandId *string `json:"brandId" yaml:"brandId"`
    Category   *  CategoryEntity `json:"category" yaml:"category"    gorm:"foreignKey:CategoryId;references:UniqueId"     `
    // Datenano also has a text representation
        CategoryId *string `json:"categoryId" yaml:"categoryId"`
    Tags   []*  TagEntity `json:"tags" yaml:"tags"    gorm:"many2many:productSubmission_tags;foreignKey:UniqueId;references:UniqueId"     `
    // Datenano also has a text representation
    TagsListId []string `json:"tagsListId" yaml:"tagsListId" gorm:"-" sql:"-"`
    Children []*ProductSubmissionEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *ProductSubmissionEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var ProductSubmissionPreloadRelations []string = []string{}
var PRODUCT_SUBMISSION_EVENT_CREATED = "productSubmission.created"
var PRODUCT_SUBMISSION_EVENT_UPDATED = "productSubmission.updated"
var PRODUCT_SUBMISSION_EVENT_DELETED = "productSubmission.deleted"
var PRODUCT_SUBMISSION_EVENTS = []string{
	PRODUCT_SUBMISSION_EVENT_CREATED,
	PRODUCT_SUBMISSION_EVENT_UPDATED,
	PRODUCT_SUBMISSION_EVENT_DELETED,
}
type ProductSubmissionFieldMap struct {
		Product workspaces.TranslatedString `yaml:"product"`
		Data workspaces.TranslatedString `yaml:"data"`
		Values workspaces.TranslatedString `yaml:"values"`
		Name workspaces.TranslatedString `yaml:"name"`
		Price workspaces.TranslatedString `yaml:"price"`
		Image workspaces.TranslatedString `yaml:"image"`
		Description workspaces.TranslatedString `yaml:"description"`
		Sku workspaces.TranslatedString `yaml:"sku"`
		Brand workspaces.TranslatedString `yaml:"brand"`
		Category workspaces.TranslatedString `yaml:"category"`
		Tags workspaces.TranslatedString `yaml:"tags"`
}
var ProductSubmissionEntityMetaConfig map[string]int64 = map[string]int64{
            "DescriptionExcerptSize": 100,
}
var ProductSubmissionEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&ProductSubmissionEntity{}))
func ProductSubmissionValuesActionCreate(
  dto *ProductSubmissionValues ,
  query workspaces.QueryDSL,
) (*ProductSubmissionValues , *workspaces.IError) {
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
func ProductSubmissionValuesActionUpdate(
    query workspaces.QueryDSL,
    dto *ProductSubmissionValues,
) (*ProductSubmissionValues, *workspaces.IError) {
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
func ProductSubmissionValuesActionGetOne(
    query workspaces.QueryDSL,
) (*ProductSubmissionValues , *workspaces.IError) {
    refl := reflect.ValueOf(&ProductSubmissionValues {})
    item, err := workspaces.GetOneEntity[ProductSubmissionValues ](query, refl)
    return item, err
}
func ProductSubmissionPriceActionCreate(
  dto *ProductSubmissionPrice ,
  query workspaces.QueryDSL,
) (*ProductSubmissionPrice , *workspaces.IError) {
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
func ProductSubmissionPriceActionUpdate(
    query workspaces.QueryDSL,
    dto *ProductSubmissionPrice,
) (*ProductSubmissionPrice, *workspaces.IError) {
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
func ProductSubmissionPriceActionGetOne(
    query workspaces.QueryDSL,
) (*ProductSubmissionPrice , *workspaces.IError) {
    refl := reflect.ValueOf(&ProductSubmissionPrice {})
    item, err := workspaces.GetOneEntity[ProductSubmissionPrice ](query, refl)
    return item, err
}
func entityProductSubmissionFormatter(dto *ProductSubmissionEntity, query workspaces.QueryDSL) {
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
func ProductSubmissionMockEntity() *ProductSubmissionEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &ProductSubmissionEntity{
      Name : &stringHolder,
      Sku : &stringHolder,
	}
	return entity
}
func ProductSubmissionActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := ProductSubmissionMockEntity()
		_, err := ProductSubmissionActionCreate(entity, query)
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
  func ProductSubmissionActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*ProductSubmissionEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &ProductSubmissionEntity{
          Values: []*ProductSubmissionValues{{}},
          Name: &tildaRef,
          Price: &ProductSubmissionPrice{},
          ImageListId: []string{"~"},
          Image: []*workspaces.FileEntity{{}},
          Sku: &tildaRef,
          TagsListId: []string{"~"},
          Tags: []*TagEntity{{}},
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
  func ProductSubmissionAssociationCreate(dto *ProductSubmissionEntity, query workspaces.QueryDSL) error {
      {
        if dto.ImageListId != nil && len(dto.ImageListId) > 0 {
          var items []workspaces.FileEntity
          err := query.Tx.Where(dto.ImageListId).Find(&items).Error
          if err != nil {
              return err
          }
          err = query.Tx.Model(dto).Association("Image").Replace(items)
          if err != nil {
              return err
          }
        }
      }
      {
        if dto.TagsListId != nil && len(dto.TagsListId) > 0 {
          var items []TagEntity
          err := query.Tx.Where(dto.TagsListId).Find(&items).Error
          if err != nil {
              return err
          }
          err = query.Tx.Model(dto).Association("Tags").Replace(items)
          if err != nil {
              return err
          }
        }
      }
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func ProductSubmissionRelationContentCreate(dto *ProductSubmissionEntity, query workspaces.QueryDSL) error {
return nil
}
func ProductSubmissionRelationContentUpdate(dto *ProductSubmissionEntity, query workspaces.QueryDSL) error {
	return nil
}
func ProductSubmissionPolyglotCreateHandler(dto *ProductSubmissionEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func ProductSubmissionValidator(dto *ProductSubmissionEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
        if dto != nil && dto.Values != nil {
          workspaces.AppendSliceErrors(dto.Values, isPatch, "values", err)
        }
    return err
  }
func ProductSubmissionEntityPreSanitize(dto *ProductSubmissionEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
			if (dto.Description != nil ) {
          Description := *dto.Description
          DescriptionExcerpt := stripPolicy.Sanitize(*dto.Description)
            Description = ugcPolicy.Sanitize(Description)
            DescriptionExcerpt = stripPolicy.Sanitize(DescriptionExcerpt)
        DescriptionExcerptSize, DescriptionExcerptSizeExists := ProductSubmissionEntityMetaConfig["DescriptionExcerptSize"]
        if DescriptionExcerptSizeExists {
          DescriptionExcerpt = workspaces.PickFirstNWords(DescriptionExcerpt, int(DescriptionExcerptSize))
        } else {
          DescriptionExcerpt = workspaces.PickFirstNWords(DescriptionExcerpt, 30)
        }
        dto.DescriptionExcerpt = &DescriptionExcerpt
        dto.Description = &Description
      }
}
  func ProductSubmissionEntityBeforeCreateAppend(dto *ProductSubmissionEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    ProductSubmissionRecursiveAddUniqueId(dto, query)
  }
  func ProductSubmissionRecursiveAddUniqueId(dto *ProductSubmissionEntity, query workspaces.QueryDSL) {
      if dto.Values != nil && len(dto.Values) > 0 {
        for index0 := range dto.Values {
          if (dto.Values[index0].UniqueId == "") {
            dto.Values[index0].UniqueId = workspaces.UUID()
          }
        }
    }
        if dto.Price != nil {
          dto.Price.UniqueId = workspaces.UUID()
        }
  }
func ProductSubmissionActionBatchCreateFn(dtos []*ProductSubmissionEntity, query workspaces.QueryDSL) ([]*ProductSubmissionEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*ProductSubmissionEntity{}
		for _, item := range dtos {
			s, err := ProductSubmissionActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func ProductSubmissionDeleteEntireChildren(query workspaces.QueryDSL, dto *ProductSubmissionEntity) (*workspaces.IError) {
  // intentionally removed this. It's hard to implement it, and probably wrong without
  // proper on delete cascade
  return nil
}
func ProductSubmissionActionCreateFn(dto *ProductSubmissionEntity, query workspaces.QueryDSL) (*ProductSubmissionEntity, *workspaces.IError) {
    ProductSubmissionCastFieldsToEavAndValidate(dto, query)
	// 1. Validate always
	if iError := ProductSubmissionValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	ProductSubmissionEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	ProductSubmissionEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	ProductSubmissionPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	ProductSubmissionRelationContentCreate(dto, query)
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
	ProductSubmissionAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(PRODUCT_SUBMISSION_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&ProductSubmissionEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func ProductSubmissionActionGetOne(query workspaces.QueryDSL) (*ProductSubmissionEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&ProductSubmissionEntity{})
    item, err := workspaces.GetOneEntity[ProductSubmissionEntity](query, refl)
    entityProductSubmissionFormatter(item, query)
    return item, err
  }
  func ProductSubmissionActionQuery(query workspaces.QueryDSL) ([]*ProductSubmissionEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&ProductSubmissionEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[ProductSubmissionEntity](query, refl)
    for _, item := range items {
      entityProductSubmissionFormatter(item, query)
    }
    return items, meta, err
  }
  func ProductSubmissionUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *ProductSubmissionEntity) (*ProductSubmissionEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = PRODUCT_SUBMISSION_EVENT_UPDATED
    ProductSubmissionEntityPreSanitize(fields, query)
    var item ProductSubmissionEntity
    q := dbref.
      Where(&ProductSubmissionEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    ProductSubmissionRelationContentUpdate(fields, query)
    ProductSubmissionPolyglotCreateHandler(fields, query)
    if ero := ProductSubmissionDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
        if fields.Price != nil {
          linkerId := uniqueId
          q := dbref.
            Model(&item.Price).
            Where(&ProductSubmissionPrice{LinkerId: &linkerId}).
            UpdateColumns(fields.Price)
          err := q.Error
          if err != nil {
            return &item, workspaces.GormErrorToIError(err)
          }
          if q.RowsAffected == 0 {
            fields.Price.UniqueId = workspaces.UUID()
            fields.Price.LinkerId = &linkerId
            err := dbref.
              Model(&item.Price).Create(fields.Price).Error
            if err != nil {
              return &item, workspaces.GormErrorToIError(err)
            }
          }
        }
    // @meta(update has many)
        if fields.ImageListId  != nil {
          var items []workspaces.FileEntity
          if len(fields.ImageListId ) > 0 {
            dbref.
              Where(&fields.ImageListId ).
              Find(&items)
          }
          dbref.
            Model(&ProductSubmissionEntity{UniqueId: uniqueId}).
            Association("Image").
            Replace(&items)
        }
        if fields.TagsListId  != nil {
          var items []TagEntity
          if len(fields.TagsListId ) > 0 {
            dbref.
              Where(&fields.TagsListId ).
              Find(&items)
          }
          dbref.
            Model(&ProductSubmissionEntity{UniqueId: uniqueId}).
            Association("Tags").
            Replace(&items)
        }
      if fields.Values != nil {
       linkerId := uniqueId;
        dbref.
          Where(&ProductSubmissionValues {LinkerId: &linkerId}).
          Delete(&ProductSubmissionValues {})
        for _, newItem := range fields.Values {
          newItem.UniqueId = workspaces.UUID()
          newItem.LinkerId = &linkerId
          dbref.Create(&newItem)
        }
      }
    err = dbref.
      Preload(clause.Associations).
      Where(&ProductSubmissionEntity{UniqueId: uniqueId}).
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
  func ProductSubmissionActionUpdateFn(query workspaces.QueryDSL, fields *ProductSubmissionEntity) (*ProductSubmissionEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
      ProductSubmissionCastFieldsToEavAndValidate(fields, query)
    // 1. Validate always
    if iError := ProductSubmissionValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // ProductSubmissionRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *ProductSubmissionEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = ProductSubmissionUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return ProductSubmissionUpdateExec(dbref, query, fields)
    }
  }
var ProductSubmissionWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire productsubmissions ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_SUBMISSION_DELETE},
    })
		count, _ := ProductSubmissionActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func ProductSubmissionActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&ProductSubmissionEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_SUBMISSION_DELETE}
	return workspaces.RemoveEntity[ProductSubmissionEntity](query, refl)
}
func ProductSubmissionActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
			{
				subCount, subErr := workspaces.WipeCleanEntity[ProductSubmissionValues]()
				if (subErr != nil) {
					fmt.Println("Error while wiping 'ProductSubmissionValues'", subErr)
					return count, subErr
				} else {
					count += subCount
				}
			}
			{
				subCount, subErr := workspaces.WipeCleanEntity[ProductSubmissionPrice]()
				if (subErr != nil) {
					fmt.Println("Error while wiping 'ProductSubmissionPrice'", subErr)
					return count, subErr
				} else {
					count += subCount
				}
			}
	{
		subCount, subErr := workspaces.WipeCleanEntity[ProductSubmissionEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'ProductSubmissionEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func ProductSubmissionActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[ProductSubmissionEntity]) (
    *workspaces.BulkRecordRequest[ProductSubmissionEntity], *workspaces.IError,
  ) {
    result := []*ProductSubmissionEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := ProductSubmissionActionUpdate(query, record)
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
func (x *ProductSubmissionEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var ProductSubmissionEntityMeta = workspaces.TableMetaData{
	EntityName:    "ProductSubmission",
	ExportKey:    "product-submissions",
	TableNameInDb: "fb_product-submission_entities",
	EntityObject:  &ProductSubmissionEntity{},
	ExportStream: ProductSubmissionActionExportT,
	ImportQuery: ProductSubmissionActionImport,
}
func ProductSubmissionActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[ProductSubmissionEntity](query, ProductSubmissionActionQuery, ProductSubmissionPreloadRelations)
}
func ProductSubmissionActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[ProductSubmissionEntity](query, ProductSubmissionActionQuery, ProductSubmissionPreloadRelations)
}
func ProductSubmissionActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content ProductSubmissionEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := ProductSubmissionActionCreate(&content, query)
	return err
}
var ProductSubmissionCommonCliFlags = []cli.Flag{
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
      Name:     "product-id",
      Required: true,
      Usage:    "product",
    },
    &cli.StringSliceFlag{
      Name:     "values",
      Required: false,
      Usage:    "values",
    },
    &cli.StringFlag{
      Name:     "name",
      Required: false,
      Usage:    "name",
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
    &cli.StringFlag{
      Name:     "string-representation-value",
      Required: false,
      Usage:    "stringRepresentationValue",
    },
    &cli.StringSliceFlag{
      Name:     "variations",
      Required: false,
      Usage:    "variations",
    },
    &cli.StringSliceFlag{
      Name:     "image",
      Required: false,
      Usage:    "image",
    },
    &cli.StringFlag{
      Name:     "description",
      Required: false,
      Usage:    "Detailed description of the product",
    },
    &cli.StringFlag{
      Name:     "sku",
      Required: false,
      Usage:    "Stock Keeping Unit code for the product",
    },
    &cli.StringFlag{
      Name:     "brand-id",
      Required: false,
      Usage:    "brand",
    },
    &cli.StringFlag{
      Name:     "category-id",
      Required: false,
      Usage:    "category",
    },
    &cli.StringSliceFlag{
      Name:     "tags",
      Required: false,
      Usage:    "tags",
    },
}
var ProductSubmissionCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "name",
		StructField:     "Name",
		Required: false,
		Usage:    "name",
		Type: "string",
	},
	{
		Name:     "sku",
		StructField:     "Sku",
		Required: false,
		Usage:    "Stock Keeping Unit code for the product",
		Type: "string",
	},
}
var ProductSubmissionCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "product-id",
      Required: true,
      Usage:    "product",
    },
    &cli.StringSliceFlag{
      Name:     "values",
      Required: false,
      Usage:    "values",
    },
    &cli.StringFlag{
      Name:     "name",
      Required: false,
      Usage:    "name",
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
    &cli.StringFlag{
      Name:     "string-representation-value",
      Required: false,
      Usage:    "stringRepresentationValue",
    },
    &cli.StringSliceFlag{
      Name:     "variations",
      Required: false,
      Usage:    "variations",
    },
    &cli.StringSliceFlag{
      Name:     "image",
      Required: false,
      Usage:    "image",
    },
    &cli.StringFlag{
      Name:     "description",
      Required: false,
      Usage:    "Detailed description of the product",
    },
    &cli.StringFlag{
      Name:     "sku",
      Required: false,
      Usage:    "Stock Keeping Unit code for the product",
    },
    &cli.StringFlag{
      Name:     "brand-id",
      Required: false,
      Usage:    "brand",
    },
    &cli.StringFlag{
      Name:     "category-id",
      Required: false,
      Usage:    "category",
    },
    &cli.StringSliceFlag{
      Name:     "tags",
      Required: false,
      Usage:    "tags",
    },
}
  var ProductSubmissionCreateCmd cli.Command = PRODUCT_SUBMISSION_ACTION_POST_ONE.ToCli()
  var ProductSubmissionCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_SUBMISSION_CREATE},
      })
      entity := &ProductSubmissionEntity{}
      for _, item := range ProductSubmissionCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := ProductSubmissionActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var ProductSubmissionUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: ProductSubmissionCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_SUBMISSION_UPDATE},
      })
      entity := CastProductSubmissionFromCli(c)
      if entity, err := ProductSubmissionActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* ProductSubmissionEntity) FromCli(c *cli.Context) *ProductSubmissionEntity {
	return CastProductSubmissionFromCli(c)
}
func CastProductSubmissionFromCli (c *cli.Context) *ProductSubmissionEntity {
	template := &ProductSubmissionEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("product-id") {
        value := c.String("product-id")
        template.ProductId = &value
      }
      if c.IsSet("name") {
        value := c.String("name")
        template.Name = &value
      }
      if c.IsSet("image") {
        value := c.String("image")
        template.ImageListId = strings.Split(value, ",")
      }
      if c.IsSet("description") {
        value := c.String("description")
        template.Description = &value
      }
      if c.IsSet("sku") {
        value := c.String("sku")
        template.Sku = &value
      }
      if c.IsSet("brand-id") {
        value := c.String("brand-id")
        template.BrandId = &value
      }
      if c.IsSet("category-id") {
        value := c.String("category-id")
        template.CategoryId = &value
      }
      if c.IsSet("tags") {
        value := c.String("tags")
        template.TagsListId = strings.Split(value, ",")
      }
	return template
}
  func ProductSubmissionSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      ProductSubmissionActionCreate,
      reflect.ValueOf(&ProductSubmissionEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func ProductSubmissionImportMocks() {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      ProductSubmissionActionCreate,
      reflect.ValueOf(&ProductSubmissionEntity{}).Elem(),
      &mocks.ViewsFs,
      []string{},
      false,
    )
  }
  func ProductSubmissionWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := ProductSubmissionActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "ProductSubmission", result)
    }
  }
var ProductSubmissionImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_SUBMISSION_CREATE},
      })
			ProductSubmissionActionSeeder(query, c.Int("count"))
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
				Value: "product-submission-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_SUBMISSION_CREATE},
      })
			ProductSubmissionActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "product-submission-seeder-product-submission.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of product-submissions, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]ProductSubmissionEntity{}
			workspaces.ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
		cli.Command{
			Name:  "mocks",
			Usage: "Prints the list of mocks",
			Action: func(c *cli.Context) error {
				if entity, err := workspaces.GetSeederFilenames(&mocks.ViewsFs, ""); err != nil {
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
				workspaces.CommonCliImportEmbedCmd(c,
					ProductSubmissionActionCreate,
					reflect.ValueOf(&ProductSubmissionEntity{}).Elem(),
					&mocks.ViewsFs,
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
			ProductSubmissionCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				ProductSubmissionActionCreate,
				reflect.ValueOf(&ProductSubmissionEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_SUBMISSION_CREATE},
				},
        func() ProductSubmissionEntity {
					v := CastProductSubmissionFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var ProductSubmissionCliCommands []cli.Command = []cli.Command{
      PRODUCT_SUBMISSION_ACTION_QUERY.ToCli(),
      PRODUCT_SUBMISSION_ACTION_TABLE.ToCli(),
      ProductSubmissionCreateCmd,
      ProductSubmissionUpdateCmd,
      ProductSubmissionCreateInteractiveCmd,
      ProductSubmissionWipeCmd,
      workspaces.GetCommonRemoveQuery(reflect.ValueOf(&ProductSubmissionEntity{}).Elem(), ProductSubmissionActionRemove),
  }
  func ProductSubmissionCliFn() cli.Command {
    ProductSubmissionCliCommands = append(ProductSubmissionCliCommands, ProductSubmissionImportExportCommands...)
    return cli.Command{
      Name:        "productSubmission",
      Description: "ProductSubmissions module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: ProductSubmissionCliCommands,
    }
  }
var PRODUCT_SUBMISSION_ACTION_TABLE = workspaces.Module2Action{
  Name:    "table",
  ActionName: "table",
  ActionAliases: []string{"t"},
  Flags:  workspaces.CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: ProductSubmissionActionQuery,
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    workspaces.CommonCliTableCmd2(c,
      ProductSubmissionActionQuery,
      security,
      reflect.ValueOf(&ProductSubmissionEntity{}).Elem(),
    )
    return nil
  },
}
var PRODUCT_SUBMISSION_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/product-submissions",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_SUBMISSION_QUERY},
  },
  Group: "productSubmission",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, ProductSubmissionActionQuery)
    },
  },
  Format: "QUERY",
  Action: ProductSubmissionActionQuery,
  ResponseEntity: &[]ProductSubmissionEntity{},
  Out: workspaces.Module2ActionBody{
		Entity: "ProductSubmissionEntity",
	},
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			ProductSubmissionActionQuery,
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
var PRODUCT_SUBMISSION_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/product-submissions/export",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_SUBMISSION_QUERY},
  },
  Group: "productSubmission",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, ProductSubmissionActionExport)
    },
  },
  Format: "QUERY",
  Action: ProductSubmissionActionExport,
  ResponseEntity: &[]ProductSubmissionEntity{},
  Out: workspaces.Module2ActionBody{
		Entity: "ProductSubmissionEntity",
	},
}
var PRODUCT_SUBMISSION_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/product-submission/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_SUBMISSION_QUERY},
  },
  Group: "productSubmission",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, ProductSubmissionActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: ProductSubmissionActionGetOne,
  ResponseEntity: &ProductSubmissionEntity{},
  Out: workspaces.Module2ActionBody{
		Entity: "ProductSubmissionEntity",
	},
}
var PRODUCT_SUBMISSION_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new productSubmission",
  Flags: ProductSubmissionCommonCliFlags,
  Method: "POST",
  Url:    "/product-submission",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_SUBMISSION_CREATE},
  },
  Group: "productSubmission",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, ProductSubmissionActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, ProductSubmissionActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: ProductSubmissionActionCreate,
  Format: "POST_ONE",
  RequestEntity: &ProductSubmissionEntity{},
  ResponseEntity: &ProductSubmissionEntity{},
  Out: workspaces.Module2ActionBody{
		Entity: "ProductSubmissionEntity",
	},
  In: workspaces.Module2ActionBody{
		Entity: "ProductSubmissionEntity",
	},
}
var PRODUCT_SUBMISSION_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: ProductSubmissionCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/product-submission",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_SUBMISSION_UPDATE},
  },
  Group: "productSubmission",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, ProductSubmissionActionUpdate)
    },
  },
  Action: ProductSubmissionActionUpdate,
  RequestEntity: &ProductSubmissionEntity{},
  ResponseEntity: &ProductSubmissionEntity{},
  Format: "PATCH_ONE",
  Out: workspaces.Module2ActionBody{
		Entity: "ProductSubmissionEntity",
	},
  In: workspaces.Module2ActionBody{
		Entity: "ProductSubmissionEntity",
	},
}
var PRODUCT_SUBMISSION_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/product-submissions",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_SUBMISSION_UPDATE},
  },
  Group: "productSubmission",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, ProductSubmissionActionBulkUpdate)
    },
  },
  Action: ProductSubmissionActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[ProductSubmissionEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[ProductSubmissionEntity]{},
  Out: workspaces.Module2ActionBody{
		Entity: "ProductSubmissionEntity",
	},
  In: workspaces.Module2ActionBody{
		Entity: "ProductSubmissionEntity",
	},
}
var PRODUCT_SUBMISSION_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/product-submission",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_SUBMISSION_DELETE},
  },
  Group: "productSubmission",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, ProductSubmissionActionRemove)
    },
  },
  Action: ProductSubmissionActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &ProductSubmissionEntity{},
}
    var PRODUCT_SUBMISSION_VALUES_ACTION_PATCH = workspaces.Module2Action{
      Method: "PATCH",
      Url:    "/product-submission/:linkerId/values/:uniqueId",
      SecurityModel: &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_SUBMISSION_UPDATE},
      },
      Group: "productSubmission",
      Handlers: []gin.HandlerFunc{
        func (
          c *gin.Context,
        ) {
          workspaces.HttpUpdateEntity(c, ProductSubmissionValuesActionUpdate)
        },
      },
      Action: ProductSubmissionValuesActionUpdate,
      Format: "PATCH_ONE",
      RequestEntity: &ProductSubmissionValues{},
      ResponseEntity: &ProductSubmissionValues{},
      Out: workspaces.Module2ActionBody{
        Entity: "ProductSubmissionValues",
      },
      In: workspaces.Module2ActionBody{
        Entity: "ProductSubmissionValues",
      },
    }
    var PRODUCT_SUBMISSION_VALUES_ACTION_GET = workspaces.Module2Action {
      Method: "GET",
      Url:    "/product-submission/values/:linkerId/:uniqueId",
      SecurityModel: &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_SUBMISSION_QUERY},
      },
      Group: "productSubmission",
      Handlers: []gin.HandlerFunc{
        func (
          c *gin.Context,
        ) {
          workspaces.HttpGetEntity(c, ProductSubmissionValuesActionGetOne)
        },
      },
      Action: ProductSubmissionValuesActionGetOne,
      Format: "GET_ONE",
      ResponseEntity: &ProductSubmissionValues{},
      Out: workspaces.Module2ActionBody{
        Entity: "ProductSubmissionValues",
      },
    }
    var PRODUCT_SUBMISSION_VALUES_ACTION_POST = workspaces.Module2Action{
      Method: "POST",
      Url:    "/product-submission/:linkerId/values",
      SecurityModel: &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_SUBMISSION_CREATE},
      },
      Group: "productSubmission",
      Handlers: []gin.HandlerFunc{
        func (
          c *gin.Context,
        ) {
          workspaces.HttpPostEntity(c, ProductSubmissionValuesActionCreate)
        },
      },
      Action: ProductSubmissionValuesActionCreate,
      Format: "POST_ONE",
      RequestEntity: &ProductSubmissionValues{},
      ResponseEntity: &ProductSubmissionValues{},
      Out: workspaces.Module2ActionBody{
        Entity: "ProductSubmissionValues",
      },
      In: workspaces.Module2ActionBody{
        Entity: "ProductSubmissionValues",
      },
    }
    var PRODUCT_SUBMISSION_PRICE_ACTION_PATCH = workspaces.Module2Action{
      Method: "PATCH",
      Url:    "/product-submission/:linkerId/price/:uniqueId",
      SecurityModel: &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_SUBMISSION_UPDATE},
      },
      Group: "productSubmission",
      Handlers: []gin.HandlerFunc{
        func (
          c *gin.Context,
        ) {
          workspaces.HttpUpdateEntity(c, ProductSubmissionPriceActionUpdate)
        },
      },
      Action: ProductSubmissionPriceActionUpdate,
      Format: "PATCH_ONE",
      RequestEntity: &ProductSubmissionPrice{},
      ResponseEntity: &ProductSubmissionPrice{},
      Out: workspaces.Module2ActionBody{
        Entity: "ProductSubmissionPrice",
      },
      In: workspaces.Module2ActionBody{
        Entity: "ProductSubmissionPrice",
      },
    }
    var PRODUCT_SUBMISSION_PRICE_ACTION_GET = workspaces.Module2Action {
      Method: "GET",
      Url:    "/product-submission/price/:linkerId/:uniqueId",
      SecurityModel: &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_SUBMISSION_QUERY},
      },
      Group: "productSubmission",
      Handlers: []gin.HandlerFunc{
        func (
          c *gin.Context,
        ) {
          workspaces.HttpGetEntity(c, ProductSubmissionPriceActionGetOne)
        },
      },
      Action: ProductSubmissionPriceActionGetOne,
      Format: "GET_ONE",
      ResponseEntity: &ProductSubmissionPrice{},
      Out: workspaces.Module2ActionBody{
        Entity: "ProductSubmissionPrice",
      },
    }
    var PRODUCT_SUBMISSION_PRICE_ACTION_POST = workspaces.Module2Action{
      Method: "POST",
      Url:    "/product-submission/:linkerId/price",
      SecurityModel: &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_PRODUCT_SUBMISSION_CREATE},
      },
      Group: "productSubmission",
      Handlers: []gin.HandlerFunc{
        func (
          c *gin.Context,
        ) {
          workspaces.HttpPostEntity(c, ProductSubmissionPriceActionCreate)
        },
      },
      Action: ProductSubmissionPriceActionCreate,
      Format: "POST_ONE",
      RequestEntity: &ProductSubmissionPrice{},
      ResponseEntity: &ProductSubmissionPrice{},
      Out: workspaces.Module2ActionBody{
        Entity: "ProductSubmissionPrice",
      },
      In: workspaces.Module2ActionBody{
        Entity: "ProductSubmissionPrice",
      },
    }
  /**
  *	Override this function on ProductSubmissionEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendProductSubmissionRouter = func(r *[]workspaces.Module2Action) {}
  func GetProductSubmissionModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      PRODUCT_SUBMISSION_ACTION_QUERY,
      PRODUCT_SUBMISSION_ACTION_EXPORT,
      PRODUCT_SUBMISSION_ACTION_GET_ONE,
      PRODUCT_SUBMISSION_ACTION_POST_ONE,
      PRODUCT_SUBMISSION_ACTION_PATCH,
      PRODUCT_SUBMISSION_ACTION_PATCH_BULK,
      PRODUCT_SUBMISSION_ACTION_DELETE,
          PRODUCT_SUBMISSION_VALUES_ACTION_PATCH,
          PRODUCT_SUBMISSION_VALUES_ACTION_GET,
          PRODUCT_SUBMISSION_VALUES_ACTION_POST,
          PRODUCT_SUBMISSION_PRICE_ACTION_PATCH,
          PRODUCT_SUBMISSION_PRICE_ACTION_GET,
          PRODUCT_SUBMISSION_PRICE_ACTION_POST,
    }
    // Append user defined functions
    AppendProductSubmissionRouter(&routes)
    return routes
  }
  func CreateProductSubmissionRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetProductSubmissionModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, ProductSubmissionEntityJsonSchema, "product-submission-http", "shop")
    workspaces.WriteEntitySchema("ProductSubmissionEntity", ProductSubmissionEntityJsonSchema, "shop")
    return httpRoutes
  }
var PERM_ROOT_PRODUCT_SUBMISSION_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/product-submission/delete",
  Name: "Delete product submission",
}
var PERM_ROOT_PRODUCT_SUBMISSION_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/product-submission/create",
  Name: "Create product submission",
}
var PERM_ROOT_PRODUCT_SUBMISSION_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/product-submission/update",
  Name: "Update product submission",
}
var PERM_ROOT_PRODUCT_SUBMISSION_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/shop/product-submission/query",
  Name: "Query product submission",
}
var PERM_ROOT_PRODUCT_SUBMISSION = workspaces.PermissionInfo{
  CompleteKey: "root/shop/product-submission/*",
  Name: "Entire product submission actions (*)",
}
var ALL_PRODUCT_SUBMISSION_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_PRODUCT_SUBMISSION_DELETE,
	PERM_ROOT_PRODUCT_SUBMISSION_CREATE,
	PERM_ROOT_PRODUCT_SUBMISSION_UPDATE,
	PERM_ROOT_PRODUCT_SUBMISSION_QUERY,
	PERM_ROOT_PRODUCT_SUBMISSION,
}