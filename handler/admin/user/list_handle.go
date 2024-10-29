package user

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/admin/user"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// ListHandle 用户列表
func ListHandle(c *gin.Context) {
	var req types.UserListRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := user.List(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}
