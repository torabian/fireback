package resolve

import (
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
)

func WithAuthorizationPureDefault(context *AuthContextDto) (*AuthResultDto, *fireback.IError) {
	result := &AuthResultDto{}

	// workspaceId := context.WorkspaceId
	token := context.Token

	if token == "" {
		return nil, fireback.Create401Error(&AbacMessages.ProvideTokenInAuthorization, []string{})
	}

	user, err := GetUserFromToken(token)

	if err != nil {
		return nil, fireback.Create401Error(&AbacMessages.TokenNotFound, []string{
			maskToken(token),
		})
	}

	if user == nil {
		return nil, fireback.Create401Error(&AbacMessages.UserNotFoundOrDeleted, []string{})
	}

	access, accessError := GetUserAccessLevels(QueryDSL{UserId: user.UniqueId})

	if accessError != nil {
		return nil, accessError
	}

	query := QueryDSL{
		UserAccessPerWorkspace: access.UserAccessPerWorkspace,
		ActionRequires:         context.Capabilities,
		WorkspaceId:            context.WorkspaceId,
	}

	// MeetsAccessLevel checks query.UserHas/WorkspaceHas directly - they don't
	// come from UserAccessPerWorkspace automatically, so they need to be flattened
	// out of it for the *active* workspace (context.WorkspaceId) first. Without
	// this, both are always empty, which - combined with the now-fixed
	// MeetsAccessLevel actually enforcing its verdict instead of always returning
	// true - would deny every capability-gated action for everyone, root included.
	query.WorkspaceHas, query.UserHas = GetWorkspaceAndUserAccesses(query)

	meets, missing := MeetsAccessLevel(query, false)

	if !meets {
		return nil, fireback.Create401Error(&AbacMessages.NotEnoughPermission, missing)
	}

	result.UserId = emigo.NullableOf(user.UniqueId)
	result.User = user
	result.UserAccessPerWorkspace = access.UserAccessPerWorkspace
	result.SqlContext = GetSqlContext(access.UserAccessPerWorkspace, context.WorkspaceId, context.AllowCascade)

	// some actions could be restricted to happen only on root workspaces
	// here we check if user belongs to root, and the workspace selected needs to be root
	// as well. In some cases, user is in root workspace, but also exploring
	// another workspace for maintenance, should not be able to create root level content
	// in another workspace.

	// Fix this allow only on root.
	if context.Security != nil && context.Security.AllowOnRoot {
		if context.WorkspaceId != fireback.ROOT_VAR {
			return nil, &fireback.IError{
				HttpCode: 400,
				Message:  AbacMessages.ActionOnlyInRoot,
			}
		}
	}

	return result, nil
}
