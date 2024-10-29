package user

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/admin/user"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// AddHandle 添加用户
func AddHandle(c *gin.Context) {
	var req types.UserAddRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	err := user.Add(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
