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
type DiscountCodeEntity struct {
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
    Series   *string `json:"series" yaml:"series"       `
    // Datenano also has a text representation
    Limit   *int64 `json:"limit" yaml:"limit"       `
    // Datenano also has a text representation
    // Datenano also has a text representation
    // Date range is a complex date storage
    ValidStart workspaces.XDate `json:"validStart" yaml:"validStart"`
    ValidEnd workspaces.XDate `json:"validEnd" yaml:"validEnd"`
    Valid workspaces.XDateComputed `json:"valid" yaml:"valid" gorm:"-" sql:"-"`
    AppliedProducts   []*  ProductSubmissionEntity `json:"appliedProducts" yaml:"appliedProducts"    gorm:"many2many:discountCode_appliedProducts;foreignKey:UniqueId;references:UniqueId"     `
    // Datenano also has a text representation
    AppliedProductsListId []string `json:"appliedProductsListId" yaml:"appliedProductsListId" gorm:"-" sql:"-"`
    ExcludedProducts   []*  ProductSubmissionEntity `json:"excludedProducts" yaml:"excludedProducts"    gorm:"many2many:discountCode_excludedProducts;foreignKey:UniqueId;references:UniqueId"     `
    // Datenano also has a text representation
    ExcludedProductsListId []string `json:"excludedProductsListId" yaml:"excludedProductsListId" gorm:"-" sql:"-"`
    AppliedCategories   []*  CategoryEntity `json:"appliedCategories" yaml:"appliedCategories"    gorm:"many2many:discountCode_appliedCategories;foreignKey:UniqueId;references:UniqueId"     `
    // Datenano also has a text representation
    AppliedCategoriesListId []string `json:"appliedCategoriesListId" yaml:"appliedCategoriesListId" gorm:"-" sql:"-"`
    ExcludedCategories   []*  CategoryEntity `json:"excludedCategories" yaml:"excludedCategories"    gorm:"many2many:discountCode_excludedCategories;foreignKey:UniqueId;references:UniqueId"     `
    // Datenano also has a text representation
    ExcludedCategoriesListId []string `json:"excludedCategoriesListId" yaml:"excludedCategoriesListId" gorm:"-" sql:"-"`
    Children []*DiscountCodeEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *DiscountCodeEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var DiscountCodePreloadRelations []string = []string{}
var DISCOUNT_CODE_EVENT_CREATED = "discountCode.created"
var DISCOUNT_CODE_EVENT_UPDATED = "discountCode.updated"
var DISCOUNT_CODE_EVENT_DELETED = "discountCode.deleted"
var DISCOUNT_CODE_EVENTS = []string{
	DISCOUNT_CODE_EVENT_CREATED,
	DISCOUNT_CODE_EVENT_UPDATED,
	DISCOUNT_CODE_EVENT_DELETED,
}
type DiscountCodeFieldMap struct {
		Series workspaces.TranslatedString `yaml:"series"`
		Limit workspaces.TranslatedString `yaml:"limit"`
		Valid workspaces.TranslatedString `yaml:"valid"`
		AppliedProducts workspaces.TranslatedString `yaml:"appliedProducts"`
		ExcludedProducts workspaces.TranslatedString `yaml:"excludedProducts"`
		AppliedCategories workspaces.TranslatedString `yaml:"appliedCategories"`
		ExcludedCategories workspaces.TranslatedString `yaml:"excludedCategories"`
}
var DiscountCodeEntityMetaConfig map[string]int64 = map[string]int64{
}
var DiscountCodeEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&DiscountCodeEntity{}))
func entityDiscountCodeFormatter(dto *DiscountCodeEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
			dto.Valid = workspaces.ComputeDateRange(dto.ValidStart, dto.ValidEnd , query)
	if dto.Created > 0 {
		dto.CreatedFormatted = workspaces.FormatDateBasedOnQuery(dto.Created, query)
	}
	if dto.Updated > 0 {
		dto.CreatedFormatted = workspaces.FormatDateBasedOnQuery(dto.Updated, query)
	}
}
func DiscountCodeMockEntity() *DiscountCodeEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &DiscountCodeEntity{
      Series : &stringHolder,
      Limit : &int64Holder,
	}
	return entity
}
func DiscountCodeActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := DiscountCodeMockEntity()
		_, err := DiscountCodeActionCreate(entity, query)
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
  func DiscountCodeActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*DiscountCodeEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &DiscountCodeEntity{
          Series: &tildaRef,
          AppliedProductsListId: []string{"~"},
          AppliedProducts: []*ProductSubmissionEntity{{}},
          ExcludedProductsListId: []string{"~"},
          ExcludedProducts: []*ProductSubmissionEntity{{}},
          AppliedCategoriesListId: []string{"~"},
          AppliedCategories: []*CategoryEntity{{}},
          ExcludedCategoriesListId: []string{"~"},
          ExcludedCategories: []*CategoryEntity{{}},
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
  func DiscountCodeAssociationCreate(dto *DiscountCodeEntity, query workspaces.QueryDSL) error {
      {
        if dto.AppliedProductsListId != nil && len(dto.AppliedProductsListId) > 0 {
          var items []ProductSubmissionEntity
          err := query.Tx.Where(dto.AppliedProductsListId).Find(&items).Error
          if err != nil {
              return err
          }
          err = query.Tx.Model(dto).Association("AppliedProducts").Replace(items)
          if err != nil {
              return err
          }
        }
      }
      {
        if dto.ExcludedProductsListId != nil && len(dto.ExcludedProductsListId) > 0 {
          var items []ProductSubmissionEntity
          err := query.Tx.Where(dto.ExcludedProductsListId).Find(&items).Error
          if err != nil {
              return err
          }
          err = query.Tx.Model(dto).Association("ExcludedProducts").Replace(items)
          if err != nil {
              return err
          }
        }
      }
      {
        if dto.AppliedCategoriesListId != nil && len(dto.AppliedCategoriesListId) > 0 {
          var items []CategoryEntity
          err := query.Tx.Where(dto.AppliedCategoriesListId).Find(&items).Error
          if err != nil {
              return err
          }
          err = query.Tx.Model(dto).Association("AppliedCategories").Replace(items)
          if err != nil {
              return err
          }
        }
      }
      {
        if dto.ExcludedCategoriesListId != nil && len(dto.ExcludedCategoriesListId) > 0 {
          var items []CategoryEntity
          err := query.Tx.Where(dto.ExcludedCategoriesListId).Find(&items).Error
          if err != nil {
              return err
          }
          err = query.Tx.Model(dto).Association("ExcludedCategories").Replace(items)
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
func DiscountCodeRelationContentCreate(dto *DiscountCodeEntity, query workspaces.QueryDSL) error {
return nil
}
func DiscountCodeRelationContentUpdate(dto *DiscountCodeEntity, query workspaces.QueryDSL) error {
	return nil
}
func DiscountCodePolyglotCreateHandler(dto *DiscountCodeEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func DiscountCodeValidator(dto *DiscountCodeEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func DiscountCodeEntityPreSanitize(dto *DiscountCodeEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func DiscountCodeEntityBeforeCreateAppend(dto *DiscountCodeEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    DiscountCodeRecursiveAddUniqueId(dto, query)
  }
  func DiscountCodeRecursiveAddUniqueId(dto *DiscountCodeEntity, query workspaces.QueryDSL) {
  }
func DiscountCodeActionBatchCreateFn(dtos []*DiscountCodeEntity, query workspaces.QueryDSL) ([]*DiscountCodeEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*DiscountCodeEntity{}
		for _, item := range dtos {
			s, err := DiscountCodeActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func DiscountCodeDeleteEntireChildren(query workspaces.QueryDSL, dto *DiscountCodeEntity) (*workspaces.IError) {
  return nil
}
func DiscountCodeActionCreateFn(dto *DiscountCodeEntity, query workspaces.QueryDSL) (*DiscountCodeEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := DiscountCodeValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	DiscountCodeEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	DiscountCodeEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	DiscountCodePolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	DiscountCodeRelationContentCreate(dto, query)
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
	DiscountCodeAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(DISCOUNT_CODE_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&DiscountCodeEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func DiscountCodeActionGetOne(query workspaces.QueryDSL) (*DiscountCodeEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&DiscountCodeEntity{})
    item, err := workspaces.GetOneEntity[DiscountCodeEntity](query, refl)
    entityDiscountCodeFormatter(item, query)
    return item, err
  }
  func DiscountCodeActionQuery(query workspaces.QueryDSL) ([]*DiscountCodeEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&DiscountCodeEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[DiscountCodeEntity](query, refl)
    for _, item := range items {
      entityDiscountCodeFormatter(item, query)
    }
    return items, meta, err
  }
  func DiscountCodeUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *DiscountCodeEntity) (*DiscountCodeEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = DISCOUNT_CODE_EVENT_UPDATED
    DiscountCodeEntityPreSanitize(fields, query)
    var item DiscountCodeEntity
    q := dbref.
      Where(&DiscountCodeEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    DiscountCodeRelationContentUpdate(fields, query)
    DiscountCodePolyglotCreateHandler(fields, query)
    if ero := DiscountCodeDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
        if fields.AppliedProductsListId  != nil {
          var items []ProductSubmissionEntity
          if len(fields.AppliedProductsListId ) > 0 {
            dbref.
              Where(&fields.AppliedProductsListId ).
              Find(&items)
          }
          dbref.
            Model(&DiscountCodeEntity{UniqueId: uniqueId}).
            Association("AppliedProducts").
            Replace(&items)
        }
        if fields.ExcludedProductsListId  != nil {
          var items []ProductSubmissionEntity
          if len(fields.ExcludedProductsListId ) > 0 {
            dbref.
              Where(&fields.ExcludedProductsListId ).
              Find(&items)
          }
          dbref.
            Model(&DiscountCodeEntity{UniqueId: uniqueId}).
            Association("ExcludedProducts").
            Replace(&items)
        }
        if fields.AppliedCategoriesListId  != nil {
          var items []CategoryEntity
          if len(fields.AppliedCategoriesListId ) > 0 {
            dbref.
              Where(&fields.AppliedCategoriesListId ).
              Find(&items)
          }
          dbref.
            Model(&DiscountCodeEntity{UniqueId: uniqueId}).
            Association("AppliedCategories").
            Replace(&items)
        }
        if fields.ExcludedCategoriesListId  != nil {
          var items []CategoryEntity
          if len(fields.ExcludedCategoriesListId ) > 0 {
            dbref.
              Where(&fields.ExcludedCategoriesListId ).
              Find(&items)
          }
          dbref.
            Model(&DiscountCodeEntity{UniqueId: uniqueId}).
            Association("ExcludedCategories").
            Replace(&items)
        }
    err = dbref.
      Preload(clause.Associations).
      Where(&DiscountCodeEntity{UniqueId: uniqueId}).
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
  func DiscountCodeActionUpdateFn(query workspaces.QueryDSL, fields *DiscountCodeEntity) (*DiscountCodeEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := DiscountCodeValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // DiscountCodeRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *DiscountCodeEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = DiscountCodeUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return DiscountCodeUpdateExec(dbref, query, fields)
    }
  }
var DiscountCodeWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire discountcodes ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_DISCOUNT_CODE_DELETE},
    })
		count, _ := DiscountCodeActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func DiscountCodeActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&DiscountCodeEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_DISCOUNT_CODE_DELETE}
	return workspaces.RemoveEntity[DiscountCodeEntity](query, refl)
}
func DiscountCodeActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[DiscountCodeEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'DiscountCodeEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func DiscountCodeActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[DiscountCodeEntity]) (
    *workspaces.BulkRecordRequest[DiscountCodeEntity], *workspaces.IError,
  ) {
    result := []*DiscountCodeEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := DiscountCodeActionUpdate(query, record)
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
func (x *DiscountCodeEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var DiscountCodeEntityMeta = workspaces.TableMetaData{
	EntityName:    "DiscountCode",
	ExportKey:    "discount-codes",
	TableNameInDb: "fb_discount-code_entities",
	EntityObject:  &DiscountCodeEntity{},
	ExportStream: DiscountCodeActionExportT,
	ImportQuery: DiscountCodeActionImport,
}
func DiscountCodeActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[DiscountCodeEntity](query, DiscountCodeActionQuery, DiscountCodePreloadRelations)
}
func DiscountCodeActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[DiscountCodeEntity](query, DiscountCodeActionQuery, DiscountCodePreloadRelations)
}
func DiscountCodeActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content DiscountCodeEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := DiscountCodeActionCreate(&content, query)
	return err
}
var DiscountCodeCommonCliFlags = []cli.Flag{
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
      Name:     "series",
      Required: false,
      Usage:    "series",
    },
    &cli.Int64Flag{
      Name:     "limit",
      Required: false,
      Usage:    "limit",
    },
    &cli.StringFlag{
      Name:     "valid-start",
      Required: false,
      Usage:    "valid",
    },
    &cli.StringFlag{
      Name:     "valid-end",
      Required: false,
      Usage:    "valid",
    },
    &cli.StringSliceFlag{
      Name:     "applied-products",
      Required: false,
      Usage:    "appliedProducts",
    },
    &cli.StringSliceFlag{
      Name:     "excluded-products",
      Required: false,
      Usage:    "excludedProducts",
    },
    &cli.StringSliceFlag{
      Name:     "applied-categories",
      Required: false,
      Usage:    "appliedCategories",
    },
    &cli.StringSliceFlag{
      Name:     "excluded-categories",
      Required: false,
      Usage:    "excludedCategories",
    },
}
var DiscountCodeCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "series",
		StructField:     "Series",
		Required: false,
		Usage:    "series",
		Type: "string",
	},
	{
		Name:     "limit",
		StructField:     "Limit",
		Required: false,
		Usage:    "limit",
		Type: "int64",
	},
}
var DiscountCodeCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "series",
      Required: false,
      Usage:    "series",
    },
    &cli.Int64Flag{
      Name:     "limit",
      Required: false,
      Usage:    "limit",
    },
    &cli.StringFlag{
      Name:     "valid-start",
      Required: false,
      Usage:    "valid",
    },
    &cli.StringFlag{
      Name:     "valid-end",
      Required: false,
      Usage:    "valid",
    },
    &cli.StringSliceFlag{
      Name:     "applied-products",
      Required: false,
      Usage:    "appliedProducts",
    },
    &cli.StringSliceFlag{
      Name:     "excluded-products",
      Required: false,
      Usage:    "excludedProducts",
    },
    &cli.StringSliceFlag{
      Name:     "applied-categories",
      Required: false,
      Usage:    "appliedCategories",
    },
    &cli.StringSliceFlag{
      Name:     "excluded-categories",
      Required: false,
      Usage:    "excludedCategories",
    },
}
  var DiscountCodeCreateCmd cli.Command = DISCOUNT_CODE_ACTION_POST_ONE.ToCli()
  var DiscountCodeCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_DISCOUNT_CODE_CREATE},
      })
      entity := &DiscountCodeEntity{}
      for _, item := range DiscountCodeCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := DiscountCodeActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var DiscountCodeUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: DiscountCodeCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_DISCOUNT_CODE_UPDATE},
      })
      entity := CastDiscountCodeFromCli(c)
      if entity, err := DiscountCodeActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* DiscountCodeEntity) FromCli(c *cli.Context) *DiscountCodeEntity {
	return CastDiscountCodeFromCli(c)
}
func CastDiscountCodeFromCli (c *cli.Context) *DiscountCodeEntity {
	template := &DiscountCodeEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("series") {
        value := c.String("series")
        template.Series = &value
      }
      if c.IsSet("valid-start") {
        value := c.String("valid-start")
        template.ValidStart.Scan(value)
      }
      if c.IsSet("valid-end") {
        value := c.String("valid-end")
        template.ValidEnd.Scan(value)
      }
      if c.IsSet("applied-products") {
        value := c.String("applied-products")
        template.AppliedProductsListId = strings.Split(value, ",")
      }
      if c.IsSet("excluded-products") {
        value := c.String("excluded-products")
        template.ExcludedProductsListId = strings.Split(value, ",")
      }
      if c.IsSet("applied-categories") {
        value := c.String("applied-categories")
        template.AppliedCategoriesListId = strings.Split(value, ",")
      }
      if c.IsSet("excluded-categories") {
        value := c.String("excluded-categories")
        template.ExcludedCategoriesListId = strings.Split(value, ",")
      }
	return template
}
  func DiscountCodeSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      DiscountCodeActionCreate,
      reflect.ValueOf(&DiscountCodeEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func DiscountCodeWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := DiscountCodeActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "DiscountCode", result)
    }
  }
var DiscountCodeImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_DISCOUNT_CODE_CREATE},
      })
			DiscountCodeActionSeeder(query, c.Int("count"))
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
				Value: "discount-code-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_DISCOUNT_CODE_CREATE},
      })
			DiscountCodeActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "discount-code-seeder-discount-code.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of discount-codes, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]DiscountCodeEntity{}
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
			DiscountCodeCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				DiscountCodeActionCreate,
				reflect.ValueOf(&DiscountCodeEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_DISCOUNT_CODE_CREATE},
				},
        func() DiscountCodeEntity {
					v := CastDiscountCodeFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var DiscountCodeCliCommands []cli.Command = []cli.Command{
      DISCOUNT_CODE_ACTION_QUERY.ToCli(),
      DISCOUNT_CODE_ACTION_TABLE.ToCli(),
      DiscountCodeCreateCmd,
      DiscountCodeUpdateCmd,
      DiscountCodeCreateInteractiveCmd,
      DiscountCodeWipeCmd,
      workspaces.GetCommonRemoveQuery(reflect.ValueOf(&DiscountCodeEntity{}).Elem(), DiscountCodeActionRemove),
  }
  func DiscountCodeCliFn() cli.Command {
    DiscountCodeCliCommands = append(DiscountCodeCliCommands, DiscountCodeImportExportCommands...)
    return cli.Command{
      Name:        "discountCode",
      Description: "DiscountCodes module actions (sample module to handle complex entities)",
      Usage:       "List of all discount codes inside the application",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: DiscountCodeCliCommands,
    }
  }
var DISCOUNT_CODE_ACTION_TABLE = workspaces.Module2Action{
  Name:    "table",
  ActionAliases: []string{"t"},
  Flags:  workspaces.CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: DiscountCodeActionQuery,
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    workspaces.CommonCliTableCmd2(c,
      DiscountCodeActionQuery,
      security,
      reflect.ValueOf(&DiscountCodeEntity{}).Elem(),
    )
    return nil
  },
}
var DISCOUNT_CODE_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/discount-codes",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_DISCOUNT_CODE_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, DiscountCodeActionQuery)
    },
  },
  Format: "QUERY",
  Action: DiscountCodeActionQuery,
  ResponseEntity: &[]DiscountCodeEntity{},
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			DiscountCodeActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionAliases: []string{"q"},
	Flags:         workspaces.CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var DISCOUNT_CODE_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/discount-codes/export",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_DISCOUNT_CODE_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, DiscountCodeActionExport)
    },
  },
  Format: "QUERY",
  Action: DiscountCodeActionExport,
  ResponseEntity: &[]DiscountCodeEntity{},
}
var DISCOUNT_CODE_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/discount-code/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_DISCOUNT_CODE_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, DiscountCodeActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: DiscountCodeActionGetOne,
  ResponseEntity: &DiscountCodeEntity{},
}
var DISCOUNT_CODE_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new discountCode",
  Flags: DiscountCodeCommonCliFlags,
  Method: "POST",
  Url:    "/discount-code",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_DISCOUNT_CODE_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, DiscountCodeActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, DiscountCodeActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: DiscountCodeActionCreate,
  Format: "POST_ONE",
  RequestEntity: &DiscountCodeEntity{},
  ResponseEntity: &DiscountCodeEntity{},
}
var DISCOUNT_CODE_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: DiscountCodeCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/discount-code",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_DISCOUNT_CODE_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, DiscountCodeActionUpdate)
    },
  },
  Action: DiscountCodeActionUpdate,
  RequestEntity: &DiscountCodeEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &DiscountCodeEntity{},
}
var DISCOUNT_CODE_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/discount-codes",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_DISCOUNT_CODE_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, DiscountCodeActionBulkUpdate)
    },
  },
  Action: DiscountCodeActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[DiscountCodeEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[DiscountCodeEntity]{},
}
var DISCOUNT_CODE_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/discount-code",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_DISCOUNT_CODE_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, DiscountCodeActionRemove)
    },
  },
  Action: DiscountCodeActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &DiscountCodeEntity{},
}
  /**
  *	Override this function on DiscountCodeEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendDiscountCodeRouter = func(r *[]workspaces.Module2Action) {}
  func GetDiscountCodeModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      DISCOUNT_CODE_ACTION_QUERY,
      DISCOUNT_CODE_ACTION_EXPORT,
      DISCOUNT_CODE_ACTION_GET_ONE,
      DISCOUNT_CODE_ACTION_POST_ONE,
      DISCOUNT_CODE_ACTION_PATCH,
      DISCOUNT_CODE_ACTION_PATCH_BULK,
      DISCOUNT_CODE_ACTION_DELETE,
    }
    // Append user defined functions
    AppendDiscountCodeRouter(&routes)
    return routes
  }
  func CreateDiscountCodeRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetDiscountCodeModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, DiscountCodeEntityJsonSchema, "discount-code-http", "shop")
    workspaces.WriteEntitySchema("DiscountCodeEntity", DiscountCodeEntityJsonSchema, "shop")
    return httpRoutes
  }
var PERM_ROOT_DISCOUNT_CODE_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/discount-code/delete",
  Name: "Delete discount code",
}
var PERM_ROOT_DISCOUNT_CODE_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/discount-code/create",
  Name: "Create discount code",
}
var PERM_ROOT_DISCOUNT_CODE_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/discount-code/update",
  Name: "Update discount code",
}
var PERM_ROOT_DISCOUNT_CODE_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/shop/discount-code/query",
  Name: "Query discount code",
}
var PERM_ROOT_DISCOUNT_CODE = workspaces.PermissionInfo{
  CompleteKey: "root/shop/discount-code/*",
  Name: "Entire discount code actions (*)",
}
var ALL_DISCOUNT_CODE_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_DISCOUNT_CODE_DELETE,
	PERM_ROOT_DISCOUNT_CODE_CREATE,
	PERM_ROOT_DISCOUNT_CODE_UPDATE,
	PERM_ROOT_DISCOUNT_CODE_QUERY,
	PERM_ROOT_DISCOUNT_CODE,
}