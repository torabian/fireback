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
type PersonEntity struct {
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
    FirstName   *string `json:"firstName" yaml:"firstName"  validate:"required"       `
    // Datenano also has a text representation
    LastName   *string `json:"lastName" yaml:"lastName"  validate:"required"       `
    // Datenano also has a text representation
    Photo   *string `json:"photo" yaml:"photo"       `
    // Datenano also has a text representation
    Gender   *string `json:"gender" yaml:"gender"       `
    // Datenano also has a text representation
    Title   *string `json:"title" yaml:"title"       `
    // Datenano also has a text representation
    BirthDate   XDate `json:"birthDate" yaml:"birthDate"       `
    // Datenano also has a text representation
    // Date range is a complex date storage
    BirthDateDateInfo XDateMetaData `json:"birthDateDateInfo" yaml:"birthDateDateInfo" sql:"-" gorm:"-"`
    Children []*PersonEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *PersonEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var PersonPreloadRelations []string = []string{}
var PERSON_EVENT_CREATED = "person.created"
var PERSON_EVENT_UPDATED = "person.updated"
var PERSON_EVENT_DELETED = "person.deleted"
var PERSON_EVENTS = []string{
	PERSON_EVENT_CREATED,
	PERSON_EVENT_UPDATED,
	PERSON_EVENT_DELETED,
}
type PersonFieldMap struct {
		FirstName TranslatedString `yaml:"firstName"`
		LastName TranslatedString `yaml:"lastName"`
		Photo TranslatedString `yaml:"photo"`
		Gender TranslatedString `yaml:"gender"`
		Title TranslatedString `yaml:"title"`
		BirthDate TranslatedString `yaml:"birthDate"`
}
var PersonEntityMetaConfig map[string]int64 = map[string]int64{
}
var PersonEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&PersonEntity{}))
func entityPersonFormatter(dto *PersonEntity, query QueryDSL) {
	if dto == nil {
		return
	}
			dto.BirthDateDateInfo = ComputeXDateMetaData(&dto.BirthDate, query)
	if dto.Created > 0 {
		dto.CreatedFormatted = FormatDateBasedOnQuery(dto.Created, query)
	}
	if dto.Updated > 0 {
		dto.CreatedFormatted = FormatDateBasedOnQuery(dto.Updated, query)
	}
}
func PersonMockEntity() *PersonEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &PersonEntity{
      FirstName : &stringHolder,
      LastName : &stringHolder,
      Photo : &stringHolder,
      Gender : &stringHolder,
      Title : &stringHolder,
	}
	return entity
}
func PersonActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := PersonMockEntity()
		_, err := PersonActionCreate(entity, query)
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
  func PersonActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*PersonEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &PersonEntity{
          FirstName: &tildaRef,
          LastName: &tildaRef,
          Photo: &tildaRef,
          Gender: &tildaRef,
          Title: &tildaRef,
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
  func PersonAssociationCreate(dto *PersonEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func PersonRelationContentCreate(dto *PersonEntity, query QueryDSL) error {
return nil
}
func PersonRelationContentUpdate(dto *PersonEntity, query QueryDSL) error {
	return nil
}
func PersonPolyglotCreateHandler(dto *PersonEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func PersonValidator(dto *PersonEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func PersonEntityPreSanitize(dto *PersonEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func PersonEntityBeforeCreateAppend(dto *PersonEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    PersonRecursiveAddUniqueId(dto, query)
  }
  func PersonRecursiveAddUniqueId(dto *PersonEntity, query QueryDSL) {
  }
func PersonActionBatchCreateFn(dtos []*PersonEntity, query QueryDSL) ([]*PersonEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*PersonEntity{}
		for _, item := range dtos {
			s, err := PersonActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func PersonDeleteEntireChildren(query QueryDSL, dto *PersonEntity) (*IError) {
  return nil
}
func PersonActionCreateFn(dto *PersonEntity, query QueryDSL) (*PersonEntity, *IError) {
	// 1. Validate always
	if iError := PersonValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	PersonEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	PersonEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	PersonPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	PersonRelationContentCreate(dto, query)
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
	PersonAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(PERSON_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&PersonEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func PersonActionGetOne(query QueryDSL) (*PersonEntity, *IError) {
    refl := reflect.ValueOf(&PersonEntity{})
    item, err := GetOneEntity[PersonEntity](query, refl)
    entityPersonFormatter(item, query)
    return item, err
  }
  func PersonActionQuery(query QueryDSL) ([]*PersonEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&PersonEntity{})
    items, meta, err := QueryEntitiesPointer[PersonEntity](query, refl)
    for _, item := range items {
      entityPersonFormatter(item, query)
    }
    return items, meta, err
  }
  func PersonUpdateExec(dbref *gorm.DB, query QueryDSL, fields *PersonEntity) (*PersonEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = PERSON_EVENT_UPDATED
    PersonEntityPreSanitize(fields, query)
    var item PersonEntity
    q := dbref.
      Where(&PersonEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    PersonRelationContentUpdate(fields, query)
    PersonPolyglotCreateHandler(fields, query)
    if ero := PersonDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&PersonEntity{UniqueId: uniqueId}).
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
  func PersonActionUpdateFn(query QueryDSL, fields *PersonEntity) (*PersonEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := PersonValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // PersonRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      var item *PersonEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *IError
        item, err = PersonUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, CastToIError(vf)
    } else {
      dbref = query.Tx
      return PersonUpdateExec(dbref, query, fields)
    }
  }
var PersonWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire people ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
      ActionRequires: []PermissionInfo{PERM_ROOT_PERSON_DELETE},
    })
		count, _ := PersonActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func PersonActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&PersonEntity{})
	query.ActionRequires = []PermissionInfo{PERM_ROOT_PERSON_DELETE}
	return RemoveEntity[PersonEntity](query, refl)
}
func PersonActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[PersonEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'PersonEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func PersonActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[PersonEntity]) (
    *BulkRecordRequest[PersonEntity], *IError,
  ) {
    result := []*PersonEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := PersonActionUpdate(query, record)
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
func (x *PersonEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var PersonEntityMeta = TableMetaData{
	EntityName:    "Person",
	ExportKey:    "people",
	TableNameInDb: "fb_person_entities",
	EntityObject:  &PersonEntity{},
	ExportStream: PersonActionExportT,
	ImportQuery: PersonActionImport,
}
func PersonActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[PersonEntity](query, PersonActionQuery, PersonPreloadRelations)
}
func PersonActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[PersonEntity](query, PersonActionQuery, PersonPreloadRelations)
}
func PersonActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content PersonEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := PersonActionCreate(&content, query)
	return err
}
var PersonCommonCliFlags = []cli.Flag{
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
      Name:     "first-name",
      Required: true,
      Usage:    "firstName",
    },
    &cli.StringFlag{
      Name:     "last-name",
      Required: true,
      Usage:    "lastName",
    },
    &cli.StringFlag{
      Name:     "photo",
      Required: false,
      Usage:    "photo",
    },
    &cli.StringFlag{
      Name:     "gender",
      Required: false,
      Usage:    "gender",
    },
    &cli.StringFlag{
      Name:     "title",
      Required: false,
      Usage:    "title",
    },
    &cli.StringFlag{
      Name:     "birth-date",
      Required: false,
      Usage:    "birthDate",
    },
}
var PersonCommonInteractiveCliFlags = []CliInteractiveFlag{
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
		Name:     "photo",
		StructField:     "Photo",
		Required: false,
		Usage:    "photo",
		Type: "string",
	},
	{
		Name:     "gender",
		StructField:     "Gender",
		Required: false,
		Usage:    "gender",
		Type: "string",
	},
	{
		Name:     "title",
		StructField:     "Title",
		Required: false,
		Usage:    "title",
		Type: "string",
	},
}
var PersonCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "first-name",
      Required: true,
      Usage:    "firstName",
    },
    &cli.StringFlag{
      Name:     "last-name",
      Required: true,
      Usage:    "lastName",
    },
    &cli.StringFlag{
      Name:     "photo",
      Required: false,
      Usage:    "photo",
    },
    &cli.StringFlag{
      Name:     "gender",
      Required: false,
      Usage:    "gender",
    },
    &cli.StringFlag{
      Name:     "title",
      Required: false,
      Usage:    "title",
    },
    &cli.StringFlag{
      Name:     "birth-date",
      Required: false,
      Usage:    "birthDate",
    },
}
  var PersonCreateCmd cli.Command = PERSON_ACTION_POST_ONE.ToCli()
  var PersonCreateInteractiveCmd cli.Command = cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_PERSON_CREATE},
      })
      entity := &PersonEntity{}
      for _, item := range PersonCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := PersonActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var PersonUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: PersonCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_PERSON_UPDATE},
      })
      entity := CastPersonFromCli(c)
      if entity, err := PersonActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* PersonEntity) FromCli(c *cli.Context) *PersonEntity {
	return CastPersonFromCli(c)
}
func CastPersonFromCli (c *cli.Context) *PersonEntity {
	template := &PersonEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("first-name") {
        value := c.String("first-name")
        template.FirstName = &value
      }
      if c.IsSet("last-name") {
        value := c.String("last-name")
        template.LastName = &value
      }
      if c.IsSet("photo") {
        value := c.String("photo")
        template.Photo = &value
      }
      if c.IsSet("gender") {
        value := c.String("gender")
        template.Gender = &value
      }
      if c.IsSet("title") {
        value := c.String("title")
        template.Title = &value
      }
      if c.IsSet("birth-date") {
        value := c.String("birth-date")
        template.BirthDate.Scan(value)
      }
	return template
}
  func PersonSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      PersonActionCreate,
      reflect.ValueOf(&PersonEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func PersonWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := PersonActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "Person", result)
    }
  }
var PersonImportExportCommands = []cli.Command{
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
        ActionRequires: []PermissionInfo{PERM_ROOT_PERSON_CREATE},
      })
			PersonActionSeeder(query, c.Int("count"))
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
				Value: "person-seeder.yml",
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
        ActionRequires: []PermissionInfo{PERM_ROOT_PERSON_CREATE},
      })
			PersonActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "person-seeder-person.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of people, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]PersonEntity{}
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
			PersonCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				PersonActionCreate,
				reflect.ValueOf(&PersonEntity{}).Elem(),
				c.String("file"),
        &SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_PERSON_CREATE},
				},
        func() PersonEntity {
					v := CastPersonFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var PersonCliCommands []cli.Command = []cli.Command{
      PERSON_ACTION_QUERY.ToCli(),
      PERSON_ACTION_TABLE.ToCli(),
      PersonCreateCmd,
      PersonUpdateCmd,
      PersonCreateInteractiveCmd,
      PersonWipeCmd,
      GetCommonRemoveQuery(reflect.ValueOf(&PersonEntity{}).Elem(), PersonActionRemove),
  }
  func PersonCliFn() cli.Command {
    PersonCliCommands = append(PersonCliCommands, PersonImportExportCommands...)
    return cli.Command{
      Name:        "person",
      Description: "Persons module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: PersonCliCommands,
    }
  }
var PERSON_ACTION_TABLE = Module2Action{
  Name:    "table",
  ActionAliases: []string{"t"},
  Flags:  CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: PersonActionQuery,
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    CommonCliTableCmd2(c,
      PersonActionQuery,
      security,
      reflect.ValueOf(&PersonEntity{}).Elem(),
    )
    return nil
  },
}
var PERSON_ACTION_QUERY = Module2Action{
  Method: "GET",
  Url:    "/people",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PERSON_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpQueryEntity(c, PersonActionQuery)
    },
  },
  Format: "QUERY",
  Action: PersonActionQuery,
  ResponseEntity: &[]PersonEntity{},
  CliAction: func(c *cli.Context, security *SecurityModel) error {
		CommonCliQueryCmd2(
			c,
			PersonActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionAliases: []string{"q"},
	Flags:         CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var PERSON_ACTION_EXPORT = Module2Action{
  Method: "GET",
  Url:    "/people/export",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PERSON_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpStreamFileChannel(c, PersonActionExport)
    },
  },
  Format: "QUERY",
  Action: PersonActionExport,
  ResponseEntity: &[]PersonEntity{},
}
var PERSON_ACTION_GET_ONE = Module2Action{
  Method: "GET",
  Url:    "/person/:uniqueId",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PERSON_QUERY},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpGetEntity(c, PersonActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: PersonActionGetOne,
  ResponseEntity: &PersonEntity{},
}
var PERSON_ACTION_POST_ONE = Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new person",
  Flags: PersonCommonCliFlags,
  Method: "POST",
  Url:    "/person",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PERSON_CREATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpPostEntity(c, PersonActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    result, err := CliPostEntity(c, PersonActionCreate, security)
    HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: PersonActionCreate,
  Format: "POST_ONE",
  RequestEntity: &PersonEntity{},
  ResponseEntity: &PersonEntity{},
}
var PERSON_ACTION_PATCH = Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: PersonCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/person",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PERSON_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntity(c, PersonActionUpdate)
    },
  },
  Action: PersonActionUpdate,
  RequestEntity: &PersonEntity{},
  Format: "PATCH_ONE",
  ResponseEntity: &PersonEntity{},
}
var PERSON_ACTION_PATCH_BULK = Module2Action{
  Method: "PATCH",
  Url:    "/people",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PERSON_UPDATE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntities(c, PersonActionBulkUpdate)
    },
  },
  Action: PersonActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &BulkRecordRequest[PersonEntity]{},
  ResponseEntity: &BulkRecordRequest[PersonEntity]{},
}
var PERSON_ACTION_DELETE = Module2Action{
  Method: "DELETE",
  Url:    "/person",
  Format: "DELETE_DSL",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_PERSON_DELETE},
  },
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpRemoveEntity(c, PersonActionRemove)
    },
  },
  Action: PersonActionRemove,
  RequestEntity: &DeleteRequest{},
  ResponseEntity: &DeleteResponse{},
  TargetEntity: &PersonEntity{},
}
  /**
  *	Override this function on PersonEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendPersonRouter = func(r *[]Module2Action) {}
  func GetPersonModule2Actions() []Module2Action {
    routes := []Module2Action{
      PERSON_ACTION_QUERY,
      PERSON_ACTION_EXPORT,
      PERSON_ACTION_GET_ONE,
      PERSON_ACTION_POST_ONE,
      PERSON_ACTION_PATCH,
      PERSON_ACTION_PATCH_BULK,
      PERSON_ACTION_DELETE,
    }
    // Append user defined functions
    AppendPersonRouter(&routes)
    return routes
  }
  func CreatePersonRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetPersonModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, PersonEntityJsonSchema, "person-http", "workspaces")
    WriteEntitySchema("PersonEntity", PersonEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_PERSON_DELETE = PermissionInfo{
  CompleteKey: "root/workspaces/person/delete",
}
var PERM_ROOT_PERSON_CREATE = PermissionInfo{
  CompleteKey: "root/workspaces/person/create",
}
var PERM_ROOT_PERSON_UPDATE = PermissionInfo{
  CompleteKey: "root/workspaces/person/update",
}
var PERM_ROOT_PERSON_QUERY = PermissionInfo{
  CompleteKey: "root/workspaces/person/query",
}
var PERM_ROOT_PERSON = PermissionInfo{
  CompleteKey: "root/workspaces/person/*",
}
var ALL_PERSON_PERMISSIONS = []PermissionInfo{
	PERM_ROOT_PERSON_DELETE,
	PERM_ROOT_PERSON_CREATE,
	PERM_ROOT_PERSON_UPDATE,
	PERM_ROOT_PERSON_QUERY,
	PERM_ROOT_PERSON,
}