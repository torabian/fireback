package workspaces

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	ui "github.com/gizak/termui/v3"

	"github.com/gizak/termui/v3/widgets"
)

func monitor() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	sl := widgets.NewSparkline()
	sl.Title = "Active connections"
	sl.Data = []float64{}
	sl.LineColor = ui.ColorCyan
	sl.TitleStyle.Fg = ui.ColorWhite

	data := []float64{}

	draw := func(count int) {

		if len(data) > 70 {
			data = data[1:]
		}
		mutex.Lock()
		data = append(data, float64(activeConnections)*10)
		mutex.Unlock()
		sl.Data = data
		sl.Title = fmt.Sprintf("Active: %v,   max: %v", activeConnections, maxConnections)

		slg := widgets.NewSparklineGroup(sl)
		slg.Title = "Http connections"
		slg.SetRect(0, 0, 100, 10)

		ui.Render(slg)
	}

	tickerCount := 1
	draw(tickerCount)
	tickerCount++
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(250 * time.Millisecond).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-ticker:

			draw(tickerCount)
			tickerCount++
		}
	}
}

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
