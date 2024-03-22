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
	seeders "github.com/torabian/fireback/modules/shop/seeders/DiscountType"
	metas "github.com/torabian/fireback/modules/shop/metas"
)
type DiscountTypeEntity struct {
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
    Translations     []*DiscountTypeEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    Children []*DiscountTypeEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *DiscountTypeEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var DiscountTypePreloadRelations []string = []string{}
var DISCOUNTTYPE_EVENT_CREATED = "discountType.created"
var DISCOUNTTYPE_EVENT_UPDATED = "discountType.updated"
var DISCOUNTTYPE_EVENT_DELETED = "discountType.deleted"
var DISCOUNTTYPE_EVENTS = []string{
	DISCOUNTTYPE_EVENT_CREATED,
	DISCOUNTTYPE_EVENT_UPDATED,
	DISCOUNTTYPE_EVENT_DELETED,
}
type DiscountTypeFieldMap struct {
		Name workspaces.TranslatedString `yaml:"name"`
		Description workspaces.TranslatedString `yaml:"description"`
}
var DiscountTypeEntityMetaConfig map[string]int64 = map[string]int64{
}
var DiscountTypeEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&DiscountTypeEntity{}))
  type DiscountTypeEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        Name string `yaml:"name" json:"name"`
        Description string `yaml:"description" json:"description"`
  }
func entityDiscountTypeFormatter(dto *DiscountTypeEntity, query workspaces.QueryDSL) {
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
func DiscountTypeMockEntity() *DiscountTypeEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &DiscountTypeEntity{
      Name : &stringHolder,
      Description : &stringHolder,
	}
	return entity
}
func DiscountTypeActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := DiscountTypeMockEntity()
		_, err := DiscountTypeActionCreate(entity, query)
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
    func (x*DiscountTypeEntity) GetNameTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Name
          }
        }
      }
      return ""
    }
    func (x*DiscountTypeEntity) GetDescriptionTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Description
          }
        }
      }
      return ""
    }
  func DiscountTypeActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*DiscountTypeEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &DiscountTypeEntity{
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
  func DiscountTypeAssociationCreate(dto *DiscountTypeEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func DiscountTypeRelationContentCreate(dto *DiscountTypeEntity, query workspaces.QueryDSL) error {
return nil
}
func DiscountTypeRelationContentUpdate(dto *DiscountTypeEntity, query workspaces.QueryDSL) error {
	return nil
}
func DiscountTypePolyglotCreateHandler(dto *DiscountTypeEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &DiscountTypeEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func DiscountTypeValidator(dto *DiscountTypeEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func DiscountTypeEntityPreSanitize(dto *DiscountTypeEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func DiscountTypeEntityBeforeCreateAppend(dto *DiscountTypeEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    DiscountTypeRecursiveAddUniqueId(dto, query)
  }
  func DiscountTypeRecursiveAddUniqueId(dto *DiscountTypeEntity, query workspaces.QueryDSL) {
  }
func DiscountTypeActionBatchCreateFn(dtos []*DiscountTypeEntity, query workspaces.QueryDSL) ([]*DiscountTypeEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*DiscountTypeEntity{}
		for _, item := range dtos {
			s, err := DiscountTypeActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func DiscountTypeDeleteEntireChildren(query workspaces.QueryDSL, dto *DiscountTypeEntity) (*workspaces.IError) {
  return nil
}
func DiscountTypeActionCreateFn(dto *DiscountTypeEntity, query workspaces.QueryDSL) (*DiscountTypeEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := DiscountTypeValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	DiscountTypeEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	DiscountTypeEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	DiscountTypePolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	DiscountTypeRelationContentCreate(dto, query)
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
	DiscountTypeAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(DISCOUNTTYPE_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&DiscountTypeEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func DiscountTypeActionGetOne(query workspaces.QueryDSL) (*DiscountTypeEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&DiscountTypeEntity{})
    item, err := workspaces.GetOneEntity[DiscountTypeEntity](query, refl)
    entityDiscountTypeFormatter(item, query)
    return item, err
  }
  func DiscountTypeActionQuery(query workspaces.QueryDSL) ([]*DiscountTypeEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&DiscountTypeEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[DiscountTypeEntity](query, refl)
    for _, item := range items {
      entityDiscountTypeFormatter(item, query)
    }
    return items, meta, err
  }
  func DiscountTypeUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *DiscountTypeEntity) (*DiscountTypeEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = DISCOUNTTYPE_EVENT_UPDATED
    DiscountTypeEntityPreSanitize(fields, query)
    var item DiscountTypeEntity
    q := dbref.
      Where(&DiscountTypeEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    DiscountTypeRelationContentUpdate(fields, query)
    DiscountTypePolyglotCreateHandler(fields, query)
    if ero := DiscountTypeDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&DiscountTypeEntity{UniqueId: uniqueId}).
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
  func DiscountTypeActionUpdateFn(query workspaces.QueryDSL, fields *DiscountTypeEntity) (*DiscountTypeEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := DiscountTypeValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // DiscountTypeRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *DiscountTypeEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = DiscountTypeUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return DiscountTypeUpdateExec(dbref, query, fields)
    }
  }
var DiscountTypeWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire discounttypes ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []string{PERM_ROOT_DISCOUNTTYPE_DELETE},
    })
		count, _ := DiscountTypeActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func DiscountTypeActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&DiscountTypeEntity{})
	query.ActionRequires = []string{PERM_ROOT_DISCOUNTTYPE_DELETE}
	return workspaces.RemoveEntity[DiscountTypeEntity](query, refl)
}
func DiscountTypeActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[DiscountTypeEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'DiscountTypeEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func DiscountTypeActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[DiscountTypeEntity]) (
    *workspaces.BulkRecordRequest[DiscountTypeEntity], *workspaces.IError,
  ) {
    result := []*DiscountTypeEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := DiscountTypeActionUpdate(query, record)
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
func (x *DiscountTypeEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var DiscountTypeEntityMeta = workspaces.TableMetaData{
	EntityName:    "DiscountType",
	ExportKey:    "discount-types",
	TableNameInDb: "fb_discounttype_entities",
	EntityObject:  &DiscountTypeEntity{},
	ExportStream: DiscountTypeActionExportT,
	ImportQuery: DiscountTypeActionImport,
}
func DiscountTypeActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[DiscountTypeEntity](query, DiscountTypeActionQuery, DiscountTypePreloadRelations)
}
func DiscountTypeActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[DiscountTypeEntity](query, DiscountTypeActionQuery, DiscountTypePreloadRelations)
}
func DiscountTypeActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content DiscountTypeEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := DiscountTypeActionCreate(&content, query)
	return err
}
var DiscountTypeCommonCliFlags = []cli.Flag{
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
var DiscountTypeCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
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
var DiscountTypeCommonCliFlagsOptional = []cli.Flag{
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
  var DiscountTypeCreateCmd cli.Command = DISCOUNTTYPE_ACTION_POST_ONE.ToCli()
  var DiscountTypeCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []string{PERM_ROOT_DISCOUNTTYPE_CREATE},
      })
      entity := &DiscountTypeEntity{}
      for _, item := range DiscountTypeCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := DiscountTypeActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var DiscountTypeUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: DiscountTypeCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []string{PERM_ROOT_DISCOUNTTYPE_UPDATE},
      })
      entity := CastDiscountTypeFromCli(c)
      if entity, err := DiscountTypeActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* DiscountTypeEntity) FromCli(c *cli.Context) *DiscountTypeEntity {
	return CastDiscountTypeFromCli(c)
}
func CastDiscountTypeFromCli (c *cli.Context) *DiscountTypeEntity {
	template := &DiscountTypeEntity{}
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
  func DiscountTypeSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      DiscountTypeActionCreate,
      reflect.ValueOf(&DiscountTypeEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func DiscountTypeSyncSeeders() {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{WorkspaceId: workspaces.USER_SYSTEM},
      DiscountTypeActionCreate,
      reflect.ValueOf(&DiscountTypeEntity{}).Elem(),
      &seeders.ViewsFs,
      []string{},
      true,
    )
  }
  func DiscountTypeWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := DiscountTypeActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "DiscountType", result)
    }
  }
var DiscountTypeImportExportCommands = []cli.Command{
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
        ActionRequires: []string{PERM_ROOT_DISCOUNTTYPE_CREATE},
      })
			DiscountTypeActionSeeder(query, c.Int("count"))
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
				Value: "discount-type-seeder.yml",
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
        ActionRequires: []string{PERM_ROOT_DISCOUNTTYPE_CREATE},
      })
			DiscountTypeActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "discount-type-seeder-discount-type.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of discount-types, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]DiscountTypeEntity{}
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
				DiscountTypeActionCreate,
				reflect.ValueOf(&DiscountTypeEntity{}).Elem(),
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
				DiscountTypeActionQuery,
				reflect.ValueOf(&DiscountTypeEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"DiscountTypeFieldMap.yml",
				DiscountTypePreloadRelations,
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
			DiscountTypeCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				DiscountTypeActionCreate,
				reflect.ValueOf(&DiscountTypeEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []string{PERM_ROOT_DISCOUNTTYPE_CREATE},
				},
        func() DiscountTypeEntity {
					v := CastDiscountTypeFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var DiscountTypeCliCommands []cli.Command = []cli.Command{
      workspaces.GetCommonQuery2(DiscountTypeActionQuery, &workspaces.SecurityModel{
        ActionRequires: []string{PERM_ROOT_DISCOUNTTYPE_CREATE},
      }),
      workspaces.GetCommonTableQuery(reflect.ValueOf(&DiscountTypeEntity{}).Elem(), DiscountTypeActionQuery),
          DiscountTypeCreateCmd,
          DiscountTypeUpdateCmd,
          DiscountTypeCreateInteractiveCmd,
          DiscountTypeWipeCmd,
          workspaces.GetCommonRemoveQuery(reflect.ValueOf(&DiscountTypeEntity{}).Elem(), DiscountTypeActionRemove),
  }
  func DiscountTypeCliFn() cli.Command {
    DiscountTypeCliCommands = append(DiscountTypeCliCommands, DiscountTypeImportExportCommands...)
    return cli.Command{
      Name:        "discountType",
      Description: "DiscountTypes module actions (sample module to handle complex entities)",
      Usage:       "Types of the discounts",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: DiscountTypeCliCommands,
    }
  }
var DISCOUNTTYPE_ACTION_POST_ONE = workspaces.Module2Action{
    ActionName:    "create",
    ActionAliases: []string{"c"},
    Description: "Create new discountType",
    Flags: DiscountTypeCommonCliFlags,
    Method: "POST",
    Url:    "/discount-type",
    SecurityModel: &workspaces.SecurityModel{
      ActionRequires: []string{PERM_ROOT_DISCOUNTTYPE_CREATE},
    },
    Handlers: []gin.HandlerFunc{
      func (c *gin.Context) {
        workspaces.HttpPostEntity(c, DiscountTypeActionCreate)
      },
    },
    CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
      result, err := workspaces.CliPostEntity(c, DiscountTypeActionCreate, security)
      workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
      return err
    },
    Action: DiscountTypeActionCreate,
    Format: "POST_ONE",
    RequestEntity: &DiscountTypeEntity{},
    ResponseEntity: &DiscountTypeEntity{},
  }
  /**
  *	Override this function on DiscountTypeEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendDiscountTypeRouter = func(r *[]workspaces.Module2Action) {}
  func GetDiscountTypeModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
       {
        Method: "GET",
        Url:    "/discount-types",
        SecurityModel: &workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_DISCOUNTTYPE_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpQueryEntity(c, DiscountTypeActionQuery)
          },
        },
        Format: "QUERY",
        Action: DiscountTypeActionQuery,
        ResponseEntity: &[]DiscountTypeEntity{},
      },
      {
        Method: "GET",
        Url:    "/discount-types/export",
        SecurityModel: &workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_DISCOUNTTYPE_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpStreamFileChannel(c, DiscountTypeActionExport)
          },
        },
        Format: "QUERY",
        Action: DiscountTypeActionExport,
        ResponseEntity: &[]DiscountTypeEntity{},
      },
      {
        Method: "GET",
        Url:    "/discount-type/:uniqueId",
        SecurityModel: &workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_DISCOUNTTYPE_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpGetEntity(c, DiscountTypeActionGetOne)
          },
        },
        Format: "GET_ONE",
        Action: DiscountTypeActionGetOne,
        ResponseEntity: &DiscountTypeEntity{},
      },
      DISCOUNTTYPE_ACTION_POST_ONE,
      {
        ActionName:    "update",
        ActionAliases: []string{"u"},
        Flags: DiscountTypeCommonCliFlagsOptional,
        Method: "PATCH",
        Url:    "/discount-type",
        SecurityModel: &workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_DISCOUNTTYPE_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpUpdateEntity(c, DiscountTypeActionUpdate)
          },
        },
        Action: DiscountTypeActionUpdate,
        RequestEntity: &DiscountTypeEntity{},
        Format: "PATCH_ONE",
        ResponseEntity: &DiscountTypeEntity{},
      },
      {
        Method: "PATCH",
        Url:    "/discount-types",
        SecurityModel: &workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_DISCOUNTTYPE_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpUpdateEntities(c, DiscountTypeActionBulkUpdate)
          },
        },
        Action: DiscountTypeActionBulkUpdate,
        Format: "PATCH_BULK",
        RequestEntity:  &workspaces.BulkRecordRequest[DiscountTypeEntity]{},
        ResponseEntity: &workspaces.BulkRecordRequest[DiscountTypeEntity]{},
      },
      {
        Method: "DELETE",
        Url:    "/discount-type",
        Format: "DELETE_DSL",
        SecurityModel: &workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_DISCOUNTTYPE_DELETE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpRemoveEntity(c, DiscountTypeActionRemove)
          },
        },
        Action: DiscountTypeActionRemove,
        RequestEntity: &workspaces.DeleteRequest{},
        ResponseEntity: &workspaces.DeleteResponse{},
        TargetEntity: &DiscountTypeEntity{},
      },
    }
    // Append user defined functions
    AppendDiscountTypeRouter(&routes)
    return routes
  }
  func CreateDiscountTypeRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetDiscountTypeModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, DiscountTypeEntityJsonSchema, "discount-type-http", "shop")
    workspaces.WriteEntitySchema("DiscountTypeEntity", DiscountTypeEntityJsonSchema, "shop")
    return httpRoutes
  }
var PERM_ROOT_DISCOUNTTYPE_DELETE = "root/discounttype/delete"
var PERM_ROOT_DISCOUNTTYPE_CREATE = "root/discounttype/create"
var PERM_ROOT_DISCOUNTTYPE_UPDATE = "root/discounttype/update"
var PERM_ROOT_DISCOUNTTYPE_QUERY = "root/discounttype/query"
var PERM_ROOT_DISCOUNTTYPE = "root/discounttype"
var ALL_DISCOUNTTYPE_PERMISSIONS = []string{
	PERM_ROOT_DISCOUNTTYPE_DELETE,
	PERM_ROOT_DISCOUNTTYPE_CREATE,
	PERM_ROOT_DISCOUNTTYPE_UPDATE,
	PERM_ROOT_DISCOUNTTYPE_QUERY,
	PERM_ROOT_DISCOUNTTYPE,
}