package keyboardActions

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
	metas "github.com/torabian/fireback/modules/keyboardActions/metas"
	seeders "github.com/torabian/fireback/modules/keyboardActions/seeders/KeyboardShortcut"
	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type KeyboardShortcutDefaultCombination struct {
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
	AltKey           *bool   `json:"altKey" yaml:"altKey"       `
	// Datenano also has a text representation
	Key *string `json:"key" yaml:"key"       `
	// Datenano also has a text representation
	MetaKey *bool `json:"metaKey" yaml:"metaKey"       `
	// Datenano also has a text representation
	ShiftKey *bool `json:"shiftKey" yaml:"shiftKey"       `
	// Datenano also has a text representation
	CtrlKey *bool `json:"ctrlKey" yaml:"ctrlKey"       `
	// Datenano also has a text representation
	LinkedTo *KeyboardShortcutEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

func (x *KeyboardShortcutDefaultCombination) RootObjectName() string {
	return "KeyboardShortcutEntity"
}

type KeyboardShortcutUserCombination struct {
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
	AltKey           *bool   `json:"altKey" yaml:"altKey"       `
	// Datenano also has a text representation
	Key *string `json:"key" yaml:"key"       `
	// Datenano also has a text representation
	MetaKey *bool `json:"metaKey" yaml:"metaKey"       `
	// Datenano also has a text representation
	ShiftKey *bool `json:"shiftKey" yaml:"shiftKey"       `
	// Datenano also has a text representation
	CtrlKey *bool `json:"ctrlKey" yaml:"ctrlKey"       `
	// Datenano also has a text representation
	LinkedTo *KeyboardShortcutEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

func (x *KeyboardShortcutUserCombination) RootObjectName() string {
	return "KeyboardShortcutEntity"
}

type KeyboardShortcutEntity struct {
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
	Os               *string `json:"os" yaml:"os"       `
	// Datenano also has a text representation
	Host *string `json:"host" yaml:"host"       `
	// Datenano also has a text representation
	DefaultCombination *KeyboardShortcutDefaultCombination `json:"defaultCombination" yaml:"defaultCombination"    gorm:"foreignKey:LinkerId;references:UniqueId"     `
	// Datenano also has a text representation
	UserCombination *KeyboardShortcutUserCombination `json:"userCombination" yaml:"userCombination"    gorm:"foreignKey:LinkerId;references:UniqueId"     `
	// Datenano also has a text representation
	Action *string `json:"action" yaml:"action"        translate:"true" `
	// Datenano also has a text representation
	ActionKey *string `json:"actionKey" yaml:"actionKey"       `
	// Datenano also has a text representation
	Translations []*KeyboardShortcutEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
	Children     []*KeyboardShortcutEntity         `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
	LinkedTo     *KeyboardShortcutEntity           `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

var KeyboardShortcutPreloadRelations []string = []string{}
var KEYBOARD_SHORTCUT_EVENT_CREATED = "keyboardShortcut.created"
var KEYBOARD_SHORTCUT_EVENT_UPDATED = "keyboardShortcut.updated"
var KEYBOARD_SHORTCUT_EVENT_DELETED = "keyboardShortcut.deleted"
var KEYBOARD_SHORTCUT_EVENTS = []string{
	KEYBOARD_SHORTCUT_EVENT_CREATED,
	KEYBOARD_SHORTCUT_EVENT_UPDATED,
	KEYBOARD_SHORTCUT_EVENT_DELETED,
}

type KeyboardShortcutFieldMap struct {
	Os                 workspaces.TranslatedString `yaml:"os"`
	Host               workspaces.TranslatedString `yaml:"host"`
	DefaultCombination workspaces.TranslatedString `yaml:"defaultCombination"`
	UserCombination    workspaces.TranslatedString `yaml:"userCombination"`
	Action             workspaces.TranslatedString `yaml:"action"`
	ActionKey          workspaces.TranslatedString `yaml:"actionKey"`
}

var KeyboardShortcutEntityMetaConfig map[string]int64 = map[string]int64{}
var KeyboardShortcutEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&KeyboardShortcutEntity{}))

type KeyboardShortcutEntityPolyglot struct {
	LinkerId   string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
	LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
	Action     string `yaml:"action" json:"action"`
}

func KeyboardShortcutDefaultCombinationActionCreate(
	dto *KeyboardShortcutDefaultCombination,
	query workspaces.QueryDSL,
) (*KeyboardShortcutDefaultCombination, *workspaces.IError) {
	dto.LinkerId = &query.LinkerId
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()
	} else {
		dbref = query.Tx
	}
	query.Tx = dbref
	if dto.UniqueId == "" {
		dto.UniqueId = workspaces.UUID()
	}
	err := dbref.Create(&dto).Error
	if err != nil {
		err := workspaces.GormErrorToIError(err)
		return dto, err
	}
	return dto, nil
}
func KeyboardShortcutDefaultCombinationActionUpdate(
	query workspaces.QueryDSL,
	dto *KeyboardShortcutDefaultCombination,
) (*KeyboardShortcutDefaultCombination, *workspaces.IError) {
	dto.LinkerId = &query.LinkerId
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()
	} else {
		dbref = query.Tx
	}
	query.Tx = dbref
	err := dbref.UpdateColumns(&dto).Error
	if err != nil {
		err := workspaces.GormErrorToIError(err)
		return dto, err
	}
	return dto, nil
}
func KeyboardShortcutDefaultCombinationActionGetOne(
	query workspaces.QueryDSL,
) (*KeyboardShortcutDefaultCombination, *workspaces.IError) {
	refl := reflect.ValueOf(&KeyboardShortcutDefaultCombination{})
	item, err := workspaces.GetOneEntity[KeyboardShortcutDefaultCombination](query, refl)
	return item, err
}
func KeyboardShortcutUserCombinationActionCreate(
	dto *KeyboardShortcutUserCombination,
	query workspaces.QueryDSL,
) (*KeyboardShortcutUserCombination, *workspaces.IError) {
	dto.LinkerId = &query.LinkerId
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()
	} else {
		dbref = query.Tx
	}
	query.Tx = dbref
	if dto.UniqueId == "" {
		dto.UniqueId = workspaces.UUID()
	}
	err := dbref.Create(&dto).Error
	if err != nil {
		err := workspaces.GormErrorToIError(err)
		return dto, err
	}
	return dto, nil
}
func KeyboardShortcutUserCombinationActionUpdate(
	query workspaces.QueryDSL,
	dto *KeyboardShortcutUserCombination,
) (*KeyboardShortcutUserCombination, *workspaces.IError) {
	dto.LinkerId = &query.LinkerId
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()
	} else {
		dbref = query.Tx
	}
	query.Tx = dbref
	err := dbref.UpdateColumns(&dto).Error
	if err != nil {
		err := workspaces.GormErrorToIError(err)
		return dto, err
	}
	return dto, nil
}
func KeyboardShortcutUserCombinationActionGetOne(
	query workspaces.QueryDSL,
) (*KeyboardShortcutUserCombination, *workspaces.IError) {
	refl := reflect.ValueOf(&KeyboardShortcutUserCombination{})
	item, err := workspaces.GetOneEntity[KeyboardShortcutUserCombination](query, refl)
	return item, err
}
func entityKeyboardShortcutFormatter(dto *KeyboardShortcutEntity, query workspaces.QueryDSL) {
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
func KeyboardShortcutMockEntity() *KeyboardShortcutEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &KeyboardShortcutEntity{
		Os:        &stringHolder,
		Host:      &stringHolder,
		Action:    &stringHolder,
		ActionKey: &stringHolder,
	}
	return entity
}
func KeyboardShortcutActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := KeyboardShortcutMockEntity()
		_, err := KeyboardShortcutActionCreate(entity, query)
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
func (x *KeyboardShortcutEntity) GetActionTranslated(language string) string {
	if x.Translations != nil && len(x.Translations) > 0 {
		for _, item := range x.Translations {
			if item.LanguageId == language {
				return item.Action
			}
		}
	}
	return ""
}
func KeyboardShortcutActionSeederInit(query workspaces.QueryDSL, file string, format string) {
	body := []byte{}
	var err error
	data := []*KeyboardShortcutEntity{}
	tildaRef := "~"
	_ = tildaRef
	entity := &KeyboardShortcutEntity{
		Os:                 &tildaRef,
		Host:               &tildaRef,
		DefaultCombination: &KeyboardShortcutDefaultCombination{},
		UserCombination:    &KeyboardShortcutUserCombination{},
		Action:             &tildaRef,
		ActionKey:          &tildaRef,
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
func KeyboardShortcutAssociationCreate(dto *KeyboardShortcutEntity, query workspaces.QueryDSL) error {
	return nil
}

/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func KeyboardShortcutRelationContentCreate(dto *KeyboardShortcutEntity, query workspaces.QueryDSL) error {
	return nil
}
func KeyboardShortcutRelationContentUpdate(dto *KeyboardShortcutEntity, query workspaces.QueryDSL) error {
	return nil
}
func KeyboardShortcutPolyglotCreateHandler(dto *KeyboardShortcutEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
	workspaces.PolyglotCreateHandler(dto, &KeyboardShortcutEntityPolyglot{}, query)
}

/**
 * This will be validating your entity fully. Important note is that, you add validate:* tag
 * in your entity, it will automatically work here. For slices inside entity, make sure you add
 * extra line of AppendSliceErrors, otherwise they won't be detected
 */
func KeyboardShortcutValidator(dto *KeyboardShortcutEntity, isPatch bool) *workspaces.IError {
	err := workspaces.CommonStructValidatorPointer(dto, isPatch)
	return err
}
func KeyboardShortcutEntityPreSanitize(dto *KeyboardShortcutEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
func KeyboardShortcutEntityBeforeCreateAppend(dto *KeyboardShortcutEntity, query workspaces.QueryDSL) {
	if dto.UniqueId == "" {
		dto.UniqueId = workspaces.UUID()
	}
	dto.WorkspaceId = &query.WorkspaceId
	dto.UserId = &query.UserId
	KeyboardShortcutRecursiveAddUniqueId(dto, query)
}
func KeyboardShortcutRecursiveAddUniqueId(dto *KeyboardShortcutEntity, query workspaces.QueryDSL) {
	if dto.DefaultCombination != nil {
		dto.DefaultCombination.UniqueId = workspaces.UUID()
	}
	if dto.UserCombination != nil {
		dto.UserCombination.UniqueId = workspaces.UUID()
	}
}
func KeyboardShortcutActionBatchCreateFn(dtos []*KeyboardShortcutEntity, query workspaces.QueryDSL) ([]*KeyboardShortcutEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*KeyboardShortcutEntity{}
		for _, item := range dtos {
			s, err := KeyboardShortcutActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil
}
func KeyboardShortcutDeleteEntireChildren(query workspaces.QueryDSL, dto *KeyboardShortcutEntity) *workspaces.IError {
	if dto.DefaultCombination != nil {
		q := query.Tx.
			Model(&dto.DefaultCombination).
			Where(&KeyboardShortcutDefaultCombination{LinkerId: &dto.UniqueId}).
			Delete(&KeyboardShortcutDefaultCombination{})
		err := q.Error
		if err != nil {
			return workspaces.GormErrorToIError(err)
		}
	}
	if dto.UserCombination != nil {
		q := query.Tx.
			Model(&dto.UserCombination).
			Where(&KeyboardShortcutUserCombination{LinkerId: &dto.UniqueId}).
			Delete(&KeyboardShortcutUserCombination{})
		err := q.Error
		if err != nil {
			return workspaces.GormErrorToIError(err)
		}
	}
	return nil
}
func KeyboardShortcutActionCreateFn(dto *KeyboardShortcutEntity, query workspaces.QueryDSL) (*KeyboardShortcutEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := KeyboardShortcutValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	KeyboardShortcutEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	KeyboardShortcutEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	KeyboardShortcutPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	KeyboardShortcutRelationContentCreate(dto, query)
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
	KeyboardShortcutAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(KEYBOARD_SHORTCUT_EVENT_CREATED, event.M{
		"entity":    dto,
		"entityKey": workspaces.GetTypeString(&KeyboardShortcutEntity{}),
		"target":    "workspace",
		"unqiueId":  query.WorkspaceId,
	})
	return dto, nil
}
func KeyboardShortcutActionGetOne(query workspaces.QueryDSL) (*KeyboardShortcutEntity, *workspaces.IError) {
	refl := reflect.ValueOf(&KeyboardShortcutEntity{})
	item, err := workspaces.GetOneEntity[KeyboardShortcutEntity](query, refl)
	entityKeyboardShortcutFormatter(item, query)
	return item, err
}
func KeyboardShortcutActionQuery(query workspaces.QueryDSL) ([]*KeyboardShortcutEntity, *workspaces.QueryResultMeta, error) {
	refl := reflect.ValueOf(&KeyboardShortcutEntity{})
	items, meta, err := workspaces.QueryEntitiesPointer[KeyboardShortcutEntity](query, refl)
	for _, item := range items {
		entityKeyboardShortcutFormatter(item, query)
	}
	return items, meta, err
}
func KeyboardShortcutUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *KeyboardShortcutEntity) (*KeyboardShortcutEntity, *workspaces.IError) {
	uniqueId := fields.UniqueId
	query.TriggerEventName = KEYBOARD_SHORTCUT_EVENT_UPDATED
	KeyboardShortcutEntityPreSanitize(fields, query)
	var item KeyboardShortcutEntity
	q := dbref.
		Where(&KeyboardShortcutEntity{UniqueId: uniqueId}).
		FirstOrCreate(&item)
	err := q.UpdateColumns(fields).Error
	if err != nil {
		return nil, workspaces.GormErrorToIError(err)
	}
	query.Tx = dbref
	KeyboardShortcutRelationContentUpdate(fields, query)
	KeyboardShortcutPolyglotCreateHandler(fields, query)
	if ero := KeyboardShortcutDeleteEntireChildren(query, fields); ero != nil {
		return nil, ero
	}
	if fields.DefaultCombination != nil {
		linkerId := uniqueId
		q := dbref.
			Model(&item.DefaultCombination).
			Where(&KeyboardShortcutDefaultCombination{LinkerId: &linkerId}).
			UpdateColumns(fields.DefaultCombination)
		err := q.Error
		if err != nil {
			return &item, workspaces.GormErrorToIError(err)
		}
		if q.RowsAffected == 0 {
			fields.DefaultCombination.UniqueId = workspaces.UUID()
			fields.DefaultCombination.LinkerId = &linkerId
			err := dbref.
				Model(&item.DefaultCombination).Create(fields.DefaultCombination).Error
			if err != nil {
				return &item, workspaces.GormErrorToIError(err)
			}
		}
	}
	if fields.UserCombination != nil {
		linkerId := uniqueId
		q := dbref.
			Model(&item.UserCombination).
			Where(&KeyboardShortcutUserCombination{LinkerId: &linkerId}).
			UpdateColumns(fields.UserCombination)
		err := q.Error
		if err != nil {
			return &item, workspaces.GormErrorToIError(err)
		}
		if q.RowsAffected == 0 {
			fields.UserCombination.UniqueId = workspaces.UUID()
			fields.UserCombination.LinkerId = &linkerId
			err := dbref.
				Model(&item.UserCombination).Create(fields.UserCombination).Error
			if err != nil {
				return &item, workspaces.GormErrorToIError(err)
			}
		}
	}
	// @meta(update has many)
	err = dbref.
		Preload(clause.Associations).
		Where(&KeyboardShortcutEntity{UniqueId: uniqueId}).
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
func KeyboardShortcutActionUpdateFn(query workspaces.QueryDSL, fields *KeyboardShortcutEntity) (*KeyboardShortcutEntity, *workspaces.IError) {
	if fields == nil {
		return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
	}
	// 1. Validate always
	if iError := KeyboardShortcutValidator(fields, true); iError != nil {
		return nil, iError
	}
	// Let's not add this. I am not sure of the consequences
	// KeyboardShortcutRecursiveAddUniqueId(fields, query)
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()
		var item *KeyboardShortcutEntity
		vf := dbref.Transaction(func(tx *gorm.DB) error {
			dbref = tx
			var err *workspaces.IError
			item, err = KeyboardShortcutUpdateExec(dbref, query, fields)
			if err == nil {
				return nil
			} else {
				return err
			}
		})
		return item, workspaces.CastToIError(vf)
	} else {
		dbref = query.Tx
		return KeyboardShortcutUpdateExec(dbref, query, fields)
	}
}

var KeyboardShortcutWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire keyboardshortcuts ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
			ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_KEYBOARD_SHORTCUT_DELETE},
		})
		count, _ := KeyboardShortcutActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}

func KeyboardShortcutActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&KeyboardShortcutEntity{})
	query.ActionRequires = []workspaces.PermissionInfo{PERM_ROOT_KEYBOARD_SHORTCUT_DELETE}
	return workspaces.RemoveEntity[KeyboardShortcutEntity](query, refl)
}
func KeyboardShortcutActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error
	var count int64 = 0
	{
		subCount, subErr := workspaces.WipeCleanEntity[KeyboardShortcutDefaultCombination]()
		if subErr != nil {
			fmt.Println("Error while wiping 'KeyboardShortcutDefaultCombination'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	{
		subCount, subErr := workspaces.WipeCleanEntity[KeyboardShortcutUserCombination]()
		if subErr != nil {
			fmt.Println("Error while wiping 'KeyboardShortcutUserCombination'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	{
		subCount, subErr := workspaces.WipeCleanEntity[KeyboardShortcutEntity]()
		if subErr != nil {
			fmt.Println("Error while wiping 'KeyboardShortcutEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
func KeyboardShortcutActionBulkUpdate(
	query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[KeyboardShortcutEntity]) (
	*workspaces.BulkRecordRequest[KeyboardShortcutEntity], *workspaces.IError,
) {
	result := []*KeyboardShortcutEntity{}
	err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
		query.Tx = tx
		for _, record := range dto.Records {
			item, err := KeyboardShortcutActionUpdate(query, record)
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
func (x *KeyboardShortcutEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

var KeyboardShortcutEntityMeta = workspaces.TableMetaData{
	EntityName:    "KeyboardShortcut",
	ExportKey:     "keyboard-shortcuts",
	TableNameInDb: "fb_keyboard-shortcut_entities",
	EntityObject:  &KeyboardShortcutEntity{},
	ExportStream:  KeyboardShortcutActionExportT,
	ImportQuery:   KeyboardShortcutActionImport,
}

func KeyboardShortcutActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[KeyboardShortcutEntity](query, KeyboardShortcutActionQuery, KeyboardShortcutPreloadRelations)
}
func KeyboardShortcutActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[KeyboardShortcutEntity](query, KeyboardShortcutActionQuery, KeyboardShortcutPreloadRelations)
}
func KeyboardShortcutActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content KeyboardShortcutEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := KeyboardShortcutActionCreate(&content, query)
	return err
}

var KeyboardShortcutCommonCliFlags = []cli.Flag{
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
		Name:     "os",
		Required: false,
		Usage:    "os",
	},
	&cli.StringFlag{
		Name:     "host",
		Required: false,
		Usage:    "host",
	},
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
	&cli.BoolFlag{
		Name:     "alt-key",
		Required: false,
		Usage:    "altKey",
	},
	&cli.StringFlag{
		Name:     "key",
		Required: false,
		Usage:    "key",
	},
	&cli.BoolFlag{
		Name:     "meta-key",
		Required: false,
		Usage:    "metaKey",
	},
	&cli.BoolFlag{
		Name:     "shift-key",
		Required: false,
		Usage:    "shiftKey",
	},
	&cli.BoolFlag{
		Name:     "ctrl-key",
		Required: false,
		Usage:    "ctrlKey",
	},
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
	&cli.BoolFlag{
		Name:     "alt-key",
		Required: false,
		Usage:    "altKey",
	},
	&cli.StringFlag{
		Name:     "key",
		Required: false,
		Usage:    "key",
	},
	&cli.BoolFlag{
		Name:     "meta-key",
		Required: false,
		Usage:    "metaKey",
	},
	&cli.BoolFlag{
		Name:     "shift-key",
		Required: false,
		Usage:    "shiftKey",
	},
	&cli.BoolFlag{
		Name:     "ctrl-key",
		Required: false,
		Usage:    "ctrlKey",
	},
	&cli.StringFlag{
		Name:     "action",
		Required: false,
		Usage:    "action",
	},
	&cli.StringFlag{
		Name:     "action-key",
		Required: false,
		Usage:    "actionKey",
	},
}
var KeyboardShortcutCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:        "os",
		StructField: "Os",
		Required:    false,
		Usage:       "os",
		Type:        "string",
	},
	{
		Name:        "host",
		StructField: "Host",
		Required:    false,
		Usage:       "host",
		Type:        "string",
	},
	{
		Name:        "action",
		StructField: "Action",
		Required:    false,
		Usage:       "action",
		Type:        "string",
	},
	{
		Name:        "actionKey",
		StructField: "ActionKey",
		Required:    false,
		Usage:       "actionKey",
		Type:        "string",
	},
}
var KeyboardShortcutCommonCliFlagsOptional = []cli.Flag{
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
		Name:     "os",
		Required: false,
		Usage:    "os",
	},
	&cli.StringFlag{
		Name:     "host",
		Required: false,
		Usage:    "host",
	},
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
	&cli.BoolFlag{
		Name:     "alt-key",
		Required: false,
		Usage:    "altKey",
	},
	&cli.StringFlag{
		Name:     "key",
		Required: false,
		Usage:    "key",
	},
	&cli.BoolFlag{
		Name:     "meta-key",
		Required: false,
		Usage:    "metaKey",
	},
	&cli.BoolFlag{
		Name:     "shift-key",
		Required: false,
		Usage:    "shiftKey",
	},
	&cli.BoolFlag{
		Name:     "ctrl-key",
		Required: false,
		Usage:    "ctrlKey",
	},
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
	&cli.BoolFlag{
		Name:     "alt-key",
		Required: false,
		Usage:    "altKey",
	},
	&cli.StringFlag{
		Name:     "key",
		Required: false,
		Usage:    "key",
	},
	&cli.BoolFlag{
		Name:     "meta-key",
		Required: false,
		Usage:    "metaKey",
	},
	&cli.BoolFlag{
		Name:     "shift-key",
		Required: false,
		Usage:    "shiftKey",
	},
	&cli.BoolFlag{
		Name:     "ctrl-key",
		Required: false,
		Usage:    "ctrlKey",
	},
	&cli.StringFlag{
		Name:     "action",
		Required: false,
		Usage:    "action",
	},
	&cli.StringFlag{
		Name:     "action-key",
		Required: false,
		Usage:    "actionKey",
	},
}
var KeyboardShortcutCreateCmd cli.Command = KEYBOARD_SHORTCUT_ACTION_POST_ONE.ToCli()
var KeyboardShortcutCreateInteractiveCmd cli.Command = cli.Command{
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
			ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_KEYBOARD_SHORTCUT_CREATE},
		})
		entity := &KeyboardShortcutEntity{}
		for _, item := range KeyboardShortcutCommonInteractiveCliFlags {
			if !item.Required && c.Bool("all") == false {
				continue
			}
			result := workspaces.AskForInput(item.Name, "")
			workspaces.SetFieldString(entity, item.StructField, result)
		}
		if entity, err := KeyboardShortcutActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}
var KeyboardShortcutUpdateCmd cli.Command = cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Flags:   KeyboardShortcutCommonCliFlagsOptional,
	Usage:   "Updates a template by passing the parameters",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, &workspaces.SecurityModel{
			ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_KEYBOARD_SHORTCUT_UPDATE},
		})
		entity := CastKeyboardShortcutFromCli(c)
		if entity, err := KeyboardShortcutActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {
			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
		return nil
	},
}

func (x *KeyboardShortcutEntity) FromCli(c *cli.Context) *KeyboardShortcutEntity {
	return CastKeyboardShortcutFromCli(c)
}
func CastKeyboardShortcutFromCli(c *cli.Context) *KeyboardShortcutEntity {
	template := &KeyboardShortcutEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
	if c.IsSet("os") {
		value := c.String("os")
		template.Os = &value
	}
	if c.IsSet("host") {
		value := c.String("host")
		template.Host = &value
	}
	if c.IsSet("action") {
		value := c.String("action")
		template.Action = &value
	}
	if c.IsSet("action-key") {
		value := c.String("action-key")
		template.ActionKey = &value
	}
	return template
}
func KeyboardShortcutSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		KeyboardShortcutActionCreate,
		reflect.ValueOf(&KeyboardShortcutEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}
func KeyboardShortcutSyncSeeders() {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{WorkspaceId: workspaces.USER_SYSTEM},
		KeyboardShortcutActionCreate,
		reflect.ValueOf(&KeyboardShortcutEntity{}).Elem(),
		&seeders.ViewsFs,
		[]string{},
		true,
	)
}
func KeyboardShortcutWriteQueryMock(ctx workspaces.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := KeyboardShortcutActionQuery(f)
		result := workspaces.QueryEntitySuccessResult(f, items, count)
		workspaces.WriteMockDataToFile(lang, "", "KeyboardShortcut", result)
	}
}

var KeyboardShortcutImportExportCommands = []cli.Command{
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
				ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_KEYBOARD_SHORTCUT_CREATE},
			})
			KeyboardShortcutActionSeeder(query, c.Int("count"))
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
				Value: "keyboard-shortcut-seeder.yml",
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
				ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_KEYBOARD_SHORTCUT_CREATE},
			})
			KeyboardShortcutActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "keyboard-shortcut-seeder-keyboard-shortcut.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of keyboard-shortcuts, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]KeyboardShortcutEntity{}
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
				KeyboardShortcutActionCreate,
				reflect.ValueOf(&KeyboardShortcutEntity{}).Elem(),
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
				KeyboardShortcutActionQuery,
				reflect.ValueOf(&KeyboardShortcutEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"KeyboardShortcutFieldMap.yml",
				KeyboardShortcutPreloadRelations,
			)
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
			KeyboardShortcutCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmdAuthorized(c,
				KeyboardShortcutActionCreate,
				reflect.ValueOf(&KeyboardShortcutEntity{}).Elem(),
				c.String("file"),
				&workspaces.SecurityModel{
					ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_KEYBOARD_SHORTCUT_CREATE},
				},
				func() KeyboardShortcutEntity {
					v := CastKeyboardShortcutFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
var KeyboardShortcutCliCommands []cli.Command = []cli.Command{
	workspaces.GetCommonQuery2(KeyboardShortcutActionQuery, &workspaces.SecurityModel{
		ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_KEYBOARD_SHORTCUT_CREATE},
	}),
	workspaces.GetCommonTableQuery(reflect.ValueOf(&KeyboardShortcutEntity{}).Elem(), KeyboardShortcutActionQuery),
	KeyboardShortcutCreateCmd,
	KeyboardShortcutUpdateCmd,
	KeyboardShortcutCreateInteractiveCmd,
	KeyboardShortcutWipeCmd,
	workspaces.GetCommonRemoveQuery(reflect.ValueOf(&KeyboardShortcutEntity{}).Elem(), KeyboardShortcutActionRemove),
}

func KeyboardShortcutCliFn() cli.Command {
	KeyboardShortcutCliCommands = append(KeyboardShortcutCliCommands, KeyboardShortcutImportExportCommands...)
	return cli.Command{
		Name:        "keyboardShortcut",
		ShortName:   "kbshort",
		Description: "KeyboardShortcuts module actions (sample module to handle complex entities)",
		Usage:       "Manage the keyboard shortcuts in web and desktop apps (accessibility)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: KeyboardShortcutCliCommands,
	}
}

var KEYBOARD_SHORTCUT_ACTION_POST_ONE = workspaces.Module2Action{
	ActionName:    "create",
	ActionAliases: []string{"c"},
	Description:   "Create new keyboardShortcut",
	Flags:         KeyboardShortcutCommonCliFlags,
	Method:        "POST",
	Url:           "/keyboard-shortcut",
	SecurityModel: &workspaces.SecurityModel{},
	Handlers: []gin.HandlerFunc{
		func(c *gin.Context) {
			workspaces.HttpPostEntity(c, KeyboardShortcutActionCreate)
		},
	},
	CliAction: func(c *cli.Context, security *workspaces.SecurityModel) error {
		result, err := workspaces.CliPostEntity(c, KeyboardShortcutActionCreate, security)
		workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
		return err
	},
	Action:         KeyboardShortcutActionCreate,
	Format:         "POST_ONE",
	RequestEntity:  &KeyboardShortcutEntity{},
	ResponseEntity: &KeyboardShortcutEntity{},
}

/**
 *	Override this function on KeyboardShortcutEntityHttp.go,
 *	In order to add your own http
 **/
var AppendKeyboardShortcutRouter = func(r *[]workspaces.Module2Action) {}

func GetKeyboardShortcutModule2Actions() []workspaces.Module2Action {
	routes := []workspaces.Module2Action{
		{
			Method:        "GET",
			Url:           "/keyboard-shortcuts",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpQueryEntity(c, KeyboardShortcutActionQuery)
				},
			},
			Format:         "QUERY",
			Action:         KeyboardShortcutActionQuery,
			ResponseEntity: &[]KeyboardShortcutEntity{},
		},
		{
			Method:        "GET",
			Url:           "/keyboard-shortcuts/export",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpStreamFileChannel(c, KeyboardShortcutActionExport)
				},
			},
			Format:         "QUERY",
			Action:         KeyboardShortcutActionExport,
			ResponseEntity: &[]KeyboardShortcutEntity{},
		},
		{
			Method:        "GET",
			Url:           "/keyboard-shortcut/:uniqueId",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpGetEntity(c, KeyboardShortcutActionGetOne)
				},
			},
			Format:         "GET_ONE",
			Action:         KeyboardShortcutActionGetOne,
			ResponseEntity: &KeyboardShortcutEntity{},
		},
		KEYBOARD_SHORTCUT_ACTION_POST_ONE,
		{
			ActionName:    "update",
			ActionAliases: []string{"u"},
			Flags:         KeyboardShortcutCommonCliFlagsOptional,
			Method:        "PATCH",
			Url:           "/keyboard-shortcut",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpUpdateEntity(c, KeyboardShortcutActionUpdate)
				},
			},
			Action:         KeyboardShortcutActionUpdate,
			RequestEntity:  &KeyboardShortcutEntity{},
			Format:         "PATCH_ONE",
			ResponseEntity: &KeyboardShortcutEntity{},
		},
		{
			Method:        "PATCH",
			Url:           "/keyboard-shortcuts",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpUpdateEntities(c, KeyboardShortcutActionBulkUpdate)
				},
			},
			Action:         KeyboardShortcutActionBulkUpdate,
			Format:         "PATCH_BULK",
			RequestEntity:  &workspaces.BulkRecordRequest[KeyboardShortcutEntity]{},
			ResponseEntity: &workspaces.BulkRecordRequest[KeyboardShortcutEntity]{},
		},
		{
			Method:        "DELETE",
			Url:           "/keyboard-shortcut",
			Format:        "DELETE_DSL",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					workspaces.HttpRemoveEntity(c, KeyboardShortcutActionRemove)
				},
			},
			Action:         KeyboardShortcutActionRemove,
			RequestEntity:  &workspaces.DeleteRequest{},
			ResponseEntity: &workspaces.DeleteResponse{},
			TargetEntity:   &KeyboardShortcutEntity{},
		},
		{
			Method:        "PATCH",
			Url:           "/keyboard-shortcut/:linkerId/default_combination/:uniqueId",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(
					c *gin.Context,
				) {
					workspaces.HttpUpdateEntity(c, KeyboardShortcutDefaultCombinationActionUpdate)
				},
			},
			Action:         KeyboardShortcutDefaultCombinationActionUpdate,
			Format:         "PATCH_ONE",
			RequestEntity:  &KeyboardShortcutDefaultCombination{},
			ResponseEntity: &KeyboardShortcutDefaultCombination{},
		},
		{
			Method:        "GET",
			Url:           "/keyboard-shortcut/default_combination/:linkerId/:uniqueId",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(
					c *gin.Context,
				) {
					workspaces.HttpGetEntity(c, KeyboardShortcutDefaultCombinationActionGetOne)
				},
			},
			Action:         KeyboardShortcutDefaultCombinationActionGetOne,
			Format:         "GET_ONE",
			ResponseEntity: &KeyboardShortcutDefaultCombination{},
		},
		{
			Method:        "POST",
			Url:           "/keyboard-shortcut/:linkerId/default_combination",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(
					c *gin.Context,
				) {
					workspaces.HttpPostEntity(c, KeyboardShortcutDefaultCombinationActionCreate)
				},
			},
			Action:         KeyboardShortcutDefaultCombinationActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &KeyboardShortcutDefaultCombination{},
			ResponseEntity: &KeyboardShortcutDefaultCombination{},
		},
		{
			Method:        "PATCH",
			Url:           "/keyboard-shortcut/:linkerId/user_combination/:uniqueId",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(
					c *gin.Context,
				) {
					workspaces.HttpUpdateEntity(c, KeyboardShortcutUserCombinationActionUpdate)
				},
			},
			Action:         KeyboardShortcutUserCombinationActionUpdate,
			Format:         "PATCH_ONE",
			RequestEntity:  &KeyboardShortcutUserCombination{},
			ResponseEntity: &KeyboardShortcutUserCombination{},
		},
		{
			Method:        "GET",
			Url:           "/keyboard-shortcut/user_combination/:linkerId/:uniqueId",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(
					c *gin.Context,
				) {
					workspaces.HttpGetEntity(c, KeyboardShortcutUserCombinationActionGetOne)
				},
			},
			Action:         KeyboardShortcutUserCombinationActionGetOne,
			Format:         "GET_ONE",
			ResponseEntity: &KeyboardShortcutUserCombination{},
		},
		{
			Method:        "POST",
			Url:           "/keyboard-shortcut/:linkerId/user_combination",
			SecurityModel: &workspaces.SecurityModel{},
			Handlers: []gin.HandlerFunc{
				func(
					c *gin.Context,
				) {
					workspaces.HttpPostEntity(c, KeyboardShortcutUserCombinationActionCreate)
				},
			},
			Action:         KeyboardShortcutUserCombinationActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &KeyboardShortcutUserCombination{},
			ResponseEntity: &KeyboardShortcutUserCombination{},
		},
	}
	// Append user defined functions
	AppendKeyboardShortcutRouter(&routes)
	return routes
}
func CreateKeyboardShortcutRouter(r *gin.Engine) []workspaces.Module2Action {
	httpRoutes := GetKeyboardShortcutModule2Actions()
	workspaces.CastRoutes(httpRoutes, r)
	workspaces.WriteHttpInformationToFile(&httpRoutes, KeyboardShortcutEntityJsonSchema, "keyboard-shortcut-http", "keyboardActions")
	workspaces.WriteEntitySchema("KeyboardShortcutEntity", KeyboardShortcutEntityJsonSchema, "keyboardActions")
	return httpRoutes
}

var PERM_ROOT_KEYBOARD_SHORTCUT_DELETE = workspaces.PermissionInfo{
	CompleteKey: "root/keyboardActions/keyboard-shortcut/delete",
}
var PERM_ROOT_KEYBOARD_SHORTCUT_CREATE = workspaces.PermissionInfo{
	CompleteKey: "root/keyboardActions/keyboard-shortcut/create",
}
var PERM_ROOT_KEYBOARD_SHORTCUT_UPDATE = workspaces.PermissionInfo{
	CompleteKey: "root/keyboardActions/keyboard-shortcut/update",
}
var PERM_ROOT_KEYBOARD_SHORTCUT_QUERY = workspaces.PermissionInfo{
	CompleteKey: "root/keyboardActions/keyboard-shortcut/query",
}
var PERM_ROOT_KEYBOARD_SHORTCUT = workspaces.PermissionInfo{
	CompleteKey: "root/keyboardActions/keyboard-shortcut/*",
}
var ALL_KEYBOARD_SHORTCUT_PERMISSIONS = []workspaces.PermissionInfo{
	PERM_ROOT_KEYBOARD_SHORTCUT_DELETE,
	PERM_ROOT_KEYBOARD_SHORTCUT_CREATE,
	PERM_ROOT_KEYBOARD_SHORTCUT_UPDATE,
	PERM_ROOT_KEYBOARD_SHORTCUT_QUERY,
	PERM_ROOT_KEYBOARD_SHORTCUT,
}
