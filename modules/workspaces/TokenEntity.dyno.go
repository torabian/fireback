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
type TokenEntity struct {
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
    ValidUntil   *string `json:"validUntil" yaml:"validUntil"       `
    // Datenano also has a text representation
    Children []*TokenEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *TokenEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var TokenPreloadRelations []string = []string{}
var TOKEN_EVENT_CREATED = "token.created"
var TOKEN_EVENT_UPDATED = "token.updated"
var TOKEN_EVENT_DELETED = "token.deleted"
var TOKEN_EVENTS = []string{
	TOKEN_EVENT_CREATED,
	TOKEN_EVENT_UPDATED,
	TOKEN_EVENT_DELETED,
}
type TokenFieldMap struct {
		User TranslatedString `yaml:"user"`
		ValidUntil TranslatedString `yaml:"validUntil"`
}
var TokenEntityMetaConfig map[string]int64 = map[string]int64{
}
var TokenEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&TokenEntity{}))
func entityTokenFormatter(dto *TokenEntity, query QueryDSL) {
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
func TokenMockEntity() *TokenEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &TokenEntity{
      ValidUntil : &stringHolder,
	}
	return entity
}
func TokenActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := TokenMockEntity()
		_, err := TokenActionCreate(entity, query)
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
  func TokenActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*TokenEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &TokenEntity{
          ValidUntil: &tildaRef,
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
  func TokenAssociationCreate(dto *TokenEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func TokenRelationContentCreate(dto *TokenEntity, query QueryDSL) error {
return nil
}
func TokenRelationContentUpdate(dto *TokenEntity, query QueryDSL) error {
	return nil
}
func TokenPolyglotCreateHandler(dto *TokenEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func TokenValidator(dto *TokenEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func TokenEntityPreSanitize(dto *TokenEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func TokenEntityBeforeCreateAppend(dto *TokenEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    TokenRecursiveAddUniqueId(dto, query)
  }
  func TokenRecursiveAddUniqueId(dto *TokenEntity, query QueryDSL) {
  }
func TokenActionBatchCreateFn(dtos []*TokenEntity, query QueryDSL) ([]*TokenEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*TokenEntity{}
		for _, item := range dtos {
			s, err := TokenActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func TokenDeleteEntireChildren(query QueryDSL, dto *TokenEntity) (*IError) {
  return nil
}
func TokenActionCreateFn(dto *TokenEntity, query QueryDSL) (*TokenEntity, *IError) {
	// 1. Validate always
	if iError := TokenValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	TokenEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	TokenEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	TokenPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	TokenRelationContentCreate(dto, query)
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
	TokenAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(TOKEN_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&TokenEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func TokenActionGetOne(query QueryDSL) (*TokenEntity, *IError) {
    refl := reflect.ValueOf(&TokenEntity{})
    item, err := GetOneEntity[TokenEntity](query, refl)
    entityTokenFormatter(item, query)
    return item, err
  }
  func TokenActionQuery(query QueryDSL) ([]*TokenEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&TokenEntity{})
    items, meta, err := QueryEntitiesPointer[TokenEntity](query, refl)
    for _, item := range items {
      entityTokenFormatter(item, query)
    }
    return items, meta, err
  }
  func TokenUpdateExec(dbref *gorm.DB, query QueryDSL, fields *TokenEntity) (*TokenEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = TOKEN_EVENT_UPDATED
    TokenEntityPreSanitize(fields, query)
    var item TokenEntity
    q := dbref.
      Where(&TokenEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    TokenRelationContentUpdate(fields, query)
    TokenPolyglotCreateHandler(fields, query)
    if ero := TokenDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&TokenEntity{UniqueId: uniqueId}).
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
  func TokenActionUpdateFn(query QueryDSL, fields *TokenEntity) (*TokenEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := TokenValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // TokenRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      var item *TokenEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *IError
        item, err = TokenUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, CastToIError(vf)
    } else {
      dbref = query.Tx
      return TokenUpdateExec(dbref, query, fields)
    }
  }
var TokenWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire tokens ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
      ActionRequires: []PermissionInfo{PERM_ROOT_TOKEN_DELETE},
    })
		count, _ := TokenActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func TokenActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&TokenEntity{})
	query.ActionRequires = []PermissionInfo{PERM_ROOT_TOKEN_DELETE}
	return RemoveEntity[TokenEntity](query, refl)
}
func TokenActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[TokenEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'TokenEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func TokenActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[TokenEntity]) (
    *BulkRecordRequest[TokenEntity], *IError,
  ) {
    result := []*TokenEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := TokenActionUpdate(query, record)
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
func (x *TokenEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var TokenEntityMeta = TableMetaData{
	EntityName:    "Token",
	ExportKey:    "tokens",
	TableNameInDb: "fb_token_entities",
	EntityObject:  &TokenEntity{},
	ExportStream: TokenActionExportT,
	ImportQuery: TokenActionImport,
}
func TokenActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[TokenEntity](query, TokenActionQuery, TokenPreloadRelations)
}
func TokenActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[TokenEntity](query, TokenActionQuery, TokenPreloadRelations)
}
func TokenActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content TokenEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := TokenActionCreate(&content, query)
	return err
}
var TokenCommonCliFlags = []cli.Flag{
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
      Name:     "valid-until",
      Required: false,
      Usage:    "validUntil",
    },
}
var TokenCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:     "validUntil",
		StructField:     "ValidUntil",
		Required: false,
		Usage:    "validUntil",
		Type: "string",
	},
}
var TokenCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "valid-until",
      Required: false,
      Usage:    "validUntil",
    },
}
  var TokenCreateCmd cli.Command = TOKEN_ACTION_POST_ONE.ToCli()
  var TokenCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_TOKEN_CREATE},
      })
      entity := &TokenEntity{}
      for _, item := range TokenCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := TokenActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var TokenUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: TokenCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_TOKEN_UPDATE},
      })
      entity := CastTokenFromCli(c)
      if entity, err := TokenActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* TokenEntity) FromCli(c *cli.Context) *TokenEntity {
	return CastTokenFromCli(c)
}
func CastTokenFromCli (c *cli.Context) *TokenEntity {
	template := &TokenEntity{}
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
      if c.IsSet("valid-until") {
        value := c.String("valid-until")
        template.ValidUntil = &value
      }
	return template
}
  func TokenSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      TokenActionCreate,
      reflect.ValueOf(&TokenEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func TokenWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := TokenActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "Token", result)
    }
  }
var TokenImportExportCommands = []cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_TOKEN_CREATE},
      })
			TokenActionSeeder(query, c.Int("count"))
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
				Value: "token-seeder.yml",
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
        ActionRequires: []PermissionInfo{PERM_ROOT_TOKEN_CREATE},
      })
			TokenActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "token-seeder-token.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of tokens, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]TokenEntity{}
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
			TokenCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				TokenActionCreate,
				reflect.ValueOf(&TokenEntity{}).Elem(),
				c.String("file"),
        &SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_TOKEN_CREATE},
				},
        func() TokenEntity {
					v := CastTokenFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var TokenCliCommands []cli.Command = []cli.Command{
      TOKEN_ACTION_QUERY.ToCli(),
      TOKEN_ACTION_TABLE.ToCli(),
      GetCommonTableQuery(reflect.ValueOf(&TokenEntity{}).Elem(), TokenActionQuery),
      TokenCreateCmd,
      TokenUpdateCmd,
      TokenCreateInteractiveCmd,
      TokenWipeCmd,
      GetCommonRemoveQuery(reflect.ValueOf(&TokenEntity{}).Elem(), TokenActionRemove),
  }
  func TokenCliFn() cli.Command {
    TokenCliCommands = append(TokenCliCommands, TokenImportExportCommands...)
    return cli.Command{
      Name:        "token",
      Description: "Tokens module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: TokenCliCommands,
    }
  }
var TOKEN_ACTION_TABLE = Module2Action{
  Name:    "table",
  ActionAliases: []string{"t"},
  Flags:  CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: TokenActionQuery,
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    CommonCliTableCmd2(c,
      TokenActionQuery,
      security,
      reflect.ValueOf(&TokenEntity{}).Elem(),
    )
    return nil
  },
}
var TOKEN_ACTION_QUERY = Module2Action{
  Method: "GET",
  Url:    "/tokens",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_TOKEN_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpQueryEntity(c, TokenActionQuery)
    },
  },
  Format: "QUERY",
  Action: TokenActionQuery,
  ResponseEntity: &[]TokenEntity{},
  CliAction: func(c *cli.Context, security *SecurityModel) error {
		CommonCliQueryCmd2(
			c,
			TokenActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionAliases: []string{"q"},
	Flags:         CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var TOKEN_ACTION_EXPORT = Module2Action{
  Method: "GET",
  Url:    "/tokens/export",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_TOKEN_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpStreamFileChannel(c, TokenActionExport)
    },
  },
  Format: "QUERY",
  Action: TokenActionExport,
  ResponseEntity: &[]TokenEntity{},
}
var TOKEN_ACTION_GET_ONE = Module2Action{
  Method: "GET",
  Url:    "/token/:uniqueId",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_TOKEN_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpGetEntity(c, TokenActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: TokenActionGetOne,
  ResponseEntity: &TokenEntity{},
}
var TOKEN_ACTION_POST_ONE = Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new token",
  Flags: TokenCommonCliFlags,
  Method: "POST",
  Url:    "/token",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_TOKEN_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpPostEntity(c, TokenActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    result, err := CliPostEntity(c, TokenActionCreate, security)
    HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: TokenActionCreate,
  Format: "POST_ONE",
  RequestEntity: &TokenEntity{},
  ResponseEntity: &TokenEntity{},
}
var TOKEN_ACTION_PATCH = Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: TokenCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/token",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_TOKEN_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntity(c, TokenActionUpdate)
    },
  },
  Action: TokenActionUpdate,
  RequestEntity: &TokenEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &TokenEntity{},
}
var TOKEN_ACTION_PATCH_BULK = Module2Action{
  Method: "PATCH",
  Url:    "/tokens",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_TOKEN_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntities(c, TokenActionBulkUpdate)
    },
  },
  Action: TokenActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &BulkRecordRequest[TokenEntity]{},
  ResponseEntity: &BulkRecordRequest[TokenEntity]{},
}
var TOKEN_ACTION_DELETE = Module2Action{
  Method: "DELETE",
  Url:    "/token",
  Format: "DELETE_DSL",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_TOKEN_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpRemoveEntity(c, TokenActionRemove)
    },
  },
  Action: TokenActionRemove,
  RequestEntity: &DeleteRequest{},
  ResponseEntity: &DeleteResponse{},
  TargetEntity: &TokenEntity{},
}
  /**
  *	Override this function on TokenEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendTokenRouter = func(r *[]Module2Action) {}
  func GetTokenModule2Actions() []Module2Action {
    routes := []Module2Action{
      TOKEN_ACTION_QUERY,
      TOKEN_ACTION_EXPORT,
      TOKEN_ACTION_GET_ONE,
      TOKEN_ACTION_POST_ONE,
      TOKEN_ACTION_PATCH,
      TOKEN_ACTION_PATCH_BULK,
      TOKEN_ACTION_DELETE,
    }
    // Append user defined functions
    AppendTokenRouter(&routes)
    return routes
  }
  func CreateTokenRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetTokenModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, TokenEntityJsonSchema, "token-http", "workspaces")
    WriteEntitySchema("TokenEntity", TokenEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_TOKEN_DELETE = PermissionInfo{
  CompleteKey: "root/workspaces/token/delete",
}
var PERM_ROOT_TOKEN_CREATE = PermissionInfo{
  CompleteKey: "root/workspaces/token/create",
}
var PERM_ROOT_TOKEN_UPDATE = PermissionInfo{
  CompleteKey: "root/workspaces/token/update",
}
var PERM_ROOT_TOKEN_QUERY = PermissionInfo{
  CompleteKey: "root/workspaces/token/query",
}
var PERM_ROOT_TOKEN = PermissionInfo{
  CompleteKey: "root/workspaces/token/*",
}
var ALL_TOKEN_PERMISSIONS = []PermissionInfo{
	PERM_ROOT_TOKEN_DELETE,
	PERM_ROOT_TOKEN_CREATE,
	PERM_ROOT_TOKEN_UPDATE,
	PERM_ROOT_TOKEN_QUERY,
	PERM_ROOT_TOKEN,
}