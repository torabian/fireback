package queries

import "embed"

//go:embed *.sql *.vsql
var QueriesFs embed.FS
