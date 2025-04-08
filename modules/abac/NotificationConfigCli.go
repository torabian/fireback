package abac

import (
	"fmt"
	reflect "reflect"

	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
)

var NotificationConfigTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the notificationConfig",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	NotificationConfigCliCommands = append(NotificationConfigCliCommands, NotificationConfigTestCmd)
}
func NotificationConfigActionGetOneByWorkspace(query workspaces.QueryDSL) (*NotificationConfigEntity, *workspaces.IError) {
	refl := reflect.ValueOf(&NotificationConfigEntity{})
	item, err := workspaces.GetOneEntityByWorkspace[NotificationConfigEntity](query, refl)
	entityNotificationConfigFormatter(item, query)
	return item, err
}

func GetRootNotificationConfig() (*NotificationConfigEntity, *workspaces.IError) {
	return NotificationConfigActionGetOneByWorkspace(workspaces.QueryDSL{WorkspaceId: ROOT_VAR})

}
