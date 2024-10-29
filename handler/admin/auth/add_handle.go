package auth

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/admin/auth"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// AddHandle 添加权限
func AddHandle(c *gin.Context) {
	var req types.AuthAddRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	if req.IsMenu == 0 && (len(req.API) == 0 || len(req.Action) == 0) {
		response.HandleResponse(c, nil, errors.New("非菜单权限必填操作方法和接口"))
		return
	}

	err := auth.Add(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
