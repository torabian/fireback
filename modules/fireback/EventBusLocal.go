package fireback

import (
	"context"
	"encoding/json"

	"github.com/gookit/event"
)

type LocalEventManager struct{}

func NewLocalEventManager() *LocalEventManager {
	return &LocalEventManager{}
}

// Local registry of instanceId => userIds (like Redis sets)
var localInstanceUsers = map[string]map[string]bool{}

func (l *LocalEventManager) key(instanceId string) string {
	return instanceId
}

func (l *LocalEventManager) AddUser(instanceId, userID string) error {
	if localInstanceUsers[instanceId] == nil {
		localInstanceUsers[instanceId] = make(map[string]bool)
	}
	localInstanceUsers[instanceId][userID] = true
	return nil
}

func (l *LocalEventManager) RemoveUser(instanceId, userID string) error {
	if users, ok := localInstanceUsers[instanceId]; ok {
		delete(users, userID)
		if len(users) == 0 {
			delete(localInstanceUsers, instanceId)
		}
	}
	return nil
}

func (l *LocalEventManager) ListUsers(instanceId string) ([]string, error) {
	var users []string
	for uid := range localInstanceUsers[instanceId] {
		users = append(users, uid)
	}
	return users, nil
}

func (l *LocalEventManager) IsUserIn(instanceId, userID string) (bool, error) {
	_, ok := localInstanceUsers[instanceId][userID]
	return ok, nil
}

func (l *LocalEventManager) FireEvent(q QueryDSL, e Event) {

	event.MustFire("locale_events", map[string]any{
		"content": e.Json(),
	})
}

func (l *LocalEventManager) Subscribe(ctx context.Context, channel string) {
	event.On("locale_events", event.ListenerFunc(func(e event.Event) error {
		payload := e.Data()["content"]
		var event *Event
		json.Unmarshal([]byte(payload.(string)), &event)

		RouteEvent(*event)
		return nil
	}))
}
