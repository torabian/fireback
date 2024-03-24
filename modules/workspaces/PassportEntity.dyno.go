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
type PassportEntity struct {
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
    User   *  UserEntity `json:"user" yaml:"user"    gorm:"foreignKey:UserId;references:UniqueId"     `
    // Datenano also has a text representation
    Value   *string `json:"value" yaml:"value"  validate:"required"    gorm:"unique"     `
    // Datenano also has a text representation
    Password   *string `json:"-" yaml:"-"       `
    // Datenano also has a text representation
    Confirmed   *bool `json:"confirmed" yaml:"confirmed"       `
    // Datenano also has a text representation
    AccessToken   *string `json:"accessToken" yaml:"accessToken"       `
    // Datenano also has a text representation
    Children []*PassportEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *PassportEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var PassportPreloadRelations []string = []string{}
var PASSPORT_EVENT_CREATED = "passport.created"
var PASSPORT_EVENT_UPDATED = "passport.updated"
var PASSPORT_EVENT_DELETED = "passport.deleted"
var PASSPORT_EVENTS = []string{
	PASSPORT_EVENT_CREATED,
	PASSPORT_EVENT_UPDATED,
	PASSPORT_EVENT_DELETED,
}
type PassportFieldMap struct {
		Type TranslatedString `yaml:"type"`
		User TranslatedString `yaml:"user"`
		Value TranslatedString `yaml:"value"`
		Password TranslatedString `yaml:"password"`
		Confirmed TranslatedString `yaml:"confirmed"`
		AccessToken TranslatedString `yaml:"accessToken"`
}
var PassportEntityMetaConfig map[string]int64 = map[string]int64{
}
var PassportEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&PassportEntity{}))
func entityPassportFormatter(dto *PassportEntity, query QueryDSL) {
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
func PassportMockEntity() *PassportEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &PassportEntity{
      Type : &stringHolder,
      Value : &stringHolder,
      Password : &stringHolder,
      AccessToken : &stringHolder,
	}
	return entity
}
func PassportActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := PassportMockEntity()
		_, err := PassportActionCreate(entity, query)
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
  func PassportActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*PassportEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &PassportEntity{
          Type: &tildaRef,
          Value: &tildaRef,
          Password: &tildaRef,
          AccessToken: &tildaRef,
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
  func PassportAssociationCreate(dto *PassportEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func PassportRelationContentCreate(dto *PassportEntity, query QueryDSL) error {
return nil
}
func PassportRelationContentUpdate(dto *PassportEntity, query QueryDSL) error {
	return nil
}
func PassportPolyglotCreateHandler(dto *PassportEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func PassportValidator(dto *PassportEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func PassportEntityPreSanitize(dto *PassportEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func PassportEntityBeforeCreateAppend(dto *PassportEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    PassportRecursiveAddUniqueId(dto, query)
  }
  func PassportRecursiveAddUniqueId(dto *PassportEntity, query QueryDSL) {
  }
func PassportActionBatchCreateFn(dtos []*PassportEntity, query QueryDSL) ([]*PassportEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*PassportEntity{}
		for _, item := range dtos {
			s, err := PassportActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func PassportDeleteEntireChildren(query QueryDSL, dto *PassportEntity) (*IError) {
  return nil
}
func PassportActionCreateFn(dto *PassportEntity, query QueryDSL) (*PassportEntity, *IError) {
	// 1. Validate always
	if iError := PassportValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	PassportEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	PassportEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	PassportPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	PassportRelationContentCreate(dto, query)
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
	PassportAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(PASSPORT_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&PassportEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func PassportActionGetOne(query QueryDSL) (*PassportEntity, *IError) {
    refl := reflect.ValueOf(&PassportEntity{})
    item, err := GetOneEntity[PassportEntity](query, refl)
    entityPassportFormatter(item, query)
    return item, err
  }
  func PassportActionQuery(query QueryDSL) ([]*PassportEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&PassportEntity{})
    items, meta, err := QueryEntitiesPointer[PassportEntity](query, refl)
    for _, item := range items {
      entityPassportFormatter(item, query)
    }
    return items, meta, err
  }
  func PassportUpdateExec(dbref *gorm.DB, query QueryDSL, fields *PassportEntity) (*PassportEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = PASSPORT_EVENT_UPDATED
    PassportEntityPreSanitize(fields, query)
    var item PassportEntity
    q := dbref.
      Where(&PassportEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    PassportRelationContentUpdate(fields, query)
    PassportPolyglotCreateHandler(fields, query)
    if ero := PassportDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&PassportEntity{UniqueId: uniqueId}).
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
  func PassportActionUpdateFn(query QueryDSL, fields *PassportEntity) (*PassportEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := PassportValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // PassportRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      var item *PassportEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *IError
        item, err = PassportUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, CastToIError(vf)
    } else {
      dbref = query.Tx
      return PassportUpdateExec(dbref, query, fields)
    }
  }
var PassportWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire passports ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
      ActionRequires: []PermissionInfo{PERM_ROOT_PASSPORT_DELETE},
    })
		count, _ := PassportActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func PassportActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&PassportEntity{})
	query.ActionRequires = []PermissionInfo{PERM_ROOT_PASSPORT_DELETE}
	return RemoveEntity[PassportEntity](query, refl)
}
func PassportActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[PassportEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'PassportEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func PassportActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[PassportEntity]) (
    *BulkRecordRequest[PassportEntity], *IError,
  ) {
    result := []*PassportEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := PassportActionUpdate(query, record)
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
func (x *PassportEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var PassportEntityMeta = TableMetaData{
	EntityName:    "Passport",
	ExportKey:    "passports",
	TableNameInDb: "fb_passport_entities",
	EntityObject:  &PassportEntity{},
	ExportStream: PassportActionExportT,
	ImportQuery: PassportActionImport,
}
func PassportActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[PassportEntity](query, PassportActionQuery, PassportPreloadRelations)
}
func PassportActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[PassportEntity](query, PassportActionQuery, PassportPreloadRelations)
}
func PassportActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content PassportEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := PassportActionCreate(&content, query)
	return err
}
var PassportCommonCliFlags = []cli.Flag{
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
      Usage:    "type",
    },
    &cli.StringFlag{
      Name:     "user-id",
      Required: false,
      Usage:    "user",
    },
    &cli.StringFlag{
      Name:     "value",
      Required: true,
      Usage:    "value",
    },
    &cli.StringFlag{
      Name:     "password",
      Required: false,
      Usage:    "password",
    },
    &cli.BoolFlag{
      Name:     "confirmed",
      Required: false,
      Usage:    "confirmed",
    },
    &cli.StringFlag{
      Name:     "access-token",
      Required: false,
      Usage:    "accessToken",
    },
}
var PassportCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:     "type",
		StructField:     "Type",
		Required: true,
		Usage:    "type",
		Type: "string",
	},
	{
		Name:     "value",
		StructField:     "Value",
		Required: true,
		Usage:    "value",
		Type: "string",
	},
	{
		Name:     "password",
		StructField:     "Password",
		Required: false,
		Usage:    "password",
		Type: "string",
	},
	{
		Name:     "confirmed",
		StructField:     "Confirmed",
		Required: false,
		Usage:    "confirmed",
		Type: "bool",
	},
	{
		Name:     "accessToken",
		StructField:     "AccessToken",
		Required: false,
		Usage:    "accessToken",
		Type: "string",
	},
}
var PassportCommonCliFlagsOptional = []cli.Flag{
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
      Usage:    "type",
    },
    &cli.StringFlag{
      Name:     "user-id",
      Required: false,
      Usage:    "user",
    },
    &cli.StringFlag{
      Name:     "value",
      Required: true,
      Usage:    "value",
    },
    &cli.StringFlag{
      Name:     "password",
      Required: false,
      Usage:    "password",
    },
    &cli.BoolFlag{
      Name:     "confirmed",
      Required: false,
      Usage:    "confirmed",
    },
    &cli.StringFlag{
      Name:     "access-token",
      Required: false,
      Usage:    "accessToken",
    },
}
  var PassportCreateCmd cli.Command = PASSPORT_ACTION_POST_ONE.ToCli()
  var PassportCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_PASSPORT_CREATE},
      })
      entity := &PassportEntity{}
      for _, item := range PassportCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := PassportActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var PassportUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: PassportCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_PASSPORT_UPDATE},
      })
      entity := CastPassportFromCli(c)
      if entity, err := PassportActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* PassportEntity) FromCli(c *cli.Context) *PassportEntity {
	return CastPassportFromCli(c)
}
func CastPassportFromCli (c *cli.Context) *PassportEntity {
	template := &PassportEntity{}
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
      if c.IsSet("user-id") {
        value := c.String("user-id")
        template.UserId = &value
      }
      if c.IsSet("value") {
        value := c.String("value")
        template.Value = &value
      }
      if c.IsSet("password") {
        value := c.String("password")
        template.Password = &value
      }
      if c.IsSet("access-token") {
        value := c.String("access-token")
        template.AccessToken = &value
      }
	return template
}
  func PassportSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      PassportActionCreate,
      reflect.ValueOf(&PassportEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func PassportWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := PassportActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "Passport", result)
    }
  }
var PassportImportExportCommands = []cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_PASSPORT_CREATE},
      })
			PassportActionSeeder(query, c.Int("count"))
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
				Value: "passport-seeder.yml",
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
        ActionRequires: []PermissionInfo{PERM_ROOT_PASSPORT_CREATE},
      })
			PassportActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "passport-seeder-passport.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of passports, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]PassportEntity{}
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
			PassportCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				PassportActionCreate,
				reflect.ValueOf(&PassportEntity{}).Elem(),
				c.String("file"),
        &SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_PASSPORT_CREATE},
				},
        func() PassportEntity {
					v := CastPassportFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var PassportCliCommands []cli.Command = []cli.Command{
      GetCommonQuery2(PassportActionQuery, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_PASSPORT_QUERY},
      }),
      GetCommonTableQuery(reflect.ValueOf(&PassportEntity{}).Elem(), PassportActionQuery),
          PassportCreateCmd,
          PassportUpdateCmd,
          PassportCreateInteractiveCmd,
          PassportWipeCmd,
          GetCommonRemoveQuery(reflect.ValueOf(&PassportEntity{}).Elem(), PassportActionRemove),
  }
  func PassportCliFn() cli.Command {
    PassportCliCommands = append(PassportCliCommands, PassportImportExportCommands...)
    return cli.Command{
      Name:        "passport",
      Description: "Passports module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: PassportCliCommands,
    }
  }
var PASSPORT_ACTION_QUERY = Module2Action{
  Method: "GET",
  Url:    "/passports",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PASSPORT_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpQueryEntity(c, PassportActionQuery)
    },
  },
  Format: "QUERY",
  Action: PassportActionQuery,
  ResponseEntity: &[]PassportEntity{},
}
var PASSPORT_ACTION_EXPORT = Module2Action{
  Method: "GET",
  Url:    "/passports/export",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PASSPORT_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpStreamFileChannel(c, PassportActionExport)
    },
  },
  Format: "QUERY",
  Action: PassportActionExport,
  ResponseEntity: &[]PassportEntity{},
}
var PASSPORT_ACTION_GET_ONE = Module2Action{
  Method: "GET",
  Url:    "/passport/:uniqueId",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PASSPORT_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpGetEntity(c, PassportActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: PassportActionGetOne,
  ResponseEntity: &PassportEntity{},
}
var PASSPORT_ACTION_POST_ONE = Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new passport",
  Flags: PassportCommonCliFlags,
  Method: "POST",
  Url:    "/passport",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PASSPORT_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpPostEntity(c, PassportActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    result, err := CliPostEntity(c, PassportActionCreate, security)
    HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: PassportActionCreate,
  Format: "POST_ONE",
  RequestEntity: &PassportEntity{},
  ResponseEntity: &PassportEntity{},
}
var PASSPORT_ACTION_PATCH = Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: PassportCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/passport",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PASSPORT_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntity(c, PassportActionUpdate)
    },
  },
  Action: PassportActionUpdate,
  RequestEntity: &PassportEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &PassportEntity{},
}
var PASSPORT_ACTION_PATCH_BULK = Module2Action{
  Method: "PATCH",
  Url:    "/passports",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PASSPORT_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntities(c, PassportActionBulkUpdate)
    },
  },
  Action: PassportActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &BulkRecordRequest[PassportEntity]{},
  ResponseEntity: &BulkRecordRequest[PassportEntity]{},
}
var PASSPORT_ACTION_DELETE = Module2Action{
  Method: "DELETE",
  Url:    "/passport",
  Format: "DELETE_DSL",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PASSPORT_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpRemoveEntity(c, PassportActionRemove)
    },
  },
  Action: PassportActionRemove,
  RequestEntity: &DeleteRequest{},
  ResponseEntity: &DeleteResponse{},
  TargetEntity: &PassportEntity{},
}
  /**
  *	Override this function on PassportEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendPassportRouter = func(r *[]Module2Action) {}
  func GetPassportModule2Actions() []Module2Action {
    routes := []Module2Action{
      PASSPORT_ACTION_QUERY,
      PASSPORT_ACTION_EXPORT,
      PASSPORT_ACTION_GET_ONE,
      PASSPORT_ACTION_POST_ONE,
      PASSPORT_ACTION_PATCH,
      PASSPORT_ACTION_PATCH_BULK,
      PASSPORT_ACTION_DELETE,
    }
    // Append user defined functions
    AppendPassportRouter(&routes)
    return routes
  }
  func CreatePassportRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetPassportModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, PassportEntityJsonSchema, "passport-http", "workspaces")
    WriteEntitySchema("PassportEntity", PassportEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_PASSPORT_DELETE = PermissionInfo{
  CompleteKey: "root/workspaces/passport/delete",
}
var PERM_ROOT_PASSPORT_CREATE = PermissionInfo{
  CompleteKey: "root/workspaces/passport/create",
}
var PERM_ROOT_PASSPORT_UPDATE = PermissionInfo{
  CompleteKey: "root/workspaces/passport/update",
}
var PERM_ROOT_PASSPORT_QUERY = PermissionInfo{
  CompleteKey: "root/workspaces/passport/query",
}
var PERM_ROOT_PASSPORT = PermissionInfo{
  CompleteKey: "root/workspaces/passport/*",
}
var ALL_PASSPORT_PERMISSIONS = []PermissionInfo{
	PERM_ROOT_PASSPORT_DELETE,
	PERM_ROOT_PASSPORT_CREATE,
	PERM_ROOT_PASSPORT_UPDATE,
	PERM_ROOT_PASSPORT_QUERY,
	PERM_ROOT_PASSPORT,
}