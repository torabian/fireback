package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"os/exec"
	"strings"
	"testing"
)

// extractLastJSONObject pulls the last top-level {...} object out of s. The CLI prints
// gorm's query trace and other log lines to the same stream ahead of the actual
// emigo.HandleActionInCli JSON payload (see CheckPassportMethodsAction.dyno.go's
// CliHandler), so a plain json.Unmarshal(stdout) would fail on that leading noise -
// this scans for the final balanced-brace object instead of assuming stdout is clean
// JSON on its own.
func extractLastJSONObject(t *testing.T, s string) string {
	t.Helper()

	start := -1
	depth := 0
	end := -1

	for i, r := range s {
		switch r {
		case '{':
			if depth == 0 {
				start = i
			}
			depth++
		case '}':
			if depth == 0 {
				continue
			}
			depth--
			if depth == 0 && start >= 0 {
				end = i + 1
				// Keep scanning: we want the *last* top-level object, in case the
				// action prints more than one JSON blob.
			}
		}
	}

	if start < 0 || end < 0 {
		t.Fatalf("no JSON object found in CLI output:\n%s", s)
	}

	return s[start:end]
}

// TestCheckPassportMethods_CLI runs the compiled fireback binary's
// `check-passport-methods` command (the same action as GET /passports/available-methods,
// wired for CLI instead of HTTP - see CheckPassportMethodsActionCliHandler) and checks
// it prints the same shape of payload the HTTP test expects.
func TestCheckPassportMethods_CLI(t *testing.T) {
	cfg := LoadTestConfig(t)
	bin := cfg.ResolveAppBinary(t)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.CLITimeout)
	defer cancel()

	// check-passport-methods is nested under the "passport" cli group (see
	// PassportCli.go's PassportCli.Commands), not a bare top-level command.
	cmd := exec.CommandContext(ctx, bin, "passport", "check-passport-methods")
	// Must run from the repo root: the binary resolves its .env/config relative to its
	// own working directory, which under `go test` defaults to this test package's
	// directory rather than the repo root - see TestConfig.WorkDir's doc comment.
	cmd.Dir = cfg.WorkDir(t)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		t.Fatalf(
			"`%s passport check-passport-methods` failed: %v\nstdout:\n%s\nstderr:\n%s",
			bin, err, stdout.String(), stderr.String(),
		)
	}

	jsonBlob := extractLastJSONObject(t, stdout.String())

	var item checkPassportMethodsRes
	if err := json.Unmarshal([]byte(jsonBlob), &item); err != nil {
		t.Fatalf("failed to decode CLI JSON output: %v\nextracted blob: %s\nfull stdout:\n%s", err, jsonBlob, stdout.String())
	}

	if !item.Email && !item.Phone && !item.Google && !item.Facebook {
		t.Errorf("expected at least one authentication method to be enabled, got none: %+v", item)
	}
}

// TestCheckPassportMethods_CLI_Help is a cheap smoke test that the command is actually
// registered in the CLI tree at all (catches it silently disappearing from
// PassportCli.go/AbacModule.go's command wiring without needing a live database).
func TestCheckPassportMethods_CLI_Help(t *testing.T) {
	cfg := LoadTestConfig(t)
	bin := cfg.ResolveAppBinary(t)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.CLITimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, bin, "passport", "check-passport-methods", "--help")
	cmd.Dir = cfg.WorkDir(t)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("`%s passport check-passport-methods --help` failed: %v\noutput:\n%s", bin, err, out)
	}

	if !strings.Contains(string(out), "check-passport-methods") {
		t.Errorf("expected --help output to mention the command name, got:\n%s", out)
	}
}
