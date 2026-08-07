package eventbus

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/torabian/fireback/modules/fireback"
)

// newTestSocketPair spins up a throwaway httptest websocket server and dials it. The
// server-side connection is what production code plugs into SocketConnection.Connection
// (what HandleEventForSocketConnections writes to); the client-side connection stands
// in for the browser/app on the other end, used here to read back whatever got
// written.
func newTestSocketPair(t *testing.T) (server *websocket.Conn, client *websocket.Conn) {
	t.Helper()

	upgrader := websocket.Upgrader{}
	serverConnCh := make(chan *websocket.Conn, 1)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Errorf("server upgrade: %v", err)
			return
		}
		serverConnCh <- conn
	}))
	t.Cleanup(ts.Close)

	wsURL := "ws" + strings.TrimPrefix(ts.URL, "http")
	c, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("client dial: %v", err)
	}
	t.Cleanup(func() { c.Close() })

	select {
	case s := <-serverConnCh:
		t.Cleanup(func() { s.Close() })
		return s, c
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for server-side websocket upgrade")
		return nil, nil
	}
}

// registerSocket adds a connection to SocketSessionPool the same way
// addUserToEventBus does (see EventBusSubscriptionActionImplementation.go), and
// removes it again once the test finishes.
func registerSocket(t *testing.T, workspaceId, userId string, conn *websocket.Conn, urw fireback.UserAccessPerWorkspaceDto) {
	t.Helper()

	socketMutex.Lock()
	if SocketSessionPool[workspaceId] == nil {
		SocketSessionPool[workspaceId] = make(map[string][]*SocketConnection)
	}
	SocketSessionPool[workspaceId][userId] = append(SocketSessionPool[workspaceId][userId], &SocketConnection{
		UserId:     userId,
		Connection: conn,
		URW:        urw,
		UniqueId:   "test-" + workspaceId + "-" + userId,
	})
	socketMutex.Unlock()

	t.Cleanup(func() {
		socketMutex.Lock()
		delete(SocketSessionPool, workspaceId)
		socketMutex.Unlock()
	})
}

func readNotification(t *testing.T, client *websocket.Conn) Notification {
	t.Helper()

	client.SetReadDeadline(time.Now().Add(2 * time.Second))
	_, data, err := client.ReadMessage()
	if err != nil {
		t.Fatalf("reading notification: %v", err)
	}

	var n Notification
	if err := json.Unmarshal(data, &n); err != nil {
		t.Fatalf("decoding notification %q: %v", data, err)
	}
	return n
}

func expectNoMessage(t *testing.T, client *websocket.Conn) {
	t.Helper()

	client.SetReadDeadline(time.Now().Add(300 * time.Millisecond))
	_, _, err := client.ReadMessage()
	if err == nil {
		t.Fatal("expected no message to be delivered, but one arrived")
	}
}

func TestLocalEventManager_UserRegistry(t *testing.T) {
	m := NewLocalEventManager()

	if in, _ := m.IsUserIn("inst-1", "u1"); in {
		t.Fatal("expected u1 not to be registered in inst-1 yet")
	}

	if err := m.AddUser("inst-1", "u1"); err != nil {
		t.Fatalf("AddUser: %v", err)
	}
	if err := m.AddUser("inst-1", "u2"); err != nil {
		t.Fatalf("AddUser: %v", err)
	}

	if in, _ := m.IsUserIn("inst-1", "u1"); !in {
		t.Fatal("expected u1 to be registered in inst-1")
	}

	users, _ := m.ListUsers("inst-1")
	if len(users) != 2 {
		t.Fatalf("expected 2 users, got %d (%v)", len(users), users)
	}

	if err := m.RemoveUser("inst-1", "u1"); err != nil {
		t.Fatalf("RemoveUser: %v", err)
	}
	if in, _ := m.IsUserIn("inst-1", "u1"); in {
		t.Fatal("expected u1 to have been removed from inst-1")
	}

	users, _ = m.ListUsers("inst-1")
	if len(users) != 1 || users[0] != "u2" {
		t.Fatalf("expected only u2 left in inst-1, got %v", users)
	}

	// A different instanceId is its own independent registry.
	if in, _ := m.IsUserIn("inst-2", "u2"); in {
		t.Fatal("expected u2 not to be registered in the unrelated inst-2")
	}
}

// TestHandleEventForSocketConnections_RoutesByWorkspace verifies an event only
// reaches connections whose workspace matches event.SourceContext.workspaceId - a
// connection on a different workspace must not see it.
func TestHandleEventForSocketConnections_RoutesByWorkspace(t *testing.T) {
	serverA, clientA := newTestSocketPair(t)
	registerSocket(t, "ws-a", "user-a", serverA, fireback.UserAccessPerWorkspaceDto{})

	serverB, clientB := newTestSocketPair(t)
	registerSocket(t, "ws-b", "user-b", serverB, fireback.UserAccessPerWorkspaceDto{})

	event := Event{
		Name:    "SomethingHappened",
		Payload: map[string]any{"hello": "world"},
		SourceContext: map[string]string{
			"workspaceId": "ws-a",
			"userId":      "user-a",
		},
	}

	HandleEventForSocketConnections(event)

	got := readNotification(t, clientA)
	if got.Name != "SomethingHappened" {
		t.Fatalf("expected ws-a's connection to receive the notification, got %+v", got)
	}

	expectNoMessage(t, clientB)
}

// TestHandleEventForSocketConnections_AccessControl verifies Event.Security gates
// delivery through meetsAccessLevel: no check wired up fails closed, a check that
// denies drops the message, and a check that allows delivers it.
func TestHandleEventForSocketConnections_AccessControl(t *testing.T) {
	prevMeets := meetsAccessLevel
	prevFireback := fireback.MeetsAccessLevel
	t.Cleanup(func() {
		meetsAccessLevel = prevMeets
		fireback.MeetsAccessLevel = prevFireback
	})

	securedEvent := Event{
		Name: "SecretHappened",
		SourceContext: map[string]string{
			"workspaceId": "ws-secure",
			"userId":      "user-1",
		},
		Security: &fireback.SecurityModel{
			ActionRequires: []fireback.PermissionInfo{{CompleteKey: "root.some.perm"}},
		},
	}

	// A gorilla/websocket connection is considered broken by the library after a
	// read times out, so each sub-case below gets its own fresh connection pair
	// instead of reusing one across an expectNoMessage/readNotification sequence.

	t.Run("fails closed when nothing is wired up", func(t *testing.T) {
		server, client := newTestSocketPair(t)
		registerSocket(t, "ws-secure", "user-1", server, fireback.UserAccessPerWorkspaceDto{})

		meetsAccessLevel = nil
		fireback.MeetsAccessLevel = nil
		HandleEventForSocketConnections(securedEvent)
		expectNoMessage(t, client)
	})

	t.Run("drops the message when the check denies", func(t *testing.T) {
		server, client := newTestSocketPair(t)
		registerSocket(t, "ws-secure", "user-1", server, fireback.UserAccessPerWorkspaceDto{})

		meetsAccessLevel = func(query fireback.QueryDSL, onlyRoot bool) (bool, []string) {
			return false, []string{"root.some.perm"}
		}
		HandleEventForSocketConnections(securedEvent)
		expectNoMessage(t, client)
	})

	t.Run("delivers the message when the check allows", func(t *testing.T) {
		server, client := newTestSocketPair(t)
		registerSocket(t, "ws-secure", "user-1", server, fireback.UserAccessPerWorkspaceDto{})

		meetsAccessLevel = func(query fireback.QueryDSL, onlyRoot bool) (bool, []string) {
			return true, nil
		}
		HandleEventForSocketConnections(securedEvent)
		got := readNotification(t, client)
		if got.Name != "SecretHappened" {
			t.Fatalf("expected the secured event to be delivered once allowed, got %+v", got)
		}
	})

	t.Run("falls back to fireback.MeetsAccessLevel when not overridden", func(t *testing.T) {
		server, client := newTestSocketPair(t)
		registerSocket(t, "ws-secure", "user-1", server, fireback.UserAccessPerWorkspaceDto{})

		meetsAccessLevel = nil
		fireback.MeetsAccessLevel = func(query fireback.QueryDSL, onlyRoot bool) (bool, []string) {
			return true, nil
		}
		HandleEventForSocketConnections(securedEvent)
		got := readNotification(t, client)
		if got.Name != "SecretHappened" {
			t.Fatalf("expected the fireback.MeetsAccessLevel fallback to deliver, got %+v", got)
		}
	})

	t.Run("skips the check entirely when Security is unset", func(t *testing.T) {
		server, client := newTestSocketPair(t)
		registerSocket(t, "ws-secure", "user-1", server, fireback.UserAccessPerWorkspaceDto{})

		meetsAccessLevel = nil
		fireback.MeetsAccessLevel = nil
		HandleEventForSocketConnections(Event{
			Name: "PublicHappened",
			SourceContext: map[string]string{
				"workspaceId": "ws-secure",
				"userId":      "user-1",
			},
		})
		got := readNotification(t, client)
		if got.Name != "PublicHappened" {
			t.Fatalf("expected the unsecured event to be delivered regardless, got %+v", got)
		}
	})
}

// TestStartEventBus_Local_FireEventDelivers exercises the full local (non-redis) path
// end to end: StartEventBus("") selects the in-process backend and starts listening,
// then GetEventBusInstance().FireEvent(...) round-trips through it and reaches a
// connected socket exactly like a real subscriber would. Runs after the package's
// other tests so SocketSessionPool/meetsAccessLevel are back to their zero state
// (StartEventBus itself is only ever called once for the whole test binary - see
// TestMain - since gookit/event's global registry has no way to unregister a listener,
// calling it twice would double-deliver every subsequent event).
func TestStartEventBus_Local_FireEventDelivers(t *testing.T) {
	startEventBusOnce(t)

	if _, ok := GetEventBusInstance().(*LocalEventManager); !ok {
		t.Fatalf("expected StartEventBus(\"\") to select the local backend, got %T", GetEventBusInstance())
	}

	server, client := newTestSocketPair(t)
	registerSocket(t, "ws-fire", "user-fire", server, fireback.UserAccessPerWorkspaceDto{})

	event := Event{
		Name:    "FiredThroughTheBus",
		Payload: "payload",
		SourceContext: map[string]string{
			"workspaceId": "ws-fire",
			"userId":      "user-fire",
		},
	}

	GetEventBusInstance().FireEvent(fireback.QueryDSL{}, event)

	got := readNotification(t, client)
	if got.Name != "FiredThroughTheBus" {
		t.Fatalf("expected the fired event to arrive over the local bus, got %+v", got)
	}
}

var startEventBusOnceGuard = false

// startEventBusOnce calls StartEventBus("") exactly once for the whole test binary -
// gookit/event's DefaultEM has no listener-removal API, so calling it (and therefore
// Subscribe) more than once would register a second "locale_events" listener and
// double-deliver every event fired afterwards. Safe to call synchronously with no
// extra wait: StartEventBus("") runs LocalEventManager.Subscribe synchronously (see
// EventBus.go), so its listener is already registered by the time it returns.
func startEventBusOnce(t *testing.T) {
	t.Helper()
	if startEventBusOnceGuard {
		return
	}
	startEventBusOnceGuard = true
	StartEventBus("")
}
