// messagingConfig black-box CLI tests - real subprocess invocations of the compiled
// fireback binary (`messaging config get`/`messaging config update`), not in-process
// calls into modules/abac/messaging.MessagingConfigGetAction/MessagingConfigUpdateAction
// directly - see testconfig.go's doc comment and passport_methods_cli_test.go for the
// same convention.
//
// MessagingConfig (see Messaging.emi.yml's features.actions:false) is a single, global
// row shared by the whole running server - there is no per-workspace scoping and no
// generic Browse/Create/AwareDelete exposed for it, only the two hand-declared
// "distinct" actions this file exercises. Every test here captures whatever value is
// already on the row before touching it and restores it afterwards (same discipline
// workspace_config_test.go's TestWorkspaceConfigDistinct* tests use), so these are safe
// to run in any order, repeatedly, without leaving global side effects for the rest of
// the suite - and, since there's no delete action for this row, restoring is also the
// only way to get back to a "nothing set" state after the very first write ever made to
// it.
package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"testing"
	"time"
)

// randomSuffix keeps checkendpointtests-* CLI flag values unique across test runs,
// following this package's own time.Now().UnixNano() convention (see e.g.
// capability_test.go).
func randomSuffix() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// messagingConfigRes mirrors messagingdefs.MessagingConfigDto's JSON shape - duplicated
// here rather than imported, since these are black-box tests against the CLI surface,
// not the Go package itself.
type messagingConfigRes struct {
	UniqueId                   string  `json:"uniqueId"`
	GeneralEmailProviderId     *string `json:"generalEmailProviderId"`
	GeneralGsmProviderId       *string `json:"generalGsmProviderId"`
	InviteToWorkspaceContentId *string `json:"inviteToWorkspaceContentId"`
	EmailOtpContentId          *string `json:"emailOtpContentId"`
	SmsOtpContentId            *string `json:"smsOtpContentId"`
}

// runMessagingConfigCLI runs `<bin> messaging config <args...>` and returns its raw
// stdout, failing the test (with both stdout and stderr attached) if the process itself
// errors - a non-zero exit here always means the command genuinely failed, never a
// legitimate "empty" result (see MessagingConfigGetAction, which returns 200 with an
// empty dto rather than any kind of error when there's no row yet).
func runMessagingConfigCLI(t *testing.T, cfg TestConfig, args ...string) string {
	t.Helper()
	bin := cfg.ResolveAppBinary(t)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.CLITimeout)
	defer cancel()

	fullArgs := append([]string{"messaging", "config"}, args...)
	cmd := exec.CommandContext(ctx, bin, fullArgs...)
	// Must run from the repo root: the binary resolves its .env/config relative to its
	// own working directory - see TestConfig.WorkDir's doc comment.
	cmd.Dir = cfg.WorkDir(t)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		t.Fatalf(
			"`%s messaging config %s` failed: %v\nstdout:\n%s\nstderr:\n%s",
			bin, strings.Join(args, " "), err, stdout.String(), stderr.String(),
		)
	}

	return stdout.String()
}

func messagingConfigDecode(t *testing.T, raw string) messagingConfigRes {
	t.Helper()
	blob := extractLastJSONObject(t, raw)
	var env googleResponseEnvelope[messagingConfigRes]
	if err := json.Unmarshal([]byte(blob), &env); err != nil {
		t.Fatalf("failed to decode messagingConfig CLI JSON output: %v\nextracted blob: %s\nfull output:\n%s", err, blob, raw)
	}
	return env.Data.Item
}

func messagingConfigGet(t *testing.T, cfg TestConfig) messagingConfigRes {
	t.Helper()
	return messagingConfigDecode(t, runMessagingConfigCLI(t, cfg, "get"))
}

func messagingConfigUpdate(t *testing.T, cfg TestConfig, flags ...string) messagingConfigRes {
	t.Helper()
	args := append([]string{"update"}, flags...)
	return messagingConfigDecode(t, runMessagingConfigCLI(t, cfg, args...))
}

// nullableFlagValue renders a *string back into the CLI flag value that reproduces it:
// the literal "null" sentinel CastMessagingConfigOptionalDtoFromCli/emigo.ParseNullable
// recognizes for "set this Nullable field to null" when the field is unset, or the
// value itself otherwise. An empty string is deliberately never produced here -
// ParseNullable treats "" as "flag not meaningfully set" (a no-op, unlike "null"), so a
// genuinely empty captured value would silently fail to round-trip.
func nullableFlagValue(v *string) string {
	if v == nil || *v == "" {
		return "null"
	}
	return *v
}

// restoreMessagingConfig writes every field of original back onto the live row,
// including nulling out whatever was unset - see nullableFlagValue.
func restoreMessagingConfig(t *testing.T, cfg TestConfig, original messagingConfigRes) {
	t.Helper()
	messagingConfigUpdate(t, cfg,
		"--general-email-provider-id", nullableFlagValue(original.GeneralEmailProviderId),
		"--general-gsm-provider-id", nullableFlagValue(original.GeneralGsmProviderId),
		"--invite-to-workspace-content-id", nullableFlagValue(original.InviteToWorkspaceContentId),
		"--email-otp-content-id", nullableFlagValue(original.EmailOtpContentId),
		"--sms-otp-content-id", nullableFlagValue(original.SmsOtpContentId),
	)
}

func strPtr(s string) *string { return &s }

// TestMessagingConfigGet_CLI_Succeeds covers the plain "does this even work" case -
// `messaging config get` must always come back as a well-formed, successful response,
// whether or not anything has ever been configured (see the package doc comment above:
// there's no delete action to force a genuinely record-less state after the first
// write, but MessagingConfigGetAction's gorm.ErrRecordNotFound branch and its normal
// branch return the identical dto shape either way - TestMessagingConfigUpdate_CLI_
// CanNullFieldsAgain below is what actually proves the "empty dto, not an error" claim
// for every field).
func TestMessagingConfigGet_CLI_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)

	item := messagingConfigGet(t, cfg)
	_ = item // reaching here at all (no t.Fatalf from runMessagingConfigCLI/decode) is the assertion
}

// TestMessagingConfigGet_CLI_Help is a cheap smoke test that both commands are actually
// registered in the CLI tree (catches them silently disappearing from
// MessagingModule.go's command wiring without needing a live database).
func TestMessagingConfigGet_CLI_Help(t *testing.T) {
	cfg := LoadTestConfig(t)
	bin := cfg.ResolveAppBinary(t)

	for _, sub := range [][]string{{"get", "--help"}, {"update", "--help"}} {
		ctx, cancel := context.WithTimeout(context.Background(), cfg.CLITimeout)
		args := append([]string{"messaging", "config"}, sub...)
		cmd := exec.CommandContext(ctx, bin, args...)
		cmd.Dir = cfg.WorkDir(t)
		out, err := cmd.CombinedOutput()
		cancel()
		if err != nil {
			t.Fatalf("`%s %s` failed: %v\noutput:\n%s", bin, strings.Join(args, " "), err, out)
		}
		if !strings.Contains(string(out), sub[0]) {
			t.Errorf("expected --help output for %q to mention the command name, got:\n%s", sub[0], out)
		}
	}
}

// TestMessagingConfigUpdate_CLI_SingleField updates exactly one field and verifies it's
// the only one that changed - catches an Update implementation that clobbers untouched
// fields (e.g. building a full-column overwrite instead of diffing on IsSet()).
func TestMessagingConfigUpdate_CLI_SingleField(t *testing.T) {
	cfg := LoadTestConfig(t)

	original := messagingConfigGet(t, cfg)
	defer restoreMessagingConfig(t, cfg, original)

	newValue := "checkendpointtests-single-" + randomSuffix()
	updated := messagingConfigUpdate(t, cfg, "--general-email-provider-id", newValue)

	if updated.GeneralEmailProviderId == nil || *updated.GeneralEmailProviderId != newValue {
		t.Fatalf("expected generalEmailProviderId %q after single-field update, got %+v", newValue, updated.GeneralEmailProviderId)
	}
	if !equalStrPtr(updated.GeneralGsmProviderId, original.GeneralGsmProviderId) {
		t.Errorf("expected generalGsmProviderId to stay untouched (%v), got %v", strDeref(original.GeneralGsmProviderId), strDeref(updated.GeneralGsmProviderId))
	}
	if !equalStrPtr(updated.InviteToWorkspaceContentId, original.InviteToWorkspaceContentId) {
		t.Errorf("expected inviteToWorkspaceContentId to stay untouched (%v), got %v", strDeref(original.InviteToWorkspaceContentId), strDeref(updated.InviteToWorkspaceContentId))
	}
	if !equalStrPtr(updated.EmailOtpContentId, original.EmailOtpContentId) {
		t.Errorf("expected emailOtpContentId to stay untouched (%v), got %v", strDeref(original.EmailOtpContentId), strDeref(updated.EmailOtpContentId))
	}
	if !equalStrPtr(updated.SmsOtpContentId, original.SmsOtpContentId) {
		t.Errorf("expected smsOtpContentId to stay untouched (%v), got %v", strDeref(original.SmsOtpContentId), strDeref(updated.SmsOtpContentId))
	}

	// Re-fetch independently of the update response, so this also proves the write was
	// actually persisted (not just echoed back by the update call itself).
	reGet := messagingConfigGet(t, cfg)
	if reGet.GeneralEmailProviderId == nil || *reGet.GeneralEmailProviderId != newValue {
		t.Errorf("expected a follow-up get to reflect the update, got %+v", reGet.GeneralEmailProviderId)
	}
}

// TestMessagingConfigUpdate_CLI_AllFieldsAtOnce sets every field in a single update
// call and verifies every one of them actually persisted - the "check all fields are
// being saved" requirement.
func TestMessagingConfigUpdate_CLI_AllFieldsAtOnce(t *testing.T) {
	cfg := LoadTestConfig(t)

	original := messagingConfigGet(t, cfg)
	defer restoreMessagingConfig(t, cfg, original)

	suffix := randomSuffix()
	want := messagingConfigRes{
		GeneralEmailProviderId:     strPtr("checkendpointtests-provider-email-" + suffix),
		GeneralGsmProviderId:       strPtr("checkendpointtests-provider-gsm-" + suffix),
		InviteToWorkspaceContentId: strPtr("checkendpointtests-content-invite-" + suffix),
		EmailOtpContentId:          strPtr("checkendpointtests-content-email-otp-" + suffix),
		SmsOtpContentId:            strPtr("checkendpointtests-content-sms-otp-" + suffix),
	}

	updated := messagingConfigUpdate(t, cfg,
		"--general-email-provider-id", *want.GeneralEmailProviderId,
		"--general-gsm-provider-id", *want.GeneralGsmProviderId,
		"--invite-to-workspace-content-id", *want.InviteToWorkspaceContentId,
		"--email-otp-content-id", *want.EmailOtpContentId,
		"--sms-otp-content-id", *want.SmsOtpContentId,
	)

	assertMessagingConfigFields(t, "update response", updated, want)

	// Same independent re-fetch check as the single-field test above.
	reGet := messagingConfigGet(t, cfg)
	assertMessagingConfigFields(t, "follow-up get", reGet, want)
}

// TestMessagingConfigUpdate_CLI_CanNullFieldsAgain proves fields can be cleared back to
// null (not just set to a new non-empty value), and that a fully-nulled row is exactly
// the "empty dto, no error" shape MessagingConfigGetAction promises for a brand new
// install with no row at all yet (see the package doc comment and
// TestMessagingConfigGet_CLI_Succeeds above).
func TestMessagingConfigUpdate_CLI_CanNullFieldsAgain(t *testing.T) {
	cfg := LoadTestConfig(t)

	original := messagingConfigGet(t, cfg)
	defer restoreMessagingConfig(t, cfg, original)

	suffix := randomSuffix()
	// First set every field to something non-null, so nulling them back out afterwards
	// is a real change to prove, not a no-op against fields that were already null.
	messagingConfigUpdate(t, cfg,
		"--general-email-provider-id", "checkendpointtests-nulltest-email-"+suffix,
		"--general-gsm-provider-id", "checkendpointtests-nulltest-gsm-"+suffix,
		"--invite-to-workspace-content-id", "checkendpointtests-nulltest-invite-"+suffix,
		"--email-otp-content-id", "checkendpointtests-nulltest-emailotp-"+suffix,
		"--sms-otp-content-id", "checkendpointtests-nulltest-smsotp-"+suffix,
	)

	nulled := messagingConfigUpdate(t, cfg,
		"--general-email-provider-id", "null",
		"--general-gsm-provider-id", "null",
		"--invite-to-workspace-content-id", "null",
		"--email-otp-content-id", "null",
		"--sms-otp-content-id", "null",
	)

	assertMessagingConfigFields(t, "null update response", nulled, messagingConfigRes{})

	reGet := messagingConfigGet(t, cfg)
	assertMessagingConfigFields(t, "follow-up get after nulling", reGet, messagingConfigRes{})
}

func assertMessagingConfigFields(t *testing.T, label string, got, want messagingConfigRes) {
	t.Helper()
	if !equalStrPtr(got.GeneralEmailProviderId, want.GeneralEmailProviderId) {
		t.Errorf("%s: generalEmailProviderId: expected %v, got %v", label, strDeref(want.GeneralEmailProviderId), strDeref(got.GeneralEmailProviderId))
	}
	if !equalStrPtr(got.GeneralGsmProviderId, want.GeneralGsmProviderId) {
		t.Errorf("%s: generalGsmProviderId: expected %v, got %v", label, strDeref(want.GeneralGsmProviderId), strDeref(got.GeneralGsmProviderId))
	}
	if !equalStrPtr(got.InviteToWorkspaceContentId, want.InviteToWorkspaceContentId) {
		t.Errorf("%s: inviteToWorkspaceContentId: expected %v, got %v", label, strDeref(want.InviteToWorkspaceContentId), strDeref(got.InviteToWorkspaceContentId))
	}
	if !equalStrPtr(got.EmailOtpContentId, want.EmailOtpContentId) {
		t.Errorf("%s: emailOtpContentId: expected %v, got %v", label, strDeref(want.EmailOtpContentId), strDeref(got.EmailOtpContentId))
	}
	if !equalStrPtr(got.SmsOtpContentId, want.SmsOtpContentId) {
		t.Errorf("%s: smsOtpContentId: expected %v, got %v", label, strDeref(want.SmsOtpContentId), strDeref(got.SmsOtpContentId))
	}
}

func equalStrPtr(a, b *string) bool {
	if a == nil || b == nil {
		return a == b
	}
	return *a == *b
}

func strDeref(s *string) string {
	if s == nil {
		return "<nil>"
	}
	return *s
}
