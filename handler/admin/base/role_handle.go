package base

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/admin/base"
	"k8s-deploy/svc"

	"github.com/gin-gonic/gin"
)

// RoleHandle 获取所有角色
func RoleHandle(c *gin.Context) {
	resp, err := base.Role(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
