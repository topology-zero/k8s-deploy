package project

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/project"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// AddHandle 添加项目
func AddHandle(c *gin.Context) {
	var req types.ProjectAddRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	err := project.Add(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
