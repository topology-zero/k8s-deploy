package role

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/admin/role"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// AddHandle 添加角色
func AddHandle(c *gin.Context) {
	var req types.RoleAddRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	err := role.Add(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
