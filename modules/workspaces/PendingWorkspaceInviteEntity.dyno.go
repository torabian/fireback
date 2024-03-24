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
type PendingWorkspaceInviteEntity struct {
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
    Value   *string `json:"value" yaml:"value"       `
    // Datenano also has a text representation
    Type   *string `json:"type" yaml:"type"       `
    // Datenano also has a text representation
    CoverLetter   *string `json:"coverLetter" yaml:"coverLetter"       `
    // Datenano also has a text representation
    WorkspaceName   *string `json:"workspaceName" yaml:"workspaceName"       `
    // Datenano also has a text representation
    Role   *  RoleEntity `json:"role" yaml:"role"    gorm:"foreignKey:RoleId;references:UniqueId"     `
    // Datenano also has a text representation
        RoleId *string `json:"roleId" yaml:"roleId"`
    Children []*PendingWorkspaceInviteEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *PendingWorkspaceInviteEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var PendingWorkspaceInvitePreloadRelations []string = []string{}
var PENDING_WORKSPACE_INVITE_EVENT_CREATED = "pendingWorkspaceInvite.created"
var PENDING_WORKSPACE_INVITE_EVENT_UPDATED = "pendingWorkspaceInvite.updated"
var PENDING_WORKSPACE_INVITE_EVENT_DELETED = "pendingWorkspaceInvite.deleted"
var PENDING_WORKSPACE_INVITE_EVENTS = []string{
	PENDING_WORKSPACE_INVITE_EVENT_CREATED,
	PENDING_WORKSPACE_INVITE_EVENT_UPDATED,
	PENDING_WORKSPACE_INVITE_EVENT_DELETED,
}
type PendingWorkspaceInviteFieldMap struct {
		Value TranslatedString `yaml:"value"`
		Type TranslatedString `yaml:"type"`
		CoverLetter TranslatedString `yaml:"coverLetter"`
		WorkspaceName TranslatedString `yaml:"workspaceName"`
		Role TranslatedString `yaml:"role"`
}
var PendingWorkspaceInviteEntityMetaConfig map[string]int64 = map[string]int64{
}
var PendingWorkspaceInviteEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&PendingWorkspaceInviteEntity{}))
func entityPendingWorkspaceInviteFormatter(dto *PendingWorkspaceInviteEntity, query QueryDSL) {
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
func PendingWorkspaceInviteMockEntity() *PendingWorkspaceInviteEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &PendingWorkspaceInviteEntity{
      Value : &stringHolder,
      Type : &stringHolder,
      CoverLetter : &stringHolder,
      WorkspaceName : &stringHolder,
	}
	return entity
}
func PendingWorkspaceInviteActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := PendingWorkspaceInviteMockEntity()
		_, err := PendingWorkspaceInviteActionCreate(entity, query)
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
  func PendingWorkspaceInviteActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*PendingWorkspaceInviteEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &PendingWorkspaceInviteEntity{
          Value: &tildaRef,
          Type: &tildaRef,
          CoverLetter: &tildaRef,
          WorkspaceName: &tildaRef,
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
  func PendingWorkspaceInviteAssociationCreate(dto *PendingWorkspaceInviteEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func PendingWorkspaceInviteRelationContentCreate(dto *PendingWorkspaceInviteEntity, query QueryDSL) error {
return nil
}
func PendingWorkspaceInviteRelationContentUpdate(dto *PendingWorkspaceInviteEntity, query QueryDSL) error {
	return nil
}
func PendingWorkspaceInvitePolyglotCreateHandler(dto *PendingWorkspaceInviteEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func PendingWorkspaceInviteValidator(dto *PendingWorkspaceInviteEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func PendingWorkspaceInviteEntityPreSanitize(dto *PendingWorkspaceInviteEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func PendingWorkspaceInviteEntityBeforeCreateAppend(dto *PendingWorkspaceInviteEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    PendingWorkspaceInviteRecursiveAddUniqueId(dto, query)
  }
  func PendingWorkspaceInviteRecursiveAddUniqueId(dto *PendingWorkspaceInviteEntity, query QueryDSL) {
  }
func PendingWorkspaceInviteActionBatchCreateFn(dtos []*PendingWorkspaceInviteEntity, query QueryDSL) ([]*PendingWorkspaceInviteEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*PendingWorkspaceInviteEntity{}
		for _, item := range dtos {
			s, err := PendingWorkspaceInviteActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func PendingWorkspaceInviteDeleteEntireChildren(query QueryDSL, dto *PendingWorkspaceInviteEntity) (*IError) {
  return nil
}
func PendingWorkspaceInviteActionCreateFn(dto *PendingWorkspaceInviteEntity, query QueryDSL) (*PendingWorkspaceInviteEntity, *IError) {
	// 1. Validate always
	if iError := PendingWorkspaceInviteValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	PendingWorkspaceInviteEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	PendingWorkspaceInviteEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	PendingWorkspaceInvitePolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	PendingWorkspaceInviteRelationContentCreate(dto, query)
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
	PendingWorkspaceInviteAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(PENDING_WORKSPACE_INVITE_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&PendingWorkspaceInviteEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func PendingWorkspaceInviteActionGetOne(query QueryDSL) (*PendingWorkspaceInviteEntity, *IError) {
    refl := reflect.ValueOf(&PendingWorkspaceInviteEntity{})
    item, err := GetOneEntity[PendingWorkspaceInviteEntity](query, refl)
    entityPendingWorkspaceInviteFormatter(item, query)
    return item, err
  }
  func PendingWorkspaceInviteActionQuery(query QueryDSL) ([]*PendingWorkspaceInviteEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&PendingWorkspaceInviteEntity{})
    items, meta, err := QueryEntitiesPointer[PendingWorkspaceInviteEntity](query, refl)
    for _, item := range items {
      entityPendingWorkspaceInviteFormatter(item, query)
    }
    return items, meta, err
  }
  func PendingWorkspaceInviteUpdateExec(dbref *gorm.DB, query QueryDSL, fields *PendingWorkspaceInviteEntity) (*PendingWorkspaceInviteEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = PENDING_WORKSPACE_INVITE_EVENT_UPDATED
    PendingWorkspaceInviteEntityPreSanitize(fields, query)
    var item PendingWorkspaceInviteEntity
    q := dbref.
      Where(&PendingWorkspaceInviteEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    PendingWorkspaceInviteRelationContentUpdate(fields, query)
    PendingWorkspaceInvitePolyglotCreateHandler(fields, query)
    if ero := PendingWorkspaceInviteDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&PendingWorkspaceInviteEntity{UniqueId: uniqueId}).
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
  func PendingWorkspaceInviteActionUpdateFn(query QueryDSL, fields *PendingWorkspaceInviteEntity) (*PendingWorkspaceInviteEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := PendingWorkspaceInviteValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // PendingWorkspaceInviteRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      var item *PendingWorkspaceInviteEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *IError
        item, err = PendingWorkspaceInviteUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, CastToIError(vf)
    } else {
      dbref = query.Tx
      return PendingWorkspaceInviteUpdateExec(dbref, query, fields)
    }
  }
var PendingWorkspaceInviteWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire pendingworkspaceinvites ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
      ActionRequires: []string{PERM_ROOT_PENDING_WORKSPACE_INVITE_DELETE},
    })
		count, _ := PendingWorkspaceInviteActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func PendingWorkspaceInviteActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&PendingWorkspaceInviteEntity{})
	query.ActionRequires = []string{PERM_ROOT_PENDING_WORKSPACE_INVITE_DELETE}
	return RemoveEntity[PendingWorkspaceInviteEntity](query, refl)
}
func PendingWorkspaceInviteActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[PendingWorkspaceInviteEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'PendingWorkspaceInviteEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func PendingWorkspaceInviteActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[PendingWorkspaceInviteEntity]) (
    *BulkRecordRequest[PendingWorkspaceInviteEntity], *IError,
  ) {
    result := []*PendingWorkspaceInviteEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := PendingWorkspaceInviteActionUpdate(query, record)
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
func (x *PendingWorkspaceInviteEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var PendingWorkspaceInviteEntityMeta = TableMetaData{
	EntityName:    "PendingWorkspaceInvite",
	ExportKey:    "pending-workspace-invites",
	TableNameInDb: "fb_pending-workspace-invite_entities",
	EntityObject:  &PendingWorkspaceInviteEntity{},
	ExportStream: PendingWorkspaceInviteActionExportT,
	ImportQuery: PendingWorkspaceInviteActionImport,
}
func PendingWorkspaceInviteActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[PendingWorkspaceInviteEntity](query, PendingWorkspaceInviteActionQuery, PendingWorkspaceInvitePreloadRelations)
}
func PendingWorkspaceInviteActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[PendingWorkspaceInviteEntity](query, PendingWorkspaceInviteActionQuery, PendingWorkspaceInvitePreloadRelations)
}
func PendingWorkspaceInviteActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content PendingWorkspaceInviteEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := PendingWorkspaceInviteActionCreate(&content, query)
	return err
}
var PendingWorkspaceInviteCommonCliFlags = []cli.Flag{
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
      Name:     "value",
      Required: false,
      Usage:    "value",
    },
    &cli.StringFlag{
      Name:     "type",
      Required: false,
      Usage:    "type",
    },
    &cli.StringFlag{
      Name:     "cover-letter",
      Required: false,
      Usage:    "coverLetter",
    },
    &cli.StringFlag{
      Name:     "workspace-name",
      Required: false,
      Usage:    "workspaceName",
    },
    &cli.StringFlag{
      Name:     "role-id",
      Required: false,
      Usage:    "role",
    },
}
var PendingWorkspaceInviteCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:     "value",
		StructField:     "Value",
		Required: false,
		Usage:    "value",
		Type: "string",
	},
	{
		Name:     "type",
		StructField:     "Type",
		Required: false,
		Usage:    "type",
		Type: "string",
	},
	{
		Name:     "coverLetter",
		StructField:     "CoverLetter",
		Required: false,
		Usage:    "coverLetter",
		Type: "string",
	},
	{
		Name:     "workspaceName",
		StructField:     "WorkspaceName",
		Required: false,
		Usage:    "workspaceName",
		Type: "string",
	},
}
var PendingWorkspaceInviteCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "value",
      Required: false,
      Usage:    "value",
    },
    &cli.StringFlag{
      Name:     "type",
      Required: false,
      Usage:    "type",
    },
    &cli.StringFlag{
      Name:     "cover-letter",
      Required: false,
      Usage:    "coverLetter",
    },
    &cli.StringFlag{
      Name:     "workspace-name",
      Required: false,
      Usage:    "workspaceName",
    },
    &cli.StringFlag{
      Name:     "role-id",
      Required: false,
      Usage:    "role",
    },
}
  var PendingWorkspaceInviteCreateCmd cli.Command = PENDING_WORKSPACE_INVITE_ACTION_POST_ONE.ToCli()
  var PendingWorkspaceInviteCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []string{PERM_ROOT_PENDING_WORKSPACE_INVITE_CREATE},
      })
      entity := &PendingWorkspaceInviteEntity{}
      for _, item := range PendingWorkspaceInviteCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := PendingWorkspaceInviteActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var PendingWorkspaceInviteUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: PendingWorkspaceInviteCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []string{PERM_ROOT_PENDING_WORKSPACE_INVITE_UPDATE},
      })
      entity := CastPendingWorkspaceInviteFromCli(c)
      if entity, err := PendingWorkspaceInviteActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* PendingWorkspaceInviteEntity) FromCli(c *cli.Context) *PendingWorkspaceInviteEntity {
	return CastPendingWorkspaceInviteFromCli(c)
}
func CastPendingWorkspaceInviteFromCli (c *cli.Context) *PendingWorkspaceInviteEntity {
	template := &PendingWorkspaceInviteEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("value") {
        value := c.String("value")
        template.Value = &value
      }
      if c.IsSet("type") {
        value := c.String("type")
        template.Type = &value
      }
      if c.IsSet("cover-letter") {
        value := c.String("cover-letter")
        template.CoverLetter = &value
      }
      if c.IsSet("workspace-name") {
        value := c.String("workspace-name")
        template.WorkspaceName = &value
      }
      if c.IsSet("role-id") {
        value := c.String("role-id")
        template.RoleId = &value
      }
	return template
}
  func PendingWorkspaceInviteSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      PendingWorkspaceInviteActionCreate,
      reflect.ValueOf(&PendingWorkspaceInviteEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func PendingWorkspaceInviteWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := PendingWorkspaceInviteActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "PendingWorkspaceInvite", result)
    }
  }
var PendingWorkspaceInviteImportExportCommands = []cli.Command{
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
        ActionRequires: []string{PERM_ROOT_PENDING_WORKSPACE_INVITE_CREATE},
      })
			PendingWorkspaceInviteActionSeeder(query, c.Int("count"))
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
				Value: "pending-workspace-invite-seeder.yml",
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
        ActionRequires: []string{PERM_ROOT_PENDING_WORKSPACE_INVITE_CREATE},
      })
			PendingWorkspaceInviteActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "pending-workspace-invite-seeder-pending-workspace-invite.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of pending-workspace-invites, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]PendingWorkspaceInviteEntity{}
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
			PendingWorkspaceInviteCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				PendingWorkspaceInviteActionCreate,
				reflect.ValueOf(&PendingWorkspaceInviteEntity{}).Elem(),
				c.String("file"),
        &SecurityModel{
					ActionRequires: []string{PERM_ROOT_PENDING_WORKSPACE_INVITE_CREATE},
				},
        func() PendingWorkspaceInviteEntity {
					v := CastPendingWorkspaceInviteFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var PendingWorkspaceInviteCliCommands []cli.Command = []cli.Command{
      GetCommonQuery2(PendingWorkspaceInviteActionQuery, &SecurityModel{
        ActionRequires: []string{PERM_ROOT_PENDING_WORKSPACE_INVITE_CREATE},
      }),
      GetCommonTableQuery(reflect.ValueOf(&PendingWorkspaceInviteEntity{}).Elem(), PendingWorkspaceInviteActionQuery),
          PendingWorkspaceInviteCreateCmd,
          PendingWorkspaceInviteUpdateCmd,
          PendingWorkspaceInviteCreateInteractiveCmd,
          PendingWorkspaceInviteWipeCmd,
          GetCommonRemoveQuery(reflect.ValueOf(&PendingWorkspaceInviteEntity{}).Elem(), PendingWorkspaceInviteActionRemove),
  }
  func PendingWorkspaceInviteCliFn() cli.Command {
    PendingWorkspaceInviteCliCommands = append(PendingWorkspaceInviteCliCommands, PendingWorkspaceInviteImportExportCommands...)
    return cli.Command{
      Name:        "pendingWorkspaceInvite",
      Description: "PendingWorkspaceInvites module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: PendingWorkspaceInviteCliCommands,
    }
  }
var PENDING_WORKSPACE_INVITE_ACTION_POST_ONE = Module2Action{
    ActionName:    "create",
    ActionAliases: []string{"c"},
    Description: "Create new pendingWorkspaceInvite",
    Flags: PendingWorkspaceInviteCommonCliFlags,
    Method: "POST",
    Url:    "/pending-workspace-invite",
    SecurityModel: &SecurityModel{
      ActionRequires: []string{PERM_ROOT_PENDING_WORKSPACE_INVITE_CREATE},
    },
    Handlers: []gin.HandlerFunc{
      func (c *gin.Context) {
        HttpPostEntity(c, PendingWorkspaceInviteActionCreate)
      },
    },
    CliAction: func(c *cli.Context, security *SecurityModel) error {
      result, err := CliPostEntity(c, PendingWorkspaceInviteActionCreate, security)
      HandleActionInCli(c, result, err, map[string]map[string]string{})
      return err
    },
    Action: PendingWorkspaceInviteActionCreate,
    Format: "POST_ONE",
    RequestEntity: &PendingWorkspaceInviteEntity{},
    ResponseEntity: &PendingWorkspaceInviteEntity{},
  }
  /**
  *	Override this function on PendingWorkspaceInviteEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendPendingWorkspaceInviteRouter = func(r *[]Module2Action) {}
  func GetPendingWorkspaceInviteModule2Actions() []Module2Action {
    routes := []Module2Action{
       {
        Method: "GET",
        Url:    "/pending-workspace-invites",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_PENDING_WORKSPACE_INVITE_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpQueryEntity(c, PendingWorkspaceInviteActionQuery)
          },
        },
        Format: "QUERY",
        Action: PendingWorkspaceInviteActionQuery,
        ResponseEntity: &[]PendingWorkspaceInviteEntity{},
      },
      {
        Method: "GET",
        Url:    "/pending-workspace-invites/export",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_PENDING_WORKSPACE_INVITE_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpStreamFileChannel(c, PendingWorkspaceInviteActionExport)
          },
        },
        Format: "QUERY",
        Action: PendingWorkspaceInviteActionExport,
        ResponseEntity: &[]PendingWorkspaceInviteEntity{},
      },
      {
        Method: "GET",
        Url:    "/pending-workspace-invite/:uniqueId",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_PENDING_WORKSPACE_INVITE_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpGetEntity(c, PendingWorkspaceInviteActionGetOne)
          },
        },
        Format: "GET_ONE",
        Action: PendingWorkspaceInviteActionGetOne,
        ResponseEntity: &PendingWorkspaceInviteEntity{},
      },
      PENDING_WORKSPACE_INVITE_ACTION_POST_ONE,
      {
        ActionName:    "update",
        ActionAliases: []string{"u"},
        Flags: PendingWorkspaceInviteCommonCliFlagsOptional,
        Method: "PATCH",
        Url:    "/pending-workspace-invite",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_PENDING_WORKSPACE_INVITE_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpUpdateEntity(c, PendingWorkspaceInviteActionUpdate)
          },
        },
        Action: PendingWorkspaceInviteActionUpdate,
        RequestEntity: &PendingWorkspaceInviteEntity{},
        Format: "PATCH_ONE",
        ResponseEntity: &PendingWorkspaceInviteEntity{},
      },
      {
        Method: "PATCH",
        Url:    "/pending-workspace-invites",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_PENDING_WORKSPACE_INVITE_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpUpdateEntities(c, PendingWorkspaceInviteActionBulkUpdate)
          },
        },
        Action: PendingWorkspaceInviteActionBulkUpdate,
        Format: "PATCH_BULK",
        RequestEntity:  &BulkRecordRequest[PendingWorkspaceInviteEntity]{},
        ResponseEntity: &BulkRecordRequest[PendingWorkspaceInviteEntity]{},
      },
      {
        Method: "DELETE",
        Url:    "/pending-workspace-invite",
        Format: "DELETE_DSL",
        SecurityModel: &SecurityModel{
          ActionRequires: []string{PERM_ROOT_PENDING_WORKSPACE_INVITE_DELETE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpRemoveEntity(c, PendingWorkspaceInviteActionRemove)
          },
        },
        Action: PendingWorkspaceInviteActionRemove,
        RequestEntity: &DeleteRequest{},
        ResponseEntity: &DeleteResponse{},
        TargetEntity: &PendingWorkspaceInviteEntity{},
      },
    }
    // Append user defined functions
    AppendPendingWorkspaceInviteRouter(&routes)
    return routes
  }
  func CreatePendingWorkspaceInviteRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetPendingWorkspaceInviteModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, PendingWorkspaceInviteEntityJsonSchema, "pending-workspace-invite-http", "workspaces")
    WriteEntitySchema("PendingWorkspaceInviteEntity", PendingWorkspaceInviteEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_PENDING_WORKSPACE_INVITE_DELETE = "root/workspaces/pending-workspace-invite/delete"
var PERM_ROOT_PENDING_WORKSPACE_INVITE_CREATE = "root/workspaces/pending-workspace-invite/create"
var PERM_ROOT_PENDING_WORKSPACE_INVITE_UPDATE = "root/workspaces/pending-workspace-invite/update"
var PERM_ROOT_PENDING_WORKSPACE_INVITE_QUERY = "root/workspaces/pending-workspace-invite/query"
var PERM_ROOT_PENDING_WORKSPACE_INVITE = "root/workspaces/pending-workspace-invite/*"
var ALL_PENDING_WORKSPACE_INVITE_PERMISSIONS = []string{
	PERM_ROOT_PENDING_WORKSPACE_INVITE_DELETE,
	PERM_ROOT_PENDING_WORKSPACE_INVITE_CREATE,
	PERM_ROOT_PENDING_WORKSPACE_INVITE_UPDATE,
	PERM_ROOT_PENDING_WORKSPACE_INVITE_QUERY,
	PERM_ROOT_PENDING_WORKSPACE_INVITE,
}