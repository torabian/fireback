package fireback

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	activeConnections int
	maxConnections    int
	mutex             sync.Mutex
)

// Middleware to track active connections
func trackConnectionsMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Increment the active connection count

		mutex.Lock()
		activeConnections++

		if activeConnections > maxConnections {
			maxConnections = activeConnections
		}

		mutex.Unlock()

		// Decrement after the request is processed
		defer func() {

			mutex.Lock()
			activeConnections--
			mutex.Unlock()
		}()

		// Continue with the request
		c.Next()
	}
}
