package tests

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os/exec"
	"strings"
	"testing"
)

// capabilityInfoDto mirrors abac.CapabilityInfoDto (see modules/abac/CapabilityInfoDto.go)
// - duplicated here rather than imported, since these are black-box tests against the
// HTTP surface, not the Go package itself.
type capabilityInfoDto struct {
	UniqueId string              `json:"uniqueId"`
	Name     string              `json:"name"`
	Children []capabilityInfoDto `json:"children"`
}

// capabilitiesTreeRes mirrors abac.CapabilitiesTreeActionRes (see
// modules/abac/CapabilitiesTreeAction.go).
type capabilitiesTreeRes struct {
	Capabilities []capabilityInfoDto `json:"capabilities"`
	Nested       []capabilityInfoDto `json:"nested"`
}

// TestCapabilitiesTree_HTTP hits GET /capabilitiesTree (abac's CapabilitiesTreeAction) as
// an authenticated caller and checks it actually returns the capability catalog instead
// of two empty arrays.
//
// Regression guard: CapabilitiesTreeAction used to hardcode query.WorkspaceId/UserId to
// "system" right after ResolveActionContext had already resolved the caller's real
// workspace/user - GetWorkspaceAndUserAccesses then looked up
// UserAccessPerWorkspace["system"] (almost never a key the caller actually has data
// under), got back empty access lists, and MeetsCheck rejected every capability against
// an empty allowed-list, so both "capabilities" and "nested" always came back [] no
// matter how much was actually in the database. Fixed by leaving
// query.WorkspaceId/UserId as whatever ResolveActionContext resolved from the request.
func TestCapabilitiesTree_HTTP(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/capabilitiesTree"), nil)
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", cfg.WorkspaceID)

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request to %s failed: %v", req.URL, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d %s\nbody: %s", resp.StatusCode, resp.Status, body)
	}

	var out googleResponseEnvelope[capabilitiesTreeRes]
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode JSON response: %v\nbody: %s", err, body)
	}

	item := out.Data.Item

	// Every fireback app registers at least its own module wildcard/CRUD permissions
	// (see SyncPermissionsInDatabase, run on every migration apply and app startup), so
	// an authenticated caller should never see an empty catalog - this is the exact bug
	// this test guards against.
	if len(item.Capabilities) == 0 {
		t.Errorf("expected at least one capability, got an empty \"capabilities\" array")
	}
	if len(item.Nested) == 0 {
		t.Errorf("expected at least one nested capability node, got an empty \"nested\" array")
	}

	// Spot-check: this permission is always seeded (it's one of Capability's own CRUD
	// permissions, see abac.PERM_ROOT_CAPABILITY_QUERY), so it should always be present
	// for an authorized caller.
	found := false
	for _, c := range item.Capabilities {
		if c.UniqueId == "root.manage.fireback.capability.query" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected %q among the returned capabilities, got %d capabilities: %+v",
			"root.manage.fireback.capability.query", len(item.Capabilities), item.Capabilities)
	}
}

// TestCapabilitiesTree_CLI_Help is a cheap smoke test that the "treex" command is
// actually registered in the CLI tree at all (catches it silently disappearing from
// AbacModule.go's command wiring without needing a live database/authenticated
// session).
func TestCapabilitiesTree_CLI_Help(t *testing.T) {
	cfg := LoadTestConfig(t)
	bin := cfg.ResolveAppBinary(t)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.CLITimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, bin, "treex", "--help")
	cmd.Dir = cfg.WorkDir(t)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("`%s treex --help` failed: %v\noutput:\n%s", bin, err, out)
	}

	if !strings.Contains(string(out), "treex") {
		t.Errorf("expected --help output to mention the command name, got:\n%s", out)
	}
}
