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
type EmailSenderEntity struct {
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
    FromName   *string `json:"fromName" yaml:"fromName"  validate:"required"       `
    // Datenano also has a text representation
    FromEmailAddress   *string `json:"fromEmailAddress" yaml:"fromEmailAddress"  validate:"required"    gorm:"unique"     `
    // Datenano also has a text representation
    ReplyTo   *string `json:"replyTo" yaml:"replyTo"  validate:"required"       `
    // Datenano also has a text representation
    NickName   *string `json:"nickName" yaml:"nickName"  validate:"required"       `
    // Datenano also has a text representation
    Children []*EmailSenderEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *EmailSenderEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var EmailSenderPreloadRelations []string = []string{}
var EMAILSENDER_EVENT_CREATED = "emailSender.created"
var EMAILSENDER_EVENT_UPDATED = "emailSender.updated"
var EMAILSENDER_EVENT_DELETED = "emailSender.deleted"
var EMAILSENDER_EVENTS = []string{
	EMAILSENDER_EVENT_CREATED,
	EMAILSENDER_EVENT_UPDATED,
	EMAILSENDER_EVENT_DELETED,
}
type EmailSenderFieldMap struct {
		FromName TranslatedString `yaml:"fromName"`
		FromEmailAddress TranslatedString `yaml:"fromEmailAddress"`
		ReplyTo TranslatedString `yaml:"replyTo"`
		NickName TranslatedString `yaml:"nickName"`
}
var EmailSenderEntityMetaConfig map[string]int64 = map[string]int64{
}
var EmailSenderEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&EmailSenderEntity{}))
func entityEmailSenderFormatter(dto *EmailSenderEntity, query QueryDSL) {
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
func EmailSenderMockEntity() *EmailSenderEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &EmailSenderEntity{
      FromName : &stringHolder,
      FromEmailAddress : &stringHolder,
      ReplyTo : &stringHolder,
      NickName : &stringHolder,
	}
	return entity
}
func EmailSenderActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := EmailSenderMockEntity()
		_, err := EmailSenderActionCreate(entity, query)
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
  func EmailSenderActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*EmailSenderEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &EmailSenderEntity{
          FromName: &tildaRef,
          FromEmailAddress: &tildaRef,
          ReplyTo: &tildaRef,
          NickName: &tildaRef,
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
  func EmailSenderAssociationCreate(dto *EmailSenderEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func EmailSenderRelationContentCreate(dto *EmailSenderEntity, query QueryDSL) error {
return nil
}
func EmailSenderRelationContentUpdate(dto *EmailSenderEntity, query QueryDSL) error {
	return nil
}
func EmailSenderPolyglotCreateHandler(dto *EmailSenderEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func EmailSenderValidator(dto *EmailSenderEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func EmailSenderEntityPreSanitize(dto *EmailSenderEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func EmailSenderEntityBeforeCreateAppend(dto *EmailSenderEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    EmailSenderRecursiveAddUniqueId(dto, query)
  }
  func EmailSenderRecursiveAddUniqueId(dto *EmailSenderEntity, query QueryDSL) {
  }
func EmailSenderActionBatchCreateFn(dtos []*EmailSenderEntity, query QueryDSL) ([]*EmailSenderEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*EmailSenderEntity{}
		for _, item := range dtos {
			s, err := EmailSenderActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func EmailSenderDeleteEntireChildren(query QueryDSL, dto *EmailSenderEntity) (*IError) {
  return nil
}
func EmailSenderActionCreateFn(dto *EmailSenderEntity, query QueryDSL) (*EmailSenderEntity, *IError) {
	// 1. Validate always
	if iError := EmailSenderValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	EmailSenderEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	EmailSenderEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	EmailSenderPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	EmailSenderRelationContentCreate(dto, query)
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
	EmailSenderAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(EMAILSENDER_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&EmailSenderEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func EmailSenderActionGetOne(query QueryDSL) (*EmailSenderEntity, *IError) {
    refl := reflect.ValueOf(&EmailSenderEntity{})
    item, err := GetOneEntity[EmailSenderEntity](query, refl)
    entityEmailSenderFormatter(item, query)
    return item, err
  }
  func EmailSenderActionQuery(query QueryDSL) ([]*EmailSenderEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&EmailSenderEntity{})
    items, meta, err := QueryEntitiesPointer[EmailSenderEntity](query, refl)
    for _, item := range items {
      entityEmailSenderFormatter(item, query)
    }
    return items, meta, err
  }
  func EmailSenderUpdateExec(dbref *gorm.DB, query QueryDSL, fields *EmailSenderEntity) (*EmailSenderEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = EMAILSENDER_EVENT_UPDATED
    EmailSenderEntityPreSanitize(fields, query)
    var item EmailSenderEntity
    q := dbref.
      Where(&EmailSenderEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    EmailSenderRelationContentUpdate(fields, query)
    EmailSenderPolyglotCreateHandler(fields, query)
    if ero := EmailSenderDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&EmailSenderEntity{UniqueId: uniqueId}).
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
  func EmailSenderActionUpdateFn(query QueryDSL, fields *EmailSenderEntity) (*EmailSenderEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := EmailSenderValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // EmailSenderRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      var item *EmailSenderEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *IError
        item, err = EmailSenderUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, CastToIError(vf)
    } else {
      dbref = query.Tx
      return EmailSenderUpdateExec(dbref, query, fields)
    }
  }
var EmailSenderWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire emailsenders ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
      ActionRequires: []string{PERM_ROOT_EMAILSENDER_DELETE},
    })
		count, _ := EmailSenderActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func EmailSenderActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&EmailSenderEntity{})
	query.ActionRequires = []string{PERM_ROOT_EMAILSENDER_DELETE}
	return RemoveEntity[EmailSenderEntity](query, refl)
}
func EmailSenderActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[EmailSenderEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'EmailSenderEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func EmailSenderActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[EmailSenderEntity]) (
    *BulkRecordRequest[EmailSenderEntity], *IError,
  ) {
    result := []*EmailSenderEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := EmailSenderActionUpdate(query, record)
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
func (x *EmailSenderEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var EmailSenderEntityMeta = TableMetaData{
	EntityName:    "EmailSender",
	ExportKey:    "email-senders",
	TableNameInDb: "fb_emailsender_entities",
	EntityObject:  &EmailSenderEntity{},
	ExportStream: EmailSenderActionExportT,
	ImportQuery: EmailSenderActionImport,
}
func EmailSenderActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[EmailSenderEntity](query, EmailSenderActionQuery, EmailSenderPreloadRelations)
}
func EmailSenderActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[EmailSenderEntity](query, EmailSenderActionQuery, EmailSenderPreloadRelations)
}
func EmailSenderActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content EmailSenderEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := EmailSenderActionCreate(&content, query)
	return err
}
var EmailSenderCommonCliFlags = []cli.Flag{
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
      Name:     "from-name",
      Required: true,
      Usage:    "fromName",
    },
    &cli.StringFlag{
      Name:     "from-email-address",
      Required: true,
      Usage:    "fromEmailAddress",
    },
    &cli.StringFlag{
      Name:     "reply-to",
      Required: true,
      Usage:    "replyTo",
    },
    &cli.StringFlag{
      Name:     "nick-name",
      Required: true,
      Usage:    "nickName",
    },
}
var EmailSenderCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:     "fromName",
		StructField:     "FromName",
		Required: true,
		Usage:    "fromName",
		Type: "string",
	},
	{
		Name:     "fromEmailAddress",
		StructField:     "FromEmailAddress",
		Required: true,
		Usage:    "fromEmailAddress",
		Type: "string",
	},
	{
		Name:     "replyTo",
		StructField:     "ReplyTo",
		Required: true,
		Usage:    "replyTo",
		Type: "string",
	},
	{
		Name:     "nickName",
		StructField:     "NickName",
		Required: true,
		Usage:    "nickName",
		Type: "string",
	},
}
var EmailSenderCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "from-name",
      Required: true,
      Usage:    "fromName",
    },
    &cli.StringFlag{
      Name:     "from-email-address",
      Required: true,
      Usage:    "fromEmailAddress",
    },
    &cli.StringFlag{
      Name:     "reply-to",
      Required: true,
      Usage:    "replyTo",
    },
    &cli.StringFlag{
      Name:     "nick-name",
      Required: true,
      Usage:    "nickName",
    },
}
  var EmailSenderCreateCmd cli.Command = EMAILSENDER_ACTION_POST_ONE.ToCli()
  var EmailSenderCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []string{PERM_ROOT_EMAILSENDER_CREATE},
      })
      entity := &EmailSenderEntity{}
      for _, item := range EmailSenderCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := EmailSenderActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var EmailSenderUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: EmailSenderCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []string{PERM_ROOT_EMAILSENDER_UPDATE},
      })
      entity := CastEmailSenderFromCli(c)
      if entity, err := EmailSenderActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* EmailSenderEntity) FromCli(c *cli.Context) *EmailSenderEntity {
	return CastEmailSenderFromCli(c)
}
func CastEmailSenderFromCli (c *cli.Context) *EmailSenderEntity {
	template := &EmailSenderEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("from-name") {
        value := c.String("from-name")
        template.FromName = &value
      }
      if c.IsSet("from-email-address") {
        value := c.String("from-email-address")
        template.FromEmailAddress = &value
      }
      if c.IsSet("reply-to") {
        value := c.String("reply-to")
        template.ReplyTo = &value
      }
      if c.IsSet("nick-name") {
        value := c.String("nick-name")
        template.NickName = &value
      }
	return template
}
  func EmailSenderSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      EmailSenderActionCreate,
      reflect.ValueOf(&EmailSenderEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func EmailSenderWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := EmailSenderActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "EmailSender", result)
    }
  }
var EmailSenderImportExportCommands = []cli.Command{
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
        ActionRequires: []string{PERM_ROOT_EMAILSENDER_CREATE},
      })
			EmailSenderActionSeeder(query, c.Int("count"))
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
				Value: "email-sender-seeder.yml",
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
        ActionRequires: []string{PERM_ROOT_EMAILSENDER_CREATE},
      })
			EmailSenderActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "email-sender-seeder-email-sender.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of email-senders, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]EmailSenderEntity{}
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
			EmailSenderCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				EmailSenderActionCreate,
				reflect.ValueOf(&EmailSenderEntity{}).Elem(),
				c.String("file"),
        &SecurityModel{
					ActionRequires: []string{PERM_ROOT_EMAILSENDER_CREATE},
				},
        func() EmailSenderEntity {
					v := CastEmailSenderFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var EmailSenderCliCommands []cli.Command = []cli.Command{
      GetCommonQuery2(EmailSenderActionQuery, &SecurityModel{
        ActionRequires: []string{PERM_ROOT_EMAILSENDER_CREATE},
      }),
      GetCommonTableQuery(reflect.ValueOf(&EmailSenderEntity{}).Elem(), EmailSenderActionQuery),
          EmailSenderCreateCmd,
          EmailSenderUpdateCmd,
          EmailSenderCreateInteractiveCmd,
          EmailSenderWipeCmd,
          GetCommonRemoveQuery(reflect.ValueOf(&EmailSenderEntity{}).Elem(), EmailSenderActionRemove),
  }
  func EmailSenderCliFn() cli.Command {
    EmailSenderCliCommands = append(EmailSenderCliCommands, EmailSenderImportExportCommands...)
    return cli.Command{
      Name:        "emailSender",
      Description: "EmailSenders module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: EmailSenderCliCommands,
    }
  }
var EMAILSENDER_ACTION_POST_ONE = Module2Action{
    ActionName:    "create",
    ActionAliases: []string{"c"},
    Description: "Create new emailSender",
    Flags: EmailSenderCommonCliFlags,
    Method: "POST",
    Url:    "/email-sender",
    SecurityModel: &SecurityModel{
      ActionRequires: []string{PERM_ROOT_EMAILSENDER_CREATE},
    },
    Handlers: []gin.HandlerFunc{
      func (c *gin.Context) {
        HttpPostEntity(c, EmailSenderActionCreate)
      },
    },
    CliAction: func(c *cli.Context, security *SecurityModel) error {
      result, err := CliPostEntity(c, EmailSenderActionCreate, security)
      HandleActionInCli(c, result, err, map[string]map[string]string{})
      return err
    },
    Action: EmailSenderActionCreate,
    Format: "POST_ONE",
    RequestEntity: &EmailSenderEntity{},
    ResponseEntity: &EmailSenderEntity{},
  }
  /**
  *	Override this function on EmailSenderEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendEmailSenderRouter = func(r *[]Module2Action) {}
  func GetEmailSenderModule2Actions() []Module2Action {
    routes := []Module2Action{
       {
        Method: "GET",
        Url:    "/email-senders",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_EMAILSENDER_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpQueryEntity(c, EmailSenderActionQuery)
          },
        },
        Format: "QUERY",
        Action: EmailSenderActionQuery,
        ResponseEntity: &[]EmailSenderEntity{},
      },
      {
        Method: "GET",
        Url:    "/email-senders/export",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_EMAILSENDER_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpStreamFileChannel(c, EmailSenderActionExport)
          },
        },
        Format: "QUERY",
        Action: EmailSenderActionExport,
        ResponseEntity: &[]EmailSenderEntity{},
      },
      {
        Method: "GET",
        Url:    "/email-sender/:uniqueId",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_EMAILSENDER_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpGetEntity(c, EmailSenderActionGetOne)
          },
        },
        Format: "GET_ONE",
        Action: EmailSenderActionGetOne,
        ResponseEntity: &EmailSenderEntity{},
      },
      EMAILSENDER_ACTION_POST_ONE,
      {
        ActionName:    "update",
        ActionAliases: []string{"u"},
        Flags: EmailSenderCommonCliFlagsOptional,
        Method: "PATCH",
        Url:    "/email-sender",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_EMAILSENDER_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpUpdateEntity(c, EmailSenderActionUpdate)
          },
        },
        Action: EmailSenderActionUpdate,
        RequestEntity: &EmailSenderEntity{},
        Format: "PATCH_ONE",
        ResponseEntity: &EmailSenderEntity{},
      },
      {
        Method: "PATCH",
        Url:    "/email-senders",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_EMAILSENDER_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpUpdateEntities(c, EmailSenderActionBulkUpdate)
          },
        },
        Action: EmailSenderActionBulkUpdate,
        Format: "PATCH_BULK",
        RequestEntity:  &BulkRecordRequest[EmailSenderEntity]{},
        ResponseEntity: &BulkRecordRequest[EmailSenderEntity]{},
      },
      {
        Method: "DELETE",
        Url:    "/email-sender",
        Format: "DELETE_DSL",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_EMAILSENDER_DELETE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpRemoveEntity(c, EmailSenderActionRemove)
          },
        },
        Action: EmailSenderActionRemove,
        RequestEntity: &DeleteRequest{},
        ResponseEntity: &DeleteResponse{},
        TargetEntity: &EmailSenderEntity{},
      },
    }
    // Append user defined functions
    AppendEmailSenderRouter(&routes)
    return routes
  }
  func CreateEmailSenderRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetEmailSenderModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, EmailSenderEntityJsonSchema, "email-sender-http", "workspaces")
    WriteEntitySchema("EmailSenderEntity", EmailSenderEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_EMAILSENDER_DELETE = "root/emailsender/delete"
var PERM_ROOT_EMAILSENDER_CREATE = "root/emailsender/create"
var PERM_ROOT_EMAILSENDER_UPDATE = "root/emailsender/update"
var PERM_ROOT_EMAILSENDER_QUERY = "root/emailsender/query"
var PERM_ROOT_EMAILSENDER = "root/emailsender"
var ALL_EMAILSENDER_PERMISSIONS = []string{
	PERM_ROOT_EMAILSENDER_DELETE,
	PERM_ROOT_EMAILSENDER_CREATE,
	PERM_ROOT_EMAILSENDER_UPDATE,
	PERM_ROOT_EMAILSENDER_QUERY,
	PERM_ROOT_EMAILSENDER,
}