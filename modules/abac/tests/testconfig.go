// Package tests holds black-box endpoint/CLI tests for the abac module: real HTTP
// requests against a running `./app start` server, and real subprocess invocations of
// the compiled CLI binary - not in-process unit tests. That means every test here needs
// a live server and/or a live database to actually exercise anything, which is exactly
// what these tests are for (there were previously none at all for abac), but it also
// means they must be configurable rather than hardcoded to one developer's machine -
// see TestConfig below. Every test skips (rather than fails) when its target isn't
// reachable/buildable, so `go test ./...` stays green in environments with no server or
// binary available (CI without a Postgres instance, for example) while still doing real
// verification wherever one is.
package tests

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// TestConfig carries every knob these tests read from the environment. All of it has a
// default matching this repo's own checked-in .env (see repo root), so `go test
// ./modules/abac/tests/...` works out of the box against a `./app start` you already
// have running locally the same way the manual curl/CLI checks in this repo's history
// did - override any of it to point at a different host, a CI-built binary, or an
// authenticated account for tests that need one.
type TestConfig struct {
	// BaseURL is the abac/fireback HTTP server's address, no trailing slash.
	// Override: ABAC_TEST_BASE_URL.
	BaseURL string

	// AppBinary is the path to the compiled fireback CLI binary (`go build -o app
	// ./cmd/fireback`). May be relative (resolved against the repo root) or absolute.
	// Override: ABAC_TEST_APP_BINARY.
	AppBinary string

	// CliToken/WorkspaceID are an authenticated session's token and the workspace it
	// was issued for (see appConfig.CliToken/CliWorkspace, e.g. after `./app auth`) -
	// only needed by tests against endpoints/commands that require authorization.
	// Unset by default: tests requiring it must skip cleanly when it's empty, not fail.
	// Override: ABAC_TEST_CLI_TOKEN, ABAC_TEST_WORKSPACE_ID.
	CliToken    string
	WorkspaceID string

	// HTTPTimeout bounds every HTTP request these tests make.
	// Override: ABAC_TEST_HTTP_TIMEOUT (Go duration string, e.g. "5s").
	HTTPTimeout time.Duration

	// CLITimeout bounds every CLI subprocess invocation.
	// Override: ABAC_TEST_CLI_TIMEOUT (Go duration string, e.g. "20s").
	CLITimeout time.Duration
}

func envOr(key, def string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return def
}

// LoadTestConfig reads TestConfig from the environment, applying repo-relative
// defaults. It never fails - individual tests decide what to do when a required piece
// (a token, a reachable server, a built binary) is actually missing, via t.Skip.
func LoadTestConfig(t *testing.T) TestConfig {
	t.Helper()

	cfg := TestConfig{
		BaseURL:     strings.TrimRight(envOr("ABAC_TEST_BASE_URL", "http://localhost:4500"), "/"),
		AppBinary:   envOr("ABAC_TEST_APP_BINARY", "app"),
		CliToken:    os.Getenv("ABAC_TEST_CLI_TOKEN"),
		WorkspaceID: envOr("ABAC_TEST_WORKSPACE_ID", "root"),
	}

	cfg.HTTPTimeout = parseDurationOr(t, "ABAC_TEST_HTTP_TIMEOUT", 5*time.Second)
	cfg.CLITimeout = parseDurationOr(t, "ABAC_TEST_CLI_TIMEOUT", 20*time.Second)

	return cfg
}

func parseDurationOr(t *testing.T, key string, def time.Duration) time.Duration {
	t.Helper()
	raw := os.Getenv(key)
	if raw == "" {
		return def
	}
	d, err := time.ParseDuration(raw)
	if err != nil {
		t.Fatalf("%s=%q is not a valid Go duration (e.g. \"5s\"): %v", key, raw, err)
	}
	return d
}

// repoRoot walks up from the current working directory (package directory, under `go
// test`) until it finds go.mod, so AppBinary's relative default ("app", the repo root's
// build output) resolves regardless of which directory `go test` happens to be invoked
// from.
func repoRoot(t *testing.T) string {
	t.Helper()

	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("cannot determine working directory: %v", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatalf("could not locate repo root (go.mod) walking up from %s", dir)
		}
		dir = parent
	}
}

// ResolveAppBinary returns AppBinary as an absolute path (relative paths are resolved
// against the repo root, not the test package's own directory) and skips the test if
// nothing exists there - with a message telling the developer/CI exactly how to build
// it, rather than failing with a bare "no such file" from exec.Command.
func (c TestConfig) ResolveAppBinary(t *testing.T) string {
	t.Helper()

	bin := c.AppBinary
	if !filepath.IsAbs(bin) {
		bin = filepath.Join(repoRoot(t), bin)
	}

	if info, err := os.Stat(bin); err != nil || info.IsDir() {
		t.Skipf(
			"fireback CLI binary not found at %s - build it first: GOEXPERIMENT=jsonv2 go build -o app ./cmd/fireback (or set ABAC_TEST_APP_BINARY to an existing binary)",
			bin,
		)
	}

	return bin
}

// WorkDir returns the repo root, which callers must set as the CLI subprocess's
// cmd.Dir. The binary resolves its .env/config relative to its working directory (see
// modules/fireback's config loading), not relative to itself - under `go test`, the
// process's own cwd is the test package's directory (modules/abac/tests/), not the repo
// root, so without this every CLI invocation would silently pick up an empty/default
// config instead of the repo's real .env.
func (c TestConfig) WorkDir(t *testing.T) string {
	t.Helper()
	return repoRoot(t)
}

// RequireServer skips the test (with an actionable message) unless BaseURL is actually
// accepting connections - these are black-box tests, not in-process ones, so there's no
// way to run them at all without a live `./app start`.
func (c TestConfig) RequireServer(t *testing.T) {
	t.Helper()

	u := c.BaseURL
	host := strings.TrimPrefix(strings.TrimPrefix(u, "https://"), "http://")
	if idx := strings.Index(host, "/"); idx >= 0 {
		host = host[:idx]
	}
	if !strings.Contains(host, ":") {
		if strings.HasPrefix(u, "https://") {
			host += ":443"
		} else {
			host += ":80"
		}
	}

	conn, err := net.DialTimeout("tcp", host, 2*time.Second)
	if err != nil {
		t.Skipf(
			"abac/fireback server not reachable at %s (%v) - start it first: ./app start (or set ABAC_TEST_BASE_URL to point at a running instance)",
			c.BaseURL, err,
		)
		return
	}
	conn.Close()
}

// RequireAuth skips the test unless a CLI token has been configured - it's opt-in
// because most of the repo's existing dev setups don't have one lying around in the
// environment, and a test that needs to authenticate has no sensible fallback.
func (c TestConfig) RequireAuth(t *testing.T) {
	t.Helper()
	if c.CliToken == "" {
		t.Skip("no authenticated session configured - set ABAC_TEST_CLI_TOKEN (and ABAC_TEST_WORKSPACE_ID if not \"root\") to run this test; see ./app auth")
	}
}

// NewHTTPClient returns an *http.Client bounded by cfg.HTTPTimeout.
func (c TestConfig) NewHTTPClient() *http.Client {
	return &http.Client{Timeout: c.HTTPTimeout}
}

// URL joins BaseURL with a path that must start with "/".
func (c TestConfig) URL(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return fmt.Sprintf("%s%s", c.BaseURL, path)
}
