package demo

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/torabian/fireback/modules/workspaces"
)

func sendStringWithInterval(ids []string, ctx context.Context, interval time.Duration, out chan *string) {

	for {
		select {
		case <-ctx.Done():
			return
		default:
			activities := []*UserActivityActivities{}
			for _, v := range ids {
				v := v
				activityState := int64(rand.Intn(3) + 1)
				activities = append(activities, &UserActivityActivities{

					UniqueId: &v,
					Activity: &activityState,
				})

			}

			data := UserActivityDto{
				Activities: activities,
			}

			js := data.Json()
			out <- &js
			time.Sleep(time.Millisecond * 1000)
		}
	}

}

func TestAction(
	query workspaces.QueryDSL,
	done chan bool,
	read chan map[string]interface{},
) (chan *string, error) {

	controlSheetStreamParsed := make(chan *string)

	go func() {
		var ctx context.Context = nil
		var cancel context.CancelFunc = nil

		for {
			select {
			case <-done:
				fmt.Println("Completed actually")
				return

			case row, more := <-read:

				data, _ := json.MarshalIndent(row, "", "  ")
				var dto UserActivityFocusDto
				json.Unmarshal(data, &dto)
				fmt.Println(dto)

				if cancel != nil {
					cancel()
				}

				ctx, cancel = context.WithCancel(context.Background())
				defer cancel()

				go sendStringWithInterval(dto.Ids, ctx, 1000, controlSheetStreamParsed)

				if !more {
					return
				}
			}
		}
	}()

	return controlSheetStreamParsed, nil
}
