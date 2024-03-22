package shop

import (
	"github.com/torabian/fireback/modules/workspaces"
	"gorm.io/gorm"
)

func init() {
	ConfirmPurchaseActionImp = ConfirmPurchaseAction
	ShoppingCartCliCommands = append(ShoppingCartCliCommands, ShopCustomActionsCli...)
}

func ShopingCartItemsToOrder(items []*ShoppingCartItems) []*OrderItems {
	result := []*OrderItems{}
	price := 0.0
	for _, item := range items {
		snapshotJs := &workspaces.JSON{}
		snapshotJs.Scan(item.Product.Json())
		result = append(result, &OrderItems{
			Quantity:        item.Quantity,
			Price:           &price,
			ProductSnapshot: snapshotJs,
		})
	}

	return result
}

func ConfirmPurchaseAction(
	req *ConfirmPurchaseActionReqDto, q workspaces.QueryDSL,
) (*OrderEntity, *workspaces.IError) {
	order := &OrderEntity{}
	return workspaces.RunTransaction[OrderEntity](order, q, func(tx *gorm.DB) error {
		q.Tx = tx
		currencyId := "USD"

		q.UniqueId = *req.BasketId
		q.Deep = true
		q.WithPreloads = []string{"Items.Product.Price.Variations"}
		cart, err := ShoppingCartActionGetOne(q)

		if err != nil {
			return err
		}

		total, err2 := GetCartTotal(cart, *req.CurrencyId)
		if err2 != nil {
			return err
		}

		shipping := "One apple park, 20202, 2092"
		invoiceNumber := "KJW38892"

		order = &OrderEntity{
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
			Items:           ShopingCartItemsToOrder(cart.Items),
		}

		orderResult, err3 := OrderActionCreate(order, q)

		if err3 != nil {
			return err3
		}

		order = orderResult

		q.Query = "unique_id = " + q.UniqueId
		_, err4 := ShoppingCartActionRemove(q)
		if err4 != nil {
			return err4
		}
		return nil
	})

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
