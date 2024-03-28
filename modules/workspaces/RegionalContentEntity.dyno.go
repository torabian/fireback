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
	seeders "github.com/torabian/fireback/modules/workspaces/seeders/RegionalContent"
	metas "github.com/torabian/fireback/modules/workspaces/metas"
)
type RegionalContentEntity struct {
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
    Content   *string `json:"content" yaml:"content"  validate:"required"       `
    // Datenano also has a text representation
    ContentExcerpt * string `json:"contentExcerpt" yaml:"contentExcerpt"`
    Region   *string `json:"region" yaml:"region"  validate:"required"       `
    // Datenano also has a text representation
    Title   *string `json:"title" yaml:"title"       `
    // Datenano also has a text representation
    LanguageId   *string `json:"languageId" yaml:"languageId"  validate:"required"    gorm:"index:regional_content_index,unique"     `
    // Datenano also has a text representation
    KeyGroup   *string `json:"keyGroup" yaml:"keyGroup"  validate:"required"    gorm:"index:regional_content_index,unique"     `
    // Datenano also has a text representation
    Children []*RegionalContentEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *RegionalContentEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var RegionalContentPreloadRelations []string = []string{}
var REGIONAL_CONTENT_EVENT_CREATED = "regionalContent.created"
var REGIONAL_CONTENT_EVENT_UPDATED = "regionalContent.updated"
var REGIONAL_CONTENT_EVENT_DELETED = "regionalContent.deleted"
var REGIONAL_CONTENT_EVENTS = []string{
	REGIONAL_CONTENT_EVENT_CREATED,
	REGIONAL_CONTENT_EVENT_UPDATED,
	REGIONAL_CONTENT_EVENT_DELETED,
}
type RegionalContentFieldMap struct {
		Content TranslatedString `yaml:"content"`
		Region TranslatedString `yaml:"region"`
		Title TranslatedString `yaml:"title"`
		LanguageId TranslatedString `yaml:"languageId"`
		KeyGroup TranslatedString `yaml:"keyGroup"`
}
var RegionalContentEntityMetaConfig map[string]int64 = map[string]int64{
            "ContentExcerptSize": 100,
}
var RegionalContentEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&RegionalContentEntity{}))
func entityRegionalContentFormatter(dto *RegionalContentEntity, query QueryDSL) {
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
func RegionalContentMockEntity() *RegionalContentEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &RegionalContentEntity{
      Region : &stringHolder,
      Title : &stringHolder,
      LanguageId : &stringHolder,
      KeyGroup : &stringHolder,
	}
	return entity
}
func RegionalContentActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := RegionalContentMockEntity()
		_, err := RegionalContentActionCreate(entity, query)
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
  func RegionalContentActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*RegionalContentEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &RegionalContentEntity{
          Region: &tildaRef,
          Title: &tildaRef,
          LanguageId: &tildaRef,
          KeyGroup: &tildaRef,
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
  func RegionalContentAssociationCreate(dto *RegionalContentEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func RegionalContentRelationContentCreate(dto *RegionalContentEntity, query QueryDSL) error {
return nil
}
func RegionalContentRelationContentUpdate(dto *RegionalContentEntity, query QueryDSL) error {
	return nil
}
func RegionalContentPolyglotCreateHandler(dto *RegionalContentEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func RegionalContentValidator(dto *RegionalContentEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func RegionalContentEntityPreSanitize(dto *RegionalContentEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
			if (dto.Content != nil ) {
          Content := *dto.Content
          ContentExcerpt := stripPolicy.Sanitize(*dto.Content)
            Content = ugcPolicy.Sanitize(Content)
            ContentExcerpt = stripPolicy.Sanitize(ContentExcerpt)
        ContentExcerptSize, ContentExcerptSizeExists := RegionalContentEntityMetaConfig["ContentExcerptSize"]
        if ContentExcerptSizeExists {
          ContentExcerpt = PickFirstNWords(ContentExcerpt, int(ContentExcerptSize))
        } else {
          ContentExcerpt = PickFirstNWords(ContentExcerpt, 30)
        }
        dto.ContentExcerpt = &ContentExcerpt
        dto.Content = &Content
      }
}
  func RegionalContentEntityBeforeCreateAppend(dto *RegionalContentEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    RegionalContentRecursiveAddUniqueId(dto, query)
  }
  func RegionalContentRecursiveAddUniqueId(dto *RegionalContentEntity, query QueryDSL) {
  }
func RegionalContentActionBatchCreateFn(dtos []*RegionalContentEntity, query QueryDSL) ([]*RegionalContentEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*RegionalContentEntity{}
		for _, item := range dtos {
			s, err := RegionalContentActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func RegionalContentDeleteEntireChildren(query QueryDSL, dto *RegionalContentEntity) (*IError) {
  return nil
}
func RegionalContentActionCreateFn(dto *RegionalContentEntity, query QueryDSL) (*RegionalContentEntity, *IError) {
	// 1. Validate always
	if iError := RegionalContentValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	RegionalContentEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	RegionalContentEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	RegionalContentPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	RegionalContentRelationContentCreate(dto, query)
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
	RegionalContentAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(REGIONAL_CONTENT_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&RegionalContentEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func RegionalContentActionGetOne(query QueryDSL) (*RegionalContentEntity, *IError) {
    refl := reflect.ValueOf(&RegionalContentEntity{})
    item, err := GetOneEntity[RegionalContentEntity](query, refl)
    entityRegionalContentFormatter(item, query)
    return item, err
  }
  func RegionalContentActionQuery(query QueryDSL) ([]*RegionalContentEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&RegionalContentEntity{})
    items, meta, err := QueryEntitiesPointer[RegionalContentEntity](query, refl)
    for _, item := range items {
      entityRegionalContentFormatter(item, query)
    }
    return items, meta, err
  }
  func RegionalContentUpdateExec(dbref *gorm.DB, query QueryDSL, fields *RegionalContentEntity) (*RegionalContentEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = REGIONAL_CONTENT_EVENT_UPDATED
    RegionalContentEntityPreSanitize(fields, query)
    var item RegionalContentEntity
    q := dbref.
      Where(&RegionalContentEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    RegionalContentRelationContentUpdate(fields, query)
    RegionalContentPolyglotCreateHandler(fields, query)
    if ero := RegionalContentDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&RegionalContentEntity{UniqueId: uniqueId}).
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
  func RegionalContentActionUpdateFn(query QueryDSL, fields *RegionalContentEntity) (*RegionalContentEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := RegionalContentValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // RegionalContentRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      var item *RegionalContentEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *IError
        item, err = RegionalContentUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, CastToIError(vf)
    } else {
      dbref = query.Tx
      return RegionalContentUpdateExec(dbref, query, fields)
    }
  }
var RegionalContentWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire regionalcontents ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
      ActionRequires: []PermissionInfo{PERM_ROOT_REGIONAL_CONTENT_DELETE},
    })
		count, _ := RegionalContentActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func RegionalContentActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&RegionalContentEntity{})
	query.ActionRequires = []PermissionInfo{PERM_ROOT_REGIONAL_CONTENT_DELETE}
	return RemoveEntity[RegionalContentEntity](query, refl)
}
func RegionalContentActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[RegionalContentEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'RegionalContentEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func RegionalContentActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[RegionalContentEntity]) (
    *BulkRecordRequest[RegionalContentEntity], *IError,
  ) {
    result := []*RegionalContentEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := RegionalContentActionUpdate(query, record)
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
func (x *RegionalContentEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var RegionalContentEntityMeta = TableMetaData{
	EntityName:    "RegionalContent",
	ExportKey:    "regional-contents",
	TableNameInDb: "fb_regional-content_entities",
	EntityObject:  &RegionalContentEntity{},
	ExportStream: RegionalContentActionExportT,
	ImportQuery: RegionalContentActionImport,
}
func RegionalContentActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[RegionalContentEntity](query, RegionalContentActionQuery, RegionalContentPreloadRelations)
}
func RegionalContentActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[RegionalContentEntity](query, RegionalContentActionQuery, RegionalContentPreloadRelations)
}
func RegionalContentActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content RegionalContentEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := RegionalContentActionCreate(&content, query)
	return err
}
var RegionalContentCommonCliFlags = []cli.Flag{
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
      Name:     "content",
      Required: true,
      Usage:    "content",
    },
    &cli.StringFlag{
      Name:     "region",
      Required: true,
      Usage:    "region",
    },
    &cli.StringFlag{
      Name:     "title",
      Required: false,
      Usage:    "title",
    },
    &cli.StringFlag{
      Name:     "language-id",
      Required: true,
      Usage:    "languageId",
    },
    &cli.StringFlag{
      Name:     "key-group",
      Required: true,
      Usage:    "One of: 'SMS_OTP', 'EMAIL_OTP'",
    },
}
var RegionalContentCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:     "region",
		StructField:     "Region",
		Required: true,
		Usage:    "region",
		Type: "string",
	},
	{
		Name:     "title",
		StructField:     "Title",
		Required: false,
		Usage:    "title",
		Type: "string",
	},
	{
		Name:     "languageId",
		StructField:     "LanguageId",
		Required: true,
		Usage:    "languageId",
		Type: "string",
	},
	{
		Name:     "keyGroup",
		StructField:     "KeyGroup",
		Required: true,
		Usage:    "One of: 'SMS_OTP', 'EMAIL_OTP'",
		Type: "string",
	},
}
var RegionalContentCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "content",
      Required: true,
      Usage:    "content",
    },
    &cli.StringFlag{
      Name:     "region",
      Required: true,
      Usage:    "region",
    },
    &cli.StringFlag{
      Name:     "title",
      Required: false,
      Usage:    "title",
    },
    &cli.StringFlag{
      Name:     "language-id",
      Required: true,
      Usage:    "languageId",
    },
    &cli.StringFlag{
      Name:     "key-group",
      Required: true,
      Usage:    "One of: 'SMS_OTP', 'EMAIL_OTP'",
    },
}
  var RegionalContentCreateCmd cli.Command = REGIONAL_CONTENT_ACTION_POST_ONE.ToCli()
  var RegionalContentCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_REGIONAL_CONTENT_CREATE},
      })
      entity := &RegionalContentEntity{}
      for _, item := range RegionalContentCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := RegionalContentActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var RegionalContentUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: RegionalContentCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_REGIONAL_CONTENT_UPDATE},
      })
      entity := CastRegionalContentFromCli(c)
      if entity, err := RegionalContentActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* RegionalContentEntity) FromCli(c *cli.Context) *RegionalContentEntity {
	return CastRegionalContentFromCli(c)
}
func CastRegionalContentFromCli (c *cli.Context) *RegionalContentEntity {
	template := &RegionalContentEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("content") {
        value := c.String("content")
        template.Content = &value
      }
      if c.IsSet("region") {
        value := c.String("region")
        template.Region = &value
      }
      if c.IsSet("title") {
        value := c.String("title")
        template.Title = &value
      }
      if c.IsSet("language-id") {
        value := c.String("language-id")
        template.LanguageId = &value
      }
      if c.IsSet("key-group") {
        value := c.String("key-group")
        template.KeyGroup = &value
      }
	return template
}
  func RegionalContentSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      RegionalContentActionCreate,
      reflect.ValueOf(&RegionalContentEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func RegionalContentSyncSeeders() {
    SeederFromFSImport(
      QueryDSL{WorkspaceId: USER_SYSTEM},
      RegionalContentActionCreate,
      reflect.ValueOf(&RegionalContentEntity{}).Elem(),
      &seeders.ViewsFs,
      []string{},
      true,
    )
  }
  func RegionalContentWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := RegionalContentActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "RegionalContent", result)
    }
  }
var RegionalContentImportExportCommands = []cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_REGIONAL_CONTENT_CREATE},
      })
			RegionalContentActionSeeder(query, c.Int("count"))
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
				Value: "regional-content-seeder.yml",
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
        ActionRequires: []PermissionInfo{PERM_ROOT_REGIONAL_CONTENT_CREATE},
      })
			RegionalContentActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "regional-content-seeder-regional-content.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of regional-contents, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]RegionalContentEntity{}
			ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
	cli.Command{
		Name:  "list",
		Usage: "Prints the list of files attached to this module for syncing or bootstrapping project",
		Action: func(c *cli.Context) error {
			if entity, err := GetSeederFilenames(&seeders.ViewsFs, ""); err != nil {
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
			CommonCliImportEmbedCmd(c,
				RegionalContentActionCreate,
				reflect.ValueOf(&RegionalContentEntity{}).Elem(),
				&seeders.ViewsFs,
			)
			return nil
		},
	},
	cli.Command{
		Name:    "export",
		Aliases: []string{"e"},
		Flags: append(CommonQueryFlags,
			&cli.StringFlag{
				Name:     "file",
				Usage:    "The address of file you want the csv/yaml/json be exported to",
				Required: true,
			}),
		Usage: "Exports a query results into the csv/yaml/json format",
		Action: func(c *cli.Context) error {
			CommonCliExportCmd(c,
				RegionalContentActionQuery,
				reflect.ValueOf(&RegionalContentEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"RegionalContentFieldMap.yml",
				RegionalContentPreloadRelations,
			)
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
			RegionalContentCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				RegionalContentActionCreate,
				reflect.ValueOf(&RegionalContentEntity{}).Elem(),
				c.String("file"),
        &SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_REGIONAL_CONTENT_CREATE},
				},
        func() RegionalContentEntity {
					v := CastRegionalContentFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var RegionalContentCliCommands []cli.Command = []cli.Command{
      REGIONAL_CONTENT_ACTION_QUERY.ToCli(),
      REGIONAL_CONTENT_ACTION_TABLE.ToCli(),
      RegionalContentCreateCmd,
      RegionalContentUpdateCmd,
      RegionalContentCreateInteractiveCmd,
      RegionalContentWipeCmd,
      GetCommonRemoveQuery(reflect.ValueOf(&RegionalContentEntity{}).Elem(), RegionalContentActionRemove),
  }
  func RegionalContentCliFn() cli.Command {
    RegionalContentCliCommands = append(RegionalContentCliCommands, RegionalContentImportExportCommands...)
    return cli.Command{
      Name:        "regionalContent",
      ShortName:   "rc",
      Description: "RegionalContents module actions (sample module to handle complex entities)",
      Usage:       "Email templates, sms templates or other textual content which can be accessed.",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: RegionalContentCliCommands,
    }
  }
var REGIONAL_CONTENT_ACTION_TABLE = Module2Action{
  Name:    "table",
  ActionAliases: []string{"t"},
  Flags:  CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: RegionalContentActionQuery,
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    CommonCliTableCmd2(c,
      RegionalContentActionQuery,
      security,
      reflect.ValueOf(&RegionalContentEntity{}).Elem(),
    )
    return nil
  },
}
var REGIONAL_CONTENT_ACTION_QUERY = Module2Action{
  Method: "GET",
  Url:    "/regional-contents",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_REGIONAL_CONTENT_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpQueryEntity(c, RegionalContentActionQuery)
    },
  },
  Format: "QUERY",
  Action: RegionalContentActionQuery,
  ResponseEntity: &[]RegionalContentEntity{},
  CliAction: func(c *cli.Context, security *SecurityModel) error {
		CommonCliQueryCmd2(
			c,
			RegionalContentActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionAliases: []string{"q"},
	Flags:         CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var REGIONAL_CONTENT_ACTION_EXPORT = Module2Action{
  Method: "GET",
  Url:    "/regional-contents/export",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_REGIONAL_CONTENT_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpStreamFileChannel(c, RegionalContentActionExport)
    },
  },
  Format: "QUERY",
  Action: RegionalContentActionExport,
  ResponseEntity: &[]RegionalContentEntity{},
}
var REGIONAL_CONTENT_ACTION_GET_ONE = Module2Action{
  Method: "GET",
  Url:    "/regional-content/:uniqueId",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_REGIONAL_CONTENT_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpGetEntity(c, RegionalContentActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: RegionalContentActionGetOne,
  ResponseEntity: &RegionalContentEntity{},
}
var REGIONAL_CONTENT_ACTION_POST_ONE = Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new regionalContent",
  Flags: RegionalContentCommonCliFlags,
  Method: "POST",
  Url:    "/regional-content",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_REGIONAL_CONTENT_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpPostEntity(c, RegionalContentActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    result, err := CliPostEntity(c, RegionalContentActionCreate, security)
    HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: RegionalContentActionCreate,
  Format: "POST_ONE",
  RequestEntity: &RegionalContentEntity{},
  ResponseEntity: &RegionalContentEntity{},
}
var REGIONAL_CONTENT_ACTION_PATCH = Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: RegionalContentCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/regional-content",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_REGIONAL_CONTENT_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntity(c, RegionalContentActionUpdate)
    },
  },
  Action: RegionalContentActionUpdate,
  RequestEntity: &RegionalContentEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &RegionalContentEntity{},
}
var REGIONAL_CONTENT_ACTION_PATCH_BULK = Module2Action{
  Method: "PATCH",
  Url:    "/regional-contents",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_REGIONAL_CONTENT_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntities(c, RegionalContentActionBulkUpdate)
    },
  },
  Action: RegionalContentActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &BulkRecordRequest[RegionalContentEntity]{},
  ResponseEntity: &BulkRecordRequest[RegionalContentEntity]{},
}
var REGIONAL_CONTENT_ACTION_DELETE = Module2Action{
  Method: "DELETE",
  Url:    "/regional-content",
  Format: "DELETE_DSL",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_REGIONAL_CONTENT_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpRemoveEntity(c, RegionalContentActionRemove)
    },
  },
  Action: RegionalContentActionRemove,
  RequestEntity: &DeleteRequest{},
  ResponseEntity: &DeleteResponse{},
  TargetEntity: &RegionalContentEntity{},
}
  /**
  *	Override this function on RegionalContentEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendRegionalContentRouter = func(r *[]Module2Action) {}
  func GetRegionalContentModule2Actions() []Module2Action {
    routes := []Module2Action{
      REGIONAL_CONTENT_ACTION_QUERY,
      REGIONAL_CONTENT_ACTION_EXPORT,
      REGIONAL_CONTENT_ACTION_GET_ONE,
      REGIONAL_CONTENT_ACTION_POST_ONE,
      REGIONAL_CONTENT_ACTION_PATCH,
      REGIONAL_CONTENT_ACTION_PATCH_BULK,
      REGIONAL_CONTENT_ACTION_DELETE,
    }
    // Append user defined functions
    AppendRegionalContentRouter(&routes)
    return routes
  }
  func CreateRegionalContentRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetRegionalContentModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, RegionalContentEntityJsonSchema, "regional-content-http", "workspaces")
    WriteEntitySchema("RegionalContentEntity", RegionalContentEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_REGIONAL_CONTENT_DELETE = PermissionInfo{
  CompleteKey: "root/workspaces/regional-content/delete",
  Name: "Delete regional content",
}
var PERM_ROOT_REGIONAL_CONTENT_CREATE = PermissionInfo{
  CompleteKey: "root/workspaces/regional-content/create",
  Name: "Create regional content",
}
var PERM_ROOT_REGIONAL_CONTENT_UPDATE = PermissionInfo{
  CompleteKey: "root/workspaces/regional-content/update",
  Name: "Update regional content",
}
var PERM_ROOT_REGIONAL_CONTENT_QUERY = PermissionInfo{
  CompleteKey: "root/workspaces/regional-content/query",
  Name: "Query regional content",
}
var PERM_ROOT_REGIONAL_CONTENT = PermissionInfo{
  CompleteKey: "root/workspaces/regional-content/*",
  Name: "Entire regional content actions (*)",
}
var ALL_REGIONAL_CONTENT_PERMISSIONS = []PermissionInfo{
	PERM_ROOT_REGIONAL_CONTENT_DELETE,
	PERM_ROOT_REGIONAL_CONTENT_CREATE,
	PERM_ROOT_REGIONAL_CONTENT_UPDATE,
	PERM_ROOT_REGIONAL_CONTENT_QUERY,
	PERM_ROOT_REGIONAL_CONTENT,
}
var RegionalContentKeyGroup = newRegionalContentKeyGroup()
func newRegionalContentKeyGroup() *xRegionalContentKeyGroup {
	return &xRegionalContentKeyGroup{
      SMS_OTP: "SMS_OTP",
      EMAIL_OTP: "EMAIL_OTP",
	}
}
type xRegionalContentKeyGroup struct {
    SMS_OTP string
    EMAIL_OTP string
}