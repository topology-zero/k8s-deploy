package auth

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/admin/auth"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// EditHandle 编辑权限
func EditHandle(c *gin.Context) {
	var req types.AuthEditRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	err := auth.Edit(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
