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
type ForgetPasswordEntity struct {
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
    User   *  UserEntity `json:"false" yaml:"user"    gorm:"foreignKey:UserId;references:UniqueId"     `
    // Datenano also has a text representation
    Passport   *  PassportEntity `json:"false" yaml:"passport"    gorm:"foreignKey:PassportId;references:UniqueId"     `
    // Datenano also has a text representation
        PassportId *string `json:"passportId" yaml:"passportId"`
    Status   *string `json:"false" yaml:"status"       `
    // Datenano also has a text representation
    ValidUntil   int64 `json:"validUntil" yaml:"validUntil"       `
    // Datenano also has a text representation
    ValidUntilFormatted string `json:"validUntilFormatted" yaml:"validUntilFormatted"`
    BlockedUntil   int64 `json:"blockedUntil" yaml:"blockedUntil"       `
    // Datenano also has a text representation
    BlockedUntilFormatted string `json:"blockedUntilFormatted" yaml:"blockedUntilFormatted"`
    SecondsToUnblock   *int64 `json:"secondsToUnblock" yaml:"secondsToUnblock"       `
    // Datenano also has a text representation
    Otp   *string `json:"false" yaml:"otp"       `
    // Datenano also has a text representation
    RecoveryAbsoluteUrl   *string `json:"false" yaml:"recoveryAbsoluteUrl"       sql:"false"  `
    // Datenano also has a text representation
    Children []*ForgetPasswordEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *ForgetPasswordEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var ForgetPasswordPreloadRelations []string = []string{}
var FORGET_PASSWORD_EVENT_CREATED = "forgetPassword.created"
var FORGET_PASSWORD_EVENT_UPDATED = "forgetPassword.updated"
var FORGET_PASSWORD_EVENT_DELETED = "forgetPassword.deleted"
var FORGET_PASSWORD_EVENTS = []string{
	FORGET_PASSWORD_EVENT_CREATED,
	FORGET_PASSWORD_EVENT_UPDATED,
	FORGET_PASSWORD_EVENT_DELETED,
}
type ForgetPasswordFieldMap struct {
		User TranslatedString `yaml:"user"`
		Passport TranslatedString `yaml:"passport"`
		Status TranslatedString `yaml:"status"`
		ValidUntil TranslatedString `yaml:"validUntil"`
		BlockedUntil TranslatedString `yaml:"blockedUntil"`
		SecondsToUnblock TranslatedString `yaml:"secondsToUnblock"`
		Otp TranslatedString `yaml:"otp"`
		RecoveryAbsoluteUrl TranslatedString `yaml:"recoveryAbsoluteUrl"`
}
var ForgetPasswordEntityMetaConfig map[string]int64 = map[string]int64{
}
var ForgetPasswordEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&ForgetPasswordEntity{}))
func entityForgetPasswordFormatter(dto *ForgetPasswordEntity, query QueryDSL) {
	if dto == nil {
		return
	}
			dto.ValidUntilFormatted = FormatDateBasedOnQuery(dto.ValidUntil, query)
			dto.BlockedUntilFormatted = FormatDateBasedOnQuery(dto.BlockedUntil, query)
	if dto.Created > 0 {
		dto.CreatedFormatted = FormatDateBasedOnQuery(dto.Created, query)
	}
	if dto.Updated > 0 {
		dto.CreatedFormatted = FormatDateBasedOnQuery(dto.Updated, query)
	}
}
func ForgetPasswordMockEntity() *ForgetPasswordEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &ForgetPasswordEntity{
      Status : &stringHolder,
      SecondsToUnblock : &int64Holder,
      Otp : &stringHolder,
      RecoveryAbsoluteUrl : &stringHolder,
	}
	return entity
}
func ForgetPasswordActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := ForgetPasswordMockEntity()
		_, err := ForgetPasswordActionCreate(entity, query)
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
  func ForgetPasswordActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*ForgetPasswordEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &ForgetPasswordEntity{
          Status: &tildaRef,
          Otp: &tildaRef,
          RecoveryAbsoluteUrl: &tildaRef,
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
  func ForgetPasswordAssociationCreate(dto *ForgetPasswordEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func ForgetPasswordRelationContentCreate(dto *ForgetPasswordEntity, query QueryDSL) error {
return nil
}
func ForgetPasswordRelationContentUpdate(dto *ForgetPasswordEntity, query QueryDSL) error {
	return nil
}
func ForgetPasswordPolyglotCreateHandler(dto *ForgetPasswordEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func ForgetPasswordValidator(dto *ForgetPasswordEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func ForgetPasswordEntityPreSanitize(dto *ForgetPasswordEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func ForgetPasswordEntityBeforeCreateAppend(dto *ForgetPasswordEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    ForgetPasswordRecursiveAddUniqueId(dto, query)
  }
  func ForgetPasswordRecursiveAddUniqueId(dto *ForgetPasswordEntity, query QueryDSL) {
  }
func ForgetPasswordActionBatchCreateFn(dtos []*ForgetPasswordEntity, query QueryDSL) ([]*ForgetPasswordEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*ForgetPasswordEntity{}
		for _, item := range dtos {
			s, err := ForgetPasswordActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func ForgetPasswordDeleteEntireChildren(query QueryDSL, dto *ForgetPasswordEntity) (*IError) {
  return nil
}
func ForgetPasswordActionCreateFn(dto *ForgetPasswordEntity, query QueryDSL) (*ForgetPasswordEntity, *IError) {
	// 1. Validate always
	if iError := ForgetPasswordValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	ForgetPasswordEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	ForgetPasswordEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	ForgetPasswordPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	ForgetPasswordRelationContentCreate(dto, query)
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
	ForgetPasswordAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(FORGET_PASSWORD_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&ForgetPasswordEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func ForgetPasswordActionGetOne(query QueryDSL) (*ForgetPasswordEntity, *IError) {
    refl := reflect.ValueOf(&ForgetPasswordEntity{})
    item, err := GetOneEntity[ForgetPasswordEntity](query, refl)
    entityForgetPasswordFormatter(item, query)
    return item, err
  }
  func ForgetPasswordActionQuery(query QueryDSL) ([]*ForgetPasswordEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&ForgetPasswordEntity{})
    items, meta, err := QueryEntitiesPointer[ForgetPasswordEntity](query, refl)
    for _, item := range items {
      entityForgetPasswordFormatter(item, query)
    }
    return items, meta, err
  }
  func ForgetPasswordUpdateExec(dbref *gorm.DB, query QueryDSL, fields *ForgetPasswordEntity) (*ForgetPasswordEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = FORGET_PASSWORD_EVENT_UPDATED
    ForgetPasswordEntityPreSanitize(fields, query)
    var item ForgetPasswordEntity
    q := dbref.
      Where(&ForgetPasswordEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    ForgetPasswordRelationContentUpdate(fields, query)
    ForgetPasswordPolyglotCreateHandler(fields, query)
    if ero := ForgetPasswordDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&ForgetPasswordEntity{UniqueId: uniqueId}).
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
  func ForgetPasswordActionUpdateFn(query QueryDSL, fields *ForgetPasswordEntity) (*ForgetPasswordEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := ForgetPasswordValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // ForgetPasswordRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      var item *ForgetPasswordEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *IError
        item, err = ForgetPasswordUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, CastToIError(vf)
    } else {
      dbref = query.Tx
      return ForgetPasswordUpdateExec(dbref, query, fields)
    }
  }
var ForgetPasswordWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire forgetpasswords ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
      ActionRequires: []PermissionInfo{PERM_ROOT_FORGET_PASSWORD_DELETE},
    })
		count, _ := ForgetPasswordActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func ForgetPasswordActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&ForgetPasswordEntity{})
	query.ActionRequires = []PermissionInfo{PERM_ROOT_FORGET_PASSWORD_DELETE}
	return RemoveEntity[ForgetPasswordEntity](query, refl)
}
func ForgetPasswordActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[ForgetPasswordEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'ForgetPasswordEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func ForgetPasswordActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[ForgetPasswordEntity]) (
    *BulkRecordRequest[ForgetPasswordEntity], *IError,
  ) {
    result := []*ForgetPasswordEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := ForgetPasswordActionUpdate(query, record)
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
func (x *ForgetPasswordEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var ForgetPasswordEntityMeta = TableMetaData{
	EntityName:    "ForgetPassword",
	ExportKey:    "forget-passwords",
	TableNameInDb: "fb_forget-password_entities",
	EntityObject:  &ForgetPasswordEntity{},
	ExportStream: ForgetPasswordActionExportT,
	ImportQuery: ForgetPasswordActionImport,
}
func ForgetPasswordActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[ForgetPasswordEntity](query, ForgetPasswordActionQuery, ForgetPasswordPreloadRelations)
}
func ForgetPasswordActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[ForgetPasswordEntity](query, ForgetPasswordActionQuery, ForgetPasswordPreloadRelations)
}
func ForgetPasswordActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content ForgetPasswordEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := ForgetPasswordActionCreate(&content, query)
	return err
}
var ForgetPasswordCommonCliFlags = []cli.Flag{
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
      Name:     "passport-id",
      Required: false,
      Usage:    "passport",
    },
    &cli.StringFlag{
      Name:     "status",
      Required: false,
      Usage:    "status",
    },
    &cli.Int64Flag{
      Name:     "seconds-to-unblock",
      Required: false,
      Usage:    "secondsToUnblock",
    },
    &cli.StringFlag{
      Name:     "otp",
      Required: false,
      Usage:    "otp",
    },
    &cli.StringFlag{
      Name:     "recovery-absolute-url",
      Required: false,
      Usage:    "recoveryAbsoluteUrl",
    },
}
var ForgetPasswordCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:     "status",
		StructField:     "Status",
		Required: false,
		Usage:    "status",
		Type: "string",
	},
	{
		Name:     "secondsToUnblock",
		StructField:     "SecondsToUnblock",
		Required: false,
		Usage:    "secondsToUnblock",
		Type: "int64",
	},
	{
		Name:     "otp",
		StructField:     "Otp",
		Required: false,
		Usage:    "otp",
		Type: "string",
	},
	{
		Name:     "recoveryAbsoluteUrl",
		StructField:     "RecoveryAbsoluteUrl",
		Required: false,
		Usage:    "recoveryAbsoluteUrl",
		Type: "string",
	},
}
var ForgetPasswordCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "passport-id",
      Required: false,
      Usage:    "passport",
    },
    &cli.StringFlag{
      Name:     "status",
      Required: false,
      Usage:    "status",
    },
    &cli.Int64Flag{
      Name:     "seconds-to-unblock",
      Required: false,
      Usage:    "secondsToUnblock",
    },
    &cli.StringFlag{
      Name:     "otp",
      Required: false,
      Usage:    "otp",
    },
    &cli.StringFlag{
      Name:     "recovery-absolute-url",
      Required: false,
      Usage:    "recoveryAbsoluteUrl",
    },
}
  var ForgetPasswordCreateCmd cli.Command = FORGET_PASSWORD_ACTION_POST_ONE.ToCli()
  var ForgetPasswordCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_FORGET_PASSWORD_CREATE},
      })
      entity := &ForgetPasswordEntity{}
      for _, item := range ForgetPasswordCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := ForgetPasswordActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var ForgetPasswordUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: ForgetPasswordCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_FORGET_PASSWORD_UPDATE},
      })
      entity := CastForgetPasswordFromCli(c)
      if entity, err := ForgetPasswordActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* ForgetPasswordEntity) FromCli(c *cli.Context) *ForgetPasswordEntity {
	return CastForgetPasswordFromCli(c)
}
func CastForgetPasswordFromCli (c *cli.Context) *ForgetPasswordEntity {
	template := &ForgetPasswordEntity{}
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
      if c.IsSet("passport-id") {
        value := c.String("passport-id")
        template.PassportId = &value
      }
      if c.IsSet("status") {
        value := c.String("status")
        template.Status = &value
      }
      if c.IsSet("otp") {
        value := c.String("otp")
        template.Otp = &value
      }
      if c.IsSet("recovery-absolute-url") {
        value := c.String("recovery-absolute-url")
        template.RecoveryAbsoluteUrl = &value
      }
	return template
}
  func ForgetPasswordSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      ForgetPasswordActionCreate,
      reflect.ValueOf(&ForgetPasswordEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func ForgetPasswordWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := ForgetPasswordActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "ForgetPassword", result)
    }
  }
var ForgetPasswordImportExportCommands = []cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_FORGET_PASSWORD_CREATE},
      })
			ForgetPasswordActionSeeder(query, c.Int("count"))
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
				Value: "forget-password-seeder.yml",
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
        ActionRequires: []PermissionInfo{PERM_ROOT_FORGET_PASSWORD_CREATE},
      })
			ForgetPasswordActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "forget-password-seeder-forget-password.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of forget-passwords, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]ForgetPasswordEntity{}
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
			ForgetPasswordCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				ForgetPasswordActionCreate,
				reflect.ValueOf(&ForgetPasswordEntity{}).Elem(),
				c.String("file"),
        &SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_FORGET_PASSWORD_CREATE},
				},
        func() ForgetPasswordEntity {
					v := CastForgetPasswordFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var ForgetPasswordCliCommands []cli.Command = []cli.Command{
      FORGET_PASSWORD_ACTION_QUERY.ToCli(),
      FORGET_PASSWORD_ACTION_TABLE.ToCli(),
      GetCommonTableQuery(reflect.ValueOf(&ForgetPasswordEntity{}).Elem(), ForgetPasswordActionQuery),
      ForgetPasswordCreateCmd,
      ForgetPasswordUpdateCmd,
      ForgetPasswordCreateInteractiveCmd,
      ForgetPasswordWipeCmd,
      GetCommonRemoveQuery(reflect.ValueOf(&ForgetPasswordEntity{}).Elem(), ForgetPasswordActionRemove),
  }
  func ForgetPasswordCliFn() cli.Command {
    ForgetPasswordCliCommands = append(ForgetPasswordCliCommands, ForgetPasswordImportExportCommands...)
    return cli.Command{
      Name:        "forgetPassword",
      Description: "ForgetPasswords module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: ForgetPasswordCliCommands,
    }
  }
var FORGET_PASSWORD_ACTION_TABLE = Module2Action{
  Name:    "table",
  ActionAliases: []string{"t"},
  Flags:  CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: ForgetPasswordActionQuery,
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    CommonCliTableCmd2(c,
      ForgetPasswordActionQuery,
      security,
      reflect.ValueOf(&ForgetPasswordEntity{}).Elem(),
    )
    return nil
  },
}
var FORGET_PASSWORD_ACTION_QUERY = Module2Action{
  Method: "GET",
  Url:    "/forget-passwords",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_FORGET_PASSWORD_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpQueryEntity(c, ForgetPasswordActionQuery)
    },
  },
  Format: "QUERY",
  Action: ForgetPasswordActionQuery,
  ResponseEntity: &[]ForgetPasswordEntity{},
  CliAction: func(c *cli.Context, security *SecurityModel) error {
		CommonCliQueryCmd2(
			c,
			ForgetPasswordActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionAliases: []string{"q"},
	Flags:         CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var FORGET_PASSWORD_ACTION_EXPORT = Module2Action{
  Method: "GET",
  Url:    "/forget-passwords/export",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_FORGET_PASSWORD_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpStreamFileChannel(c, ForgetPasswordActionExport)
    },
  },
  Format: "QUERY",
  Action: ForgetPasswordActionExport,
  ResponseEntity: &[]ForgetPasswordEntity{},
}
var FORGET_PASSWORD_ACTION_GET_ONE = Module2Action{
  Method: "GET",
  Url:    "/forget-password/:uniqueId",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_FORGET_PASSWORD_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpGetEntity(c, ForgetPasswordActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: ForgetPasswordActionGetOne,
  ResponseEntity: &ForgetPasswordEntity{},
}
var FORGET_PASSWORD_ACTION_POST_ONE = Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new forgetPassword",
  Flags: ForgetPasswordCommonCliFlags,
  Method: "POST",
  Url:    "/forget-password",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_FORGET_PASSWORD_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpPostEntity(c, ForgetPasswordActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    result, err := CliPostEntity(c, ForgetPasswordActionCreate, security)
    HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: ForgetPasswordActionCreate,
  Format: "POST_ONE",
  RequestEntity: &ForgetPasswordEntity{},
  ResponseEntity: &ForgetPasswordEntity{},
}
var FORGET_PASSWORD_ACTION_PATCH = Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: ForgetPasswordCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/forget-password",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_FORGET_PASSWORD_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntity(c, ForgetPasswordActionUpdate)
    },
  },
  Action: ForgetPasswordActionUpdate,
  RequestEntity: &ForgetPasswordEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &ForgetPasswordEntity{},
}
var FORGET_PASSWORD_ACTION_PATCH_BULK = Module2Action{
  Method: "PATCH",
  Url:    "/forget-passwords",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_FORGET_PASSWORD_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntities(c, ForgetPasswordActionBulkUpdate)
    },
  },
  Action: ForgetPasswordActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &BulkRecordRequest[ForgetPasswordEntity]{},
  ResponseEntity: &BulkRecordRequest[ForgetPasswordEntity]{},
}
var FORGET_PASSWORD_ACTION_DELETE = Module2Action{
  Method: "DELETE",
  Url:    "/forget-password",
  Format: "DELETE_DSL",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_FORGET_PASSWORD_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpRemoveEntity(c, ForgetPasswordActionRemove)
    },
  },
  Action: ForgetPasswordActionRemove,
  RequestEntity: &DeleteRequest{},
  ResponseEntity: &DeleteResponse{},
  TargetEntity: &ForgetPasswordEntity{},
}
  /**
  *	Override this function on ForgetPasswordEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendForgetPasswordRouter = func(r *[]Module2Action) {}
  func GetForgetPasswordModule2Actions() []Module2Action {
    routes := []Module2Action{
      FORGET_PASSWORD_ACTION_QUERY,
      FORGET_PASSWORD_ACTION_EXPORT,
      FORGET_PASSWORD_ACTION_GET_ONE,
      FORGET_PASSWORD_ACTION_POST_ONE,
      FORGET_PASSWORD_ACTION_PATCH,
      FORGET_PASSWORD_ACTION_PATCH_BULK,
      FORGET_PASSWORD_ACTION_DELETE,
    }
    // Append user defined functions
    AppendForgetPasswordRouter(&routes)
    return routes
  }
  func CreateForgetPasswordRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetForgetPasswordModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, ForgetPasswordEntityJsonSchema, "forget-password-http", "workspaces")
    WriteEntitySchema("ForgetPasswordEntity", ForgetPasswordEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_FORGET_PASSWORD_DELETE = PermissionInfo{
  CompleteKey: "root/workspaces/forget-password/delete",
}
var PERM_ROOT_FORGET_PASSWORD_CREATE = PermissionInfo{
  CompleteKey: "root/workspaces/forget-password/create",
}
var PERM_ROOT_FORGET_PASSWORD_UPDATE = PermissionInfo{
  CompleteKey: "root/workspaces/forget-password/update",
}
var PERM_ROOT_FORGET_PASSWORD_QUERY = PermissionInfo{
  CompleteKey: "root/workspaces/forget-password/query",
}
var PERM_ROOT_FORGET_PASSWORD = PermissionInfo{
  CompleteKey: "root/workspaces/forget-password/*",
}
var ALL_FORGET_PASSWORD_PERMISSIONS = []PermissionInfo{
	PERM_ROOT_FORGET_PASSWORD_DELETE,
	PERM_ROOT_FORGET_PASSWORD_CREATE,
	PERM_ROOT_FORGET_PASSWORD_UPDATE,
	PERM_ROOT_FORGET_PASSWORD_QUERY,
	PERM_ROOT_FORGET_PASSWORD,
}