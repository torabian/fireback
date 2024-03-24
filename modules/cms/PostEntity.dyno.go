package cms

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	reflect "reflect"
	"strings"

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
)

type PostEntity struct {
	Visibility       *string `json:"visibility,omitempty" yaml:"visibility"`
	WorkspaceId      *string `json:"workspaceId,omitempty" yaml:"workspaceId"`
	LinkerId         *string `json:"linkerId,omitempty" yaml:"linkerId"`
	ParentId         *string `json:"parentId,omitempty" yaml:"parentId"`
	UniqueId         string  `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
	UserId           *string `json:"userId,omitempty" yaml:"userId"`
	Rank             int64   `json:"rank,omitempty" gorm:"type:int;name:rank"`
	Updated          int64   `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
	Created          int64   `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
	CreatedFormatted string  `json:"createdFormatted,omitempty" sql:"-" gorm:"-"`
	UpdatedFormatted string  `json:"updatedFormatted,omitempty" sql:"-" gorm:"-"`
	Title            *string `json:"title" yaml:"title"       `
	// Datenano also has a text representation
	Content *string `json:"content" yaml:"content"       `
	// Datenano also has a text representation
	ContentExcerpt *string             `json:"contentExcerpt" yaml:"contentExcerpt"`
	Category       *PostCategoryEntity `json:"category" yaml:"category"    gorm:"foreignKey:CategoryId;references:UniqueId"     `
	// Datenano also has a text representation
	CategoryId *string          `json:"categoryId" yaml:"categoryId"`
	Tags       []*PostTagEntity `json:"tags" yaml:"tags"    gorm:"many2many:post_tags;foreignKey:UniqueId;references:UniqueId"     `
	// Datenano also has a text representation
	TagsListId []string      `json:"tagsListId" yaml:"tagsListId" gorm:"-" sql:"-"`
	Children   []*PostEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
	LinkedTo   *PostEntity   `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

var PostPreloadRelations []string = []string{}
var POST_EVENT_CREATED = "post.created"
var POST_EVENT_UPDATED = "post.updated"
var POST_EVENT_DELETED = "post.deleted"
var POST_EVENTS = []string{
	POST_EVENT_CREATED,
	POST_EVENT_UPDATED,
	POST_EVENT_DELETED,
}

type PostFieldMap struct {
	Title    workspaces.TranslatedString `yaml:"title"`
	Content  workspaces.TranslatedString `yaml:"content"`
	Category workspaces.TranslatedString `yaml:"category"`
	Tags     workspaces.TranslatedString `yaml:"tags"`
}

var PostEntityMetaConfig map[string]int64 = map[string]int64{
	"ContentExcerptSize": 100,
}
var PostEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&PostEntity{}))

func entityPostFormatter(dto *PostEntity, query workspaces.QueryDSL) {
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
func PostMockEntity() *PostEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &PostEntity{
		Title: &stringHolder,
	}
	return entity
}
func PostActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := PostMockEntity()
		_, err := PostActionCreate(entity, query)
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
func PostActionSeederInit(query workspaces.QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*PostEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &PostEntity{
		Title:      &tildaRef,
		TagsListId: []string{"~"},
		Tags:       []*PostTagEntity{{}},
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
func PostAssociationCreate(dto *PostEntity, query workspaces.QueryDSL) error {
	{
		if dto.TagsListId != nil && len(dto.TagsListId) > 0 {
			var items []PostTagEntity
			err := query.Tx.Where(dto.TagsListId).Find(&items).Error
			if err != nil {
				return err
			}
			err = query.Tx.Model(dto).Association("Tags").Replace(items)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func PostRelationContentCreate(dto *PostEntity, query workspaces.QueryDSL) error {
	return nil
}
func PostRelationContentUpdate(dto *PostEntity, query workspaces.QueryDSL) error {
	return nil
}
func PostPolyglotCreateHandler(dto *PostEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
}

/**
 * This will be validating your entity fully. Important note is that, you add validate:* tag
 * in your entity, it will automatically work here. For slices inside entity, make sure you add
 * extra line of AppendSliceErrors, otherwise they won't be detected
 */
func PostValidator(dto *PostEntity, isPatch bool) *workspaces.IError {
	err := workspaces.CommonStructValidatorPointer(dto, isPatch)
	return err
}
func PostEntityPreSanitize(dto *PostEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
	if dto.Content != nil {
		Content := *dto.Content
		ContentExcerpt := stripPolicy.Sanitize(*dto.Content)
		Content = ugcPolicy.Sanitize(Content)
		ContentExcerpt = stripPolicy.Sanitize(ContentExcerpt)
		ContentExcerptSize, ContentExcerptSizeExists := PostEntityMetaConfig["ContentExcerptSize"]
		if ContentExcerptSizeExists {
			ContentExcerpt = workspaces.PickFirstNWords(ContentExcerpt, int(ContentExcerptSize))
		} else {
			ContentExcerpt = workspaces.PickFirstNWords(ContentExcerpt, 30)
		}
		dto.ContentExcerpt = &ContentExcerpt
		dto.Content = &Content
	}
}
func PostEntityBeforeCreateAppend(dto *PostEntity, query workspaces.QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = workspaces.UUID()
	}
	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId
	PostRecursiveAddUniqueId(dto, query)
}
func PostRecursiveAddUniqueId(dto *PostEntity, query workspaces.QueryDSL) {
}
func PostActionBatchCreateFn(dtos []*PostEntity, query workspaces.QueryDSL) ([]*PostEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*PostEntity{}
		for _, item := range dtos {
			s, err := PostActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil
}
func PostDeleteEntireChildren(query workspaces.QueryDSL, dto *PostEntity) *workspaces.IError {
	return nil
}
func PostActionCreateFn(dto *PostEntity, query workspaces.QueryDSL) (*PostEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := PostValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	PostEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	PostEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	PostPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	PostRelationContentCreate(dto, query)
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
	PostAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(POST_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": workspaces.GetTypeString(&PostEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})
	return dto, nil
}
func PostActionGetOne(query workspaces.QueryDSL) (*PostEntity, *workspaces.IError) {
	refl := reflect.ValueOf(&PostEntity{})
	item, err := workspaces.GetOneEntity[PostEntity](query, refl)
	entityPostFormatter(item, query)
	return item, err
}
func PostActionQuery(query workspaces.QueryDSL) ([]*PostEntity, *workspaces.QueryResultMeta, error) {
	refl := reflect.ValueOf(&PostEntity{})
	items, meta, err := workspaces.QueryEntitiesPointer[PostEntity](query, refl)
	for _, item := range items {
		entityPostFormatter(item, query)
	}
	return items, meta, err
}
func PostUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *PostEntity) (*PostEntity, *workspaces.IError) {
	uniqueId := fields.UniqueId
	query.TriggerEventName = POST_EVENT_UPDATED
	PostEntityPreSanitize(fields, query)
	var item PostEntity
	q := dbref.
		Where(&PostEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)
	err := q.UpdateColumns(fields).Error
	if err != nil {
		return nil, workspaces.GormErrorToIError(err)
	}
	query.Tx = dbref
	PostRelationContentUpdate(fields, query)
	PostPolyglotCreateHandler(fields, query)
	if ero := PostDeleteEntireChildren(query, fields); ero != nil {
		return nil, ero
	}
	// @meta(update has many)
	if fields.TagsListId != nil {
		var items []PostTagEntity
		if len(fields.TagsListId) > 0 {
			dbref.
				Where(&fields.TagsListId).
				Find(&items)
		}
		dbref.
			Model(&PostEntity{UniqueId: uniqueId}).
			Association("Tags").
			Replace(&items)
	}
	err = dbref.
		Preload(clause.Associations).
		Where(&PostEntity{UniqueId: uniqueId}).
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
func PostActionUpdateFn(query workspaces.QueryDSL, fields *PostEntity) (*PostEntity, *workspaces.IError) {
	if fields == nil {
		return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}
	// 1. Validate always
	if iError := PostValidator(fields, true); iError != nil {
		return nil, iError
	}
	// Let's not add this. I am not sure of the consequences
	// PostRecursiveAddUniqueId(fields, query)
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()
		var item *PostEntity
		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			var err *workspaces.IError
			item, err = PostUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}
		})
		return item, workspaces.CastToIError(vf)
	} else {
		dbref = query.Tx
		return PostUpdateExec(dbref, query, fields)
	}
}

var PostWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire posts ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
			ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_DELETE},
		})
		count, _ := PostActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}

func PostActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&PostEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_POST_DELETE}
	return workspaces.RemoveEntity[PostEntity](query, refl)
}
func PostActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error
	var count int64 = 0
	{
		subCount, subErr := workspaces.WipeCleanEntity[PostEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'PostEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
func PostActionBulkUpdate(
	query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[PostEntity]) (
	*workspaces.BulkRecordRequest[PostEntity], *workspaces.IError,
) {
	result := []*PostEntity{}
	err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := PostActionUpdate(query, record)
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
func (x *PostEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

var PostEntityMeta = workspaces.TableMetaData{
	EntityName:    "Post",
	ExportKey:     "posts",
	TableNameInDb: "fb_post_entities",
	EntityObject:  &PostEntity{},
	ExportStream:  PostActionExportT,
	ImportQuery:   PostActionImport,
}

func PostActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[PostEntity](query, PostActionQuery, PostPreloadRelations)
}
func PostActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[PostEntity](query, PostActionQuery, PostPreloadRelations)
}
func PostActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content PostEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := PostActionCreate(&content, query)
	return err
}

var PostCommonCliFlags = []cli.Flag{
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
		Name:     "title",
		Required: false,
		Usage:    "title",
	},
	&cli.StringFlag{
		Name:     "content",
		Required: false,
		Usage:    "content",
	},
	&cli.StringFlag{
		Name:     "category-id",
		Required: false,
		Usage:    "category",
	},
	&cli.StringSliceFlag{
		Name:     "tags",
		Required: false,
		Usage:    "tags",
	},
}
var PostCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:        "title",
		StructField: "Title",
		Required:    false,
		Usage:       "title",
		Type:        "string",
	},
}
var PostCommonCliFlagsOptional = []cli.Flag{
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
		Name:     "title",
		Required: false,
		Usage:    "title",
	},
	&cli.StringFlag{
		Name:     "content",
		Required: false,
		Usage:    "content",
	},
	&cli.StringFlag{
		Name:     "category-id",
		Required: false,
		Usage:    "category",
	},
	&cli.StringSliceFlag{
		Name:     "tags",
		Required: false,
		Usage:    "tags",
	},
}
var PostCreateCmd cli.Command = POST_ACTION_POST_ONE.ToCli()
var PostCreateInteractiveCmd cli.Command = cli.Command{
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
			ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_CREATE},
		})
		entity := &PostEntity{}
		for _, item := range PostCommonInteractiveCliFlags {
			if !item.Required && c.Bool("all") == false {
				continue
			}
			result := workspaces.AskForInput(item.Name, "")
			workspaces.SetFieldString(entity, item.StructField, result)
		}
		if entity, err := PostActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var PostUpdateCmd cli.Command = cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Flags:   PostCommonCliFlagsOptional,
	Usage:   "Updates a template by passing the parameters",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
			ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_UPDATE},
		})
		entity := CastPostFromCli(c)
		if entity, err := PostActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
		return nil
	},
}

func (x *PostEntity) FromCli(c *cli.Context) *PostEntity {
	return CastPostFromCli(c)
}
func CastPostFromCli(c *cli.Context) *PostEntity {
	template := &PostEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
	if c.IsSet("title") {
		value := c.String("title")
		template.Title = &value
	}
	if c.IsSet("content") {
		value := c.String("content")
		template.Content = &value
	}
	if c.IsSet("category-id") {
		value := c.String("category-id")
		template.CategoryId = &value
	}
	if c.IsSet("tags") {
		value := c.String("tags")
		template.TagsListId = strings.Split(value, ",")
	}
	return template
}
func PostSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		PostActionCreate,
		reflect.ValueOf(&PostEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}
func PostWriteQueryMock(ctx workspaces.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := PostActionQuery(f)
		result := workspaces.QueryEntitySuccessResult(f, items, count)
		workspaces.WriteMockDataToFile(lang, "", "Post", result)
	}
}

var PostImportExportCommands = []cli.Command{
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
				ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_CREATE},
			})
			PostActionSeeder(query, c.Int("count"))
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
				Value: "post-seeder.yml",
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
				ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_CREATE},
			})
			PostActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "post-seeder-post.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of posts, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]PostEntity{}
			workspaces.ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
	cli.Command{
		Name: "import",
		Flags: append(
			append(
				workspaces.CommonQueryFlags,
				&cli.StringFlag{
					Name:     "file",
					Usage:    "The address of file you want the csv be imported from",
					Required: true,
				}),
			PostCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				PostActionCreate,
				reflect.ValueOf(&PostEntity{}).Elem(),
				c.String("file"),
				&workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_CREATE},
				},
				func() PostEntity {
					v := CastPostFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
var PostCliCommands []cli.Command = []cli.Command{
	workspaces.GetCommonQuery2(PostActionQuery, &workspaces.SecurityModel{
		ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_CREATE},
	}),
	workspaces.GetCommonTableQuery(reflect.ValueOf(&PostEntity{}).Elem(), PostActionQuery),
	PostCreateCmd,
	PostUpdateCmd,
	PostCreateInteractiveCmd,
	PostWipeCmd,
	workspaces.GetCommonRemoveQuery(reflect.ValueOf(&PostEntity{}).Elem(), PostActionRemove),
}

func PostCliFn() cli.Command {
	PostCliCommands = append(PostCliCommands, PostImportExportCommands...)
	return cli.Command{
		Name:        "post",
		Description: "Posts module actions (sample module to handle complex entities)",
		Usage:       "",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: PostCliCommands,
	}
}

var POST_ACTION_POST_ONE = workspaces.Module2Action{
	ActionName:    "create",
	ActionAliases: []string{"c"},
	Description:   "Create new post",
	Flags:         PostCommonCliFlags,
	Method:        "POST",
	Url:           "/post",
	SecurityModel: &workspaces.SecurityModel{
		ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_CREATE},
	},
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			workspaces.HttpPostEntity(c, PostActionCreate)
		},
	},
	CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		result, err := workspaces.CliPostEntity(c, PostActionCreate, security)
		workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
		return err
	},
	Action:         PostActionCreate,
	Format:         "POST_ONE",
	RequestEntity:  &PostEntity{},
	ResponseEntity: &PostEntity{},
}

/**
 *	Override this function on PostEntityHttp.go,
 *	In order to add your own http
 **/
var AppendPostRouter = func(r *[]workspaces.Module2Action) {}

func GetPostModule2Actions() []workspaces.Module2Action {
	routes := []workspaces.Module2Action{
		{
			Method: "GET",
			Url:    "/posts",
			SecurityModel: &workspaces.SecurityModel{
				ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpQueryEntity(c, PostActionQuery)
				},
			},
			Format:         "QUERY",
			Action:         PostActionQuery,
			ResponseEntity: &[]PostEntity{},
		},
		{
			Method: "GET",
			Url:    "/posts/export",
			SecurityModel: &workspaces.SecurityModel{
				ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpStreamFileChannel(c, PostActionExport)
				},
			},
			Format:         "QUERY",
			Action:         PostActionExport,
			ResponseEntity: &[]PostEntity{},
		},
		{
			Method: "GET",
			Url:    "/post/:uniqueId",
			SecurityModel: &workspaces.SecurityModel{
				ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpGetEntity(c, PostActionGetOne)
				},
			},
			Format:         "GET_ONE",
			Action:         PostActionGetOne,
			ResponseEntity: &PostEntity{},
		},
		POST_ACTION_POST_ONE,
		{
			ActionName:    "update",
			ActionAliases: []string{"u"},
			Flags:         PostCommonCliFlagsOptional,
			Method:        "PATCH",
			Url:           "/post",
			SecurityModel: &workspaces.SecurityModel{
				ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpUpdateEntity(c, PostActionUpdate)
				},
			},
			Action:         PostActionUpdate,
			RequestEntity:  &PostEntity{},
			Format:         "PATCH_ONE",
			ResponseEntity: &PostEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/posts",
			SecurityModel: &workspaces.SecurityModel{
				ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpUpdateEntities(c, PostActionBulkUpdate)
				},
			},
			Action:         PostActionBulkUpdate,
			Format:         "PATCH_BULK",
			RequestEntity:  &workspaces.BulkRecordRequest[PostEntity]{},
			ResponseEntity: &workspaces.BulkRecordRequest[PostEntity]{},
		},
		{
			Method: "DELETE",
			Url:    "/post",
			Format: "DELETE_DSL",
			SecurityModel: &workspaces.SecurityModel{
				ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_POST_DELETE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpRemoveEntity(c, PostActionRemove)
				},
			},
			Action:         PostActionRemove,
			RequestEntity:  &workspaces.DeleteRequest{},
			ResponseEntity: &workspaces.DeleteResponse{},
			TargetEntity:   &PostEntity{},
		},
	}
	// Append user defined functions
	AppendPostRouter(&routes)
	return routes
}
func CreatePostRouter(r *gin.Engine) []workspaces.Module2Action {
	httpRoutes := GetPostModule2Actions()
	workspaces.CastRoutes(httpRoutes, r)
	workspaces.WriteHttpInformationToFile(&httpRoutes, PostEntityJsonSchema, "post-http", "cms")
	workspaces.WriteEntitySchema("PostEntity", PostEntityJsonSchema, "cms")
	return httpRoutes
}

var PERM_ROOT_POST_DELETE = workspaces.PermissionInfo{
	CompleteKey: "root/cms/post/delete",
}
var PERM_ROOT_POST_CREATE = workspaces.PermissionInfo{
	CompleteKey: "root/cms/post/create",
}
var PERM_ROOT_POST_UPDATE = workspaces.PermissionInfo{
	CompleteKey: "root/cms/post/update",
}
var PERM_ROOT_POST_QUERY = workspaces.PermissionInfo{
	CompleteKey: "root/cms/post/query",
}
var PERM_ROOT_POST = workspaces.PermissionInfo{
	CompleteKey: "root/cms/post/*",
}
var ALL_POST_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_POST_DELETE,
	PERM_ROOT_POST_CREATE,
	PERM_ROOT_POST_UPDATE,
	PERM_ROOT_POST_QUERY,
	PERM_ROOT_POST,
}
