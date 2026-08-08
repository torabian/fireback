package fireback

var DATABASE_TYPE_MYSQL string = "mysql"
var DATABASE_TYPE_SQLITE string = "sqlite"
var DATABASE_TYPE_SQLITE_MEMORY string = "sqlite (:memory:)"
var DATABASE_TYPE_POSTGRES string = "postgres"
var DATABASE_TYPE_MARIADB string = "mariadb"

type Database struct {
	Username string `yaml:"username,omitempty"`
	Port     string `yaml:"port,omitempty"`
	Host     string `yaml:"host,omitempty"`
	Password string `yaml:"password,omitempty"`
	Database string `yaml:"database,omitempty"`
	Vendor   string `yaml:"vendor,omitempty"`
	Dsn      string `yaml:"dsn,omitempty"`
}
