//go:build wasm

// Package wasmpgdriver is a database/sql driver that speaks to a Postgres-ish
// backend through a single JS-side function instead of a socket:
//
//	window.queryDatabase(sql string, args []any) => Promise<string /* JSON WasmQueryResult */>
//
// This is the exact bridge shape used by the emi in-browser-server example
// (browser/database-bridge.js, wired to pglite). By implementing the
// database/sql/driver interfaces here, gorm never has to know it's talking
// to something other than a real Postgres connection: gorm.io/driver/postgres
// is handed a normal *sql.DB (see application.ConnectWasmPostgres), and every
// gorm feature (AutoMigrate, associations, hooks, transactions, ...) keeps
// working unmodified. Only the transport underneath sql.DB changes.
package wasmpgdriver

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"sync"
	"syscall/js"
	"time"
)

// Well-known Postgres OIDs we need to special-case when turning the JSON
// values that come back over the bridge into driver.Value. Everything else
// is left as the JSON-decoded bool/float64/string/nil, which database/sql's
// convertAssign already knows how to fit into common Go destination kinds.
const (
	oidDate        = 1082
	oidTime        = 1083
	oidTimestamp   = 1114
	oidTimestampTz = 1184
	oidBytea       = 17
)

type wasmField struct {
	Name       string `json:"name"`
	DataTypeID int    `json:"dataTypeID"`
}

type wasmError struct {
	Message  string `json:"Message"`
	Code     string `json:"Code"`
	Severity string `json:"Severity"`
}

func (e *wasmError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("%s (%s): %s", e.Severity, e.Code, e.Message)
	}
	return e.Message
}

type wasmResult struct {
	Rows         []map[string]any `json:"rows"`
	Fields       []wasmField      `json:"fields"`
	AffectedRows int64            `json:"affectedRows"`
	Error        *wasmError       `json:"error"`
}

// Open wraps queryFunc (normally js.Global().Get("queryDatabase")) into a
// *sql.DB. The bridge fronts exactly one logical pglite/postgres session, so
// the pool is pinned to a single connection the same way fireback already
// pins sqlite to one connection in DirectConnectToDb.
func Open(queryFunc js.Value) *sql.DB {
	db := sql.OpenDB(NewConnector(queryFunc))
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	return db
}

// Connector implements database/sql/driver.Connector so callers can go
// straight from a live js.Value to a *sql.DB via sql.OpenDB, without a
// string DSN or a process-wide sql.Register call (queryFunc is only known
// at runtime, once database-bridge.js has run).
type Connector struct {
	queryFunc js.Value
	mu        *sync.Mutex
}

func NewConnector(queryFunc js.Value) *Connector {
	return &Connector{queryFunc: queryFunc, mu: &sync.Mutex{}}
}

func (c *Connector) Connect(ctx context.Context) (driver.Conn, error) {
	return &wasmConn{queryFunc: c.queryFunc, mu: c.mu}, nil
}

// Driver satisfies driver.Connector. database/sql never calls Open on it
// when the *sql.DB was created via sql.OpenDB(connector), but the interface
// requires a value to exist.
func (c *Connector) Driver() driver.Driver { return &wasmDriver{connector: c} }

type wasmDriver struct{ connector *Connector }

func (d *wasmDriver) Open(name string) (driver.Conn, error) {
	return d.connector.Connect(context.Background())
}

// wasmConn is one logical connection. mu is shared with sibling connections
// from the same Connector (there should only ever be one, see Open above)
// so BEGIN/COMMIT/ROLLBACK sequences can't interleave with unrelated
// statements if the pool size guard is ever relaxed.
type wasmConn struct {
	queryFunc js.Value
	mu        *sync.Mutex
	closed    bool
}

var (
	_ driver.Conn              = (*wasmConn)(nil)
	_ driver.ExecerContext     = (*wasmConn)(nil)
	_ driver.QueryerContext    = (*wasmConn)(nil)
	_ driver.ConnBeginTx       = (*wasmConn)(nil)
	_ driver.Pinger            = (*wasmConn)(nil)
	_ driver.NamedValueChecker = (*wasmConn)(nil)
)

func (c *wasmConn) Prepare(query string) (driver.Stmt, error) {
	return &wasmStmt{c: c, query: query}, nil
}

func (c *wasmConn) Close() error {
	c.closed = true
	return nil
}

func (c *wasmConn) Begin() (driver.Tx, error) {
	return c.BeginTx(context.Background(), driver.TxOptions{})
}

func (c *wasmConn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	if _, err := c.exec(ctx, "BEGIN", nil); err != nil {
		return nil, err
	}
	return &wasmTx{c: c}, nil
}

func (c *wasmConn) Ping(ctx context.Context) error {
	_, err := c.exec(ctx, "SELECT 1", nil)
	return err
}

// CheckNamedValue accepts every value as-is. By the time an argument reaches
// here, database/sql has already reduced it (via driver.Valuer / the default
// parameter converter) to one of int64, float64, bool, []byte, string,
// time.Time or nil — the set toJSValue below knows how to hand to
// syscall/js.
func (c *wasmConn) CheckNamedValue(nv *driver.NamedValue) error {
	return nil
}

func (c *wasmConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	res, err := c.exec(ctx, query, args)
	if err != nil {
		return nil, err
	}
	return driver.RowsAffected(res.AffectedRows), nil
}

func (c *wasmConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	res, err := c.exec(ctx, query, args)
	if err != nil {
		return nil, err
	}
	return &wasmRows{fields: res.Fields, data: res.Rows}, nil
}

func (c *wasmConn) exec(ctx context.Context, query string, args []driver.NamedValue) (*wasmResult, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return nil, driver.ErrBadConn
	}

	jsArgs := make([]any, len(args))
	for i, a := range args {
		jsArgs[i] = toJSValue(a.Value)
	}

	promise := c.queryFunc.Invoke(query, js.ValueOf(jsArgs))

	payload, err := awaitPromise(ctx, promise)
	if err != nil {
		return nil, err
	}

	var result wasmResult
	if err := json.Unmarshal(payload, &result); err != nil {
		return nil, fmt.Errorf("wasmpgdriver: decoding bridge response: %w", err)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &result, nil
}

// toJSValue narrows the handful of types database/sql may hand us into
// something syscall/js.ValueOf accepts. time.Time is the only one of the
// driver.Value primitives js.ValueOf itself panics on.
func toJSValue(v driver.Value) any {
	if t, ok := v.(time.Time); ok {
		return t.UTC().Format(time.RFC3339Nano)
	}
	return v
}

func awaitPromise(ctx context.Context, promise js.Value) ([]byte, error) {
	resultCh := make(chan []byte, 1)
	errCh := make(chan error, 1)

	then := js.FuncOf(func(this js.Value, jsArgs []js.Value) any {
		resultCh <- []byte(jsArgs[0].String())
		return nil
	})
	defer then.Release()

	catch := js.FuncOf(func(this js.Value, jsArgs []js.Value) any {
		msg := "unknown JS error"
		if len(jsArgs) > 0 {
			if m := jsArgs[0].Get("message"); !m.IsUndefined() {
				msg = m.String()
			} else {
				msg = jsArgs[0].String()
			}
		}
		errCh <- fmt.Errorf("wasmpgdriver: bridge rejected: %s", msg)
		return nil
	})
	defer catch.Release()

	promise.Call("then", then).Call("catch", catch)

	select {
	case res := <-resultCh:
		return res, nil
	case err := <-errCh:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// wasmStmt is a thin adapter: Query/ExecerContext on wasmConn already do the
// real work, Prepare exists only because driver.Conn requires it. NumInput
// of -1 tells database/sql to skip arg-count validation, since the bridge
// (not this package) is what ultimately rejects a bad parameter count.
type wasmStmt struct {
	c     *wasmConn
	query string
}

func (s *wasmStmt) Close() error  { return nil }
func (s *wasmStmt) NumInput() int { return -1 }

func (s *wasmStmt) Exec(args []driver.Value) (driver.Result, error) {
	return s.c.ExecContext(context.Background(), s.query, valuesToNamed(args))
}

func (s *wasmStmt) Query(args []driver.Value) (driver.Rows, error) {
	return s.c.QueryContext(context.Background(), s.query, valuesToNamed(args))
}

func (s *wasmStmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	return s.c.ExecContext(ctx, s.query, args)
}

func (s *wasmStmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	return s.c.QueryContext(ctx, s.query, args)
}

func valuesToNamed(args []driver.Value) []driver.NamedValue {
	nv := make([]driver.NamedValue, len(args))
	for i, v := range args {
		nv[i] = driver.NamedValue{Ordinal: i + 1, Value: v}
	}
	return nv
}

type wasmTx struct{ c *wasmConn }

func (t *wasmTx) Commit() error {
	_, err := t.c.exec(context.Background(), "COMMIT", nil)
	return err
}

func (t *wasmTx) Rollback() error {
	_, err := t.c.exec(context.Background(), "ROLLBACK", nil)
	return err
}

type wasmRows struct {
	fields []wasmField
	data   []map[string]any
	idx    int
}

func (r *wasmRows) Columns() []string {
	cols := make([]string, len(r.fields))
	for i, f := range r.fields {
		cols[i] = f.Name
	}
	return cols
}

func (r *wasmRows) Close() error { return nil }

func (r *wasmRows) Next(dest []driver.Value) error {
	if r.idx >= len(r.data) {
		return io.EOF
	}
	row := r.data[r.idx]
	r.idx++

	for i, f := range r.fields {
		if i >= len(dest) {
			break
		}
		dest[i] = decodeValue(row[f.Name], f.DataTypeID)
	}
	return nil
}

// decodeValue turns one JSON-decoded field (bool/float64/string/nil/map/
// slice, per encoding/json's default unmarshaling into `any`) into a
// database/sql-legal driver.Value. Dates/times need OID-directed parsing —
// the bridge hands them back as strings (JSON has no native timestamp type),
// but gorm model fields are typically time.Time and database/sql's
// convertAssign won't parse a string into one on its own, only pass through
// an already-typed time.Time. bytea similarly needs unescaping from
// Postgres's "\x..." hex text form into []byte. Everything else is left as
// the primitive JSON gave us, or re-encoded to a JSON string for composite
// types (json/jsonb/arrays), same as any driver returning text-mode values.
func decodeValue(v any, oid int) driver.Value {
	switch t := v.(type) {
	case nil, bool, float64:
		return t
	case string:
		switch oid {
		case oidTimestamp, oidTimestampTz:
			if parsed, err := time.Parse(time.RFC3339Nano, t); err == nil {
				return parsed
			}
		case oidDate:
			if parsed, err := time.Parse("2006-01-02", t); err == nil {
				return parsed
			}
		case oidTime:
			return t
		case oidBytea:
			if b, ok := decodeHexBytea(t); ok {
				return b
			}
		}
		return t
	default:
		b, err := json.Marshal(t)
		if err != nil {
			return nil
		}
		return string(b)
	}
}

func decodeHexBytea(s string) ([]byte, bool) {
	if !strings.HasPrefix(s, "\\x") {
		return nil, false
	}
	b, err := hex.DecodeString(s[2:])
	if err != nil {
		return nil, false
	}
	return b, true
}
