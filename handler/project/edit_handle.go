package project

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/project"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// EditHandle 编辑项目
func EditHandle(c *gin.Context) {
	var req types.ProjectEditRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	err := project.Edit(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
