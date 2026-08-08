package fireback

import (
	"bytes"
	"embed"
	"text/template"

	"github.com/gin-gonic/gin"
)

/*
*	This is tools for parsing a nativesql file, they are normal sql
*	Files but with golang template expression, and some extra special anotation
*	To taget specific vendor or version
 */

// Information about the sql vendor, and version,
// Or even other stuff over time.
// Template will be modified based on this
type SqlExecuteContext struct {
	Vendor string
}

func (x *SqlExecuteContext) Mysql() bool {
	if x.Vendor == "mysql" {
		return true
	}

	return false
}

func (x *SqlExecuteContext) Sqlite() bool {
	if x.Vendor == "sqlite" {
		return true
	}

	return false
}

// This is a new implementation of the native query.
// It aims to make sql file compatible with all possible different sql dialects
// Basically it's an sql file, writen in golang template.
func NativeQuery[T any](fsRef *embed.FS, queryName string, query QueryDSL, values ...interface{}) ([]*T, *QueryResultMeta, error) {
	qrm := &QueryResultMeta{
		TotalItems:          -1,
		TotalAvailableItems: -1,
	}

	ctx := SqlExecuteContext{
		Vendor: config.DbVendor,
	}

	result, err := NativeQueryResolver(ctx, fsRef, queryName, query, values)

	if err != nil {
		return []*T{}, qrm, err
	}

	return UnsafeQuerySql[T](result.SqlBody, result.SqlCounter, query, "", values...)
}

type ResolvedSql struct {
	SqlBody    string
	SqlCounter string
}

func CompileString(fs *embed.FS, fname string, params gin.H) (string, error) {
	t, err := template.New("").Funcs(CommonMap).ParseFS(fs, fname)
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer

	err = t.ExecuteTemplate(&tpl, fname, params)

	return tpl.String(), err
}

func NativeQueryResolver(ctx SqlExecuteContext, fsRef *embed.FS, queryName string, query QueryDSL, values ...interface{}) (*ResolvedSql, error) {
	result := ResolvedSql{}

	sqlQuery, readSqlSourceError := CompileString(fsRef, queryName, gin.H{
		"ctx":     &ctx,
		"counter": false,
	})
	if readSqlSourceError != nil {
		return nil, GormErrorToIError(readSqlSourceError)
	}

	sqlQueryCounter, readSqlCounterSourceError := CompileString(fsRef, queryName, gin.H{
		"ctx":     &ctx,
		"counter": true,
	})
	if readSqlCounterSourceError != nil {
		return nil, GormErrorToIError(readSqlCounterSourceError)
	}

	result.SqlBody = sqlQuery
	result.SqlCounter = sqlQueryCounter

	return &result, nil
}
