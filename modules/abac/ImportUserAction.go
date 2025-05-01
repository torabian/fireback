package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	ImportUserActionImp = ImportUserAction
}
func ImportUserAction(
	req *ImportUserActionReqDto,
	q fireback.QueryDSL) (*OkayResponseDto,
	*fireback.IError,
) {
	// Implement the logic here.
	return nil, nil
}
