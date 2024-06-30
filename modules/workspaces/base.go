package workspaces

import (
	"time"

	"gorm.io/gorm"
)

// @meta(include)
type Model struct {
	UniqueId  string         `gorm:"primaryKey;autoIncrement:false" json:"uniqueId"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
