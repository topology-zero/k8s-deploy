package namespace

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/namespace"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// DelHandle 删除命名空间
func DelHandle(c *gin.Context) {
	var req types.NamespaceDelRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	err := namespace.Del(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
