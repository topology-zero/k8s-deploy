package base

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/admin/base"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// ChangeSelfPwdHandle 修改自己的密码
func ChangeSelfPwdHandle(c *gin.Context) {
	var req types.ChangeSelfPwdRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	err := base.ChangeSelfPwd(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
