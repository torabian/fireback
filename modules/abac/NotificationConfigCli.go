package workspaces

import (
	"fmt"
	reflect "reflect"

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
func NotificationConfigActionGetOneByWorkspace(query QueryDSL) (*NotificationConfigEntity, *IError) {
	refl := reflect.ValueOf(&NotificationConfigEntity{})
	item, err := GetOneEntityByWorkspace[NotificationConfigEntity](query, refl)
	entityNotificationConfigFormatter(item, query)
	return item, err
}

func GetRootNotificationConfig() (*NotificationConfigEntity, *IError) {
	return NotificationConfigActionGetOneByWorkspace(QueryDSL{WorkspaceId: ROOT_VAR})

}
