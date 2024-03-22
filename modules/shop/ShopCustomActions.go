package shop

import (
	"fmt"

	"github.com/torabian/fireback/modules/workspaces"
)

func init() {
	ConfirmPurchaseActionImp = ConfirmPurchaseAction
	ShoppingCartCliCommands = append(ShoppingCartCliCommands, ShopCustomActionsCli...)
}

func ConfirmPurchaseAction(
	req *ConfirmPurchaseActionReqDto, q workspaces.QueryDSL,
) (*OrderEntity, *workspaces.IError) {

	fmt.Println(q.Json())
	currencyId := "USD"

	q.UniqueId = *req.BasketId
	q.Deep = true
	q.WithPreloads = []string{"Items.Product.Price.Variations"}
	cart, err := ShoppingCartActionGetOne(q)
	q.UniqueId = ""

	if err != nil {
		return nil, err
	}

	total, err2 := GetCartTotal(cart, *req.CurrencyId)
	if err2 != nil {
		return nil, err
	}

	shipping := "One apple park, 20202, 2092"
	invoiceNumber := "KJW38892"

	order := &OrderEntity{
		UserId:          &q.UserId,
		WorkspaceId:     &q.WorkspaceId,
		OrderStatusId:   &ORDER_STATUS_PENDING,
		PaymentStatusId: &PAYMENT_STATUS_PENDING,
		TotalPrice: &OrderTotalPrice{
			Amount:     &total,
			CurrencyId: &currencyId,
		},
		ShippingAddress: &shipping,
		InvoiceNumber:   &invoiceNumber,
	}

	return OrderActionCreate(order, q)

}

func GetCartTotal(cart *ShoppingCartEntity, currency string) (float64, error) {
	total := 0.0
	for _, item := range cart.Items {
		if item.Product == nil || item.Product.Price == nil {
			continue
		}

		for _, variation := range item.Product.Price.Variations {
			if *variation.CurrencyId == currency {
				total += *variation.Amount * *item.Quantity
			}
		}
	}

	return total, nil

}
