package workspaces

/*
*	Generated by fireback 1.1.27
*	Written by Ali Torabi.
*	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
 */
import (
	"embed"
	"encoding/json"
	"fmt"
	reflect "reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gookit/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/schollz/progressbar/v3"
	metas "github.com/torabian/fireback/modules/workspaces/metas"
	mocks "github.com/torabian/fireback/modules/workspaces/mocks/Workspace"
	queries "github.com/torabian/fireback/modules/workspaces/queries"
	seeders "github.com/torabian/fireback/modules/workspaces/seeders/Workspace"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var workspaceSeedersFs = &seeders.ViewsFs

func ResetWorkspaceSeeders(fs *embed.FS) {
	workspaceSeedersFs = fs
}

type WorkspaceEntity struct {
	Visibility       *string              `json:"visibility,omitempty" yaml:"visibility,omitempty"`
	WorkspaceId      *string              `json:"workspaceId,omitempty" yaml:"workspaceId,omitempty"`
	LinkerId         *string              `json:"linkerId,omitempty" yaml:"linkerId,omitempty"`
	ParentId         *string              `json:"parentId,omitempty" yaml:"parentId,omitempty"`
	IsDeletable      *bool                `json:"isDeletable,omitempty" yaml:"isDeletable,omitempty" gorm:"default:true"`
	IsUpdatable      *bool                `json:"isUpdatable,omitempty" yaml:"isUpdatable,omitempty" gorm:"default:true"`
	UserId           *string              `json:"userId,omitempty" yaml:"userId,omitempty"`
	Rank             int64                `json:"rank,omitempty" gorm:"type:int;name:rank"`
	ID               uint                 `gorm:"primaryKey;autoIncrement" json:"id,omitempty" yaml:"id,omitempty"`
	UniqueId         string               `json:"uniqueId,omitempty" gorm:"unique;not null;size:100;" yaml:"uniqueId,omitempty"`
	Created          int64                `json:"created,omitempty" yaml:"created,omitempty" gorm:"autoUpdateTime:nano"`
	Updated          int64                `json:"updated,omitempty" yaml:"updated,omitempty"`
	Deleted          int64                `json:"deleted,omitempty" yaml:"deleted,omitempty"`
	CreatedFormatted string               `json:"createdFormatted,omitempty" yaml:"createdFormatted,omitempty" sql:"-" gorm:"-"`
	UpdatedFormatted string               `json:"updatedFormatted,omitempty" yaml:"updatedFormatted,omitempty" sql:"-" gorm:"-"`
	Description      *string              `json:"description" yaml:"description"        `
	Name             *string              `json:"name" yaml:"name"  validate:"required"        `
	Type             *WorkspaceTypeEntity `json:"type" yaml:"type"    gorm:"foreignKey:TypeId;references:UniqueId"      `
	TypeId           *string              `json:"typeId" yaml:"typeId" validate:"required" `
	Children         []*WorkspaceEntity   `csv:"-" gorm:"-" sql:"-" json:"children,omitempty" yaml:"children,omitempty"`
	LinkedTo         *WorkspaceEntity     `csv:"-" yaml:"-" gorm:"-" json:"-" sql:"-"`
}

func WorkspaceEntityStream(q QueryDSL) (chan []*WorkspaceEntity, *QueryResultMeta, error) {
	cn := make(chan []*WorkspaceEntity)
	q.ItemsPerPage = 50
	q.StartIndex = 0
	_, qrm, err := WorkspaceActionQuery(q)
	if err != nil {
		return nil, nil, err
	}
	go func() {
		for i := 0; i <= int(qrm.TotalAvailableItems)-1; i++ {
			items, _, _ := WorkspaceActionQuery(q)
			i += q.ItemsPerPage
			q.StartIndex = i
			cn <- items
		}
	}()
	return cn, qrm, nil
}

type WorkspaceEntityList struct {
	Items []*WorkspaceEntity
}

func NewWorkspaceEntityList(items []*WorkspaceEntity) *WorkspaceEntityList {
	return &WorkspaceEntityList{
		Items: items,
	}
}
func (x *WorkspaceEntityList) ToTree() *TreeOperation[WorkspaceEntity] {
	return NewTreeOperation(
		x.Items,
		func(t *WorkspaceEntity) string {
			if t.ParentId == nil {
				return ""
			}
			return *t.ParentId
		},
		func(t *WorkspaceEntity) string {
			return t.UniqueId
		},
	)
}

var WorkspacePreloadRelations []string = []string{}
var WORKSPACE_EVENT_CREATED = "workspace.created"
var WORKSPACE_EVENT_UPDATED = "workspace.updated"
var WORKSPACE_EVENT_DELETED = "workspace.deleted"
var WORKSPACE_EVENTS = []string{
	WORKSPACE_EVENT_CREATED,
	WORKSPACE_EVENT_UPDATED,
	WORKSPACE_EVENT_DELETED,
}

type WorkspaceFieldMap struct {
	Description TranslatedString `yaml:"description"`
	Name        TranslatedString `yaml:"name"`
	Type        TranslatedString `yaml:"type"`
}

var WorkspaceEntityMetaConfig map[string]int64 = map[string]int64{}
var WorkspaceEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&WorkspaceEntity{}))

func entityWorkspaceFormatter(dto *WorkspaceEntity, query QueryDSL) {
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
func WorkspaceMockEntity() *WorkspaceEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &WorkspaceEntity{
		Description: &stringHolder,
		Name:        &stringHolder,
	}
	return entity
}
func WorkspaceActionSeederMultiple(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	batchSize := 100
	bar := progressbar.Default(int64(count))
	// Collect entities in batches
	var entitiesBatch []*WorkspaceEntity
	for i := 1; i <= count; i++ {
		entity := WorkspaceMockEntity()
		entitiesBatch = append(entitiesBatch, entity)
		// When batch size is reached, perform the batch insert
		if len(entitiesBatch) == batchSize || i == count {
			// Insert batch
			_, err := WorkspaceMultiInsert(entitiesBatch, query)
			if err == nil {
				successInsert += len(entitiesBatch)
			} else {
				fmt.Println(err)
				failureInsert += len(entitiesBatch)
			}
			// Clear the batch after insert
			entitiesBatch = nil
		}
		bar.Add(1)
	}
	fmt.Println("Success", successInsert, "Failure", failureInsert)
}
func WorkspaceActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := WorkspaceMockEntity()
		_, err := WorkspaceActionCreate(entity, query)
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
func (x *WorkspaceEntity) Seeder() string {
	obj := WorkspaceActionSeederInit()
	v, _ := json.MarshalIndent(obj, "", "  ")
	return string(v)
}
func WorkspaceActionSeederInit() *WorkspaceEntity {
	tildaRef := "~"
	_ = tildaRef
	entity := &WorkspaceEntity{
		Description: &tildaRef,
		Name:        &tildaRef,
	}
	return entity
}
func WorkspaceAssociationCreate(dto *WorkspaceEntity, query QueryDSL) error {
	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func WorkspaceRelationContentCreate(dto *WorkspaceEntity, query QueryDSL) error {
	return nil
}
func WorkspaceRelationContentUpdate(dto *WorkspaceEntity, query QueryDSL) error {
	return nil
}
func WorkspacePolyglotCreateHandler(dto *WorkspaceEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}

/**
 * This will be validating your entity fully. Important note is that, you add validate:* tag
 * in your entity, it will automatically work here. For slices inside entity, make sure you add
 * extra line of AppendSliceErrors, otherwise they won't be detected
 */
func WorkspaceValidator(dto *WorkspaceEntity, isPatch bool) *IError {
	err := CommonStructValidatorPointer(dto, isPatch)
	return err
}

// Creates a set of natural language queries, which can be used with
// AI tools to create content or help with some tasks
var WorkspaceAskCmd cli.Command = cli.Command{
	Name:  "nlp",
	Usage: "Set of natural language queries which helps creating content or data",
	Subcommands: []cli.Command{
		{
			Name:  "sample",
			Usage: "Asks for generating sample by giving an example data",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "format",
					Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json'",
					Value: "yaml",
				},
				&cli.IntFlag{
					Name:  "count",
					Usage: "How many samples to ask",
					Value: 30,
				},
			},
			Action: func(c *cli.Context) error {
				v := &WorkspaceEntity{}
				format := c.String("format")
				request := "\033[1m" + `
I need you to create me an array of exact signature as the example given below,
with at least ` + fmt.Sprint(c.String("count")) + ` items, mock the content with few words, and guess the possible values
based on the common sense. I need the output to be a valid ` + format + ` file.
Make sure you wrap the entire array in 'items' field. Also before that, I provide some explanation of each field:
Description: (type: string) Description: 
Name: (type: string) Description: 
Type: (type: one) Description: 
And here is the actual object signature:
` + v.Seeder() + `
`
				fmt.Println(request)
				return nil
			},
		},
	},
}

func WorkspaceEntityPreSanitize(dto *WorkspaceEntity, query QueryDSL) {
}
func WorkspaceEntityBeforeCreateAppend(dto *WorkspaceEntity, query QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = UUID()
	}
	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId
	WorkspaceRecursiveAddUniqueId(dto, query)
}
func WorkspaceRecursiveAddUniqueId(dto *WorkspaceEntity, query QueryDSL) {
}

/*
*

		Batch inserts, do not have all features that create
		operation does. Use it with unnormalized content,
		or read the source code carefully.
	  This is not marked as an action, because it should not be available publicly
	  at this moment.

*
*/
func WorkspaceMultiInsert(dtos []*WorkspaceEntity, query QueryDSL) ([]*WorkspaceEntity, *IError) {
	if len(dtos) > 0 {
		for index := range dtos {
			WorkspaceEntityPreSanitize(dtos[index], query)
			WorkspaceEntityBeforeCreateAppend(dtos[index], query)
		}
		var dbref *gorm.DB = nil
		if query.Tx == nil {
			dbref = GetDbRef()
		} else {
			dbref = query.Tx
		}
		query.Tx = dbref
		err := dbref.Create(&dtos).Error
		if err != nil {
			return nil, GormErrorToIError(err)
		}
	}
	return dtos, nil
}
func WorkspaceActionBatchCreateFn(dtos []*WorkspaceEntity, query QueryDSL) ([]*WorkspaceEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*WorkspaceEntity{}
		for _, item := range dtos {
			s, err := WorkspaceActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil
}
func WorkspaceDeleteEntireChildren(query QueryDSL, dto *WorkspaceEntity) *IError {
	// intentionally removed this. It's hard to implement it, and probably wrong without
	// proper on delete cascade
	return nil
}
func WorkspaceActionCreateFn(dto *WorkspaceEntity, query QueryDSL) (*WorkspaceEntity, *IError) {
	// 1. Validate always
	if iError := WorkspaceValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	WorkspaceEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	WorkspaceEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	WorkspacePolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	WorkspaceRelationContentCreate(dto, query)
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
	WorkspaceAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(WORKSPACE_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": GetTypeString(&WorkspaceEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})
	return dto, nil
}
func WorkspaceActionGetOne(query QueryDSL) (*WorkspaceEntity, *IError) {
	refl := reflect.ValueOf(&WorkspaceEntity{})
	item, err := GetOneEntity[WorkspaceEntity](query, refl)
	entityWorkspaceFormatter(item, query)
	return item, err
}
func WorkspaceActionGetByWorkspace(query QueryDSL) (*WorkspaceEntity, *IError) {
	refl := reflect.ValueOf(&WorkspaceEntity{})
	item, err := GetOneByWorkspaceEntity[WorkspaceEntity](query, refl)
	entityWorkspaceFormatter(item, query)
	return item, err
}
func WorkspaceActionQuery(query QueryDSL) ([]*WorkspaceEntity, *QueryResultMeta, error) {
	refl := reflect.ValueOf(&WorkspaceEntity{})
	items, meta, err := QueryEntitiesPointer[WorkspaceEntity](query, refl)
	for _, item := range items {
		entityWorkspaceFormatter(item, query)
	}
	return items, meta, err
}

var workspaceMemoryItems []*WorkspaceEntity = []*WorkspaceEntity{}

func WorkspaceEntityIntoMemory() {
	q := QueryDSL{
		ItemsPerPage: 500,
		StartIndex:   0,
	}
	_, qrm, _ := WorkspaceActionQuery(q)
	for i := 0; i <= int(qrm.TotalAvailableItems)-1; i++ {
		items, _, _ := WorkspaceActionQuery(q)
		workspaceMemoryItems = append(workspaceMemoryItems, items...)
		i += q.ItemsPerPage
		q.StartIndex = i
	}
}
func WorkspaceMemGet(id uint) *WorkspaceEntity {
	for _, item := range workspaceMemoryItems {
		if item.ID == id {
			return item
		}
	}
	return nil
}
func WorkspaceMemJoin(items []uint) []*WorkspaceEntity {
	res := []*WorkspaceEntity{}
	for _, item := range items {
		v := WorkspaceMemGet(item)
		if v != nil {
			res = append(res, v)
		}
	}
	return res
}
func (dto *WorkspaceEntity) Size() int {
	var size int = len(dto.Children)
	for _, c := range dto.Children {
		size += c.Size()
	}
	return size
}
func (dto *WorkspaceEntity) Add(nodes ...*WorkspaceEntity) bool {
	var size = dto.Size()
	for _, n := range nodes {
		if n.ParentId != nil && *n.ParentId == dto.UniqueId {
			dto.Children = append(dto.Children, n)
		} else {
			for _, c := range dto.Children {
				if c.Add(n) {
					break
				}
			}
		}
	}
	return dto.Size() == size+len(nodes)
}
func WorkspaceActionCommonPivotQuery(query QueryDSL) ([]*PivotResult, *QueryResultMeta, error) {
	items, meta, err := UnsafeQuerySqlFromFs[PivotResult](
		&queries.QueriesFs, "WorkspaceCommonPivot.sqlite.dyno", query,
	)
	return items, meta, err
}
func WorkspaceActionCteQuery(query QueryDSL) ([]*WorkspaceEntity, *QueryResultMeta, error) {
	refl := reflect.ValueOf(&WorkspaceEntity{})
	items, meta, err := ContextAwareVSqlOperation[WorkspaceEntity](
		refl, &queries.QueriesFs, "WorkspaceCte.vsql", query,
	)
	for _, item := range items {
		entityWorkspaceFormatter(item, query)
	}
	var tree []*WorkspaceEntity
	for _, item := range items {
		if item.ParentId == nil {
			root := item
			root.Add(items...)
			tree = append(tree, root)
		}
	}
	return tree, meta, err
}
func WorkspaceUpdateExec(dbref *gorm.DB, query QueryDSL, fields *WorkspaceEntity) (*WorkspaceEntity, *IError) {
	uniqueId := fields.UniqueId
	query.TriggerEventName = WORKSPACE_EVENT_UPDATED
	WorkspaceEntityPreSanitize(fields, query)
	var item WorkspaceEntity
	q := dbref.
		Where(&WorkspaceEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)
	err := q.UpdateColumns(fields).Error
	if err != nil {
		return nil, GormErrorToIError(err)
	}
	query.Tx = dbref
	WorkspaceRelationContentUpdate(fields, query)
	WorkspacePolyglotCreateHandler(fields, query)
	if ero := WorkspaceDeleteEntireChildren(query, fields); ero != nil {
		return nil, ero
	}
	// @meta(update has many)
	err = dbref.
		Preload(clause.Associations).
		Where(&WorkspaceEntity{UniqueId: uniqueId}).
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
func WorkspaceActionUpdateFn(query QueryDSL, fields *WorkspaceEntity) (*WorkspaceEntity, *IError) {
	if fields == nil {
		return nil, Create401Error(&WorkspacesMessages.BodyIsMissing, []string{})
	}
	// 1. Validate always
	if iError := WorkspaceValidator(fields, true); iError != nil {
		return nil, iError
	}
	// Let's not add this. I am not sure of the consequences
	// WorkspaceRecursiveAddUniqueId(fields, query)
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = GetDbRef()
		var item *WorkspaceEntity
		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			var err *IError
			item, err = WorkspaceUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}
		})
		return item, CastToIError(vf)
	} else {
		dbref = query.Tx
		return WorkspaceUpdateExec(dbref, query, fields)
	}
}

var WorkspaceWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire workspaces ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
			ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_DELETE},
		})
		count, _ := WorkspaceActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}

func WorkspaceActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&WorkspaceEntity{})
	query.ActionRequires = []PermissionInfo{PERM_ROOT_WORKSPACE_DELETE}
	return RemoveEntity[WorkspaceEntity](query, refl)
}
func WorkspaceActionWipeClean(query QueryDSL) (int64, error) {
	var err error
	var count int64 = 0
	{
		subCount, subErr := WipeCleanEntity[WorkspaceEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'WorkspaceEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
func WorkspaceActionBulkUpdate(
	query QueryDSL, dto *BulkRecordRequest[WorkspaceEntity]) (
	*BulkRecordRequest[WorkspaceEntity], *IError,
) {
	result := []*WorkspaceEntity{}
	err := GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := WorkspaceActionUpdate(query, record)
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
func (x *WorkspaceEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

var WorkspaceEntityMeta = TableMetaData{
	EntityName:    "Workspace",
	ExportKey:     "workspaces",
	TableNameInDb: "fb_workspace_entities",
	EntityObject:  &WorkspaceEntity{},
	ExportStream:  WorkspaceActionExportT,
	ImportQuery:   WorkspaceActionImport,
}

func WorkspaceActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[WorkspaceEntity](query, WorkspaceActionQuery, WorkspacePreloadRelations)
}
func WorkspaceActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[WorkspaceEntity](query, WorkspaceActionQuery, WorkspacePreloadRelations)
}
func WorkspaceActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content WorkspaceEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return Create401Error(&WorkspacesMessages.InvalidContent, []string{})
	}
	json.Unmarshal(cx, &content)
	_, err := WorkspaceActionCreate(&content, query)
	return err
}

var WorkspaceCommonCliFlags = []cli.Flag{
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
		Name:     "description",
		Required: false,
		Usage:    `description`,
	},
	&cli.StringFlag{
		Name:     "name",
		Required: true,
		Usage:    `name`,
	},
	&cli.StringFlag{
		Name:     "type-id",
		Required: true,
		Usage:    `type`,
	},
}
var WorkspaceCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:        "description",
		StructField: "Description",
		Required:    false,
		Recommended: false,
		Usage:       `description`,
		Type:        "string",
	},
	{
		Name:        "name",
		StructField: "Name",
		Required:    true,
		Recommended: false,
		Usage:       `name`,
		Type:        "string",
	},
}
var WorkspaceCommonCliFlagsOptional = []cli.Flag{
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
		Name:     "description",
		Required: false,
		Usage:    `description`,
	},
	&cli.StringFlag{
		Name:     "name",
		Required: true,
		Usage:    `name`,
	},
	&cli.StringFlag{
		Name:     "type-id",
		Required: true,
		Usage:    `type`,
	},
}
var WorkspaceCreateCmd cli.Command = WORKSPACE_ACTION_POST_ONE.ToCli()
var WorkspaceCreateInteractiveCmd cli.Command = cli.Command{
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
			ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_CREATE},
		})
		entity := &WorkspaceEntity{}
		PopulateInteractively(entity, c, WorkspaceCommonInteractiveCliFlags)
		if entity, err := WorkspaceActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := yaml.Marshal(entity)
			fmt.Println(FormatYamlKeys(string(f)))
		}
	},
}
var WorkspaceUpdateCmd cli.Command = cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Flags:   WorkspaceCommonCliFlagsOptional,
	Usage:   "Updates a template by passing the parameters",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
			ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_UPDATE},
		})
		entity := CastWorkspaceFromCli(c)
		if entity, err := WorkspaceActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
		return nil
	},
}

func (x *WorkspaceEntity) FromCli(c *cli.Context) *WorkspaceEntity {
	return CastWorkspaceFromCli(c)
}
func CastWorkspaceFromCli(c *cli.Context) *WorkspaceEntity {
	template := &WorkspaceEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
	if c.IsSet("description") {
		value := c.String("description")
		template.Description = &value
	}
	if c.IsSet("name") {
		value := c.String("name")
		template.Name = &value
	}
	if c.IsSet("type-id") {
		value := c.String("type-id")
		template.TypeId = &value
	}
	return template
}
func WorkspaceSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	SeederFromFSImport(
		QueryDSL{},
		WorkspaceActionCreate,
		reflect.ValueOf(&WorkspaceEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}
func WorkspaceSyncSeeders() {
	SeederFromFSImport(
		QueryDSL{WorkspaceId: USER_SYSTEM},
		WorkspaceActionCreate,
		reflect.ValueOf(&WorkspaceEntity{}).Elem(),
		workspaceSeedersFs,
		[]string{},
		true,
	)
}
func WorkspaceImportMocks() {
	SeederFromFSImport(
		QueryDSL{},
		WorkspaceActionCreate,
		reflect.ValueOf(&WorkspaceEntity{}).Elem(),
		&mocks.ViewsFs,
		[]string{},
		false,
	)
}
func WorkspaceWriteQueryMock(ctx MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := WorkspaceActionQuery(f)
		result := QueryEntitySuccessResult(f, items, count)
		WriteMockDataToFile(lang, "", "Workspace", result)
	}
}

var WorkspaceImportExportCommands = []cli.Command{
	{
		Name:  "mock",
		Usage: "Generates mock records based on the entity definition",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "count",
				Usage: "how many activation key do you need to be generated and stored in database",
				Value: 10,
			},
			&cli.BoolFlag{
				Name:  "batch",
				Usage: "Multiple insert into database mode. Might miss children and relations at the moment",
			},
		},
		Action: func(c *cli.Context) error {
			query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
				ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_CREATE},
			})
			if c.Bool("batch") {
				WorkspaceActionSeederMultiple(query, c.Int("count"))
			} else {
				WorkspaceActionSeeder(query, c.Int("count"))
			}
			return nil
		},
	},
	{
		Name:    "init",
		Aliases: []string{"i"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json'",
				Value: "yaml",
			},
		},
		Usage: "Creates a basic seeder file for you, based on the definition module we have. You can populate this file as an example",
		Action: func(c *cli.Context) error {
			seed := WorkspaceActionSeederInit()
			CommonInitSeeder(strings.TrimSpace(c.String("format")), seed)
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
				Value: "workspace-seeder-workspace.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of workspaces, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]WorkspaceEntity{}
			ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
	cli.Command{
		Name:  "list",
		Usage: "Prints the list of files attached to this module for syncing or bootstrapping project",
		Action: func(c *cli.Context) error {
			if entity, err := GetSeederFilenames(workspaceSeedersFs, ""); err != nil {
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
			CommonCliImportEmbedCmd(c,
				WorkspaceActionCreate,
				reflect.ValueOf(&WorkspaceEntity{}).Elem(),
				workspaceSeedersFs,
			)
			return nil
		},
	},
	cli.Command{
		Name:  "mocks",
		Usage: "Prints the list of mocks",
		Action: func(c *cli.Context) error {
			if entity, err := GetSeederFilenames(&mocks.ViewsFs, ""); err != nil {
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
			CommonCliImportEmbedCmd(c,
				WorkspaceActionCreate,
				reflect.ValueOf(&WorkspaceEntity{}).Elem(),
				&mocks.ViewsFs,
			)
			return nil
		},
	},
	cli.Command{
		Name:    "export",
		Aliases: []string{"e"},
		Flags: append(CommonQueryFlags,
			&cli.StringFlag{
				Name:     "file",
				Usage:    "The address of file you want the csv/yaml/json be exported to",
				Required: true,
			}),
		Usage: "Exports a query results into the csv/yaml/json format",
		Action: func(c *cli.Context) error {
			if strings.Contains(c.String("file"), ".csv") {
				CommonCliExportCmd2(c,
					WorkspaceEntityStream,
					reflect.ValueOf(&WorkspaceEntity{}).Elem(),
					c.String("file"),
					&metas.MetaFs,
					"WorkspaceFieldMap.yml",
					WorkspacePreloadRelations,
				)
			} else {
				CommonCliExportCmd(c,
					WorkspaceActionQuery,
					reflect.ValueOf(&WorkspaceEntity{}).Elem(),
					c.String("file"),
					&metas.MetaFs,
					"WorkspaceFieldMap.yml",
					WorkspacePreloadRelations,
				)
			}
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
			WorkspaceCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				WorkspaceActionCreate,
				reflect.ValueOf(&WorkspaceEntity{}).Elem(),
				c.String("file"),
				&SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_CREATE},
				},
				func() WorkspaceEntity {
					v := CastWorkspaceFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
var WorkspaceCliCommands []cli.Command = []cli.Command{
	WORKSPACE_ACTION_QUERY.ToCli(),
	WORKSPACE_ACTION_TABLE.ToCli(),
	WorkspaceCreateCmd,
	WorkspaceUpdateCmd,
	WorkspaceAskCmd,
	WorkspaceCreateInteractiveCmd,
	WorkspaceWipeCmd,
	GetCommonRemoveQuery(reflect.ValueOf(&WorkspaceEntity{}).Elem(), WorkspaceActionRemove),
	GetCommonCteQuery(WorkspaceActionCteQuery),
	GetCommonPivotQuery(WorkspaceActionCommonPivotQuery),
}

func WorkspaceCliFn() cli.Command {
	WorkspaceCliCommands = append(WorkspaceCliCommands, WorkspaceImportExportCommands...)
	return cli.Command{
		Name:        "ws",
		Description: "Workspaces module actions",
		Usage:       `Fireback general user role, workspaces services.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: WorkspaceCliCommands,
	}
}

var WORKSPACE_ACTION_TABLE = Module2Action{
	Name:          "table",
	ActionName:    "table",
	ActionAliases: []string{"t"},
	Flags:         CommonQueryFlags,
	Description:   "Table formatted queries all of the entities in database based on the standard query format",
	Action:        WorkspaceActionQuery,
	CliAction: func(c *cli.Context, security *SecurityModel) error {
		CommonCliTableCmd2(c,
			WorkspaceActionQuery,
			security,
			reflect.ValueOf(&WorkspaceEntity{}).Elem(),
		)
		return nil
	},
}
var WORKSPACE_ACTION_QUERY = Module2Action{
	Method: "GET",
	Url:    "/workspaces",
	SecurityModel: &SecurityModel{
		ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_QUERY},
	},
	Group: "workspace",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpQueryEntity(c, WorkspaceActionQuery)
		},
	},
	Format:         "QUERY",
	Action:         WorkspaceActionQuery,
	ResponseEntity: &[]WorkspaceEntity{},
	Out: &Module2ActionBody{
		Entity: "WorkspaceEntity",
	},
	CliAction: func(c *cli.Context, security *SecurityModel) error {
		CommonCliQueryCmd2(
			c,
			WorkspaceActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionName:    "query",
	ActionAliases: []string{"q"},
	Flags:         CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var WORKSPACE_ACTION_QUERY_CTE = Module2Action{
	Method: "GET",
	Url:    "/cte-workspaces",
	SecurityModel: &SecurityModel{
		ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_QUERY},
	},
	Group: "workspace",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpQueryEntity(c, WorkspaceActionCteQuery)
		},
	},
	Format:         "QUERY",
	Action:         WorkspaceActionCteQuery,
	ResponseEntity: &[]WorkspaceEntity{},
	Out: &Module2ActionBody{
		Entity: "WorkspaceEntity",
	},
}
var WORKSPACE_ACTION_EXPORT = Module2Action{
	Method: "GET",
	Url:    "/workspaces/export",
	SecurityModel: &SecurityModel{
		ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_QUERY},
	},
	Group: "workspace",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpStreamFileChannel(c, WorkspaceActionExport)
		},
	},
	Format:         "QUERY",
	Action:         WorkspaceActionExport,
	ResponseEntity: &[]WorkspaceEntity{},
	Out: &Module2ActionBody{
		Entity: "WorkspaceEntity",
	},
}
var WORKSPACE_ACTION_GET_ONE = Module2Action{
	Method: "GET",
	Url:    "/workspace/:uniqueId",
	SecurityModel: &SecurityModel{
		ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_QUERY},
	},
	Group: "workspace",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpGetEntity(c, WorkspaceActionGetOne)
		},
	},
	Format:         "GET_ONE",
	Action:         WorkspaceActionGetOne,
	ResponseEntity: &WorkspaceEntity{},
	Out: &Module2ActionBody{
		Entity: "WorkspaceEntity",
	},
}
var WORKSPACE_ACTION_POST_ONE = Module2Action{
	ActionName:    "create",
	ActionAliases: []string{"c"},
	Description:   "Create new workspace",
	Flags:         WorkspaceCommonCliFlags,
	Method:        "POST",
	Url:           "/workspace",
	SecurityModel: &SecurityModel{
		ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_CREATE},
	},
	Group: "workspace",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpPostEntity(c, WorkspaceActionCreate)
		},
	},
	CliAction: func(c *cli.Context, security *SecurityModel) error {
		result, err := CliPostEntity(c, WorkspaceActionCreate, security)
		HandleActionInCli(c, result, err, map[string]map[string]string{})
		return err
	},
	Action:         WorkspaceActionCreate,
	Format:         "POST_ONE",
	RequestEntity:  &WorkspaceEntity{},
	ResponseEntity: &WorkspaceEntity{},
	Out: &Module2ActionBody{
		Entity: "WorkspaceEntity",
	},
	In: &Module2ActionBody{
		Entity: "WorkspaceEntity",
	},
}
var WORKSPACE_ACTION_PATCH = Module2Action{
	ActionName:    "update",
	ActionAliases: []string{"u"},
	Flags:         WorkspaceCommonCliFlagsOptional,
	Method:        "PATCH",
	Url:           "/workspace",
	SecurityModel: &SecurityModel{
		ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_UPDATE},
	},
	Group: "workspace",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpUpdateEntity(c, WorkspaceActionUpdate)
		},
	},
	Action:         WorkspaceActionUpdate,
	RequestEntity:  &WorkspaceEntity{},
	ResponseEntity: &WorkspaceEntity{},
	Format:         "PATCH_ONE",
	Out: &Module2ActionBody{
		Entity: "WorkspaceEntity",
	},
	In: &Module2ActionBody{
		Entity: "WorkspaceEntity",
	},
}
var WORKSPACE_ACTION_PATCH_BULK = Module2Action{
	Method: "PATCH",
	Url:    "/workspaces",
	SecurityModel: &SecurityModel{
		ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_UPDATE},
	},
	Group: "workspace",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpUpdateEntities(c, WorkspaceActionBulkUpdate)
		},
	},
	Action:         WorkspaceActionBulkUpdate,
	Format:         "PATCH_BULK",
	RequestEntity:  &BulkRecordRequest[WorkspaceEntity]{},
	ResponseEntity: &BulkRecordRequest[WorkspaceEntity]{},
	Out: &Module2ActionBody{
		Entity: "WorkspaceEntity",
	},
	In: &Module2ActionBody{
		Entity: "WorkspaceEntity",
	},
}
var WORKSPACE_ACTION_DELETE = Module2Action{
	Method: "DELETE",
	Url:    "/workspace",
	Format: "DELETE_DSL",
	SecurityModel: &SecurityModel{
		ActionRequires: []PermissionInfo{PERM_ROOT_WORKSPACE_DELETE},
	},
	Group: "workspace",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpRemoveEntity(c, WorkspaceActionRemove)
		},
	},
	Action:         WorkspaceActionRemove,
	RequestEntity:  &DeleteRequest{},
	ResponseEntity: &DeleteResponse{},
	TargetEntity:   &WorkspaceEntity{},
}

/**
 *	Override this function on WorkspaceEntityHttp.go,
 *	In order to add your own http
 **/
var AppendWorkspaceRouter = func(r *[]Module2Action) {}

func GetWorkspaceModule2Actions() []Module2Action {
	routes := []Module2Action{
		WORKSPACE_ACTION_QUERY_CTE,
		WORKSPACE_ACTION_QUERY,
		WORKSPACE_ACTION_EXPORT,
		WORKSPACE_ACTION_GET_ONE,
		WORKSPACE_ACTION_POST_ONE,
		WORKSPACE_ACTION_PATCH,
		WORKSPACE_ACTION_PATCH_BULK,
		WORKSPACE_ACTION_DELETE,
	}
	// Append user defined functions
	AppendWorkspaceRouter(&routes)
	return routes
}

var PERM_ROOT_WORKSPACE_DELETE = PermissionInfo{
	CompleteKey: "root/modules/workspaces/workspace/delete",
	Name:        "Delete workspace",
}
var PERM_ROOT_WORKSPACE_CREATE = PermissionInfo{
	CompleteKey: "root/modules/workspaces/workspace/create",
	Name:        "Create workspace",
}
var PERM_ROOT_WORKSPACE_UPDATE = PermissionInfo{
	CompleteKey: "root/modules/workspaces/workspace/update",
	Name:        "Update workspace",
}
var PERM_ROOT_WORKSPACE_QUERY = PermissionInfo{
	CompleteKey: "root/modules/workspaces/workspace/query",
	Name:        "Query workspace",
}
var PERM_ROOT_WORKSPACE = PermissionInfo{
	CompleteKey: "root/modules/workspaces/workspace/*",
	Name:        "Entire workspace actions (*)",
}
var ALL_WORKSPACE_PERMISSIONS = []PermissionInfo{
	PERM_ROOT_WORKSPACE_DELETE,
	PERM_ROOT_WORKSPACE_CREATE,
	PERM_ROOT_WORKSPACE_UPDATE,
	PERM_ROOT_WORKSPACE_QUERY,
	PERM_ROOT_WORKSPACE,
}
var WorkspaceEntityBundle = EntityBundle{
	Permissions: ALL_WORKSPACE_PERMISSIONS,
	// Cli command has been exluded, since we use module to wrap all the entities
	// to be more easier to wrap up.
	// Create your own bundle if you need with Cli
	//CliCommands: []cli.Command{
	//	WorkspaceCliFn(),
	//},
	Actions:      GetWorkspaceModule2Actions(),
	MockProvider: WorkspaceImportMocks,
	AutoMigrationEntities: []interface{}{
		&WorkspaceEntity{},
	},
}
