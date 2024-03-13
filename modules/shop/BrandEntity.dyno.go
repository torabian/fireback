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
	mocks "github.com/torabian/fireback/modules/shop/mocks/Brand"
)
type BrandEntity struct {
    Visibility       *string                         `json:"visibility,omitempty" yaml:"visibility"`
    WorkspaceId      *string                         `json:"workspaceId,omitempty" yaml:"workspaceId"`
    LinkerId         *string                         `json:"linkerId,omitempty" yaml:"linkerId"`
    ParentId         *string                         `json:"parentId,omitempty" yaml:"parentId"`
    UniqueId         string                          `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
    UserId           *string                         `json:"userId,omitempty" yaml:"userId"`
    Rank             int64                           `json:"rank,omitempty" gorm:"type:int;name:rank"`
    Updated          int64                           `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
    Created          int64                           `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
    CreatedFormatted string                          `json:"createdFormatted,omitempty" sql:"-"`
    UpdatedFormatted string                          `json:"updatedFormatted,omitempty" sql:"-"`
    Name   *string `json:"name" yaml:"name"        translate:"true" `
    // Datenano also has a text representation
    Translations     []*BrandEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    Children []*BrandEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *BrandEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var BrandPreloadRelations []string = []string{}
var BRAND_EVENT_CREATED = "brand.created"
var BRAND_EVENT_UPDATED = "brand.updated"
var BRAND_EVENT_DELETED = "brand.deleted"
var BRAND_EVENTS = []string{
	BRAND_EVENT_CREATED,
	BRAND_EVENT_UPDATED,
	BRAND_EVENT_DELETED,
}
type BrandFieldMap struct {
		Name workspaces.TranslatedString `yaml:"name"`
}
var BrandEntityMetaConfig map[string]int64 = map[string]int64{
}
var BrandEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&BrandEntity{}))
  type BrandEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        Name string `yaml:"name" json:"name"`
  }
func entityBrandFormatter(dto *BrandEntity, query workspaces.QueryDSL) {
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
func BrandMockEntity() *BrandEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &BrandEntity{
      Name : &stringHolder,
	}
	return entity
}
func BrandActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := BrandMockEntity()
		_, err := BrandActionCreate(entity, query)
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
    func (x*BrandEntity) GetNameTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Name
          }
        }
      }
      return ""
    }
  func BrandActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*BrandEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &BrandEntity{
          Name: &tildaRef,
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
  func BrandAssociationCreate(dto *BrandEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func BrandRelationContentCreate(dto *BrandEntity, query workspaces.QueryDSL) error {
return nil
}
func BrandRelationContentUpdate(dto *BrandEntity, query workspaces.QueryDSL) error {
	return nil
}
func BrandPolyglotCreateHandler(dto *BrandEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &BrandEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func BrandValidator(dto *BrandEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func BrandEntityPreSanitize(dto *BrandEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func BrandEntityBeforeCreateAppend(dto *BrandEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    BrandRecursiveAddUniqueId(dto, query)
  }
  func BrandRecursiveAddUniqueId(dto *BrandEntity, query workspaces.QueryDSL) {
  }
func BrandActionBatchCreateFn(dtos []*BrandEntity, query workspaces.QueryDSL) ([]*BrandEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*BrandEntity{}
		for _, item := range dtos {
			s, err := BrandActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func BrandActionCreateFn(dto *BrandEntity, query workspaces.QueryDSL) (*BrandEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := BrandValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	BrandEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	BrandEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	BrandPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	BrandRelationContentCreate(dto, query)
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
	BrandAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(BRAND_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&BrandEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func BrandActionGetOne(query workspaces.QueryDSL) (*BrandEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&BrandEntity{})
    item, err := workspaces.GetOneEntity[BrandEntity](query, refl)
    entityBrandFormatter(item, query)
    return item, err
  }
  func BrandActionQuery(query workspaces.QueryDSL) ([]*BrandEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&BrandEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[BrandEntity](query, refl)
    for _, item := range items {
      entityBrandFormatter(item, query)
    }
    return items, meta, err
  }
  func BrandUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *BrandEntity) (*BrandEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = BRAND_EVENT_UPDATED
    BrandEntityPreSanitize(fields, query)
    var item BrandEntity
    q := dbref.
      Where(&BrandEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    BrandRelationContentUpdate(fields, query)
    BrandPolyglotCreateHandler(fields, query)
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&BrandEntity{UniqueId: uniqueId}).
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
  func BrandActionUpdateFn(query workspaces.QueryDSL, fields *BrandEntity) (*BrandEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := BrandValidator(fields, true); iError != nil {
      return nil, iError
    }
    BrandRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        _, err := BrandUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return nil, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return BrandUpdateExec(dbref, query, fields)
    }
  }
var BrandWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire brands ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		count, _ := BrandActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func BrandActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&BrandEntity{})
	query.ActionRequires = []string{PERM_ROOT_BRAND_DELETE}
	return workspaces.RemoveEntity[BrandEntity](query, refl)
}
func BrandActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := workspaces.WipeCleanEntity[BrandEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'BrandEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func BrandActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[BrandEntity]) (
    *workspaces.BulkRecordRequest[BrandEntity], *workspaces.IError,
  ) {
    result := []*BrandEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := BrandActionUpdate(query, record)
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
func (x *BrandEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var BrandEntityMeta = workspaces.TableMetaData{
	EntityName:    "Brand",
	ExportKey:    "brands",
	TableNameInDb: "fb_brand_entities",
	EntityObject:  &BrandEntity{},
	ExportStream: BrandActionExportT,
	ImportQuery: BrandActionImport,
}
func BrandActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[BrandEntity](query, BrandActionQuery, BrandPreloadRelations)
}
func BrandActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[BrandEntity](query, BrandActionQuery, BrandPreloadRelations)
}
func BrandActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content BrandEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := BrandActionCreate(&content, query)
	return err
}
var BrandCommonCliFlags = []cli.Flag{
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
}
var BrandCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "name",
		StructField:     "Name",
		Required: false,
		Usage:    "name",
		Type: "string",
	},
}
var BrandCommonCliFlagsOptional = []cli.Flag{
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
}
  var BrandCreateCmd cli.Command = cli.Command{
    Name:    "create",
    Aliases: []string{"c"},
    Flags: BrandCommonCliFlags,
    Usage: "Create a new template",
    Action: func(c *cli.Context) {
      query := workspaces.CommonCliQueryDSLBuilder(c)
      entity := CastBrandFromCli(c)
      if entity, err := BrandActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var BrandCreateInteractiveCmd cli.Command = cli.Command{
    Name:  "ic",
    Usage: "Creates a new template, using requied fields in an interactive name",
    Flags: []cli.Flag{
      &cli.BoolFlag{
        Name:  "all",
        Usage: "Interactively asks for all inputs, not only required ones",
      },
    },
    Action: func(c *cli.Context) {
      query := workspaces.CommonCliQueryDSLBuilder(c)
      entity := &BrandEntity{}
      for _, item := range BrandCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := BrandActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var BrandUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: BrandCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilder(c)
      entity := CastBrandFromCli(c)
      if entity, err := BrandActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x BrandEntity) FromCli(c *cli.Context) *BrandEntity {
	return CastBrandFromCli(c)
}
func CastBrandFromCli (c *cli.Context) *BrandEntity {
	template := &BrandEntity{}
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
	return template
}
  func BrandSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      BrandActionCreate,
      reflect.ValueOf(&BrandEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func BrandImportMocks() {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      BrandActionCreate,
      reflect.ValueOf(&BrandEntity{}).Elem(),
      &mocks.ViewsFs,
      []string{},
      false,
    )
  }
  func BrandWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := BrandActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "Brand", result)
    }
  }
var BrandImportExportCommands = []cli.Command{
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
			query := workspaces.CommonCliQueryDSLBuilder(c)
			BrandActionSeeder(query, c.Int("count"))
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
				Value: "brand-seeder.yml",
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
			f := workspaces.CommonCliQueryDSLBuilder(c)
			BrandActionSeederInit(f, c.String("file"), c.String("format"))
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
				Value: "brand-seeder-brand.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of brands, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]BrandEntity{}
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
					BrandActionCreate,
					reflect.ValueOf(&BrandEntity{}).Elem(),
					&mocks.ViewsFs,
				)
				return nil
			},
		},
	cli.Command{
		Name:    "import",
		Flags: append(workspaces.CommonQueryFlags,
			&cli.StringFlag{
				Name:     "file",
				Usage:    "The address of file you want the csv be imported from",
				Required: true,
			}),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmd(c,
				BrandActionCreate,
				reflect.ValueOf(&BrandEntity{}).Elem(),
				c.String("file"),
			)
			return nil
		},
	},
}
    var BrandCliCommands []cli.Command = []cli.Command{
      workspaces.GetCommonQuery(BrandActionQuery),
      workspaces.GetCommonTableQuery(reflect.ValueOf(&BrandEntity{}).Elem(), BrandActionQuery),
          BrandCreateCmd,
          BrandUpdateCmd,
          BrandCreateInteractiveCmd,
          BrandWipeCmd,
          workspaces.GetCommonRemoveQuery(reflect.ValueOf(&BrandEntity{}).Elem(), BrandActionRemove),
  }
  func BrandCliFn() cli.Command {
    BrandCliCommands = append(BrandCliCommands, BrandImportExportCommands...)
    return cli.Command{
      Name:        "brand",
      Description: "Brands module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: BrandCliCommands,
    }
  }
  /**
  *	Override this function on BrandEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendBrandRouter = func(r *[]workspaces.Module2Action) {}
  func GetBrandModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
       {
        Method: "GET",
        Url:    "/brands",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_BRAND_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpQueryEntity(c, BrandActionQuery)
          },
        },
        Format: "QUERY",
        Action: BrandActionQuery,
        ResponseEntity: &[]BrandEntity{},
      },
      {
        Method: "GET",
        Url:    "/brands/export",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_BRAND_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpStreamFileChannel(c, BrandActionExport)
          },
        },
        Format: "QUERY",
        Action: BrandActionExport,
        ResponseEntity: &[]BrandEntity{},
      },
      {
        Method: "GET",
        Url:    "/brand/:uniqueId",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_BRAND_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpGetEntity(c, BrandActionGetOne)
          },
        },
        Format: "GET_ONE",
        Action: BrandActionGetOne,
        ResponseEntity: &BrandEntity{},
      },
      {
        ActionName:    "create",
        ActionAliases: []string{"c"},
        Flags: BrandCommonCliFlags,
        Method: "POST",
        Url:    "/brand",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_BRAND_CREATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpPostEntity(c, BrandActionCreate)
          },
        },
        Action: BrandActionCreate,
        Format: "POST_ONE",
        RequestEntity: &BrandEntity{},
        ResponseEntity: &BrandEntity{},
      },
      {
        ActionName:    "update",
        ActionAliases: []string{"u"},
        Flags: BrandCommonCliFlagsOptional,
        Method: "PATCH",
        Url:    "/brand",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_BRAND_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpUpdateEntity(c, BrandActionUpdate)
          },
        },
        Action: BrandActionUpdate,
        RequestEntity: &BrandEntity{},
        Format: "PATCH_ONE",
        ResponseEntity: &BrandEntity{},
      },
      {
        Method: "PATCH",
        Url:    "/brands",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_BRAND_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpUpdateEntities(c, BrandActionBulkUpdate)
          },
        },
        Action: BrandActionBulkUpdate,
        Format: "PATCH_BULK",
        RequestEntity:  &workspaces.BulkRecordRequest[BrandEntity]{},
        ResponseEntity: &workspaces.BulkRecordRequest[BrandEntity]{},
      },
      {
        Method: "DELETE",
        Url:    "/brand",
        Format: "DELETE_DSL",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_BRAND_DELETE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpRemoveEntity(c, BrandActionRemove)
          },
        },
        Action: BrandActionRemove,
        RequestEntity: &workspaces.DeleteRequest{},
        ResponseEntity: &workspaces.DeleteResponse{},
        TargetEntity: &BrandEntity{},
      },
    }
    // Append user defined functions
    AppendBrandRouter(&routes)
    return routes
  }
  func CreateBrandRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetBrandModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, BrandEntityJsonSchema, "brand-http", "shop")
    workspaces.WriteEntitySchema("BrandEntity", BrandEntityJsonSchema, "shop")
    return httpRoutes
  }
var PERM_ROOT_BRAND_DELETE = "root/brand/delete"
var PERM_ROOT_BRAND_CREATE = "root/brand/create"
var PERM_ROOT_BRAND_UPDATE = "root/brand/update"
var PERM_ROOT_BRAND_QUERY = "root/brand/query"
var PERM_ROOT_BRAND = "root/brand"
var ALL_BRAND_PERMISSIONS = []string{
	PERM_ROOT_BRAND_DELETE,
	PERM_ROOT_BRAND_CREATE,
	PERM_ROOT_BRAND_UPDATE,
	PERM_ROOT_BRAND_QUERY,
	PERM_ROOT_BRAND,
}