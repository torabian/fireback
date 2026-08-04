package migrations

import "embed"

//go:embed *.sql
var MigrationsFs embed.FS
