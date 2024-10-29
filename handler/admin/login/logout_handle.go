package login

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/admin/login"
	"k8s-deploy/svc"

	"github.com/gin-gonic/gin"
)

// LogoutHandle 退出登录
func LogoutHandle(c *gin.Context) {
	err := login.Logout(svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}
