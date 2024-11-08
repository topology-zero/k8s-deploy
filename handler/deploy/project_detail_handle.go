package deploy

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/deploy"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// ProjectDetailHandle 项目详情
func ProjectDetailHandle(c *gin.Context) {
	var req types.PathID
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := deploy.ProjectDetail(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}
