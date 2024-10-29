package login

import (
	"k8s-deploy/config"
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/admin/login"
	"k8s-deploy/pkg/jwt"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// LoginHandle 登录
func LoginHandle(c *gin.Context) {
	var req types.LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := login.Login(svc.NewServiceContext(c), &req)
	if err == nil {
		c.SetCookie(jwt.JwtName, resp.Jwt, config.JwtConf.Expire*3600, "/", "127.0.0.1", true, true)
		c.SetCookie(jwt.JwtName, resp.Jwt, config.JwtConf.Expire*3600, "/", "localhost", true, true)
	}
	response.HandleResponse(c, resp, err)
}
