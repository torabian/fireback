// Package dbdsn is fireback's single place for turning a database
// connection string into its component pieces and back, one vendor at a
// time.
//
// Fireback used to store a database connection two ways at once: a DB_DSN
// string, and a parallel set of DB_HOST/DB_PORT/DB_USERNAME/DB_PASSWORD/
// DB_NAME fields that got assembled into a DSN on the fly whenever DB_DSN
// itself was empty (see GetDatabaseDsn's history). That meant two sources
// of truth that could disagree, and every caller that wanted individual
// pieces (the interactive setup wizard, the backup module's Postgres
// fallback) either read the scattered fields directly or hand-rolled its
// own DSN parsing.
//
// Now DB_DSN is the only thing fireback persists for a database connection
// (DB_VENDOR alongside it, to know which format the DSN is even in). This
// package is where anything that needs the individual pieces - to prefill a
// prompt, to build a libpq/go-sql-driver connection string one field at a
// time - does that parsing/building, exactly once per vendor.
package dbdsn

import (
	"fmt"
	"regexp"
	"strings"
)

// Vendor names, matching fireback.DATABASE_TYPE_* (duplicated here rather
// than imported, so this package stays a leaf with no dependency on the
// rest of fireback - see the package doc in modules/fireback/DatabaseCommon.go).
const (
	VendorPostgres = "postgres"
	VendorMysql    = "mysql"
	VendorMariadb  = "mariadb"
	VendorSqlite   = "sqlite"
)

// ConnectionInfo is the parsed-out form of a Postgres or MySQL/MariaDB DSN:
// the individual pieces a human types into the setup wizard, or that a
// fallback connection (backup module's PGHOST/PGPORT/... derivation) needs
// discretely rather than as one opaque string.
//
// Sqlite has no equivalent breakdown - its "dsn" is just a file path (or
// ":memory:") - so it's never represented as ConnectionInfo; use the dsn
// string directly.
type ConnectionInfo struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSL      bool
}

// Parse breaks a vendor's DSN down into ConnectionInfo. It returns an error
// for sqlite (nothing to parse; use the dsn directly as a path) and for any
// vendor it doesn't recognize.
func Parse(vendor string, dsn string) (ConnectionInfo, error) {
	switch vendor {
	case VendorPostgres:
		return ParsePostgres(dsn)
	case VendorMysql, VendorMariadb:
		return ParseMysql(dsn)
	default:
		return ConnectionInfo{}, fmt.Errorf("dbdsn: no dsn parser for vendor %q", vendor)
	}
}

// Build assembles a vendor's DSN from ConnectionInfo - the inverse of Parse.
func Build(vendor string, info ConnectionInfo) (string, error) {
	switch vendor {
	case VendorPostgres:
		return BuildPostgres(info), nil
	case VendorMysql, VendorMariadb:
		return BuildMysql(info), nil
	default:
		return "", fmt.Errorf("dbdsn: no dsn builder for vendor %q", vendor)
	}
}

// ParsePostgres reads a libpq keyword/value connection string
// (`host=... user=... password=... dbname=... port=... sslmode=...`) into
// ConnectionInfo. Unknown keywords are ignored; SSL is true for any
// sslmode other than "disable" (or unset).
func ParsePostgres(dsn string) (ConnectionInfo, error) {
	info := ConnectionInfo{}

	for _, token := range splitPostgresTokens(dsn) {
		key, value, ok := strings.Cut(token, "=")
		if !ok {
			continue
		}
		value = unquotePostgresValue(value)

		switch key {
		case "host":
			info.Host = value
		case "port":
			info.Port = value
		case "user":
			info.Username = value
		case "password":
			info.Password = value
		case "dbname":
			info.Database = value
		case "sslmode":
			info.SSL = value != "" && value != "disable"
		}
	}

	return info, nil
}

// BuildPostgres is the inverse of ParsePostgres. Keywords with an empty
// value are omitted entirely rather than written as e.g. `password=` -
// pgx.ParseConfig silently drops every keyword after an empty one
// (confirmed directly: it falls back to Postgres's "no database given"
// default, the OS username, instead of the dbname actually requested),
// which matters most for passwordless/trust-auth local setups.
func BuildPostgres(info ConnectionInfo) string {
	parts := []string{}

	if info.Host != "" {
		parts = append(parts, "host="+quotePostgresValue(info.Host))
	}
	if info.Port != "" {
		parts = append(parts, "port="+quotePostgresValue(info.Port))
	}
	if info.Username != "" {
		parts = append(parts, "user="+quotePostgresValue(info.Username))
	}
	if info.Password != "" {
		parts = append(parts, "password="+quotePostgresValue(info.Password))
	}
	if info.Database != "" {
		parts = append(parts, "dbname="+quotePostgresValue(info.Database))
	}

	sslmode := "disable"
	if info.SSL {
		sslmode = "require"
	}
	parts = append(parts, "sslmode="+sslmode)

	return strings.Join(parts, " ")
}

// splitPostgresTokens splits a libpq keyword/value string on whitespace,
// respecting single-quoted values that may themselves contain spaces (the
// format libpq itself accepts for e.g. `host='my host'`).
func splitPostgresTokens(dsn string) []string {
	var tokens []string
	var current strings.Builder
	inQuotes := false

	for _, r := range dsn {
		switch {
		case r == '\'':
			inQuotes = !inQuotes
			current.WriteRune(r)
		case r == ' ' && !inQuotes:
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
		default:
			current.WriteRune(r)
		}
	}
	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}

	return tokens
}

func unquotePostgresValue(v string) string {
	if len(v) >= 2 && v[0] == '\'' && v[len(v)-1] == '\'' {
		v = v[1 : len(v)-1]
		v = strings.ReplaceAll(v, `\'`, `'`)
		v = strings.ReplaceAll(v, `\\`, `\`)
	}
	return v
}

func quotePostgresValue(v string) string {
	if v == "" {
		return "''"
	}
	if !strings.ContainsAny(v, " '\\") {
		return v
	}
	v = strings.ReplaceAll(v, `\`, `\\`)
	v = strings.ReplaceAll(v, `'`, `\'`)
	return "'" + v + "'"
}

// mysqlDsnPattern matches the go-sql-driver/mysql DSN shape fireback always
// writes: `[username[:password]]@tcp(host:port)/dbname[?params]`.
var mysqlDsnPattern = regexp.MustCompile(`^(?:([^:@]*)(?::([^@]*))?@)?tcp\(([^)]*)\)/([^?]*)(?:\?.*)?$`)

// ParseMysql reads a go-sql-driver/mysql DSN into ConnectionInfo.
func ParseMysql(dsn string) (ConnectionInfo, error) {
	m := mysqlDsnPattern.FindStringSubmatch(dsn)
	if m == nil {
		return ConnectionInfo{}, fmt.Errorf("dbdsn: not a recognized mysql dsn: %q", dsn)
	}

	info := ConnectionInfo{
		Username: m[1],
		Password: m[2],
		Database: m[4],
	}

	host, port, hasPort := strings.Cut(m[3], ":")
	info.Host = host
	if hasPort {
		info.Port = port
	}

	return info, nil
}

// BuildMysql is the inverse of ParseMysql. charset/parseTime/loc are always
// set the same way fireback has always connected - not something the
// wizard asks about - so they aren't part of ConnectionInfo.
func BuildMysql(info ConnectionInfo) string {
	address := info.Host
	if info.Port != "" {
		address += ":" + info.Port
	}

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		info.Username, info.Password, address, info.Database)
}
