//go:build cgo

package workspaces

import (

	// This is go version (glebarez), use it for desktop and server, and use gorm.io/driver/sqlite for mobile
	// on Android and IOS the golang version does not work. On the server, it's better not use cgo
	// to make it portable on more operating systems.

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetSQLiteDialector(dsn string) gorm.Dialector {
	return sqlite.Open(dsn)
}
