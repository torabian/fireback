package shop

import (
	"embed"
	"fmt"

	// "github.com/urfave/cli"
	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

//go:embed *Module3.yml
var Module2Definitions embed.FS

func ShopModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name:        "shop",
		Definitions: &Module2Definitions,
	}

	module.ProvideMockImportHandler(func() {
		// FormImportMocks()
	})

	module.ProvideSeederImportHandler(func() {
		DiscountTypeSyncSeeders()
		DiscountScopeSyncSeeders()
		OrderStatusSyncSeeders()
		PaymentStatusSyncSeeders()
		PaymentMethodSyncSeeders()
	})

	module.ProvideMockWriterHandler(func(languages []string) {

	})

	module.ProvidePermissionHandler(
		ALL_PRODUCT_PERMISSIONS,
		ALL_PRODUCT_SUBMISSION_PERMISSIONS,
		ALL_BRAND_PERMISSIONS,
		ALL_CATEGORY_PERMISSIONS,
		ALL_TAG_PERMISSIONS,
		ALL_DISCOUNT_TYPE_PERMISSIONS,
		ALL_DISCOUNT_SCOPE_PERMISSIONS,
		ALL_DISCOUNT_CODE_PERMISSIONS,
		ALL_ORDER_STATUS_PERMISSIONS,
		ALL_ORDER_PERMISSIONS,
		ALL_PAYMENT_METHOD_PERMISSIONS,
		ALL_PAYMENT_STATUS_PERMISSIONS,
		ALL_SHOPPING_CART_PERMISSIONS,
		ALL_PERM_SHOP_MODULE,
	)

	module.Actions = [][]workspaces.Module2Action{
		GetProductSubmissionModule2Actions(),
		GetProductModule2Actions(),
		GetTagModule2Actions(),
		GetCategoryModule2Actions(),
		GetBrandModule2Actions(),
		GetDiscountTypeModule2Actions(),
		GetDiscountScopeModule2Actions(),
		GetDiscountCodeModule2Actions(),
		GetOrderStatusModule2Actions(),
		GetPaymentMethodModule2Actions(),
		GetPaymentStatusModule2Actions(),
		GetShoppingCartModule2Actions(),
		GetOrderModule2Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) {
		if err := dbref.AutoMigrate(
			&ProductEntity{},
			&ProductFields{},
			&ProductSubmissionEntity{},
			&ProductSubmissionValues{},
			&ProductSubmissionPrice{},
			&ProductSubmissionPriceVariations{},
			&CategoryEntity{},
			&CategoryEntityPolyglot{},
			&TagEntity{},
			&TagEntityPolyglot{},
			&BrandEntity{},
			&BrandEntityPolyglot{},
			&DiscountTypeEntity{},
			&DiscountTypeEntityPolyglot{},
			&DiscountScopeEntity{},
			&DiscountScopeEntityPolyglot{},
			&DiscountCodeEntity{},
			&OrderEntity{},
			&OrderItems{},
			&OrderTotalPrice{},
			&OrderStatusEntity{},
			&OrderStatusEntityPolyglot{},
			&PaymentMethodEntity{},
			&PaymentMethodEntityPolyglot{},
			&PaymentStatusEntity{},
			&PaymentStatusEntityPolyglot{},
			&ShoppingCartEntity{},
			&ShoppingCartItems{},
		); err != nil {
			fmt.Println(err.Error())
		}

	})

	module.ProvideCliHandlers([]cli.Command{
		ProductCliFn(),
		ProductSubmissionCliFn(),
		CategoryCliFn(),
		TagCliFn(),
		BrandCliFn(),
		DiscountTypeCliFn(),
		DiscountScopeCliFn(),
		DiscountCodeCliFn(),
		OrderStatusCliFn(),
		PaymentMethodCliFn(),
		PaymentStatusCliFn(),
		OrderCliFn(),
		ShoppingCartCliFn(),
	})

	return module
}

func QueryProductSubmissionsReact(query workspaces.QueryDSL, chanStream chan *workspaces.ReactiveSearchResultDto) {
	actionFnNavigate := "navigate"

	query.Query = "name %" + query.SearchPhrase + "%"
	items, _, _ := ProductSubmissionActionQuery(query)

	product := "$product"
	products := "products"
	for _, item := range items {
		loc := "/product-submission/" + item.UniqueId

		uid := workspaces.UUID()

		chanStream <- &workspaces.ReactiveSearchResultDto{
			Phrase:      item.Name,
			Icon:        &product,
			Description: item.Name,
			Group:       &products,
			ActionFn:    &actionFnNavigate,
			UiLocation:  &loc,
			UniqueId:    &uid,
		}
	}

}
