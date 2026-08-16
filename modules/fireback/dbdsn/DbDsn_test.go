package dbdsn

import "testing"

func TestParsePostgres(t *testing.T) {
	cases := []struct {
		name string
		dsn  string
		want ConnectionInfo
	}{
		{
			name: "full dsn, ssl disabled",
			dsn:  "host=127.0.0.1 port=5432 user=postgres password=secret dbname=fireback sslmode=disable",
			want: ConnectionInfo{Host: "127.0.0.1", Port: "5432", Username: "postgres", Password: "secret", Database: "fireback", SSL: false},
		},
		{
			name: "ssl required",
			dsn:  "host=db.internal port=5432 user=app dbname=app sslmode=require",
			want: ConnectionInfo{Host: "db.internal", Port: "5432", Username: "app", Database: "app", SSL: true},
		},
		{
			name: "no password keyword at all (trust auth)",
			dsn:  "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable",
			want: ConnectionInfo{Host: "localhost", Port: "5432", Username: "postgres", Database: "postgres", SSL: false},
		},
		{
			name: "quoted value with a space",
			dsn:  "host=localhost port=5432 user=postgres password='my pass' dbname='my db' sslmode=disable",
			want: ConnectionInfo{Host: "localhost", Port: "5432", Username: "postgres", Password: "my pass", Database: "my db", SSL: false},
		},
		{
			name: "unknown sslmode still counts as SSL on",
			dsn:  "host=localhost port=5432 user=postgres dbname=postgres sslmode=verify-full",
			want: ConnectionInfo{Host: "localhost", Port: "5432", Username: "postgres", Database: "postgres", SSL: true},
		},
		{
			name: "empty dsn",
			dsn:  "",
			want: ConnectionInfo{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParsePostgres(tc.dsn)
			if err != nil {
				t.Fatalf("ParsePostgres(%q) returned error: %v", tc.dsn, err)
			}
			if got != tc.want {
				t.Fatalf("ParsePostgres(%q) = %+v, want %+v", tc.dsn, got, tc.want)
			}
		})
	}
}

func TestBuildPostgres(t *testing.T) {
	cases := []struct {
		name string
		info ConnectionInfo
		want string
	}{
		{
			name: "full info, ssl off",
			info: ConnectionInfo{Host: "127.0.0.1", Port: "5432", Username: "postgres", Password: "secret", Database: "fireback"},
			want: "host=127.0.0.1 port=5432 user=postgres password=secret dbname=fireback sslmode=disable",
		},
		{
			name: "ssl on",
			info: ConnectionInfo{Host: "db.internal", Port: "5432", Username: "app", Database: "app", SSL: true},
			want: "host=db.internal port=5432 user=app dbname=app sslmode=require",
		},
		{
			// The whole point of building password= only when non-empty: an
			// empty "password=" keyword/value pair trips up pgx.ParseConfig,
			// silently dropping dbname too. See the comment on BuildPostgres.
			name: "empty password is omitted, not written as password=",
			info: ConnectionInfo{Host: "localhost", Port: "5432", Username: "postgres", Database: "postgres"},
			want: "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable",
		},
		{
			name: "value with a space gets quoted",
			info: ConnectionInfo{Host: "localhost", Port: "5432", Username: "postgres", Password: "my pass", Database: "postgres"},
			want: "host=localhost port=5432 user=postgres password='my pass' dbname=postgres sslmode=disable",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := BuildPostgres(tc.info)
			if got != tc.want {
				t.Fatalf("BuildPostgres(%+v) = %q, want %q", tc.info, got, tc.want)
			}
		})
	}
}

func TestPostgresRoundTrip(t *testing.T) {
	cases := []ConnectionInfo{
		{Host: "127.0.0.1", Port: "5432", Username: "postgres", Password: "secret", Database: "fireback", SSL: false},
		{Host: "db.internal", Port: "5432", Username: "app", Database: "app", SSL: true},
		{Host: "localhost", Port: "5432", Username: "postgres", Database: "postgres"},
	}

	for _, info := range cases {
		dsn := BuildPostgres(info)
		got, err := ParsePostgres(dsn)
		if err != nil {
			t.Fatalf("ParsePostgres(%q) returned error: %v", dsn, err)
		}
		if got != info {
			t.Fatalf("round trip mismatch: built %q, parsed back %+v, want %+v", dsn, got, info)
		}
	}
}

func TestParseMysql(t *testing.T) {
	cases := []struct {
		name string
		dsn  string
		want ConnectionInfo
	}{
		{
			name: "full dsn with query params",
			dsn:  "root:secret@tcp(127.0.0.1:3306)/fireback?charset=utf8mb4&parseTime=True&loc=Local",
			want: ConnectionInfo{Host: "127.0.0.1", Port: "3306", Username: "root", Password: "secret", Database: "fireback"},
		},
		{
			name: "no password",
			dsn:  "root@tcp(localhost:3306)/fireback",
			want: ConnectionInfo{Host: "localhost", Port: "3306", Username: "root", Database: "fireback"},
		},
		{
			name: "no database selected (admin connection)",
			dsn:  "root:secret@tcp(localhost:3306)/",
			want: ConnectionInfo{Host: "localhost", Port: "3306", Username: "root", Password: "secret"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseMysql(tc.dsn)
			if err != nil {
				t.Fatalf("ParseMysql(%q) returned error: %v", tc.dsn, err)
			}
			if got != tc.want {
				t.Fatalf("ParseMysql(%q) = %+v, want %+v", tc.dsn, got, tc.want)
			}
		})
	}
}

func TestParseMysqlInvalid(t *testing.T) {
	if _, err := ParseMysql("not a dsn at all"); err == nil {
		t.Fatal("expected an error for an unrecognized mysql dsn, got nil")
	}
}

func TestBuildMysql(t *testing.T) {
	info := ConnectionInfo{Host: "127.0.0.1", Port: "3306", Username: "root", Password: "secret", Database: "fireback"}
	want := "root:secret@tcp(127.0.0.1:3306)/fireback?charset=utf8mb4&parseTime=True&loc=Local"

	if got := BuildMysql(info); got != want {
		t.Fatalf("BuildMysql(%+v) = %q, want %q", info, got, want)
	}
}

func TestMysqlRoundTrip(t *testing.T) {
	info := ConnectionInfo{Host: "127.0.0.1", Port: "3306", Username: "root", Password: "secret", Database: "fireback"}
	dsn := BuildMysql(info)

	got, err := ParseMysql(dsn)
	if err != nil {
		t.Fatalf("ParseMysql(%q) returned error: %v", dsn, err)
	}
	if got != info {
		t.Fatalf("round trip mismatch: built %q, parsed back %+v, want %+v", dsn, got, info)
	}
}

func TestParseVendorDispatch(t *testing.T) {
	if _, err := Parse(VendorSqlite, "/tmp/db.sqlite"); err == nil {
		t.Fatal("expected Parse to error for sqlite - it has no ConnectionInfo breakdown")
	}
	if _, err := Parse("oracle", "whatever"); err == nil {
		t.Fatal("expected Parse to error for an unsupported vendor")
	}

	for _, vendor := range []string{VendorMysql, VendorMariadb} {
		info, err := Parse(vendor, "root:secret@tcp(localhost:3306)/fireback")
		if err != nil {
			t.Fatalf("Parse(%q, ...) returned error: %v", vendor, err)
		}
		if info.Database != "fireback" {
			t.Fatalf("Parse(%q, ...) = %+v, want Database=fireback", vendor, info)
		}
	}
}

func TestBuildVendorDispatch(t *testing.T) {
	if _, err := Build("oracle", ConnectionInfo{}); err == nil {
		t.Fatal("expected Build to error for an unsupported vendor")
	}

	dsn, err := Build(VendorPostgres, ConnectionInfo{Host: "localhost", Port: "5432", Username: "postgres", Database: "postgres"})
	if err != nil {
		t.Fatalf("Build(postgres, ...) returned error: %v", err)
	}
	if dsn != "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable" {
		t.Fatalf("Build(postgres, ...) = %q", dsn)
	}
}
