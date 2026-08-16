package fireback

var DATABASE_TYPE_MYSQL string = "mysql"
var DATABASE_TYPE_SQLITE string = "sqlite"
var DATABASE_TYPE_SQLITE_MEMORY string = "sqlite (:memory:)"
var DATABASE_TYPE_POSTGRES string = "postgres"
var DATABASE_TYPE_MARIADB string = "mariadb"

// Database is the interactive setup wizard's scratch state while it collects
// a connection one piece at a time (host, port, credentials, database name,
// ...). Only Vendor and Dsn ever get persisted to Config - see
// modules/fireback/dbdsn, which is what turns the rest of these fields into
// Dsn (and, going the other way, what recovers them from an existing Dsn to
// prefill the wizard on a re-run).
type Database struct {
	Username string `yaml:"username,omitempty"`
	Port     string `yaml:"port,omitempty"`
	Host     string `yaml:"host,omitempty"`
	Password string `yaml:"password,omitempty"`
	Database string `yaml:"database,omitempty"`
	Vendor   string `yaml:"vendor,omitempty"`
	Dsn      string `yaml:"dsn,omitempty"`
	// SSL is only meaningful for postgres today - see dbdsn.ConnectionInfo.
	SSL bool `yaml:"ssl,omitempty"`
}
