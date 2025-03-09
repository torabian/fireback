package product

import "github.com/torabian/fireback/modules/workspaces"

func init() {
	// Override the implementation with our actual code.
	GetProductsCountActionImp = GetProductsCountAction
}
func GetProductsCountAction(
	q workspaces.QueryDSL) (*GetProductsCountActionResDto,
	*workspaces.IError,
) {
	// Implement the logic here.
	return &GetProductsCountActionResDto{
		TotalProduct: 4,
	}, nil
}
