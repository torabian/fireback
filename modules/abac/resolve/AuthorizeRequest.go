package resolve

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/torabian/fireback/modules/fireback"
)

func AuthorizeRequest(securityModel *SecurityModel, c *gin.Context) (*AuthResultDto, error) {
	q := fireback.ExtractQueryDslFromGinContext(c)

	// A WebSocket handshake can't carry custom headers from the browser, so
	// clients of a reactive action (e.g. eventbus's /ws, opened by the UI's
	// useFirebackSocket) send token/workspace-id/role-id as query params
	// instead. And by the time a reactive action's factory calls into this
	// function (see EventBusSubscriptionActionSig), gorilla-websocket has
	// already hijacked the connection to complete the upgrade - gin's
	// response writer can no longer send an HTTP response onto it, so the
	// c.AbortWithStatusJSON below would crash with "response.Write on
	// hijacked connection" if reached. Both quirks trace back to the same
	// signal, checked once here: is this request a websocket upgrade.
	isWebsocketUpgrade := websocket.IsWebSocketUpgrade(c.Request)

	wi := c.GetHeader("Workspace-id")
	ri := c.GetHeader("Role-id")
	tk := c.GetHeader("Authorization")

	if isWebsocketUpgrade {
		if wi == "" {
			wi = c.Query("workspaceId")
		}
		if ri == "" {
			ri = c.Query("roleId")
		}
		if tk == "" {
			tk = c.Query("token")
		}
	}

	ck, ckerr := c.Cookie("authorization")

	if ckerr == nil && ck != "" && tk == "" {
		// If on secure cookie we have the authorization, and it's not defined on headers or query params.
		tk = ck
	}

	context := &AuthContextDto{
		WorkspaceId:  wi,
		Token:        tk,
		Capabilities: securityModel.ActionRequires,
		Security:     securityModel,
	}

	result, err := WithAuthorizationPureDefault(context)

	if err != nil {
		if isWebsocketUpgrade {
			// Already hijacked - can't send an HTTP response. The caller
			// (the reactive action's generated Gin handler) reports the
			// failure as a websocket frame instead, using the real
			// connection rather than gin's now-unusable writer.
			return nil, nil
		}
		c.AbortWithStatusJSON(int(err.HttpCode), gin.H{"error": err.ToPublicEndUser(&q)})
		return nil, nil
	}

	return result, nil
}
