package login

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/admin/login"
	"k8s-deploy/svc"

	"github.com/gin-gonic/gin"
)

// CodeHandle 获取验证码
func CodeHandle(c *gin.Context) {
	resp, err := login.Code(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
