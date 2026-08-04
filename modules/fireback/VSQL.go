package fireback

// Module3Query represents an SQL query configuration used by Fireback.
// Fireback will generate Golang functions to run the queries and will
// replace placeholders such as the current user and workspaces in the SQL query.
type Module3Query struct {

	// Name is the identifier for the query. It will be used to generate controller
	// code and should uniquely identify the query.
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name is the identifier for the query. It will be used to generate controller code and should uniquely identify the query."`

	// Description provides a detailed explanation of the query. It helps other
	// developers or API consumers understand what the query does and its purpose.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description provides a detailed explanation of the query. It helps other developers or API consumers understand what the query does and its purpose."`

	// Columns defines the structure of the result set returned by the query.
	// It lists the expected columns in the result when the query is executed.
	Columns *Module3ActionBody `yaml:"columns,omitempty" json:"columns,omitempty" jsonschema:"description=Columns defines the structure of the result set returned by the query. It lists the expected columns in the result when the query is executed."`

	// The actual SQL or VSQL query. There are some special placeholders and this is infact a golang template
	// which will be converted in the end to SQL and will be sent to ORM.
	Query string `yaml:"query,omitempty" json:"query,omitempty" jsonschema:"description=The actual SQL or VSQL query. There are some special placeholders and this is infact a golang template which will be converted in the end to SQL and will be sent to ORM."`
}
