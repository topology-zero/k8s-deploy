package socket

import (
	"sync"

	"github.com/gorilla/websocket"
)

var AllConn = make(map[string]Conn)

type Conn struct {
	WebsocketConn *websocket.Conn
	Lock          *sync.Mutex
}

func SendJsonMessage(data any) {
	for _, conn := range AllConn {
		conn.Lock.Lock()
		conn.WebsocketConn.WriteJSON(data)
		conn.Lock.Unlock()
	}
}
func SendMessage(data string) {
	for _, conn := range AllConn {
		conn.Lock.Lock()
		conn.WebsocketConn.WriteMessage(websocket.TextMessage, []byte(data))
		conn.Lock.Unlock()
	}
}
