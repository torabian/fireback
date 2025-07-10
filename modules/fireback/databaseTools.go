//go:build !cgo

package fireback

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

func GetSQLiteDialector(dsn string) gorm.Dialector {
	fmt.Println("Called")
	return sqlite.Open(dsn)
}
