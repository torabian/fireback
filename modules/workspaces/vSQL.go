package workspaces

// Module3Query represents an SQL query configuration used by Fireback.
// Fireback will generate Golang functions to run the queries and will
// replace placeholders such as the current user and workspaces in the SQL query.
type Module3Query struct {

	// Name is the identifier for the query. It will be used to generate controller
	// code and should uniquely identify the query.
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	// Description provides a detailed explanation of the query. It helps other
	// developers or API consumers understand what the query does and its purpose.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	// Affects indicates whether the query modifies the database. This is a
	// safety flag to signal if the query performs insert, update, or delete
	// operations. Fireback may also analyze the SQL statement to verify if
	// it contains such operations.
	Affects bool `yaml:"affects,omitempty" json:"affects,omitempty"`

	// Columns defines the structure of the result set returned by the query.
	// It lists the expected columns in the result when the query is executed.
	Columns *Module3ActionBody `yaml:"columns,omitempty" json:"columns,omitempty"`
}
