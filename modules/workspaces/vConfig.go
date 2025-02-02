package workspaces

// Module3ConfigField represents a configuration field, typically a variable definition,
// which is converted into a Go struct. It provides functionality to read the value from
// YAML, environment variables, and CLI flags, with the option to save it as a .env file.
// This struct offers a type-safe way to use environment variables while documenting
// them in a YAML configuration file.
type Module3ConfigField struct {

	// Name is the identifier for the configuration field, used both in Go code and
	// the environment file. By default, the name will be converted to uppercase with
	// underscores to reference the environment variable, unless overridden by the 'env' field.
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name is the identifier for the configuration field used both in Go code and the environment file. By default the name will be converted to uppercase with underscores to reference the environment variable unless overridden by the 'env' field."`

	// Type defines the data type for the environment variable. It supports standard Go types
	// such as string, bool, int64, and others, along with custom Fireback types.
	// Ensure that the chosen type is supported.
	Type string `yaml:"type,omitempty" json:"type,omitempty" jsonschema:"description=Type defines the data type for the environment variable. It supports standard Go types such as string - bool - int64 - and others - along with custom Fireback types. Ensure that the chosen type is supported."`

	// Description explains the purpose of the configuration field. It can be helpful for developers
	// and also used in CLI for interactive configuration.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description explains the purpose of the configuration field. It can be helpful for developers and also used in CLI for interactive configuration."`

	// Default specifies the default value for the configuration field if it is not defined.
	Default string `yaml:"default,omitempty" json:"default,omitempty" jsonschema:"description=Default specifies the default value for the configuration field if it is not defined."`

	// Env allows you to override the default environment variable name, which is automatically
	// generated from the Name field. Use this field if you want to manually specify the environment variable name.
	Env string `yaml:"env,omitempty" json:"env,omitempty" jsonschema:"description=Env allows you to override the default environment variable name, which is automatically generated from the Name field. Use this field if you want to manually specify the environment variable name."`

	// Fields defines child configuration fields in case the current field represents an object
	// or an array of subfields. Note that support for nested fields may be limited.
	Fields []Module3ConfigField `yaml:"fields,omitempty" json:"fields,omitempty" jsonschema:"description=Fields defines child configuration fields in case the current field represents an object or an array of subfields. Note that support for nested fields may be limited."`
}
