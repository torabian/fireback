package licenses
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
	mocks "github.com/torabian/fireback/modules/licenses/mocks/LicensableProduct"
)
var licensableProductSeedersFs *embed.FS = nil
func ResetLicensableProductSeeders(fs *embed.FS) {
	licensableProductSeedersFs = fs
}
type LicensableProductEntity struct {
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
    Name   *string `json:"name" yaml:"name"  validate:"required,omitempty,min=1,max=100"        translate:"true" `
    // Datenano also has a text representation
    PrivateKey   *string `json:"privateKey" yaml:"privateKey"       `
    // Datenano also has a text representation
    PublicKey   *string `json:"publicKey" yaml:"publicKey"       `
    // Datenano also has a text representation
    Translations     []*LicensableProductEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    Children []*LicensableProductEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *LicensableProductEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var LicensableProductPreloadRelations []string = []string{}
var LICENSABLE_PRODUCT_EVENT_CREATED = "licensableProduct.created"
var LICENSABLE_PRODUCT_EVENT_UPDATED = "licensableProduct.updated"
var LICENSABLE_PRODUCT_EVENT_DELETED = "licensableProduct.deleted"
var LICENSABLE_PRODUCT_EVENTS = []string{
	LICENSABLE_PRODUCT_EVENT_CREATED,
	LICENSABLE_PRODUCT_EVENT_UPDATED,
	LICENSABLE_PRODUCT_EVENT_DELETED,
}
type LicensableProductFieldMap struct {
		Name workspaces.TranslatedString `yaml:"name"`
		PrivateKey workspaces.TranslatedString `yaml:"privateKey"`
		PublicKey workspaces.TranslatedString `yaml:"publicKey"`
}
var LicensableProductEntityMetaConfig map[string]int64 = map[string]int64{
}
var LicensableProductEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&LicensableProductEntity{}))
  type LicensableProductEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        Name string `yaml:"name" json:"name"`
  }
func entityLicensableProductFormatter(dto *LicensableProductEntity, query workspaces.QueryDSL) {
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
func LicensableProductMockEntity() *LicensableProductEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &LicensableProductEntity{
      Name : &stringHolder,
      PrivateKey : &stringHolder,
      PublicKey : &stringHolder,
	}
	return entity
}
func LicensableProductActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := LicensableProductMockEntity()
		_, err := LicensableProductActionCreate(entity, query)
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
    func (x*LicensableProductEntity) GetNameTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Name
          }
        }
      }
      return ""
    }
  func LicensableProductActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*LicensableProductEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &LicensableProductEntity{
          Name: &tildaRef,
          PrivateKey: &tildaRef,
          PublicKey: &tildaRef,
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
  func LicensableProductAssociationCreate(dto *LicensableProductEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func LicensableProductRelationContentCreate(dto *LicensableProductEntity, query workspaces.QueryDSL) error {
return nil
}
func LicensableProductRelationContentUpdate(dto *LicensableProductEntity, query workspaces.QueryDSL) error {
	return nil
}
func LicensableProductPolyglotCreateHandler(dto *LicensableProductEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &LicensableProductEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func LicensableProductValidator(dto *LicensableProductEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func LicensableProductEntityPreSanitize(dto *LicensableProductEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func LicensableProductEntityBeforeCreateAppend(dto *LicensableProductEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    LicensableProductRecursiveAddUniqueId(dto, query)
  }
  func LicensableProductRecursiveAddUniqueId(dto *LicensableProductEntity, query workspaces.QueryDSL) {
  }
func LicensableProductActionBatchCreateFn(dtos []*LicensableProductEntity, query workspaces.QueryDSL) ([]*LicensableProductEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*LicensableProductEntity{}
		for _, item := range dtos {
			s, err := LicensableProductActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func LicensableProductDeleteEntireChildren(query workspaces.QueryDSL, dto *LicensableProductEntity) (*workspaces.IError) {
  // intentionally removed this. It's hard to implement it, and probably wrong without
  // proper on delete cascade
  return nil
}
func LicensableProductActionCreateFn(dto *LicensableProductEntity, query workspaces.QueryDSL) (*LicensableProductEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := LicensableProductValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	LicensableProductEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	LicensableProductEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	LicensableProductPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	LicensableProductRelationContentCreate(dto, query)
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
	LicensableProductAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(LICENSABLE_PRODUCT_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&LicensableProductEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func LicensableProductActionGetOne(query workspaces.QueryDSL) (*LicensableProductEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&LicensableProductEntity{})
    item, err := workspaces.GetOneEntity[LicensableProductEntity](query, refl)
    entityLicensableProductFormatter(item, query)
    return item, err
  }
  func LicensableProductActionQuery(query workspaces.QueryDSL) ([]*LicensableProductEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&LicensableProductEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[LicensableProductEntity](query, refl)
    for _, item := range items {
      entityLicensableProductFormatter(item, query)
    }
    return items, meta, err
  }
  func LicensableProductUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *LicensableProductEntity) (*LicensableProductEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = LICENSABLE_PRODUCT_EVENT_UPDATED
    LicensableProductEntityPreSanitize(fields, query)
    var item LicensableProductEntity
    q := dbref.
      Where(&LicensableProductEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    LicensableProductRelationContentUpdate(fields, query)
    LicensableProductPolyglotCreateHandler(fields, query)
    if ero := LicensableProductDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&LicensableProductEntity{UniqueId: uniqueId}).
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
  func LicensableProductActionUpdateFn(query workspaces.QueryDSL, fields *LicensableProductEntity) (*LicensableProductEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := LicensableProductValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // LicensableProductRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      var item *LicensableProductEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *workspaces.IError
        item, err = LicensableProductUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return LicensableProductUpdateExec(dbref, query, fields)
    }
  }
var LicensableProductWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire licensableproducts ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
      ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_LICENSABLE_PRODUCT_DELETE},
    })
		count, _ := LicensableProductActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func LicensableProductActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&LicensableProductEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_LICENSABLE_PRODUCT_DELETE}
	return workspaces.RemoveEntity[LicensableProductEntity](query, refl)
}
func LicensableProductActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[LicensableProductEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'LicensableProductEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func LicensableProductActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[LicensableProductEntity]) (
    *workspaces.BulkRecordRequest[LicensableProductEntity], *workspaces.IError,
  ) {
    result := []*LicensableProductEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := LicensableProductActionUpdate(query, record)
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
func (x *LicensableProductEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var LicensableProductEntityMeta = workspaces.TableMetaData{
	EntityName:    "LicensableProduct",
	ExportKey:    "licensable-products",
	TableNameInDb: "fb_licensable-product_entities",
	EntityObject:  &LicensableProductEntity{},
	ExportStream: LicensableProductActionExportT,
	ImportQuery: LicensableProductActionImport,
}
func LicensableProductActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[LicensableProductEntity](query, LicensableProductActionQuery, LicensableProductPreloadRelations)
}
func LicensableProductActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[LicensableProductEntity](query, LicensableProductActionQuery, LicensableProductPreloadRelations)
}
func LicensableProductActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content LicensableProductEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := LicensableProductActionCreate(&content, query)
	return err
}
var LicensableProductCommonCliFlags = []cli.Flag{
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
      Required: true,
      Usage:    "name",
    },
    &cli.StringFlag{
      Name:     "private-key",
      Required: false,
      Usage:    "privateKey",
    },
    &cli.StringFlag{
      Name:     "public-key",
      Required: false,
      Usage:    "publicKey",
    },
}
var LicensableProductCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "name",
		StructField:     "Name",
		Required: true,
		Usage:    "name",
		Type: "string",
	},
	{
		Name:     "privateKey",
		StructField:     "PrivateKey",
		Required: false,
		Usage:    "privateKey",
		Type: "string",
	},
	{
		Name:     "publicKey",
		StructField:     "PublicKey",
		Required: false,
		Usage:    "publicKey",
		Type: "string",
	},
}
var LicensableProductCommonCliFlagsOptional = []cli.Flag{
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
      Required: true,
      Usage:    "name",
    },
    &cli.StringFlag{
      Name:     "private-key",
      Required: false,
      Usage:    "privateKey",
    },
    &cli.StringFlag{
      Name:     "public-key",
      Required: false,
      Usage:    "publicKey",
    },
}
  var LicensableProductCreateCmd cli.Command = LICENSABLE_PRODUCT_ACTION_POST_ONE.ToCli()
  var LicensableProductCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_LICENSABLE_PRODUCT_CREATE},
      })
      entity := &LicensableProductEntity{}
      for _, item := range LicensableProductCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := LicensableProductActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var LicensableProductUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: LicensableProductCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_LICENSABLE_PRODUCT_UPDATE},
      })
      entity := CastLicensableProductFromCli(c)
      if entity, err := LicensableProductActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* LicensableProductEntity) FromCli(c *cli.Context) *LicensableProductEntity {
	return CastLicensableProductFromCli(c)
}
func CastLicensableProductFromCli (c *cli.Context) *LicensableProductEntity {
	template := &LicensableProductEntity{}
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
      if c.IsSet("private-key") {
        value := c.String("private-key")
        template.PrivateKey = &value
      }
      if c.IsSet("public-key") {
        value := c.String("public-key")
        template.PublicKey = &value
      }
	return template
}
  func LicensableProductSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      LicensableProductActionCreate,
      reflect.ValueOf(&LicensableProductEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func LicensableProductImportMocks() {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      LicensableProductActionCreate,
      reflect.ValueOf(&LicensableProductEntity{}).Elem(),
      &mocks.ViewsFs,
      []string{},
      false,
    )
  }
  func LicensableProductWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := LicensableProductActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "LicensableProduct", result)
    }
  }
var LicensableProductImportExportCommands = []cli.Command{
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_LICENSABLE_PRODUCT_CREATE},
      })
			LicensableProductActionSeeder(query, c.Int("count"))
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
				Value: "licensable-product-seeder.yml",
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
        ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_LICENSABLE_PRODUCT_CREATE},
      })
			LicensableProductActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "licensable-product-seeder-licensable-product.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of licensable-products, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]LicensableProductEntity{}
			workspaces.ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
		cli.Command{
			Name:  "mocks",
			Usage: "Prints the list of mocks",
			Action: func(c *cli.Context) error {
				if entity, err := workspaces.GetSeederFilenames(&mocks.ViewsFs, ""); err != nil {
					fmt.Println(err.Error())
				} else {
					f, _ := json.MarshalIndent(entity, "", "  ")
					fmt.Println(string(f))
				}
				return nil
			},
		},
		cli.Command{
			Name:  "msync",
			Usage: "Tries to sync mocks into the system",
			Action: func(c *cli.Context) error {
				workspaces.CommonCliImportEmbedCmd(c,
					LicensableProductActionCreate,
					reflect.ValueOf(&LicensableProductEntity{}).Elem(),
					&mocks.ViewsFs,
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
			LicensableProductCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				LicensableProductActionCreate,
				reflect.ValueOf(&LicensableProductEntity{}).Elem(),
				c.String("file"),
        &workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_LICENSABLE_PRODUCT_CREATE},
				},
        func() LicensableProductEntity {
					v := CastLicensableProductFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var LicensableProductCliCommands []cli.Command = []cli.Command{
      LICENSABLE_PRODUCT_ACTION_QUERY.ToCli(),
      LICENSABLE_PRODUCT_ACTION_TABLE.ToCli(),
      LicensableProductCreateCmd,
      LicensableProductUpdateCmd,
      LicensableProductCreateInteractiveCmd,
      LicensableProductWipeCmd,
      workspaces.GetCommonRemoveQuery(reflect.ValueOf(&LicensableProductEntity{}).Elem(), LicensableProductActionRemove),
  }
  func LicensableProductCliFn() cli.Command {
    LicensableProductCliCommands = append(LicensableProductCliCommands, LicensableProductImportExportCommands...)
    return cli.Command{
      Name:        "product",
      Description: "LicensableProducts module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: LicensableProductCliCommands,
    }
  }
var LICENSABLE_PRODUCT_ACTION_TABLE = workspaces.Module2Action{
  Name:    "table",
  ActionName: "table",
  ActionAliases: []string{"t"},
  Flags:  workspaces.CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: LicensableProductActionQuery,
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    workspaces.CommonCliTableCmd2(c,
      LicensableProductActionQuery,
      security,
      reflect.ValueOf(&LicensableProductEntity{}).Elem(),
    )
    return nil
  },
}
var LICENSABLE_PRODUCT_ACTION_QUERY = workspaces.Module2Action{
  Method: "GET",
  Url:    "/licensable-products",
  SecurityModel: &workspaces.SecurityModel{
  },
  Group: "licensableProduct",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpQueryEntity(c, LicensableProductActionQuery)
    },
  },
  Format: "QUERY",
  Action: LicensableProductActionQuery,
  ResponseEntity: &[]LicensableProductEntity{},
  Out: workspaces.Module2ActionBody{
		Entity: "LicensableProductEntity",
	},
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		workspaces.CommonCliQueryCmd2(
			c,
			LicensableProductActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionName:    "query",
	ActionAliases: []string{"q"},
	Flags:         workspaces.CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var LICENSABLE_PRODUCT_ACTION_EXPORT = workspaces.Module2Action{
  Method: "GET",
  Url:    "/licensable-products/export",
  SecurityModel: &workspaces.SecurityModel{
  },
  Group: "licensableProduct",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpStreamFileChannel(c, LicensableProductActionExport)
    },
  },
  Format: "QUERY",
  Action: LicensableProductActionExport,
  ResponseEntity: &[]LicensableProductEntity{},
  Out: workspaces.Module2ActionBody{
		Entity: "LicensableProductEntity",
	},
}
var LICENSABLE_PRODUCT_ACTION_GET_ONE = workspaces.Module2Action{
  Method: "GET",
  Url:    "/licensable-product/:uniqueId",
  SecurityModel: &workspaces.SecurityModel{
  },
  Group: "licensableProduct",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpGetEntity(c, LicensableProductActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: LicensableProductActionGetOne,
  ResponseEntity: &LicensableProductEntity{},
  Out: workspaces.Module2ActionBody{
		Entity: "LicensableProductEntity",
	},
}
var LICENSABLE_PRODUCT_ACTION_POST_ONE = workspaces.Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new licensableProduct",
  Flags: LicensableProductCommonCliFlags,
  Method: "POST",
  Url:    "/licensable-product",
  SecurityModel: &workspaces.SecurityModel{
  },
  Group: "licensableProduct",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpPostEntity(c, LicensableProductActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
    result, err := workspaces.CliPostEntity(c, LicensableProductActionCreate, security)
    workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: LicensableProductActionCreate,
  Format: "POST_ONE",
  RequestEntity: &LicensableProductEntity{},
  ResponseEntity: &LicensableProductEntity{},
  Out: workspaces.Module2ActionBody{
		Entity: "LicensableProductEntity",
	},
  In: workspaces.Module2ActionBody{
		Entity: "LicensableProductEntity",
	},
}
var LICENSABLE_PRODUCT_ACTION_PATCH = workspaces.Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: LicensableProductCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/licensable-product",
  SecurityModel: &workspaces.SecurityModel{
  },
  Group: "licensableProduct",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntity(c, LicensableProductActionUpdate)
    },
  },
  Action: LicensableProductActionUpdate,
  RequestEntity: &LicensableProductEntity{},
  ResponseEntity: &LicensableProductEntity{},
  Format: "PATCH_ONE",
  Out: workspaces.Module2ActionBody{
		Entity: "LicensableProductEntity",
	},
  In: workspaces.Module2ActionBody{
		Entity: "LicensableProductEntity",
	},
}
var LICENSABLE_PRODUCT_ACTION_PATCH_BULK = workspaces.Module2Action{
  Method: "PATCH",
  Url:    "/licensable-products",
  SecurityModel: &workspaces.SecurityModel{
  },
  Group: "licensableProduct",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpUpdateEntities(c, LicensableProductActionBulkUpdate)
    },
  },
  Action: LicensableProductActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &workspaces.BulkRecordRequest[LicensableProductEntity]{},
  ResponseEntity: &workspaces.BulkRecordRequest[LicensableProductEntity]{},
  Out: workspaces.Module2ActionBody{
		Entity: "LicensableProductEntity",
	},
  In: workspaces.Module2ActionBody{
		Entity: "LicensableProductEntity",
	},
}
var LICENSABLE_PRODUCT_ACTION_DELETE = workspaces.Module2Action{
  Method: "DELETE",
  Url:    "/licensable-product",
  Format: "DELETE_DSL",
  SecurityModel: &workspaces.SecurityModel{
  },
  Group: "licensableProduct",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      workspaces.HttpRemoveEntity(c, LicensableProductActionRemove)
    },
  },
  Action: LicensableProductActionRemove,
  RequestEntity: &workspaces.DeleteRequest{},
  ResponseEntity: &workspaces.DeleteResponse{},
  TargetEntity: &LicensableProductEntity{},
}
  /**
  *	Override this function on LicensableProductEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendLicensableProductRouter = func(r *[]workspaces.Module2Action) {}
  func GetLicensableProductModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
      LICENSABLE_PRODUCT_ACTION_QUERY,
      LICENSABLE_PRODUCT_ACTION_EXPORT,
      LICENSABLE_PRODUCT_ACTION_GET_ONE,
      LICENSABLE_PRODUCT_ACTION_POST_ONE,
      LICENSABLE_PRODUCT_ACTION_PATCH,
      LICENSABLE_PRODUCT_ACTION_PATCH_BULK,
      LICENSABLE_PRODUCT_ACTION_DELETE,
    }
    // Append user defined functions
    AppendLicensableProductRouter(&routes)
    return routes
  }
  func CreateLicensableProductRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetLicensableProductModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, LicensableProductEntityJsonSchema, "licensable-product-http", "licenses")
    workspaces.WriteEntitySchema("LicensableProductEntity", LicensableProductEntityJsonSchema, "licenses")
    return httpRoutes
  }
var PERM_ROOT_LICENSABLE_PRODUCT_DELETE = workspaces.PermissionInfo{
  CompleteKey: "root/licenses/licensable-product/delete",
  Name: "Delete licensable product",
}
var PERM_ROOT_LICENSABLE_PRODUCT_CREATE = workspaces.PermissionInfo{
  CompleteKey: "root/licenses/licensable-product/create",
  Name: "Create licensable product",
}
var PERM_ROOT_LICENSABLE_PRODUCT_UPDATE = workspaces.PermissionInfo{
  CompleteKey: "root/licenses/licensable-product/update",
  Name: "Update licensable product",
}
var PERM_ROOT_LICENSABLE_PRODUCT_QUERY = workspaces.PermissionInfo{
  CompleteKey: "root/licenses/licensable-product/query",
  Name: "Query licensable product",
}
var PERM_ROOT_LICENSABLE_PRODUCT = workspaces.PermissionInfo{
  CompleteKey: "root/licenses/licensable-product/*",
  Name: "Entire licensable product actions (*)",
}
var ALL_LICENSABLE_PRODUCT_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_LICENSABLE_PRODUCT_DELETE,
	PERM_ROOT_LICENSABLE_PRODUCT_CREATE,
	PERM_ROOT_LICENSABLE_PRODUCT_UPDATE,
	PERM_ROOT_LICENSABLE_PRODUCT_QUERY,
	PERM_ROOT_LICENSABLE_PRODUCT,
}