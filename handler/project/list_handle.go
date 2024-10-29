package project

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/project"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// ListHandle 项目列表
func ListHandle(c *gin.Context) {
	var req types.ProjectListRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := project.List(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}
