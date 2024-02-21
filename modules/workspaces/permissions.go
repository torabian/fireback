package workspaces

import (
	"strings"
)

func MeetsAccessLevel(query QueryDSL, rootOrSystem bool) (bool, []string) {
	if rootOrSystem && (query.WorkspaceId != "root" && query.WorkspaceId != "system") {
		return false, []string{"SYSTEM_OR_ROOT_ALLOWED"}
	}

	meets := true
	missingPerms := []string{}

	if Contains(query.UserHas, "root/*") {
		return meets, missingPerms
	}

	for _, requiredPermission := range query.ActionRequires {

		// Two things needs to be checked, first if it does contain exact capability
		hasExactKey := Contains(query.UserHas, requiredPermission)

		hasParentalKey := false
		for _, a := range query.UserHas {
			if strings.Contains(requiredPermission, strings.ReplaceAll(a, "*", "")) {
				hasParentalKey = true
				continue
			}
		}

		if !hasExactKey && !hasParentalKey {
			meets = false
		}
	}

	if !meets {
		for _, perm := range query.ActionRequires {
			if Contains(query.UserHas, perm) {
				continue
			}

			missingPerms = append(missingPerms, perm)
		}
	}

	return meets, missingPerms
}
