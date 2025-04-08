package abac

import "github.com/torabian/fireback/modules/workspaces"

func init() {
	// Override the implementation with our actual code.
	ImportUserActionImp = ImportUserAction
}
func ImportUserAction(
	req *ImportUserActionReqDto,
	q workspaces.QueryDSL) (*OkayResponseDto,
	*workspaces.IError,
) {
	// Implement the logic here.
	return nil, nil
}
