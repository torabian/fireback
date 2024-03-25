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
type WorkspaceInviteEntity struct {
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
    CoverLetter   *string `json:"coverLetter" yaml:"coverLetter"       `
    // Datenano also has a text representation
    TargetUserLocale   *string `json:"targetUserLocale" yaml:"targetUserLocale"       `
    // Datenano also has a text representation
    Value   *string `json:"value" yaml:"value"  validate:"required"       `
    // Datenano also has a text representation
    Workspace   *  WorkspaceEntity `json:"workspace" yaml:"workspace"    gorm:"foreignKey:WorkspaceId;references:UniqueId"     `
    // Datenano also has a text representation
    FirstName   *string `json:"firstName" yaml:"firstName"  validate:"required"       `
    // Datenano also has a text representation
    LastName   *string `json:"lastName" yaml:"lastName"  validate:"required"       `
    // Datenano also has a text representation
    Used   *bool `json:"used" yaml:"used"       `
    // Datenano also has a text representation
    Role   *  RoleEntity `json:"role" yaml:"role"    gorm:"foreignKey:RoleId;references:UniqueId"     `
    // Datenano also has a text representation
        RoleId *string `json:"roleId" yaml:"roleId" validate:"required" `
    Children []*WorkspaceInviteEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *WorkspaceInviteEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var WorkspaceInvitePreloadRelations []string = []string{}
var WORKSPACE_INVITE_EVENT_CREATED = "workspaceInvite.created"
var WORKSPACE_INVITE_EVENT_UPDATED = "workspaceInvite.updated"
var WORKSPACE_INVITE_EVENT_DELETED = "workspaceInvite.deleted"
var WORKSPACE_INVITE_EVENTS = []string{
	WORKSPACE_INVITE_EVENT_CREATED,
	WORKSPACE_INVITE_EVENT_UPDATED,
	WORKSPACE_INVITE_EVENT_DELETED,
}
type WorkspaceInviteFieldMap struct {
		CoverLetter TranslatedString `yaml:"coverLetter"`
		TargetUserLocale TranslatedString `yaml:"targetUserLocale"`
		Value TranslatedString `yaml:"value"`
		Workspace TranslatedString `yaml:"workspace"`
		FirstName TranslatedString `yaml:"firstName"`
		LastName TranslatedString `yaml:"lastName"`
		Used TranslatedString `yaml:"used"`
		Role TranslatedString `yaml:"role"`
}
var WorkspaceInviteEntityMetaConfig map[string]int64 = map[string]int64{
}
var WorkspaceInviteEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&WorkspaceInviteEntity{}))
func entityWorkspaceInviteFormatter(dto *WorkspaceInviteEntity, query QueryDSL) {
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
func WorkspaceInviteMockEntity() *WorkspaceInviteEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &WorkspaceInviteEntity{
      CoverLetter : &stringHolder,
      TargetUserLocale : &stringHolder,
      Value : &stringHolder,
      FirstName : &stringHolder,
      LastName : &stringHolder,
	}
	return entity
}
func WorkspaceInviteActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := WorkspaceInviteMockEntity()
		_, err := WorkspaceInviteActionCreate(entity, query)
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
  func WorkspaceInviteActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*WorkspaceInviteEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &WorkspaceInviteEntity{
          CoverLetter: &tildaRef,
          TargetUserLocale: &tildaRef,
          Value: &tildaRef,
          FirstName: &tildaRef,
          LastName: &tildaRef,
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
  func WorkspaceInviteAssociationCreate(dto *WorkspaceInviteEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func WorkspaceInviteRelationContentCreate(dto *WorkspaceInviteEntity, query QueryDSL) error {
return nil
}
func WorkspaceInviteRelationContentUpdate(dto *WorkspaceInviteEntity, query QueryDSL) error {
	return nil
}
func WorkspaceInvitePolyglotCreateHandler(dto *WorkspaceInviteEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func WorkspaceInviteValidator(dto *WorkspaceInviteEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func WorkspaceInviteEntityPreSanitize(dto *WorkspaceInviteEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func WorkspaceInviteEntityBeforeCreateAppend(dto *WorkspaceInviteEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    WorkspaceInviteRecursiveAddUniqueId(dto, query)
  }
  func WorkspaceInviteRecursiveAddUniqueId(dto *WorkspaceInviteEntity, query QueryDSL) {
  }
func WorkspaceInviteActionBatchCreateFn(dtos []*WorkspaceInviteEntity, query QueryDSL) ([]*WorkspaceInviteEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*WorkspaceInviteEntity{}
		for _, item := range dtos {
			s, err := WorkspaceInviteActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func WorkspaceInviteDeleteEntireChildren(query QueryDSL, dto *WorkspaceInviteEntity) (*IError) {
  return nil
}
func WorkspaceInviteActionCreateFn(dto *WorkspaceInviteEntity, query QueryDSL) (*WorkspaceInviteEntity, *IError) {
	// 1. Validate always
	if iError := WorkspaceInviteValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	WorkspaceInviteEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	WorkspaceInviteEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	WorkspaceInvitePolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	WorkspaceInviteRelationContentCreate(dto, query)
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
	WorkspaceInviteAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(WORKSPACE_INVITE_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&WorkspaceInviteEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func WorkspaceInviteActionGetOne(query QueryDSL) (*WorkspaceInviteEntity, *IError) {
    refl := reflect.ValueOf(&WorkspaceInviteEntity{})
    item, err := GetOneEntity[WorkspaceInviteEntity](query, refl)
    entityWorkspaceInviteFormatter(item, query)
    return item, err
  }
  func WorkspaceInviteActionQuery(query QueryDSL) ([]*WorkspaceInviteEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&WorkspaceInviteEntity{})
    items, meta, err := QueryEntitiesPointer[WorkspaceInviteEntity](query, refl)
    for _, item := range items {
      entityWorkspaceInviteFormatter(item, query)
    }
    return items, meta, err
  }
  func WorkspaceInviteUpdateExec(dbref *gorm.DB, query QueryDSL, fields *WorkspaceInviteEntity) (*WorkspaceInviteEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = WORKSPACE_INVITE_EVENT_UPDATED
    WorkspaceInviteEntityPreSanitize(fields, query)
    var item WorkspaceInviteEntity
    q := dbref.
      Where(&WorkspaceInviteEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    WorkspaceInviteRelationContentUpdate(fields, query)
    WorkspaceInvitePolyglotCreateHandler(fields, query)
    if ero := WorkspaceInviteDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&WorkspaceInviteEntity{UniqueId: uniqueId}).
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
  func WorkspaceInviteActionUpdateFn(query QueryDSL, fields *WorkspaceInviteEntity) (*WorkspaceInviteEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := WorkspaceInviteValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // WorkspaceInviteRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      var item *WorkspaceInviteEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *IError
        item, err = WorkspaceInviteUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, CastToIError(vf)
    } else {
      dbref = query.Tx
      return WorkspaceInviteUpdateExec(dbref, query, fields)
    }
  }
var WorkspaceInviteWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire workspaceinvites ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
      ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_INVITE_DELETE},
    })
		count, _ := WorkspaceInviteActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func WorkspaceInviteActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&WorkspaceInviteEntity{})
	query.ActionRequires = []PermissionInfo{PERM_ROOT_WORKSPACE_INVITE_DELETE}
	return RemoveEntity[WorkspaceInviteEntity](query, refl)
}
func WorkspaceInviteActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[WorkspaceInviteEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'WorkspaceInviteEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func WorkspaceInviteActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[WorkspaceInviteEntity]) (
    *BulkRecordRequest[WorkspaceInviteEntity], *IError,
  ) {
    result := []*WorkspaceInviteEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := WorkspaceInviteActionUpdate(query, record)
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
func (x *WorkspaceInviteEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var WorkspaceInviteEntityMeta = TableMetaData{
	EntityName:    "WorkspaceInvite",
	ExportKey:    "workspace-invites",
	TableNameInDb: "fb_workspace-invite_entities",
	EntityObject:  &WorkspaceInviteEntity{},
	ExportStream: WorkspaceInviteActionExportT,
	ImportQuery: WorkspaceInviteActionImport,
}
func WorkspaceInviteActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[WorkspaceInviteEntity](query, WorkspaceInviteActionQuery, WorkspaceInvitePreloadRelations)
}
func WorkspaceInviteActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[WorkspaceInviteEntity](query, WorkspaceInviteActionQuery, WorkspaceInvitePreloadRelations)
}
func WorkspaceInviteActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content WorkspaceInviteEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := WorkspaceInviteActionCreate(&content, query)
	return err
}
var WorkspaceInviteCommonCliFlags = []cli.Flag{
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
      Name:     "cover-letter",
      Required: false,
      Usage:    "coverLetter",
    },
    &cli.StringFlag{
      Name:     "target-user-locale",
      Required: false,
      Usage:    "targetUserLocale",
    },
    &cli.StringFlag{
      Name:     "value",
      Required: true,
      Usage:    "value",
    },
    &cli.StringFlag{
      Name:     "workspace-id",
      Required: true,
      Usage:    "workspace",
    },
    &cli.StringFlag{
      Name:     "first-name",
      Required: true,
      Usage:    "firstName",
    },
    &cli.StringFlag{
      Name:     "last-name",
      Required: true,
      Usage:    "lastName",
    },
    &cli.BoolFlag{
      Name:     "used",
      Required: false,
      Usage:    "used",
    },
    &cli.StringFlag{
      Name:     "role-id",
      Required: true,
      Usage:    "role",
    },
}
var WorkspaceInviteCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:     "coverLetter",
		StructField:     "CoverLetter",
		Required: false,
		Usage:    "coverLetter",
		Type: "string",
	},
	{
		Name:     "targetUserLocale",
		StructField:     "TargetUserLocale",
		Required: false,
		Usage:    "targetUserLocale",
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
		Name:     "firstName",
		StructField:     "FirstName",
		Required: true,
		Usage:    "firstName",
		Type: "string",
	},
	{
		Name:     "lastName",
		StructField:     "LastName",
		Required: true,
		Usage:    "lastName",
		Type: "string",
	},
	{
		Name:     "used",
		StructField:     "Used",
		Required: false,
		Usage:    "used",
		Type: "bool",
	},
}
var WorkspaceInviteCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "cover-letter",
      Required: false,
      Usage:    "coverLetter",
    },
    &cli.StringFlag{
      Name:     "target-user-locale",
      Required: false,
      Usage:    "targetUserLocale",
    },
    &cli.StringFlag{
      Name:     "value",
      Required: true,
      Usage:    "value",
    },
    &cli.StringFlag{
      Name:     "workspace-id",
      Required: true,
      Usage:    "workspace",
    },
    &cli.StringFlag{
      Name:     "first-name",
      Required: true,
      Usage:    "firstName",
    },
    &cli.StringFlag{
      Name:     "last-name",
      Required: true,
      Usage:    "lastName",
    },
    &cli.BoolFlag{
      Name:     "used",
      Required: false,
      Usage:    "used",
    },
    &cli.StringFlag{
      Name:     "role-id",
      Required: true,
      Usage:    "role",
    },
}
  var WorkspaceInviteCreateCmd cli.Command = WORKSPACE_INVITE_ACTION_POST_ONE.ToCli()
  var WorkspaceInviteCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_INVITE_CREATE},
      })
      entity := &WorkspaceInviteEntity{}
      for _, item := range WorkspaceInviteCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := WorkspaceInviteActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var WorkspaceInviteUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: WorkspaceInviteCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_INVITE_UPDATE},
      })
      entity := CastWorkspaceInviteFromCli(c)
      if entity, err := WorkspaceInviteActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* WorkspaceInviteEntity) FromCli(c *cli.Context) *WorkspaceInviteEntity {
	return CastWorkspaceInviteFromCli(c)
}
func CastWorkspaceInviteFromCli (c *cli.Context) *WorkspaceInviteEntity {
	template := &WorkspaceInviteEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("cover-letter") {
        value := c.String("cover-letter")
        template.CoverLetter = &value
      }
      if c.IsSet("target-user-locale") {
        value := c.String("target-user-locale")
        template.TargetUserLocale = &value
      }
      if c.IsSet("value") {
        value := c.String("value")
        template.Value = &value
      }
      if c.IsSet("workspace-id") {
        value := c.String("workspace-id")
        template.WorkspaceId = &value
      }
      if c.IsSet("first-name") {
        value := c.String("first-name")
        template.FirstName = &value
      }
      if c.IsSet("last-name") {
        value := c.String("last-name")
        template.LastName = &value
      }
      if c.IsSet("role-id") {
        value := c.String("role-id")
        template.RoleId = &value
      }
	return template
}
  func WorkspaceInviteSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      WorkspaceInviteActionCreate,
      reflect.ValueOf(&WorkspaceInviteEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func WorkspaceInviteWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := WorkspaceInviteActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "WorkspaceInvite", result)
    }
  }
var WorkspaceInviteImportExportCommands = []cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_INVITE_CREATE},
      })
			WorkspaceInviteActionSeeder(query, c.Int("count"))
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
				Value: "workspace-invite-seeder.yml",
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
        ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_INVITE_CREATE},
      })
			WorkspaceInviteActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "workspace-invite-seeder-workspace-invite.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of workspace-invites, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]WorkspaceInviteEntity{}
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
			WorkspaceInviteCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				WorkspaceInviteActionCreate,
				reflect.ValueOf(&WorkspaceInviteEntity{}).Elem(),
				c.String("file"),
        &SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_INVITE_CREATE},
				},
        func() WorkspaceInviteEntity {
					v := CastWorkspaceInviteFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var WorkspaceInviteCliCommands []cli.Command = []cli.Command{
      WORKSPACE_INVITE_ACTION_QUERY.ToCli(),
      WORKSPACE_INVITE_ACTION_TABLE.ToCli(),
      GetCommonTableQuery(reflect.ValueOf(&WorkspaceInviteEntity{}).Elem(), WorkspaceInviteActionQuery),
      WorkspaceInviteCreateCmd,
      WorkspaceInviteUpdateCmd,
      WorkspaceInviteCreateInteractiveCmd,
      WorkspaceInviteWipeCmd,
      GetCommonRemoveQuery(reflect.ValueOf(&WorkspaceInviteEntity{}).Elem(), WorkspaceInviteActionRemove),
  }
  func WorkspaceInviteCliFn() cli.Command {
    WorkspaceInviteCliCommands = append(WorkspaceInviteCliCommands, WorkspaceInviteImportExportCommands...)
    return cli.Command{
      Name:        "workspaceInvite",
      Description: "WorkspaceInvites module actions (sample module to handle complex entities)",
      Usage:       "Active invitations for non-users or already users to join an specific workspace",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: WorkspaceInviteCliCommands,
    }
  }
var WORKSPACE_INVITE_ACTION_TABLE = Module2Action{
  Name:    "table",
  ActionAliases: []string{"t"},
  Flags:  CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: WorkspaceInviteActionQuery,
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    CommonCliTableCmd2(c,
      WorkspaceInviteActionQuery,
      security,
      reflect.ValueOf(&WorkspaceInviteEntity{}).Elem(),
    )
    return nil
  },
}
var WORKSPACE_INVITE_ACTION_QUERY = Module2Action{
  Method: "GET",
  Url:    "/workspace-invites",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_INVITE_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpQueryEntity(c, WorkspaceInviteActionQuery)
    },
  },
  Format: "QUERY",
  Action: WorkspaceInviteActionQuery,
  ResponseEntity: &[]WorkspaceInviteEntity{},
  CliAction: func(c *cli.Context, security *SecurityModel) error {
		CommonCliQueryCmd2(
			c,
			WorkspaceInviteActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionAliases: []string{"q"},
	Flags:         CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var WORKSPACE_INVITE_ACTION_EXPORT = Module2Action{
  Method: "GET",
  Url:    "/workspace-invites/export",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_INVITE_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpStreamFileChannel(c, WorkspaceInviteActionExport)
    },
  },
  Format: "QUERY",
  Action: WorkspaceInviteActionExport,
  ResponseEntity: &[]WorkspaceInviteEntity{},
}
var WORKSPACE_INVITE_ACTION_GET_ONE = Module2Action{
  Method: "GET",
  Url:    "/workspace-invite/:uniqueId",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_INVITE_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpGetEntity(c, WorkspaceInviteActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: WorkspaceInviteActionGetOne,
  ResponseEntity: &WorkspaceInviteEntity{},
}
var WORKSPACE_INVITE_ACTION_POST_ONE = Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new workspaceInvite",
  Flags: WorkspaceInviteCommonCliFlags,
  Method: "POST",
  Url:    "/workspace-invite",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_INVITE_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpPostEntity(c, WorkspaceInviteActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    result, err := CliPostEntity(c, WorkspaceInviteActionCreate, security)
    HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: WorkspaceInviteActionCreate,
  Format: "POST_ONE",
  RequestEntity: &WorkspaceInviteEntity{},
  ResponseEntity: &WorkspaceInviteEntity{},
}
var WORKSPACE_INVITE_ACTION_PATCH = Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: WorkspaceInviteCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/workspace-invite",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_INVITE_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntity(c, WorkspaceInviteActionUpdate)
    },
  },
  Action: WorkspaceInviteActionUpdate,
  RequestEntity: &WorkspaceInviteEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &WorkspaceInviteEntity{},
}
var WORKSPACE_INVITE_ACTION_PATCH_BULK = Module2Action{
  Method: "PATCH",
  Url:    "/workspace-invites",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_INVITE_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntities(c, WorkspaceInviteActionBulkUpdate)
    },
  },
  Action: WorkspaceInviteActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &BulkRecordRequest[WorkspaceInviteEntity]{},
  ResponseEntity: &BulkRecordRequest[WorkspaceInviteEntity]{},
}
var WORKSPACE_INVITE_ACTION_DELETE = Module2Action{
  Method: "DELETE",
  Url:    "/workspace-invite",
  Format: "DELETE_DSL",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_INVITE_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpRemoveEntity(c, WorkspaceInviteActionRemove)
    },
  },
  Action: WorkspaceInviteActionRemove,
  RequestEntity: &DeleteRequest{},
  ResponseEntity: &DeleteResponse{},
  TargetEntity: &WorkspaceInviteEntity{},
}
  /**
  *	Override this function on WorkspaceInviteEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendWorkspaceInviteRouter = func(r *[]Module2Action) {}
  func GetWorkspaceInviteModule2Actions() []Module2Action {
    routes := []Module2Action{
      WORKSPACE_INVITE_ACTION_QUERY,
      WORKSPACE_INVITE_ACTION_EXPORT,
      WORKSPACE_INVITE_ACTION_GET_ONE,
      WORKSPACE_INVITE_ACTION_POST_ONE,
      WORKSPACE_INVITE_ACTION_PATCH,
      WORKSPACE_INVITE_ACTION_PATCH_BULK,
      WORKSPACE_INVITE_ACTION_DELETE,
    }
    // Append user defined functions
    AppendWorkspaceInviteRouter(&routes)
    return routes
  }
  func CreateWorkspaceInviteRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetWorkspaceInviteModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, WorkspaceInviteEntityJsonSchema, "workspace-invite-http", "workspaces")
    WriteEntitySchema("WorkspaceInviteEntity", WorkspaceInviteEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_WORKSPACE_INVITE_DELETE = PermissionInfo{
  CompleteKey: "root/workspaces/workspace-invite/delete",
}
var PERM_ROOT_WORKSPACE_INVITE_CREATE = PermissionInfo{
  CompleteKey: "root/workspaces/workspace-invite/create",
}
var PERM_ROOT_WORKSPACE_INVITE_UPDATE = PermissionInfo{
  CompleteKey: "root/workspaces/workspace-invite/update",
}
var PERM_ROOT_WORKSPACE_INVITE_QUERY = PermissionInfo{
  CompleteKey: "root/workspaces/workspace-invite/query",
}
var PERM_ROOT_WORKSPACE_INVITE = PermissionInfo{
  CompleteKey: "root/workspaces/workspace-invite/*",
}
var ALL_WORKSPACE_INVITE_PERMISSIONS = []PermissionInfo{
	PERM_ROOT_WORKSPACE_INVITE_DELETE,
	PERM_ROOT_WORKSPACE_INVITE_CREATE,
	PERM_ROOT_WORKSPACE_INVITE_UPDATE,
	PERM_ROOT_WORKSPACE_INVITE_QUERY,
	PERM_ROOT_WORKSPACE_INVITE,
}