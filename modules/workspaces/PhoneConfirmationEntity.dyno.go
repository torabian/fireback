package workspaces
import (
    "github.com/gin-gonic/gin"
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
type PhoneConfirmationEntity struct {
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
    User   *  UserEntity `json:"user" yaml:"user"    gorm:"foreignKey:UserId;references:UniqueId"     `
    // Datenano also has a text representation
    Status   *string `json:"status" yaml:"status"       `
    // Datenano also has a text representation
    PhoneNumber   *string `json:"phoneNumber" yaml:"phoneNumber"       `
    // Datenano also has a text representation
    Key   *string `json:"key" yaml:"key"       `
    // Datenano also has a text representation
    ExpiresAt   *string `json:"expiresAt" yaml:"expiresAt"       `
    // Datenano also has a text representation
    Children []*PhoneConfirmationEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *PhoneConfirmationEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var PhoneConfirmationPreloadRelations []string = []string{}
var PHONE_CONFIRMATION_EVENT_CREATED = "phoneConfirmation.created"
var PHONE_CONFIRMATION_EVENT_UPDATED = "phoneConfirmation.updated"
var PHONE_CONFIRMATION_EVENT_DELETED = "phoneConfirmation.deleted"
var PHONE_CONFIRMATION_EVENTS = []string{
	PHONE_CONFIRMATION_EVENT_CREATED,
	PHONE_CONFIRMATION_EVENT_UPDATED,
	PHONE_CONFIRMATION_EVENT_DELETED,
}
type PhoneConfirmationFieldMap struct {
		User TranslatedString `yaml:"user"`
		Status TranslatedString `yaml:"status"`
		PhoneNumber TranslatedString `yaml:"phoneNumber"`
		Key TranslatedString `yaml:"key"`
		ExpiresAt TranslatedString `yaml:"expiresAt"`
}
var PhoneConfirmationEntityMetaConfig map[string]int64 = map[string]int64{
}
var PhoneConfirmationEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&PhoneConfirmationEntity{}))
func entityPhoneConfirmationFormatter(dto *PhoneConfirmationEntity, query QueryDSL) {
	if dto == nil {
		return
	}
	if dto.Created > 0 {
		dto.CreatedFormatted = FormatDateBasedOnQuery(dto.Created, query)
	}
	if dto.Updated > 0 {
		dto.CreatedFormatted = FormatDateBasedOnQuery(dto.Updated, query)
	}
}
func PhoneConfirmationMockEntity() *PhoneConfirmationEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &PhoneConfirmationEntity{
      Status : &stringHolder,
      PhoneNumber : &stringHolder,
      Key : &stringHolder,
      ExpiresAt : &stringHolder,
	}
	return entity
}
func PhoneConfirmationActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := PhoneConfirmationMockEntity()
		_, err := PhoneConfirmationActionCreate(entity, query)
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
  func PhoneConfirmationActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*PhoneConfirmationEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &PhoneConfirmationEntity{
          Status: &tildaRef,
          PhoneNumber: &tildaRef,
          Key: &tildaRef,
          ExpiresAt: &tildaRef,
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
  func PhoneConfirmationAssociationCreate(dto *PhoneConfirmationEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func PhoneConfirmationRelationContentCreate(dto *PhoneConfirmationEntity, query QueryDSL) error {
return nil
}
func PhoneConfirmationRelationContentUpdate(dto *PhoneConfirmationEntity, query QueryDSL) error {
	return nil
}
func PhoneConfirmationPolyglotCreateHandler(dto *PhoneConfirmationEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func PhoneConfirmationValidator(dto *PhoneConfirmationEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func PhoneConfirmationEntityPreSanitize(dto *PhoneConfirmationEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func PhoneConfirmationEntityBeforeCreateAppend(dto *PhoneConfirmationEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    PhoneConfirmationRecursiveAddUniqueId(dto, query)
  }
  func PhoneConfirmationRecursiveAddUniqueId(dto *PhoneConfirmationEntity, query QueryDSL) {
  }
func PhoneConfirmationActionBatchCreateFn(dtos []*PhoneConfirmationEntity, query QueryDSL) ([]*PhoneConfirmationEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*PhoneConfirmationEntity{}
		for _, item := range dtos {
			s, err := PhoneConfirmationActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func PhoneConfirmationDeleteEntireChildren(query QueryDSL, dto *PhoneConfirmationEntity) (*IError) {
  return nil
}
func PhoneConfirmationActionCreateFn(dto *PhoneConfirmationEntity, query QueryDSL) (*PhoneConfirmationEntity, *IError) {
	// 1. Validate always
	if iError := PhoneConfirmationValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	PhoneConfirmationEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	PhoneConfirmationEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	PhoneConfirmationPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	PhoneConfirmationRelationContentCreate(dto, query)
	// 4. Create the entity
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = GetDbRef()
	} else {
		dbref = query.Tx
	}
	query.Tx = dbref;
	err := dbref.Create(&dto).Error
	if err != nil {
		err := GormErrorToIError(err)
		return dto, err
	}
	// 5. Create sub entities, objects or arrays, association to other entities
	PhoneConfirmationAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(PHONE_CONFIRMATION_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&PhoneConfirmationEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func PhoneConfirmationActionGetOne(query QueryDSL) (*PhoneConfirmationEntity, *IError) {
    refl := reflect.ValueOf(&PhoneConfirmationEntity{})
    item, err := GetOneEntity[PhoneConfirmationEntity](query, refl)
    entityPhoneConfirmationFormatter(item, query)
    return item, err
  }
  func PhoneConfirmationActionQuery(query QueryDSL) ([]*PhoneConfirmationEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&PhoneConfirmationEntity{})
    items, meta, err := QueryEntitiesPointer[PhoneConfirmationEntity](query, refl)
    for _, item := range items {
      entityPhoneConfirmationFormatter(item, query)
    }
    return items, meta, err
  }
  func PhoneConfirmationUpdateExec(dbref *gorm.DB, query QueryDSL, fields *PhoneConfirmationEntity) (*PhoneConfirmationEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = PHONE_CONFIRMATION_EVENT_UPDATED
    PhoneConfirmationEntityPreSanitize(fields, query)
    var item PhoneConfirmationEntity
    q := dbref.
      Where(&PhoneConfirmationEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    PhoneConfirmationRelationContentUpdate(fields, query)
    PhoneConfirmationPolyglotCreateHandler(fields, query)
    if ero := PhoneConfirmationDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&PhoneConfirmationEntity{UniqueId: uniqueId}).
      First(&item).Error
    event.MustFire(query.TriggerEventName, event.M{
      "entity":   &item,
      "target":   "workspace",
      "unqiueId": query.WorkspaceId,
    })
    if err != nil {
      return &item, GormErrorToIError(err)
    }
    return &item, nil
  }
  func PhoneConfirmationActionUpdateFn(query QueryDSL, fields *PhoneConfirmationEntity) (*PhoneConfirmationEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := PhoneConfirmationValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // PhoneConfirmationRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      var item *PhoneConfirmationEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *IError
        item, err = PhoneConfirmationUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, CastToIError(vf)
    } else {
      dbref = query.Tx
      return PhoneConfirmationUpdateExec(dbref, query, fields)
    }
  }
var PhoneConfirmationWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire phoneconfirmations ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
      ActionRequires: []PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_DELETE},
    })
		count, _ := PhoneConfirmationActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func PhoneConfirmationActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&PhoneConfirmationEntity{})
	query.ActionRequires = []PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_DELETE}
	return RemoveEntity[PhoneConfirmationEntity](query, refl)
}
func PhoneConfirmationActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[PhoneConfirmationEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'PhoneConfirmationEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func PhoneConfirmationActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[PhoneConfirmationEntity]) (
    *BulkRecordRequest[PhoneConfirmationEntity], *IError,
  ) {
    result := []*PhoneConfirmationEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := PhoneConfirmationActionUpdate(query, record)
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
    return nil, err.(*IError)
  }
func (x *PhoneConfirmationEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var PhoneConfirmationEntityMeta = TableMetaData{
	EntityName:    "PhoneConfirmation",
	ExportKey:    "phone-confirmations",
	TableNameInDb: "fb_phone-confirmation_entities",
	EntityObject:  &PhoneConfirmationEntity{},
	ExportStream: PhoneConfirmationActionExportT,
	ImportQuery: PhoneConfirmationActionImport,
}
func PhoneConfirmationActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[PhoneConfirmationEntity](query, PhoneConfirmationActionQuery, PhoneConfirmationPreloadRelations)
}
func PhoneConfirmationActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[PhoneConfirmationEntity](query, PhoneConfirmationActionQuery, PhoneConfirmationPreloadRelations)
}
func PhoneConfirmationActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content PhoneConfirmationEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := PhoneConfirmationActionCreate(&content, query)
	return err
}
var PhoneConfirmationCommonCliFlags = []cli.Flag{
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
      Name:     "user-id",
      Required: false,
      Usage:    "user",
    },
    &cli.StringFlag{
      Name:     "status",
      Required: false,
      Usage:    "status",
    },
    &cli.StringFlag{
      Name:     "phone-number",
      Required: false,
      Usage:    "phoneNumber",
    },
    &cli.StringFlag{
      Name:     "key",
      Required: false,
      Usage:    "key",
    },
    &cli.StringFlag{
      Name:     "expires-at",
      Required: false,
      Usage:    "expiresAt",
    },
}
var PhoneConfirmationCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:     "status",
		StructField:     "Status",
		Required: false,
		Usage:    "status",
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
		Name:     "key",
		StructField:     "Key",
		Required: false,
		Usage:    "key",
		Type: "string",
	},
	{
		Name:     "expiresAt",
		StructField:     "ExpiresAt",
		Required: false,
		Usage:    "expiresAt",
		Type: "string",
	},
}
var PhoneConfirmationCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "user-id",
      Required: false,
      Usage:    "user",
    },
    &cli.StringFlag{
      Name:     "status",
      Required: false,
      Usage:    "status",
    },
    &cli.StringFlag{
      Name:     "phone-number",
      Required: false,
      Usage:    "phoneNumber",
    },
    &cli.StringFlag{
      Name:     "key",
      Required: false,
      Usage:    "key",
    },
    &cli.StringFlag{
      Name:     "expires-at",
      Required: false,
      Usage:    "expiresAt",
    },
}
  var PhoneConfirmationCreateCmd cli.Command = PHONE_CONFIRMATION_ACTION_POST_ONE.ToCli()
  var PhoneConfirmationCreateInteractiveCmd cli.Command = cli.Command{
    Name:  "ic",
    Usage: "Creates a new template, using requied fields in an interactive name",
    Flags: []cli.Flag{
      &cli.BoolFlag{
        Name:  "all",
        Usage: "Interactively asks for all inputs, not only required ones",
      },
    },
    Action: func(c *cli.Context) {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_CREATE},
      })
      entity := &PhoneConfirmationEntity{}
      for _, item := range PhoneConfirmationCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := PhoneConfirmationActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var PhoneConfirmationUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: PhoneConfirmationCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_UPDATE},
      })
      entity := CastPhoneConfirmationFromCli(c)
      if entity, err := PhoneConfirmationActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* PhoneConfirmationEntity) FromCli(c *cli.Context) *PhoneConfirmationEntity {
	return CastPhoneConfirmationFromCli(c)
}
func CastPhoneConfirmationFromCli (c *cli.Context) *PhoneConfirmationEntity {
	template := &PhoneConfirmationEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("user-id") {
        value := c.String("user-id")
        template.UserId = &value
      }
      if c.IsSet("status") {
        value := c.String("status")
        template.Status = &value
      }
      if c.IsSet("phone-number") {
        value := c.String("phone-number")
        template.PhoneNumber = &value
      }
      if c.IsSet("key") {
        value := c.String("key")
        template.Key = &value
      }
      if c.IsSet("expires-at") {
        value := c.String("expires-at")
        template.ExpiresAt = &value
      }
	return template
}
  func PhoneConfirmationSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      PhoneConfirmationActionCreate,
      reflect.ValueOf(&PhoneConfirmationEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func PhoneConfirmationWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := PhoneConfirmationActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "PhoneConfirmation", result)
    }
  }
var PhoneConfirmationImportExportCommands = []cli.Command{
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
			query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_CREATE},
      })
			PhoneConfirmationActionSeeder(query, c.Int("count"))
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
				Value: "phone-confirmation-seeder.yml",
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
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_CREATE},
      })
			PhoneConfirmationActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "phone-confirmation-seeder-phone-confirmation.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of phone-confirmations, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]PhoneConfirmationEntity{}
			ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
	cli.Command{
		Name:    "import",
    Flags: append(
			append(
				CommonQueryFlags,
				&cli.StringFlag{
					Name:     "file",
					Usage:    "The address of file you want the csv be imported from",
					Required: true,
				}),
			PhoneConfirmationCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				PhoneConfirmationActionCreate,
				reflect.ValueOf(&PhoneConfirmationEntity{}).Elem(),
				c.String("file"),
        &SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_CREATE},
				},
        func() PhoneConfirmationEntity {
					v := CastPhoneConfirmationFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var PhoneConfirmationCliCommands []cli.Command = []cli.Command{
      PHONE_CONFIRMATION_ACTION_QUERY.ToCli(),
      PHONE_CONFIRMATION_ACTION_TABLE.ToCli(),
      PhoneConfirmationCreateCmd,
      PhoneConfirmationUpdateCmd,
      PhoneConfirmationCreateInteractiveCmd,
      PhoneConfirmationWipeCmd,
      GetCommonRemoveQuery(reflect.ValueOf(&PhoneConfirmationEntity{}).Elem(), PhoneConfirmationActionRemove),
  }
  func PhoneConfirmationCliFn() cli.Command {
    PhoneConfirmationCliCommands = append(PhoneConfirmationCliCommands, PhoneConfirmationImportExportCommands...)
    return cli.Command{
      Name:        "phoneConfirmation",
      Description: "PhoneConfirmations module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: PhoneConfirmationCliCommands,
    }
  }
var PHONE_CONFIRMATION_ACTION_TABLE = Module2Action{
  Name:    "table",
  ActionAliases: []string{"t"},
  Flags:  CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: PhoneConfirmationActionQuery,
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    CommonCliTableCmd2(c,
      PhoneConfirmationActionQuery,
      security,
      reflect.ValueOf(&PhoneConfirmationEntity{}).Elem(),
    )
    return nil
  },
}
var PHONE_CONFIRMATION_ACTION_QUERY = Module2Action{
  Method: "GET",
  Url:    "/phone-confirmations",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpQueryEntity(c, PhoneConfirmationActionQuery)
    },
  },
  Format: "QUERY",
  Action: PhoneConfirmationActionQuery,
  ResponseEntity: &[]PhoneConfirmationEntity{},
  CliAction: func(c *cli.Context, security *SecurityModel) error {
		CommonCliQueryCmd2(
			c,
			PhoneConfirmationActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionAliases: []string{"q"},
	Flags:         CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var PHONE_CONFIRMATION_ACTION_EXPORT = Module2Action{
  Method: "GET",
  Url:    "/phone-confirmations/export",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpStreamFileChannel(c, PhoneConfirmationActionExport)
    },
  },
  Format: "QUERY",
  Action: PhoneConfirmationActionExport,
  ResponseEntity: &[]PhoneConfirmationEntity{},
}
var PHONE_CONFIRMATION_ACTION_GET_ONE = Module2Action{
  Method: "GET",
  Url:    "/phone-confirmation/:uniqueId",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpGetEntity(c, PhoneConfirmationActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: PhoneConfirmationActionGetOne,
  ResponseEntity: &PhoneConfirmationEntity{},
}
var PHONE_CONFIRMATION_ACTION_POST_ONE = Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new phoneConfirmation",
  Flags: PhoneConfirmationCommonCliFlags,
  Method: "POST",
  Url:    "/phone-confirmation",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpPostEntity(c, PhoneConfirmationActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    result, err := CliPostEntity(c, PhoneConfirmationActionCreate, security)
    HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: PhoneConfirmationActionCreate,
  Format: "POST_ONE",
  RequestEntity: &PhoneConfirmationEntity{},
  ResponseEntity: &PhoneConfirmationEntity{},
}
var PHONE_CONFIRMATION_ACTION_PATCH = Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: PhoneConfirmationCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/phone-confirmation",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntity(c, PhoneConfirmationActionUpdate)
    },
  },
  Action: PhoneConfirmationActionUpdate,
  RequestEntity: &PhoneConfirmationEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &PhoneConfirmationEntity{},
}
var PHONE_CONFIRMATION_ACTION_PATCH_BULK = Module2Action{
  Method: "PATCH",
  Url:    "/phone-confirmations",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntities(c, PhoneConfirmationActionBulkUpdate)
    },
  },
  Action: PhoneConfirmationActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &BulkRecordRequest[PhoneConfirmationEntity]{},
  ResponseEntity: &BulkRecordRequest[PhoneConfirmationEntity]{},
}
var PHONE_CONFIRMATION_ACTION_DELETE = Module2Action{
  Method: "DELETE",
  Url:    "/phone-confirmation",
  Format: "DELETE_DSL",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpRemoveEntity(c, PhoneConfirmationActionRemove)
    },
  },
  Action: PhoneConfirmationActionRemove,
  RequestEntity: &DeleteRequest{},
  ResponseEntity: &DeleteResponse{},
  TargetEntity: &PhoneConfirmationEntity{},
}
  /**
  *	Override this function on PhoneConfirmationEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendPhoneConfirmationRouter = func(r *[]Module2Action) {}
  func GetPhoneConfirmationModule2Actions() []Module2Action {
    routes := []Module2Action{
      PHONE_CONFIRMATION_ACTION_QUERY,
      PHONE_CONFIRMATION_ACTION_EXPORT,
      PHONE_CONFIRMATION_ACTION_GET_ONE,
      PHONE_CONFIRMATION_ACTION_POST_ONE,
      PHONE_CONFIRMATION_ACTION_PATCH,
      PHONE_CONFIRMATION_ACTION_PATCH_BULK,
      PHONE_CONFIRMATION_ACTION_DELETE,
    }
    // Append user defined functions
    AppendPhoneConfirmationRouter(&routes)
    return routes
  }
  func CreatePhoneConfirmationRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetPhoneConfirmationModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, PhoneConfirmationEntityJsonSchema, "phone-confirmation-http", "workspaces")
    WriteEntitySchema("PhoneConfirmationEntity", PhoneConfirmationEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_PHONE_CONFIRMATION_DELETE = PermissionInfo{
  CompleteKey: "root/workspaces/phone-confirmation/delete",
}
var PERM_ROOT_PHONE_CONFIRMATION_CREATE = PermissionInfo{
  CompleteKey: "root/workspaces/phone-confirmation/create",
}
var PERM_ROOT_PHONE_CONFIRMATION_UPDATE = PermissionInfo{
  CompleteKey: "root/workspaces/phone-confirmation/update",
}
var PERM_ROOT_PHONE_CONFIRMATION_QUERY = PermissionInfo{
  CompleteKey: "root/workspaces/phone-confirmation/query",
}
var PERM_ROOT_PHONE_CONFIRMATION = PermissionInfo{
  CompleteKey: "root/workspaces/phone-confirmation/*",
}
var ALL_PHONE_CONFIRMATION_PERMISSIONS = []PermissionInfo{
	PERM_ROOT_PHONE_CONFIRMATION_DELETE,
	PERM_ROOT_PHONE_CONFIRMATION_CREATE,
	PERM_ROOT_PHONE_CONFIRMATION_UPDATE,
	PERM_ROOT_PHONE_CONFIRMATION_QUERY,
	PERM_ROOT_PHONE_CONFIRMATION,
}