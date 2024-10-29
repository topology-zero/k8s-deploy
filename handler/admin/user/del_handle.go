package user

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/admin/user"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// DelHandle 删除用户
func DelHandle(c *gin.Context) {
	var req types.PathID
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	err := user.Del(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
