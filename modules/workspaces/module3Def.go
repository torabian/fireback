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
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

type ErrorItem map[string]string

// Module3 struct represents the entire file tree
type Module3 struct {

	// Represents where is the location of the module in app tree. Similar to PHP namespacing sytem it be used to explicitly as export path of the actions for client frameworks
	Namespace string `yaml:"namespace,omitempty" json:"namespace,omitempty" jsonschema:"description=Represents where is the location of the module in app tree. Similar to PHP namespacing sytem it be used to explicitly as export path of the actions for client frameworks"`

	// Description of module and it's purpose. Used in code gen and creating documents.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description of module and it's purpose. Used in code gen and creating documents."`

	// Version of the module. Helpful for different code generation phases but it's not necessary.
	Version string `yaml:"version,omitempty" json:"version,omitempty" jsonschema:"description=Version of the module. Helpful for different code generation phases but it's not necessary."`

	// Magic property for Fireback WorkspacesModule3.yml file. It's gonna be true only in a single file internally in Fireback
	MetaWorkspace bool `yaml:"meta-workspace,omitempty" json:"meta-workspace,omitempty" jsonschema:"description=Magic property for Fireback WorkspacesModule3.yml file. It's gonna be true only in a single file internally in Fireback"`

	// Name of the module. Needs to be lower camel case and Module.go and Module.dyno.go will be generated based on this name.
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the module. Needs to be lower camel case and Module.go and Module.dyno.go will be generated based on this name."`

	// List of entities that module contains. Entities are basically tables in database with their mapping on golang and general actions generated for them
	Entities []Module3Entity `yaml:"entities,omitempty" json:"entities,omitempty" jsonschema:"description=List of entities that module contains. Entities are basically tables in database with their mapping on golang and general actions generated for them"`

	// Tasks are actions which are triggered by a queue message or a cron job.
	Tasks []*Module3Task `yaml:"tasks,omitempty" json:"tasks,omitempty" jsonschema:"description=Tasks are actions which are triggered by a queue message or a cron job."`

	// Dtos are basically golang structs with some additional functionality which can be used for request/response actions
	Dto []Module3Dto `yaml:"dtos,omitempty" json:"dtos,omitempty" jsonschema:"description=Dtos are basically golang structs with some additional functionality which can be used for request/response actions"`

	// Actions are similar to controllers in other frameworks. They are custom functionality available via CLI or Http requests and developer need to implement their logic
	Actions []*Module3Action `yaml:"actions,omitempty" json:"actions,omitempty" jsonschema:"description=Actions are similar to controllers in other frameworks. They are custom functionality available via CLI or Http requests and developer need to implement their logic"`

	// Macros are extra definition or templates which will modify the module and able to add extra fields or tables before the codegen occures.
	Macros []Module3Macro `yaml:"macros,omitempty" json:"macros,omitempty" jsonschema:"description=Macros are extra definition or templates which will modify the module and able to add extra fields or tables before the codegen occures."`

	// Remotes are definition of external services which could be contacted via http and Fireback developer can make them typesafe by defining them here.
	Remotes []*Module3Remote `yaml:"remotes,omitempty" json:"remotes,omitempty" jsonschema:"description=Remotes are definition of external services which could be contacted via http and Fireback developer can make them typesafe by defining them here."`

	// Queries are set of SQL queries that developer writes and Fireback generates tools for fetching them from database to golang code.
	Queries []*Module3Query `yaml:"queries,omitempty" json:"queries,omitempty" jsonschema:"description=Queries are set of SQL queries that developer writes and Fireback generates tools for fetching them from database to golang code."`

	// An interesting way of defining env variables
	Config []*Module3ConfigField `yaml:"config,omitempty" json:"config,omitempty" jsonschema:"description=An interesting way of defining env variables"`

	// Messages are translatable strings which will be used as errors and other types of messages and become automatically picked via user locale.
	Messages Module3Message `yaml:"messages,omitempty" json:"messages,omitempty" jsonschema:"description=Messages are translatable strings which will be used as errors and other types of messages and become automatically picked via user locale."`
}

// Trigger is an automatic mechanism of task to be automatically run
// At the moment cron jobs are the only supported method.
type Module3Trigger struct {

	// The 5-6 star standard cronjob described in https://en.wikipedia.org/wiki/Cron
	Cron string `yaml:"cron,omitempty" json:"cron,omitempty" jsonschema:"description=The 5-6 star standard cronjob described in https://en.wikipedia.org/wiki/Cron"`
}

// Task is a general function or similarly Fireback Action, which has no results
// and could be run via Queue services or cronjobs
// Developer needs to implement the functionality manually, Fireback only generates the func signature
// Tasks are only available internally and not exported via http or client sdks
type Module3Task struct {

	// List of triggers such as cronjobs which can make this task run automatically.
	Triggers []Module3Trigger `yaml:"triggers,omitempty" json:"triggers,omitempty" jsonschema:"description=List of triggers such as cronjobs which can make this task run automatically."`

	// Name of the task is general identifier and golang functions will be generated based on it.
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the task is general identifier and golang functions will be generated based on it."`

	// Description of the task useful for developers and generated documentations.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description of the task useful for developers and generated documentations."`

	// Parameters that can be sent to this task. Since tasks are runnable in the golang as well
	// they can get parameters in go and cli if necessary. For cronjobs might make no sense.
	In *Module3ActionBody `yaml:"in,omitempty" json:"in,omitempty" jsonschema:"description=Parameters that can be sent to this task. Since tasks are runnable in the golang as well they can get parameters in go and cli if necessary. For cronjobs might make no sense."`
}

// This is a fireback remote definition, you can make the external API calls typesafe using
// definitions. This feature is documented in docs/remotes.md
type Module3Remote struct {
	// Remote action name, it will become the Golang function that you will call
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Remote action name, it will become the Golang function that you will call"`

	// Standard HTTP methods
	Method string `yaml:"method,omitempty" json:"method,omitempty" jsonschema:"enum=get,enum=post,enum=put,enum=delete,enum=patch,enum=options,enum=head,description=Standard HTTP methods"`

	// The url which will be requested. You need to add full url here, but maybe you could add a prefix
	// also in the client from your Go code - There might be a prefix for remotes later version of fireback
	Url string `yaml:"url,omitempty" json:"url,omitempty" jsonschema:"description=The url which will be requested. You need to add full url here, but maybe you could add a prefix also in the client from your Go code - There might be a prefix for remotes later version of fireback"`

	// Standard Module3ActionBody object. Could have fields, entity, dto as content and you
	// can define the output to cast automatically into them. If the response could be different objects, add them all
	// and create custom dtos and manually map them
	Out *Module3ActionBody `yaml:"out,omitempty" json:"out,omitempty" jsonschema:"description=Standard Module3ActionBody object. Could have fields, entity, dto as content and you can define the output to cast automatically into them. If the response could be different objects, add them all and create custom dtos and manually map them."`

	// Standard Module3ActionBody object. Could have fields, entity, dto as content and you
	// can define the input parameters as struct in Go and fireback will convert it into
	// json.
	In *Module3ActionBody `yaml:"in,omitempty" json:"in,omitempty" jsonschema:"description=Standard Module3ActionBody object. Could have fields entity dto as content and you can define the input parameters as struct in Go and fireback will convert it into json."`

	// Query params for the address if you want to define them in Golang dynamically instead of URL
	Query []*Module3Field `yaml:"query,omitempty" json:"query,omitempty" jsonschema:"description=Query params for the address if you want to define them in Golang dynamically instead of URL."`
}

// Used in Module3Field as the definition of enum items
type Module3Enum struct {
	// Enum key which will be used in golang generation and validation
	Key string `yaml:"k,omitempty" json:"k,omitempty" jsonschema:"description=Enum key which will be used in golang generation and validation"`

	// Description of the enum for developers. It's not translated or meant to be shown to end users.
	Descrtipion string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description of the enum for developers. It's not translated or meant to be shown to end users."`
}

// Macros is a pre-compile mechanism in Fireback, and it will modify the module definition
// before it's given to the compiler. The idea is for example, you can add extra entities
// on some modules with it.
// Until version 1.1.28, there is a single macro for EAV database model, which would create
// All of the necessary tables and fields.
// Custom macros can be indefintely useful, but need to be very well defined and documented since
// the parameters are interface{}
type Module3Macro struct {

	// The macro name which you are using. Fireback developers need to add the macros name here as reference.
	Using string `yaml:"using,omitempty" json:"using,omitempty" jsonschema:"enum=eav,description=The macro name which you are using. Fireback developers need to add the macros name here as reference."`

	// Params are the macro configuration which are dynamically set based on each macro itself.
	// They will be passed as interface{} to macro and function itself will decide what to do next.
	Params interface{} `yaml:"params,omitempty" json:"params,omitempty" jsonschema:"description=Params are the macro configuration which are dynamically set based on each macro itself. They will be passed as interface{} to macro and function itself will decide what to do next."`
}

func ConvertParams(params interface{}) interface{} {
	if params == nil {
		return nil
	}

	// Handle map[interface{}]interface{} case
	if rawMap, ok := params.(map[interface{}]interface{}); ok {
		converted := make(map[string]interface{})
		for k, v := range rawMap {
			if key, isString := k.(string); isString {
				converted[key] = v
			}
		}
		return converted
	}
	return params
}

// Useful for calling when writing a custom macro.
func Module3MacroCastParams[T any](m *Module3Macro) (*T, error) {
	m.Params = ConvertParams(m.Params)

	data, err := json.Marshal(m.Params)
	if err != nil {
		return nil, err
	}

	var result T
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type Module3FieldMatch struct {

	// The dto name from Fireback which will be matched. Might be also work with any other go struct but check the generated code.
	Dto *string `yaml:"dto,omitempty" json:"dto,omitempty" jsonschema:"description=The dto name from Fireback which will be matched. Might be also work with any other go struct but check the generated code."`
}

// Used in Module code generation to customized the generated code for gorm tags on Fireback
// Data management fields such as workspace or user id. For example, you can add extra indexes on these
// fields.
type GormOverrideMap struct {

	// Override the workspace id configuration for gorm instead of default config. Useful for adding extra constraints or indexes.
	WorkspaceId string `yaml:"workspaceId,omitempty" json:"workspaceId,omitempty" jsonschema:"description=Override the workspace id configuration for gorm instead of default config. Useful for adding extra constraints or indexes."`

	// Override the user id configuration for gorm instead of default config. Useful for adding extra constraints or indexes.
	UserId string `yaml:"userId,omitempty" json:"userId,omitempty" jsonschema:"description=Override the user id configuration for gorm instead of default config. Useful for adding extra constraints or indexes."`
}

// Permission is an access key to limit the usages of a feature.
type Module3Permission struct {
	// Name of the permission which will be used in golang and external ui
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the permission which will be used in golang and external ui"`

	// Key of the permission, separated with dots such as root.feature.action
	Key string `yaml:"key,omitempty" json:"key,omitempty" jsonschema:"description=Key of the permission, separated with dots such as root.feature.action"`

	// Description of the permission for developers or users. Not translated at this moment.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description of the permission for developers or users. Not translated at this moment."`
}

type Module3Message map[string]map[string]string

type Module3DataFields struct {

	// Essential is a set of the fields which fireback uses to give userId and workspaceId
	Essentials bool `yaml:"essentials,omitempty" json:"essentials,omitempty" jsonschema:"default=true,description=Essential is a set of the fields which fireback uses to give userId and workspaceId"`

	// Adds a int primary key auto increment
	PrimaryId bool `yaml:"primaryId,omitempty" json:"primaryId,omitempty" jsonschema:"default=true,description=Adds a int primary key auto increment"`

	// adds created - updated - delete as nano seconds to the database
	NumericTimestamp bool `yaml:"numericTimestamp,omitempty" json:"numericTimestamp,omitempty" jsonschema:"default=true,description=adds created - updated - delete as nano seconds to the database"`

	// adds created - updated - deleted fields as timestamps
	DateTimestamp bool `yaml:"dateTimestamp,omitempty" json:"dateTimestamp,omitempty" jsonschema:"default=false,description=adds created - updated - deleted fields as timestamps"`
}

// Used to adjust the features generated for each entity.
type Module3EntityFeatures struct {

	// Adds a CLI task to automatically generate mocks.
	// Disable this if it is not relevant for the feature
	// or if validations are required, making a custom mock necessary.
	// IMPORTANT: This is a pointer because it is enabled by default.
	Mock *bool `yaml:"mock,omitempty" json:"mock,omitempty" jsonschema:"Adds a CLI task to automatically generate mocks Disable this if it is not relevant for the feature or if validations are required making a custom mock necessary IMPORTANT This is a pointer because it is enabled by default"`

	// Enables embedded mock files for an entity.
	// Disables the 'msync' and 'mlist' commands.
	MSync *bool `yaml:"msync,omitempty" json:"msync,omitempty" jsonschema:"Enables embedded mock files for an entity disabling the 'msync' and 'mlist' commands"`
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
	Permissions []Module3Permission `yaml:"permissions,omitempty" json:"permissions,omitempty" jsonschema:"description=Extra permissions that an entity might need. You can add extra permissions that you will need in your business logic related to entity in itself to make it easier become as a group and document later"`

	// Actions or extra actions (on top of default actions which automatically is generated) these are
	// the same actions that you can define for a module, but defining them on entity level make it easier
	// to relate them and group them. Also permission might be added automatically (need to clearify)
	Actions []*Module3Action `yaml:"actions,omitempty" json:"actions,omitempty" jsonschema:"description=Actions or extra actions (on top of default actions which automatically is generated) these are the same actions that you can define for a module but defining them on entity level make it easier to relate them and group them. Also permission might be added automatically (need to clearify)"`

	// The entity name is crucial as it determines database table names and is used by Fireback's Go and code generation tools; note that changing an entity name does not delete previously created entities requiring manual file deletion and only camelCase naming is supported.
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=The entity name is crucial as it determines database table names and is used by Fireback's Go and code generation tools; note that changing an entity name does not delete previously created entities requiring manual file deletion and only camelCase naming is supported."`

	// You can make sure there is only one record of the entity per user or workspace using this option.
	// for example, if you want only one credit card per workspace, you can set distinctBy: workspace
	// and it will do the job
	DistinctBy string `yaml:"distinctBy,omitempty" json:"distinctBy,omitempty" jsonschema:"enum=workspace,enum=user,description=You can ensure there is only one record of the entity per user or workspace using this option for example if you want only one credit card per workspace set distinctBy: workspace and it will do the job"`

	// Customize the features generated for entity, less common  changes goes to this object
	Features Module3EntityFeatures `yaml:"features,omitempty" json:"features,omitempty" jsonschema:"description=Customize the features generated for entity, less common  changes goes to this object"`

	// Changes the default table name based on project prefix (fb_ by default) and entity name useful for times that you want to connect project to an existing database
	Table string `yaml:"table,omitempty" json:"table,omitempty" jsonschema:"description=Changes the default table name based on project prefix (fb_ by default) and entity name useful for times that you want to connect project to an existing database"`

	// Use fields allows you to customize the entity default generated fields.
	UseFields *Module3DataFields `yaml:"useFields,omitempty" json:"useFields,omitempty" jsonschema:"description=Use fields allows you to customize the entity default generated fields."`

	// Manages the entity models
	SecurityModel *EntitySecurityModel `yaml:"security,omitempty" json:"security,omitempty" jsonschema:"description=Manages the entity models"`

	// Adds a golang code to the geenrated code in very top location of the file after imports and before any code.
	PrependScript string `yaml:"prependScript,omitempty" json:"prependScript,omitempty" jsonschema:"description=Adds a golang code to the geenrated code in very top location of the file after imports and before any code."`

	// Messages are translatable strings which will be used as errors and other types of messages and become automatically picked via user locale.
	Messages Module3Message `yaml:"messages,omitempty" json:"messages,omitempty" jsonschema:"description=Messages are translatable strings which will be used as errors and other types of messages and become automatically picked via user locale."`

	// Adds a extra code before the create action in the entity. This is pure golang code.
	// Use it with caution such meta codes make module unreadable overtime. You can add script on non-dyno file of the entity.
	PrependCreateScript string `yaml:"prependCreateScript,omitempty" json:"prependCreateScript,omitempty" jsonschema:"description=Adds a extra code before the create action in the entity. This is pure golang code. Use it with caution such meta codes make module unreadable overtime. You can add script on non-dyno file of the entity."`

	// Adds a extra code before the update action in the entity. This is pure golang code.
	// Use it with caution such meta codes make module unreadable overtime. You can add script on non-dyno file of the entity.
	PrependUpdateScript string `yaml:"prependUpdateScript,omitempty" json:"prependUpdateScript,omitempty" jsonschema:"description=Adds a extra code before the update action in the entity. This is pure golang code. Use it with caution such meta codes make module unreadable overtime. You can add script on non-dyno file of the entity."`

	// Access is a method of limiting which type offunctionality will be created for the entity. For example access read will remove all create functionality from code and public API.
	Access string `yaml:"access,omitempty" json:"access,omitempty" jsonschema:"description=Access is a method of limiting which type offunctionality will be created for the entity. For example access read will remove all create functionality from code and public API."`

	// For entities, if the query scope is public the query action will become automatically public and without authentication
	QueryScope string `yaml:"queryScope,omitempty" json:"queryScope,omitempty" jsonschema:"enum=public,enum=specific,description=For entities, if the query scope is public the query action will become automatically public and without authentication"`

	// A list of extra queries that Fireback can generate for the the entity. Fireback might offer some extra queries to be generated so they will be listed here.
	Queries []string `yaml:"queries,omitempty" json:"queries,omitempty" jsonschema:"description=A list of extra queries that Fireback can generate for the the entity. Fireback might offer some extra queries to be generated so they will be listed here."`

	// Override the some default Fireback generated fields gorm configuration.
	GormMap GormOverrideMap `yaml:"gormMap,omitempty" json:"gormMap,omitempty" jsonschema:"description=Override the some default Fireback generated fields gorm configuration."`

	// Define the fields that this entity will have both in golang and database columns.
	Fields []*Module3Field `yaml:"fields,omitempty" json:"fields,omitempty" jsonschema:"description=Define the fields that this entity will have both in golang and database columns."`

	// The name of the entity which will appear in CLI. By default the name of the entity will be used with dashes.
	CliName string `yaml:"cliName,omitempty" json:"cliName,omitempty" jsonschema:"description=The name of the entity which will appear in CLI. By default the name of the entity will be used with dashes."`

	// The alternative shortcut in the CLI. By default it's empty and only the entity name or CliName.
	CliShort string `yaml:"cliShort,omitempty" json:"cliShort,omitempty" jsonschema:"description=The alternative shortcut in the CLI. By default it's empty and only the entity name or CliName."`

	// Description about the purpose of the entity. It will be used in CLI and codegen documentation.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description about the purpose of the entity. It will be used in CLI and codegen documentation."`

	// CTE is a common recursive feature of an entity; enabling it generates SQL for recursive parent-child CTE queries and makes it available in Golang.
	Cte bool `yaml:"cte,omitempty" json:"cte,omitempty" jsonschema:"description=CTE is a common recursive feature of an entity; enabling it generates SQL for recursive parent-child CTE queries and makes it available in Golang."`

	// The name of the golang function which will recieve entity pointer to make some modification
	// upon query, get or other details.
	PostFormatter string `yaml:"postFormatter,omitempty" json:"postFormatter,omitempty" jsonschema:"description=The name of the golang function which will recieve entity pointer to make some modification upon query, get or other details."`
}

// Represents a dto in an application. Can be used for variety of reasons,
// request response of an action, or even internally. Fireback generates bunch of
// helpers for each dto, so it might make sense to define them in Module3 instead
// of pure struct in golang.
type Module3Dto struct {

	// Name of the dto, in camel case, the rest of the code related to this dto is being generated based on this
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the dto in camel case the rest of the code related to this dto is being generated based on this"`

	// List of fields and body definitions of the dto
	Fields []*Module3Field `yaml:"fields,omitempty" json:"fields,omitempty" jsonschema:"description=List of fields and body definitions of the dto"`
}

// Represents any action request or response DTO.
// Used in multiple places as the request/response signature.
type Module3ActionBody struct {

	// Defines the fields directly, and DTO will be generated
	// and assigned automatically.
	Fields []*Module3Field `yaml:"fields,omitempty" json:"fields,omitempty" jsonschema:"Defines the fields directly and DTO will be generated and assigned automatically"`

	// Selects the DTO existing in the module from Golang.
	// It can also be a pure Go struct, but those do not compile.
	Dto string `yaml:"dto,omitempty" json:"dto,omitempty" jsonschema:"Selects the DTO existing in the module from Golang It can also be a pure Go struct but those do not compile"`

	// Use a entity which is generated by Fireback instead.
	Entity string `yaml:"entity,omitempty" json:"entity,omitempty" jsonschema:"Generates the entity name in the module"`

	// Uses a primitive type instead, such as a string or int64.
	Primitive string `yaml:"primitive,omitempty" json:"primitive,omitempty" jsonschema:"Uses a primitive type instead such as a string or int64"`
}

type Module3Action struct {

	// General name of the action used for generating code and CLI commands.
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=General name of the action used for generating code and CLI commands"`

	// Overrides the CLI action name if specified otherwise defaults to Name.
	CliName string `yaml:"cliName,omitempty" json:"cliName,omitempty" jsonschema:"description=Overrides the CLI action name if specified otherwise defaults to Name"`

	// CLI command aliases for shorter action names.
	ActionAliases []string `yaml:"actionAliases,omitempty" json:"actionAliases,omitempty" jsonschema:"description=CLI command aliases for shorter action names"`

	// HTTP route of the action; if not specified the action is CLI-only.
	Url string `yaml:"url,omitempty" json:"url,omitempty" jsonschema:"description=HTTP route of the action; if not specified the action is CLI-only"`

	// HTTP method type including standard and Fireback-specific methods.
	Method string `yaml:"method,omitempty" json:"method,omitempty" jsonschema:"enum=post,enum=get,enum=delete,enum=reactive,description=HTTP method type including standard and Fireback-specific methods"`

	// Type-safe query parameters for CLI and HTTP requests.
	Query []*Module3Field `yaml:"query,omitempty" json:"query,omitempty" jsonschema:"description=Type-safe query parameters for CLI and HTTP requests"`

	// Action description used in API specs and documentation.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Action description used in API specs and documentation"`

	// Higher-level request format such as POST_ONE PATCH_ONE, QUERY, and PATCH_BULK.
	Format string `yaml:"format,omitempty" json:"format,omitempty" jsonschema:"enum=reactive,enum=query,description=Higher-level request format such as POST_ONE PATCH_ONE, QUERY, and PATCH_BULK"`

	// Request body definition similar to HTTP request body.
	In *Module3ActionBody `yaml:"in,omitempty" json:"in,omitempty" jsonschema:"description=Request body definition similar to HTTP request body"`

	// Response body definition similar to HTTP response body.
	Out *Module3ActionBody `yaml:"out,omitempty" json:"out,omitempty" jsonschema:"description=Response body definition similar to HTTP response body"`

	// Defines access control similar to middleware checking permissions, tokens, and roles.
	SecurityModel *SecurityModel `yaml:"security,omitempty" json:"security,omitempty" jsonschema:"description=Defines access control similar to middleware checking permissions, tokens, and roles"`

	// CLI implementation of the action.
	CliAction func(c *cli.Context, security *SecurityModel) error `jsonschema:"-"`

	// HTTP implementation using Gin handlers.
	Handlers []gin.HandlerFunc `yaml:"-" json:"-" jsonschema:"-"`

	// CLI action flags similar to HTTP request body fields.
	Flags []cli.Flag `yaml:"-" json:"-" jsonschema:"-"`

	// External function name used in generated code.
	ExternFuncName string `yaml:"-" json:"-" jsonschema:"-"`

	// Struct representing the request body used for generating RPC code.
	RequestEntity any `yaml:"-" json:"-" jsonschema:"-"`

	// Struct representing the response body used for generating RPC code.
	ResponseEntity any `yaml:"-" json:"-" jsonschema:"-"`

	// Function representing the action's implementation.
	Action any `yaml:"-" json:"-" jsonschema:"-"`

	// Pointer to the struct representing the entity being operated on.
	TargetEntity any `yaml:"-" json:"-" jsonschema:"-"`

	// Internal metadata for code generation.
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

	if x.UseFields == nil {
		data = Module3DataFields{
			Essentials:       true,
			PrimaryId:        true,
			NumericTimestamp: true,
			DateTimestamp:    false,
		}

		return data
	} else {
		data = *x.UseFields
	}

	return data
}
