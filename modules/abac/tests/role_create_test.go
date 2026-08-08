package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
	"testing"
	"time"
)

// roleDtoPayload mirrors abac.RoleDto's POST /role request body shape (see
// modules/abac/RoleDto.go) - only the fields these tests actually send.
type roleDtoPayload struct {
	Name               string   `json:"name"`
	UniqueId           string   `json:"uniqueId,omitempty"`
	CapabilitiesListId []string `json:"capabilitiesListId"`
}

// roleRes mirrors abac.RoleEntity's JSON response shape - only the fields these tests
// actually read.
type roleRes struct {
	UniqueId           string   `json:"uniqueId"`
	Name               string   `json:"name"`
	CapabilitiesListId []string `json:"capabilitiesListId"`
}

// postRole POSTs to /role (abac.RoleCreateAction) and returns the raw response - callers
// decide what status/body they expect, since several of these tests exist specifically to
// verify a rejection.
func postRole(t *testing.T, cfg TestConfig, payload roleDtoPayload) (*http.Response, []byte) {
	t.Helper()

	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/role"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", cfg.WorkspaceID)

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request to %s failed: %v", req.URL, err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	return resp, respBody
}

// deleteRole cleans up a role created by these tests via POST /role/delete
// (abac.RoleAwareDeleteAction) - best-effort, logs rather than fails on error, since
// cleanup problems shouldn't mask the actual test result.
func deleteRole(t *testing.T, cfg TestConfig, uniqueId string) {
	t.Helper()
	if uniqueId == "" {
		return
	}

	body, _ := json.Marshal(map[string][]string{"uniqueIds": {uniqueId}})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/role/delete"), bytes.NewReader(body))
	if err != nil {
		t.Logf("cleanup: failed to build delete request for role %s: %v", uniqueId, err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", cfg.WorkspaceID)

	resp, err := client.Do(req)
	if err != nil {
		t.Logf("cleanup: failed to delete role %s: %v", uniqueId, err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		t.Logf("cleanup: deleting role %s returned %d: %s", uniqueId, resp.StatusCode, b)
	}
}

// TestRoleCreate_HTTP_RejectsReservedName covers RoleCreateAction's guard against ever
// creating a second role literally named (or uniqueId'd) "root", case-insensitively on
// the name - the seeded super-admin role is unique and reserved, see RoleActions.go's
// RoleMessages.RoleNameReserved.
//
// Note: the generated Gin wiring always answers with HTTP 500 regardless of the
// IError's own HttpCode (see RoleCreateActionHandler - `if err != nil { m.JSON(500,
// ...) }`), so these assertions key off the error message content, not the status code.
func TestRoleCreate_HTTP_RejectsReservedName(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	cases := []struct {
		name    string
		payload roleDtoPayload
	}{
		{"lowercase name", roleDtoPayload{Name: "root", CapabilitiesListId: []string{"root.*"}}},
		{"uppercase name", roleDtoPayload{Name: "ROOT", CapabilitiesListId: []string{"root.*"}}},
		{"mixed-case name", roleDtoPayload{Name: "RoOt", CapabilitiesListId: []string{"root.*"}}},
		{"reserved uniqueId", roleDtoPayload{Name: "Totally Fine Name", UniqueId: "root", CapabilitiesListId: []string{"root.*"}}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			resp, body := postRole(t, cfg, tc.payload)

			if resp.StatusCode == http.StatusOK {
				// Shouldn't have been allowed to create - clean up if it somehow was, so
				// a failing test run doesn't leave a role squatting on "root" behind.
				var out googleResponseEnvelope[roleRes]
				if err := json.Unmarshal(body, &out); err == nil {
					deleteRole(t, cfg, out.Data.Item.UniqueId)
				}
				t.Fatalf("expected role creation to be rejected, got 200 OK: %s", body)
			}

			if !strings.Contains(string(body), "RoleNameReserved") {
				t.Errorf("expected the RoleNameReserved error, got %d: %s", resp.StatusCode, body)
			}
		})
	}
}

// TestRoleCreate_HTTP_RequiresOneCapability covers the pre-existing
// RoleNeedsOneCapability guard (see RoleActions.go) - a role that ends up with zero
// capabilities must not be creatable.
func TestRoleCreate_HTTP_RequiresOneCapability(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	resp, body := postRole(t, cfg, roleDtoPayload{
		Name:               "checkendpointtests role with no caps",
		CapabilitiesListId: []string{},
	})

	if resp.StatusCode == http.StatusOK {
		var out googleResponseEnvelope[roleRes]
		if err := json.Unmarshal(body, &out); err == nil {
			deleteRole(t, cfg, out.Data.Item.UniqueId)
		}
		t.Fatalf("expected role creation with no capabilities to be rejected, got 200 OK: %s", body)
	}

	if !strings.Contains(string(body), "RoleNeedsOneCapability") {
		t.Errorf("expected the RoleNeedsOneCapability error, got %d: %s", resp.StatusCode, body)
	}
}

// TestRoleCreate_HTTP_Succeeds is the positive counterpart to the guards above: a role
// with an ordinary name and at least one capability the caller actually has should be
// created normally, and none of the reserved-name/needs-a-capability guards should
// false-positive on it.
func TestRoleCreate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	name := fmt.Sprintf("checkendpointtests role %d", time.Now().UnixNano())
	resp, body := postRole(t, cfg, roleDtoPayload{
		Name:               name,
		CapabilitiesListId: []string{"root.*"},
	})

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d: %s", resp.StatusCode, body)
	}

	var out googleResponseEnvelope[roleRes]
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode JSON response: %v\nbody: %s", err, body)
	}
	defer deleteRole(t, cfg, out.Data.Item.UniqueId)

	if out.Data.Item.UniqueId == "" {
		t.Errorf("expected a generated uniqueId, got none: %+v", out.Data.Item)
	}
	if out.Data.Item.Name != name {
		t.Errorf("expected name %q, got %q", name, out.Data.Item.Name)
	}
	if len(out.Data.Item.CapabilitiesListId) == 0 {
		t.Errorf("expected at least one capability on the created role, got none")
	}
}

// TestRoleCreate_CLI_Help is a cheap smoke test that the role-create command is actually
// registered in the CLI tree at all (catches it silently disappearing from
// AbacModule.go's command wiring without needing a live database/authenticated session).
func TestRoleCreate_CLI_Help(t *testing.T) {
	cfg := LoadTestConfig(t)
	bin := cfg.ResolveAppBinary(t)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.CLITimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, bin, "role-create-action", "--help")
	cmd.Dir = cfg.WorkDir(t)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("`%s role-create-action --help` failed: %v\noutput:\n%s", bin, err, out)
	}

	if !strings.Contains(string(out), "role-create-action") {
		t.Errorf("expected --help output to mention the command name, got:\n%s", out)
	}
}
