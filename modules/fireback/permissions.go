package fireback

import (
	"strings"
)

func MeetsCheck(actionRequires []PermissionInfo, perms []string) bool {
	meets := true
	for _, requiredPermission := range actionRequires {

		// Two things needs to be checked, first if it does contain exact capability
		hasExactKey := Contains(perms, requiredPermission.CompleteKey)

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

func MeetsAccessLevel(query QueryDSL, onlyRoot bool) (bool, []string) {
	if onlyRoot && (query.WorkspaceId != ROOT_VAR && query.WorkspaceId != "system") {
		return false, []string{"SYSTEM_OR_ROOT_ALLOWED"}
	}

	missingPerms := []string{}

	if Contains(query.UserHas, ROOT_ALL_ACCESS) && Contains(query.WorkspaceHas, ROOT_ALL_ACCESS) {
		return false, missingPerms
	}

	meetsUser := MeetsCheck(query.ActionRequires, query.UserHas)
	meetsWorkspace := MeetsCheck(query.ActionRequires, query.WorkspaceHas)

	if !meetsUser || !meetsWorkspace {
		for _, perm := range query.ActionRequires {
			if Contains(query.UserHas, perm.CompleteKey) {
				continue
			}

			missingPerms = append(missingPerms, perm.CompleteKey)
		}
	}

	return true, missingPerms
}
