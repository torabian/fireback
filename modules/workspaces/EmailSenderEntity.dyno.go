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
	"github.com/gin-gonic/gin"
	"github.com/gookit/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/schollz/progressbar/v3"
	metas "github.com/torabian/fireback/modules/workspaces/metas"
	mocks "github.com/torabian/fireback/modules/workspaces/mocks/EmailSender"
	seeders "github.com/torabian/fireback/modules/workspaces/seeders/EmailSender"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	reflect "reflect"
	"strings"
)

var emailSenderSeedersFs = &seeders.ViewsFs

func ResetEmailSenderSeeders(fs *embed.FS) {
	emailSenderSeedersFs = fs
}

type EmailSenderEntity struct {
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
	FromName         *string              `json:"fromName" yaml:"fromName"  validate:"required"        `
	FromEmailAddress *string              `json:"fromEmailAddress" yaml:"fromEmailAddress"  validate:"required"    gorm:"unique"      `
	ReplyTo          *string              `json:"replyTo" yaml:"replyTo"  validate:"required"        `
	NickName         *string              `json:"nickName" yaml:"nickName"  validate:"required"        `
	Children         []*EmailSenderEntity `csv:"-" gorm:"-" sql:"-" json:"children,omitempty" yaml:"children,omitempty"`
	LinkedTo         *EmailSenderEntity   `csv:"-" yaml:"-" gorm:"-" json:"-" sql:"-"`
}

func EmailSenderEntityStream(q QueryDSL) (chan []*EmailSenderEntity, *QueryResultMeta, error) {
	cn := make(chan []*EmailSenderEntity)
	q.ItemsPerPage = 50
	q.StartIndex = 0
	_, qrm, err := EmailSenderActionQuery(q)
	if err != nil {
		return nil, nil, err
	}
	go func() {
		for i := 0; i <= int(qrm.TotalAvailableItems)-1; i++ {
			items, _, _ := EmailSenderActionQuery(q)
			i += q.ItemsPerPage
			q.StartIndex = i
			cn <- items
		}
	}()
	return cn, qrm, nil
}

type EmailSenderEntityList struct {
	Items []*EmailSenderEntity
}

func NewEmailSenderEntityList(items []*EmailSenderEntity) *EmailSenderEntityList {
	return &EmailSenderEntityList{
		Items: items,
	}
}
func (x *EmailSenderEntityList) ToTree() *TreeOperation[EmailSenderEntity] {
	return NewTreeOperation(
		x.Items,
		func(t *EmailSenderEntity) string {
			if t.ParentId == nil {
				return ""
			}
			return *t.ParentId
		},
		func(t *EmailSenderEntity) string {
			return t.UniqueId
		},
	)
}

var EmailSenderPreloadRelations []string = []string{}
var EMAIL_SENDER_EVENT_CREATED = "emailSender.created"
var EMAIL_SENDER_EVENT_UPDATED = "emailSender.updated"
var EMAIL_SENDER_EVENT_DELETED = "emailSender.deleted"
var EMAIL_SENDER_EVENTS = []string{
	EMAIL_SENDER_EVENT_CREATED,
	EMAIL_SENDER_EVENT_UPDATED,
	EMAIL_SENDER_EVENT_DELETED,
}

type EmailSenderFieldMap struct {
	FromName         TranslatedString `yaml:"fromName"`
	FromEmailAddress TranslatedString `yaml:"fromEmailAddress"`
	ReplyTo          TranslatedString `yaml:"replyTo"`
	NickName         TranslatedString `yaml:"nickName"`
}

var EmailSenderEntityMetaConfig map[string]int64 = map[string]int64{}
var EmailSenderEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&EmailSenderEntity{}))

func entityEmailSenderFormatter(dto *EmailSenderEntity, query QueryDSL) {
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
func EmailSenderMockEntity() *EmailSenderEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &EmailSenderEntity{
		FromName:         &stringHolder,
		FromEmailAddress: &stringHolder,
		ReplyTo:          &stringHolder,
		NickName:         &stringHolder,
	}
	return entity
}
func EmailSenderActionSeederMultiple(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	batchSize := 100
	bar := progressbar.Default(int64(count))
	// Collect entities in batches
	var entitiesBatch []*EmailSenderEntity
	for i := 1; i <= count; i++ {
		entity := EmailSenderMockEntity()
		entitiesBatch = append(entitiesBatch, entity)
		// When batch size is reached, perform the batch insert
		if len(entitiesBatch) == batchSize || i == count {
			// Insert batch
			_, err := EmailSenderMultiInsert(entitiesBatch, query)
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
func EmailSenderActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := EmailSenderMockEntity()
		_, err := EmailSenderActionCreate(entity, query)
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
func (x *EmailSenderEntity) Seeder() string {
	obj := EmailSenderActionSeederInit()
	v, _ := json.MarshalIndent(obj, "", "  ")
	return string(v)
}
func EmailSenderActionSeederInit() *EmailSenderEntity {
	tildaRef := "~"
	_ = tildaRef
	entity := &EmailSenderEntity{
		FromName:         &tildaRef,
		FromEmailAddress: &tildaRef,
		ReplyTo:          &tildaRef,
		NickName:         &tildaRef,
	}
	return entity
}
func EmailSenderAssociationCreate(dto *EmailSenderEntity, query QueryDSL) error {
	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func EmailSenderRelationContentCreate(dto *EmailSenderEntity, query QueryDSL) error {
	return nil
}
func EmailSenderRelationContentUpdate(dto *EmailSenderEntity, query QueryDSL) error {
	return nil
}
func EmailSenderPolyglotCreateHandler(dto *EmailSenderEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}

/**
 * This will be validating your entity fully. Important note is that, you add validate:* tag
 * in your entity, it will automatically work here. For slices inside entity, make sure you add
 * extra line of AppendSliceErrors, otherwise they won't be detected
 */
func EmailSenderValidator(dto *EmailSenderEntity, isPatch bool) *IError {
	err := CommonStructValidatorPointer(dto, isPatch)
	return err
}

// Creates a set of natural language queries, which can be used with
// AI tools to create content or help with some tasks
var EmailSenderAskCmd cli.Command = cli.Command{
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
				v := &EmailSenderEntity{}
				format := c.String("format")
				request := "\033[1m" + `
I need you to create me an array of exact signature as the example given below,
with at least ` + fmt.Sprint(c.String("count")) + ` items, mock the content with few words, and guess the possible values
based on the common sense. I need the output to be a valid ` + format + ` file.
Make sure you wrap the entire array in 'items' field. Also before that, I provide some explanation of each field:
FromName: (type: string) Description: 
FromEmailAddress: (type: string) Description: 
ReplyTo: (type: string) Description: 
NickName: (type: string) Description: 
And here is the actual object signature:
` + v.Seeder() + `
`
				fmt.Println(request)
				return nil
			},
		},
	},
}

func EmailSenderEntityPreSanitize(dto *EmailSenderEntity, query QueryDSL) {
}
func EmailSenderEntityBeforeCreateAppend(dto *EmailSenderEntity, query QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = UUID()
	}
	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId
	EmailSenderRecursiveAddUniqueId(dto, query)
}
func EmailSenderRecursiveAddUniqueId(dto *EmailSenderEntity, query QueryDSL) {
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
func EmailSenderMultiInsert(dtos []*EmailSenderEntity, query QueryDSL) ([]*EmailSenderEntity, *IError) {
	if len(dtos) > 0 {
		for index := range dtos {
			EmailSenderEntityPreSanitize(dtos[index], query)
			EmailSenderEntityBeforeCreateAppend(dtos[index], query)
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
func EmailSenderActionBatchCreateFn(dtos []*EmailSenderEntity, query QueryDSL) ([]*EmailSenderEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*EmailSenderEntity{}
		for _, item := range dtos {
			s, err := EmailSenderActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil
}
func EmailSenderDeleteEntireChildren(query QueryDSL, dto *EmailSenderEntity) *IError {
	// intentionally removed this. It's hard to implement it, and probably wrong without
	// proper on delete cascade
	return nil
}
func EmailSenderActionCreateFn(dto *EmailSenderEntity, query QueryDSL) (*EmailSenderEntity, *IError) {
	// 1. Validate always
	if iError := EmailSenderValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	EmailSenderEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	EmailSenderEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	EmailSenderPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	EmailSenderRelationContentCreate(dto, query)
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
	EmailSenderAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(EMAIL_SENDER_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": GetTypeString(&EmailSenderEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})
	return dto, nil
}
func EmailSenderActionGetOne(query QueryDSL) (*EmailSenderEntity, *IError) {
	refl := reflect.ValueOf(&EmailSenderEntity{})
	item, err := GetOneEntity[EmailSenderEntity](query, refl)
	entityEmailSenderFormatter(item, query)
	return item, err
}
func EmailSenderActionGetByWorkspace(query QueryDSL) (*EmailSenderEntity, *IError) {
	refl := reflect.ValueOf(&EmailSenderEntity{})
	item, err := GetOneByWorkspaceEntity[EmailSenderEntity](query, refl)
	entityEmailSenderFormatter(item, query)
	return item, err
}
func EmailSenderActionQuery(query QueryDSL) ([]*EmailSenderEntity, *QueryResultMeta, error) {
	refl := reflect.ValueOf(&EmailSenderEntity{})
	items, meta, err := QueryEntitiesPointer[EmailSenderEntity](query, refl)
	for _, item := range items {
		entityEmailSenderFormatter(item, query)
	}
	return items, meta, err
}

var emailSenderMemoryItems []*EmailSenderEntity = []*EmailSenderEntity{}

func EmailSenderEntityIntoMemory() {
	q := QueryDSL{
		ItemsPerPage: 500,
		StartIndex:   0,
	}
	_, qrm, _ := EmailSenderActionQuery(q)
	for i := 0; i <= int(qrm.TotalAvailableItems)-1; i++ {
		items, _, _ := EmailSenderActionQuery(q)
		emailSenderMemoryItems = append(emailSenderMemoryItems, items...)
		i += q.ItemsPerPage
		q.StartIndex = i
	}
}
func EmailSenderMemGet(id uint) *EmailSenderEntity {
	for _, item := range emailSenderMemoryItems {
		if item.ID == id {
			return item
		}
	}
	return nil
}
func EmailSenderMemJoin(items []uint) []*EmailSenderEntity {
	res := []*EmailSenderEntity{}
	for _, item := range items {
		v := EmailSenderMemGet(item)
		if v != nil {
			res = append(res, v)
		}
	}
	return res
}
func EmailSenderUpdateExec(dbref *gorm.DB, query QueryDSL, fields *EmailSenderEntity) (*EmailSenderEntity, *IError) {
	uniqueId := fields.UniqueId
	query.TriggerEventName = EMAIL_SENDER_EVENT_UPDATED
	EmailSenderEntityPreSanitize(fields, query)
	var item EmailSenderEntity
	q := dbref.
		Where(&EmailSenderEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)
	err := q.UpdateColumns(fields).Error
	if err != nil {
		return nil, GormErrorToIError(err)
	}
	query.Tx = dbref
	EmailSenderRelationContentUpdate(fields, query)
	EmailSenderPolyglotCreateHandler(fields, query)
	if ero := EmailSenderDeleteEntireChildren(query, fields); ero != nil {
		return nil, ero
	}
	// @meta(update has many)
	err = dbref.
		Preload(clause.Associations).
		Where(&EmailSenderEntity{UniqueId: uniqueId}).
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
func EmailSenderActionUpdateFn(query QueryDSL, fields *EmailSenderEntity) (*EmailSenderEntity, *IError) {
	if fields == nil {
		return nil, Create401Error(&WorkspacesMessages.BodyIsMissing, []string{})
	}
	// 1. Validate always
	if iError := EmailSenderValidator(fields, true); iError != nil {
		return nil, iError
	}
	// Let's not add this. I am not sure of the consequences
	// EmailSenderRecursiveAddUniqueId(fields, query)
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = GetDbRef()
		var item *EmailSenderEntity
		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			var err *IError
			item, err = EmailSenderUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}
		})
		return item, CastToIError(vf)
	} else {
		dbref = query.Tx
		return EmailSenderUpdateExec(dbref, query, fields)
	}
}

var EmailSenderWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire emailsenders ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
			ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_SENDER_DELETE},
		})
		count, _ := EmailSenderActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}

func EmailSenderActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&EmailSenderEntity{})
	query.ActionRequires = []PermissionInfo{PERM_ROOT_EMAIL_SENDER_DELETE}
	return RemoveEntity[EmailSenderEntity](query, refl)
}
func EmailSenderActionWipeClean(query QueryDSL) (int64, error) {
	var err error
	var count int64 = 0
	{
		subCount, subErr := WipeCleanEntity[EmailSenderEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'EmailSenderEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
func EmailSenderActionBulkUpdate(
	query QueryDSL, dto *BulkRecordRequest[EmailSenderEntity]) (
	*BulkRecordRequest[EmailSenderEntity], *IError,
) {
	result := []*EmailSenderEntity{}
	err := GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := EmailSenderActionUpdate(query, record)
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
func (x *EmailSenderEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

var EmailSenderEntityMeta = TableMetaData{
	EntityName:    "EmailSender",
	ExportKey:     "email-senders",
	TableNameInDb: "fb_email-sender_entities",
	EntityObject:  &EmailSenderEntity{},
	ExportStream:  EmailSenderActionExportT,
	ImportQuery:   EmailSenderActionImport,
}

func EmailSenderActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[EmailSenderEntity](query, EmailSenderActionQuery, EmailSenderPreloadRelations)
}
func EmailSenderActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[EmailSenderEntity](query, EmailSenderActionQuery, EmailSenderPreloadRelations)
}
func EmailSenderActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content EmailSenderEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return Create401Error(&WorkspacesMessages.InvalidContent, []string{})
	}
	json.Unmarshal(cx, &content)
	_, err := EmailSenderActionCreate(&content, query)
	return err
}

var EmailSenderCommonCliFlags = []cli.Flag{
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
		Name:     "from-name",
		Required: true,
		Usage:    `fromName`,
	},
	&cli.StringFlag{
		Name:     "from-email-address",
		Required: true,
		Usage:    `fromEmailAddress`,
	},
	&cli.StringFlag{
		Name:     "reply-to",
		Required: true,
		Usage:    `replyTo`,
	},
	&cli.StringFlag{
		Name:     "nick-name",
		Required: true,
		Usage:    `nickName`,
	},
}
var EmailSenderCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:        "fromName",
		StructField: "FromName",
		Required:    true,
		Recommended: false,
		Usage:       `fromName`,
		Type:        "string",
	},
	{
		Name:        "fromEmailAddress",
		StructField: "FromEmailAddress",
		Required:    true,
		Recommended: false,
		Usage:       `fromEmailAddress`,
		Type:        "string",
	},
	{
		Name:        "replyTo",
		StructField: "ReplyTo",
		Required:    true,
		Recommended: false,
		Usage:       `replyTo`,
		Type:        "string",
	},
	{
		Name:        "nickName",
		StructField: "NickName",
		Required:    true,
		Recommended: false,
		Usage:       `nickName`,
		Type:        "string",
	},
}
var EmailSenderCommonCliFlagsOptional = []cli.Flag{
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
		Name:     "from-name",
		Required: true,
		Usage:    `fromName`,
	},
	&cli.StringFlag{
		Name:     "from-email-address",
		Required: true,
		Usage:    `fromEmailAddress`,
	},
	&cli.StringFlag{
		Name:     "reply-to",
		Required: true,
		Usage:    `replyTo`,
	},
	&cli.StringFlag{
		Name:     "nick-name",
		Required: true,
		Usage:    `nickName`,
	},
}
var EmailSenderCreateCmd cli.Command = EMAIL_SENDER_ACTION_POST_ONE.ToCli()
var EmailSenderCreateInteractiveCmd cli.Command = cli.Command{
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
			ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_SENDER_CREATE},
		})
		entity := &EmailSenderEntity{}
		PopulateInteractively(entity, c, EmailSenderCommonInteractiveCliFlags)
		if entity, err := EmailSenderActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := yaml.Marshal(entity)
			fmt.Println(FormatYamlKeys(string(f)))
		}
	},
}
var EmailSenderUpdateCmd cli.Command = cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Flags:   EmailSenderCommonCliFlagsOptional,
	Usage:   "Updates a template by passing the parameters",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
			ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_SENDER_UPDATE},
		})
		entity := CastEmailSenderFromCli(c)
		if entity, err := EmailSenderActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
		return nil
	},
}

func (x *EmailSenderEntity) FromCli(c *cli.Context) *EmailSenderEntity {
	return CastEmailSenderFromCli(c)
}
func CastEmailSenderFromCli(c *cli.Context) *EmailSenderEntity {
	template := &EmailSenderEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
	if c.IsSet("from-name") {
		value := c.String("from-name")
		template.FromName = &value
	}
	if c.IsSet("from-email-address") {
		value := c.String("from-email-address")
		template.FromEmailAddress = &value
	}
	if c.IsSet("reply-to") {
		value := c.String("reply-to")
		template.ReplyTo = &value
	}
	if c.IsSet("nick-name") {
		value := c.String("nick-name")
		template.NickName = &value
	}
	return template
}
func EmailSenderSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	SeederFromFSImport(
		QueryDSL{},
		EmailSenderActionCreate,
		reflect.ValueOf(&EmailSenderEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}
func EmailSenderSyncSeeders() {
	SeederFromFSImport(
		QueryDSL{WorkspaceId: USER_SYSTEM},
		EmailSenderActionCreate,
		reflect.ValueOf(&EmailSenderEntity{}).Elem(),
		emailSenderSeedersFs,
		[]string{},
		true,
	)
}
func EmailSenderImportMocks() {
	SeederFromFSImport(
		QueryDSL{},
		EmailSenderActionCreate,
		reflect.ValueOf(&EmailSenderEntity{}).Elem(),
		&mocks.ViewsFs,
		[]string{},
		false,
	)
}
func EmailSenderWriteQueryMock(ctx MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := EmailSenderActionQuery(f)
		result := QueryEntitySuccessResult(f, items, count)
		WriteMockDataToFile(lang, "", "EmailSender", result)
	}
}

var EmailSenderImportExportCommands = []cli.Command{
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
				ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_SENDER_CREATE},
			})
			if c.Bool("batch") {
				EmailSenderActionSeederMultiple(query, c.Int("count"))
			} else {
				EmailSenderActionSeeder(query, c.Int("count"))
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
			seed := EmailSenderActionSeederInit()
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
				Value: "email-sender-seeder-email-sender.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of email-senders, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]EmailSenderEntity{}
			ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
	cli.Command{
		Name:  "list",
		Usage: "Prints the list of files attached to this module for syncing or bootstrapping project",
		Action: func(c *cli.Context) error {
			if entity, err := GetSeederFilenames(emailSenderSeedersFs, ""); err != nil {
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
				EmailSenderActionCreate,
				reflect.ValueOf(&EmailSenderEntity{}).Elem(),
				emailSenderSeedersFs,
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
				EmailSenderActionCreate,
				reflect.ValueOf(&EmailSenderEntity{}).Elem(),
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
					EmailSenderEntityStream,
					reflect.ValueOf(&EmailSenderEntity{}).Elem(),
					c.String("file"),
					&metas.MetaFs,
					"EmailSenderFieldMap.yml",
					EmailSenderPreloadRelations,
				)
			} else {
				CommonCliExportCmd(c,
					EmailSenderActionQuery,
					reflect.ValueOf(&EmailSenderEntity{}).Elem(),
					c.String("file"),
					&metas.MetaFs,
					"EmailSenderFieldMap.yml",
					EmailSenderPreloadRelations,
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
			EmailSenderCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				EmailSenderActionCreate,
				reflect.ValueOf(&EmailSenderEntity{}).Elem(),
				c.String("file"),
				&SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_SENDER_CREATE},
				},
				func() EmailSenderEntity {
					v := CastEmailSenderFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
var EmailSenderCliCommands []cli.Command = []cli.Command{
	EMAIL_SENDER_ACTION_QUERY.ToCli(),
	EMAIL_SENDER_ACTION_TABLE.ToCli(),
	EmailSenderCreateCmd,
	EmailSenderUpdateCmd,
	EmailSenderAskCmd,
	EmailSenderCreateInteractiveCmd,
	EmailSenderWipeCmd,
	GetCommonRemoveQuery(reflect.ValueOf(&EmailSenderEntity{}).Elem(), EmailSenderActionRemove),
}

func EmailSenderCliFn() cli.Command {
	EmailSenderCliCommands = append(EmailSenderCliCommands, EmailSenderImportExportCommands...)
	return cli.Command{
		Name:        "emailsender",
		Description: "EmailSenders module actions",
		Usage:       ``,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: EmailSenderCliCommands,
	}
}

var EMAIL_SENDER_ACTION_TABLE = Module2Action{
	Name:          "table",
	ActionName:    "table",
	ActionAliases: []string{"t"},
	Flags:         CommonQueryFlags,
	Description:   "Table formatted queries all of the entities in database based on the standard query format",
	Action:        EmailSenderActionQuery,
	CliAction: func(c *cli.Context, security *SecurityModel) error {
		CommonCliTableCmd2(c,
			EmailSenderActionQuery,
			security,
			reflect.ValueOf(&EmailSenderEntity{}).Elem(),
		)
		return nil
	},
}
var EMAIL_SENDER_ACTION_QUERY = Module2Action{
	Method: "GET",
	Url:    "/email-senders",
	SecurityModel: &SecurityModel{
		ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_SENDER_QUERY},
	},
	Group: "emailSender",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpQueryEntity(c, EmailSenderActionQuery)
		},
	},
	Format:         "QUERY",
	Action:         EmailSenderActionQuery,
	ResponseEntity: &[]EmailSenderEntity{},
	Out: &Module2ActionBody{
		Entity: "EmailSenderEntity",
	},
	CliAction: func(c *cli.Context, security *SecurityModel) error {
		CommonCliQueryCmd2(
			c,
			EmailSenderActionQuery,
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
var EMAIL_SENDER_ACTION_EXPORT = Module2Action{
	Method: "GET",
	Url:    "/email-senders/export",
	SecurityModel: &SecurityModel{
		ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_SENDER_QUERY},
	},
	Group: "emailSender",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpStreamFileChannel(c, EmailSenderActionExport)
		},
	},
	Format:         "QUERY",
	Action:         EmailSenderActionExport,
	ResponseEntity: &[]EmailSenderEntity{},
	Out: &Module2ActionBody{
		Entity: "EmailSenderEntity",
	},
}
var EMAIL_SENDER_ACTION_GET_ONE = Module2Action{
	Method: "GET",
	Url:    "/email-sender/:uniqueId",
	SecurityModel: &SecurityModel{
		ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_SENDER_QUERY},
	},
	Group: "emailSender",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpGetEntity(c, EmailSenderActionGetOne)
		},
	},
	Format:         "GET_ONE",
	Action:         EmailSenderActionGetOne,
	ResponseEntity: &EmailSenderEntity{},
	Out: &Module2ActionBody{
		Entity: "EmailSenderEntity",
	},
}
var EMAIL_SENDER_ACTION_POST_ONE = Module2Action{
	ActionName:    "create",
	ActionAliases: []string{"c"},
	Description:   "Create new emailSender",
	Flags:         EmailSenderCommonCliFlags,
	Method:        "POST",
	Url:           "/email-sender",
	SecurityModel: &SecurityModel{
		ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_SENDER_CREATE},
	},
	Group: "emailSender",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpPostEntity(c, EmailSenderActionCreate)
		},
	},
	CliAction: func(c *cli.Context, security *SecurityModel) error {
		result, err := CliPostEntity(c, EmailSenderActionCreate, security)
		HandleActionInCli(c, result, err, map[string]map[string]string{})
		return err
	},
	Action:         EmailSenderActionCreate,
	Format:         "POST_ONE",
	RequestEntity:  &EmailSenderEntity{},
	ResponseEntity: &EmailSenderEntity{},
	Out: &Module2ActionBody{
		Entity: "EmailSenderEntity",
	},
	In: &Module2ActionBody{
		Entity: "EmailSenderEntity",
	},
}
var EMAIL_SENDER_ACTION_PATCH = Module2Action{
	ActionName:    "update",
	ActionAliases: []string{"u"},
	Flags:         EmailSenderCommonCliFlagsOptional,
	Method:        "PATCH",
	Url:           "/email-sender",
	SecurityModel: &SecurityModel{
		ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_SENDER_UPDATE},
	},
	Group: "emailSender",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpUpdateEntity(c, EmailSenderActionUpdate)
		},
	},
	Action:         EmailSenderActionUpdate,
	RequestEntity:  &EmailSenderEntity{},
	ResponseEntity: &EmailSenderEntity{},
	Format:         "PATCH_ONE",
	Out: &Module2ActionBody{
		Entity: "EmailSenderEntity",
	},
	In: &Module2ActionBody{
		Entity: "EmailSenderEntity",
	},
}
var EMAIL_SENDER_ACTION_PATCH_BULK = Module2Action{
	Method: "PATCH",
	Url:    "/email-senders",
	SecurityModel: &SecurityModel{
		ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_SENDER_UPDATE},
	},
	Group: "emailSender",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpUpdateEntities(c, EmailSenderActionBulkUpdate)
		},
	},
	Action:         EmailSenderActionBulkUpdate,
	Format:         "PATCH_BULK",
	RequestEntity:  &BulkRecordRequest[EmailSenderEntity]{},
	ResponseEntity: &BulkRecordRequest[EmailSenderEntity]{},
	Out: &Module2ActionBody{
		Entity: "EmailSenderEntity",
	},
	In: &Module2ActionBody{
		Entity: "EmailSenderEntity",
	},
}
var EMAIL_SENDER_ACTION_DELETE = Module2Action{
	Method: "DELETE",
	Url:    "/email-sender",
	Format: "DELETE_DSL",
	SecurityModel: &SecurityModel{
		ActionRequires: []PermissionInfo{PERM_ROOT_EMAIL_SENDER_DELETE},
	},
	Group: "emailSender",
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			HttpRemoveEntity(c, EmailSenderActionRemove)
		},
	},
	Action:         EmailSenderActionRemove,
	RequestEntity:  &DeleteRequest{},
	ResponseEntity: &DeleteResponse{},
	TargetEntity:   &EmailSenderEntity{},
}

/**
 *	Override this function on EmailSenderEntityHttp.go,
 *	In order to add your own http
 **/
var AppendEmailSenderRouter = func(r *[]Module2Action) {}

func GetEmailSenderModule2Actions() []Module2Action {
	routes := []Module2Action{
		EMAIL_SENDER_ACTION_QUERY,
		EMAIL_SENDER_ACTION_EXPORT,
		EMAIL_SENDER_ACTION_GET_ONE,
		EMAIL_SENDER_ACTION_POST_ONE,
		EMAIL_SENDER_ACTION_PATCH,
		EMAIL_SENDER_ACTION_PATCH_BULK,
		EMAIL_SENDER_ACTION_DELETE,
	}
	// Append user defined functions
	AppendEmailSenderRouter(&routes)
	return routes
}

var PERM_ROOT_EMAIL_SENDER_DELETE = PermissionInfo{
	CompleteKey: "root/modules/workspaces/email-sender/delete",
	Name:        "Delete email sender",
}
var PERM_ROOT_EMAIL_SENDER_CREATE = PermissionInfo{
	CompleteKey: "root/modules/workspaces/email-sender/create",
	Name:        "Create email sender",
}
var PERM_ROOT_EMAIL_SENDER_UPDATE = PermissionInfo{
	CompleteKey: "root/modules/workspaces/email-sender/update",
	Name:        "Update email sender",
}
var PERM_ROOT_EMAIL_SENDER_QUERY = PermissionInfo{
	CompleteKey: "root/modules/workspaces/email-sender/query",
	Name:        "Query email sender",
}
var PERM_ROOT_EMAIL_SENDER = PermissionInfo{
	CompleteKey: "root/modules/workspaces/email-sender/*",
	Name:        "Entire email sender actions (*)",
}
var ALL_EMAIL_SENDER_PERMISSIONS = []PermissionInfo{
	PERM_ROOT_EMAIL_SENDER_DELETE,
	PERM_ROOT_EMAIL_SENDER_CREATE,
	PERM_ROOT_EMAIL_SENDER_UPDATE,
	PERM_ROOT_EMAIL_SENDER_QUERY,
	PERM_ROOT_EMAIL_SENDER,
}
var EmailSenderEntityBundle = EntityBundle{
	Permissions: ALL_EMAIL_SENDER_PERMISSIONS,
	// Cli command has been exluded, since we use module to wrap all the entities
	// to be more easier to wrap up.
	// Create your own bundle if you need with Cli
	//CliCommands: []cli.Command{
	//	EmailSenderCliFn(),
	//},
	Actions:      GetEmailSenderModule2Actions(),
	MockProvider: EmailSenderImportMocks,
	AutoMigrationEntities: []interface{}{
		&EmailSenderEntity{},
	},
}
