package resolve

import (
	"slices"
	"strings"

	"github.com/torabian/fireback/modules/fireback/application"
)

func MeetsAccessLevel(query MsgContext, onlyRoot bool) (bool, []string) {
	if onlyRoot && (query.WorkspaceId != ROOT_VAR) {
		return false, []string{"SYSTEM_OR_ROOT_ALLOWED"}
	}

	missingPerms := []string{}

	// A user (and their active workspace) holding the full root.* wildcard always
	// meets every requirement - no need to check individual capabilities. This was
	// previously inverted (returned false, i.e. denied, for exactly this case) and
	// the real verdict below was discarded entirely (always returned true) - see
	// the fix note on the final return.
	if slices.Contains(query.UserHas, ROOT_ALL_ACCESS) && slices.Contains(query.WorkspaceHas, ROOT_ALL_ACCESS) {
		return true, missingPerms
	}

	meetsUser := MeetsCheck(query.ActionRequires, query.UserHas)
	meetsWorkspace := MeetsCheck(query.ActionRequires, query.WorkspaceHas)

	if !meetsUser || !meetsWorkspace {
		for _, perm := range query.ActionRequires {
			if slices.Contains(query.UserHas, perm.CompleteKey) {
				continue
			}

			missingPerms = append(missingPerms, perm.CompleteKey)
		}
	}

	// Bug fix: this always returned true regardless of meetsUser/meetsWorkspace,
	// so every action requiring a capability the caller didn't actually have was
	// silently allowed through - the only real gate left was the onlyRoot check
	// above. missingPerms was still computed correctly, just never acted on.
	return meetsUser && meetsWorkspace, missingPerms
}

func MeetsCheck(actionRequires []application.PermissionInfo, perms []string) bool {
	meets := true
	for _, requiredPermission := range actionRequires {

		// Two things needs to be checked, first if it does contain exact capability
		hasExactKey := slices.Contains(perms, requiredPermission.CompleteKey)
		hasParentalKey := false

		for _, a := range perms {
			if strings.Contains(requiredPermission.CompleteKey, strings.ReplaceAll(a, "*", "")) {
				hasParentalKey = true
				continue
			}
		}

		if !hasExactKey && !hasParentalKey {
			meets = false
		}
	}
	return meets
}
