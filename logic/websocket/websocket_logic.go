package websocket

import (
	"net/http"
	"sync"

	"k8s-deploy/logic/websocket/data_chain"
	"k8s-deploy/pkg/socket"
	"k8s-deploy/svc"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Websocket websocket
func Websocket(ctx *svc.ServiceContext) {
	conn, err := upgrader.Upgrade(ctx.GinContext.Writer, ctx.GinContext.Request, nil)
	if err != nil {
		ctx.Log.Errorf("%+v", err)
		return
	}

	defer conn.Close()
	uid := uuid.New().String()
	myConn := socket.Conn{WebsocketConn: conn, Lock: &sync.Mutex{}}
	socket.AllConn[uid] = myConn
	defer delete(socket.AllConn, uid)

	for {
		mt, data, err := conn.ReadMessage()
		if err != nil {
			delete(socket.AllConn, uid)
			if _, ok := err.(*websocket.CloseError); ok {
				ctx.Log.Errorf("%s 正常关闭", uid)
				return
			}
			ctx.Log.Errorf("非正常关闭\n%+v", errors.WithStack(err))
			return
		}
		dataStr := string(data)
		if dataStr == "ping" {
			myConn.Lock.Lock()
			err := myConn.WebsocketConn.WriteMessage(mt, []byte("pong"))
			myConn.Lock.Unlock()
			if err != nil {
				ctx.Log.Errorf("%+v", errors.WithStack(err))
				return
			}
		}

		context := data_chain.ChainContext{Ctx: ctx}
		data_chain.ParseData(&context, data)
	}
}
