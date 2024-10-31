package deploy

import (
	"k8s-deploy/internal/response"
	"k8s-deploy/logic/deploy"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/gin-gonic/gin"
)

// AddHandle 添加部署
func AddHandle(c *gin.Context) {
	var req types.DeployAddRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	err := deploy.Add(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
