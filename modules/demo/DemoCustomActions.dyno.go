package demo
import (
	"github.com/gin-gonic/gin"
    "github.com/urfave/cli"
	"github.com/torabian/fireback/modules/workspaces"
)
var CustomerActivitySecurityModel *workspaces.SecurityModel = nil
type CustomerActivityActionReqDto struct {
    UniqueId   []string `json:"uniqueId" yaml:"uniqueId"  validate:"required"       `
    // Datenano also has a text representation
}
func ( x * CustomerActivityActionReqDto) RootObjectName() string {
	return "demo"
}
var CustomerActivityCommonCliFlagsOptional = []cli.Flag{
}
func CustomerActivityActionReqValidator(dto *CustomerActivityActionReqDto) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, false)
    return err
  }
func CastCustomerActivityFromCli (c *cli.Context) *CustomerActivityActionReqDto {
	template := &CustomerActivityActionReqDto{}
	return template
}
type customerActivityActionImpSig func(
    req *CustomerActivityActionReqDto, 
    q workspaces.QueryDSL) (*UserActivityDto,
    *workspaces.IError,
)
var CustomerActivityActionImp customerActivityActionImpSig
func CustomerActivityActionFn(
    req *CustomerActivityActionReqDto, 
    q workspaces.QueryDSL,
) (
    *UserActivityDto,
    *workspaces.IError,
) {
    if CustomerActivityActionImp == nil {
        return nil,  nil
    }
    return CustomerActivityActionImp(req,  q)
}
var CustomerActivityActionCmd cli.Command = cli.Command{
	Name:  "activity",
	Usage: "Returns the customer status regarding their activity",
	Flags: CustomerActivityCommonCliFlagsOptional,
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, CustomerActivitySecurityModel)
		dto := CastCustomerActivityFromCli(c)
		result, err := CustomerActivityActionFn(dto, query)
		workspaces.HandleActionInCli(c, result, err, map[string]map[string]string{})
	},
}
func DemoCustomActions() []workspaces.Module2Action {
	routes := []workspaces.Module2Action{
		{
			Method: "REACTIVE",
			Url:    "/customer/activity",
            SecurityModel: CustomerActivitySecurityModel,
			Handlers: []gin.HandlerFunc{
                workspaces.ReactiveSocketHandler(TestAction),
				func(c *gin.Context) {
                    // POST_ONE - reactive
                },
			},
			Format:         "POST_ONE",
			Action:         CustomerActivityActionFn,
			ResponseEntity: &UserActivityDto{},
			RequestEntity: &CustomerActivityActionReqDto{},
		},
	}
	return routes
}
var DemoCustomActionsCli = []cli.Command {
    CustomerActivityActionCmd,
}