package abac

import (
	"context"
	"fmt"

	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli/v3"
)

var NotificationConfigTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the notificationConfig",
	Action: func(ctx context.Context, c *cli.Command) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func NotificationConfigActionGetOneByWorkspace(query fireback.QueryDSL) (*NotificationConfigEntity, *fireback.IError) {
	return NotificationConfigActions.GetByWorkspace(query)
}

func GetRootNotificationConfig() (*NotificationConfigEntity, *fireback.IError) {
	return NotificationConfigActionGetOneByWorkspace(fireback.QueryDSL{WorkspaceId: ROOT_VAR})

}
