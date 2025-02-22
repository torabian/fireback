package workspaces

func init() {
	// Override the implementation with our actual code.
	ImportUserActionImp = ImportUserAction
}
func ImportUserAction(
	req *ImportUserActionReqDto,
	q QueryDSL) (*OkayResponseDto,
	*IError,
) {
	// Implement the logic here.
	return nil, nil
}
