package namespace

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/namespace"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// AddHandle 添加命名空间
func AddHandle(c *gin.Context) {
	var req types.NamespaceAddRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	err := namespace.Add(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
