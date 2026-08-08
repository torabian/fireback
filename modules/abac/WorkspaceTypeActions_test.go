package abac

import (
	"strings"
	"testing"
)

// TestRejectRootWorkspaceTypeDeletion covers the guard added to
// WorkspaceTypeAwareDeleteAction/WorkspaceTypeAwareDeletePreviewAction that stops the
// seeded "root" workspace type (see UserCli.go) from ever being deleted, while leaving
// ordinary workspace types (and batched deletes that merely happen to not include
// "root") untouched.
func TestRejectRootWorkspaceTypeDeletion(t *testing.T) {
	t.Run("rejects root alone", func(t *testing.T) {
		err := RejectRootWorkspaceTypeDeletion([]string{"root"})
		if err == nil {
			t.Fatal("expected an error when deleting the root workspace type, got nil")
		}
		if err.Message["$"] != "CannotDeleteRootWorkspaceType" {
			t.Errorf("expected CannotDeleteRootWorkspaceType, got %v", err.Message)
		}
	})

	t.Run("rejects root mixed in with other ids", func(t *testing.T) {
		err := RejectRootWorkspaceTypeDeletion([]string{"cms", "root", "customer"})
		if err == nil {
			t.Fatal("expected an error when the batch includes the root workspace type, got nil")
		}
	})

	t.Run("allows ordinary workspace types", func(t *testing.T) {
		if err := RejectRootWorkspaceTypeDeletion([]string{"cms", "customer"}); err != nil {
			t.Errorf("expected non-root workspace types to be deletable, got error: %v", err.Message)
		}
	})

	t.Run("allows empty input", func(t *testing.T) {
		if err := RejectRootWorkspaceTypeDeletion(nil); err != nil {
			t.Errorf("expected nil input to pass through, got error: %v", err.Message)
		}
	})
}

// TestValidateWorkspaceTypeSlug_Format covers the format rules ValidateWorkspaceTypeSlug
// enforces before ever touching the database (must start with "/", only lowercase a-z
// and dashes after that, at most 50 characters) - this is the fix for the bug report
// reproduced via curl POST /workspaceType, where none of these rules existed yet.
// Inputs that fail one of these checks return before workspaceTypeSlugTaken's DB query
// runs, so they're safe to exercise without a live database; the uniqueness half (and
// the happy path, which does reach the DB) is covered by the curl reproduction in the
// PR/commit description instead.
func TestValidateWorkspaceTypeSlug_Format(t *testing.T) {
	cases := []struct {
		name        string
		slug        string
		wantMessage string
	}{
		{"missing leading slash", "wewe", "SlugMustStartWithSlash"},
		{"empty string", "", "SlugMustStartWithSlash"},
		{"uppercase letters", "/WeWe", "SlugInvalidFormat"},
		{"digits", "/wewe123", "SlugInvalidFormat"},
		{"underscore", "/we_we", "SlugInvalidFormat"},
		{"trailing slash", "/wewe/", "SlugInvalidFormat"},
		{"double leading slash", "//wewe", "SlugInvalidFormat"},
		{"too long", "/" + strings.Repeat("a", 60), "SlugTooLong"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			errs := ValidateWorkspaceTypeSlug(tc.slug, "")
			if len(errs) != 1 {
				t.Fatalf("expected exactly one error, got %d: %+v", len(errs), errs)
			}
			if (*errs[0].Message)["$"] != tc.wantMessage {
				t.Errorf("expected %s, got %v", tc.wantMessage, errs[0].Message)
			}
			if errs[0].Location != "slug" {
				t.Errorf("expected the error located at \"slug\", got %q", errs[0].Location)
			}
		})
	}

	t.Run("valid formats pass the format checks themselves", func(t *testing.T) {
		for _, slug := range []string{"/wewe", "/back-office", "/a", "/" + strings.Repeat("a", 49)} {
			if !workspaceTypeSlugPattern.MatchString(slug) || len(slug) > 50 {
				t.Errorf("expected %q to satisfy the format rules", slug)
			}
		}
	})
}

// TestRoleHasRootWildcardCapability covers the guard ValidateTheWorkspaceTypeEntity uses
// to keep a workspace type from being tied to a role that's effectively super-admin
// (holds the root.* wildcard capability), even when that role isn't literally the
// reserved "root" role by uniqueId.
func TestRoleHasRootWildcardCapability(t *testing.T) {
	cases := []struct {
		name         string
		capabilities []string
		want         bool
	}{
		{"has the wildcard", []string{"root.modules.abac.role", "root.*"}, true},
		{"wildcard only", []string{"root.*"}, true},
		{"ordinary capabilities only", []string{"root.modules.abac.role", "root.modules.abac.public-join-key"}, false},
		{"no capabilities", nil, false},
		{"similar but not the literal wildcard", []string{"root.modules.*"}, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			role := &RoleEntity{CapabilitiesListId: RoleCapabilitiesListIdOf(tc.capabilities)}
			if got := roleHasRootWildcardCapability(role); got != tc.want {
				t.Errorf("roleHasRootWildcardCapability(%v) = %v, want %v", tc.capabilities, got, tc.want)
			}
		})
	}
}
