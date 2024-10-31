package deploy

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/deploy"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// ListHandle 部署列表
func ListHandle(c *gin.Context) {
	var req types.DeployListRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := deploy.List(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}
