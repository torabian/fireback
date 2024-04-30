package shop
import (
	"github.com/gin-gonic/gin"
    "github.com/urfave/cli"
	"github.com/torabian/fireback/modules/workspaces"
)
var ConfirmPurchaseSecurityModel *workspaces.SecurityModel = nil
type ConfirmPurchaseActionReqDto struct {
    BasketId   *string `json:"basketId" yaml:"basketId"  validate:"required"       `
    // Datenano also has a text representation
    CurrencyId   *string `json:"currencyId" yaml:"currencyId"  validate:"required"       `
    // Datenano also has a text representation
}
func ( x * ConfirmPurchaseActionReqDto) RootObjectName() string {
	return "shop"
}
var ConfirmPurchaseCommonCliFlagsOptional = []cli.Flag{
    &cli.StringFlag{
      Name:     "basket-id",
      Required: true,
      Usage:    "basketId",
    },
    &cli.StringFlag{
      Name:     "currency-id",
      Required: true,
      Usage:    "currencyId",
    },
}
func ConfirmPurchaseActionReqValidator(dto *ConfirmPurchaseActionReqDto) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, false)
    return err
  }
func CastConfirmPurchaseFromCli (c *cli.Context) *ConfirmPurchaseActionReqDto {
	template := &ConfirmPurchaseActionReqDto{}
      if c.IsSet("basket-id") {
        value := c.String("basket-id")
        template.BasketId = &value
      }
      if c.IsSet("currency-id") {
        value := c.String("currency-id")
        template.CurrencyId = &value
      }
	return template
}
type confirmPurchaseActionImpSig func(
    req *ConfirmPurchaseActionReqDto, 
    q workspaces.QueryDSL) (*OrderEntity,
    *workspaces.IError,
)
var ConfirmPurchaseActionImp confirmPurchaseActionImpSig
func ConfirmPurchaseActionFn(
    req *ConfirmPurchaseActionReqDto, 
    q workspaces.QueryDSL,
) (
    *OrderEntity,
    *workspaces.IError,
) {
    if ConfirmPurchaseActionImp == nil {
        return nil,  nil
    }
    return ConfirmPurchaseActionImp(req,  q)
}
var ConfirmPurchaseActionCmd cli.Command = cli.Command{
	Name:  "purchase",
	Usage: "Confirms a purchase, from a basket and converts it into an order",
	Flags: ConfirmPurchaseCommonCliFlagsOptional,
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, ConfirmPurchaseSecurityModel)
		dto := CastConfirmPurchaseFromCli(c)
		result, err := ConfirmPurchaseActionFn(dto, query)
		workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
	},
}
func ShopCustomActions() []workspaces.Module2Action {
	routes := []workspaces.Module2Action{
		{
			Method: "POST",
			Url:    "/purchase/confirm",
            SecurityModel: ConfirmPurchaseSecurityModel,
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
                    // POST_ONE - post
                        workspaces.HttpPostEntity(c, ConfirmPurchaseActionFn)
                },
			},
			Format:         "POST_ONE",
			Action:         ConfirmPurchaseActionFn,
			ResponseEntity: &OrderEntity{},
			RequestEntity: &ConfirmPurchaseActionReqDto{},
		},
	}
	return routes
}
var ShopCustomActionsCli = []cli.Command {
    ConfirmPurchaseActionCmd,
}