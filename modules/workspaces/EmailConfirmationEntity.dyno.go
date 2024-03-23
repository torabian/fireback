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
type EmailConfirmationEntity struct {
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
    User   *  UserEntity `json:"user" yaml:"user"    gorm:"foreignKey:UserId;references:UniqueId"     `
    // Datenano also has a text representation
    Status   *string `json:"status" yaml:"status"       `
    // Datenano also has a text representation
    Email   *string `json:"email" yaml:"email"       `
    // Datenano also has a text representation
    Key   *string `json:"key" yaml:"key"       `
    // Datenano also has a text representation
    ExpiresAt   *string `json:"expiresAt" yaml:"expiresAt"       `
    // Datenano also has a text representation
    Children []*EmailConfirmationEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *EmailConfirmationEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var EmailConfirmationPreloadRelations []string = []string{}
var EMAILCONFIRMATION_EVENT_CREATED = "emailConfirmation.created"
var EMAILCONFIRMATION_EVENT_UPDATED = "emailConfirmation.updated"
var EMAILCONFIRMATION_EVENT_DELETED = "emailConfirmation.deleted"
var EMAILCONFIRMATION_EVENTS = []string{
	EMAILCONFIRMATION_EVENT_CREATED,
	EMAILCONFIRMATION_EVENT_UPDATED,
	EMAILCONFIRMATION_EVENT_DELETED,
}
type EmailConfirmationFieldMap struct {
		User TranslatedString `yaml:"user"`
		Status TranslatedString `yaml:"status"`
		Email TranslatedString `yaml:"email"`
		Key TranslatedString `yaml:"key"`
		ExpiresAt TranslatedString `yaml:"expiresAt"`
}
var EmailConfirmationEntityMetaConfig map[string]int64 = map[string]int64{
}
var EmailConfirmationEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&EmailConfirmationEntity{}))
func entityEmailConfirmationFormatter(dto *EmailConfirmationEntity, query QueryDSL) {
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
func EmailConfirmationMockEntity() *EmailConfirmationEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &EmailConfirmationEntity{
      Status : &stringHolder,
      Email : &stringHolder,
      Key : &stringHolder,
      ExpiresAt : &stringHolder,
	}
	return entity
}
func EmailConfirmationActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := EmailConfirmationMockEntity()
		_, err := EmailConfirmationActionCreate(entity, query)
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
  func EmailConfirmationActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*EmailConfirmationEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &EmailConfirmationEntity{
          Status: &tildaRef,
          Email: &tildaRef,
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
  func EmailConfirmationAssociationCreate(dto *EmailConfirmationEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func EmailConfirmationRelationContentCreate(dto *EmailConfirmationEntity, query QueryDSL) error {
return nil
}
func EmailConfirmationRelationContentUpdate(dto *EmailConfirmationEntity, query QueryDSL) error {
	return nil
}
func EmailConfirmationPolyglotCreateHandler(dto *EmailConfirmationEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func EmailConfirmationValidator(dto *EmailConfirmationEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func EmailConfirmationEntityPreSanitize(dto *EmailConfirmationEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func EmailConfirmationEntityBeforeCreateAppend(dto *EmailConfirmationEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    EmailConfirmationRecursiveAddUniqueId(dto, query)
  }
  func EmailConfirmationRecursiveAddUniqueId(dto *EmailConfirmationEntity, query QueryDSL) {
  }
func EmailConfirmationActionBatchCreateFn(dtos []*EmailConfirmationEntity, query QueryDSL) ([]*EmailConfirmationEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*EmailConfirmationEntity{}
		for _, item := range dtos {
			s, err := EmailConfirmationActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func EmailConfirmationDeleteEntireChildren(query QueryDSL, dto *EmailConfirmationEntity) (*IError) {
  return nil
}
func EmailConfirmationActionCreateFn(dto *EmailConfirmationEntity, query QueryDSL) (*EmailConfirmationEntity, *IError) {
	// 1. Validate always
	if iError := EmailConfirmationValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	EmailConfirmationEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	EmailConfirmationEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	EmailConfirmationPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	EmailConfirmationRelationContentCreate(dto, query)
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
	EmailConfirmationAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(EMAILCONFIRMATION_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&EmailConfirmationEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func EmailConfirmationActionGetOne(query QueryDSL) (*EmailConfirmationEntity, *IError) {
    refl := reflect.ValueOf(&EmailConfirmationEntity{})
    item, err := GetOneEntity[EmailConfirmationEntity](query, refl)
    entityEmailConfirmationFormatter(item, query)
    return item, err
  }
  func EmailConfirmationActionQuery(query QueryDSL) ([]*EmailConfirmationEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&EmailConfirmationEntity{})
    items, meta, err := QueryEntitiesPointer[EmailConfirmationEntity](query, refl)
    for _, item := range items {
      entityEmailConfirmationFormatter(item, query)
    }
    return items, meta, err
  }
  func EmailConfirmationUpdateExec(dbref *gorm.DB, query QueryDSL, fields *EmailConfirmationEntity) (*EmailConfirmationEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = EMAILCONFIRMATION_EVENT_UPDATED
    EmailConfirmationEntityPreSanitize(fields, query)
    var item EmailConfirmationEntity
    q := dbref.
      Where(&EmailConfirmationEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    EmailConfirmationRelationContentUpdate(fields, query)
    EmailConfirmationPolyglotCreateHandler(fields, query)
    if ero := EmailConfirmationDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&EmailConfirmationEntity{UniqueId: uniqueId}).
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
  func EmailConfirmationActionUpdateFn(query QueryDSL, fields *EmailConfirmationEntity) (*EmailConfirmationEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := EmailConfirmationValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // EmailConfirmationRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      var item *EmailConfirmationEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *IError
        item, err = EmailConfirmationUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, CastToIError(vf)
    } else {
      dbref = query.Tx
      return EmailConfirmationUpdateExec(dbref, query, fields)
    }
  }
var EmailConfirmationWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire emailconfirmations ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
      ActionRequires: []string{PERM_ROOT_EMAILCONFIRMATION_DELETE},
    })
		count, _ := EmailConfirmationActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func EmailConfirmationActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&EmailConfirmationEntity{})
	query.ActionRequires = []string{PERM_ROOT_EMAILCONFIRMATION_DELETE}
	return RemoveEntity[EmailConfirmationEntity](query, refl)
}
func EmailConfirmationActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[EmailConfirmationEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'EmailConfirmationEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func EmailConfirmationActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[EmailConfirmationEntity]) (
    *BulkRecordRequest[EmailConfirmationEntity], *IError,
  ) {
    result := []*EmailConfirmationEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := EmailConfirmationActionUpdate(query, record)
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
func (x *EmailConfirmationEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var EmailConfirmationEntityMeta = TableMetaData{
	EntityName:    "EmailConfirmation",
	ExportKey:    "email-confirmations",
	TableNameInDb: "fb_emailconfirmation_entities",
	EntityObject:  &EmailConfirmationEntity{},
	ExportStream: EmailConfirmationActionExportT,
	ImportQuery: EmailConfirmationActionImport,
}
func EmailConfirmationActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[EmailConfirmationEntity](query, EmailConfirmationActionQuery, EmailConfirmationPreloadRelations)
}
func EmailConfirmationActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[EmailConfirmationEntity](query, EmailConfirmationActionQuery, EmailConfirmationPreloadRelations)
}
func EmailConfirmationActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content EmailConfirmationEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := EmailConfirmationActionCreate(&content, query)
	return err
}
var EmailConfirmationCommonCliFlags = []cli.Flag{
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
      Name:     "email",
      Required: false,
      Usage:    "email",
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
var EmailConfirmationCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:     "status",
		StructField:     "Status",
		Required: false,
		Usage:    "status",
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
var EmailConfirmationCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "email",
      Required: false,
      Usage:    "email",
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
  var EmailConfirmationCreateCmd cli.Command = EMAILCONFIRMATION_ACTION_POST_ONE.ToCli()
  var EmailConfirmationCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []string{PERM_ROOT_EMAILCONFIRMATION_CREATE},
      })
      entity := &EmailConfirmationEntity{}
      for _, item := range EmailConfirmationCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := EmailConfirmationActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var EmailConfirmationUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: EmailConfirmationCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []string{PERM_ROOT_EMAILCONFIRMATION_UPDATE},
      })
      entity := CastEmailConfirmationFromCli(c)
      if entity, err := EmailConfirmationActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* EmailConfirmationEntity) FromCli(c *cli.Context) *EmailConfirmationEntity {
	return CastEmailConfirmationFromCli(c)
}
func CastEmailConfirmationFromCli (c *cli.Context) *EmailConfirmationEntity {
	template := &EmailConfirmationEntity{}
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
      if c.IsSet("email") {
        value := c.String("email")
        template.Email = &value
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
  func EmailConfirmationSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      EmailConfirmationActionCreate,
      reflect.ValueOf(&EmailConfirmationEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func EmailConfirmationWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := EmailConfirmationActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "EmailConfirmation", result)
    }
  }
var EmailConfirmationImportExportCommands = []cli.Command{
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
        ActionRequires: []string{PERM_ROOT_EMAILCONFIRMATION_CREATE},
      })
			EmailConfirmationActionSeeder(query, c.Int("count"))
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
				Value: "email-confirmation-seeder.yml",
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
        ActionRequires: []string{PERM_ROOT_EMAILCONFIRMATION_CREATE},
      })
			EmailConfirmationActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "email-confirmation-seeder-email-confirmation.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of email-confirmations, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]EmailConfirmationEntity{}
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
			EmailConfirmationCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				EmailConfirmationActionCreate,
				reflect.ValueOf(&EmailConfirmationEntity{}).Elem(),
				c.String("file"),
        &SecurityModel{
					ActionRequires: []string{PERM_ROOT_EMAILCONFIRMATION_CREATE},
				},
        func() EmailConfirmationEntity {
					v := CastEmailConfirmationFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var EmailConfirmationCliCommands []cli.Command = []cli.Command{
      GetCommonQuery2(EmailConfirmationActionQuery, &SecurityModel{
        ActionRequires: []string{PERM_ROOT_EMAILCONFIRMATION_CREATE},
      }),
      GetCommonTableQuery(reflect.ValueOf(&EmailConfirmationEntity{}).Elem(), EmailConfirmationActionQuery),
          EmailConfirmationCreateCmd,
          EmailConfirmationUpdateCmd,
          EmailConfirmationCreateInteractiveCmd,
          EmailConfirmationWipeCmd,
          GetCommonRemoveQuery(reflect.ValueOf(&EmailConfirmationEntity{}).Elem(), EmailConfirmationActionRemove),
  }
  func EmailConfirmationCliFn() cli.Command {
    EmailConfirmationCliCommands = append(EmailConfirmationCliCommands, EmailConfirmationImportExportCommands...)
    return cli.Command{
      Name:        "emailConfirmation",
      Description: "EmailConfirmations module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: EmailConfirmationCliCommands,
    }
  }
var EMAILCONFIRMATION_ACTION_POST_ONE = Module2Action{
    ActionName:    "create",
    ActionAliases: []string{"c"},
    Description: "Create new emailConfirmation",
    Flags: EmailConfirmationCommonCliFlags,
    Method: "POST",
    Url:    "/email-confirmation",
    SecurityModel: &SecurityModel{
      ActionRequires: []string{PERM_ROOT_EMAILCONFIRMATION_CREATE},
    },
    Handlers: []gin.HandlerFunc{
      func (c *gin.Context) {
        HttpPostEntity(c, EmailConfirmationActionCreate)
      },
    },
    CliAction: func(c *cli.Context, security *SecurityModel) error {
      result, err := CliPostEntity(c, EmailConfirmationActionCreate, security)
      HandleActionInCli(c, result, err, map[string]map[string]string{})
      return err
    },
    Action: EmailConfirmationActionCreate,
    Format: "POST_ONE",
    RequestEntity: &EmailConfirmationEntity{},
    ResponseEntity: &EmailConfirmationEntity{},
  }
  /**
  *	Override this function on EmailConfirmationEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendEmailConfirmationRouter = func(r *[]Module2Action) {}
  func GetEmailConfirmationModule2Actions() []Module2Action {
    routes := []Module2Action{
       {
        Method: "GET",
        Url:    "/email-confirmations",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_EMAILCONFIRMATION_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpQueryEntity(c, EmailConfirmationActionQuery)
          },
        },
        Format: "QUERY",
        Action: EmailConfirmationActionQuery,
        ResponseEntity: &[]EmailConfirmationEntity{},
      },
      {
        Method: "GET",
        Url:    "/email-confirmations/export",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_EMAILCONFIRMATION_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpStreamFileChannel(c, EmailConfirmationActionExport)
          },
        },
        Format: "QUERY",
        Action: EmailConfirmationActionExport,
        ResponseEntity: &[]EmailConfirmationEntity{},
      },
      {
        Method: "GET",
        Url:    "/email-confirmation/:uniqueId",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_EMAILCONFIRMATION_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpGetEntity(c, EmailConfirmationActionGetOne)
          },
        },
        Format: "GET_ONE",
        Action: EmailConfirmationActionGetOne,
        ResponseEntity: &EmailConfirmationEntity{},
      },
      EMAILCONFIRMATION_ACTION_POST_ONE,
      {
        ActionName:    "update",
        ActionAliases: []string{"u"},
        Flags: EmailConfirmationCommonCliFlagsOptional,
        Method: "PATCH",
        Url:    "/email-confirmation",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_EMAILCONFIRMATION_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpUpdateEntity(c, EmailConfirmationActionUpdate)
          },
        },
        Action: EmailConfirmationActionUpdate,
        RequestEntity: &EmailConfirmationEntity{},
        Format: "PATCH_ONE",
        ResponseEntity: &EmailConfirmationEntity{},
      },
      {
        Method: "PATCH",
        Url:    "/email-confirmations",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_EMAILCONFIRMATION_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpUpdateEntities(c, EmailConfirmationActionBulkUpdate)
          },
        },
        Action: EmailConfirmationActionBulkUpdate,
        Format: "PATCH_BULK",
        RequestEntity:  &BulkRecordRequest[EmailConfirmationEntity]{},
        ResponseEntity: &BulkRecordRequest[EmailConfirmationEntity]{},
      },
      {
        Method: "DELETE",
        Url:    "/email-confirmation",
        Format: "DELETE_DSL",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_EMAILCONFIRMATION_DELETE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpRemoveEntity(c, EmailConfirmationActionRemove)
          },
        },
        Action: EmailConfirmationActionRemove,
        RequestEntity: &DeleteRequest{},
        ResponseEntity: &DeleteResponse{},
        TargetEntity: &EmailConfirmationEntity{},
      },
    }
    // Append user defined functions
    AppendEmailConfirmationRouter(&routes)
    return routes
  }
  func CreateEmailConfirmationRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetEmailConfirmationModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, EmailConfirmationEntityJsonSchema, "email-confirmation-http", "workspaces")
    WriteEntitySchema("EmailConfirmationEntity", EmailConfirmationEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_EMAILCONFIRMATION_DELETE = "root/emailconfirmation/delete"
var PERM_ROOT_EMAILCONFIRMATION_CREATE = "root/emailconfirmation/create"
var PERM_ROOT_EMAILCONFIRMATION_UPDATE = "root/emailconfirmation/update"
var PERM_ROOT_EMAILCONFIRMATION_QUERY = "root/emailconfirmation/query"
var PERM_ROOT_EMAILCONFIRMATION = "root/emailconfirmation"
var ALL_EMAILCONFIRMATION_PERMISSIONS = []string{
	PERM_ROOT_EMAILCONFIRMATION_DELETE,
	PERM_ROOT_EMAILCONFIRMATION_CREATE,
	PERM_ROOT_EMAILCONFIRMATION_UPDATE,
	PERM_ROOT_EMAILCONFIRMATION_QUERY,
	PERM_ROOT_EMAILCONFIRMATION,
}