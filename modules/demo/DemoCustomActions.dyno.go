package demo

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
)

var CustomerActivitySecurityModel *workspaces.SecurityModel = nil

type CustomerActivityActionReqDto struct {
	UniqueId []string `json:"uniqueId" yaml:"uniqueId"  validate:"required"       `
	// Datenano also has a text representation
}

func (x *CustomerActivityActionReqDto) RootObjectName() string {
	return "demo"
}

var CustomerActivityCommonCliFlagsOptional = []cli.Flag{}

func CustomerActivityActionReqValidator(dto *CustomerActivityActionReqDto) *workspaces.IError {
	err := workspaces.CommonStructValidatorPointer(dto, false)
	return err
}
func CastCustomerActivityFromCli(c *cli.Context) *CustomerActivityActionReqDto {
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
		return nil, nil
	}
	return CustomerActivityActionImp(req, q)
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
			Method:        "REACTIVE",
			Url:           "/customer/activity",
			SecurityModel: CustomerActivitySecurityModel,
			Handlers: []gin.HandlerFunc{
				workspaces.WithSocketAuthorization(&workspaces.SecurityModel{}, true),
				func(ctx *gin.Context) {
					writer := func(string) {}
					workspaces.HttpSocketRequest(ctx, func(query workspaces.QueryDSL, write func(string)) {
						writer = write

					}, func(query workspaces.QueryDSL, i interface{}) {

						// Generate a random number between 1 and 3

						var dto UserActivityFocusDto

						dat, _ := json.Marshal(i.(map[string]interface{}))

						fmt.Println(string(dat))
						json.Unmarshal(dat, &dto)

						for i := 0; i <= 0; i++ {

							activities := []*UserActivityActivities{}
							for _, v := range dto.Ids {
								v := v
								activityState := int64(rand.Intn(3) + 1)
								activities = append(activities, &UserActivityActivities{

									UniqueId: &v,
									Activity: &activityState,
								})

							}

							data := UserActivityDto{
								Activities: activities,
							}

							writer(data.Json())
							time.Sleep(time.Millisecond * 1000)

						}
					})

					// workspaces.HttpReactiveQuery(ctx,
					// 	func(query workspaces.QueryDSL, j chan bool, read chan map[string]interface{}) chan *UserActivityDto {

					// 		chanStream := make(chan *UserActivityDto)

					// 		go func() {
					// 			data := <-read

					// 			fmt.Println("Incomcing data", data)

					// 			for i := 0; i <= 10; i++ {
					// 				newUniq := "xxx"
					// 				activityState := int64(1)
					// 				chanStream <- &UserActivityDto{
					// 					Activities: []*UserActivityActivities{
					// 						{
					// 							UniqueId: &newUniq,
					// 							Activity: &activityState,
					// 						},
					// 					},
					// 				}
					// 			}
					// 		}()

					// 		return chanStream
					// 	},
					// )
				},
			},
			Format:         "POST_ONE",
			Action:         CustomerActivityActionFn,
			ResponseEntity: &UserActivityDto{},
			RequestEntity:  &CustomerActivityActionReqDto{},
		},
	}
	return routes
}

var DemoCustomActionsCli = []cli.Command{
	CustomerActivityActionCmd,
}
