package user

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/admin/user"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// DetailHandle 用户详情
func DetailHandle(c *gin.Context) {
	var req types.PathID
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := user.Detail(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}
