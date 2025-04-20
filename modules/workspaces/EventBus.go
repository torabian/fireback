package workspaces

import (
	"context"
	"fmt"
	"log"
)

var ctx = context.Background()
var instance InstanceUserSocketManager
var EVENT_BUS_TOPIC string = "workspace.notifications"

func GetEventBusInstance() InstanceUserSocketManager {
	return instance
}

func StartEventBus() {

	if config.RedisEventsUrl == "" {
		log.Default().Println("Event bus is not enabled (redisEventsUrl is missing). Using internal")
		instance = NewLocalEventManager()

		return
	}

	// Try to use redis. If fails fallback to internal
	if redis, err := NewRedisManager(config.RedisEventsUrl); err == nil {
		instance = redis
	} else {
		instance = NewLocalEventManager()
	}

	go instance.Subscribe(ctx, EVENT_BUS_TOPIC)
}

// When a event bus is being sent over socket, we cast it to this struct
type EventBusSocketMessage struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Payload any    `json:"payload"`
}

// Replace this, based on how you want do handle the source of the event
// Fireback ABAC is also using this - it's enough
var ApplyQueryDslContextToEvent = func(event *Event, query QueryDSL) {
	event.InstanceSourceId = SERVER_INSTANCE
	event.SourceContext = map[string]string{
		"userId":      query.UserId,
		"workspaceId": query.WorkspaceId,
	}
}

// Event default source context
type SourceContext struct {
	UserId      string `json:"userId"`
	WorkspaceId string `json:"workspaceId"`
}

func ToSourceContext(i interface{}) (SourceContext, error) {
	m, ok := i.(map[string]interface{})
	if !ok {
		return SourceContext{}, fmt.Errorf("invalid type")
	}

	return SourceContext{
		UserId:      fmt.Sprintf("%v", m["userId"]),
		WorkspaceId: fmt.Sprintf("%v", m["workspaceId"]),
	}, nil
}

// Critical function which routes an event as notification and sends them via difference channels
// one important channel is Socket, the Other web push notification
// The biggest question is, how to understand a user has to be notified about an event, at all.
// Because the messages might contain the info that user has no access to it.
var RouteEvent = func(event Event) {
	HandleEventForSocketConnections(event)
}

// This logic seems to work, but I am not sure if it's directing the events deeply correctly.
func HandleEventForSocketConnections(event Event) {

	sourceContext, _ := ToSourceContext(event.SourceContext)
	if len(SocketSessionPool) > 0 {
		for workspaceId, workspace := range SocketSessionPool {

			if workspaceId != sourceContext.WorkspaceId {
				// Does not belong to this workspace, continue
				continue
			}

			for _, userConnections := range workspace {

				for _, connection := range userConnections {

					if event.Security != nil {

						query := QueryDSL{
							UserAccessPerWorkspace: &connection.URW,
							ActionRequires:         event.Security.ActionRequires,
						}

						meets, _ := MeetsAccessLevel(query, event.Security.AllowOnRoot)
						if !meets {
							continue
						}
					}

					connection.Connection.WriteJSON(event.ToNotification())
				}
			}
		}
	}
}
