package resolve

import (
	"strings"

	"github.com/torabian/fireback/modules/fireback"
)

// It would convert the current selected role_id and workspace_id into a sql
// with given permissions to make the queries do not need check that again
func GetSqlContext(x *UserAccessPerWorkspaceDto, activeWorkspaceId string, allowCascade bool) string {
	conditions := []string{}

	// Let's allow the user to see everything which they belong to
	// but usually it's not necessary, because they are focused on one workspace at the moment
	if allowCascade {
		for workspaceId := range *x {
			conditions = append(conditions, fireback.RealEscape("workspace_id in (?)", workspaceId))
		}
	} else {
		userBelongsToWorkspace := false
		for workspaceId := range *x {
			if workspaceId == activeWorkspaceId {
				userBelongsToWorkspace = true

				// Important to break, otherwise can show other workspaces
				break
			}
		}

		if userBelongsToWorkspace {
			conditions = append(conditions, fireback.RealEscape("workspace_id in (?)", activeWorkspaceId))
		}
	}

	sql := strings.Join(conditions, " or ")

	return sql
}
