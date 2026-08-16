//go:build !wasm

package clitools

import (
	"net"
	"testing"

	"github.com/torabian/fireback/modules/fireback"
)

// withConfig temporarily overrides the shared config's DbVendor/DbDsn for a
// test, restoring the previous values afterward - config is the same
// package-level pointer production code reads, so tests must not leak
// changes into each other.
func withConfig(t *testing.T, vendor string, dsn string, fn func()) {
	t.Helper()
	prevVendor, prevDsn := config.DbVendor, config.DbDsn
	config.DbVendor, config.DbDsn = vendor, dsn
	defer func() { config.DbVendor, config.DbDsn = prevVendor, prevDsn }()
	fn()
}

func TestDefaultDatabasePort(t *testing.T) {
	cases := map[string]string{
		fireback.DATABASE_TYPE_POSTGRES: "5432",
		fireback.DATABASE_TYPE_MYSQL:    "3306",
		fireback.DATABASE_TYPE_MARIADB:  "3306",
		"":                              "3306",
	}
	for vendor, want := range cases {
		if got := defaultDatabasePort(vendor); got != want {
			t.Errorf("defaultDatabasePort(%q) = %q, want %q", vendor, got, want)
		}
	}
}

func TestIndexOf(t *testing.T) {
	items := []string{"a", "b", "c"}

	if got := indexOf(items, "b"); got != 1 {
		t.Errorf("indexOf(items, %q) = %d, want 1", "b", got)
	}
	if got := indexOf(items, "missing"); got != 0 {
		t.Errorf("indexOf(items, %q) = %d, want 0 (not found falls back to first item)", "missing", got)
	}
	if got := indexOf(items, ""); got != 0 {
		t.Errorf("indexOf(items, \"\") = %d, want 0", got)
	}
}

func TestOrDefault(t *testing.T) {
	if got := orDefault("value", "fallback"); got != "value" {
		t.Errorf("orDefault with a non-empty value = %q, want %q", got, "value")
	}
	if got := orDefault("", "fallback"); got != "fallback" {
		t.Errorf("orDefault with an empty value = %q, want %q", got, "fallback")
	}
}

func TestConfiguredDsn(t *testing.T) {
	withConfig(t, fireback.DATABASE_TYPE_POSTGRES, "host=localhost port=5432 user=postgres dbname=app sslmode=disable", func() {
		if got := configuredDsn(fireback.DATABASE_TYPE_POSTGRES); got != config.DbDsn {
			t.Errorf("configuredDsn for the currently configured vendor = %q, want %q", got, config.DbDsn)
		}
		if got := configuredDsn(fireback.DATABASE_TYPE_MYSQL); got != "" {
			t.Errorf("configuredDsn for a different vendor = %q, want empty - a postgres dsn is not a mysql default", got)
		}
	})
}

func TestConfiguredDbDefaults(t *testing.T) {
	withConfig(t, fireback.DATABASE_TYPE_POSTGRES, "host=db.internal port=5432 user=app password=secret dbname=app sslmode=require", func() {
		defaults := configuredDbDefaults(fireback.DATABASE_TYPE_POSTGRES)
		if defaults.Host != "db.internal" || defaults.Port != "5432" || defaults.Username != "app" || defaults.Password != "secret" || defaults.Database != "app" || !defaults.SSL {
			t.Errorf("configuredDbDefaults(postgres) = %+v, want the parsed dsn pieces", defaults)
		}

		// A vendor switch must not leak the previous vendor's connection
		// details into the new vendor's prompts.
		if mysqlDefaults := configuredDbDefaults(fireback.DATABASE_TYPE_MYSQL); mysqlDefaults.Host != "" {
			t.Errorf("configuredDbDefaults(mysql) while configured for postgres = %+v, want zero value", mysqlDefaults)
		}
	})

	withConfig(t, "", "", func() {
		if defaults := configuredDbDefaults(fireback.DATABASE_TYPE_POSTGRES); defaults.Host != "" {
			t.Errorf("configuredDbDefaults with nothing configured = %+v, want zero value", defaults)
		}
	})
}

func TestPostgresAdminDsn(t *testing.T) {
	got := postgresAdminDsn("localhost", "5432", "postgres", "", false)
	want := "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable"
	if got != want {
		t.Errorf("postgresAdminDsn(..., ssl=false) = %q, want %q", got, want)
	}

	got = postgresAdminDsn("localhost", "5432", "postgres", "secret", true)
	want = "host=localhost port=5432 user=postgres password=secret dbname=postgres sslmode=require"
	if got != want {
		t.Errorf("postgresAdminDsn(..., ssl=true) = %q, want %q", got, want)
	}
}

func TestMysqlAdminDsn(t *testing.T) {
	got := mysqlAdminDsn("localhost", "3306", "root", "secret")
	want := "root:secret@tcp(localhost:3306)/?charset=utf8mb4&parseTime=True&loc=Local"
	if got != want {
		t.Errorf("mysqlAdminDsn(...) = %q, want %q", got, want)
	}
}

func TestPgQuoteIdentifier(t *testing.T) {
	if got := pgQuoteIdentifier("app"); got != `"app"` {
		t.Errorf("pgQuoteIdentifier(%q) = %q, want %q", "app", got, `"app"`)
	}
	if got := pgQuoteIdentifier(`weird"name`); got != `"weird""name"` {
		t.Errorf("pgQuoteIdentifier with an embedded quote = %q, want %q", got, `"weird""name"`)
	}
}

func TestDetectDefaultDatabasePort(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("could not open a local listener for the test: %v", err)
	}
	_, port, _ := net.SplitHostPort(listener.Addr().String())

	if !detectDefaultDatabasePort("127.0.0.1", port) {
		t.Errorf("detectDefaultDatabasePort did not find the listener on 127.0.0.1:%s", port)
	}

	listener.Close()

	if detectDefaultDatabasePort("127.0.0.1", port) {
		t.Errorf("detectDefaultDatabasePort reported something listening on 127.0.0.1:%s after it was closed", port)
	}
}
