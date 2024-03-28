package commonprofile
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
type CommonProfileEntity struct {
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
    LastName   *string `json:"lastName" yaml:"lastName"       `
    // Datenano also has a text representation
    PhoneNumber   *string `json:"phoneNumber" yaml:"phoneNumber"       `
    // Datenano also has a text representation
    Email   *string `json:"email" yaml:"email"       `
    // Datenano also has a text representation
    Company   *string `json:"company" yaml:"company"       `
    // Datenano also has a text representation
    Street   *string `json:"street" yaml:"street"       `
    // Datenano also has a text representation
    HouseNumber   *string `json:"houseNumber" yaml:"houseNumber"       `
    // Datenano also has a text representation
    ZipCode   *string `json:"zipCode" yaml:"zipCode"       `
    // Datenano also has a text representation
    City   *string `json:"city" yaml:"city"       `
    // Datenano also has a text representation
    Gender   *string `json:"gender" yaml:"gender"       `
    // Datenano also has a text representation
    Children []*CommonProfileEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *CommonProfileEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var CommonProfilePreloadRelations []string = []string{}
var COMMON_PROFILE_EVENT_CREATED = "commonProfile.created"
var COMMON_PROFILE_EVENT_UPDATED = "commonProfile.updated"
var COMMON_PROFILE_EVENT_DELETED = "commonProfile.deleted"
var COMMON_PROFILE_EVENTS = []string{
	COMMON_PROFILE_EVENT_CREATED,
	COMMON_PROFILE_EVENT_UPDATED,
	COMMON_PROFILE_EVENT_DELETED,
}
type CommonProfileFieldMap struct {
		FirstName workspaces.TranslatedString `yaml:"firstName"`
		LastName workspaces.TranslatedString `yaml:"lastName"`
		PhoneNumber workspaces.TranslatedString `yaml:"phoneNumber"`
		Email workspaces.TranslatedString `yaml:"email"`
		Company workspaces.TranslatedString `yaml:"company"`
		Street workspaces.TranslatedString `yaml:"street"`
		HouseNumber workspaces.TranslatedString `yaml:"houseNumber"`
		ZipCode workspaces.TranslatedString `yaml:"zipCode"`
		City workspaces.TranslatedString `yaml:"city"`
		Gender workspaces.TranslatedString `yaml:"gender"`
}
var CommonProfileEntityMetaConfig map[string]int64 = map[string]int64{
}
var CommonProfileEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&CommonProfileEntity{}))
func entityCommonProfileFormatter(dto *CommonProfileEntity, query workspaces.QueryDSL) {
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
func CommonProfileMockEntity() *CommonProfileEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &CommonProfileEntity{
      FirstName : &stringHolder,
      LastName : &stringHolder,
      PhoneNumber : &stringHolder,
      Email : &stringHolder,
      Company : &stringHolder,
      Street : &stringHolder,
      HouseNumber : &stringHolder,
      ZipCode : &stringHolder,
      City : &stringHolder,
      Gender : &stringHolder,
	}
	return entity
}
func CommonProfileActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := CommonProfileMockEntity()
		_, err := CommonProfileActionCreate(entity, query)
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
  func CommonProfileActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*CommonProfileEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &CommonProfileEntity{
          FirstName: &tildaRef,
          LastName: &tildaRef,
          PhoneNumber: &tildaRef,
          Email: &tildaRef,
          Company: &tildaRef,
          Street: &tildaRef,
          HouseNumber: &tildaRef,
          ZipCode: &tildaRef,
          City: &tildaRef,
          Gender: &tildaRef,
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
  func CommonProfileAssociationCreate(dto *CommonProfileEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func CommonProfileRelationContentCreate(dto *CommonProfileEntity, query workspaces.QueryDSL) error {
return nil
}
func CommonProfileRelationContentUpdate(dto *CommonProfileEntity, query workspaces.QueryDSL) error {
	return nil
}
func CommonProfilePolyglotCreateHandler(dto *CommonProfileEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func CommonProfileValidator(dto *CommonProfileEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func CommonProfileEntityPreSanitize(dto *CommonProfileEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func CommonProfileEntityBeforeCreateAppend(dto *CommonProfileEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    CommonProfileRecursiveAddUniqueId(dto, query)
  }
  func CommonProfileRecursiveAddUniqueId(dto *CommonProfileEntity, query workspaces.QueryDSL) {
  }
func CommonProfileActionBatchCreateFn(dtos []*CommonProfileEntity, query workspaces.QueryDSL) ([]*CommonProfileEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*CommonProfileEntity{}
		for _, item := range dtos {
			s, err := CommonProfileActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func CommonProfileDeleteEntireChildren(query workspaces.QueryDSL, dto *CommonProfileEntity) (*workspaces.IError) {
  return nil
}
func CommonProfileActionCreateFn(dto *CommonProfileEntity, query workspaces.QueryDSL) (*CommonProfileEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := CommonProfileValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	CommonProfileEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	CommonProfileEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	CommonProfilePolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	CommonProfileRelationContentCreate(dto, query)
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
	CommonProfileAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(COMMON_PROFILE_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&CommonProfileEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func CommonProfileActionGetOne(query workspaces.QueryDSL) (*CommonProfileEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&CommonProfileEntity{})
    item, err := workspaces.GetOneEntity[CommonProfileEntity](query, refl)
    entityCommonProfileFormatter(item, query)
    return item, err
  }
  func CommonProfileActionQuery(query workspaces.QueryDSL) ([]*CommonProfileEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&CommonProfileEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[CommonProfileEntity](query, refl)
    for _, item := range items {
      entityCommonProfileFormatter(item, query)
    }
    return items, meta, err
  }
  func CommonProfileUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *CommonProfileEntity) (*CommonProfileEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = COMMON_PROFILE_EVENT_UPDATED
    CommonProfileEntityPreSanitize(fields, query)
    var item CommonProfileEntity
    q := dbref.
      Where(&CommonProfileEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    CommonProfileRelationContentUpdate(fields, query)
    CommonProfilePolyglotCreateHandler(fields, query)
    if ero := CommonProfileDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&CommonProfileEntity{UniqueId: uniqueId}).
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
  func CommonProfileActionUpdateFn(query workspaces.QueryDSL, fields *CommonProfileEntity) (*CommonProfileEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := CommonProfileValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // CommonProfileRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *CommonProfileEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = CommonProfileUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return CommonProfileUpdateExec(dbref, query, fields)
    }
  }
var CommonProfileWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire commonprofiles ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_COMMON_PROFILE_DELETE},
    })
		count, _ := CommonProfileActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func CommonProfileActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&CommonProfileEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_COMMON_PROFILE_DELETE}
	return workspaces.RemoveEntity[CommonProfileEntity](query, refl)
}
func CommonProfileActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[CommonProfileEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'CommonProfileEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func CommonProfileActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[CommonProfileEntity]) (
    *workspaces.BulkRecordRequest[CommonProfileEntity], *workspaces.IError,
  ) {
    result := []*CommonProfileEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := CommonProfileActionUpdate(query, record)
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
func (x *CommonProfileEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var CommonProfileEntityMeta = workspaces.TableMetaData{
	EntityName:    "CommonProfile",
	ExportKey:    "common-profiles",
	TableNameInDb: "fb_common-profile_entities",
	EntityObject:  &CommonProfileEntity{},
	ExportStream: CommonProfileActionExportT,
	ImportQuery: CommonProfileActionImport,
}
func CommonProfileActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[CommonProfileEntity](query, CommonProfileActionQuery, CommonProfilePreloadRelations)
}
func CommonProfileActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[CommonProfileEntity](query, CommonProfileActionQuery, CommonProfilePreloadRelations)
}
func CommonProfileActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content CommonProfileEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := CommonProfileActionCreate(&content, query)
	return err
}
var CommonProfileCommonCliFlags = []cli.Flag{
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
      Name:     "last-name",
      Required: false,
      Usage:    "lastName",
    },
    &cli.StringFlag{
      Name:     "phone-number",
      Required: false,
      Usage:    "phoneNumber",
    },
    &cli.StringFlag{
      Name:     "email",
      Required: false,
      Usage:    "email",
    },
    &cli.StringFlag{
      Name:     "company",
      Required: false,
      Usage:    "company",
    },
    &cli.StringFlag{
      Name:     "street",
      Required: false,
      Usage:    "street",
    },
    &cli.StringFlag{
      Name:     "house-number",
      Required: false,
      Usage:    "houseNumber",
    },
    &cli.StringFlag{
      Name:     "zip-code",
      Required: false,
      Usage:    "zipCode",
    },
    &cli.StringFlag{
      Name:     "city",
      Required: false,
      Usage:    "city",
    },
    &cli.StringFlag{
      Name:     "gender",
      Required: false,
      Usage:    "gender",
    },
}
var CommonProfileCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "firstName",
		StructField:     "FirstName",
		Required: false,
		Usage:    "firstName",
		Type: "string",
	},
	{
		Name:     "lastName",
		StructField:     "LastName",
		Required: false,
		Usage:    "lastName",
		Type: "string",
	},
	{
		Name:     "phoneNumber",
		StructField:     "PhoneNumber",
		Required: false,
		Usage:    "phoneNumber",
		Type: "string",
	},
	{
		Name:     "email",
		StructField:     "Email",
		Required: false,
		Usage:    "email",
		Type: "string",
	},
	{
		Name:     "company",
		StructField:     "Company",
		Required: false,
		Usage:    "company",
		Type: "string",
	},
	{
		Name:     "street",
		StructField:     "Street",
		Required: false,
		Usage:    "street",
		Type: "string",
	},
	{
		Name:     "houseNumber",
		StructField:     "HouseNumber",
		Required: false,
		Usage:    "houseNumber",
		Type: "string",
	},
	{
		Name:     "zipCode",
		StructField:     "ZipCode",
		Required: false,
		Usage:    "zipCode",
		Type: "string",
	},
	{
		Name:     "city",
		StructField:     "City",
		Required: false,
		Usage:    "city",
		Type: "string",
	},
	{
		Name:     "gender",
		StructField:     "Gender",
		Required: false,
		Usage:    "gender",
		Type: "string",
	},
}
var CommonProfileCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "last-name",
      Required: false,
      Usage:    "lastName",
    },
    &cli.StringFlag{
      Name:     "phone-number",
      Required: false,
      Usage:    "phoneNumber",
    },
    &cli.StringFlag{
      Name:     "email",
      Required: false,
      Usage:    "email",
    },
    &cli.StringFlag{
      Name:     "company",
      Required: false,
      Usage:    "company",
    },
    &cli.StringFlag{
      Name:     "street",
      Required: false,
      Usage:    "street",
    },
    &cli.StringFlag{
      Name:     "house-number",
      Required: false,
      Usage:    "houseNumber",
    },
    &cli.StringFlag{
      Name:     "zip-code",
      Required: false,
      Usage:    "zipCode",
    },
    &cli.StringFlag{
      Name:     "city",
      Required: false,
      Usage:    "city",
    },
    &cli.StringFlag{
      Name:     "gender",
      Required: false,
      Usage:    "gender",
    },
}
  var CommonProfileCreateCmd cli.Command = COMMON_PROFILE_ACTION_POST_ONE.ToCli()
  var CommonProfileCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_COMMON_PROFILE_CREATE},
      })
      entity := &CommonProfileEntity{}
      for _, item := range CommonProfileCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := CommonProfileActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var CommonProfileUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: CommonProfileCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_COMMON_PROFILE_UPDATE},
      })
      entity := CastCommonProfileFromCli(c)
      if entity, err := CommonProfileActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* CommonProfileEntity) FromCli(c *cli.Context) *CommonProfileEntity {
	return CastCommonProfileFromCli(c)
}
func CastCommonProfileFromCli (c *cli.Context) *CommonProfileEntity {
	template := &CommonProfileEntity{}
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
      if c.IsSet("last-name") {
        value := c.String("last-name")
        template.LastName = &value
      }
      if c.IsSet("phone-number") {
        value := c.String("phone-number")
        template.PhoneNumber = &value
      }
      if c.IsSet("email") {
        value := c.String("email")
        template.Email = &value
      }
      if c.IsSet("company") {
        value := c.String("company")
        template.Company = &value
      }
      if c.IsSet("street") {
        value := c.String("street")
        template.Street = &value
      }
      if c.IsSet("house-number") {
        value := c.String("house-number")
        template.HouseNumber = &value
      }
      if c.IsSet("zip-code") {
        value := c.String("zip-code")
        template.ZipCode = &value
      }
      if c.IsSet("city") {
        value := c.String("city")
        template.City = &value
      }
      if c.IsSet("gender") {
        value := c.String("gender")
        template.Gender = &value
      }
	return template
}
  func CommonProfileSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      CommonProfileActionCreate,
      reflect.ValueOf(&CommonProfileEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func CommonProfileWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := CommonProfileActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "CommonProfile", result)
    }
  }
var CommonProfileImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_COMMON_PROFILE_CREATE},
      })
			CommonProfileActionSeeder(query, c.Int("count"))
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
				Value: "common-profile-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_COMMON_PROFILE_CREATE},
      })
			CommonProfileActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "common-profile-seeder-common-profile.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of common-profiles, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]CommonProfileEntity{}
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
			CommonProfileCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				CommonProfileActionCreate,
				reflect.ValueOf(&CommonProfileEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_COMMON_PROFILE_CREATE},
				},
        func() CommonProfileEntity {
					v := CastCommonProfileFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var CommonProfileCliCommands []cli.Command = []cli.Command{
      COMMON_PROFILE_ACTION_QUERY.ToCli(),
      COMMON_PROFILE_ACTION_TABLE.ToCli(),
      CommonProfileCreateCmd,
      CommonProfileUpdateCmd,
      CommonProfileCreateInteractiveCmd,
      CommonProfileWipeCmd,
      workspaces.GetCommonRemoveQuery(reflect.ValueOf(&CommonProfileEntity{}).Elem(), CommonProfileActionRemove),
  }
  func CommonProfileCliFn() cli.Command {
    CommonProfileCliCommands = append(CommonProfileCliCommands, CommonProfileImportExportCommands...)
    return cli.Command{
      Name:        "commonProfile",
      Description: "CommonProfiles module actions (sample module to handle complex entities)",
      Usage:       "A common profile issues for every user (Set the living address, etc)",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: CommonProfileCliCommands,
    }
  }
var COMMON_PROFILE_ACTION_TABLE = workspaces.Module2Action{
  Name:    "table",
  ActionAliases: []string{"t"},
  Flags:  workspaces.CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: CommonProfileActionQuery,
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    workspaces.CommonCliTableCmd2(c,
      CommonProfileActionQuery,
      security,
      reflect.ValueOf(&CommonProfileEntity{}).Elem(),
    )
    return nil
  },
}
var COMMON_PROFILE_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/common-profiles",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_COMMON_PROFILE_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, CommonProfileActionQuery)
    },
  },
  Format: "QUERY",
  Action: CommonProfileActionQuery,
  ResponseEntity: &[]CommonProfileEntity{},
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			CommonProfileActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionAliases: []string{"q"},
	Flags:         workspaces.CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var COMMON_PROFILE_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/common-profiles/export",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_COMMON_PROFILE_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, CommonProfileActionExport)
    },
  },
  Format: "QUERY",
  Action: CommonProfileActionExport,
  ResponseEntity: &[]CommonProfileEntity{},
}
var COMMON_PROFILE_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/common-profile/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_COMMON_PROFILE_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, CommonProfileActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: CommonProfileActionGetOne,
  ResponseEntity: &CommonProfileEntity{},
}
var COMMON_PROFILE_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new commonProfile",
  Flags: CommonProfileCommonCliFlags,
  Method: "POST",
  Url:    "/common-profile",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_COMMON_PROFILE_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, CommonProfileActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, CommonProfileActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: CommonProfileActionCreate,
  Format: "POST_ONE",
  RequestEntity: &CommonProfileEntity{},
  ResponseEntity: &CommonProfileEntity{},
}
var COMMON_PROFILE_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: CommonProfileCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/common-profile",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_COMMON_PROFILE_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, CommonProfileActionUpdate)
    },
  },
  Action: CommonProfileActionUpdate,
  RequestEntity: &CommonProfileEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &CommonProfileEntity{},
}
var COMMON_PROFILE_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/common-profiles",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_COMMON_PROFILE_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, CommonProfileActionBulkUpdate)
    },
  },
  Action: CommonProfileActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[CommonProfileEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[CommonProfileEntity]{},
}
var COMMON_PROFILE_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/common-profile",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_COMMON_PROFILE_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, CommonProfileActionRemove)
    },
  },
  Action: CommonProfileActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &CommonProfileEntity{},
}
var COMMON_PROFILE_ACTION_DISTINCT_PATCH_ONE = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/common-profile/distinct",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_COMMON_PROFILE_UPDATE_DISTINCT_USER},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, CommonProfileDistinctActionUpdate)
    },
  },
  Action: CommonProfileDistinctActionUpdate,
  Format: "PATCH_ONE",
  RequestEntity: &CommonProfileEntity{},
  ResponseEntity: &CommonProfileEntity{},
}
var COMMON_PROFILE_ACTION_DISTINCT_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/common-profile/distinct",
  SecurityModel: &workspaces.SecurityModel{
    ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_COMMON_PROFILE_GET_DISTINCT_USER},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, CommonProfileDistinctActionGetOne)
    },
  },
  Action: CommonProfileDistinctActionGetOne,
  Format: "GET_ONE",
  ResponseEntity: &CommonProfileEntity{},
}
  /**
  *	Override this function on CommonProfileEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendCommonProfileRouter = func(r *[]workspaces.Module2Action) {}
  func GetCommonProfileModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      COMMON_PROFILE_ACTION_QUERY,
      COMMON_PROFILE_ACTION_EXPORT,
      COMMON_PROFILE_ACTION_GET_ONE,
      COMMON_PROFILE_ACTION_POST_ONE,
      COMMON_PROFILE_ACTION_PATCH,
      COMMON_PROFILE_ACTION_PATCH_BULK,
      COMMON_PROFILE_ACTION_DELETE,
      COMMON_PROFILE_ACTION_DISTINCT_PATCH_ONE,
      COMMON_PROFILE_ACTION_DISTINCT_GET_ONE,
    }
    // Append user defined functions
    AppendCommonProfileRouter(&routes)
    return routes
  }
  func CreateCommonProfileRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetCommonProfileModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, CommonProfileEntityJsonSchema, "common-profile-http", "commonprofile")
    workspaces.WriteEntitySchema("CommonProfileEntity", CommonProfileEntityJsonSchema, "commonprofile")
    return httpRoutes
  }
var PERM_ROOT_COMMON_PROFILE_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/commonprofile/common-profile/delete",
  Name: "Delete common profile",
}
var PERM_ROOT_COMMON_PROFILE_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/commonprofile/common-profile/create",
  Name: "Create common profile",
}
var PERM_ROOT_COMMON_PROFILE_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/commonprofile/common-profile/update",
  Name: "Update common profile",
}
var PERM_ROOT_COMMON_PROFILE_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/commonprofile/common-profile/query",
  Name: "Query common profile",
}
  var PERM_ROOT_COMMON_PROFILE_GET_DISTINCT_USER = workspaces.PermissionInfo{
    CompleteKey: "root/commonprofile/common-profile/get-distinct-user",
    Name: "Get common profile Distinct",
  }
  var PERM_ROOT_COMMON_PROFILE_UPDATE_DISTINCT_USER = workspaces.PermissionInfo{
    CompleteKey: "root/commonprofile/common-profile/update-distinct-user",
    Name: "Update common profile Distinct",
  }
var PERM_ROOT_COMMON_PROFILE = workspaces.PermissionInfo{
  CompleteKey: "root/commonprofile/common-profile/*",
  Name: "Entire common profile actions (*)",
}
var ALL_COMMON_PROFILE_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_COMMON_PROFILE_DELETE,
	PERM_ROOT_COMMON_PROFILE_CREATE,
	PERM_ROOT_COMMON_PROFILE_UPDATE,
    PERM_ROOT_COMMON_PROFILE_GET_DISTINCT_USER,
    PERM_ROOT_COMMON_PROFILE_UPDATE_DISTINCT_USER,
	PERM_ROOT_COMMON_PROFILE_QUERY,
	PERM_ROOT_COMMON_PROFILE,
}
  func CommonProfileDistinctActionUpdate(
    query workspaces.QueryDSL,
    fields *CommonProfileEntity,
  ) (*CommonProfileEntity, *workspaces.IError) {
    query.UniqueId = query.UserId
    entity, err := CommonProfileActionGetOne(query)
    if err != nil || entity.UniqueId == "" {
      fields.UniqueId = query.UserId
      return CommonProfileActionCreateFn(fields, query)
    } else {
      fields.UniqueId = query.UniqueId
      return CommonProfileActionUpdateFn(query, fields)
    }
  }
  func CommonProfileDistinctActionGetOne(
    query workspaces.QueryDSL,
  ) (*CommonProfileEntity, *workspaces.IError) {
    query.UniqueId = query.UserId
    entity, err := CommonProfileActionGetOne(query)
    if err != nil && err.HttpCode == 404 {
      return &CommonProfileEntity{}, nil
    }
    return entity, err
  }