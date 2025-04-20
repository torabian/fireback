package workspaces

import "context"

type InstanceUserSocketManager interface {
	AddUser(instanceId, userId string) error
	RemoveUser(instanceId, userId string) error
	ListUsers(instanceId string) ([]string, error)
	IsUserIn(instanceId string, userId string) (bool, error)
	FireEvent(q QueryDSL, event Event)
	Subscribe(ctx context.Context, channel string)
}
