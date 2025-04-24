package abac

import (
	"fmt"

	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	// StreamAudio2ActionImp = func(query fireback.QueryDSL, done chan bool, read chan fireback.SocketReadChan) (chan []byte, error) {
	// 	return fireback.SocketStreamToFile(done, read, "file_"+time.Now().Format("20060102_150405")+".webm")
	// }
	StreamAudio2ActionImp = func(query fireback.QueryDSL, done chan bool, read chan fireback.SocketReadChan) (chan []byte, error) {
		out := make(chan []byte)

		go func() {
			defer close(out)
			for {
				select {
				case msg, ok := <-read:
					fmt.Println(msg, ok)
					if !ok {
						return
					}

				case <-done:
					return
				}
			}
		}()

		return out, nil
	}
}
