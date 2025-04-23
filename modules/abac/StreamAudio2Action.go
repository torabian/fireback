package abac

import (
	"time"

	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	StreamAudio2ActionImp = func(query fireback.QueryDSL, done chan bool, read chan []byte) (chan []byte, error) {
		return fireback.SocketStreamToFile(done, read, "file_"+time.Now().Format("20060102_150405")+".webm")
	}
}
