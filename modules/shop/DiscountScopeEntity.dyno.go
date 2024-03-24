package shop
import (
    "github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/workspaces"
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
	seeders "github.com/torabian/fireback/modules/shop/seeders/DiscountScope"
	metas "github.com/torabian/fireback/modules/shop/metas"
)
type DiscountScopeEntity struct {
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
    Name   *string `json:"name" yaml:"name"        translate:"true" `
    // Datenano also has a text representation
    Description   *string `json:"description" yaml:"description"        translate:"true" `
    // Datenano also has a text representation
    Translations     []*DiscountScopeEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    Children []*DiscountScopeEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *DiscountScopeEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var DiscountScopePreloadRelations []string = []string{}
var DISCOUNT_SCOPE_EVENT_CREATED = "discountScope.created"
var DISCOUNT_SCOPE_EVENT_UPDATED = "discountScope.updated"
var DISCOUNT_SCOPE_EVENT_DELETED = "discountScope.deleted"
var DISCOUNT_SCOPE_EVENTS = []string{
	DISCOUNT_SCOPE_EVENT_CREATED,
	DISCOUNT_SCOPE_EVENT_UPDATED,
	DISCOUNT_SCOPE_EVENT_DELETED,
}
type DiscountScopeFieldMap struct {
		Name workspaces.TranslatedString `yaml:"name"`
		Description workspaces.TranslatedString `yaml:"description"`
}
var DiscountScopeEntityMetaConfig map[string]int64 = map[string]int64{
}
var DiscountScopeEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&DiscountScopeEntity{}))
  type DiscountScopeEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        Name string `yaml:"name" json:"name"`
        Description string `yaml:"description" json:"description"`
  }
func entityDiscountScopeFormatter(dto *DiscountScopeEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
	if dto.Created > 0 {
		dto.CreatedFormatted = workspaces.FormatDateBasedOnQuery(dto.Created, query)
	}
	if dto.Updated > 0 {
		dto.CreatedFormatted = workspaces.FormatDateBasedOnQuery(dto.Updated, query)
	}
}
func DiscountScopeMockEntity() *DiscountScopeEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &DiscountScopeEntity{
      Name : &stringHolder,
      Description : &stringHolder,
	}
	return entity
}
func DiscountScopeActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := DiscountScopeMockEntity()
		_, err := DiscountScopeActionCreate(entity, query)
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
    func (x*DiscountScopeEntity) GetNameTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Name
          }
        }
      }
      return ""
    }
    func (x*DiscountScopeEntity) GetDescriptionTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Description
          }
        }
      }
      return ""
    }
  func DiscountScopeActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*DiscountScopeEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &DiscountScopeEntity{
          Name: &tildaRef,
          Description: &tildaRef,
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
  func DiscountScopeAssociationCreate(dto *DiscountScopeEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func DiscountScopeRelationContentCreate(dto *DiscountScopeEntity, query workspaces.QueryDSL) error {
return nil
}
func DiscountScopeRelationContentUpdate(dto *DiscountScopeEntity, query workspaces.QueryDSL) error {
	return nil
}
func DiscountScopePolyglotCreateHandler(dto *DiscountScopeEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &DiscountScopeEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func DiscountScopeValidator(dto *DiscountScopeEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func DiscountScopeEntityPreSanitize(dto *DiscountScopeEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func DiscountScopeEntityBeforeCreateAppend(dto *DiscountScopeEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    DiscountScopeRecursiveAddUniqueId(dto, query)
  }
  func DiscountScopeRecursiveAddUniqueId(dto *DiscountScopeEntity, query workspaces.QueryDSL) {
  }
func DiscountScopeActionBatchCreateFn(dtos []*DiscountScopeEntity, query workspaces.QueryDSL) ([]*DiscountScopeEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*DiscountScopeEntity{}
		for _, item := range dtos {
			s, err := DiscountScopeActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func DiscountScopeDeleteEntireChildren(query workspaces.QueryDSL, dto *DiscountScopeEntity) (*workspaces.IError) {
  return nil
}
func DiscountScopeActionCreateFn(dto *DiscountScopeEntity, query workspaces.QueryDSL) (*DiscountScopeEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := DiscountScopeValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	DiscountScopeEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	DiscountScopeEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	DiscountScopePolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	DiscountScopeRelationContentCreate(dto, query)
	// 4. Create the entity
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()
	} else {
		dbref = query.Tx
	}
	query.Tx = dbref;
	err := dbref.Create(&dto).Error
	if err != nil {
		err := workspaces.GormErrorToIError(err)
		return dto, err
	}
	// 5. Create sub entities, objects or arrays, association to other entities
	DiscountScopeAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(DISCOUNT_SCOPE_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&DiscountScopeEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func DiscountScopeActionGetOne(query workspaces.QueryDSL) (*DiscountScopeEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&DiscountScopeEntity{})
    item, err := workspaces.GetOneEntity[DiscountScopeEntity](query, refl)
    entityDiscountScopeFormatter(item, query)
    return item, err
  }
  func DiscountScopeActionQuery(query workspaces.QueryDSL) ([]*DiscountScopeEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&DiscountScopeEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[DiscountScopeEntity](query, refl)
    for _, item := range items {
      entityDiscountScopeFormatter(item, query)
    }
    return items, meta, err
  }
  func DiscountScopeUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *DiscountScopeEntity) (*DiscountScopeEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = DISCOUNT_SCOPE_EVENT_UPDATED
    DiscountScopeEntityPreSanitize(fields, query)
    var item DiscountScopeEntity
    q := dbref.
      Where(&DiscountScopeEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    DiscountScopeRelationContentUpdate(fields, query)
    DiscountScopePolyglotCreateHandler(fields, query)
    if ero := DiscountScopeDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&DiscountScopeEntity{UniqueId: uniqueId}).
      First(&item).Error
    event.MustFire(query.TriggerEventName, event.M{
      "entity":   &item,
      "target":   "workspace",
      "unqiueId": query.WorkspaceId,
    })
    if err != nil {
      return &item, workspaces.GormErrorToIError(err)
    }
    return &item, nil
  }
  func DiscountScopeActionUpdateFn(query workspaces.QueryDSL, fields *DiscountScopeEntity) (*DiscountScopeEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := DiscountScopeValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // DiscountScopeRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *DiscountScopeEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = DiscountScopeUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return DiscountScopeUpdateExec(dbref, query, fields)
    }
  }
var DiscountScopeWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire discountscopes ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []string{PERM_ROOT_DISCOUNT_SCOPE_DELETE},
    })
		count, _ := DiscountScopeActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func DiscountScopeActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&DiscountScopeEntity{})
	query.ActionRequires = []string{PERM_ROOT_DISCOUNT_SCOPE_DELETE}
	return workspaces.RemoveEntity[DiscountScopeEntity](query, refl)
}
func DiscountScopeActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[DiscountScopeEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'DiscountScopeEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func DiscountScopeActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[DiscountScopeEntity]) (
    *workspaces.BulkRecordRequest[DiscountScopeEntity], *workspaces.IError,
  ) {
    result := []*DiscountScopeEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := DiscountScopeActionUpdate(query, record)
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
    return nil, err.(*workspaces.IError)
  }
func (x *DiscountScopeEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var DiscountScopeEntityMeta = workspaces.TableMetaData{
	EntityName:    "DiscountScope",
	ExportKey:    "discount-scopes",
	TableNameInDb: "fb_discount-scope_entities",
	EntityObject:  &DiscountScopeEntity{},
	ExportStream: DiscountScopeActionExportT,
	ImportQuery: DiscountScopeActionImport,
}
func DiscountScopeActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[DiscountScopeEntity](query, DiscountScopeActionQuery, DiscountScopePreloadRelations)
}
func DiscountScopeActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[DiscountScopeEntity](query, DiscountScopeActionQuery, DiscountScopePreloadRelations)
}
func DiscountScopeActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content DiscountScopeEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := DiscountScopeActionCreate(&content, query)
	return err
}
var DiscountScopeCommonCliFlags = []cli.Flag{
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
      Name:     "name",
      Required: false,
      Usage:    "name",
    },
    &cli.StringFlag{
      Name:     "description",
      Required: false,
      Usage:    "description",
    },
}
var DiscountScopeCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "name",
		StructField:     "Name",
		Required: false,
		Usage:    "name",
		Type: "string",
	},
	{
		Name:     "description",
		StructField:     "Description",
		Required: false,
		Usage:    "description",
		Type: "string",
	},
}
var DiscountScopeCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "name",
      Required: false,
      Usage:    "name",
    },
    &cli.StringFlag{
      Name:     "description",
      Required: false,
      Usage:    "description",
    },
}
  var DiscountScopeCreateCmd cli.Command = DISCOUNT_SCOPE_ACTION_POST_ONE.ToCli()
  var DiscountScopeCreateInteractiveCmd cli.Command = cli.Command{
    Name:  "ic",
    Usage: "Creates a new template, using requied fields in an interactive name",
    Flags: []cli.Flag{
      &cli.BoolFlag{
        Name:  "all",
        Usage: "Interactively asks for all inputs, not only required ones",
      },
    },
    Action: func(c *cli.Context) {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []string{PERM_ROOT_DISCOUNT_SCOPE_CREATE},
      })
      entity := &DiscountScopeEntity{}
      for _, item := range DiscountScopeCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := DiscountScopeActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var DiscountScopeUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: DiscountScopeCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []string{PERM_ROOT_DISCOUNT_SCOPE_UPDATE},
      })
      entity := CastDiscountScopeFromCli(c)
      if entity, err := DiscountScopeActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* DiscountScopeEntity) FromCli(c *cli.Context) *DiscountScopeEntity {
	return CastDiscountScopeFromCli(c)
}
func CastDiscountScopeFromCli (c *cli.Context) *DiscountScopeEntity {
	template := &DiscountScopeEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("name") {
        value := c.String("name")
        template.Name = &value
      }
      if c.IsSet("description") {
        value := c.String("description")
        template.Description = &value
      }
	return template
}
  func DiscountScopeSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      DiscountScopeActionCreate,
      reflect.ValueOf(&DiscountScopeEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func DiscountScopeSyncSeeders() {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{WorkspaceId: workspaces.USER_SYSTEM},
      DiscountScopeActionCreate,
      reflect.ValueOf(&DiscountScopeEntity{}).Elem(),
      &seeders.ViewsFs,
      []string{},
      true,
    )
  }
  func DiscountScopeWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := DiscountScopeActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "DiscountScope", result)
    }
  }
var DiscountScopeImportExportCommands = []cli.Command{
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
			query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []string{PERM_ROOT_DISCOUNT_SCOPE_CREATE},
      })
			DiscountScopeActionSeeder(query, c.Int("count"))
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
				Value: "discount-scope-seeder.yml",
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
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []string{PERM_ROOT_DISCOUNT_SCOPE_CREATE},
      })
			DiscountScopeActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "discount-scope-seeder-discount-scope.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of discount-scopes, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]DiscountScopeEntity{}
			workspaces.ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
	cli.Command{
		Name:  "list",
		Usage: "Prints the list of files attached to this module for syncing or bootstrapping project",
		Action: func(c *cli.Context) error {
			if entity, err := workspaces.GetSeederFilenames(&seeders.ViewsFs, ""); err != nil {
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
			workspaces.CommonCliImportEmbedCmd(c,
				DiscountScopeActionCreate,
				reflect.ValueOf(&DiscountScopeEntity{}).Elem(),
				&seeders.ViewsFs,
			)
			return nil
		},
	},
	cli.Command{
		Name:    "export",
		Aliases: []string{"e"},
		Flags: append(workspaces.CommonQueryFlags,
			&cli.StringFlag{
				Name:     "file",
				Usage:    "The address of file you want the csv/yaml/json be exported to",
				Required: true,
			}),
		Usage: "Exports a query results into the csv/yaml/json format",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliExportCmd(c,
				DiscountScopeActionQuery,
				reflect.ValueOf(&DiscountScopeEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"DiscountScopeFieldMap.yml",
				DiscountScopePreloadRelations,
			)
			return nil
		},
	},
	cli.Command{
		Name:    "import",
    Flags: append(
			append(
				workspaces.CommonQueryFlags,
				&cli.StringFlag{
					Name:     "file",
					Usage:    "The address of file you want the csv be imported from",
					Required: true,
				}),
			DiscountScopeCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				DiscountScopeActionCreate,
				reflect.ValueOf(&DiscountScopeEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []string{PERM_ROOT_DISCOUNT_SCOPE_CREATE},
				},
        func() DiscountScopeEntity {
					v := CastDiscountScopeFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var DiscountScopeCliCommands []cli.Command = []cli.Command{
      workspaces.GetCommonQuery2(DiscountScopeActionQuery, &workspaces.SecurityModel{
        ActionRequires: []string{PERM_ROOT_DISCOUNT_SCOPE_CREATE},
      }),
      workspaces.GetCommonTableQuery(reflect.ValueOf(&DiscountScopeEntity{}).Elem(), DiscountScopeActionQuery),
          DiscountScopeCreateCmd,
          DiscountScopeUpdateCmd,
          DiscountScopeCreateInteractiveCmd,
          DiscountScopeWipeCmd,
          workspaces.GetCommonRemoveQuery(reflect.ValueOf(&DiscountScopeEntity{}).Elem(), DiscountScopeActionRemove),
  }
  func DiscountScopeCliFn() cli.Command {
    DiscountScopeCliCommands = append(DiscountScopeCliCommands, DiscountScopeImportExportCommands...)
    return cli.Command{
      Name:        "discountScope",
      Description: "DiscountScopes module actions (sample module to handle complex entities)",
      Usage:       "Determine if the discount applies to the entire basket (total order) or per item, etc",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: DiscountScopeCliCommands,
    }
  }
var DISCOUNT_SCOPE_ACTION_POST_ONE = workspaces.Module2Action{
    ActionName:    "create",
    ActionAliases: []string{"c"},
    Description: "Create new discountScope",
    Flags: DiscountScopeCommonCliFlags,
    Method: "POST",
    Url:    "/discount-scope",
    SecurityModel: &workspaces.SecurityModel{
      ActionRequires: []string{PERM_ROOT_DISCOUNT_SCOPE_CREATE},
    },
    Handlers: []gin.HandlerFunc{
      func (c *gin.Context) {
        workspaces.HttpPostEntity(c, DiscountScopeActionCreate)
      },
    },
    CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
      result, err := workspaces.CliPostEntity(c, DiscountScopeActionCreate, security)
      workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
      return err
    },
    Action: DiscountScopeActionCreate,
    Format: "POST_ONE",
    RequestEntity: &DiscountScopeEntity{},
    ResponseEntity: &DiscountScopeEntity{},
  }
  /**
  *	Override this function on DiscountScopeEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendDiscountScopeRouter = func(r *[]workspaces.Module2Action) {}
  func GetDiscountScopeModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
       {
        Method: "GET",
        Url:    "/discount-scopes",
        SecurityModel: &workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_DISCOUNT_SCOPE_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpQueryEntity(c, DiscountScopeActionQuery)
          },
        },
        Format: "QUERY",
        Action: DiscountScopeActionQuery,
        ResponseEntity: &[]DiscountScopeEntity{},
      },
      {
        Method: "GET",
        Url:    "/discount-scopes/export",
        SecurityModel: &workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_DISCOUNT_SCOPE_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpStreamFileChannel(c, DiscountScopeActionExport)
          },
        },
        Format: "QUERY",
        Action: DiscountScopeActionExport,
        ResponseEntity: &[]DiscountScopeEntity{},
      },
      {
        Method: "GET",
        Url:    "/discount-scope/:uniqueId",
        SecurityModel: &workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_DISCOUNT_SCOPE_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpGetEntity(c, DiscountScopeActionGetOne)
          },
        },
        Format: "GET_ONE",
        Action: DiscountScopeActionGetOne,
        ResponseEntity: &DiscountScopeEntity{},
      },
      DISCOUNT_SCOPE_ACTION_POST_ONE,
      {
        ActionName:    "update",
        ActionAliases: []string{"u"},
        Flags: DiscountScopeCommonCliFlagsOptional,
        Method: "PATCH",
        Url:    "/discount-scope",
        SecurityModel: &workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_DISCOUNT_SCOPE_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpUpdateEntity(c, DiscountScopeActionUpdate)
          },
        },
        Action: DiscountScopeActionUpdate,
        RequestEntity: &DiscountScopeEntity{},
        Format: "PATCH_ONE",
        ResponseEntity: &DiscountScopeEntity{},
      },
      {
        Method: "PATCH",
        Url:    "/discount-scopes",
        SecurityModel: &workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_DISCOUNT_SCOPE_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpUpdateEntities(c, DiscountScopeActionBulkUpdate)
          },
        },
        Action: DiscountScopeActionBulkUpdate,
        Format: "PATCH_BULK",
        RequestEntity:  &workspaces.BulkRecordRequest[DiscountScopeEntity]{},
        ResponseEntity: &workspaces.BulkRecordRequest[DiscountScopeEntity]{},
      },
      {
        Method: "DELETE",
        Url:    "/discount-scope",
        Format: "DELETE_DSL",
        SecurityModel: &workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_DISCOUNT_SCOPE_DELETE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpRemoveEntity(c, DiscountScopeActionRemove)
          },
        },
        Action: DiscountScopeActionRemove,
        RequestEntity: &workspaces.DeleteRequest{},
        ResponseEntity: &workspaces.DeleteResponse{},
        TargetEntity: &DiscountScopeEntity{},
      },
    }
    // Append user defined functions
    AppendDiscountScopeRouter(&routes)
    return routes
  }
  func CreateDiscountScopeRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetDiscountScopeModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, DiscountScopeEntityJsonSchema, "discount-scope-http", "shop")
    workspaces.WriteEntitySchema("DiscountScopeEntity", DiscountScopeEntityJsonSchema, "shop")
    return httpRoutes
  }
var PERM_ROOT_DISCOUNT_SCOPE_DELETE = "root/shop/discount-scope/delete"
var PERM_ROOT_DISCOUNT_SCOPE_CREATE = "root/shop/discount-scope/create"
var PERM_ROOT_DISCOUNT_SCOPE_UPDATE = "root/shop/discount-scope/update"
var PERM_ROOT_DISCOUNT_SCOPE_QUERY = "root/shop/discount-scope/query"
var PERM_ROOT_DISCOUNT_SCOPE = "root/shop/discount-scope/*"
var ALL_DISCOUNT_SCOPE_PERMISSIONS = []string{
	PERM_ROOT_DISCOUNT_SCOPE_DELETE,
	PERM_ROOT_DISCOUNT_SCOPE_CREATE,
	PERM_ROOT_DISCOUNT_SCOPE_UPDATE,
	PERM_ROOT_DISCOUNT_SCOPE_QUERY,
	PERM_ROOT_DISCOUNT_SCOPE,
}