//go:build wasm

package application

import (
	"embed"

	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback/gintools"
	"gorm.io/gorm"
)

type Application struct {

	// modules created by fireback are injected here
	Modules []*ModuleProvider

	SetupWebServerHook func(*gin.Engine, *Application)
	SeedersSync        func()
	MockSync           func()
	PublicFolders      []gintools.PublicFolderInfo
}

type MigrationScript struct {
	Exec func() error
}

// Entities also can be bundled into one
type EntityBundle struct {
	Permissions           []PermissionInfo
	AutoMigrationEntities []interface{}
	MockProvider          func()
	MigrationScripts      []MigrationScript
}

type ModuleProvider struct {
	EntityProvider    func(*gorm.DB) error
	MockHandler       func()
	SeederHandler     func()
	Namespace         string
	MigrationFunction func(x *Application, db *gorm.DB)

	MockWriterHandler   func(languages []string)
	PermissionsProvider []PermissionInfo
	Name                string

	GoMigrateDirectory *embed.FS

	// On goose or go migrate, there is such table to keep knowledge of how much migration
	// already has happened. If empty, module name is used, but when it's duplicate name or
	// nested modules then it would be good to change it.
	DatabaseMigrationHistoryName string

	Children []*ModuleProvider
}

type PermissionInfo struct {
	Name        string `yaml:"name,omitempty" json:"name,omitempty"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
	CompleteKey string `yaml:"completeKey,omitempty" json:"completeKey,omitempty"`
	GoVariable  string `yaml:"-" json:"-"`
}
