package auth

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/admin/auth"
	"k8s-deploy/svc"

	"github.com/gin-gonic/gin"
)

// ListHandle 权限列表
func ListHandle(c *gin.Context) {
	resp, err := auth.List(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
