package user

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/admin/user"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// EditHandle 编辑用户
func EditHandle(c *gin.Context) {
	var req types.UserEditRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	pwdLen := len(req.Password)
	if pwdLen > 0 && pwdLen < 6 {
		response.HandleResponse(c, nil, errors.New("密码不得小于6位数"))
		return
	}

	err := user.Edit(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
