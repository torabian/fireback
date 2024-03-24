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
type EmailProviderEntity struct {
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
    Type   *string `json:"type" yaml:"type"  validate:"required"       `
    // Datenano also has a text representation
    ApiKey   *string `json:"apiKey" yaml:"apiKey"       `
    // Datenano also has a text representation
    Children []*EmailProviderEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *EmailProviderEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var EmailProviderPreloadRelations []string = []string{}
var EMAIL_PROVIDER_EVENT_CREATED = "emailProvider.created"
var EMAIL_PROVIDER_EVENT_UPDATED = "emailProvider.updated"
var EMAIL_PROVIDER_EVENT_DELETED = "emailProvider.deleted"
var EMAIL_PROVIDER_EVENTS = []string{
	EMAIL_PROVIDER_EVENT_CREATED,
	EMAIL_PROVIDER_EVENT_UPDATED,
	EMAIL_PROVIDER_EVENT_DELETED,
}
type EmailProviderFieldMap struct {
		Type TranslatedString `yaml:"type"`
		ApiKey TranslatedString `yaml:"apiKey"`
}
var EmailProviderEntityMetaConfig map[string]int64 = map[string]int64{
}
var EmailProviderEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&EmailProviderEntity{}))
func entityEmailProviderFormatter(dto *EmailProviderEntity, query QueryDSL) {
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
func EmailProviderMockEntity() *EmailProviderEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &EmailProviderEntity{
      Type : &stringHolder,
      ApiKey : &stringHolder,
	}
	return entity
}
func EmailProviderActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := EmailProviderMockEntity()
		_, err := EmailProviderActionCreate(entity, query)
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
  func EmailProviderActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*EmailProviderEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &EmailProviderEntity{
          Type: &tildaRef,
          ApiKey: &tildaRef,
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
  func EmailProviderAssociationCreate(dto *EmailProviderEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func EmailProviderRelationContentCreate(dto *EmailProviderEntity, query QueryDSL) error {
return nil
}
func EmailProviderRelationContentUpdate(dto *EmailProviderEntity, query QueryDSL) error {
	return nil
}
func EmailProviderPolyglotCreateHandler(dto *EmailProviderEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func EmailProviderValidator(dto *EmailProviderEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func EmailProviderEntityPreSanitize(dto *EmailProviderEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func EmailProviderEntityBeforeCreateAppend(dto *EmailProviderEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    EmailProviderRecursiveAddUniqueId(dto, query)
  }
  func EmailProviderRecursiveAddUniqueId(dto *EmailProviderEntity, query QueryDSL) {
  }
func EmailProviderActionBatchCreateFn(dtos []*EmailProviderEntity, query QueryDSL) ([]*EmailProviderEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*EmailProviderEntity{}
		for _, item := range dtos {
			s, err := EmailProviderActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func EmailProviderDeleteEntireChildren(query QueryDSL, dto *EmailProviderEntity) (*IError) {
  return nil
}
func EmailProviderActionCreateFn(dto *EmailProviderEntity, query QueryDSL) (*EmailProviderEntity, *IError) {
	// 1. Validate always
	if iError := EmailProviderValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	EmailProviderEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	EmailProviderEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	EmailProviderPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	EmailProviderRelationContentCreate(dto, query)
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
	EmailProviderAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(EMAIL_PROVIDER_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&EmailProviderEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func EmailProviderActionGetOne(query QueryDSL) (*EmailProviderEntity, *IError) {
    refl := reflect.ValueOf(&EmailProviderEntity{})
    item, err := GetOneEntity[EmailProviderEntity](query, refl)
    entityEmailProviderFormatter(item, query)
    return item, err
  }
  func EmailProviderActionQuery(query QueryDSL) ([]*EmailProviderEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&EmailProviderEntity{})
    items, meta, err := QueryEntitiesPointer[EmailProviderEntity](query, refl)
    for _, item := range items {
      entityEmailProviderFormatter(item, query)
    }
    return items, meta, err
  }
  func EmailProviderUpdateExec(dbref *gorm.DB, query QueryDSL, fields *EmailProviderEntity) (*EmailProviderEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = EMAIL_PROVIDER_EVENT_UPDATED
    EmailProviderEntityPreSanitize(fields, query)
    var item EmailProviderEntity
    q := dbref.
      Where(&EmailProviderEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    EmailProviderRelationContentUpdate(fields, query)
    EmailProviderPolyglotCreateHandler(fields, query)
    if ero := EmailProviderDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&EmailProviderEntity{UniqueId: uniqueId}).
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
  func EmailProviderActionUpdateFn(query QueryDSL, fields *EmailProviderEntity) (*EmailProviderEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := EmailProviderValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // EmailProviderRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      var item *EmailProviderEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *IError
        item, err = EmailProviderUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, CastToIError(vf)
    } else {
      dbref = query.Tx
      return EmailProviderUpdateExec(dbref, query, fields)
    }
  }
var EmailProviderWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire emailproviders ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
      ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_PROVIDER_DELETE},
    })
		count, _ := EmailProviderActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func EmailProviderActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&EmailProviderEntity{})
	query.ActionRequires = []PermissionInfo{PERM_ROOT_EMAIL_PROVIDER_DELETE}
	return RemoveEntity[EmailProviderEntity](query, refl)
}
func EmailProviderActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[EmailProviderEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'EmailProviderEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func EmailProviderActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[EmailProviderEntity]) (
    *BulkRecordRequest[EmailProviderEntity], *IError,
  ) {
    result := []*EmailProviderEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := EmailProviderActionUpdate(query, record)
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
func (x *EmailProviderEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var EmailProviderEntityMeta = TableMetaData{
	EntityName:    "EmailProvider",
	ExportKey:    "email-providers",
	TableNameInDb: "fb_email-provider_entities",
	EntityObject:  &EmailProviderEntity{},
	ExportStream: EmailProviderActionExportT,
	ImportQuery: EmailProviderActionImport,
}
func EmailProviderActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[EmailProviderEntity](query, EmailProviderActionQuery, EmailProviderPreloadRelations)
}
func EmailProviderActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[EmailProviderEntity](query, EmailProviderActionQuery, EmailProviderPreloadRelations)
}
func EmailProviderActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content EmailProviderEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := EmailProviderActionCreate(&content, query)
	return err
}
var EmailProviderCommonCliFlags = []cli.Flag{
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
      Name:     "type",
      Required: true,
      Usage:    "One of: 'terminal', 'sendgrid'",
    },
    &cli.StringFlag{
      Name:     "api-key",
      Required: false,
      Usage:    "apiKey",
    },
}
var EmailProviderCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:     "type",
		StructField:     "Type",
		Required: true,
		Usage:    "One of: 'terminal', 'sendgrid'",
		Type: "string",
	},
	{
		Name:     "apiKey",
		StructField:     "ApiKey",
		Required: false,
		Usage:    "apiKey",
		Type: "string",
	},
}
var EmailProviderCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "type",
      Required: true,
      Usage:    "One of: 'terminal', 'sendgrid'",
    },
    &cli.StringFlag{
      Name:     "api-key",
      Required: false,
      Usage:    "apiKey",
    },
}
  var EmailProviderCreateCmd cli.Command = EMAIL_PROVIDER_ACTION_POST_ONE.ToCli()
  var EmailProviderCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_PROVIDER_CREATE},
      })
      entity := &EmailProviderEntity{}
      for _, item := range EmailProviderCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := EmailProviderActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var EmailProviderUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: EmailProviderCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_PROVIDER_UPDATE},
      })
      entity := CastEmailProviderFromCli(c)
      if entity, err := EmailProviderActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* EmailProviderEntity) FromCli(c *cli.Context) *EmailProviderEntity {
	return CastEmailProviderFromCli(c)
}
func CastEmailProviderFromCli (c *cli.Context) *EmailProviderEntity {
	template := &EmailProviderEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("type") {
        value := c.String("type")
        template.Type = &value
      }
      if c.IsSet("api-key") {
        value := c.String("api-key")
        template.ApiKey = &value
      }
	return template
}
  func EmailProviderSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      EmailProviderActionCreate,
      reflect.ValueOf(&EmailProviderEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func EmailProviderWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := EmailProviderActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "EmailProvider", result)
    }
  }
var EmailProviderImportExportCommands = []cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_PROVIDER_CREATE},
      })
			EmailProviderActionSeeder(query, c.Int("count"))
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
				Value: "email-provider-seeder.yml",
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
        ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_PROVIDER_CREATE},
      })
			EmailProviderActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "email-provider-seeder-email-provider.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of email-providers, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]EmailProviderEntity{}
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
			EmailProviderCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				EmailProviderActionCreate,
				reflect.ValueOf(&EmailProviderEntity{}).Elem(),
				c.String("file"),
        &SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_PROVIDER_CREATE},
				},
        func() EmailProviderEntity {
					v := CastEmailProviderFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var EmailProviderCliCommands []cli.Command = []cli.Command{
      GetCommonQuery2(EmailProviderActionQuery, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_PROVIDER_QUERY},
      }),
      GetCommonTableQuery(reflect.ValueOf(&EmailProviderEntity{}).Elem(), EmailProviderActionQuery),
          EmailProviderCreateCmd,
          EmailProviderUpdateCmd,
          EmailProviderCreateInteractiveCmd,
          EmailProviderWipeCmd,
          GetCommonRemoveQuery(reflect.ValueOf(&EmailProviderEntity{}).Elem(), EmailProviderActionRemove),
  }
  func EmailProviderCliFn() cli.Command {
    EmailProviderCliCommands = append(EmailProviderCliCommands, EmailProviderImportExportCommands...)
    return cli.Command{
      Name:        "emailProvider",
      Description: "EmailProviders module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: EmailProviderCliCommands,
    }
  }
var EMAIL_PROVIDER_ACTION_QUERY = Module2Action{
  Method: "GET",
  Url:    "/email-providers",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_PROVIDER_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpQueryEntity(c, EmailProviderActionQuery)
    },
  },
  Format: "QUERY",
  Action: EmailProviderActionQuery,
  ResponseEntity: &[]EmailProviderEntity{},
}
var EMAIL_PROVIDER_ACTION_EXPORT = Module2Action{
  Method: "GET",
  Url:    "/email-providers/export",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_PROVIDER_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpStreamFileChannel(c, EmailProviderActionExport)
    },
  },
  Format: "QUERY",
  Action: EmailProviderActionExport,
  ResponseEntity: &[]EmailProviderEntity{},
}
var EMAIL_PROVIDER_ACTION_GET_ONE = Module2Action{
  Method: "GET",
  Url:    "/email-provider/:uniqueId",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_PROVIDER_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpGetEntity(c, EmailProviderActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: EmailProviderActionGetOne,
  ResponseEntity: &EmailProviderEntity{},
}
var EMAIL_PROVIDER_ACTION_POST_ONE = Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new emailProvider",
  Flags: EmailProviderCommonCliFlags,
  Method: "POST",
  Url:    "/email-provider",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_PROVIDER_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpPostEntity(c, EmailProviderActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    result, err := CliPostEntity(c, EmailProviderActionCreate, security)
    HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: EmailProviderActionCreate,
  Format: "POST_ONE",
  RequestEntity: &EmailProviderEntity{},
  ResponseEntity: &EmailProviderEntity{},
}
var EMAIL_PROVIDER_ACTION_PATCH = Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: EmailProviderCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/email-provider",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_PROVIDER_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntity(c, EmailProviderActionUpdate)
    },
  },
  Action: EmailProviderActionUpdate,
  RequestEntity: &EmailProviderEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &EmailProviderEntity{},
}
var EMAIL_PROVIDER_ACTION_PATCH_BULK = Module2Action{
  Method: "PATCH",
  Url:    "/email-providers",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_PROVIDER_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntities(c, EmailProviderActionBulkUpdate)
    },
  },
  Action: EmailProviderActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &BulkRecordRequest[EmailProviderEntity]{},
  ResponseEntity: &BulkRecordRequest[EmailProviderEntity]{},
}
var EMAIL_PROVIDER_ACTION_DELETE = Module2Action{
  Method: "DELETE",
  Url:    "/email-provider",
  Format: "DELETE_DSL",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_PROVIDER_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpRemoveEntity(c, EmailProviderActionRemove)
    },
  },
  Action: EmailProviderActionRemove,
  RequestEntity: &DeleteRequest{},
  ResponseEntity: &DeleteResponse{},
  TargetEntity: &EmailProviderEntity{},
}
  /**
  *	Override this function on EmailProviderEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendEmailProviderRouter = func(r *[]Module2Action) {}
  func GetEmailProviderModule2Actions() []Module2Action {
    routes := []Module2Action{
      EMAIL_PROVIDER_ACTION_QUERY,
      EMAIL_PROVIDER_ACTION_EXPORT,
      EMAIL_PROVIDER_ACTION_GET_ONE,
      EMAIL_PROVIDER_ACTION_POST_ONE,
      EMAIL_PROVIDER_ACTION_PATCH,
      EMAIL_PROVIDER_ACTION_PATCH_BULK,
      EMAIL_PROVIDER_ACTION_DELETE,
    }
    // Append user defined functions
    AppendEmailProviderRouter(&routes)
    return routes
  }
  func CreateEmailProviderRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetEmailProviderModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, EmailProviderEntityJsonSchema, "email-provider-http", "workspaces")
    WriteEntitySchema("EmailProviderEntity", EmailProviderEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_EMAIL_PROVIDER_DELETE = PermissionInfo{
  CompleteKey: "root/workspaces/email-provider/delete",
}
var PERM_ROOT_EMAIL_PROVIDER_CREATE = PermissionInfo{
  CompleteKey: "root/workspaces/email-provider/create",
}
var PERM_ROOT_EMAIL_PROVIDER_UPDATE = PermissionInfo{
  CompleteKey: "root/workspaces/email-provider/update",
}
var PERM_ROOT_EMAIL_PROVIDER_QUERY = PermissionInfo{
  CompleteKey: "root/workspaces/email-provider/query",
}
var PERM_ROOT_EMAIL_PROVIDER = PermissionInfo{
  CompleteKey: "root/workspaces/email-provider/*",
}
var ALL_EMAIL_PROVIDER_PERMISSIONS = []PermissionInfo{
	PERM_ROOT_EMAIL_PROVIDER_DELETE,
	PERM_ROOT_EMAIL_PROVIDER_CREATE,
	PERM_ROOT_EMAIL_PROVIDER_UPDATE,
	PERM_ROOT_EMAIL_PROVIDER_QUERY,
	PERM_ROOT_EMAIL_PROVIDER,
}
var EmailProviderType = newEmailProviderType()
func newEmailProviderType() *xEmailProviderType {
	return &xEmailProviderType{
      Terminal: "terminal",
      Sendgrid: "sendgrid",
	}
}
type xEmailProviderType struct {
    Terminal string
    Sendgrid string
}