package auth

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/admin/auth"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// DelHandle 删除权限
func DelHandle(c *gin.Context) {
	var req types.PathID
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	err := auth.Del(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
