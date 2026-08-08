package eventbus

import (
	"context"
	"fmt"

	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/owner"
)

var ctx = context.Background()
var instance InstanceUserSocketManager
var EVENT_BUS_TOPIC string = "workspace.notifications"

// SourceContext satisfies owner.Owner.
var _ owner.Owner = SourceContext{}

// meetsAccessLevel filters which connections get to see an Event with Security set -
// wired from EventBusModuleConfig.MeetsAccessLevel by ModuleSetup (see
// EventBusModule.go), since it depends on a project's own access-control model (e.g.
// abac's), which this module can't import without creating a dependency in the wrong
// direction. If left unset, any Event with Security set is dropped for every
// connection (fails closed) rather than panicking.
var meetsAccessLevel func(query fireback.QueryDSL, onlyRoot bool) (bool, []string)

type InstanceUserSocketManager interface {
	AddUser(instanceId, userId string) error
	RemoveUser(instanceId, userId string) error
	ListUsers(instanceId string) ([]string, error)
	IsUserIn(instanceId string, userId string) (bool, error)
	FireEvent(q fireback.QueryDSL, event Event)
	Subscribe(ctx context.Context, channel string)
}

func GetEventBusInstance() InstanceUserSocketManager {
	return instance
}

func init() {
	instance = NewLocalEventManager()
}

// StartEventBus is called from EventBusModuleConfig's OnAppStart hook (see
// EventBusModule.go) - it's no longer started automatically by fireback itself, so a
// project only pays for this goroutine if it actually registers the eventbus module.
// redisEventsUrl empty means events stay local to this instance (no distribution
// across other instances) - but Subscribe still has to run either way, otherwise
// FireEvent on the local backend just publishes into a void nothing is listening to.
func StartEventBus(redisEventsUrl string) {

	local := NewLocalEventManager()
	instance = local

	if redisEventsUrl != "" {
		// Try to use redis. If fails fallback to internal
		if redis, err := NewRedisManager(redisEventsUrl); err == nil {
			instance = redis
			// RedisManager.Subscribe blocks forever reading its pubsub channel, so
			// it has to run in the background - a caller firing events right after
			// StartEventBus returns has to tolerate the small async window before
			// this goroutine actually gets scheduled and subscribes.
			go instance.Subscribe(ctx, EVENT_BUS_TOPIC)
			return
		}
	}

	// LocalEventManager.Subscribe only registers a listener and returns
	// immediately - unlike the redis case above, there's no need to run it in the
	// background, and running it synchronously means FireEvent calls made right
	// after StartEventBus returns are guaranteed to reach it.
	local.Subscribe(ctx, EVENT_BUS_TOPIC)
}

// When a event bus is being sent over socket, we cast it to this struct
type EventBusSocketMessage struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Payload any    `json:"payload"`
}

// ApplyOwnerContextToEvent stamps event with who it belongs to. Takes owner.Owner
// rather than fireback.QueryDSL directly - all it ever needed from QueryDSL was
// GetWorkspaceId/GetUserId, so any caller can pass its own type here (a QueryDSL, or
// anything else satisfying owner.Owner) without eventbus needing to know about it.
// Replace this, based on how you want do handle the source of the event.
var ApplyOwnerContextToEvent = func(event *Event, o owner.Owner) {
	event.InstanceSourceId = fireback.SERVER_INSTANCE
	event.SourceContext = map[string]string{
		"userId":      o.GetUserId(),
		"workspaceId": o.GetWorkspaceId(),
	}
}

// Event default source context - also satisfies owner.Owner (see GetWorkspaceId/
// GetUserId below), since it's reconstructed from an Event's SourceContext and is
// itself just an identity: who this event belongs to.
type SourceContext struct {
	UserId      string `json:"userId"`
	WorkspaceId string `json:"workspaceId"`
}

func (s SourceContext) GetWorkspaceId() string {
	return s.WorkspaceId
}

func (s SourceContext) GetUserId() string {
	return s.UserId
}

// ToSourceContext reads back the workspace/user an event belongs to.
// event.SourceContext arrives as map[string]interface{} whenever the event went
// through a JSON round-trip first (the redis-backed manager, or the local manager's
// own gookit/event listener - see EventBusLocal.go/EventBusRedis.go's Subscribe) -
// but as the literal map[string]string ApplyOwnerContextToEvent builds when
// RouteEvent/HandleEventForSocketConnections is called directly, in-process, without
// going through either. Both are accepted so calling either path works the same way.
func ToSourceContext(i interface{}) (SourceContext, error) {
	switch m := i.(type) {
	case map[string]interface{}:
		return SourceContext{
			UserId:      fmt.Sprintf("%v", m["userId"]),
			WorkspaceId: fmt.Sprintf("%v", m["workspaceId"]),
		}, nil
	case map[string]string:
		return SourceContext{
			UserId:      m["userId"],
			WorkspaceId: m["workspaceId"],
		}, nil
	default:
		return SourceContext{}, fmt.Errorf("invalid type")
	}
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

						check := meetsAccessLevel
						if check == nil {
							// EventBusModuleConfig.MeetsAccessLevel wasn't overridden -
							// fall back to fireback.MeetsAccessLevel (wired by abac, see
							// AbacModule.go), read lazily here rather than snapshotted at
							// ModuleSetup time so it doesn't matter whether eventbus or
							// abac's own module got set up first.
							check = fireback.MeetsAccessLevel
						}
						if check == nil {
							continue
						}

						query := fireback.QueryDSL{
							UserAccessPerWorkspace: &connection.URW,
							ActionRequires:         event.Security.ActionRequires,
						}

						meets, _ := check(query, event.Security.AllowOnRoot)
						if !meets {
							continue
						}
					}

					if connection.Connection != nil {
						connection.Connection.WriteJSON(event.ToNotification())
					}

				}
			}
		}
	}
}
