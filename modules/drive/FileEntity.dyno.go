package drive

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/microcosm-cc/bluemonday"
	"github.com/schollz/progressbar/v3"
	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"os"
	reflect "reflect"
	"strings"
)

type FileEntity struct {
	Visibility       *string `json:"visibility,omitempty" yaml:"visibility"`
	WorkspaceId      *string `json:"workspaceId,omitempty" yaml:"workspaceId"`
	LinkerId         *string `json:"linkerId,omitempty" yaml:"linkerId"`
	ParentId         *string `json:"parentId,omitempty" yaml:"parentId"`
	UniqueId         string  `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
	UserId           *string `json:"userId,omitempty" yaml:"userId"`
	Rank             int64   `json:"rank,omitempty" gorm:"type:int;name:rank"`
	Updated          int64   `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
	Created          int64   `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
	CreatedFormatted string  `json:"createdFormatted,omitempty" sql:"-"`
	UpdatedFormatted string  `json:"updatedFormatted,omitempty" sql:"-"`
	Name             *string `json:"name" yaml:"name"       `
	// Datenano also has a text representation
	DiskPath *string `json:"diskPath" yaml:"diskPath"       `
	// Datenano also has a text representation
	Size *int64 `json:"size" yaml:"size"       `
	// Datenano also has a text representation
	VirtualPath *string `json:"virtualPath" yaml:"virtualPath"       `
	// Datenano also has a text representation
	Type *string `json:"type" yaml:"type"       `
	// Datenano also has a text representation
	Children []*FileEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
	LinkedTo *FileEntity   `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

var FilePreloadRelations []string = []string{}
var FILE_EVENT_CREATED = "file.created"
var FILE_EVENT_UPDATED = "file.updated"
var FILE_EVENT_DELETED = "file.deleted"
var FILE_EVENTS = []string{
	FILE_EVENT_CREATED,
	FILE_EVENT_UPDATED,
	FILE_EVENT_DELETED,
}

type FileFieldMap struct {
	Name        workspaces.TranslatedString `yaml:"name"`
	DiskPath    workspaces.TranslatedString `yaml:"diskPath"`
	Size        workspaces.TranslatedString `yaml:"size"`
	VirtualPath workspaces.TranslatedString `yaml:"virtualPath"`
	Type        workspaces.TranslatedString `yaml:"type"`
}

var FileEntityMetaConfig map[string]int64 = map[string]int64{}
var FileEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&FileEntity{}))

func entityFileFormatter(dto *FileEntity, query workspaces.QueryDSL) {
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
func FileMockEntity() *FileEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &FileEntity{
		Name:        &stringHolder,
		DiskPath:    &stringHolder,
		Size:        &int64Holder,
		VirtualPath: &stringHolder,
		Type:        &stringHolder,
	}
	return entity
}
func FileActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := FileMockEntity()
		_, err := FileActionCreate(entity, query)
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
func FileActionSeederInit(query workspaces.QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*FileEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &FileEntity{
		Name:        &tildaRef,
		DiskPath:    &tildaRef,
		VirtualPath: &tildaRef,
		Type:        &tildaRef,
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
func FileAssociationCreate(dto *FileEntity, query workspaces.QueryDSL) error {
	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func FileRelationContentCreate(dto *FileEntity, query workspaces.QueryDSL) error {
	return nil
}
func FileRelationContentUpdate(dto *FileEntity, query workspaces.QueryDSL) error {
	return nil
}
func FilePolyglotCreateHandler(dto *FileEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
}

/**
 * This will be validating your entity fully. Important note is that, you add validate:* tag
 * in your entity, it will automatically work here. For slices inside entity, make sure you add
 * extra line of AppendSliceErrors, otherwise they won't be detected
 */
func FileValidator(dto *FileEntity, isPatch bool) *workspaces.IError {
	err := workspaces.CommonStructValidatorPointer(dto, isPatch)
	return err
}
func FileEntityPreSanitize(dto *FileEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
func FileEntityBeforeCreateAppend(dto *FileEntity, query workspaces.QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = workspaces.UUID()
	}
	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId
	FileRecursiveAddUniqueId(dto, query)
}
func FileRecursiveAddUniqueId(dto *FileEntity, query workspaces.QueryDSL) {
}
func FileActionBatchCreateFn(dtos []*FileEntity, query workspaces.QueryDSL) ([]*FileEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*FileEntity{}
		for _, item := range dtos {
			s, err := FileActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil
}
func FileActionCreateFn(dto *FileEntity, query workspaces.QueryDSL) (*FileEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := FileValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	FileEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	FileEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	FilePolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	FileRelationContentCreate(dto, query)
	// 4. Create the entity
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()
	} else {
		dbref = query.Tx
	}
	query.Tx = dbref
	err := dbref.Create(&dto).Error
	if err != nil {
		err := workspaces.GormErrorToIError(err)
		return dto, err
	}
	// 5. Create sub entities, objects or arrays, association to other entities
	FileAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(FILE_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": workspaces.GetTypeString(&FileEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})
	return dto, nil
}
func FileActionGetOne(query workspaces.QueryDSL) (*FileEntity, *workspaces.IError) {
	refl := reflect.ValueOf(&FileEntity{})
	item, err := workspaces.GetOneEntity[FileEntity](query, refl)
	entityFileFormatter(item, query)
	return item, err
}
func FileActionQuery(query workspaces.QueryDSL) ([]*FileEntity, *workspaces.QueryResultMeta, error) {
	refl := reflect.ValueOf(&FileEntity{})
	items, meta, err := workspaces.QueryEntitiesPointer[FileEntity](query, refl)
	for _, item := range items {
		entityFileFormatter(item, query)
	}
	return items, meta, err
}
func FileUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *FileEntity) (*FileEntity, *workspaces.IError) {
	uniqueId := fields.UniqueId
	query.TriggerEventName = FILE_EVENT_UPDATED
	FileEntityPreSanitize(fields, query)
	var item FileEntity
	q := dbref.
		Where(&FileEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)
	err := q.UpdateColumns(fields).Error
	if err != nil {
		return nil, workspaces.GormErrorToIError(err)
	}
	query.Tx = dbref
	FileRelationContentUpdate(fields, query)
	FilePolyglotCreateHandler(fields, query)
	// @meta(update has many)
	err = dbref.
		Preload(clause.Associations).
		Where(&FileEntity{UniqueId: uniqueId}).
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
func FileActionUpdateFn(query workspaces.QueryDSL, fields *FileEntity) (*FileEntity, *workspaces.IError) {
	if fields == nil {
		return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}
	// 1. Validate always
	if iError := FileValidator(fields, true); iError != nil {
		return nil, iError
	}
	FileRecursiveAddUniqueId(fields, query)
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()
		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			_, err := FileUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}
		})
		return nil, workspaces.CastToIError(vf)
	} else {
		dbref = query.Tx
		return FileUpdateExec(dbref, query, fields)
	}
}

var FileWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire files ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		count, _ := FileActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}

func FileActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&FileEntity{})
	query.ActionRequires = []string{PERM_ROOT_FILE_DELETE}
	return workspaces.RemoveEntity[FileEntity](query, refl)
}
func FileActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error
	var count int64 = 0
	{
		subCount, subErr := workspaces.WipeCleanEntity[FileEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'FileEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
func FileActionBulkUpdate(
	query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[FileEntity]) (
	*workspaces.BulkRecordRequest[FileEntity], *workspaces.IError,
) {
	result := []*FileEntity{}
	err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := FileActionUpdate(query, record)
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
func (x *FileEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

var FileEntityMeta = workspaces.TableMetaData{
	EntityName:    "File",
	ExportKey:     "files",
	TableNameInDb: "fb_file_entities",
	EntityObject:  &FileEntity{},
	ExportStream:  FileActionExportT,
	ImportQuery:   FileActionImport,
}

func FileActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[FileEntity](query, FileActionQuery, FilePreloadRelations)
}
func FileActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[FileEntity](query, FileActionQuery, FilePreloadRelations)
}
func FileActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content FileEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := FileActionCreate(&content, query)
	return err
}

var FileCommonCliFlags = []cli.Flag{
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
		Name:     "disk-path",
		Required: false,
		Usage:    "diskPath",
	},
	&cli.Int64Flag{
		Name:     "size",
		Required: false,
		Usage:    "size",
	},
	&cli.StringFlag{
		Name:     "virtual-path",
		Required: false,
		Usage:    "virtualPath",
	},
	&cli.StringFlag{
		Name:     "type",
		Required: false,
		Usage:    "type",
	},
}
var FileCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:        "name",
		StructField: "Name",
		Required:    false,
		Usage:       "name",
		Type:        "string",
	},
	{
		Name:        "diskPath",
		StructField: "DiskPath",
		Required:    false,
		Usage:       "diskPath",
		Type:        "string",
	},
	{
		Name:        "size",
		StructField: "Size",
		Required:    false,
		Usage:       "size",
		Type:        "int64",
	},
	{
		Name:        "virtualPath",
		StructField: "VirtualPath",
		Required:    false,
		Usage:       "virtualPath",
		Type:        "string",
	},
	{
		Name:        "type",
		StructField: "Type",
		Required:    false,
		Usage:       "type",
		Type:        "string",
	},
}
var FileCommonCliFlagsOptional = []cli.Flag{
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
		Name:     "disk-path",
		Required: false,
		Usage:    "diskPath",
	},
	&cli.Int64Flag{
		Name:     "size",
		Required: false,
		Usage:    "size",
	},
	&cli.StringFlag{
		Name:     "virtual-path",
		Required: false,
		Usage:    "virtualPath",
	},
	&cli.StringFlag{
		Name:     "type",
		Required: false,
		Usage:    "type",
	},
}
var FileCreateCmd cli.Command = cli.Command{
	Name:    "create",
	Aliases: []string{"c"},
	Flags:   FileCommonCliFlags,
	Usage:   "Create a new template",
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastFileFromCli(c)
		if entity, err := FileActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var FileCreateInteractiveCmd cli.Command = cli.Command{
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
		entity := &FileEntity{}
		for _, item := range FileCommonInteractiveCliFlags {
			if !item.Required && c.Bool("all") == false {
				continue
			}
			result := workspaces.AskForInput(item.Name, "")
			workspaces.SetFieldString(entity, item.StructField, result)
		}
		if entity, err := FileActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var FileUpdateCmd cli.Command = cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Flags:   FileCommonCliFlagsOptional,
	Usage:   "Updates a template by passing the parameters",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastFileFromCli(c)
		if entity, err := FileActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
		return nil
	},
}

func CastFileFromCli(c *cli.Context) *FileEntity {
	template := &FileEntity{}
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
	if c.IsSet("disk-path") {
		value := c.String("disk-path")
		template.DiskPath = &value
	}
	if c.IsSet("virtual-path") {
		value := c.String("virtual-path")
		template.VirtualPath = &value
	}
	if c.IsSet("type") {
		value := c.String("type")
		template.Type = &value
	}
	return template
}
func FileSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		FileActionCreate,
		reflect.ValueOf(&FileEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}
func FileWriteQueryMock(ctx workspaces.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := FileActionQuery(f)
		result := workspaces.QueryEntitySuccessResult(f, items, count)
		workspaces.WriteMockDataToFile(lang, "", "File", result)
	}
}

var FileImportExportCommands = []cli.Command{
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
			FileActionSeeder(query, c.Int("count"))
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
				Value: "file-seeder.yml",
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
			FileActionSeederInit(f, c.String("file"), c.String("format"))
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
				Value: "file-seeder-file.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of files, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]FileEntity{}
			workspaces.ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
	cli.Command{
		Name: "import",
		Flags: append(workspaces.CommonQueryFlags,
			&cli.StringFlag{
				Name:     "file",
				Usage:    "The address of file you want the csv be imported from",
				Required: true,
			}),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmd(c,
				FileActionCreate,
				reflect.ValueOf(&FileEntity{}).Elem(),
				c.String("file"),
			)
			return nil
		},
	},
}
var FileCliCommands []cli.Command = []cli.Command{
	workspaces.GetCommonQuery(FileActionQuery),
	workspaces.GetCommonTableQuery(reflect.ValueOf(&FileEntity{}).Elem(), FileActionQuery),
	FileCreateCmd,
	FileUpdateCmd,
	FileCreateInteractiveCmd,
	FileWipeCmd,
	workspaces.GetCommonRemoveQuery(reflect.ValueOf(&FileEntity{}).Elem(), FileActionRemove),
}

func FileCliFn() cli.Command {
	FileCliCommands = append(FileCliCommands, FileImportExportCommands...)
	return cli.Command{
		Name:        "file",
		Description: "Files module actions (sample module to handle complex entities)",
		Usage:       "",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: FileCliCommands,
	}
}

/**
 *	Override this function on FileEntityHttp.go,
 *	In order to add your own http
 **/
var AppendFileRouter = func(r *[]workspaces.Module2Action) {}

func GetFileModule2Actions() []workspaces.Module2Action {
	routes := []workspaces.Module2Action{
		{
			Method: "GET",
			Url:    "/files",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FILE_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpQueryEntity(c, FileActionQuery)
				},
			},
			Format:         "QUERY",
			Action:         FileActionQuery,
			ResponseEntity: &[]FileEntity{},
		},
		{
			Method: "GET",
			Url:    "/files/export",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FILE_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpStreamFileChannel(c, FileActionExport)
				},
			},
			Format:         "QUERY",
			Action:         FileActionExport,
			ResponseEntity: &[]FileEntity{},
		},
		{
			Method: "GET",
			Url:    "/file/:uniqueId",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FILE_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpGetEntity(c, FileActionGetOne)
				},
			},
			Format:         "GET_ONE",
			Action:         FileActionGetOne,
			ResponseEntity: &FileEntity{},
		},
		{
			Method: "POST",
			Url:    "/file",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FILE_CREATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpPostEntity(c, FileActionCreate)
				},
			},
			Action:         FileActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &FileEntity{},
			ResponseEntity: &FileEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/file",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FILE_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpUpdateEntity(c, FileActionUpdate)
				},
			},
			Action:         FileActionUpdate,
			RequestEntity:  &FileEntity{},
			Format:         "PATCH_ONE",
			ResponseEntity: &FileEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/files",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FILE_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpUpdateEntities(c, FileActionBulkUpdate)
				},
			},
			Action:         FileActionBulkUpdate,
			Format:         "PATCH_BULK",
			RequestEntity:  &workspaces.BulkRecordRequest[FileEntity]{},
			ResponseEntity: &workspaces.BulkRecordRequest[FileEntity]{},
		},
		{
			Method: "DELETE",
			Url:    "/file",
			Format: "DELETE_DSL",
			SecurityModel: workspaces.SecurityModel{
				ActionRequires: []string{PERM_ROOT_FILE_DELETE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpRemoveEntity(c, FileActionRemove)
				},
			},
			Action:         FileActionRemove,
			RequestEntity:  &workspaces.DeleteRequest{},
			ResponseEntity: &workspaces.DeleteResponse{},
			TargetEntity:   &FileEntity{},
		},
	}
	// Append user defined functions
	AppendFileRouter(&routes)
	return routes
}
func CreateFileRouter(r *gin.Engine) []workspaces.Module2Action {
	httpRoutes := GetFileModule2Actions()
	workspaces.CastRoutes(httpRoutes, r)
	workspaces.WriteHttpInformationToFile(&httpRoutes, FileEntityJsonSchema, "file-http", "drive")
	workspaces.WriteEntitySchema("FileEntity", FileEntityJsonSchema, "drive")
	return httpRoutes
}

var PERM_ROOT_FILE_DELETE = "root/file/delete"
var PERM_ROOT_FILE_CREATE = "root/file/create"
var PERM_ROOT_FILE_UPDATE = "root/file/update"
var PERM_ROOT_FILE_QUERY = "root/file/query"
var PERM_ROOT_FILE = "root/file"
var ALL_FILE_PERMISSIONS = []string{
	PERM_ROOT_FILE_DELETE,
	PERM_ROOT_FILE_CREATE,
	PERM_ROOT_FILE_UPDATE,
	PERM_ROOT_FILE_QUERY,
	PERM_ROOT_FILE,
}
