package abac

import (
	"encoding/json"

	"github.com/torabian/emi/emigo"
)

// These helpers back the generated CLI flag casting for dto fields which relate to an
// entity (type: one/collection). They are invoked by emigo's CapturePossible* helpers,
// which currently JSON-decode the raw flag value directly rather than calling the
// generator, but the generator still has to exist with the right signature for the
// generated code (UserSessionDto, WorkspaceInvitationDto, ...) to compile.

func CastPassportEntityFromCli(c emigo.CliCastable) PassportEntity {
	var result PassportEntity
	json.Unmarshal([]byte(c.String("passport")), &result)
	return result
}

func CastUserWorkspaceEntityFromCli(c emigo.CliCastable) UserWorkspaceEntity {
	var result UserWorkspaceEntity
	json.Unmarshal([]byte(c.String("user-workspace")), &result)
	return result
}

func CastUserEntityFromCli(c emigo.CliCastable) UserEntity {
	var result UserEntity
	json.Unmarshal([]byte(c.String("user")), &result)
	return result
}

func CastWorkspaceEntityFromCli(c emigo.CliCastable) WorkspaceEntity {
	var result WorkspaceEntity
	json.Unmarshal([]byte(c.String("workspace")), &result)
	return result
}
