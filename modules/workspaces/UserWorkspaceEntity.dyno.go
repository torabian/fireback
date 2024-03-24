package workspaces

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
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserWorkspaceEntity struct {
	Visibility       *string     `json:"visibility,omitempty" yaml:"visibility"`
	WorkspaceId      *string     `json:"workspaceId,omitempty" yaml:"workspaceId" gorm:"index:userworkspace_idx,unique" `
	LinkerId         *string     `json:"linkerId,omitempty" yaml:"linkerId"`
	ParentId         *string     `json:"parentId,omitempty" yaml:"parentId"`
	UniqueId         string      `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
	UserId           *string     `json:"userId,omitempty" yaml:"userId" gorm:"index:userworkspace_idx,unique" `
	Rank             int64       `json:"rank,omitempty" gorm:"type:int;name:rank"`
	Updated          int64       `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
	Created          int64       `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
	CreatedFormatted string      `json:"createdFormatted,omitempty" sql:"-" gorm:"-"`
	UpdatedFormatted string      `json:"updatedFormatted,omitempty" sql:"-" gorm:"-"`
	User             *UserEntity `json:"user" yaml:"user"    gorm:"foreignKey:UserId;references:UniqueId"     `
	// Datenano also has a text representation
	Workspace *WorkspaceEntity `json:"workspace" yaml:"workspace"    gorm:"foreignKey:WorkspaceId;references:UniqueId"     `
	// Datenano also has a text representation
	Children []*UserWorkspaceEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
	LinkedTo *UserWorkspaceEntity   `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

var UserWorkspacePreloadRelations []string = []string{}
var USER_WORKSPACE_EVENT_CREATED = "userWorkspace.created"
var USER_WORKSPACE_EVENT_UPDATED = "userWorkspace.updated"
var USER_WORKSPACE_EVENT_DELETED = "userWorkspace.deleted"
var USER_WORKSPACE_EVENTS = []string{
	USER_WORKSPACE_EVENT_CREATED,
	USER_WORKSPACE_EVENT_UPDATED,
	USER_WORKSPACE_EVENT_DELETED,
}

type UserWorkspaceFieldMap struct {
	User      TranslatedString `yaml:"user"`
	Workspace TranslatedString `yaml:"workspace"`
}

var UserWorkspaceEntityMetaConfig map[string]int64 = map[string]int64{}
var UserWorkspaceEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&UserWorkspaceEntity{}))

func entityUserWorkspaceFormatter(dto *UserWorkspaceEntity, query QueryDSL) {
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
func UserWorkspaceMockEntity() *UserWorkspaceEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &UserWorkspaceEntity{}
	return entity
}
func UserWorkspaceActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := UserWorkspaceMockEntity()
		_, err := UserWorkspaceActionCreate(entity, query)
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
func UserWorkspaceActionSeederInit(query QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*UserWorkspaceEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &UserWorkspaceEntity{}
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
func UserWorkspaceAssociationCreate(dto *UserWorkspaceEntity, query QueryDSL) error {
	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func UserWorkspaceRelationContentCreate(dto *UserWorkspaceEntity, query QueryDSL) error {
	return nil
}
func UserWorkspaceRelationContentUpdate(dto *UserWorkspaceEntity, query QueryDSL) error {
	return nil
}
func UserWorkspacePolyglotCreateHandler(dto *UserWorkspaceEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}

/**
 * This will be validating your entity fully. Important note is that, you add validate:* tag
 * in your entity, it will automatically work here. For slices inside entity, make sure you add
 * extra line of AppendSliceErrors, otherwise they won't be detected
 */
func UserWorkspaceValidator(dto *UserWorkspaceEntity, isPatch bool) *IError {
	err := CommonStructValidatorPointer(dto, isPatch)
	return err
}
func UserWorkspaceEntityPreSanitize(dto *UserWorkspaceEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
func UserWorkspaceEntityBeforeCreateAppend(dto *UserWorkspaceEntity, query QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = UUID()
	}
	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId
	UserWorkspaceRecursiveAddUniqueId(dto, query)
}
func UserWorkspaceRecursiveAddUniqueId(dto *UserWorkspaceEntity, query QueryDSL) {
}
func UserWorkspaceActionBatchCreateFn(dtos []*UserWorkspaceEntity, query QueryDSL) ([]*UserWorkspaceEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*UserWorkspaceEntity{}
		for _, item := range dtos {
			s, err := UserWorkspaceActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil
}
func UserWorkspaceDeleteEntireChildren(query QueryDSL, dto *UserWorkspaceEntity) *IError {
	return nil
}
func UserWorkspaceActionCreateFn(dto *UserWorkspaceEntity, query QueryDSL) (*UserWorkspaceEntity, *IError) {
	// 1. Validate always
	if iError := UserWorkspaceValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	UserWorkspaceEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	UserWorkspaceEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	UserWorkspacePolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	UserWorkspaceRelationContentCreate(dto, query)
	// 4. Create the entity
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = GetDbRef()
	} else {
		dbref = query.Tx
	}
	query.Tx = dbref
	err := dbref.Create(&dto).Error
	if err != nil {
		err := GormErrorToIError(err)
		return dto, err
	}
	// 5. Create sub entities, objects or arrays, association to other entities
	UserWorkspaceAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(USER_WORKSPACE_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": GetTypeString(&UserWorkspaceEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})
	return dto, nil
}
func UserWorkspaceActionGetOne(query QueryDSL) (*UserWorkspaceEntity, *IError) {
	refl := reflect.ValueOf(&UserWorkspaceEntity{})
	item, err := GetOneEntity[UserWorkspaceEntity](query, refl)
	entityUserWorkspaceFormatter(item, query)
	return item, err
}
func UserWorkspaceActionQuery(query QueryDSL) ([]*UserWorkspaceEntity, *QueryResultMeta, error) {
	query.WorkspaceId = ""
	refl := reflect.ValueOf(&UserWorkspaceEntity{})
	items, meta, err := QueryEntitiesPointer[UserWorkspaceEntity](query, refl)
	for _, item := range items {
		entityUserWorkspaceFormatter(item, query)
	}
	return items, meta, err
}
func UserWorkspaceUpdateExec(dbref *gorm.DB, query QueryDSL, fields *UserWorkspaceEntity) (*UserWorkspaceEntity, *IError) {
	uniqueId := fields.UniqueId
	query.TriggerEventName = USER_WORKSPACE_EVENT_UPDATED
	UserWorkspaceEntityPreSanitize(fields, query)
	var item UserWorkspaceEntity
	q := dbref.
		Where(&UserWorkspaceEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)
	err := q.UpdateColumns(fields).Error
	if err != nil {
		return nil, GormErrorToIError(err)
	}
	query.Tx = dbref
	UserWorkspaceRelationContentUpdate(fields, query)
	UserWorkspacePolyglotCreateHandler(fields, query)
	if ero := UserWorkspaceDeleteEntireChildren(query, fields); ero != nil {
		return nil, ero
	}
	// @meta(update has many)
	err = dbref.
		Preload(clause.Associations).
		Where(&UserWorkspaceEntity{UniqueId: uniqueId}).
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
func UserWorkspaceActionUpdateFn(query QueryDSL, fields *UserWorkspaceEntity) (*UserWorkspaceEntity, *IError) {
	if fields == nil {
		return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}
	// 1. Validate always
	if iError := UserWorkspaceValidator(fields, true); iError != nil {
		return nil, iError
	}
	// Let's not add this. I am not sure of the consequences
	// UserWorkspaceRecursiveAddUniqueId(fields, query)
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = GetDbRef()
		var item *UserWorkspaceEntity
		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			var err *IError
			item, err = UserWorkspaceUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}
		})
		return item, CastToIError(vf)
	} else {
		dbref = query.Tx
		return UserWorkspaceUpdateExec(dbref, query, fields)
	}
}

var UserWorkspaceWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire userworkspaces ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
			ActionRequires: []PermissionInfo{PERM_ROOT_USER_WORKSPACE_DELETE},
		})
		count, _ := UserWorkspaceActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}

func UserWorkspaceActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&UserWorkspaceEntity{})
	query.ActionRequires = []PermissionInfo{PERM_ROOT_USER_WORKSPACE_DELETE}
	return RemoveEntity[UserWorkspaceEntity](query, refl)
}
func UserWorkspaceActionWipeClean(query QueryDSL) (int64, error) {
	var err error
	var count int64 = 0
	{
		subCount, subErr := WipeCleanEntity[UserWorkspaceEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'UserWorkspaceEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
func UserWorkspaceActionBulkUpdate(
	query QueryDSL, dto *BulkRecordRequest[UserWorkspaceEntity]) (
	*BulkRecordRequest[UserWorkspaceEntity], *IError,
) {
	result := []*UserWorkspaceEntity{}
	err := GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := UserWorkspaceActionUpdate(query, record)
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
func (x *UserWorkspaceEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

var UserWorkspaceEntityMeta = TableMetaData{
	EntityName:    "UserWorkspace",
	ExportKey:     "user-workspaces",
	TableNameInDb: "fb_user-workspace_entities",
	EntityObject:  &UserWorkspaceEntity{},
	ExportStream:  UserWorkspaceActionExportT,
	ImportQuery:   UserWorkspaceActionImport,
}

func UserWorkspaceActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[UserWorkspaceEntity](query, UserWorkspaceActionQuery, UserWorkspacePreloadRelations)
}
func UserWorkspaceActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[UserWorkspaceEntity](query, UserWorkspaceActionQuery, UserWorkspacePreloadRelations)
}
func UserWorkspaceActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content UserWorkspaceEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := UserWorkspaceActionCreate(&content, query)
	return err
}

var UserWorkspaceCommonCliFlags = []cli.Flag{
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
		Name:     "user-id",
		Required: false,
		Usage:    "user",
	},
	&cli.StringFlag{
		Name:     "workspace-id",
		Required: false,
		Usage:    "workspace",
	},
}
var UserWorkspaceCommonInteractiveCliFlags = []CliInteractiveFlag{}
var UserWorkspaceCommonCliFlagsOptional = []cli.Flag{
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
		Name:     "user-id",
		Required: false,
		Usage:    "user",
	},
	&cli.StringFlag{
		Name:     "workspace-id",
		Required: false,
		Usage:    "workspace",
	},
}
var UserWorkspaceCreateCmd cli.Command = USER_WORKSPACE_ACTION_POST_ONE.ToCli()
var UserWorkspaceCreateInteractiveCmd cli.Command = cli.Command{
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
			ActionRequires: []PermissionInfo{PERM_ROOT_USER_WORKSPACE_CREATE},
		})
		entity := &UserWorkspaceEntity{}
		for _, item := range UserWorkspaceCommonInteractiveCliFlags {
			if !item.Required && c.Bool("all") == false {
				continue
			}
			result := AskForInput(item.Name, "")
			SetFieldString(entity, item.StructField, result)
		}
		if entity, err := UserWorkspaceActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var UserWorkspaceUpdateCmd cli.Command = cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Flags:   UserWorkspaceCommonCliFlagsOptional,
	Usage:   "Updates a template by passing the parameters",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
			ActionRequires: []PermissionInfo{PERM_ROOT_USER_WORKSPACE_UPDATE},
		})
		entity := CastUserWorkspaceFromCli(c)
		if entity, err := UserWorkspaceActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
		return nil
	},
}

func (x *UserWorkspaceEntity) FromCli(c *cli.Context) *UserWorkspaceEntity {
	return CastUserWorkspaceFromCli(c)
}
func CastUserWorkspaceFromCli(c *cli.Context) *UserWorkspaceEntity {
	template := &UserWorkspaceEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
	if c.IsSet("user-id") {
		value := c.String("user-id")
		template.UserId = &value
	}
	if c.IsSet("workspace-id") {
		value := c.String("workspace-id")
		template.WorkspaceId = &value
	}
	return template
}
func UserWorkspaceSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	SeederFromFSImport(
		QueryDSL{},
		UserWorkspaceActionCreate,
		reflect.ValueOf(&UserWorkspaceEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}
func UserWorkspaceWriteQueryMock(ctx MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := UserWorkspaceActionQuery(f)
		result := QueryEntitySuccessResult(f, items, count)
		WriteMockDataToFile(lang, "", "UserWorkspace", result)
	}
}

var UserWorkspaceImportExportCommands = []cli.Command{
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
				ActionRequires: []PermissionInfo{PERM_ROOT_USER_WORKSPACE_CREATE},
			})
			UserWorkspaceActionSeeder(query, c.Int("count"))
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
				Value: "user-workspace-seeder.yml",
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
				ActionRequires: []PermissionInfo{PERM_ROOT_USER_WORKSPACE_CREATE},
			})
			UserWorkspaceActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "user-workspace-seeder-user-workspace.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of user-workspaces, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]UserWorkspaceEntity{}
			ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
	cli.Command{
		Name: "import",
		Flags: append(
			append(
				CommonQueryFlags,
				&cli.StringFlag{
					Name:     "file",
					Usage:    "The address of file you want the csv be imported from",
					Required: true,
				}),
			UserWorkspaceCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				UserWorkspaceActionCreate,
				reflect.ValueOf(&UserWorkspaceEntity{}).Elem(),
				c.String("file"),
				&SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_USER_WORKSPACE_CREATE},
				},
				func() UserWorkspaceEntity {
					v := CastUserWorkspaceFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
var UserWorkspaceCliCommands []cli.Command = []cli.Command{
	GetCommonQuery2(UserWorkspaceActionQuery, &SecurityModel{
		ActionRequires:  []PermissionInfo{PERM_ROOT_USER_WORKSPACE_QUERY},
		ResolveStrategy: ResolveStrategyUser,
	}),
	GetCommonTableQuery(reflect.ValueOf(&UserWorkspaceEntity{}).Elem(), UserWorkspaceActionQuery),
	UserWorkspaceCreateCmd,
	UserWorkspaceUpdateCmd,
	UserWorkspaceCreateInteractiveCmd,
	UserWorkspaceWipeCmd,
	GetCommonRemoveQuery(reflect.ValueOf(&UserWorkspaceEntity{}).Elem(), UserWorkspaceActionRemove),
}

func UserWorkspaceCliFn() cli.Command {
	UserWorkspaceCliCommands = append(UserWorkspaceCliCommands, UserWorkspaceImportExportCommands...)
	return cli.Command{
		Name:        "userWorkspace",
		ShortName:   "user",
		Description: "UserWorkspaces module actions (sample module to handle complex entities)",
		Usage:       "Manage the workspaces that user belongs to (either its himselves or adding by invitation)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: UserWorkspaceCliCommands,
	}
}

var USER_WORKSPACE_ACTION_POST_ONE = Module2Action{
	ActionName:    "create",
	ActionAliases: []string{"c"},
	Description:   "Create new userWorkspace",
	Flags:         UserWorkspaceCommonCliFlags,
	Method:        "POST",
	Url:           "/user-workspace",
	SecurityModel: &SecurityModel{
		ActionRequires: []PermissionInfo{PERM_ROOT_USER_WORKSPACE_CREATE},
	},
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpPostEntity(c, UserWorkspaceActionCreate)
		},
	},
	CliAction: func(c *cli.Context, security *SecurityModel) error {
		result, err := CliPostEntity(c, UserWorkspaceActionCreate, security)
		HandleActionInCli(c, result, err, map[string]map[string]string{})
		return err
	},
	Action:         UserWorkspaceActionCreate,
	Format:         "POST_ONE",
	RequestEntity:  &UserWorkspaceEntity{},
	ResponseEntity: &UserWorkspaceEntity{},
}

/**
 *	Override this function on UserWorkspaceEntityHttp.go,
 *	In order to add your own http
 **/
var AppendUserWorkspaceRouter = func(r *[]Module2Action) {}

func GetUserWorkspaceModule2Actions() []Module2Action {
	routes := []Module2Action{
		{
			Method: "GET",
			Url:    "/user-workspaces",
			SecurityModel: &SecurityModel{
				ActionRequires:  []PermissionInfo{PERM_ROOT_USER_WORKSPACE_QUERY},
				ResolveStrategy: ResolveStrategyUser,
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpQueryEntity(c, UserWorkspaceActionQuery)
				},
			},
			Format:         "QUERY",
			Action:         UserWorkspaceActionQuery,
			ResponseEntity: &[]UserWorkspaceEntity{},
		},
		{
			Method: "GET",
			Url:    "/user-workspaces/export",
			SecurityModel: &SecurityModel{
				ActionRequires: []PermissionInfo{PERM_ROOT_USER_WORKSPACE_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpStreamFileChannel(c, UserWorkspaceActionExport)
				},
			},
			Format:         "QUERY",
			Action:         UserWorkspaceActionExport,
			ResponseEntity: &[]UserWorkspaceEntity{},
		},
		{
			Method: "GET",
			Url:    "/user-workspace/:uniqueId",
			SecurityModel: &SecurityModel{
				ActionRequires: []PermissionInfo{PERM_ROOT_USER_WORKSPACE_QUERY},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpGetEntity(c, UserWorkspaceActionGetOne)
				},
			},
			Format:         "GET_ONE",
			Action:         UserWorkspaceActionGetOne,
			ResponseEntity: &UserWorkspaceEntity{},
		},
		USER_WORKSPACE_ACTION_POST_ONE,
		{
			ActionName:    "update",
			ActionAliases: []string{"u"},
			Flags:         UserWorkspaceCommonCliFlagsOptional,
			Method:        "PATCH",
			Url:           "/user-workspace",
			SecurityModel: &SecurityModel{
				ActionRequires: []PermissionInfo{PERM_ROOT_USER_WORKSPACE_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpUpdateEntity(c, UserWorkspaceActionUpdate)
				},
			},
			Action:         UserWorkspaceActionUpdate,
			RequestEntity:  &UserWorkspaceEntity{},
			Format:         "PATCH_ONE",
			ResponseEntity: &UserWorkspaceEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/user-workspaces",
			SecurityModel: &SecurityModel{
				ActionRequires: []PermissionInfo{PERM_ROOT_USER_WORKSPACE_UPDATE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpUpdateEntities(c, UserWorkspaceActionBulkUpdate)
				},
			},
			Action:         UserWorkspaceActionBulkUpdate,
			Format:         "PATCH_BULK",
			RequestEntity:  &BulkRecordRequest[UserWorkspaceEntity]{},
			ResponseEntity: &BulkRecordRequest[UserWorkspaceEntity]{},
		},
		{
			Method: "DELETE",
			Url:    "/user-workspace",
			Format: "DELETE_DSL",
			SecurityModel: &SecurityModel{
				ActionRequires: []PermissionInfo{PERM_ROOT_USER_WORKSPACE_DELETE},
			},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					HttpRemoveEntity(c, UserWorkspaceActionRemove)
				},
			},
			Action:         UserWorkspaceActionRemove,
			RequestEntity:  &DeleteRequest{},
			ResponseEntity: &DeleteResponse{},
			TargetEntity:   &UserWorkspaceEntity{},
		},
	}
	// Append user defined functions
	AppendUserWorkspaceRouter(&routes)
	return routes
}
func CreateUserWorkspaceRouter(r *gin.Engine) []Module2Action {
	httpRoutes := GetUserWorkspaceModule2Actions()
	CastRoutes(httpRoutes, r)
	WriteHttpInformationToFile(&httpRoutes, UserWorkspaceEntityJsonSchema, "user-workspace-http", "workspaces")
	WriteEntitySchema("UserWorkspaceEntity", UserWorkspaceEntityJsonSchema, "workspaces")
	return httpRoutes
}

var PERM_ROOT_USER_WORKSPACE_DELETE = PermissionInfo{
	CompleteKey: "root/workspaces/user-workspace/delete",
}
var PERM_ROOT_USER_WORKSPACE_CREATE = PermissionInfo{
	CompleteKey: "root/workspaces/user-workspace/create",
}
var PERM_ROOT_USER_WORKSPACE_UPDATE = PermissionInfo{
	CompleteKey: "root/workspaces/user-workspace/update",
}
var PERM_ROOT_USER_WORKSPACE_QUERY = PermissionInfo{
	CompleteKey: "root/workspaces/user-workspace/query",
}
var PERM_ROOT_USER_WORKSPACE = PermissionInfo{
	CompleteKey: "root/workspaces/user-workspace/*",
}
var ALL_USER_WORKSPACE_PERMISSIONS = []PermissionInfo{
	PERM_ROOT_USER_WORKSPACE_DELETE,
	PERM_ROOT_USER_WORKSPACE_CREATE,
	PERM_ROOT_USER_WORKSPACE_UPDATE,
	PERM_ROOT_USER_WORKSPACE_QUERY,
	PERM_ROOT_USER_WORKSPACE,
}
