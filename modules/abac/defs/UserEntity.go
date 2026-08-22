package abacdefs

import (
	"encoding/json"
	"fmt"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/emigorm"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
	"github.com/torabian/fireback/modules/fireback/complexes"
	"gorm.io/gorm"
)

// The base class definition for userEntity
type UserEntity struct {
	Id        int64               `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId  string              `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	FirstName string              `json:"firstName" validate:"required" yaml:"firstName"`
	LastName  string              `json:"lastName" validate:"required" yaml:"lastName"`
	Photo     string              `json:"photo" yaml:"photo"`
	Gender    emigo.Nullable[int] `json:"gender" yaml:"gender"`
	Title     string              `json:"title" yaml:"title"`
	BirthDate complexes.XDate     `json:"birthDate" yaml:"birthDate"`
	Avatar    string              `json:"avatar" yaml:"avatar"`
	// User last connecting ip address
	LastIpAddress string `json:"lastIpAddress" yaml:"lastIpAddress"`
	// User primary address location. Can be useful for simple projects that a user is associated with a single address.
	PrimaryAddress emigo.Nullable[UserEntityPrimaryAddress] `gorm:"-" json:"primaryAddress" yaml:"primaryAddress"`
	// Contact phone number for this user (separate from any passport used to sign in).
	PhoneNumber emigo.Nullable[string] `json:"phoneNumber" yaml:"phoneNumber"`
	// The user's job title/role, e.g. "Support Engineer".
	JobTitle emigo.Nullable[string] `json:"jobTitle" yaml:"jobTitle"`
	// The company or organization the user is associated with.
	Company emigo.Nullable[string] `json:"company" yaml:"company"`
	// Free-form biography/notes about the user.
	Bio               emigo.Nullable[string]    `gorm:"text" json:"bio" yaml:"bio"`
	UserId            emigo.Nullable[string]    `json:"userId" yaml:"userId"`
	CreatedAt         abaccomplexes.PlainTime   `json:"createdAt" yaml:"createdAt"`
	UpdatedAt         abaccomplexes.PlainTime   `json:"updatedAt" yaml:"updatedAt"`
	PrimaryAddressRow *UserEntityPrimaryAddress `gorm:"embedded" json:"-" yaml:"-"`
}

// The base class definition for primaryAddress
type UserEntityPrimaryAddress struct {
	// Street address, building number
	AddressLine1 string `json:"addressLine1" yaml:"addressLine1"`
	// Apartment, suite, floor (optional)
	AddressLine2 emigo.Nullable[string] `json:"addressLine2" yaml:"addressLine2"`
	// City or locality
	City emigo.Nullable[string] `json:"city" yaml:"city"`
	// State, region, or province
	StateOrProvince emigo.Nullable[string] `json:"stateOrProvince" yaml:"stateOrProvince"`
	// ZIP or postal code
	PostalCode emigo.Nullable[string] `json:"postalCode" yaml:"postalCode"`
	// ISO 3166-1 alpha-2 (e.g., "US", "DE")
	CountryCode emigo.Nullable[string] `json:"countryCode" yaml:"countryCode"`
}

func (x *UserEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
// UserEntityCreateFn creates a new UserEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func UserEntityCreateFn(tx *gorm.DB, dto *UserEntity) (*UserEntity, error) {
	err := tx.Transaction(func(tx *gorm.DB) error {
		if v, ok := dto.PrimaryAddress.Get(); ok && v != nil {
			dto.PrimaryAddressRow = v
		}
		if err := tx.Create(dto).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return dto, nil
}

// UserEntityUpdateFn applies a partial update to the UserEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// UserEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func UserEntityUpdateFn(tx *gorm.DB, uniqueId string, input UserOptionalDto) (*UserEntity, error) {
	var entity UserEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.FirstName.IsSet() {
			changes["FirstName"] = input.FirstName
		}
		if input.LastName.IsSet() {
			changes["LastName"] = input.LastName
		}
		if input.Photo.IsSet() {
			changes["Photo"] = input.Photo
		}
		if input.Gender.IsSet() {
			changes["Gender"] = input.Gender
		}
		if input.Title.IsSet() {
			changes["Title"] = input.Title
		}
		changes["BirthDate"] = input.BirthDate
		if input.Avatar.IsSet() {
			changes["Avatar"] = input.Avatar
		}
		if input.LastIpAddress.IsSet() {
			changes["LastIpAddress"] = input.LastIpAddress
		}
		if input.PrimaryAddress.IsSet() {
			if v, ok := input.PrimaryAddress.Get(); ok && v != nil {
				if v.AddressLine1.IsSet() {
					changes["AddressLine1"] = v.AddressLine1
				}
				if v.AddressLine2.IsSet() {
					changes["AddressLine2"] = v.AddressLine2
				}
				if v.City.IsSet() {
					changes["City"] = v.City
				}
				if v.StateOrProvince.IsSet() {
					changes["StateOrProvince"] = v.StateOrProvince
				}
				if v.PostalCode.IsSet() {
					changes["PostalCode"] = v.PostalCode
				}
				if v.CountryCode.IsSet() {
					changes["CountryCode"] = v.CountryCode
				}
			}
		}
		if input.PhoneNumber.IsSet() {
			changes["PhoneNumber"] = input.PhoneNumber
		}
		if input.JobTitle.IsSet() {
			changes["JobTitle"] = input.JobTitle
		}
		if input.Company.IsSet() {
			changes["Company"] = input.Company
		}
		if input.Bio.IsSet() {
			changes["Bio"] = input.Bio
		}
		if input.UserId.IsSet() {
			changes["UserId"] = input.UserId
		}
		changes["CreatedAt"] = input.CreatedAt
		changes["UpdatedAt"] = input.UpdatedAt
		if len(changes) > 0 {
			if err := tx.Model(&entity).Updates(changes).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	var updated UserEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// UserEntityGetFn looks up a single UserEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func UserEntityGetFn(tx *gorm.DB, uniqueId string) (*UserEntity, error) {
	var entity UserEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// UserEntityBrowseFn returns UserEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func UserEntityBrowseFn(tx *gorm.DB, qs UserBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*UserEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&UserEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*UserEntity
	paged := emigorm.ApplyQueryPage(emigorm.ApplyQueryCursor(emigorm.ApplyQuerySort(filtered, qs.Sort), qs.Cursor), qs.StartIndex, qs.ItemsPerPage)
	if err := paged.Find(&items).Error; err != nil {
		return nil, nil, err
	}
	meta := &emigo.QueryResultMeta{
		TotalItems: total,
		Cursor:     emigorm.BuildQueryCursor(items),
	}
	return items, meta, nil
}

// UserEntityAwareDeleteAffected reports one relation of UserEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on UserEntity itself, so deleting UserEntity doesn't cascade into them.
type UserEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// UserEntityAwareDeletePreview is the result of UserEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts UserEntityAwareDeleteFn would delete/clear
// alongside the UserEntity row(s) themselves.
type UserEntityAwareDeletePreview struct {
	Message  string                          `json:"message"`
	Affected []UserEntityAwareDeleteAffected `json:"affected"`
}

// UserEntityAwareDeletePreviewFn looks up the UserEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// UserEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling UserEntityAwareDeleteFn.
func UserEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*UserEntityAwareDeletePreview, error) {
	var rows []*UserEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &UserEntityAwareDeletePreview{Message: "No matching UserEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []UserEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d UserEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &UserEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// UserEntityAwareDeleteFn deletes the UserEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation UserEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func UserEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*UserEntity
		if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
			return err
		}
		if len(rows) == 0 {
			return nil
		}
		ids := make([]int64, len(rows))
		for i := range rows {
			ids[i] = rows[i].Id
		}
		return tx.Where("id IN ?", ids).Delete(&UserEntity{}).Error
	})
}

// UserEntityActionsSig bundles the actions available for UserEntity. Extend this (and
// UserEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type UserEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *UserEntity) (*UserEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input UserOptionalDto) (*UserEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*UserEntity, error)
	Browse             func(tx *gorm.DB, qs UserBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*UserEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*UserEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var UserEntityActions UserEntityActionsSig = UserEntityActionsSig{
	Create:             UserEntityCreateFn,
	Update:             UserEntityUpdateFn,
	Get:                UserEntityGetFn,
	Browse:             UserEntityBrowseFn,
	AwareDeletePreview: UserEntityAwareDeletePreviewFn,
	AwareDelete:        UserEntityAwareDeleteFn,
}
