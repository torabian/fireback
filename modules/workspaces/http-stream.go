package workspaces

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ContentType = "text/event-stream"

var contentType = []string{}
var noCache = []string{"no-cache"}

type stringWriter interface {
	io.Writer
	WriteString(string) (int, error)
}

type stringWrapper struct {
	io.Writer
}

func (w stringWrapper) WriteString(str string) (int, error) {
	return w.Writer.Write([]byte(str))
}

func checkWriter(writer io.Writer) stringWriter {
	if w, ok := writer.(stringWriter); ok {
		return w
	} else {
		return stringWrapper{writer}
	}
}

type Event struct {
	Event string
	Id    string
	Retry uint
	Data  []byte
}

func Encode(writer io.Writer, event Event) error {
	w := checkWriter(writer)
	w.WriteString(string(event.Data))

	return nil
}

func (r Event) Render(w http.ResponseWriter) error {
	return Encode(w, r)
}

func (r Event) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	header["Content-Type"] = contentType

	if _, exist := header["Cache-Control"]; !exist {
		header["Cache-Control"] = noCache
	}
}

func Stream(c *gin.Context, step func(w io.Writer) bool) bool {

	w := c.Writer
	clientGone := w.CloseNotify()
	for {
		select {
		case <-clientGone:
			return true
		default:
			keepOpen := step(w)
			w.Flush()
			if !keepOpen {
				return false
			}
		}
	}
}

func WriteToStream(c *gin.Context, data []byte) {
	c.Render(-1, Event{
		Data: data,
	})
}
