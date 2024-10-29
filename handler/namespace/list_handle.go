package namespace

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/namespace"
	"k8s-deploy/svc"

	"github.com/gin-gonic/gin"
)

// ListHandle 命名空间列表
func ListHandle(c *gin.Context) {
	resp, err := namespace.List(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
