package role

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/admin/role"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// ListHandle 角色列表
func ListHandle(c *gin.Context) {
	var req types.RoleListRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := role.List(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}
