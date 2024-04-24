package demo
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
	mocks "github.com/torabian/fireback/modules/demo/mocks/Customer"
)
type CustomerAddress struct {
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
    ZipCode   *string `json:"zipCode" yaml:"zipCode"       `
    // Datenano also has a text representation
    Country   *string `json:"country" yaml:"country"       `
    // Datenano also has a text representation
    Street   *string `json:"street" yaml:"street"       `
    // Datenano also has a text representation
    City   *string `json:"city" yaml:"city"       `
    // Datenano also has a text representation
	LinkedTo *CustomerEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
func ( x * CustomerAddress) RootObjectName() string {
	return "CustomerEntity"
}
type CustomerEntity struct {
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
    FirstName   *string `json:"firstName" yaml:"firstName"       `
    // Datenano also has a text representation
    Avatar   *string `json:"avatar" yaml:"avatar"       `
    // Datenano also has a text representation
    Sex   *string `json:"sex" yaml:"sex"       `
    // Datenano also has a text representation
    SubscriptionTier   *string `json:"subscriptionTier" yaml:"subscriptionTier"       `
    // Datenano also has a text representation
    Birthday   *string `json:"birthday" yaml:"birthday"       `
    // Datenano also has a text representation
    LastName   *string `json:"lastName" yaml:"lastName"       `
    // Datenano also has a text representation
    Address   *  CustomerAddress `json:"address" yaml:"address"    gorm:"foreignKey:LinkerId;references:UniqueId"     `
    // Datenano also has a text representation
    Children []*CustomerEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *CustomerEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var CustomerPreloadRelations []string = []string{}
var CUSTOMER_EVENT_CREATED = "customer.created"
var CUSTOMER_EVENT_UPDATED = "customer.updated"
var CUSTOMER_EVENT_DELETED = "customer.deleted"
var CUSTOMER_EVENTS = []string{
	CUSTOMER_EVENT_CREATED,
	CUSTOMER_EVENT_UPDATED,
	CUSTOMER_EVENT_DELETED,
}
type CustomerFieldMap struct {
		FirstName workspaces.TranslatedString `yaml:"firstName"`
		Avatar workspaces.TranslatedString `yaml:"avatar"`
		Sex workspaces.TranslatedString `yaml:"sex"`
		SubscriptionTier workspaces.TranslatedString `yaml:"subscriptionTier"`
		Birthday workspaces.TranslatedString `yaml:"birthday"`
		LastName workspaces.TranslatedString `yaml:"lastName"`
		Address workspaces.TranslatedString `yaml:"address"`
}
var CustomerEntityMetaConfig map[string]int64 = map[string]int64{
}
var CustomerEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&CustomerEntity{}))
func CustomerAddressActionCreate(
  dto *CustomerAddress ,
  query workspaces.QueryDSL,
) (*CustomerAddress , *workspaces.IError) {
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
func CustomerAddressActionUpdate(
    query workspaces.QueryDSL,
    dto *CustomerAddress,
) (*CustomerAddress, *workspaces.IError) {
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
func CustomerAddressActionGetOne(
    query workspaces.QueryDSL,
) (*CustomerAddress , *workspaces.IError) {
    refl := reflect.ValueOf(&CustomerAddress {})
    item, err := workspaces.GetOneEntity[CustomerAddress ](query, refl)
    return item, err
}
func entityCustomerFormatter(dto *CustomerEntity, query workspaces.QueryDSL) {
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
func CustomerMockEntity() *CustomerEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &CustomerEntity{
      FirstName : &stringHolder,
      Avatar : &stringHolder,
      Sex : &stringHolder,
      SubscriptionTier : &stringHolder,
      Birthday : &stringHolder,
      LastName : &stringHolder,
	}
	return entity
}
func CustomerActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := CustomerMockEntity()
		_, err := CustomerActionCreate(entity, query)
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
  func CustomerActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*CustomerEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &CustomerEntity{
          FirstName: &tildaRef,
          Avatar: &tildaRef,
          Sex: &tildaRef,
          SubscriptionTier: &tildaRef,
          Birthday: &tildaRef,
          LastName: &tildaRef,
          Address: &CustomerAddress{},
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
  func CustomerAssociationCreate(dto *CustomerEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func CustomerRelationContentCreate(dto *CustomerEntity, query workspaces.QueryDSL) error {
return nil
}
func CustomerRelationContentUpdate(dto *CustomerEntity, query workspaces.QueryDSL) error {
	return nil
}
func CustomerPolyglotCreateHandler(dto *CustomerEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func CustomerValidator(dto *CustomerEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func CustomerEntityPreSanitize(dto *CustomerEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func CustomerEntityBeforeCreateAppend(dto *CustomerEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    CustomerRecursiveAddUniqueId(dto, query)
  }
  func CustomerRecursiveAddUniqueId(dto *CustomerEntity, query workspaces.QueryDSL) {
        if dto.Address != nil {
          dto.Address.UniqueId = workspaces.UUID()
        }
  }
func CustomerActionBatchCreateFn(dtos []*CustomerEntity, query workspaces.QueryDSL) ([]*CustomerEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*CustomerEntity{}
		for _, item := range dtos {
			s, err := CustomerActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func CustomerDeleteEntireChildren(query workspaces.QueryDSL, dto *CustomerEntity) (*workspaces.IError) {
  if dto.Address != nil {
    q := query.Tx.
      Model(&dto.Address).
      Where(&CustomerAddress{LinkerId: &dto.UniqueId }).
      Delete(&CustomerAddress{})
    err := q.Error
    if err != nil {
      return workspaces.GormErrorToIError(err)
    }
  }
  return nil
}
func CustomerActionCreateFn(dto *CustomerEntity, query workspaces.QueryDSL) (*CustomerEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := CustomerValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	CustomerEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	CustomerEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	CustomerPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	CustomerRelationContentCreate(dto, query)
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
	CustomerAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(CUSTOMER_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&CustomerEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func CustomerActionGetOne(query workspaces.QueryDSL) (*CustomerEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&CustomerEntity{})
    item, err := workspaces.GetOneEntity[CustomerEntity](query, refl)
    entityCustomerFormatter(item, query)
    return item, err
  }
  func CustomerActionQuery(query workspaces.QueryDSL) ([]*CustomerEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&CustomerEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[CustomerEntity](query, refl)
    for _, item := range items {
      entityCustomerFormatter(item, query)
    }
    return items, meta, err
  }
  func CustomerUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *CustomerEntity) (*CustomerEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = CUSTOMER_EVENT_UPDATED
    CustomerEntityPreSanitize(fields, query)
    var item CustomerEntity
    q := dbref.
      Where(&CustomerEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    CustomerRelationContentUpdate(fields, query)
    CustomerPolyglotCreateHandler(fields, query)
    if ero := CustomerDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
        if fields.Address != nil {
          linkerId := uniqueId
          q := dbref.
            Model(&item.Address).
            Where(&CustomerAddress{LinkerId: &linkerId}).
            UpdateColumns(fields.Address)
          err := q.Error
          if err != nil {
            return &item, workspaces.GormErrorToIError(err)
          }
          if q.RowsAffected == 0 {
            fields.Address.UniqueId = workspaces.UUID()
            fields.Address.LinkerId = &linkerId
            err := dbref.
              Model(&item.Address).Create(fields.Address).Error
            if err != nil {
              return &item, workspaces.GormErrorToIError(err)
            }
          }
        }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&CustomerEntity{UniqueId: uniqueId}).
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
  func CustomerActionUpdateFn(query workspaces.QueryDSL, fields *CustomerEntity) (*CustomerEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := CustomerValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // CustomerRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *CustomerEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = CustomerUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return CustomerUpdateExec(dbref, query, fields)
    }
  }
var CustomerWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire customers ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CUSTOMER_DELETE},
    })
		count, _ := CustomerActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func CustomerActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&CustomerEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_CUSTOMER_DELETE}
	return workspaces.RemoveEntity[CustomerEntity](query, refl)
}
func CustomerActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
			{
				subCount, subErr := workspaces.WipeCleanEntity[CustomerAddress]()
				if (subErr != nil) {
					fmt.Println("Error while wiping 'CustomerAddress'", subErr)
					return count, subErr
				} else {
					count += subCount
				}
			}
	{
		subCount, subErr := workspaces.WipeCleanEntity[CustomerEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'CustomerEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func CustomerActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[CustomerEntity]) (
    *workspaces.BulkRecordRequest[CustomerEntity], *workspaces.IError,
  ) {
    result := []*CustomerEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := CustomerActionUpdate(query, record)
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
func (x *CustomerEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var CustomerEntityMeta = workspaces.TableMetaData{
	EntityName:    "Customer",
	ExportKey:    "customers",
	TableNameInDb: "fb_customer_entities",
	EntityObject:  &CustomerEntity{},
	ExportStream: CustomerActionExportT,
	ImportQuery: CustomerActionImport,
}
func CustomerActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[CustomerEntity](query, CustomerActionQuery, CustomerPreloadRelations)
}
func CustomerActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[CustomerEntity](query, CustomerActionQuery, CustomerPreloadRelations)
}
func CustomerActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content CustomerEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := CustomerActionCreate(&content, query)
	return err
}
var CustomerCommonCliFlags = []cli.Flag{
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
      Name:     "first-name",
      Required: false,
      Usage:    "firstName",
    },
    &cli.StringFlag{
      Name:     "avatar",
      Required: false,
      Usage:    "avatar",
    },
    &cli.StringFlag{
      Name:     "sex",
      Required: false,
      Usage:    "sex",
    },
    &cli.StringFlag{
      Name:     "subscription-tier",
      Required: false,
      Usage:    "subscriptionTier",
    },
    &cli.StringFlag{
      Name:     "birthday",
      Required: false,
      Usage:    "birthday",
    },
    &cli.StringFlag{
      Name:     "last-name",
      Required: false,
      Usage:    "lastName",
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
      Name:     "zip-code",
      Required: false,
      Usage:    "zipCode",
    },
    &cli.StringFlag{
      Name:     "country",
      Required: false,
      Usage:    "country",
    },
    &cli.StringFlag{
      Name:     "street",
      Required: false,
      Usage:    "street",
    },
    &cli.StringFlag{
      Name:     "city",
      Required: false,
      Usage:    "city",
    },
}
var CustomerCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "firstName",
		StructField:     "FirstName",
		Required: false,
		Usage:    "firstName",
		Type: "string",
	},
	{
		Name:     "avatar",
		StructField:     "Avatar",
		Required: false,
		Usage:    "avatar",
		Type: "string",
	},
	{
		Name:     "sex",
		StructField:     "Sex",
		Required: false,
		Usage:    "sex",
		Type: "string",
	},
	{
		Name:     "subscriptionTier",
		StructField:     "SubscriptionTier",
		Required: false,
		Usage:    "subscriptionTier",
		Type: "string",
	},
	{
		Name:     "birthday",
		StructField:     "Birthday",
		Required: false,
		Usage:    "birthday",
		Type: "string",
	},
	{
		Name:     "lastName",
		StructField:     "LastName",
		Required: false,
		Usage:    "lastName",
		Type: "string",
	},
}
var CustomerCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "first-name",
      Required: false,
      Usage:    "firstName",
    },
    &cli.StringFlag{
      Name:     "avatar",
      Required: false,
      Usage:    "avatar",
    },
    &cli.StringFlag{
      Name:     "sex",
      Required: false,
      Usage:    "sex",
    },
    &cli.StringFlag{
      Name:     "subscription-tier",
      Required: false,
      Usage:    "subscriptionTier",
    },
    &cli.StringFlag{
      Name:     "birthday",
      Required: false,
      Usage:    "birthday",
    },
    &cli.StringFlag{
      Name:     "last-name",
      Required: false,
      Usage:    "lastName",
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
      Name:     "zip-code",
      Required: false,
      Usage:    "zipCode",
    },
    &cli.StringFlag{
      Name:     "country",
      Required: false,
      Usage:    "country",
    },
    &cli.StringFlag{
      Name:     "street",
      Required: false,
      Usage:    "street",
    },
    &cli.StringFlag{
      Name:     "city",
      Required: false,
      Usage:    "city",
    },
}
  var CustomerCreateCmd cli.Command = CUSTOMER_ACTION_POST_ONE.ToCli()
  var CustomerCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CUSTOMER_CREATE},
      })
      entity := &CustomerEntity{}
      for _, item := range CustomerCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := CustomerActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var CustomerUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: CustomerCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CUSTOMER_UPDATE},
      })
      entity := CastCustomerFromCli(c)
      if entity, err := CustomerActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* CustomerEntity) FromCli(c *cli.Context) *CustomerEntity {
	return CastCustomerFromCli(c)
}
func CastCustomerFromCli (c *cli.Context) *CustomerEntity {
	template := &CustomerEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("first-name") {
        value := c.String("first-name")
        template.FirstName = &value
      }
      if c.IsSet("avatar") {
        value := c.String("avatar")
        template.Avatar = &value
      }
      if c.IsSet("sex") {
        value := c.String("sex")
        template.Sex = &value
      }
      if c.IsSet("subscription-tier") {
        value := c.String("subscription-tier")
        template.SubscriptionTier = &value
      }
      if c.IsSet("birthday") {
        value := c.String("birthday")
        template.Birthday = &value
      }
      if c.IsSet("last-name") {
        value := c.String("last-name")
        template.LastName = &value
      }
	return template
}
  func CustomerSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      CustomerActionCreate,
      reflect.ValueOf(&CustomerEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func CustomerImportMocks() {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      CustomerActionCreate,
      reflect.ValueOf(&CustomerEntity{}).Elem(),
      &mocks.ViewsFs,
      []string{},
      false,
    )
  }
  func CustomerWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := CustomerActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "Customer", result)
    }
  }
var CustomerImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CUSTOMER_CREATE},
      })
			CustomerActionSeeder(query, c.Int("count"))
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
				Value: "customer-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CUSTOMER_CREATE},
      })
			CustomerActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "customer-seeder-customer.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of customers, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]CustomerEntity{}
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
					CustomerActionCreate,
					reflect.ValueOf(&CustomerEntity{}).Elem(),
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
			CustomerCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				CustomerActionCreate,
				reflect.ValueOf(&CustomerEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CUSTOMER_CREATE},
				},
        func() CustomerEntity {
					v := CastCustomerFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var CustomerCliCommands []cli.Command = []cli.Command{
      CUSTOMER_ACTION_QUERY.ToCli(),
      CUSTOMER_ACTION_TABLE.ToCli(),
      CustomerCreateCmd,
      CustomerUpdateCmd,
      CustomerCreateInteractiveCmd,
      CustomerWipeCmd,
      workspaces.GetCommonRemoveQuery(reflect.ValueOf(&CustomerEntity{}).Elem(), CustomerActionRemove),
  }
  func CustomerCliFn() cli.Command {
    CustomerCliCommands = append(CustomerCliCommands, CustomerImportExportCommands...)
    return cli.Command{
      Name:        "customer",
      Description: "Customers module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: CustomerCliCommands,
    }
  }
var CUSTOMER_ACTION_TABLE = workspaces.Module2Action{
  Name:    "table",
  ActionName: "table",
  ActionAliases: []string{"t"},
  Flags:  workspaces.CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: CustomerActionQuery,
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    workspaces.CommonCliTableCmd2(c,
      CustomerActionQuery,
      security,
      reflect.ValueOf(&CustomerEntity{}).Elem(),
    )
    return nil
  },
}
var CUSTOMER_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/customers",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CUSTOMER_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, CustomerActionQuery)
    },
  },
  Format: "QUERY",
  Action: CustomerActionQuery,
  ResponseEntity: &[]CustomerEntity{},
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			CustomerActionQuery,
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
var CUSTOMER_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/customers/export",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CUSTOMER_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, CustomerActionExport)
    },
  },
  Format: "QUERY",
  Action: CustomerActionExport,
  ResponseEntity: &[]CustomerEntity{},
}
var CUSTOMER_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/customer/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CUSTOMER_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, CustomerActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: CustomerActionGetOne,
  ResponseEntity: &CustomerEntity{},
}
var CUSTOMER_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new customer",
  Flags: CustomerCommonCliFlags,
  Method: "POST",
  Url:    "/customer",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CUSTOMER_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, CustomerActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, CustomerActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: CustomerActionCreate,
  Format: "POST_ONE",
  RequestEntity: &CustomerEntity{},
  ResponseEntity: &CustomerEntity{},
}
var CUSTOMER_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: CustomerCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/customer",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CUSTOMER_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, CustomerActionUpdate)
    },
  },
  Action: CustomerActionUpdate,
  RequestEntity: &CustomerEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &CustomerEntity{},
}
var CUSTOMER_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/customers",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CUSTOMER_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, CustomerActionBulkUpdate)
    },
  },
  Action: CustomerActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[CustomerEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[CustomerEntity]{},
}
var CUSTOMER_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/customer",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CUSTOMER_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, CustomerActionRemove)
    },
  },
  Action: CustomerActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &CustomerEntity{},
}
    var CUSTOMER_ADDRESS_ACTION_PATCH = workspaces.Module2Action{
      Method: "PATCH",
      Url:    "/customer/:linkerId/address/:uniqueId",
      SecurityModel: &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CUSTOMER_UPDATE},
      },
      Handlers: []gin.HandlerFunc{
        func (
          c *gin.Context,
        ) {
          workspaces.HttpUpdateEntity(c, CustomerAddressActionUpdate)
        },
      },
      Action: CustomerAddressActionUpdate,
      Format: "PATCH_ONE",
      RequestEntity: &CustomerAddress{},
      ResponseEntity: &CustomerAddress{},
    }
    var CUSTOMER_ADDRESS_ACTION_GET = workspaces.Module2Action {
      Method: "GET",
      Url:    "/customer/address/:linkerId/:uniqueId",
      SecurityModel: &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CUSTOMER_QUERY},
      },
      Handlers: []gin.HandlerFunc{
        func (
          c *gin.Context,
        ) {
          workspaces.HttpGetEntity(c, CustomerAddressActionGetOne)
        },
      },
      Action: CustomerAddressActionGetOne,
      Format: "GET_ONE",
      ResponseEntity: &CustomerAddress{},
    }
    var CUSTOMER_ADDRESS_ACTION_POST = workspaces.Module2Action{
      Method: "POST",
      Url:    "/customer/:linkerId/address",
      SecurityModel: &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_CUSTOMER_CREATE},
      },
      Handlers: []gin.HandlerFunc{
        func (
          c *gin.Context,
        ) {
          workspaces.HttpPostEntity(c, CustomerAddressActionCreate)
        },
      },
      Action: CustomerAddressActionCreate,
      Format: "POST_ONE",
      RequestEntity: &CustomerAddress{},
      ResponseEntity: &CustomerAddress{},
    }
  /**
  *	Override this function on CustomerEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendCustomerRouter = func(r *[]workspaces.Module2Action) {}
  func GetCustomerModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      CUSTOMER_ACTION_QUERY,
      CUSTOMER_ACTION_EXPORT,
      CUSTOMER_ACTION_GET_ONE,
      CUSTOMER_ACTION_POST_ONE,
      CUSTOMER_ACTION_PATCH,
      CUSTOMER_ACTION_PATCH_BULK,
      CUSTOMER_ACTION_DELETE,
          CUSTOMER_ADDRESS_ACTION_PATCH,
          CUSTOMER_ADDRESS_ACTION_GET,
          CUSTOMER_ADDRESS_ACTION_POST,
    }
    // Append user defined functions
    AppendCustomerRouter(&routes)
    return routes
  }
  func CreateCustomerRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetCustomerModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, CustomerEntityJsonSchema, "customer-http", "demo")
    workspaces.WriteEntitySchema("CustomerEntity", CustomerEntityJsonSchema, "demo")
    return httpRoutes
  }
var PERM_ROOT_CUSTOMER_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/demo/customer/delete",
  Name: "Delete customer",
}
var PERM_ROOT_CUSTOMER_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/demo/customer/create",
  Name: "Create customer",
}
var PERM_ROOT_CUSTOMER_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/demo/customer/update",
  Name: "Update customer",
}
var PERM_ROOT_CUSTOMER_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/demo/customer/query",
  Name: "Query customer",
}
var PERM_ROOT_CUSTOMER = workspaces.PermissionInfo{
  CompleteKey: "root/demo/customer/*",
  Name: "Entire customer actions (*)",
}
var ALL_CUSTOMER_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_CUSTOMER_DELETE,
	PERM_ROOT_CUSTOMER_CREATE,
	PERM_ROOT_CUSTOMER_UPDATE,
	PERM_ROOT_CUSTOMER_QUERY,
	PERM_ROOT_CUSTOMER,
}