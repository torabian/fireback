package izaTheme

import "database/sql"

type Manifest struct {
	DB            *sql.DB
	FilterResolver func(string) (string, error)
}


