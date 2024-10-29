package base

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/admin/base"
	"k8s-deploy/svc"

	"github.com/gin-gonic/gin"
)

// AuthHandle 获取所有权限
func AuthHandle(c *gin.Context) {
	resp, err := base.Auth(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
