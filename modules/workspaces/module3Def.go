/**
Current file is set of definitions, to create Module2 yaml files.
Module2 is a declarative way of creating backend entities, crud actions on them,
and many complex operation. Fireback would generate those codes for many languages
both for backend and front-end purposes.

Backend code can be generated in: C and Golang
Front-end code can be generated in: Angular, React, Pure TypeScript, Android Java, Swift
*/

package workspaces

import (
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

type ErrorItem map[string]string

type Module2EntityConfig struct {
	UseFields     []string       `yaml:"useFields,omitempty" json:"useFields,omitempty"`
	SecurityModel *SecurityModel `yaml:"security,omitempty" json:"security,omitempty"`
}

// Module2 struct represents the entire file tree
type Module2 struct {
	Entity Module2EntityConfig `yaml:"entity,omitempty" json:"entity,omitempty"`

	// represents where is the location of the module in app tree, similar to PHP namespacing sytem
	// it be used to explicitly as export path of the actions for client frameworks
	Namespace     string           `yaml:"namespace,omitempty" json:"namespace,omitempty"`
	Path          string           `yaml:"path,omitempty" json:"path,omitempty"`
	Description   string           `yaml:"description,omitempty" json:"description,omitempty"`
	Version       string           `yaml:"version,omitempty" json:"version,omitempty"`
	MetaWorkspace bool             `yaml:"meta-workspace,omitempty" json:"meta-workspace,omitempty"`
	Name          string           `yaml:"name,omitempty" json:"name,omitempty"`
	Entities      []Module2Entity  `yaml:"entities,omitempty" json:"entities,omitempty"`
	Tasks         []*Module2Task   `yaml:"tasks,omitempty" json:"tasks,omitempty"`
	Dto           []Module2DtoBase `yaml:"dtos,omitempty" json:"dtos,omitempty"`
	Actions       []*Module2Action `yaml:"actions,omitempty" json:"actions,omitempty"`
	Macros        []Module2Macro   `yaml:"macros,omitempty" json:"macros,omitempty"`
	Remotes       []*Module2Remote `yaml:"remotes,omitempty" json:"remotes,omitempty"`
	Queries       []*Module2Query  `yaml:"queries,omitempty" json:"queries,omitempty"`

	// An interesting way of defining env variables
	Config []*Module2ConfigField `yaml:"config,omitempty" json:"config,omitempty"`

	Messages Module2Message `yaml:"messages,omitempty" json:"messages,omitempty"`
}

type Module2Trigger struct {
	Cron *string `yaml:"cron,omitempty" json:"cron,omitempty"`
}

type Module2Task struct {
	Triggers    []Module2Trigger   `yaml:"triggers,omitempty" json:"triggers,omitempty"`
	Name        string             `yaml:"name,omitempty" json:"name,omitempty"`
	Description string             `yaml:"description,omitempty" json:"description,omitempty"`
	In          *Module2ActionBody `yaml:"in,omitempty" json:"in,omitempty"`
}

// This is a fireback remote definition, you can make the external API calls typesafe using
// definitions. This feature is documented in docs/docs/definitions/remotes.md
type Module2Remote struct {

	// Http method, lower case post, delete, ...
	Method string `yaml:"method,omitempty" json:"method,omitempty"`

	// The url which will be requested. You need to add full url here, but maybe you could add a prefix
	// also in the client from your Go code - There might be a prefix for remotes later version of fireback
	Url string `yaml:"url,omitempty" json:"url,omitempty"`

	// Standard Module2ActionBody object. Could have fields, entity, dto as content and you
	// can define the output to cast automatically into them.
	Out *Module2ActionBody `yaml:"out,omitempty" json:"out,omitempty"`

	// Standard Module2ActionBody object. Could have fields, entity, dto as content and you
	// can define the input parameters as struct in Go and fireback will convert it into
	// json.
	In *Module2ActionBody `yaml:"in,omitempty" json:"in,omitempty"`

	// Query params for the address, if you want to define them in Golang dynamically instead of URL
	Query []*Module2Field `yaml:"query,omitempty" json:"query,omitempty"`

	// Remote action name, it will become the Golang function that you will call
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	ResponseFields []*Module2Field `yaml:"-" json:"-"`
}

type Module2FieldOf struct {
	Key string `yaml:"k,omitempty" json:"k,omitempty"`
}

type Module2Macro struct {
	Using string `yaml:"using,omitempty" json:"using,omitempty"`
	Name  string `yaml:"name,omitempty" json:"name,omitempty"`
	// Might be used on some macros
	Fields []*Module2Field `yaml:"fields,omitempty" json:"fields,omitempty"`
}

type Module2FieldMatch struct {
	Dto *string `yaml:"dto,omitempty" json:"dto,omitempty"`
}

type GormOverrideMap struct {
	WorkspaceId string `yaml:"workspaceId,omitempty" json:"workspaceId,omitempty"`
	UserId      string `yaml:"userId,omitempty" json:"userId,omitempty"`
}

type Security struct {
	Model string `yaml:"model,omitempty" json:"model,omitempty"`
}

type Module2Http struct {
	Query bool `yaml:"query,omitempty" json:"query,omitempty"`
}

type Module2Permission struct {
	Name        string `yaml:"name,omitempty" json:"name,omitempty"`
	Key         string `yaml:"key,omitempty" json:"key,omitempty"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
}

type Module2Message map[string]map[string]string

type Module2DataFields struct {
	// Essential is a set of the fields which fireback uses to give userId and workspaceId
	Essentials bool

	// Adds a int primary key auto increment
	PrimaryId bool

	// adds created, updated, delete as nano seconds to the database
	NumericTimestamp bool

	// adds created, updated, deleted fields as timestamps
	DateTimestamp bool
}

// Represents Entities in Fireback. An entity in Fireback is a table in database, with addition general
// features such as permissions, actions, security, and common actions which might be created or extra
// queries based on the type
type Module2Entity struct {
	// Extra permissions that an entity might need. You can add extra permissions that you will need in your
	// business logic related to entity in itself, to make it easier become as a group and document
	// later
	Permissions []Module2Permission `yaml:"permissions,omitempty" json:"permissions,omitempty"`

	// Actions or extra actions (on top of default actions which automatically is generated) these are
	// the same actions that you can define for a module, but defining them on entity level make it easier
	// to relate them and group them. Also permission might be added automatically (need to clearify)
	Actions []*Module2Action `yaml:"actions,omitempty" json:"actions,omitempty"`

	// Entity name is very important, based on this entity the tables on data base will be created
	// and Go and other codegeneration tool related to Fireback will be using it.
	// Important: Changing entity name would not delete the previously created entities,
	// you need to delete previous files manually. Fireback would consider camelCase
	// names only for the entity name
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	// You can make sure there is only one record of the entity per user or workspace using this option.
	// for example, if you want only one credit card per workspace, you can set distinctBy: workspace
	// and it will do the job
	DistinctBy string `yaml:"distinctBy,omitempty" json:"distinctBy,omitempty"`

	// Changes the default table name based on project prefix (fb_ by default) and entity name
	// useful for times that you want to connect project to an existing database
	Table string `yaml:"table,omitempty" json:"table,omitempty"`

	UseFields           []string             `yaml:"useFields,omitempty" json:"useFields,omitempty"`
	SecurityModel       *EntitySecurityModel `yaml:"security,omitempty" json:"security,omitempty"`
	PrependScript       string               `yaml:"prependScript,omitempty" json:"prependScript,omitempty"`
	Messages            Module2Message       `yaml:"messages,omitempty" json:"messages,omitempty"`
	PrependCreateScript string               `yaml:"prependCreateScript,omitempty" json:"prependCreateScript,omitempty"`
	PrependUpdateScript string               `yaml:"prependUpdateScript,omitempty" json:"prependUpdateScript,omitempty"`
	NoQuery             bool                 `yaml:"noQuery,omitempty" json:"noQuery,omitempty"`
	Access              string               `yaml:"access,omitempty" json:"access,omitempty"`
	QueryScope          string               `yaml:"queryScope,omitempty" json:"queryScope,omitempty"`
	Http                Module2Http          `yaml:"http,omitempty" json:"http,omitempty"`
	Patch               bool                 `yaml:"patch,omitempty" json:"patch,omitempty"`
	Queries             []string             `yaml:"queries,omitempty" json:"queries,omitempty"`
	Get                 bool                 `yaml:"get,omitempty" json:"get,omitempty"`
	GormMap             GormOverrideMap      `yaml:"gormMap,omitempty" json:"gormMap,omitempty"`
	Query               bool                 `yaml:"query,omitempty" json:"query,omitempty"`
	Post                bool                 `yaml:"post,omitempty" json:"post,omitempty"`
	ImportList          []string             `yaml:"importList,omitempty" json:"importList,omitempty"`
	Fields              []*Module2Field      `yaml:"fields,omitempty" json:"fields,omitempty"`
	C                   bool                 `yaml:"c,omitempty" json:"c,omitempty"`
	CliName             string               `yaml:"cliName,omitempty" json:"cliName,omitempty"`
	CliShort            string               `yaml:"cliShort,omitempty" json:"cliShort,omitempty"`
	CliDescription      string               `yaml:"cliDescription,omitempty" json:"cliDescription,omitempty"`
	Cte                 bool                 `yaml:"cte,omitempty" json:"cte,omitempty"`
	PostFormatter       string               `yaml:"postFormatter,omitempty" json:"postFormatter,omitempty"`
}

// This is the new dto version
type Module2DtoBase struct {
	Name       string          `yaml:"name,omitempty" json:"name,omitempty"`
	ImportList []string        `yaml:"importList,omitempty" json:"importList,omitempty"`
	Fields     []*Module2Field `yaml:"fields,omitempty" json:"fields,omitempty"`
}

type Module2ActionBody struct {
	Fields []*Module2Field `yaml:"fields,omitempty" json:"fields,omitempty"`
	Dto    string          `yaml:"dto,omitempty" json:"dto,omitempty"`
	Entity string          `yaml:"entity,omitempty" json:"entity,omitempty"`
}

type Module2Action struct {
	ActionName    string          `yaml:"actionName,omitempty" json:"actionName,omitempty"`
	CliName       string          `yaml:"cliName,omitempty" json:"cliName,omitempty"`
	ActionAliases []string        `yaml:"actionAliases,omitempty" json:"actionAliases,omitempty"`
	Name          string          `yaml:"name,omitempty" json:"name,omitempty"`
	Url           string          `yaml:"url,omitempty" json:"url,omitempty"`
	Method        string          `yaml:"method,omitempty" json:"method,omitempty"`
	Query         []*Module2Field `yaml:"query,omitempty" json:"query,omitempty"`
	Fn            string          `yaml:"fn,omitempty" json:"fn,omitempty"`
	Description   string          `yaml:"description,omitempty" json:"description,omitempty"`

	Group           string             `yaml:"group,omitempty" json:"group,omitempty"`
	Format          string             `yaml:"format,omitempty" json:"format,omitempty"`
	In              *Module2ActionBody `yaml:"in,omitempty" json:"in,omitempty"`
	Out             *Module2ActionBody `yaml:"out,omitempty" json:"out,omitempty"`
	SecurityModel   *SecurityModel     `yaml:"security,omitempty" json:"security,omitempty"`
	CastBodyFromCli func(c *cli.Context) any
	CliAction       func(c *cli.Context, security *SecurityModel) error
	Flags           []cli.Flag
	Virtual         bool
	Handlers        []gin.HandlerFunc `yaml:"-" json:"-"`
	ExternFuncName  string            `yaml:"-" json:"-"`
	RequestEntity   any               `yaml:"-" json:"-"`
	ResponseEntity  any               `yaml:"-" json:"-"`
	Action          any               `yaml:"-" json:"-"`
	TargetEntity    any               `yaml:"-" json:"-"`
	RootModule      *Module2          `yaml:"-" json:"-"`
}

func (x Module2Action) MethodUpper() string {
	return strings.ToUpper(x.Method)
}

func (x Module2Action) ToCli() cli.Command {

	return cli.Command{
		Name:        x.ActionName,
		Aliases:     x.ActionAliases,
		Description: x.Description,
		Usage:       x.Description,
		Action: func(c *cli.Context) error {
			return x.CliAction(c, x.SecurityModel)
		},
		Flags: x.Flags,
	}
}

func (x *Module2) ToModuleProvider() *ModuleProvider {

	actions := []Module2Action{}
	for _, item := range x.Actions {
		actions = append(actions, *item)
	}
	return &ModuleProvider{
		Name: x.Name,
		Actions: [][]Module2Action{
			actions,
		},
	}
}

func (x *Module2Entity) DataFields() Module2DataFields {
	data := Module2DataFields{}

	if len(x.UseFields) == 0 {
		data = Module2DataFields{
			Essentials:       true,
			PrimaryId:        true,
			NumericTimestamp: true,
		}

		return data
	}

	if slices.Contains(x.UseFields, "essentials") {
		data.Essentials = true
	}
	if slices.Contains(x.UseFields, "datetime") {
		data.DateTimestamp = true
	}
	if slices.Contains(x.UseFields, "primary") {
		data.PrimaryId = true
	}
	if slices.Contains(x.UseFields, "nanotime") {
		data.NumericTimestamp = true
	}

	return data
}
