package base

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/admin/base"
	"k8s-deploy/svc"

	"github.com/gin-gonic/gin"
)

// UserInfoHandle 获取用户信息
func UserInfoHandle(c *gin.Context) {
	resp, err := base.UserInfo(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
