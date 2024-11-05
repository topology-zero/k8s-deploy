package websocket

import (
	"k8s-deploy/logic/websocket"
	"k8s-deploy/svc"

	"github.com/gin-gonic/gin"
)

// WebsocketHandle websocket
func WebsocketHandle(c *gin.Context) {
	websocket.Websocket(svc.NewServiceContext(c))
}
