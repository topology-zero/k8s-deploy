package run

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/run"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// DeployHandle 运行
func DeployHandle(c *gin.Context) {
	var req types.RunDeployRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	err := run.Deploy(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
