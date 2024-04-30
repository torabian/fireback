package worldtimezone
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
	seeders "github.com/torabian/fireback/modules/worldtimezone/seeders/TimezoneGroup"
	metas "github.com/torabian/fireback/modules/worldtimezone/metas"
)
type TimezoneGroupUtcItems struct {
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
    Name   *string `json:"name" yaml:"name"  validate:"required"        translate:"true" `
    // Datenano also has a text representation
	LinkedTo *TimezoneGroupEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
func ( x * TimezoneGroupUtcItems) RootObjectName() string {
	return "TimezoneGroupEntity"
}
type TimezoneGroupEntity struct {
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
    Value   *string `json:"value" yaml:"value"        translate:"true" `
    // Datenano also has a text representation
    Abbr   *string `json:"abbr" yaml:"abbr"       `
    // Datenano also has a text representation
    Offset   *int64 `json:"offset" yaml:"offset"       `
    // Datenano also has a text representation
    Isdst   *bool `json:"isdst" yaml:"isdst"       `
    // Datenano also has a text representation
    Text   *string `json:"text" yaml:"text"        translate:"true" `
    // Datenano also has a text representation
    UtcItems   []*  TimezoneGroupUtcItems `json:"utcItems" yaml:"utcItems"    gorm:"foreignKey:LinkerId;references:UniqueId"     `
    // Datenano also has a text representation
    Translations     []*TimezoneGroupEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    Children []*TimezoneGroupEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *TimezoneGroupEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var TimezoneGroupPreloadRelations []string = []string{}
var TIMEZONE_GROUP_EVENT_CREATED = "timezoneGroup.created"
var TIMEZONE_GROUP_EVENT_UPDATED = "timezoneGroup.updated"
var TIMEZONE_GROUP_EVENT_DELETED = "timezoneGroup.deleted"
var TIMEZONE_GROUP_EVENTS = []string{
	TIMEZONE_GROUP_EVENT_CREATED,
	TIMEZONE_GROUP_EVENT_UPDATED,
	TIMEZONE_GROUP_EVENT_DELETED,
}
type TimezoneGroupFieldMap struct {
		Value workspaces.TranslatedString `yaml:"value"`
		Abbr workspaces.TranslatedString `yaml:"abbr"`
		Offset workspaces.TranslatedString `yaml:"offset"`
		Isdst workspaces.TranslatedString `yaml:"isdst"`
		Text workspaces.TranslatedString `yaml:"text"`
		UtcItems workspaces.TranslatedString `yaml:"utcItems"`
}
var TimezoneGroupEntityMetaConfig map[string]int64 = map[string]int64{
}
var TimezoneGroupEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&TimezoneGroupEntity{}))
  type TimezoneGroupEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        Value string `yaml:"value" json:"value"`
        Text string `yaml:"text" json:"text"`
  }
func TimezoneGroupUtcItemsActionCreate(
  dto *TimezoneGroupUtcItems ,
  query workspaces.QueryDSL,
) (*TimezoneGroupUtcItems , *workspaces.IError) {
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
func TimezoneGroupUtcItemsActionUpdate(
    query workspaces.QueryDSL,
    dto *TimezoneGroupUtcItems,
) (*TimezoneGroupUtcItems, *workspaces.IError) {
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
func TimezoneGroupUtcItemsActionGetOne(
    query workspaces.QueryDSL,
) (*TimezoneGroupUtcItems , *workspaces.IError) {
    refl := reflect.ValueOf(&TimezoneGroupUtcItems {})
    item, err := workspaces.GetOneEntity[TimezoneGroupUtcItems ](query, refl)
    return item, err
}
func entityTimezoneGroupFormatter(dto *TimezoneGroupEntity, query workspaces.QueryDSL) {
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
func TimezoneGroupMockEntity() *TimezoneGroupEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &TimezoneGroupEntity{
      Value : &stringHolder,
      Abbr : &stringHolder,
      Offset : &int64Holder,
      Text : &stringHolder,
	}
	return entity
}
func TimezoneGroupActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := TimezoneGroupMockEntity()
		_, err := TimezoneGroupActionCreate(entity, query)
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
    func (x*TimezoneGroupEntity) GetValueTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Value
          }
        }
      }
      return ""
    }
    func (x*TimezoneGroupEntity) GetTextTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Text
          }
        }
      }
      return ""
    }
  func TimezoneGroupActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*TimezoneGroupEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &TimezoneGroupEntity{
          Value: &tildaRef,
          Abbr: &tildaRef,
          Text: &tildaRef,
          UtcItems: []*TimezoneGroupUtcItems{{}},
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
  func TimezoneGroupAssociationCreate(dto *TimezoneGroupEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func TimezoneGroupRelationContentCreate(dto *TimezoneGroupEntity, query workspaces.QueryDSL) error {
return nil
}
func TimezoneGroupRelationContentUpdate(dto *TimezoneGroupEntity, query workspaces.QueryDSL) error {
	return nil
}
func TimezoneGroupPolyglotCreateHandler(dto *TimezoneGroupEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &TimezoneGroupEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func TimezoneGroupValidator(dto *TimezoneGroupEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
        if dto != nil && dto.UtcItems != nil {
          workspaces.AppendSliceErrors(dto.UtcItems, isPatch, "utcItems", err)
        }
    return err
  }
func TimezoneGroupEntityPreSanitize(dto *TimezoneGroupEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func TimezoneGroupEntityBeforeCreateAppend(dto *TimezoneGroupEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    TimezoneGroupRecursiveAddUniqueId(dto, query)
  }
  func TimezoneGroupRecursiveAddUniqueId(dto *TimezoneGroupEntity, query workspaces.QueryDSL) {
      if dto.UtcItems != nil && len(dto.UtcItems) > 0 {
        for index0 := range dto.UtcItems {
          if (dto.UtcItems[index0].UniqueId == "") {
            dto.UtcItems[index0].UniqueId = workspaces.UUID()
          }
        }
    }
  }
func TimezoneGroupActionBatchCreateFn(dtos []*TimezoneGroupEntity, query workspaces.QueryDSL) ([]*TimezoneGroupEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*TimezoneGroupEntity{}
		for _, item := range dtos {
			s, err := TimezoneGroupActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func TimezoneGroupDeleteEntireChildren(query workspaces.QueryDSL, dto *TimezoneGroupEntity) (*workspaces.IError) {
  if dto.UtcItems != nil {
    q := query.Tx.
      Model(&dto.UtcItems).
      Where(&TimezoneGroupUtcItems{LinkerId: &dto.UniqueId }).
      Delete(&TimezoneGroupUtcItems{})
    err := q.Error
    if err != nil {
      return workspaces.GormErrorToIError(err)
    }
  }
  return nil
}
func TimezoneGroupActionCreateFn(dto *TimezoneGroupEntity, query workspaces.QueryDSL) (*TimezoneGroupEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := TimezoneGroupValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	TimezoneGroupEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	TimezoneGroupEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	TimezoneGroupPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	TimezoneGroupRelationContentCreate(dto, query)
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
	TimezoneGroupAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(TIMEZONE_GROUP_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&TimezoneGroupEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func TimezoneGroupActionGetOne(query workspaces.QueryDSL) (*TimezoneGroupEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&TimezoneGroupEntity{})
    item, err := workspaces.GetOneEntity[TimezoneGroupEntity](query, refl)
    entityTimezoneGroupFormatter(item, query)
    return item, err
  }
  func TimezoneGroupActionQuery(query workspaces.QueryDSL) ([]*TimezoneGroupEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&TimezoneGroupEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[TimezoneGroupEntity](query, refl)
    for _, item := range items {
      entityTimezoneGroupFormatter(item, query)
    }
    return items, meta, err
  }
  func TimezoneGroupUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *TimezoneGroupEntity) (*TimezoneGroupEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = TIMEZONE_GROUP_EVENT_UPDATED
    TimezoneGroupEntityPreSanitize(fields, query)
    var item TimezoneGroupEntity
    q := dbref.
      Where(&TimezoneGroupEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    TimezoneGroupRelationContentUpdate(fields, query)
    TimezoneGroupPolyglotCreateHandler(fields, query)
    if ero := TimezoneGroupDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
      if fields.UtcItems != nil {
       linkerId := uniqueId;
        dbref.
          Where(&TimezoneGroupUtcItems {LinkerId: &linkerId}).
          Delete(&TimezoneGroupUtcItems {})
        for _, newItem := range fields.UtcItems {
          newItem.UniqueId = workspaces.UUID()
          newItem.LinkerId = &linkerId
          dbref.Create(&newItem)
        }
      }
    err = dbref.
      Preload(clause.Associations).
      Where(&TimezoneGroupEntity{UniqueId: uniqueId}).
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
  func TimezoneGroupActionUpdateFn(query workspaces.QueryDSL, fields *TimezoneGroupEntity) (*TimezoneGroupEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := TimezoneGroupValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // TimezoneGroupRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *TimezoneGroupEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = TimezoneGroupUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return TimezoneGroupUpdateExec(dbref, query, fields)
    }
  }
var TimezoneGroupWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire timezonegroups ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_TIMEZONE_GROUP_DELETE},
    })
		count, _ := TimezoneGroupActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func TimezoneGroupActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&TimezoneGroupEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_TIMEZONE_GROUP_DELETE}
	return workspaces.RemoveEntity[TimezoneGroupEntity](query, refl)
}
func TimezoneGroupActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
			{
				subCount, subErr := workspaces.WipeCleanEntity[TimezoneGroupUtcItems]()
				if (subErr != nil) {
					fmt.Println("Error while wiping 'TimezoneGroupUtcItems'", subErr)
					return count, subErr
				} else {
					count += subCount
				}
			}
	{
		subCount, subErr := workspaces.WipeCleanEntity[TimezoneGroupEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'TimezoneGroupEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func TimezoneGroupActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[TimezoneGroupEntity]) (
    *workspaces.BulkRecordRequest[TimezoneGroupEntity], *workspaces.IError,
  ) {
    result := []*TimezoneGroupEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := TimezoneGroupActionUpdate(query, record)
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
func (x *TimezoneGroupEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var TimezoneGroupEntityMeta = workspaces.TableMetaData{
	EntityName:    "TimezoneGroup",
	ExportKey:    "timezone-groups",
	TableNameInDb: "fb_timezone-group_entities",
	EntityObject:  &TimezoneGroupEntity{},
	ExportStream: TimezoneGroupActionExportT,
	ImportQuery: TimezoneGroupActionImport,
}
func TimezoneGroupActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[TimezoneGroupEntity](query, TimezoneGroupActionQuery, TimezoneGroupPreloadRelations)
}
func TimezoneGroupActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[TimezoneGroupEntity](query, TimezoneGroupActionQuery, TimezoneGroupPreloadRelations)
}
func TimezoneGroupActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content TimezoneGroupEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := TimezoneGroupActionCreate(&content, query)
	return err
}
var TimezoneGroupCommonCliFlags = []cli.Flag{
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
      Name:     "value",
      Required: false,
      Usage:    "value",
    },
    &cli.StringFlag{
      Name:     "abbr",
      Required: false,
      Usage:    "abbr",
    },
    &cli.Int64Flag{
      Name:     "offset",
      Required: false,
      Usage:    "offset",
    },
    &cli.BoolFlag{
      Name:     "isdst",
      Required: false,
      Usage:    "isdst",
    },
    &cli.StringFlag{
      Name:     "text",
      Required: false,
      Usage:    "text",
    },
    &cli.StringSliceFlag{
      Name:     "utc-items",
      Required: false,
      Usage:    "utcItems",
    },
}
var TimezoneGroupCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "value",
		StructField:     "Value",
		Required: false,
		Usage:    "value",
		Type: "string",
	},
	{
		Name:     "abbr",
		StructField:     "Abbr",
		Required: false,
		Usage:    "abbr",
		Type: "string",
	},
	{
		Name:     "offset",
		StructField:     "Offset",
		Required: false,
		Usage:    "offset",
		Type: "int64",
	},
	{
		Name:     "isdst",
		StructField:     "Isdst",
		Required: false,
		Usage:    "isdst",
		Type: "bool",
	},
	{
		Name:     "text",
		StructField:     "Text",
		Required: false,
		Usage:    "text",
		Type: "string",
	},
}
var TimezoneGroupCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "value",
      Required: false,
      Usage:    "value",
    },
    &cli.StringFlag{
      Name:     "abbr",
      Required: false,
      Usage:    "abbr",
    },
    &cli.Int64Flag{
      Name:     "offset",
      Required: false,
      Usage:    "offset",
    },
    &cli.BoolFlag{
      Name:     "isdst",
      Required: false,
      Usage:    "isdst",
    },
    &cli.StringFlag{
      Name:     "text",
      Required: false,
      Usage:    "text",
    },
    &cli.StringSliceFlag{
      Name:     "utc-items",
      Required: false,
      Usage:    "utcItems",
    },
}
  var TimezoneGroupCreateCmd cli.Command = TIMEZONE_GROUP_ACTION_POST_ONE.ToCli()
  var TimezoneGroupCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_TIMEZONE_GROUP_CREATE},
      })
      entity := &TimezoneGroupEntity{}
      for _, item := range TimezoneGroupCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := TimezoneGroupActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var TimezoneGroupUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: TimezoneGroupCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_TIMEZONE_GROUP_UPDATE},
      })
      entity := CastTimezoneGroupFromCli(c)
      if entity, err := TimezoneGroupActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* TimezoneGroupEntity) FromCli(c *cli.Context) *TimezoneGroupEntity {
	return CastTimezoneGroupFromCli(c)
}
func CastTimezoneGroupFromCli (c *cli.Context) *TimezoneGroupEntity {
	template := &TimezoneGroupEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("value") {
        value := c.String("value")
        template.Value = &value
      }
      if c.IsSet("abbr") {
        value := c.String("abbr")
        template.Abbr = &value
      }
      if c.IsSet("text") {
        value := c.String("text")
        template.Text = &value
      }
	return template
}
  func TimezoneGroupSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      TimezoneGroupActionCreate,
      reflect.ValueOf(&TimezoneGroupEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func TimezoneGroupSyncSeeders() {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{WorkspaceId: workspaces.USER_SYSTEM},
      TimezoneGroupActionCreate,
      reflect.ValueOf(&TimezoneGroupEntity{}).Elem(),
      &seeders.ViewsFs,
      []string{},
      true,
    )
  }
  func TimezoneGroupWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := TimezoneGroupActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "TimezoneGroup", result)
    }
  }
var TimezoneGroupImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_TIMEZONE_GROUP_CREATE},
      })
			TimezoneGroupActionSeeder(query, c.Int("count"))
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
				Value: "timezone-group-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_TIMEZONE_GROUP_CREATE},
      })
			TimezoneGroupActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "timezone-group-seeder-timezone-group.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of timezone-groups, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]TimezoneGroupEntity{}
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
				TimezoneGroupActionCreate,
				reflect.ValueOf(&TimezoneGroupEntity{}).Elem(),
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
				TimezoneGroupActionQuery,
				reflect.ValueOf(&TimezoneGroupEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"TimezoneGroupFieldMap.yml",
				TimezoneGroupPreloadRelations,
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
			TimezoneGroupCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				TimezoneGroupActionCreate,
				reflect.ValueOf(&TimezoneGroupEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_TIMEZONE_GROUP_CREATE},
				},
        func() TimezoneGroupEntity {
					v := CastTimezoneGroupFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var TimezoneGroupCliCommands []cli.Command = []cli.Command{
      TIMEZONE_GROUP_ACTION_QUERY.ToCli(),
      TIMEZONE_GROUP_ACTION_TABLE.ToCli(),
      TimezoneGroupCreateCmd,
      TimezoneGroupUpdateCmd,
      TimezoneGroupCreateInteractiveCmd,
      TimezoneGroupWipeCmd,
      workspaces.GetCommonRemoveQuery(reflect.ValueOf(&TimezoneGroupEntity{}).Elem(), TimezoneGroupActionRemove),
  }
  func TimezoneGroupCliFn() cli.Command {
    TimezoneGroupCliCommands = append(TimezoneGroupCliCommands, TimezoneGroupImportExportCommands...)
    return cli.Command{
      Name:        "tz",
      Description: "TimezoneGroups module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: TimezoneGroupCliCommands,
    }
  }
var TIMEZONE_GROUP_ACTION_TABLE = workspaces.Module2Action{
  Name:    "table",
  ActionName: "table",
  ActionAliases: []string{"t"},
  Flags:  workspaces.CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: TimezoneGroupActionQuery,
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    workspaces.CommonCliTableCmd2(c,
      TimezoneGroupActionQuery,
      security,
      reflect.ValueOf(&TimezoneGroupEntity{}).Elem(),
    )
    return nil
  },
}
var TIMEZONE_GROUP_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/timezone-groups",
  SecurityModel: &workspaces.SecurityModel{
  },
  Group: "timezoneGroup",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, TimezoneGroupActionQuery)
    },
  },
  Format: "QUERY",
  Action: TimezoneGroupActionQuery,
  ResponseEntity: &[]TimezoneGroupEntity{},
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			TimezoneGroupActionQuery,
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
var TIMEZONE_GROUP_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/timezone-groups/export",
  SecurityModel: &workspaces.SecurityModel{
  },
  Group: "timezoneGroup",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, TimezoneGroupActionExport)
    },
  },
  Format: "QUERY",
  Action: TimezoneGroupActionExport,
  ResponseEntity: &[]TimezoneGroupEntity{},
}
var TIMEZONE_GROUP_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/timezone-group/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
  },
  Group: "timezoneGroup",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, TimezoneGroupActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: TimezoneGroupActionGetOne,
  ResponseEntity: &TimezoneGroupEntity{},
}
var TIMEZONE_GROUP_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new timezoneGroup",
  Flags: TimezoneGroupCommonCliFlags,
  Method: "POST",
  Url:    "/timezone-group",
  SecurityModel: &workspaces.SecurityModel{
  },
  Group: "timezoneGroup",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, TimezoneGroupActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, TimezoneGroupActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: TimezoneGroupActionCreate,
  Format: "POST_ONE",
  RequestEntity: &TimezoneGroupEntity{},
  ResponseEntity: &TimezoneGroupEntity{},
}
var TIMEZONE_GROUP_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: TimezoneGroupCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/timezone-group",
  SecurityModel: &workspaces.SecurityModel{
  },
  Group: "timezoneGroup",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, TimezoneGroupActionUpdate)
    },
  },
  Action: TimezoneGroupActionUpdate,
  RequestEntity: &TimezoneGroupEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &TimezoneGroupEntity{},
}
var TIMEZONE_GROUP_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/timezone-groups",
  SecurityModel: &workspaces.SecurityModel{
  },
  Group: "timezoneGroup",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, TimezoneGroupActionBulkUpdate)
    },
  },
  Action: TimezoneGroupActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[TimezoneGroupEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[TimezoneGroupEntity]{},
}
var TIMEZONE_GROUP_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/timezone-group",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
  },
  Group: "timezoneGroup",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, TimezoneGroupActionRemove)
    },
  },
  Action: TimezoneGroupActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &TimezoneGroupEntity{},
}
    var TIMEZONE_GROUP_UTC_ITEMS_ACTION_PATCH = workspaces.Module2Action{
      Method: "PATCH",
      Url:    "/timezone-group/:linkerId/utc_items/:uniqueId",
      SecurityModel: &workspaces.SecurityModel{
      },
      Group: "timezoneGroup",
      Handlers: []gin.HandlerFunc{
        func (
          c *gin.Context,
        ) {
          workspaces.HttpUpdateEntity(c, TimezoneGroupUtcItemsActionUpdate)
        },
      },
      Action: TimezoneGroupUtcItemsActionUpdate,
      Format: "PATCH_ONE",
      RequestEntity: &TimezoneGroupUtcItems{},
      ResponseEntity: &TimezoneGroupUtcItems{},
    }
    var TIMEZONE_GROUP_UTC_ITEMS_ACTION_GET = workspaces.Module2Action {
      Method: "GET",
      Url:    "/timezone-group/utc_items/:linkerId/:uniqueId",
      SecurityModel: &workspaces.SecurityModel{
      },
      Group: "timezoneGroup",
      Handlers: []gin.HandlerFunc{
        func (
          c *gin.Context,
        ) {
          workspaces.HttpGetEntity(c, TimezoneGroupUtcItemsActionGetOne)
        },
      },
      Action: TimezoneGroupUtcItemsActionGetOne,
      Format: "GET_ONE",
      ResponseEntity: &TimezoneGroupUtcItems{},
    }
    var TIMEZONE_GROUP_UTC_ITEMS_ACTION_POST = workspaces.Module2Action{
      Method: "POST",
      Url:    "/timezone-group/:linkerId/utc_items",
      SecurityModel: &workspaces.SecurityModel{
      },
      Group: "timezoneGroup",
      Handlers: []gin.HandlerFunc{
        func (
          c *gin.Context,
        ) {
          workspaces.HttpPostEntity(c, TimezoneGroupUtcItemsActionCreate)
        },
      },
      Action: TimezoneGroupUtcItemsActionCreate,
      Format: "POST_ONE",
      RequestEntity: &TimezoneGroupUtcItems{},
      ResponseEntity: &TimezoneGroupUtcItems{},
    }
  /**
  *	Override this function on TimezoneGroupEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendTimezoneGroupRouter = func(r *[]workspaces.Module2Action) {}
  func GetTimezoneGroupModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      TIMEZONE_GROUP_ACTION_QUERY,
      TIMEZONE_GROUP_ACTION_EXPORT,
      TIMEZONE_GROUP_ACTION_GET_ONE,
      TIMEZONE_GROUP_ACTION_POST_ONE,
      TIMEZONE_GROUP_ACTION_PATCH,
      TIMEZONE_GROUP_ACTION_PATCH_BULK,
      TIMEZONE_GROUP_ACTION_DELETE,
          TIMEZONE_GROUP_UTC_ITEMS_ACTION_PATCH,
          TIMEZONE_GROUP_UTC_ITEMS_ACTION_GET,
          TIMEZONE_GROUP_UTC_ITEMS_ACTION_POST,
    }
    // Append user defined functions
    AppendTimezoneGroupRouter(&routes)
    return routes
  }
  func CreateTimezoneGroupRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetTimezoneGroupModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, TimezoneGroupEntityJsonSchema, "timezone-group-http", "worldtimezone")
    workspaces.WriteEntitySchema("TimezoneGroupEntity", TimezoneGroupEntityJsonSchema, "worldtimezone")
    return httpRoutes
  }
var PERM_ROOT_TIMEZONE_GROUP_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/worldtimezone/timezone-group/delete",
  Name: "Delete timezone group",
}
var PERM_ROOT_TIMEZONE_GROUP_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/worldtimezone/timezone-group/create",
  Name: "Create timezone group",
}
var PERM_ROOT_TIMEZONE_GROUP_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/worldtimezone/timezone-group/update",
  Name: "Update timezone group",
}
var PERM_ROOT_TIMEZONE_GROUP_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/worldtimezone/timezone-group/query",
  Name: "Query timezone group",
}
var PERM_ROOT_TIMEZONE_GROUP = workspaces.PermissionInfo{
  CompleteKey: "root/worldtimezone/timezone-group/*",
  Name: "Entire timezone group actions (*)",
}
var ALL_TIMEZONE_GROUP_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_TIMEZONE_GROUP_DELETE,
	PERM_ROOT_TIMEZONE_GROUP_CREATE,
	PERM_ROOT_TIMEZONE_GROUP_UPDATE,
	PERM_ROOT_TIMEZONE_GROUP_QUERY,
	PERM_ROOT_TIMEZONE_GROUP,
}