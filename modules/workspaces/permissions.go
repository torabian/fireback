package workspaces

import (
	"strings"
)

func meetsCheck(actionRequires []PermissionInfo, perms []string) bool {
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

func MeetsAccessLevel(query QueryDSL, rootOrSystem bool) (bool, []string) {
	if rootOrSystem && (query.WorkspaceId != "root" && query.WorkspaceId != "system") {
		return false, []string{"SYSTEM_OR_ROOT_ALLOWED"}
	}

	missingPerms := []string{}

	if Contains(query.UserHas, "root/*") && Contains(query.WorkspaceHas, "root/*") {
		return false, missingPerms
	}

	meetsUser := meetsCheck(query.ActionRequires, query.UserHas)
	meetsWorkspace := meetsCheck(query.ActionRequires, query.WorkspaceHas)

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
