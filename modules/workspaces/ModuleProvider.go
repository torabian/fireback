package workspaces

import (
	"embed"
	"reflect"

	"github.com/urfave/cli"
	"golang.org/x/exp/maps"
	"gorm.io/gorm"
)

type TableMetaData struct {
	EntityName    string
	TableNameInDb string
	EntityObject  any
	ExportKey     string
	ExportStream  func(query QueryDSL) (chan []interface{}, *IError)
	ImportQuery   func(dto interface{}, query QueryDSL) *IError
}

type Report struct {
	Title        string
	Description  string
	UniqueId     string
	RowEntity    any
	Query        string
	QueryCounter string
	V            reflect.Value
	Fn           func(path string, query QueryDSL, report *Report, refl reflect.Value) *IError
}

// Entities also can be bundled into one
type EntityBundle struct {
	Permissions           []PermissionInfo
	Tests                 []Test
	Actions               []Module2Action
	AutoMigrationEntities []interface{}
	CliCommands           []cli.Command
	MockProvider          func()
}

type ModuleProvider struct {
	EntityProvider      func(*gorm.DB) error
	MockHandler         func()
	Reports             []Report
	SeederHandler       func()
	MockWriterHandler   func(languages []string)
	PermissionsProvider []PermissionInfo
	Name                string
	CliHandlers         []cli.Command
	BackupTables        []TableMetaData
	Tasks               []*TaskAction
	Definitions         *embed.FS
	Actions             [][]Module2Action
	Translations        map[string]map[string]string
	Tests               []Test

	EntityBundles []EntityBundle
}

func (x *ModuleProvider) ToModule2() Module2 {
	return Module2{
		Name: x.Name,
		Path: x.Name,
	}
}

func (x *ModuleProvider) ProvideMockImportHandler(t func()) {
	x.MockHandler = t
}

func (x *ModuleProvider) AppenedTasks(tasks ...[]*TaskAction) {
	for _, task := range tasks {
		x.Tasks = append(x.Tasks, task...)
	}
}

func (x *ModuleProvider) ProvideSeederImportHandler(t func()) {
	x.SeederHandler = t
}

func (x *ModuleProvider) ProvideMockWriterHandler(t func(languages []string)) {
	x.MockWriterHandler = t
}

func (x *ModuleProvider) ProvideTests(tests ...[]Test) {
	for _, t := range tests {
		x.Tests = append(x.Tests, t...)
	}
}

func (x *ModuleProvider) ProvideEntityHandlers(t func(*gorm.DB) error) {
	x.EntityProvider = t
}

func (x *ModuleProvider) ProvidePermissionHandler(items ...[]PermissionInfo) {
	for _, item := range items {
		x.PermissionsProvider = append(x.PermissionsProvider, item...)
	}
}

func (x *ModuleProvider) ProvideTranslationList(items ...map[string]map[string]string) {
	if x.Translations == nil {
		x.Translations = map[string]map[string]string{}
	}

	for _, item := range items {
		maps.Copy(x.Translations, item)
	}
}

func (x *ModuleProvider) ProvideCliHandlers(t []cli.Command) {
	x.CliHandlers = t
}
