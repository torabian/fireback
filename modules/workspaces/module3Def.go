/**
Current file is set of definitions, to create Module3 yaml files.
Module3 is a declarative way of creating backend entities, crud actions on them,
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

type Module3EntityConfig struct {
	UseFields     []string       `yaml:"useFields,omitempty" json:"useFields,omitempty"`
	SecurityModel *SecurityModel `yaml:"security,omitempty" json:"security,omitempty"`
}

// Module3 struct represents the entire file tree
type Module3 struct {
	Entity Module3EntityConfig `yaml:"entity,omitempty" json:"entity,omitempty"`

	// represents where is the location of the module in app tree, similar to PHP namespacing sytem
	// it be used to explicitly as export path of the actions for client frameworks
	Namespace     string           `yaml:"namespace,omitempty" json:"namespace,omitempty"`
	Path          string           `yaml:"path,omitempty" json:"path,omitempty"`
	Description   string           `yaml:"description,omitempty" json:"description,omitempty"`
	Version       string           `yaml:"version,omitempty" json:"version,omitempty"`
	MetaWorkspace bool             `yaml:"meta-workspace,omitempty" json:"meta-workspace,omitempty"`
	Name          string           `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the module"`
	Entities      []Module3Entity  `yaml:"entities,omitempty" json:"entities,omitempty"`
	Tasks         []*Module3Task   `yaml:"tasks,omitempty" json:"tasks,omitempty"`
	Dto           []Module3Dto     `yaml:"dtos,omitempty" json:"dtos,omitempty"`
	Actions       []*Module3Action `yaml:"actions,omitempty" json:"actions,omitempty"`
	Macros        []Module3Macro   `yaml:"macros,omitempty" json:"macros,omitempty"`
	Remotes       []*Module3Remote `yaml:"remotes,omitempty" json:"remotes,omitempty"`
	Queries       []*Module3Query  `yaml:"queries,omitempty" json:"queries,omitempty"`

	// An interesting way of defining env variables
	Config []*Module3ConfigField `yaml:"config,omitempty" json:"config,omitempty"`

	Messages Module3Message `yaml:"messages,omitempty" json:"messages,omitempty"`
}

type Module3Trigger struct {
	Cron *string `yaml:"cron,omitempty" json:"cron,omitempty"`
}

type Module3Task struct {
	Triggers    []Module3Trigger   `yaml:"triggers,omitempty" json:"triggers,omitempty"`
	Name        string             `yaml:"name,omitempty" json:"name,omitempty"`
	Description string             `yaml:"description,omitempty" json:"description,omitempty"`
	In          *Module3ActionBody `yaml:"in,omitempty" json:"in,omitempty"`
}

// This is a fireback remote definition, you can make the external API calls typesafe using
// definitions. This feature is documented in docs/remotes.md
type Module3Remote struct {

	// Http method, lower case post, delete, ...
	Method string `yaml:"method,omitempty" json:"method,omitempty"`

	// The url which will be requested. You need to add full url here, but maybe you could add a prefix
	// also in the client from your Go code - There might be a prefix for remotes later version of fireback
	Url string `yaml:"url,omitempty" json:"url,omitempty"`

	// Standard Module3ActionBody object. Could have fields, entity, dto as content and you
	// can define the output to cast automatically into them.
	Out *Module3ActionBody `yaml:"out,omitempty" json:"out,omitempty"`

	// Standard Module3ActionBody object. Could have fields, entity, dto as content and you
	// can define the input parameters as struct in Go and fireback will convert it into
	// json.
	In *Module3ActionBody `yaml:"in,omitempty" json:"in,omitempty"`

	// Query params for the address, if you want to define them in Golang dynamically instead of URL
	Query []*Module3Field `yaml:"query,omitempty" json:"query,omitempty"`

	// Remote action name, it will become the Golang function that you will call
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	ResponseFields []*Module3Field `yaml:"-" json:"-"`
}

type Module3FieldOf struct {
	Key string `yaml:"k,omitempty" json:"k,omitempty"`
}

type Module3Macro struct {
	Using string `yaml:"using,omitempty" json:"using,omitempty"`
	Name  string `yaml:"name,omitempty" json:"name,omitempty"`
	// Might be used on some macros
	Fields []*Module3Field `yaml:"fields,omitempty" json:"fields,omitempty"`
}

type Module3FieldMatch struct {
	Dto *string `yaml:"dto,omitempty" json:"dto,omitempty"`
}

type GormOverrideMap struct {
	WorkspaceId string `yaml:"workspaceId,omitempty" json:"workspaceId,omitempty"`
	UserId      string `yaml:"userId,omitempty" json:"userId,omitempty"`
}

type Security struct {
	Model string `yaml:"model,omitempty" json:"model,omitempty"`
}

type Module3Http struct {
	Query bool `yaml:"query,omitempty" json:"query,omitempty"`
}

type Module3Permission struct {
	Name        string `yaml:"name,omitempty" json:"name,omitempty"`
	Key         string `yaml:"key,omitempty" json:"key,omitempty"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
}

type Module3Message map[string]map[string]string

type Module3DataFields struct {
	// Essential is a set of the fields which fireback uses to give userId and workspaceId
	Essentials bool

	// Adds a int primary key auto increment
	PrimaryId bool

	// adds created, updated, delete as nano seconds to the database
	NumericTimestamp bool

	// adds created, updated, deleted fields as timestamps
	DateTimestamp bool
}

// Used to adjust the features generated for each entity.
type Module3EntityFeatures struct {

	// Adds a CLI task to make automatic mock, you can disable this
	// if it makes no sense for the feature or has validations required
	// so only a custom mock makes sense - IMPORTANT this is pointer
	// because by default it's enabled
	Mock *bool `yaml:"mock,omitempty" json:"mock,omitempty"`

	// Msync enables to have the embedded mock files for an entity
	// it would disable the 'msync' and 'mlist' commands
	MSync *bool `yaml:"msync,omitempty" json:"msync,omitempty"`
}

// Checks if codegen needs to print a default mocking tools for the entity
func (x Module3EntityFeatures) HasMockAction() bool {

	if x.Mock != nil && !*x.Mock {
		return false
	}

	return true
}

// Checks id module3 definition enabled or disabled the feature
func (x Module3EntityFeatures) HasMsyncActions() bool {

	if x.MSync != nil && !*x.MSync {
		return false
	}

	return true
}

// Represents Entities in Fireback. An entity in Fireback is a table in database, with addition general
// features such as permissions, actions, security, and common actions which might be created or extra
// queries based on the type
type Module3Entity struct {
	// Extra permissions that an entity might need. You can add extra permissions that you will need in your
	// business logic related to entity in itself, to make it easier become as a group and document
	// later
	Permissions []Module3Permission `yaml:"permissions,omitempty" json:"permissions,omitempty"`

	// Actions or extra actions (on top of default actions which automatically is generated) these are
	// the same actions that you can define for a module, but defining them on entity level make it easier
	// to relate them and group them. Also permission might be added automatically (need to clearify)
	Actions []*Module3Action `yaml:"actions,omitempty" json:"actions,omitempty"`

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

	// Customize the features generated for entity, less common  changes goes to this object
	Features Module3EntityFeatures `yaml:"features,omitempty" json:"features,omitempty"`

	// Changes the default table name based on project prefix (fb_ by default) and entity name
	// useful for times that you want to connect project to an existing database
	Table string `yaml:"table,omitempty" json:"table,omitempty"`

	UseFields           []string             `yaml:"useFields,omitempty" json:"useFields,omitempty"`
	SecurityModel       *EntitySecurityModel `yaml:"security,omitempty" json:"security,omitempty"`
	PrependScript       string               `yaml:"prependScript,omitempty" json:"prependScript,omitempty"`
	Messages            Module3Message       `yaml:"messages,omitempty" json:"messages,omitempty"`
	PrependCreateScript string               `yaml:"prependCreateScript,omitempty" json:"prependCreateScript,omitempty"`
	PrependUpdateScript string               `yaml:"prependUpdateScript,omitempty" json:"prependUpdateScript,omitempty"`
	NoQuery             bool                 `yaml:"noQuery,omitempty" json:"noQuery,omitempty"`
	Access              string               `yaml:"access,omitempty" json:"access,omitempty"`
	QueryScope          string               `yaml:"queryScope,omitempty" json:"queryScope,omitempty"`
	Http                Module3Http          `yaml:"http,omitempty" json:"http,omitempty"`
	Patch               bool                 `yaml:"patch,omitempty" json:"patch,omitempty"`
	Queries             []string             `yaml:"queries,omitempty" json:"queries,omitempty"`
	Get                 bool                 `yaml:"get,omitempty" json:"get,omitempty"`
	GormMap             GormOverrideMap      `yaml:"gormMap,omitempty" json:"gormMap,omitempty"`
	Query               bool                 `yaml:"query,omitempty" json:"query,omitempty"`
	Post                bool                 `yaml:"post,omitempty" json:"post,omitempty"`
	ImportList          []string             `yaml:"importList,omitempty" json:"importList,omitempty"`
	Fields              []*Module3Field      `yaml:"fields,omitempty" json:"fields,omitempty"`
	C                   bool                 `yaml:"c,omitempty" json:"c,omitempty"`
	CliName             string               `yaml:"cliName,omitempty" json:"cliName,omitempty"`
	CliShort            string               `yaml:"cliShort,omitempty" json:"cliShort,omitempty"`
	Description         string               `yaml:"description,omitempty" json:"description,omitempty"`
	Cte                 bool                 `yaml:"cte,omitempty" json:"cte,omitempty"`
	PostFormatter       string               `yaml:"postFormatter,omitempty" json:"postFormatter,omitempty"`
}

// Represents a dto in an application. Can be used for variety of reasons,
// request response of an action, or even internally. Fireback generates bunch of
// helpers for each dto, so it might make sense to define them in Module3 instead
// of pure struct in golang.
type Module3Dto struct {

	// Name of the dto, in camel case, the rest of the code related to this dto is being generated
	// based on this
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	// List of fields and body definitions of the dto
	Fields []*Module3Field `yaml:"fields,omitempty" json:"fields,omitempty"`
}

type Module3ActionBody struct {
	Fields    []*Module3Field `yaml:"fields,omitempty" json:"fields,omitempty"`
	Dto       string          `yaml:"dto,omitempty" json:"dto,omitempty"`
	Entity    string          `yaml:"entity,omitempty" json:"entity,omitempty"`
	Primitive string          `yaml:"primitive,omitempty" json:"primitive,omitempty"`
}

// Defines an action, very similar to http action (controller) on other framework,
// the difference is it's accessbile both on cli and http, and it's less tight to
// http definitions, and able to operate on socket directy.
type Module3Action struct {

	// General name of the action, which will be used to create golang code, req/res bodies,
	// and other places. Besides, it would become available on the cli using this by default
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	// Overrides the cli action name, if not specified the 'name' would be used instead,
	// and might be casted to dashed instead of camel case
	CliName string `yaml:"cliName,omitempty" json:"cliName,omitempty"`

	// A list of aliases on cli command only. If the action name is too long, you can specify
	// some shorter characters to make it easier for the cli user. aka 'u' for update which
	// fireback internally uses
	ActionAliases []string `yaml:"actionAliases,omitempty" json:"actionAliases,omitempty"`

	// The address of action on http server. similar to traditional /api/address/etc style.
	// just notice that the address can be prefixed by module or nested modules, but the last
	// part would be url.
	// If url is not specified, the action won't be available on the http router and becomes cli only
	// implicitly
	Url string `yaml:"url,omitempty" json:"url,omitempty"`

	// similar to standard http methods, post, get, delete, and others.
	// method: reactive is also added by fireback for opening web sockets connection directly
	Method string `yaml:"method,omitempty" json:"method,omitempty"`

	// Type-safe query params which will become available with --qs in cli, and normal
	// query strings in the http requests.
	Query []*Module3Field `yaml:"query,omitempty" json:"query,omitempty"`

	// Description of the action, which would become to available on different locations,
	// on comments, api spec, describing features, and many more.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	// Format is a higher level of the method definition for each request.
	// The formats available at this moment:
	// POST_ONE: single post body
	// PATCH_ONE: single patch body
	// QUERY: intended to search and return an array of items always,
	// and is compatible with infinite scroll
	// PATCH_BULK it's capable of sending array of entity and return array of entity
	// after patching them.
	// Other formats can be easily added to fireback source
	Format string `yaml:"format,omitempty" json:"format,omitempty"`

	// Similar to body in http request, (post,patch,...) and can contain fields, entity, dto
	// Check Module3ActionBody struct definition for further details
	In *Module3ActionBody `yaml:"in,omitempty" json:"in,omitempty"`

	// Similar to response in http request, (post,patch,...) and can contain fields, entity, dto
	// Check Module3ActionBody struct definition for further details
	Out *Module3ActionBody `yaml:"out,omitempty" json:"out,omitempty"`

	// Security model defines how the action is accessible for users. Similar to guard or middlewears
	// to check the access level, permission, and token.
	// Security model is a wider topic and you need to check SecurityModel struct for further definitions
	SecurityModel *SecurityModel `yaml:"security,omitempty" json:"security,omitempty"`

	// The action implementation in cli. Fireback generated code handles that in a way
	// to have the same functionality for both cli and http, but in action level you can
	// define it's own implementation regardless of the http and vice versa
	CliAction func(c *cli.Context, security *SecurityModel) error `jsonschema:"-"`

	// http implementation of the action. You need to provide gin handlers (gin framework)
	// one by one. Fireback security is being checked before these handlers, you do not need to
	// check them again here. You can pass as many as handlers you want.
	Handlers []gin.HandlerFunc `yaml:"-" json:"-" jsonschema:"-"`

	// The flags that the CLI action should accept, similar to the http request body json definition.
	// check the urfave/cli library to understand more, we are using that directly.
	Flags []cli.Flag `yaml:"-" json:"-" jsonschema:"-"`

	// Used to create the external functions on code generation such as react (typescript)
	// if left empty, it would be calculated automatically url and some other logics.
	// search for: func (route Module3Action) GetFuncName() string
	// in the code base to understand the logic
	ExternFuncName string `yaml:"-" json:"-" jsonschema:"-"`

	// A pointer to empty struct which represents the request body, it would be used to create
	// code for rpc calls on typescript, swift.
	RequestEntity any `yaml:"-" json:"-" jsonschema:"-"`

	// A pointer to empty struct which represents the response body, it would be used to create
	// code for rpc calls on typescript, swift.
	ResponseEntity any `yaml:"-" json:"-" jsonschema:"-"`

	// The actual function which would represent the implementation of the action. This is only
	// for code generation purpose, cli action and http implementation is done via
	// Handlers and CliAction fields
	Action any `yaml:"-" json:"-" jsonschema:"-"`

	// Pointer to the struct which would be operating on the object. Some actions such as
	// deletion do not have request or response, therefor a TargetEntiy pointer is being
	// used to detect the classes generated
	TargetEntity any `yaml:"-" json:"-" jsonschema:"-"`

	// Meta data used in code gen internally only. It would attach the module3 instance to
	RootModule *Module3 `yaml:"-" json:"-" jsonschema:"-"`
}

func (x Module3Action) MethodUpper() string {
	return strings.ToUpper(x.Method)
}

func (x Module3Action) ToCli() cli.Command {

	return cli.Command{
		Name:        x.Name,
		Aliases:     x.ActionAliases,
		Description: x.Description,
		Usage:       x.Description,
		Action: func(c *cli.Context) error {
			return x.CliAction(c, x.SecurityModel)
		},
		Flags: x.Flags,
	}
}

func (x *Module3) ToModuleProvider() *ModuleProvider {

	actions := []Module3Action{}
	for _, item := range x.Actions {
		actions = append(actions, *item)
	}
	return &ModuleProvider{
		Name: x.Name,
		Actions: [][]Module3Action{
			actions,
		},
	}
}

func (x *Module3Entity) DataFields() Module3DataFields {
	data := Module3DataFields{}

	if len(x.UseFields) == 0 {
		data = Module3DataFields{
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
