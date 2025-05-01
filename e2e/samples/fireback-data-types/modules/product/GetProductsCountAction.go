package product

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	GetProductsCountActionImp = GetProductsCountAction
}
func GetProductsCountAction(
	q fireback.QueryDSL) (*GetProductsCountActionResDto,
	*fireback.IError,
) {
	// Implement the logic here.
	return &GetProductsCountActionResDto{
		TotalProduct: 4,
	}, nil
}
