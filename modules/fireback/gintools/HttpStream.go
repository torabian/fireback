package gintools

import (
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

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

type HttpEvent struct {
	Event string
	Id    string
	Retry uint
	Data  []byte
}

func Encode(writer io.Writer, event HttpEvent) error {
	w := checkWriter(writer)
	w.WriteString(string(event.Data))

	return nil
}

func (r HttpEvent) Render(w http.ResponseWriter) error {
	return Encode(w, r)
}

func (r HttpEvent) WriteContentType(w http.ResponseWriter) {
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
	c.Render(-1, HttpEvent{
		Data: data,
	})
}

func GinStreamFromChannel(c *gin.Context, chanStream chan []byte) {
	rc := http.NewResponseController(c.Writer)
	rc.SetWriteDeadline(time.Time{})

	c.Header("Content-Type", "application/x-yaml")
	c.Header("Connection", "Keep-Alive")
	c.Header("Transfer-Encoding", "chunked")
	c.Header("Content-Disposition", `inline; filename="myfile.txt"`)
	c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

	Stream(c, func(w io.Writer) bool {
		if msg, ok := <-chanStream; ok {
			WriteToStream(c, msg)
			return true
		}
		return false
	})
}
