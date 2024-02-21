package workspaces

import "time"

var startTime time.Time = time.Now()

func Uptime() time.Duration {
	return time.Since(startTime)
}
