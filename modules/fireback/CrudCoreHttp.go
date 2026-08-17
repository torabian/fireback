package fireback

import (
	"net/http"
	"reflect"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	//Solve "request origin not allowed by Upgrader.CheckOrigin"
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type QueryableAction[T any] func(query QueryDSL) ([]*T, *QueryResultMeta, error)

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GetType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}
