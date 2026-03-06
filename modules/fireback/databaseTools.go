//go:build !cgo

package fireback

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func GetSQLiteDialector(dsn string) gorm.Dialector {

	return sqlite.Open(dsn)
}
